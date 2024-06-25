package cmd

import (
	"fmt"
	"os"

	"github.com/project-chip/alchemy/config"
	"github.com/spf13/cobra"
)

var versionCommand = &cobra.Command{
	Use:   "version",
	Short: "displays the current version",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		fmt.Fprintf(os.Stdout, "version: %v\n", config.Version())
		return nil
	},
}

func init() {
	rootCmd.AddCommand(versionCommand)
}
