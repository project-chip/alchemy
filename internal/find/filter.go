package find

import "iter"

func Filter[T any](list []T, filter func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, l := range list {
			if filter(l) && !yield(l) {
				return
			}
		}
	}
}
