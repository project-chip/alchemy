package render

import (
	"slices"

	"github.com/beevik/etree"
	"github.com/project-chip/alchemy/matter/types"
)

func (cr *configuratorRenderer) reorderConfigurator(configuratorElement *etree.Element) error {
	entityOrder := make(map[types.Entity]int)
	var order int
	for _, d := range cr.configurator.Docs {
		entities, err := d.OrderedEntities()
		if err != nil {
			return err
		}
		for _, e := range entities {
			entityOrder[e] = order
			order++
		}
	}
	cr.reorderElement(configuratorElement, entityOrder)
	return nil
}

func (cr *configuratorRenderer) reorderElement(parent *etree.Element, entityOrder map[types.Entity]int) {
	if len(parent.Child) == 0 {
		return
	}
	type tokenChunk struct {
		tokens  []etree.Token
		element *etree.Element
		entity  types.Entity
		order   int
	}
	lastChunk := &tokenChunk{}
	var chunks []*tokenChunk
	for _, t := range parent.Child {
		element, ok := t.(*etree.Element)
		if !ok {
			lastChunk.tokens = append(lastChunk.tokens, t)
			continue
		}
		cr.reorderElement(element, entityOrder)
		entity, ok := cr.elementMap[element]
		if !ok {
			lastChunk.tokens = append(lastChunk.tokens, t)
			continue
		}
		order, ok := entityOrder[entity]
		if !ok {
			lastChunk.tokens = append(lastChunk.tokens, t)
			continue
		}
		lastChunk.element = element
		lastChunk.entity = entity
		lastChunk.order = order
		chunks = append(chunks, lastChunk)
		lastChunk = &tokenChunk{}
	}
	slices.SortStableFunc(chunks, func(a *tokenChunk, b *tokenChunk) int {
		if a.order < b.order {
			return -1
		} else if a.order > b.order {
			return 1
		}
		return 0
	})
	var children []etree.Token
	for _, tc := range chunks {
		children = append(children, tc.tokens...)
		children = append(children, tc.element)
	}
	children = append(children, lastChunk.tokens...)
	parent.Child = children
}
