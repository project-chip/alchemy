package tests

import (
	"testing"

	"github.com/project-chip/alchemy/asciidoc"
)

func TestParser(t *testing.T) {
	parserTests.run(t)
}

var parserTests = parseTests{

	{"parse name with more than 3 parts in author attribute", "asciidoctor/parser_test_parse_name_with_more_than_3_parts_in_author_attribute.adoc", parserTestParseNameWithMoreThan3PartsInAuthorAttribute},

	{"use explicit authorinitials if set after author attribute", "asciidoctor/parser_test_use_explicit_authorinitials_if_set_after_author_attribute.adoc", parserTestUseExplicitAuthorinitialsIfSetAfterAuthorAttribute},

	{"use implicit authors if value of authors attribute matches computed value", "asciidoctor/parser_test_use_implicit_authors_if_value_of_authors_attribute_matches_computed_value.adoc", parserTestUseImplicitAuthorsIfValueOfAuthorsAttributeMatchesComputedValue},

	{"replace implicit authors if value of authors attribute does not match computed value", "asciidoctor/parser_test_replace_implicit_authors_if_value_of_authors_attribute_does_not_match_computed_value.adoc", parserTestReplaceImplicitAuthorsIfValueOfAuthorsAttributeDoesNotMatchComputedValue},

	{"sets authorcount to 0 if document has no authors", "asciidoctor/parser_test_sets_authorcount_to_0_if_document_has_no_authors.adoc", parserTestSetsAuthorcountTo0IfDocumentHasNoAuthors},

	{"parse rev number date remark", "asciidoctor/parser_test_parse_rev_number_date_remark.adoc", parserTestParseRevNumberDateRemark},

	{"parse rev number, data, and remark as attribute references", "asciidoctor/parser_test_parse_rev_number_data_and_remark_as_attribute_references.adoc", parserTestParseRevNumberDataAndRemarkAsAttributeReferences},

	{"should resolve attribute references in rev number, data, and remark", "asciidoctor/parser_test_should_resolve_attribute_references_in_rev_number_data_and_remark.adoc", parserTestShouldResolveAttributeReferencesInRevNumberDataAndRemark},

	{"parse rev date", "asciidoctor/parser_test_parse_rev_date.adoc", parserTestParseRevDate},

	{"parse rev number with trailing comma", "asciidoctor/parser_test_parse_rev_number_with_trailing_comma.adoc", parserTestParseRevNumberWithTrailingComma},

	{"parse rev number", "asciidoctor/parser_test_parse_rev_number.adoc", parserTestParseRevNumber},

	{"treats arbitrary text on rev line as revdate", "asciidoctor/parser_test_treats_arbitrary_text_on_rev_line_as_revdate.adoc", parserTestTreatsArbitraryTextOnRevLineAsRevdate},

	{"parse rev date remark", "asciidoctor/parser_test_parse_rev_date_remark.adoc", parserTestParseRevDateRemark},

	{"should not mistake attribute entry as rev remark", "asciidoctor/parser_test_should_not_mistake_attribute_entry_as_rev_remark.adoc", parserTestShouldNotMistakeAttributeEntryAsRevRemark},

	{"parse rev remark only", "asciidoctor/parser_test_parse_rev_remark_only.adoc", parserTestParseRevRemarkOnly},

	{"skip line comments before author", "asciidoctor/parser_test_skip_line_comments_before_author.adoc", parserTestSkipLineCommentsBeforeAuthor},

	{"skip block comment before author", "asciidoctor/parser_test_skip_block_comment_before_author.adoc", parserTestSkipBlockCommentBeforeAuthor},

	{"skip block comment before rev", "asciidoctor/parser_test_skip_block_comment_before_rev.adoc", parserTestSkipBlockCommentBeforeRev},

	{"break header at line with three forward slashes", "asciidoctor/parser_test_break_header_at_line_with_three_forward_slashes.adoc", parserTestBreakHeaderAtLineWithThreeForwardSlashes},

	{"expands tabs to spaces", "asciidoctor/parser_test_expands_tabs_to_spaces.adoc", parserTestExpandsTabsToSpaces},

	{"adjust indentation handles empty lines gracefully", "asciidoctor/parser_test_adjust_indentation_handles_empty_lines_gracefully.adoc", parserTestAdjustIndentationHandlesEmptyLinesGracefully},
}

var parserTestParseNameWithMoreThan3PartsInAuthorAttribute = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "Jean-Claude Van Damme",
		},
		&asciidoc.NewLine{},
		&asciidoc.AttributeEntry{
			Name: "authorinitials",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "JCVD",
				},
			},
		},
	},
}

var parserTestUseExplicitAuthorinitialsIfSetAfterAuthorAttribute = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "author",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Jean-Claude Van Damme",
				},
			},
		},
		&asciidoc.AttributeEntry{
			Name: "authorinitials",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "JCVD",
				},
			},
		},
	},
}

var parserTestUseImplicitAuthorsIfValueOfAuthorsAttributeMatchesComputedValue = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "Doc Writer; Junior Writer",
		},
		&asciidoc.NewLine{},
		&asciidoc.AttributeEntry{
			Name: "authors",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Doc Writer, Junior Writer",
				},
			},
		},
	},
}

