package zapdiff

import (
	"testing"
	"github.com/beevik/etree"
)

func TestParentAndSelfAttrUniqueness(t *testing.T) {
	doc := etree.NewDocument()
	err := doc.ReadFromString(`<configurator>
  <cluster>
    <name>Test</name>
    <attribute/>
    <attribute/>
  </cluster>
</configurator>`)
	if err != nil {
		t.Fatal(err)
	}

	cluster := doc.FindElement("/configurator/cluster")
	if cluster == nil {
		t.Fatal("cluster not found")
	}

	attrs := cluster.SelectElements("attribute")
	if len(attrs) != 2 {
		t.Fatalf("expected 2 attributes, got %d", len(attrs))
	}

	id1 := getEntityUniqueIdentifier(attrs[0])
	id2 := getEntityUniqueIdentifier(attrs[1])

	if id1 == id2 {
		t.Errorf("IDs should be unique, but both are: %s", id1)
	}

	expected1 := "configurator/cluster[name='Test']/attribute[1]"
	expected2 := "configurator/cluster[name='Test']/attribute[2]"

	if id1 != expected1 {
		t.Errorf("Expected ID1 %s, got %s", expected1, id1)
	}
	if id2 != expected2 {
		t.Errorf("Expected ID2 %s, got %s", expected2, id2)
	}
}

func TestParentAndSelfTextUniqueness(t *testing.T) {
	doc := etree.NewDocument()
	err := doc.ReadFromString(`<configurator>
  <cluster>
    <name>Test</name>
    <domain>General</domain>
    <domain></domain>
  </cluster>
</configurator>`)
	if err != nil {
		t.Fatal(err)
	}

	cluster := doc.FindElement("/configurator/cluster")
	if cluster == nil {
		t.Fatal("cluster not found")
	}

	domains := cluster.SelectElements("domain")
	if len(domains) != 2 {
		t.Fatalf("expected 2 domains, got %d", len(domains))
	}

	id1 := getEntityUniqueIdentifier(domains[0])
	id2 := getEntityUniqueIdentifier(domains[1])

	if id1 == id2 {
		t.Errorf("IDs should be unique, but both are: %s", id1)
	}

	expected1 := "configurator/cluster[name='Test']/domain[text()='General']"
	expected2 := "configurator/cluster[name='Test']/domain[2]"

	if id1 != expected1 {
		t.Errorf("Expected ID1 %s, got %s", expected1, id1)
	}
	if id2 != expected2 {
		t.Errorf("Expected ID2 %s, got %s", expected2, id2)
	}
}
