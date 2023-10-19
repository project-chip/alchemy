package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/hasty/matterfmt/render/adoc"
)

func Format(cxt context.Context, filepaths []string, dryRun bool, serial bool) error {

	return processFiles(cxt, filepaths, serial, dryRun, func(cxt context.Context, file string, index int, total int) (result string, outPath string, err error) {
		outPath = file
		result, err = format(cxt, file)
		if err != nil {
			return
		}
		fmt.Fprintf(os.Stderr, "Formatted %s (%d of %d)...\n", file, index, total)
		return
	})

}

/*
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
}*/

func format(errCxt context.Context, file string) (string, error) {
	out, err := getOutputContext(errCxt, file)
	if err != nil {
		return "", err
	}
	return adoc.Render(errCxt, out.Doc)
}
