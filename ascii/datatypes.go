package ascii

import (
	"fmt"
	"log/slog"
	"regexp"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/alchemy/conformance"
	"github.com/hasty/alchemy/constraint"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/parse"
)

func (s *Section) toDataTypes(d *Doc, cluster *matter.Cluster) (err error) {

	for _, s := range parse.Skim[*Section](s.Elements) {
		switch s.SecType {
		case matter.SectionDataTypeBitmap:
			var mb *matter.Bitmap
			mb, err = s.toBitmap(d)
			if err != nil {
				return
			}
			cluster.Bitmaps = append(cluster.Bitmaps, mb)
		case matter.SectionDataTypeEnum:
			var me *matter.Enum
			me, err = s.toEnum(d)
			if err != nil {
				return
			}
			cluster.Enums = append(cluster.Enums, me)
		case matter.SectionDataTypeStruct:
			var me *matter.Struct
			me, err = s.toStruct(d)
			if err != nil {
				return
			}
			cluster.Structs = append(cluster.Structs, me)
		default:
		}
	}
	return
}

func (d *Doc) readFields(headerRowIndex int, rows []*types.TableRow, columnMap ColumnIndex) (fields []*matter.Field, err error) {
	for i := headerRowIndex + 1; i < len(rows); i++ {
		row := rows[i]
		f := &matter.Field{}
		f.Name, err = readRowValue(row, columnMap, matter.TableColumnName)
		if err != nil {
			return
		}
		f.Type = d.ReadRowDataType(row, columnMap, matter.TableColumnType)
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
		f.Conformance = d.getRowConformance(row, columnMap, matter.TableColumnConformance)
		if err != nil {
			return
		}
		var a string
		a, err = readRowValue(row, columnMap, matter.TableColumnAccess)
		if err != nil {
			return
		}
		f.Access = ParseAccess(a)
		f.ID, err = readRowID(row, columnMap, matter.TableColumnID)
		if err != nil {
			return
		}

		fields = append(fields, f)
	}
	return
}

var listDataTypeDefinitionPattern = regexp.MustCompile(`list\[([^\]]+)\]`)
var asteriskPattern = regexp.MustCompile(`\^[0-9]+\^\s*$`)

func (d *Doc) ReadRowDataType(row *types.TableRow, columnMap ColumnIndex, column matter.TableColumn) *matter.DataType {
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
		return nil
	}
	if len(p.Elements) == 0 {
		return nil
	}
	var name string
	var isArray bool
	if len(p.Elements) == 1 {
		el := p.Elements[0]
		switch v := el.(type) {
		case *types.StringElement:
			content := v.Content
			content = asteriskPattern.ReplaceAllString(content, "")
			match := listDataTypeDefinitionPattern.FindStringSubmatch(content)
			if match != nil {
				name = match[1]
				isArray = true
			} else {
				name = v.Content
			}
		case *types.InternalCrossReference:
			anchor, _ := d.getAnchor(v.ID.(string))
			if anchor != nil {
				name = ReferenceName(anchor.Element)
			} else {
				name = v.ID.(string)
			}
			break
		default:
			slog.Warn("unknown data type value element", "type", fmt.Sprintf("%T", el))
		}
	} else {
		for _, el := range p.Elements {
			switch v := el.(type) {
			case *types.StringElement:
				content := v.Content
				content = asteriskPattern.ReplaceAllString(content, "")

				if content == "list[" {
					slog.Debug("isArray", "content", content)
					isArray = true
				} else if name == "" {
					slog.Debug("inner list", "content", content)
					anchor, _ := d.getAnchor(content)
					if anchor != nil {
						name = ReferenceName(anchor.Element)
					} else {
						slog.Debug("inner list", "no anchor", content)
						name = strings.TrimPrefix(content, "_")
					}
				}
			case *types.InternalCrossReference:
				anchor, _ := d.getAnchor(v.ID.(string))
				slog.Debug("inner list", "icdr", v.ID.(string))
				if anchor != nil {
					slog.Debug("inner list", "anchor", anchor.Element)
					name = ReferenceName(anchor.Element)
				} else {
					slog.Debug("inner list", "no anchor", v.ID.(string))
					name = strings.TrimPrefix(v.ID.(string), "_")
					name = strings.TrimPrefix(name, "ref_") // Trim, and hope someone else has it defined
				}
				break
			default:
				slog.Warn("unknown data type value element", "type", v)
			}
		}
	}
	name = strings.TrimSuffix(name, " Type")
	return matter.NewDataType(name, isArray)
}

