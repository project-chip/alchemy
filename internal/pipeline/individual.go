package pipeline

import "context"

type IndividualProcessor[I, O any] interface {
	Processor
	Process(cxt context.Context, input *Data[I], index int32, total int32) (outputs []*Data[O], extra []*Data[I], err error)
}

type IndividualProcess[I, O any] func(cxt context.Context, input *Data[I], index int32, total int32) (outputs []*Data[O], extra []*Data[I], err error)
