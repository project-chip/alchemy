package spec

import (
	"fmt"
	"iter"
	"log/slog"
	"regexp"
	"strconv"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/constraint"
	"github.com/project-chip/alchemy/matter/types"
)

type ColumnIndex map[matter.TableColumn]int

func (ci ColumnIndex) HasAny(columns ...matter.TableColumn) bool {
	for _, col := range columns {
		_, ok := ci[col]
		if ok {
			return true
		}
	}
	return false
}

func (ci ColumnIndex) HasAll(columns ...matter.TableColumn) bool {
	if len(columns) == 0 {
		return false
	}
	for _, col := range columns {
		_, ok := ci[col]
		if !ok {
			return false
		}
	}
	return true
}

type ExtraColumn struct {
	Name   string
	Offset int
}

type TableInfo struct {
	Doc            *Doc
	Element        *asciidoc.Table
	Rows           []*asciidoc.TableRow
	HeaderRowIndex int
	ColumnMap      ColumnIndex
	ExtraColumns   []ExtraColumn
}

func (ti *TableInfo) ColumnIndex(columns ...matter.TableColumn) (index int, ok bool) {
	for _, column := range columns {
		index, ok = ti.ColumnMap[column]
		if ok {
			return
		}
	}
	return
}

func (ti *TableInfo) Rescan(doc *Doc) (err error) {
	ti.HeaderRowIndex, ti.ColumnMap, ti.ExtraColumns, err = mapTableColumns(doc, ti.Rows)
	return
}

func (ti *TableInfo) Body() iter.Seq[*asciidoc.TableRow] {
	return func(yield func(*asciidoc.TableRow) bool) {
		for i := ti.HeaderRowIndex + 1; i < len(ti.Rows); i++ {
			if !yield(ti.Rows[i]) {
				return
			}
		}

	}
}

func (ti *TableInfo) ReadString(row *asciidoc.TableRow, columns ...matter.TableColumn) (string, error) {
	for _, column := range columns {
		offset, ok := ti.ColumnMap[column]
		if !ok {
			continue
		}
		cell := row.Cell(offset)
		val, err := RenderTableCell(cell)
		if err != nil {
			return "", err
		}
		val = asteriskPattern.ReplaceAllString(val, "")
		return val, nil
	}
	return "", nil
}

func (ti *TableInfo) ReadStringAtOffset(row *asciidoc.TableRow, offset int) (string, error) {
	cell := row.Cell(offset)
	val, err := RenderTableCell(cell)
	if err != nil {
		return "", err
	}
	val = asteriskPattern.ReplaceAllString(val, "")
	return val, nil
}

func (ti *TableInfo) ReadID(row *asciidoc.TableRow, columns ...matter.TableColumn) (*matter.Number, error) {
	id, err := ti.ReadString(row, columns...)
	if err != nil {
		return matter.InvalidID, err
	}
	return matter.ParseNumber(id), nil
}

func (ti *TableInfo) ReadName(row *asciidoc.TableRow, columns ...matter.TableColumn) (name string, xref *asciidoc.CrossReference, err error) {
	for _, column := range columns {
		offset, ok := ti.ColumnMap[column]
		if !ok {
			continue
		}
		cell := row.Cell(offset)
		cellElements := cell.Elements()
		for _, el := range cellElements {
			switch el := el.(type) {
			case *asciidoc.CrossReference:
				xref = el
			}
			if xref != nil {
				break
			}
		}
		var value strings.Builder
		err = readRowCellValueElements(ti.Doc, cellElements, &value)
		if err != nil {
			return "", nil, err
		}
		return strings.TrimSpace(value.String()), xref, nil
	}
	return "", nil, nil
}

func (ti *TableInfo) ReadValue(row *asciidoc.TableRow, columns ...matter.TableColumn) (string, error) {
	for _, column := range columns {
		offset, ok := ti.ColumnMap[column]
		if !ok {
			continue
		}
		return ti.ReadValueByIndex(row, offset)
	}
	return "", nil
}

func (ti *TableInfo) ReadValueByIndex(row *asciidoc.TableRow, offset int) (string, error) {
	cell := row.Cell(offset)
	cellElements := cell.Elements()
	if len(cellElements) == 0 {
		return "", nil
	}
	var value strings.Builder
	err := readRowCellValueElements(ti.Doc, cellElements, &value)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(value.String()), nil
}

type CellRenderer func(cellElements asciidoc.Set, sb *strings.Builder) (source asciidoc.Element)

