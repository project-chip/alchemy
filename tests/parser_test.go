package tests

import (
	"testing"

	"github.com/project-chip/alchemy/asciidoc"
)

func TestParser(t *testing.T) {
	parserTests.run(t)
}

var parserTests = parseTests{

	{"parse name with more than 3 parts in author attribute", "asciidoctor/parser_test_parse_name_with_more_than_3_parts_in_author_attribute.adoc", parserTestParseNameWithMoreThan3PartsInAuthorAttribute, nil},

	{"use explicit authorinitials if set after author attribute", "asciidoctor/parser_test_use_explicit_authorinitials_if_set_after_author_attribute.adoc", parserTestUseExplicitAuthorinitialsIfSetAfterAuthorAttribute, nil},

	{"use implicit authors if value of authors attribute matches computed value", "asciidoctor/parser_test_use_implicit_authors_if_value_of_authors_attribute_matches_computed_value.adoc", parserTestUseImplicitAuthorsIfValueOfAuthorsAttributeMatchesComputedValue, nil},

	{"replace implicit authors if value of authors attribute does not match computed value", "asciidoctor/parser_test_replace_implicit_authors_if_value_of_authors_attribute_does_not_match_computed_value.adoc", parserTestReplaceImplicitAuthorsIfValueOfAuthorsAttributeDoesNotMatchComputedValue, nil},

	{"sets authorcount to 0 if document has no authors", "asciidoctor/parser_test_sets_authorcount_to_0_if_document_has_no_authors.adoc", parserTestSetsAuthorcountTo0IfDocumentHasNoAuthors, nil},

	{"parse rev number date remark", "asciidoctor/parser_test_parse_rev_number_date_remark.adoc", parserTestParseRevNumberDateRemark, nil},

	{"parse rev number, data, and remark as attribute references", "asciidoctor/parser_test_parse_rev_number_data_and_remark_as_attribute_references.adoc", parserTestParseRevNumberDataAndRemarkAsAttributeReferences, nil},

	{"should resolve attribute references in rev number, data, and remark", "asciidoctor/parser_test_should_resolve_attribute_references_in_rev_number_data_and_remark.adoc", parserTestShouldResolveAttributeReferencesInRevNumberDataAndRemark, nil},

	{"parse rev date", "asciidoctor/parser_test_parse_rev_date.adoc", parserTestParseRevDate, nil},

	{"parse rev number with trailing comma", "asciidoctor/parser_test_parse_rev_number_with_trailing_comma.adoc", parserTestParseRevNumberWithTrailingComma, nil},

	{"parse rev number", "asciidoctor/parser_test_parse_rev_number.adoc", parserTestParseRevNumber, nil},

	{"treats arbitrary text on rev line as revdate", "asciidoctor/parser_test_treats_arbitrary_text_on_rev_line_as_revdate.adoc", parserTestTreatsArbitraryTextOnRevLineAsRevdate, nil},

	{"parse rev date remark", "asciidoctor/parser_test_parse_rev_date_remark.adoc", parserTestParseRevDateRemark, nil},

	{"should not mistake attribute entry as rev remark", "asciidoctor/parser_test_should_not_mistake_attribute_entry_as_rev_remark.adoc", parserTestShouldNotMistakeAttributeEntryAsRevRemark, nil},

	{"parse rev remark only", "asciidoctor/parser_test_parse_rev_remark_only.adoc", parserTestParseRevRemarkOnly, nil},

	{"skip line comments before author", "asciidoctor/parser_test_skip_line_comments_before_author.adoc", parserTestSkipLineCommentsBeforeAuthor, nil},

	{"skip block comment before author", "asciidoctor/parser_test_skip_block_comment_before_author.adoc", parserTestSkipBlockCommentBeforeAuthor, nil},

	{"skip block comment before rev", "asciidoctor/parser_test_skip_block_comment_before_rev.adoc", parserTestSkipBlockCommentBeforeRev, nil},

	{"break header at line with three forward slashes", "asciidoctor/parser_test_break_header_at_line_with_three_forward_slashes.adoc", parserTestBreakHeaderAtLineWithThreeForwardSlashes, nil},

	{"expands tabs to spaces", "asciidoctor/parser_test_expands_tabs_to_spaces.adoc", parserTestExpandsTabsToSpaces, nil},

	{"adjust indentation handles empty lines gracefully", "asciidoctor/parser_test_adjust_indentation_handles_empty_lines_gracefully.adoc", parserTestAdjustIndentationHandlesEmptyLinesGracefully, nil},
}

