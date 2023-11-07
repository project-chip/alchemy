package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/parse"
)

type dumper struct {
	asciiParser

	dumpAscii bool
	dumpJSON  bool
}

func Dump(cxt context.Context, filepaths []string, options ...Option) error {
	d := &dumper{}
	for _, opt := range options {
		err := opt(d)
		if err != nil {
			return err
		}
	}
	return d.run(cxt, filepaths)
}

func (d *dumper) run(cxt context.Context, filepaths []string) error {
	files, err := getFilePaths(filepaths)
	if err != nil {
		return err
	}
	for i, f := range files {
		if len(files) > 0 {
			fmt.Fprintf(os.Stderr, "Dumping %s (%d of %d)...\n", f, (i + 1), len(files))
		}
		doc, err := ascii.Open(f, d.settings...)
		docType, err := doc.DocType()
		if err != nil {
			return err
		}
		if d.dumpAscii {
			for _, top := range parse.Skim[*ascii.Section](doc.Elements) {
				ascii.AssignSectionTypes(docType, top)
			}
			dumpElements(doc, doc.Elements, 0)
		} else if d.dumpJSON {
			models, err := doc.ToModel()
			if err != nil {
				return err
			}
			encoder := json.NewEncoder(os.Stdout)
			//encoder.SetIndent("", "\t")
			return encoder.Encode(models)
		} else {
			dumpElements(doc, doc.Base.Elements, 0)
		}
	}
	return nil
}

func dumpElements(doc *ascii.Doc, elements []interface{}, indent int) {

	for _, e := range elements {
		fmt.Print(strings.Repeat("\t", indent))
		as, ok := e.(*ascii.Section)
		if ok {
			fmt.Printf("{SEC %d (%s)}:\n", as.Base.Level, matter.SectionTypeString(as.SecType))

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
			dumpElements(doc, []interface{}{el.Value}, indent+1)
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
				dumpElements(doc, []interface{}{a}, indent+1)
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
			dumpElements(doc, []interface{}{el.Element}, indent+1)
		case *types.PredefinedAttribute:
			fmt.Printf("{predef %s}", el.Name)
		default:
			fmt.Printf("unknown element type: %T\n", el)
		}
	}
}

func dumpAttributes(attributes types.Attributes, indent int) {
	if len(attributes) == 0 {
		return
	}
	fmt.Print(strings.Repeat("\t", indent))
	fmt.Print("{attr:\n")
	for key, val := range attributes {
		fmt.Print(strings.Repeat("\t", indent+1))
		fmt.Printf(" %s=", key)
		switch v := val.(type) {
		case *types.StringElement:
			fmt.Print(v.Content)
		case string:
			fmt.Print(v)
		case types.Options:
			dumpAttributeVals(v, indent+1)
		case []interface{}:
			dumpAttributeVals(v, indent+1)
		default:
			fmt.Printf("unknown type: %T", val)
		}
		fmt.Print("\n")
	}
	fmt.Print(strings.Repeat("\t", indent))
	fmt.Print("}\n")
}

func dumpAttributeVals(attributes []interface{}, indent int) {
	fmt.Print("{\n")
	for _, val := range attributes {
		fmt.Print(strings.Repeat("\t", indent+1))
		switch v := val.(type) {
		case *types.StringElement:
			fmt.Print(v.Content)
		case string:
			fmt.Print(v)
		case *types.TableColumn:
			fmt.Printf("{col:\n")
			fmt.Print(strings.Repeat("\t", indent+2))
			fmt.Printf("multiplier: %d\n", v.Multiplier)
			fmt.Print(strings.Repeat("\t", indent+2))
			fmt.Printf("halign: %s\n", v.HAlign)
			fmt.Print(strings.Repeat("\t", indent+2))
			fmt.Printf("valign: %s\n", v.VAlign)
			fmt.Print(strings.Repeat("\t", indent+2))
			fmt.Printf("weight: %d\n", v.Weight)
			fmt.Print(strings.Repeat("\t", indent+2))
			fmt.Printf("style: %s\n", v.Style)
			fmt.Print(strings.Repeat("\t", indent+2))
			fmt.Printf("autowidth: %v\n", v.Autowidth)
			fmt.Print(strings.Repeat("\t", indent+2))
			fmt.Printf("width: %s\n", v.Width)
			fmt.Print(strings.Repeat("\t", indent+1))
			fmt.Printf("}\n")
		default:
			fmt.Printf("unknown type: %T", val)
		}
		fmt.Print(",\n")
	}
	fmt.Print(strings.Repeat("\t", indent))
	fmt.Print("}\n")
}

func dumpTOC(tocs []*types.ToCSection, indent int) {
	for _, toc := range tocs {
		fmt.Print(strings.Repeat("\t", indent))
		fmt.Printf("{toc %d} %s.%s\n", toc.Level, toc.Number, toc.Title)
		if len(toc.Children) > 0 {
			dumpTOC(toc.Children, indent+1)
		}
	}

}

func dumpTable(doc *ascii.Doc, tbl *types.Table, indent int) {
	fmt.Print(strings.Repeat("\t", indent))
	fmt.Print("{head}:\n")
	dumpTableCells(doc, tbl.Header.Cells, indent+1)
	fmt.Print(strings.Repeat("\t", indent))
	fmt.Print("{body}:\n")
	for _, row := range tbl.Rows {
		dumpTableRow(doc, row, indent+1)
	}
}

func dumpTableRow(doc *ascii.Doc, row *types.TableRow, indent int) {
	fmt.Print(strings.Repeat("\t", indent))
	fmt.Print("{row}:\n")
	dumpTableCells(doc, row.Cells, indent+1)
}

func dumpTableCells(doc *ascii.Doc, cells []*types.TableCell, indent int) {
	for _, c := range cells {
		fmt.Print(strings.Repeat("\t", indent))
		if c.Blank {
			fmt.Print("{cellblank}:\n")
		} else {
			fmt.Print("{cell}:\n")
			if c.Formatter != nil {
				fmt.Print(strings.Repeat("\t", indent+1))
				fmt.Printf("{format: %s (cell %d row %d)}\n", c.Formatter.Content, c.Formatter.ColumnSpan, c.Formatter.RowSpan)
			}
			dumpElements(doc, c.Elements, indent+1)

		}
	}

}

func dumpLocation(l *types.Location) {
	if l != nil {
		fmt.Printf("%s %s}", l.Scheme, l.Path.(string))
	} else {
		fmt.Printf("missing location")
	}
}

func snippet(str string) string {
	v := []rune(str)
	if 42 < len(v) {
		str = string(v[:20]) + "…" + string(v[len(v)-20:])
	}
	return strings.ReplaceAll(str, "\n", "")
}
