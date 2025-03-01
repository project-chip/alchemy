package text

import "strings"

func HasCaseInsensitiveSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) &&
		strings.EqualFold(s[len(s)-len(suffix):], suffix)
}

func HasCaseInsensitivePrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && strings.EqualFold(s[0:len(prefix)], prefix)
}

func TrimCaseInsensitiveSuffix(s, suffix string) string {
	if HasCaseInsensitiveSuffix(s, suffix) {
		return s[:len(s)-len(suffix)]
	}
	return s
}

func TrimCaseInsensitivePrefix(s, prefix string) string {
	if HasCaseInsensitivePrefix(s, prefix) {
		return s[len(prefix):]
	}
	return s
}

func TrimUnnecessaryParens(as string) string {
	if len(as) > 2 {
		if as[0] == '(' && as[len(as)-1] == ')' {
			as = as[1 : len(as)-1]
		}
	}
	return as
}
