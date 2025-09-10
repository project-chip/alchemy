package tests

import (
	"testing"

	"github.com/project-chip/alchemy/asciidoc"
)

func TestLinks(t *testing.T) {
	linksTests.run(t)
}

var linksTests = parseTests{

	{"unescapes square bracket in reftext of anchor macro", "asciidoctor/links_test_unescapes_square_bracket_in_reftext_of_anchor_macro.adoc", linksTestUnescapesSquareBracketInReftextOfAnchorMacro},

	{"xref using angled bracket syntax with label", "asciidoctor/links_test_xref_using_angled_bracket_syntax_with_label.adoc", linksTestXrefUsingAngledBracketSyntaxWithLabel},

	{"xref should use title of target as link text when no explicit reftext is specified", "asciidoctor/links_test_xref_should_use_title_of_target_as_link_text_when_no_explicit_reftext_is_specified.adoc", linksTestXrefShouldUseTitleOfTargetAsLinkTextWhenNoExplicitReftextIsSpecified},

	{"xref should use title of target as link text when explicit link text is empty", "asciidoctor/links_test_xref_should_use_title_of_target_as_link_text_when_explicit_link_text_is_empty.adoc", linksTestXrefShouldUseTitleOfTargetAsLinkTextWhenExplicitLinkTextIsEmpty},

	{"xref using angled bracket syntax with quoted label", "asciidoctor/links_test_xref_using_angled_bracket_syntax_with_quoted_label.adoc", linksTestXrefUsingAngledBracketSyntaxWithQuotedLabel},

	{"xref using angled bracket syntax inline with text", "asciidoctor/links_test_xref_using_angled_bracket_syntax_inline_with_text.adoc", linksTestXrefUsingAngledBracketSyntaxInlineWithText},

	{"xref using angled bracket syntax with multi-line label inline with text", "asciidoctor/links_test_xref_using_angled_bracket_syntax_with_multi_line_label_inline_with_text.adoc", linksTestXrefUsingAngledBracketSyntaxWithMultiLineLabelInlineWithText},

	{"xref with escaped text", "asciidoctor/links_test_xref_with_escaped_text.adoc", linksTestXrefWithEscapedText},

	{"xref using macro syntax", "asciidoctor/links_test_xref_using_macro_syntax.adoc", linksTestXrefUsingMacroSyntax},

	{"xref using macro syntax with explicit hash", "asciidoctor/links_test_xref_using_macro_syntax_with_explicit_hash.adoc", linksTestXrefUsingMacroSyntaxWithExplicitHash},

	{"xref using macro syntax inline with text", "asciidoctor/links_test_xref_using_macro_syntax_inline_with_text.adoc", linksTestXrefUsingMacroSyntaxInlineWithText},

	{"xref using macro syntax with multi-line label inline with text", "asciidoctor/links_test_xref_using_macro_syntax_with_multi_line_label_inline_with_text.adoc", linksTestXrefUsingMacroSyntaxWithMultiLineLabelInlineWithText},

	{"xref using macro syntax with text that ends with an escaped closing bracket", "asciidoctor/links_test_xref_using_macro_syntax_with_text_that_ends_with_an_escaped_closing_bracket.adoc", linksTestXrefUsingMacroSyntaxWithTextThatEndsWithAnEscapedClosingBracket},

	{"xref using macro syntax with text that contains an escaped closing bracket", "asciidoctor/links_test_xref_using_macro_syntax_with_text_that_contains_an_escaped_closing_bracket.adoc", linksTestXrefUsingMacroSyntaxWithTextThatContainsAnEscapedClosingBracket},

	{"unescapes square bracket in reftext used by xref", "asciidoctor/links_test_unescapes_square_bracket_in_reftext_used_by_xref.adoc", linksTestUnescapesSquareBracketInReftextUsedByXref},

	{"xref using invalid macro syntax does not create link", "asciidoctor/links_test_xref_using_invalid_macro_syntax_does_not_create_link.adoc", linksTestXrefUsingInvalidMacroSyntaxDoesNotCreateLink},

	{"should not warn if verbose flag is set and reference is found in compat mode", "asciidoctor/links_test_should_not_warn_if_verbose_flag_is_set_and_reference_is_found_in_compat_mode.adoc", linksTestShouldNotWarnIfVerboseFlagIsSetAndReferenceIsFoundInCompatMode},

	{"should warn and create link if verbose flag is set and reference using # notation is not found", "asciidoctor/links_test_should_warn_and_create_link_if_verbose_flag_is_set_and_reference_using_#_notation_is_not_found.adoc", linksTestShouldWarnAndCreateLinkIfVerboseFlagIsSetAndReferenceUsingNotationIsNotFound},

	{"should produce an internal anchor from an inter-document xref to file included into current file", "asciidoctor/links_test_should_produce_an_internal_anchor_from_an_inter_document_xref_to_file_included_into_current_file.adoc", linksTestShouldProduceAnInternalAnchorFromAnInterDocumentXrefToFileIncludedIntoCurrentFile},

	{"should produce an internal anchor from an inter-document xref to file included entirely into current file using tags", "asciidoctor/links_test_should_produce_an_internal_anchor_from_an_inter_document_xref_to_file_included_entirely_into_current_file_using_tags.adoc", linksTestShouldProduceAnInternalAnchorFromAnInterDocumentXrefToFileIncludedEntirelyIntoCurrentFileUsingTags},

	{"should not produce an internal anchor for inter-document xref to file partially included into current file", "asciidoctor/links_test_should_not_produce_an_internal_anchor_for_inter_document_xref_to_file_partially_included_into_current_file.adoc", linksTestShouldNotProduceAnInternalAnchorForInterDocumentXrefToFilePartiallyIncludedIntoCurrentFile},

	{"should produce an internal anchor for inter-document xref to file included fully and partially", "asciidoctor/links_test_should_produce_an_internal_anchor_for_inter_document_xref_to_file_included_fully_and_partially.adoc", linksTestShouldProduceAnInternalAnchorForInterDocumentXrefToFileIncludedFullyAndPartially},

	{"should warn and create link if debug mode is enabled, inter-document xref points to current doc, and reference not found", "asciidoctor/links_test_should_warn_and_create_link_if_debug_mode_is_enabled_inter_document_xref_points_to_current_doc_and_reference_not_found.adoc", linksTestShouldWarnAndCreateLinkIfDebugModeIsEnabledInterDocumentXrefPointsToCurrentDocAndReferenceNotFound},

	{"should use doctitle as fallback link text if inter-document xref points to current doc and no link text is provided", "asciidoctor/links_test_should_use_doctitle_as_fallback_link_text_if_inter_document_xref_points_to_current_doc_and_no_link_text_is_provided.adoc", linksTestShouldUseDoctitleAsFallbackLinkTextIfInterDocumentXrefPointsToCurrentDocAndNoLinkTextIsProvided},

	{"should use doctitle of root document as fallback link text for inter-document xref in AsciiDoc table cell that resolves to current doc", "asciidoctor/links_test_should_use_doctitle_of_root_document_as_fallback_link_text_for_inter_document_xref_in_ascii_doc_table_cell_that_resolves_to_current_doc.adoc", linksTestShouldUseDoctitleOfRootDocumentAsFallbackLinkTextForInterDocumentXrefInAsciiDocTableCellThatResolvesToCurrentDoc},

	{"should use reftext on document as fallback link text if inter-document xref points to current doc and no link text is provided", "asciidoctor/links_test_should_use_reftext_on_document_as_fallback_link_text_if_inter_document_xref_points_to_current_doc_and_no_link_text_is_provided.adoc", linksTestShouldUseReftextOnDocumentAsFallbackLinkTextIfInterDocumentXrefPointsToCurrentDocAndNoLinkTextIsProvided},

	{"should use reftext on document as fallback link text if xref points to empty fragment and no link text is provided", "asciidoctor/links_test_should_use_reftext_on_document_as_fallback_link_text_if_xref_points_to_empty_fragment_and_no_link_text_is_provided.adoc", linksTestShouldUseReftextOnDocumentAsFallbackLinkTextIfXrefPointsToEmptyFragmentAndNoLinkTextIsProvided},

	{"should use fallback link text if inter-document xref points to current doc without header and no link text is provided", "asciidoctor/links_test_should_use_fallback_link_text_if_inter_document_xref_points_to_current_doc_without_header_and_no_link_text_is_provided.adoc", linksTestShouldUseFallbackLinkTextIfInterDocumentXrefPointsToCurrentDocWithoutHeaderAndNoLinkTextIsProvided},

	{"should use fallback link text if fragment of internal xref is empty and no link text is provided", "asciidoctor/links_test_should_use_fallback_link_text_if_fragment_of_internal_xref_is_empty_and_no_link_text_is_provided.adoc", linksTestShouldUseFallbackLinkTextIfFragmentOfInternalXrefIsEmptyAndNoLinkTextIsProvided},

	{"should use document id as linkend for self xref in DocBook backend", "asciidoctor/links_test_should_use_document_id_as_linkend_for_self_xref_in_doc_book_backend.adoc", linksTestShouldUseDocumentIdAsLinkendForSelfXrefInDocBookBackend},

	{"should auto-generate document id to use as linkend for self xref in DocBook backend", "asciidoctor/links_test_should_auto_generate_document_id_to_use_as_linkend_for_self_xref_in_doc_book_backend.adoc", linksTestShouldAutoGenerateDocumentIdToUseAsLinkendForSelfXrefInDocBookBackend},

	{"should produce an internal anchor for inter-document xref to file outside of base directory", "asciidoctor/links_test_should_produce_an_internal_anchor_for_inter_document_xref_to_file_outside_of_base_directory.adoc", linksTestShouldProduceAnInternalAnchorForInterDocumentXrefToFileOutsideOfBaseDirectory},

	{"xref uses title of target as label for forward and backward references in html output", "asciidoctor/links_test_xref_uses_title_of_target_as_label_for_forward_and_backward_references_in_html_output.adoc", linksTestXrefUsesTitleOfTargetAsLabelForForwardAndBackwardReferencesInHtmlOutput},

	{"should not fail to resolve broken xref in title of block with ID", "asciidoctor/links_test_should_not_fail_to_resolve_broken_xref_in_title_of_block_with_id.adoc", linksTestShouldNotFailToResolveBrokenXrefInTitleOfBlockWithId},

	{"should resolve forward xref in title of block with ID", "asciidoctor/links_test_should_resolve_forward_xref_in_title_of_block_with_id.adoc", linksTestShouldResolveForwardXrefInTitleOfBlockWithId},

	{"should not fail to resolve broken xref in section title", "asciidoctor/links_test_should_not_fail_to_resolve_broken_xref_in_section_title.adoc", linksTestShouldNotFailToResolveBrokenXrefInSectionTitle},

	{"should break circular xref reference in section title", "asciidoctor/links_test_should_break_circular_xref_reference_in_section_title.adoc", linksTestShouldBreakCircularXrefReferenceInSectionTitle},

	{"should drop nested anchor in xreftext", "asciidoctor/links_test_should_drop_nested_anchor_in_xreftext.adoc", linksTestShouldDropNestedAnchorInXreftext},

	{"should not resolve forward xref evaluated during parsing", "asciidoctor/links_test_should_not_resolve_forward_xref_evaluated_during_parsing.adoc", linksTestShouldNotResolveForwardXrefEvaluatedDuringParsing},

	{"should not resolve forward natural xref evaluated during parsing", "asciidoctor/links_test_should_not_resolve_forward_natural_xref_evaluated_during_parsing.adoc", linksTestShouldNotResolveForwardNaturalXrefEvaluatedDuringParsing},

	{"should resolve first matching natural xref", "asciidoctor/links_test_should_resolve_first_matching_natural_xref.adoc", linksTestShouldResolveFirstMatchingNaturalXref},

	{"should not match numeric character references while searching for fragment in xref target", "asciidoctor/links_test_should_not_match_numeric_character_references_while_searching_for_fragment_in_xref_target.adoc", linksTestShouldNotMatchNumericCharacterReferencesWhileSearchingForFragmentInXrefTarget},

	{"should not match numeric character references in path of interdocument xref", "asciidoctor/links_test_should_not_match_numeric_character_references_in_path_of_interdocument_xref.adoc", linksTestShouldNotMatchNumericCharacterReferencesInPathOfInterdocumentXref},
}