var parserTestParseNameWithMoreThan3PartsInAuthorAttribute = &asciidoc.Document{
	Set: asciidoc.Set{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "Jean-Claude Van Damme",
		},
		&asciidoc.NewLine{},
		&asciidoc.AttributeEntry{
			Name: "authorinitials",
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "JCVD",
				},
			},
		},
	},
}

var parserTestUseExplicitAuthorinitialsIfSetAfterAuthorAttribute = &asciidoc.Document{
	Set: asciidoc.Set{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "author",
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Jean-Claude Van Damme",
				},
			},
		},
		&asciidoc.AttributeEntry{
			Name: "authorinitials",
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "JCVD",
				},
			},
		},
	},
}

var parserTestUseImplicitAuthorsIfValueOfAuthorsAttributeMatchesComputedValue = &asciidoc.Document{
	Set: asciidoc.Set{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "Doc Writer; Junior Writer",
		},
		&asciidoc.NewLine{},
		&asciidoc.AttributeEntry{
			Name: "authors",
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Doc Writer, Junior Writer",
				},
			},
		},
	},
}

var parserTestReplaceImplicitAuthorsIfValueOfAuthorsAttributeDoesNotMatchComputedValue = &asciidoc.Document{
	Set: asciidoc.Set{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "Doc Writer; Junior Writer",
		},
		&asciidoc.NewLine{},
		&asciidoc.AttributeEntry{
			Name: "authors",
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Stuart Rackham; Dan Allen; Sarah White",
				},
			},
		},
	},
}

var parserTestSetsAuthorcountTo0IfDocumentHasNoAuthors = &asciidoc.Document{
	Set: asciidoc.Set{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "Kismet Chameleon; Johnny Bravo; Lazarus het_Draeke",
		},
		&asciidoc.NewLine{},
		&asciidoc.AttributeEntry{
			Name: "author_2",
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Danger Mouse",
				},
			},
		},
	},
}

var parserTestParseRevNumberDateRemark = &asciidoc.Document{
	Set: asciidoc.Set{
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
	Set: asciidoc.Set{
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
	Set: asciidoc.Set{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Set: asciidoc.Set{
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
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Document Title",
				},
			},
			Level: 0,
		},
	},
}

var parserTestParseRevDate = &asciidoc.Document{
	Set: asciidoc.Set{
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
	Set: asciidoc.Set{
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
	Set: asciidoc.Set{
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
	Set: asciidoc.Set{
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
	Set: asciidoc.Set{
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
	Set: asciidoc.Set{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "Joe Cool",
		},
		&asciidoc.NewLine{},
		&asciidoc.AttributeEntry{
			Name: "page-layout",
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "post",
				},
			},
		},
	},
}

var parserTestParseRevRemarkOnly = &asciidoc.Document{
	Set: asciidoc.Set{
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
	Set: asciidoc.Set{
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
	Set: asciidoc.Set{
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
	Set: asciidoc.Set{
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
	Set: asciidoc.Set{
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
	Set: asciidoc.Set{
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
	Set: asciidoc.Set{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Paragraph{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: nil,
					ID: &asciidoc.ShorthandID{
						Set: asciidoc.Set{
							&asciidoc.String{
								Value: "in-use",
							},
						},
					},
					Roles:   nil,
					Options: nil,
				},
			},
			Set: asciidoc.Set{
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
			ID:  "in-use",
			Set: nil,
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
