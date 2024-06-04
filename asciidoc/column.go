package asciidoc

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type TableColumnWidth int

var TableColumnWidthAuto TableColumnWidth = -1

func (tcv TableColumnWidth) String() string {
	if tcv == TableColumnWidthAuto {
		return "auto"
	}
	return strconv.Itoa(int(tcv))
}

func (tcv TableColumnWidth) AsciiDocString() string {
	if tcv == TableColumnWidthAuto {
		return "~"
	}
	return strconv.Itoa(int(tcv))
}

type TableColumn struct {
	Multiplier      Optional[int]
	HorizontalAlign Optional[TableCellHorizontalAlign]
	VerticalAlign   Optional[TableCellVerticalAlign]
	Width           Optional[TableColumnWidth]
	Percentage      Optional[int]
	Style           Optional[TableCellStyle]
}

func (tcf *TableColumn) Equals(otcf *TableColumn) bool {
	if !tcf.Multiplier.Equals(otcf.Multiplier) {
		return false
	}
	if !tcf.HorizontalAlign.Equals(otcf.HorizontalAlign) {
		return false
	}
	if !tcf.VerticalAlign.Equals(otcf.VerticalAlign) {
		return false
	}
	if !tcf.Width.Equals(otcf.Width) {
		return false
	}
	if !tcf.Percentage.Equals(otcf.Percentage) {
		return false
	}
	if !tcf.Style.Equals(otcf.Style) {
		return false
	}

	return true
}

func NewTableColumn() *TableColumn {
	return &TableColumn{
		Multiplier:      Default(1),
		HorizontalAlign: Default(TableCellHorizontalAlignLeft),
		VerticalAlign:   Default(TableCellVerticalAlignTop),
		Width:           Default(TableColumnWidth(1)),
		Percentage:      Default(0),
		Style:           Default(TableCellStyleDefault),
	}
}

type TableColumnsAttribute struct {
	attribute

	Columns []*TableColumn
}

func (tca *TableColumnsAttribute) AsciiDocString() string {
	var sb strings.Builder
	for i, col := range tca.Columns {
		if i > 0 {
			sb.WriteRune(',')
		}
		if col.Multiplier.IsSet {
			sb.WriteString(strconv.Itoa(col.Multiplier.Value))
			sb.WriteRune('*')
		}
		if col.HorizontalAlign.IsSet {
			sb.WriteString(col.HorizontalAlign.Value.AsciiDocString())
		}
		if col.VerticalAlign.IsSet {
			sb.WriteString(col.VerticalAlign.Value.AsciiDocString())
		}
		if col.Percentage.IsSet {
			sb.WriteString(strconv.Itoa(col.Percentage.Value))
			sb.WriteRune('%')
		}
		if col.Width.IsSet {
			sb.WriteString(col.Width.Value.AsciiDocString())
		}
		if col.Style.IsSet {
			sb.WriteString(col.Style.Value.AsciiDocString())
		}
	}
	return sb.String()
}

func (tca *TableColumnsAttribute) Equals(a Attribute) bool {
	otca, ok := a.(*TableColumnsAttribute)
	if !ok {
		return false
	}
	if len(otca.Columns) != len(tca.Columns) {
		return false
	}
	for i, tc := range tca.Columns {
		otc := otca.Columns[i]
		if !tc.Equals(otc) {
			return false
		}
	}
	return true
}

func (tca *TableColumnsAttribute) Value() any {
	return tca.Columns
}

func (na *TableColumnsAttribute) SetValue(v any) error {
	if v, ok := v.([]*TableColumn); ok {
		na.Columns = v
		return nil
	}
	return fmt.Errorf("invalid type for TableColumnsAttribute: %T", v)
}

func (TableColumnsAttribute) AttributeType() AttributeType {
	return AttributeTypeColumns
}

func (TableColumnsAttribute) QuoteType() AttributeQuoteType {
	return AttributeQuoteTypeDouble
}

