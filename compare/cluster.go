package compare

import (
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/spec"
	"github.com/hasty/alchemy/matter/types"
)

type ClusterDifferences struct {
	IdentifiedDiff

	Features   []Diff `json:"features,omitempty"`
	Bitmaps    []Diff `json:"bitmaps,omitempty"`
	Enums      []Diff `json:"enums,omitempty"`
	Structs    []Diff `json:"structs,omitempty"`
	Attributes []Diff `json:"attributes,omitempty"`
	Events     []Diff `json:"events,omitempty"`
	Commands   []Diff `json:"commands,omitempty"`
}

func compareClusters(spec *spec.Specification, specCluster *matter.Cluster, zapCluster *matter.Cluster) (*ClusterDifferences, error) {

	var err error
	cd := &ClusterDifferences{IdentifiedDiff: IdentifiedDiff{ID: specCluster.ID, Name: specCluster.Name, Entity: types.EntityTypeCluster}}
	if !namesEqual(specCluster.Name, zapCluster.Name) {
		cd.Diffs = append(cd.Diffs, &StringDiff{Type: DiffTypeMismatch, Property: DiffPropertyName, Spec: specCluster.Name, ZAP: zapCluster.Name})
	}

	cd.Features = compareFeatures(specCluster.Features, zapCluster.Features)
	cd.Attributes, err = compareFields(types.EntityTypeAttribute, specCluster.Attributes, zapCluster.Attributes)
	cd.Bitmaps = compareBitmaps(specCluster.Bitmaps, zapCluster.Bitmaps)
	cd.Enums = compareEnums(spec, specCluster, zapCluster.Enums)
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
