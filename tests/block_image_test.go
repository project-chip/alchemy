package tests

import "github.com/hasty/alchemy/asciidoc"

var blockImage = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.BlockImage{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.PositionalAttribute{
					Offset:      0,
					ImpliedName: "alt",
					Val: asciidoc.Set{
						&asciidoc.String{
							Value: "image",
						},
					},
				},
				&asciidoc.NamedAttribute{
					Name: "height",
					Val: asciidoc.Set{
						&asciidoc.String{
							Value: "85",
						},
					},
				},
				&asciidoc.NamedAttribute{
					Name: "width",
					Val: asciidoc.Set{
						&asciidoc.String{
							Value: "167",
						},
					},
				},
			},
			Path: asciidoc.Set{
				&asciidoc.String{
					Value: "./images/lighting/media/image3.png",
				},
			},
		},
		asciidoc.EmptyLine{
			Text: "",
		},
	},
}
