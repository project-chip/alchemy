package cmd

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/hasty/matterfmt/ascii"
	"github.com/hasty/matterfmt/disco"
	"github.com/hasty/matterfmt/render/adoc"
)

func DiscoBall(cxt context.Context, filepaths []string, dryRun bool, serial bool, options ...disco.Option) error {
	return processFiles(cxt, filepaths, serial, dryRun, func(cxt context.Context, file string, index int, total int) (result string, outPath string, err error) {
		outPath = file
		result, err = discoBall(cxt, file, options...)
		if err != nil {
			return
		}
		fmt.Fprintf(os.Stderr, "Disco-balled %s (%d of %d)...\n", file, index, total)
		return
	})
}

func discoBall(cxt context.Context, file string, options ...disco.Option) (string, error) {
	doc, err := ascii.Open(file)
	if err != nil {
		return "", err
	}
	b := disco.NewBall(doc)
	for _, option := range options {
		option(b)
	}
	err = b.Run(cxt)
	if err != nil {
		slog.Error("error disco balling", "file", file, "error", err)
		return "", nil
	}
	return adoc.Render(cxt, doc)
}
