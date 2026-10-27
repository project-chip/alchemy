package render

import (
	"maps"
	"testing"

	"github.com/beevik/etree"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/types"
	"github.com/project-chip/alchemy/zap"
)

func newQualityTestRenderer(extendedQuality bool) *configuratorRenderer {
	generator := &TemplateGenerator{options: TemplateOptions{ExtendedQuality: extendedQuality}}
	return newConfiguratorRenderer(generator, &zap.Configurator{Errata: &errata.SDK{}})
}

// qualityAttributes returns the attributes of the single quality child of parent, or nil if there is none.
func qualityAttributes(t *testing.T, parent *etree.Element) map[string]string {
	t.Helper()
	qualities := parent.SelectElements("quality")
	if len(qualities) == 0 {
		return nil
	}
	if len(qualities) > 1 {
		t.Fatalf("quality elements = %d, want at most 1", len(qualities))
	}
	attributes := make(map[string]string)
	for _, a := range qualities[0].Attr {
		attributes[a.Key] = a.Value
	}
	return attributes
}

func checkQualityAttributes(t *testing.T, parent *etree.Element, want map[string]string) {
	t.Helper()
	got := qualityAttributes(t, parent)
	if got == nil && want != nil {
		t.Fatalf("quality element missing, want %v", want)
	}
	if got != nil && want == nil {
		t.Fatalf("quality = %v, want no quality element", got)
	}
	if !maps.Equal(got, want) {
		t.Fatalf("quality = %v, want %v", got, want)
	}
}

func TestAttributeQualityPersistence(t *testing.T) {
	for _, test := range []struct {
		name     string
		extended bool
		quality  matter.Quality
		existing string
		want     map[string]string
	}{
		{"fixed", false, matter.QualityFixed, "", map[string]string{"persistence": "fixed"}},
		{"fixed_change_omitted", false, matter.QualityFixed | matter.QualityChangedOmitted, "", map[string]string{"changeOmitted": "true", "persistence": "fixed"}},
		{"fixed_nullable", false, matter.QualityFixed | matter.QualityNullable, "", map[string]string{"persistence": "fixed"}},
		{"change_omitted", false, matter.QualityChangedOmitted, "", map[string]string{"changeOmitted": "true"}},
		{"non_volatile", false, matter.QualityNonVolatile, "", nil},
		{"none", false, matter.QualityNone, "", nil},
		{"existing_change_omitted_now_fixed", false, matter.QualityFixed | matter.QualityChangedOmitted, `<attribute><quality changeOmitted="true"/></attribute>`, map[string]string{"changeOmitted": "true", "persistence": "fixed"}},
		{"existing_fixed_no_longer_fixed", false, matter.QualityChangedOmitted, `<attribute><quality changeOmitted="true" persistence="fixed"/></attribute>`, map[string]string{"changeOmitted": "true"}},
		// Default mode writes persistence only for F, so any other persistence value is removed.
		{"existing_non_volatile", false, matter.QualityChangedOmitted | matter.QualityNonVolatile, `<attribute><quality changeOmitted="true" persistence="nonVolatile"/></attribute>`, map[string]string{"changeOmitted": "true"}},
		{"existing_fixed_quality_removed", false, matter.QualityNone, `<attribute><quality persistence="fixed"/></attribute>`, nil},
		{"extended_fixed", true, matter.QualityFixed, "", map[string]string{"fixed": "true"}},
		{"extended_non_volatile", true, matter.QualityNonVolatile, "", map[string]string{"persistence": "nonVolatile"}},
		{"extended_existing_persistence_fixed", true, matter.QualityFixed, `<attribute><quality persistence="fixed"/></attribute>`, map[string]string{"fixed": "true"}},
	} {
		t.Run(test.name, func(t *testing.T) {
			cluster := matter.NewCluster(nil)
			attribute := matter.NewAttribute(nil, cluster)
			attribute.ID = matter.NewNumber(1)
			attribute.Name = "SyntheticAttribute"
			attribute.Quality = test.quality
			attribute.Conformance = conformance.Set{&conformance.Mandatory{}}
			cluster.Attributes = matter.FieldSet{attribute}
			element := etree.NewElement("attribute")
			if test.existing != "" {
				document := etree.NewDocument()
				if err := document.ReadFromString(test.existing); err != nil {
					t.Fatal(err)
				}
				element = document.Root()
			}
			if err := newQualityTestRenderer(test.extended).populateAttribute(element, attribute, cluster, ""); err != nil {
				t.Fatal(err)
			}
			checkQualityAttributes(t, element, test.want)
		})
	}
}

func TestFieldQualityPersistence(t *testing.T) {
	for _, test := range []struct {
		name       string
		parentType types.EntityType
		fieldType  types.EntityType
	}{
		{"struct_field", types.EntityTypeStruct, types.EntityTypeStructField},
		{"command_field", types.EntityTypeCommand, types.EntityTypeCommandField},
		{"event_field", types.EntityTypeEvent, types.EntityTypeEventField},
	} {
		t.Run(test.name, func(t *testing.T) {
			field := matter.NewField(nil, nil, test.fieldType)
			field.ID = matter.NewNumber(0)
			field.Name = "SyntheticField"
			field.Quality = matter.QualityFixed
			field.Conformance = conformance.Set{&conformance.Mandatory{}}
			element := etree.NewElement("item")
			newQualityTestRenderer(false).setFieldAttributes(element, test.parentType, field, matter.FieldSet{field}, false)
			checkQualityAttributes(t, element, map[string]string{"persistence": "fixed"})
		})
	}
}
