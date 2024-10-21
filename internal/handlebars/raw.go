package handlebars

import "github.com/mailgun/raymond/v2"

func RawHelper(value string) raymond.SafeString {
	return raymond.SafeString(value)
}
