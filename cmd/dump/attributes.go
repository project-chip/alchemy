package dump

import (
	"fmt"
	"strings"

	"github.com/hasty/adoc/elements"
)

func dumpAttributes(attributes elements.AttributeList, indent int) {
	if len(attributes) == 0 {
		return
	}
	fmt.Print(strings.Repeat("\t", indent))
	fmt.Print("{attr:\n")
	for key, val := range attributes {
		fmt.Print(strings.Repeat("\t", indent+1))
		fmt.Printf(" %s=", key)
		switch v := val.(type) {
		case *elements.String:
			fmt.Print(v.Content)
		case string:
			fmt.Print(v)
		case elements.Options:
			dumpAttributeVals(v, indent+1)
		case []any:
			dumpAttributeVals(v, indent+1)
		default:
			fmt.Printf("unknown type: %T", val)
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
		case *elements.String:
			fmt.Print(v.Content)
		case string:
			fmt.Print(v)
		case *elements.TableColumn:
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
