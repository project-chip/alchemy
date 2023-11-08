package conformance

import "github.com/hasty/alchemy/matter"

type DisallowedConformance struct {
	raw string
}

func (dc *DisallowedConformance) RawText() string {
	return dc.raw
}

func (dc *DisallowedConformance) String() string {
	return "disallowed"
}

func (id *DisallowedConformance) Eval(context matter.ConformanceContext) (matter.ConformanceState, error) {
	return matter.ConformanceStateDisallowed, nil
}
