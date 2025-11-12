package constraint

import "strings"

func isReservedWord(s string) bool {
	switch strings.ToLower(s) {
	case "to", "max", "min", "any", "all", "true", "false", "desc":
		return true
	}
	return false
}
