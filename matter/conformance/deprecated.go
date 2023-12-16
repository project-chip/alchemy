package conformance

type DeprecatedConformance struct {
	raw string
}

func (dc *DeprecatedConformance) RawText() string {
	return dc.raw
}

func (dc *DeprecatedConformance) String() string {
	return "deprecated"
}

func (dc *DeprecatedConformance) Eval(context ConformanceContext) (ConformanceState, error) {
	return ConformanceStateDeprecated, nil
}
