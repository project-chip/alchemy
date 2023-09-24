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

	switch el := cf.OriginalID.(type) {
	case string:
		if strings.HasPrefix(el, "_") {
			return
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
