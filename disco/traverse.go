package disco

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/matterfmt/ascii"
)

func find[T any](elements []interface{}, callback func(t T) bool) bool {
	for _, e := range elements {
		if ae, ok := e.(*ascii.Element); ok {
			e = ae.Base
		}
		var shortCircuit bool
		switch el := e.(type) {
		case T:
			shortCircuit = callback(el)
		case types.WithElements:
			shortCircuit = find(el.GetElements(), callback)
		case *ascii.Section:
			shortCircuit = find(el.Elements, callback)
		}
		if shortCircuit {
			return true
		}

	}
	return false
}

func traverse(parent types.WithElements, elements []interface{}, callback func(interface{}, types.WithElements, int) bool) bool {
	for i, e := range elements {
		if callback(e, parent, i) {
			return true
		}

		switch el := e.(type) {
		case types.WithElements:
			if traverse(el, el.GetElements(), callback) {
				return true
			}
		}
	}
	return false
}
