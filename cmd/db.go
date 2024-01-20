//go:build db

package cmd

import "github.com/hasty/alchemy/cmd/database"

func init() {
	rootCmd.AddCommand(database.Command)

}
