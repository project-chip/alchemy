package matter

import (
	"fmt"
	"strconv"
	"strings"
)

type ID struct {
	text  string
	value int64
}

func NewID(id uint64) *ID {
	return &ID{value: int64(id)}
}

func ParseID(s string) *ID {
	id, err := strconv.ParseUint(s, 10, 64)
	if err == nil {
		return &ID{
			text:  s,
			value: int64(id),
		}
	}

	id, err = strconv.ParseUint(strings.TrimPrefix(s, "0x"), 16, 64)
	if err == nil {
		return &ID{
			text:  s,
			value: int64(id),
		}
	}
	return &ID{
		text:  s,
		value: -1,
	}
}

func (id *ID) Equals(oid *ID) bool {
	if id.value >= 0 && oid.value >= 0 {
		return id.value == oid.value
	}
	return false
}

func (id *ID) Is(oid uint64) bool {
	if id.value >= 0 {
		return int64(oid) == id.value
	}
	return false
}

func (id *ID) Valid() bool {
	return id.value >= 0
}

func (id *ID) Value() uint64 {
	return uint64(id.value)
}

func (id *ID) IntString() string {
	if !id.Valid() {
		return id.text
	}
	return strconv.FormatInt(id.value, 10)
}

func (id *ID) HexString() string {
	if !id.Valid() {
		return id.text
	}
	return fmt.Sprintf("0x%04X", id.value)
}

func (id *ID) ShortHexString() string {
	if !id.Valid() {
		return id.text
	}
	return fmt.Sprintf("0x%02X", id.value)
}

var InvalidID = &ID{value: -1}