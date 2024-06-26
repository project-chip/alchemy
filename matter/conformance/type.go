package conformance

import (
	"encoding/json"
	"fmt"
)

type Type uint8

const (
	TypeUnknown Type = iota
	TypeMandatory
	TypeOptional
	TypeProvisional
	TypeDeprecated
	TypeDisallowed
	TypeDescribed
	TypeGeneric
	TypeSet
)

var TypeNames = map[Type]string{
	TypeUnknown:     "unknown",
	TypeMandatory:   "mandatory",
	TypeOptional:    "optional",
	TypeProvisional: "provisional",
	TypeDeprecated:  "deprecated",
	TypeDisallowed:  "disallowed",
	TypeDescribed:   "described",
	TypeGeneric:     "generic",
	TypeSet:         "set",
}

func (ct Type) String() string {
	return TypeNames[ct]
}

var typeNameMap map[string]Type

func init() {
	typeNameMap = make(map[string]Type, len(TypeNames))
	for p, n := range TypeNames {
		typeNameMap[n] = p
	}
}

func (ct Type) MarshalJSON() ([]byte, error) {
	return json.Marshal(TypeNames[ct])
}

func (ct *Type) UnmarshalJSON(data []byte) error {
	var t string
	if err := json.Unmarshal(data, &t); err != nil {
		return fmt.Errorf("error parsing conformance type %s: %w", string(data), err)
	}
	var ok bool
	*ct, ok = typeNameMap[t]
	if !ok {
		return fmt.Errorf("unknown conformance type: %s", t)
	}
	return nil
}
