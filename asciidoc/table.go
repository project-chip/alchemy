package asciidoc

import (
	"fmt"
	"strconv"
	"strings"
)

type TableCellHorizontalAlign uint8

const (
	TableCellHorizontalAlignLeft TableCellHorizontalAlign = iota
	TableCellHorizontalAlignRight
	TableCellHorizontalAlignCenter
)

func (ha TableCellHorizontalAlign) String() string {
	switch ha {
	case TableCellHorizontalAlignLeft:
		return "left"
	case TableCellHorizontalAlignCenter:
		return "center"
	case TableCellHorizontalAlignRight:
		return "right"
	default:
		return "unknown"
	}
}

func (ha TableCellHorizontalAlign) AsciiDocString() string {
	switch ha {
	case TableCellHorizontalAlignLeft:
		return "<"
	case TableCellHorizontalAlignCenter:
		return "^"
	case TableCellHorizontalAlignRight:
		return ">"
	default:
		return ""
	}
}

type TableCellVerticalAlign uint8

const (
	TableCellVerticalAlignTop TableCellVerticalAlign = iota
	TableCellVerticalAlignBottom
	TableCellVerticalAlignMiddle
)

func (ha TableCellVerticalAlign) String() string {
	switch ha {
	case TableCellVerticalAlignTop:
		return "top"
	case TableCellVerticalAlignBottom:
		return "bottom"
	case TableCellVerticalAlignMiddle:
		return "middle"
	default:
		return "unknown"
	}
}

func (ha TableCellVerticalAlign) AsciiDocString() string {
	switch ha {
	case TableCellVerticalAlignTop:
		return ".<"
	case TableCellVerticalAlignMiddle:
		return ".^"
	case TableCellVerticalAlignBottom:
		return ".>"
	default:
		return ""
	}
}

type TableCellStyle uint8

const (
	TableCellStyleDefault TableCellStyle = iota
	TableCellStyleAsciiDoc
	TableCellStyleEmphasis
	TableCellStyleHeader
	TableCellStyleLiteral
	TableCellStyleMonospace
	TableCellStyleStrong
)

func (ha TableCellStyle) String() string {
	switch ha {
	case TableCellStyleDefault:
		return "default"
	case TableCellStyleAsciiDoc:
		return "asciidoc"
	case TableCellStyleEmphasis:
		return "emphasis"
	case TableCellStyleHeader:
		return "header"
	case TableCellStyleLiteral:
		return "literal"
	case TableCellStyleMonospace:
		return "monospace"
	case TableCellStyleStrong:
		return "strong"
	default:
		return "unknown"
	}
}

func (ha TableCellStyle) AsciiDocString() string {
	switch ha {
	case TableCellStyleDefault:
		return "d"
	case TableCellStyleAsciiDoc:
		return "a"
	case TableCellStyleEmphasis:
		return "e"
	case TableCellStyleHeader:
		return "h"
	case TableCellStyleLiteral:
		return "l"
	case TableCellStyleMonospace:
		return "m"
	case TableCellStyleStrong:
		return "s"
	default:
		return ""
	}
}

type TableCellSpan struct {
	Column Optional[int]
	Row    Optional[int]
}

func NewTableCellSpan() TableCellSpan {
	return TableCellSpan{
		Column: Default(1),
		Row:    Default(1),
	}
}

type TableCellFormat struct {
	Multiplier      Optional[int]
	Span            TableCellSpan
	HorizontalAlign Optional[TableCellHorizontalAlign]
	VerticalAlign   Optional[TableCellVerticalAlign]
	Style           Optional[TableCellStyle]
}

func NewTableCellFormat() *TableCellFormat {
	return &TableCellFormat{
		Multiplier:      Default(1),
		Span:            NewTableCellSpan(),
		HorizontalAlign: Default(TableCellHorizontalAlignLeft),
		VerticalAlign:   Default(TableCellVerticalAlignTop),
		Style:           Default(TableCellStyleDefault),
	}
}

func (tcf *TableCellFormat) Equals(otcf *TableCellFormat) bool {
	if !tcf.Multiplier.Equals(otcf.Multiplier) {
		return false
	}
	if !tcf.HorizontalAlign.Equals(otcf.HorizontalAlign) {
		return false
	}
	if !tcf.VerticalAlign.Equals(otcf.VerticalAlign) {
		return false
	}
	if !tcf.Style.Equals(otcf.Style) {
		return false
	}
	if !tcf.Span.Column.Equals(otcf.Span.Column) {
		return false
	}
	if !tcf.Span.Row.Equals(otcf.Span.Row) {
		return false
	}
	return true
}

