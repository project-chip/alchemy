package format

import (
	"context"
	"fmt"
	"os"

	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/ascii/render"
	"github.com/hasty/alchemy/cmd/files"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "format",
	Short: "format Matter spec documents",
	RunE: func(cmd *cobra.Command, args []string) (err error) {

		var fileOptions = files.Flags(cmd)
		return files.Save(context.Background(), args, func(cxt context.Context, file string, index, total int) (result string, outPath string, err error) {
			outPath = file
			var doc *ascii.Doc
			doc, err = ascii.OpenFile(file)
			if err != nil {
				return
			}
			result, err = render.Render(cxt, doc)
			if err != nil {
				return
			}
			if fileOptions.Serial {
				fmt.Fprintf(os.Stderr, "Formatted %s (%d of %d)...\n", file, index+1, total)
			}
			return
		},
			fileOptions)
	},
}
