package files

import (
	"context"
	"os"
	"sync/atomic"

	"github.com/hasty/alchemy/ascii"
	"golang.org/x/sync/errgroup"
)

func Process(cxt context.Context, filepaths []string, processor func(cxt context.Context, file string, index int, total int) error, options Options) error {
	files, err := Paths(filepaths)
	if err != nil {
		return err
	}
	if options.Serial {
		for i, file := range files {
			err = processor(cxt, file, i, len(files))
			if err != nil {
				return err
			}
		}
		return nil
	}
	var complete int32
	g, errCxt := errgroup.WithContext(cxt)
	for i, f := range files {
		func(file string, index int) {
			g.Go(func() error {
				done := atomic.AddInt32(&complete, 1)
				return processor(errCxt, file, int(done), len(files))
			})
		}(f, i)

	}
	return g.Wait()
}

func ProcessDocs(cxt context.Context, docs []*ascii.Doc, processor func(cxt context.Context, doc *ascii.Doc, index int, total int) error, options Options) (err error) {
	if options.Serial {
		for i, d := range docs {
			err = processor(cxt, d, i, len(docs))
			if err != nil {
				return err
			}
		}
		return nil
	}
	var complete int32
	g, errCxt := errgroup.WithContext(cxt)
	for i, d := range docs {
		func(doc *ascii.Doc, index int) {
			g.Go(func() error {
				done := atomic.AddInt32(&complete, 1)
				return processor(errCxt, doc, int(done), len(docs))
			})
		}(d, i)

	}
	return g.Wait()
}

func Save(cxt context.Context, filepaths []string, processor func(cxt context.Context, file string, index int, total int) (result string, outPath string, err error), options Options) error {
	return Process(cxt, filepaths, func(cxt context.Context, file string, index, total int) error {
		result, outPath, err := processor(cxt, file, index, total)
		if err != nil {
			return err
		}
		if !options.DryRun {
			err = os.WriteFile(outPath, []byte(result), os.ModeAppend|0644)
			if err != nil {
				return err
			}
		}
		return nil
	}, options)
}
