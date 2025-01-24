package conformance

import (
	"encoding/json"
	"fmt"
	"strings"
)

type ReferenceExpression struct {
	Reference string `json:"ref"`
	Field     string `json:"field,omitempty"`
	Label     string `json:"label,omitempty"`
	Not       bool   `json:"not,omitempty"`
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
	if re.Field != "" {
		s.WriteRune('.')
		s.WriteString(re.Field)
	}
	return s.String()
}

func (re *ReferenceExpression) Description() string {
	if re.Not {
		return fmt.Sprintf("not %s", re.Reference)
	}
	return re.Reference
}

func (re *ReferenceExpression) Eval(context Context) (bool, error) {
	return evalReference(context, re.Reference, re.Not)
}

func evalReference(context Context, id string, not bool) (bool, error) {

	if context.References != nil {
		if context.VisitedReferences == nil {
			context.VisitedReferences = make(map[string]struct{})
		} else if _, ok := context.VisitedReferences[id]; ok {
			return false, nil
		}
		ref, ok := context.References.Reference(id)
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
	Reference string `json:"ref"`
	Field     string `json:"field,omitempty"`
	Label     string `json:"label,omitempty"`
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
	if re.Field != "" {
		s.WriteRune('.')
		s.WriteString(re.Field)
	}
	return s.String()
}

func (ie *ReferenceValue) Description() string {

	return fmt.Sprintf("the value of %s", ie.Reference)
}

func (ie *ReferenceValue) Compare(context Context, other ComparisonValue, op ComparisonOperator) (bool, error) {
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

func (ie *ReferenceValue) Value(context Context) (any, error) {
	return identifierValue(context, ie.Reference)
}
