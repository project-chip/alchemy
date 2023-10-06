package cmd

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"sync/atomic"

	"github.com/hasty/matterfmt/disco"
	"github.com/hasty/matterfmt/render"
	"golang.org/x/sync/errgroup"
)

func DiscoBall(cxt context.Context, filepaths []string, dryRun bool, serial bool, linkAttributes bool) error {
	if serial {
		return discoBallSerial(cxt, filepaths, dryRun, linkAttributes)
	}
	return discoBallParallel(cxt, filepaths, dryRun, linkAttributes)
}

func discoBallSerial(cxt context.Context, filepaths []string, dryRun bool, linkAttributes bool) error {
	files, err := getFilePaths(filepaths)
	if err != nil {
		return err
	}
	for i, file := range files {
		err := discoBall(cxt, file, dryRun, linkAttributes)
		if err != nil {
			return err
		}
		fmt.Fprintf(os.Stderr, "Formatted %s (%d of %d)...\n", file, (i + 1), len(files))
	}
	return nil
}

func discoBallParallel(cxt context.Context, filepaths []string, dryRun bool, linkAttributes bool) error {
	files, err := getFilePaths(filepaths)
	if err != nil {
		return err
	}
	g, errCxt := errgroup.WithContext(cxt)
	var complete int32
	for i, f := range files {
		func(file string, index int) {
			g.Go(func() error {
				err := discoBall(errCxt, file, dryRun, linkAttributes)
				if err != nil {
					return err
				}
				done := atomic.AddInt32(&complete, 1)
				fmt.Fprintf(os.Stderr, "Disco-balled %s (%d of %d)...\n", file, done, len(files))
				return nil
			})
		}(f, i)

	}
	return g.Wait()
}

func discoBall(cxt context.Context, file string, dryRun bool, linkAttributes bool) error {
	out, err := getOutputContext(cxt, file)
	if err != nil {
		return err
	}
	b := disco.NewBall(out.Doc)
	b.ShouldLinkAttributes = linkAttributes
	err = b.Run(cxt)
	if err != nil {
		slog.Error("error disco balling", "file", file, "error", err)
		return nil
	}
	result, err := render.Render(cxt, out.Doc)
	if err != nil {
		return err
	}
	if dryRun {
		return nil
	}
	err = os.WriteFile(file, []byte(result), os.ModeAppend)
	if err != nil {
		return err
	}
	return nil
}
