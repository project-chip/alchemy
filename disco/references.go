package disco

import (
	"log/slog"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/asciidoc/parse"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
)

func (an AnchorNormalizer) rewriteCrossReferences(doc *spec.Doc) {
	for id, xrefs := range doc.CrossReferences() {
		if len(xrefs) == 0 {
			continue
		}

		anchor := doc.FindAnchor(id, xrefs[0].Reference)
		if anchor == nil {

			anchors := doc.FindAnchors(id)
			allAnchorsSkippable := true
			for _, a := range anchors {
				if !allAnchorsSkippable || !skipAnchor(a) {
					allAnchorsSkippable = false
				}
			}
			if allAnchorsSkippable {
				continue
			}
			sources := []any{slog.String("id", id)}
			for _, xref := range xrefs {
				sources = append(sources, log.Path("source", xref.Source))
			}
			switch len(anchors) {
			case 0:
				slog.Warn("cross reference points to non-existent anchor", sources...)
			default:
				for _, a := range anchors {
					sources = append(sources, log.Path("target", a.Source))
				}
				slog.Warn("cross reference points to ambiguous anchor", slog.Group("sources", sources...))
			}
			continue
		}
		anchorLabel := labelText(anchor.LabelElements)
		if anchorLabel == "" {
			section, isSection := anchor.Element.(*asciidoc.Section)
			if isSection {
				anchorLabel = doc.SectionName(section)
			}
		}
		// We're going to be modifying the underlying array, so we need to make a copy of the slice
		xrefsToChange := make([]*spec.CrossReference, len(xrefs))
		copy(xrefsToChange, xrefs)
		for _, xref := range xrefsToChange {
			existingAnchorID := anchor.Identifier(asciidoc.RawReader)
			existingXrefID := xref.Identifier(asciidoc.RawReader)
			if existingAnchorID != existingXrefID {
				xref.SyncToDoc(anchor.ID)
			}
			if len(xref.Reference.Elements) > 0 {
				// If the cross reference has a label that's the same as the one we generated for the anchor, remove it
				refText := labelText(xref.Reference.Elements)
				if anchorLabel == refText {
					xref.Reference.Elements = nil
				}
			}
		}
	}
	if an.options.NormalizeAnchors {
		parse.Search(asciidoc.RawReader, nil, doc.Base.Children(), func(el asciidoc.Element, parent asciidoc.Parent, index int) parse.SearchShould {
			switch el := el.(type) {
			case *asciidoc.CrossReference:
				normalizeCrossReference(doc, el)
				removeCrossReferenceStutter(doc, el, parent, index)
				return parse.SearchShouldContinue
			case *asciidoc.Table:
				an.normalizeTypeCrossReferencesInTable(doc, el)
				return parse.SearchShouldSkip
			default:
				return parse.SearchShouldContinue
			}
		})
	}
}

func removeCrossReferenceStutter(doc *spec.Doc, icr *asciidoc.CrossReference, parent asciidoc.Parent, index int) {
	if len(icr.Elements) > 0 {
		return
	}
	anchor := doc.FindAnchorByID(icr.ID, icr, icr)
	if anchor == nil {
		return
	}
	section, isSection := anchor.Element.(*asciidoc.Section)
	if !isSection {
		return
	}
	sectionName := section.Name()
	elements := parent.Children()
	if index >= len(elements)-1 {
		return
	}
	nextElement := elements[index+1]
	nextString, ok := nextElement.(*asciidoc.String)
	if !ok {
		return
	}
	lastSpaceIndex := strings.LastIndex(sectionName, " ")
	if lastSpaceIndex == -1 {
		return
	}
	suffix := sectionName[lastSpaceIndex:]
	if !text.HasCaseInsensitivePrefix(nextString.Value, suffix) {
		return
	}
	replacement := text.TrimCaseInsensitivePrefix(nextString.Value, suffix)
	nextString.Value = replacement
}

func (an AnchorNormalizer) normalizeTypeCrossReferencesInTable(doc *spec.Doc, table *asciidoc.Table) {
	parse.Search(asciidoc.RawReader, table, table.Elements, func(icr *asciidoc.CrossReference, parent asciidoc.Parent, index int) parse.SearchShould {
		normalizeCrossReference(doc, icr)
		return parse.SearchShouldContinue
	})

}

func normalizeCrossReference(doc *spec.Doc, icr *asciidoc.CrossReference) {
	if len(icr.Elements) > 0 {
		// Don't touch existing labels
		return
	}
	anchor := doc.FindAnchorByID(icr.ID, icr, icr)
	if anchor == nil {
		return
	}
	section, isSection := anchor.Element.(*asciidoc.Section)
	if !isSection {
		return
	}
	entities, ok := anchor.Document.EntitiesForSection(section)
	if !ok {
		return
	}
	if len(entities) != 1 {
		return
	}
	entity := entities[0]
	if !types.IsDataTypeEntity(entity.EntityType()) {
		return
	}
	normalizedLabel := normalizeAnchorLabel(section.Name(), section)
	if labelText(normalizedLabel) != section.Name() {
		icr.Elements = normalizedLabel
		slog.Debug("Added label to type xref in table", matter.LogEntity("type", entity), "label", labelText(icr.Elements))
	}

}

func findRefSection(parent any) *asciidoc.Section {
	switch p := parent.(type) {
	case *asciidoc.Section:
		return p
	case asciidoc.HasParent:
		return findRefSection(p.Parent())
	case *spec.Doc:
		return nil
	default:
		return nil
	}

}
