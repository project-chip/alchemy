package cmd

import (
	"context"

	"github.com/hasty/alchemy/cmd/dump"
	"github.com/spf13/cobra"
)

var dumpCommand = &cobra.Command{
	Use:   "dump",
	Short: "dump the parse tree of Matter documents",
	RunE: func(cmd *cobra.Command, args []string) (err error) {

		var options dump.Options
		options.AsciiSettings = getAsciiAttributes()
		options.Ascii, _ = cmd.Flags().GetBool("ascii")
		options.Json, _ = cmd.Flags().GetBool("json")
		return dump.Dump(context.Background(), args, options)
	},
}

func init() {
	rootCmd.AddCommand(dumpCommand)
	dumpCommand.Flags().Bool("ascii", false, "dump asciidoc object model")
	dumpCommand.Flags().Bool("json", false, "dump json object model")
}
