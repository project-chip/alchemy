package matter

type ModelType uint8

const (
	ModelTypeUnknown ModelType = iota
	ModelTypeCluster
	ModelTypeBitmap
	ModelTypeEnum
	ModelTypeStruct
	ModelTypeCommand
	ModelTypeEvent
)

type Model interface {
	ModelType() ModelType
}
