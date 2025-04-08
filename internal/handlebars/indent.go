package handlebars

import (
	"strings"

	"github.com/mailgun/raymond/v2"
)

func IndentLengthHelper(s string) raymond.SafeString {
	return raymond.SafeString(strings.Repeat(" ", len(s)))
}
