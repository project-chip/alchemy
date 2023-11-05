package matter

import "fmt"

type ConformanceState uint8

const (
	ConformanceStateUnknown ConformanceState = iota
	ConformanceStateMandatory
	ConformanceStateOptional
	ConformanceStateProvisional
	ConformanceStateDeprecated
	ConformanceStateDisallowed
)

var ConformanceStateNames = map[ConformanceState]string{
	ConformanceStateUnknown:     "Unknown",
	ConformanceStateMandatory:   "Mandatory",
	ConformanceStateOptional:    "Optional",
	ConformanceStateProvisional: "Provisional",
	ConformanceStateDeprecated:  "Deprecated",
	ConformanceStateDisallowed:  "Disallowed",
}

func (cs ConformanceState) String() string {
	return ConformanceStateNames[cs]
}

type ConformanceContext map[string]any

type Conformance interface {
	fmt.Stringer

	Eval(context ConformanceContext) (ConformanceState, error)
}
