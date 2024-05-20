package parse

import (
	"fmt"

	"github.com/hasty/adoc/elements"
)

func Position(el elements.Element) string {
	p, ok := el.(elements.HasPosition)
	if !ok {
		return "unknown"
	}
	line, col, _ := p.Position()
	return fmt.Sprintf("%d:%d", line, col)
}
