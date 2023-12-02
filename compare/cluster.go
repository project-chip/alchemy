package compare

import "github.com/hasty/alchemy/matter"

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
	if specCluster.Name != zapCluster.Name {
		cd.Diffs = append(cd.Diffs, &StringDiff{Type: DiffTypeMismatch, Property: DiffPropertyName, Spec: specCluster.Name, ZAP: zapCluster.Name})
	}
	/*if c.Hierarchy != oc.Hierarchy {
		cd.Diffs = append(cd.Diffs, &StringDiff{Type: DiffTypeMismatch, Property: DiffPropertyHierarchy, Spec: c.Hierarchy, ZAP: oc.Hierarchy})
	}
	if c.Role != oc.Role {
		cd.Diffs = append(cd.Diffs, &StringDiff{Type: DiffTypeMismatch, Property: DiffPropertyRole, Spec: c.Role, ZAP: oc.Role})
	}
	if c.PICS != oc.PICS {
		cd.Diffs = append(cd.Diffs, &StringDiff{Type: DiffTypeMismatch, Property: DiffPropertyPICS, Spec: c.PICS, ZAP: oc.PICS})
	}*/

	cd.Features = compareFeatures(specCluster.Features, zapCluster.Features)
	cd.Attributes, err = compareFields(specCluster.Attributes, zapCluster.Attributes)
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
	return nil, nil
}
