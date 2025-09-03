//go:build db

package database

import (
	"fmt"

	"github.com/dolthub/go-mysql-server/sql"
	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/cmd/cli"
	"github.com/project-chip/alchemy/cmd/common"
	"github.com/project-chip/alchemy/db"
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
	Raw     bool   `default:"false" hidden:"" help:"parse the sections directly, bypassing entity building"`
}

func (cmd *Command) Run(cc *cli.Context) (err error) {

	var specDocs spec.DocSet
	var specification *spec.Specification
	specification, specDocs, err = spec.Parse(cc, cmd.ParserOptions, cmd.ProcessingOptions, cmd.BuilderOptions.List(), cmd.ASCIIDocAttributes.ToList())
	if err != nil {
		return
	}

	docs := make([]*asciidoc.Document, 0, specDocs.Size())
	specDocs.Range(func(key string, value *pipeline.Data[*asciidoc.Document]) bool {
		docs = append(docs, value.Content)
		return true
	})

	sc := sql.NewContext(cc)
	sc.SetCurrentDatabase("matter")

	h := db.New()
	err = h.Build(sc, specification, docs)
	if err != nil {
		return fmt.Errorf("error building DB: %w", err)
	}
	return h.Run(cmd.Address, cmd.Port)
}