func (d *Doc) getAnchor(id string) (*Anchor, error) {
	anchors, err := d.Anchors()
	if err != nil {
		return nil, err
	}
	if a, ok := anchors[id]; ok {
		return a, nil
	}
	for _, p := range d.Parents() {
		slog.Debug("checking parents for anchor", "id", id)
		a, err := p.getAnchor(id)
		if err != nil {
			return nil, err
		}
		if a != nil {
			return a, nil
		}
	}
	return nil, nil
}

func (d *Doc) getRowConstraint(row *types.TableRow, columnMap ColumnIndex, column matter.TableColumn, parentDataType *matter.DataType) matter.Constraint {
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
		slog.Debug("unexpected non-paragraph in constraints cell", "type", fmt.Sprintf("%T", cell.Elements[0]))
		return nil
	}
	if len(p.Elements) == 0 {
		return nil
	}
	var val matter.Constraint
	if len(p.Elements) == 1 {
		switch e := p.Elements[0].(type) {
		case *types.StringElement:
			val = constraint.ParseConstraint(e.Content)
		case *types.InternalCrossReference:
			anchor, _ := d.getAnchor(e.ID.(string))
			var name string
			if anchor != nil {
				name = ReferenceName(anchor.Element)
			} else {
				name = strings.TrimPrefix(e.ID.(string), "_")
			}
			val = constraint.ParseConstraint(StripTypeSuffixes(name))
		}
	} else {
		var sb strings.Builder
		for _, el := range p.Elements {
			switch v := el.(type) {
			case *types.StringElement:
				sb.WriteString(v.Content)
			case *types.InternalCrossReference:
				anchor, _ := d.getAnchor(v.ID.(string))
				var name string
				if anchor != nil {
					name = ReferenceName(anchor.Element)
				} else {
					name = strings.TrimPrefix(v.ID.(string), "_")
				}
				sb.WriteString(name)
			case *types.QuotedText:
				// This is usually an asterisk, and should be ignored
			default:
				slog.Warn("unknown constraint value element", "doc", d.Path, "type", fmt.Sprintf("%T", el))
			}
		}
		val = constraint.ParseConstraint(StripTypeSuffixes(sb.String()))
	}
	return val
}

func (d *Doc) getRowConformance(row *types.TableRow, columnMap ColumnIndex, column matter.TableColumn) matter.Conformance {
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
		slog.Debug("unexpected non-paragraph in constraints cell", "doc", d.Path, "type", fmt.Sprintf("%T", cell.Elements[0]))
		return nil
	}
	if len(p.Elements) == 0 {
		return nil
	}
	var val matter.Conformance
	if len(p.Elements) == 1 {
		switch e := p.Elements[0].(type) {
		case *types.StringElement:
			val = conformance.ParseConformance(e.Content)
		case *types.InternalCrossReference:
			anchor, _ := d.getAnchor(e.ID.(string))
			var name string
			if anchor != nil {
				name = ReferenceName(anchor.Element)
			} else {
				name = strings.TrimPrefix(e.ID.(string), "_")
			}
			val = conformance.ParseConformance(StripTypeSuffixes(name))
		}
	} else {
		var sb strings.Builder
		for _, el := range p.Elements {
			switch v := el.(type) {
			case *types.StringElement:
				sb.WriteString(v.Content)
			case *types.InternalCrossReference:
				anchor, _ := d.getAnchor(v.ID.(string))
				var name string
				if anchor != nil {
					name = ReferenceName(anchor.Element)
				} else {
					name = strings.TrimPrefix(v.ID.(string), "_")
				}
				sb.WriteString(name)
			case *types.SpecialCharacter:
				sb.WriteString(v.Name)
			case *types.QuotedText:
				// This is usually an asterisk, and should be ignored
			default:
				slog.Warn("unknown conformance value element", "doc", d.Path, "type", fmt.Sprintf("%T", el))
			}
		}

		val = conformance.ParseConformance(StripTypeSuffixes(sb.String()))
	}
	return val
}

var typeSuffixes = []string{" Attribute", " Type", " Field", " Command", " Attribute", " Event"}

func StripTypeSuffixes(dataType string) string {
	for _, suffix := range typeSuffixes {
		if strings.HasSuffix(dataType, suffix) {
			dataType = dataType[0 : len(dataType)-len(suffix)]
			break
		}
	}
	return dataType
}
