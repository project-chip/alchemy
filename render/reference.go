package render

import (
	"fmt"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/matterfmt/output"
)

func renderInternalCrossReference(cxt *output.Context, cf *types.InternalCrossReference) {
	//fmt.Printf("icf ID type: %T -> %v\n", cf.ID, cf.ID)
	//fmt.Printf("icf Label type: %T -> %v\n", cf.Label, cf.Label)

	switch el := cf.ID.(type) {
	case string:
		if strings.HasPrefix(el, "_") {
			ref, ok := cxt.Doc.Base.ElementReferences[el]
			if ok {
				//fmt.Printf("id %s => icf ref type: %T\n", el, ref)
				switch idref := ref.(type) {
				case []interface{}:
					for _, i := range idref {
						//fmt.Printf("\ticf ref child type: %T\n", i)
						switch iv := i.(type) {
						case *types.StringElement:
							t, _ := iv.RawText()

							//fmt.Printf("val: %s\n", t)
							el = t
						}
					}
				}
			}
		}
		cxt.WriteString("<<")
		cxt.WriteString(el)
		if label, ok := cf.Label.(string); ok {
			cxt.WriteString(",")
			cxt.WriteString(label)
		}
		cxt.WriteString(">>")
	default:
		panic(fmt.Errorf("unknown internal cross reference ID type: %T", el))
	}
}