var linksTestUnescapesSquareBracketInReftextOfAnchorMacro = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "see ",
		},
		&asciidoc.CrossReference{
			AttributeList: nil,
			Elements:      nil,
			ID: asciidoc.Elements{
				&asciidoc.String{
					Value: "foo",
				},
			},
			Format: 0,
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "anchor:foo[b[a\\]r]tex",
		},
		&asciidoc.NewLine{},
	},
}

var linksTestXrefUsingAngledBracketSyntaxWithLabel = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.CrossReference{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "About Tigers",
				},
			},
			ID: asciidoc.Elements{
				&asciidoc.String{
					Value: "tigers",
				},
			},
			Format: 0,
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: nil,
					ID: &asciidoc.ShorthandID{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "tigers",
							},
						},
					},
					Roles:   nil,
					Options: nil,
				},
			},
			Elements: nil,
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Tigers",
				},
			},
			Level: 1,
		},
	},
}

var linksTestXrefShouldUseTitleOfTargetAsLinkTextWhenNoExplicitReftextIsSpecified = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.CrossReference{
			AttributeList: nil,
			Elements:      nil,
			ID: asciidoc.Elements{
				&asciidoc.String{
					Value: "tigers",
				},
			},
			Format: 0,
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: nil,
					ID: &asciidoc.ShorthandID{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "tigers",
							},
						},
					},
					Roles:   nil,
					Options: nil,
				},
			},
			Elements: nil,
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Tigers",
				},
			},
			Level: 1,
		},
	},
}

