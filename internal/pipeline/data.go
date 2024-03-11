package pipeline

import (
	"slices"
	"strings"
)

type Data[T any] struct {
	Path    string
	Content T
}

func SortData[T any](data []*Data[T]) {
	slices.SortFunc[[]*Data[T], *Data[T]](data, func(a *Data[T], b *Data[T]) int {
		return strings.Compare(a.Path, b.Path)
	})
}
