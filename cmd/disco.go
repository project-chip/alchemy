package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/cmd/files"
	"github.com/hasty/alchemy/disco"
	"github.com/hasty/alchemy/render/adoc"
	"github.com/spf13/cobra"
)

type discoBall struct {
	processor

	options []disco.Option
}

var discoCommand = &cobra.Command{
	Use:   "disco",
	Short: "disco ball Matter spec documents",
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

		options := getOptions()
		linkAttributes, _ := cmd.Flags().GetBool("linkAttributes")
		options = append(options, Disco(disco.LinkAttributes(linkAttributes)))

		return files.Save(context.Background(), args, func(cxt context.Context, file string, index, total int) (result string, outPath string, err error) {
			outPath = file
			var doc *ascii.Doc
			doc, err = ascii.Open(file)
			if err != nil {
				return
			}
			b := disco.NewBall(doc)
			for _, option := range options {
				option(b)
			}
			err = b.Run(cxt)
			if err != nil {
				return
			}
			result, err = adoc.Render(cxt, doc)
			if err != nil {
				return
			}

			fmt.Fprintf(os.Stderr, "Disco-balled %s (%d of %d)...\n", file, index, total)
			return
		},
			getFilesOptions())

	},
}

func init() {
	rootCmd.AddCommand(discoCommand)
	discoCommand.Flags().Bool("linkAttributes", false, "whether or not to link attributes table to individual attribute sections")
}
