package parse

import (
	"strconv"
	"strings"
)

type HasElements interface {
	SetElements([]interface{}) error
	GetElements() []interface{}
}

type HasBase interface {
	GetBase() interface{}
}

func HexOrDec(s string) (uint64, error) {
	id, err := strconv.ParseUint(s, 10, 64)
	if err == nil {
		return id, err
	}

	s = strings.TrimPrefix(s, "0x")
	return strconv.ParseUint(s, 16, 64)
}
