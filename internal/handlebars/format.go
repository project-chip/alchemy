package handlebars

import (
	"fmt"

	"github.com/mailgun/raymond/v2"
)

func FormatHelper(value any, format string) raymond.SafeString {
	return raymond.SafeString(fmt.Sprintf(format, value))
}