var linksTestXrefShouldUseTitleOfTargetAsLinkTextWhenExplicitLinkTextIsEmpty = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "<<tigers,>>",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: nil,
					ID: &asciidoc.ShorthandID{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "tigers",
							},
						},
					},
					Roles:   nil,
					Options: nil,
				},
			},
			Elements: nil,
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Tigers",
				},
			},
			Level: 1,
		},
	},
}

var linksTestXrefUsingAngledBracketSyntaxWithQuotedLabel = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.CrossReference{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "\"About Tigers\"",
				},
			},
			ID: asciidoc.Elements{
				&asciidoc.String{
					Value: "tigers",
				},
			},
			Format: 0,
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: nil,
					ID: &asciidoc.ShorthandID{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "tigers",
							},
						},
					},
					Roles:   nil,
					Options: nil,
				},
			},
			Elements: nil,
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Tigers",
				},
			},
			Level: 1,
		},
	},
}

var linksTestXrefUsingAngledBracketSyntaxInlineWithText = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "Want to learn ",
		},
		&asciidoc.CrossReference{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "about tigers",
				},
			},
			ID: asciidoc.Elements{
				&asciidoc.String{
					Value: "tigers",
				},
			},
			Format: 0,
		},
		&asciidoc.String{
			Value: "?",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: nil,
					ID: &asciidoc.ShorthandID{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "tigers",
							},
						},
					},
					Roles:   nil,
					Options: nil,
				},
			},
			Elements: nil,
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Tigers",
				},
			},
			Level: 1,
		},
	},
}

