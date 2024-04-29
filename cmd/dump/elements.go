package dump

import (
	"fmt"
	"strings"

	"github.com/hasty/adoc/elements"
	"github.com/hasty/alchemy/ascii"
)

func dumpElements(doc *ascii.Doc, els []elements.Element, indent int) {

	for _, e := range els {
		fmt.Print(strings.Repeat("\t", indent))
		as, ok := e.(*ascii.Section)
		if ok {
			fmt.Printf("{SEC %d (%s)}:\n", as.Base.Level, as.SecType)

			dumpAttributes(as.Base.Attributes(), indent+1)
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
		case *elements.EmptyLine:
			fmt.Print("{empty}\n")
		/*case *elements.DelimitedBlock:
		fmt.Printf("{delim kind=%s}:\n", el.Kind)
		dumpAttributes(el.Attributes, indent+1)
		dumpElements(doc, el.Elements, indent+1)*/
		case *elements.AttributeEntry:
			fmt.Printf("{attrib}: %s", el.Name)
			dumpElements(doc, el.Elements(), indent+1)
			fmt.Print("\n")
		case *elements.Paragraph:
			fmt.Print("{para}: ")
			fmt.Print("\n")
			dumpAttributes(el.Attributes(), indent+1)
			dumpElements(doc, el.Elements(), indent+1)
		case *elements.Section:
			fmt.Printf("{sec %d}:\n", el.Level)
			dumpAttributes(el.Attributes(), indent+1)
			fmt.Print(strings.Repeat("\t", indent+1))
			fmt.Printf("{title:}\n")
			dumpElements(doc, el.Title, indent+2)
			fmt.Print(strings.Repeat("\t", indent+1))
			fmt.Printf("{body:}\n")
			dumpElements(doc, el.Elements(), indent+2)
		case *elements.String:
			fmt.Print("{str}: ", snippet(el.Value))
			fmt.Print("\n")
		case elements.FormattedTextElement:
			fmt.Printf("{formatted text %d}:\n", el.TextFormat())
			if a, ok := el.(elements.Attributable); ok {
				dumpAttributes(a.Attributes(), indent+1)
			}
			fmt.Print(strings.Repeat("\t", indent+1))
			fmt.Printf("{body:}\n")
			dumpElements(doc, el.Elements(), indent+2)
		case *elements.Table:
			fmt.Print("{tab}:\n")
			dumpAttributes(el.Attributes(), indent+1)
			dumpTable(doc, el, indent+1)
		/*case *elements.List:
		fmt.Print("{list}:\n")
		dumpAttributes(el.Attributes, indent+1)
		dumpElements(doc, el.GetElements(), indent+1)*/
		case *elements.OrderedListItem:
			fmt.Print("{ole}:\n")
			dumpAttributes(el.Attributes(), indent+1)
			dumpElements(doc, el.Elements(), indent+1)
		case *elements.UnorderedListItem:
			fmt.Printf("{uole bs=%s cl=%v}:\n", el.Marker, el.Checklist)
			dumpAttributes(el.Attributes(), indent+1)
			dumpElements(doc, el.Elements(), indent+1)
		case *elements.CrossReference:
			fmt.Printf("{xref id:%v label %v}\n", el.ID, el.Set)
		case *elements.SpecialCharacter:
			fmt.Printf("{sc: %s}\n", el.Character)
		case *elements.Link:
			fmt.Printf("{link: ")
			dumpLocation(el.URL)
			fmt.Print("}\n")
			dumpAttributes(el.Attributes(), indent+1)
		/*case *elements.DocumentHeader:
			fmt.Printf("{head}\n")
			fmt.Print(strings.Repeat("\t", indent+1))
			dumpAttributes(el.Attributes, indent+1)
			fmt.Printf("{title:}\n")
			dumpElements(doc, el.Title, indent+2)
			fmt.Print(strings.Repeat("\t", indent+1))
			fmt.Printf("{body:}\n")
			dumpElements(doc, el.Elements, indent+2)

		case *elements.Preamble:
			fmt.Printf("{preamble}\n")
			fmt.Print(strings.Repeat("\t", indent+1))
			fmt.Printf("{body:}\n")
			dumpElements(doc, el.Elements, indent+2)
			if el.TableOfContents != nil {
				fmt.Print(strings.Repeat("\t", indent+1))
				dumpTOC(el.TableOfContents.Sections, indent+2)
			}
		case elements.DocumentAuthors:
			fmt.Print("{authors}\n")
			for _, a := range el {
				dumpElements(doc, []any{a}, indent+1)
			}
		case *elements.DocumentAuthor:
			fmt.Printf("{author %s", el.Email)
			if el.DocumentAuthorFullName != nil {
				fmt.Printf("( %s %s %s)", el.DocumentAuthorFullName.FirstName, el.DocumentAuthorFullName.MiddleName, el.DocumentAuthorFullName.LastName)
			}
			fmt.Print("}\n")
		case *elements.FootnoteReference:
			fmt.Printf("{footnote ID=%d, Ref=%s}\n", el.ID, el.Ref)
			var fn *elements.Footnote
			for _, f := range doc.Footnotes() {
				if f.ID == el.ID {
					fn = f
					break
				}
			}
			if fn != nil {
				dumpElements(doc, fn.Elements, indent+1)

			}
		*/
		case *elements.InlineImage:
			fmt.Printf("{image: ")
			dumpElements(doc, el.Path, indent+1)
			fmt.Print("}\n")
			dumpAttributes(el.Attributes(), indent+1)
		case *elements.BlockImage:
			fmt.Printf("{imageblock: ")
			dumpElements(doc, el.Path, indent+1)
			fmt.Print("}\n")
			dumpAttributes(el.Attributes(), indent+1)
		case *elements.AttributeReset:
			fmt.Printf("{attr_reset: %s}\n", el.Name)
		/*case *elements.ListElements:
		fmt.Printf("{list els}\n")
		dumpElements(doc, el.Elements, indent+1)*/
		case *elements.ListContinuation:
			fmt.Printf("{list con}\n")
			dumpElements(doc, []elements.Element{el.Child()}, indent+1)
		case *elements.CharacterReplacementReference:
			fmt.Printf("{predef %s}", el.Name)
		default:
			fmt.Printf("unknown render element type: %T\n", el)
		}
	}
}
