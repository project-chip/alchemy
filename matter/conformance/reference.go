package conformance

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/project-chip/alchemy/matter/types"
)

type ReferenceExpression struct {
	Reference string          `json:"ref"`
	Field     ComparisonValue `json:"field,omitempty"`
	Label     string          `json:"label,omitempty"`
	Not       bool            `json:"not,omitempty"`
	Entity    types.Entity    `json:"-"`
}

func (re *ReferenceExpression) ASCIIDocString() string {
	var s strings.Builder
	if re.Not {
		s.WriteRune('!')
	}
	s.WriteString("<<")
	s.WriteString(re.Reference)
	if len(re.Label) > 0 {
		s.WriteString(", ")
		s.WriteString(re.Label)
	}
	s.WriteString(">>")
	if re.Field != nil {
		s.WriteRune('.')
		s.WriteString(re.Field.ASCIIDocString())
	}
	return s.String()
}

func (re *ReferenceExpression) Description() string {
	if re.Not {
		return fmt.Sprintf("not %s", re.Reference)
	}
	return re.Reference
}

func (re *ReferenceExpression) Eval(context Context) (ExpressionResult, error) {

	if re.Entity == nil {
		return &expressionResult{confidence: ConfidenceImpossible}, nil
	}
	if context.VisitedReferences == nil {
		context.VisitedReferences = make(map[string]struct{})
	} else if _, ok := context.VisitedReferences[re.Reference]; ok {
		return &expressionResult{confidence: ConfidenceImpossible}, nil
	}
	context.VisitedReferences[re.Reference] = struct{}{}
	if ref, ok := re.Entity.(HasConformance); ok {
		conf := ref.GetConformance()
		if conf != nil {
			cs, err := conf.Eval(context)
			if err != nil {
				return nil, err
			}
			switch cs.State {
			case StateMandatory:
				return &expressionResult{value: !re.Not, confidence: cs.Confidence}, nil
			case StateOptional, StateProvisional, StateDeprecated:
				return &expressionResult{value: !re.Not, confidence: cs.Confidence}, nil
			case StateDisallowed:
				return &expressionResult{value: re.Not, confidence: cs.Confidence}, nil
			default:
				return nil, fmt.Errorf("unexpected conformance state evaluating identifier conformance: %s", cs.State.String())

			}
		}
	}
	return &expressionResult{value: !re.Not, confidence: ConfidencePossible}, nil
}

func (re *ReferenceExpression) Equal(e Expression) bool {
	if re == nil {
		return e == nil
	} else if e == nil {
		return false
	}
	ore, ok := e.(*ReferenceExpression)
	if !ok {
		return false
	}
	if re.Not != ore.Not {
		return false
	}
	if re.Reference != ore.Reference {
		return false
	}
	return true
}

func (re *ReferenceExpression) MarshalJSON() ([]byte, error) {
	js := map[string]any{
		"type": "identifier",
		"ref":  re.Reference,
	}
	if re.Not {
		js["not"] = true
	}
	return json.Marshal(js)
}

func (re *ReferenceExpression) Clone() Expression {
	return &ReferenceExpression{Not: re.Not, Reference: re.Reference}
}

type ReferenceValue struct {
	Reference string          `json:"ref"`
	Field     ComparisonValue `json:"field,omitempty"`
	Label     string          `json:"label,omitempty"`
	Entity    types.Entity    `json:"-"`
}

func (re *ReferenceValue) ASCIIDocString() string {
	var s strings.Builder
	s.WriteString("<<")
	s.WriteString(re.Reference)
	if len(re.Label) > 0 {
		s.WriteString(", ")
		s.WriteString(re.Label)
	}
	s.WriteString(">>")
	if re.Field != nil {
		s.WriteRune('.')
		s.WriteString(re.Field.ASCIIDocString())
	}
	return s.String()
}

func (ie *ReferenceValue) Description() string {

	return fmt.Sprintf("the value of %s", ie.Reference)
}

func (ie *ReferenceValue) Compare(context Context, other ComparisonValue, op ComparisonOperator) (ExpressionResult, error) {
	return compare(context, op, ie, other)
}

func (re *ReferenceValue) Equal(e ComparisonValue) bool {
	if re == nil {
		return e == nil
	} else if e == nil {
		return false
	}
	ore, ok := e.(*ReferenceValue)
	if !ok {
		return false
	}
	if re.Reference != ore.Reference {
		return false
	}
	return true
}

func (ie *ReferenceValue) Clone() ComparisonValue {
	return &ReferenceValue{Reference: ie.Reference, Label: ie.Label}
}

func (ie *ReferenceValue) Value(context Context) (ExpressionResult, error) {
	return identifierValue(context, ie.Reference)
}