func (ti *TableInfo) RenderColumn(row *asciidoc.TableRow, renderer CellRenderer, columns ...matter.TableColumn) (value string, source asciidoc.Element, ok bool) {
	for _, column := range columns {
		var offset int
		offset, ok = ti.ColumnMap[column]
		if !ok {
			continue
		}
		cell := row.Cell(offset)
		cellElements := cell.Elements()
		if len(cellElements) == 0 {
			ok = false
			return
		}
		source = cell
		var sb strings.Builder
		sourceOverride := renderer(cellElements, &sb)
		if sourceOverride != nil {
			source = sourceOverride
		}
		value = sb.String()
		return
	}
	return
}

var newLineReplacer = strings.NewReplacer("\r\n", "", "\r", "", "\n", "")

func (ti *TableInfo) ReadConformance(row *asciidoc.TableRow, column matter.TableColumn) conformance.Set {
	val, source, ok := ti.RenderColumn(row, ti.buildRowConformance, column)
	if !ok {
		return nil
	}

	s := strings.TrimSpace(val)
	if len(s) == 0 {
		return conformance.Set{&conformance.Mandatory{}}
	}
	s = newLineReplacer.Replace(s)
	s = matter.StripTypeSuffixes(s)
	conf := conformance.ParseConformance(s)
	if conformance.IsGeneric(conf) {
		slog.Error("failed parsing conformance cell", log.Element("source", ti.Doc.Path, source), slog.String("value", s))
	}
	return conf
}

func (ti *TableInfo) buildRowConformance(cellElements asciidoc.Set, sb *strings.Builder) (source asciidoc.Element) {
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
				anchor := ti.Doc.FindAnchor(v.ID)
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
			ti.buildRowConformance(v.URL.Path, sb)
		case *asciidoc.LinkMacro:
			sb.WriteString("link:")
			sb.WriteString(v.URL.Scheme)
			ti.buildRowConformance(v.URL.Path, sb)
		case *asciidoc.CharacterReplacementReference:
			switch v.Name() {
			case "nbsp":
				sb.WriteRune(' ')
			default:
				slog.Warn("unknown predefined attribute", log.Element("source", ti.Doc.Path, el), "name", v.Name)
			}
		case *asciidoc.NewLine:
			sb.WriteRune(' ')
		case asciidoc.HasElements:
			ti.buildRowConformance(v.Elements(), sb)
		default:
			slog.Warn("unknown conformance value element", log.Element("source", ti.Doc.Path, el), "type", fmt.Sprintf("%T", el))
		}
	}
	return
}

func (ti *TableInfo) ReadConstraint(row *asciidoc.TableRow, columns ...matter.TableColumn) constraint.Constraint {
	val, source, _ := ti.RenderColumn(row, ti.buildConstraintValue, columns...)
	s := strings.TrimSpace(val)
	s = strings.ReplaceAll(s, "\n", " ")
	var c constraint.Constraint
	c, err := constraint.TryParseString(s)
	if err != nil {
		slog.Error("failed parsing constraint cell", log.Element("source", ti.Doc.Path, source), slog.String("constraint", val))
		return &constraint.GenericConstraint{Value: val}
	}
	return c
}

func (ti *TableInfo) ReadLimit(row *asciidoc.TableRow, columns ...matter.TableColumn) constraint.Limit {
	val, source, _ := ti.RenderColumn(row, ti.buildConstraintValue, columns...)
	s := strings.TrimSpace(val)
	s = strings.ReplaceAll(s, "\n", " ")
	l, err := constraint.TryParseLimit(s)
	if err != nil {
		slog.Error("failed parsing constraint cell", log.Element("source", ti.Doc.Path, source), slog.String("constraint", val))
		return &constraint.GenericLimit{Value: val}
	}
	return l
}

func (ti *TableInfo) buildConstraintValue(els asciidoc.Set, sb *strings.Builder) (source asciidoc.Element) {
	for _, el := range els {
		switch v := el.(type) {
		case *asciidoc.String:
			sb.WriteString(v.Value)
		case *asciidoc.CrossReference:
			anchor := ti.Doc.FindAnchor(v.ID)
			var name string
			if anchor != nil {
				name = matter.StripReferenceSuffixes(ReferenceName(anchor.Element))
			} else {
				name = strings.TrimPrefix(v.ID, "_")
			}
			sb.WriteString(name)
		case *asciidoc.Superscript:
			var qt strings.Builder
			ti.buildConstraintValue(v.Elements(), &qt)
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
			ti.buildConstraintValue(v.Elements(), sb)
		case asciidoc.AttributeReference:
			sb.WriteString(fmt.Sprintf("{%s}", v.Name()))
		default:
			slog.Warn("unknown constraint value element", log.Element("source", ti.Doc.Path, el), "type", fmt.Sprintf("%T", el))
		}
	}
	return
}

