package render

import (
	"fmt"
	"strings"
)

func renderInternalCrossReference(cxt *Context, cf *elements.InternalCrossReference) (err error) {
	switch el := cf.OriginalID.(type) {
	case string:
		if strings.HasPrefix(el, "_") {
			return
		}
		cxt.WriteString("<<")
		cxt.WriteString(el)
		if label, ok := cf.Label.(string); ok && len(label) > 0 {
			cxt.WriteString(",")
			cxt.WriteString(label)
		}
		cxt.WriteString(">>")
	case nil:
	default:
		err = fmt.Errorf("unknown internal cross reference ID type: %T", el)
	}
	return
}
