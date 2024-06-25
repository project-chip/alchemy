package log

import (
	"log/slog"
	"strconv"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/matter"
)

func Path(name string, source matter.Source) slog.Attr {
	if source == nil {
		return slog.String(name, "unknown")
	}
	var path strings.Builder
	p, l := source.Origin()
	path.WriteString(p)
	if l >= 0 {
		path.WriteRune(':')
		path.WriteString(strconv.Itoa(l))
	}
	return slog.String(name, path.String())
}

func Element(name string, path string, element asciidoc.Element) slog.Attr {
	var arg strings.Builder
	arg.WriteString(path)
	if hp, ok := element.(asciidoc.HasPosition); ok {
		l, _, _ := hp.Position()
		if l >= 0 {
			arg.WriteRune(':')
			arg.WriteString(strconv.Itoa(l))
		}
	}
	return slog.String(name, arg.String())
}
