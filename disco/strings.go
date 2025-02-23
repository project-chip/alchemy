package disco

import (
	"regexp"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/internal/parse"
	"github.com/project-chip/alchemy/matter/spec"
)

var missingSpaceAfterPunctuationPattern = regexp.MustCompile(`([a-z])([.?!,])([A-Z])`)
var multipleSpacesPattern = regexp.MustCompile(`([\w.?!,\(\)\-":]) {2,}([\w.?!,\(\)\-":])`)
var lowercaseHexPattern = regexp.MustCompile(`(\b0x[0-9a-f]*[a-f][0-9a-f]*\b)`)
var lowercasePattern = regexp.MustCompile(`[a-f]+`)

func precleanStrings(els asciidoc.Set) {
	parse.Search(els, func(t *asciidoc.String) parse.SearchShould {
		t.Value = strings.ReplaceAll(t.Value, "\t", "  ")
		return parse.SearchShouldContinue
	})
}

func (b *Baller) postCleanUpStrings(els asciidoc.Set) {
	parse.Search(els, func(el any) parse.SearchShould {
		se, ok := el.(*spec.Element)
		if ok {
			el = se.Base
		}
		switch el := el.(type) {
		case *asciidoc.Monospace, *asciidoc.DoubleMonospace:
			return parse.SearchShouldSkip
		case *asciidoc.Table:
			return parse.SearchShouldSkip
		case *asciidoc.String:
			if b.options.addSpaceAfterPunctuation {
				el.Value = missingSpaceAfterPunctuationPattern.ReplaceAllString(el.Value, "$1$2 $3")
			}
		}
		return parse.SearchShouldContinue
	})
	parse.Search(els, func(t *asciidoc.String) parse.SearchShould {
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
