package types

import (
	"encoding/json"
)

type EntityType uint8

const (
	EntityTypeUnknown EntityType = iota
	EntityTypeCluster
	EntityTypeBitmap
	EntityTypeBitmapValue
	EntityTypeEnum
	EntityTypeEnumValue
	EntityTypeStruct
	EntityTypeAttribute
	EntityTypeCommand
	EntityTypeCommandField
	EntityTypeEvent
	EntityTypeFeature
	EntityTypeDeviceType
	EntityTypeCondition
	EntityTypeField
	EntityTypeElementRequirement
)

type Entity interface {
	EntityType() EntityType
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
		EntityTypeCommandField:       "field",
		EntityTypeEvent:              "event",
		EntityTypeFeature:            "feature",
		EntityTypeDeviceType:         "deviceType",
		EntityTypeCondition:          "condition",
		EntityTypeField:              "field",
		EntityTypeElementRequirement: "elementRequirement",
	}
)

func (s EntityType) MarshalJSON() ([]byte, error) {
	return json.Marshal(entityTypeNames[s])
}

type EntityStore interface {
	Entities() ([]Entity, error)
}
