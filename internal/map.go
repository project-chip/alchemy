package internal

import (
	"iter"
	"slices"
)

type OrderedMap[K comparable, V any] struct {
	items map[K]V
	keys  []K
}

func NewOrderedMap[K comparable, V any]() *OrderedMap[K, V] {
	return &OrderedMap[K, V]{
		items: make(map[K]V),
	}
}

func (om *OrderedMap[K, V]) Get(key K) (value V, ok bool) {
	value, ok = om.items[key]
	return
}

func (om *OrderedMap[K, V]) Set(key K, value V) {
	if _, exists := om.items[key]; exists {
		om.items[key] = value
		return
	}
	om.items[key] = value
	om.keys = append(om.keys, key)
}

func (om *OrderedMap[K, V]) Remove(key K) (value V, ok bool) {
	if value, ok = om.items[key]; ok {
		for i, k := range om.keys {
			if k == key {
				om.keys = append(om.keys[:i], om.keys[i+1:]...)
				break
			}
		}
		delete(om.items, key)
		return
	}
	return
}

func (om *OrderedMap[K, V]) Keys() []K {
	return om.keys
}

func (om *OrderedMap[K, V]) Values() iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, key := range om.keys {
			if !yield(om.items[key]) {
				return
			}
		}
	}
}

func (om *OrderedMap[K, V]) Pairs() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for _, key := range om.keys {
			if !yield(key, om.items[key]) {
				return
			}
		}
	}
}

func (om *OrderedMap[K, V]) Len() int {
	return len(om.items)
}

func (om *OrderedMap[K, V]) Sort(compare func(a, b K) int) {
	slices.SortFunc(om.keys, compare)
}

func (om *OrderedMap[K, V]) SortKeys(sort func(keys []K)) {
	sort(om.keys)
}
