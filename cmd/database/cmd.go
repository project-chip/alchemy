//go:build db

package database

import (
	"context"
	"fmt"

	"github.com/dolthub/go-mysql-server/sql"
	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/cmd/common"
	"github.com/hasty/alchemy/db"
	"github.com/hasty/alchemy/internal/files"
	"github.com/hasty/alchemy/internal/pipeline"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "db",
	Short: "run a local MySQL DB containing the contents of the Matter spec or the ZAP templates",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		cxt := context.Background()
		specRoot, _ := cmd.Flags().GetString("specRoot")

		asciiSettings := common.AsciiDocAttributes(cmd)

		address, _ := cmd.Flags().GetString("address")
		port, _ := cmd.Flags().GetInt("port")
		raw, _ := cmd.Flags().GetBool("raw")

		pipelineOptions := pipeline.Flags(cmd)

		specFiles, err := pipeline.Start[struct{}](cxt, files.SpecTargeter(specRoot))
		if err != nil {
			return err
		}

		docParser := ascii.NewParser(asciiSettings)
		specDocMap, err := pipeline.Process[struct{}, *ascii.Doc](cxt, pipelineOptions, docParser, specFiles)
		if err != nil {
			return err
		}
		var specParser files.SpecParser
		specDocMap, err = pipeline.Process[*ascii.Doc, *ascii.Doc](cxt, pipelineOptions, &specParser, specDocMap)
		if err != nil {
			return err
		}

		specDocs := make([]*ascii.Doc, 0, specDocMap.Size())
		specDocMap.Range(func(key string, value *pipeline.Data[*ascii.Doc]) bool {
			specDocs = append(specDocs, value.Content)
			return true
		})

		sc := sql.NewContext(cxt)
		sc.SetCurrentDatabase("matter")

		h := db.New()
		err = h.Build(sc, specParser.Spec, specDocs, raw)
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
