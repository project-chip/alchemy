package disco

import (
	"context"

	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/disco"
	"github.com/hasty/alchemy/internal/files"
	"github.com/hasty/alchemy/internal/pipeline"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:     "disco",
	Short:   "disco ball Matter spec documents",
	Long:    ``,
	Aliases: []string{"discoball", "disco-ball"},
	RunE:    discoBall,
}

func discoBall(cmd *cobra.Command, args []string) (err error) {
	cxt := context.Background()

	var targeter pipeline.Targeter
	var filter *files.PathFilter[*ascii.Doc]
	specRoot, _ := cmd.Flags().GetString("specRoot")
	if specRoot != "" {
		targeter = files.SpecTargeter(specRoot)
		if len(args) > 0 {
			filter = files.NewPathFilter[*ascii.Doc](args)
		}
	} else {
		targeter = files.PathsTargeter(args...)
	}
	if err != nil {
		return err
	}

	pipelineOptions := pipeline.Flags(cmd)
	fileOptions := files.Flags(cmd)
	discoOptions := getDiscoOptions(cmd)

	writer := files.NewWriter[string]("Writing disco-balled docs", fileOptions)

	err = disco.Pipeline(cxt, targeter, pipelineOptions, discoOptions, filter, writer)

	return
}

func init() {
	Command.Flags().Bool("linkIndexTables", false, "link index tables to child sections")
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
		"linkIndexTables":          disco.LinkIndexTables,
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
