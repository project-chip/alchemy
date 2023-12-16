package conformance

type GenericConformance struct {
	raw string
}

func (dc *GenericConformance) RawText() string {
	return dc.raw
}

func (dc *GenericConformance) String() string {
	return dc.raw
}

func (id *GenericConformance) Eval(context ConformanceContext) (ConformanceState, error) {
	return ConformanceStateUnknown, nil
}
