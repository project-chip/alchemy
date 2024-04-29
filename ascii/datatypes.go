package ascii

import (
	"fmt"
	"log/slog"
	"regexp"
	"strings"

	"github.com/hasty/adoc/elements"
	"github.com/hasty/alchemy/internal/parse"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/conformance"
	"github.com/hasty/alchemy/matter/constraint"
	"github.com/hasty/alchemy/matter/types"
)

func (s *Section) toDataTypes(d *Doc, entityMap map[elements.Attributable][]types.Entity) (bitmaps matter.BitmapSet, enums matter.EnumSet, structs matter.StructSet, err error) {

	for _, s := range parse.Skim[*Section](s.Elements) {
		switch s.SecType {
		case matter.SectionDataTypeBitmap:
			var mb *matter.Bitmap
			mb, err = s.toBitmap(d, entityMap)
			if err != nil {
				return
			}
			bitmaps = append(bitmaps, mb)
		case matter.SectionDataTypeEnum:
			var me *matter.Enum
			me, err = s.toEnum(d, entityMap)
			if err != nil {
				return
			}
			enums = append(enums, me)
		case matter.SectionDataTypeStruct:
			var me *matter.Struct
			me, err = s.toStruct(d, entityMap)
			if err != nil {
				return
			}
			structs = append(structs, me)
		default:
		}
	}
	return
}

