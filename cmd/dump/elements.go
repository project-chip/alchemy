package dump

import (
	"fmt"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/alchemy/ascii"
)

func dumpElements(doc *ascii.Doc, elements []any, indent int) {

	for _, e := range elements {
		fmt.Print(strings.Repeat("\t", indent))
		as, ok := e.(*ascii.Section)
		if ok {
			fmt.Printf("{SEC %d (%s)}:\n", as.Base.Level, as.SecType)

			dumpAttributes(as.Base.Attributes, indent+1)
			fmt.Print(strings.Repeat("\t", indent+1))
			fmt.Printf("{title:}\n")
			dumpElements(doc, as.Base.Title, indent+2)
			fmt.Print(strings.Repeat("\t", indent+1))
			fmt.Printf("{body:}\n")
			dumpElements(doc, as.Elements, indent+2)
			continue
		}
		ae, ok := e.(*ascii.Element)
		if ok {
			e = ae.Base
		}
		switch el := e.(type) {
		case *types.BlankLine:
			fmt.Print("{blank}\n")
		case *types.DelimitedBlock:
			fmt.Printf("{delim kind=%s}:\n", el.Kind)
			dumpAttributes(el.Attributes, indent+1)
			dumpElements(doc, el.Elements, indent+1)
		case *types.AttributeDeclaration:
			fmt.Printf("{attrib}: %s", el.Name)
			dumpElements(doc, []any{el.Value}, indent+1)
			fmt.Print("\n")
		case *types.Paragraph:
			fmt.Print("{para}: ")
			fmt.Print("\n")
			dumpAttributes(el.Attributes, indent+1)
			dumpElements(doc, el.Elements, indent+1)
		case *types.Section:
			fmt.Printf("{sec %d}:\n", el.Level)
			dumpAttributes(el.Attributes, indent+1)
			fmt.Print(strings.Repeat("\t", indent+1))
			fmt.Printf("{title:}\n")
			dumpElements(doc, el.Title, indent+2)
			fmt.Print(strings.Repeat("\t", indent+1))
			fmt.Printf("{body:}\n")
			dumpElements(doc, el.Elements, indent+2)
		case *types.StringElement:
			fmt.Print("{se}: ", snippet(el.Content))
			fmt.Print("\n")
		case string:
			fmt.Print("{str}: ", snippet(el))
			fmt.Print("\n")
		case *types.QuotedText:
			fmt.Printf("{qt %s}:\n", el.Kind)
			dumpAttributes(el.Attributes, indent+1)
			fmt.Print(strings.Repeat("\t", indent+1))
			fmt.Printf("{body:}\n")
			dumpElements(doc, el.Elements, indent+2)
		case *types.Table:
			fmt.Print("{tab}:\n")
			dumpAttributes(el.Attributes, indent+1)
			dumpTable(doc, el, indent+1)
		case *types.List:
			fmt.Print("{list}:\n")
			dumpAttributes(el.Attributes, indent+1)
			dumpElements(doc, el.GetElements(), indent+1)
		case *types.OrderedListElement:
			fmt.Print("{ole}:\n")
			dumpAttributes(el.Attributes, indent+1)
			dumpElements(doc, el.GetElements(), indent+1)
		case *types.UnorderedListElement:
			fmt.Printf("{uole bs=%s cs=%s}:\n", el.BulletStyle, el.CheckStyle)
			dumpAttributes(el.Attributes, indent+1)
			dumpElements(doc, el.GetElements(), indent+1)
		case *types.InternalCrossReference:
			fmt.Printf("{xref id:%v label %v}\n", el.ID, el.Label)
		case *types.SpecialCharacter:
			fmt.Printf("{sc: %s}\n", el.Name)
		case *types.Symbol:
			fmt.Printf("{sym: %s}\n", el.Name)
		case *types.InlineLink:
			fmt.Printf("{link: ")
			dumpLocation(el.Location)
			fmt.Print("}\n")
			dumpAttributes(el.Attributes, indent+1)
		case *types.DocumentHeader:
			fmt.Printf("{head}\n")
			fmt.Print(strings.Repeat("\t", indent+1))
			dumpAttributes(el.Attributes, indent+1)
			fmt.Printf("{title:}\n")
			dumpElements(doc, el.Title, indent+2)
			fmt.Print(strings.Repeat("\t", indent+1))
			fmt.Printf("{body:}\n")
			dumpElements(doc, el.Elements, indent+2)

		case *types.Preamble:
			fmt.Printf("{preamble}\n")
			fmt.Print(strings.Repeat("\t", indent+1))
			fmt.Printf("{body:}\n")
			dumpElements(doc, el.Elements, indent+2)
			if el.TableOfContents != nil {
				fmt.Print(strings.Repeat("\t", indent+1))
				dumpTOC(el.TableOfContents.Sections, indent+2)
			}
		case types.DocumentAuthors:
			fmt.Print("{authors}\n")
			for _, a := range el {
				dumpElements(doc, []any{a}, indent+1)
			}
		case *types.DocumentAuthor:
			fmt.Printf("{author %s", el.Email)
			if el.DocumentAuthorFullName != nil {
				fmt.Printf("( %s %s %s)", el.DocumentAuthorFullName.FirstName, el.DocumentAuthorFullName.MiddleName, el.DocumentAuthorFullName.LastName)
			}
			fmt.Print("}\n")
		case *types.FootnoteReference:
			fmt.Printf("{footnote ID=%d, Ref=%s}\n", el.ID, el.Ref)
			var fn *types.Footnote
			for _, f := range doc.Footnotes() {
				if f.ID == el.ID {
					fn = f
					break
				}
			}
			if fn != nil {
				dumpElements(doc, fn.Elements, indent+1)

			}
		case *types.InlineImage:
			fmt.Printf("{image: ")
			dumpLocation(el.Location)
			fmt.Print("}\n")
			dumpAttributes(el.Attributes, indent+1)
		case *types.ImageBlock:
			fmt.Printf("{imageblock: ")
			dumpLocation(el.Location)
			fmt.Print("}\n")
			dumpAttributes(el.Attributes, indent+1)
		case *types.AttributeReset:
			fmt.Printf("{attr_reset: %s}\n", el.Name)
		case *types.ListElements:
			fmt.Printf("{list els}\n")
			dumpElements(doc, el.Elements, indent+1)
		case *types.ListContinuation:
			fmt.Printf("{list con %d}\n", el.Offset)
			dumpElements(doc, []any{el.Element}, indent+1)
		case *types.PredefinedAttribute:
			fmt.Printf("{predef %s}", el.Name)
		default:
			fmt.Printf("unknown render element type: %T\n", el)
		}
	}
}