var parserTestReplaceImplicitAuthorsIfValueOfAuthorsAttributeDoesNotMatchComputedValue = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "Doc Writer; Junior Writer",
		},
		&asciidoc.NewLine{},
		&asciidoc.AttributeEntry{
			Name: "authors",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Stuart Rackham; Dan Allen; Sarah White",
				},
			},
		},
	},
}

var parserTestSetsAuthorcountTo0IfDocumentHasNoAuthors = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "Kismet Chameleon; Johnny Bravo; Lazarus het_Draeke",
		},
		&asciidoc.NewLine{},
		&asciidoc.AttributeEntry{
			Name: "author_2",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Danger Mouse",
				},
			},
		},
	},
}

var parserTestParseRevNumberDateRemark = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "Ryan Waldron",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "v0.0.7, 2013-12-18: The first release you can stand on",
		},
		&asciidoc.NewLine{},
	},
}

var parserTestParseRevNumberDataAndRemarkAsAttributeReferences = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "Author Name",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "v",
		},
		&asciidoc.UserAttributeReference{
			Value: "project-version",
		},
		&asciidoc.String{
			Value: ", ",
		},
		&asciidoc.UserAttributeReference{
			Value: "release-date",
		},
		&asciidoc.String{
			Value: ": ",
		},
		&asciidoc.UserAttributeReference{
			Value: "release-summary",
		},
		&asciidoc.NewLine{},
	},
}

var parserTestShouldResolveAttributeReferencesInRevNumberDataAndRemark = &asciidoc.Document{
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
				&asciidoc.UserAttributeReference{
					Value: "project-version",
				},
				&asciidoc.String{
					Value: ", ",
				},
				&asciidoc.UserAttributeReference{
					Value: "release-date",
				},
				&asciidoc.String{
					Value: ": ",
				},
				&asciidoc.UserAttributeReference{
					Value: "release-summary",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Document Title",
				},
			},
			Level: 0,
		},
	},
}

var parserTestParseRevDate = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "Ryan Waldron",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "2013-12-18",
		},
		&asciidoc.NewLine{},
	},
}

var parserTestParseRevNumberWithTrailingComma = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "Stuart Rackham",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "v8.6.8,",
		},
		&asciidoc.NewLine{},
	},
}

var parserTestParseRevNumber = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "Stuart Rackham",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "v8.6.8",
		},
		&asciidoc.NewLine{},
	},
}

var parserTestTreatsArbitraryTextOnRevLineAsRevdate = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "Ryan Waldron",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "foobar",
		},
		&asciidoc.NewLine{},
	},
}

var parserTestParseRevDateRemark = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "Ryan Waldron",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "2013-12-18:  The first release you can stand on",
		},
		&asciidoc.NewLine{},
	},
}

var parserTestShouldNotMistakeAttributeEntryAsRevRemark = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "Joe Cool",
		},
		&asciidoc.NewLine{},
		&asciidoc.AttributeEntry{
			Name: "page-layout",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "post",
				},
			},
		},
	},
}

var parserTestParseRevRemarkOnly = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "Joe Cool",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: " :Must start revremark-only line with space",
		},
		&asciidoc.NewLine{},
	},
}

var parserTestSkipLineCommentsBeforeAuthor = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.SingleLineComment{
			Value: " Asciidoctor",
		},
		&asciidoc.SingleLineComment{
			Value: " release artist",
		},
		&asciidoc.String{
			Value: "Ryan Waldron",
		},
		&asciidoc.NewLine{},
	},
}

var parserTestSkipBlockCommentBeforeAuthor = &asciidoc.Document{
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
				"Asciidoctor",
				"release artist",
			},
		},
		&asciidoc.String{
			Value: "Ryan Waldron",
		},
		&asciidoc.NewLine{},
	},
}

var parserTestSkipBlockCommentBeforeRev = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "Ryan Waldron",
		},
		&asciidoc.NewLine{},
		&asciidoc.MultiLineComment{
			Delimiter: asciidoc.Delimiter{
				Type:   2,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"Asciidoctor",
				"release info",
			},
		},
		&asciidoc.String{
			Value: "v0.0.7, 2013-12-18",
		},
		&asciidoc.NewLine{},
	},
}

var parserTestBreakHeaderAtLineWithThreeForwardSlashes = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "Joe Cool",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "v1.0",
		},
		&asciidoc.NewLine{},
		&asciidoc.SingleLineComment{
			Value: "/",
		},
		&asciidoc.String{
			Value: "stuff",
		},
		&asciidoc.NewLine{},
	},
}

var parserTestExpandsTabsToSpaces = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "Filesystem\t\t\t\tSize\tUsed\tAvail\tUse%\tMounted on",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "Filesystem              Size    Used    Avail   Use%    Mounted on",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "devtmpfs\t\t\t\t3.9G\t   0\t 3.9G\t  0%\t/dev",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "/dev/mapper/fedora-root\t 48G\t 18G\t  29G\t 39%\t/",
		},
		&asciidoc.NewLine{},
	},
}

var parserTestAdjustIndentationHandlesEmptyLinesGracefully = &asciidoc.Document{
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
								Value: "in-use",
							},
						},
					},
					Roles:   nil,
					Options: nil,
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "A paragraph with an id.",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 0,
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "Another paragraph",
		},
		&asciidoc.NewLine{},
		&asciidoc.Anchor{
			ID:       "in-use",
			Elements: nil,
		},
		&asciidoc.String{
			Value: "that uses an id",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "which is already in use.",
		},
		&asciidoc.NewLine{},
	},
}
