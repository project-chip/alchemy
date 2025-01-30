package conformance

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/project-chip/alchemy/matter/types"
)

type IdentifierExpression struct {
	ID     string          `json:"id"`
	Not    bool            `json:"not,omitempty"`
	Field  ComparisonValue `json:"field,omitempty"`
	Entity types.Entity    `json:"-"`
}

func (ie *IdentifierExpression) ASCIIDocString() string {
	var sb strings.Builder
	if ie.Not {
		sb.WriteRune('!')
	}
	sb.WriteString(ie.ID)
	if ie.Field != nil {
		sb.WriteString(ie.Field.ASCIIDocString())
	}
	return sb.String()
}

func (ie *IdentifierExpression) Description() string {
	if ie.Not {
		return fmt.Sprintf("%s is not indicated", ie.ID)
	}
	return fmt.Sprintf("%s is indicated", ie.ID)
}

func (ie *IdentifierExpression) Eval(context Context) (bool, error) {
	return evalIdentifier(context, ie.ID, ie.Not)
}

func evalIdentifier(context Context, id string, not bool) (bool, error) {
	if context.Values != nil {
		v, ok := context.Values[id]
		if ok {
			if b, ok := v.(bool); ok {
				return b != not, nil
			}
		}
	}
	if context.Identifiers != nil {
		if context.VisitedReferences == nil {
			context.VisitedReferences = make(map[string]struct{})
		} else if _, ok := context.VisitedReferences[id]; ok {
			return false, nil
		}
		ref, ok := context.Identifiers.Identifier(id)
		context.VisitedReferences[id] = struct{}{}
		if !ok {
			return false, nil
		}
		if ref, ok := ref.(HasConformance); ok {
			conf := ref.GetConformance()
			if conf != nil {
				cs, err := conf.Eval(context)
				if err != nil {
					return false, err
				}
				return (cs == StateMandatory || cs == StateOptional || cs == StateProvisional || cs == StateDeprecated) != not, nil
			}
		}
	}
	return not, nil
}

func (ie *IdentifierExpression) Equal(e Expression) bool {
	if ie == nil {
		return e == nil
	} else if e == nil {
		return false
	}
	oie, ok := e.(*IdentifierExpression)
	if !ok {
		return false
	}
	if ie.Not != oie.Not {
		return false
	}
	if ie.ID != oie.ID {
		return false
	}
	return true
}

func (ie *IdentifierExpression) MarshalJSON() ([]byte, error) {
	js := map[string]any{
		"type": "identifier",
		"id":   ie.ID,
	}
	if ie.Not {
		js["not"] = true
	}
	return json.Marshal(js)
}

func (ie *IdentifierExpression) Clone() Expression {
	return &IdentifierExpression{Not: ie.Not, ID: ie.ID}
}

type IdentifierValue struct {
	ID     string          `json:"id"`
	Field  ComparisonValue `json:"field,omitempty"`
	Entity types.Entity    `json:"-"`
}

func (ie *IdentifierValue) ASCIIDocString() string {
	var sb strings.Builder
	sb.WriteString(ie.ID)
	if ie.Field != nil {
		sb.WriteString(ie.Field.ASCIIDocString())
	}
	return sb.String()
}

func (ie *IdentifierValue) Description() string {

	return fmt.Sprintf("the value of %s", ie.ID)
}

func (ie *IdentifierValue) Compare(context Context, other ComparisonValue, op ComparisonOperator) (bool, error) {
	return compare(context, op, ie, other)
}

func (ie *IdentifierValue) Equal(e ComparisonValue) bool {
	if ie == nil {
		return e == nil
	} else if e == nil {
		return false
	}
	oie, ok := e.(*IdentifierValue)
	if !ok {
		return false
	}
	if ie.ID != oie.ID {
		return false
	}
	return true
}

func (ie *IdentifierValue) Clone() ComparisonValue {
	return &IdentifierValue{ID: ie.ID}
}

func (ie *IdentifierValue) Value(context Context) (any, error) {
	return identifierValue(context, ie.ID)
}

func identifierValue(context Context, id string) (any, error) {
	if context.Values != nil {
		v, ok := context.Values[id]
		if ok {
			return v, nil
		}
	}
	return nil, fmt.Errorf("unrecognized identifier: %s", id)
}
