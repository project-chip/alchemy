//go:build db

package cmd

import (
	"github.com/project-chip/alchemy/cmd/database"
)

func init() {
	rootCmd.PersistentFlags().StringSliceP("attribute", "a", []string{}, "attribute for pre-processing asciidoc; this flag can be provided more than once")
	rootCmd.AddCommand(database.Command)
	defaultCommand = "db"
}
