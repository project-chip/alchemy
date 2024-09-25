package spec

import (
	"context"
	"fmt"
	"iter"
	"log/slog"
	"strconv"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/asciidoc/render"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/internal/parse"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/constraint"
)

type ColumnIndex map[matter.TableColumn]int

func (ci ColumnIndex) HasAny(columns ...matter.TableColumn) bool {
	for _, col := range columns {
		_, ok := ci[col]
		if ok {
			return true
		}
	}
	return false
}

func (ci ColumnIndex) HasAll(columns ...matter.TableColumn) bool {
	if len(columns) == 0 {
		return false
	}
	for _, col := range columns {
		_, ok := ci[col]
		if !ok {
			return false
		}
	}
	return true
}

type ExtraColumn struct {
	Name   string
	Offset int
}

type TableInfo struct {
	Doc            *Doc
	Element        *asciidoc.Table
	Rows           []*asciidoc.TableRow
	HeaderRowIndex int
	ColumnMap      ColumnIndex
	ExtraColumns   []ExtraColumn
}

var ErrNoTableFound = fmt.Errorf("no table found")

func parseFirstTable(doc *Doc, section *Section) (ti *TableInfo, err error) {
	t := FindFirstTable(section)
	if t == nil {
		err = ErrNoTableFound
		return
	}
	return parseTable(doc, section, t)
}

func parseTable(doc *Doc, section *Section, t *asciidoc.Table) (ti *TableInfo, err error) {

	ti, err = ReadTable(doc, t)
	if err != nil {
		err = fmt.Errorf("failed mapping table columns for first table in section %s: %w", section.Name, err)
		return
	}
	if len(ti.Rows) < ti.HeaderRowIndex+2 {
		err = fmt.Errorf("not enough value rows in table")
		return
	}
	if ti.ColumnMap == nil {
		err = fmt.Errorf("can't read table without columns")
	}
	return
}

func (ti *TableInfo) ReadString(row *asciidoc.TableRow, columns ...matter.TableColumn) (string, error) {
	for _, column := range columns {
		offset, ok := ti.ColumnMap[column]
		if !ok {
			continue
		}
		cell := row.Cell(offset)
		val, err := RenderTableCell(cell)
		if err != nil {
			return "", err
		}
		val = asteriskPattern.ReplaceAllString(val, "")
		return val, nil
	}
	return "", nil
}

func (ti *TableInfo) ReadStringAtOffset(row *asciidoc.TableRow, offset int) (string, error) {
	cell := row.Cell(offset)
	val, err := RenderTableCell(cell)
	if err != nil {
		return "", err
	}
	val = asteriskPattern.ReplaceAllString(val, "")
	return val, nil
}

func (ti *TableInfo) ReadID(row *asciidoc.TableRow, columns ...matter.TableColumn) (*matter.Number, error) {
	id, err := ti.ReadString(row, columns...)
	if err != nil {
		return matter.InvalidID, err
	}
	return matter.ParseNumber(id), nil
}

func (ti *TableInfo) ReadName(row *asciidoc.TableRow, columns ...matter.TableColumn) (name string, xref *asciidoc.CrossReference, err error) {
	for _, column := range columns {
		offset, ok := ti.ColumnMap[column]
		if !ok {
			continue
		}
		cell := row.Cell(offset)
		cellElements := cell.Elements()
		for _, el := range cellElements {
			switch el := el.(type) {
			case *asciidoc.CrossReference:
				xref = el
			}
			if xref != nil {
				break
			}
		}
		var value strings.Builder
		err = readRowCellValueElements(ti.Doc, cellElements, &value)
		if err != nil {
			return "", nil, err
		}
		return strings.TrimSpace(value.String()), xref, nil
	}
	return "", nil, nil
}

func (ti *TableInfo) ReadValue(row *asciidoc.TableRow, columns ...matter.TableColumn) (string, error) {
	for _, column := range columns {
		offset, ok := ti.ColumnMap[column]
		if !ok {
			continue
		}
		return ti.ReadValueByIndex(row, offset)
	}
	return "", nil
}

func (ti *TableInfo) ReadValueByIndex(row *asciidoc.TableRow, offset int) (string, error) {
	cell := row.Cell(offset)
	cellElements := cell.Elements()
	if len(cellElements) == 0 {
		return "", nil
	}
	var value strings.Builder
	err := readRowCellValueElements(ti.Doc, cellElements, &value)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(value.String()), nil
}

type CellRenderer func(cellElements asciidoc.Set, sb *strings.Builder) (source asciidoc.Element)

