package matter

type Entity uint8

const (
	EntityUnknown Entity = iota
	EntityCluster
	EntityBitmap
	EntityBitmapValue
	EntityEnum
	EntityEnumValue
	EntityStruct
	EntityAttribute
	EntityCommand
	EntityCommandField
	EntityEvent
	EntityFeature
	EntityDeviceType
	EntityCondition
	EntityField
)

type Model interface {
	Entity() Entity
}