var linksTestXrefUsingAngledBracketSyntaxWithMultiLineLabelInlineWithText = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "Want to learn ",
		},
		&asciidoc.CrossReference{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "about\ntigers",
				},
			},
			ID: asciidoc.Elements{
				&asciidoc.String{
					Value: "tigers",
				},
			},
			Format: 0,
		},
		&asciidoc.String{
			Value: "?",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: nil,
					ID: &asciidoc.ShorthandID{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "tigers",
							},
						},
					},
					Roles:   nil,
					Options: nil,
				},
			},
			Elements: nil,
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Tigers",
				},
			},
			Level: 1,
		},
	},
}

var linksTestXrefWithEscapedText = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "See the <<tigers, ",
		},
		&asciidoc.Monospace{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "+[tigers]+",
				},
			},
		},
		&asciidoc.String{
			Value: ">> section for details about tigers.",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: nil,
					ID: &asciidoc.ShorthandID{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "tigers",
							},
						},
					},
					Roles:   nil,
					Options: nil,
				},
			},
			Elements: nil,
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Tigers",
				},
			},
			Level: 1,
		},
	},
}

var linksTestXrefUsingMacroSyntax = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "This document has two sections, ",
		},
		&asciidoc.CrossReference{
			AttributeList: nil,
			Elements:      nil,
			ID: asciidoc.Elements{
				&asciidoc.String{
					Value: "sect-a",
				},
			},
			Format: 1,
		},
		&asciidoc.String{
			Value: " and ",
		},
		&asciidoc.CrossReference{
			AttributeList: nil,
			Elements:      nil,
			ID: asciidoc.Elements{
				&asciidoc.String{
					Value: "sect-b",
				},
			},
			Format: 1,
		},
		&asciidoc.String{
			Value: ".",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: nil,
					ID: &asciidoc.ShorthandID{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "sect-a",
							},
						},
					},
					Roles:   nil,
					Options: nil,
				},
			},
			Elements: asciidoc.Elements{
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
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: nil,
					ID: &asciidoc.ShorthandID{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "sect-b",
							},
						},
					},
					Roles:   nil,
					Options: nil,
				},
			},
			Elements: nil,
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Section B",
				},
			},
			Level: 1,
		},
	},
}

var linksTestXrefUsingMacroSyntaxWithExplicitHash = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.CrossReference{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "About Tigers",
				},
			},
			ID: asciidoc.Elements{
				&asciidoc.String{
					Value: "tigers",
				},
			},
			Format: 1,
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: nil,
					ID: &asciidoc.ShorthandID{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "tigers",
							},
						},
					},
					Roles:   nil,
					Options: nil,
				},
			},
			Elements: nil,
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Tigers",
				},
			},
			Level: 1,
		},
	},
}

var linksTestXrefUsingMacroSyntaxInlineWithText = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "Want to learn ",
		},
		&asciidoc.CrossReference{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "about tigers",
				},
			},
			ID: asciidoc.Elements{
				&asciidoc.String{
					Value: "tigers",
				},
			},
			Format: 1,
		},
		&asciidoc.String{
			Value: "?",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: nil,
					ID: &asciidoc.ShorthandID{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "tigers",
							},
						},
					},
					Roles:   nil,
					Options: nil,
				},
			},
			Elements: nil,
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Tigers",
				},
			},
			Level: 1,
		},
	},
}

var linksTestXrefUsingMacroSyntaxWithMultiLineLabelInlineWithText = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "Want to learn ",
		},
		&asciidoc.CrossReference{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "about",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "tigers",
				},
			},
			ID: asciidoc.Elements{
				&asciidoc.String{
					Value: "tigers",
				},
			},
			Format: 1,
		},
		&asciidoc.String{
			Value: "?",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: nil,
					ID: &asciidoc.ShorthandID{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "tigers",
							},
						},
					},
					Roles:   nil,
					Options: nil,
				},
			},
			Elements: nil,
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Tigers",
				},
			},
			Level: 1,
		},
	},
}

var linksTestXrefUsingMacroSyntaxWithTextThatEndsWithAnEscapedClosingBracket = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.CrossReference{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "[tigers\\]",
				},
			},
			ID: asciidoc.Elements{
				&asciidoc.String{
					Value: "tigers",
				},
			},
			Format: 1,
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: nil,
					ID: &asciidoc.ShorthandID{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "tigers",
							},
						},
					},
					Roles:   nil,
					Options: nil,
				},
			},
			Elements: nil,
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Tigers",
				},
			},
			Level: 1,
		},
	},
}

