package generate

import (
	"fmt"
	"log/slog"

	"github.com/hasty/alchemy/matter"
)

func findDependencies(spec *matter.Spec, models []matter.Model, dependencies *concurrentMap[bool]) {
	for _, m := range models {
		switch m := m.(type) {
		case *matter.Cluster:
			findClusterDependencies(spec, m, dependencies)
		case *matter.Struct:
			findFieldSetDependencies(spec, m.Fields, dependencies)
		}
	}
}

func findClusterDependencies(spec *matter.Spec, c *matter.Cluster, dependencies *concurrentMap[bool]) {
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

func findFieldSetDependencies(spec *matter.Spec, fs matter.FieldSet, dependencies *concurrentMap[bool]) {
	for _, f := range fs {
		findDataTypeDependencies(spec, f.Type, dependencies)
	}
}

func findDataTypeDependencies(spec *matter.Spec, dt *matter.DataType, dependencies *concurrentMap[bool]) {
	if dt == nil {
		return
	}
	if dt.IsArray() {
		findDataTypeDependencies(spec, dt.EntryType, dependencies)
		return
	}
	if dt.Model != nil {
		path, ok := spec.DocRefs[dt.Model]
		if !ok {
			slog.Warn("missing document for data type", "name", dt.Name, "model", dt.Model, "pointer", fmt.Sprintf("%p", dt.Model))
			return
		}
		dependencies.Lock()
		_, ok = dependencies.Map[path]
		if !ok {
			slog.Warn("dependency found", "name", dt.Name, "path", path)
			dependencies.Map[path] = false
		}
		dependencies.Unlock()
	}
}
