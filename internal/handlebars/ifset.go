package handlebars

import (
	"github.com/mailgun/raymond/v2"
)

func IfSetHelper(value any, options *raymond.Options) string {
	switch value.(type) {
	case nil:
		return options.Inverse()
	default:
		return options.Fn()
	}
}

func IfEqualHelper(a, b any, options *raymond.Options) string {
	if raymond.Str(a) == raymond.Str(b) {
		return options.Fn()
	}
	return options.Inverse()
}
