package render

import (
	"slices"
	"testing"

	"github.com/beevik/etree"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
)

func TestMatterNamespaceName(t *testing.T) {
	names := []string{"Common Closure", "Closure", "Common Area Namespace", "Common Location", "Common Relative Position", "Closure Panel"}
	specification := &spec.Specification{}
	for i, name := range names {
		specification.Namespaces = append(specification.Namespaces, &matter.Namespace{ID: matter.NewNumber(uint64(i)), Name: name})
	}

	want := []string{"CommonClosureTag", "ClosureTag", "CommonAreaTag", "CommonLocationTag", "CommonRelativePositionTag", "ClosurePanelTag"}
	for i, ns := range specification.Namespaces {
		if got := matterNamespaceName(specification, ns); got != want[i] {
			t.Errorf("matterNamespaceName(%q) = %q, want %q", ns.Name, got, want[i])
		}
	}
}

func TestGeneratedNamespacesSkipsSkipFileErrataAndSortsByID(t *testing.T) {
	closure := &matter.Namespace{ID: matter.NewNumber(0x44), Name: "Closure"}
	commonClosure := &matter.Namespace{ID: matter.NewNumber(0x01), Name: "Common Closure"}
	location := &matter.Namespace{ID: matter.NewNumber(0x06), Name: "Common Location"}
	skipFile := &errata.Errata{SDK: errata.SDK{SkipFile: true}}
	errataFor := func(ns *matter.Namespace) *errata.Errata {
		if ns == commonClosure {
			return skipFile
		}
		return errata.DefaultErrata
	}

	got := generatedNamespaces([]*matter.Namespace{closure, commonClosure, location}, errataFor)

	if !slices.Equal(got, []*matter.Namespace{location, closure}) {
		var names []string
		for _, ns := range got {
			names = append(names, ns.Name)
		}
		t.Errorf("generatedNamespaces returned %v, want [Common Location Closure]", names)
	}
}

func TestRemoveEnumsReportsEnumsNoNamespaceGenerates(t *testing.T) {
	doc := etree.NewDocument()
	err := doc.ReadFromString(`<configurator><enum name="SemanticTagNamespaceID"/><enum name="ClosureTag"/><enum name="HandTypedTag"/><struct name="Kept"/></configurator>`)
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

func TestSemanticTagNameAppliesErrataOverrideName(t *testing.T) {
	nsErrata := &errata.Errata{SDK: errata.SDK{Types: &errata.SDKTypes{Enums: map[string]*errata.SDKType{
		"CommonAreaTag": {Fields: []*errata.SDKType{{Name: "MediaTvRoom", OverrideName: "Media/TV Room"}}},
	}}}}

	tests := []struct {
		nsErrata *errata.Errata
		enumName string
		tagName  string
		want     string
	}{
		{nsErrata, "CommonAreaTag", "MediaTvRoom", "Media/TV Room"},
		{nsErrata, "CommonAreaTag", "Kitchen", "Kitchen"},
		{nsErrata, "CommonLandmarkTag", "MediaTvRoom", "MediaTvRoom"},
		{errata.DefaultErrata, "CommonAreaTag", "MediaTvRoom", "MediaTvRoom"},
	}
	for _, tt := range tests {
		tag := &matter.SemanticTag{Name: tt.tagName}
		if got := semanticTagName(tt.nsErrata, tt.enumName, tag); got != tt.want {
			t.Errorf("semanticTagName(%s, %s) = %q, want %q", tt.enumName, tt.tagName, got, tt.want)
		}
	}
}
