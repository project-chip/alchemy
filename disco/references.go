package disco

import (
	"log/slog"
	"slices"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/matterfmt/ascii"
	"github.com/hasty/matterfmt/matter"
)

func getUnrecognizedReference(els []interface{}) string {
	if len(els) < 5 {
		return ""
	}
	sc, ok := els[0].(*types.SpecialCharacter)
	if !ok || sc.Name != "<" {
		return ""
	}
	sc, ok = els[1].(*types.SpecialCharacter)
	if !ok || sc.Name != "<" {
		return ""
	}
	s, ok := els[2].(*types.StringElement)
	if !ok {
		return ""
	}
	sc, ok = els[3].(*types.SpecialCharacter)
	if !ok || sc.Name != ">" {
		return ""
	}
	sc, ok = els[4].(*types.SpecialCharacter)
	if !ok || sc.Name != ">" {
		return ""
	}
	return s.Content
}

// The parser sometimes doesn't recognize inline references, and just returns them as "<", "<", "ref", ">", ">"
// Rather than doing the right thing and fixing the parser, we'll just find these little suckers after the fact
func fixUnrecognizedReferences(doc *ascii.Doc) {
	var elementsWithUnrecognizedReferences []ascii.HasElements
	ascii.Traverse(nil, doc.Base.Elements, func(el interface{}, parent ascii.HasElements, index int) bool {
		sc, ok := el.(*types.SpecialCharacter)
		if !ok {
			return false
		}
		if sc.Name != "<" {
			return false
		}
		els := parent.GetElements()[index:]
		if getUnrecognizedReference(els) == "" {
			return false
		}
		elementsWithUnrecognizedReferences = append(elementsWithUnrecognizedReferences, parent)
		return false
	})
	for _, e := range elementsWithUnrecognizedReferences {
		els := e.GetElements()
		for i := 0; i < len(els); i++ {
			sub := els[i:]
			ref := getUnrecognizedReference(sub)
			if ref != "" {
				var id string
				var label string
				parts := strings.Split(ref, ",")
				id = parts[0]
				if len(parts) > 0 {
					label = parts[1]
				}
				icr, _ := types.NewInternalCrossReference(id, label)
				n := slices.Replace(els, i, i+5, interface{}(icr))
				e.SetElements(n)

				break
			}
		}
	}
}

func findCrossReferences(doc *ascii.Doc) map[string][]*types.InternalCrossReference {
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

func rewriteCrossReferences(crossReferences map[string][]*types.InternalCrossReference, anchors map[string]*anchorInfo) {
	for id, xrefs := range crossReferences {
		info, ok := anchors[id]
		if !ok {
			slog.Warn("cross reference points to non-existent anchor", "name", id)
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

func stripReferenceSuffixes(newId string) string {
	for _, suffix := range matter.DisallowedReferenceSuffixes {
		if strings.HasSuffix(newId, suffix) {
			newId = newId[0 : len(newId)-len(suffix)]
			break
		}
	}
	return newId
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
	case *types.Table:
		if el.Attributes != nil {

			if title, ok := el.Attributes["title"]; ok {
				if name, ok := title.(string); ok {
					return name
				}
			}
		}
	default:
		//slog.
		slog.Debug("Unknown type to get reference name", "type", element)
	}
	return ""
}
