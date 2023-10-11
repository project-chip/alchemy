package cmd

import (
	"context"
	"fmt"
	"os"
	"sync/atomic"

	"github.com/dolthub/go-mysql-server/sql"
	"github.com/hasty/matterfmt/db"
	"golang.org/x/sync/errgroup"
)

func Database(cxt context.Context, filepaths []string, serial bool) error {
	files, err := getFilePaths(filepaths)
	if err != nil {
		return err
	}
	fmt.Fprintf(os.Stderr, "Reading %d files...\n", len(files))

	sc := sql.NewContext(cxt)

	h := db.New()

	if serial {
		err = readDBSerial(cxt, h, files)
	} else {
		err = readDBParallel(cxt, h, files)
	}
	if err != nil {
		return err
	}
	err = h.Build(sc)
	if err != nil {
		return err
	}
	return h.Run()
}

func readDBSerial(cxt context.Context, h *db.Host, files []string) error {
	for i, file := range files {
		fmt.Fprintf(os.Stderr, "Loading %s (%d of %d)...\n", file, (i + 1), len(files))
		doc, err := getDoc(file)
		if err != nil {
			return err
		}
		err = h.Load(doc)
		if err != nil {
			return err
		}
	}
	return nil
}

func readDBParallel(cxt context.Context, h *db.Host, files []string) error {
	var complete int32
	g, _ := errgroup.WithContext(cxt)
	for i, f := range files {
		func(file string, index int) {
			g.Go(func() error {
				doc, err := getDoc(file)
				if err != nil {
					return err
				}
				err = h.Load(doc)
				if err != nil {
					return err
				}
				done := atomic.AddInt32(&complete, 1)
				fmt.Fprintf(os.Stderr, "Loaded %s (%d of %d)...\n", file, done, len(files))
				return nil
			})
		}(f, i)

	}
	err := g.Wait()
	if err != nil {
		return err
	}
	return nil
}
