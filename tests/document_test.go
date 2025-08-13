package tests

import (
	"testing"

	"github.com/project-chip/alchemy/asciidoc"
)

func TestDocument(t *testing.T) {
	documentTests.run(t)
}

var documentTests = parseTests{

	{"should be able to disable toc and sectnums in document header in DocBook backend", "asciidoctor/document_test_should_be_able_to_disable_toc_and_sectnums_in_document_header_in_doc_book_backend.adoc", documentTestShouldBeAbleToDisableTocAndSectnumsInDocumentHeaderInDocBookBackend},

	{"noheader attribute should suppress info element when converting to DocBook", "asciidoctor/document_test_noheader_attribute_should_suppress_info_element_when_converting_to_doc_book.adoc", documentTestNoheaderAttributeShouldSuppressInfoElementWhenConvertingToDocBook},

	{"should be able to disable section numbering using numbered attribute in document header in DocBook backend", "asciidoctor/document_test_should_be_able_to_disable_section_numbering_using_numbered_attribute_in_document_header_in_doc_book_backend.adoc", documentTestShouldBeAbleToDisableSectionNumberingUsingNumberedAttributeInDocumentHeaderInDocBookBackend},

	{"convert methods on built-in converter are registered by default", "asciidoctor/document_test_convert_methods_on_built_in_converter_are_registered_by_default.adoc", documentTestConvertMethodsOnBuiltInConverterAreRegisteredByDefault},

	{"should not enable compat mode for document with legacy doctitle if compat mode disable by header", "asciidoctor/document_test_should_not_enable_compat_mode_for_document_with_legacy_doctitle_if_compat_mode_disable_by_header.adoc", documentTestShouldNotEnableCompatModeForDocumentWithLegacyDoctitleIfCompatModeDisableByHeader},

	{"should not enable compat mode for document with legacy doctitle if compat mode is locked by API", "asciidoctor/document_test_should_not_enable_compat_mode_for_document_with_legacy_doctitle_if_compat_mode_is_locked_by_api.adoc", documentTestShouldNotEnableCompatModeForDocumentWithLegacyDoctitleIfCompatModeIsLockedByApi},

	{"should apply max-width to each top-level container", "asciidoctor/document_test_should_apply_max_width_to_each_top_level_container.adoc", documentTestShouldApplyMaxWidthToEachTopLevelContainer},

	{"title partition API with default separator", "asciidoctor/document_test_title_partition_api_with_default_separator.adoc", documentTestTitlePartitionApiWithDefaultSeparator},

	{"document with subtitle and custom separator", "asciidoctor/document_test_document_with_subtitle_and_custom_separator.adoc", documentTestDocumentWithSubtitleAndCustomSeparator},

	{"should not honor custom separator for doctitle if attribute is locked by API", "asciidoctor/document_test_should_not_honor_custom_separator_for_doctitle_if_attribute_is_locked_by_api.adoc", documentTestShouldNotHonorCustomSeparatorForDoctitleIfAttributeIsLockedByApi},

	{"document with doctitle defined as attribute entry", "asciidoctor/document_test_document_with_doctitle_defined_as_attribute_entry.adoc", documentTestDocumentWithDoctitleDefinedAsAttributeEntry},

	{"document with doctitle defined as attribute entry followed by block with title", "asciidoctor/document_test_document_with_doctitle_defined_as_attribute_entry_followed_by_block_with_title.adoc", documentTestDocumentWithDoctitleDefinedAsAttributeEntryFollowedByBlockWithTitle},

	{"document with title attribute entry overrides doctitle", "asciidoctor/document_test_document_with_title_attribute_entry_overrides_doctitle.adoc", documentTestDocumentWithTitleAttributeEntryOverridesDoctitle},

	{"document with blank title attribute entry overrides doctitle", "asciidoctor/document_test_document_with_blank_title_attribute_entry_overrides_doctitle.adoc", documentTestDocumentWithBlankTitleAttributeEntryOverridesDoctitle},

	{"document header can reference intrinsic doctitle attribute", "asciidoctor/document_test_document_header_can_reference_intrinsic_doctitle_attribute.adoc", documentTestDocumentHeaderCanReferenceIntrinsicDoctitleAttribute},

	{"document with title attribute entry overrides doctitle attribute entry", "asciidoctor/document_test_document_with_title_attribute_entry_overrides_doctitle_attribute_entry.adoc", documentTestDocumentWithTitleAttributeEntryOverridesDoctitleAttributeEntry},

	{"document with doctitle attribute entry overrides implicit doctitle", "asciidoctor/document_test_document_with_doctitle_attribute_entry_overrides_implicit_doctitle.adoc", documentTestDocumentWithDoctitleAttributeEntryOverridesImplicitDoctitle},

	{"doctitle attribute entry above header overrides implicit doctitle", "asciidoctor/document_test_doctitle_attribute_entry_above_header_overrides_implicit_doctitle.adoc", documentTestDoctitleAttributeEntryAboveHeaderOverridesImplicitDoctitle},

	{"should apply header substitutions to value of the doctitle attribute assigned from implicit doctitle", "asciidoctor/document_test_should_apply_header_substitutions_to_value_of_the_doctitle_attribute_assigned_from_implicit_doctitle.adoc", documentTestShouldApplyHeaderSubstitutionsToValueOfTheDoctitleAttributeAssignedFromImplicitDoctitle},

	{"should substitute attribute reference in implicit document title for attribute defined earlier in header", "asciidoctor/document_test_should_substitute_attribute_reference_in_implicit_document_title_for_attribute_defined_earlier_in_header.adoc", documentTestShouldSubstituteAttributeReferenceInImplicitDocumentTitleForAttributeDefinedEarlierInHeader},

	{"should not warn if implicit document title contains attribute reference for attribute defined later in header", "asciidoctor/document_test_should_not_warn_if_implicit_document_title_contains_attribute_reference_for_attribute_defined_later_in_header.adoc", documentTestShouldNotWarnIfImplicitDocumentTitleContainsAttributeReferenceForAttributeDefinedLaterInHeader},

	{"should recognize document title when preceded by blank lines", "asciidoctor/document_test_should_recognize_document_title_when_preceded_by_blank_lines.adoc", documentTestShouldRecognizeDocumentTitleWhenPrecededByBlankLines},

	{"should recognize document title when preceded by blank lines introduced by a preprocessor conditional", "asciidoctor/document_test_should_recognize_document_title_when_preceded_by_blank_lines_introduced_by_a_preprocessor_conditional.adoc", documentTestShouldRecognizeDocumentTitleWhenPrecededByBlankLinesIntroducedByAPreprocessorConditional},

	{"should recognize document title when preceded by blank lines after an attribute entry", "asciidoctor/document_test_should_recognize_document_title_when_preceded_by_blank_lines_after_an_attribute_entry.adoc", documentTestShouldRecognizeDocumentTitleWhenPrecededByBlankLinesAfterAnAttributeEntry},

	{"should recognize document title in include file when preceded by blank lines", "asciidoctor/document_test_should_recognize_document_title_in_include_file_when_preceded_by_blank_lines.adoc", documentTestShouldRecognizeDocumentTitleInIncludeFileWhenPrecededByBlankLines},

	{"should include specified lines even when leading lines are skipped", "asciidoctor/document_test_should_include_specified_lines_even_when_leading_lines_are_skipped.adoc", documentTestShouldIncludeSpecifiedLinesEvenWhenLeadingLinesAreSkipped},

	{"document with multiline attribute entry but only one line should not crash", "asciidoctor/document_test_document_with_multiline_attribute_entry_but_only_one_line_should_not_crash.adoc", documentTestDocumentWithMultilineAttributeEntryButOnlyOneLineShouldNotCrash},

	{"should not choke on empty source", "asciidoctor/document_test_should_not_choke_on_empty_source.adoc", documentTestShouldNotChokeOnEmptySource},

	{"should parse revision line if date is empty", "asciidoctor/document_test_should_parse_revision_line_if_date_is_empty.adoc", documentTestShouldParseRevisionLineIfDateIsEmpty},

	{"should include revision history in DocBook output if revdate and revnumber is set", "asciidoctor/document_test_should_include_revision_history_in_doc_book_output_if_revdate_and_revnumber_is_set.adoc", documentTestShouldIncludeRevisionHistoryInDocBookOutputIfRevdateAndRevnumberIsSet},

	{"should include revision history in DocBook output if revdate and revremark is set", "asciidoctor/document_test_should_include_revision_history_in_doc_book_output_if_revdate_and_revremark_is_set.adoc", documentTestShouldIncludeRevisionHistoryInDocBookOutputIfRevdateAndRevremarkIsSet},

	{"should not include revision history in DocBook output if revdate is not set", "asciidoctor/document_test_should_not_include_revision_history_in_doc_book_output_if_revdate_is_not_set.adoc", documentTestShouldNotIncludeRevisionHistoryInDocBookOutputIfRevdateIsNotSet},

	{"with metadata to DocBook 5", "asciidoctor/document_test_with_metadata_to_doc_book_5.adoc", documentTestWithMetadataToDocBook5},

	{"with document ID to Docbook 5", "asciidoctor/document_test_with_document_id_to_docbook_5.adoc", documentTestWithDocumentIdToDocbook5},

	{"with author defined using attribute entry to DocBook", "asciidoctor/document_test_with_author_defined_using_attribute_entry_to_doc_book.adoc", documentTestWithAuthorDefinedUsingAttributeEntryToDocBook},

	{"should substitute replacements in author names in HTML output", "asciidoctor/document_test_should_substitute_replacements_in_author_names_in_html_output.adoc", documentTestShouldSubstituteReplacementsInAuthorNamesInHtmlOutput},

	{"should substitute replacements in author names in DocBook output", "asciidoctor/document_test_should_substitute_replacements_in_author_names_in_doc_book_output.adoc", documentTestShouldSubstituteReplacementsInAuthorNamesInDocBookOutput},

	{"should sanitize content of HTML meta authors tag", "asciidoctor/document_test_should_sanitize_content_of_html_meta_authors_tag.adoc", documentTestShouldSanitizeContentOfHtmlMetaAuthorsTag},

	{"should not double escape ampersand in author attribute", "asciidoctor/document_test_should_not_double_escape_ampersand_in_author_attribute.adoc", documentTestShouldNotDoubleEscapeAmpersandInAuthorAttribute},

	{"should include multiple authors in HTML output", "asciidoctor/document_test_should_include_multiple_authors_in_html_output.adoc", documentTestShouldIncludeMultipleAuthorsInHtmlOutput},

	{"should create authorgroup in DocBook when multiple authors", "asciidoctor/document_test_should_create_authorgroup_in_doc_book_when_multiple_authors.adoc", documentTestShouldCreateAuthorgroupInDocBookWhenMultipleAuthors},

	{"should process author defined by attribute when implicit doctitle is absent", "asciidoctor/document_test_should_process_author_defined_by_attribute_when_implicit_doctitle_is_absent.adoc", documentTestShouldProcessAuthorDefinedByAttributeWhenImplicitDoctitleIsAbsent},

	{"should process author and authorinitials defined by attribute when implicit doctitle is absent", "asciidoctor/document_test_should_process_author_and_authorinitials_defined_by_attribute_when_implicit_doctitle_is_absent.adoc", documentTestShouldProcessAuthorAndAuthorinitialsDefinedByAttributeWhenImplicitDoctitleIsAbsent},

	{"should process authors defined by attribute when implicit doctitle is absent", "asciidoctor/document_test_should_process_authors_defined_by_attribute_when_implicit_doctitle_is_absent.adoc", documentTestShouldProcessAuthorsDefinedByAttributeWhenImplicitDoctitleIsAbsent},

	{"should process authors and authorinitials defined by attribute when implicit doctitle is absent", "asciidoctor/document_test_should_process_authors_and_authorinitials_defined_by_attribute_when_implicit_doctitle_is_absent.adoc", documentTestShouldProcessAuthorsAndAuthorinitialsDefinedByAttributeWhenImplicitDoctitleIsAbsent},

	{"should set authorcount to 0 if document has no header", "asciidoctor/document_test_should_set_authorcount_to_0_if_document_has_no_header.adoc", documentTestShouldSetAuthorcountTo0IfDocumentHasNoHeader},

	{"should set authorcount to 0 if author not set by attribute and document starts with level-0 section with style", "asciidoctor/document_test_should_set_authorcount_to_0_if_author_not_set_by_attribute_and_document_starts_with_level_0_section_with_style.adoc", documentTestShouldSetAuthorcountTo0IfAuthorNotSetByAttributeAndDocumentStartsWithLevel0SectionWithStyle},

	{"with author defined by indexed attribute name", "asciidoctor/document_test_with_author_defined_by_indexed_attribute_name.adoc", documentTestWithAuthorDefinedByIndexedAttributeName},

	{"with authors defined using attribute entry to DocBook", "asciidoctor/document_test_with_authors_defined_using_attribute_entry_to_doc_book.adoc", documentTestWithAuthorsDefinedUsingAttributeEntryToDocBook},

	{"should populate copyright element in DocBook output if copyright attribute is defined", "asciidoctor/document_test_should_populate_copyright_element_in_doc_book_output_if_copyright_attribute_is_defined.adoc", documentTestShouldPopulateCopyrightElementInDocBookOutputIfCopyrightAttributeIsDefined},

	{"should populate copyright element in DocBook output if copyright attribute is defined with year", "asciidoctor/document_test_should_populate_copyright_element_in_doc_book_output_if_copyright_attribute_is_defined_with_year.adoc", documentTestShouldPopulateCopyrightElementInDocBookOutputIfCopyrightAttributeIsDefinedWithYear},

	{"should populate copyright element in DocBook output if copyright attribute is defined with year range", "asciidoctor/document_test_should_populate_copyright_element_in_doc_book_output_if_copyright_attribute_is_defined_with_year_range.adoc", documentTestShouldPopulateCopyrightElementInDocBookOutputIfCopyrightAttributeIsDefinedWithYearRange},

	{"with header footer", "asciidoctor/document_test_with_header_footer.adoc", documentTestWithHeaderFooter},

	{"can disable last updated in footer", "asciidoctor/document_test_can_disable_last_updated_in_footer.adoc", documentTestCanDisableLastUpdatedInFooter},

	{"parse header only", "asciidoctor/document_test_parse_header_only.adoc", documentTestParseHeaderOnly},

	{"should parse header only when docytpe is manpage", "asciidoctor/document_test_should_parse_header_only_when_docytpe_is_manpage.adoc", documentTestShouldParseHeaderOnlyWhenDocytpeIsManpage},

	{"should not warn when parsing header only when docytpe is manpage and body is empty", "asciidoctor/document_test_should_not_warn_when_parsing_header_only_when_docytpe_is_manpage_and_body_is_empty.adoc", documentTestShouldNotWarnWhenParsingHeaderOnlyWhenDocytpeIsManpageAndBodyIsEmpty},

	{"outputs footnotes in footer", "asciidoctor/document_test_outputs_footnotes_in_footer.adoc", documentTestOutputsFootnotesInFooter},

	{"outputs footnotes block in embedded document by default", "asciidoctor/document_test_outputs_footnotes_block_in_embedded_document_by_default.adoc", documentTestOutputsFootnotesBlockInEmbeddedDocumentByDefault},

	{"should return empty :ids table", "asciidoctor/document_test_should_return_empty_ids_table.adoc", documentTestShouldReturnEmptyidsTable},

	{"honor htmlsyntax attribute in document header if followed by backend attribute", "asciidoctor/document_test_honor_htmlsyntax_attribute_in_document_header_if_followed_by_backend_attribute.adoc", documentTestHonorHtmlsyntaxAttributeInDocumentHeaderIfFollowedByBackendAttribute},

	{"does not honor htmlsyntax attribute in document header if not followed by backend attribute", "asciidoctor/document_test_does_not_honor_htmlsyntax_attribute_in_document_header_if_not_followed_by_backend_attribute.adoc", documentTestDoesNotHonorHtmlsyntaxAttributeInDocumentHeaderIfNotFollowedByBackendAttribute},

	{"should close all short tags when htmlsyntax is xml", "asciidoctor/document_test_should_close_all_short_tags_when_htmlsyntax_is_xml.adoc", documentTestShouldCloseAllShortTagsWhenHtmlsyntaxIsXml},

	{"xhtml backend should emit elements in proper namespace", "asciidoctor/document_test_xhtml_backend_should_emit_elements_in_proper_namespace.adoc", documentTestXhtmlBackendShouldEmitElementsInProperNamespace},

	{"should be able to set doctype to article when converting to DocBook", "asciidoctor/document_test_should_be_able_to_set_doctype_to_article_when_converting_to_doc_book.adoc", documentTestShouldBeAbleToSetDoctypeToArticleWhenConvertingToDocBook},

	{"should set doctype to article by default for document with no title when converting to DocBook", "asciidoctor/document_test_should_set_doctype_to_article_by_default_for_document_with_no_title_when_converting_to_doc_book.adoc", documentTestShouldSetDoctypeToArticleByDefaultForDocumentWithNoTitleWhenConvertingToDocBook},

	{"should output non-breaking space for source and manual in docbook manpage output if absent from source", "asciidoctor/document_test_should_output_non_breaking_space_for_source_and_manual_in_docbook_manpage_output_if_absent_from_source.adoc", documentTestShouldOutputNonBreakingSpaceForSourceAndManualInDocbookManpageOutputIfAbsentFromSource},

	{"should apply replacements substitution to value of mantitle attribute used in DocBook output", "asciidoctor/document_test_should_apply_replacements_substitution_to_value_of_mantitle_attribute_used_in_doc_book_output.adoc", documentTestShouldApplyReplacementsSubstitutionToValueOfMantitleAttributeUsedInDocBookOutput},

	{"should be able to set doctype to book when converting to DocBook", "asciidoctor/document_test_should_be_able_to_set_doctype_to_book_when_converting_to_doc_book.adoc", documentTestShouldBeAbleToSetDoctypeToBookWhenConvertingToDocBook},

	{"should be able to set doctype to book for document with no title when converting to DocBook", "asciidoctor/document_test_should_be_able_to_set_doctype_to_book_for_document_with_no_title_when_converting_to_doc_book.adoc", documentTestShouldBeAbleToSetDoctypeToBookForDocumentWithNoTitleWhenConvertingToDocBook},

	{"adds a front and back cover image to DocBook 5 when doctype is book", "asciidoctor/document_test_adds_a_front_and_back_cover_image_to_doc_book_5_when_doctype_is_book.adoc", documentTestAddsAFrontAndBackCoverImageToDocBook5WhenDoctypeIsBook},

	{"should be able to set backend using :backend option key", "asciidoctor/document_test_should_be_able_to_set_backend_using_backend_option_key.adoc", documentTestShouldBeAbleToSetBackendUsingbackendOptionKey},

	{"attribute entry can appear immediately after document title", "asciidoctor/document_test_attribute_entry_can_appear_immediately_after_document_title.adoc", documentTestAttributeEntryCanAppearImmediatelyAfterDocumentTitle},

	{"attribute entry can appear before author line under document title", "asciidoctor/document_test_attribute_entry_can_appear_before_author_line_under_document_title.adoc", documentTestAttributeEntryCanAppearBeforeAuthorLineUnderDocumentTitle},

	{"should parse mantitle and manvolnum from document title for manpage doctype", "asciidoctor/document_test_should_parse_mantitle_and_manvolnum_from_document_title_for_manpage_doctype.adoc", documentTestShouldParseMantitleAndManvolnumFromDocumentTitleForManpageDoctype},

	{"should perform attribute substitution on mantitle in manpage doctype", "asciidoctor/document_test_should_perform_attribute_substitution_on_mantitle_in_manpage_doctype.adoc", documentTestShouldPerformAttributeSubstitutionOnMantitleInManpageDoctype},

	{"should consume name section as manname and manpurpose for manpage doctype", "asciidoctor/document_test_should_consume_name_section_as_manname_and_manpurpose_for_manpage_doctype.adoc", documentTestShouldConsumeNameSectionAsMannameAndManpurposeForManpageDoctype},

	{"should set docname and outfilesuffix from manname and manvolnum for manpage backend and doctype", "asciidoctor/document_test_should_set_docname_and_outfilesuffix_from_manname_and_manvolnum_for_manpage_backend_and_doctype.adoc", documentTestShouldSetDocnameAndOutfilesuffixFromMannameAndManvolnumForManpageBackendAndDoctype},

	{"should mark synopsis as special section in manpage doctype", "asciidoctor/document_test_should_mark_synopsis_as_special_section_in_manpage_doctype.adoc", documentTestShouldMarkSynopsisAsSpecialSectionInManpageDoctype},

	{"should output special header block in HTML for manpage doctype", "asciidoctor/document_test_should_output_special_header_block_in_html_for_manpage_doctype.adoc", documentTestShouldOutputSpecialHeaderBlockInHtmlForManpageDoctype},

	{"should output special header block in embeddable HTML for manpage doctype", "asciidoctor/document_test_should_output_special_header_block_in_embeddable_html_for_manpage_doctype.adoc", documentTestShouldOutputSpecialHeaderBlockInEmbeddableHtmlForManpageDoctype},

	{"should output all mannames in name section in man page output", "asciidoctor/document_test_should_output_all_mannames_in_name_section_in_man_page_output.adoc", documentTestShouldOutputAllMannamesInNameSectionInManPageOutput},

	{"allows us to specify a path relative to the current dir", "asciidoctor/document_test_allows_us_to_specify_a_path_relative_to_the_current_dir.adoc", documentTestAllowsUsToSpecifyAPathRelativeToTheCurrentDir},

	{"should raise an exception when a converter cannot be resolved while parsing", "asciidoctor/document_test_should_raise_an_exception_when_a_converter_cannot_be_resolved_while_parsing.adoc", documentTestShouldRaiseAnExceptionWhenAConverterCannotBeResolvedWhileParsing},
}

