package conformance

import "github.com/hasty/matterfmt/matter"

type ProvisionalConformance struct {
}

func (id *ProvisionalConformance) Eval(context matter.ConformanceContext) (matter.ConformanceState, error) {
	return matter.ConformanceStateProvisional, nil
}
