package disco

import (
	"context"
	"fmt"
	"os"

	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/ascii/render"
	"github.com/hasty/alchemy/cmd/files"
	"github.com/hasty/alchemy/disco"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "disco",
	Short: "disco ball Matter spec documents",
	Long:  ``,
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

		linkAttributes, _ := cmd.Flags().GetBool("linkAttributes")

		var discoOptions []disco.Option
		discoOptions = append(discoOptions, disco.LinkAttributes(linkAttributes))

		var fileOptions = files.Flags(cmd)

		return files.Save(context.Background(), args, func(cxt context.Context, file string, index, total int) (result string, outPath string, err error) {
			outPath = file
			var doc *ascii.Doc
			doc, err = ascii.Open(file)
			if err != nil {
				return
			}
			b := disco.NewBall(doc)
			for _, option := range discoOptions {
				option(b)
			}
			err = b.Run(cxt)
			if err != nil {
				return
			}
			result, err = render.Render(cxt, doc)
			if err != nil {
				return
			}
			if fileOptions.Serial {
				fmt.Fprintf(os.Stderr, "Disco-balled %s (%d of %d)...\n", file, index+1, total)
			}
			return
		},
			fileOptions)

	},
}

func init() {
	Command.Flags().Bool("linkAttributes", false, "whether or not to link attributes table to individual attribute sections")
}
