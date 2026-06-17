package idl

import (
	"testing"

	"github.com/project-chip/alchemy/matter"
)

func TestEntityPath(t *testing.T) {
	cluster := &matter.Cluster{
		Name: "DoorLock",
	}
	features := matter.NewFeatures(nil, cluster)
	featureBit := matter.NewFeature(nil, "0", "PinGeneration", "PIN", "Supports PINs", nil)
	features.AddFeatureBit(featureBit)

	t.Run("Features", func(t *testing.T) {
		got := entityPath(features)
		want := "DoorLock.Feature"
		if got != want {
			t.Errorf("entityPath(features) = %q, want %q", got, want)
		}
	})

	t.Run("FeatureBit", func(t *testing.T) {
		got := entityPath(featureBit)
		want := "DoorLock.Feature.PinGeneration"
		if got != want {
			t.Errorf("entityPath(featureBit) = %q, want %q", got, want)
		}
	})
}
