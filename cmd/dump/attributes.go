package dump

import (
	"fmt"
	"strings"

	"github.com/hasty/adoc/asciidoc"
)

func dumpAttributes(attributes []asciidoc.Attribute, indent int) {
	if len(attributes) == 0 {
		return
	}
	fmt.Print(strings.Repeat("\t", indent))
	fmt.Print("{attr:\n")
	for _, a := range attributes {
		fmt.Print(strings.Repeat("\t", indent+1))
		switch a := a.(type) {
		case *asciidoc.NamedAttribute:
			fmt.Printf(" %s=", a.Name)
		}
		switch v := a.Value().(type) {
		case *asciidoc.String:
			fmt.Print(v.Value)
		case string:
			fmt.Print(v)
		/*case asciidoc.Options:
		dumpAttributeVals(v, indent+1)*/
		case []any:
			dumpAttributeVals(v, indent+1)
		default:
			fmt.Printf("unknown type: %T", v)
		}
		fmt.Print("\n")
	}
	fmt.Print(strings.Repeat("\t", indent))
	fmt.Print("}\n")
}

func dumpAttributeVals(attributes []any, indent int) {
	fmt.Print("{\n")
	for _, val := range attributes {
		fmt.Print(strings.Repeat("\t", indent+1))
		switch v := val.(type) {
		case asciidoc.String:
			fmt.Print(v)
		case string:
			fmt.Print(v)
		case *asciidoc.TableColumn:
			fmt.Printf("{col:\n")
			fmt.Print(strings.Repeat("\t", indent+2))
			fmt.Printf("multiplier: %d\n", v.Multiplier.Value)
			fmt.Print(strings.Repeat("\t", indent+2))
			fmt.Printf("halign: %s\n", v.HorizontalAlign.Value.String())
			fmt.Print(strings.Repeat("\t", indent+2))
			fmt.Printf("valign: %s\n", v.VerticalAlign.Value.String())
			fmt.Print(strings.Repeat("\t", indent+2))
			fmt.Printf("style: %s\n", v.Style.Value.String())
			fmt.Print(strings.Repeat("\t", indent+2))
			fmt.Printf("width: %s\n", v.Width.Value.String())
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
