package disco

import (
	"slices"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/matterfmt/ascii"
)

type HasElements interface {
	SetElements([]interface{}) error
	GetElements() []interface{}
}

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

func filter(parent HasElements, callback func(i interface{}) (remove bool, shortCircuit bool)) (shortCircuit bool) {
	i := 0
	elements := parent.GetElements()
	var removed bool
	for i < len(elements) {
		e := elements[i]
		if ae, ok := e.(*ascii.Element); ok {
			e = ae.Base
		}
		switch el := e.(type) {
		case HasElements:
			shortCircuit = filter(el, callback)
		}
		if shortCircuit {
			break
		}
		remove, shortCircuit := callback(e)
		var empty []interface{}
		if remove {
			elements = slices.Replace(elements, i, i+1, empty...)
			removed = true
			remove = false
		} else {
			i++
		}
		if shortCircuit {
			break
		}
	}
	if removed {
		parent.SetElements(elements)
	}
	return
}

func traverse(parent HasElements, elements []interface{}, callback func(el interface{}, parent HasElements, index int) bool) bool {
	for i, e := range elements {
		if callback(e, parent, i) {
			return true
		}

	}
	for _, e := range elements {

		switch el := e.(type) {
		case HasElements:
			if traverse(el, el.GetElements(), callback) {
				return true
			}
		}
	}
	return false
}
