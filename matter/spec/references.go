package spec

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/hasty/alchemy/asciidoc"
	"github.com/hasty/alchemy/internal/parse"
)

func (doc *Doc) CrossReferences() map[string][]*asciidoc.CrossReference {
	doc.Lock()
	if doc.crossReferences != nil {
		doc.Unlock()
		return doc.crossReferences
	}
	doc.crossReferences = make(map[string][]*asciidoc.CrossReference)
	parse.Traverse(nil, doc.Base.Elements(), func(el any, parent parse.HasElements, index int) parse.SearchShould {
		if icr, ok := el.(*asciidoc.CrossReference); ok {
			doc.crossReferences[icr.ID] = append(doc.crossReferences[icr.ID], icr)
		}
		return parse.SearchShouldContinue
	})
	doc.Unlock()
	return doc.crossReferences
}

func ReferenceName(element any) string {
	switch el := element.(type) {
	case *asciidoc.Section:
		var val strings.Builder
		for _, el := range el.Title {
			switch el := el.(type) {
			case *asciidoc.String:
				val.WriteString(el.Value)
			case *asciidoc.SpecialCharacter:
				var char string
				switch el.Character {
				case "&", ">", "<":
					char = el.Character
				default:
					slog.Warn("unrecognized special character", "char", el.Character, "context", val.String())
				}
				val.WriteString(char)
			case asciidoc.Attributable:
				val.WriteString(referenceNameFromAttributes(el))
			default:
				slog.Warn("unknown section title element", "element", el, "type", fmt.Sprintf("%T", el))
			}
		}
		return val.String()
	case asciidoc.Attributable:
		return referenceNameFromAttributes(el)
	default:
		slog.Warn("Unknown type to get reference name", "type", element)
	}
	return ""
}

func referenceNameFromAttributes(el asciidoc.Attributable) string {
	for _, a := range el.Attributes() {
		switch a := a.(type) {
		case *asciidoc.AnchorAttribute:
		case *asciidoc.TitleAttribute:
			return a.AsciiDocString()
		case *asciidoc.NamedAttribute:
			if a.Name == asciidoc.AttributeNameTitle {
				title, ok := a.Value().(string)
				if !ok {
					slog.Warn("empty section title attribute")
				}
				return title
			}
		default:
			slog.Warn("Unknown attribute type to get reference name", "type", a, "val", a.Value())
		}
	}
	return ""
}
