package ascii

import (
	"log/slog"
	"slices"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/alchemy/parse"
)

func (doc *Doc) CrossReferences() map[string][]*types.InternalCrossReference {
	crossReferences := make(map[string][]*types.InternalCrossReference)
	parse.Traverse(nil, doc.Base.Elements, func(el interface{}, parent parse.HasElements, index int) bool {
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

// The parser sometimes doesn't recognize inline references, and just returns them as "<", "<", "ref", ">", ">"
// Rather than doing the right thing and fixing the parser, we'll just find these little suckers after the fact
func PatchUnrecognizedReferences(doc *Doc) {
	var elementsWithUnrecognizedReferences []parse.HasElements
	parse.Traverse(nil, doc.Base.Elements, func(el interface{}, parent parse.HasElements, index int) bool {
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
				if len(parts) > 1 {
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

func ReferenceName(element interface{}) string {
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
