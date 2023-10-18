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

func DiscoBall(cxt context.Context, filepaths []string, dryRun bool, serial bool, options ...disco.Option) error {
	if serial {
		return discoBallSerial(cxt, filepaths, dryRun, options...)
	}
	return discoBallParallel(cxt, filepaths, dryRun, options...)
}

func discoBallSerial(cxt context.Context, filepaths []string, dryRun bool, options ...disco.Option) error {
	files, err := getFilePaths(filepaths)
	if err != nil {
		return err
	}
	for i, file := range files {
		result, err := discoBall(cxt, file, options...)
		if err != nil {
			return err
		}
		if !dryRun {
			err = os.WriteFile(file, []byte(result), os.ModeAppend)
			if err != nil {
				return err
			}
		}
		fmt.Fprintf(os.Stderr, "Formatted %s (%d of %d)...\n", file, (i + 1), len(files))
	}
	return nil
}

func discoBallParallel(cxt context.Context, filepaths []string, dryRun bool, options ...disco.Option) error {
	files, err := getFilePaths(filepaths)
	if err != nil {
		return err
	}
	g, errCxt := errgroup.WithContext(cxt)
	var complete int32
	for i, f := range files {
		func(file string, index int) {
			g.Go(func() error {
				result, err := discoBall(errCxt, file, options...)
				if err != nil {
					return err
				}
				if !dryRun {
					err = os.WriteFile(file, []byte(result), os.ModeAppend)
					if err != nil {
						return err
					}
				}
				done := atomic.AddInt32(&complete, 1)
				fmt.Fprintf(os.Stderr, "Disco-balled %s (%d of %d)...\n", file, done, len(files))
				return nil
			})
		}(f, i)

	}
	return g.Wait()
}

func discoBall(cxt context.Context, file string, options ...disco.Option) (string, error) {
	out, err := getOutputContext(cxt, file)
	if err != nil {
		return "", err
	}
	b := disco.NewBall(out.Doc)
	for _, option := range options {
		option(b)
	}
	err = b.Run(cxt)
	if err != nil {
		slog.Error("error disco balling", "file", file, "error", err)
		return "", nil
	}
	return render.Render(cxt, out.Doc)
}