func (d *Doc) readFields(headerRowIndex int, rows []*elements.TableRow, columnMap ColumnIndex, entityType types.EntityType) (fields []*matter.Field, err error) {
	ids := make(map[uint64]struct{})
	for i := headerRowIndex + 1; i < len(rows); i++ {
		row := rows[i]
		f := matter.NewField()
		f.Name, err = ReadRowValue(d, row, columnMap, matter.TableColumnName)
		if err != nil {
			return
		}
		f.Name = matter.StripTypeSuffixes(f.Name)
		f.Conformance = d.getRowConformance(row, columnMap, matter.TableColumnConformance)
		if err != nil {
			return
		}
		f.Type, err = d.ReadRowDataType(row, columnMap, matter.TableColumnType)
		if err != nil {
			slog.Debug("error reading field data type", slog.String("path", d.Path), slog.String("name", f.Name), slog.Any("error", err))
			err = nil
		}

		f.Constraint = d.getRowConstraint(row, columnMap, matter.TableColumnConstraint, f.Type)
		if err != nil {
			return
		}
		var q string
		q, err = readRowASCIIDocString(row, columnMap, matter.TableColumnQuality)
		if err != nil {
			return
		}
		f.Quality = matter.ParseQuality(q)
		if err != nil {
			return
		}
		f.Default, err = readRowASCIIDocString(row, columnMap, matter.TableColumnDefault)
		if err != nil {
			return
		}

		var a string
		a, err = readRowASCIIDocString(row, columnMap, matter.TableColumnAccess)
		if err != nil {
			return
		}
		f.Access, _ = ParseAccess(a, entityType)
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
			var cs constraint.Set
			switch f.Type.BaseType {
			case types.BaseDataTypeMessageID:
				cs = []constraint.Constraint{&constraint.ExactConstraint{Value: &constraint.IntLimit{Value: 16}}}
			case types.BaseDataTypeIPAddress:
				cs = []constraint.Constraint{&constraint.ExactConstraint{Value: &constraint.IntLimit{Value: 4}}, &constraint.ExactConstraint{Value: &constraint.IntLimit{Value: 16}}}
			case types.BaseDataTypeIPv4Address:
				cs = []constraint.Constraint{&constraint.ExactConstraint{Value: &constraint.IntLimit{Value: 4}}}
			case types.BaseDataTypeIPv6Address:
				cs = []constraint.Constraint{&constraint.ExactConstraint{Value: &constraint.IntLimit{Value: 16}}}
			case types.BaseDataTypeIPv6Prefix:
				cs = []constraint.Constraint{&constraint.RangeConstraint{Minimum: &constraint.IntLimit{Value: 1}, Maximum: &constraint.IntLimit{Value: 17}}}
			case types.BaseDataTypeHardwareAddress:
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

func (d *Doc) ReadRowDataType(row *elements.TableRow, columnMap ColumnIndex, column matter.TableColumn) (*types.DataType, error) {
	i, ok := columnMap[column]
	if !ok {
		return nil, fmt.Errorf("missing %s column for data type", column)
	}
	cell := row.TableCells[i]
	cellElements := cell.Elements()
	if len(cellElements) == 0 {
		return nil, fmt.Errorf("empty %s cell for data type", column)
	}
	p, ok := cellElements[0].(*elements.Paragraph)
	if !ok {
		return nil, fmt.Errorf("missing paragraph in %s cell for data type", column)
	}
	pElements := p.Elements()
	if len(pElements) == 0 {
		return nil, fmt.Errorf("empty paragraph in %s cell for data type", column)
	}
	var isArray bool

	var sb strings.Builder
	for _, el := range pElements {
		switch v := el.(type) {
		case *elements.String:
			sb.WriteString(v.Value)
		case *elements.CrossReference:
			var name string
			anchor, _ := d.getAnchor(v.ID)
			if anchor != nil {
				name = ReferenceName(anchor.Element)
				if len(name) == 0 {
					name = anchor.Label
				}
			}
			if len(name) == 0 {
				name = strings.TrimPrefix(v.ID, "_")
			}
			sb.WriteString(name)
		case *elements.SpecialCharacter:
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
	dt := types.ParseDataType(name, isArray)
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

func (d *Doc) getRowConstraint(row *elements.TableRow, columnMap ColumnIndex, column matter.TableColumn, parentDataType *types.DataType) constraint.Constraint {
	var val string
	offset, ok := columnMap[column]
	if ok {
		cell := row.TableCells[offset]
		if len(cell.Elements()) > 0 {
			el := cell.Elements()[0]
			para, ok := el.(*elements.Paragraph)
			if ok {
				var sb strings.Builder
				d.buildConstraintValue(para.Elements(), &sb)
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

func (d *Doc) buildConstraintValue(els []elements.Element, sb *strings.Builder) {
	for _, el := range els {
		switch v := el.(type) {
		case *elements.String:
			sb.WriteString(v.Value)
		case *elements.CrossReference:
			anchor, _ := d.getAnchor(v.ID)
			var name string
			if anchor != nil {
				name = matter.StripReferenceSuffixes(ReferenceName(anchor.Element))
			} else {
				name = strings.TrimPrefix(v.ID, "_")
			}
			sb.WriteString(name)
		case *elements.Superscript:
			var qt strings.Builder
			d.buildConstraintValue(v.Elements(), &qt)
			val := qt.String()
			if val == "*" { // We ignore asterisks here
				continue
			}
			sb.WriteString("^")
			sb.WriteString(val)
			sb.WriteString("^")
		case *elements.Bold:
			// This is usually an asterisk, and should be ignored
		case elements.HasElements:
		default:
			slog.Warn("unknown constraint value element", "doc", d.Path, "type", fmt.Sprintf("%T", el))
		}
	}
}

func (d *Doc) getRowConformance(row *elements.TableRow, columnMap ColumnIndex, column matter.TableColumn) conformance.Set {
	i, ok := columnMap[column]
	if !ok {
		return nil
	}
	cell := row.TableCells[i]
	if len(cell.Elements()) == 0 {
		return nil
	}
	p, ok := cell.Elements()[0].(*elements.Paragraph)
	if !ok {
		slog.Debug("unexpected non-paragraph in constraints cell", "doc", d.Path, "type", fmt.Sprintf("%T", cell.Elements()[0]))
		return nil
	}
	if len(p.Elements()) == 0 {
		return nil
	}

	var sb strings.Builder
	for _, el := range p.Elements() {
		switch v := el.(type) {
		case *elements.String:
			sb.WriteString(v.Value)
		case *elements.CrossReference:
			id := v.ID
			if strings.HasPrefix(id, "ref_") {
				// This is a proper reference; allow the conformance parser to recognize it
				sb.WriteString(fmt.Sprintf("<<%s>>", id))
			} else {
				anchor, _ := d.getAnchor(v.ID)
				var name string
				if anchor != nil {
					name = ReferenceName(anchor.Element)
				} else {
					name = strings.TrimPrefix(v.ID, "_")
				}
				sb.WriteString(name)
			}
		case *elements.SpecialCharacter:
			sb.WriteString(v.Character)
		case *elements.Bold:
			// This is usually an asterisk, and should be ignored
		case *elements.Link:
			if len(v.URL.Scheme) > 0 {
				sb.WriteString(v.URL.Scheme)
			} else {
				sb.WriteString("link:")
			}
			if path, ok := v.URL.Path.(string); ok {
				sb.WriteString(path)
			}

		case *elements.CharacterReplacementReference:
			switch v.Name() {
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
	return conformance.ParseConformance(matter.StripTypeSuffixes(s))
}

var newLineReplacer = strings.NewReplacer("\r\n", "", "\r", "", "\n", "")
