package testplan

import (
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/types"
)

type optionality[T types.Entity] struct {
	Mandatory []T
	Optional  []T
	Features  []conformanceOptional[T]
}

type conformanceOptional[T types.Entity] struct {
	Entity      T
	Conformance conformance.Set
}

func getOptionality[T types.Entity](list []T, getName func(e T) string, getConformance func(e T) conformance.Set) (o optionality[T]) {
	for _, el := range list {
		for _, c := range getConformance(el) {
			switch c := c.(type) {
			case *conformance.Mandatory:
				if c.Expression == nil {
					o.Mandatory = append(o.Mandatory, el)
					break
				}
				co := conformanceOptional[T]{Entity: el}
				co.Conformance = append(co.Conformance, c)
				o.Features = append(o.Features, co)
			case *conformance.Optional:
				if c.Expression == nil {
					o.Optional = append(o.Optional, el)
					continue
				}
				co := conformanceOptional[T]{Entity: el}
				co.Conformance = append(co.Conformance, &conformance.Mandatory{
					Expression: &conformance.LogicalExpression{
						Operand: "&", Left: c.Expression,
						Right: []conformance.Expression{&conformance.IdentifierExpression{
							ID: getName(el),
						}},
					},
				})
				o.Features = append(o.Features, co)
			default:
				continue
			}
			break
		}
	}
	return
}
