package internal

import (
	"iter"
	"slices"
	"strings"
)

func Iterate[T any](list []T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, t := range list {
			if !yield(t) {
				return
			}
		}
	}
}

func List[T any](i iter.Seq[T]) (list []T) {
	for v := range i {
		list = append(list, v)
	}
	return
}

func IterateMapAlphabetically[K comparable, V any](m map[K]V, nameExtractor func(K) string) iter.Seq2[K, V] {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	slices.SortStableFunc(keys, func(a, b K) int {
		return strings.Compare(nameExtractor(a), nameExtractor(b))
	})
	return func(yield func(K, V) bool) {
		for _, k := range keys {
			if !yield(k, m[k]) {
				return
			}
		}
	}
}
