package cmd

import (
	"context"

	"github.com/hasty/alchemy/cmd/database"
	"github.com/spf13/cobra"
)

var dbCommand = &cobra.Command{
	Use:   "db",
	Short: "run a local MySQL DB containing the contents of the Matter spec or the ZAP templates",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		var paths []string
		specRoot, _ := cmd.Flags().GetString("specRoot")
		if specRoot != "" {
			paths = append(paths, specRoot)
		} else {
			zclRoot, _ := cmd.Flags().GetString("zclRoot")
			if zclRoot != "" {
				paths = append(paths, zclRoot)
			} else {
				paths = args
			}
		}

		var options database.Options
		options.FilesOptions = getFilesOptions()
		options.AsciiSettings = getAsciiAttributes()

		options.Address, _ = cmd.Flags().GetString("address")
		options.Port, _ = cmd.Flags().GetInt("port")
		options.Raw, _ = cmd.Flags().GetBool("raw")
		return database.Run(context.Background(), paths, options)
	},
}

func init() {
	rootCmd.AddCommand(dbCommand)
	dbCommand.Flags().String("specRoot", "", "the src root of your clone of CHIP-Specifications/connectedhomeip-spec")
	dbCommand.Flags().String("zclRoot", "", "the src root of your clone of project-chip/connectedhomeip")
	dbCommand.Flags().String("address", "localhost", "the address to host the database server on")
	dbCommand.Flags().Int("port", 3306, "the port to run the database server on")
	dbCommand.Flags().Bool("raw", false, "parse the sections directly, bypassing model building")
}
