package types

import (
	"encoding/json"
	"iter"

	"github.com/project-chip/alchemy/asciidoc"
)

type EntityType uint8

const (
	EntityTypeUnknown EntityType = iota
	EntityTypeCluster
	EntityTypeClusterGroup
	EntityTypeBitmap
	EntityTypeBitmapValue
	EntityTypeEnum
	EntityTypeEnumValue
	EntityTypeStruct
	EntityTypeAttribute
	EntityTypeCommand
	EntityTypeCommandField
	EntityTypeEvent
	EntityTypeEventField
	EntityTypeFeature
	EntityTypeDeviceType
	EntityTypeCondition
	EntityTypeStructField
	EntityTypeElementRequirement
	EntityTypeDeviceTypeRequirement
	EntityTypeClusterRequirement
	EntityTypeConditionRequirement
	EntityTypeDef
	EntityTypeNamespace
	EntityTypeConstant
	EntityTypeSemanticTag
)

type Entity interface {
	EntityType() EntityType
	Source() asciidoc.Element
	Origin() (path string, line int)
	Parent() Entity
	Equals(Entity) bool
}

func (et EntityType) String() string {
	return entityTypeNames[et]
}

var (
	entityTypeNames = map[EntityType]string{
		EntityTypeUnknown:               "unknown",
		EntityTypeCluster:               "cluster",
		EntityTypeBitmap:                "bitmap",
		EntityTypeBitmapValue:           "bit",
		EntityTypeEnum:                  "enum",
		EntityTypeEnumValue:             "enumValue",
		EntityTypeStruct:                "struct",
		EntityTypeAttribute:             "attribute",
		EntityTypeCommand:               "command",
		EntityTypeCommandField:          "commandField",
		EntityTypeEventField:            "eventField",
		EntityTypeEvent:                 "event",
		EntityTypeFeature:               "feature",
		EntityTypeDeviceType:            "deviceType",
		EntityTypeCondition:             "condition",
		EntityTypeStructField:           "structField",
		EntityTypeElementRequirement:    "elementRequirement",
		EntityTypeDeviceTypeRequirement: "deviceTypeRequirement",
		EntityTypeClusterRequirement:    "clusterRequirement",
		EntityTypeConditionRequirement:  "conditionRequirement",
		EntityTypeDef:                   "typeDef",
		EntityTypeNamespace:             "namespace",
		EntityTypeConstant:              "constant",
		EntityTypeSemanticTag:           "semanticTag",
	}
)

func (et EntityType) MarshalJSON() ([]byte, error) {
	return json.Marshal(entityTypeNames[et])
}

type EntityStore interface {
	Entities() ([]Entity, error)
}

func IsDataTypeEntity(entityType EntityType) bool {
	switch entityType {
	case EntityTypeBitmap, EntityTypeEnum, EntityTypeStruct, EntityTypeCommand, EntityTypeEvent, EntityTypeBitmapValue, EntityTypeEnumValue, EntityTypeStructField, EntityTypeCommandField, EntityTypeEventField, EntityTypeAttribute:
		return true
	}
	return false
}

func Filter[T Entity](list []Entity) (out []T) {
	for _, e := range list {
		t, ok := e.(T)
		if ok {
			out = append(out, t)
		}
	}
	return
}

type EntitySet[S any] map[Entity]S

func (es EntitySet[S]) Iterate() iter.Seq[Entity] {
	return func(yield func(Entity) bool) {
		for e := range es {
			if !yield(e) {
				return
			}
		}
	}
}

func FilterSet[T Entity, S any](set EntitySet[S]) (out []T) {
	for e := range set {
		t, ok := e.(T)
		if ok {
			out = append(out, t)
		}
	}
	return
}
