package disco

import (
	"regexp"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/internal/parse"
)

var missingSpaceAfterPunctuationPattern = regexp.MustCompile(`([a-z])([.?!,])([A-Z])`)
var multipleSpacesPattern = regexp.MustCompile(`([\w.?!,\(\)\-":]) {2,}([\w.?!,\(\)\-":])`)
var lowercaseHexPattern = regexp.MustCompile(`(\b0x[0-9a-f]*[a-f][0-9a-f]*\b)`)
var lowercasePattern = regexp.MustCompile(`[a-f]+`)

func precleanStrings(doc asciidoc.ParentElement) {
	parse.Traverse(doc, doc.Children(), func(t *asciidoc.String, parent parse.HasElements, index int) parse.SearchShould {
		t.Value = strings.ReplaceAll(t.Value, "\t", "  ")
		return parse.SearchShouldContinue
	})
}

func (b *Baller) postCleanUpStrings(root asciidoc.ParentElement) {
	parse.Traverse(root, root.Children(), func(el any, parent parse.HasElements, index int) parse.SearchShould {
		switch el := el.(type) {
		case *asciidoc.Monospace, *asciidoc.DoubleMonospace:
			return parse.SearchShouldSkip
		case *asciidoc.Table:
			return parse.SearchShouldSkip
		case *asciidoc.String:
			if b.options.AddSpaceAfterPunctuation {
				el.Value = missingSpaceAfterPunctuationPattern.ReplaceAllString(el.Value, "$1$2 $3")
			}
		}
		return parse.SearchShouldContinue
	})
	parse.Traverse(root, root.Children(), func(t *asciidoc.String, parent parse.HasElements, index int) parse.SearchShould {
		if b.options.RemoveExtraSpaces {
			t.Value = multipleSpacesPattern.ReplaceAllString(t.Value, "$1 $2")
		}
		if b.options.UppercaseHex {
			t.Value = lowercaseHexPattern.ReplaceAllStringFunc(t.Value, func(s string) string {
				return lowercasePattern.ReplaceAllStringFunc(s, func(s string) string {
					return strings.ToUpper(s)
				})
			})
		}
		return parse.SearchShouldContinue
	})
}
