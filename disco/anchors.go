package disco

import (
	"context"
	"fmt"
	"log/slog"
	"regexp"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
)

var properAnchorPattern = regexp.MustCompile(`^ref_[A-Z0-9]+[a-z0-9]*(?:[A-Z]+[a-z]*)*([A-Z0-9]+[a-z0-9]*(?:[A-Z0-9]+[a-z0-9]*)*)*$`)
var improperAnchorCharactersPattern = regexp.MustCompile(`[^A-Za-z0-9]+`)

type anchorLibrary struct {
	library          *spec.Library
	updatedAnchors   map[string][]*spec.Anchor
	rewrittenAnchors map[string][]*spec.Anchor
}

type AnchorNormalizer struct {
	spec    *spec.Specification
	options DiscoOptions
}

func newAnchorNormalizer(spec *spec.Specification, discoOptions DiscoOptions) AnchorNormalizer {
	an := AnchorNormalizer{spec: spec, options: discoOptions}
	return an
}

func (r AnchorNormalizer) Name() string {
	return "Normalizing anchors"
}

func (p AnchorNormalizer) Process(cxt context.Context, inputs []*pipeline.Data[*asciidoc.Document]) (outputs []*pipeline.Data[*asciidoc.Document], err error) {
	var anchorLibraries map[*spec.Library]*anchorLibrary
	anchorLibraries, err = p.normalizeAnchors(inputs)
	if err != nil {
		return
	}
	extraDocs := make(map[*asciidoc.Document]struct{})
	for _, al := range anchorLibraries {
		for id, infos := range al.updatedAnchors {
			if len(infos) == 1 {
				al.library.SyncToDoc(infos[0], asciidoc.NewStringElements(id))
			} else if len(infos) > 1 { // We ended up with some duplicate anchors
				var skip bool
				for _, info := range infos {
					if skipAnchor(info) {
						skip = true
					}
				}
				if skip {
					continue
				}
				var disambiguatedIDs []string
				disambiguatedIDs, err = disambiguateAnchorSet(infos, id, al)
				if err != nil {
					var args []any
					args = append(args, slog.String("id", id), slog.Any("error", err))
					for _, info := range infos {
						args = append(args, log.Element("source", info.Document.Path, info.Element))
					}

					slog.Warn("failed disambiguating anchor", args...)
					err = nil
					continue
				}
				for i, info := range infos {
					al.library.SyncToDoc(info, asciidoc.NewStringElements(disambiguatedIDs[i]))
				}
			}
		}
		for from, to := range al.rewrittenAnchors {
			xrefs := al.library.CrossReferences(from)
			// We're going to be modifying the underlying array, so we need to make a copy of the slice
			xrefsToChange := make([]*spec.CrossReference, len(xrefs))
			copy(xrefsToChange, xrefs)
			if len(to) == 1 {
				a := to[0]
				for _, xref := range xrefsToChange {
					al.library.SyncCrossReference(xref, a.ID)
					extraDocs[xref.Document] = struct{}{}
				}
			} else {
				docs := make(map[*asciidoc.Document][]*spec.Anchor)
				for _, info := range to {
					docs[info.Document] = append(docs[info.Document], info)
				}
				for _, xref := range xrefsToChange {
					info, ok := docs[xref.Document]
					if ok && len(info) == 1 {
						al.library.SyncCrossReference(xref, info[0].ID)
					} else {
						var logArgs []any
						logArgs = append(logArgs, slog.Any("id", al.library.Identifier(xref.Reference, xref.Reference, xref.Reference.Elements)), log.Path("origin", xref.Source))
						for _, info := range to {
							logArgs = append(logArgs, slog.Any("target", al.library.Identifier(info.Parent, info.Element, info.ID)), log.Path("targetPath", info.Source))
						}
						slog.Warn("rewritten xref points to multiple anchors", logArgs...)
					}
				}
			}
		}
	}
	for _, input := range inputs {
		doc := input.Content
		p.rewriteCrossReferences(doc)
		delete(extraDocs, doc)
		outputs = append(outputs, pipeline.NewData[*asciidoc.Document](input.Path, input.Content))
	}
	for doc := range extraDocs {
		outputs = append(outputs, pipeline.NewData[*asciidoc.Document](doc.Path.Relative, doc))
	}
	return
}

