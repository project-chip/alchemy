package parse

import (
	"strconv"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
)

type HasElements interface {
	SetElements(asciidoc.Set) error
	Elements() asciidoc.Set
}

type HasBase interface {
	GetBase() asciidoc.Element
}

func HexOrDec(s string) (uint64, error) {
	id, err := strconv.ParseUint(s, 10, 64)
	if err == nil {
		return id, err
	}

	s = strings.TrimPrefix(s, "0x")
	return strconv.ParseUint(s, 16, 64)
}
