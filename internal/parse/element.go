package parse

import (
	"strconv"
	"strings"

	"github.com/hasty/adoc/elements"
)

type HasElements interface {
	SetElements(elements.Set) error
	Elements() elements.Set
}

type HasBase interface {
	GetBase() elements.Element
}

func HexOrDec(s string) (uint64, error) {
	id, err := strconv.ParseUint(s, 10, 64)
	if err == nil {
		return id, err
	}

	s = strings.TrimPrefix(s, "0x")
	return strconv.ParseUint(s, 16, 64)
}
