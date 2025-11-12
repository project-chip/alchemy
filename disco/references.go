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

func (an AnchorNormalizer) rewriteCrossReferences(doc *asciidoc.Document) {

	library, ok := an.spec.LibraryForDocument(doc)
	if !ok {
		return
	}

	for id, xrefs := range library.CrossReferencesForDoc(doc) {
		if len(xrefs) == 0 {
			continue
		}

		anchor := library.FindAnchor(id, xrefs[0].Reference)
		if anchor == nil {

			anchors := library.FindAnchors(id)
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
				anchorLabel = library.SectionName(section)
			}
		}
		// We're going to be modifying the underlying array, so we need to make a copy of the slice
		xrefsToChange := make([]*spec.CrossReference, len(xrefs))
		copy(xrefsToChange, xrefs)
		for _, xref := range xrefsToChange {
			existingAnchorID := library.Identifier(anchor.Parent, anchor.Element, anchor.ID)
			existingXrefID := library.Identifier(xref.Reference, xref.Reference, xref.Reference.Elements)
			if existingAnchorID != existingXrefID {
				library.SyncToDoc(anchor, anchor.ID)
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
		parse.Search(doc, asciidoc.RawReader, nil, doc.Children(), func(doc *asciidoc.Document, el asciidoc.Element, parent asciidoc.ParentElement, index int) parse.SearchShould {
			switch el := el.(type) {
			case *asciidoc.CrossReference:
				normalizeCrossReference(library, doc, el)
				removeCrossReferenceStutter(library, doc, el, parent, index)
				return parse.SearchShouldContinue
			case *asciidoc.Table:
				an.normalizeTypeCrossReferencesInTable(library, doc, el)
				return parse.SearchShouldSkip
			default:
				return parse.SearchShouldContinue
			}
		})
	}
}

func removeCrossReferenceStutter(library *spec.Library, doc *asciidoc.Document, icr *asciidoc.CrossReference, parent asciidoc.Parent, index int) {
	if len(icr.Elements) > 0 {
		return
	}
	anchor := library.FindAnchorByID(icr.ID, icr, icr)
	if anchor == nil {
		return
	}
	section, isSection := anchor.Element.(*asciidoc.Section)
	if !isSection {
		return
	}
	sectionName := library.SectionName(section)
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

func (an AnchorNormalizer) normalizeTypeCrossReferencesInTable(library *spec.Library, doc *asciidoc.Document, table *asciidoc.Table) {
	parse.Search(doc, asciidoc.RawReader, table, table.Elements, func(doc *asciidoc.Document, icr *asciidoc.CrossReference, parent asciidoc.ParentElement, index int) parse.SearchShould {
		normalizeCrossReference(library, doc, icr)
		return parse.SearchShouldContinue
	})

}

func normalizeCrossReference(library *spec.Library, doc *asciidoc.Document, icr *asciidoc.CrossReference) {
	if len(icr.Elements) > 0 {
		// Don't touch existing labels
		return
	}
	anchor := library.FindAnchorByID(icr.ID, icr, icr)
	if anchor == nil {
		return
	}
	section, isSection := anchor.Element.(*asciidoc.Section)
	if !isSection {
		return
	}
	entities, ok := library.EntitiesForElement(section)
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
	sectionName := library.SectionName(section)
	normalizedLabel := normalizeAnchorLabel(sectionName, section)
	if labelText(normalizedLabel) != sectionName {
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
	case *asciidoc.Document:
		return nil
	default:
		return nil
	}

}
