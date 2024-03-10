package pipeline

import "context"

type ProcessorType int

const (
	ProcessorTypeParallel ProcessorType = iota
	ProcessorTypeSerial                 = iota
)

type Processor[I, O any] interface {
	Name() string
	Type() ProcessorType
	Process(cxt context.Context, input *Data[I], index int32, total int32) (outputs []*Data[O], extra []*Data[I], err error)
	ProcessAll(cxt context.Context, inputs []*Data[I]) (outputs []*Data[O], err error)
}
