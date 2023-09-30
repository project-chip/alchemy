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

func DiscoBall(cxt context.Context, filepath string, dryRun bool, serial bool) error {
	if serial {
		return discoBallSerial(cxt, filepath, dryRun)
	}
	return discoBallParallel(cxt, filepath, dryRun)
}

func discoBallSerial(cxt context.Context, filepath string, dryRun bool) error {
	files, err := getFilePaths(filepath)
	if err != nil {
		return err
	}
	for i, file := range files {
		err := discoBall(cxt, file, dryRun)
		if err != nil {
			return err
		}
		fmt.Fprintf(os.Stderr, "Formatted %s (%d of %d)...\n", file, (i + 1), len(files))
	}
	return nil
}

func discoBallParallel(cxt context.Context, filepath string, dryRun bool) error {
	files, err := getFilePaths(filepath)
	if err != nil {
		return err
	}
	g, errCxt := errgroup.WithContext(cxt)
	var complete int32
	for i, f := range files {
		func(file string, index int) {
			g.Go(func() error {
				err := discoBall(errCxt, file, dryRun)
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

func discoBall(errCxt context.Context, file string, dryRun bool) error {
	out, err := getOutputContext(errCxt, file)
	if err != nil {
		return err
	}
	err = disco.Ball(disco.NewContext(errCxt), out.Doc)
	if err != nil {
		slog.Error("error disco balling", "file", file, "error", err)
		return nil
	}
	result, err := render.Render(errCxt, out.Doc)
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
