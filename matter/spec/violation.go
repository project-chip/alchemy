package spec

import (
	"strings"

	"github.com/project-chip/alchemy/matter/types"
)

type ViolationType uint8

const (
	ViolationTypeNone ViolationType = 0

	ViolationTypeNonProvisional ViolationType = 1 << (iota - 1)
	ViolationTypeNotIfDefd

	ViolationNewParseError

	ViolationMasterList
)

func (vt ViolationType) String() string {
	var sb strings.Builder
	if (vt & ViolationTypeNonProvisional) != ViolationTypeNone {
		sb.WriteString("non-provisional")
	}
	if (vt & ViolationTypeNotIfDefd) != ViolationTypeNone {
		if sb.Len() > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString("not-if-def'd")
	}
	if (vt & ViolationNewParseError) != ViolationTypeNone {
		if sb.Len() > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString("new-parse-error")
	}
	if (vt & ViolationMasterList) != ViolationTypeNone {
		if sb.Len() > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString("master-list-incompatible")
	}
	return sb.String()
}

func (vt ViolationType) Has(o ViolationType) bool {
	return (vt & o) == o
}

type Violation struct {
	Type   ViolationType
	Entity types.Entity
	Path   string
	Line   int
	Text   string
}

func MergeViolations(vvs ...map[string][]Violation) (v map[string][]Violation) {
	v = make(map[string][]Violation)

	for _, vv := range vvs {
		for key, valueToAppend := range vv {
			v[key] = append(v[key], valueToAppend...)
		}
	}

	return
}
