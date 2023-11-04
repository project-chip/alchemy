package conformance

import "testing"

func TestOptional(t *testing.T) {
	conformance, err := ParseConformance("[!AB & (CD != EF)], O")
	if err != nil {

		t.Errorf("failed parsing: %v", err)
	}
	t.Logf("conformance: %s", conformance.String())
}
