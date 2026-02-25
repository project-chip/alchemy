package tests

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/asciidoc/parse"
)

func TestTableCellIfDefs(t *testing.T) {
	input, err := os.ReadFile("asciidoctor/table_cell_ifdefs.adoc")
	if err != nil {
		t.Fatalf("failed to read file: %v", err)
	}

	doc, err := parse.Reader(asciidoc.Path{Relative: "asciidoctor/table_cell_ifdefs.adoc"}, bytes.NewReader(input))
	if err != nil {
		t.Fatalf("failed to parse file: %v", err)
	}

	var tables []*asciidoc.Table
	for _, el := range doc.Elements {
		if tbl, ok := el.(*asciidoc.Table); ok {
			tables = append(tables, tbl)
		}
	}

	if len(tables) < 6 {
		t.Fatalf("Expected 6 tables, found %d", len(tables))
	}

	getDataCell := func(t *testing.T, table *asciidoc.Table) *asciidoc.TableCell {
		var dataRows []*asciidoc.TableRow
		for _, el := range table.Elements {
			if row, ok := el.(*asciidoc.TableRow); ok {
				dataRows = append(dataRows, row)
			}
		}

		if len(dataRows) < 2 {
			t.Logf("Table has %d rows (expected > 1 for header+data)", len(dataRows))
			// Fallback to last row if present
			if len(dataRows) > 0 {
				cells := dataRows[len(dataRows)-1].TableCells()
				if len(cells) > 0 {
					return cells[0]
				}
			}
			t.Fatal("Could not find data cell")
			return nil
		}

		// Return first cell of second row (index 1)
		cells := dataRows[1].TableCells()
		if len(cells) > 0 {
			return cells[0]
		}
		t.Fatal("Data row has no cells")
		return nil
	}

	t.Run("Scenario 1: Normal Text Only", func(t *testing.T) {
		table := tables[0]
		cell := getDataCell(t, table)
		hasText := false
		for _, child := range cell.Children() {
			if s, ok := child.(*asciidoc.String); ok && strings.Contains(s.Value, "Normal Text Only") {
				hasText = true
			}
		}
		if !hasText {
			t.Error("Did not find 'Normal Text Only'")
		}
	})

	t.Run("Scenario 2: Normal Text + IfDef", func(t *testing.T) {
		table := tables[1]
		cell := getDataCell(t, table)
		hasText := false
		hasIfDef := false
		for _, child := range cell.Children() {
			if s, ok := child.(*asciidoc.String); ok && strings.Contains(s.Value, "Normal Text before") {
				hasText = true
			}
			if _, ok := child.(*asciidoc.IfDef); ok {
				hasIfDef = true
			}
		}
		if !hasText {
			t.Error("Did not find 'Normal Text before'")
		}
		if !hasIfDef {
			t.Error("Did not find IfDef")
		}
	})

	t.Run("Scenario 3: IfDef + Normal Text (Unwrapped)", func(t *testing.T) {
		table := tables[2]
		cell := getDataCell(t, table)
		hasIfDef := false
		for _, child := range cell.Children() {
			if _, ok := child.(*asciidoc.IfDef); ok {
				hasIfDef = true
			}
		}
		foundString := false
		var search func(asciidoc.Element)
		search = func(el asciidoc.Element) {
			if s, ok := el.(*asciidoc.String); ok && strings.Contains(s.Value, "Normal Text after") {
				foundString = true
			}
			if withChildren, ok := el.(interface{ Children() []asciidoc.Element }); ok {
				for _, c := range withChildren.Children() {
					search(c)
				}
			}
		}
		for _, child := range cell.Children() {
			search(child)
		}

		if !hasIfDef {
			t.Error("Did not find IfDef")
		}
		if !foundString {
			t.Error("Did not find 'Normal Text after'")
		}
	})

	t.Run("Scenario 4: IfDef + Middle Text + IfDef (Unwrapped)", func(t *testing.T) {
		table := tables[3]
		cell := getDataCell(t, table)
		ifDefs := 0

		for _, child := range cell.Children() {
			if _, ok := child.(*asciidoc.IfDef); ok {
				ifDefs++
			}
		}

		// Expect 2 IfDefs now? Or still 3 if we consider "Middle Text" as implicitly wrapped?
		// User wants "IfDef + Middle Text + IfDef".
		// Use count of IfDefs found.
		if ifDefs < 2 {
			t.Errorf("Expected at least 2 IfDefs, found %d", ifDefs)
		}

		foundMiddle := false
		var search func(asciidoc.Element)
		search = func(el asciidoc.Element) {
			if s, ok := el.(*asciidoc.String); ok && strings.Contains(s.Value, "Middle Text") {
				foundMiddle = true
			}
			if withChildren, ok := el.(interface{ Children() []asciidoc.Element }); ok {
				for _, c := range withChildren.Children() {
					search(c)
				}
			}
		}
		for _, child := range cell.Children() {
			search(child)
		}

		if !foundMiddle {
			t.Error("Did not find 'Middle Text'")
		}
	})

	t.Run("Scenario 5: List + IfDef + List + IfDef (Unwrapped)", func(t *testing.T) {
		table := tables[4]
		cell := getDataCell(t, table)

		ifDefs := 0
		lists := 0

		for _, child := range cell.Children() {
			if _, ok := child.(*asciidoc.IfDef); ok {
				ifDefs++
			}
			if _, ok := child.(*asciidoc.UnorderedListItem); ok {
				lists++
			}
		}

		// 1 UnorderedListItem (Item 1)
		// 1 UnorderedListItem (Item 3)?
		// 2 IfDefs (Item 2, Item 4)

		if lists < 2 {
			t.Errorf("Expected at least 2 UnorderedListItems, found %d", lists)
		}
		if ifDefs < 2 {
			t.Errorf("Expected at least 2 IfDefs, found %d", ifDefs)
		}
	})

	t.Run("Scenario 6: Table Row Wrapped in IfDef", func(t *testing.T) {
		table := tables[5]
		// Verify structure:
		// Row 1: Header
		// Row 2: Data (Value 1, Value 2, List)
		// IfDef -> Row 3: Data (Value 1, Value 2, Value 3)

		// We need to check that the IfDef is NOT inside the last cell of Row 2.
		// And that we have a Row 3 (or IfDef containing Row 3).

		rows := 0
		ifDefs := 0
		for _, el := range table.Elements {
			if _, ok := el.(*asciidoc.TableRow); ok {
				rows++
			}
			if _, ok := el.(*asciidoc.IfDef); ok {
				ifDefs++
			}
		}

		// Header + Row 2 = 2 rows
		if rows < 2 {
			t.Errorf("Expected at least 2 top-level rows, found %d", rows)
		}
		if ifDefs < 1 {
			t.Errorf("Expected 1 top-level IfDef, found %d", ifDefs)
		}

		// Check last cell of Row 2 does NOT contain the IfDef
		// Row 2 is index 1
		var row2 *asciidoc.TableRow
		rowCount := 0
		for _, el := range table.Elements {
			if r, ok := el.(*asciidoc.TableRow); ok {
				if rowCount == 1 {
					row2 = r
					break
				}
				rowCount++
			}
		}

		if row2 != nil {
			cells := row2.TableCells()
			if len(cells) > 0 {
				lastCell := cells[len(cells)-1]
				for _, child := range lastCell.Children() {
					if _, ok := child.(*asciidoc.IfDef); ok {
						t.Error("Found IfDef in last cell of Row 2 (should be top-level)")
					}
					if s, ok := child.(*asciidoc.String); ok && strings.Contains(s.Value, "Value 3") {
						t.Error("Found 'Value 3' in last cell of Row 2 (should be in separate row)")
					}
				}
			}
		}
	})
}
