package cli

import (
	"fmt"
	"log/slog"
	"path/filepath"

	"github.com/project-chip/alchemy/cmd/common"
	"github.com/project-chip/alchemy/dm"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal/files"
	"github.com/project-chip/alchemy/internal/paths"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter/spec"
)

type DataModel struct {
	common.ASCIIDocAttributes  `embed:""`
	spec.ParserOptions         `embed:""`
	pipeline.ProcessingOptions `embed:""`
	files.OutputOptions        `embed:""`
	dm.DataModelOptions        `embed:""`

	Paths   []string `arg:"" optional:""`
	Exclude []string `short:"e"  help:"exclude files matching this file pattern" group:"Data Model:"`
}

func (c *DataModel) Run(cc *Context) (err error) {
	specParser, err := spec.NewParser(c.ASCIIDocAttributes.ToList(), c.ParserOptions)
	if err != nil {
		return err
	}

	err = errata.LoadErrataConfig(c.ParserOptions.Root)
	if err != nil {
		return
	}

	specBuilder := spec.NewBuilder(c.ParserOptions.Root, spec.IgnoreHierarchy(true))

	specFiles, err := pipeline.Start(cc, specParser.Targets)
	if err != nil {
		return err
	}

	specDocs, err := pipeline.Parallel(cc, c.ProcessingOptions, specParser, specFiles)
	if err != nil {
		return err
	}

	specDocs, err = pipeline.Collective(cc, c.ProcessingOptions, &specBuilder, specDocs)
	if err != nil {
		return err
	}

	if len(c.Paths) > 0 {
		filter := paths.NewFilter[*spec.Doc](c.ParserOptions.Root, c.Paths)
		specDocs, err = pipeline.Collective(cc, c.ProcessingOptions, filter, specDocs)
		if err != nil {
			return err
		}
	}

	if len(c.Exclude) > 0 {
		filter := paths.NewFilter[*spec.Doc](c.ParserOptions.Root, c.Exclude)
		filter.Exclude = true
		specDocs, err = pipeline.Collective(cc, c.ProcessingOptions, filter, specDocs)
		if err != nil {
			return
		}
	}

	if len(specBuilder.Spec.Errors) > 0 && !c.Force {
		if c.Force {
			slog.Warn("Ignoring parse errors; proceed with caution")
		} else {
			for _, specError := range specBuilder.Spec.Errors {
				path, _ := specError.Origin()
				path = filepath.Join(specBuilder.Spec.Root, path)
				_, ok := specDocs.Load(path)
				if ok {
					err = fmt.Errorf("specified document has errors: %s %s", path, specError.Error())
					return
				}
			}
		}
	}

	dataModelRenderer := dm.NewRenderer(c.DmRoot, specBuilder.Spec)

	dataModelDocs, err := pipeline.Parallel(cc, c.ProcessingOptions, dataModelRenderer, specDocs)
	if err != nil {
		return err
	}

	globalDocs, err := dataModelRenderer.GenerateGlobalObjects()
	if err != nil {
		return err
	}

	globalDocs.Range(func(key string, value *pipeline.Data[string]) bool {
		dataModelDocs.Store(key, value)
		return true
	})

	clusterIDJSON, err := dataModelRenderer.GenerateClusterIDsJson()
	if err != nil {
		return err
	}
	dataModelDocs.Store(clusterIDJSON.Path, clusterIDJSON)

	writer := files.NewWriter[string]("Writing data model", c.OutputOptions)
	err = writer.Write(cc, dataModelDocs, c.ProcessingOptions)
	if err != nil {
		return err
	}
	return
}