var linksTestXrefUsingMacroSyntaxWithTextThatContainsAnEscapedClosingBracket = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.CrossReference{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "[tigers\\] are cats",
				},
			},
			ID: asciidoc.Elements{
				&asciidoc.String{
					Value: "tigers",
				},
			},
			Format: 1,
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: nil,
					ID: &asciidoc.ShorthandID{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "tigers",
							},
						},
					},
					Roles:   nil,
					Options: nil,
				},
			},
			Elements: nil,
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Tigers",
				},
			},
			Level: 1,
		},
	},
}

var linksTestUnescapesSquareBracketInReftextUsedByXref = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "anchor:foo[b[a\\]r]about",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "see ",
		},
		&asciidoc.CrossReference{
			AttributeList: nil,
			Elements:      nil,
			ID: asciidoc.Elements{
				&asciidoc.String{
					Value: "foo",
				},
			},
			Format: 0,
		},
		&asciidoc.NewLine{},
	},
}

var linksTestXrefUsingInvalidMacroSyntaxDoesNotCreateLink = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: nil,
					ID: &asciidoc.ShorthandID{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "foobar",
							},
						},
					},
					Roles:   nil,
					Options: nil,
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.EmptyLine{
					Text: "",
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Foobar",
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
					Value: "See ",
				},
				&asciidoc.CrossReference{
					AttributeList: nil,
					Elements:      nil,
					ID: asciidoc.Elements{
						&asciidoc.String{
							Value: "foobaz",
						},
					},
					Format: 0,
				},
				&asciidoc.String{
					Value: ".",
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
}

var linksTestShouldNotWarnIfVerboseFlagIsSetAndReferenceIsFoundInCompatMode = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.AnchorAttribute{
					ID: asciidoc.Elements{
						&asciidoc.String{
							Value: "foobar",
						},
					},
					Label: nil,
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.EmptyLine{
					Text: "",
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Foobar",
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
					Value: "See ",
				},
				&asciidoc.CrossReference{
					AttributeList: nil,
					Elements:      nil,
					ID: asciidoc.Elements{
						&asciidoc.String{
							Value: "foobar",
						},
					},
					Format: 0,
				},
				&asciidoc.String{
					Value: ".",
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
}

var linksTestShouldWarnAndCreateLinkIfVerboseFlagIsSetAndReferenceUsingNotationIsNotFound = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: nil,
					ID: &asciidoc.ShorthandID{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "foobar",
							},
						},
					},
					Roles:   nil,
					Options: nil,
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.EmptyLine{
					Text: "",
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Foobar",
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
					Value: "See <<#foobaz>>.",
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
}

var linksTestShouldProduceAnInternalAnchorFromAnInterDocumentXrefToFileIncludedIntoCurrentFile = &asciidoc.Document{
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
				&asciidoc.Section{
					AttributeList: asciidoc.AttributeList{
						&asciidoc.ShorthandAttribute{
							Style: nil,
							ID: &asciidoc.ShorthandID{
								Elements: asciidoc.Elements{
									&asciidoc.String{
										Value: "ch1",
									},
								},
							},
							Roles:   nil,
							Options: nil,
						},
					},
					Elements: asciidoc.Elements{
						&asciidoc.EmptyLine{
							Text: "",
						},
						&asciidoc.String{
							Value: "So it begins.",
						},
						&asciidoc.NewLine{},
						&asciidoc.EmptyLine{
							Text: "",
						},
						&asciidoc.String{
							Value: "Read <<other-chapters.adoc#ch2>> to find out what happens next!",
						},
						&asciidoc.NewLine{},
						&asciidoc.EmptyLine{
							Text: "",
						},
						&asciidoc.FileInclude{
							AttributeList: nil,
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "other-chapters.adoc",
								},
							},
						},
					},
					Title: asciidoc.Elements{
						&asciidoc.String{
							Value: "Chapter 1",
						},
					},
					Level: 1,
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Book Title",
				},
			},
			Level: 0,
		},
	},
}

var linksTestShouldProduceAnInternalAnchorFromAnInterDocumentXrefToFileIncludedEntirelyIntoCurrentFileUsingTags = &asciidoc.Document{
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
				&asciidoc.Section{
					AttributeList: asciidoc.AttributeList{
						&asciidoc.ShorthandAttribute{
							Style: nil,
							ID: &asciidoc.ShorthandID{
								Elements: asciidoc.Elements{
									&asciidoc.String{
										Value: "ch1",
									},
								},
							},
							Roles:   nil,
							Options: nil,
						},
					},
					Elements: asciidoc.Elements{
						&asciidoc.EmptyLine{
							Text: "",
						},
						&asciidoc.String{
							Value: "So it begins.",
						},
						&asciidoc.NewLine{},
						&asciidoc.EmptyLine{
							Text: "",
						},
						&asciidoc.String{
							Value: "Read <<other-chapters.adoc#ch2>> to find out what happens next!",
						},
						&asciidoc.NewLine{},
						&asciidoc.EmptyLine{
							Text: "",
						},
						&asciidoc.FileInclude{
							AttributeList: asciidoc.AttributeList{
								&asciidoc.NamedAttribute{
									Name: "tags",
									Val: asciidoc.Elements{
										&asciidoc.String{
											Value: "**",
										},
									},
									Quote: 0,
								},
							},
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "other-chapters.adoc",
								},
							},
						},
					},
					Title: asciidoc.Elements{
						&asciidoc.String{
							Value: "Chapter 1",
						},
					},
					Level: 1,
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Book Title",
				},
			},
			Level: 0,
		},
	},
}

