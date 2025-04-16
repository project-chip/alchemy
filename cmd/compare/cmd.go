package compare

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/project-chip/alchemy/cmd/cli"
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
)

type Command struct {
	spec.ParserOptions         `embed:""`
	sdk.SDKOptions             `embed:""`
	common.ASCIIDocAttributes  `embed:""`
	pipeline.ProcessingOptions `embed:""`
	files.OutputOptions        `embed:""`

	Text bool
}

func (c *Command) Run(alchemy *cli.Alchemy) (err error) {
	err = sdk.CheckAlchemyVersion(c.SdkRoot)
	if err != nil {
		return
	}

	specParser, err := spec.NewParser(c.ASCIIDocAttributes.ToList(), c.ToOptions()...)
	if err != nil {
		return err
	}

	err = errata.LoadErrataConfig(specParser.Root)
	if err != nil {
		return
	}

	specFiles, err := pipeline.Start(alchemy, spec.Targeter(specParser.Root))
	if err != nil {
		return err
	}

	specDocs, err := pipeline.Parallel(alchemy, c.ProcessingOptions, specParser, specFiles)
	if err != nil {
		return err
	}

	specBuilder := spec.NewBuilder(specParser.Root)
	specDocs, err = pipeline.Collective(alchemy, c.ProcessingOptions, &specBuilder, specDocs)
	if err != nil {
		return err
	}

	xmlPaths, err := pipeline.Start(alchemy, paths.NewTargeter(filepath.Join(c.SdkRoot, "src/app/zap-templates/zcl/data-model/chip/*.xml")))
	if err != nil {
		return err
	}

	var xmlFiles pipeline.FileSet
	xmlFiles, err = pipeline.Parallel(alchemy, c.ProcessingOptions, files.NewReader("Reading ZAP templates"), xmlPaths)

	if err != nil {
		return
	}

	var zapEntities pipeline.Map[string, *pipeline.Data[[]types.Entity]]
	zapParser := parse.NewZapParser()
	zapEntities, err = pipeline.Parallel(alchemy, c.ProcessingOptions, zapParser, xmlFiles)
	if err != nil {
		return
	}
	zapParser.ResolveReferences()

	var specEntities pipeline.Map[string, *pipeline.Data[[]types.Entity]]
	specEntities, err = pipeline.Collective(alchemy, c.ProcessingOptions, &common.EntityFilter[*spec.Doc, types.Entity]{}, specDocs)

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

		destinations := render.ZAPTemplateDestinations(c.SdkRoot, path, entities.Content, errata)
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

	if c.DryRun {
		return nil
	}

	if c.Text {
		writeText(os.Stdout, diffs)
		return
	}

	jm := json.NewEncoder(os.Stdout)
	jm.SetIndent("", "\t")
	return jm.Encode(diffs)
}