func (ti *TableInfo) RenderColumn(row *asciidoc.TableRow, renderer CellRenderer, columns ...matter.TableColumn) (value string, source asciidoc.Element, ok bool) {
	for _, column := range columns {
		var offset int
		offset, ok = ti.ColumnMap[column]
		if !ok {
			continue
		}
		cell := row.Cell(offset)
		cellElements := cell.Elements()
		if len(cellElements) == 0 {
			ok = false
			return
		}
		source = cell
		var sb strings.Builder
		sourceOverride := renderer(cellElements, &sb)
		if sourceOverride != nil {
			source = sourceOverride
		}
		value = sb.String()
		return
	}
	return
}

var newLineReplacer = strings.NewReplacer("\r\n", "", "\r", "", "\n", "")

func (ti *TableInfo) ReadConformance(row *asciidoc.TableRow, column matter.TableColumn) conformance.Set {
	val, _, ok := ti.RenderColumn(row, ti.buildRowConformance, column)
	if !ok {
		return nil
	}

	s := strings.TrimSpace(val)
	if len(s) == 0 {
		return conformance.Set{&conformance.Mandatory{}}
	}
	s = newLineReplacer.Replace(s)
	return conformance.ParseConformance(matter.StripTypeSuffixes(s))
}

func (ti *TableInfo) buildRowConformance(cellElements asciidoc.Set, sb *strings.Builder) (source asciidoc.Element) {
	for _, el := range cellElements {
		switch v := el.(type) {
		case *asciidoc.String:
			sb.WriteString(v.Value)
		case *asciidoc.CrossReference:
			id := v.ID
			if strings.HasPrefix(id, "ref_") {
				// This is a proper reference; allow the conformance parser to recognize it
				sb.WriteString(fmt.Sprintf("<<%s>>", id))
			} else {
				anchor := ti.Doc.FindAnchor(v.ID)
				var name string
				if anchor != nil {
					name = ReferenceName(anchor.Element)
				} else {
					name = strings.TrimPrefix(v.ID, "_")
				}
				sb.WriteString(name)
			}
		case *asciidoc.SpecialCharacter:
			sb.WriteString(v.Character)
		case *asciidoc.Superscript:
			// This is usually an asterisk, and should be ignored
		case *asciidoc.Link:
			sb.WriteString(v.URL.Scheme)
			ti.buildRowConformance(v.URL.Path, sb)
		case *asciidoc.LinkMacro:
			sb.WriteString("link:")
			sb.WriteString(v.URL.Scheme)
			ti.buildRowConformance(v.URL.Path, sb)
		case *asciidoc.CharacterReplacementReference:
			switch v.Name() {
			case "nbsp":
				sb.WriteRune(' ')
			default:
				slog.Warn("unknown predefined attribute", log.Element("path", ti.Doc.Path, el), "name", v.Name)
			}
		case *asciidoc.NewLine:
			sb.WriteRune(' ')
		case asciidoc.HasElements:
			ti.buildRowConformance(v.Elements(), sb)
		default:
			slog.Warn("unknown conformance value element", log.Element("path", ti.Doc.Path, el), "type", fmt.Sprintf("%T", el))
		}
	}
	return
}

func (ti *TableInfo) ReadConstraint(row *asciidoc.TableRow, column matter.TableColumn) constraint.Constraint {

	val, source, _ := ti.RenderColumn(row, ti.buildConstraintValue, column)
	s := strings.TrimSpace(val)
	s = strings.ReplaceAll(s, "\n", " ")
	var c constraint.Constraint
	c, err := constraint.ParseString(s)
	if err != nil {
		slog.Error("failed parsing constraint cell", log.Element("path", ti.Doc.Path, source), slog.String("constraint", val))
		return &constraint.GenericConstraint{Value: val}
	}
	return c
}

