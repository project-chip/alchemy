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

func (ti *TableInfo) Rescan(doc *Doc, reader asciidoc.Reader) (err error) {
	ti.HeaderRowIndex, ti.ColumnMap, ti.ExtraColumns, err = mapTableColumns(doc, reader, ti.Rows)
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

func (ti *TableInfo) ContentRows() iter.Seq[*asciidoc.TableRow] {
	return func(yield func(*asciidoc.TableRow) bool) {
		for i := ti.HeaderRowIndex + 1; i < len(ti.Rows); i++ {
			row := ti.Rows[i]
			if len(row.Elements) > 0 {
				firstCell := row.Cell(0)
				if firstCell.Format.Span.Column.IsSet && firstCell.Format.Span.Column.Value == len(row.Elements) {
					continue
				}
			}
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
		return ti.ReadNameAtOffset(row, offset)
	}
	return "", nil, nil
}

func (ti *TableInfo) ReadNameAtOffset(row *asciidoc.TableRow, offset int) (name string, xref *asciidoc.CrossReference, err error) {
	cell := row.Cell(offset)
	cellElements := cell.Children()
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
	err = readRowCellValueElements(ti.Doc, ti.Doc.Reader(), row, cell, cellElements, &value)
	if err != nil {
		return "", nil, err
	}
	return strings.TrimSpace(value.String()), xref, nil
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
	cellElements := cell.Children()
	if len(cellElements) == 0 {
		return "", nil
	}
	var value strings.Builder
	err := readRowCellValueElements(ti.Doc, ti.Doc.Reader(), row, cell, cellElements, &value)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(value.String()), nil
}

type CellRenderer func(cellElements asciidoc.Elements, sb *strings.Builder) (source asciidoc.Element)

func (ti *TableInfo) RenderColumn(row *asciidoc.TableRow, renderer CellRenderer, columns ...matter.TableColumn) (value string, source asciidoc.Element, ok bool) {
	for _, column := range columns {
		var offset int
		offset, ok = ti.ColumnMap[column]
		if !ok {
			continue
		}
		cell := row.Cell(offset)
		cellElements := cell.Children()
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
	conf := conformance.ParseConformance(s)
	if conformance.IsGeneric(conf) {
		slog.Error("failed parsing conformance cell", log.Element("source", ti.Doc.Path, source), slog.String("value", s))
		ti.Doc.spec.addError(&InvalidConformanceError{Source: row, Conformance: s})

	}
	return conf
}

func (ti *TableInfo) buildRowConformance(cellElements asciidoc.Elements, sb *strings.Builder) (source asciidoc.Element) {
	for _, el := range cellElements {
		switch v := el.(type) {
		case *asciidoc.String:
			sb.WriteString(v.Value)
		case *asciidoc.CrossReference:
			sb.WriteString("<<")
			ti.buildRowConformance(v.ID, sb)
			if !v.Elements.IsWhitespace() {
				sb.WriteString(",")
				ti.buildRowConformance(v.Elements, sb)
			}
			sb.WriteString(">>")
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
		case asciidoc.ParentElement:
			ti.buildRowConformance(v.Children(), sb)
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
		ti.Doc.spec.addError(&InvalidConstraintError{Source: row, Constraint: val})
		return &constraint.GenericConstraint{Value: val}
	}
	return c
}

func (ti *TableInfo) ReadFallback(row *asciidoc.TableRow, columns ...matter.TableColumn) constraint.Limit {
	val, source, _ := ti.RenderColumn(row, ti.buildConstraintValue, columns...)
	s := strings.TrimSpace(val)
	s = strings.ReplaceAll(s, "\n", " ")
	l, err := constraint.TryParseLimit(s)
	if err != nil {
		slog.Error("failed parsing limit cell", log.Element("source", ti.Doc.Path, source), slog.String("limit", val))
		ti.Doc.spec.addError(&InvalidFallbackError{Source: row, Fallback: val})
		return &constraint.GenericLimit{Value: val}
	}
	if l == nil {
		return &constraint.GenericLimit{Value: val}
	}
	return l
}

func (ti *TableInfo) buildConstraintValue(els asciidoc.Elements, sb *strings.Builder) (source asciidoc.Element) {
	for _, el := range els {
		switch v := el.(type) {
		case *asciidoc.String:
			sb.WriteString(v.Value)
		case *asciidoc.CrossReference:
			sb.WriteString("<<")
			ti.buildConstraintValue(v.ID, sb)
			if !v.Elements.IsWhitespace() {
				sb.WriteString(",")
				ti.buildConstraintValue(v.Elements, sb)
			}
			sb.WriteString(">>")
		case *asciidoc.Superscript:
			var qt strings.Builder
			ti.buildConstraintValue(v.Children(), &qt)
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
		case *asciidoc.Monospace:
			sb.WriteRune('`')
			ti.buildConstraintValue(v.Children(), sb)
			sb.WriteRune('`')
		case asciidoc.ParentElement:
			ti.buildConstraintValue(v.Children(), sb)
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

func (ti *TableInfo) ReadLocation(row *asciidoc.TableRow, columns ...matter.TableColumn) (relation matter.DeviceTypeRequirementLocation, err error) {
	var rs string
	rs, err = ti.ReadString(row, columns...)
	if err != nil {
		return
	}
	switch rs {
	case "":
		relation = matter.DeviceTypeRequirementLocationUnknown
	case "Self":
		relation = matter.DeviceTypeRequirementLocationDeviceEndpoint
	case "Child":
		relation = matter.DeviceTypeRequirementLocationChildEndpoint
	case "Root":
		relation = matter.DeviceTypeRequirementLocationRootEndpoint
	case "Descendant":
		relation = matter.DeviceTypeRequirementLocationDescendantEndpoint
	default:
		err = newGenericParseError(row, "unknown location: %s", rs)
	}
	return
}

func readRowCellValueElements(doc *Doc, reader asciidoc.Reader, row *asciidoc.TableRow, parent asciidoc.Parent, els asciidoc.Elements, value *strings.Builder) (err error) {
	for el := range reader.Iterate(parent, els) {
		switch el := el.(type) {
		case *asciidoc.String:
			value.WriteString(el.Value)
		case asciidoc.FormattedTextElement:
			err = readRowCellValueElements(doc, reader, row, el, el.Children(), value)
		case *asciidoc.Paragraph:
			err = readRowCellValueElements(doc, reader, row, el, el.Children(), value)
		case *asciidoc.CrossReference:
			if len(el.Elements) > 0 {
				err = readRowCellValueElements(doc, reader, row, el, el.Elements, value)
				if err != nil {
					return
				}
			} else {
				var val string
				anchor := doc.FindAnchorByID(el.ID, el, el)
				if anchor != nil {
					val = matter.StripTypeSuffixes(ReferenceName(reader, anchor.Element))
				} else {
					val = strings.TrimPrefix(val, "ref_") // Trim, and hope someone else has it defined
				}
				value.WriteString(val)
			}
		case *asciidoc.Link:
			value.WriteString(el.URL.Scheme)
			err = readRowCellValueElements(doc, reader, row, &el.URL.Path, el.URL.Path, value)
		case *asciidoc.LinkMacro:
			value.WriteString(el.URL.Scheme)
			err = readRowCellValueElements(doc, reader, row, &el.URL.Path, el.URL.Path, value)
		case *asciidoc.Superscript:
			// In the special case of superscript elements, we do checks to make sure it's not an asterisk or a footnote, which should be ignored
			var quotedText strings.Builder
			err = readRowCellValueElements(doc, reader, row, el, el.Children(), &quotedText)
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
			err = readRowCellValueElements(doc, reader, row, el, el.Children(), value)
		case *asciidoc.InlineDoublePassthrough:
			value.WriteString("++")
			err = readRowCellValueElements(doc, reader, row, el, el.Children(), value)
		case *asciidoc.ThematicBreak:
		case *asciidoc.EmptyLine:
		case *asciidoc.NewLine:
			value.WriteString(" ")
		case asciidoc.ParentElement:
			err = readRowCellValueElements(doc, reader, row, el, el.Children(), value)
		case *asciidoc.LineBreak:
			value.WriteString(" ")
		default:
			return newGenericParseError(row, "unexpected element type in cell: %T", el)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

var listDataTypeDefinitionPattern = regexp.MustCompile(`(?:list|List|DataTypeList)\[([^]]+)]`)
var listDataTypeEntryPattern = regexp.MustCompile(`\[([^]]+)]`)
var asteriskPattern = regexp.MustCompile(`\^[0-9]+\^\s*$`)

func crossReferenceToDataType(doc *Doc, cr *asciidoc.CrossReference, isArray bool) *types.DataType {
	var dt *types.DataType
	id := doc.anchorId(doc.Reader(), cr, cr, cr.ID)
	if len(cr.Elements) > 0 {
		var label strings.Builder
		for el := range doc.Reader().Iterate(cr, cr.Children()) {
			switch el := el.(type) {
			case *asciidoc.String:
				label.WriteString(el.Value)
			default:
				slog.Warn("unexpected type in cross reference label", log.Type("type", el), log.Path("source", cr))
			}
		}
		id = strings.TrimSpace(label.String())
	}
	id = strings.TrimPrefix(id, "ref_")
	id = strings.TrimPrefix(id, "DataType")
	baseType, nameOverride := types.ParseDataTypeName(id)
	switch baseType {
	case types.BaseDataTypeCustom:
		if nameOverride != "" {
			dt = &types.DataType{
				BaseType: baseType,
				Name:     nameOverride,
			}
		} else {
			dt = &types.DataType{
				BaseType: types.BaseDataTypeCustom,
				Source:   cr,
			}
		}
	default:
		dt = &types.DataType{
			BaseType: baseType,
			Name:     id,
		}
	}
	if isArray {
		return &types.DataType{
			BaseType:  types.BaseDataTypeList,
			Name:      "list",
			EntryType: dt,
		}
	}
	return dt
}

type dataTypePattern func(doc *Doc, row *asciidoc.TableRow, elements []asciidoc.Element) (*types.DataType, bool)

func simpleDataTypePattern(doc *Doc, row *asciidoc.TableRow, elements []asciidoc.Element) (dt *types.DataType, empty bool) {
	if len(elements) != 1 {
		return
	}
	switch el := elements[0].(type) {
	case *asciidoc.String:
		if el.Value == "" {
			empty = true
			return
		}
		var name string
		var isArray bool
		var content = asteriskPattern.ReplaceAllString(el.Value, "")
		match := listDataTypeDefinitionPattern.FindStringSubmatch(content)
		if match != nil {
			name = match[1]
			isArray = true
		} else {
			name = content
		}

		dt = types.ParseDataType(name, isArray)
		if dt == nil {
			slog.Warn("unable to parse data type", slog.String("dataType", el.Value), log.Path("source", row))
		}
	case *asciidoc.CrossReference:
		dt = crossReferenceToDataType(doc, el, false)
	default:
		slog.Warn("unexpected type in data type cell", log.Type("type", el), log.Path("source", row))
	}
	return
}

func listDataTypePattern(doc *Doc, row *asciidoc.TableRow, elements []asciidoc.Element) (dt *types.DataType, empty bool) {
	switch len(elements) {
	case 2:
		listTypeElement, ok := elements[0].(*asciidoc.CrossReference)
		if !ok {
			slog.Warn("unexpected type in list data type cell", log.Type("type", elements[1]), log.Path("source", row))
			return
		}
		listType := crossReferenceToDataType(doc, listTypeElement, false)
		if !listType.IsArray() {
			slog.Warn("unexpected non-list type in list data type cell", slog.String("listType", listType.Name), slog.Any("id", listTypeElement.ID), log.Type("type", elements[0]), log.Path("source", row))
			return
		}
		switch el := elements[1].(type) {
		case *asciidoc.String:
			if el.Value == "" {
				empty = true
				return
			}
			var name string
			var content = asteriskPattern.ReplaceAllString(el.Value, "")
			match := listDataTypeEntryPattern.FindStringSubmatch(content)
			if match != nil {
				name = match[1]
				dt = types.ParseDataType(name, true)
				if dt == nil {
					slog.Warn("unable to parse data type", slog.String("dataType", el.Value), log.Path("source", row))
				}
			}
		default:
			slog.Warn("unexpected type in list data type cell", log.Type("type", el), log.Path("source", row))
		}
	case 3:
		if !strings.EqualFold(asciidoc.StringValue(elements[0]), "list[") {
			return
		}
		if !strings.EqualFold(asciidoc.StringValue(elements[2]), "]") {
			return
		}
		switch el := elements[1].(type) {
		case *asciidoc.String:
			dt = types.ParseDataType(el.Value, true)
		case *asciidoc.CrossReference:
			return crossReferenceToDataType(doc, el, true), false
		default:
			slog.Warn("unexpected type in list data type cell", log.Type("type", el), log.Path("source", row))
		}
	case 4:
		if !strings.EqualFold(asciidoc.StringValue(elements[1]), "[") {
			return
		}
		if !strings.EqualFold(asciidoc.StringValue(elements[3]), "]") {
			return
		}
		listTypeElement, ok := elements[0].(*asciidoc.CrossReference)
		if !ok {
			slog.Warn("unexpected type in list data type cell", log.Type("type", elements[1]), log.Path("source", row))
			return
		}
		listType := crossReferenceToDataType(doc, listTypeElement, false)
		if !listType.IsArray() {
			slog.Warn("unexpected non-list type in list data type cell", slog.String("listType", listType.Name), slog.Any("id", listTypeElement.ID), log.Type("type", elements[0]), log.Path("source", row))
			return
		}

		switch el := elements[2].(type) {
		case *asciidoc.String:
			dt = types.ParseDataType(el.Value, true)
		case *asciidoc.CrossReference:
			listType.EntryType = crossReferenceToDataType(doc, el, false)
			dt = listType
		default:
			slog.Warn("unexpected type in list data type cell", log.Type("type", elements[2]), log.Path("source", row))
			return
		}
	default:
		slog.Warn("unexpected number of elements in list data type cell", slog.Int("count", len(elements)), log.Path("source", row))
	}
	return
}

func (ti *TableInfo) ReadDataType(reader asciidoc.Reader, row *asciidoc.TableRow, column matter.TableColumn) (*types.DataType, error) {
	if !ti.Doc.anchorsParsed {
		ti.Doc.findAnchors()
	}
	i, ok := ti.ColumnMap[column]
	if !ok {
		return nil, newGenericParseError(row, "missing %s column for data type", column)
	}
	cell := row.Cell(i)
	cellElements := reader.Iterate(cell, cell.Children()).List()

	if len(cellElements) == 0 {
		return nil, newGenericParseError(row, "empty %s cell for data type", column)
	}

	var dt *types.DataType
	var dataTypePatterns = []dataTypePattern{simpleDataTypePattern, listDataTypePattern}
	for _, pattern := range dataTypePatterns {
		var empty bool
		dt, empty = pattern(ti.Doc, row, cellElements)
		if dt != nil || empty {
			return dt, nil
		}
	}

	slog.Warn("no matching data type patterns", log.Path("source", row))
	return dt, nil
}

func buildDataTypeString(d *Doc, cellElements asciidoc.Elements, sb *strings.Builder) (source asciidoc.Element) {
	for _, el := range cellElements {
		switch v := el.(type) {
		case *asciidoc.String:
			sb.WriteString(v.Value)
		case *asciidoc.CrossReference:
			if len(v.Elements) > 0 {
				buildDataTypeString(d, v.Elements, sb)
			} else {
				var name string
				anchor := d.FindAnchorByID(v.ID, v, v)
				if anchor != nil {
					name = ReferenceName(d.Reader(), anchor.Element)
					if len(name) == 0 {
						name = asciidoc.AttributeAsciiDocString(anchor.LabelElements)
					}
				} else {
					slog.Warn("data type references unknown or ambiguous anchor", slog.Any("name", v.ID), log.Path("source", NewSource(d, v)))
				}
				if len(name) == 0 {
					name = d.anchorId(d.Reader(), v, v, v.ID)
				}
				sb.WriteString(name)
			}
			source = el
		case *asciidoc.SpecialCharacter:
		case *asciidoc.Paragraph:
			source = buildDataTypeString(d, v.Children(), sb)
		default:
			slog.Warn("unknown data type value element", log.Element("source", d.Path, el), "type", fmt.Sprintf("%T", v))
		}
	}
	return
}
