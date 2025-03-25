package tests

import (
	"testing"

	"github.com/project-chip/alchemy/asciidoc"
)

func TestDocument(t *testing.T) {
	documentTests.run(t)
}

var documentTests = parseTests{

	{"should be able to disable toc and sectnums in document header in DocBook backend", "asciidoctor/document_test_should_be_able_to_disable_toc_and_sectnums_in_document_header_in_doc_book_backend.adoc", shouldBeAbleToDisableTocAndSectnumsInDocumentHeaderInDocBookBackend, nil},

	{"noheader attribute should suppress info element when converting to DocBook", "asciidoctor/document_test_noheader_attribute_should_suppress_info_element_when_converting_to_doc_book.adoc", noheaderAttributeShouldSuppressInfoElementWhenConvertingToDocBook, nil},

	{"should be able to disable section numbering using numbered attribute in document header in DocBook backend", "asciidoctor/document_test_should_be_able_to_disable_section_numbering_using_numbered_attribute_in_document_header_in_doc_book_backend.adoc", shouldBeAbleToDisableSectionNumberingUsingNumberedAttributeInDocumentHeaderInDocBookBackend, nil},

	{"convert methods on built-in converter are registered by default", "asciidoctor/document_test_convert_methods_on_built_in_converter_are_registered_by_default.adoc", convertMethodsOnBuiltInConverterAreRegisteredByDefault, nil},

	{"should not enable compat mode for document with legacy doctitle if compat mode disable by header", "asciidoctor/document_test_should_not_enable_compat_mode_for_document_with_legacy_doctitle_if_compat_mode_disable_by_header.adoc", shouldNotEnableCompatModeForDocumentWithLegacyDoctitleIfCompatModeDisableByHeader, nil},

	{"should not enable compat mode for document with legacy doctitle if compat mode is locked by API", "asciidoctor/document_test_should_not_enable_compat_mode_for_document_with_legacy_doctitle_if_compat_mode_is_locked_by_api.adoc", shouldNotEnableCompatModeForDocumentWithLegacyDoctitleIfCompatModeIsLockedByApi, nil},

	{"should apply max-width to each top-level container", "asciidoctor/document_test_should_apply_max_width_to_each_top_level_container.adoc", shouldApplyMaxWidthToEachTopLevelContainer, nil},

	{"title partition API with default separator", "asciidoctor/document_test_title_partition_api_with_default_separator.adoc", titlePartitionApiWithDefaultSeparator, nil},

	{"document with subtitle and custom separator", "asciidoctor/document_test_document_with_subtitle_and_custom_separator.adoc", documentWithSubtitleAndCustomSeparator, nil},

	{"should not honor custom separator for doctitle if attribute is locked by API", "asciidoctor/document_test_should_not_honor_custom_separator_for_doctitle_if_attribute_is_locked_by_api.adoc", shouldNotHonorCustomSeparatorForDoctitleIfAttributeIsLockedByApi, nil},

	{"document with doctitle defined as attribute entry", "asciidoctor/document_test_document_with_doctitle_defined_as_attribute_entry.adoc", documentWithDoctitleDefinedAsAttributeEntry, nil},

	{"document with doctitle defined as attribute entry followed by block with title", "asciidoctor/document_test_document_with_doctitle_defined_as_attribute_entry_followed_by_block_with_title.adoc", documentWithDoctitleDefinedAsAttributeEntryFollowedByBlockWithTitle, nil},

	{"document with title attribute entry overrides doctitle", "asciidoctor/document_test_document_with_title_attribute_entry_overrides_doctitle.adoc", documentWithTitleAttributeEntryOverridesDoctitle, nil},

	{"document with blank title attribute entry overrides doctitle", "asciidoctor/document_test_document_with_blank_title_attribute_entry_overrides_doctitle.adoc", documentWithBlankTitleAttributeEntryOverridesDoctitle, nil},

	{"document header can reference intrinsic doctitle attribute", "asciidoctor/document_test_document_header_can_reference_intrinsic_doctitle_attribute.adoc", documentHeaderCanReferenceIntrinsicDoctitleAttribute, nil},

	{"document with title attribute entry overrides doctitle attribute entry", "asciidoctor/document_test_document_with_title_attribute_entry_overrides_doctitle_attribute_entry.adoc", documentWithTitleAttributeEntryOverridesDoctitleAttributeEntry, nil},

	{"document with doctitle attribute entry overrides implicit doctitle", "asciidoctor/document_test_document_with_doctitle_attribute_entry_overrides_implicit_doctitle.adoc", documentWithDoctitleAttributeEntryOverridesImplicitDoctitle, nil},

	{"doctitle attribute entry above header overrides implicit doctitle", "asciidoctor/document_test_doctitle_attribute_entry_above_header_overrides_implicit_doctitle.adoc", doctitleAttributeEntryAboveHeaderOverridesImplicitDoctitle, nil},

	{"should apply header substitutions to value of the doctitle attribute assigned from implicit doctitle", "asciidoctor/document_test_should_apply_header_substitutions_to_value_of_the_doctitle_attribute_assigned_from_implicit_doctitle.adoc", shouldApplyHeaderSubstitutionsToValueOfTheDoctitleAttributeAssignedFromImplicitDoctitle, nil},

	{"should substitute attribute reference in implicit document title for attribute defined earlier in header", "asciidoctor/document_test_should_substitute_attribute_reference_in_implicit_document_title_for_attribute_defined_earlier_in_header.adoc", shouldSubstituteAttributeReferenceInImplicitDocumentTitleForAttributeDefinedEarlierInHeader, nil},

	{"should not warn if implicit document title contains attribute reference for attribute defined later in header", "asciidoctor/document_test_should_not_warn_if_implicit_document_title_contains_attribute_reference_for_attribute_defined_later_in_header.adoc", shouldNotWarnIfImplicitDocumentTitleContainsAttributeReferenceForAttributeDefinedLaterInHeader, nil},

	{"should recognize document title when preceded by blank lines", "asciidoctor/document_test_should_recognize_document_title_when_preceded_by_blank_lines.adoc", shouldRecognizeDocumentTitleWhenPrecededByBlankLines, nil},

	{"should recognize document title when preceded by blank lines introduced by a preprocessor conditional", "asciidoctor/document_test_should_recognize_document_title_when_preceded_by_blank_lines_introduced_by_a_preprocessor_conditional.adoc", shouldRecognizeDocumentTitleWhenPrecededByBlankLinesIntroducedByAPreprocessorConditional, nil},

	{"should recognize document title when preceded by blank lines after an attribute entry", "asciidoctor/document_test_should_recognize_document_title_when_preceded_by_blank_lines_after_an_attribute_entry.adoc", shouldRecognizeDocumentTitleWhenPrecededByBlankLinesAfterAnAttributeEntry, nil},

	{"should recognize document title in include file when preceded by blank lines", "asciidoctor/document_test_should_recognize_document_title_in_include_file_when_preceded_by_blank_lines.adoc", shouldRecognizeDocumentTitleInIncludeFileWhenPrecededByBlankLines, nil},

	{"should include specified lines even when leading lines are skipped", "asciidoctor/document_test_should_include_specified_lines_even_when_leading_lines_are_skipped.adoc", shouldIncludeSpecifiedLinesEvenWhenLeadingLinesAreSkipped, nil},

	{"document with multiline attribute entry but only one line should not crash", "asciidoctor/document_test_document_with_multiline_attribute_entry_but_only_one_line_should_not_crash.adoc", documentWithMultilineAttributeEntryButOnlyOneLineShouldNotCrash, nil},

	{"should not choke on empty source", "asciidoctor/document_test_should_not_choke_on_empty_source.adoc", shouldNotChokeOnEmptySource, nil},

	{"should parse revision line if date is empty", "asciidoctor/document_test_should_parse_revision_line_if_date_is_empty.adoc", shouldParseRevisionLineIfDateIsEmpty, nil},

	{"should include revision history in DocBook output if revdate and revnumber is set", "asciidoctor/document_test_should_include_revision_history_in_doc_book_output_if_revdate_and_revnumber_is_set.adoc", shouldIncludeRevisionHistoryInDocBookOutputIfRevdateAndRevnumberIsSet, nil},

	{"should include revision history in DocBook output if revdate and revremark is set", "asciidoctor/document_test_should_include_revision_history_in_doc_book_output_if_revdate_and_revremark_is_set.adoc", shouldIncludeRevisionHistoryInDocBookOutputIfRevdateAndRevremarkIsSet, nil},

	{"should not include revision history in DocBook output if revdate is not set", "asciidoctor/document_test_should_not_include_revision_history_in_doc_book_output_if_revdate_is_not_set.adoc", shouldNotIncludeRevisionHistoryInDocBookOutputIfRevdateIsNotSet, nil},

	{"with metadata to DocBook 5", "asciidoctor/document_test_with_metadata_to_doc_book_5.adoc", withMetadataToDocBook5, nil},

	{"with document ID to Docbook 5", "asciidoctor/document_test_with_document_id_to_docbook_5.adoc", withDocumentIdToDocbook5, nil},

	{"with author defined using attribute entry to DocBook", "asciidoctor/document_test_with_author_defined_using_attribute_entry_to_doc_book.adoc", withAuthorDefinedUsingAttributeEntryToDocBook, nil},

	{"should substitute replacements in author names in HTML output", "asciidoctor/document_test_should_substitute_replacements_in_author_names_in_html_output.adoc", shouldSubstituteReplacementsInAuthorNamesInHtmlOutput, nil},

	{"should substitute replacements in author names in DocBook output", "asciidoctor/document_test_should_substitute_replacements_in_author_names_in_doc_book_output.adoc", shouldSubstituteReplacementsInAuthorNamesInDocBookOutput, nil},

	{"should sanitize content of HTML meta authors tag", "asciidoctor/document_test_should_sanitize_content_of_html_meta_authors_tag.adoc", shouldSanitizeContentOfHtmlMetaAuthorsTag, nil},

	{"should not double escape ampersand in author attribute", "asciidoctor/document_test_should_not_double_escape_ampersand_in_author_attribute.adoc", shouldNotDoubleEscapeAmpersandInAuthorAttribute, nil},

	{"should include multiple authors in HTML output", "asciidoctor/document_test_should_include_multiple_authors_in_html_output.adoc", shouldIncludeMultipleAuthorsInHtmlOutput, nil},

	{"should create authorgroup in DocBook when multiple authors", "asciidoctor/document_test_should_create_authorgroup_in_doc_book_when_multiple_authors.adoc", shouldCreateAuthorgroupInDocBookWhenMultipleAuthors, nil},

	{"should process author defined by attribute when implicit doctitle is absent", "asciidoctor/document_test_should_process_author_defined_by_attribute_when_implicit_doctitle_is_absent.adoc", shouldProcessAuthorDefinedByAttributeWhenImplicitDoctitleIsAbsent, nil},

	{"should process author and authorinitials defined by attribute when implicit doctitle is absent", "asciidoctor/document_test_should_process_author_and_authorinitials_defined_by_attribute_when_implicit_doctitle_is_absent.adoc", shouldProcessAuthorAndAuthorinitialsDefinedByAttributeWhenImplicitDoctitleIsAbsent, nil},

	{"should process authors defined by attribute when implicit doctitle is absent", "asciidoctor/document_test_should_process_authors_defined_by_attribute_when_implicit_doctitle_is_absent.adoc", shouldProcessAuthorsDefinedByAttributeWhenImplicitDoctitleIsAbsent, nil},

	{"should process authors and authorinitials defined by attribute when implicit doctitle is absent", "asciidoctor/document_test_should_process_authors_and_authorinitials_defined_by_attribute_when_implicit_doctitle_is_absent.adoc", shouldProcessAuthorsAndAuthorinitialsDefinedByAttributeWhenImplicitDoctitleIsAbsent, nil},

	{"should set authorcount to 0 if document has no header", "asciidoctor/document_test_should_set_authorcount_to_0_if_document_has_no_header.adoc", shouldSetAuthorcountTo0IfDocumentHasNoHeader, nil},

	{"should set authorcount to 0 if author not set by attribute and document starts with level-0 section with style", "asciidoctor/document_test_should_set_authorcount_to_0_if_author_not_set_by_attribute_and_document_starts_with_level_0_section_with_style.adoc", shouldSetAuthorcountTo0IfAuthorNotSetByAttributeAndDocumentStartsWithLevel0SectionWithStyle, nil},

	{"with author defined by indexed attribute name", "asciidoctor/document_test_with_author_defined_by_indexed_attribute_name.adoc", withAuthorDefinedByIndexedAttributeName, nil},

	{"with authors defined using attribute entry to DocBook", "asciidoctor/document_test_with_authors_defined_using_attribute_entry_to_doc_book.adoc", withAuthorsDefinedUsingAttributeEntryToDocBook, nil},

	{"should populate copyright element in DocBook output if copyright attribute is defined", "asciidoctor/document_test_should_populate_copyright_element_in_doc_book_output_if_copyright_attribute_is_defined.adoc", shouldPopulateCopyrightElementInDocBookOutputIfCopyrightAttributeIsDefined, nil},

	{"should populate copyright element in DocBook output if copyright attribute is defined with year", "asciidoctor/document_test_should_populate_copyright_element_in_doc_book_output_if_copyright_attribute_is_defined_with_year.adoc", shouldPopulateCopyrightElementInDocBookOutputIfCopyrightAttributeIsDefinedWithYear, nil},

	{"should populate copyright element in DocBook output if copyright attribute is defined with year range", "asciidoctor/document_test_should_populate_copyright_element_in_doc_book_output_if_copyright_attribute_is_defined_with_year_range.adoc", shouldPopulateCopyrightElementInDocBookOutputIfCopyrightAttributeIsDefinedWithYearRange, nil},

	{"with header footer", "asciidoctor/document_test_with_header_footer.adoc", withHeaderFooter, nil},

	{"can disable last updated in footer", "asciidoctor/document_test_can_disable_last_updated_in_footer.adoc", canDisableLastUpdatedInFooter, nil},

	{"parse header only", "asciidoctor/document_test_parse_header_only.adoc", parseHeaderOnly, nil},

	{"should parse header only when docytpe is manpage", "asciidoctor/document_test_should_parse_header_only_when_docytpe_is_manpage.adoc", shouldParseHeaderOnlyWhenDocytpeIsManpage, nil},

	{"should not warn when parsing header only when docytpe is manpage and body is empty", "asciidoctor/document_test_should_not_warn_when_parsing_header_only_when_docytpe_is_manpage_and_body_is_empty.adoc", shouldNotWarnWhenParsingHeaderOnlyWhenDocytpeIsManpageAndBodyIsEmpty, nil},

	{"outputs footnotes in footer", "asciidoctor/document_test_outputs_footnotes_in_footer.adoc", outputsFootnotesInFooter, nil},

	{"outputs footnotes block in embedded document by default", "asciidoctor/document_test_outputs_footnotes_block_in_embedded_document_by_default.adoc", outputsFootnotesBlockInEmbeddedDocumentByDefault, nil},

	{"should return empty :ids table", "asciidoctor/document_test_should_return_empty_:ids_table.adoc", shouldReturnEmptyidsTable, nil},

	{"honor htmlsyntax attribute in document header if followed by backend attribute", "asciidoctor/document_test_honor_htmlsyntax_attribute_in_document_header_if_followed_by_backend_attribute.adoc", honorHtmlsyntaxAttributeInDocumentHeaderIfFollowedByBackendAttribute, nil},

	{"does not honor htmlsyntax attribute in document header if not followed by backend attribute", "asciidoctor/document_test_does_not_honor_htmlsyntax_attribute_in_document_header_if_not_followed_by_backend_attribute.adoc", doesNotHonorHtmlsyntaxAttributeInDocumentHeaderIfNotFollowedByBackendAttribute, nil},

	{"should close all short tags when htmlsyntax is xml", "asciidoctor/document_test_should_close_all_short_tags_when_htmlsyntax_is_xml.adoc", shouldCloseAllShortTagsWhenHtmlsyntaxIsXml, nil},

	{"xhtml backend should emit elements in proper namespace", "asciidoctor/document_test_xhtml_backend_should_emit_elements_in_proper_namespace.adoc", xhtmlBackendShouldEmitElementsInProperNamespace, nil},

	{"should be able to set doctype to article when converting to DocBook", "asciidoctor/document_test_should_be_able_to_set_doctype_to_article_when_converting_to_doc_book.adoc", shouldBeAbleToSetDoctypeToArticleWhenConvertingToDocBook, nil},

	{"should set doctype to article by default for document with no title when converting to DocBook", "asciidoctor/document_test_should_set_doctype_to_article_by_default_for_document_with_no_title_when_converting_to_doc_book.adoc", shouldSetDoctypeToArticleByDefaultForDocumentWithNoTitleWhenConvertingToDocBook, nil},

	{"should output non-breaking space for source and manual in docbook manpage output if absent from source", "asciidoctor/document_test_should_output_non_breaking_space_for_source_and_manual_in_docbook_manpage_output_if_absent_from_source.adoc", shouldOutputNonBreakingSpaceForSourceAndManualInDocbookManpageOutputIfAbsentFromSource, nil},

	{"should apply replacements substitution to value of mantitle attribute used in DocBook output", "asciidoctor/document_test_should_apply_replacements_substitution_to_value_of_mantitle_attribute_used_in_doc_book_output.adoc", shouldApplyReplacementsSubstitutionToValueOfMantitleAttributeUsedInDocBookOutput, nil},

	{"should be able to set doctype to book when converting to DocBook", "asciidoctor/document_test_should_be_able_to_set_doctype_to_book_when_converting_to_doc_book.adoc", shouldBeAbleToSetDoctypeToBookWhenConvertingToDocBook, nil},

	{"should be able to set doctype to book for document with no title when converting to DocBook", "asciidoctor/document_test_should_be_able_to_set_doctype_to_book_for_document_with_no_title_when_converting_to_doc_book.adoc", shouldBeAbleToSetDoctypeToBookForDocumentWithNoTitleWhenConvertingToDocBook, nil},

	{"adds a front and back cover image to DocBook 5 when doctype is book", "asciidoctor/document_test_adds_a_front_and_back_cover_image_to_doc_book_5_when_doctype_is_book.adoc", addsAFrontAndBackCoverImageToDocBook5WhenDoctypeIsBook, nil},

	{"should be able to set backend using :backend option key", "asciidoctor/document_test_should_be_able_to_set_backend_using_:backend_option_key.adoc", shouldBeAbleToSetBackendUsingbackendOptionKey, nil},

	{"attribute entry can appear immediately after document title", "asciidoctor/document_test_attribute_entry_can_appear_immediately_after_document_title.adoc", attributeEntryCanAppearImmediatelyAfterDocumentTitle, nil},

	{"attribute entry can appear before author line under document title", "asciidoctor/document_test_attribute_entry_can_appear_before_author_line_under_document_title.adoc", attributeEntryCanAppearBeforeAuthorLineUnderDocumentTitle, nil},

	{"should parse mantitle and manvolnum from document title for manpage doctype", "asciidoctor/document_test_should_parse_mantitle_and_manvolnum_from_document_title_for_manpage_doctype.adoc", shouldParseMantitleAndManvolnumFromDocumentTitleForManpageDoctype, nil},

	{"should perform attribute substitution on mantitle in manpage doctype", "asciidoctor/document_test_should_perform_attribute_substitution_on_mantitle_in_manpage_doctype.adoc", shouldPerformAttributeSubstitutionOnMantitleInManpageDoctype, nil},

	{"should consume name section as manname and manpurpose for manpage doctype", "asciidoctor/document_test_should_consume_name_section_as_manname_and_manpurpose_for_manpage_doctype.adoc", shouldConsumeNameSectionAsMannameAndManpurposeForManpageDoctype, nil},

	{"should set docname and outfilesuffix from manname and manvolnum for manpage backend and doctype", "asciidoctor/document_test_should_set_docname_and_outfilesuffix_from_manname_and_manvolnum_for_manpage_backend_and_doctype.adoc", shouldSetDocnameAndOutfilesuffixFromMannameAndManvolnumForManpageBackendAndDoctype, nil},

	{"should mark synopsis as special section in manpage doctype", "asciidoctor/document_test_should_mark_synopsis_as_special_section_in_manpage_doctype.adoc", shouldMarkSynopsisAsSpecialSectionInManpageDoctype, nil},

	{"should output special header block in HTML for manpage doctype", "asciidoctor/document_test_should_output_special_header_block_in_html_for_manpage_doctype.adoc", shouldOutputSpecialHeaderBlockInHtmlForManpageDoctype, nil},

	{"should output special header block in embeddable HTML for manpage doctype", "asciidoctor/document_test_should_output_special_header_block_in_embeddable_html_for_manpage_doctype.adoc", shouldOutputSpecialHeaderBlockInEmbeddableHtmlForManpageDoctype, nil},

	{"should output all mannames in name section in man page output", "asciidoctor/document_test_should_output_all_mannames_in_name_section_in_man_page_output.adoc", shouldOutputAllMannamesInNameSectionInManPageOutput, nil},

	{"allows us to specify a path relative to the current dir", "asciidoctor/document_test_allows_us_to_specify_a_path_relative_to_the_current_dir.adoc", allowsUsToSpecifyAPathRelativeToTheCurrentDir, nil},

	{"should raise an exception when a converter cannot be resolved while parsing", "asciidoctor/document_test_should_raise_an_exception_when_a_converter_cannot_be_resolved_while_parsing.adoc", shouldRaiseAnExceptionWhenAConverterCannotBeResolvedWhileParsing, nil},
}

