package parse

import (
	"fmt"

	"github.com/hasty/adoc/asciidoc"
)

func Position(el asciidoc.Element) string {
	p, ok := el.(asciidoc.HasPosition)
	if !ok {
		return "unknown"
	}
	line, col, _ := p.Position()
	return fmt.Sprintf("%d:%d", line, col)
}