func (an AnchorNormalizer) normalizeAnchors(inputs []*pipeline.Data[*asciidoc.Document]) (anchorGroups map[*spec.Library]*anchorLibrary, err error) {
	anchorGroups = make(map[*spec.Library]*anchorLibrary)
	docs := make(map[*asciidoc.Document]struct{})
	for _, input := range inputs {
		doc := input.Content
		library, ok := an.spec.LibraryForDocument(doc)
		if !ok {
			continue
		}
		_, ok = anchorGroups[library]
		if !ok {
			anchorLibrary := &anchorLibrary{
				library:          library,
				updatedAnchors:   make(map[string][]*spec.Anchor),
				rewrittenAnchors: make(map[string][]*spec.Anchor),
			}
			anchorGroups[library] = anchorLibrary
		}
		docs[doc] = struct{}{}
	}
	for library, anchorLibrary := range anchorGroups {
		var da map[string][]*spec.Anchor
		da, err = library.Anchors(asciidoc.RawReader)
		if err != nil {
			err = fmt.Errorf("error fetching anchors in %w", err)
			return
		}
		for _, as := range da {
			for _, a := range as {
				if _, ok := docs[a.Document]; !ok {
					continue
				}
				if isDynamicAnchor(a) {
					continue
				}
				var id string
				id, err = library.StringValue(a.Parent, a.ID)
				newID := an.normalizeAnchor(library, a)
				if id == newID {
					anchorLibrary.updatedAnchors[id] = append(anchorLibrary.updatedAnchors[id], a)
					continue
				}
				if _, existingAnchor := da[newID]; existingAnchor {
					slog.Warn("Attempting to rename anchor to existing anchor", slog.Any("oldAnchor", id), slog.String("newAnchor", newID), log.Element("source", a.Document.Path, a.Element))
					continue
				}
				anchorLibrary.updatedAnchors[newID] = append(anchorLibrary.updatedAnchors[newID], a)
				slog.Debug("rewrote anchor", "from", id, "to", newID)
				anchorLibrary.rewrittenAnchors[id] = append(anchorLibrary.rewrittenAnchors[id], a)
			}
		}
	}
	return
}

func (an AnchorNormalizer) normalizeAnchor(library *spec.Library, info *spec.Anchor) (id string) {
	id = library.Identifier(info.Parent, info.Element, info.ID)
	if skipAnchor(info) {
		return
	}
	name := info.Name(asciidoc.RawReader)
	if properAnchorPattern.MatchString(id) {
		if len(info.LabelElements) == 0 || labelText(info.LabelElements) == name {
			info.LabelElements = normalizeAnchorLabel(info.Name(asciidoc.RawReader), info.Element)
		}
	} else {
		normalizedID, normalized := quickNormalizeAnchorID(id)
		if normalized {
			id = normalizedID
			if len(info.LabelElements) == 0 || labelText(info.LabelElements) == name {
				info.LabelElements = normalizeAnchorLabel(info.Name(asciidoc.RawReader), info.Element)
			}
		} else {
			if len(name) == 0 {
				name = labelText(info.LabelElements)
			}
			if len(info.LabelElements) == 0 || labelText(info.LabelElements) == name {
				info.LabelElements = normalizeAnchorLabel(name, info.Element)
			}
			id = normalizeAnchorID(id, info.LabelElements)
		}
	}
	if labelText(info.LabelElements) == name {
		info.LabelElements = nil
	}
	if an.options.NormalizeAnchors {
		_, isSection := info.Element.(*asciidoc.Section)
		if isSection {
			info.LabelElements = nil
		}
	}
	return
}

func skipAnchor(info *spec.Anchor) bool {
	if isDynamicAnchor(info) {
		return true
	}
	section, ok := info.Element.(*asciidoc.Section)
	if !ok {
		return false
	}
	if info.Library.DiscoErrata(info.Document.Path.Relative).IgnoreSection(info.Library.SectionName(section), errata.DiscoPurposeNormalizeAnchor) {
		return true
	}
	return false
}

