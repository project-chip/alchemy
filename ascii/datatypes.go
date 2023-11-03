package ascii

import (
	"fmt"
	"log/slog"
	"regexp"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/matterfmt/matter"
	"github.com/hasty/matterfmt/parse"
)

func (s *Section) toDataTypes(d *Doc) (dataTypes []interface{}, err error) {

	for _, s := range parse.Skim[*Section](s.Elements) {
		switch s.SecType {
		case matter.SectionDataTypeBitmap:
			var mb *matter.Bitmap
			mb, err = s.toBitmap()
			if err != nil {
				return
			}
			dataTypes = append(dataTypes, mb)
		case matter.SectionDataTypeEnum:
			var me *matter.Enum
			me, err = s.toEnum()
			if err != nil {
				return
			}
			dataTypes = append(dataTypes, me)
		case matter.SectionDataTypeStruct:
			var me *matter.Struct
			me, err = s.toStruct(d)
			if err != nil {
				return
			}
			dataTypes = append(dataTypes, me)
		default:
		}
	}
	return
}

func (s *Section) toEnum() (e *matter.Enum, err error) {
	var rows []*types.TableRow
	var headerRowIndex int
	var columnMap map[matter.TableColumn]int
	rows, headerRowIndex, columnMap, _, err = parseFirstTable(s)
	if err != nil {
		return nil, fmt.Errorf("failed reading enum: %w", err)
	}
	name := strings.TrimSuffix(s.Name, " Type")
	e = &matter.Enum{
		Name: name,
		Type: s.GetDataType(),
	}

	for i := headerRowIndex + 1; i < len(rows); i++ {
		row := rows[i]
		ev := &matter.EnumValue{}
		ev.Name, err = readRowValue(row, columnMap, matter.TableColumnName)
		if err != nil {
			return
		}
		ev.Summary, err = readRowValue(row, columnMap, matter.TableColumnSummary)
		if err != nil {
			return
		}
		ev.Conformance, err = readRowValue(row, columnMap, matter.TableColumnConformance)
		if err != nil {
			return
		}
		ev.Value, err = readRowValue(row, columnMap, matter.TableColumnValue)
		if err != nil {
			return
		}
		e.Values = append(e.Values, ev)
	}
	return
}

func (s *Section) toBitmap() (e *matter.Bitmap, err error) {
	var rows []*types.TableRow
	var headerRowIndex int
	var columnMap map[matter.TableColumn]int
	rows, headerRowIndex, columnMap, _, err = parseFirstTable(s)
	if err != nil {
		return nil, fmt.Errorf("failed reading bitmap: %w", err)
	}
	name := strings.TrimSuffix(s.Name, " Type")
	e = &matter.Bitmap{
		Name: name,
		Type: s.GetDataType(),
	}

	for i := headerRowIndex + 1; i < len(rows); i++ {
		row := rows[i]
		bv := &matter.BitmapValue{}
		bv.Name, err = readRowValue(row, columnMap, matter.TableColumnName)
		if err != nil {
			return
		}
		bv.Summary, err = readRowValue(row, columnMap, matter.TableColumnSummary)
		if err != nil {
			return
		}
		bv.Conformance, err = readRowValue(row, columnMap, matter.TableColumnConformance)
		if err != nil {
			return
		}
		bv.Bit, err = readRowValue(row, columnMap, matter.TableColumnValue)
		if err != nil {
			return
		}
		e.Bits = append(e.Bits, bv)
	}
	return
}

func (s *Section) toStruct(d *Doc) (ms *matter.Struct, err error) {
	var rows []*types.TableRow
	var headerRowIndex int
	var columnMap map[matter.TableColumn]int
	rows, headerRowIndex, columnMap, _, err = parseFirstTable(s)
	if err != nil {
		return nil, fmt.Errorf("failed reading struct: %w", err)
	}
	name := strings.TrimSuffix(s.Name, " Type")
	ms = &matter.Struct{
		Name: name,
	}

	ms.Fields, err = d.readFields(headerRowIndex, rows, columnMap)
	return
}