func (ti *TableInfo) buildConstraintValue(els asciidoc.Set, sb *strings.Builder) (source asciidoc.Element) {
	for _, el := range els {
		switch v := el.(type) {
		case *asciidoc.String:
			sb.WriteString(v.Value)
		case *asciidoc.CrossReference:
			anchor := ti.Doc.FindAnchor(v.ID)
			var name string
			if anchor != nil {
				name = matter.StripReferenceSuffixes(ReferenceName(anchor.Element))
			} else {
				name = strings.TrimPrefix(v.ID, "_")
			}
			sb.WriteString(name)
		case *asciidoc.Superscript:
			var qt strings.Builder
			ti.buildConstraintValue(v.Elements(), &qt)
			val := qt.String()
			if val == "*" { // We ignore asterisks here
				continue
			}
			sb.WriteString("^")
			sb.WriteString(val)
			sb.WriteString("^")
		case *asciidoc.Bold: // This is usually an asterisk, and should be ignored
		case *asciidoc.NewLine, *asciidoc.LineBreak:
			sb.WriteRune(' ')
		case asciidoc.HasElements:
			ti.buildConstraintValue(v.Elements(), sb)
		case asciidoc.AttributeReference:
			sb.WriteString(fmt.Sprintf("{%s}", v.Name()))
		default:
			slog.Warn("unknown constraint value element", log.Element("path", ti.Doc.Path, el), "type", fmt.Sprintf("%T", el))
		}
	}
	return
}

