package database

import (
	"context"
	"fmt"
	"os"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/dolthub/go-mysql-server/sql"
	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/cmd/files"
	"github.com/hasty/alchemy/db"
)

type Options struct {
	AsciiSettings []configuration.Setting
	Raw           bool
	FilesOptions  files.Options

	Address string
	Port    int
}

func Run(cxt context.Context, filepaths []string, options Options) error {
	sc := sql.NewContext(cxt)
	sc.SetCurrentDatabase("matter")

	h := db.New()
	err := files.Process(cxt, filepaths, func(cxt context.Context, file string, index, total int) error {
		fmt.Fprintf(os.Stderr, "Loading %s (%d of %d)...\n", file, index, total)
		doc, err := ascii.Open(file, options.AsciiSettings...)
		if err != nil {
			return err
		}
		err = h.Load(doc)
		if err != nil {
			return err
		}
		return nil
	}, options.FilesOptions)
	if err != nil {
		return err
	}
	err = h.Build(sc, options.Raw)
	if err != nil {
		return err
	}
	return h.Run(options.Address, options.Port)
}
