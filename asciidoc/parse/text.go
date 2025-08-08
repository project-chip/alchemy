package parse

import (
	"fmt"
	"log/slog"
	"strings"
	"unicode"

	"github.com/project-chip/alchemy/asciidoc"
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
	case nil:
	default:
		sb.WriteString(fmt.Sprintf("unknown string type: %T", e))
	}
}

func mergeStrings[T any](els []T) (out asciidoc.Elements) {
	var s strings.Builder
	out = mergeStringsInternal(els, &s)
	if s.Len() > 0 {
		out = append(out, asciidoc.NewString(s.String()))
	}
	return
}

func mergeStringsInternal[T any](els []T, s *strings.Builder) (out asciidoc.Elements) {
	for _, e := range els {
		switch e := any(e).(type) {
		case string:
			s.WriteString(e)
		case []byte:
			s.WriteString(string(e))
		case *asciidoc.String:
			s.WriteString(e.Value)
		case asciidoc.Element:
			if s.Len() > 0 {
				out = append(out, asciidoc.NewString(s.String()))
				s.Reset()
			}
			out = append(out, e)
		case []any:
			out = append(out, mergeStringsInternal(e, s)...)
		case nil:
		default:
			slog.Warn("unexpected type in string merge", slog.String("type", fmt.Sprintf("%T", e)))
		}
	}
	return
}

func trim(in asciidoc.Elements) (out asciidoc.Elements) {
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
		out = asciidoc.Elements{asciidoc.NewString("")}
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
		out = asciidoc.Elements{asciidoc.NewString("")}
	}
	return
}

func trimRight(in asciidoc.Elements) (out asciidoc.Elements) {
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
		out = asciidoc.Elements{asciidoc.NewString("")}
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
