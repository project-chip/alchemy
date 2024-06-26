//go:generate go run -tags generate generate.go generate_release.go
//go:build generate && !debugParser

package main

import (
	"log/slog"
	"os"

	"github.com/project-chip/alchemy/internal/generate"
)

var debugParser = false

func main() {
	slog.Info("Generating Asciidoc parser...")
	err := generate.Parser("grammar/grammar.json", debugParser, parserPatch)
	if err != nil {
		slog.Error("error generating asciidoc parser", slog.Any("error", err))
		os.Exit(1)
		return
	}
	os.Exit(0)
	return
}
