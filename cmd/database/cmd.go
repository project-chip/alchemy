package database

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/dolthub/go-mysql-server/sql"
	"github.com/hasty/alchemy/cmd/common"
	"github.com/hasty/alchemy/cmd/files"
	"github.com/hasty/alchemy/db"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "db",
	Short: "run a local MySQL DB containing the contents of the Matter spec or the ZAP templates",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		specRoot, _ := cmd.Flags().GetString("specRoot")

		filesOptions := files.Flags(cmd)
		asciiSettings := common.AsciiDocAttributes(cmd)

		address, _ := cmd.Flags().GetString("address")
		port, _ := cmd.Flags().GetInt("port")
		raw, _ := cmd.Flags().GetBool("raw")

		sc := sql.NewContext(context.Background())
		sc.SetCurrentDatabase("matter")

		slog.Info("Loading spec...")
		spec, docs, err := files.LoadSpec(sc, specRoot, filesOptions, asciiSettings)
		if err != nil {
			return err
		}

		h := db.New()
		err = h.Build(sc, spec, docs, raw)
		if err != nil {
			return fmt.Errorf("error building DB: %w", err)
		}
		return h.Run(address, port)
	},
}

func init() {
	Command.Flags().String("specRoot", "", "the src root of your clone of CHIP-Specifications/connectedhomeip-spec")
	Command.Flags().String("sdkRoot", "", "the src root of your clone of project-chip/connectedhomeip")
	Command.Flags().String("address", "localhost", "the address to host the database server on")
	Command.Flags().Int("port", 3306, "the port to run the database server on")
	Command.Flags().Bool("raw", false, "parse the sections directly, bypassing entity building")
}
