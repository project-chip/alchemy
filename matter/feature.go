package matter

type Feature struct {
	Bit         string      `json:"bit,omitempty"`
	Code        string      `json:"code,omitempty"`
	Name        string      `json:"name,omitempty"`
	Summary     string      `json:"summary,omitempty"`
	Conformance Conformance `json:"conformance,omitempty"`
}

func (f *Feature) Entity() Entity {
	return EntityFeature
}

func (f *Feature) GetConformance() Conformance {
	return f.Conformance
}

type FeatureSet []*Feature

func (fs FeatureSet) ConformanceReference(id string) HasConformance {
	if len(fs) == 0 {
		return nil
	}
	for _, f := range fs {
		if f.Code == id {
			return f
		}
	}
	return nil
}
