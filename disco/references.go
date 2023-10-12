package disco

import (
	"log/slog"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/matterfmt/ascii"
	"github.com/hasty/matterfmt/matter"
)

func (b *Ball) findCrossReferences(doc *ascii.Doc) map[string][]*types.InternalCrossReference {
	crossReferences := make(map[string][]*types.InternalCrossReference)
	ascii.Traverse(nil, doc.Base.Elements, func(el interface{}, parent ascii.HasElements, index int) bool {
		if icr, ok := el.(*types.InternalCrossReference); ok {
			id, ok := icr.ID.(string)
			if !ok {
				return false
			}
			crossReferences[id] = append(crossReferences[id], icr)
		}
		return false
	})
	return crossReferences
}

func (b *Ball) rewriteCrossReferences(crossReferences map[string][]*types.InternalCrossReference, anchors map[string]*anchorInfo) {
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
			xref.OriginalID = info.id
			xref.ID = info.id
			// If the cross reference has a label that's the same as the one we generated for the anchor, remove it
			if label, ok := xref.Label.(string); ok && label == info.label {
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

func getReferenceName(element interface{}) string {
	switch el := element.(type) {
	case *types.Section:
		name := types.Reduce(el.Title)
		if s, ok := name.(string); ok {
			return s
		}
	case types.WithAttributes:
		attr := el.GetAttributes()
		if attr != nil {
			if title, ok := attr.GetAsString("title"); ok {
				return title
			}
		}
	default:
		//slog.
		slog.Debug("Unknown type to get reference name", "type", element)
	}
	return ""
}
