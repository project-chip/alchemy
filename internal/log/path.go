package log

import (
	"fmt"
	"log/slog"
	"strconv"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
)

type Source interface {
	Origin() (path string, line int)
}

func Origin(source Source) string {
	var path strings.Builder
	p, l := source.Origin()
	path.WriteString(p)
	if l >= 0 {
		path.WriteRune(':')
		path.WriteString(strconv.Itoa(l))
	}
	return path.String()
}

func Path(name string, source Source) slog.Attr {
	if source == nil {
		return slog.String(name, "unknown")
	}
	return slog.String(name, Origin(source))
}

func Element(name string, path fmt.Stringer, element asciidoc.Element) slog.Attr {
	var arg strings.Builder
	arg.WriteString(path.String())
	if hp, ok := element.(asciidoc.HasPosition); ok {
		l, _, _ := hp.Position()
		if l >= 0 {
			arg.WriteRune(':')
			arg.WriteString(strconv.Itoa(l))
		}
	}
	return slog.String(name, arg.String())
}

func Elements(name string, path fmt.Stringer, elements asciidoc.Set) slog.Attr {
	var arg strings.Builder
	arg.WriteString(path.String())
	for _, element := range elements {
		if hp, ok := element.(asciidoc.HasPosition); ok {
			l, _, _ := hp.Position()
			if l >= 0 {
				arg.WriteRune(':')
				arg.WriteString(strconv.Itoa(l))
				break
			}
		}

	}
	return slog.String(name, arg.String())
}