var shouldBeAbleToDisableTocAndSectnumsInDocumentHeaderInDocBookBackend = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Set: asciidoc.Set{
				&asciidoc.AttributeReset{
					Name: "toc",
				},
				&asciidoc.AttributeReset{
					Name: "sectnums",
				},
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

var noheaderAttributeShouldSuppressInfoElementWhenConvertingToDocBook = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Set: asciidoc.Set{
				&asciidoc.AttributeEntry{
					Name: "noheader",
					Set:  nil,
				},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "content",
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

var shouldBeAbleToDisableSectionNumberingUsingNumberedAttributeInDocumentHeaderInDocBookBackend = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Set: asciidoc.Set{
				&asciidoc.AttributeReset{
					Name: "numbered",
				},
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

var convertMethodsOnBuiltInConverterAreRegisteredByDefault = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.InlinePassthrough{
			AttributeList: nil,
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "content",
				},
			},
		},
		&asciidoc.NewLine{},
	},
}

var shouldNotEnableCompatModeForDocumentWithLegacyDoctitleIfCompatModeDisableByHeader = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.InlinePassthrough{
			AttributeList: nil,
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "content",
				},
			},
		},
		&asciidoc.NewLine{},
	},
}

var shouldNotEnableCompatModeForDocumentWithLegacyDoctitleIfCompatModeIsLockedByApi = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.InlinePassthrough{
			AttributeList: nil,
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "content",
				},
			},
		},
		&asciidoc.NewLine{},
	},
}

