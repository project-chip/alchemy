package tests

import (
	"testing"

	"github.com/project-chip/alchemy/asciidoc"
)

func TestText(t *testing.T) {
	textTests.run(t)
}

var textTests = parseTests{

	{"line breaks", "asciidoctor/text_test_line_breaks.adoc", textTestLineBreaks},
}

var textTestLineBreaks = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "This line is separated by a horizontal rule...",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.ThematicBreak{
			AttributeList: nil,
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "...from this line.",
		},
		&asciidoc.NewLine{},
	},
}
