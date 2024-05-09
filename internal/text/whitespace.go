package text

import "strings"

func TrimWhitespace(s string) string {
	return strings.Trim(s, " \t")
}
