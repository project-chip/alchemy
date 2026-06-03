package handlebars

import (
	"strings"

	"github.com/mailgun/raymond/v2"
)

func ToUpperHelper(s string) raymond.SafeString {
	return raymond.SafeString(strings.ToUpper(s))
}

func ToLowerHelper(s string) raymond.SafeString {
	return raymond.SafeString(strings.ToLower(s))
}
