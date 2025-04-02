package handlebars

import (
	"strconv"
	"time"

	"github.com/mailgun/raymond/v2"
)

func YearHelper() raymond.SafeString {
	return raymond.SafeString(strconv.Itoa(time.Now().Year()))
}
