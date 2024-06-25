package generate

import (
	"fmt"
	"log/slog"

	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
)

func findDependencies(spec *spec.Specification, entities []types.Entity, dependencies pipeline.Map[string, bool]) {
	for _, m := range entities {
		switch m := m.(type) {
		case *matter.ClusterGroup:
			for _, c := range m.Clusters {
				findClusterDependencies(spec, c, dependencies)
			}
		case *matter.Cluster:
			findClusterDependencies(spec, m, dependencies)
		case *matter.Struct:
			findFieldSetDependencies(spec, m.Fields, dependencies)
		}
	}
}

func findClusterDependencies(spec *spec.Specification, c *matter.Cluster, dependencies pipeline.Map[string, bool]) {
	findFieldSetDependencies(spec, c.Attributes, dependencies)
	for _, s := range c.Structs {
		findFieldSetDependencies(spec, s.Fields, dependencies)
	}
	for _, s := range c.Events {
		findFieldSetDependencies(spec, s.Fields, dependencies)
	}
	for _, s := range c.Commands {
		findFieldSetDependencies(spec, s.Fields, dependencies)
	}
}

func findFieldSetDependencies(spec *spec.Specification, fs matter.FieldSet, dependencies pipeline.Map[string, bool]) {
	for _, f := range fs {
		findDataTypeDependencies(spec, f.Type, dependencies)
	}
}

func findDataTypeDependencies(spec *spec.Specification, dt *types.DataType, dependencies pipeline.Map[string, bool]) {
	if dt == nil {
		return
	}
	if dt.IsArray() {
		findDataTypeDependencies(spec, dt.EntryType, dependencies)
		return
	}
	if dt.Entity != nil {
		path, ok := spec.DocRefs[dt.Entity]
		if !ok {
			slog.Warn("missing document for data type", "name", dt.Name, "entity", dt.Entity, "pointer", fmt.Sprintf("%p", dt.Entity))
			return
		}
		_, loaded := dependencies.LoadOrStore(path, false)
		if !loaded {
			slog.Debug("dependency found", "name", dt.Name, "path", path)
		}
	}
}
