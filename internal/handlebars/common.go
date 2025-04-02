package handlebars

import "github.com/mailgun/raymond/v2"

func RegisterCommonHelpers(t *raymond.Template) {
	t.RegisterHelper("raw", RawHelper)
	t.RegisterHelper("ifSet", IfSetHelper)
	t.RegisterHelper("ifEqual", IfEqualHelper)
	t.RegisterHelper("quote", QuoteHelper)
	t.RegisterHelper("add", AddHelper)
	t.RegisterHelper("brace", BraceHelper)
	t.RegisterHelper("format", FormatHelper)
	t.RegisterHelper("year", YearHelper)
}
