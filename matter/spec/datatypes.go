package spec

import (
	"fmt"
	"log/slog"
	"regexp"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/internal/parse"
	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/constraint"
	"github.com/project-chip/alchemy/matter/types"
)

func (s *Section) toDataTypes(d *Doc, entityMap map[asciidoc.Attributable][]types.Entity) (bitmaps matter.BitmapSet, enums matter.EnumSet, structs matter.StructSet, err error) {

	parse.Traverse(s, s.Elements(), func(s *Section, parent parse.HasElements, index int) parse.SearchShould {
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

func (d *Doc) readFields(headerRowIndex int, rows []*asciidoc.TableRow, columnMap ColumnIndex, entityType types.EntityType) (fields []*matter.Field, err error) {
	ids := make(map[uint64]struct{})
	for i := headerRowIndex + 1; i < len(rows); i++ {
		row := rows[i]
		f := matter.NewField(newSource(d, row))
		f.Name, err = ReadRowValue(d, row, columnMap, matter.TableColumnName)
		if err != nil {
			return
		}
		f.Name = matter.StripTypeSuffixes(f.Name)
		f.Conformance = d.getRowConformance(row, columnMap, matter.TableColumnConformance)
		f.Type, err = d.ReadRowDataType(row, columnMap, matter.TableColumnType)
		if err != nil {
			slog.Debug("error reading field data type", slog.String("path", d.Path), slog.String("name", f.Name), slog.Any("error", err))
			err = nil
		}

		f.Constraint = d.getRowConstraint(row, columnMap, matter.TableColumnConstraint)
		if err != nil {
			return
		}
		var q string
		q, err = readRowASCIIDocString(row, columnMap, matter.TableColumnQuality)
		if err != nil {
			return
		}
		f.Quality = matter.ParseQuality(q)
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
				slog.Warn("duplicate field ID", log.Path("source", f.Source), slog.String("name", f.Name), slog.Uint64("id", id))
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
		f.Name = specName(f.Name)
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
	d.buildDataTypeString(cellElements, &sb)
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
	return dt, nil
}

func (d *Doc) buildDataTypeString(cellElements asciidoc.Set, sb *strings.Builder) {
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
					slog.Warn("data type references unknown or ambiguous anchor", slog.String("name", v.ID), log.Path("source", newSource(d, v)))
				}
				if len(name) == 0 {
					name = strings.TrimPrefix(v.ID, "_")
				}
				sb.WriteString(name)
			}
		case *asciidoc.SpecialCharacter:
		case *asciidoc.Paragraph:
			d.buildDataTypeString(v.Elements(), sb)
		default:
			slog.Warn("unknown data type value element", log.Element("path", d.Path, el), "type", fmt.Sprintf("%T", v))
		}
	}
}

func (d *Doc) getRowConstraint(row *asciidoc.TableRow, columnMap ColumnIndex, column matter.TableColumn) constraint.Constraint {
	var val string
	var cell *asciidoc.TableCell
	offset, ok := columnMap[column]
	if ok {
		cell = row.Cell(offset)
		cellElements := cell.Elements()
		if len(cellElements) > 0 {
			var sb strings.Builder
			d.buildConstraintValue(cellElements, &sb)
			val = strings.TrimSpace(sb.String())

		}
	}
	val = strings.ReplaceAll(val, "\n", " ")
	var c constraint.Constraint
	c, err := constraint.ParseString(val)
	if err != nil {
		slog.Warn("failed parsing constraint cell", log.Element("path", d.Path, cell), slog.String("constraint", val))
		return &constraint.GenericConstraint{Value: val}
	}
	return c
}

func (d *Doc) buildConstraintValue(els asciidoc.Set, sb *strings.Builder) {
	for _, el := range els {
		switch v := el.(type) {
		case *asciidoc.String:
			sb.WriteString(v.Value)
		case *asciidoc.CrossReference:
			anchor := d.FindAnchor(v.ID)
			var name string
			if anchor != nil {
				name = matter.StripReferenceSuffixes(ReferenceName(anchor.Element))
			} else {
				name = strings.TrimPrefix(v.ID, "_")
			}
			sb.WriteString(name)
		case *asciidoc.Superscript:
			var qt strings.Builder
			d.buildConstraintValue(v.Elements(), &qt)
			val := qt.String()
			if val == "*" { // We ignore asterisks here
				continue
			}
			sb.WriteString("^")
			sb.WriteString(val)
			sb.WriteString("^")
		case *asciidoc.Bold: // This is usually an asterisk, and should be ignored
		case *asciidoc.NewLine, *asciidoc.LineBreak:
			sb.WriteRune(' ')
		case asciidoc.HasElements:
			d.buildConstraintValue(v.Elements(), sb)
		default:
			slog.Warn("unknown constraint value element", log.Element("path", d.Path, el), "type", fmt.Sprintf("%T", el))
		}
	}
}

func (d *Doc) getRowConformance(row *asciidoc.TableRow, columnMap ColumnIndex, column matter.TableColumn) conformance.Set {
	i, ok := columnMap[column]
	if !ok {
		return nil
	}
	cell := row.Cell(i)
	cellElements := cell.Elements()
	if len(cellElements) == 0 {
		return nil
	}
	var sb strings.Builder
	d.buildRowConformance(cellElements, &sb)

	s := strings.TrimSpace(sb.String())
	if len(s) == 0 {
		return conformance.Set{&conformance.Mandatory{}}
	}
	s = newLineReplacer.Replace(s)
	return conformance.ParseConformance(matter.StripTypeSuffixes(s))
}

func (d *Doc) buildRowConformance(cellElements asciidoc.Set, sb *strings.Builder) {
	for _, el := range cellElements {
		switch v := el.(type) {
		case *asciidoc.String:
			sb.WriteString(v.Value)
		case *asciidoc.CrossReference:
			id := v.ID
			if strings.HasPrefix(id, "ref_") {
				// This is a proper reference; allow the conformance parser to recognize it
				sb.WriteString(fmt.Sprintf("<<%s>>", id))
			} else {
				anchor := d.FindAnchor(v.ID)
				var name string
				if anchor != nil {
					name = ReferenceName(anchor.Element)
				} else {
					name = strings.TrimPrefix(v.ID, "_")
				}
				sb.WriteString(name)
			}
		case *asciidoc.SpecialCharacter:
			sb.WriteString(v.Character)
		case *asciidoc.Superscript:
			// This is usually an asterisk, and should be ignored
		case *asciidoc.Link:
			sb.WriteString(v.URL.Scheme)
			d.buildRowConformance(v.URL.Path, sb)
		case *asciidoc.LinkMacro:
			sb.WriteString("link:")
			sb.WriteString(v.URL.Scheme)
			d.buildRowConformance(v.URL.Path, sb)
		case *asciidoc.CharacterReplacementReference:
			switch v.Name() {
			case "nbsp":
				sb.WriteRune(' ')
			default:
				slog.Warn("unknown predefined attribute", log.Element("path", d.Path, el), "name", v.Name)
			}
		case *asciidoc.NewLine:
			sb.WriteRune(' ')
		case asciidoc.HasElements:
			d.buildRowConformance(v.Elements(), sb)
		default:
			slog.Warn("unknown conformance value element", log.Element("path", d.Path, el), "type", fmt.Sprintf("%T", el))
		}
	}
}

var newLineReplacer = strings.NewReplacer("\r\n", "", "\r", "", "\n", "")
