package parse

import (
	"fmt"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/matterfmt/output"
	"github.com/hasty/matterfmt/render"
)

func Dump(cxt *output.Context) {
	DumpElements(cxt, cxt.Doc.Base.Elements, 0)
}

func DumpElements(cxt *output.Context, elements []interface{}, indent int) {
	for _, e := range elements {
		fmt.Print(strings.Repeat("\t", indent))
		switch el := e.(type) {
		case *types.BlankLine:
			fmt.Print("{blank}\n")
		case *types.DelimitedBlock:
			fmt.Print("{delim}: ")
			dumpAttributes(cxt, el.Attributes, indent+1)
			switch el.Kind {
			case "comment":
				renderContext := output.NewContext(cxt, cxt.Doc)
				render.RenderElements(renderContext, "", []interface{}{el})
				fmt.Print(snippet(renderContext.String()))
			}
			fmt.Print("\n")
		case *types.AttributeDeclaration:
			fmt.Printf("{attrib}: %s", el.Name)
			DumpElements(cxt, []interface{}{el.Value}, indent+1)
			fmt.Print("\n")
		case *types.Paragraph:

			fmt.Print("{para}: ")
			fmt.Print("\n")
			dumpAttributes(cxt, el.Attributes, indent+1)
			DumpElements(cxt, el.Elements, indent+1)
		case *types.Section:
			fmt.Printf("{sec %d}:\n", el.Level)
			dumpAttributes(cxt, el.Attributes, indent+1)
			fmt.Print(strings.Repeat("\t", indent+1))
			fmt.Printf("{title:}\n")
			DumpElements(cxt, el.Title, indent+2)
			fmt.Print(strings.Repeat("\t", indent+1))
			fmt.Printf("{body:}\n")
			DumpElements(cxt, el.Elements, indent+2)
		case *types.StringElement:
			fmt.Print("{se}: ", snippet(el.Content))
			fmt.Print("\n")
		case string:
			fmt.Print("{str}: ", snippet(el))
			fmt.Print("\n")
		case *types.QuotedText:
			fmt.Printf("{qt %s}:\n", el.Kind)
			dumpAttributes(cxt, el.Attributes, indent+1)
			fmt.Print(strings.Repeat("\t", indent+1))
			fmt.Printf("{body:}\n")
			DumpElements(cxt, el.Elements, indent+2)
		case *types.Table:
			fmt.Print("{tab}:\n")
			dumpAttributes(cxt, el.Attributes, indent+1)
			dumpTable(cxt, el, indent+1)
		case *types.List:
			fmt.Print("{list}:\n")
			dumpAttributes(cxt, el.Attributes, indent+1)
			DumpElements(cxt, el.GetElements(), indent+1)
		case *types.OrderedListElement:
			fmt.Print("{ole}:\n")
			dumpAttributes(cxt, el.Attributes, indent+1)
			DumpElements(cxt, el.GetElements(), indent+1)
		case *types.UnorderedListElement:
			fmt.Print("{uole}:\n")
			dumpAttributes(cxt, el.Attributes, indent+1)
			DumpElements(cxt, el.GetElements(), indent+1)
		case *types.InternalCrossReference:
			fmt.Print("{xref}\n")
		case *types.SpecialCharacter:
			fmt.Printf("{sc: %s}\n", el.Name)
		case *types.Symbol:
			fmt.Printf("{sym: %s}\n", el.Name)
		case *types.InlineLink:
			fmt.Printf("{link: ")
			if el.Location != nil {
				fmt.Printf("%s %s}", el.Location.Scheme, el.Location.Path.(string))
			} else {
				fmt.Printf("missing location")
			}
			fmt.Print("}\n")
			dumpAttributes(cxt, el.Attributes, indent+1)
		case *types.DocumentHeader:
			fmt.Printf("{head}\n")
			fmt.Print(strings.Repeat("\t", indent+1))
			dumpAttributes(cxt, el.Attributes, indent+1)
			fmt.Printf("{title:}\n")
			DumpElements(cxt, el.Title, indent+2)
			fmt.Print(strings.Repeat("\t", indent+1))
			fmt.Printf("{body:}\n")
			DumpElements(cxt, el.Elements, indent+2)

		case *types.Preamble:
			fmt.Printf("{preamble}\n")
			fmt.Print(strings.Repeat("\t", indent+1))
			fmt.Printf("{body:}\n")
			DumpElements(cxt, el.Elements, indent+2)
			if el.TableOfContents != nil {
				fmt.Print(strings.Repeat("\t", indent+1))
				dumpTOC(cxt, el.TableOfContents.Sections, indent+2)
			}
		case types.DocumentAuthors:
			fmt.Print("{authors}\n")
			for _, a := range el {
				DumpElements(cxt, []interface{}{a}, indent+1)
			}
		case *types.DocumentAuthor:
			fmt.Printf("{author %s", el.Email)
			if el.DocumentAuthorFullName != nil {
				fmt.Printf("( %s %s %s)", el.DocumentAuthorFullName.FirstName, el.DocumentAuthorFullName.MiddleName, el.DocumentAuthorFullName.LastName)
			}
			fmt.Print("}\n")
		default:
			fmt.Printf("unknown element type: %T\n", el)
		}
	}
}

func dumpAttributes(cxt *output.Context, attributes types.Attributes, indent int) {
	if len(attributes) == 0 {
		return
	}
	fmt.Print(strings.Repeat("\t", indent))
	fmt.Print("{attr:")
	for key, val := range attributes {
		fmt.Printf(" %s=", key)
		switch v := val.(type) {
		case *types.StringElement:
			fmt.Print(v.Content)
		case string:
			fmt.Print(v)
		default:
			fmt.Printf("unknown type: %T", val)
		}
	}
	fmt.Print("}\n")
}

func dumpTOC(cxt *output.Context, tocs []*types.ToCSection, indent int) {
	for _, toc := range tocs {
		fmt.Print(strings.Repeat("\t", indent))
		fmt.Printf("{toc %d} %s.%s\n", toc.Level, toc.Number, toc.Title)
		if len(toc.Children) > 0 {
			dumpTOC(cxt, toc.Children, indent+1)
		}
	}

}

func dumpTable(cxt *output.Context, tbl *types.Table, indent int) {
	fmt.Print(strings.Repeat(" ", indent*5))
	fmt.Print("{head}:\n")
	dumpTableCells(cxt, tbl.Header.Cells, indent+1)
	fmt.Print(strings.Repeat("\t", indent))
	fmt.Print("{body}:\n")
	for _, row := range tbl.Rows {
		dumpTableRow(cxt, row, indent+1)
	}
}

func dumpTableCells(cxt *output.Context, cells []*types.TableCell, indent int) {
	for _, c := range cells {
		fmt.Print(strings.Repeat("\t", indent))
		fmt.Print("{cell}:\n")
		DumpElements(cxt, c.Elements, indent+1)
	}

}

func dumpTableRow(cxt *output.Context, row *types.TableRow, indent int) {
	fmt.Print(strings.Repeat("\t", indent))
	fmt.Print("{row}:\n")
	dumpTableCells(cxt, row.Cells, indent+1)
}

func snippet(str string) string {
	v := []rune(str)
	if 25 < len(v) {
		str = string(v[:25])
	}
	return strings.ReplaceAll(str, "\n", "")
}
