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

func iterateOverDataTypes(spec *Specification, callback func(cluster *matter.Cluster, parent types.Entity, entity types.Entity)) {
	for _, c := range spec.ClustersByName {
		if c.Features != nil {
			callback(c, c, c.Features)
		}
		for _, en := range c.Bitmaps {
			callback(c, c, en)
		}
		for _, en := range c.Enums {
			callback(c, c, en)
		}
		for _, en := range c.Structs {
			callback(c, c, en)
			for _, f := range en.Fields {
				iterateOverFieldDataTypes(spec, c, f, callback)
			}
		}
		for _, cmd := range c.Commands {
			callback(c, c, cmd)
			for _, f := range cmd.Fields {
				iterateOverFieldDataTypes(spec, c, f, callback)
			}
			if cmd.Response != nil && cmd.Response.Entity != nil && cmd.Response.Name == "" {
				callback(c, cmd, cmd.Response.Entity)
			}
		}
		for _, ev := range c.Events {
			callback(c, c, ev)
			for _, f := range ev.Fields {
				iterateOverFieldDataTypes(spec, c, f, callback)
			}
		}
		for _, a := range c.Attributes {
			callback(c, c, a)
			iterateOverFieldDataTypes(spec, c, a, callback)
		}
	}
	for e := range spec.GlobalObjects {
		switch en := e.(type) {
		case *matter.Struct:
			for _, f := range en.Fields {
				callback(nil, en, f)
				iterateOverFieldDataTypes(spec, nil, f, callback)
			}
		case *matter.Command:
			for _, f := range en.Fields {
				iterateOverFieldDataTypes(spec, nil, f, callback)
			}
			if en.Response != nil && en.Response.Entity != nil && en.Response.Name == "" {
				callback(nil, en, en.Response.Entity)
			}
		case *matter.Event:
			for _, f := range en.Fields {
				iterateOverFieldDataTypes(spec, nil, f, callback)
			}
		}
	}
}

func iterateOverFieldDataTypes(spec *Specification, cluster *matter.Cluster, field *matter.Field, callback func(cluster *matter.Cluster, parent types.Entity, entity types.Entity)) {
	if field.Type == nil {
		return
	}
	var entity types.Entity
	if field.Type.IsArray() {
		if field.Type.EntryType == nil {
			return
		}
		if field.Type.EntryType.Entity == nil {
			return
		}
		entity = field.Type.EntryType.Entity
	} else {
		entity = field.Type.Entity
	}
	if entity == nil {
		return
	}
	callback(cluster, field, entity)
	switch entity := entity.(type) {
	case *matter.Struct:
		for _, f := range entity.Fields {
			iterateOverFieldDataTypes(spec, cluster, f, callback)
		}
	}
}
