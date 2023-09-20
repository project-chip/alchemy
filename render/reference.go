package render

import (
	"fmt"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/matterfmt/output"
)

func renderInternalCrossReference(cxt *output.Context, cf *types.InternalCrossReference) {
	fmt.Printf("icf ID type: %T -> %v\n", cf.ID, cf.ID)
	fmt.Printf("icf Label type: %T -> %v\n", cf.Label, cf.Label)

	switch el := cf.ID.(type) {
	case string:
		id := el
		//fmt.Printf("id %s\n", id)
		ref, ok := cxt.Doc.Base.ElementReferences[id]
		if ok {
			fmt.Printf("id %s => icf ref type: %T\n", id, ref)
			switch idref := ref.(type) {
			case []interface{}:
				for _, i := range idref {
					fmt.Printf("\ticf ref child type: %T\n", i)
					switch iv := i.(type) {
					case *types.StringElement:
						t, _ := iv.RawText()

						//fmt.Printf("val: %s\n", t)
						el = t
					}
				}
			}
		}
		cxt.WriteString("<<")
		cxt.WriteString(el)
		cxt.WriteString(">>")
	default:
		fmt.Printf("unknown internal cross reference ID type: %T\n", el)
	}
}
