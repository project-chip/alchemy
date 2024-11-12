package pipeline

import "context"

type IndividualProcessor[I, O any] interface {
	Processor
	Process(cxt context.Context, input *Data[I], index int32, total int32) (outputs []*Data[O], extra []*Data[I], err error)
}

type IndividualProcess[I, O any] func(cxt context.Context, input *Data[I], index int32, total int32) (outputs []*Data[O], extra []*Data[I], err error)

type anonymousParallelProcessor[I, O any] struct {
	name    string
	process IndividualProcess[I, O]
}

func (acp *anonymousParallelProcessor[I, O]) Name() string {
	return acp.name
}

func (acp *anonymousParallelProcessor[I, O]) Process(cxt context.Context, input *Data[I], index int32, total int32) (outputs []*Data[O], extra []*Data[I], err error) {
	return acp.process(cxt, input, index, total)
}

func ParallelFunc[I, O any](name string, process IndividualProcess[I, O]) IndividualProcessor[I, O] {
	return &anonymousParallelProcessor[I, O]{process: process}
}
