package tests

import "github.com/project-chip/alchemy/asciidoc"

var inlineAdmonition = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.Paragraph{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "If you want the output to look familiar, copy (or link) the AsciiDoc stylesheet, asciiasciidoc.css, to the output directory.",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 2,
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Paragraph{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Items marked with TODO are either not yet supported or work in progress.",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 1,
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "If the lock already has an Aliro Reader configuration defined,",
				},
			},
			AttributeList: nil,
			Marker:        ".",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "If the lock already has an Aliro Reader configuration defined,",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "  (i.e. the AliroReaderVerificationKey attribute is not null),",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "  the response SHALL be INVALID_IN_STATE.",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: " ",
				},
				&asciidoc.LineBreak{},
				&asciidoc.NewLine{},
				&asciidoc.Paragraph{
					AttributeList: nil,
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "This avoids accidentally overwriting values that were just set by a different administrator.",
						},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "This avoids accidentally overwriting values that were just set by a different administrator.",
						},
					},
					Admonition: 1,
				},
			},
			AttributeList: nil,
			Marker:        ".",
		},
	},
}
