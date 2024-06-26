package dump

import (
	"fmt"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/matter/spec"
)

/*
	func dumpTOC(tocs []*asciidoc.ToCSection, indent int) {
		for _, toc := range tocs {
			fmt.Print(strings.Repeat("\t", indent))
			fmt.Printf("{toc %d} %s.%s\n", toc.Level, toc.Number, toc.Title)
			if len(toc.Children) > 0 {
				dumpTOC(toc.Children, indent+1)
			}
		}

}
*/
func dumpLocation(doc *spec.Doc, l asciidoc.URL, indent int) {
	fmt.Printf("%s", l.Scheme)
	dumpElements(doc, l.Path, indent)
}

func snippet(str string) string {
	v := []rune(str)
	if 42 < len(v) {
		str = string(v[:20]) + "…" + string(v[len(v)-20:])
	}
	return strings.ReplaceAll(str, "\n", "\\n")
}