var documentTestShouldBeAbleToDisableTocAndSectnumsInDocumentHeaderInDocBookBackend = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.AttributeReset{
					Name: "toc",
				},
				&asciidoc.AttributeReset{
					Name: "sectnums",
				},
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

var documentTestNoheaderAttributeShouldSuppressInfoElementWhenConvertingToDocBook = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.AttributeEntry{
					Name:     "noheader",
					Elements: nil,
				},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "content",
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

var documentTestShouldBeAbleToDisableSectionNumberingUsingNumberedAttributeInDocumentHeaderInDocBookBackend = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.AttributeReset{
					Name: "numbered",
				},
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

var documentTestConvertMethodsOnBuiltInConverterAreRegisteredByDefault = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "Document Title",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "==============",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.InlinePassthrough{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "content",
				},
			},
		},
		&asciidoc.NewLine{},
	},
}

var documentTestShouldNotEnableCompatModeForDocumentWithLegacyDoctitleIfCompatModeDisableByHeader = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "Document Title",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "==============",
		},
		&asciidoc.NewLine{},
		&asciidoc.AttributeReset{
			Name: "compat-mode",
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.InlinePassthrough{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "content",
				},
			},
		},
		&asciidoc.NewLine{},
	},
}

var documentTestShouldNotEnableCompatModeForDocumentWithLegacyDoctitleIfCompatModeIsLockedByApi = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "Document Title",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "==============",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.InlinePassthrough{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "content",
				},
			},
		},
		&asciidoc.NewLine{},
	},
}

