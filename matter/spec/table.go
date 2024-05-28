package spec

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/hasty/alchemy/asciidoc"
	"github.com/hasty/alchemy/asciidoc/render"
	"github.com/hasty/alchemy/internal/parse"
	"github.com/hasty/alchemy/matter"
)

type ColumnIndex map[matter.TableColumn]int

var ErrNoTableFound = fmt.Errorf("no table found")

func parseFirstTable(doc *Doc, section *Section) (rows []*asciidoc.TableRow, headerRowIndex int, columnMap ColumnIndex, extraColumns []ExtraColumn, err error) {
	t := FindFirstTable(section)
	if t == nil {
		err = ErrNoTableFound
		return
	}
	return parseTable(doc, section, t)
}

func parseTable(doc *Doc, section *Section, t *asciidoc.Table) (rows []*asciidoc.TableRow, headerRowIndex int, columnMap ColumnIndex, extraColumns []ExtraColumn, err error) {

	rows = t.TableRows()
	if len(rows) < 2 {
		err = fmt.Errorf("not enough rows in table")
		return
	}
	headerRowIndex, columnMap, extraColumns, err = MapTableColumns(doc, rows)
	if err != nil {
		err = fmt.Errorf("failed mapping table columns for first table in section %s: %w", section.Name, err)
		return
	}
	if len(rows) < headerRowIndex+2 {
		err = fmt.Errorf("not enough value rows in table")
		return
	}
	if columnMap == nil {
		err = fmt.Errorf("can't read table without columns")
	}
	return
}

func readRowASCIIDocString(row *asciidoc.TableRow, columnMap ColumnIndex, columns ...matter.TableColumn) (string, error) {
	for _, column := range columns {
		offset, ok := columnMap[column]
		if !ok {
			continue
		}
		return readRowCellASCIIDocString(row, offset)
	}
	return "", nil
}

func readRowCellASCIIDocString(row *asciidoc.TableRow, offset int) (string, error) {
	cell := row.Cell(offset)
	val, err := RenderTableCell(cell)
	if err != nil {
		return "", err
	}
	val = asteriskPattern.ReplaceAllString(val, "")
	return val, nil
}

func readRowID(row *asciidoc.TableRow, columnMap ColumnIndex, column matter.TableColumn) (*matter.Number, error) {
	id, err := readRowASCIIDocString(row, columnMap, column)
	if err != nil {
		return matter.InvalidID, err
	}
	return matter.ParseNumber(id), nil
}

func ReadRowValue(doc *Doc, row *asciidoc.TableRow, columnMap ColumnIndex, columns ...matter.TableColumn) (string, error) {
	for _, column := range columns {
		offset, ok := columnMap[column]
		if !ok {
			continue
		}
		return readRowCellValue(doc, row, offset)
	}
	return "", nil
}

func readRowCellValue(doc *Doc, row *asciidoc.TableRow, offset int) (string, error) {
	cell := row.Cell(offset)
	cellElements := cell.Elements()
	if len(cellElements) == 0 {
		return "", nil
	}
	var value strings.Builder
	err := readRowCellValueElements(doc, cellElements, &value)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(value.String()), nil
}

func readRowCellValueElements(doc *Doc, els asciidoc.Set, value *strings.Builder) error {
	for _, el := range els {
		switch el := el.(type) {
		case *asciidoc.String:
			value.WriteString(el.Value)
		case *asciidoc.CrossReference:
			var val string
			anchor, _ := doc.getAnchor(el.ID)
			if anchor != nil {
				val = matter.StripTypeSuffixes(ReferenceName(anchor.Element))
			} else {
				val = strings.TrimPrefix(el.ID, "_")
				val = strings.TrimPrefix(val, "ref_") // Trim, and hope someone else has it defined
			}
			value.WriteString(val)
		case *asciidoc.Link:
			value.WriteString(el.URL.Scheme)
			l, ok := el.URL.Path.(string)
			if ok {
				value.WriteString(l)
			}
		case *asciidoc.Superscript:
			// In the special case of superscript elements, we do checks to make sure it's not an asterisk or a footnote, which should be ignored
			var quotedText strings.Builder
			err := readRowCellValueElements(doc, el.Elements(), &quotedText)
			if err != nil {
				return err
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
			err := readRowCellValueElements(doc, el.Elements(), value)
			if err != nil {
				return err
			}
		case *asciidoc.InlineDoublePassthrough:
			value.WriteString("++")
			err := readRowCellValueElements(doc, el.Elements(), value)
			if err != nil {
				return err
			}
		case *asciidoc.ThematicBreak:
		case *asciidoc.NewLine:
			value.WriteString(" ")
		case asciidoc.HasElements:
			err := readRowCellValueElements(doc, el.Elements(), value)
			if err != nil {
				return err
			}
		default:
			return fmt.Errorf("unexpected type in cell: %T", el)
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
	out := render.NewContext(context.Background(), nil)
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
	err := d.readCellContent(cellElements, &v)
	if err != nil {
		return "", err
	}
	return v.String(), nil
}

func (d *Doc) readCellContent(els asciidoc.Set, content *strings.Builder) (err error) {
	for _, s := range els {
		switch s := s.(type) {
		case *asciidoc.String:
			content.WriteString(s.Value)
		case asciidoc.FormattedTextElement:
			return d.readCellContent(s.Elements(), content)
		case *asciidoc.CrossReference:
			var name string
			anchor, _ := d.getAnchor(s.ID)
			if anchor != nil {
				name = ReferenceName(anchor.Element)
			} else {
				name = s.ID
			}
			content.WriteString(name)
		case *asciidoc.SpecialCharacter:
			content.WriteString(s.Character)
		case *asciidoc.NewLine:
			content.WriteString(" ")
		case *asciidoc.Paragraph:
			return d.readCellContent(s.Elements(), content)
		default:
			return fmt.Errorf("unknown element in table header cell: %T", s)
		}
	}
	return nil
}

type ExtraColumn struct {
	Name   string
	Offset int
}

func MapTableColumns(doc *Doc, rows []*asciidoc.TableRow) (headerRow int, columnMap ColumnIndex, extraColumns []ExtraColumn, err error) {
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
	case "pics code", "pics":
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
	case "cluster":
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
	}
	return matter.TableColumnUnknown
}