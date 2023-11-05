package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/dolthub/go-mysql-server/sql"
	"github.com/hasty/matterfmt/ascii"
	"github.com/hasty/matterfmt/db"
)

type databaseReader struct {
	processor
	asciiParser

	raw bool

	address string
	port    int
}

func Database(cxt context.Context, filepaths []string, options ...Option) error {
	dbr := &databaseReader{}
	for _, opt := range options {
		err := opt(dbr)
		if err != nil {
			return err
		}
	}
	return dbr.run(cxt, filepaths)
}

func (dbr *databaseReader) run(cxt context.Context, filepaths []string) error {
	sc := sql.NewContext(cxt)
	sc.SetCurrentDatabase("matter")

	h := db.New()
	err := dbr.processFiles(cxt, filepaths, func(cxt context.Context, file string, index, total int) error {
		fmt.Fprintf(os.Stderr, "Loading %s (%d of %d)...\n", file, index, total)
		doc, err := ascii.Open(file)
		if err != nil {
			return err
		}
		err = h.Load(doc)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	err = h.Build(sc, dbr.raw)
	if err != nil {
		return err
	}
	return h.Run("localhost", 3306)
}
