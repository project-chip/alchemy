package conformance

type State uint8

const (
	StateUnknown State = iota
	StateMandatory
	StateOptional
	StateProvisional
	StateDeprecated
	StateDisallowed
)

var StateNames = map[State]string{
	StateUnknown:     "Unknown",
	StateMandatory:   "Mandatory",
	StateOptional:    "Optional",
	StateProvisional: "Provisional",
	StateDeprecated:  "Deprecated",
	StateDisallowed:  "Disallowed",
}

func (cs State) String() string {
	return StateNames[cs]
}

type ConformanceState struct {
	State      State
	Confidence Confidence
}
