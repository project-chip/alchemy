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

type ConformanceValueStore interface {
	ConformanceReference(id string) HasConformance
}

type ConformanceContext struct {
	Values            map[string]any
	Store             ConformanceValueStore
	VisitedReferences map[string]struct{}
}

type Conformance interface {
	fmt.Stringer

	Eval(context ConformanceContext) (ConformanceState, error)
}

type ConformanceExpression interface {
	fmt.Stringer

	Eval(context ConformanceContext) (bool, error)
}

type HasConformance interface {
	GetConformance() Conformance
}

type HasExpression interface {
	GetExpression() ConformanceExpression
}
