package compare

import (
	"encoding/json"
	"fmt"

	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/conformance"
	"github.com/hasty/alchemy/matter/constraint"
	"github.com/hasty/alchemy/matter/types"
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
	DiffPropertyNullable
	DiffPropertyReadAccess
	DiffPropertyWriteAccess
	DiffPropertyInvokeAccess
	DiffPropertyOptionalWrite
	DiffPropertyFabricScoping
	DiffPropertyFabricSensitivity
	DiffPropertyTiming
	DiffPropertyBit
	DiffPropertyValue
	DiffPropertyCommandResponse
	DiffPropertyCommandDirection
	DiffPropertyPriority
	DiffPropertyLength
	DiffPropertyMinLength
	DiffPropertyMax
	DiffPropertyMin
)

var (
	diffPropertyNames = map[DiffProperty]string{
		DiffPropertyUnknown:           "unknown",
		DiffPropertyName:              "name",
		DiffPropertyHierarchy:         "hierarchy",
		DiffPropertyPICS:              "pics",
		DiffPropertyRole:              "role",
		DiffPropertyType:              "type",
		DiffPropertyIsArray:           "isArray",
		DiffPropertyConstraint:        "constraint",
		DiffPropertyDefault:           "default",
		DiffPropertyConformance:       "conformance",
		DiffPropertyNullable:          "nullable",
		DiffPropertyReadAccess:        "readAccess",
		DiffPropertyWriteAccess:       "writeAccess",
		DiffPropertyInvokeAccess:      "invokeAccess",
		DiffPropertyOptionalWrite:     "optionalWrite",
		DiffPropertyBit:               "bit",
		DiffPropertyValue:             "value",
		DiffPropertyFabricScoping:     "fabricScoping",
		DiffPropertyFabricSensitivity: "fabricSensitivity",
		DiffPropertyTiming:            "timing",
		DiffPropertyCommandResponse:   "commandResponse",
		DiffPropertyCommandDirection:  "commandDirection",
		DiffPropertyPriority:          "priority",
		DiffPropertyLength:            "length",
		DiffPropertyMinLength:         "minLength",
		DiffPropertyMax:               "max",
		DiffPropertyMin:               "min",
	}
	diffPropertyValues = map[string]DiffProperty{
		"unknown":           DiffPropertyUnknown,
		"name":              DiffPropertyName,
		"hierarchy":         DiffPropertyHierarchy,
		"pics":              DiffPropertyPICS,
		"role":              DiffPropertyRole,
		"type":              DiffPropertyType,
		"isArray":           DiffPropertyIsArray,
		"constraint":        DiffPropertyConstraint,
		"default":           DiffPropertyDefault,
		"conformance":       DiffPropertyConformance,
		"nullable":          DiffPropertyNullable,
		"readAccess":        DiffPropertyReadAccess,
		"writeAccess":       DiffPropertyWriteAccess,
		"invokeAccess":      DiffPropertyInvokeAccess,
		"optionalWrite":     DiffPropertyOptionalWrite,
		"bit":               DiffPropertyBit,
		"value":             DiffPropertyValue,
		"fabricScoping":     DiffPropertyFabricScoping,
		"fabricSensitivity": DiffPropertyFabricSensitivity,
		"timing":            DiffPropertyTiming,
		"commandResponse":   DiffPropertyCommandResponse,
		"commandDirection":  DiffPropertyCommandDirection,
		"priority":          DiffPropertyPriority,
		"length":            DiffPropertyLength,
		"minLength":         DiffPropertyMinLength,
		"max":               DiffPropertyMax,
		"min":               DiffPropertyMin,
	}
)

func (s DiffProperty) String() string {
	return diffPropertyNames[s]
}

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

type IdentifiedDiff struct {
	Type   DiffType         `json:"type,omitempty"`
	Entity types.EntityType `json:"entity,omitempty"`
	ID     *matter.Number   `json:"id,omitempty"`
	Name   string           `json:"name,omitempty"`
	Diffs  []Diff           `json:"diffs,omitempty"`
}

func (d IdentifiedDiff) String() string {
	return "identified"
}

type MissingDiff struct {
	Type     DiffType         `json:"type"`
	Entity   types.EntityType `json:"entity,omitempty"`
	Source   Source           `json:"source,omitempty"`
	Property DiffProperty     `json:"property,omitempty"`
	ID       *matter.Number   `json:"id,omitempty"`
	Name     string           `json:"name,omitempty"`
	Code     string           `json:"code,omitempty"`
}

func (md MissingDiff) String() string {
	return "missing"
}

func newMissingDiff(name string, props ...any) *MissingDiff {
	diff := &MissingDiff{Type: DiffTypeMissing, Name: name}
	for _, prop := range props {
		switch prop := prop.(type) {
		case string:
			diff.Code = prop
		case types.EntityType:
			diff.Entity = prop
		case Source:
			diff.Source = prop
		case *matter.Number:
			diff.ID = prop
		case DiffProperty:
			diff.Property = prop
		}
	}
	return diff
}

type StringDiff struct {
	Type     DiffType     `json:"type"`
	Property DiffProperty `json:"property"`
	Spec     string       `json:"spec"`
	ZAP      string       `json:"zap"`
}

func (d StringDiff) String() string {
	return "string"
}

type BoolDiff struct {
	Type     DiffType     `json:"type"`
	Property DiffProperty `json:"property"`
	Spec     bool         `json:"spec"`
	ZAP      bool         `json:"zap"`
}

func (d BoolDiff) String() string {
	return "bool"
}

type ConformanceDiff struct {
	Type            DiffType          `json:"type"`
	Property        DiffProperty      `json:"property"`
	Spec            conformance.State `json:"spec"`
	ZAP             conformance.State `json:"zap"`
	SpecConfornance conformance.Set   `json:"specConformance"`
}

func (d ConformanceDiff) String() string {
	return "conformance"
}

type ConstraintDiff struct {
	Type     DiffType              `json:"type"`
	Property DiffProperty          `json:"property"`
	Spec     constraint.Constraint `json:"spec"`
	ZAP      constraint.Constraint `json:"zap"`
}

func (d ConstraintDiff) String() string {
	return "constraint"
}

type QualityDiff struct {
	Type     DiffType       `json:"type"`
	Property DiffProperty   `json:"property"`
	Spec     matter.Quality `json:"spec"`
	ZAP      matter.Quality `json:"zap"`
}

func (d QualityDiff) String() string {
	return "quality"
}

type Diff interface {
	String() string
}

type PropertyDiff[T ~uint8] struct {
	Type     DiffType     `json:"type"`
	Property DiffProperty `json:"property"`
	Spec     T            `json:"spec"`
	ZAP      T            `json:"zap"`
}

func (d PropertyDiff[T]) String() string {
	return diffPropertyNames[d.Property]
}

func NewPropertyDiff[T ~uint8](diffType DiffType, property DiffProperty, spec T, zap T) *PropertyDiff[T] {
	return &PropertyDiff[T]{
		Type:     diffType,
		Property: property,
		Spec:     spec,
		ZAP:      zap,
	}
}
