//go:generate go run -tags generate,debugParser generate.go  generate_debug.go
//go:build generate && debugParser

package main

import (
	"log/slog"
	"os"

	"github.com/project-chip/alchemy/internal/generate"
)

func main() {
	slog.Info("Generating Asciidoc parser...")
	err := generate.Parser("grammar/grammar.json", true, parserPatch)
	if err != nil {
		slog.Error("error generating asciidoc parser", slog.Any("error", err))
		os.Exit(1)
		return
	}
	os.Exit(0)
	return
}
