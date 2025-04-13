package spec

import (
	"sync"

	"github.com/project-chip/alchemy/matter/types"
)

type EntityRefs[T comparable] struct {
	sync.RWMutex
	refs map[types.Entity]map[T]struct{}
}

func NewEntityRefs[T comparable]() EntityRefs[T] {
	return EntityRefs[T]{
		refs: make(map[types.Entity]map[T]struct{}),
	}
}

func (cr *EntityRefs[T]) Add(c T, m types.Entity) {
	cr.Lock()
	cm, ok := cr.refs[m]
	if !ok {
		cm = make(map[T]struct{})
		cr.refs[m] = cm
	}
	cm[c] = struct{}{}
	cr.Unlock()
}

func (cr *EntityRefs[T]) Get(m types.Entity) (map[T]struct{}, bool) {
	cr.RLock()
	cm, ok := cr.refs[m]
	cr.RUnlock()
	return cm, ok
}
