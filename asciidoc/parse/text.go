package parse

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/hasty/alchemy/asciidoc"
)

func toString(e any) string {
	var sb strings.Builder
	toStringBuilder(e, &sb)
	return sb.String()
}

func toStringBuilder(e any, sb *strings.Builder) {
	switch e := e.(type) {
	case []any:
		for _, c := range e {
			toStringBuilder(c, sb)
		}
	case string:
		sb.WriteString(e)
	case []byte:
		sb.WriteString(string(e))
	case *asciidoc.String:
		sb.WriteString(e.Value)
	default:
		sb.WriteString(fmt.Sprintf("unknown string type: %T", e))
	}
}

func mergeStrings[T any](els []T) (out asciidoc.Set) {
	var s strings.Builder
	for _, e := range els {
		switch e := any(e).(type) {
		case *asciidoc.String:
			s.WriteString(e.Value)
		case asciidoc.Element:
			if s.Len() > 0 {
				out = append(out, asciidoc.NewString(s.String()))
				s.Reset()
			}
			out = append(out, e)
		default:
			fmt.Printf("unexpected type in string merge: %T\n", e)
		}
	}
	if s.Len() > 0 {
		out = append(out, asciidoc.NewString(s.String()))
	}
	return
}

func trim(in asciidoc.Set) (out asciidoc.Set) {
	var first int = -1
	var last = len(in)
	for i, e := range in {
		s, ok := e.(*asciidoc.String)
		if !ok {
			first = i
			break
		}
		if !isWhitespace(s.Value) {
			s.Value = strings.TrimLeft(s.Value, " \t")
			first = i
			break
		}

	}
	if first == -1 { // Every element is a whitespace string
		out = asciidoc.Set{asciidoc.NewString("")}
		return
	}

	for i := last - 1; i >= first; i-- {
		e := in[i]
		s, ok := e.(*asciidoc.String)
		if !ok {
			last = i + 1
			break
		}
		if !isWhitespace(s.Value) {
			s.Value = strings.TrimRight(s.Value, " \t")
			last = i + 1
			break
		}
	}
	out = in[first:last]
	if len(out) == 0 {
		out = asciidoc.Set{asciidoc.NewString("")}
	}
	return
}

func trimRight(in asciidoc.Set) (out asciidoc.Set) {
	var last = len(in)

	for i := last - 1; i >= 0; i-- {
		e := in[i]
		s, ok := e.(*asciidoc.String)
		if !ok {
			last = i + 1
			break
		}
		if !isWhitespace(s.Value) {
			s.Value = strings.TrimRight(s.Value, " \t")
			last = i + 1
			break
		}
	}
	if last == 0 {
		out = asciidoc.Set{asciidoc.NewString("")}
		return
	}
	out = in[0:last]
	return
}

func isWhitespace(s string) bool {
	for _, r := range s {
		if !unicode.IsSpace(r) {
			return false
		}
	}
	return true
}