var shouldApplyMaxWidthToEachTopLevelContainer = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Set: asciidoc.Set{
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "contentfootnote:[placeholder]",
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

var titlePartitionApiWithDefaultSeparator = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Author Name",
				},
				&asciidoc.NewLine{},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "content",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Main Title: *Subtitle*",
				},
			},
			Level: 0,
		},
	},
}

var documentWithSubtitleAndCustomSeparator = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.NamedAttribute{
					Name: "separator",
					Val: asciidoc.Set{
						&asciidoc.String{
							Value: "::",
						},
					},
					Quote: 0,
				},
			},
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Author Name",
				},
				&asciidoc.NewLine{},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "content",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Main Title:: *Subtitle*",
				},
			},
			Level: 0,
		},
	},
}

var shouldNotHonorCustomSeparatorForDoctitleIfAttributeIsLockedByApi = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.NamedAttribute{
					Name: "separator",
					Val: asciidoc.Set{
						&asciidoc.String{
							Value: "::",
						},
					},
					Quote: 0,
				},
			},
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Author Name",
				},
				&asciidoc.NewLine{},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "content",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Main Title - *Subtitle*",
				},
			},
			Level: 0,
		},
	},
}

var documentWithDoctitleDefinedAsAttributeEntry = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "doctitle",
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Document Title",
				},
			},
		},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "preamble",
		},
		&asciidoc.NewLine{},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Set:           nil,
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "First Section",
				},
			},
			Level: 1,
		},
	},
}

var documentWithDoctitleDefinedAsAttributeEntryFollowedByBlockWithTitle = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "doctitle",
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Document Title",
				},
			},
		},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Paragraph{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.TitleAttribute{
					Val: asciidoc.Set{
						&asciidoc.String{
							Value: "Block title",
						},
					},
				},
			},
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Block content",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 0,
		},
	},
}

