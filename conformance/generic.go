package conformance

import "github.com/hasty/alchemy/matter"

type GenericConformance struct {
	raw string
}

func (dc *GenericConformance) RawText() string {
	return dc.raw
}

func (dc *GenericConformance) String() string {
	return dc.raw
}

func (id *GenericConformance) Eval(context matter.ConformanceContext) (matter.ConformanceState, error) {
	return matter.ConformanceStateUnknown, nil
}
