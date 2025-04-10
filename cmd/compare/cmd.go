package compare

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/project-chip/alchemy/cmd/common"
	"github.com/project-chip/alchemy/compare"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal/files"
	"github.com/project-chip/alchemy/internal/paths"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
	"github.com/project-chip/alchemy/sdk"
	"github.com/project-chip/alchemy/zap/parse"
	"github.com/project-chip/alchemy/zap/render"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "compare",
	Short: "compare the spec to zap-templates and output a JSON diff",
	RunE:  compareSpec,
}

func init() {
	flags := Command.Flags()
	flags.String("specRoot", "connectedhomeip-spec", "the src root of your clone of CHIP-Specifications/connectedhomeip-spec")
	flags.String("sdkRoot", "connectedhomeip", "the src root of your clone of project-chip/connectedhomeip")
	flags.Bool("text", false, "output as text")
}

func compareSpec(cmd *cobra.Command, args []string) (err error) {

	cxt := cmd.Context()
	flags := cmd.Flags()

	asciiSettings := common.ASCIIDocAttributes(flags)

	sdkRoot, _ := flags.GetString("sdkRoot")
	text, _ := flags.GetBool("text")

	pipelineOptions := pipeline.PipelineOptions(flags)
	parserOptions := spec.ParserOptions(flags)
	outputOptions := files.OutputOptions(flags)

	err = sdk.CheckAlchemyVersion(sdkRoot)
	if err != nil {
		return
	}

	specParser, err := spec.NewParser(asciiSettings, parserOptions...)
	if err != nil {
		return err
	}

	err = errata.LoadErrataConfig(specParser.Root)
	if err != nil {
		return
	}

	specFiles, err := pipeline.Start(cxt, spec.Targeter(specParser.Root))
	if err != nil {
		return err
	}

	specDocs, err := pipeline.Parallel(cxt, pipelineOptions, specParser, specFiles)
	if err != nil {
		return err
	}

	specBuilder := spec.NewBuilder(specParser.Root)
	specDocs, err = pipeline.Collective(cxt, pipelineOptions, &specBuilder, specDocs)
	if err != nil {
		return err
	}

	xmlPaths, err := pipeline.Start(cxt, paths.NewTargeter(filepath.Join(sdkRoot, "src/app/zap-templates/zcl/data-model/chip/*.xml")))
	if err != nil {
		return err
	}

	var xmlFiles pipeline.FileSet
	xmlFiles, err = pipeline.Parallel(cxt, pipelineOptions, files.NewReader("Reading ZAP templates"), xmlPaths)

	if err != nil {
		return
	}

	var zapEntities pipeline.Map[string, *pipeline.Data[[]types.Entity]]
	zapParser := parse.NewZapParser()
	zapEntities, err = pipeline.Parallel(cxt, pipelineOptions, zapParser, xmlFiles)
	if err != nil {
		return
	}
	zapParser.ResolveReferences()

	var specEntities pipeline.Map[string, *pipeline.Data[[]types.Entity]]
	specEntities, err = pipeline.Collective(cxt, pipelineOptions, &common.EntityFilter[*spec.Doc, types.Entity]{}, specDocs)

	if err != nil {
		return
	}

	zapEntityMap := make(map[string][]types.Entity, zapEntities.Size())
	zapEntities.Range(func(path string, entities *pipeline.Data[[]types.Entity]) bool {
		zapEntityMap[path] = entities.Content
		return true
	})

	specEntityMap := make(map[string][]types.Entity, specEntities.Size())
	specEntities.Range(func(path string, entities *pipeline.Data[[]types.Entity]) bool {

		errata := errata.GetSDK(path)

		destinations := render.ZAPTemplateDestinations(sdkRoot, path, entities.Content, errata)
		for templatePath, entities := range destinations {
			var clusters []types.Entity
			for _, e := range entities {
				switch e := e.(type) {
				case *matter.ClusterGroup:
					for _, c := range e.Clusters {
						clusters = append(clusters, c)
					}
				case *matter.Cluster:
					clusters = append(clusters, e)
				}
			}
			if len(clusters) > 0 {
				specEntityMap[templatePath] = clusters
			}
		}
		return true
	})

	var diffs []*compare.ClusterDifferences
	diffs, err = compare.Entities(specBuilder.Spec, specEntityMap, zapEntityMap)
	if err != nil {
		return
	}

	if outputOptions.DryRun {
		return nil
	}

	if text {
		writeText(os.Stdout, diffs)
		return
	}

	jm := json.NewEncoder(os.Stdout)
	jm.SetIndent("", "\t")
	return jm.Encode(diffs)
}
