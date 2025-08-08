package tests

import (
	"testing"

	"github.com/project-chip/alchemy/asciidoc"
)

func TestLogger(t *testing.T) {
	loggerTests.run(t)
}

var loggerTests = parseTests{

	{"provides access to logger via static logger method", "asciidoctor/logger_test_provides_access_to_logger_via_static_logger_method.adoc", loggerTestProvidesAccessToLoggerViaStaticLoggerMethod, nil},
}

var loggerTestProvidesAccessToLoggerViaStaticLoggerMethod = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Paragraph{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: nil,
					ID: &asciidoc.ShorthandID{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "first",
							},
						},
					},
					Roles:   nil,
					Options: nil,
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "first paragraph",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 0,
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Paragraph{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: nil,
					ID: &asciidoc.ShorthandID{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "first",
							},
						},
					},
					Roles:   nil,
					Options: nil,
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "another first paragraph",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 0,
		},
	},
}
