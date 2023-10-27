package disco

import (
	"log/slog"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/matterfmt/ascii"
	"github.com/hasty/matterfmt/matter"
)

func (b *Ball) rewriteCrossReferences(crossReferences map[string][]*types.InternalCrossReference, anchors map[string]*ascii.Anchor) {
	for id, xrefs := range crossReferences {
		info, ok := anchors[id]
		if !ok {
			dt, _ := b.doc.DocType()
			switch dt {
			case matter.DocTypeAppCluster, matter.DocTypeDeviceType:
				slog.Warn("cross reference points to non-existent anchor", "name", id)
			}
			continue
		}

		for _, xref := range xrefs {
			xref.OriginalID = info.ID
			xref.ID = info.ID
			// If the cross reference has a label that's the same as the one we generated for the anchor, remove it
			if label, ok := xref.Label.(string); ok && label == info.Label {
				xref.Label = nil
			}
		}
	}
}

func findRefSection(parent interface{}) *ascii.Section {
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
