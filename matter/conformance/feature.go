package conformance

import (
	"encoding/json"
	"fmt"

	"github.com/project-chip/alchemy/matter/types"
)

type FeatureExpression struct {
	Feature string `json:"id"`
	Not     bool   `json:"not"`
	Entity  types.Entity
}

func (fe *FeatureExpression) ASCIIDocString() string {
	if fe.Not {
		return fmt.Sprintf("!%s", fe.Feature)
	}
	return fe.Feature
}

func (fe *FeatureExpression) Description() string {
	if fe.Not {
		return fmt.Sprintf("feature %s is not enabled", fe.Feature)
	}
	return fmt.Sprintf("feature %s is enabled", fe.Feature)
}

func (fe *FeatureExpression) Eval(context Context) (bool, error) {
	return evalIdentifier(context, fe.Feature, fe.Not)
}

func (fe *FeatureExpression) Value(context Context) (any, error) {
	return identifierValue(context, fe.Feature)
}

func (fe *FeatureExpression) Equal(e Expression) bool {
	if fe == nil {
		return e == nil
	} else if e == nil {
		return false
	}
	ofe, ok := e.(*FeatureExpression)
	if !ok {
		return false
	}
	if fe.Not != ofe.Not {
		return false
	}
	if fe.Feature != ofe.Feature {
		return false
	}
	return true
}

func (fe *FeatureExpression) MarshalJSON() ([]byte, error) {
	js := map[string]any{
		"type": "feature",
		"id":   fe.Feature,
	}
	if fe.Not {
		js["not"] = true
	}
	return json.Marshal(js)
}

func (fe *FeatureExpression) Clone() Expression {
	return &FeatureExpression{Not: fe.Not, Feature: fe.Feature}
}

type FeatureValue struct {
	Feature string `json:"id"`
}

func (ie *FeatureValue) ASCIIDocString() string {
	return ie.Feature
}

func (ie *FeatureValue) Description() string {

	return fmt.Sprintf("the value of %s", ie.Feature)
}

func (ie *FeatureValue) Compare(context Context, other ComparisonValue, op ComparisonOperator) (bool, error) {
	return compare(context, op, ie, other)
}

func (ie *FeatureValue) Equal(e ComparisonValue) bool {
	if ie == nil {
		return e == nil
	} else if e == nil {
		return false
	}
	oie, ok := e.(*FeatureValue)
	if !ok {
		return false
	}
	if ie.Feature != oie.Feature {
		return false
	}
	return true
}

func (ie *FeatureValue) Clone() ComparisonValue {
	return &FeatureValue{Feature: ie.Feature}
}

func (ie *FeatureValue) Value(context Context) (any, error) {
	return identifierValue(context, ie.Feature)
}
