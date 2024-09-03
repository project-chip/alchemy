package disco

import (
	"context"
	"fmt"
	"log/slog"
	"regexp"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/asciidoc/render"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
)

var properAnchorPattern = regexp.MustCompile(`^ref_[A-Z0-9]+[a-z0-9]*(?:[A-Z]+[a-z]*)*([A-Z0-9]+[a-z0-9]*(?:[A-Z0-9]+[a-z0-9]*)*)*$`)
var improperAnchorCharactersPattern = regexp.MustCompile(`[^A-Za-z0-9]+`)

type anchorGroup struct {
	group            *spec.DocGroup
	updatedAnchors   map[string][]*spec.Anchor
	rewrittenAnchors map[string][]*spec.Anchor
}

type AnchorNormalizer struct {
	discoOptions []Option
}

func newAnchorNormalizer(discoOptions []Option, pipelineOptions pipeline.Options) AnchorNormalizer {
	return AnchorNormalizer{discoOptions: discoOptions}
}

func (r AnchorNormalizer) Name() string {
	return "Normalizing anchors"
}

func (p AnchorNormalizer) Type() pipeline.ProcessorType {
	return pipeline.ProcessorTypeCollective
}

func (p AnchorNormalizer) Process(cxt context.Context, inputs []*pipeline.Data[*spec.Doc]) (outputs []*pipeline.Data[render.InputDocument], err error) {
	var anchorGroups map[*spec.DocGroup]*anchorGroup
	anchorGroups, err = p.normalizeAnchors(inputs)
	if err != nil {
		return
	}
	extraDocs := make(map[*spec.Doc]struct{})
	for _, ag := range anchorGroups {
		for id, infos := range ag.updatedAnchors {
			if len(infos) == 1 {
				infos[0].SyncToDoc(id)
			} else if len(infos) > 1 { // We ended up with some duplicate anchors
				var disambiguatedIDs []string
				disambiguatedIDs, err = disambiguateAnchorSet(infos, id, ag)
				if err != nil {
					var args []any
					args = append(args, slog.String("id", id), slog.Any("error", err))
					for _, info := range infos {
						args = append(args, log.Element("path", info.Document.Path, info.Element))
					}

					slog.Warn("failed disambiguating anchor", args...)
					continue
				}
				for i, info := range infos {
					info.SyncToDoc(disambiguatedIDs[i])
				}
			}
		}
		for from, to := range ag.rewrittenAnchors {
			xrefs := ag.group.CrossReferences(from)
			// We're going to be modifying the underlying array, so we need to make a copy of the slice
			xrefsToChange := make([]*spec.CrossReference, len(xrefs))
			copy(xrefsToChange, xrefs)
			if len(to) == 1 {
				a := to[0]
				for _, xref := range xrefsToChange {
					xref.SyncToDoc(a.ID)
					extraDocs[xref.Document] = struct{}{}
				}
			} else {
				docs := make(map[*spec.Doc][]*spec.Anchor)
				for _, info := range to {
					docs[info.Document] = append(docs[info.Document], info)
				}
				for _, xref := range xrefsToChange {
					info, ok := docs[xref.Document]
					if ok && len(info) == 1 {
						xref.SyncToDoc(info[0].ID)
					} else {
						var logArgs []any
						logArgs = append(logArgs, slog.String("id", xref.Reference.ID), log.Path("origin", xref.Source))
						for _, info := range to {
							logArgs = append(logArgs, slog.String("target", info.ID), log.Path("targetPath", info.Source))
						}
						slog.Warn("rewritten xref points to multiple anchors", logArgs...)
					}
				}
			}
		}
	}
	for _, input := range inputs {
		doc := input.Content
		rewriteCrossReferences(doc)
		delete(extraDocs, doc)
		outputs = append(outputs, pipeline.NewData[render.InputDocument](input.Path, input.Content))
	}
	for doc := range extraDocs {
		outputs = append(outputs, pipeline.NewData[render.InputDocument](doc.Path.Relative, doc))
	}
	return
}