var documentTestShouldApplyMaxWidthToEachTopLevelContainer = &asciidoc.Document{
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
					Value: "contentfootnote:[placeholder]",
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

var documentTestTitlePartitionApiWithDefaultSeparator = &asciidoc.Document{
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
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "content",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Main Title: *Subtitle*",
				},
			},
			Level: 0,
		},
	},
}

var documentTestDocumentWithSubtitleAndCustomSeparator = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.NamedAttribute{
					Name: "separator",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "::",
						},
					},
					Quote: 0,
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Author Name",
				},
				&asciidoc.NewLine{},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "content",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Main Title:: *Subtitle*",
				},
			},
			Level: 0,
		},
	},
}

var documentTestShouldNotHonorCustomSeparatorForDoctitleIfAttributeIsLockedByApi = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.NamedAttribute{
					Name: "separator",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "::",
						},
					},
					Quote: 0,
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Author Name",
				},
				&asciidoc.NewLine{},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "content",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Main Title - *Subtitle*",
				},
			},
			Level: 0,
		},
	},
}

var documentTestDocumentWithDoctitleDefinedAsAttributeEntry = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "doctitle",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Document Title",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "preamble",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Elements:      nil,
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "First Section",
				},
			},
			Level: 1,
		},
	},
}

var documentTestDocumentWithDoctitleDefinedAsAttributeEntryFollowedByBlockWithTitle = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "doctitle",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Document Title",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Paragraph{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.TitleAttribute{
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Block title",
						},
					},
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Block content",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 0,
		},
	},
}

