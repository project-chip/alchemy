package dump

import (
	"fmt"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
)

func dumpTOC(tocs []*types.ToCSection, indent int) {
	for _, toc := range tocs {
		fmt.Print(strings.Repeat("\t", indent))
		fmt.Printf("{toc %d} %s.%s\n", toc.Level, toc.Number, toc.Title)
		if len(toc.Children) > 0 {
			dumpTOC(toc.Children, indent+1)
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
	return strings.ReplaceAll(str, "\n", "\\n")
}
