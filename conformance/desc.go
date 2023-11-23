package conformance

import (
	"github.com/hasty/alchemy/matter"
)

type DescribedConformance struct {
}

func (cc *DescribedConformance) String() string {
	return "desc"
}

func (oc *DescribedConformance) Eval(context matter.ConformanceContext) (matter.ConformanceState, error) {
	return matter.ConformanceStateUnknown, nil
}
