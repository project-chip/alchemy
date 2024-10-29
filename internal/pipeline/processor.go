package pipeline

type ProcessorType int

const (
	ProcessorTypeIndividual ProcessorType = iota // Individual processors can handle one input at a time, and are parallelizable
	ProcessorTypeCollective                      // Collective processors require the entire set of inputs at once
)

type Processor interface {
	Name() string
	Type() ProcessorType
}
