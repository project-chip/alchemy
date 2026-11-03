package render

import (
	"testing"

	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
)

func TestMatterNamespaceName(t *testing.T) {
	names := []string{"Common Closure", "Closure", "Common Area", "Common Location", "Common Relative Position", "Closure Panel"}
	specification := &spec.Specification{}
	for i, name := range names {
		specification.Namespaces = append(specification.Namespaces, &matter.Namespace{ID: matter.NewNumber(uint64(i)), Name: name})
	}

	want := []string{"CommonClosureTag", "ClosureTag", "AreaTypeTag", "LocationTag", "RelativePositionTag", "ClosurePanelTag"}
	for i, ns := range specification.Namespaces {
		if got := matterNamespaceName(specification, ns); got != want[i] {
			t.Errorf("matterNamespaceName(%q) = %q, want %q", ns.Name, got, want[i])
		}
	}
}