var documentWithTitleAttributeEntryOverridesDoctitle = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{ // p0
			AttributeList: nil,
			Set: asciidoc.Set{
				&asciidoc.AttributeEntry{
					Name: "title",
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "Override",
						},
					},
				},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.UserAttributeReference{
					Value: "doctitle",
				},
				&asciidoc.NewLine{},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.Section{
					AttributeList: nil,
					Set:           nil,
					Title: asciidoc.Set{
						&asciidoc.String{
							Value: "First Section",
						},
					},
					Level: 1,
				},
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

var documentWithBlankTitleAttributeEntryOverridesDoctitle = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{ // p0
			AttributeList: nil,
			Set: asciidoc.Set{
				&asciidoc.AttributeEntry{
					Name: "title",
					Set:  nil,
				},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.UserAttributeReference{
					Value: "doctitle",
				},
				&asciidoc.NewLine{},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.Section{
					AttributeList: nil,
					Set:           nil,
					Title: asciidoc.Set{
						&asciidoc.String{
							Value: "First Section",
						},
					},
					Level: 1,
				},
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

var documentHeaderCanReferenceIntrinsicDoctitleAttribute = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Set: asciidoc.Set{
				&asciidoc.AttributeEntry{
					Name: "intro",
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "Welcome to the {doctitle}!",
						},
					},
				},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.UserAttributeReference{
					Value: "intro",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "ACME Documentation",
				},
			},
			Level: 0,
		},
	},
}

var documentWithTitleAttributeEntryOverridesDoctitleAttributeEntry = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{ // p0
			AttributeList: nil,
			Set: asciidoc.Set{
				&asciidoc.AttributeEntry{
					Name: "snapshot",
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "{doctitle}",
						},
					},
				},
				&asciidoc.AttributeEntry{
					Name: "doctitle",
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "doctitle",
						},
					},
				},
				&asciidoc.AttributeEntry{
					Name: "title",
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "Override",
						},
					},
				},
				asciidoc.EmptyLine{
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
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.Section{
					AttributeList: nil,
					Set:           nil,
					Title: asciidoc.Set{
						&asciidoc.String{
							Value: "First Section",
						},
					},
					Level: 1,
				},
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

var documentWithDoctitleAttributeEntryOverridesImplicitDoctitle = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{ // p0
			AttributeList: nil,
			Set: asciidoc.Set{
				&asciidoc.AttributeEntry{
					Name: "snapshot",
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "{doctitle}",
						},
					},
				},
				&asciidoc.AttributeEntry{
					Name: "doctitle",
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "Override",
						},
					},
				},
				asciidoc.EmptyLine{
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
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.Section{
					AttributeList: nil,
					Set:           nil,
					Title: asciidoc.Set{
						&asciidoc.String{
							Value: "First Section",
						},
					},
					Level: 1,
				},
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

var doctitleAttributeEntryAboveHeaderOverridesImplicitDoctitle = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "doctitle",
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Override",
				},
			},
		},
		&asciidoc.Section{ // p0
			AttributeList: nil,
			Set: asciidoc.Set{
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.UserAttributeReference{
					Value: "doctitle",
				},
				&asciidoc.NewLine{},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.Section{
					AttributeList: nil,
					Set:           nil,
					Title: asciidoc.Set{
						&asciidoc.String{
							Value: "First Section",
						},
					},
					Level: 1,
				},
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

var shouldApplyHeaderSubstitutionsToValueOfTheDoctitleAttributeAssignedFromImplicitDoctitle = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Set: asciidoc.Set{
				asciidoc.EmptyLine{
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
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "<Foo> {plus} <Bar>",
				},
			},
			Level: 0,
		},
	},
}

var shouldSubstituteAttributeReferenceInImplicitDocumentTitleForAttributeDefinedEarlierInHeader = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "project-name",
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "ACME",
				},
			},
		},
		&asciidoc.Section{
			AttributeList: nil,
			Set: asciidoc.Set{
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.UserAttributeReference{
					Value: "doctitle",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "{project-name} Docs",
				},
			},
			Level: 0,
		},
	},
}

var shouldNotWarnIfImplicitDocumentTitleContainsAttributeReferenceForAttributeDefinedLaterInHeader = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Set: asciidoc.Set{
				&asciidoc.AttributeEntry{
					Name: "project-name",
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "ACME",
						},
					},
				},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.UserAttributeReference{
					Value: "doctitle",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "{project-name} Docs",
				},
			},
			Level: 0,
		},
	},
}

var shouldRecognizeDocumentTitleWhenPrecededByBlankLines = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{ // p0
			AttributeList: nil,
			Set: asciidoc.Set{
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "preamble",
				},
				&asciidoc.NewLine{},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.Section{
					AttributeList: nil,
					Set: asciidoc.Set{
						asciidoc.EmptyLine{
							Text: "",
						},
						&asciidoc.String{
							Value: "text",
						},
						&asciidoc.NewLine{},
					},
					Title: asciidoc.Set{
						&asciidoc.String{
							Value: "Section 1",
						},
					},
					Level: 1,
				},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Title",
				},
			},
			Level: 0,
		},
	},
}

var shouldRecognizeDocumentTitleWhenPrecededByBlankLinesIntroducedByAPreprocessorConditional = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.IfDef{
			Attributes: asciidoc.AttributeNames{
				"sectids",
			},
			Union: 0,
		},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "foo",
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "bar",
				},
			},
		},
		&asciidoc.EndIf{
			Attributes: nil,
			Union:      0,
		},
		&asciidoc.Section{ // p0
			AttributeList: nil,
			Set: asciidoc.Set{
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "preamble",
				},
				&asciidoc.NewLine{},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.Section{
					AttributeList: nil,
					Set: asciidoc.Set{
						asciidoc.EmptyLine{
							Text: "",
						},
						&asciidoc.String{
							Value: "text",
						},
						&asciidoc.NewLine{},
					},
					Title: asciidoc.Set{
						&asciidoc.String{
							Value: "Section 1",
						},
					},
					Level: 1,
				},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Title",
				},
			},
			Level: 0,
		},
	},
}

var shouldRecognizeDocumentTitleWhenPrecededByBlankLinesAfterAnAttributeEntry = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "doctype",
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "book",
				},
			},
		},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{ // p0
			AttributeList: nil,
			Set: asciidoc.Set{
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "preamble",
				},
				&asciidoc.NewLine{},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.Section{
					AttributeList: nil,
					Set: asciidoc.Set{
						asciidoc.EmptyLine{
							Text: "",
						},
						&asciidoc.String{
							Value: "text",
						},
						&asciidoc.NewLine{},
					},
					Title: asciidoc.Set{
						&asciidoc.String{
							Value: "Section 1",
						},
					},
					Level: 1,
				},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Title",
				},
			},
			Level: 0,
		},
	},
}

var shouldRecognizeDocumentTitleInIncludeFileWhenPrecededByBlankLines = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.FileInclude{
			AttributeList: nil,
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "fixtures/include-with-leading-blank-line.adoc",
				},
			},
		},
	},
}

var shouldIncludeSpecifiedLinesEvenWhenLeadingLinesAreSkipped = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.FileInclude{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.NamedAttribute{
					Name: "lines",
					Val: asciidoc.Set{
						&asciidoc.String{
							Value: "6",
						},
					},
					Quote: 0,
				},
			},
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "fixtures/include-with-leading-blank-line.adoc",
				},
			},
		},
	},
}

var documentWithMultilineAttributeEntryButOnlyOneLineShouldNotCrash = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Set: asciidoc.Set{
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "content",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "*Document* image:logo.png[] _Title_ image:another-logo.png[another logo]",
				},
			},
			Level: 0,
		},
	},
}

var shouldNotChokeOnEmptySource = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{ // p0
			AttributeList: nil,
			Set: asciidoc.Set{
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
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "AsciiDoc user guide",
						},
					},
				},
				&asciidoc.AttributeEntry{
					Name: "keywords",
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "asciidoc,documentation",
						},
					},
				},
				&asciidoc.AttributeEntry{
					Name: "copyright",
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "Stuart Rackham",
						},
					},
				},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.Section{
					AttributeList: nil,
					Set: asciidoc.Set{
						asciidoc.EmptyLine{
							Text: "",
						},
						&asciidoc.String{
							Value: "more info...",
						},
						&asciidoc.NewLine{},
					},
					Title: asciidoc.Set{
						&asciidoc.String{
							Value: "Version 8.6.8",
						},
					},
					Level: 1,
				},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "AsciiDoc",
				},
			},
			Level: 0,
		},
	},
}