func (ti *TableInfo) ReadQuality(row *asciidoc.TableRow, entityType types.EntityType, columns ...matter.TableColumn) (quality matter.Quality, err error) {
	for _, column := range columns {
		offset, ok := ti.ColumnMap[column]
		if !ok {
			continue
		}
		var q string
		q, err = ti.ReadStringAtOffset(row, offset)
		if err != nil {
			return
		}
		quality = parseQuality(q, entityType, ti.Doc, row)
		return
	}
	return
}

func readRowCellValueElements(doc *Doc, els asciidoc.Set, value *strings.Builder) (err error) {
	for _, el := range els {
		switch el := el.(type) {
		case *asciidoc.String:
			value.WriteString(el.Value)
		case asciidoc.FormattedTextElement:
			err = readRowCellValueElements(doc, el.Elements(), value)
		case *asciidoc.Paragraph:
			err = readRowCellValueElements(doc, el.Elements(), value)
		case *asciidoc.CrossReference:
			if len(el.Set) > 0 {
				err = readRowCellValueElements(doc, el.Set, value)
				if err != nil {
					return
				}
			} else {
				var val string
				anchor := doc.FindAnchor(el.ID)
				if anchor != nil {
					val = matter.StripTypeSuffixes(ReferenceName(anchor.Element))
				} else {
					val = strings.TrimPrefix(el.ID, "_")
					val = strings.TrimPrefix(val, "ref_") // Trim, and hope someone else has it defined
				}
				value.WriteString(val)
			}
		case *asciidoc.Link:
			value.WriteString(el.URL.Scheme)
			err = readRowCellValueElements(doc, el.URL.Path, value)
		case *asciidoc.LinkMacro:
			value.WriteString(el.URL.Scheme)
			err = readRowCellValueElements(doc, el.URL.Path, value)
		case *asciidoc.Superscript:
			// In the special case of superscript elements, we do checks to make sure it's not an asterisk or a footnote, which should be ignored
			var quotedText strings.Builder
			err = readRowCellValueElements(doc, el.Elements(), &quotedText)
			if err != nil {
				return
			}
			qt := quotedText.String()
			if qt == "*" { // It's an asterisk, so we'll ignore it
				continue
			}
			_, parseErr := strconv.Atoi(qt)
			if parseErr == nil {
				// This is probably a footnote
				// The similar buildConstraintValue method does not do this, as there are exponential values in contraints
				continue
			}
			value.WriteString(qt)
		case *asciidoc.SpecialCharacter:
			value.WriteString(el.Character)
		case *asciidoc.InlinePassthrough:
			value.WriteString("+")
			err = readRowCellValueElements(doc, el.Elements(), value)
		case *asciidoc.InlineDoublePassthrough:
			value.WriteString("++")
			err = readRowCellValueElements(doc, el.Elements(), value)
		case *asciidoc.ThematicBreak:
		case asciidoc.EmptyLine:
		case *asciidoc.NewLine:
			value.WriteString(" ")
		case asciidoc.HasElements:
			err = readRowCellValueElements(doc, el.Elements(), value)
		case *asciidoc.LineBreak:
			value.WriteString(" ")
		default:
			return fmt.Errorf("unexpected type in cell: %T", el)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

var listDataTypeDefinitionPattern = regexp.MustCompile(`(?:list|List|DataTypeList)\[([^]]+)]`)
var asteriskPattern = regexp.MustCompile(`\^[0-9]+\^\s*$`)

func (ti *TableInfo) ReadDataType(row *asciidoc.TableRow, column matter.TableColumn) (*types.DataType, error) {
	if !ti.Doc.anchorsParsed {
		ti.Doc.findAnchors()
	}
	i, ok := ti.ColumnMap[column]
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
	source := buildDataTypeString(ti.Doc, cellElements, &sb)
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

func buildDataTypeString(d *Doc, cellElements asciidoc.Set, sb *strings.Builder) (source asciidoc.Element) {
	for _, el := range cellElements {
		switch v := el.(type) {
		case *asciidoc.String:
			sb.WriteString(v.Value)
		case *asciidoc.CrossReference:
			if len(v.Set) > 0 {
				buildDataTypeString(d, v.Set, sb)

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
			source = buildDataTypeString(d, v.Elements(), sb)
		default:
			slog.Warn("unknown data type value element", log.Element("source", d.Path, el), "type", fmt.Sprintf("%T", v))
		}
	}
	return
}
