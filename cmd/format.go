package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/hasty/matterfmt/ascii"
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

func format(errCxt context.Context, file string) (string, error) {
	doc, err := ascii.Open(file)
	if err != nil {
		return "", err
	}
	return adoc.Render(errCxt, doc)
}