var shouldParseRevisionLineIfDateIsEmpty = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Author Name",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "v1.0.0,:remark",
				},
				&asciidoc.NewLine{},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "content",
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

var shouldIncludeRevisionHistoryInDocBookOutputIfRevdateAndRevnumberIsSet = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Author Name",
				},
				&asciidoc.NewLine{},
				&asciidoc.AttributeEntry{
					Name: "revdate",
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "2011-11-11",
						},
					},
				},
				&asciidoc.AttributeEntry{
					Name: "revnumber",
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "1.0",
						},
					},
				},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "content",
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

var shouldIncludeRevisionHistoryInDocBookOutputIfRevdateAndRevremarkIsSet = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Author Name",
				},
				&asciidoc.NewLine{},
				&asciidoc.AttributeEntry{
					Name: "revdate",
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "2011-11-11",
						},
					},
				},
				&asciidoc.AttributeEntry{
					Name: "revremark",
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "features!",
						},
					},
				},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "content",
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

var shouldNotIncludeRevisionHistoryInDocBookOutputIfRevdateIsNotSet = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Author Name",
				},
				&asciidoc.NewLine{},
				&asciidoc.AttributeEntry{
					Name: "revnumber",
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "1.0",
						},
					},
				},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "content",
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

var withMetadataToDocBook5 = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{ // p0
			AttributeList: nil,
			Set: asciidoc.Set{
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
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.Section{
					AttributeList: nil,
					Set: asciidoc.Set{
						asciidoc.EmptyLine{
							Text: "",
						},
						&asciidoc.String{
							Value: "more info...",
						},
						&asciidoc.NewLine{},
					},
					Title: asciidoc.Set{
						&asciidoc.String{
							Value: "Version 8.6.8",
						},
					},
					Level: 1,
				},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "AsciiDoc",
				},
			},
			Level: 0,
		},
	},
}

var withDocumentIdToDocbook5 = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.AnchorAttribute{
					ID: &asciidoc.String{
						Value: "document-id",
					},
					Label: nil,
				},
			},
			Set: asciidoc.Set{
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "more info...",
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

var withAuthorDefinedUsingAttributeEntryToDocBook = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Set: asciidoc.Set{
				&asciidoc.AttributeEntry{
					Name: "author",
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "Doc Writer",
						},
					},
				},
				&asciidoc.AttributeEntry{
					Name: "email",
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "thedoctor@asciidoc.org",
						},
					},
				},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "content",
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

var shouldSubstituteReplacementsInAuthorNamesInHtmlOutput = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Set: asciidoc.Set{
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
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "content",
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

var shouldSubstituteReplacementsInAuthorNamesInDocBookOutput = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Set: asciidoc.Set{
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
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "content",
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

var shouldSanitizeContentOfHtmlMetaAuthorsTag = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Set: asciidoc.Set{
				&asciidoc.AttributeEntry{
					Name: "author",
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "pass:n[http://example.org/community/team.html[Ze *Product* team]]",
						},
					},
				},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "content",
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

var shouldNotDoubleEscapeAmpersandInAuthorAttribute = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "R&D Lab",
				},
				&asciidoc.NewLine{},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.UserAttributeReference{
					Value: "author",
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

var shouldIncludeMultipleAuthorsInHtmlOutput = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Set: asciidoc.Set{
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
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "content",
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

var shouldCreateAuthorgroupInDocBookWhenMultipleAuthors = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Set: asciidoc.Set{
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
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "content",
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

var shouldProcessAuthorDefinedByAttributeWhenImplicitDoctitleIsAbsent = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "author",
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Doc Writer",
				},
			},
		},
		asciidoc.EmptyLine{
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

var shouldProcessAuthorAndAuthorinitialsDefinedByAttributeWhenImplicitDoctitleIsAbsent = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "authorinitials",
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "DOC",
				},
			},
		},
		&asciidoc.AttributeEntry{
			Name: "author",
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Doc Writer",
				},
			},
		},
		asciidoc.EmptyLine{
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

var shouldProcessAuthorsDefinedByAttributeWhenImplicitDoctitleIsAbsent = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "authors",
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Doc Writer; Other Author",
				},
			},
		},
		asciidoc.EmptyLine{
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

var shouldProcessAuthorsAndAuthorinitialsDefinedByAttributeWhenImplicitDoctitleIsAbsent = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "authorinitials",
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "DOC",
				},
			},
		},
		&asciidoc.AttributeEntry{
			Name: "authors",
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Doc Writer; Other Author",
				},
			},
		},
		asciidoc.EmptyLine{
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

var shouldSetAuthorcountTo0IfDocumentHasNoHeader = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "idprefix",
			Set:  nil,
		},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Set: asciidoc.Set{
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "content",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Section Title",
				},
			},
			Level: 1,
		},
	},
}

var shouldSetAuthorcountTo0IfAuthorNotSetByAttributeAndDocumentStartsWithLevel0SectionWithStyle = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "doctype",
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "book",
				},
			},
		},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Set: asciidoc.Set{
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
			Set: asciidoc.Set{
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "content",
				},
				&asciidoc.NewLine{},
				asciidoc.EmptyLine{
					Text: "",
				},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Preface",
				},
			},
			Level: 0,
		},
		&asciidoc.Section{ // p0
			AttributeList: nil,
			Set: asciidoc.Set{
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.Section{
					AttributeList: nil,
					Set: asciidoc.Set{
						asciidoc.EmptyLine{
							Text: "",
						},
						&asciidoc.String{
							Value: "content",
						},
						&asciidoc.NewLine{},
					},
					Title: asciidoc.Set{
						&asciidoc.String{
							Value: "Chapter",
						},
					},
					Level: 1,
				},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Part",
				},
			},
			Level: 0,
		},
	},
}

var withAuthorDefinedByIndexedAttributeName = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Set: asciidoc.Set{
				&asciidoc.AttributeEntry{
					Name: "author_1",
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "Doc Writer",
						},
					},
				},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.UserAttributeReference{
					Value: "author",
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

var withAuthorsDefinedUsingAttributeEntryToDocBook = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Set: asciidoc.Set{
				&asciidoc.AttributeEntry{
					Name: "authors",
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "Doc Writer; Junior Writer",
						},
					},
				},
				&asciidoc.AttributeEntry{
					Name: "email_1",
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "thedoctor@asciidoc.org",
						},
					},
				},
				&asciidoc.AttributeEntry{
					Name: "email_2",
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "junior@asciidoc.org",
						},
					},
				},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "content",
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

var shouldPopulateCopyrightElementInDocBookOutputIfCopyrightAttributeIsDefined = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Set: asciidoc.Set{
				&asciidoc.AttributeEntry{
					Name: "copyright",
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "ACME, Inc.",
						},
					},
				},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "Essential for catching road runners.",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Jet Bike",
				},
			},
			Level: 0,
		},
	},
}

var shouldPopulateCopyrightElementInDocBookOutputIfCopyrightAttributeIsDefinedWithYear = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Set: asciidoc.Set{
				&asciidoc.AttributeEntry{
					Name: "copyright",
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "ACME, Inc. 1956",
						},
					},
				},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "Essential for catching road runners.",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Jet Bike",
				},
			},
			Level: 0,
		},
	},
}

var shouldPopulateCopyrightElementInDocBookOutputIfCopyrightAttributeIsDefinedWithYearRange = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Set: asciidoc.Set{
				&asciidoc.AttributeEntry{
					Name: "copyright",
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "ACME, Inc. 1956-2018",
						},
					},
				},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "Essential for catching road runners.",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Jet Bike",
				},
			},
			Level: 0,
		},
	},
}

var withHeaderFooter = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "nofooter",
			Set:  nil,
		},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "content",
		},
		&asciidoc.NewLine{},
	},
}

var canDisableLastUpdatedInFooter = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Set: asciidoc.Set{
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "content",
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

var parseHeaderOnly = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Author Name",
				},
				&asciidoc.NewLine{},
				&asciidoc.AttributeEntry{
					Name: "foo",
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "bar",
						},
					},
				},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "preamble",
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

