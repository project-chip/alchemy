package zap

import (
	"testing"

	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
)

func newTestCluster(name string, id *matter.Number) *matter.Cluster {
	c := matter.NewCluster(nil)
	c.Name = name
	c.ID = id
	return c
}

func numbersToStrings(numbers []*matter.Number) (s []string) {
	for _, n := range numbers {
		s = append(s, n.HexString())
	}
	return
}

func checkClusterCodes(t *testing.T, name string, got []*matter.Number, want ...string) {
	t.Helper()
	gs := numbersToStrings(got)
	if len(gs) != len(want) {
		t.Fatalf("%s cluster codes = %v, want %v", name, gs, want)
	}
	for i := range want {
		if gs[i] != want[i] {
			t.Fatalf("%s cluster codes = %v, want %v", name, gs, want)
		}
	}
}

// Mirrors Mode Base: a base cluster without an ID, skipped via errata, whose structs are shared with derived clusters
func TestSkippedClusterRendersSharedTypes(t *testing.T) {
	s := &spec.Specification{ClusterRefs: spec.NewEntityRefs[*matter.Cluster]()}

	base := newTestCluster("Mode Base", matter.InvalidID)
	sharedStruct := matter.NewStruct(nil, base)
	sharedStruct.Name = "ModeOptionStruct"
	sharedEnum := matter.NewEnum(nil, base)
	sharedEnum.Name = "SharedEnum"
	sharedBitmap := matter.NewBitmap(nil, base)
	sharedBitmap.Name = "SharedBitmap"
	baseOnlyEnum := matter.NewEnum(nil, base)
	baseOnlyEnum.Name = "StatusCodeEnum"
	base.Structs = append(base.Structs, sharedStruct)
	base.Enums = append(base.Enums, sharedEnum, baseOnlyEnum)
	base.Bitmaps = append(base.Bitmaps, sharedBitmap)

	derivedA := newTestCluster("RVC Run Mode", matter.NewNumber(0x0054))
	derivedB := newTestCluster("Laundry Washer Mode", matter.NewNumber(0x0051))
	noIDCluster := newTestCluster("Cluster Without ID", matter.InvalidID)

	for _, e := range []types.Entity{sharedStruct, sharedEnum, sharedBitmap, baseOnlyEnum} {
		s.ClusterRefs.Add(base, e)
	}
	for _, e := range []types.Entity{sharedStruct, sharedEnum, sharedBitmap} {
		s.ClusterRefs.Add(derivedA, e)
		s.ClusterRefs.Add(derivedB, e)
	}
	// References from clusters without a valid ID can't produce a cluster code
	s.ClusterRefs.Add(noIDCluster, baseOnlyEnum)

	sdkErrata := &errata.SDK{ClusterSkip: []string{"Mode Base"}}
	c, err := NewConfigurator(s, nil, []types.Entity{base}, "mode-base-cluster.xml", sdkErrata, false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if c.IsEmpty() {
		t.Fatalf("configurator for skipped cluster with shared types should not be empty")
	}
	if len(c.Clusters) != 0 || len(c.ClusterIDs) != 0 || c.Features != nil {
		t.Fatalf("skipped cluster should not be rendered: clusters=%d clusterIDs=%v features=%v", len(c.Clusters), c.ClusterIDs, c.Features)
	}

	if len(c.Structs) != 1 {
		t.Fatalf("structs = %d, want 1", len(c.Structs))
	}
	checkClusterCodes(t, "struct", c.Structs[sharedStruct], "0x0051", "0x0054")

	if len(c.Enums) != 1 {
		t.Fatalf("enums = %d, want 1 (base-only enum must not be added)", len(c.Enums))
	}
	checkClusterCodes(t, "enum", c.Enums[sharedEnum], "0x0051", "0x0054")

	if len(c.Bitmaps) != 1 {
		t.Fatalf("bitmaps = %d, want 1", len(c.Bitmaps))
	}
	checkClusterCodes(t, "bitmap", c.Bitmaps[sharedBitmap], "0x0051", "0x0054")
}

// Mirrors Label: a skipped cluster with no shared types should still produce an empty configurator
func TestSkippedClusterWithoutSharedTypesIsEmpty(t *testing.T) {
	s := &spec.Specification{ClusterRefs: spec.NewEntityRefs[*matter.Cluster]()}

	base := newTestCluster("Label", matter.InvalidID)
	ownStruct := matter.NewStruct(nil, base)
	ownStruct.Name = "OwnStruct"
	base.Structs = append(base.Structs, ownStruct)
	s.ClusterRefs.Add(base, ownStruct)

	sdkErrata := &errata.SDK{ClusterSkip: []string{"Label"}}
	c, err := NewConfigurator(s, nil, []types.Entity{base}, "label-cluster.xml", sdkErrata, false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !c.IsEmpty() {
		t.Fatalf("configurator should be empty, got structs=%d enums=%d bitmaps=%d", len(c.Structs), len(c.Enums), len(c.Bitmaps))
	}
}
