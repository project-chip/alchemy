package ascii

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/hasty/adoc/elements"
	"github.com/hasty/alchemy/internal/parse"
)

func (doc *Doc) CrossReferences() map[string][]*elements.CrossReference {
	doc.Lock()
	if doc.crossReferences != nil {
		doc.Unlock()
		return doc.crossReferences
	}
	doc.crossReferences = make(map[string][]*elements.CrossReference)
	parse.Traverse(nil, doc.Base.Elements(), func(el any, parent parse.HasElements, index int) parse.SearchShould {
		if icr, ok := el.(*elements.CrossReference); ok {
			doc.crossReferences[icr.ID] = append(doc.crossReferences[icr.ID], icr)
		}
		return parse.SearchShouldContinue
	})
	doc.Unlock()
	return doc.crossReferences
}

func ReferenceName(element any) string {
	switch el := element.(type) {
	case *elements.Section:
		var val strings.Builder
		for _, el := range el.Title {
			switch el := el.(type) {
			case *elements.String:
				val.WriteString(el.Value)
			case *elements.SpecialCharacter:
				var char string
				switch el.Character {
				case "&", ">", "<":
					char = el.Character
				default:
					slog.Warn("unrecognized special character", "char", el.Character, "context", val.String())
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
	for _, a := range el.Attributes() {
		switch a := a.(type) {
		case *elements.TitleAttribute:
			return a.AsciiDocString()
		case *elements.NamedAttribute:
			if a.Name == elements.AttributeNameTitle {
				title, ok := a.Value().(string)
				if !ok {
					slog.Warn("empty section title attribute")
				}
				return title
			}
		}
	}
	return ""
}