func readRowCellValueElements(doc *Doc, els asciidoc.Set, value *strings.Builder) (err error) {
	for _, el := range els {
		switch el := el.(type) {
		case *asciidoc.String:
			value.WriteString(el.Value)
		case asciidoc.FormattedTextElement:
			err = readRowCellValueElements(doc, el.Elements(), value)
		case *asciidoc.Paragraph:
			err = readRowCellValueElements(doc, el.Elements(), value)
		case *asciidoc.CrossReference:
			if len(el.Set) > 0 {
				readRowCellValueElements(doc, el.Set, value)
			} else {
				var val string
				anchor := doc.FindAnchor(el.ID)
				if anchor != nil {
					val = matter.StripTypeSuffixes(ReferenceName(anchor.Element))
				} else {
					val = strings.TrimPrefix(el.ID, "_")
					val = strings.TrimPrefix(val, "ref_") // Trim, and hope someone else has it defined
				}
				value.WriteString(val)
			}
		case *asciidoc.Link:
			value.WriteString(el.URL.Scheme)
			readRowCellValueElements(doc, el.URL.Path, value)
		case *asciidoc.LinkMacro:
			value.WriteString(el.URL.Scheme)
			readRowCellValueElements(doc, el.URL.Path, value)
		case *asciidoc.Superscript:
			// In the special case of superscript elements, we do checks to make sure it's not an asterisk or a footnote, which should be ignored
			var quotedText strings.Builder
			err = readRowCellValueElements(doc, el.Elements(), &quotedText)
			if err != nil {
				return
			}
			qt := quotedText.String()
			if qt == "*" { //
				continue
			}
			_, parseErr := strconv.Atoi(qt)
			if parseErr == nil {
				// This is probably a footnote
				// The similar buildConstraintValue method does not do this, as there are exponential values in contraints
				continue
			}
			value.WriteString(qt)
		case *asciidoc.SpecialCharacter:
			value.WriteString(el.Character)
		case *asciidoc.InlinePassthrough:
			value.WriteString("+")
			err = readRowCellValueElements(doc, el.Elements(), value)
		case *asciidoc.InlineDoublePassthrough:
			value.WriteString("++")
			err = readRowCellValueElements(doc, el.Elements(), value)
		case *asciidoc.ThematicBreak:
		case asciidoc.EmptyLine:
		case *asciidoc.NewLine:
			value.WriteString(" ")
		case asciidoc.HasElements:
			err = readRowCellValueElements(doc, el.Elements(), value)
		case *asciidoc.LineBreak:
			value.WriteString(" ")
		default:
			return fmt.Errorf("unexpected type in cell: %T", el)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func FindFirstTable(section *Section) *asciidoc.Table {
	var table *asciidoc.Table
	parse.SkimFunc(section.Elements(), func(t *asciidoc.Table) bool {
		table = t
		return true
	})
	return table
}

func RenderTableCell(cell *asciidoc.TableCell) (string, error) {
	cellElements := cell.Elements()
	if len(cellElements) == 0 {
		return "", nil
	}
	out := render.NewUnwrappedTarget(context.Background())
	err := render.Elements(out, "", cellElements...)
	if err != nil {
		return "", err
	}
	return out.String(), nil
}

func (d *Doc) GetHeaderCellString(cell *asciidoc.TableCell) (string, error) {
	cellElements := cell.Elements()
	if len(cellElements) == 0 {
		return "", nil
	}
	var v strings.Builder
	err := readRowCellValueElements(d, cellElements, &v)
	if err != nil {
		return "", fmt.Errorf("error reading table header cell: %w", err)
	}
	return v.String(), nil
}

func (ti *TableInfo) ColumnIndex(columns ...matter.TableColumn) (index int, ok bool) {
	for _, column := range columns {
		index, ok = ti.ColumnMap[column]
		if ok {
			return
		}
	}
	return
}

func (ti *TableInfo) Rescan(doc *Doc) (err error) {
	ti.HeaderRowIndex, ti.ColumnMap, ti.ExtraColumns, err = mapTableColumns(doc, ti.Rows)
	return
}

func (ti *TableInfo) Body() iter.Seq[*asciidoc.TableRow] {
	return func(yield func(*asciidoc.TableRow) bool) {
		for i := ti.HeaderRowIndex + 1; i < len(ti.Rows); i++ {
			if !yield(ti.Rows[i]) {
				return
			}
		}

	}
}

func ReadTable(doc *Doc, table *asciidoc.Table) (ti *TableInfo, err error) {
	ti = &TableInfo{Doc: doc, Element: table, Rows: table.TableRows()}
	ti.HeaderRowIndex, ti.ColumnMap, ti.ExtraColumns, err = mapTableColumns(doc, ti.Rows)
	return
}

func mapTableColumns(doc *Doc, rows []*asciidoc.TableRow) (headerRow int, columnMap ColumnIndex, extraColumns []ExtraColumn, err error) {
	var cellCount = -1
	headerRow = -1
	for i, row := range rows {
		tableCells := row.TableCells()
		if cellCount == -1 {
			cellCount = len(tableCells)
		} else if cellCount != len(tableCells) {
			return -1, nil, nil, fmt.Errorf("can't map table columns with unequal cell counts between rows; row %d has %d cells, expected %d", i, len(tableCells), cellCount)
		}
		if columnMap != nil { // We've already processed the columns
			continue
		}
		var spares []ExtraColumn
		for j, cell := range tableCells {
			var val string
			val, err = doc.GetHeaderCellString(cell)
			if err != nil {
				return
			}
			attributeColumn := getTableColumn(val)
			if attributeColumn == matter.TableColumnUnknown {
				spares = append(spares, ExtraColumn{Name: val, Offset: j})
				continue
			}
			if columnMap == nil {
				headerRow = i
				columnMap = make(ColumnIndex)
			}
			if _, ok := columnMap[attributeColumn]; ok {
				return -1, nil, nil, fmt.Errorf("can't map table columns with duplicate columns")
			}
			columnMap[attributeColumn] = j

		}
		if columnMap != nil { // Don't return extra columns if we were unable to parse any regular columns
			extraColumns = spares
		}
	}
	return headerRow, columnMap, extraColumns, nil
}

func getTableColumn(val string) matter.TableColumn {
	switch strings.ToLower(val) {
	case "id", "identifier":
		return matter.TableColumnID
	case "name", "field":
		return matter.TableColumnName
	case "type":
		return matter.TableColumnType
	case "constraint":
		return matter.TableColumnConstraint
	case "quality":
		return matter.TableColumnQuality
	case "default":
		return matter.TableColumnDefault
	case "access":
		return matter.TableColumnAccess
	case "conformance":
		return matter.TableColumnConformance
	case "priority":
		return matter.TableColumnPriority
	case "hierarchy":
		return matter.TableColumnHierarchy
	case "role":
		return matter.TableColumnRole
	case "context":
		return matter.TableColumnContext
	case "pics code":
		return matter.TableColumnPICSCode
	case "pics":
		return matter.TableColumnPICS
	case "scope":
		return matter.TableColumnScope
	case "value":
		return matter.TableColumnValue
	case "bit", "bit index":
		return matter.TableColumnBit
	case "code":
		return matter.TableColumnCode
	case "feature":
		return matter.TableColumnFeature
	case "cluster id":
		return matter.TableColumnClusterID
	case "device id", "device type id":
		return matter.TableColumnDeviceID
	case "device name":
		return matter.TableColumnDeviceName
	case "superset":
		return matter.TableColumnSuperset
	case "class":
		return matter.TableColumnClass
	case "direction":
		return matter.TableColumnDirection
	case "response":
		return matter.TableColumnResponse
	case "description":
		return matter.TableColumnDescription
	case "summary":
		return matter.TableColumnSummary
	case "cluster", "cluster name", "clustername":
		return matter.TableColumnCluster
	case "client/server":
		return matter.TableColumnClientServer
	case "revision", "rev":
		return matter.TableColumnRevision
	case "element":
		return matter.TableColumnElement
	case "condition":
		return matter.TableColumnCondition
	case "mode tag value":
		return matter.TableColumnModeTagValue
	case "status code":
		return matter.TableColumnStatusCode
	}
	return matter.TableColumnUnknown
}
