//go:build db

package database

import (
	"context"
	"fmt"

	"github.com/dolthub/go-mysql-server/sql"
	"github.com/project-chip/alchemy/cmd/common"
	"github.com/project-chip/alchemy/db"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "db",
	Short: "run a local MySQL DB containing the contents of the Matter spec or the ZAP templates",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		cxt := context.Background()
		specRoot, _ := cmd.Flags().GetString("specRoot")

		asciiSettings := common.ASCIIDocAttributes(cmd)

		address, _ := cmd.Flags().GetString("address")
		port, _ := cmd.Flags().GetInt("port")
		raw, _ := cmd.Flags().GetBool("raw")

		errata.LoadErrataConfig(specRoot)

		pipelineOptions := pipeline.Flags(cmd)

		specFiles, err := pipeline.Start(cxt, spec.Targeter(specRoot))
		if err != nil {
			return err
		}

		docParser, err := spec.NewParser(specRoot, asciiSettings)
		if err != nil {
			return err
		}
		specDocs, err := pipeline.Parallel(cxt, pipelineOptions, docParser, specFiles)
		if err != nil {
			return err
		}
		var specBuilder spec.Builder
		specDocs, err = pipeline.Collective(cxt, pipelineOptions, &specBuilder, specDocs)
		if err != nil {
			return err
		}

		docs := make([]*spec.Doc, 0, specDocs.Size())
		specDocs.Range(func(key string, value *pipeline.Data[*spec.Doc]) bool {
			docs = append(docs, value.Content)
			return true
		})

		sc := sql.NewContext(cxt)
		sc.SetCurrentDatabase("matter")

		h := db.New()
		err = h.Build(sc, specBuilder.Spec, docs, raw)
		if err != nil {
			return fmt.Errorf("error building DB: %w", err)
		}
		return h.Run(address, port)
	},
}

func init() {
	Command.Flags().String("specRoot", "connectedhomeip-spec", "the src root of your clone of CHIP-Specifications/connectedhomeip-spec")
	Command.Flags().String("address", "localhost", "the address to host the database server on")
	Command.Flags().Int("port", 3306, "the port to run the database server on")
	Command.Flags().Bool("raw", false, "parse the sections directly, bypassing entity building")
}
