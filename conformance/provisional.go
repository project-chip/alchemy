package conformance

import "github.com/hasty/alchemy/matter"

type ProvisionalConformance struct {
}

func (dc *ProvisionalConformance) String() string {
	return "provisional"
}

func (id *ProvisionalConformance) Eval(context matter.ConformanceContext) (matter.ConformanceState, error) {
	return matter.ConformanceStateProvisional, nil
}
