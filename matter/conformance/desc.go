package conformance

type DescribedConformance struct {
}

func (cc *DescribedConformance) String() string {
	return "desc"
}

func (oc *DescribedConformance) Eval(context ConformanceContext) (ConformanceState, error) {
	return ConformanceStateUnknown, nil
}
