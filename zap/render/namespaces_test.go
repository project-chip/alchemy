package render

import (
	"slices"
	"testing"

	"github.com/beevik/etree"
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

func TestRemoveEnumsReportsEnumsNoNamespaceGenerates(t *testing.T) {
	doc := etree.NewDocument()
	err := doc.ReadFromString(`<configurator><enum name="SemanticTagNamespaceId"/><enum name="ClosureTag"/><enum name="HandTypedTag"/><struct name="Kept"/></configurator>`)
	if err != nil {
		t.Fatal(err)
	}
	configurator := doc.SelectElement("configurator")

	unknown := removeEnums(configurator, map[string]*matter.Namespace{"ClosureTag": {}})

	if !slices.Equal(unknown, []string{"HandTypedTag"}) {
		t.Errorf("removeEnums reported %v, want [HandTypedTag]", unknown)
	}
	if n := len(configurator.SelectElements("enum")); n != 0 {
		t.Errorf("%d enums left after removeEnums, want 0", n)
	}
	if configurator.SelectElement("struct") == nil {
		t.Error("removeEnums removed a non-enum element")
	}
}
