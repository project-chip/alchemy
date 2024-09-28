package spec

import (
	"fmt"
	"log/slog"
	"regexp"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/internal/parse"
	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/constraint"
	"github.com/project-chip/alchemy/matter/types"
)

func (s *Section) toDataTypes(d *Doc, entityMap map[asciidoc.Attributable][]types.Entity) (bitmaps matter.BitmapSet, enums matter.EnumSet, structs matter.StructSet, err error) {

	traverse(d, s, errata.SpecPurposeDataTypes, func(s *Section, parent parse.HasElements, index int) parse.SearchShould {
		switch s.SecType {
		case matter.SectionDataTypeBitmap:
			var mb *matter.Bitmap
			mb, err = s.toBitmap(d, entityMap)
			if err != nil {
				slog.Warn("Error converting section to bitmap", log.Element("path", d.Path, s.Base), slog.Any("error", err))
				err = nil
			} else {
				bitmaps = append(bitmaps, mb)
			}
		case matter.SectionDataTypeEnum:
			var me *matter.Enum
			me, err = s.toEnum(d, entityMap)
			if err != nil {
				slog.Warn("Error converting section to enum", log.Element("path", d.Path, s.Base), slog.Any("error", err))
				err = nil
			} else {
				enums = append(enums, me)
			}
		case matter.SectionDataTypeStruct:
			var me *matter.Struct
			me, err = s.toStruct(d, entityMap)
			if err != nil {
				slog.Warn("Error converting section to struct", log.Element("path", d.Path, s.Base), slog.Any("error", err))
				err = nil
			} else {
				structs = append(structs, me)
			}
		default:
		}
		return parse.SearchShouldContinue
	})

	return
}

func (d *Doc) readFields(ti *TableInfo, entityType types.EntityType) (fields []*matter.Field, err error) {
	ids := make(map[uint64]*matter.Field)
	for row := range ti.Body() {
		f := matter.NewField(row)
		f.Name, err = ti.ReadValue(row, matter.TableColumnName)
		if err != nil {
			return
		}
		f.Name = matter.StripTypeSuffixes(f.Name)
		f.Conformance = ti.ReadConformance(row, matter.TableColumnConformance)
		f.Type, err = d.ReadRowDataType(row, ti.ColumnMap, matter.TableColumnType)
		if err != nil {
			slog.Debug("error reading field data type", slog.String("path", d.Path.String()), slog.String("name", f.Name), slog.Any("error", err))
			err = nil
		}

		f.Constraint = ti.ReadConstraint(row, matter.TableColumnConstraint)
		if err != nil {
			return
		}
		var q string
		q, err = ti.ReadString(row, matter.TableColumnQuality)
		if err != nil {
			return
		}
		f.Quality = parseQuality(q, entityType, d, row)
		f.Default, err = ti.ReadString(row, matter.TableColumnDefault)
		if err != nil {
			return
		}

		var a string
		a, err = ti.ReadString(row, matter.TableColumnAccess)
		if err != nil {
			return
		}
		f.Access, _ = ParseAccess(a, entityType)
		f.ID, err = ti.ReadID(row, matter.TableColumnID)
		if err != nil {
			return
		}
		if f.ID.Valid() {
			id := f.ID.Value()
			existing, ok := ids[id]
			if ok {
				slog.Error("duplicate field ID", log.Path("source", f), slog.String("name", f.Name), slog.Uint64("id", id), log.Path("original", existing))
				continue
			}
			ids[id] = f
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
		f.Name = CanonicalName(f.Name)
		fields = append(fields, f)
	}
	return
}

var listDataTypeDefinitionPattern = regexp.MustCompile(`(?:list|List|DataTypeList)\[([^\]]+)\]`)
var asteriskPattern = regexp.MustCompile(`\^[0-9]+\^\s*$`)

func (d *Doc) ReadRowDataType(row *asciidoc.TableRow, columnMap ColumnIndex, column matter.TableColumn) (*types.DataType, error) {
	if !d.anchorsParsed {
		d.findAnchors()
	}
	i, ok := columnMap[column]
	if !ok {
		return nil, fmt.Errorf("missing %s column for data type", column)
	}
	cell := row.Cell(i)
	cellElements := cell.Elements()
	if len(cellElements) == 0 {
		return nil, fmt.Errorf("empty %s cell for data type", column)
	}

	var isArray bool

	var sb strings.Builder
	source := d.buildDataTypeString(cellElements, &sb)
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
	name = text.TrimCaseInsensitiveSuffix(name, " Type")
	dt := types.ParseDataType(name, isArray)
	if dt != nil {
		dt.Source = source
	}
	return dt, nil
}

func (d *Doc) buildDataTypeString(cellElements asciidoc.Set, sb *strings.Builder) (source asciidoc.Element) {
	for _, el := range cellElements {
		switch v := el.(type) {
		case *asciidoc.String:
			sb.WriteString(v.Value)
		case *asciidoc.CrossReference:
			if len(v.Set) > 0 {
				d.buildDataTypeString(v.Set, sb)

			} else {
				var name string
				anchor := d.FindAnchor(v.ID)
				if anchor != nil {
					name = ReferenceName(anchor.Element)
					if len(name) == 0 {
						name = asciidoc.AttributeAsciiDocString(anchor.LabelElements)
					}
				} else {
					slog.Warn("data type references unknown or ambiguous anchor", slog.String("name", v.ID), log.Path("source", NewSource(d, v)))
				}
				if len(name) == 0 {
					name = strings.TrimPrefix(v.ID, "_")
				}
				sb.WriteString(name)
			}
			source = el
		case *asciidoc.SpecialCharacter:
		case *asciidoc.Paragraph:
			source = d.buildDataTypeString(v.Elements(), sb)
		default:
			slog.Warn("unknown data type value element", log.Element("path", d.Path, el), "type", fmt.Sprintf("%T", v))
		}
	}
	return
}
