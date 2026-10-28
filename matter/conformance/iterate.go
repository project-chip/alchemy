package conformance

import "iter"

func Iterate(c Conformance) iter.Seq[Conformance] {
	return func(yield func(Conformance) bool) {
		switch c := c.(type) {
		case Set:
			for _, con := range c {
				if !yield(con) {
					return
				}
			}
		default:
			if !yield(c) {
				return
			}
		}
	}
}
