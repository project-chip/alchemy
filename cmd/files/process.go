package files

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sync/atomic"
	"unicode/utf8"

	"github.com/hasty/alchemy/ascii"
	"github.com/schollz/progressbar/v3"
	"golang.org/x/sync/errgroup"
)

type FileProcessor func(cxt context.Context, file string, index int, total int) error

func process(cxt context.Context, filepaths []string, options Options, processor FileProcessor, showProgress bool) error {
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
	bar := progressbar.Default(int64(len(files)))
	for i, f := range files {
		func(file string, index int) {
			g.Go(func() error {
				done := atomic.AddInt32(&complete, 1)
				err := processor(errCxt, file, int(done), len(files))
				bar.Describe(ProgressFileName(file))
				bar.Add(1)
				return err
			})
		}(f, i)

	}
	return g.Wait()
}

func Process(cxt context.Context, filepaths []string, processor FileProcessor, options Options) error {
	return process(cxt, filepaths, options, processor, true)
}

type DocProcessor func(cxt context.Context, doc *ascii.Doc, index int, total int) error

func ProcessDocs(cxt context.Context, docs []*ascii.Doc, processor DocProcessor, options Options) (err error) {
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

type FileSaver func(cxt context.Context, file string, index int, total int) (result string, outPath string, err error)

func Save(cxt context.Context, filepaths []string, saver FileSaver, options Options) error {
	return process(cxt, filepaths, options, func(cxt context.Context, file string, index, total int) error {
		result, outPath, err := saver(cxt, file, index, total)
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
	}, false)
}

func ProgressFileName(file string) string {
	file = filepath.Base(file)
	if utf8.RuneCountInString(file) > 20 {
		file = string([]rune(file)[:20]) + "..."
	}
	return fmt.Sprintf("%-20s", file)
}