func (tcf *TableCellFormat) AsciiDocString() string {
	var sb strings.Builder

	if tcf.Multiplier.IsSet {
		sb.WriteString(strconv.Itoa(tcf.Multiplier.Value))
		sb.WriteRune('*')
	}
	if tcf.Span.Column.IsSet || tcf.Span.Row.IsSet {
		if tcf.Span.Column.IsSet {
			sb.WriteString(strconv.Itoa(tcf.Span.Column.Value))
		}
		if tcf.Span.Row.IsSet {
			sb.WriteRune('.')
			sb.WriteString(strconv.Itoa(tcf.Span.Row.Value))
		}
		sb.WriteRune('+')

	}
	if tcf.HorizontalAlign.IsSet {
		sb.WriteString(tcf.HorizontalAlign.Value.AsciiDocString())
	}
	if tcf.VerticalAlign.IsSet {
		sb.WriteString(tcf.VerticalAlign.Value.AsciiDocString())
	}
	if tcf.Style.IsSet {
		sb.WriteString(tcf.Style.Value.AsciiDocString())
	}

	return sb.String()
}

type TableCell struct {
	position
	Format *TableCellFormat
	Parent *TableRow

	Set

	Blank bool
}

func (TableCell) Type() ElementType {
	return ElementTypeBlock
}

func NewTableCell(format *TableCellFormat) *TableCell {
	if format == nil {
		format = NewTableCellFormat()
	}
	return &TableCell{
		Format: format,
	}
}

func (tc *TableCell) Equals(e Element) bool {
	otc, ok := e.(*TableCell)
	if !ok {
		return false
	}
	if otc.Blank != tc.Blank {
		return false
	}
	if otc.Format == nil {
		if tc.Format != nil {
			return false
		}
	} else if tc.Format == nil {
		return false
	}
	if otc.Format != nil && !otc.Format.Equals(tc.Format) {
		return false
	}
	if !tc.Set.Equals(otc.Set) {
		return false
	}
	return true
}

type TableCells []*TableCell

func (trs TableCells) Elements() Set {
	els := make(Set, 0, len(trs))
	for _, tr := range trs {
		els = append(els, tr)
	}
	return els
}

func (trs *TableCells) Append(e Element) error {
	tr, ok := e.(*TableCell)
	if !ok {
		return fmt.Errorf("invalid element for TableCells: %T", e)
	}
	*trs = append(*trs, tr)
	return nil
}

func (trs *TableCells) SetElements(els Set) error {
	ntrs := make([]*TableCell, 0, len(els))
	for _, e := range els {
		tr, ok := e.(*TableCell)
		if !ok {
			return fmt.Errorf("invalid element for TableCells: %T", e)
		}
		ntrs = append(ntrs, tr)
	}
	*trs = ntrs
	return nil
}

type TableRow struct {
	position
	Parent *Table

	Set
}

func (TableRow) Type() ElementType {
	return ElementTypeBlock
}

func (tr *TableRow) Equals(e Element) bool {
	otr, ok := e.(*TableRow)
	if !ok {
		return false
	}
	return tr.Set.Equals(otr.Set)
}

func (tr *TableRow) TableCells() []*TableCell {
	tcs := make([]*TableCell, 0, len(tr.Set))
	for _, el := range tr.Set {
		if tc, ok := el.(*TableCell); ok {
			tcs = append(tcs, tc)
		}
	}
	return tcs
}

func (tr *TableRow) Cell(i int) *TableCell {
	return tr.Set[i].(*TableCell)
}

type TableRows []*TableRow

func (trs TableRows) Elements() Set {
	els := make(Set, 0, len(trs))
	for _, tr := range trs {
		els = append(els, tr)
	}
	return els
}

func (trs *TableRows) Append(e Element) error {
	tr, ok := e.(*TableRow)
	if !ok {
		return fmt.Errorf("invalid element for TableRows: %T", e)
	}
	*trs = append(*trs, tr)
	return nil
}

func (trs *TableRows) SetElements(els Set) error {
	ntrs := make([]*TableRow, 0, len(els))
	for _, e := range els {
		tr, ok := e.(*TableRow)
		if !ok {
			return fmt.Errorf("invalid element for TableRows: %T", e)
		}
		ntrs = append(ntrs, tr)
	}
	*trs = ntrs
	return nil
}

type Table struct {
	position
	AttributeList

	ColumnCount int
	Set
}

func (Table) Type() ElementType {
	return ElementTypeBlock
}

func (t *Table) Equals(e Element) bool {
	ot, ok := e.(*Table)
	if !ok {
		return false
	}
	if t.ColumnCount != ot.ColumnCount {
		return false
	}
	if !t.AttributeList.Equals(ot.AttributeList) {
		return false
	}
	return t.Set.Equals(ot.Set)
}

func (t *Table) TableRows() []*TableRow {
	trs := make([]*TableRow, 0, len(t.Set))
	for _, el := range t.Set {
		if tr, ok := el.(*TableRow); ok {
			trs = append(trs, tr)
		}
	}
	return trs
}
