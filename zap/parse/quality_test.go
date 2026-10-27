package parse

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/project-chip/alchemy/matter"
)

func parseQualityString(t *testing.T, s string) matter.Quality {
	t.Helper()
	d := xml.NewDecoder(strings.NewReader(s))
	for {
		tok, err := d.Token()
		if err != nil {
			t.Fatal(err)
		}
		if e, ok := tok.(xml.StartElement); ok {
			q, err := parseQuality(d, e)
			if err != nil {
				t.Fatal(err)
			}
			return q
		}
	}
}

func TestParseQualityPersistence(t *testing.T) {
	for _, test := range []struct {
		name string
		xml  string
		want matter.Quality
	}{
		{"persistence_fixed", `<quality persistence="fixed"/>`, matter.QualityFixed},
		{"fixed_attribute", `<quality fixed="true"/>`, matter.QualityFixed},
		{"persistence_non_volatile", `<quality persistence="nonVolatile"/>`, matter.QualityNonVolatile},
		{"change_omitted_persistence_fixed", `<quality changeOmitted="true" persistence="fixed"/>`, matter.QualityChangedOmitted | matter.QualityFixed},
	} {
		t.Run(test.name, func(t *testing.T) {
			if got := parseQualityString(t, test.xml); got != test.want {
				t.Fatalf("quality = %v, want %v", got, test.want)
			}
		})
	}
}
