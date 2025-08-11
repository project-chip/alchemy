package cli

import (
	"github.com/project-chip/alchemy/cmd/common"
	"github.com/project-chip/alchemy/dm"
	"github.com/project-chip/alchemy/internal/files"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter/spec"
)

type DataModel struct {
	common.ASCIIDocAttributes  `embed:""`
	spec.ParserOptions         `embed:""`
	pipeline.ProcessingOptions `embed:""`
	files.OutputOptions        `embed:""`
	dm.DataModelOptions        `embed:""`

	spec.FilterOptions `embed:""`
}

func (c *DataModel) Run(cc *Context) (err error) {

	var specification *spec.Specification
	var specDocs spec.DocSet
	specification, specDocs, err = spec.Parse(cc, c.ParserOptions, c.ProcessingOptions, c.ASCIIDocAttributes.ToList())

	if err != nil {
		return err
	}

	specDocs, err = filterSpecDocs(cc, specDocs, specification, c.FilterOptions, c.ProcessingOptions)
	if err != nil {
		return
	}

	specDocs, err = filterSpecErrors(cc, specDocs, specification, c.FilterOptions, c.ProcessingOptions)
	if err != nil {
		return
	}

	err = checkSpecErrors(cc, specification, c.FilterOptions, specDocs)
	if err != nil {
		return
	}

	dataModelRenderer := dm.NewRenderer(c.DmRoot, specification)

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
