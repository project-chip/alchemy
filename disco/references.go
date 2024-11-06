package disco

import (
	"log/slog"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/internal/parse"
	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter/spec"
)

func (p AnchorNormalizer) rewriteCrossReferences(doc *spec.Doc) {
	for id, xrefs := range doc.CrossReferences() {
		anchor := doc.FindAnchor(id)
		if anchor == nil {
			sources := []any{slog.String("id", id)}
			for _, xref := range xrefs {
				sources = append(sources, log.Path("source", xref.Source))
			}
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
			switch len(anchors) {
			case 0:
				slog.Warn("cross reference points to non-existent anchor", sources...)
			default:
				for _, a := range anchors {
					sources = append(sources, log.Path("target", a.Source))
				}
				slog.Warn("cross reference points to ambiguous anchor", sources...)
			}
			continue
		}
		anchorLabel := labelText(anchor.LabelElements)
		if anchorLabel == "" {
			section, isSection := anchor.Element.(*asciidoc.Section)
			if isSection {
				anchorLabel = section.Name()
			}
		}
		// We're going to be modifying the underlying array, so we need to make a copy of the slice
		xrefsToChange := make([]*spec.CrossReference, len(xrefs))
		copy(xrefsToChange, xrefs)
		for _, xref := range xrefsToChange {
			if anchor.ID != xref.Reference.ID {
				xref.SyncToDoc(anchor.ID)
			}
			if len(xref.Reference.Set) > 0 {
				// If the cross reference has a label that's the same as the one we generated for the anchor, remove it
				refText := labelText(xref.Reference.Set)
				if anchorLabel == refText {
					xref.Reference.Set = nil
				}
			}
		}
	}
	if p.options.normalizeAnchors {
		parse.Traverse(nil, doc.Base.Elements(), func(icr *asciidoc.CrossReference, parent parse.HasElements, index int) parse.SearchShould {
			if len(icr.Set) > 0 {
				return parse.SearchShouldContinue
			}
			anchor := doc.FindAnchor(icr.ID)
			if anchor == nil {
				return parse.SearchShouldContinue
			}
			section, isSection := anchor.Element.(*asciidoc.Section)
			if !isSection {
				return parse.SearchShouldContinue
			}
			sectionName := section.Name()
			elements := parent.Elements()
			if index >= len(elements)-1 {
				return parse.SearchShouldContinue
			}
			nextElement := elements[index+1]
			nextString, ok := nextElement.(*asciidoc.String)
			if !ok {
				return parse.SearchShouldContinue
			}
			lastSpaceIndex := strings.LastIndex(sectionName, " ")
			if lastSpaceIndex == -1 {
				return parse.SearchShouldContinue
			}
			suffix := sectionName[lastSpaceIndex:]
			if !text.HasCaseInsensitivePrefix(nextString.Value, suffix) {
				return parse.SearchShouldContinue
			}
			replacement := text.TrimCaseInsensitivePrefix(nextString.Value, suffix)
			nextString.Value = replacement
			return parse.SearchShouldContinue
		})

	}
}

func findRefSection(parent any) *spec.Section {
	switch p := parent.(type) {
	case *spec.Section:
		return p
	case *spec.Element:
		return findRefSection(p.Parent)
	case *spec.Doc:
		return nil
	default:
		return nil
	}

}
