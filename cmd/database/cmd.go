package database

import (
	"context"
	"fmt"
	"os"

	"github.com/dolthub/go-mysql-server/sql"
	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/cmd/common"
	"github.com/hasty/alchemy/cmd/files"
	"github.com/hasty/alchemy/db"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "db",
	Short: "run a local MySQL DB containing the contents of the Matter spec or the ZAP templates",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		var paths []string
		specRoot, _ := cmd.Flags().GetString("specRoot")
		if specRoot != "" {
			paths = append(paths, specRoot)
		} else {
			sdkRoot, _ := cmd.Flags().GetString("sdkRoot")
			if sdkRoot != "" {
				paths = append(paths, sdkRoot)
			} else {
				paths = args
			}
		}

		filesOptions := files.Flags(cmd)
		asciiSettings := common.AsciiDocAttributes(cmd)

		address, _ := cmd.Flags().GetString("address")
		port, _ := cmd.Flags().GetInt("port")
		raw, _ := cmd.Flags().GetBool("raw")

		sc := sql.NewContext(context.Background())
		sc.SetCurrentDatabase("matter")

		h := db.New()
		err = files.Process(sc, paths, func(cxt context.Context, file string, index, total int) error {
			fmt.Fprintf(os.Stderr, "Loading %s (%d of %d)...\n", file, index, total)
			doc, err := ascii.OpenFile(file, asciiSettings...)
			if err != nil {
				return fmt.Errorf("error opening file %s: %w", file, err)
			}
			err = h.Load(doc)
			if err != nil {
				return fmt.Errorf("error loading file %s: %w", file, err)
			}
			return nil
		}, filesOptions)
		if err != nil {
			return fmt.Errorf("error processing files: %w", err)
		}
		err = h.Build(sc, raw)
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
