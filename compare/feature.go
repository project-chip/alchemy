package compare

import (
	"strings"

	"github.com/hasty/matterfmt/matter"
)

type FeatureDiffs struct {
	ID    string `json:"id"`
	Diffs []any  `json:"diffs,omitempty"`
}

func compareFeature(specFeature *matter.Feature, zapFeature *matter.Feature) (diffs []any) {
	if specFeature.Bit != zapFeature.Bit {
		diffs = append(diffs, &StringDiff{Type: DiffTypeMismatch, Property: DiffPropertyAccess, Spec: specFeature.Bit, ZAP: zapFeature.Bit})
	}
	if specFeature.Name != zapFeature.Name {
		diffs = append(diffs, &StringDiff{Type: DiffTypeMismatch, Property: DiffPropertyName, Spec: specFeature.Name, ZAP: zapFeature.Name})

	}
	if specFeature.Conformance != zapFeature.Conformance {
		diffs = append(diffs, &StringDiff{Type: DiffTypeMismatch, Property: DiffPropertyConformance, Spec: specFeature.Conformance, ZAP: zapFeature.Conformance})
	}
	return
}

func compareFeatures(specFeatures []*matter.Feature, zapFeatures []*matter.Feature) (diffs []any) {
	specFeatureMap := make(map[string]*matter.Feature)
	for _, f := range specFeatures {
		specFeatureMap[strings.ToLower(f.Code)] = f
	}

	zapFeatureMap := make(map[string]*matter.Feature)
	for _, f := range zapFeatures {
		zapFeatureMap[strings.ToLower(f.Code)] = f
	}

	for code, zapFeature := range zapFeatureMap {
		specFeature, ok := specFeatureMap[code]
		if !ok {
			continue
		}
		delete(zapFeatureMap, code)
		delete(specFeatureMap, code)
		featureDiffs := compareFeature(specFeature, zapFeature)
		if len(featureDiffs) > 0 {
			diffs = append(diffs, &FeatureDiffs{ID: code, Diffs: featureDiffs})
		}
	}
	for _, f := range specFeatureMap {
		diffs = append(diffs, &MissingDiff{Type: DiffTypeMissing, ID: f.Code, Source: SourceZAP})
	}
	for _, f := range zapFeatureMap {
		diffs = append(diffs, &MissingDiff{Type: DiffTypeMissing, ID: f.Code, Source: SourceSpec})
	}
	return
}
