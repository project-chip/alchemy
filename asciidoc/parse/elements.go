package parse

import (
	"fmt"
	"io"
	"log/slog"

	"github.com/project-chip/alchemy/asciidoc"
)

func Elements(path string, reader io.Reader, opts ...Option) (elements asciidoc.Elements, err error) {
	var vals any
	vals, err = ParseReader(path, reader, opts...)
	if err != nil {
		slog.Error("error parsing file", slog.String("path", path), slog.Any("error", err))
		return nil, err
	}
	var ok bool
	elements, ok = vals.(asciidoc.Elements)
	if !ok {
		return nil, fmt.Errorf("unexpected type in UnifiedParse: %T", vals)
	}
	return
}
