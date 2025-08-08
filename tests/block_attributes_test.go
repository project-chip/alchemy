package tests

import "github.com/project-chip/alchemy/asciidoc"

var blockAttributes = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.Paragraph{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.NamedAttribute{
					Name: "role",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "lead",
						},
					},
					Quote: 2,
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "This is a demonstration of ",
				},
				&asciidoc.UserAttributeReference{
					Value: "library",
				},
				&asciidoc.String{
					Value: ". And this is the preamble of this document.",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 0,
		},
	},
}
