package zapdiff

import (
	"os"
	"strings"
	"testing"

	"github.com/pmezard/go-difflib/difflib"
)

func TestElementTargetedDiff(t *testing.T) {
	refContent := `<configurator>
  <cluster>
    <name>Energy EVSE</name>
    <attribute code="0x0002" side="server" name="FaultState" type="FaultStateEnum" define="FAULT_STATE" min="0x00" max="0xFF">
      <mandatoryConform/>
    </attribute>
  </cluster>
</configurator>`

	genContent := `<configurator>
  <cluster>
    <name>Energy EVSE</name>
    <attribute code="0x0002" side="server" name="FaultState" type="FaultStateEnum" define="FAULT_STATE" min="0x00" max="0xFF">
    </attribute>
  </cluster>
</configurator>`

	refFile, err := os.CreateTemp("", "ref_*.xml")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(refFile.Name())
	refFile.WriteString(refContent)
	refFile.Close()

	genFile, err := os.CreateTemp("", "gen_*.xml")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(genFile.Name())
	genFile.WriteString(genContent)
	genFile.Close()

	id := "configurator/cluster[name='Energy EVSE']/attribute[@name='FaultState']/mandatoryConform"
	parentID := getParentID(id)
	if parentID != "configurator/cluster[name='Energy EVSE']/attribute[@name='FaultState']" {
		t.Fatalf("Expected parent ID to be configurator/cluster[name='Energy EVSE']/attribute[@name='FaultState'], got %s", parentID)
	}

	lines1, err := findElementLines(refFile.Name(), parentID)
	if err != nil {
		t.Fatal(err)
	}
	lines2, err := findElementLines(genFile.Name(), parentID)
	if err != nil {
		t.Fatal(err)
	}

	if len(lines1) == 0 || len(lines2) == 0 {
		t.Fatalf("Expected to find lines for parent element")
	}

	diff := difflib.UnifiedDiff{
		A:        lines1,
		B:        lines2,
		FromFile: "Ref",
		ToFile:   "Generated",
		Context:  3,
	}
	diffStr, err := difflib.GetUnifiedDiffString(diff)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Generated Diff:\n%s", diffStr)

	// Verify that the diff contains the expected changes!
	if !strings.Contains(diffStr, "-      <mandatoryConform/>") {
		t.Errorf("Expected diff to contains removal of mandatoryConform")
	}
}
