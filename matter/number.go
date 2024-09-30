package matter

import (
	"fmt"
	"math"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/project-chip/alchemy/matter/types"
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

func NumberFromExtreme(de *types.DataTypeExtreme) *Number {
	n := &Number{}
	switch de.Type {
	case types.DataTypeExtremeTypeInt64:
		n.value = de.Int64
		n.format = de.Format
	case types.DataTypeExtremeTypeUInt64:
		n.value = int64(de.UInt64)
		n.format = de.Format
	}
	return n
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

func (n Number) MarshalJSON() ([]byte, error) {
	var val strings.Builder
	val.WriteRune('"')
	switch n.format {
	case types.NumberFormatHex:
		val.WriteString(n.HexString())
	default:
		val.WriteString(n.IntString())
	}
	val.WriteRune('"')
	return []byte(val.String()), nil
}

func (n *Number) UnmarshalJSON(data []byte) error {
	s := string(data)
	if s == "null" || s == `""` {
		return nil
	}
	fn, _ := ParseFormattedNumber(s)
	*n = *fn
	return nil
}

func (n *Number) Equals(oid *Number) bool {
	if n.value >= 0 && oid.value >= 0 {
		return n.value == oid.value
	}
	return false
}

func (n *Number) Compare(oid *Number) int {
	if n.value < oid.value {
		return -1
	} else if n.value > oid.value {
		return 1
	}
	return 0
}

func (n *Number) Clone() *Number {
	return &Number{text: n.text, value: n.value, format: n.format}
}

func (n *Number) Is(oid uint64) bool {
	if n.value >= 0 {
		return int64(oid) == n.value
	}
	return false
}

func (n *Number) Valid() bool {
	return n != nil && n.value >= 0
}

func (n *Number) Value() uint64 {
	return uint64(n.value)
}

func (n *Number) Text() string {
	if n != nil {
		return n.text
	}
	return ""
}

func (n *Number) Format() types.NumberFormat {
	if n != nil {
		return n.format
	}
	return types.NumberFormatUndefined
}

func (n *Number) IntString() string {
	if !n.Valid() {
		return n.Text()
	}
	return strconv.FormatInt(n.value, 10)
}

func (n *Number) HexString() string {
	if !n.Valid() {
		return n.text
	}
	return fmt.Sprintf("0x%04X", n.value)
}

func (n *Number) ShortHexString() string {
	if !n.Valid() {
		return n.text
	}
	return fmt.Sprintf("0x%02X", n.value)
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
	slices.SortStableFunc[[]*Number, *Number](s, func(a *Number, b *Number) int {
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

func NonGlobalIDInvalidForEntity(n *Number, entityType types.EntityType) bool {
	if !n.Valid() {
		// If the number isn't valid, just let it pass through
		return false
	}
	val := n.Value()
	if val > math.MaxUint32 {
		// Matter IDs are 32-bits, max
		return true
	}
	prefix := (val & 0xFFFF0000) >> 16
	var manufacturerCode bool
	if prefix > 0 {
		manufacturerCode = true
	}

	switch entityType {
	case types.EntityTypeDeviceType:
		return val > 0xBFFF
	case types.EntityTypeCluster:
		if manufacturerCode {
			return val < 0xFC00 && val > 0xFFFE
		}
		return val > 0x7FFF
	case types.EntityTypeAttribute:
		return val > 0x4FFF
	case types.EntityTypeCommand:
		if manufacturerCode {
			return val > 0xFF
		}
		return val > 0xDF
	case types.EntityTypeEvent:
		return val > 0xFF
	case types.EntityTypeCommandField, types.EntityTypeEventField, types.EntityTypeStructField:
		return val > 0xDF
	}
	return false
}
