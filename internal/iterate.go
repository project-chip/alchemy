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
