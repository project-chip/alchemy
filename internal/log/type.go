package log

import (
	"fmt"
	"log/slog"
)

func Type(name string, t any) slog.Attr {
	return slog.String(name, fmt.Sprintf("%T", t))
}
