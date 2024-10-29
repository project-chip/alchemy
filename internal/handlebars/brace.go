package handlebars

import "github.com/mailgun/raymond/v2"

func BraceHelper(options *raymond.Options) raymond.SafeString {
	return raymond.SafeString("{" + options.Fn() + "}")
}
