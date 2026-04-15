package zapdiff

import (
	"os"
	"strings"
	"testing"
)

func TestParseEntityUniqueIdentifierWithText(t *testing.T) {
	id := "configurator/cluster[name='Test']/domain[text()='General']"
	segments, err := parseEntityUniqueIdentifier(id)
	if err != nil {
		t.Fatal(err)
	}

	if len(segments) != 3 {
		t.Fatalf("expected 3 segments, got %d", len(segments))
	}

	if segments[2].tag != "domain" {
		t.Errorf("expected tag domain, got %s", segments[2].tag)
	}
	if segments[2].attr != "text()" {
		t.Errorf("expected attr text(), got %s", segments[2].attr)
	}
	if segments[2].value != "General" {
		t.Errorf("expected value General, got %s", segments[2].value)
	}
	if segments[2].isAttr {
		t.Errorf("expected isAttr false, got true")
	}
}

func TestFindElementLinesWithText(t *testing.T) {
	content := `<configurator>
  <cluster>
    <name>Test</name>
    <domain>General</domain>
    <domain>Special</domain>
  </cluster>
</configurator>`

	tmpFile, err := os.CreateTemp("", "zapdiff_test_*.xml")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.WriteString(content); err != nil {
		t.Fatal(err)
	}
	tmpFile.Close()

	id := "configurator/cluster[name='Test']/domain[text()='Special']"
	lines, startLine, err := findElementLines(tmpFile.Name(), id)
	if err != nil {
		t.Fatal(err)
	}

	if len(lines) != 1 {
		t.Fatalf("expected 1 line, got %d", len(lines))
	}

	if startLine != 5 { // 1-indexed, line 5 is "<domain>Special</domain>"
		t.Errorf("expected startLine 5, got %d", startLine)
	}

	if !strings.Contains(lines[0], "<domain>Special</domain>") {
		t.Errorf("expected line to contain <domain>Special</domain>, got %s", lines[0])
	}
}

func TestFindElementLinesWithChildAttributes(t *testing.T) {
	content := `<configurator>
  <cluster>
    <name>Test</name>
    <domain attr="val">Special</domain>
  </cluster>
</configurator>`

	tmpFile, err := os.CreateTemp("", "zapdiff_test_*.xml")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.WriteString(content); err != nil {
		t.Fatal(err)
	}
	tmpFile.Close()

	id := "configurator/cluster[domain='Special']"
	lines, startLine, err := findElementLines(tmpFile.Name(), id)
	if err != nil {
		t.Fatal(err)
	}

	if len(lines) != 4 {
		t.Fatalf("expected 4 lines, got %d", len(lines))
	}

	if startLine != 2 {
		t.Errorf("expected startLine 2, got %d", startLine)
	}
}

func TestFindElementLinesWithMultiLineSelfClosing(t *testing.T) {
	content := `<configurator>
  <cluster name="Test"
    code="0x1234"
  />
</configurator>`

	tmpFile, err := os.CreateTemp("", "zapdiff_test_*.xml")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.WriteString(content); err != nil {
		t.Fatal(err)
	}
	tmpFile.Close()

	id := "configurator/cluster[@name='Test']"
	lines, startLine, err := findElementLines(tmpFile.Name(), id)
	if err != nil {
		t.Fatal(err)
	}

	if len(lines) != 3 {
		t.Fatalf("expected 3 lines, got %d", len(lines))
	}

	if startLine != 2 {
		t.Errorf("expected startLine 2, got %d", startLine)
	}
}

func TestGenerateUnifiedDiff(t *testing.T) {
	lines1 := []string{"line1\n", "line2\n", "line3\n"}
	lines2 := []string{"line1\n", "line2 modified\n", "line3\n"}

	diffStr, err := GenerateUnifiedDiff(lines1, lines2, 10, 20)
	if err != nil {
		t.Fatal(err)
	}

	expected := "--- Ref\n+++ Generated\n@@ -10,3 +20,3 @@\n line1\n-line2\n+line2 modified\n line3\n"
	if diffStr != expected {
		t.Errorf("expected diff:\n%s\ngot:\n%s", expected, diffStr)
	}
}

func TestGenerateUnifiedDiffGrouped(t *testing.T) {
	lines1 := []string{"A\n", "B\n", "C\n", "D\n"}
	lines2 := []string{"A\n", "X\n", "Y\n", "D\n"}

	diffStr, err := GenerateUnifiedDiff(lines1, lines2, 10, 20)
	if err != nil {
		t.Fatal(err)
	}

	expected := "--- Ref\n+++ Generated\n@@ -10,4 +20,4 @@\n A\n-B\n-C\n+X\n+Y\n D\n"
	if diffStr != expected {
		t.Errorf("expected diff:\n%s\ngot:\n%s", expected, diffStr)
	}
}

func TestGetCustomDiffLines(t *testing.T) {
	content := `<configurator>
  <cluster code="0x0028"/>
  <attribute code="0x0017" side="server"/>
</configurator>`

	tmpFile, err := os.CreateTemp("", "test_diff_*.xml")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.WriteString(content); err != nil {
		t.Fatal(err)
	}
	tmpFile.Close()

	// We need to mock getParentID or use a targetID that works with it.
	// getParentID removes the last segment.
	// If targetID is "configurator/cluster[@code='0x0028']/attribute[@code='0x0017']"
	// parentID will be "configurator/cluster[@code='0x0028']"
	// This matches the structure in our content (textually for findElementLines).

	targetID := "configurator/cluster[@code='0x0028']/attribute[@code='0x0017']"
	
	lines, _, err := getCustomDiffLines(tmpFile.Name(), targetID)
	if err != nil {
		t.Fatal(err)
	}

	// Expected lines:
	// 1. Parent start: <cluster code="0x0028"/>
	// 2. Target: <attribute code="0x0017" side="server"/>
	// Parent close should NOT be appended because it was self-closing (len=1).

	if len(lines) != 2 {
		t.Fatalf("expected 2 lines, got %d: %v", len(lines), lines)
	}

	if !strings.Contains(lines[0], "<cluster") {
		t.Errorf("expected cluster line, got %s", lines[0])
	}
	if !strings.Contains(lines[1], "<attribute") {
		t.Errorf("expected attribute line, got %s", lines[1])
	}
}
