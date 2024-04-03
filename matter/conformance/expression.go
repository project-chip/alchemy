package conformance

import (
	"encoding/json"
	"fmt"
)

type Expression interface {
	ASCIIDocString() string
	Description() string

	Eval(context Context) (bool, error)
	Equal(e Expression) bool
	Clone() Expression
}

type HasExpression interface {
	GetExpression() Expression
}

type ExpressionType uint8

const (
	ExpressionTypeUnknown ExpressionType = iota
	ExpressionTypeMandatory
	ExpressionTypeOptional
	ExpressionTypeProvisional
	ExpressionTypeDeprecated
	ExpressionTypeDisallowed
	ExpressionTypeDescribed
	ExpressionTypeGeneric
	ExpressionTypeSet
)

var ExpressionTypeNames = map[ExpressionType]string{
	ExpressionTypeUnknown:     "unknown",
	ExpressionTypeMandatory:   "mandatory",
	ExpressionTypeOptional:    "optional",
	ExpressionTypeProvisional: "provisional",
	ExpressionTypeDeprecated:  "deprecated",
	ExpressionTypeDisallowed:  "disallowed",
	ExpressionTypeDescribed:   "described",
	ExpressionTypeGeneric:     "generic",
	ExpressionTypeSet:         "set",
}

func (et ExpressionType) String() string {
	return ExpressionTypeNames[et]
}

var expressionTypeNameMap map[string]ExpressionType

func init() {
	expressionTypeNameMap = make(map[string]ExpressionType, len(TypeNames))
	for p, n := range ExpressionTypeNames {
		expressionTypeNameMap[n] = p
	}
}

func (et ExpressionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(ExpressionTypeNames[et])
}

func (et *ExpressionType) UnmarshalJSON(data []byte) error {
	var t string
	if err := json.Unmarshal(data, &t); err != nil {
		return fmt.Errorf("error parsing conformance expression type %s: %w", string(data), err)
	}
	var ok bool
	*et, ok = expressionTypeNameMap[t]
	if !ok {
		return fmt.Errorf("unknown conformance expression type: %s", t)
	}
	return nil
}

func unmarshalExpression(js []byte) (e Expression, err error) {
	var rm map[string]json.RawMessage
	err = json.Unmarshal(js, &rm)
	if err != nil {
		return nil, err
	}
	etb, ok := rm["type"]
	if !ok {
		return nil, fmt.Errorf("missing \"type\" field on expression object")
	}
	var et string
	err = json.Unmarshal(etb, &et)
	if err != nil {
		return nil, err
	}

	switch et {
	case "equality":
		var ee EqualityExpression
		err = json.Unmarshal(js, &ee)
		e = &ee
	default:
		err = fmt.Errorf("unknown expression type: %s", et)
	}
	if err != nil {
		return nil, err
	}
	return
}
