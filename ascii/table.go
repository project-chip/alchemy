package ascii

import (
	"context"
	"fmt"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/matterfmt/matter"
	"github.com/hasty/matterfmt/output"
	"github.com/hasty/matterfmt/parse"
	"github.com/hasty/matterfmt/render/adoc"
)

var NoTableFound = fmt.Errorf("no table found")

func parseFirstTable(section *Section) (rows []*types.TableRow, headerRowIndex int, columnMap map[matter.TableColumn]int, extraColumns []ExtraColumn, err error) {
	t := FindFirstTable(section)
	if t == nil {
		err = NoTableFound
		return
	}
	rows = TableRows(t)
	if len(rows) < 2 {
		err = fmt.Errorf("not enough rows in table")
		return
	}
	headerRowIndex, columnMap, extraColumns, err = MapTableColumns(rows)
	if err != nil {
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

func readRowValue(row *types.TableRow, columnMap map[matter.TableColumn]int, column matter.TableColumn) (string, error) {
	i, ok := columnMap[column]
	if !ok {
		return "", nil
	}
	cell := row.Cells[i]
	return GetTableCellValue(cell)
}

func FindFirstTable(section *Section) *types.Table {
	var table *types.Table
	parse.Search(section.Elements, func(t *types.Table) bool {
		table = t
		return true
	})

	return table
}

func TableRows(t *types.Table) (rows []*types.TableRow) {
	rows = make([]*types.TableRow, 0, len(t.Rows)+2)
	if t.Header != nil {
		rows = append(rows, t.Header)
	}
	rows = append(rows, t.Rows...)
	if t.Footer != nil {
		rows = append(rows, t.Footer)
	}
	return
}

func GetTableCellValue(cell *types.TableCell) (string, error) {
	if len(cell.Elements) == 0 {
		return "", fmt.Errorf("missing table cell elements")
	}
	p, ok := cell.Elements[0].(*types.Paragraph)
	if !ok {
		return "", fmt.Errorf("missing paragraph in table cell")
	}
	if len(p.Elements) == 0 {
		return "", nil
	}
	out := output.NewContext(context.Background(), nil)
	err := adoc.RenderElements(out, "", p.Elements)
	if err != nil {
		return "", err
	}
	return out.String(), nil
}

type ExtraColumn struct {
	Name   string
	Offset int
}

func MapTableColumns(rows []*types.TableRow) (headerRow int, columnMap map[matter.TableColumn]int, extraColumns []ExtraColumn, err error) {
	var cellCount = -1
	headerRow = -1
	for i, row := range rows {
		if cellCount == -1 {
			cellCount = len(row.Cells)
		} else if cellCount != len(row.Cells) {
			return -1, nil, nil, fmt.Errorf("can't rearrange attributes table with unequal cell counts between rows")
		}
		if columnMap == nil {
			var spares []ExtraColumn
			for j, cell := range row.Cells {
				val, _ := GetTableCellValue(cell)
				attributeColumn := getTableColumn(val)
				if attributeColumn != matter.TableColumnUnknown {
					if columnMap == nil {
						headerRow = i
						columnMap = make(map[matter.TableColumn]int)
					}
					if _, ok := columnMap[attributeColumn]; ok {
						return -1, nil, nil, fmt.Errorf("can't rearrange attributes table duplicate columns")
					}
					columnMap[attributeColumn] = j
				} else {
					spares = append(spares, ExtraColumn{Name: val, Offset: j})
				}
			}
			if columnMap != nil {
				extraColumns = spares
			}
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
	case "bit":
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
	}
	return matter.TableColumnUnknown
}
