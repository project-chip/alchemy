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
	EntityTypeDeviceTypeClusterRequirement
	EntityTypeDeviceTypeElementRequirement
	EntityTypeDeviceTypeTagRequirement
	EntityTypeDef
	EntityTypeNamespace
	EntityTypeConstant
	EntityTypeSemanticTag
	EntityTypeTagRequirement
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

func (et EntityType) FriendlyString() string {
	return entityTypeFriendlyNames[et]
}

var (
	entityTypeNames = map[EntityType]string{
		EntityTypeUnknown:                      "unknown",
		EntityTypeCluster:                      "cluster",
		EntityTypeClusterGroup:                 "clusterGroup",
		EntityTypeBitmap:                       "bitmap",
		EntityTypeBitmapValue:                  "bit",
		EntityTypeEnum:                         "enum",
		EntityTypeEnumValue:                    "enumValue",
		EntityTypeStruct:                       "struct",
		EntityTypeAttribute:                    "attribute",
		EntityTypeCommand:                      "command",
		EntityTypeCommandField:                 "commandField",
		EntityTypeEventField:                   "eventField",
		EntityTypeEvent:                        "event",
		EntityTypeFeature:                      "feature",
		EntityTypeDeviceType:                   "deviceType",
		EntityTypeCondition:                    "condition",
		EntityTypeStructField:                  "structField",
		EntityTypeElementRequirement:           "elementRequirement",
		EntityTypeDeviceTypeRequirement:        "deviceTypeRequirement",
		EntityTypeClusterRequirement:           "clusterRequirement",
		EntityTypeConditionRequirement:         "conditionRequirement",
		EntityTypeDef:                          "typeDef",
		EntityTypeNamespace:                    "namespace",
		EntityTypeConstant:                     "constant",
		EntityTypeSemanticTag:                  "semanticTag",
		EntityTypeDeviceTypeClusterRequirement: "deviceTypeClusterRequirement",
		EntityTypeDeviceTypeElementRequirement: "deviceTypeElementRequirement",
		EntityTypeDeviceTypeTagRequirement:     "deviceTypeTagRequirement",
		EntityTypeTagRequirement:               "tagRequirement",
	}

	entityTypeFriendlyNames = map[EntityType]string{
		EntityTypeUnknown:                      "Unknown",
		EntityTypeCluster:                      "Cluster",
		EntityTypeClusterGroup:                 "Cluster Group",
		EntityTypeBitmap:                       "Bitmap",
		EntityTypeBitmapValue:                  "Bit",
		EntityTypeEnum:                         "Enum",
		EntityTypeEnumValue:                    "Enum Value",
		EntityTypeStruct:                       "Struct",
		EntityTypeAttribute:                    "Attribute",
		EntityTypeCommand:                      "Command",
		EntityTypeCommandField:                 "Command Field",
		EntityTypeEventField:                   "Event Field",
		EntityTypeEvent:                        "Event",
		EntityTypeFeature:                      "Feature",
		EntityTypeDeviceType:                   "Device Type",
		EntityTypeCondition:                    "Condition",
		EntityTypeStructField:                  "Struct Field",
		EntityTypeElementRequirement:           "Element Requirement",
		EntityTypeDeviceTypeRequirement:        "Device Type Requirement",
		EntityTypeClusterRequirement:           "Cluster Requirement",
		EntityTypeConditionRequirement:         "Condition Requirement",
		EntityTypeDef:                          "TypeDef",
		EntityTypeNamespace:                    "Namespace",
		EntityTypeConstant:                     "Constant",
		EntityTypeSemanticTag:                  "Semantic Tag",
		EntityTypeDeviceTypeClusterRequirement: "Device Type Cluster Requirement",
		EntityTypeDeviceTypeElementRequirement: "Device Type Element Requirement",
		EntityTypeDeviceTypeTagRequirement:     "Device Type Tag Requirement",
		EntityTypeTagRequirement:               "Tag Requirement",
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

func (es EntitySet[S]) OfType[T Entity]() (out []T) {
	for e := range es {
		t, ok := e.(T)
		if ok {
			out = append(out, t)
		}
	}
	return
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
