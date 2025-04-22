//go:build db

package database

import (
	"fmt"

	"github.com/dolthub/go-mysql-server/sql"
	"github.com/project-chip/alchemy/cmd/cli"
	"github.com/project-chip/alchemy/cmd/common"
	"github.com/project-chip/alchemy/db"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter/spec"
)

type Command struct {
	common.ASCIIDocAttributes  `embed:""`
	spec.ParserOptions         `embed:""`
	spec.BuilderOptions        `embed:""`
	pipeline.ProcessingOptions `embed:""`

	Address string `default:"localhost" help:"the address to host the database server on"`
	Port    int    `default:"3306" help:"the port to run the database server on"`
	Raw     bool   `default:"false" help:"parse the sections directly, bypassing entity building"`
}

func (cmd *Command) Run(cc *cli.Context) (err error) {
	specParser, err := spec.NewParser(cmd.ASCIIDocAttributes.ToList(), cmd.ParserOptions)
	if err != nil {
		return err
	}

	errata.LoadErrataConfig(cmd.ParserOptions.Root)

	specFiles, err := pipeline.Start(cc, specParser.Targets)
	if err != nil {
		return err
	}

	specDocs, err := pipeline.Parallel(cc, cmd.ProcessingOptions, specParser, specFiles)
	if err != nil {
		return err
	}
	specBuilder := spec.NewBuilder(cmd.ParserOptions.Root, spec.IgnoreHierarchy(cmd.IgnoreHierarchy))
	specDocs, err = pipeline.Collective(cc, cmd.ProcessingOptions, &specBuilder, specDocs)
	if err != nil {
		return err
	}

	docs := make([]*spec.Doc, 0, specDocs.Size())
	specDocs.Range(func(key string, value *pipeline.Data[*spec.Doc]) bool {
		docs = append(docs, value.Content)
		return true
	})

	sc := sql.NewContext(cc)
	sc.SetCurrentDatabase("matter")

	h := db.New()
	err = h.Build(sc, specBuilder.Spec, docs, cmd.Raw)
	if err != nil {
		return fmt.Errorf("error building DB: %w", err)
	}
	return h.Run(cmd.Address, cmd.Port)
}
