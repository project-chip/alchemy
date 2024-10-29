package find

import "iter"

func FirstPairFunc[K comparable, V any](m map[K]V, filter func(K) bool) (key K, value V, ok bool) {
	for key, value = range FindPairFunc(m, filter) {
		ok = true
		return
	}
	return
}

func FindPairFunc[K comparable, V any](m map[K]V, filter func(K) bool) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range m {
			if filter(k) && yield(k, v) {
				return
			}
		}
	}
}

func DeleteFunc[K comparable, V any](m map[K]V, filter func(K) bool) {
	for k := range m {
		if filter(k) {
			delete(m, k)
		}
	}
}
