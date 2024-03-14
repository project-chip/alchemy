package compare

import (
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/types"
)

type ClusterDifferences struct {
	ID   *matter.Number
	Name string

	Diffs []any `json:"diffs,omitempty"`

	Features   []any `json:"features,omitempty"`
	Bitmaps    []any `json:"bitmaps,omitempty"`
	Enums      []any `json:"enums,omitempty"`
	Structs    []any `json:"structs,omitempty"`
	Attributes []any `json:"attributes,omitempty"`
	Events     []any `json:"events,omitempty"`
	Commands   []any `json:"commands,omitempty"`
}

func compareClusters(specCluster *matter.Cluster, zapCluster *matter.Cluster) (*ClusterDifferences, error) {

	var err error
	cd := &ClusterDifferences{ID: specCluster.ID, Name: specCluster.Name}
	if !namesEqual(specCluster.Name, zapCluster.Name) {
		cd.Diffs = append(cd.Diffs, &StringDiff{Type: DiffTypeMismatch, Property: DiffPropertyName, Spec: specCluster.Name, ZAP: zapCluster.Name})
	}

	var specFeatures, zapFeatures *matter.Bitmap
	if specCluster.Features != nil {
		specFeatures = &specCluster.Features.Bitmap
	}
	if zapCluster.Features != nil {
		zapFeatures = &zapCluster.Features.Bitmap

	}
	cd.Features = compareBitmapsByMask(specFeatures, zapFeatures, types.EntityTypeFeature)
	cd.Attributes, err = compareFields(specCluster.Attributes, zapCluster.Attributes)
	cd.Bitmaps = compareBitmaps(specCluster.Bitmaps, zapCluster.Bitmaps)
	cd.Enums = compareEnums(specCluster.Enums, zapCluster.Enums)
	cd.Structs = compareStructs(specCluster.Structs, zapCluster.Structs)
	cd.Commands = compareCommands(specCluster.Commands, zapCluster.Commands)
	cd.Events = compareEvents(specCluster.Events, zapCluster.Events)
	if err != nil {
		return nil, err
	}
	if len(cd.Diffs) > 0 {
		return cd, nil
	}
	if len(cd.Attributes) > 0 {
		return cd, nil
	}
	if len(cd.Features) > 0 {
		return cd, nil
	}
	if len(cd.Bitmaps) > 0 {
		return cd, nil
	}
	if len(cd.Enums) > 0 {
		return cd, nil
	}
	if len(cd.Commands) > 0 {
		return cd, nil
	}
	if len(cd.Structs) > 0 {
		return cd, nil
	}
	if len(cd.Events) > 0 {
		return cd, nil
	}
	return nil, nil
}
