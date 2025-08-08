package tests

import "github.com/project-chip/alchemy/asciidoc"

var inlineImage = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Paragraph{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.TitleAttribute{
					Val: asciidoc.Elements{
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
			Elements: asciidoc.Elements{
				&asciidoc.InlineImage{
					AttributeList: asciidoc.AttributeList{
						&asciidoc.PositionalAttribute{
							Offset:      0,
							ImpliedName: "alt",
							Val: asciidoc.Elements{
								&asciidoc.String{
									Value: "Priority Scheme of Pump Operation and Control.jpg",
								},
							},
						},
						&asciidoc.NamedAttribute{
							Name: "height",
							Val: asciidoc.Elements{
								&asciidoc.String{
									Value: "524",
								},
							},
						},
						&asciidoc.NamedAttribute{
							Name: "width",
							Val: asciidoc.Elements{
								&asciidoc.String{
									Value: "576",
								},
							},
						},
					},
					ImagePath: asciidoc.Elements{
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
