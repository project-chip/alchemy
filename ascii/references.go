package ascii

import (
	"fmt"
	"log/slog"
	"slices"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/alchemy/parse"
)

func (doc *Doc) CrossReferences() map[string][]*types.InternalCrossReference {
	doc.Lock()
	if doc.crossReferences != nil {
		doc.Unlock()
		return doc.crossReferences
	}
	doc.crossReferences = make(map[string][]*types.InternalCrossReference)
	parse.Traverse(nil, doc.Base.Elements, func(el interface{}, parent parse.HasElements, index int) bool {
		if icr, ok := el.(*types.InternalCrossReference); ok {
			id, ok := icr.ID.(string)
			if !ok {
				return false
			}
			doc.crossReferences[id] = append(doc.crossReferences[id], icr)
		}
		return false
	})
	doc.Unlock()
	return doc.crossReferences
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
				_ = e.SetElements(n)

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
		switch name := name.(type) {
		case string:
			return name
		case []any:
			var val strings.Builder
			for _, el := range name {
				switch el := el.(type) {
				case *types.StringElement:
					val.WriteString(el.Content)
				case *types.Symbol:
					val.WriteString(el.Name)
				case *types.SpecialCharacter:
					var char string
					switch el.Name {
					case "&":
						char = el.Name
					default:
						slog.Warn("unrecognized special character", "char", el.Name)
					}
					val.WriteString(char)
				case types.WithAttributes:
					val.WriteString(referenceNameFromAttributes(el))
				default:
					slog.Warn("unknown section title element", "element", el, "type", fmt.Sprintf("%T", el))
				}
			}
			return val.String()
		}
	case types.WithAttributes:
		return referenceNameFromAttributes(el)
	default:
		//slog.
		slog.Warn("Unknown type to get reference name", "type", element)
	}
	return ""
}

func referenceNameFromAttributes(el types.WithAttributes) string {
	attr := el.GetAttributes()
	if attr == nil {
		slog.Debug("anchor element has no attributes")
		return ""
	}
	if title, ok := attr.GetAsString("title"); ok {
		if len(title) == 0 {
			slog.Warn("empty section title attribute")
		}
		return title
	}
	slog.Debug("anchor element has no title attribute", "element", el)
	return ""
}
