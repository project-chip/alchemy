package validate

import (
	"strings"

	"github.com/project-chip/alchemy/matter/spec"
)

func Validate(spec *spec.Specification) {
	validateStructs(spec)
	validateDeviceTypes(spec)
}

func stripName(s string) string {
	var result strings.Builder
	for i := 0; i < len(s); i++ {
		b := s[i]
		if ('a' <= b && b <= 'z') ||
			('A' <= b && b <= 'Z') {
			result.WriteByte(b)
		}
	}
	return result.String()
}
