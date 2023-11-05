package conformance

import "github.com/hasty/matterfmt/matter"

type DisallowedConformance struct {
}

func (id *DisallowedConformance) Eval(context matter.ConformanceContext) (matter.ConformanceState, error) {
	return matter.ConformanceStateDisallowed, nil
}
