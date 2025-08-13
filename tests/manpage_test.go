package tests

import (
	"testing"

	"github.com/project-chip/alchemy/asciidoc"
)

func TestManpage(t *testing.T) {
	manpageTests.run(t)
}

var manpageTests = parseTests{

	{"should set proper manpage-related attributes", "asciidoctor/manpage_test_should_set_proper_manpage_related_attributes.adoc", manpageTestShouldSetProperManpageRelatedAttributes},

	{"should substitute attributes in manname and manpurpose in NAME section", "asciidoctor/manpage_test_should_substitute_attributes_in_manname_and_manpurpose_in_name_section.adoc", manpageTestShouldSubstituteAttributesInMannameAndManpurposeInNameSection},

	{"should not parse NAME section if manname and manpurpose attributes are set", "asciidoctor/manpage_test_should_not_parse_name_section_if_manname_and_manpurpose_attributes_are_set.adoc", manpageTestShouldNotParseNameSectionIfMannameAndManpurposeAttributesAreSet},

	{"should normalize whitespace and skip line comments before and inside NAME section", "asciidoctor/manpage_test_should_normalize_whitespace_and_skip_line_comments_before_and_inside_name_section.adoc", manpageTestShouldNormalizeWhitespaceAndSkipLineCommentsBeforeAndInsideNameSection},

	{"should parse malformed document with warnings", "asciidoctor/manpage_test_should_parse_malformed_document_with_warnings.adoc", manpageTestShouldParseMalformedDocumentWithWarnings},

	{"should warn if first section is not name section", "asciidoctor/manpage_test_should_warn_if_first_section_is_not_name_section.adoc", manpageTestShouldWarnIfFirstSectionIsNotNameSection},

	{"should preserve hard line breaks in verse block", "asciidoctor/manpage_test_should_preserve_hard_line_breaks_in_verse_block.adoc", manpageTestShouldPreserveHardLineBreaksInVerseBlock},
}

var manpageTestShouldSetProperManpageRelatedAttributes = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Author Name",
				},
				&asciidoc.NewLine{},
				&asciidoc.AttributeEntry{
					Name: "doctype",
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "manpage",
						},
					},
				},
				&asciidoc.String{
					Value: ":man manual: Foo Bar Manual",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: ":man source: Foo Bar 1.0",
				},
				&asciidoc.NewLine{},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.Section{
					AttributeList: nil,
					Elements: asciidoc.Elements{
						&asciidoc.EmptyLine{
							Text: "",
						},
						&asciidoc.String{
							Value: "foo-bar - puts the foo in your bar",
						},
						&asciidoc.NewLine{},
					},
					Title: asciidoc.Elements{
						&asciidoc.String{
							Value: "NAME",
						},
					},
					Level: 1,
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "foo\\--<bar> (1)",
				},
			},
			Level: 0,
		},
	},
}

var manpageTestShouldSubstituteAttributesInMannameAndManpurposeInNameSection = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Author Name",
				},
				&asciidoc.NewLine{},
				&asciidoc.AttributeEntry{
					Name: "doctype",
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "manpage",
						},
					},
				},
				&asciidoc.String{
					Value: ":man manual: Foo Bar Manual",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: ":man source: Foo Bar 1.0",
				},
				&asciidoc.NewLine{},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.Section{
					AttributeList: nil,
					Elements: asciidoc.Elements{
						&asciidoc.EmptyLine{
							Text: "",
						},
						&asciidoc.UserAttributeReference{
							Value: "cmdname",
						},
						&asciidoc.String{
							Value: " - ",
						},
						&asciidoc.UserAttributeReference{
							Value: "cmdname",
						},
						&asciidoc.String{
							Value: " puts the foo in your bar",
						},
						&asciidoc.NewLine{},
					},
					Title: asciidoc.Elements{
						&asciidoc.String{
							Value: "NAME",
						},
					},
					Level: 1,
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "{cmdname} (1)",
				},
			},
			Level: 0,
		},
	},
}

var manpageTestShouldNotParseNameSectionIfMannameAndManpurposeAttributesAreSet = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Author Name",
				},
				&asciidoc.NewLine{},
				&asciidoc.AttributeEntry{
					Name: "doctype",
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "manpage",
						},
					},
				},
				&asciidoc.String{
					Value: ":man manual: Foo Bar Manual",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: ":man source: Foo Bar 1.0",
				},
				&asciidoc.NewLine{},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.Section{
					AttributeList: nil,
					Elements: asciidoc.Elements{
						&asciidoc.EmptyLine{
							Text: "",
						},
						&asciidoc.Bold{
							AttributeList: nil,
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "foobar",
								},
							},
						},
						&asciidoc.String{
							Value: " [",
						},
						&asciidoc.Italic{
							AttributeList: nil,
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "OPTIONS",
								},
							},
						},
						&asciidoc.String{
							Value: "]...",
						},
						&asciidoc.NewLine{},
						&asciidoc.EmptyLine{
							Text: "",
						},
					},
					Title: asciidoc.Elements{
						&asciidoc.String{
							Value: "SYNOPSIS",
						},
					},
					Level: 1,
				},
				&asciidoc.Section{
					AttributeList: nil,
					Elements: asciidoc.Elements{
						&asciidoc.EmptyLine{
							Text: "",
						},
						&asciidoc.String{
							Value: "When you need to put some foo on the bar.",
						},
						&asciidoc.NewLine{},
					},
					Title: asciidoc.Elements{
						&asciidoc.String{
							Value: "DESCRIPTION",
						},
					},
					Level: 1,
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "foobar (1)",
				},
			},
			Level: 0,
		},
	},
}

