package parse

import (
	"fmt"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
)

func join[T any](els []T) (out asciidoc.Set) {
	var s strings.Builder
	//fmt.Printf("joining %d elements\n", len(els))
	for _, e := range els {
		//fmt.Printf("join; %T %v\n", e, e)
		switch e := any(e).(type) {
		case string:
			s.WriteString(e)
		case []byte:
			s.WriteString(string(e))
		case *asciidoc.String:
			//fmt.Printf("writing string: %s\n", string(e))
			s.WriteString(e.Value)
		case *asciidoc.LineContinuation:
			//fmt.Printf("writing continuation\n")
			s.WriteString(" +\n")
		case asciidoc.Element:
			if s.Len() > 0 {
				//fmt.Printf("flushing string: %s\n", s.String())
				out = append(out, asciidoc.NewString(s.String()))
				s.Reset()
			}
			out = append(out, e)
		case []asciidoc.Element:
			out = append(out, e...)
		case asciidoc.Set:
			out = append(out, e...)
		case []any:
			out = append(out, join(e)...)
		case nil:
			continue
		default:
			fmt.Printf("unknown type in join: %T\n", e)
		}
	}
	if s.Len() > 0 {
		//fmt.Printf("flushing string: %s\n", s.String())
		out = append(out, asciidoc.NewString(s.String()))
	}
	return
}
