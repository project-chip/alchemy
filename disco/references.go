package disco

import (
	"log/slog"
	"strings"

	"github.com/hasty/adoc/asciidoc"
	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/matter"
)

func (b *Ball) rewriteCrossReferences(crossReferences map[string][]*asciidoc.CrossReference, anchors map[string]*ascii.Anchor) {
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
