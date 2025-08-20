package pipeline

import (
	"github.com/puzpuzpuz/xsync/v3"
)

type StringSet Map[string, *Data[string]]
type FileSet Map[string, *Data[[]byte]]

type Map[K comparable, V any] interface {
	Load(key K) (value V, ok bool)
	Store(key K, value V)
	LoadOrStore(key K, value V) (actual V, loaded bool)
	LoadAndStore(key K, value V) (actual V, loaded bool)
	LoadAndDelete(key K) (value V, loaded bool)
	Delete(key K)
	Size() int
	Range(f func(key K, value V) bool)
}

type unsafeMap[K comparable, V any] map[K]V

func (um unsafeMap[K, V]) Load(key K) (value V, ok bool) {
	v, ok := um[key]
	return v, ok
}

func (um unsafeMap[K, V]) Store(key K, value V) {
	um[key] = value
}

func (um unsafeMap[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	actual, loaded = um[key]
	if !loaded {
		um[key] = value
		actual = value
	}
	return
}

func (um unsafeMap[K, V]) LoadAndStore(key K, value V) (actual V, loaded bool) {
	actual, loaded = um[key]
	um[key] = value
	return
}

func (um unsafeMap[K, V]) LoadAndDelete(key K) (value V, loaded bool) {
	value, loaded = um[key]
	delete(um, key)
	return
}

func (um unsafeMap[K, V]) Range(f func(key K, value V) bool) {
	for k, v := range um {
		if !f(k, v) {
			break
		}
	}
}

func (um unsafeMap[K, V]) Size() int {
	return len(um)
}

func (um unsafeMap[K, V]) Delete(key K) {
	delete(um, key)
}

func NewMap[K comparable, V any]() Map[K, V] {
	return NewMapPresized[K, V](0)
}

func NewMapPresized[K comparable, V any](size int) Map[K, V] {
	unsafeMap := make(unsafeMap[K, V], size)
	return unsafeMap
}

func NewConcurrentMap[K comparable, V any]() Map[K, V] {
	return NewConcurrentMapPresized[K, V](0)
}

func NewConcurrentMapPresized[K comparable, V any](size int) Map[K, V] {
	return xsync.NewMapOf[K, V](xsync.WithPresize(size))
}