func (d *Doc) readFields(headerRowIndex int, rows []*types.TableRow, columnMap map[matter.TableColumn]int) (fields []*matter.Field, err error) {
	for i := headerRowIndex + 1; i < len(rows); i++ {
		row := rows[i]
		f := &matter.Field{}
		f.Name, err = readRowValue(row, columnMap, matter.TableColumnName)
		if err != nil {
			return
		}
		f.Type = d.getRowDataType(row, columnMap, matter.TableColumnType)
		f.Constraint = d.getRowConstraint(row, columnMap, matter.TableColumnConstraint, f.Type)
		if err != nil {
			return
		}
		var q string
		q, err = readRowValue(row, columnMap, matter.TableColumnQuality)
		if err != nil {
			return
		}
		f.Quality = matter.ParseQuality(q)
		if err != nil {
			return
		}
		f.Default, err = readRowValue(row, columnMap, matter.TableColumnDefault)
		if err != nil {
			return
		}
		f.Conformance, err = readRowValue(row, columnMap, matter.TableColumnConformance)
		if err != nil {
			return
		}
		var a string
		a, err = readRowValue(row, columnMap, matter.TableColumnAccess)
		if err != nil {
			return
		}
		f.Access = ParseAccess(a)
		f.ID, err = readRowValue(row, columnMap, matter.TableColumnID)
		if err != nil {
			return
		}

		fields = append(fields, f)
	}
	return
}

var listDataTypeDefinitionPattern = regexp.MustCompile(`list\[([^\]]+)\]`)

func (d *Doc) getRowDataType(row *types.TableRow, columnMap map[matter.TableColumn]int, column matter.TableColumn) *matter.DataType {
	i, ok := columnMap[column]
	if !ok {
		return nil
	}
	cell := row.Cells[i]
	if len(cell.Elements) == 0 {
		return nil
	}
	p, ok := cell.Elements[0].(*types.Paragraph)
	if !ok {
		fmt.Printf("el type: %T\n", cell.Elements[0])
		return nil
	}
	if len(p.Elements) == 0 {
		return nil
	}
	val := &matter.DataType{}
	if len(p.Elements) == 1 {
		se, ok := p.Elements[0].(*types.StringElement)
		if ok {
			match := listDataTypeDefinitionPattern.FindStringSubmatch(se.Content)
			if match != nil {
				val.Name = match[1]
				val.IsArray = true
			} else {
				val.Name = se.Content
			}
		}
	} else {
		for _, el := range p.Elements {
			switch v := el.(type) {
			case *types.StringElement:
				if v.Content == "list[" {
					val.IsArray = true
				} else if val.Name == "" {
					anchor, ok := d.anchors[v.Content]
					if ok {
						fmt.Printf("array anchor: %v\n", anchor.Element)
						val.Name = ReferenceName(anchor.Element)
					} else {
						val.Name = v.Content
					}
				}
			case *types.InternalCrossReference:
				anchor, ok := d.anchors[v.ID.(string)]
				if ok {
					fmt.Printf("type anchor: %v %T\n", anchor.Element, anchor.Element)
					val.Name = ReferenceName(anchor.Element)
				} else {
					val.Name = v.ID.(string)
				}
				break
			default:
				slog.Info("unknown value element", "type", v)
			}
		}
	}
	val.Name = strings.TrimSuffix(val.Name, " Type")

	return val
}

func (d *Doc) getRowConstraint(row *types.TableRow, columnMap map[matter.TableColumn]int, column matter.TableColumn, parentDataType *matter.DataType) matter.Constraint {
	i, ok := columnMap[column]
	if !ok {
		return nil
	}
	cell := row.Cells[i]
	if len(cell.Elements) == 0 {
		return nil
	}
	p, ok := cell.Elements[0].(*types.Paragraph)
	if !ok {
		fmt.Printf("el type: %T\n", cell.Elements[0])
		return nil
	}
	if len(p.Elements) == 0 {
		return nil
	}
	var val matter.Constraint
	if len(p.Elements) == 1 {
		se, ok := p.Elements[0].(*types.StringElement)
		if ok {
			val = ParseConstraint(parentDataType, se.Content)
		}
	} else {
		for _, el := range p.Elements {
			switch v := el.(type) {
			case *types.StringElement:
				val = ParseConstraint(parentDataType, v.Content)
			default:
				slog.Info("unknown value element", "type", v)
			}
		}
	}
	return val
}
