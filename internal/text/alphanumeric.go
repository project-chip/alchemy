package text

import (
	"fmt"
	"strconv"
	"unicode"
)

func IsAlphanumeric(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			continue
		}
		return false
	}
	return true
}

func IsUpperCase(s string) bool {
	for _, r := range s {
		if unicode.IsUpper(r) {
			continue
		}
		return false
	}
	return true
}

func ParseRelativeNumber(s string) (number int, relative bool, err error) {
	if len(s) > 0 {
		switch s[0] {
		case '+':
			relative = true
		case '-':
			relative = true
		}
	}
	if len(s) == 0 {
		err = fmt.Errorf("invalid relative number: %s", s)
		return
	}
	number, err = strconv.Atoi(s)
	if err != nil {
		err = fmt.Errorf("invalid relative number: %s (%v)", s, err)
	}
	return
}

func IsRepeatedCharacter(s string, r rune) bool {
	if len(s) == 0 {
		return false
	}
	for _, c := range s {
		if c != r {
			return false
		}
	}
	return true
}
