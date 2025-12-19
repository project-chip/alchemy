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

/*func IterateDataTypes(spec *Specification, callback func(cluster *matter.Cluster, parent, entity types.Entity)) {
	for _, c := range spec.ClustersByName {
		for p, e := range c.IterateDataTypes() {
			callback(c, p, e)
		}
	}
	for e := range spec.GlobalObjects {
		callback(nil, nil, e)
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
}*/

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

/*
func traverseEntityFields(parent types.Entity, e types.Entity, callback func(cluster *matter.Cluster, parent, entity types.Entity) parse.SearchShould) parse.SearchShould {
	switch callback(nil, nil, e) {
	case parse.SearchShouldSkip:
	case parse.SearchShouldStop:
		return parse.SearchShouldStop
	case parse.SearchShouldContinue:
		switch en := parent.(type) {
		}
	}
	return parse.SearchShouldContinue
	var fields matter.FieldSet
	switch en := parent.(type) {
			case *matter.Struct:
				fields = en.Fields
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
	for _, f := range fieldSet {
		switch callback(nil, parent, f) {
		case parse.SearchShouldContinue:
		case parse.SearchShouldSkip:
		case parse.SearchShouldStop:
			return parse.SearchShouldStop
		}
		for p, e := range f.IterateDataTypes() {
			callback(nil, p, e)
		}
	}
}

func traverseFieldDataType(parent types.Entity, callback func(cluster *matter.Cluster, parent, entity types.Entity) parse.SearchShould) parse.SearchShould {
	var fields matter.FieldSet
	switch en := parent.(type) {
	case *matter.Struct:
		fields = en.Fields
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
	for _, f := range fieldSet {
		switch callback(nil, parent, f) {
		case parse.SearchShouldContinue:
		case parse.SearchShouldSkip:
		case parse.SearchShouldStop:
			return parse.SearchShouldStop
		}
		for p, e := range f.IterateDataTypes() {
			callback(nil, p, e)
		}
	}
}
*/
