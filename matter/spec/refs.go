package spec

import (
	"sync"

	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

type EntityRefs[T comparable] struct {
	sync.RWMutex
	refs map[types.Entity]pipeline.Map[T, struct{}]
}

func NewEntityRefs[T comparable]() EntityRefs[T] {
	return EntityRefs[T]{
		refs: make(map[types.Entity]pipeline.Map[T, struct{}]),
	}
}

func (cr *EntityRefs[T]) Add(c T, m types.Entity) {
	cr.Lock()
	cm, ok := cr.refs[m]
	if !ok {
		cm = pipeline.NewConcurrentMap[T, struct{}]()
		cr.refs[m] = cm
	}
	cm.Store(c, struct{}{})
	cr.Unlock()
}

func (cr *EntityRefs[T]) Remove(c T, m types.Entity) {
	cr.Lock()
	cm, ok := cr.refs[m]
	if ok {
		cm.Delete(c)
	}
	cr.Unlock()
}

func (cr *EntityRefs[T]) Get(m types.Entity) (pipeline.Map[T, struct{}], bool) {
	cr.RLock()
	cm, ok := cr.refs[m]
	cr.RUnlock()
	return cm, ok
}

func (spec *Specification) BuildDataTypeReferences() {
	iterateOverDataTypes(spec, func(cluster *matter.Cluster, parent, entity types.Entity) {
		spec.DataTypeRefs.Add(parent, entity)
	})
}

func iterateOverDataTypes(spec *Specification, callback func(cluster *matter.Cluster, parent, entity types.Entity)) {
	for _, c := range spec.ClustersByName {
		for p, e := range c.IterateDataTypes() {
			callback(c, p, e)
		}
	}
	for e := range spec.GlobalObjects {
		switch en := e.(type) {
		case *matter.Struct:
			for _, f := range en.Fields {
				for p, e := range f.IterateDataTypes() {
					callback(nil, p, e)
				}
			}
		case *matter.Command:
			for _, f := range en.Fields {
				for p, e := range f.IterateDataTypes() {
					callback(nil, p, e)
				}
			}
			if en.Response != nil && en.Response.Entity != nil && en.Response.Name == "" {
				callback(nil, en, en.Response.Entity)
			}
		case *matter.Event:
			for _, f := range en.Fields {
				for p, e := range f.IterateDataTypes() {
					callback(nil, p, e)
				}
			}
		}
	}
}
