package matter

import (
	"strings"
	"sync"
	"unicode"
	"unicode/utf8"
)

var acronyms sync.Map

func AddCaseAcronym(s string) {
	acronyms.Store(s, struct{}{})
}

func isSeparator(r rune) bool {
	return r == '.' || r == ' ' || r == '-' || r == '_' || r == '(' || r == ')'
}

// Case turns a string with spaces into a valid Matter identifier in Pascal Case
// It is Unicode-aware, and preserves acronyms
// e.g. "Level control" becomes "LevelControl", "TV set" becomes "TVSet"
func Case(s string) string {
	return caseify(s, 0)
}

// Case turns a string with spaces into a valid Matter identifier, with a custom separator rune
// e.g. "Level control" becomes "Level-Control", "TV set" becomes "TV-Set"
func CaseWithSeparator(s string, separator rune) string {
	return caseify(s, separator)
}

func caseify(s string, separator rune) string {
	runes := []rune(s)
	b := make([]byte, 0, len(runes))
	var index int
	upperCaseNextRune := true // We'll always start with a capital
	for index < len(runes) {
		r := runes[index]
		if unicode.IsUpper(r) {
			if upperCaseNextRune && separator != 0 && len(b) > 0 { // If there's a supplied separator, append it, unless the string is empty
				b = utf8.AppendRune(b, separator)
			}
			var end int
			var endedBySeparator bool
			for end = index + 1; end < len(runes); end++ { // Look for a run of upper-case letters
				if unicode.IsUpper(runes[end]) {
					if end == len(runes)-1 { // Last rune
						endedBySeparator = true
						break
					}
					continue
				}
				if isSeparator(runes[end]) {
					endedBySeparator = true
				}
				break
			}
			if end-index > 1 {

				_, isAcronym := acronyms.Load(string(runes[index:end]))
				if isAcronym || endedBySeparator {
					// If this run of upper-case letters is a known acronym, or
					// it ends the string, preserve it
					for index < end {
						b = utf8.AppendRune(b, runes[index])
						index++
					}
				} else {
					// Otherwise, lower-case the remainder of the block of upper-case letters
					b = utf8.AppendRune(b, runes[index])
					index++
					for index < end-1 {
						b = utf8.AppendRune(b, unicode.ToLower(runes[index]))
						index++
					}
					b = utf8.AppendRune(b, runes[index])
				}
				index = end
			} else {
				// It's just the one upper-case letter
				b = utf8.AppendRune(b, runes[index])
				index++
			}
			upperCaseNextRune = false
			continue
		} else if unicode.IsLower(r) {
			if upperCaseNextRune {
				if separator != 0 {
					b = utf8.AppendRune(b, separator)
				}
				b = utf8.AppendRune(b, unicode.ToUpper(r))
				upperCaseNextRune = false
			} else {
				b = utf8.AppendRune(b, r)
			}
		} else if unicode.IsNumber(r) {
			b = utf8.AppendRune(b, r)
			upperCaseNextRune = true
		} else {
			upperCaseNextRune = isSeparator(r)
		}
		index++
	}
	return string(b)
}

type charClass uint8

const (
	charClassUnknown charClass = iota
	charClassUpper
	charClassLower
	charClassDigit
)

func Uncase(s string) string {
	runes := []rune(s)
	var sb strings.Builder
	sb.Grow(len(runes))
	var last charClass
	for _, r := range runes {
		var current charClass
		if unicode.IsUpper(r) {
			current = charClassUpper
		} else if unicode.IsLower(r) {
			current = charClassLower

		} else if unicode.IsDigit(r) {
			current = charClassDigit
		}
		if last != current {
			switch last {
			case charClassUnknown:
			case charClassUpper:
				if current != charClassLower && current != charClassDigit {
					sb.WriteRune(' ')
				}
			case charClassLower:
				if current != charClassUnknown {
					sb.WriteRune(' ')
				}
			default:
				sb.WriteRune(' ')
			}
		}
		switch current {
		case charClassUpper:

		}
		sb.WriteRune(r)
		last = current
	}
	return sb.String()
}
