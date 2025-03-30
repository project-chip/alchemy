package pics

import (
	"fmt"

	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/types"
)

func ConvertConformance(entity types.Entity, c conformance.Conformance) (Expression, error) {
	switch c := c.(type) {
	case conformance.Set:
		switch len(c) {
		case 0:
			return nil, nil
		case 1:
			return ConvertConformance(entity, c[0])
		default:
			left, err := ConvertConformance(entity, c[0])
			if err != nil {
				return nil, err
			}
			var right []any
			for _, r := range c[1:] {
				re, err := ConvertConformance(entity, r)
				if err != nil {
					return nil, err
				}
				right = append(right, re)
			}
			return NewLogicalExpression(LogicalOperatorOr, left, right)
		}
	case *conformance.Mandatory:
		if c.Expression == nil {
			return nil, nil
		}
		return convertExpression(c.Expression)
	case *conformance.Optional:
		if c.Expression == nil {
			return &PICSExpression{
				PICS: EntityIdentifier(entity),
			}, nil
		}
		left := &PICSExpression{
			PICS: EntityIdentifier(entity),
		}
		right, err := convertExpression(c.Expression)
		if err != nil {
			return nil, err
		}
		return NewLogicalExpression(LogicalOperatorAnd, left, []any{right})
	default:
		return nil, fmt.Errorf("unexpected conformance converting to PICS expression: %T", c)
	}
}

func convertExpression(e conformance.Expression) (Expression, error) {
	switch e := e.(type) {
	case *conformance.IdentifierExpression:
		if e.Entity == nil {
			return nil, fmt.Errorf("conformance identifier expression missing entity: %s", e.ID)
		}
		return &PICSExpression{
			PICS: EntityIdentifier(e.Entity),
			Not:  e.Not,
		}, nil
	case *conformance.FeatureExpression:
		if e.Entity == nil {
			return nil, fmt.Errorf("conformance feature expression missing entity: %s", e.Feature)
		}
		return &PICSExpression{
			PICS: EntityIdentifier(e.Entity),
			Not:  e.Not,
		}, nil
	case *conformance.LogicalExpression:
		left, err := convertExpression(e.Left)
		if err != nil {
			return nil, err
		}
		var right []any
		for _, r := range e.Right {
			re, err := convertExpression(r)
			if err != nil {
				return nil, err
			}
			right = append(right, re)
		}
		return NewLogicalExpression(LogicalOperatorAnd, left, right)
	default:
		return nil, fmt.Errorf("unexpected conformance expression converting to PICS expression: %T", e)
	}
}
