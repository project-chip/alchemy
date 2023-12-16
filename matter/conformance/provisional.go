package conformance

type ProvisionalConformance struct {
}

func (dc *ProvisionalConformance) String() string {
	return "provisional"
}

func (id *ProvisionalConformance) Eval(context ConformanceContext) (ConformanceState, error) {
	return ConformanceStateProvisional, nil
}
