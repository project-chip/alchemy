package compare

import (
	"encoding/json"
	"fmt"
)

type Source uint8

const (
	SourceUnknown Source = iota
	SourceSpec
	SourceZAP
)

var (
	sourceNames = map[Source]string{
		SourceUnknown: "unknown",
		SourceSpec:    "spec",
		SourceZAP:     "zap",
	}
	sourceValues = map[string]Source{
		"unknown": SourceUnknown,
		"spec":    SourceSpec,
		"zap":     SourceZAP,
	}
)

func (s Source) MarshalJSON() ([]byte, error) {
	return json.Marshal(sourceNames[s])
}

func (s *Source) UnmarshalJSON(data []byte) (err error) {
	var ss string
	if err := json.Unmarshal(data, &ss); err != nil {
		return err
	}
	var ok bool
	if *s, ok = sourceValues[ss]; !ok {
		return fmt.Errorf("unknown source: %s", ss)
	}
	return nil
}

type DiffProperty uint8

const (
	DiffPropertyUnknown DiffProperty = iota
	DiffPropertyName
	DiffPropertyHierarchy
	DiffPropertyPICS
	DiffPropertyRole
	DiffPropertyType
	DiffPropertyIsArray
	DiffPropertyConstraint
	DiffPropertyDefault
	DiffPropertyConformance
	DiffPropertyQuality
	DiffPropertyAccess
	DiffPropertyBit
)

var (
	diffPropertyNames = map[DiffProperty]string{
		DiffPropertyUnknown:     "unknown",
		DiffPropertyName:        "name",
		DiffPropertyHierarchy:   "hierarchy",
		DiffPropertyPICS:        "pics",
		DiffPropertyRole:        "role",
		DiffPropertyType:        "type",
		DiffPropertyIsArray:     "isArray",
		DiffPropertyConstraint:  "constraint",
		DiffPropertyDefault:     "default",
		DiffPropertyConformance: "conformance",
		DiffPropertyQuality:     "quality",
		DiffPropertyAccess:      "access",
		DiffPropertyBit:         "bit",
	}
	diffPropertyValues = map[string]DiffProperty{
		"unknown":     DiffPropertyUnknown,
		"name":        DiffPropertyName,
		"hierarchy":   DiffPropertyHierarchy,
		"pics":        DiffPropertyPICS,
		"role":        DiffPropertyRole,
		"type":        DiffPropertyType,
		"isArray":     DiffPropertyIsArray,
		"constraint":  DiffPropertyConstraint,
		"default":     DiffPropertyDefault,
		"conformance": DiffPropertyConformance,
		"quality":     DiffPropertyQuality,
		"access":      DiffPropertyAccess,
		"bit":         DiffPropertyBit,
	}
)

func (s DiffProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(diffPropertyNames[s])
}

func (s *DiffProperty) UnmarshalJSON(data []byte) (err error) {
	var ss string
	if err := json.Unmarshal(data, &ss); err != nil {
		return err
	}
	var ok bool
	if *s, ok = diffPropertyValues[ss]; !ok {
		return fmt.Errorf("unknown diff property: %s", ss)
	}
	return nil
}

type DiffType uint8

const (
	DiffTypeUnknown DiffType = iota
	DiffTypeMismatch
	DiffTypeMissing
)

var (
	diffTypeNames = map[DiffType]string{
		DiffTypeUnknown:  "unknown",
		DiffTypeMismatch: "mismatch",
		DiffTypeMissing:  "missing",
	}
	diffTypeValues = map[string]DiffType{
		"unknown":  DiffTypeUnknown,
		"mismatch": DiffTypeMismatch,
		"missing":  DiffTypeMissing,
	}
)

func (s DiffType) MarshalJSON() ([]byte, error) {
	return json.Marshal(diffTypeNames[s])
}

func (s *DiffType) UnmarshalJSON(data []byte) (err error) {
	var ss string
	if err := json.Unmarshal(data, &ss); err != nil {
		return err
	}
	var ok bool
	if *s, ok = diffTypeValues[ss]; !ok {
		return fmt.Errorf("unknown diff type: %s", ss)
	}
	return nil
}