func (AnchorNormalizer) normalizeAnchors(inputs []*pipeline.Data[*spec.Doc]) (anchorGroups map[*spec.DocGroup]*anchorGroup, err error) {
	anchorGroups = make(map[*spec.DocGroup]*anchorGroup)
	unaffiliatedDocs := spec.NewDocGroup("")
	for _, input := range inputs {
		doc := input.Content

		group := doc.Group()
		if group == nil {
			group = unaffiliatedDocs
		}
		ag, ok := anchorGroups[group]
		if !ok {
			ag = &anchorGroup{
				group:            group,
				updatedAnchors:   make(map[string][]*spec.Anchor),
				rewrittenAnchors: make(map[string][]*spec.Anchor),
			}
			anchorGroups[group] = ag
		}

		var da map[string][]*spec.Anchor
		da, err = doc.Anchors()
		if err != nil {
			err = fmt.Errorf("error fetching anchors in %s: %w", doc.Path, err)
			return
		}

		for _, as := range da {
			for _, a := range as {
				id := a.ID
				newID := normalizeAnchor(a)
				ag.updatedAnchors[newID] = append(ag.updatedAnchors[newID], a)
				if id != newID {
					slog.Debug("rewrote anchor", "from", id, "to", newID)
					ag.rewrittenAnchors[id] = append(ag.rewrittenAnchors[id], a)
				}
			}
		}

	}
	return
}

func normalizeAnchor(info *spec.Anchor) (id string) {
	id = info.ID
	name := info.Name()
	if properAnchorPattern.MatchString(info.ID) {
		if len(info.LabelElements) == 0 || labelText(info.LabelElements) == name {
			info.LabelElements = normalizeAnchorLabel(info.Name(), info.Element)
		}
	} else {
		normalizedID, normalized := quickNormalizeAnchorID(info.ID)
		if normalized {
			id = normalizedID
			if len(info.LabelElements) == 0 || labelText(info.LabelElements) == name {
				info.LabelElements = normalizeAnchorLabel(info.Name(), info.Element)
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
	return
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

func normalizeAnchorID(existingID string, label asciidoc.Set) (id string) {

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

func normalizeAnchorLabel(name string, element any) (label asciidoc.Set) {
	switch element.(type) {
	case *asciidoc.Table:
		label = asciidoc.Set{asciidoc.NewString(strings.TrimSpace(name))}
	default:
		name = text.TrimCaseInsensitiveSuffix(name, " Type")
		label = asciidoc.Set{asciidoc.NewString(strings.TrimSpace(matter.StripReferenceSuffixes(name)))}
	}
	return
}

func disambiguateAnchorSet(conflictedAnchors []*spec.Anchor, newID string, ag *anchorGroup) (newIDs []string, err error) {
	parents := make([]any, len(conflictedAnchors))
	newIDs = make([]string, len(conflictedAnchors))
	for i, info := range conflictedAnchors {
		parents[i] = info.Parent
		newIDs[i] = newID
	}
	parentSections := make([]*spec.Section, len(conflictedAnchors))
	for {
		for i := range conflictedAnchors {
			parentSection := findRefSection(parents[i])
			if parentSection == nil {
				var errMsg strings.Builder
				errMsg.WriteString(fmt.Sprintf("duplicate anchor: %s with invalid parent,", newIDs[i]))
				for _, info := range conflictedAnchors {
					origin, line := info.Source.Origin()
					errMsg.WriteString(fmt.Sprintf(", %s:%d", origin, line))
				}
				err = fmt.Errorf(errMsg.String())
				return

			}
			parentSections[i] = parentSection
			parentName := spec.ReferenceName(parentSection.Base)
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
				parents[i] = parentSections[i].Parent
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

func labelText(label asciidoc.Set) string {
	return strings.TrimSpace(asciidoc.AttributeAsciiDocString(label))
}