var shouldParseHeaderOnlyWhenDocytpeIsManpage = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{ // p0
			AttributeList: nil,
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Author Name",
				},
				&asciidoc.NewLine{},
				&asciidoc.AttributeEntry{
					Name: "doctype",
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "manpage",
						},
					},
				},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.Section{
					AttributeList: nil,
					Set: asciidoc.Set{
						asciidoc.EmptyLine{
							Text: "",
						},
						&asciidoc.String{
							Value: "cmd - does stuff",
						},
						&asciidoc.NewLine{},
					},
					Title: asciidoc.Set{
						&asciidoc.String{
							Value: "Name",
						},
					},
					Level: 1,
				},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "cmd(1)",
				},
			},
			Level: 0,
		},
	},
}

var shouldNotWarnWhenParsingHeaderOnlyWhenDocytpeIsManpageAndBodyIsEmpty = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Author Name",
				},
				&asciidoc.NewLine{},
				&asciidoc.AttributeEntry{
					Name: "doctype",
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "manpage",
						},
					},
				},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "cmd(1)",
				},
			},
			Level: 0,
		},
	},
}

var outputsFootnotesInFooter = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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

var outputsFootnotesBlockInEmbeddedDocumentByDefault = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{ // p0
			AttributeList: nil,
			Set: asciidoc.Set{
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.Section{
					AttributeList: nil,
					Set: asciidoc.Set{
						asciidoc.EmptyLine{
							Text: "",
						},
						&asciidoc.String{
							Value: "Content",
						},
						&asciidoc.NewLine{},
						asciidoc.EmptyLine{
							Text: "",
						},
					},
					Title: asciidoc.Set{
						&asciidoc.String{
							Value: "Section A",
						},
					},
					Level: 1,
				},
				&asciidoc.Section{
					AttributeList: nil,
					Set: asciidoc.Set{
						asciidoc.EmptyLine{
							Text: "",
						},
						&asciidoc.String{
							Value: "Content.footnote:[commentary]",
						},
						&asciidoc.NewLine{},
					},
					Title: asciidoc.Set{
						&asciidoc.String{
							Value: "Section B",
						},
					},
					Level: 1,
				},
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

var shouldReturnEmptyidsTable = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.BlockImage{
			AttributeList: nil,
			Path: asciidoc.Set{
				&asciidoc.String{
					Value: "outer.png",
				},
			},
		},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Table{
			AttributeList: nil,
			ColumnCount:   1,
			Set: asciidoc.Set{
				&asciidoc.TableRow{
					Set: asciidoc.Set{
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
							Set: asciidoc.Set{
								&asciidoc.NewLine{},
								&asciidoc.BlockImage{
									AttributeList: nil,
									Path: asciidoc.Set{
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

var honorHtmlsyntaxAttributeInDocumentHeaderIfFollowedByBackendAttribute = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "htmlsyntax",
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "xml",
				},
			},
		},
		&asciidoc.AttributeEntry{
			Name: "backend",
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "html5",
				},
			},
		},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.ThematicBreak{
			AttributeList: nil,
		},
	},
}

var doesNotHonorHtmlsyntaxAttributeInDocumentHeaderIfNotFollowedByBackendAttribute = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "backend",
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "html5",
				},
			},
		},
		&asciidoc.AttributeEntry{
			Name: "htmlsyntax",
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "xml",
				},
			},
		},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.ThematicBreak{
			AttributeList: nil,
		},
	},
}

var shouldCloseAllShortTagsWhenHtmlsyntaxIsXml = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Author Name",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "v1.0, 2001-01-01",
				},
				&asciidoc.NewLine{},
				&asciidoc.AttributeEntry{
					Name: "icons",
					Set:  nil,
				},
				&asciidoc.AttributeEntry{
					Name: "favicon",
					Set:  nil,
				},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.InlineImage{
					AttributeList: nil,
					Path: asciidoc.Set{
						&asciidoc.String{
							Value: "tiger.png",
						},
					},
				},
				&asciidoc.NewLine{},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.BlockImage{
					AttributeList: nil,
					Path: asciidoc.Set{
						&asciidoc.String{
							Value: "tiger.png",
						},
					},
				},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.UnorderedListItem{
					Set: asciidoc.Set{
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
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "two",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        "*",
					Checklist:     1,
				},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.Table{
					AttributeList: nil,
					ColumnCount:   2,
					Set: asciidoc.Set{
						&asciidoc.TableRow{
							Set: asciidoc.Set{
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
									Set: asciidoc.Set{
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
									Set: asciidoc.Set{
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
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.Paragraph{
					AttributeList: asciidoc.AttributeList{
						&asciidoc.ShorthandAttribute{
							Style: &asciidoc.ShorthandStyle{
								Set: asciidoc.Set{
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
							Val: asciidoc.Set{
								&asciidoc.String{
									Value: "25%",
								},
							},
							Quote: 2,
						},
						&asciidoc.NamedAttribute{
							Name: "itemwidth",
							Val: asciidoc.Set{
								&asciidoc.String{
									Value: "75%",
								},
							},
							Quote: 2,
						},
					},
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "term:: description",
						},
						&asciidoc.NewLine{},
					},
					Admonition: 0,
				},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.Paragraph{
					AttributeList: nil,
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "note",
						},
						&asciidoc.NewLine{},
					},
					Admonition: 1,
				},
				asciidoc.EmptyLine{
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
								Set: asciidoc.Set{
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
							Val: asciidoc.Set{
								&asciidoc.String{
									Value: "Author",
								},
							},
						},
						&asciidoc.PositionalAttribute{
							Offset:      2,
							ImpliedName: "",
							Val: asciidoc.Set{
								&asciidoc.String{
									Value: "Source",
								},
							},
						},
					},
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "Quote me.",
						},
						&asciidoc.NewLine{},
					},
				},
				asciidoc.EmptyLine{
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
								Set: asciidoc.Set{
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
							Val: asciidoc.Set{
								&asciidoc.String{
									Value: "Author",
								},
							},
						},
						&asciidoc.PositionalAttribute{
							Offset:      2,
							ImpliedName: "",
							Val: asciidoc.Set{
								&asciidoc.String{
									Value: "Source",
								},
							},
						},
					},
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "A tall tale.",
						},
						&asciidoc.NewLine{},
					},
				},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.Paragraph{
					AttributeList: asciidoc.AttributeList{
						&asciidoc.NamedAttribute{
							Name: "options",
							Val: asciidoc.Set{
								&asciidoc.String{
									Value: "autoplay,loop",
								},
							},
							Quote: 2,
						},
					},
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "video::screencast.ogg[]",
						},
						&asciidoc.NewLine{},
					},
					Admonition: 0,
				},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "video::12345[vimeo]",
				},
				&asciidoc.NewLine{},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.Paragraph{
					AttributeList: asciidoc.AttributeList{
						&asciidoc.NamedAttribute{
							Name: "options",
							Val: asciidoc.Set{
								&asciidoc.String{
									Value: "autoplay,loop",
								},
							},
							Quote: 2,
						},
					},
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "audio::podcast.ogg[]",
						},
						&asciidoc.NewLine{},
					},
					Admonition: 0,
				},
				asciidoc.EmptyLine{
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
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.ThematicBreak{
					AttributeList: nil,
				},
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

var xhtmlBackendShouldEmitElementsInProperNamespace = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Set: asciidoc.Set{
				&asciidoc.AttributeEntry{
					Name: "doctype",
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "book",
						},
					},
				},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "text",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Document Title: Subtitle",
				},
			},
			Level: 0,
		},
	},
}

var shouldBeAbleToSetDoctypeToArticleWhenConvertingToDocBook = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{ // p0
			AttributeList: nil,
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Author Name",
				},
				&asciidoc.NewLine{},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "preamble",
				},
				&asciidoc.NewLine{},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.Section{
					AttributeList: nil,
					Set: asciidoc.Set{
						asciidoc.EmptyLine{
							Text: "",
						},
						&asciidoc.String{
							Value: "section body",
						},
						&asciidoc.NewLine{},
					},
					Title: asciidoc.Set{
						&asciidoc.String{
							Value: "First Section",
						},
					},
					Level: 1,
				},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Title",
				},
			},
			Level: 0,
		},
	},
}

