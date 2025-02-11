package types

import (
	"encoding/json"

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
	EntityTypeDef
	EntityTypeNamespace
	EntityTypeConstant
)

type Entity interface {
	EntityType() EntityType
	Source() asciidoc.Element
	Parent() Entity
}

func (et EntityType) String() string {
	return entityTypeNames[et]
}

var (
	entityTypeNames = map[EntityType]string{
		EntityTypeUnknown:            "unknown",
		EntityTypeCluster:            "cluster",
		EntityTypeBitmap:             "bitmap",
		EntityTypeBitmapValue:        "bit",
		EntityTypeEnum:               "enum",
		EntityTypeEnumValue:          "enumValue",
		EntityTypeStruct:             "struct",
		EntityTypeAttribute:          "attribute",
		EntityTypeCommand:            "command",
		EntityTypeCommandField:       "commandField",
		EntityTypeEventField:         "eventField",
		EntityTypeEvent:              "event",
		EntityTypeFeature:            "feature",
		EntityTypeDeviceType:         "deviceType",
		EntityTypeCondition:          "condition",
		EntityTypeStructField:        "structField",
		EntityTypeElementRequirement: "elementRequirement",
		EntityTypeDef:                "typeDef",
		EntityTypeNamespace:          "namespace",
		EntityTypeConstant:           "constant",
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
