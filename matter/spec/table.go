package spec

import (
	"context"
	"fmt"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/asciidoc/render"
	"github.com/project-chip/alchemy/internal/parse"
	"github.com/project-chip/alchemy/matter"
)

var ErrNoTableFound = fmt.Errorf("no table found")
var ErrNotEnoughRowsInTable = fmt.Errorf("not enough value rows in table")

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
		err = ErrNotEnoughRowsInTable
		return
	}
	if ti.ColumnMap == nil {
		err = fmt.Errorf("can't read table without columns")
	}
	return
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
	case "default", "fallback":
		return matter.TableColumnFallback
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
	case "device":
		return matter.TableColumnDevice
	case "location":
		return matter.TableColumnLocation
	case "device id", "device type id":
		return matter.TableColumnDeviceID
	case "device name":
		return matter.TableColumnDeviceName
	case "superset", "superset of":
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
	case "namespace":
		return matter.TableColumnNamespace
	case "mode tag value":
		return matter.TableColumnModeTagValue
	case "status code", "status code value":
		return matter.TableColumnStatusCode
	}
	return matter.TableColumnUnknown
}