var linksTestShouldNotProduceAnInternalAnchorForInterDocumentXrefToFilePartiallyIncludedIntoCurrentFile = &asciidoc.Document{
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
				&asciidoc.Section{
					AttributeList: asciidoc.AttributeList{
						&asciidoc.ShorthandAttribute{
							Style: nil,
							ID: &asciidoc.ShorthandID{
								Elements: asciidoc.Elements{
									&asciidoc.String{
										Value: "ch1",
									},
								},
							},
							Roles:   nil,
							Options: nil,
						},
					},
					Elements: asciidoc.Elements{
						&asciidoc.EmptyLine{
							Text: "",
						},
						&asciidoc.String{
							Value: "So it begins.",
						},
						&asciidoc.NewLine{},
						&asciidoc.EmptyLine{
							Text: "",
						},
						&asciidoc.String{
							Value: "Read <<other-chapters.adoc#ch2,the next chapter>> to find out what happens next!",
						},
						&asciidoc.NewLine{},
						&asciidoc.EmptyLine{
							Text: "",
						},
						&asciidoc.FileInclude{
							AttributeList: asciidoc.AttributeList{
								&asciidoc.NamedAttribute{
									Name: "tags",
									Val: asciidoc.Elements{
										&asciidoc.String{
											Value: "ch2",
										},
									},
									Quote: 0,
								},
							},
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "other-chapters.adoc",
								},
							},
						},
					},
					Title: asciidoc.Elements{
						&asciidoc.String{
							Value: "Chapter 1",
						},
					},
					Level: 1,
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Book Title",
				},
			},
			Level: 0,
		},
	},
}

var linksTestShouldProduceAnInternalAnchorForInterDocumentXrefToFileIncludedFullyAndPartially = &asciidoc.Document{
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
				&asciidoc.Section{
					AttributeList: asciidoc.AttributeList{
						&asciidoc.ShorthandAttribute{
							Style: nil,
							ID: &asciidoc.ShorthandID{
								Elements: asciidoc.Elements{
									&asciidoc.String{
										Value: "ch1",
									},
								},
							},
							Roles:   nil,
							Options: nil,
						},
					},
					Elements: asciidoc.Elements{
						&asciidoc.EmptyLine{
							Text: "",
						},
						&asciidoc.String{
							Value: "So it begins.",
						},
						&asciidoc.NewLine{},
						&asciidoc.EmptyLine{
							Text: "",
						},
						&asciidoc.String{
							Value: "Read <<other-chapters.adoc#ch2,the next chapter>> to find out what happens next!",
						},
						&asciidoc.NewLine{},
						&asciidoc.EmptyLine{
							Text: "",
						},
						&asciidoc.FileInclude{
							AttributeList: nil,
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "other-chapters.adoc",
								},
							},
						},
						&asciidoc.EmptyLine{
							Text: "",
						},
						&asciidoc.FileInclude{
							AttributeList: asciidoc.AttributeList{
								&asciidoc.NamedAttribute{
									Name: "tag",
									Val: asciidoc.Elements{
										&asciidoc.String{
											Value: "ch2-noid",
										},
									},
									Quote: 0,
								},
							},
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "other-chapters.adoc",
								},
							},
						},
					},
					Title: asciidoc.Elements{
						&asciidoc.String{
							Value: "Chapter 1",
						},
					},
					Level: 1,
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Book Title",
				},
			},
			Level: 0,
		},
	},
}

var linksTestShouldWarnAndCreateLinkIfDebugModeIsEnabledInterDocumentXrefPointsToCurrentDocAndReferenceNotFound = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: nil,
					ID: &asciidoc.ShorthandID{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "foobar",
							},
						},
					},
					Roles:   nil,
					Options: nil,
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.EmptyLine{
					Text: "",
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Foobar",
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
					Value: "See <<test.adoc#foobaz>>.",
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
}

var linksTestShouldUseDoctitleAsFallbackLinkTextIfInterDocumentXrefPointsToCurrentDocAndNoLinkTextIsProvided = &asciidoc.Document{
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
					Value: "See ",
				},
				&asciidoc.CrossReference{
					AttributeList: nil,
					Elements:      nil,
					ID: asciidoc.Elements{
						&asciidoc.String{
							Value: "test.adoc",
						},
					},
					Format: 1,
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Links & Stuff at https://example.org",
				},
			},
			Level: 0,
		},
	},
}

