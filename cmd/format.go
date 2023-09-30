package cmd

import (
	"context"
	"fmt"
	"os"
	"sync/atomic"

	"github.com/hasty/matterfmt/render"
	"golang.org/x/sync/errgroup"
)

func Format(cxt context.Context, filepath string, dryRun bool, serial bool) error {
	files, err := getFilePaths(filepath)
	if err != nil {
		return err
	}
	fmt.Fprintf(os.Stderr, "Formatting %d files...\n", len(files))
	if serial {
		return formatSerial(cxt, files, dryRun)
	}
	return formatParallel(cxt, files, dryRun)
}

func formatSerial(cxt context.Context, files []string, dryRun bool) error {
	for i, file := range files {
		fmt.Fprintf(os.Stderr, "Formatting %s (%d of %d)...\n", file, (i + 1), len(files))
		err := format(cxt, file, dryRun)
		if err != nil {
			return err
		}
	}
	return nil
}

func formatParallel(cxt context.Context, files []string, dryRun bool) error {
	var complete int32
	g, errCxt := errgroup.WithContext(cxt)
	for i, f := range files {
		func(file string, index int) {
			g.Go(func() error {

				err := format(errCxt, file, dryRun)
				if err != nil {
					return err
				}
				done := atomic.AddInt32(&complete, 1)
				fmt.Fprintf(os.Stderr, "Formatted %s (%d of %d)...\n", file, done, len(files))

				return nil
			})
		}(f, i)

	}
	return g.Wait()
}

func format(errCxt context.Context, file string, dryRun bool) error {
	out, err := getOutputContext(errCxt, file)
	if err != nil {
		return err
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
