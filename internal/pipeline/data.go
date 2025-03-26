package pipeline

import (
	"fmt"
	"iter"
	"slices"
	"strings"
)

type Data[T any] struct {
	Path    string
	Content T
}

func NewData[T any](path string, content T) *Data[T] {
	return &Data[T]{Path: path, Content: content}
}

func SortData[T any](data []*Data[T]) {
	slices.SortStableFunc(data, func(a *Data[T], b *Data[T]) int {
		return strings.Compare(a.Path, b.Path)
	})
}

func DataMapToSlice[T any](data Map[string, *Data[T]]) (slice []*Data[T]) {
	slice = make([]*Data[T], 0, int32(data.Size()))
	for value := range dataMapValues(data) {
		slice = append(slice, value)
	}
	return
}

func dataMapValues[T any](data Map[string, *Data[T]]) iter.Seq[*Data[T]] {
	return func(yield func(*Data[T]) bool) {
		data.Range(func(key string, value *Data[T]) bool {
			return yield(value)
		})
	}
}

func Cast[T any, U any](from Map[string, *Data[T]], to Map[string, *Data[U]]) (err error) {
	from.Range(func(key string, value *Data[T]) bool {
		o, ok := any(value.Content).(U)
		if !ok {
			err = fmt.Errorf("cannot convert %T to %T", value.Content, new(U))
			return false
		}
		to.Store(key, NewData(value.Path, o))
		return true
	})
	return
}
