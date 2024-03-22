package pipeline

import (
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
	slices.SortFunc[[]*Data[T], *Data[T]](data, func(a *Data[T], b *Data[T]) int {
		return strings.Compare(a.Path, b.Path)
	})
}

func DataMapToSlice[T any](data Map[string, *Data[T]]) (slice []*Data[T]) {
	data.Range(func(key string, value *Data[T]) bool {
		slice = append(slice, value)
		return true
	})
	return
}