var shouldSetDoctypeToArticleByDefaultForDocumentWithNoTitleWhenConvertingToDocBook = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{ // p0
			AttributeList: nil,
			Set: asciidoc.Set{
				&asciidoc.AttributeEntry{
					Name: "mansource",
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "Asciidoctor",
						},
					},
				},
				&asciidoc.AttributeEntry{
					Name: "manmanual",
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "Asciidoctor Manual",
						},
					},
				},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.Section{
					AttributeList: nil,
					Set: asciidoc.Set{
						asciidoc.EmptyLine{
							Text: "",
						},
						&asciidoc.String{
							Value: "asciidoctor - Process text",
						},
						&asciidoc.NewLine{},
						asciidoc.EmptyLine{
							Text: "",
						},
					},
					Title: asciidoc.Set{
						&asciidoc.String{
							Value: "NAME",
						},
					},
					Level: 1,
				},
				&asciidoc.Section{
					AttributeList: nil,
					Set: asciidoc.Set{
						asciidoc.EmptyLine{
							Text: "",
						},
						&asciidoc.String{
							Value: "some text",
						},
						&asciidoc.NewLine{},
						asciidoc.EmptyLine{
							Text: "",
						},
					},
					Title: asciidoc.Set{
						&asciidoc.String{
							Value: "SYNOPSIS",
						},
					},
					Level: 1,
				},
				&asciidoc.Section{
					AttributeList: nil,
					Set: asciidoc.Set{
						asciidoc.EmptyLine{
							Text: "",
						},
						&asciidoc.String{
							Value: "section body",
						},
						&asciidoc.NewLine{},
					},
					Title: asciidoc.Set{
						&asciidoc.String{
							Value: "First Section",
						},
					},
					Level: 1,
				},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "asciidoctor(1)",
				},
			},
			Level: 0,
		},
	},
}

var shouldOutputNonBreakingSpaceForSourceAndManualInDocbookManpageOutputIfAbsentFromSource = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{ // p0
			AttributeList: nil,
			Set: asciidoc.Set{
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.Section{
					AttributeList: nil,
					Set: asciidoc.Set{
						asciidoc.EmptyLine{
							Text: "",
						},
						&asciidoc.String{
							Value: "asciidoctor - Process text",
						},
						&asciidoc.NewLine{},
						asciidoc.EmptyLine{
							Text: "",
						},
					},
					Title: asciidoc.Set{
						&asciidoc.String{
							Value: "NAME",
						},
					},
					Level: 1,
				},
				&asciidoc.Section{
					AttributeList: nil,
					Set: asciidoc.Set{
						asciidoc.EmptyLine{
							Text: "",
						},
						&asciidoc.String{
							Value: "some text",
						},
						&asciidoc.NewLine{},
					},
					Title: asciidoc.Set{
						&asciidoc.String{
							Value: "SYNOPSIS",
						},
					},
					Level: 1,
				},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "asciidoctor(1)",
				},
			},
			Level: 0,
		},
	},
}

var shouldApplyReplacementsSubstitutionToValueOfMantitleAttributeUsedInDocBookOutput = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{ // p0
			AttributeList: nil,
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Author Name",
				},
				&asciidoc.NewLine{},
				&asciidoc.AttributeEntry{
					Name: "doctype",
					Set: asciidoc.Set{
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
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.Section{
					AttributeList: nil,
					Set: asciidoc.Set{
						asciidoc.EmptyLine{
							Text: "",
						},
						&asciidoc.String{
							Value: "foo--bar - puts the foo in your bar",
						},
						&asciidoc.NewLine{},
					},
					Title: asciidoc.Set{
						&asciidoc.String{
							Value: "NAME",
						},
					},
					Level: 1,
				},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "foo\\--bar(1)",
				},
			},
			Level: 0,
		},
	},
}

var shouldBeAbleToSetDoctypeToBookWhenConvertingToDocBook = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{ // p0
			AttributeList: nil,
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Author Name",
				},
				&asciidoc.NewLine{},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "preamble",
				},
				&asciidoc.NewLine{},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.Section{
					AttributeList: nil,
					Set: asciidoc.Set{
						asciidoc.EmptyLine{
							Text: "",
						},
						&asciidoc.String{
							Value: "chapter body",
						},
						&asciidoc.NewLine{},
					},
					Title: asciidoc.Set{
						&asciidoc.String{
							Value: "First Chapter",
						},
					},
					Level: 1,
				},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Title",
				},
			},
			Level: 0,
		},
	},
}

var shouldBeAbleToSetDoctypeToBookForDocumentWithNoTitleWhenConvertingToDocBook = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{ // p0
			AttributeList: nil,
			Set: asciidoc.Set{
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
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "manpage",
						},
					},
				},
				&asciidoc.AttributeEntry{
					Name: "manmanual",
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "EVE",
						},
					},
				},
				&asciidoc.AttributeEntry{
					Name: "mansource",
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "EVE",
						},
					},
				},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.Section{
					AttributeList: nil,
					Set: asciidoc.Set{
						asciidoc.EmptyLine{
							Text: "",
						},
						&asciidoc.String{
							Value: "eve, islifeform - analyzes an image to determine if it's a picture of a life form",
						},
						&asciidoc.NewLine{},
						asciidoc.EmptyLine{
							Text: "",
						},
					},
					Title: asciidoc.Set{
						&asciidoc.String{
							Value: "NAME",
						},
					},
					Level: 1,
				},
				&asciidoc.Section{
					AttributeList: nil,
					Set: asciidoc.Set{
						asciidoc.EmptyLine{
							Text: "",
						},
						&asciidoc.Bold{
							AttributeList: nil,
							Set: asciidoc.Set{
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
					Title: asciidoc.Set{
						&asciidoc.String{
							Value: "SYNOPSIS",
						},
					},
					Level: 1,
				},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "eve(1)",
				},
			},
			Level: 0,
		},
	},
}

var addsAFrontAndBackCoverImageToDocBook5WhenDoctypeIsBook = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{ // p0
			AttributeList: nil,
			Set: asciidoc.Set{
				&asciidoc.AttributeEntry{
					Name: "doctype",
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "book",
						},
					},
				},
				&asciidoc.AttributeEntry{
					Name: "imagesdir",
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "images",
						},
					},
				},
				&asciidoc.AttributeEntry{
					Name: "front-cover-image",
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "image:front-cover.jpg[scaledwidth=210mm]",
						},
					},
				},
				&asciidoc.AttributeEntry{
					Name: "back-cover-image",
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "image:back-cover.jpg[]",
						},
					},
				},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "preamble",
				},
				&asciidoc.NewLine{},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.Section{
					AttributeList: nil,
					Set: asciidoc.Set{
						asciidoc.EmptyLine{
							Text: "",
						},
						&asciidoc.String{
							Value: "chapter body",
						},
						&asciidoc.NewLine{},
					},
					Title: asciidoc.Set{
						&asciidoc.String{
							Value: "First Chapter",
						},
					},
					Level: 1,
				},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Title",
				},
			},
			Level: 0,
		},
	},
}

var shouldBeAbleToSetBackendUsingbackendOptionKey = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Set: asciidoc.Set{
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
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "more info...",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "AsciiDoc",
				},
			},
			Level: 0,
		},
	},
}

var attributeEntryCanAppearImmediatelyAfterDocumentTitle = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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
			Name: "toc",
			Set:  nil,
		},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "preamble",
		},
		&asciidoc.NewLine{},
	},
}

var attributeEntryCanAppearBeforeAuthorLineUnderDocumentTitle = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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
			Name: "toc",
			Set:  nil,
		},
		&asciidoc.String{
			Value: "Dan Allen",
		},
		&asciidoc.NewLine{},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "preamble",
		},
		&asciidoc.NewLine{},
	},
}

var shouldParseMantitleAndManvolnumFromDocumentTitleForManpageDoctype = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{ // p0
			AttributeList: nil,
			Set: asciidoc.Set{
				&asciidoc.AttributeEntry{
					Name: "doctype",
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "manpage",
						},
					},
				},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.Section{
					AttributeList: nil,
					Set: asciidoc.Set{
						asciidoc.EmptyLine{
							Text: "",
						},
						&asciidoc.String{
							Value: "asciidoctor - converts AsciiDoc source files to HTML, DocBook and other formats",
						},
						&asciidoc.NewLine{},
					},
					Title: asciidoc.Set{
						&asciidoc.String{
							Value: "NAME",
						},
					},
					Level: 1,
				},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "asciidoctor ( 1 )",
				},
			},
			Level: 0,
		},
	},
}

