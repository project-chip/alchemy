package handlebars

import (
	"strconv"

	"github.com/mailgun/raymond/v2"
)

func QuoteHelper(s string) raymond.SafeString {
	return raymond.SafeString(strconv.Quote(s))
}