var documentTestDocumentWithTitleAttributeEntryOverridesDoctitle = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.AttributeEntry{
					Name: "title",
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "Override",
						},
					},
				},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.UserAttributeReference{
					Value: "doctitle",
				},
				&asciidoc.NewLine{},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.Section{
					AttributeList: nil,
					Elements:      nil,
					Title: asciidoc.Elements{
						&asciidoc.String{
							Value: "First Section",
						},
					},
					Level: 1,
				},
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

var documentTestDocumentWithBlankTitleAttributeEntryOverridesDoctitle = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.AttributeEntry{
					Name:     "title",
					Elements: nil,
				},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.UserAttributeReference{
					Value: "doctitle",
				},
				&asciidoc.NewLine{},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.Section{
					AttributeList: nil,
					Elements:      nil,
					Title: asciidoc.Elements{
						&asciidoc.String{
							Value: "First Section",
						},
					},
					Level: 1,
				},
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

var documentTestDocumentHeaderCanReferenceIntrinsicDoctitleAttribute = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.AttributeEntry{
					Name: "intro",
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "Welcome to the {doctitle}!",
						},
					},
				},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.UserAttributeReference{
					Value: "intro",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "ACME Documentation",
				},
			},
			Level: 0,
		},
	},
}

var documentTestDocumentWithTitleAttributeEntryOverridesDoctitleAttributeEntry = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.AttributeEntry{
					Name: "snapshot",
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "{doctitle}",
						},
					},
				},
				&asciidoc.AttributeEntry{
					Name: "doctitle",
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "doctitle",
						},
					},
				},
				&asciidoc.AttributeEntry{
					Name: "title",
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "Override",
						},
					},
				},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.UserAttributeReference{
					Value: "snapshot",
				},
				&asciidoc.String{
					Value: ", ",
				},
				&asciidoc.UserAttributeReference{
					Value: "doctitle",
				},
				&asciidoc.NewLine{},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.Section{
					AttributeList: nil,
					Elements:      nil,
					Title: asciidoc.Elements{
						&asciidoc.String{
							Value: "First Section",
						},
					},
					Level: 1,
				},
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

var documentTestDocumentWithDoctitleAttributeEntryOverridesImplicitDoctitle = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.AttributeEntry{
					Name: "snapshot",
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "{doctitle}",
						},
					},
				},
				&asciidoc.AttributeEntry{
					Name: "doctitle",
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "Override",
						},
					},
				},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.UserAttributeReference{
					Value: "snapshot",
				},
				&asciidoc.String{
					Value: ", ",
				},
				&asciidoc.UserAttributeReference{
					Value: "doctitle",
				},
				&asciidoc.NewLine{},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.Section{
					AttributeList: nil,
					Elements:      nil,
					Title: asciidoc.Elements{
						&asciidoc.String{
							Value: "First Section",
						},
					},
					Level: 1,
				},
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

var documentTestDoctitleAttributeEntryAboveHeaderOverridesImplicitDoctitle = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "doctitle",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Override",
				},
			},
		},
		&asciidoc.Section{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.UserAttributeReference{
					Value: "doctitle",
				},
				&asciidoc.NewLine{},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.Section{
					AttributeList: nil,
					Elements:      nil,
					Title: asciidoc.Elements{
						&asciidoc.String{
							Value: "First Section",
						},
					},
					Level: 1,
				},
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

var documentTestShouldApplyHeaderSubstitutionsToValueOfTheDoctitleAttributeAssignedFromImplicitDoctitle = &asciidoc.Document{
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
					Value: "The name of the game is ",
				},
				&asciidoc.UserAttributeReference{
					Value: "doctitle",
				},
				&asciidoc.String{
					Value: ".",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "<Foo> {plus} <Bar>",
				},
			},
			Level: 0,
		},
	},
}

var documentTestShouldSubstituteAttributeReferenceInImplicitDocumentTitleForAttributeDefinedEarlierInHeader = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "project-name",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "ACME",
				},
			},
		},
		&asciidoc.Section{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.UserAttributeReference{
					Value: "doctitle",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "{project-name} Docs",
				},
			},
			Level: 0,
		},
	},
}

var documentTestShouldNotWarnIfImplicitDocumentTitleContainsAttributeReferenceForAttributeDefinedLaterInHeader = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.AttributeEntry{
					Name: "project-name",
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "ACME",
						},
					},
				},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.UserAttributeReference{
					Value: "doctitle",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "{project-name} Docs",
				},
			},
			Level: 0,
		},
	},
}

var documentTestShouldRecognizeDocumentTitleWhenPrecededByBlankLines = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
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
					Value: "preamble",
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
							Value: "text",
						},
						&asciidoc.NewLine{},
					},
					Title: asciidoc.Elements{
						&asciidoc.String{
							Value: "Section 1",
						},
					},
					Level: 1,
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Title",
				},
			},
			Level: 0,
		},
	},
}

var documentTestShouldRecognizeDocumentTitleWhenPrecededByBlankLinesIntroducedByAPreprocessorConditional = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.IfDef{
			Attributes: asciidoc.AttributeNames{
				"sectids",
			},
			Union:  0,
			Inline: false,
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "foo",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "bar",
				},
			},
		},
		&asciidoc.EndIf{
			Attributes: nil,
			Union:      0,
			Open:       nil,
		},
		&asciidoc.Section{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "preamble",
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
							Value: "text",
						},
						&asciidoc.NewLine{},
					},
					Title: asciidoc.Elements{
						&asciidoc.String{
							Value: "Section 1",
						},
					},
					Level: 1,
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Title",
				},
			},
			Level: 0,
		},
	},
}

var documentTestShouldRecognizeDocumentTitleWhenPrecededByBlankLinesAfterAnAttributeEntry = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "doctype",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "book",
				},
			},
		},
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
					Value: "preamble",
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
							Value: "text",
						},
						&asciidoc.NewLine{},
					},
					Title: asciidoc.Elements{
						&asciidoc.String{
							Value: "Section 1",
						},
					},
					Level: 1,
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Title",
				},
			},
			Level: 0,
		},
	},
}

var documentTestShouldRecognizeDocumentTitleInIncludeFileWhenPrecededByBlankLines = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.FileInclude{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "fixtures/include-with-leading-blank-line.adoc",
				},
			},
		},
	},
}

var documentTestShouldIncludeSpecifiedLinesEvenWhenLeadingLinesAreSkipped = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.FileInclude{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.NamedAttribute{
					Name: "lines",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "6",
						},
					},
					Quote: 0,
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "fixtures/include-with-leading-blank-line.adoc",
				},
			},
		},
	},
}

var documentTestDocumentWithMultilineAttributeEntryButOnlyOneLineShouldNotCrash = &asciidoc.Document{
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
					Value: "content",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "*Document* image:logo.png[] _Title_ image:another-logo.png[another logo]",
				},
			},
			Level: 0,
		},
	},
}

