package matter

import "github.com/hasty/alchemy/matter/conformance"

type Feature struct {
	Bit         string          `json:"bit,omitempty"`
	Code        string          `json:"code,omitempty"`
	Name        string          `json:"name,omitempty"`
	Summary     string          `json:"summary,omitempty"`
	Conformance conformance.Set `json:"conformance,omitempty"`
}

func (f *Feature) Entity() Entity {
	return EntityFeature
}

func (f *Feature) GetConformance() conformance.Set {
	return f.Conformance
}

type FeatureSet []*Feature

func (fs FeatureSet) Reference(id string) conformance.HasConformance {
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
