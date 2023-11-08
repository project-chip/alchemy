package cmd

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/hasty/alchemy/ascii"
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
		db := &discoBall{}
		err = applyOptions(db, options)
		if err != nil {
			return err
		}
		return db.saveFiles(context.Background(), args, func(cxt context.Context, file string, index int, total int) (result string, outPath string, err error) {
			outPath = file
			result, err = db.run(cxt, file)
			if err != nil {
				return
			}
			fmt.Fprintf(os.Stderr, "Disco-balled %s (%d of %d)...\n", file, index, total)
			return
		})
	},
}

func init() {
	rootCmd.AddCommand(discoCommand)
	discoCommand.Flags().Bool("linkAttributes", false, "whether or not to link attributes table to individual attribute sections")
}

func (db *discoBall) run(cxt context.Context, file string) (string, error) {
	doc, err := ascii.Open(file)
	if err != nil {
		return "", err
	}
	b := disco.NewBall(doc)
	for _, option := range db.options {
		option(b)
	}
	err = b.Run(cxt)
	if err != nil {
		slog.Error("error disco balling", "file", file, "error", err)
		return "", nil
	}
	return adoc.Render(cxt, doc)
}
