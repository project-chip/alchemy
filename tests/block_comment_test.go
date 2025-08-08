package tests

import "github.com/project-chip/alchemy/asciidoc"

var blockComment = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.MultiLineComment{
			Delimiter: asciidoc.Delimiter{
				Type:   2,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"Comments!",
				"",
				"Getcher comments! Right here! Pipin' hot!",
			},
		},
	},
}
