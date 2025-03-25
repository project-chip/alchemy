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
	"github.com/project-chip/alchemy/zap/generate"
	"github.com/project-chip/alchemy/zap/parse"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "compare",
	Short: "compare the spec to zap-templates and output a JSON diff",
	RunE:  compareSpec,
}

func init() {
	Command.Flags().String("specRoot", "connectedhomeip-spec", "the src root of your clone of CHIP-Specifications/connectedhomeip-spec")
	Command.Flags().String("sdkRoot", "connectedhomeip", "the src root of your clone of project-chip/connectedhomeip")
	Command.Flags().Bool("text", false, "output as text")
}

func compareSpec(cmd *cobra.Command, args []string) (err error) {

	cxt := cmd.Context()

	specRoot, _ := cmd.Flags().GetString("specRoot")
	sdkRoot, _ := cmd.Flags().GetString("sdkRoot")
	text, _ := cmd.Flags().GetBool("text")

	err = errata.LoadErrataConfig(specRoot)
	if err != nil {
		return
	}

	asciiSettings := common.ASCIIDocAttributes(cmd)
	pipelineOptions := pipeline.Flags(cmd)
	fileOptions := files.Flags(cmd)

	err = sdk.CheckAlchemyVersion(sdkRoot)
	if err != nil {
		return
	}

	specFiles, err := pipeline.Start(cxt, spec.Targeter(specRoot))
	if err != nil {
		return err
	}

	docParser, err := spec.NewParser(specRoot, asciiSettings)
	if err != nil {
		return err
	}
	specDocs, err := pipeline.Parallel(cxt, pipelineOptions, docParser, specFiles)
	if err != nil {
		return err
	}

	specBuilder := spec.NewBuilder()
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

		errata := errata.GetZAP(path)

		destinations := generate.ZAPTemplateDestinations(sdkRoot, path, entities.Content, errata)
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

	if fileOptions.DryRun {
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
