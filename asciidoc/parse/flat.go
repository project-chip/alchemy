package parse

import (
	"fmt"

	"github.com/hasty/alchemy/asciidoc"
)

func cast[F any, T any](in []F) (out []T) {
	out = make([]T, 0, len(in))
	for _, f := range in {
		if t, ok := any(f).(T); ok {
			out = append(out, t)
		} else {
			fmt.Printf("unexpected value in cast: %T\n", f)
		}
	}
	return
}

func flat(els []any) (out asciidoc.Set) {
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

func flatAppend(e any, list asciidoc.Set) asciidoc.Set {
	switch e := e.(type) {
	case asciidoc.Element:
		list = append(list, e)
	case []asciidoc.Element:
		list = append(list, e...)
	case asciidoc.Set:
		list = append(list, e...)
	case string:
		list = append(list, asciidoc.NewString(e))
	case []byte:
		list = append(list, asciidoc.NewString(string(e)))
	case []any:
		list = append(list, flat(e)...)
	case nil:

	default:
		fmt.Printf("unknown type in flat: %T\n", e)
	}
	return list
}