var documentTestShouldNotChokeOnEmptySource = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Stuart Rackham <",
				},
				asciidoc.Email{
					Address: "founder@asciidoc.org",
				},
				&asciidoc.String{
					Value: ">",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "v8.6.8, 2012-07-12: See changelog.",
				},
				&asciidoc.NewLine{},
				&asciidoc.AttributeEntry{
					Name: "description",
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "AsciiDoc user guide",
						},
					},
				},
				&asciidoc.AttributeEntry{
					Name: "keywords",
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "asciidoc,documentation",
						},
					},
				},
				&asciidoc.AttributeEntry{
					Name: "copyright",
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "Stuart Rackham",
						},
					},
				},
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
							Value: "more info...",
						},
						&asciidoc.NewLine{},
					},
					Title: asciidoc.Elements{
						&asciidoc.String{
							Value: "Version 8.6.8",
						},
					},
					Level: 1,
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "AsciiDoc",
				},
			},
			Level: 0,
		},
	},
}

var documentTestShouldParseRevisionLineIfDateIsEmpty = &asciidoc.Document{
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
				&asciidoc.String{
					Value: "v1.0.0,:remark",
				},
				&asciidoc.NewLine{},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "content",
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

var documentTestShouldIncludeRevisionHistoryInDocBookOutputIfRevdateAndRevnumberIsSet = &asciidoc.Document{
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
					Name: "revdate",
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "2011-11-11",
						},
					},
				},
				&asciidoc.AttributeEntry{
					Name: "revnumber",
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "1.0",
						},
					},
				},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "content",
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

var documentTestShouldIncludeRevisionHistoryInDocBookOutputIfRevdateAndRevremarkIsSet = &asciidoc.Document{
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
					Name: "revdate",
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "2011-11-11",
						},
					},
				},
				&asciidoc.AttributeEntry{
					Name: "revremark",
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "features!",
						},
					},
				},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "content",
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

var documentTestShouldNotIncludeRevisionHistoryInDocBookOutputIfRevdateIsNotSet = &asciidoc.Document{
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
					Name: "revnumber",
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "1.0",
						},
					},
				},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "content",
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

var documentTestWithMetadataToDocBook5 = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Stuart Rackham <",
				},
				asciidoc.Email{
					Address: "founder@asciidoc.org",
				},
				&asciidoc.String{
					Value: ">",
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
							Value: "more info...",
						},
						&asciidoc.NewLine{},
					},
					Title: asciidoc.Elements{
						&asciidoc.String{
							Value: "Version 8.6.8",
						},
					},
					Level: 1,
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "AsciiDoc",
				},
			},
			Level: 0,
		},
	},
}

var documentTestWithDocumentIdToDocbook5 = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.AnchorAttribute{
					ID: asciidoc.Elements{
						&asciidoc.String{
							Value: "document-id",
						},
					},
					Label: nil,
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "more info...",
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

var documentTestWithAuthorDefinedUsingAttributeEntryToDocBook = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.AttributeEntry{
					Name: "author",
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "Doc Writer",
						},
					},
				},
				&asciidoc.AttributeEntry{
					Name: "email",
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "thedoctor@asciidoc.org",
						},
					},
				},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "content",
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

var documentTestShouldSubstituteReplacementsInAuthorNamesInHtmlOutput = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Stephen O'Grady <",
				},
				asciidoc.Email{
					Address: "founder@redmonk.com",
				},
				&asciidoc.String{
					Value: ">",
				},
				&asciidoc.NewLine{},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "content",
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

var documentTestShouldSubstituteReplacementsInAuthorNamesInDocBookOutput = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Stephen O'Grady <",
				},
				asciidoc.Email{
					Address: "founder@redmonk.com",
				},
				&asciidoc.String{
					Value: ">",
				},
				&asciidoc.NewLine{},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "content",
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

var documentTestShouldSanitizeContentOfHtmlMetaAuthorsTag = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.AttributeEntry{
					Name: "author",
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "pass:n[http://example.org/community/team.html[Ze *Product* team]]",
						},
					},
				},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "content",
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

var documentTestShouldNotDoubleEscapeAmpersandInAuthorAttribute = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "R&D Lab",
				},
				&asciidoc.NewLine{},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.UserAttributeReference{
					Value: "author",
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

var documentTestShouldIncludeMultipleAuthorsInHtmlOutput = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Doc Writer <",
				},
				asciidoc.Email{
					Address: "thedoctor@asciidoc.org",
				},
				&asciidoc.String{
					Value: ">; Junior Writer <",
				},
				asciidoc.Email{
					Address: "junior@asciidoctor.org",
				},
				&asciidoc.String{
					Value: ">",
				},
				&asciidoc.NewLine{},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "content",
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

var documentTestShouldCreateAuthorgroupInDocBookWhenMultipleAuthors = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Doc Writer <",
				},
				asciidoc.Email{
					Address: "thedoctor@asciidoc.org",
				},
				&asciidoc.String{
					Value: ">; Junior Writer <",
				},
				asciidoc.Email{
					Address: "junior@asciidoctor.org",
				},
				&asciidoc.String{
					Value: ">",
				},
				&asciidoc.NewLine{},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "content",
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

var documentTestShouldProcessAuthorDefinedByAttributeWhenImplicitDoctitleIsAbsent = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "author",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Doc Writer",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UserAttributeReference{
			Value: "lastname",
		},
		&asciidoc.String{
			Value: ", ",
		},
		&asciidoc.UserAttributeReference{
			Value: "firstname",
		},
		&asciidoc.String{
			Value: " (",
		},
		&asciidoc.UserAttributeReference{
			Value: "authorinitials",
		},
		&asciidoc.String{
			Value: ")",
		},
		&asciidoc.NewLine{},
	},
}

var documentTestShouldProcessAuthorAndAuthorinitialsDefinedByAttributeWhenImplicitDoctitleIsAbsent = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "authorinitials",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "DOC",
				},
			},
		},
		&asciidoc.AttributeEntry{
			Name: "author",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Doc Writer",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UserAttributeReference{
			Value: "lastname",
		},
		&asciidoc.String{
			Value: ", ",
		},
		&asciidoc.UserAttributeReference{
			Value: "firstname",
		},
		&asciidoc.String{
			Value: " (",
		},
		&asciidoc.UserAttributeReference{
			Value: "authorinitials",
		},
		&asciidoc.String{
			Value: ")",
		},
		&asciidoc.NewLine{},
	},
}

var documentTestShouldProcessAuthorsDefinedByAttributeWhenImplicitDoctitleIsAbsent = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "authors",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Doc Writer; Other Author",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UserAttributeReference{
			Value: "lastname",
		},
		&asciidoc.String{
			Value: ", ",
		},
		&asciidoc.UserAttributeReference{
			Value: "firstname",
		},
		&asciidoc.String{
			Value: " (",
		},
		&asciidoc.UserAttributeReference{
			Value: "authorinitials",
		},
		&asciidoc.String{
			Value: ")",
		},
		&asciidoc.NewLine{},
	},
}

var documentTestShouldProcessAuthorsAndAuthorinitialsDefinedByAttributeWhenImplicitDoctitleIsAbsent = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "authorinitials",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "DOC",
				},
			},
		},
		&asciidoc.AttributeEntry{
			Name: "authors",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Doc Writer; Other Author",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UserAttributeReference{
			Value: "lastname",
		},
		&asciidoc.String{
			Value: ", ",
		},
		&asciidoc.UserAttributeReference{
			Value: "firstname",
		},
		&asciidoc.String{
			Value: " (",
		},
		&asciidoc.UserAttributeReference{
			Value: "authorinitials",
		},
		&asciidoc.String{
			Value: ")",
		},
		&asciidoc.NewLine{},
	},
}

var documentTestShouldSetAuthorcountTo0IfDocumentHasNoHeader = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name:     "idprefix",
			Elements: nil,
		},
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
					Value: "content",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Section Title",
				},
			},
			Level: 1,
		},
	},
}

