package matter

type Feature struct {
	Bit         string `json:"bit,omitempty"`
	Code        string `json:"code,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Conformance string `json:"conformance,omitempty"`
}

func (f *Feature) Compare(of *Feature) {
	if f.Bit != of.Bit {

	}
	if f.Name != of.Name {

	}
	if f.Conformance != of.Conformance {

	}
}

type Features []*Feature

func (c Features) compare(oc Features) {
	featureMap := make(map[string]*Feature)
	for _, f := range c {
		featureMap[f.Code] = f
	}

	oFeatureMap := make(map[string]*Feature)
	for _, f := range oc {
		oFeatureMap[f.Code] = f
	}

	for code, of := range oFeatureMap {
		f, ok := featureMap[of.Code]
		if !ok {
			continue
		}
		delete(oFeatureMap, code)
		delete(featureMap, code)
		f.Compare(of)
	}
}
