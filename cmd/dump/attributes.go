package dump

import (
	"fmt"
	"strings"

	"github.com/hasty/adoc/elements"
)

func dumpAttributes(attributes []elements.Attribute, indent int) {
	if len(attributes) == 0 {
		return
	}
	fmt.Print(strings.Repeat("\t", indent))
	fmt.Print("{attr:\n")
	for _, a := range attributes {
		fmt.Print(strings.Repeat("\t", indent+1))
		switch a := a.(type) {
		case *elements.NamedAttribute:
			fmt.Printf(" %s=", a.Name)
		}
		switch v := a.Value().(type) {
		case *elements.String:
			fmt.Print(v.Value)
		case string:
			fmt.Print(v)
		/*case elements.Options:
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
		case elements.String:
			fmt.Print(v)
		case string:
			fmt.Print(v)
		case *elements.TableColumn:
			fmt.Printf("{col:\n")
			fmt.Print(strings.Repeat("\t", indent+2))
			fmt.Printf("multiplier: %d\n", v.Multiplier.Get())
			fmt.Print(strings.Repeat("\t", indent+2))
			fmt.Printf("halign: %s\n", v.HorizontalAlign.Get().String())
			fmt.Print(strings.Repeat("\t", indent+2))
			fmt.Printf("valign: %s\n", v.VerticalAlign.Get().String())
			fmt.Print(strings.Repeat("\t", indent+2))
			fmt.Printf("style: %s\n", v.Style.Get().String())
			fmt.Print(strings.Repeat("\t", indent+2))
			fmt.Printf("width: %s\n", v.Width.Get().String())
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
