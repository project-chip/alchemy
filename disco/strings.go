package disco

import (
	"regexp"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/matterfmt/ascii"
)

func findStrings(elements []interface{}, callback func(t *types.StringElement) bool) {
	for _, e := range elements {
		if ae, ok := e.(*ascii.Element); ok {
			e = ae.Base
		}
		switch el := e.(type) {
		case *types.StringElement:
			if callback(el) {
				return
			}
		case types.WithElements:
			findStrings(el.GetElements(), callback)
		case *ascii.Section:
			findStrings(el.Elements, callback)
		case *types.Table:
			findStringsInTable(el, callback)
		}
	}
}

func findStringsInTable(t *types.Table, callback func(t *types.StringElement) bool) {
	if t.Header != nil {
		for _, c := range t.Header.Cells {
			findStrings(c.Elements, callback)
		}
	}
	for _, r := range t.Rows {
		for _, c := range r.Cells {
			findStrings(c.Elements, callback)
		}

	}
	if t.Footer != nil {
		for _, c := range t.Footer.Cells {
			findStrings(c.Elements, callback)
		}
	}
}

var missingSpaceAfterPunctuationPattern = regexp.MustCompile(`([a-z])([.?!,])([A-Z])`)
var multipleSpacesPattern = regexp.MustCompile(`([\w.?!,\(\)\-":]) {2,}([\w.?!,\(\)\-":])`)
var lowercaseHexPattern = regexp.MustCompile(`(\b0x[0-9a-f]*[a-f][0-9a-f]*\b)`)
var lowercasePattern = regexp.MustCompile(`[a-f]+`)

func precleanStrings(elements []interface{}) {
	findStrings(elements, func(t *types.StringElement) bool {
		t.Content = strings.ReplaceAll(t.Content, "\t", "  ")
		return false
	})
}

func postCleanUpStrings(elements []interface{}) {
	findStrings(elements, func(t *types.StringElement) bool {
		t.Content = missingSpaceAfterPunctuationPattern.ReplaceAllString(t.Content, "$1$2 $3")
		t.Content = multipleSpacesPattern.ReplaceAllString(t.Content, "$1 $2")
		t.Content = lowercaseHexPattern.ReplaceAllStringFunc(t.Content, func(s string) string {
			return lowercasePattern.ReplaceAllStringFunc(s, func(s string) string {
				return strings.ToUpper(s)
			})
		})
		return false
	})
}
