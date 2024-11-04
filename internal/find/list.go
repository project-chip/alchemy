package find

import "iter"

func ToList[T any](i iter.Seq[T]) (out []T) {
	for l := range i {
		out = append(out, l)
	}
	return
}

func CastList[I any, O any](list []I) (i iter.Seq[O]) {
	return func(yield func(O) bool) {
		for _, l := range list {
			switch o := any(l).(type) {
			case O:
				if !yield(o) {
					return
				}
			}
		}
	}
}
