//go:generate go run generate.go
//go:build generate

package main

import (
	"log/slog"
	"os"

	"github.com/project-chip/alchemy/internal/generate"
)

func main() {
	slog.Info("Generating PICS parser...")
	err := generate.Parser("grammar/grammar.json", false, nil)
	if err != nil {
		slog.Error("error generating PICS parser", slog.Any("error", err))
		os.Exit(1)
		return
	}
	os.Exit(0)
	return
}
