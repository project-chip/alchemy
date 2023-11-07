package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/render/adoc"
)

type formatter struct {
	processor
}

func Format(cxt context.Context, filepaths []string, options ...Option) error {
	f := &formatter{}
	for _, o := range options {
		o(f)
	}
	return f.saveFiles(cxt, filepaths, func(cxt context.Context, file string, index int, total int) (result string, outPath string, err error) {
		outPath = file
		result, err = f.format(cxt, file)
		if err != nil {
			return
		}
		fmt.Fprintf(os.Stderr, "Formatted %s (%d of %d)...\n", file, index, total)
		return
	})
}

func (f *formatter) format(errCxt context.Context, file string) (string, error) {
	doc, err := ascii.Open(file)
	if err != nil {
		return "", err
	}
	return adoc.Render(errCxt, doc)
}
