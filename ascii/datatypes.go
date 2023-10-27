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

func (s *Section) toDataTypes() (dataTypes []interface{}, err error) {

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
			me, err = s.toStruct()
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
		Name: matter.StripDataTypeSuffixes(name),
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
		Name: matter.StripDataTypeSuffixes(name),
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

func (s *Section) toStruct() (ms *matter.Struct, err error) {
	var rows []*types.TableRow
	var headerRowIndex int
	var columnMap map[matter.TableColumn]int
	rows, headerRowIndex, columnMap, _, err = parseFirstTable(s)
	if err != nil {
		return nil, fmt.Errorf("failed reading struct: %w", err)
	}
	name := strings.TrimSuffix(s.Name, " Type")
	ms = &matter.Struct{
		Name: matter.StripDataTypeSuffixes(name),
	}

	ms.Fields, err = readFields(headerRowIndex, rows, columnMap)
	return
}

func readFields(headerRowIndex int, rows []*types.TableRow, columnMap map[matter.TableColumn]int) (fields []*matter.Field, err error) {
	for i := headerRowIndex + 1; i < len(rows); i++ {
		row := rows[i]
		f := &matter.Field{}
		f.Name, err = readRowValue(row, columnMap, matter.TableColumnName)
		if err != nil {
			return
		}
		f.Type = getRowDataType(row, columnMap, matter.TableColumnType)
		f.Constraint, err = readRowValue(row, columnMap, matter.TableColumnConstraint)
		if err != nil {
			return
		}
		f.Quality, err = readRowValue(row, columnMap, matter.TableColumnQuality)
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
		f.Access = matter.ParseAccess(a)
		f.ID, err = readRowValue(row, columnMap, matter.TableColumnID)
		if err != nil {
			return
		}

		fields = append(fields, f)
	}
	return
}

var listDataTypeDefinitionPattern = regexp.MustCompile(`list\[([^\]]+)\]`)

func getRowDataType(row *types.TableRow, columnMap map[matter.TableColumn]int, column matter.TableColumn) *matter.DataType {
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
	if len(p.Elements) == 1 {
		se, ok := p.Elements[0].(*types.StringElement)
		if ok {
			match := listDataTypeDefinitionPattern.FindStringSubmatch(se.Content)
			if match != nil {
				return &matter.DataType{Name: match[1], IsArray: true}
			}
			return &matter.DataType{Name: se.Content}
		}
	}
	val := &matter.DataType{}
	for _, el := range p.Elements {
		switch v := el.(type) {
		case *types.StringElement:
			if v.Content == "list[" {
				val.IsArray = true
			} else if val.Name == "" {
				val.Name = v.Content
			}
		case *types.InternalCrossReference:
			val.Name = v.ID.(string)
			return val
		default:
			slog.Info("unknown value element", "type", v)
		}
	}
	return val
}
