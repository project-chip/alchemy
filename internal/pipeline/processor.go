package pipeline

import "context"

type ProcessorType int

const (
	ProcessorTypeIndividual ProcessorType = iota // Individual processors can handle one input at a time, and are parallelizable
	ProcessorTypeCollective                      // Collective processors require the entire set of inputs at once
)

type Processor interface {
	Name() string
	Type() ProcessorType
}

type IndividualProcessor[I, O any] interface {
	Processor
	Process(cxt context.Context, input *Data[I], index int32, total int32) (outputs []*Data[O], extra []*Data[I], err error)
}

type IndividualProcess[I, O any] func(cxt context.Context, input *Data[I], index int32, total int32) (outputs []*Data[O], extra []*Data[I], err error)
