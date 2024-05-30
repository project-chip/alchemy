package tests

import "github.com/hasty/alchemy/asciidoc"

var blockComment = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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
