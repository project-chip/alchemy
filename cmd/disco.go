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

type discoBall struct {
	processor

	options []disco.Option
}

func DiscoBall(cxt context.Context, filepaths []string, options ...Option) error {
	db := &discoBall{}
	for _, opt := range options {
		err := opt(db)
		if err != nil {
			return err
		}
	}
	return db.saveFiles(cxt, filepaths, func(cxt context.Context, file string, index int, total int) (result string, outPath string, err error) {
		outPath = file
		result, err = db.run(cxt, file)
		if err != nil {
			return
		}
		fmt.Fprintf(os.Stderr, "Disco-balled %s (%d of %d)...\n", file, index, total)
		return
	})
}

func (db *discoBall) run(cxt context.Context, file string) (string, error) {
	doc, err := ascii.Open(file)
	if err != nil {
		return "", err
	}
	b := disco.NewBall(doc)
	for _, option := range db.options {
		option(b)
	}
	err = b.Run(cxt)
	if err != nil {
		slog.Error("error disco balling", "file", file, "error", err)
		return "", nil
	}
	return adoc.Render(cxt, doc)
}
