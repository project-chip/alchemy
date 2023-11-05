package conformance

import "github.com/hasty/matterfmt/matter"

type DeprecatedConformance struct {
}

func (id *DeprecatedConformance) Eval(context matter.ConformanceContext) (matter.ConformanceState, error) {
	return matter.ConformanceStateDeprecated, nil
}
