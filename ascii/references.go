package ascii

import (
	"fmt"
	"log/slog"
	"slices"
	"strings"

	"github.com/hasty/adoc/elements"
	"github.com/hasty/alchemy/internal/parse"
)

func (doc *Doc) CrossReferences() map[string][]*elements.InternalCrossReference {
	doc.Lock()
	if doc.crossReferences != nil {
		doc.Unlock()
		return doc.crossReferences
	}
	doc.crossReferences = make(map[string][]*elements.InternalCrossReference)
	parse.Traverse(nil, doc.Base.Elements, func(el any, parent parse.HasElements, index int) bool {
		if icr, ok := el.(*elements.InternalCrossReference); ok {
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
	parse.Traverse(nil, doc.Base.Elements, func(el any, parent parse.HasElements, index int) bool {
		sc, ok := el.(*elements.SpecialCharacter)
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
				icr, _ := elements.NewInternalCrossReference(id, label)
				n := slices.Replace(els, i, i+5, any(icr))
				_ = e.SetElements(n)

				break
			}
		}
	}
}

func getUnrecognizedReference(els []any) string {
	if len(els) < 5 {
		return ""
	}
	sc, ok := els[0].(*elements.SpecialCharacter)
	if !ok || sc.Name != "<" {
		return ""
	}
	sc, ok = els[1].(*elements.SpecialCharacter)
	if !ok || sc.Name != "<" {
		return ""
	}
	s, ok := els[2].(*elements.String)
	if !ok {
		return ""
	}
	sc, ok = els[3].(*elements.SpecialCharacter)
	if !ok || sc.Name != ">" {
		return ""
	}
	sc, ok = els[4].(*elements.SpecialCharacter)
	if !ok || sc.Name != ">" {
		return ""
	}
	return s.Content
}

func ReferenceName(element any) string {
	switch el := element.(type) {
	case *elements.Section:
		var val strings.Builder
		for _, el := range el.Title {
			switch el := el.(type) {
			case *elements.String:
				val.WriteString(el.Content)
			case *elements.Symbol:
				val.WriteString(el.Name)
			case *elements.SpecialCharacter:
				var char string
				switch el.Name {
				case "&", ">", "<":
					char = el.Name
				default:
					slog.Warn("unrecognized special character", "char", el.Name, "context", val.String())
				}
				val.WriteString(char)
			case elements.Attributable:
				val.WriteString(referenceNameFromAttributes(el))
			default:
				slog.Warn("unknown section title element", "element", el, "type", fmt.Sprintf("%T", el))
			}
		}
		return val.String()
	case elements.Attributable:
		return referenceNameFromAttributes(el)
	default:
		slog.Warn("Unknown type to get reference name", "type", element)
	}
	return ""
}

func referenceNameFromAttributes(el elements.Attributable) string {
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
	return ""
}
