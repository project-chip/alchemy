package ascii

import (
	"context"
	"fmt"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/alchemy/ascii/render"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/output"
	"github.com/hasty/alchemy/parse"
)

type ColumnIndex map[matter.TableColumn]int

var NoTableFound = fmt.Errorf("no table found")

func parseFirstTable(doc *Doc, section *Section) (rows []*types.TableRow, headerRowIndex int, columnMap ColumnIndex, extraColumns []ExtraColumn, err error) {
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

func readRowValue(row *types.TableRow, columnMap ColumnIndex, columns ...matter.TableColumn) (string, error) {
	for _, column := range columns {
		offset, ok := columnMap[column]
		if !ok {
			continue
		}
		return readRowCell(row, offset)
	}
	return "", nil
}

func readRowCell(row *types.TableRow, offset int) (string, error) {
	cell := row.Cells[offset]
	val, err := GetTableCellValue(cell)
	if err != nil {
		return "", err
	}
	val = asteriskPattern.ReplaceAllString(val, "")
	return val, nil
}

func readRowID(row *types.TableRow, columnMap ColumnIndex, column matter.TableColumn) (*matter.Number, error) {
	id, err := readRowValue(row, columnMap, column)
	if err != nil {
		return matter.InvalidID, err
	}
	return matter.ParseNumber(id), nil
}

func readRowName(doc *Doc, row *types.TableRow, columnMap ColumnIndex, columns ...matter.TableColumn) (string, error) {
	for _, column := range columns {
		offset, ok := columnMap[column]
		if !ok {
			continue
		}
		return readRowCellName(doc, row, offset)
	}
	return "", nil
}

func readRowCellName(doc *Doc, row *types.TableRow, offset int) (string, error) {
	cell := row.Cells[offset]
	if len(cell.Elements) == 0 {
		return "", nil
	}
	el := cell.Elements[0]
	para, ok := el.(*types.Paragraph)
	if !ok {
		return "", fmt.Errorf("name cell missing paragraph")
	}
	for _, el := range para.Elements {
		switch el := el.(type) {
		case *types.StringElement:
			return el.Content, nil
		case *types.InternalCrossReference:
			var val string
			anchor, _ := doc.getAnchor(el.ID.(string))
			if anchor != nil {
				val = ReferenceName(anchor.Element)
			} else {
				val = strings.TrimPrefix(el.ID.(string), "_")
				val = strings.TrimPrefix(val, "ref_") // Trim, and hope someone else has it defined
			}
			return val, nil
		default:
			return "", fmt.Errorf("unexpected type in name cell: %T", el)
		}
	}
	return "", nil
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
	err := render.RenderElements(out, "", p.Elements)
	if err != nil {
		return "", err
	}
	return out.String(), nil
}

func (d *Doc) GetTableHeaderCellValue(cell *types.TableCell) (string, error) {
	if len(cell.Elements) == 0 {
		return "", fmt.Errorf("missing table header cell elements")
	}
	p, ok := cell.Elements[0].(*types.Paragraph)
	if !ok {
		return "", fmt.Errorf("missing paragraph in table cell")
	}
	if len(p.Elements) == 0 {
		return "", nil
	}
	var v strings.Builder
	err := d.readCellContent(p.Elements, &v)
	if err != nil {
		return "", err
	}
	return v.String(), nil
}

func (d *Doc) readCellContent(elements []any, content *strings.Builder) (err error) {
	for _, s := range elements {
		switch s := s.(type) {
		case *types.StringElement:
			content.WriteString(s.Content)
		case *types.QuotedText:
			return d.readCellContent(s.Elements, content)
		case *types.InternalCrossReference:
			var name string
			anchor, _ := d.getAnchor(s.ID.(string))
			if anchor != nil {
				name = ReferenceName(anchor.Element)
			} else {
				name = s.ID.(string)
			}
			content.WriteString(name)
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

func MapTableColumns(doc *Doc, rows []*types.TableRow) (headerRow int, columnMap ColumnIndex, extraColumns []ExtraColumn, err error) {
	var cellCount = -1
	headerRow = -1
	for i, row := range rows {
		if cellCount == -1 {
			cellCount = len(row.Cells)
		} else if cellCount != len(row.Cells) {
			return -1, nil, nil, fmt.Errorf("can't map table columns with unequal cell counts between rows; row %d has %d cells, expected %d", i, len(row.Cells), cellCount)
		}
		if columnMap == nil {
			var spares []ExtraColumn
			for j, cell := range row.Cells {
				var val string
				val, err = doc.GetTableHeaderCellValue(cell)
				if err != nil {
					return
				}
				attributeColumn := getTableColumn(val)
				if attributeColumn != matter.TableColumnUnknown {
					if columnMap == nil {
						headerRow = i
						columnMap = make(ColumnIndex)
					}
					if _, ok := columnMap[attributeColumn]; ok {
						return -1, nil, nil, fmt.Errorf("can't map table columns with duplicate columns")
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
	}
	return matter.TableColumnUnknown
}
