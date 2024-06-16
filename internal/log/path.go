package log

import (
	"log/slog"
	"strconv"
	"strings"

	"github.com/hasty/alchemy/matter"
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
