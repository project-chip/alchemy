package matter

type Entity uint8

const (
	EntityUnknown Entity = iota
	EntityCluster
	EntityBitmap
	EntityEnum
	EntityStruct
	EntityAttribute
	EntityCommand
	EntityEvent
	EntityFeature
	EntityDeviceType
)

type Model interface {
	ModelType() Entity
}
