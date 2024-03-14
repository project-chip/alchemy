package matter

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/hasty/alchemy/matter/types"
)

type Number struct {
	text   string
	value  int64
	format types.NumberFormat
}

func NewNumber(id uint64) *Number {
	return &Number{value: int64(id), format: types.NumberFormatAuto}
}

func ParseNumber(s string) *Number {
	id, _ := ParseFormattedNumber(s)
	return id
}

func ParseFormattedNumber(s string) (*Number, types.NumberFormat) {
	id, err := strconv.ParseUint(s, 10, 64)
	if err == nil {
		return &Number{
			text:   s,
			value:  int64(id),
			format: types.NumberFormatInt,
		}, types.NumberFormatInt
	}

	id, err = strconv.ParseUint(strings.TrimPrefix(s, "0x"), 16, 64)
	if err == nil {
		return &Number{
			text:   s,
			value:  int64(id),
			format: types.NumberFormatHex,
		}, types.NumberFormatHex
	}
	return &Number{
		text:  s,
		value: -1,
	}, types.NumberFormatAuto
}

func (t Number) MarshalJSON() ([]byte, error) {
	var val strings.Builder
	val.WriteRune('"')
	switch t.format {
	case types.NumberFormatHex:
		val.WriteString(t.HexString())
	default:
		val.WriteString(t.IntString())
	}
	val.WriteRune('"')
	return []byte(val.String()), nil
}

func (t *Number) UnmarshalJSON(data []byte) error {
	s := string(data)
	if s == "null" || s == `""` {
		return nil
	}
	n, _ := ParseFormattedNumber(s)
	*t = *n
	return nil
}

func (id *Number) Equals(oid *Number) bool {
	if id.value >= 0 && oid.value >= 0 {
		return id.value == oid.value
	}
	return false
}

func (id *Number) Compare(oid *Number) int {
	if id.value < oid.value {
		return -1
	} else if id.value > oid.value {
		return 1
	}
	return 0
}

func (id *Number) Clone() *Number {
	return &Number{text: id.text, value: id.value, format: id.format}
}

func (id *Number) Is(oid uint64) bool {
	if id.value >= 0 {
		return int64(oid) == id.value
	}
	return false
}

func (id *Number) Valid() bool {
	return id != nil && id.value >= 0
}

func (id *Number) Value() uint64 {
	return uint64(id.value)
}

func (id *Number) Text() string {
	if id != nil {
		return id.text
	}
	return ""
}

func (id *Number) Format() types.NumberFormat {
	if id != nil {
		return id.format
	}
	return types.NumberFormatUndefined
}

func (id *Number) IntString() string {
	if !id.Valid() {
		return id.Text()
	}
	return strconv.FormatInt(id.value, 10)
}

func (id *Number) HexString() string {
	if !id.Valid() {
		return id.text
	}
	return fmt.Sprintf("0x%04X", id.value)
}

func (id *Number) ShortHexString() string {
	if !id.Valid() {
		return id.text
	}
	return fmt.Sprintf("0x%02X", id.value)
}

var InvalidID = &Number{value: -1}

var idRangePattern = regexp.MustCompile(`^(?P<From>0[xX][0-9A-Fa-f]+|[0-9A-Fa-f]+|[0-9]+)\s*(?:\.\.|to)\s*(?P<To>0[xX][0-9A-Fa-f]+|[0-9A-Fa-f]+|[0-9]+)$`)

func ParseIDRange(s string) (from *Number, to *Number) {
	from, _, to, _ = ParseFormattedIDRange(s)
	return
}

func ParseFormattedIDRange(s string) (from *Number, fromFormat types.NumberFormat, to *Number, toFormat types.NumberFormat) {
	match := idRangePattern.FindStringSubmatch(s)
	if len(match) < 3 {
		return InvalidID, types.NumberFormatAuto, InvalidID, types.NumberFormatAuto
	}
	from, fromFormat = ParseFormattedNumber(match[1])
	to, toFormat = ParseFormattedNumber(match[2])
	return
}

func ContainsNumber(s []*Number, n *Number) bool {
	for _, v := range s {
		if v.Equals(n) {
			return true
		}
	}
	return false
}

func SortNumbers(s []*Number) {
	slices.SortFunc[[]*Number, *Number](s, func(a *Number, b *Number) int {
		av := a.Value()
		bv := b.Value()
		if av < bv {
			return -1
		}
		if bv > av {
			return 1
		}
		return 0
	})
}
