package disco

import (
	"regexp"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/matterfmt/ascii"
)

var missingSpaceAfterPunctuationPattern = regexp.MustCompile(`([a-z])([.?!,])([A-Z])`)
var multipleSpacesPattern = regexp.MustCompile(`([\w.?!,\(\)\-":]) {2,}([\w.?!,\(\)\-":])`)
var lowercaseHexPattern = regexp.MustCompile(`(\b0x[0-9a-f]*[a-f][0-9a-f]*\b)`)
var lowercasePattern = regexp.MustCompile(`[a-f]+`)

func precleanStrings(elements []interface{}) {
	ascii.Search(elements, func(t *types.StringElement) bool {
		t.Content = strings.ReplaceAll(t.Content, "\t", "  ")
		return false
	})
}

func (b *Ball) postCleanUpStrings(elements []interface{}) {
	ascii.Search(elements, func(t *types.StringElement) bool {
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
