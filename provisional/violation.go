package provisional

import (
	"strings"

	"github.com/project-chip/alchemy/matter/types"
)

type ViolationType uint8

const (
	ViolationTypeNone ViolationType = 0

	ViolationTypeNonProvisional ViolationType = 1 << (iota - 1)
	ViolationTypeNotIfDefd
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
}

type entityViolations map[types.Entity]ViolationType

func (ev entityViolations) add(entity types.Entity, violationType ViolationType) {
	if violationType != ViolationTypeNone {
		ev[entity] = violationType
	}
}