var linksTestShouldUseDoctitleOfRootDocumentAsFallbackLinkTextForInterDocumentXrefInAsciiDocTableCellThatResolvesToCurrentDoc = &asciidoc.Document{
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
										&asciidoc.String{
											Value: "See ",
										},
										&asciidoc.CrossReference{
											AttributeList: nil,
											Elements:      nil,
											ID: asciidoc.Elements{
												&asciidoc.String{
													Value: "test.adoc",
												},
											},
											Format: 1,
										},
									},
									Blank: false,
								},
							},
						},
					},
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

var linksTestShouldUseReftextOnDocumentAsFallbackLinkTextIfInterDocumentXrefPointsToCurrentDocAndNoLinkTextIsProvided = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.NamedAttribute{
					Name: "reftext",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Links and Stuff",
						},
					},
					Quote: 2,
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "See ",
				},
				&asciidoc.CrossReference{
					AttributeList: nil,
					Elements:      nil,
					ID: asciidoc.Elements{
						&asciidoc.String{
							Value: "test.adoc",
						},
					},
					Format: 1,
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Links & Stuff",
				},
			},
			Level: 0,
		},
	},
}

var linksTestShouldUseReftextOnDocumentAsFallbackLinkTextIfXrefPointsToEmptyFragmentAndNoLinkTextIsProvided = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.NamedAttribute{
					Name: "reftext",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Links and Stuff",
						},
					},
					Quote: 2,
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "See ",
				},
				&asciidoc.CrossReference{
					AttributeList: nil,
					Elements:      nil,
					ID: asciidoc.Elements{
						&asciidoc.String{
							Value: "#",
						},
					},
					Format: 1,
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Links & Stuff",
				},
			},
			Level: 0,
		},
	},
}

var linksTestShouldUseFallbackLinkTextIfInterDocumentXrefPointsToCurrentDocWithoutHeaderAndNoLinkTextIsProvided = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "See ",
		},
		&asciidoc.CrossReference{
			AttributeList: nil,
			Elements:      nil,
			ID: asciidoc.Elements{
				&asciidoc.String{
					Value: "test.adoc",
				},
			},
			Format: 1,
		},
		&asciidoc.NewLine{},
	},
}

var linksTestShouldUseFallbackLinkTextIfFragmentOfInternalXrefIsEmptyAndNoLinkTextIsProvided = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "See ",
		},
		&asciidoc.CrossReference{
			AttributeList: nil,
			Elements:      nil,
			ID: asciidoc.Elements{
				&asciidoc.String{
					Value: "#",
				},
			},
			Format: 1,
		},
		&asciidoc.NewLine{},
	},
}

var linksTestShouldUseDocumentIdAsLinkendForSelfXrefInDocBookBackend = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: nil,
					ID: &asciidoc.ShorthandID{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "docid",
							},
						},
					},
					Roles:   nil,
					Options: nil,
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "See ",
				},
				&asciidoc.CrossReference{
					AttributeList: nil,
					Elements:      nil,
					ID: asciidoc.Elements{
						&asciidoc.String{
							Value: "test.adoc",
						},
					},
					Format: 1,
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

var linksTestShouldAutoGenerateDocumentIdToUseAsLinkendForSelfXrefInDocBookBackend = &asciidoc.Document{
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
					Value: "See ",
				},
				&asciidoc.CrossReference{
					AttributeList: nil,
					Elements:      nil,
					ID: asciidoc.Elements{
						&asciidoc.String{
							Value: "test.adoc",
						},
					},
					Format: 1,
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

var linksTestShouldProduceAnInternalAnchorForInterDocumentXrefToFileOutsideOfBaseDirectory = &asciidoc.Document{
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
					Value: "See <<../section-a.adoc#section-a>>.",
				},
				&asciidoc.NewLine{},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.FileInclude{
					AttributeList: nil,
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "../section-a.adoc",
						},
					},
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

var linksTestXrefUsesTitleOfTargetAsLabelForForwardAndBackwardReferencesInHtmlOutput = &asciidoc.Document{
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
				&asciidoc.CrossReference{
					AttributeList: nil,
					Elements:      nil,
					ID: asciidoc.Elements{
						&asciidoc.String{
							Value: "_section_b",
						},
					},
					Format: 0,
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
				&asciidoc.CrossReference{
					AttributeList: nil,
					Elements:      nil,
					ID: asciidoc.Elements{
						&asciidoc.String{
							Value: "_section_a",
						},
					},
					Format: 0,
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
}

var linksTestShouldNotFailToResolveBrokenXrefInTitleOfBlockWithId = &asciidoc.Document{
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
								Value: "p1",
							},
						},
					},
					Roles:   nil,
					Options: nil,
				},
				&asciidoc.TitleAttribute{
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "<<DNE>>",
						},
					},
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "paragraph text",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 0,
		},
	},
}

