package parse

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
)

func join[T any](els []T) (out asciidoc.Set) {
	var s strings.Builder
	for _, e := range els {
		switch e := any(e).(type) {
		case string:
			s.WriteString(e)
		case []byte:
			s.WriteString(string(e))
		case *asciidoc.String:
			s.WriteString(e.Value)
		case *asciidoc.LineContinuation:
			s.WriteString(" +\n")
		case asciidoc.Element:
			if s.Len() > 0 {
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
			slog.Warn("unexpected type in join", slog.String("type", fmt.Sprintf("%T", e)))
		}
	}
	if s.Len() > 0 {
		out = append(out, asciidoc.NewString(s.String()))
	}
	return
}
