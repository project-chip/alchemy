package matter

type ModelType uint8

const (
	ModelTypeUnknown ModelType = iota
	ModelTypeCluster
	ModelTypeBitmap
	ModelTypeEnum
	ModelTypeStruct
	ModelTypeAttribute
	ModelTypeCommand
	ModelTypeEvent
	ModelTypeFeature
	ModelTypeDeviceType
)

type Model interface {
	ModelType() ModelType
}