var documentTestShouldSetAuthorcountTo0IfAuthorNotSetByAttributeAndDocumentStartsWithLevel0SectionWithStyle = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "doctype",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "book",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "preface",
							},
						},
					},
					ID:      nil,
					Roles:   nil,
					Options: nil,
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "content",
				},
				&asciidoc.NewLine{},
				&asciidoc.EmptyLine{
					Text: "",
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Preface",
				},
			},
			Level: 0,
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
							Value: "content",
						},
						&asciidoc.NewLine{},
					},
					Title: asciidoc.Elements{
						&asciidoc.String{
							Value: "Chapter",
						},
					},
					Level: 1,
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Part",
				},
			},
			Level: 0,
		},
	},
}

var documentTestWithAuthorDefinedByIndexedAttributeName = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.AttributeEntry{
					Name: "author_1",
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "Doc Writer",
						},
					},
				},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.UserAttributeReference{
					Value: "author",
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

var documentTestWithAuthorsDefinedUsingAttributeEntryToDocBook = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.AttributeEntry{
					Name: "authors",
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "Doc Writer; Junior Writer",
						},
					},
				},
				&asciidoc.AttributeEntry{
					Name: "email_1",
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "thedoctor@asciidoc.org",
						},
					},
				},
				&asciidoc.AttributeEntry{
					Name: "email_2",
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "junior@asciidoc.org",
						},
					},
				},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "content",
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

var documentTestShouldPopulateCopyrightElementInDocBookOutputIfCopyrightAttributeIsDefined = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.AttributeEntry{
					Name: "copyright",
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "ACME, Inc.",
						},
					},
				},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "Essential for catching road runners.",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Jet Bike",
				},
			},
			Level: 0,
		},
	},
}

var documentTestShouldPopulateCopyrightElementInDocBookOutputIfCopyrightAttributeIsDefinedWithYear = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.AttributeEntry{
					Name: "copyright",
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "ACME, Inc. 1956",
						},
					},
				},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "Essential for catching road runners.",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Jet Bike",
				},
			},
			Level: 0,
		},
	},
}

var documentTestShouldPopulateCopyrightElementInDocBookOutputIfCopyrightAttributeIsDefinedWithYearRange = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.AttributeEntry{
					Name: "copyright",
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "ACME, Inc. 1956-2018",
						},
					},
				},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "Essential for catching road runners.",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Jet Bike",
				},
			},
			Level: 0,
		},
	},
}

var documentTestWithHeaderFooter = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name:     "nofooter",
			Elements: nil,
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "content",
		},
		&asciidoc.NewLine{},
	},
}

var documentTestCanDisableLastUpdatedInFooter = &asciidoc.Document{
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
					Value: "content",
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

var documentTestParseHeaderOnly = &asciidoc.Document{
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
					Name: "foo",
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "bar",
						},
					},
				},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "preamble",
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

var documentTestShouldParseHeaderOnlyWhenDocytpeIsManpage = &asciidoc.Document{
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
							Value: "cmd - does stuff",
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
					Value: "cmd(1)",
				},
			},
			Level: 0,
		},
	},
}

var documentTestShouldNotWarnWhenParsingHeaderOnlyWhenDocytpeIsManpageAndBodyIsEmpty = &asciidoc.Document{
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
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "cmd(1)",
				},
			},
			Level: 0,
		},
	},
}

var documentTestOutputsFootnotesInFooter = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "A footnote footnote:[An example footnote.];",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "a second footnote with a reference ID footnote:note2[Second footnote.];",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "and finally a reference to the second footnote footnote:note2[].",
		},
		&asciidoc.NewLine{},
	},
}

var documentTestOutputsFootnotesBlockInEmbeddedDocumentByDefault = &asciidoc.Document{
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
							Value: "Content",
						},
						&asciidoc.NewLine{},
						&asciidoc.EmptyLine{
							Text: "",
						},
					},
					Title: asciidoc.Elements{
						&asciidoc.String{
							Value: "Section A",
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
							Value: "Content.footnote:[commentary]",
						},
						&asciidoc.NewLine{},
					},
					Title: asciidoc.Elements{
						&asciidoc.String{
							Value: "Section B",
						},
					},
					Level: 1,
				},
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

var documentTestShouldReturnEmptyidsTable = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.BlockImage{
			AttributeList: nil,
			ImagePath: asciidoc.Elements{
				&asciidoc.String{
					Value: "outer.png",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Table{
			AttributeList: nil,
			ColumnCount:   1,
			Elements: asciidoc.Elements{
				&asciidoc.TableRow{
					Elements: asciidoc.Elements{
						&asciidoc.TableCell{
							Format: &asciidoc.TableCellFormat{
								Multiplier: asciidoc.Optional[int]{
									Value: 1,
									IsSet: false,
								},
								Span: asciidoc.TableCellSpan{
									Column: asciidoc.Optional[int]{
										Value: 1,
										IsSet: false,
									},
									Row: asciidoc.Optional[int]{
										Value: 1,
										IsSet: false,
									},
								},
								HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
									Value: 0,
									IsSet: false,
								},
								VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
									Value: 0,
									IsSet: false,
								},
								Style: asciidoc.Optional[asciidoc.TableCellStyle]{
									Value: 1,
									IsSet: true,
								},
							},
							Elements: asciidoc.Elements{
								&asciidoc.NewLine{},
								&asciidoc.BlockImage{
									AttributeList: nil,
									ImagePath: asciidoc.Elements{
										&asciidoc.String{
											Value: "inner.png",
										},
									},
								},
							},
							Blank: false,
						},
					},
				},
			},
		},
	},
}

var documentTestHonorHtmlsyntaxAttributeInDocumentHeaderIfFollowedByBackendAttribute = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "htmlsyntax",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "xml",
				},
			},
		},
		&asciidoc.AttributeEntry{
			Name: "backend",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "html5",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.ThematicBreak{
			AttributeList: nil,
		},
	},
}

var documentTestDoesNotHonorHtmlsyntaxAttributeInDocumentHeaderIfNotFollowedByBackendAttribute = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "backend",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "html5",
				},
			},
		},
		&asciidoc.AttributeEntry{
			Name: "htmlsyntax",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "xml",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.ThematicBreak{
			AttributeList: nil,
		},
	},
}

