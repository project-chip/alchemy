package dump

import (
	"fmt"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
)

func dumpElements(reader Reader, parent asciidoc.Parent, els asciidoc.Elements, indent int) {

	for e := range reader.Iterate(parent, els) {
		fmt.Print(strings.Repeat("\t", indent))
		as, ok := e.(*asciidoc.Section)
		if ok {
			fmt.Printf("{SEC %d (%s)}:\n", as.Level, reader.SectionType(as))

			dumpAttributes(reader, as.Attributes(), indent+1)
			fmt.Print(strings.Repeat("\t", indent+1))
			fmt.Printf("{title:}\n")
			dumpElements(reader, as, as.Title, indent+2)
			fmt.Print(strings.Repeat("\t", indent+1))
			fmt.Printf("{body:}\n")
			dumpElements(reader, as, as.Children(), indent+2)
			continue
		}
		switch el := e.(type) {
		case *asciidoc.EmptyLine:
			fmt.Print("{empty}\n")
		case *asciidoc.NewLine:
			fmt.Printf("{newline%s}\n", dumpPosition(el))
		case *asciidoc.LineBreak:
			fmt.Printf("{linebreak%s}\n", dumpPosition(el))
		case *asciidoc.AttributeEntry:
			fmt.Printf("{attrib%s}: %s", dumpPosition(el), el.Name)
			dumpElements(reader, el, reader.Children(el), indent+1)
			fmt.Print("\n")
		case *asciidoc.Paragraph:
			fmt.Printf("{para%s}: ", dumpPosition(el))
			fmt.Print("\n")
			dumpAttributes(reader, el.Attributes(), indent+1)
			dumpElements(reader, el, reader.Children(el), indent+1)
		case *asciidoc.Section:
			fmt.Printf("{sec %d%s}:\n", el.Level, dumpPosition(el))
			dumpAttributes(reader, el.Attributes(), indent+1)
			fmt.Print(strings.Repeat("\t", indent+1))
			fmt.Printf("{title:}\n")
			dumpElements(reader, el, el.Title, indent+2)
			fmt.Print(strings.Repeat("\t", indent+1))
			fmt.Printf("{body:}\n")
			dumpElements(reader, el, reader.Children(el), indent+2)
		case *asciidoc.String:
			fmt.Print("{str}: ", snippet(el.Value))
			fmt.Print("\n")

		case asciidoc.FormattedTextElement:
			fmt.Printf("{formatted text %d%s}:\n", el.TextFormat(), dumpPosition(el))
			if a, ok := el.(asciidoc.Attributable); ok {
				dumpAttributes(reader, a.Attributes(), indent+1)
			}
			fmt.Print(strings.Repeat("\t", indent+1))
			fmt.Printf("{body:}\n")
			dumpElements(reader, el, reader.Children(el), indent+2)
		case *asciidoc.Table:
			fmt.Printf("{tab%s}:\n", dumpPosition(el))
			dumpAttributes(reader, el.Attributes(), indent+1)
			dumpTable(reader, el, indent+1)
		case *asciidoc.IfDef:
			fmt.Print("{ifdef ")
			dumpConditional(el.Attributes, el.Union, indent)
		case *asciidoc.IfNDef:
			fmt.Print("{ifndef ")
			dumpConditional(el.Attributes, el.Union, indent)
		case *asciidoc.EndIf:
			fmt.Print("{endif ")
			dumpConditional(el.Attributes, el.Union, indent)
		case *asciidoc.OrderedListItem:
			fmt.Print("{ole}:\n")
			dumpAttributes(reader, el.Attributes(), indent+1)
			dumpElements(reader, el, reader.Children(el), indent+1)
		case *asciidoc.UnorderedListItem:
			fmt.Printf("{uole bs=%s cl=%v}:\n", el.Marker, el.Checklist)
			dumpAttributes(reader, el.Attributes(), indent+1)
			dumpElements(reader, el, reader.Children(el), indent+1)
		case *asciidoc.CrossReference:
			fmt.Printf("{xref}\n")
			dumpElements(reader, el, el.ID, indent+1)
			if len(el.Elements) > 0 {
				fmt.Print(strings.Repeat("\t", indent))
				fmt.Printf("{label}\n")
				dumpElements(reader, el, el.Elements, indent+1)
			}
		case *asciidoc.DocumentCrossReference:
			fmt.Printf("{doc xref}\n")
			dumpAttributes(reader, el.Attributes(), indent+1)
			dumpElements(reader, &el.ReferencePath, el.ReferencePath, indent+1)
		case asciidoc.SpecialCharacter:
			fmt.Printf("{sc: %s}\n", el.Character)
		case *asciidoc.Link:
			fmt.Printf("{link: ")
			dumpLocation(reader, el.URL, indent+1)
			fmt.Print("}\n")
			dumpAttributes(reader, el.Attributes(), indent+1)
		case *asciidoc.LinkMacro:
			fmt.Printf("{link macro: ")
			dumpLocation(reader, el.URL, indent+1)
			fmt.Print("}\n")
			dumpAttributes(reader, el.Attributes(), indent+1)
		case *asciidoc.FileInclude:
			fmt.Printf("{include:\n")
			dumpAttributes(reader, el.Attributes(), indent+1)
			dumpElements(reader, el, reader.Children(el), indent+1)
			fmt.Print(strings.Repeat("\t", indent))
			fmt.Print("}\n")
		/*case *asciidoc.DocumentHeader:
			fmt.Printf("{head}\n")
			fmt.Print(strings.Repeat("\t", indent+1))
			dumpAttributes(el.Attributes, indent+1)
			fmt.Printf("{title:}\n")
			dumpElements( el.Title, indent+2)
			fmt.Print(strings.Repeat("\t", indent+1))
			fmt.Printf("{body:}\n")
			dumpElements( el.Elements, indent+2)

		case *asciidoc.Preamble:
			fmt.Printf("{preamble}\n")
			fmt.Print(strings.Repeat("\t", indent+1))
			fmt.Printf("{body:}\n")
			dumpElements( el.Elements, indent+2)
			if el.TableOfContents != nil {
				fmt.Print(strings.Repeat("\t", indent+1))
				dumpTOC(el.TableOfContents.Sections, indent+2)
			}
		case asciidoc.DocumentAuthors:
			fmt.Print("{authors}\n")
			for _, a := range el {
				dumpElements( []any{a}, indent+1)
			}
		case *asciidoc.DocumentAuthor:
			fmt.Printf("{author %s", el.Email)
			if el.DocumentAuthorFullName != nil {
				fmt.Printf("( %s %s %s)", el.DocumentAuthorFullName.FirstName, el.DocumentAuthorFullName.MiddleName, el.DocumentAuthorFullName.LastName)
			}
			fmt.Print("}\n")
		case *asciidoc.FootnoteReference:
			fmt.Printf("{footnote ID=%d, Ref=%s}\n", el.ID, el.Ref)
			var fn *asciidoc.Footnote
			for _, f := range doc.Footnotes() {
				if f.ID == el.ID {
					fn = f
					break
				}
			}
			if fn != nil {
				dumpElements( fn.Elements, indent+1)

			}
		*/
		case *asciidoc.InlineImage:
			fmt.Printf("{image: ")
			dumpElements(reader, &el.ImagePath, el.ImagePath, indent+1)
			fmt.Print("}\n")
			dumpAttributes(reader, el.Attributes(), indent+1)
		case *asciidoc.BlockImage:
			fmt.Printf("{imageblock: ")
			dumpElements(reader, &el.ImagePath, el.ImagePath, indent+1)
			fmt.Print("}\n")
			dumpAttributes(reader, el.Attributes(), indent+1)
		case *asciidoc.AttributeReset:
			fmt.Printf("{attr_reset: %s}\n", el.Name)
		case *asciidoc.Listing:
			fmt.Printf("{listing\n")
			dumpAttributes(reader, el.Attributes(), indent+1)
			for i, l := range el.Lines() {
				fmt.Print(strings.Repeat("\t", indent+1))
				fmt.Printf("%d: \"%s\"\n", i, l)
			}
			fmt.Print(strings.Repeat("\t", indent))
			fmt.Printf("}\n")
		case *asciidoc.SingleLineComment:
			fmt.Print("{comment}: ", snippet(el.Value))
			fmt.Print("\n")
		case *asciidoc.Anchor:
			fmt.Printf("{anchor \"%s\"", el.ID)
			if len(el.Elements) > 0 {
				fmt.Print("\n")
				dumpElements(reader, el, reader.Children(el), indent+1)
			}
			fmt.Print("}\n")
		/*case *asciidoc.ListElements:
		fmt.Printf("{list els}\n")
		dumpElements( el.Elements, indent+1)*/
		case *asciidoc.ListContinuation:
			fmt.Printf("{list continuation (new line count: %d)}\n", el.NewLineCount)
			children := asciidoc.Elements{el.Child()}
			dumpElements(reader, &children, children, indent+1)
		case *asciidoc.CharacterReplacementReference:
			fmt.Printf("{predef %s}\n", el.Name())
		case *asciidoc.QuoteBlock:
			fmt.Printf("{quote delimiter:\"%d\"}:\n", el.Delimiter.Type)
			dumpAttributes(reader, el.Attributes(), indent+1)
			dumpElements(reader, el, reader.Children(el), indent+1)
		case *asciidoc.UserAttributeReference:
			fmt.Printf("{user attribute ref:\"%s\"}:\n", el.Value)
		default:
			fmt.Printf("unknown render element type: %T\n", el)
		}
	}
}

func dumpConditional(attributes asciidoc.AttributeNames, union asciidoc.ConditionalUnion, indent int) {
	if len(attributes) == 0 {
		fmt.Print("}\n")
		return
	}
	if union == asciidoc.ConditionalUnionAll {
		fmt.Print("all of:\n")
	} else {
		fmt.Print("any of:\n")
	}
	for _, s := range attributes {
		fmt.Print(strings.Repeat("\t", indent+1))
		fmt.Printf("%s\n", s)
	}
	fmt.Print(strings.Repeat("\t", indent))
	fmt.Print("}\n")
}

func dumpPosition(el asciidoc.Element) string {
	if hp, ok := el.(asciidoc.HasPosition); ok {
		l, c, _ := hp.Position()
		return fmt.Sprintf(" %d:%d", l, c)
	}
	return ""
}
