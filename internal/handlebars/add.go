package handlebars

import (
	"strconv"

	"github.com/mailgun/raymond/v2"
)

func AddHelper(augend int, addend int) raymond.SafeString {
	return raymond.SafeString(strconv.Itoa(augend + addend))
}