var documentTestShouldCloseAllShortTagsWhenHtmlsyntaxIsXml = &asciidoc.Document{
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
				&asciidoc.String{
					Value: "v1.0, 2001-01-01",
				},
				&asciidoc.NewLine{},
				&asciidoc.AttributeEntry{
					Name:     "icons",
					Elements: nil,
				},
				&asciidoc.AttributeEntry{
					Name:     "favicon",
					Elements: nil,
				},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.InlineImage{
					AttributeList: nil,
					ImagePath: asciidoc.Elements{
						&asciidoc.String{
							Value: "tiger.png",
						},
					},
				},
				&asciidoc.NewLine{},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.BlockImage{
					AttributeList: nil,
					ImagePath: asciidoc.Elements{
						&asciidoc.String{
							Value: "tiger.png",
						},
					},
				},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.UnorderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "one",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        "*",
					Checklist:     2,
				},
				&asciidoc.UnorderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "two",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        "*",
					Checklist:     1,
				},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.Table{
					AttributeList: nil,
					ColumnCount:   2,
					Elements: asciidoc.Elements{
						&asciidoc.TableRow{
							Elements: asciidoc.Elements{
								&asciidoc.TableCell{
									Format: &asciidoc.TableCellFormat{
										Multiplier: asciidoc.Optional[int]{
											Value: 1,
											IsSet: false,
										},
										Span: asciidoc.TableCellSpan{
											Column: asciidoc.Optional[int]{
												Value: 1,
												IsSet: false,
											},
											Row: asciidoc.Optional[int]{
												Value: 1,
												IsSet: false,
											},
										},
										HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
											Value: 0,
											IsSet: false,
										},
										VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
											Value: 0,
											IsSet: false,
										},
										Style: asciidoc.Optional[asciidoc.TableCellStyle]{
											Value: 0,
											IsSet: false,
										},
									},
									Elements: asciidoc.Elements{
										&asciidoc.String{
											Value: "A",
										},
									},
									Blank: false,
								},
								&asciidoc.TableCell{
									Format: &asciidoc.TableCellFormat{
										Multiplier: asciidoc.Optional[int]{
											Value: 1,
											IsSet: false,
										},
										Span: asciidoc.TableCellSpan{
											Column: asciidoc.Optional[int]{
												Value: 1,
												IsSet: false,
											},
											Row: asciidoc.Optional[int]{
												Value: 1,
												IsSet: false,
											},
										},
										HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
											Value: 0,
											IsSet: false,
										},
										VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
											Value: 0,
											IsSet: false,
										},
										Style: asciidoc.Optional[asciidoc.TableCellStyle]{
											Value: 0,
											IsSet: false,
										},
									},
									Elements: asciidoc.Elements{
										&asciidoc.String{
											Value: "B",
										},
									},
									Blank: false,
								},
							},
						},
					},
				},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.Paragraph{
					AttributeList: asciidoc.AttributeList{
						&asciidoc.ShorthandAttribute{
							Style: &asciidoc.ShorthandStyle{
								Elements: asciidoc.Elements{
									&asciidoc.String{
										Value: "horizontal",
									},
								},
							},
							ID:      nil,
							Roles:   nil,
							Options: nil,
						},
						&asciidoc.NamedAttribute{
							Name: "labelwidth",
							Val: asciidoc.Elements{
								&asciidoc.String{
									Value: "25%",
								},
							},
							Quote: 2,
						},
						&asciidoc.NamedAttribute{
							Name: "itemwidth",
							Val: asciidoc.Elements{
								&asciidoc.String{
									Value: "75%",
								},
							},
							Quote: 2,
						},
					},
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "term:: description",
						},
						&asciidoc.NewLine{},
					},
					Admonition: 0,
				},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.Paragraph{
					AttributeList: nil,
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "note",
						},
						&asciidoc.NewLine{},
					},
					Admonition: 1,
				},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.QuoteBlock{
					Delimiter: asciidoc.Delimiter{
						Type:   11,
						Length: 4,
					},
					AttributeList: asciidoc.AttributeList{
						&asciidoc.ShorthandAttribute{
							Style: &asciidoc.ShorthandStyle{
								Elements: asciidoc.Elements{
									&asciidoc.String{
										Value: "quote",
									},
								},
							},
							ID:      nil,
							Roles:   nil,
							Options: nil,
						},
						&asciidoc.PositionalAttribute{
							Offset:      1,
							ImpliedName: "",
							Val: asciidoc.Elements{
								&asciidoc.String{
									Value: "Author",
								},
							},
						},
						&asciidoc.PositionalAttribute{
							Offset:      2,
							ImpliedName: "",
							Val: asciidoc.Elements{
								&asciidoc.String{
									Value: "Source",
								},
							},
						},
					},
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "Quote me.",
						},
						&asciidoc.NewLine{},
					},
				},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.QuoteBlock{
					Delimiter: asciidoc.Delimiter{
						Type:   11,
						Length: 4,
					},
					AttributeList: asciidoc.AttributeList{
						&asciidoc.ShorthandAttribute{
							Style: &asciidoc.ShorthandStyle{
								Elements: asciidoc.Elements{
									&asciidoc.String{
										Value: "verse",
									},
								},
							},
							ID:      nil,
							Roles:   nil,
							Options: nil,
						},
						&asciidoc.PositionalAttribute{
							Offset:      1,
							ImpliedName: "",
							Val: asciidoc.Elements{
								&asciidoc.String{
									Value: "Author",
								},
							},
						},
						&asciidoc.PositionalAttribute{
							Offset:      2,
							ImpliedName: "",
							Val: asciidoc.Elements{
								&asciidoc.String{
									Value: "Source",
								},
							},
						},
					},
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "A tall tale.",
						},
						&asciidoc.NewLine{},
					},
				},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.Paragraph{
					AttributeList: asciidoc.AttributeList{
						&asciidoc.NamedAttribute{
							Name: "options",
							Val: asciidoc.Elements{
								&asciidoc.String{
									Value: "autoplay,loop",
								},
							},
							Quote: 2,
						},
					},
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "video::screencast.ogg[]",
						},
						&asciidoc.NewLine{},
					},
					Admonition: 0,
				},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "video::12345[vimeo]",
				},
				&asciidoc.NewLine{},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.Paragraph{
					AttributeList: asciidoc.AttributeList{
						&asciidoc.NamedAttribute{
							Name: "options",
							Val: asciidoc.Elements{
								&asciidoc.String{
									Value: "autoplay,loop",
								},
							},
							Quote: 2,
						},
					},
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "audio::podcast.ogg[]",
						},
						&asciidoc.NewLine{},
					},
					Admonition: 0,
				},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "one ",
				},
				&asciidoc.LineBreak{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "two",
				},
				&asciidoc.NewLine{},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.ThematicBreak{
					AttributeList: nil,
				},
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

var documentTestXhtmlBackendShouldEmitElementsInProperNamespace = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.AttributeEntry{
					Name: "doctype",
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "book",
						},
					},
				},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "text",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Document Title: Subtitle",
				},
			},
			Level: 0,
		},
	},
}

var documentTestShouldBeAbleToSetDoctypeToArticleWhenConvertingToDocBook = &asciidoc.Document{
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
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "preamble",
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
							Value: "section body",
						},
						&asciidoc.NewLine{},
					},
					Title: asciidoc.Elements{
						&asciidoc.String{
							Value: "First Section",
						},
					},
					Level: 1,
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Title",
				},
			},
			Level: 0,
		},
	},
}

var documentTestShouldSetDoctypeToArticleByDefaultForDocumentWithNoTitleWhenConvertingToDocBook = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.AttributeEntry{
					Name: "mansource",
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "Asciidoctor",
						},
					},
				},
				&asciidoc.AttributeEntry{
					Name: "manmanual",
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "Asciidoctor Manual",
						},
					},
				},
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
							Value: "asciidoctor - Process text",
						},
						&asciidoc.NewLine{},
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
						&asciidoc.String{
							Value: "some text",
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
							Value: "section body",
						},
						&asciidoc.NewLine{},
					},
					Title: asciidoc.Elements{
						&asciidoc.String{
							Value: "First Section",
						},
					},
					Level: 1,
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "asciidoctor(1)",
				},
			},
			Level: 0,
		},
	},
}

var documentTestShouldOutputNonBreakingSpaceForSourceAndManualInDocbookManpageOutputIfAbsentFromSource = &asciidoc.Document{
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
							Value: "asciidoctor - Process text",
						},
						&asciidoc.NewLine{},
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
						&asciidoc.String{
							Value: "some text",
						},
						&asciidoc.NewLine{},
					},
					Title: asciidoc.Elements{
						&asciidoc.String{
							Value: "SYNOPSIS",
						},
					},
					Level: 1,
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "asciidoctor(1)",
				},
			},
			Level: 0,
		},
	},
}

var documentTestShouldApplyReplacementsSubstitutionToValueOfMantitleAttributeUsedInDocBookOutput = &asciidoc.Document{
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
							Value: "foo--bar - puts the foo in your bar",
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
					Value: "foo\\--bar(1)",
				},
			},
			Level: 0,
		},
	},
}

var documentTestShouldBeAbleToSetDoctypeToBookWhenConvertingToDocBook = &asciidoc.Document{
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
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "preamble",
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
							Value: "chapter body",
						},
						&asciidoc.NewLine{},
					},
					Title: asciidoc.Elements{
						&asciidoc.String{
							Value: "First Chapter",
						},
					},
					Level: 1,
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Title",
				},
			},
			Level: 0,
		},
	},
}

