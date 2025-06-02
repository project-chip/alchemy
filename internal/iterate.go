package internal

import "iter"

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
