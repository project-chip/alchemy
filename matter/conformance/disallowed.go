package conformance

type DisallowedConformance struct {
	raw string
}

func (dc *DisallowedConformance) RawText() string {
	return dc.raw
}

func (dc *DisallowedConformance) String() string {
	return "disallowed"
}

func (id *DisallowedConformance) Eval(context ConformanceContext) (ConformanceState, error) {
	return ConformanceStateDisallowed, nil
}
