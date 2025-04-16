package dm

import (
	"github.com/project-chip/alchemy/cmd/cli"
	"github.com/project-chip/alchemy/cmd/common"
	"github.com/project-chip/alchemy/dm"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal/files"
	"github.com/project-chip/alchemy/internal/paths"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter/spec"
)

type Command struct {
	common.ASCIIDocAttributes  `embed:""`
	spec.ParserOptions         `embed:""`
	pipeline.ProcessingOptions `embed:""`
	files.OutputOptions        `embed:""`
	dm.DataModelOptions        `embed:""`

	Paths []string `arg:"" optional:""`
}

func (c Command) Run(alchemy *cli.Alchemy) (err error) {
	specParser, err := spec.NewParser(c.ASCIIDocAttributes.ToList(), c.ParserOptions.ToOptions()...)
	if err != nil {
		return err
	}

	err = errata.LoadErrataConfig(specParser.Root)
	if err != nil {
		return
	}

	specBuilder := spec.NewBuilder(specParser.Root, spec.IgnoreHierarchy(true))

	specFiles, err := pipeline.Start(alchemy, specParser.Targets)
	if err != nil {
		return err
	}

	specDocs, err := pipeline.Parallel(alchemy, c.ProcessingOptions, specParser, specFiles)
	if err != nil {
		return err
	}
	specDocs, err = pipeline.Collective(alchemy, c.ProcessingOptions, &specBuilder, specDocs)
	if err != nil {
		return err
	}

	if len(c.Paths) > 0 {
		filter := paths.NewFilter[*spec.Doc](specParser.Root, c.Paths)
		specDocs, err = pipeline.Collective(alchemy, c.ProcessingOptions, filter, specDocs)
		if err != nil {
			return err
		}
	}

	dataModelRenderer := dm.NewRenderer(c.DmRoot, specBuilder.Spec)

	dataModelDocs, err := pipeline.Parallel(alchemy, c.ProcessingOptions, dataModelRenderer, specDocs)
	if err != nil {
		return err
	}

	clusterIDJSON, err := dataModelRenderer.GenerateClusterIDsJson()
	if err != nil {
		return err
	}
	dataModelDocs.Store(clusterIDJSON.Path, clusterIDJSON)

	writer := files.NewWriter[string]("Writing data model", c.OutputOptions)
	err = writer.Write(alchemy, dataModelDocs, c.ProcessingOptions)
	if err != nil {
		return err
	}
	return
}
