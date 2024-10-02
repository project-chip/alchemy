package dm

import "strings"

func scrubDescription(description string) string {
	return strings.Join(strings.Fields(description), " ")
}
