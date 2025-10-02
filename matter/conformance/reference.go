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
		// If we were unable to resolve this reference, then we can't evaluate its conformance
		return &expressionResult{confidence: ConfidenceImpossible}, nil
	}
	var err error
	er := &expressionResult{value: !re.Not, confidence: ConfidencePossible}
	if ref, ok := re.Entity.(HasConformance); ok {
		conf := ref.GetConformance()
		if conf != nil {
			if !context.MarkVisit(re.Reference) {
				return &expressionResult{confidence: ConfidenceImpossible}, nil

			}
			var cs ConformanceState
			cs, err = conf.Eval(context)
			if err == nil {
				switch cs.State {
				case StateMandatory:
					er.value = !re.Not
					er.confidence = cs.Confidence
				case StateOptional, StateProvisional, StateDeprecated:
					// Even if it's definitely optional, provisional or deprecated, we should return possible, since the server could choose not to implement
					er.value = !re.Not
					er.confidence = ConfidencePossible
				case StateDisallowed:
					er.value = re.Not
					er.confidence = cs.Confidence
				default:
					err = fmt.Errorf("unexpected conformance state evaluating identifier conformance: %s", cs.State.String())
				}
			}
			context.UnmarkVisit(re.Reference)
		}
	}
	if err != nil {
		return nil, err
	}
	return er, nil
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
