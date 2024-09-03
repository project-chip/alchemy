package disco

import (
	"log/slog"

	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter/spec"
)

func rewriteCrossReferences(doc *spec.Doc) {
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