var linksTestShouldResolveForwardXrefInTitleOfBlockWithId = &asciidoc.Document{
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
								Value: "p1",
							},
						},
					},
					Roles:   nil,
					Options: nil,
				},
				&asciidoc.TitleAttribute{
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "<<conclusion>>",
						},
					},
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "paragraph text",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 0,
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: nil,
					ID: &asciidoc.ShorthandID{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "conclusion",
							},
						},
					},
					Roles:   nil,
					Options: nil,
				},
			},
			Elements: nil,
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Conclusion",
				},
			},
			Level: 1,
		},
	},
}

var linksTestShouldNotFailToResolveBrokenXrefInSectionTitle = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: nil,
					ID: &asciidoc.ShorthandID{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "s1",
							},
						},
					},
					Roles:   nil,
					Options: nil,
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.EmptyLine{
					Text: "",
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "<<DNE>>",
				},
			},
			Level: 1,
		},
		&asciidoc.Section{
			AttributeList: nil,
			Elements:      nil,
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "<<s1>>",
				},
			},
			Level: 1,
		},
	},
}

var linksTestShouldBreakCircularXrefReferenceInSectionTitle = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: nil,
					ID: &asciidoc.ShorthandID{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "a",
							},
						},
					},
					Roles:   nil,
					Options: nil,
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.EmptyLine{
					Text: "",
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "A <<b>>",
				},
			},
			Level: 1,
		},
		&asciidoc.Section{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: nil,
					ID: &asciidoc.ShorthandID{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "b",
							},
						},
					},
					Roles:   nil,
					Options: nil,
				},
			},
			Elements: nil,
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "B <<a>>",
				},
			},
			Level: 1,
		},
	},
}

var linksTestShouldDropNestedAnchorInXreftext = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: nil,
					ID: &asciidoc.ShorthandID{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "a",
							},
						},
					},
					Roles:   nil,
					Options: nil,
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.EmptyLine{
					Text: "",
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "See <<b>>",
				},
			},
			Level: 1,
		},
		&asciidoc.Section{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: nil,
					ID: &asciidoc.ShorthandID{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "b",
							},
						},
					},
					Roles:   nil,
					Options: nil,
				},
			},
			Elements: nil,
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Consult https://google.com[Google]",
				},
			},
			Level: 1,
		},
	},
}

var linksTestShouldNotResolveForwardXrefEvaluatedDuringParsing = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: nil,
					ID: &asciidoc.ShorthandID{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "s1",
							},
						},
					},
					Roles:   nil,
					Options: nil,
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.EmptyLine{
					Text: "",
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "<<forward>>",
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
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "<<s1>>",
				},
			},
			Level: 1,
		},
		&asciidoc.Section{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: nil,
					ID: &asciidoc.ShorthandID{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "forward",
							},
						},
					},
					Roles:   nil,
					Options: nil,
				},
			},
			Elements: nil,
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Forward",
				},
			},
			Level: 1,
		},
	},
}

var linksTestShouldNotResolveForwardNaturalXrefEvaluatedDuringParsing = &asciidoc.Document{
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
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: nil,
					ID: &asciidoc.ShorthandID{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "s1",
							},
						},
					},
					Roles:   nil,
					Options: nil,
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.EmptyLine{
					Text: "",
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "<<Forward>>",
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
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "<<s1>>",
				},
			},
			Level: 1,
		},
		&asciidoc.Section{
			AttributeList: nil,
			Elements:      nil,
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Forward",
				},
			},
			Level: 1,
		},
	},
}

var linksTestShouldResolveFirstMatchingNaturalXref = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "see ",
		},
		&asciidoc.CrossReference{
			AttributeList: nil,
			Elements:      nil,
			ID: asciidoc.Elements{
				&asciidoc.String{
					Value: "Section Title",
				},
			},
			Format: 0,
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: nil,
					ID: &asciidoc.ShorthandID{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "s1",
							},
						},
					},
					Roles:   nil,
					Options: nil,
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.EmptyLine{
					Text: "",
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Section Title",
				},
			},
			Level: 1,
		},
		&asciidoc.Section{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: nil,
					ID: &asciidoc.ShorthandID{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "s2",
							},
						},
					},
					Roles:   nil,
					Options: nil,
				},
			},
			Elements: nil,
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Section Title",
				},
			},
			Level: 1,
		},
	},
}

var linksTestShouldNotMatchNumericCharacterReferencesWhileSearchingForFragmentInXrefTarget = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "see <<Cub => Tiger>>",
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
					Value: "Cub => Tiger",
				},
			},
			Level: 1,
		},
	},
}

var linksTestShouldNotMatchNumericCharacterReferencesInPathOfInterdocumentXref = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "see ",
		},
		&asciidoc.CrossReference{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.CharacterReplacementReference{
					Value: "cpp",
				},
			},
			ID: asciidoc.Elements{
				&asciidoc.CharacterReplacementReference{
					Value: "cpp",
				},
			},
			Format: 1,
		},
		&asciidoc.String{
			Value: ".",
		},
		&asciidoc.NewLine{},
	},
}
