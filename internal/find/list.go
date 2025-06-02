package find

import (
	"iter"
)

func IterateCast[I any, O any](list []I) (i iter.Seq[O]) {
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

func ListCast[I any, O any](list []I) (out []O) {
	for _, l := range list {
		switch o := any(l).(type) {
		case O:
			out = append(out, o)
		}
	}
	return
}