func parseColumnAttribute(a *NamedAttribute) (*TableColumnsAttribute, error) {
	cs := strings.TrimSpace(ValueToString(a.Value()))
	if len(cs) == 0 { // An empty cols attribute should be ignored
		return &TableColumnsAttribute{}, nil
	}
	// Sometimes we use the old cols format, where it was just a number indicating how many columns.
	oldFormat, err := strconv.Atoi(cs)
	if err == nil {
		cols := make([]*TableColumn, 0, oldFormat)
		for i := 0; i < int(oldFormat); i++ {
			cols = append(cols, NewTableColumn())
		}
		return &TableColumnsAttribute{Columns: cols}, nil
	}
	ccs := splitColsValue(cs)
	cols := make([]*TableColumn, 0, len(ccs))
	for _, c := range ccs {
		col := NewTableColumn()
		c = strings.TrimSpace(c)
		if len(c) == 0 {
			cols = append(cols, col)
			continue
		}
		matches := columnPattern.FindStringSubmatch(c)
		if matches == nil {
			return nil, fmt.Errorf("invalid column specifier: %s", c)
		}

		for i, s := range matches {
			if s == "" {
				continue
			}
			switch columnPatternMatchMap[i] {
			case columnMatchMultiplier:
				multiplier, err := strconv.Atoi(s)
				if err != nil {
					return nil, fmt.Errorf("invalid multiplier %s: %w", s, err)

				}
				col.Multiplier = One(multiplier)
			case columnMatchHorizontalAlign:
				switch s {
				case "<":
					col.HorizontalAlign = One(TableCellHorizontalAlignLeft)
				case ">":
					col.HorizontalAlign = One(TableCellHorizontalAlignRight)
				case "^":
					col.HorizontalAlign = One(TableCellHorizontalAlignCenter)
				default:
					return nil, fmt.Errorf("invalid horizontal align %s", s)
				}
			case columnMatchVerticalAlign:
				switch s {
				case ".<":
					col.VerticalAlign = One(TableCellVerticalAlignTop)
				case ".>":
					col.VerticalAlign = One(TableCellVerticalAlignBottom)
				case ".^":
					col.VerticalAlign = One(TableCellVerticalAlignMiddle)
				default:
					return nil, fmt.Errorf("invalid horizontal align %s", s)
				}
			case columnMatchWidth:
				if s == "~" {
					col.Width = One(TableColumnWidthAuto)
				} else {
					width, err := strconv.Atoi(s)
					if err != nil {
						return nil, fmt.Errorf("invalid width %s: %w", s, err)

					}
					col.Width = One(TableColumnWidth(width))
				}
			case columnMatchStyle:
				switch s {
				case "a":
					col.Style = One(TableCellStyleAsciiDoc)
				case "d":
					col.Style = One(TableCellStyleDefault)
				case "e":
					col.Style = One(TableCellStyleEmphasis)
				case "h":
					col.Style = One(TableCellStyleHeader)
				case "l":
					col.Style = One(TableCellStyleLiteral)
				case "m":
					col.Style = One(TableCellStyleMonospace)
				case "s":
					col.Style = One(TableCellStyleStrong)
				default:
					return nil, fmt.Errorf("invalid column style %s", s)
				}
			}

		}
		cols = append(cols, col)
	}

	return &TableColumnsAttribute{Columns: cols}, nil
}

func splitColsValue(val string) (cols []string) {
	for {
		commaIndex := strings.IndexAny(val, ",;")
		if commaIndex == -1 {
			cols = append(cols, val)
			break
		}
		cols = append(cols, val[:commaIndex])
		val = val[commaIndex+1:]
	}
	return
}

type columnMatch uint8

const (
	columnMatchUnknown columnMatch = iota
	columnMatchMultiplier
	columnMatchHorizontalAlign
	columnMatchVerticalAlign
	columnMatchPercentage
	columnMatchWidth
	columnMatchStyle
)

var columnPattern = regexp.MustCompile(`^(?:(?P<Multiplier>[1-9][0-9]*)\*)?(?P<HorizontalAlign><|>|\^)?(?P<VerticalAlign>\.<|\.>|\.\^)?(?:(?P<Percentage>0|[1-9][0-9]{0,1}|100)%)?(?P<Width>[1-9][0-9]*|~)?(?P<Style>[adehlms])?$`)
var columnPatternMatchMap map[int]columnMatch

func init() {
	columnPatternMatchMap = make(map[int]columnMatch)
	for i, name := range columnPattern.SubexpNames() {
		switch name {
		case "":
		case "Multiplier":
			columnPatternMatchMap[i] = columnMatchMultiplier
		case "HorizontalAlign":
			columnPatternMatchMap[i] = columnMatchHorizontalAlign
		case "VerticalAlign":
			columnPatternMatchMap[i] = columnMatchVerticalAlign
		case "Percentage":
			columnPatternMatchMap[i] = columnMatchPercentage
		case "Width":
			columnPatternMatchMap[i] = columnMatchWidth
		case "Style":
			columnPatternMatchMap[i] = columnMatchStyle
		default:
			panic(fmt.Errorf("unknown column pattern name: %s", name))
		}
	}
}
