package spec

import (
	"sync"

	"github.com/project-chip/alchemy/asciidoc/parse"
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
	TraverseEntities(spec, func(parentCluster *matter.Cluster, parent, entity types.Entity) parse.SearchShould {
		if parent != nil {
			spec.DataTypeRefs.Add(parent, entity)
		}
		return parse.SearchShouldContinue
	})
}

func TraverseEntities(spec *Specification, callback func(parentCluster *matter.Cluster, parent, entity types.Entity) parse.SearchShould) {
	for _, c := range spec.ClustersByName {
		traverseClusterEntities(c, func(parent, entity types.Entity) parse.SearchShould {
			return callback(c, parent, entity)
		})
	}
	traverseEntityList(nil, nil, spec.GlobalObjects.Iterate(), func(parent, entity types.Entity) parse.SearchShould {
		return callback(nil, parent, entity)
	})
}
