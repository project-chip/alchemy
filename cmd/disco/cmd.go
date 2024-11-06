package disco

import (
	"context"

	"github.com/project-chip/alchemy/asciidoc/render"
	"github.com/project-chip/alchemy/disco"
	"github.com/project-chip/alchemy/internal/files"
	"github.com/project-chip/alchemy/internal/pipeline"
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

	specRoot, _ := cmd.Flags().GetString("specRoot")

	pipelineOptions := pipeline.Flags(cmd)
	fileOptions := files.Flags(cmd)

	writer := files.NewWriter[string]("Writing disco-balled docs", fileOptions)

	err = disco.Pipeline(cxt, specRoot, args, pipelineOptions, getDiscoOptions(cmd), getRenderOptions(cmd), writer)

	return
}

func init() {
	Command.Flags().String("specRoot", "", "the src root of your clone of CHIP-Specifications/connectedhomeip-spec")
	Command.Flags().Bool("linkIndexTables", false, "link index tables to child sections")
	Command.Flags().Bool("addMissingColumns", true, "add standard columns missing from tables")
	Command.Flags().Bool("reorderColumns", true, "rearrange table columns into disco-ball order")
	Command.Flags().Bool("renameTableHeaders", true, "rename table headers to disco-ball standard names")
	Command.Flags().Bool("formatAccess", true, "reformat access columns in disco-ball order")
	Command.Flags().Bool("promoteDataTypes", true, "promote inline data types to Data Types section")
	Command.Flags().Bool("reorderSections", true, "reorder sections in disco-ball order")
	Command.Flags().Bool("normalizeTableOptions", true, "remove existing table options and replace with standard disco-ball options")
	Command.Flags().Bool("normalizeFeatureNames", true, "correct invalid feature names")
	Command.Flags().Bool("fixCommandDirection", true, "normalize command directions")
	Command.Flags().Bool("appendSubsectionTypes", true, "add missing suffixes to data type sections (e.g. \"Bit\", \"Value\", \"Field\", etc.)")
	Command.Flags().Bool("uppercaseHex", true, "uppercase hex values")
	Command.Flags().Bool("addSpaceAfterPunctuation", true, "add missing space after punctuation")
	Command.Flags().Bool("removeExtraSpaces", true, "remove extraneous spaces")
	Command.Flags().Bool("disambiguateConformanceChoice", false, "ensure conformance choices are only used once per document")
	Command.Flags().Bool("normalizeAnchors", false, "rewrite anchors and references without labels")
	Command.Flags().Int("wrap", 0, "the maximum length of a line")
}

type discoOption func(bool) disco.Option

func getDiscoOptions(cmd *cobra.Command) []disco.Option {
	var optionFuncs = map[string]discoOption{
		"linkIndexTables":               disco.LinkIndexTables,
		"addMissingColumns":             disco.AddMissingColumns,
		"reorderColumns":                disco.ReorderColumns,
		"renameTableHeaders":            disco.RenameTableHeaders,
		"formatAccess":                  disco.FormatAccess,
		"promoteDataTypes":              disco.PromoteDataTypes,
		"reorderSections":               disco.ReorderSections,
		"normalizeTableOptions":         disco.NormalizeTableOptions,
		"fixCommandDirection":           disco.FixCommandDirection,
		"appendSubsectionTypes":         disco.AppendSubsectionTypes,
		"uppercaseHex":                  disco.UppercaseHex,
		"addSpaceAfterPunctuation":      disco.AddSpaceAfterPunctuation,
		"removeExtraSpaces":             disco.RemoveExtraSpaces,
		"normalizeFeatureNames":         disco.NormalizeFeatureNames,
		"disambiguateConformanceChoice": disco.DisambiguateConformanceChoice,
		"normalizeAnchors":              disco.NormalizeAnchors,
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

func getRenderOptions(cmd *cobra.Command) []render.Option {
	var renderOptions []render.Option
	wrap, err := cmd.Flags().GetInt("wrap")
	if err == nil {
		renderOptions = append(renderOptions, render.Wrap(wrap))
	}
	return renderOptions
}
