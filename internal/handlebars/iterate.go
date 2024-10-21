package handlebars

import (
	"iter"
	"reflect"
)

func Iterate[T any](value any) iter.Seq[T] {
	return func(yield func(T) bool) {
		if reflect.TypeOf(value).Kind() != reflect.Slice {
			return
		}
		slice := reflect.ValueOf(value)
		for i := 0; i < slice.Len(); i++ {
			el := slice.Index(i).Interface()
			switch el := el.(type) {
			case T:
				if !yield(el) {
					return
				}
			}
		}
	}
}
