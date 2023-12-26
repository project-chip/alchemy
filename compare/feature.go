package compare

import (
	"strings"

	"github.com/hasty/alchemy/matter"
)

type FeatureDiffs struct {
	ID    string `json:"id"`
	Diffs []any  `json:"diffs,omitempty"`
}

func compareFeature(specFeature *matter.Bit, zapFeature *matter.Bit) (diffs []any) {
	if specFeature.Bit != zapFeature.Bit {
		diffs = append(diffs, &StringDiff{Type: DiffTypeMismatch, Property: DiffPropertyAccess, Spec: specFeature.Bit, ZAP: zapFeature.Bit})
	}
	if specFeature.Name != zapFeature.Name {
		diffs = append(diffs, &StringDiff{Type: DiffTypeMismatch, Property: DiffPropertyName, Spec: specFeature.Name, ZAP: zapFeature.Name})

	}
	if !specFeature.Conformance.Equal(zapFeature.Conformance) {
		diffs = append(diffs, &StringDiff{Type: DiffTypeMismatch, Property: DiffPropertyConformance, Spec: specFeature.Conformance.String(), ZAP: zapFeature.Conformance.String()})
	}
	return
}

func compareFeatures(specFeatures *matter.Bitmap, zapFeatures *matter.Bitmap) (diffs []any) {
	specFeatureMap := make(map[string]*matter.Bit)
	for _, f := range specFeatures.Bits {
		specFeatureMap[strings.ToLower(f.Code)] = f
	}

	zapFeatureMap := make(map[string]*matter.Bit)
	for _, f := range zapFeatures.Bits {
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
		diffs = append(diffs, &MissingDiff{Type: DiffTypeMissing, Code: f.Code, Source: SourceZAP})
	}
	for _, f := range zapFeatureMap {
		diffs = append(diffs, &MissingDiff{Type: DiffTypeMissing, Code: f.Code, Source: SourceSpec})
	}
	return
}
