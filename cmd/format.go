package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/cmd/files"
	"github.com/hasty/alchemy/render/adoc"
	"github.com/spf13/cobra"
)

var formatCommand = &cobra.Command{
	Use:   "format",
	Short: "format Matter spec documents",
	RunE: func(cmd *cobra.Command, args []string) (err error) {

		return files.Save(context.Background(), args, func(cxt context.Context, file string, index, total int) (result string, outPath string, err error) {
			outPath = file
			var doc *ascii.Doc
			doc, err = ascii.Open(file)
			if err != nil {
				return
			}
			result, err = adoc.Render(cxt, doc)
			if err != nil {
				return
			}
			fmt.Fprintf(os.Stderr, "Formatted %s (%d of %d)...\n", file, index, total)
			return
		},
			getFilesOptions())
	},
}

func init() {
	rootCmd.AddCommand(formatCommand)
}
