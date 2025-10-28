package zapdiff

import (
	"fmt"

	"github.com/beevik/etree"
)

func getElementXPathSegment(e *etree.Element) (s string) {
	p := e.Parent()
	if p == nil {
		return e.Tag
	}

	idx := 0
	cnt := 0
	for _, sib := range p.ChildElements() {
		if sib.Tag == e.Tag {
			cnt++
			if sib == e {
				idx = cnt
				break
			}
		}
	}

	s = e.Tag
	if cnt > 1 {
		s += fmt.Sprintf("[%d]", idx)
	}
	return s
}
