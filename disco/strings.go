package disco

import (
	"regexp"
	"strings"

	"github.com/hasty/adoc/elements"
	"github.com/hasty/alchemy/internal/parse"
)

var missingSpaceAfterPunctuationPattern = regexp.MustCompile(`([a-z])([.?!,])([A-Z])`)
var multipleSpacesPattern = regexp.MustCompile(`([\w.?!,\(\)\-":]) {2,}([\w.?!,\(\)\-":])`)
var lowercaseHexPattern = regexp.MustCompile(`(\b0x[0-9a-f]*[a-f][0-9a-f]*\b)`)
var lowercasePattern = regexp.MustCompile(`[a-f]+`)

func precleanStrings(els elements.Set) {
	parse.Search(els, func(t *elements.String) parse.SearchShould {
		t.Value = strings.ReplaceAll(t.Value, "\t", "  ")
		return parse.SearchShouldContinue
	})
}

func (b *Ball) postCleanUpStrings(els elements.Set) {
	parse.Search(els, func(t *elements.String) parse.SearchShould {
		if b.options.addSpaceAfterPunctuation {
			t.Value = missingSpaceAfterPunctuationPattern.ReplaceAllString(t.Value, "$1$2 $3")
		}
		if b.options.removeExtraSpaces {
			t.Value = multipleSpacesPattern.ReplaceAllString(t.Value, "$1 $2")
		}
		if b.options.uppercaseHex {
			t.Value = lowercaseHexPattern.ReplaceAllStringFunc(t.Value, func(s string) string {
				return lowercasePattern.ReplaceAllStringFunc(s, func(s string) string {
					return strings.ToUpper(s)
				})
			})
		}
		return parse.SearchShouldContinue
	})
}
