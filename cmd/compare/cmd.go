package compare

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/hasty/alchemy/cmd/common"
	"github.com/hasty/alchemy/compare"
	"github.com/hasty/alchemy/internal/files"
	"github.com/hasty/alchemy/internal/pipeline"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/spec"
	"github.com/hasty/alchemy/matter/types"
	"github.com/hasty/alchemy/zap"
	"github.com/hasty/alchemy/zap/generate"
	"github.com/hasty/alchemy/zap/parse"
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

	cxt := context.Background()

	specRoot, _ := cmd.Flags().GetString("specRoot")
	sdkRoot, _ := cmd.Flags().GetString("sdkRoot")
	text, _ := cmd.Flags().GetBool("text")

	asciiSettings := common.ASCIIDocAttributes(cmd)
	pipelineOptions := pipeline.Flags(cmd)
	fileOptions := files.Flags(cmd)

	asciiSettings = append(spec.GithubSettings(), asciiSettings...)

	specFiles, err := pipeline.Start[struct{}](cxt, files.SpecTargeter(specRoot))
	if err != nil {
		return err
	}

	docParser := spec.NewParser(asciiSettings)
	specDocs, err := pipeline.Process[struct{}, *spec.Doc](cxt, pipelineOptions, docParser, specFiles)
	if err != nil {
		return err
	}

	var specParser files.SpecParser
	specDocs, err = pipeline.Process[*spec.Doc, *spec.Doc](cxt, pipelineOptions, &specParser, specDocs)
	if err != nil {
		return err
	}

	xmlPaths, err := pipeline.Start[struct{}](cxt, files.PathsTargeter(filepath.Join(sdkRoot, "src/app/zap-templates/zcl/data-model/chip/*.xml")))
	if err != nil {
		return err
	}

	var xmlFiles pipeline.Map[string, *pipeline.Data[[]byte]]
	xmlFiles, err = pipeline.Process[struct{}, []byte](cxt, pipelineOptions, files.NewReader("Reading ZAP templates"), xmlPaths)

	if err != nil {
		return
	}

	var zapEntities pipeline.Map[string, *pipeline.Data[[]types.Entity]]
	zapParser := parse.NewZapParser()
	zapEntities, err = pipeline.Process[[]byte, []types.Entity](cxt, pipelineOptions, zapParser, xmlFiles)
	if err != nil {
		return
	}
	zapParser.ResolveReferences()

	var specEntities pipeline.Map[string, *pipeline.Data[[]types.Entity]]
	specEntities, err = pipeline.Process[*spec.Doc, []types.Entity](cxt, pipelineOptions, &common.EntityFilter[*spec.Doc, types.Entity]{}, specDocs)

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

		errata, ok := zap.Erratas[filepath.Base(path)]
		if !ok {
			errata = zap.DefaultErrata
		}

		destinations := generate.ZAPTemplateDestinations(sdkRoot, path, entities.Content, errata)
		for templatePath, entities := range destinations {
			var clusters []types.Entity
			for _, e := range entities {
				switch e := e.(type) {
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
	diffs, err = compare.Entities(specParser.Spec, specEntityMap, zapEntityMap)
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