var documentTestShouldBeAbleToSetDoctypeToBookForDocumentWithNoTitleWhenConvertingToDocBook = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Andrew Stanton",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "v1.0.0",
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
				&asciidoc.AttributeEntry{
					Name: "manmanual",
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "EVE",
						},
					},
				},
				&asciidoc.AttributeEntry{
					Name: "mansource",
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "EVE",
						},
					},
				},
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
							Value: "eve, islifeform - analyzes an image to determine if it's a picture of a life form",
						},
						&asciidoc.NewLine{},
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
									Value: "eve",
								},
							},
						},
						&asciidoc.String{
							Value: " ['OPTION']... 'FILE'...",
						},
						&asciidoc.NewLine{},
					},
					Title: asciidoc.Elements{
						&asciidoc.String{
							Value: "SYNOPSIS",
						},
					},
					Level: 1,
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "eve(1)",
				},
			},
			Level: 0,
		},
	},
}

var documentTestAddsAFrontAndBackCoverImageToDocBook5WhenDoctypeIsBook = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.AttributeEntry{
					Name: "doctype",
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "book",
						},
					},
				},
				&asciidoc.AttributeEntry{
					Name: "imagesdir",
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "images",
						},
					},
				},
				&asciidoc.AttributeEntry{
					Name: "front-cover-image",
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "image:front-cover.jpg[scaledwidth=210mm]",
						},
					},
				},
				&asciidoc.AttributeEntry{
					Name: "back-cover-image",
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "image:back-cover.jpg[]",
						},
					},
				},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "preamble",
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
							Value: "chapter body",
						},
						&asciidoc.NewLine{},
					},
					Title: asciidoc.Elements{
						&asciidoc.String{
							Value: "First Chapter",
						},
					},
					Level: 1,
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Title",
				},
			},
			Level: 0,
		},
	},
}

var documentTestShouldBeAbleToSetBackendUsingbackendOptionKey = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Stuart Rackham <",
				},
				asciidoc.Email{
					Address: "founder@asciidoc.org",
				},
				&asciidoc.String{
					Value: ">",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: ":Author Initials: SJR",
				},
				&asciidoc.NewLine{},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "more info...",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "AsciiDoc",
				},
			},
			Level: 0,
		},
	},
}

var documentTestAttributeEntryCanAppearImmediatelyAfterDocumentTitle = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "Reference Guide",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "===============",
		},
		&asciidoc.NewLine{},
		&asciidoc.AttributeEntry{
			Name:     "toc",
			Elements: nil,
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "preamble",
		},
		&asciidoc.NewLine{},
	},
}

var documentTestAttributeEntryCanAppearBeforeAuthorLineUnderDocumentTitle = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "Reference Guide",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "===============",
		},
		&asciidoc.NewLine{},
		&asciidoc.AttributeEntry{
			Name:     "toc",
			Elements: nil,
		},
		&asciidoc.String{
			Value: "Dan Allen",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "preamble",
		},
		&asciidoc.NewLine{},
	},
}

var documentTestShouldParseMantitleAndManvolnumFromDocumentTitleForManpageDoctype = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.AttributeEntry{
					Name: "doctype",
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "manpage",
						},
					},
				},
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
							Value: "asciidoctor - converts AsciiDoc source files to HTML, DocBook and other formats",
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
					Value: "asciidoctor ( 1 )",
				},
			},
			Level: 0,
		},
	},
}

var documentTestShouldPerformAttributeSubstitutionOnMantitleInManpageDoctype = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.AttributeEntry{
					Name: "doctype",
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "manpage",
						},
					},
				},
				&asciidoc.AttributeEntry{
					Name: "app",
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "Asciidoctor",
						},
					},
				},
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
							Value: "asciidoctor - converts AsciiDoc source files to HTML, DocBook and other formats",
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
					Value: "{app}(1)",
				},
			},
			Level: 0,
		},
	},
}

var documentTestShouldConsumeNameSectionAsMannameAndManpurposeForManpageDoctype = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.AttributeEntry{
					Name: "doctype",
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "manpage",
						},
					},
				},
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
							Value: "asciidoctor - converts AsciiDoc source files to HTML, DocBook and other formats",
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
					Value: "asciidoctor(1)",
				},
			},
			Level: 0,
		},
	},
}

var documentTestShouldSetDocnameAndOutfilesuffixFromMannameAndManvolnumForManpageBackendAndDoctype = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.AttributeEntry{
					Name: "doctype",
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "manpage",
						},
					},
				},
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
							Value: "asciidoctor - converts AsciiDoc source files to HTML, DocBook and other formats",
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
					Value: "asciidoctor(1)",
				},
			},
			Level: 0,
		},
	},
}

var documentTestShouldMarkSynopsisAsSpecialSectionInManpageDoctype = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.AttributeEntry{
					Name: "doctype",
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "manpage",
						},
					},
				},
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
							Value: "asciidoctor - converts AsciiDoc source files to HTML, DocBook and other formats",
						},
						&asciidoc.NewLine{},
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
									Value: "asciidoctor",
								},
							},
						},
						&asciidoc.String{
							Value: " ['OPTION']... 'FILE'..",
						},
						&asciidoc.NewLine{},
					},
					Title: asciidoc.Elements{
						&asciidoc.String{
							Value: "SYNOPSIS",
						},
					},
					Level: 1,
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "asciidoctor(1)",
				},
			},
			Level: 0,
		},
	},
}

var documentTestShouldOutputSpecialHeaderBlockInHtmlForManpageDoctype = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.AttributeEntry{
					Name: "doctype",
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "manpage",
						},
					},
				},
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
							Value: "asciidoctor - converts AsciiDoc source files to HTML, DocBook and other formats",
						},
						&asciidoc.NewLine{},
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
									Value: "asciidoctor",
								},
							},
						},
						&asciidoc.String{
							Value: " ['OPTION']... 'FILE'..",
						},
						&asciidoc.NewLine{},
					},
					Title: asciidoc.Elements{
						&asciidoc.String{
							Value: "SYNOPSIS",
						},
					},
					Level: 1,
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "asciidoctor(1)",
				},
			},
			Level: 0,
		},
	},
}

var documentTestShouldOutputSpecialHeaderBlockInEmbeddableHtmlForManpageDoctype = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.AttributeEntry{
					Name: "doctype",
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "manpage",
						},
					},
				},
				&asciidoc.AttributeEntry{
					Name:     "showtitle",
					Elements: nil,
				},
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
							Value: "asciidoctor - converts AsciiDoc source files to HTML, DocBook and other formats",
						},
						&asciidoc.NewLine{},
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
									Value: "asciidoctor",
								},
							},
						},
						&asciidoc.String{
							Value: " ['OPTION']... 'FILE'..",
						},
						&asciidoc.NewLine{},
					},
					Title: asciidoc.Elements{
						&asciidoc.String{
							Value: "SYNOPSIS",
						},
					},
					Level: 1,
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "asciidoctor(1)",
				},
			},
			Level: 0,
		},
	},
}

var documentTestShouldOutputAllMannamesInNameSectionInManPageOutput = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.AttributeEntry{
					Name: "doctype",
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "manpage",
						},
					},
				},
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
							Value: "eve, probe - analyzes an image to determine if it is a picture of a life form",
						},
						&asciidoc.NewLine{},
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
									Value: "eve",
								},
							},
						},
						&asciidoc.String{
							Value: " [OPTION]... FILE...",
						},
						&asciidoc.NewLine{},
					},
					Title: asciidoc.Elements{
						&asciidoc.String{
							Value: "SYNOPSIS",
						},
					},
					Level: 1,
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "eve(1)",
				},
			},
			Level: 0,
		},
	},
}

var documentTestAllowsUsToSpecifyAPathRelativeToTheCurrentDir = &asciidoc.Document{
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
					Value: "text",
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

var documentTestShouldRaiseAnExceptionWhenAConverterCannotBeResolvedWhileParsing = &asciidoc.Document{
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
							Value: "text",
						},
						&asciidoc.NewLine{},
					},
					Title: asciidoc.Elements{
						&asciidoc.String{
							Value: "A _Big_ Section",
						},
					},
					Level: 1,
				},
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