var manpageTestShouldNormalizeWhitespaceAndSkipLineCommentsBeforeAndInsideNameSection = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Author Name",
				},
				&asciidoc.NewLine{},
				&asciidoc.AttributeEntry{
					Name: "doctype",
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "manpage",
						},
					},
				},
				&asciidoc.String{
					Value: ":man manual: Foo Bar Manual",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: ":man source: Foo Bar 1.0",
				},
				&asciidoc.NewLine{},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.SingleLineComment{
					Value: " this is the name section",
				},
				&asciidoc.Section{
					AttributeList: nil,
					Elements: asciidoc.Elements{
						&asciidoc.EmptyLine{
							Text: "",
						},
						&asciidoc.SingleLineComment{
							Value: " it follows the form `name - description`",
						},
						&asciidoc.String{
							Value: "foobar - puts some foo",
						},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: " on the bar",
						},
						&asciidoc.NewLine{},
						&asciidoc.SingleLineComment{
							Value: " a little bit of this, a little bit of that",
						},
						&asciidoc.EmptyLine{
							Text: "",
						},
					},
					Title: asciidoc.Elements{
						&asciidoc.String{
							Value: "NAME",
						},
					},
					Level: 1,
				},
				&asciidoc.Section{
					AttributeList: nil,
					Elements: asciidoc.Elements{
						&asciidoc.EmptyLine{
							Text: "",
						},
						&asciidoc.Bold{
							AttributeList: nil,
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "foobar",
								},
							},
						},
						&asciidoc.String{
							Value: " [",
						},
						&asciidoc.Italic{
							AttributeList: nil,
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "OPTIONS",
								},
							},
						},
						&asciidoc.String{
							Value: "]...",
						},
						&asciidoc.NewLine{},
						&asciidoc.EmptyLine{
							Text: "",
						},
					},
					Title: asciidoc.Elements{
						&asciidoc.String{
							Value: "SYNOPSIS",
						},
					},
					Level: 1,
				},
				&asciidoc.Section{
					AttributeList: nil,
					Elements: asciidoc.Elements{
						&asciidoc.EmptyLine{
							Text: "",
						},
						&asciidoc.String{
							Value: "When you need to put some foo on the bar.",
						},
						&asciidoc.NewLine{},
					},
					Title: asciidoc.Elements{
						&asciidoc.String{
							Value: "DESCRIPTION",
						},
					},
					Level: 1,
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "foobar (1)",
				},
			},
			Level: 0,
		},
	},
}

var manpageTestShouldParseMalformedDocumentWithWarnings = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.Section{
					AttributeList: nil,
					Elements: asciidoc.Elements{
						&asciidoc.EmptyLine{
							Text: "",
						},
						&asciidoc.String{
							Value: "command - does stuff",
						},
						&asciidoc.NewLine{},
					},
					Title: asciidoc.Elements{
						&asciidoc.String{
							Value: "Name",
						},
					},
					Level: 1,
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "command",
				},
			},
			Level: 0,
		},
	},
}

var manpageTestShouldWarnIfFirstSectionIsNotNameSection = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.Section{
					AttributeList: nil,
					Elements: asciidoc.Elements{
						&asciidoc.EmptyLine{
							Text: "",
						},
						&asciidoc.String{
							Value: "Does stuff.",
						},
						&asciidoc.NewLine{},
					},
					Title: asciidoc.Elements{
						&asciidoc.String{
							Value: "Synopsis",
						},
					},
					Level: 1,
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "command(1)",
				},
			},
			Level: 0,
		},
	},
}

var manpageTestShouldPreserveHardLineBreaksInVerseBlock = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.Paragraph{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.TitleAttribute{
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "lines",
						},
					},
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "      [verse]",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "      ",
				},
				&asciidoc.Italic{
					AttributeList: nil,
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "command",
						},
					},
				},
				&asciidoc.String{
					Value: " [",
				},
				&asciidoc.Italic{
					AttributeList: nil,
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "OPTION",
						},
					},
				},
				&asciidoc.String{
					Value: "]... ",
				},
				&asciidoc.Italic{
					AttributeList: nil,
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "FILE",
						},
					},
				},
				&asciidoc.String{
					Value: "...",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 0,
		},
	},
}
