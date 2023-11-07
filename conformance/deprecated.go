package conformance

import "github.com/hasty/alchemy/matter"

type DeprecatedConformance struct {
	raw string
}

func (dc *DeprecatedConformance) RawText() string {
	return dc.raw
}

func (dc *DeprecatedConformance) Eval(context matter.ConformanceContext) (matter.ConformanceState, error) {
	return matter.ConformanceStateDeprecated, nil
}
