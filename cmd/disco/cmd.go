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

		discoOptions := getDiscoOptions(cmd)

		var fileOptions = files.Flags(cmd)

		return files.Save(context.Background(), args, func(cxt context.Context, file string, index, total int) (result string, outPath string, err error) {
			outPath = file
			var doc *ascii.Doc
			doc, err = ascii.OpenFile(file)
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
	Command.Flags().Bool("linkAttributes", false, "link attributes table to individual attribute sections")
	Command.Flags().Bool("addMissingColumns", true, "add standard columns missing from tables")
	Command.Flags().Bool("reorderColumns", true, "rearrange table columns into disco-ball order")
	Command.Flags().Bool("renameTableHeaders", true, "rename table headers to disco-ball standard names")
	Command.Flags().Bool("formatAccess", true, "reformat access columns in disco-ball order")
	Command.Flags().Bool("promoteDataTypes", true, "promote inline data types to Data Types section")
	Command.Flags().Bool("reorderSections", true, "reorder sections in disco-ball order")
	Command.Flags().Bool("normalizeTableOptions", true, "remove existing table options and replace with standard disco-ball options")
	Command.Flags().Bool("fixCommandDirection", true, "normalize command directions")
	Command.Flags().Bool("appendSubsectionTypes", true, "add missing suffixes to data type sections (e.g. \"Bit\", \"Value\", \"Field\", etc.)")
	Command.Flags().Bool("uppercaseHex", true, "uppercase hex values")
	Command.Flags().Bool("addSpaceAfterPunctuation", true, "add missing space after punctuation")
	Command.Flags().Bool("removeExtraSpaces", true, "remove extraneous spaces")
}

type discoOption func(bool) disco.Option

func getDiscoOptions(cmd *cobra.Command) []disco.Option {
	var optionFuncs = map[string]discoOption{
		"linkAttributes":           disco.LinkAttributes,
		"addMissingColumns":        disco.AddMissingColumns,
		"reorderColumns":           disco.ReorderColumns,
		"renameTableHeaders":       disco.RenameTableHeaders,
		"formatAccess":             disco.FormatAccess,
		"promoteDataTypes":         disco.PromoteDataTypes,
		"reorderSections":          disco.ReorderSections,
		"normalizeTableOptions":    disco.NormalizeTableOptions,
		"fixCommandDirection":      disco.FixCommandDirection,
		"appendSubsectionTypes":    disco.AppendSubsectionTypes,
		"uppercaseHex":             disco.UppercaseHex,
		"addSpaceAfterPunctuation": disco.AddSpaceAfterPunctuation,
		"removeExtraSpaces":        disco.RemoveExtraSpaces,
	}
	var discoOptions []disco.Option
	for name, o := range optionFuncs {
		on, err := cmd.Flags().GetBool(name)
		if err != nil {
			continue
		}
		discoOptions = append(discoOptions, o(on))
	}
	return discoOptions
}
