package ascii

import (
	"fmt"
	"log/slog"
	"regexp"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/alchemy/internal/parse"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/conformance"
	"github.com/hasty/alchemy/matter/constraint"
	mattertypes "github.com/hasty/alchemy/matter/types"
)

func (s *Section) toDataTypes(d *Doc, cluster *matter.Cluster, entityMap map[types.WithAttributes][]mattertypes.Entity) (err error) {

	for _, s := range parse.Skim[*Section](s.Elements) {
		switch s.SecType {
		case matter.SectionDataTypeBitmap:
			var mb *matter.Bitmap
			mb, err = s.toBitmap(d, entityMap)
			if err != nil {
				return
			}
			cluster.Bitmaps = append(cluster.Bitmaps, mb)
		case matter.SectionDataTypeEnum:
			var me *matter.Enum
			me, err = s.toEnum(d, entityMap)
			if err != nil {
				return
			}
			cluster.Enums = append(cluster.Enums, me)
		case matter.SectionDataTypeStruct:
			var me *matter.Struct
			me, err = s.toStruct(d, cluster, entityMap)
			if err != nil {
				return
			}
			cluster.Structs = append(cluster.Structs, me)
		default:
		}
	}
	return
}

func (d *Doc) readFields(cluster *matter.Cluster, headerRowIndex int, rows []*types.TableRow, columnMap ColumnIndex, entityType mattertypes.EntityType) (fields []*matter.Field, err error) {
	ids := make(map[uint64]struct{})
	for i := headerRowIndex + 1; i < len(rows); i++ {
		row := rows[i]
		f := matter.NewField()
		f.Name, err = readRowValue(d, row, columnMap, matter.TableColumnName)
		if err != nil {
			return
		}
		f.Name = StripTypeSuffixes(f.Name)
		f.Conformance = d.getRowConformance(row, columnMap, matter.TableColumnConformance)
		if err != nil {
			return
		}
		f.Type, err = d.ReadRowDataType(row, columnMap, matter.TableColumnType)
		if err != nil {
			if cluster != nil && cluster.Hierarchy == "Base" && !conformance.IsDeprecated(f.Conformance) {
				var clusterName string
				if cluster != nil {
					clusterName = cluster.Name
				}
				slog.Warn("error reading field data type", slog.String("path", d.Path), slog.String("cluster", clusterName), slog.String("name", f.Name), slog.Any("error", err))
			}
			err = nil
		}

		f.Constraint = d.getRowConstraint(row, columnMap, matter.TableColumnConstraint, f.Type)
		if err != nil {
			return
		}
		var q string
		q, err = readRowAsciiDocString(row, columnMap, matter.TableColumnQuality)
		if err != nil {
			return
		}
		f.Quality = matter.ParseQuality(q)
		if err != nil {
			return
		}
		f.Default, err = readRowAsciiDocString(row, columnMap, matter.TableColumnDefault)
		if err != nil {
			return
		}

		var a string
		a, err = readRowAsciiDocString(row, columnMap, matter.TableColumnAccess)
		if err != nil {
			return
		}
		f.Access = ParseAccess(a, entityType)
		f.ID, err = readRowID(row, columnMap, matter.TableColumnID)
		if err != nil {
			return
		}
		if f.ID.Valid() {
			id := f.ID.Value()
			_, ok := ids[id]
			if ok {
				slog.Warn("duplicate field ID", "doc", d.Path, "name", f.Name, "id", id)
				continue
			}
			ids[id] = struct{}{}
		}

		if f.Type != nil {
			var cs constraint.ConstraintSet
			switch f.Type.BaseType {
			case mattertypes.BaseDataTypeMessageID:
				cs = []constraint.Constraint{&constraint.ExactConstraint{Value: &constraint.IntLimit{Value: 16}}}
			case mattertypes.BaseDataTypeIPAddress:
				cs = []constraint.Constraint{&constraint.ExactConstraint{Value: &constraint.IntLimit{Value: 4}}, &constraint.ExactConstraint{Value: &constraint.IntLimit{Value: 16}}}
			case mattertypes.BaseDataTypeIPv4Address:
				cs = []constraint.Constraint{&constraint.ExactConstraint{Value: &constraint.IntLimit{Value: 4}}}
			case mattertypes.BaseDataTypeIPv6Address:
				cs = []constraint.Constraint{&constraint.ExactConstraint{Value: &constraint.IntLimit{Value: 16}}}
			case mattertypes.BaseDataTypeIPv6Prefix:
				cs = []constraint.Constraint{&constraint.RangeConstraint{Minimum: &constraint.IntLimit{Value: 1}, Maximum: &constraint.IntLimit{Value: 17}}}
			case mattertypes.BaseDataTypeHardwareAddress:
				cs = []constraint.Constraint{&constraint.ExactConstraint{Value: &constraint.IntLimit{Value: 6}}, &constraint.ExactConstraint{Value: &constraint.IntLimit{Value: 8}}}
			}
			if cs != nil {
				if f.Type.IsArray() {
					lc, ok := f.Constraint.(*constraint.ListConstraint)
					if ok {
						lc.EntryConstraint = constraint.AppendConstraint(lc.EntryConstraint, cs...)
					}
				} else {
					f.Constraint = constraint.AppendConstraint(f.Constraint, cs...)
				}

			}
		}
		fields = append(fields, f)
	}
	return
}

