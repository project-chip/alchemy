package conformance

import (
	"encoding/json"
	"fmt"
)

type ReferenceExpression struct {
	Reference string `json:"ref"`
	Not       bool   `json:"not,omitempty"`
}

func (ie *ReferenceExpression) AsciiDocString() string {
	if ie.Not {
		return fmt.Sprintf("!%s", ie.Reference)
	}
	return ie.Reference
}

func (ie *ReferenceExpression) Description() string {
	if ie.Not {
		return fmt.Sprintf("not %s", ie.Reference)
	}
	return ie.Reference
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