func isDynamicAnchor(info *spec.Anchor) bool {
	for _, idElement := range info.ID {
		switch idElement.(type) {
		case *asciidoc.String:
		default:
			return true
		}
	}
	return false
}

var anchorInvalidCharacters = strings.NewReplacer(".", "", "(", "", ")", "")

func quickNormalizeAnchorID(id string) (normalizedID string, normalized bool) {
	if !strings.HasPrefix(id, "ref_") {
		return
	}
	// If it starts with ref_, maybe we can just replace non-alphanumeric characters and get a valid ID
	normalizedID = strings.TrimPrefix(id, "ref_")
	normalizedID = "ref_" + improperAnchorCharactersPattern.ReplaceAllString(normalizedID, "")
	normalized = properAnchorPattern.MatchString(normalizedID)
	return
}

func normalizeAnchorID(existingID string, label asciidoc.Elements) (id string) {

	var ref strings.Builder
	ref.WriteString("ref_")
	existingID = strings.TrimPrefix(existingID, "ref_")
	existingID = matter.Case(anchorInvalidCharacters.Replace(existingID))
	if len(existingID) > 0 {
		ref.WriteString(existingID)
	} else {
		labelString := asciidoc.AttributeAsciiDocString(label)
		ref.WriteString(matter.Case(anchorInvalidCharacters.Replace(labelString)))
	}
	id = ref.String()
	return
}

func normalizeAnchorLabel(name string, element any) (label asciidoc.Elements) {
	switch element.(type) {
	case *asciidoc.Table:
		label = asciidoc.Elements{asciidoc.NewString(strings.TrimSpace(name))}
	default:
		name = text.TrimCaseInsensitiveSuffix(name, " Type")
		label = asciidoc.Elements{asciidoc.NewString(strings.TrimSpace(matter.StripReferenceSuffixes(name)))}
	}
	return
}

func disambiguateAnchorSet(conflictedAnchors []*spec.Anchor, newID string, ag *anchorLibrary) (newIDs []string, err error) {
	parents := make([]any, len(conflictedAnchors))
	newIDs = make([]string, len(conflictedAnchors))
	for i, info := range conflictedAnchors {
		parents[i] = info.Parent
		newIDs[i] = newID
	}
	parentSections := make([]*asciidoc.Section, len(conflictedAnchors))
	for {
		for i := range conflictedAnchors {
			parentSection := findRefSection(parents[i])
			if parentSection == nil {
				var errMsg strings.Builder
				for _, info := range conflictedAnchors {
					origin, line := info.Source.Origin()
					errMsg.WriteString(fmt.Sprintf(", %s:%d", origin, line))
				}
				err = fmt.Errorf("duplicate anchor: %s with invalid parent%s", newIDs[i], errMsg.String())
				return

			}
			parentSections[i] = parentSection
			parentName := spec.ReferenceName(asciidoc.RawReader, parentSection)
			parentName = matter.Case(matter.StripTypeSuffixes(parentName))
			newIDs[i] = "ref_" + parentName + strings.TrimPrefix(newIDs[i], "ref_")
		}
		ids := make(map[string]struct{})
		var duplicateIds bool
		for _, refID := range newIDs {
			if _, ok := ids[refID]; ok {
				duplicateIds = true
			}
			ids[refID] = struct{}{}
		}
		if duplicateIds {
			for i := range conflictedAnchors {
				parents[i] = parentSections[i].Parent()
			}
		} else {
			break
		}
	}
	for i, info := range conflictedAnchors {
		if _, ok := ag.rewrittenAnchors[newIDs[i]]; ok {
			slog.Warn("duplicate anchor target", "id", newID, "target", newIDs[i])
		}
		ag.rewrittenAnchors[newID] = append(ag.rewrittenAnchors[newID], info)
	}
	return
}

func labelText(label asciidoc.Elements) string {
	return strings.TrimSpace(asciidoc.AttributeAsciiDocString(label))
}
