package render

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
)

type DependencyTracer struct {
	spec                     *spec.Specification
	GlobalObjectDependencies spec.DocSet
}

func NewDependencyTracer(specification *spec.Specification) *DependencyTracer {
	dt := &DependencyTracer{
		spec:                     specification,
		GlobalObjectDependencies: pipeline.NewConcurrentMap[string, *pipeline.Data[*asciidoc.Document]](),
	}

	return dt
}

func (p DependencyTracer) Name() string {
	return "Tracing dependencies"
}

func (p DependencyTracer) Process(cxt context.Context, inputs []*pipeline.Data[*asciidoc.Document]) (outputs []*pipeline.Data[*asciidoc.Document], err error) {

	docMap := make(map[*asciidoc.Document]struct{}, len(inputs))
	for _, input := range inputs {
		docMap[input.Content] = struct{}{}
	}
	for _, input := range inputs {
		var entities []types.Entity
		entities, err = input.Content.Entities()
		if err != nil {
			return
		}
		p.findDependencies(p.spec, entities, docMap)
	}
	for doc := range docMap {
		outputs = append(outputs, pipeline.NewData(doc.Path.Absolute, doc))
	}
	return
}

func (tg *DependencyTracer) findDependencies(spec *spec.Specification, entities []types.Entity, dependencies map[*asciidoc.Document]struct{}) {
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

func (tg *DependencyTracer) findClusterDependencies(spec *spec.Specification, c *matter.Cluster, dependencies map[*asciidoc.Document]struct{}) {
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

func (tg *DependencyTracer) findFieldSetDependencies(spec *spec.Specification, fs matter.FieldSet, dependencies map[*asciidoc.Document]struct{}) {
	for _, f := range fs {
		tg.findDataTypeDependencies(spec, f.Type, dependencies)
	}
}

func (tg *DependencyTracer) findDataTypeDependencies(spec *spec.Specification, dt *types.DataType, dependencies map[*asciidoc.Document]struct{}) {
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
	switch entity := dt.Entity.(type) {
	case *matter.Struct:
		for _, f := range entity.Fields {
			tg.findDataTypeDependencies(spec, f.Type, dependencies)
		}
	case *matter.Command:
		for _, f := range entity.Fields {
			tg.findDataTypeDependencies(spec, f.Type, dependencies)
		}
	case *matter.Event:
		for _, f := range entity.Fields {
			tg.findDataTypeDependencies(spec, f.Type, dependencies)
		}
	}
	entityDoc, ok := spec.DocRefs[dt.Entity]
	if !ok {
		slog.Warn("missing document for data type", "name", dt.Name, "entity", dt.Entity, "pointer", fmt.Sprintf("%p", dt.Entity))
		return
	}
	_, isGlobal := spec.GlobalObjects[dt.Entity]
	if isGlobal {
		tg.GlobalObjectDependencies.Store(entityDoc.Path.Relative, pipeline.NewData(entityDoc.Path.Relative, entityDoc))
		return
	}
	if _, ok := dependencies[entityDoc]; !ok {
		return
	}
	slog.Debug("dependency found", "name", dt.Name, "path", entityDoc.Path.Relative)
	dependencies[entityDoc] = struct{}{}

}
