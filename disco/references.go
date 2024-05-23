package disco

import (
	"log/slog"
	"strings"

	"github.com/hasty/alchemy/asciidoc"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/spec"
)

func (b *Ball) rewriteCrossReferences(crossReferences map[string][]*asciidoc.CrossReference, anchors map[string]*spec.Anchor) {
	for id, xrefs := range crossReferences {
		info, ok := anchors[id]
		if !ok {
			dt, _ := b.doc.DocType()
			switch dt {
			case matter.DocTypeCluster, matter.DocTypeDeviceType:
				slog.Info("cross reference points to non-existent anchor", "name", id)
			}
			continue
		}

		for _, xref := range xrefs {
			//xref.OriginalID = info.ID
			xref.ID = info.ID

			// If the cross reference has a label that's the same as the one we generated for the anchor, remove it
			if len(xref.Set) == 1 {
				if label, ok := xref.Set[0].(*asciidoc.String); ok {
					labelString := strings.TrimSpace(asciidoc.AttributeAsciiDocString(info.LabelElements))
					if labelString == string(label.Value) {
						xref.Set = nil
					}
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