var shouldPerformAttributeSubstitutionOnMantitleInManpageDoctype = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{ // p0
			AttributeList: nil,
			Set: asciidoc.Set{
				&asciidoc.AttributeEntry{
					Name: "doctype",
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "manpage",
						},
					},
				},
				&asciidoc.AttributeEntry{
					Name: "app",
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "Asciidoctor",
						},
					},
				},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.Section{
					AttributeList: nil,
					Set: asciidoc.Set{
						asciidoc.EmptyLine{
							Text: "",
						},
						&asciidoc.String{
							Value: "asciidoctor - converts AsciiDoc source files to HTML, DocBook and other formats",
						},
						&asciidoc.NewLine{},
					},
					Title: asciidoc.Set{
						&asciidoc.String{
							Value: "NAME",
						},
					},
					Level: 1,
				},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "{app}(1)",
				},
			},
			Level: 0,
		},
	},
}

var shouldConsumeNameSectionAsMannameAndManpurposeForManpageDoctype = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{ // p0
			AttributeList: nil,
			Set: asciidoc.Set{
				&asciidoc.AttributeEntry{
					Name: "doctype",
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "manpage",
						},
					},
				},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.Section{
					AttributeList: nil,
					Set: asciidoc.Set{
						asciidoc.EmptyLine{
							Text: "",
						},
						&asciidoc.String{
							Value: "asciidoctor - converts AsciiDoc source files to HTML, DocBook and other formats",
						},
						&asciidoc.NewLine{},
					},
					Title: asciidoc.Set{
						&asciidoc.String{
							Value: "NAME",
						},
					},
					Level: 1,
				},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "asciidoctor(1)",
				},
			},
			Level: 0,
		},
	},
}

var shouldSetDocnameAndOutfilesuffixFromMannameAndManvolnumForManpageBackendAndDoctype = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{ // p0
			AttributeList: nil,
			Set: asciidoc.Set{
				&asciidoc.AttributeEntry{
					Name: "doctype",
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "manpage",
						},
					},
				},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.Section{
					AttributeList: nil,
					Set: asciidoc.Set{
						asciidoc.EmptyLine{
							Text: "",
						},
						&asciidoc.String{
							Value: "asciidoctor - converts AsciiDoc source files to HTML, DocBook and other formats",
						},
						&asciidoc.NewLine{},
					},
					Title: asciidoc.Set{
						&asciidoc.String{
							Value: "NAME",
						},
					},
					Level: 1,
				},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "asciidoctor(1)",
				},
			},
			Level: 0,
		},
	},
}

var shouldMarkSynopsisAsSpecialSectionInManpageDoctype = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{ // p0
			AttributeList: nil,
			Set: asciidoc.Set{
				&asciidoc.AttributeEntry{
					Name: "doctype",
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "manpage",
						},
					},
				},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.Section{
					AttributeList: nil,
					Set: asciidoc.Set{
						asciidoc.EmptyLine{
							Text: "",
						},
						&asciidoc.String{
							Value: "asciidoctor - converts AsciiDoc source files to HTML, DocBook and other formats",
						},
						&asciidoc.NewLine{},
						asciidoc.EmptyLine{
							Text: "",
						},
					},
					Title: asciidoc.Set{
						&asciidoc.String{
							Value: "NAME",
						},
					},
					Level: 1,
				},
				&asciidoc.Section{
					AttributeList: nil,
					Set: asciidoc.Set{
						asciidoc.EmptyLine{
							Text: "",
						},
						&asciidoc.Bold{
							AttributeList: nil,
							Set: asciidoc.Set{
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
					Title: asciidoc.Set{
						&asciidoc.String{
							Value: "SYNOPSIS",
						},
					},
					Level: 1,
				},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "asciidoctor(1)",
				},
			},
			Level: 0,
		},
	},
}

var shouldOutputSpecialHeaderBlockInHtmlForManpageDoctype = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{ // p0
			AttributeList: nil,
			Set: asciidoc.Set{
				&asciidoc.AttributeEntry{
					Name: "doctype",
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "manpage",
						},
					},
				},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.Section{
					AttributeList: nil,
					Set: asciidoc.Set{
						asciidoc.EmptyLine{
							Text: "",
						},
						&asciidoc.String{
							Value: "asciidoctor - converts AsciiDoc source files to HTML, DocBook and other formats",
						},
						&asciidoc.NewLine{},
						asciidoc.EmptyLine{
							Text: "",
						},
					},
					Title: asciidoc.Set{
						&asciidoc.String{
							Value: "NAME",
						},
					},
					Level: 1,
				},
				&asciidoc.Section{
					AttributeList: nil,
					Set: asciidoc.Set{
						asciidoc.EmptyLine{
							Text: "",
						},
						&asciidoc.Bold{
							AttributeList: nil,
							Set: asciidoc.Set{
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
					Title: asciidoc.Set{
						&asciidoc.String{
							Value: "SYNOPSIS",
						},
					},
					Level: 1,
				},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "asciidoctor(1)",
				},
			},
			Level: 0,
		},
	},
}

var shouldOutputSpecialHeaderBlockInEmbeddableHtmlForManpageDoctype = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{ // p0
			AttributeList: nil,
			Set: asciidoc.Set{
				&asciidoc.AttributeEntry{
					Name: "doctype",
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "manpage",
						},
					},
				},
				&asciidoc.AttributeEntry{
					Name: "showtitle",
					Set:  nil,
				},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.Section{
					AttributeList: nil,
					Set: asciidoc.Set{
						asciidoc.EmptyLine{
							Text: "",
						},
						&asciidoc.String{
							Value: "asciidoctor - converts AsciiDoc source files to HTML, DocBook and other formats",
						},
						&asciidoc.NewLine{},
						asciidoc.EmptyLine{
							Text: "",
						},
					},
					Title: asciidoc.Set{
						&asciidoc.String{
							Value: "NAME",
						},
					},
					Level: 1,
				},
				&asciidoc.Section{
					AttributeList: nil,
					Set: asciidoc.Set{
						asciidoc.EmptyLine{
							Text: "",
						},
						&asciidoc.Bold{
							AttributeList: nil,
							Set: asciidoc.Set{
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
					Title: asciidoc.Set{
						&asciidoc.String{
							Value: "SYNOPSIS",
						},
					},
					Level: 1,
				},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "asciidoctor(1)",
				},
			},
			Level: 0,
		},
	},
}

var shouldOutputAllMannamesInNameSectionInManPageOutput = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{ // p0
			AttributeList: nil,
			Set: asciidoc.Set{
				&asciidoc.AttributeEntry{
					Name: "doctype",
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "manpage",
						},
					},
				},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.Section{
					AttributeList: nil,
					Set: asciidoc.Set{
						asciidoc.EmptyLine{
							Text: "",
						},
						&asciidoc.String{
							Value: "eve, probe - analyzes an image to determine if it is a picture of a life form",
						},
						&asciidoc.NewLine{},
						asciidoc.EmptyLine{
							Text: "",
						},
					},
					Title: asciidoc.Set{
						&asciidoc.String{
							Value: "NAME",
						},
					},
					Level: 1,
				},
				&asciidoc.Section{
					AttributeList: nil,
					Set: asciidoc.Set{
						asciidoc.EmptyLine{
							Text: "",
						},
						&asciidoc.Bold{
							AttributeList: nil,
							Set: asciidoc.Set{
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
					Title: asciidoc.Set{
						&asciidoc.String{
							Value: "SYNOPSIS",
						},
					},
					Level: 1,
				},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "eve(1)",
				},
			},
			Level: 0,
		},
	},
}

var allowsUsToSpecifyAPathRelativeToTheCurrentDir = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Set: asciidoc.Set{
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "text",
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

var shouldRaiseAnExceptionWhenAConverterCannotBeResolvedWhileParsing = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{ // p0
			AttributeList: nil,
			Set: asciidoc.Set{
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.Section{
					AttributeList: nil,
					Set: asciidoc.Set{
						asciidoc.EmptyLine{
							Text: "",
						},
						&asciidoc.String{
							Value: "text",
						},
						&asciidoc.NewLine{},
					},
					Title: asciidoc.Set{
						&asciidoc.String{
							Value: "A _Big_ Section",
						},
					},
					Level: 1,
				},
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
