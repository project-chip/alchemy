package parse

import (
	"fmt"
	"log/slog"

	"github.com/project-chip/alchemy/asciidoc"
)

func cast[F any, T any](in []F) (out []T) {
	out = make([]T, 0, len(in))
	for _, f := range in {
		if t, ok := any(f).(T); ok {
			out = append(out, t)
		} else {
			slog.Warn("unexpected type in cast", slog.String("type", fmt.Sprintf("%T", f)))
		}
	}
	return
}

func flat(els []any) (out asciidoc.Elements) {
	for _, e := range els {
		switch e := e.(type) {
		case []any:
			out = append(out, flat(e)...)
		default:
			out = flatAppend(e, out)
		}
	}
	return
}

func flatAppend(e any, list asciidoc.Elements) asciidoc.Elements {
	switch e := e.(type) {
	case asciidoc.Element:
		list = append(list, e)
	case []asciidoc.Element:
		list = append(list, e...)
	case asciidoc.Elements:
		list = append(list, e...)
	case string:
		list = append(list, asciidoc.NewString(e))
	case []byte:
		list = append(list, asciidoc.NewString(string(e)))
	case []any:
		list = append(list, flat(e)...)
	case nil:

	default:
		slog.Warn("unexpected type in flat", slog.String("type", fmt.Sprintf("%T", e)))
	}
	return list
}
