package disco

import (
	"log/slog"

	"github.com/hasty/adoc/elements"
	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/matter"
)

func (b *Ball) rewriteCrossReferences(crossReferences map[string][]*elements.CrossReference, anchors map[string]*ascii.Anchor) {
	for id, xrefs := range crossReferences {
		info, ok := anchors[id]
		if !ok {
			dt, _ := b.doc.DocType()
			switch dt {
			case matter.DocTypeCluster, matter.DocTypeDeviceType:
				slog.Debug("cross reference points to non-existent anchor", "name", id)
			}
			continue
		}

		for _, xref := range xrefs {
			//xref.OriginalID = info.ID
			xref.ID = info.ID
			// If the cross reference has a label that's the same as the one we generated for the anchor, remove it
			if len(xref.Set) == 1 {
				if label, ok := xref.Set[0].(*elements.String); ok && info.Label == string(label.Value) {
					clear(xref.Set)
				}
			}
		}
	}
}

func findRefSection(parent any) *ascii.Section {
	switch p := parent.(type) {
	case *ascii.Section:
		return p
	case *ascii.Element:
		return findRefSection(p.Parent)
	case *ascii.Doc:
		return nil
	default:
		return nil
	}

}
