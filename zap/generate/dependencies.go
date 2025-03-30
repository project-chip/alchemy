package generate

import (
	"fmt"
	"log/slog"

	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
)

func (tg *TemplateGenerator) findDependencies(spec *spec.Specification, entities []types.Entity, dependencies pipeline.Map[string, bool]) {
	for _, m := range entities {
		switch m := m.(type) {
		case *matter.ClusterGroup:
			for _, c := range m.Clusters {
				tg.findClusterDependencies(spec, c, dependencies)
			}
		case *matter.Cluster:
			tg.findClusterDependencies(spec, m, dependencies)
		case *matter.Struct:
			tg.findFieldSetDependencies(spec, m.Fields, dependencies)
		}
	}
}

func (tg *TemplateGenerator) findClusterDependencies(spec *spec.Specification, c *matter.Cluster, dependencies pipeline.Map[string, bool]) {
	tg.findFieldSetDependencies(spec, c.Attributes, dependencies)
	for _, s := range c.Structs {
		tg.findFieldSetDependencies(spec, s.Fields, dependencies)
	}
	for _, s := range c.Events {
		tg.findFieldSetDependencies(spec, s.Fields, dependencies)
	}
	for _, s := range c.Commands {
		tg.findFieldSetDependencies(spec, s.Fields, dependencies)
	}
}

func (tg *TemplateGenerator) findFieldSetDependencies(spec *spec.Specification, fs matter.FieldSet, dependencies pipeline.Map[string, bool]) {
	for _, f := range fs {
		tg.findDataTypeDependencies(spec, f.Type, dependencies)
	}
}

func (tg *TemplateGenerator) findDataTypeDependencies(spec *spec.Specification, dt *types.DataType, dependencies pipeline.Map[string, bool]) {
	if dt == nil {
		return
	}
	if dt.IsArray() {
		tg.findDataTypeDependencies(spec, dt.EntryType, dependencies)
		return
	}
	if dt.Entity == nil {
		return
	}
	_, isGlobal := spec.GlobalObjects[dt.Entity]
	if isGlobal {
		tg.globalObjectDependencies.Store(dt.Entity, struct{}{})
		return
	}
	entityDoc, ok := spec.DocRefs[dt.Entity]
	if !ok {
		slog.Warn("missing document for data type", "name", dt.Name, "entity", dt.Entity, "pointer", fmt.Sprintf("%p", dt.Entity))
		return
	}
	_, loaded := dependencies.LoadOrStore(entityDoc.Path.Relative, false)
	if !loaded {
		slog.Debug("dependency found", "name", dt.Name, "path", entityDoc.Path.Relative)
	}

}
