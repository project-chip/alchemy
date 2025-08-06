package tests

import "github.com/project-chip/alchemy/asciidoc"

var inlineImage = &asciidoc.Document{
	Set: asciidoc.Set{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Paragraph{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.TitleAttribute{
					Val: asciidoc.Set{
						&asciidoc.String{
							Value: "Priority Scheme of Pump Operation and Control",
						},
					},
				},
				&asciidoc.AnchorAttribute{
					ID: &asciidoc.String{
						Value: "ref_PumpOperationAndControlFigure",
					},
					Label: nil,
				},
			},
			Set: asciidoc.Set{
				&asciidoc.InlineImage{
					AttributeList: asciidoc.AttributeList{
						&asciidoc.PositionalAttribute{
							Offset:      0,
							ImpliedName: "alt",
							Val: asciidoc.Set{
								&asciidoc.String{
									Value: "Priority Scheme of Pump Operation and Control.jpg",
								},
							},
						},
						&asciidoc.NamedAttribute{
							Name: "height",
							Val: asciidoc.Set{
								&asciidoc.String{
									Value: "524",
								},
							},
						},
						&asciidoc.NamedAttribute{
							Name: "width",
							Val: asciidoc.Set{
								&asciidoc.String{
									Value: "576",
								},
							},
						},
					},
					ImagePath: asciidoc.Set{
						&asciidoc.String{
							Value: "./images/hvac/media/image4.jpeg",
						},
					},
				},
				&asciidoc.NewLine{},
			},
			Admonition: 0,
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
	},
}
