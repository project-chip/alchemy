package testplan

import "github.com/project-chip/alchemy/matter/conformance"

type Feature struct {
	From        uint64
	To          uint64
	Bits        []uint64
	Code        string
	Summary     string
	Conformance conformance.Set
}