var listDataTypeDefinitionPattern = regexp.MustCompile(`(?:list|List|DataTypeList)\[([^\]]+)\]`)
var asteriskPattern = regexp.MustCompile(`\^[0-9]+\^\s*$`)

func (d *Doc) ReadRowDataType(row *types.TableRow, columnMap ColumnIndex, column matter.TableColumn) (*mattertypes.DataType, error) {
	i, ok := columnMap[column]
	if !ok {
		return nil, fmt.Errorf("missing %s column for data type", column)
	}
	cell := row.Cells[i]
	if len(cell.Elements) == 0 {
		return nil, fmt.Errorf("empty %s cell for data type", column)
	}
	p, ok := cell.Elements[0].(*types.Paragraph)
	if !ok {
		return nil, fmt.Errorf("missing paragraph in %s cell for data type", column)
	}
	if len(p.Elements) == 0 {
		return nil, fmt.Errorf("empty paragraph in %s cell for data type", column)
	}
	var isArray bool

	var sb strings.Builder
	for _, el := range p.Elements {
		switch v := el.(type) {
		case *types.StringElement:
			sb.WriteString(v.Content)
		case *types.InternalCrossReference:
			var name string
			anchor, _ := d.getAnchor(v.ID.(string))
			if anchor != nil {
				name = ReferenceName(anchor.Element)
				if len(name) == 0 {
					name = anchor.Label
				}
			}
			if len(name) == 0 {
				name = strings.TrimPrefix(v.ID.(string), "_")
			}
			sb.WriteString(name)
		case *types.SpecialCharacter:
		default:
			slog.Warn("unknown data type value element", "type", v)
		}
	}
	var name string
	var content = asteriskPattern.ReplaceAllString(sb.String(), "")
	match := listDataTypeDefinitionPattern.FindStringSubmatch(content)
	if match != nil {
		name = match[1]
		isArray = true
	} else {
		name = content
	}
	commaIndex := strings.IndexRune(name, ',')
	if commaIndex >= 0 {
		name = name[:commaIndex]
	}
	name = strings.TrimSuffix(name, " Type")
	dt := mattertypes.ParseDataType(name, isArray)
	return dt, nil
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

func (d *Doc) getRowConstraint(row *types.TableRow, columnMap ColumnIndex, column matter.TableColumn, parentDataType *mattertypes.DataType) constraint.Constraint {
	var val string
	offset, ok := columnMap[column]
	if ok {
		cell := row.Cells[offset]
		if len(cell.Elements) > 0 {
			el := cell.Elements[0]
			para, ok := el.(*types.Paragraph)
			if ok {
				var sb strings.Builder
				d.buildConstraintValue(para.Elements, &sb)
				val = strings.TrimSpace(sb.String())
			}
		}
	}
	val = strings.ReplaceAll(val, "\n", " ")
	var c constraint.Constraint
	c, err := constraint.ParseString(val)
	if err != nil {
		slog.Warn("failed parsing constraint cell", "path", d.Path, slog.String("constraint", val))
		return &constraint.GenericConstraint{Value: val}
	}
	return c
}

func (d *Doc) buildConstraintValue(elements []any, sb *strings.Builder) {
	for _, el := range elements {
		switch v := el.(type) {
		case *types.StringElement:
			sb.WriteString(v.Content)
		case *types.InternalCrossReference:
			anchor, _ := d.getAnchor(v.ID.(string))
			var name string
			if anchor != nil {
				name = matter.StripReferenceSuffixes(ReferenceName(anchor.Element))
			} else {
				name = strings.TrimPrefix(v.ID.(string), "_")
			}
			sb.WriteString(name)
		case *types.QuotedText:
			switch v.Kind {
			case types.SingleQuoteSuperscript:
				var qt strings.Builder
				d.buildConstraintValue(v.Elements, &qt)
				val := qt.String()
				if val == "*" { // We ignore asterisks here
					continue
				}
				sb.WriteString(string(v.Kind))
				sb.WriteString(val)
				sb.WriteString(string(v.Kind))
			case types.SingleQuoteBold:
				// This is usually an asterisk, and should be ignored
			default:
				slog.Warn("unexpected constraint quoted text", "doc", d.Path, "type", fmt.Sprintf("%T", v.Kind))
			}

		default:
			slog.Warn("unknown constraint value element", "doc", d.Path, "type", fmt.Sprintf("%T", el))
		}
	}
}

func (d *Doc) getRowConformance(row *types.TableRow, columnMap ColumnIndex, column matter.TableColumn) conformance.Set {
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

	var sb strings.Builder
	for _, el := range p.Elements {
		switch v := el.(type) {
		case *types.StringElement:
			sb.WriteString(v.Content)
		case *types.InternalCrossReference:
			id := v.ID.(string)
			if strings.HasPrefix(id, "ref_") {
				// This is a proper reference; allow the conformance parser to recognize it
				sb.WriteString(fmt.Sprintf("<<%s>>", id))
			} else {
				anchor, _ := d.getAnchor(v.ID.(string))
				var name string
				if anchor != nil {
					name = ReferenceName(anchor.Element)
				} else {
					name = strings.TrimPrefix(v.ID.(string), "_")
				}
				sb.WriteString(name)
			}
		case *types.SpecialCharacter:
			sb.WriteString(v.Name)
		case *types.QuotedText:
			// This is usually an asterisk, and should be ignored
		case *types.InlineLink:
			if v.Location != nil {
				if len(v.Location.Scheme) > 0 {
					sb.WriteString(v.Location.Scheme)
				} else {
					sb.WriteString("link:")
				}
				if path, ok := v.Location.Path.(string); ok {
					sb.WriteString(path)
				}
			}
		case *types.PredefinedAttribute:
			switch v.Name {
			case "nbsp":
				sb.WriteRune(' ')
			default:
				slog.Warn("unknown predefined attribute", "doc", d.Path, "name", v.Name)
			}
		default:
			slog.Warn("unknown conformance value element", "doc", d.Path, "type", fmt.Sprintf("%T", el))
		}
	}

	s := strings.TrimSpace(sb.String())
	if len(s) == 0 {
		return conformance.Set{&conformance.Mandatory{}}
	}
	s = newLineReplacer.Replace(s)
	return conformance.ParseConformance(StripTypeSuffixes(s))
}

var newLineReplacer = strings.NewReplacer("\r\n", "", "\r", "", "\n", "")

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
