package conformance

import "github.com/hasty/matterfmt/matter"

type DisallowedConformance struct {
	raw string
}

func (dc *DisallowedConformance) RawText() string {
	return dc.raw
}

func (id *DisallowedConformance) Eval(context matter.ConformanceContext) (matter.ConformanceState, error) {
	return matter.ConformanceStateDisallowed, nil
}
