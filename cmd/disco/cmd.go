package disco

import (
	"github.com/project-chip/alchemy/asciidoc/render"
	"github.com/project-chip/alchemy/disco"
	"github.com/project-chip/alchemy/internal/files"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var Command = &cobra.Command{
	Use:     "disco filename_pattern",
	Short:   "disco ball Matter spec documents specified by the filename_pattern",
	Long:    ``,
	Aliases: []string{"discoball", "disco-ball"},
	RunE:    discoBall,
}

func discoBall(cmd *cobra.Command, args []string) (err error) {
	cxt := cmd.Context()
	flags := cmd.Flags()

	specRoot, _ := flags.GetString("specRoot")

	pipelineOptions := pipeline.Flags(flags)
	fileOptions := files.Flags(flags)

	writer := files.NewWriter[string]("Writing disco-balled docs", fileOptions)

	err = disco.Pipeline(cxt, specRoot, args, pipelineOptions, getDiscoOptions(flags), getRenderOptions(flags), writer)

	return
}

func init() {
	flags := Command.Flags()
	flags.String("specRoot", "", "the src root of your clone of CHIP-Specifications/connectedhomeip-spec")
	flags.Bool("linkIndexTables", false, "link index tables to child sections")
	flags.Bool("addMissingColumns", true, "add standard columns missing from tables")
	flags.Bool("reorderColumns", true, "rearrange table columns into disco-ball order")
	flags.Bool("renameTableHeaders", true, "rename table headers to disco-ball standard names")
	flags.Bool("formatAccess", true, "reformat access columns in disco-ball order")
	flags.Bool("promoteDataTypes", true, "promote inline data types to Data Types section")
	flags.Bool("reorderSections", true, "reorder sections in disco-ball order")
	flags.Bool("normalizeTableOptions", true, "remove existing table options and replace with standard disco-ball options")
	flags.Bool("normalizeFeatureNames", true, "correct invalid feature names")
	flags.Bool("fixCommandDirection", true, "normalize command directions")
	flags.Bool("appendSubsectionTypes", true, "add missing suffixes to data type sections (e.g. \"Bit\", \"Value\", \"Field\", etc.)")
	flags.Bool("uppercaseHex", true, "uppercase hex values")
	flags.Bool("addSpaceAfterPunctuation", true, "add missing space after punctuation")
	flags.Bool("removeExtraSpaces", true, "remove extraneous spaces")
	flags.Bool("disambiguateConformanceChoice", true, "ensure conformance choices are only used once per document")
	flags.Bool("normalizeAnchors", false, "rewrite anchors and references without labels")
	flags.Bool("removeMandatoryFallbacks", true, "remove fallback values for mandatory fields")
	flags.Int("wrap", 0, "the maximum length of a line")
}

type discoOption func(bool) disco.Option

func getDiscoOptions(flags *pflag.FlagSet) []disco.Option {
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
		"removeMandatoryFallbacks":      disco.RemoveMandatoryFallbacks,
	}
	var discoOptions []disco.Option
	for name, o := range optionFuncs {
		on, err := flags.GetBool(name)
		if err != nil {
			continue
		}
		discoOptions = append(discoOptions, o(on))
	}
	return discoOptions
}

func getRenderOptions(flags *pflag.FlagSet) []render.Option {
	var renderOptions []render.Option
	wrap, err := flags.GetInt("wrap")
	if err == nil {
		renderOptions = append(renderOptions, render.Wrap(wrap))
	}
	return renderOptions
}
