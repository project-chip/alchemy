package parse

import (
	"fmt"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
)

func Dump(els ...asciidoc.Element) {
	dump(0, els...)
}

func dump(indent int, els ...asciidoc.Element) {
	for _, v := range els {
		fmt.Print(strings.Repeat("\t", indent))
		if hp, ok := v.(asciidoc.HasPosition); ok {
			l, c, o := hp.Position()
			fmt.Printf("(%d:%d:%d) ", l, c, o)
		}
		switch v := v.(type) {
		case *asciidoc.AttributeEntry:
			dumpAttributeEntry(v, indent)
		case *asciidoc.TableCell:
			fmt.Printf("[%s]", asciidoc.Describe(v))
			if v.Format != nil {
				fmt.Printf(" (format: %s)", v.Format.AsciiDocString())
			}
			fmt.Printf("\n")
			dump(indent+1, v.Children()...)
		case asciidoc.Element:
			fmt.Printf("[%s]", asciidoc.Describe(v))
			if hp, ok := v.(asciidoc.HasPosition); ok {
				line, column, offset := hp.Position()
				fmt.Printf(" (%d, %d, %d)", line, column, offset)
			}
			fmt.Printf("\n")
			if section, ok := v.(*asciidoc.Section); ok {
				fmt.Print(strings.Repeat("\t", indent+1))
				fmt.Printf("Title:\n")
				dump(indent+2, section.Title...)
			}
			dumpAttributes(v, indent+1)
			if es, ok := v.(asciidoc.ParentElement); ok {
				dump(indent+1, es.Children()...)
			}
			if es, ok := v.(asciidoc.HasChild); ok {
				dump(indent+1, es.Child())
			}
			if es, ok := v.(asciidoc.HasLines); ok {
				for _, s := range es.Lines() {
					fmt.Print(strings.Repeat("\t", indent+1))
					fmt.Printf("line: %s\n", s)
				}
			}
		default:
			fmt.Printf("%T\n", v)
			dumpAttributes(v, indent+1)

		}
	}
}

func dumpAttributeEntry(el *asciidoc.AttributeEntry, indent int) {
	fmt.Print(strings.Repeat("\t", indent))
	fmt.Printf("attribute definition %s:\n", el.Name)
	dump(indent+1, el.Children()...)

}

func dumpAttributes(el asciidoc.Element, indent int) {
	ae, ok := el.(asciidoc.Attributable)
	if !ok {
		return
	}
	attrs := ae.Attributes()
	if len(attrs) == 0 {
		return
	}
	fmt.Print(strings.Repeat("\t", indent))
	fmt.Print("attributes:\n")
	for _, a := range attrs {
		if a == nil {
			fmt.Print("nil attribute!\n")
			return
		}
		fmt.Print(strings.Repeat("\t", indent+1))
		fmt.Print("attr:")
		switch a := a.(type) {
		case *asciidoc.NamedAttribute:
			fmt.Printf(" (name: %s)", a.Name)
		case *asciidoc.AnchorAttribute:
			fmt.Printf(" (anchor:)\n")
			dump(indent+1, a.ID.Children()...)
			dump(indent+1, a.Label.Children()...)
			continue
		case *asciidoc.PositionalAttribute:
			fmt.Printf(" (position: %d)", a.Offset)
		case *asciidoc.TitleAttribute:
			fmt.Printf(" (title)")
		case *asciidoc.TableColumnsAttribute:
			fmt.Printf(" (cols)")
		default:
			fmt.Printf(" (unknown: %T)", a)
		}
		switch a.QuoteType() {
		case asciidoc.AttributeQuoteTypeDouble:
			fmt.Print(" double-quoted")
		case asciidoc.AttributeQuoteTypeSingle:
			fmt.Print(" single-quoted")
		}
		fmt.Print(" val: ")
		switch v := a.Value().(type) {
		case []asciidoc.Element:
			fmt.Println()
			dump(indent+2, v...)
		case asciidoc.Elements:
			fmt.Println()
			dump(indent+2, v...)
		case *asciidoc.String:
			fmt.Printf("\"%s\"\n", v.Value)
		case []*asciidoc.TableColumn:
			fmt.Println()
			for i, col := range v {
				fmt.Print(strings.Repeat("\t", indent+2))
				fmt.Printf("col %d:", i+1)
				if col.HorizontalAlign.IsSet {
					fmt.Printf("%s,", col.HorizontalAlign.Value.AsciiDocString())
				}
				if col.VerticalAlign.IsSet {
					fmt.Printf("%s,", col.VerticalAlign.Value.AsciiDocString())
				}
				if col.Style.IsSet {
					fmt.Printf("%s,", col.Style.Value.AsciiDocString())

				}
				fmt.Println()
			}
		default:
			fmt.Printf("unknown type: %T\n", v)
		}
	}

}
