package spec

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/internal/parse"
	"github.com/project-chip/alchemy/matter"
)

type CrossReference struct {
	Document  *Doc
	Reference *asciidoc.CrossReference
	Parent    parse.HasElements
	Source    matter.Source
}

func (cr *CrossReference) SyncToDoc(id string) {
	if id != cr.Reference.ID {
		cr.Document.changeCrossReference(cr, id)
		if cr.Document.group != nil {
			cr.Document.group.changeCrossReference(cr, id)
		}
		cr.Reference.ID = id
	}
}

func (doc *Doc) CrossReferences() map[string][]*CrossReference {
	doc.Lock()
	if doc.crossReferencesParsed {
		doc.Unlock()
		return doc.crossReferences
	}
	parse.Traverse(nil, doc.Base.Elements(), func(icr *asciidoc.CrossReference, parent parse.HasElements, index int) parse.SearchShould {
		doc.crossReferences[icr.ID] = append(doc.crossReferences[icr.ID], &CrossReference{Document: doc, Reference: icr, Parent: parent, Source: NewSource(doc, icr)})
		return parse.SearchShouldContinue
	})
	doc.crossReferencesParsed = true
	doc.Unlock()
	return doc.crossReferences
}

func ReferenceName(element any) string {
	if element == nil {
		return ""
	}
	switch el := element.(type) {
	case *asciidoc.Anchor:
		return buildReferenceName(el.Set)
	case *asciidoc.Section:
		return buildReferenceName(el.Title)
	case asciidoc.Attributable:
		return referenceNameFromAttributes(el)
	case *Element:
		if el == nil {
			slog.Warn("nil element in reference name", "type", fmt.Sprintf("%T", el))
			return ""
		}
		if el.Base == nil {
			slog.Warn("untethered element in reference name", "type", fmt.Sprintf("%T", el))
			return ""
		}
		return ReferenceName(el.Base)
	default:
		slog.Warn("Unknown type to get reference name", "type", fmt.Sprintf("%T", element))
	}
	return ""
}

func buildReferenceName(set asciidoc.Set) string {
	var val strings.Builder

	for _, el := range set {
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
		case *asciidoc.CharacterReplacementReference:
			val.WriteString(el.Value)
		case *asciidoc.Counter:
		default:
			slog.Warn("unknown reference name element", "element", el, "type", fmt.Sprintf("%T", el))
		}
	}
	return val.String()
}

func referenceNameFromAttributes(el asciidoc.Attributable) string {
	for _, a := range el.Attributes() {
		switch a := a.(type) {
		case *asciidoc.AnchorAttribute:
		case *asciidoc.TableColumnsAttribute:
		case *asciidoc.PositionalAttribute:
		case *asciidoc.ShorthandAttribute:
		case *asciidoc.TitleAttribute:
			return a.AsciiDocString()
		case *asciidoc.NamedAttribute:
			if a.Name == asciidoc.AttributeNameTitle {
				switch v := a.Value().(type) {
				case string:
					return v
				case asciidoc.Set:
					return buildReferenceName(v)
				default:
					slog.Warn("unexpected value of section title attribute", slog.String("type", fmt.Sprintf("%T", a.Value())))
				}
			}
		default:
			slog.Warn("Unknown attribute type to get reference name", "type", fmt.Sprintf("%T", a), "val", a.Value())
		}
	}
	return ""
}
