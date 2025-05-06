package spec

import "strings"

func Validate(spec *Specification) {
	validateAttributes(spec)
	validateCommands(spec)
	validateEvents(spec)
	validateStructs(spec)
	validateDeviceTypes(spec)
}

func stripNonAlphabeticalCharacters(s string) string {
	var result strings.Builder
	for i := range len(s) {
		b := s[i]
		if ('a' <= b && b <= 'z') ||
			('A' <= b && b <= 'Z') {
			result.WriteByte(b)
		}
	}
	return result.String()
}
