package spec

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/asciidoc/parse"
	"github.com/project-chip/alchemy/matter"
)

type CrossReference struct {
	Document  *Doc
	Reference *asciidoc.CrossReference
	Parent    asciidoc.Parent
	Source    matter.Source
}

func (cr *CrossReference) Identifier(reader asciidoc.Reader) string {
	return cr.Document.anchorId(reader, cr.Parent, cr.Reference, cr.Reference.ID)
}

func (cr *CrossReference) SyncToDoc(id asciidoc.Elements) {
	if !id.Equals(cr.Reference.ID) {
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
	parse.Search(doc.Reader(), nil, doc.Base.Children(), func(icr *asciidoc.CrossReference, parent asciidoc.Parent, index int) parse.SearchShould {
		referenceID := doc.anchorId(doc.Reader(), icr, icr, icr.ID)
		doc.crossReferences[referenceID] = append(doc.crossReferences[referenceID], &CrossReference{Document: doc, Reference: icr, Parent: parent, Source: NewSource(doc, icr)})
		return parse.SearchShouldContinue
	})
	doc.crossReferencesParsed = true
	doc.Unlock()
	return doc.crossReferences
}

func ReferenceName(reader asciidoc.Reader, element any) string {
	if element == nil {
		return ""
	}
	switch el := element.(type) {
	case *asciidoc.Anchor:
		return buildReferenceName(reader, el, el.Elements)
	case *asciidoc.Section:
		return buildReferenceName(reader, el, el.Title)
	case asciidoc.Attributable:
		return referenceNameFromAttributes(reader, el)
	case asciidoc.Element:
		return ReferenceName(reader, el)
	default:
		slog.Warn("Unknown type to get reference name", "type", fmt.Sprintf("%T", element))
	}
	return ""
}

func buildReferenceName(reader asciidoc.Reader, parent asciidoc.Parent, set asciidoc.Elements) string {
	var val strings.Builder

	for el := range reader.Iterate(parent, set) {
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
			val.WriteString(referenceNameFromAttributes(reader, el))
		case *asciidoc.CharacterReplacementReference:
			val.WriteString(el.Value)
		case *asciidoc.Counter:
		default:
			slog.Warn("unknown reference name element", "element", el, "type", fmt.Sprintf("%T", el))
		}
	}
	return val.String()
}

func referenceNameFromAttributes(reader asciidoc.Reader, el asciidoc.Attributable) string {
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
				case asciidoc.Elements:
					return buildReferenceName(reader, &a.Val, v)
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
