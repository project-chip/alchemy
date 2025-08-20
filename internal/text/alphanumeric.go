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
