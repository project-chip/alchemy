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

func Case(s string) string {
	return caseify(s, 0)
}

func CaseWithSeparator(s string, separator rune) string {
	return caseify(s, separator)
}

func caseify(s string, separator rune) string {
	runes := []rune(s)
	b := make([]byte, 0, len(runes))
	var index int
	nextUpper := true
	for index < len(runes) {
		r := runes[index]
		if unicode.IsUpper(r) {
			if nextUpper && separator != 0 && len(b) > 0 {
				b = utf8.AppendRune(b, separator)
			}
			var end int
			var endedBySeparator bool
			for end = index + 1; end < len(runes); end++ {
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
					for index < end {
						b = utf8.AppendRune(b, runes[index])
						index++
					}
				} else {
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
				b = utf8.AppendRune(b, runes[index])
				index++
			}
			nextUpper = false
			continue
		} else if unicode.IsLower(r) {
			if nextUpper {
				if separator != 0 {
					b = utf8.AppendRune(b, separator)
				}
				b = utf8.AppendRune(b, unicode.ToUpper(r))
				nextUpper = false
			} else {
				b = utf8.AppendRune(b, r)
			}
		} else if unicode.IsNumber(r) {
			b = utf8.AppendRune(b, r)
			nextUpper = true
		} else {
			nextUpper = isSeparator(r)
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
