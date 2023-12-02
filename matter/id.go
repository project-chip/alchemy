package matter

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Number struct {
	text   string
	value  int64
	format NumberFormat
}

func NewID(id uint64) *Number {
	return &Number{value: int64(id), format: NumberFormatAuto}
}

func ParseID(s string) *Number {
	id, _ := ParseFormattedID(s)
	return id
}

func ParseFormattedID(s string) (*Number, NumberFormat) {
	id, err := strconv.ParseUint(s, 10, 64)
	if err == nil {
		return &Number{
			text:   s,
			value:  int64(id),
			format: NumberFormatInt,
		}, NumberFormatInt
	}

	id, err = strconv.ParseUint(strings.TrimPrefix(s, "0x"), 16, 64)
	if err == nil {
		return &Number{
			text:   s,
			value:  int64(id),
			format: NumberFormatHex,
		}, NumberFormatHex
	}
	return &Number{
		text:  s,
		value: -1,
	}, NumberFormatAuto
}

func (id *Number) Equals(oid *Number) bool {
	if id.value >= 0 && oid.value >= 0 {
		return id.value == oid.value
	}
	return false
}

func (id *Number) Is(oid uint64) bool {
	if id.value >= 0 {
		return int64(oid) == id.value
	}
	return false
}

func (id *Number) Valid() bool {
	return id.value >= 0
}

func (id *Number) Value() uint64 {
	return uint64(id.value)
}

func (id *Number) IntString() string {
	if !id.Valid() {
		return id.text
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

func ParseFormattedIDRange(s string) (from *Number, fromFormat NumberFormat, to *Number, toFormat NumberFormat) {
	match := idRangePattern.FindStringSubmatch(s)
	if len(match) < 3 {
		return InvalidID, NumberFormatAuto, InvalidID, NumberFormatAuto
	}
	from, fromFormat = ParseFormattedID(match[1])
	to, toFormat = ParseFormattedID(match[2])
	return
}
