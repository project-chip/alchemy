package tests

import (
	"testing"

	"github.com/project-chip/alchemy/asciidoc"
)

func TestBlocks(t *testing.T) {
	blocksTests.run(t)
}

var blocksTests = parseTests{

	{"horizontal rule between blocks", "asciidoctor/blocks_test_horizontal_rule_between_blocks.adoc", blocksTestHorizontalRuleBetweenBlocks},

	{"line comment between paragraphs offset by blank lines", "asciidoctor/blocks_test_line_comment_between_paragraphs_offset_by_blank_lines.adoc", blocksTestLineCommentBetweenParagraphsOffsetByBlankLines},

	{"adjacent line comment between paragraphs", "asciidoctor/blocks_test_adjacent_line_comment_between_paragraphs.adoc", blocksTestAdjacentLineCommentBetweenParagraphs},

	{"comment block between paragraphs offset by blank lines", "asciidoctor/blocks_test_comment_block_between_paragraphs_offset_by_blank_lines.adoc", blocksTestCommentBlockBetweenParagraphsOffsetByBlankLines},

	{"comment block between paragraphs offset by blank lines inside delimited block", "asciidoctor/blocks_test_comment_block_between_paragraphs_offset_by_blank_lines_inside_delimited_block.adoc", blocksTestCommentBlockBetweenParagraphsOffsetByBlankLinesInsideDelimitedBlock},

	{"adjacent comment block between paragraphs", "asciidoctor/blocks_test_adjacent_comment_block_between_paragraphs.adoc", blocksTestAdjacentCommentBlockBetweenParagraphs},

	{"can convert with block comment at end of document with trailing newlines", "asciidoctor/blocks_test_can_convert_with_block_comment_at_end_of_document_with_trailing_newlines.adoc", blocksTestCanConvertWithBlockCommentAtEndOfDocumentWithTrailingNewlines},

	{"trailing newlines after block comment at end of document does not create paragraph", "asciidoctor/blocks_test_trailing_newlines_after_block_comment_at_end_of_document_does_not_create_paragraph.adoc", blocksTestTrailingNewlinesAfterBlockCommentAtEndOfDocumentDoesNotCreateParagraph},

	{"line starting with three slashes should not be line comment", "asciidoctor/blocks_test_line_starting_with_three_slashes_should_not_be_line_comment.adoc", blocksTestLineStartingWithThreeSlashesShouldNotBeLineComment},

	{"preprocessor directives should not be processed within comment block", "asciidoctor/blocks_test_preprocessor_directives_should_not_be_processed_within_comment_block.adoc", blocksTestPreprocessorDirectivesShouldNotBeProcessedWithinCommentBlock},

	{"should warn if unterminated comment block is detected in body", "asciidoctor/blocks_test_should_warn_if_unterminated_comment_block_is_detected_in_body.adoc", blocksTestShouldWarnIfUnterminatedCommentBlockIsDetectedInBody},

	{"should warn if unterminated comment block is detected inside another block", "asciidoctor/blocks_test_should_warn_if_unterminated_comment_block_is_detected_inside_another_block.adoc", blocksTestShouldWarnIfUnterminatedCommentBlockIsDetectedInsideAnotherBlock},

	{"preprocessor directives should not be processed within comment open block", "asciidoctor/blocks_test_preprocessor_directives_should_not_be_processed_within_comment_open_block.adoc", blocksTestPreprocessorDirectivesShouldNotBeProcessedWithinCommentOpenBlock},

	{"preprocessor directives should not be processed on subsequent lines of a comment paragraph", "asciidoctor/blocks_test_preprocessor_directives_should_not_be_processed_on_subsequent_lines_of_a_comment_paragraph.adoc", blocksTestPreprocessorDirectivesShouldNotBeProcessedOnSubsequentLinesOfACommentParagraph},

	{"comment style on open block should only skip block", "asciidoctor/blocks_test_comment_style_on_open_block_should_only_skip_block.adoc", blocksTestCommentStyleOnOpenBlockShouldOnlySkipBlock},

	{"comment style on paragraph should only skip paragraph", "asciidoctor/blocks_test_comment_style_on_paragraph_should_only_skip_paragraph.adoc", blocksTestCommentStyleOnParagraphShouldOnlySkipParagraph},

	{"comment style on paragraph should not cause adjacent block to be skipped", "asciidoctor/blocks_test_comment_style_on_paragraph_should_not_cause_adjacent_block_to_be_skipped.adoc", blocksTestCommentStyleOnParagraphShouldNotCauseAdjacentBlockToBeSkipped},

	{"should not drop content that follows skipped content inside a delimited block", "asciidoctor/blocks_test_should_not_drop_content_that_follows_skipped_content_inside_a_delimited_block.adoc", blocksTestShouldNotDropContentThatFollowsSkippedContentInsideADelimitedBlock},

	{"should parse sidebar block", "asciidoctor/blocks_test_should_parse_sidebar_block.adoc", blocksTestShouldParseSidebarBlock},

	{"quote block with no attribution", "asciidoctor/blocks_test_quote_block_with_no_attribution.adoc", blocksTestQuoteBlockWithNoAttribution},

	{"quote block with attribution", "asciidoctor/blocks_test_quote_block_with_attribution.adoc", blocksTestQuoteBlockWithAttribution},

	{"quote block with attribute and id and role shorthand", "asciidoctor/blocks_test_quote_block_with_attribute_and_id_and_role_shorthand.adoc", blocksTestQuoteBlockWithAttributeAndIdAndRoleShorthand},

	{"setting ID using style shorthand should not reset block style", "asciidoctor/blocks_test_setting_id_using_style_shorthand_should_not_reset_block_style.adoc", blocksTestSettingIdUsingStyleShorthandShouldNotResetBlockStyle},

	{"quote block with complex content", "asciidoctor/blocks_test_quote_block_with_complex_content.adoc", blocksTestQuoteBlockWithComplexContent},

	{"quote block with attribution converted to DocBook", "asciidoctor/blocks_test_quote_block_with_attribution_converted_to_doc_book.adoc", blocksTestQuoteBlockWithAttributionConvertedToDocBook},

	{"epigraph quote block with attribution converted to DocBook", "asciidoctor/blocks_test_epigraph_quote_block_with_attribution_converted_to_doc_book.adoc", blocksTestEpigraphQuoteBlockWithAttributionConvertedToDocBook},

	{"markdown-style quote block with single paragraph and no attribution", "asciidoctor/blocks_test_markdown_style_quote_block_with_single_paragraph_and_no_attribution.adoc", blocksTestMarkdownStyleQuoteBlockWithSingleParagraphAndNoAttribution},

	{"lazy markdown-style quote block with single paragraph and no attribution", "asciidoctor/blocks_test_lazy_markdown_style_quote_block_with_single_paragraph_and_no_attribution.adoc", blocksTestLazyMarkdownStyleQuoteBlockWithSingleParagraphAndNoAttribution},

	{"markdown-style quote block with multiple paragraphs and no attribution", "asciidoctor/blocks_test_markdown_style_quote_block_with_multiple_paragraphs_and_no_attribution.adoc", blocksTestMarkdownStyleQuoteBlockWithMultipleParagraphsAndNoAttribution},

	{"markdown-style quote block with multiple blocks and no attribution", "asciidoctor/blocks_test_markdown_style_quote_block_with_multiple_blocks_and_no_attribution.adoc", blocksTestMarkdownStyleQuoteBlockWithMultipleBlocksAndNoAttribution},

	{"markdown-style quote block with single paragraph and attribution", "asciidoctor/blocks_test_markdown_style_quote_block_with_single_paragraph_and_attribution.adoc", blocksTestMarkdownStyleQuoteBlockWithSingleParagraphAndAttribution},

	{"markdown-style quote block with only attribution", "asciidoctor/blocks_test_markdown_style_quote_block_with_only_attribution.adoc", blocksTestMarkdownStyleQuoteBlockWithOnlyAttribution},

	{"quoted paragraph-style quote block with attribution", "asciidoctor/blocks_test_quoted_paragraph_style_quote_block_with_attribution.adoc", blocksTestQuotedParagraphStyleQuoteBlockWithAttribution},

	{"should parse credit line in quoted paragraph-style quote block like positional block attributes", "asciidoctor/blocks_test_should_parse_credit_line_in_quoted_paragraph_style_quote_block_like_positional_block_attributes.adoc", blocksTestShouldParseCreditLineInQuotedParagraphStyleQuoteBlockLikePositionalBlockAttributes},

	{"single-line verse block without attribution", "asciidoctor/blocks_test_single_line_verse_block_without_attribution.adoc", blocksTestSingleLineVerseBlockWithoutAttribution},

	{"single-line verse block with attribution", "asciidoctor/blocks_test_single_line_verse_block_with_attribution.adoc", blocksTestSingleLineVerseBlockWithAttribution},

	{"single-line verse block with attribution converted to DocBook", "asciidoctor/blocks_test_single_line_verse_block_with_attribution_converted_to_doc_book.adoc", blocksTestSingleLineVerseBlockWithAttributionConvertedToDocBook},

	{"single-line epigraph verse block with attribution converted to DocBook", "asciidoctor/blocks_test_single_line_epigraph_verse_block_with_attribution_converted_to_doc_book.adoc", blocksTestSingleLineEpigraphVerseBlockWithAttributionConvertedToDocBook},

	{"multi-stanza verse block", "asciidoctor/blocks_test_multi_stanza_verse_block.adoc", blocksTestMultiStanzaVerseBlock},

	{"verse block does not contain block elements", "asciidoctor/blocks_test_verse_block_does_not_contain_block_elements.adoc", blocksTestVerseBlockDoesNotContainBlockElements},

	{"verse should have normal subs", "asciidoctor/blocks_test_verse_should_have_normal_subs.adoc", blocksTestVerseShouldHaveNormalSubs},

	{"should not recognize callouts in a verse", "asciidoctor/blocks_test_should_not_recognize_callouts_in_a_verse.adoc", blocksTestShouldNotRecognizeCalloutsInAVerse},

	{"should perform normal subs on a verse block", "asciidoctor/blocks_test_should_perform_normal_subs_on_a_verse_block.adoc", blocksTestShouldPerformNormalSubsOnAVerseBlock},

	{"can convert example block", "asciidoctor/blocks_test_can_convert_example_block.adoc", blocksTestCanConvertExampleBlock},

	{"assigns sequential numbered caption to example block with title", "asciidoctor/blocks_test_assigns_sequential_numbered_caption_to_example_block_with_title.adoc", blocksTestAssignsSequentialNumberedCaptionToExampleBlockWithTitle},

	{"assigns sequential character caption to example block with title", "asciidoctor/blocks_test_assigns_sequential_character_caption_to_example_block_with_title.adoc", blocksTestAssignsSequentialCharacterCaptionToExampleBlockWithTitle},

	{"should increment counter for example even when example-number is locked by the API", "asciidoctor/blocks_test_should_increment_counter_for_example_even_when_example_number_is_locked_by_the_api.adoc", blocksTestShouldIncrementCounterForExampleEvenWhenExampleNumberIsLockedByTheApi},

	{"should use explicit caption if specified", "asciidoctor/blocks_test_should_use_explicit_caption_if_specified.adoc", blocksTestShouldUseExplicitCaptionIfSpecified},

	{"automatic caption can be turned off and on and modified", "asciidoctor/blocks_test_automatic_caption_can_be_turned_off_and_on_and_modified.adoc", blocksTestAutomaticCaptionCanBeTurnedOffAndOnAndModified},

	{"should use explicit caption if specified even if block-specific global caption is disabled", "asciidoctor/blocks_test_should_use_explicit_caption_if_specified_even_if_block_specific_global_caption_is_disabled.adoc", blocksTestShouldUseExplicitCaptionIfSpecifiedEvenIfBlockSpecificGlobalCaptionIsDisabled},

	{"should use global caption if specified even if block-specific global caption is disabled", "asciidoctor/blocks_test_should_use_global_caption_if_specified_even_if_block_specific_global_caption_is_disabled.adoc", blocksTestShouldUseGlobalCaptionIfSpecifiedEvenIfBlockSpecificGlobalCaptionIsDisabled},

	{"should not process caption attribute on block that does not support a caption", "asciidoctor/blocks_test_should_not_process_caption_attribute_on_block_that_does_not_support_a_caption.adoc", blocksTestShouldNotProcessCaptionAttributeOnBlockThatDoesNotSupportACaption},

	{"should create details/summary set if collapsible option is set", "asciidoctor/blocks_test_should_create_details_summary_set_if_collapsible_option_is_set.adoc", blocksTestShouldCreateDetailsSummarySetIfCollapsibleOptionIsSet},

	{"should open details/summary set if collapsible and open options are set", "asciidoctor/blocks_test_should_open_details_summary_set_if_collapsible_and_open_options_are_set.adoc", blocksTestShouldOpenDetailsSummarySetIfCollapsibleAndOpenOptionsAreSet},

	{"should add default summary element if collapsible option is set and title is not specifed", "asciidoctor/blocks_test_should_add_default_summary_element_if_collapsible_option_is_set_and_title_is_not_specifed.adoc", blocksTestShouldAddDefaultSummaryElementIfCollapsibleOptionIsSetAndTitleIsNotSpecifed},

	{"should not allow collapsible block to increment example number", "asciidoctor/blocks_test_should_not_allow_collapsible_block_to_increment_example_number.adoc", blocksTestShouldNotAllowCollapsibleBlockToIncrementExampleNumber},

	{"should warn if example block is not terminated", "asciidoctor/blocks_test_should_warn_if_example_block_is_not_terminated.adoc", blocksTestShouldWarnIfExampleBlockIsNotTerminated},

	{"caption block-level attribute should be used as caption", "asciidoctor/blocks_test_caption_block_level_attribute_should_be_used_as_caption.adoc", blocksTestCaptionBlockLevelAttributeShouldBeUsedAsCaption},

	{"can override caption of admonition block using document attribute", "asciidoctor/blocks_test_can_override_caption_of_admonition_block_using_document_attribute.adoc", blocksTestCanOverrideCaptionOfAdmonitionBlockUsingDocumentAttribute},

	{"blank caption document attribute should not blank admonition block caption", "asciidoctor/blocks_test_blank_caption_document_attribute_should_not_blank_admonition_block_caption.adoc", blocksTestBlankCaptionDocumentAttributeShouldNotBlankAdmonitionBlockCaption},

	{"should separate adjacent paragraphs and listing into blocks", "asciidoctor/blocks_test_should_separate_adjacent_paragraphs_and_listing_into_blocks.adoc", blocksTestShouldSeparateAdjacentParagraphsAndListingIntoBlocks},

	{"should warn if listing block is not terminated", "asciidoctor/blocks_test_should_warn_if_listing_block_is_not_terminated.adoc", blocksTestShouldWarnIfListingBlockIsNotTerminated},

	{"should not crash when converting verbatim block that has no lines", "asciidoctor/blocks_test_should_not_crash_when_converting_verbatim_block_that_has_no_lines.adoc", blocksTestShouldNotCrashWhenConvertingVerbatimBlockThatHasNoLines},

	{"should preserve newlines in listing block", "asciidoctor/blocks_test_should_preserve_newlines_in_listing_block.adoc", blocksTestShouldPreserveNewlinesInListingBlock},

	{"should preserve newlines in verse block", "asciidoctor/blocks_test_should_preserve_newlines_in_verse_block.adoc", blocksTestShouldPreserveNewlinesInVerseBlock},

	{"should strip leading and trailing blank lines when converting verbatim block", "asciidoctor/blocks_test_should_strip_leading_and_trailing_blank_lines_when_converting_verbatim_block.adoc", blocksTestShouldStripLeadingAndTrailingBlankLinesWhenConvertingVerbatimBlock},

	{"should remove block indent if indent attribute is 0", "asciidoctor/blocks_test_should_remove_block_indent_if_indent_attribute_is_0.adoc", blocksTestShouldRemoveBlockIndentIfIndentAttributeIs0},

	{"should not remove block indent if indent attribute is -1", "asciidoctor/blocks_test_should_not_remove_block_indent_if_indent_attribute_is__1.adoc", blocksTestShouldNotRemoveBlockIndentIfIndentAttributeIs1},

	{"should set block indent to value specified by indent attribute", "asciidoctor/blocks_test_should_set_block_indent_to_value_specified_by_indent_attribute.adoc", blocksTestShouldSetBlockIndentToValueSpecifiedByIndentAttribute},

	{"should set block indent to value specified by indent document attribute", "asciidoctor/blocks_test_should_set_block_indent_to_value_specified_by_indent_document_attribute.adoc", blocksTestShouldSetBlockIndentToValueSpecifiedByIndentDocumentAttribute},

	{"literal block should honor nowrap option", "asciidoctor/blocks_test_literal_block_should_honor_nowrap_option.adoc", blocksTestLiteralBlockShouldHonorNowrapOption},

	{"literal block should set nowrap class if prewrap document attribute is disabled", "asciidoctor/blocks_test_literal_block_should_set_nowrap_class_if_prewrap_document_attribute_is_disabled.adoc", blocksTestLiteralBlockShouldSetNowrapClassIfPrewrapDocumentAttributeIsDisabled},

	{"should preserve guard in front of callout if icons are not enabled", "asciidoctor/blocks_test_should_preserve_guard_in_front_of_callout_if_icons_are_not_enabled.adoc", blocksTestShouldPreserveGuardInFrontOfCalloutIfIconsAreNotEnabled},

	{"should preserve guard around callout if icons are not enabled", "asciidoctor/blocks_test_should_preserve_guard_around_callout_if_icons_are_not_enabled.adoc", blocksTestShouldPreserveGuardAroundCalloutIfIconsAreNotEnabled},

	{"literal block should honor explicit subs list", "asciidoctor/blocks_test_literal_block_should_honor_explicit_subs_list.adoc", blocksTestLiteralBlockShouldHonorExplicitSubsList},

	{"should be able to disable callouts for literal block", "asciidoctor/blocks_test_should_be_able_to_disable_callouts_for_literal_block.adoc", blocksTestShouldBeAbleToDisableCalloutsForLiteralBlock},

	{"listing block should honor explicit subs list", "asciidoctor/blocks_test_listing_block_should_honor_explicit_subs_list.adoc", blocksTestListingBlockShouldHonorExplicitSubsList},

	{"should not mangle array that contains formatted text with role in listing block with quotes sub enabled", "asciidoctor/blocks_test_should_not_mangle_array_that_contains_formatted_text_with_role_in_listing_block_with_quotes_sub_enabled.adoc", blocksTestShouldNotMangleArrayThatContainsFormattedTextWithRoleInListingBlockWithQuotesSubEnabled},

	{"first character of block title may be a period if not followed by space", "asciidoctor/blocks_test_first_character_of_block_title_may_be_a_period_if_not_followed_by_space.adoc", blocksTestFirstCharacterOfBlockTitleMayBeAPeriodIfNotFollowedBySpace},

	{"listing block without title should generate screen element in docbook", "asciidoctor/blocks_test_listing_block_without_title_should_generate_screen_element_in_docbook.adoc", blocksTestListingBlockWithoutTitleShouldGenerateScreenElementInDocbook},

	{"listing block with title should generate screen element inside formalpara element in docbook", "asciidoctor/blocks_test_listing_block_with_title_should_generate_screen_element_inside_formalpara_element_in_docbook.adoc", blocksTestListingBlockWithTitleShouldGenerateScreenElementInsideFormalparaElementInDocbook},

	{"should not prepend caption to title of listing block with title if listing-caption attribute is not set", "asciidoctor/blocks_test_should_not_prepend_caption_to_title_of_listing_block_with_title_if_listing_caption_attribute_is_not_set.adoc", blocksTestShouldNotPrependCaptionToTitleOfListingBlockWithTitleIfListingCaptionAttributeIsNotSet},

	{"should prepend caption specified by listing-caption attribute and number to title of listing block with title", "asciidoctor/blocks_test_should_prepend_caption_specified_by_listing_caption_attribute_and_number_to_title_of_listing_block_with_title.adoc", blocksTestShouldPrependCaptionSpecifiedByListingCaptionAttributeAndNumberToTitleOfListingBlockWithTitle},

	{"should prepend caption specified by caption attribute on listing block even if listing-caption attribute is not set", "asciidoctor/blocks_test_should_prepend_caption_specified_by_caption_attribute_on_listing_block_even_if_listing_caption_attribute_is_not_set.adoc", blocksTestShouldPrependCaptionSpecifiedByCaptionAttributeOnListingBlockEvenIfListingCaptionAttributeIsNotSet},

	{"listing block without an explicit style and with a second positional argument should be promoted to a source block", "asciidoctor/blocks_test_listing_block_without_an_explicit_style_and_with_a_second_positional_argument_should_be_promoted_to_a_source_block.adoc", blocksTestListingBlockWithoutAnExplicitStyleAndWithASecondPositionalArgumentShouldBePromotedToASourceBlock},

	{"listing block without an explicit style should be promoted to a source block if source-language is set", "asciidoctor/blocks_test_listing_block_without_an_explicit_style_should_be_promoted_to_a_source_block_if_source_language_is_set.adoc", blocksTestListingBlockWithoutAnExplicitStyleShouldBePromotedToASourceBlockIfSourceLanguageIsSet},

	{"listing block with an explicit style and a second positional argument should not be promoted to a source block", "asciidoctor/blocks_test_listing_block_with_an_explicit_style_and_a_second_positional_argument_should_not_be_promoted_to_a_source_block.adoc", blocksTestListingBlockWithAnExplicitStyleAndASecondPositionalArgumentShouldNotBePromotedToASourceBlock},

	{"listing block with an explicit style should not be promoted to a source block if source-language is set", "asciidoctor/blocks_test_listing_block_with_an_explicit_style_should_not_be_promoted_to_a_source_block_if_source_language_is_set.adoc", blocksTestListingBlockWithAnExplicitStyleShouldNotBePromotedToASourceBlockIfSourceLanguageIsSet},

	{"source block with no title or language should generate screen element in docbook", "asciidoctor/blocks_test_source_block_with_no_title_or_language_should_generate_screen_element_in_docbook.adoc", blocksTestSourceBlockWithNoTitleOrLanguageShouldGenerateScreenElementInDocbook},

	{"source block with title and no language should generate screen element inside formalpara element for docbook", "asciidoctor/blocks_test_source_block_with_title_and_no_language_should_generate_screen_element_inside_formalpara_element_for_docbook.adoc", blocksTestSourceBlockWithTitleAndNoLanguageShouldGenerateScreenElementInsideFormalparaElementForDocbook},

	{"can convert open block", "asciidoctor/blocks_test_can_convert_open_block.adoc", blocksTestCanConvertOpenBlock},

	{"open block can contain another block", "asciidoctor/blocks_test_open_block_can_contain_another_block.adoc", blocksTestOpenBlockCanContainAnotherBlock},

	{"should transfer id and reftext on open block to DocBook output", "asciidoctor/blocks_test_should_transfer_id_and_reftext_on_open_block_to_doc_book_output.adoc", blocksTestShouldTransferIdAndReftextOnOpenBlockToDocBookOutput},

	{"should transfer id and reftext on open paragraph to DocBook output", "asciidoctor/blocks_test_should_transfer_id_and_reftext_on_open_paragraph_to_doc_book_output.adoc", blocksTestShouldTransferIdAndReftextOnOpenParagraphToDocBookOutput},

	{"should transfer title on open block to DocBook output", "asciidoctor/blocks_test_should_transfer_title_on_open_block_to_doc_book_output.adoc", blocksTestShouldTransferTitleOnOpenBlockToDocBookOutput},

	{"should transfer title on open paragraph to DocBook output", "asciidoctor/blocks_test_should_transfer_title_on_open_paragraph_to_doc_book_output.adoc", blocksTestShouldTransferTitleOnOpenParagraphToDocBookOutput},

	{"should transfer role on open block to DocBook output", "asciidoctor/blocks_test_should_transfer_role_on_open_block_to_doc_book_output.adoc", blocksTestShouldTransferRoleOnOpenBlockToDocBookOutput},

	{"should transfer role on open paragraph to DocBook output", "asciidoctor/blocks_test_should_transfer_role_on_open_paragraph_to_doc_book_output.adoc", blocksTestShouldTransferRoleOnOpenParagraphToDocBookOutput},

	{"can parse a passthrough block", "asciidoctor/blocks_test_can_parse_a_passthrough_block.adoc", blocksTestCanParseAPassthroughBlock},

	{"does not perform subs on a passthrough block by default", "asciidoctor/blocks_test_does_not_perform_subs_on_a_passthrough_block_by_default.adoc", blocksTestDoesNotPerformSubsOnAPassthroughBlockByDefault},

	{"does not perform subs on a passthrough block with pass style by default", "asciidoctor/blocks_test_does_not_perform_subs_on_a_passthrough_block_with_pass_style_by_default.adoc", blocksTestDoesNotPerformSubsOnAPassthroughBlockWithPassStyleByDefault},

	{"passthrough block honors explicit subs list", "asciidoctor/blocks_test_passthrough_block_honors_explicit_subs_list.adoc", blocksTestPassthroughBlockHonorsExplicitSubsList},

	{"should strip leading and trailing blank lines when converting raw block", "asciidoctor/blocks_test_should_strip_leading_and_trailing_blank_lines_when_converting_raw_block.adoc", blocksTestShouldStripLeadingAndTrailingBlankLinesWhenConvertingRawBlock},

	{"should not crash when converting stem block that has no lines", "asciidoctor/blocks_test_should_not_crash_when_converting_stem_block_that_has_no_lines.adoc", blocksTestShouldNotCrashWhenConvertingStemBlockThatHasNoLines},

	{"should return content as empty string for stem or pass block that has no lines", "asciidoctor/blocks_test_should_return_content_as_empty_string_for_stem_or_pass_block_that_has_no_lines.adoc", blocksTestShouldReturnContentAsEmptyStringForStemOrPassBlockThatHasNoLines},

	{"should not add LaTeX math delimiters around latexmath block content if already present", "asciidoctor/blocks_test_should_not_add_la_te_x_math_delimiters_around_latexmath_block_content_if_already_present.adoc", blocksTestShouldNotAddLaTeXMathDelimitersAroundLatexmathBlockContentIfAlreadyPresent},

	{"should display latexmath block in alt of equation in DocBook backend", "asciidoctor/blocks_test_should_display_latexmath_block_in_alt_of_equation_in_doc_book_backend.adoc", blocksTestShouldDisplayLatexmathBlockInAltOfEquationInDocBookBackend},

	{"should set autoNumber option for latexmath to none by default", "asciidoctor/blocks_test_should_set_auto_number_option_for_latexmath_to_none_by_default.adoc", blocksTestShouldSetAutoNumberOptionForLatexmathToNoneByDefault},

	{"should set autoNumber option for latexmath to none if eqnums is set to none", "asciidoctor/blocks_test_should_set_auto_number_option_for_latexmath_to_none_if_eqnums_is_set_to_none.adoc", blocksTestShouldSetAutoNumberOptionForLatexmathToNoneIfEqnumsIsSetToNone},

	{"should set autoNumber option for latexmath to AMS if eqnums is set", "asciidoctor/blocks_test_should_set_auto_number_option_for_latexmath_to_ams_if_eqnums_is_set.adoc", blocksTestShouldSetAutoNumberOptionForLatexmathToAmsIfEqnumsIsSet},

	{"should set autoNumber option for latexmath to all if eqnums is set to all", "asciidoctor/blocks_test_should_set_auto_number_option_for_latexmath_to_all_if_eqnums_is_set_to_all.adoc", blocksTestShouldSetAutoNumberOptionForLatexmathToAllIfEqnumsIsSetToAll},

	{"should not split equation in AsciiMath block at single newline", "asciidoctor/blocks_test_should_not_split_equation_in_ascii_math_block_at_single_newline.adoc", blocksTestShouldNotSplitEquationInAsciiMathBlockAtSingleNewline},

	{"should split equation in AsciiMath block at escaped newline", "asciidoctor/blocks_test_should_split_equation_in_ascii_math_block_at_escaped_newline.adoc", blocksTestShouldSplitEquationInAsciiMathBlockAtEscapedNewline},

	{"should split equation in AsciiMath block at sequence of escaped newlines", "asciidoctor/blocks_test_should_split_equation_in_ascii_math_block_at_sequence_of_escaped_newlines.adoc", blocksTestShouldSplitEquationInAsciiMathBlockAtSequenceOfEscapedNewlines},

	{"should split equation in AsciiMath block at newline sequence and preserve breaks", "asciidoctor/blocks_test_should_split_equation_in_ascii_math_block_at_newline_sequence_and_preserve_breaks.adoc", blocksTestShouldSplitEquationInAsciiMathBlockAtNewlineSequenceAndPreserveBreaks},

	{"should add AsciiMath delimiters around asciimath block content", "asciidoctor/blocks_test_should_add_ascii_math_delimiters_around_asciimath_block_content.adoc", blocksTestShouldAddAsciiMathDelimitersAroundAsciimathBlockContent},

	{"should not add AsciiMath delimiters around asciimath block content if already present", "asciidoctor/blocks_test_should_not_add_ascii_math_delimiters_around_asciimath_block_content_if_already_present.adoc", blocksTestShouldNotAddAsciiMathDelimitersAroundAsciimathBlockContentIfAlreadyPresent},

	{"should convert contents of asciimath block to MathML in DocBook output if asciimath gem is available", "asciidoctor/blocks_test_should_convert_contents_of_asciimath_block_to_math_ml_in_doc_book_output_if_asciimath_gem_is_available.adoc", blocksTestShouldConvertContentsOfAsciimathBlockToMathMlInDocBookOutputIfAsciimathGemIsAvailable},

	{"should output title for latexmath block if defined", "asciidoctor/blocks_test_should_output_title_for_latexmath_block_if_defined.adoc", blocksTestShouldOutputTitleForLatexmathBlockIfDefined},

	{"should output title for asciimath block if defined", "asciidoctor/blocks_test_should_output_title_for_asciimath_block_if_defined.adoc", blocksTestShouldOutputTitleForAsciimathBlockIfDefined},

	{"should add AsciiMath delimiters around stem block content if stem attribute is asciimath, empty, or not set", "asciidoctor/blocks_test_should_add_ascii_math_delimiters_around_stem_block_content_if_stem_attribute_is_asciimath_empty_or_not_set.adoc", blocksTestShouldAddAsciiMathDelimitersAroundStemBlockContentIfStemAttributeIsAsciimathEmptyOrNotSet},

	{"should add LaTeX math delimiters around stem block content if stem attribute is latexmath, latex, or tex", "asciidoctor/blocks_test_should_add_la_te_x_math_delimiters_around_stem_block_content_if_stem_attribute_is_latexmath_latex_or_tex.adoc", blocksTestShouldAddLaTeXMathDelimitersAroundStemBlockContentIfStemAttributeIsLatexmathLatexOrTex},

	{"should allow stem style to be set using second positional argument of block attributes", "asciidoctor/blocks_test_should_allow_stem_style_to_be_set_using_second_positional_argument_of_block_attributes.adoc", blocksTestShouldAllowStemStyleToBeSetUsingSecondPositionalArgumentOfBlockAttributes},

	{"should not warn if block style is unknown", "asciidoctor/blocks_test_should_not_warn_if_block_style_is_unknown.adoc", blocksTestShouldNotWarnIfBlockStyleIsUnknown},

	{"should log debug message if block style is unknown and debug level is enabled", "asciidoctor/blocks_test_should_log_debug_message_if_block_style_is_unknown_and_debug_level_is_enabled.adoc", blocksTestShouldLogDebugMessageIfBlockStyleIsUnknownAndDebugLevelIsEnabled},

	{"block title above section gets carried over to first block in section", "asciidoctor/blocks_test_block_title_above_section_gets_carried_over_to_first_block_in_section.adoc", blocksTestBlockTitleAboveSectionGetsCarriedOverToFirstBlockInSection},

	{"block title above document title demotes document title to a section title", "asciidoctor/blocks_test_block_title_above_document_title_demotes_document_title_to_a_section_title.adoc", blocksTestBlockTitleAboveDocumentTitleDemotesDocumentTitleToASectionTitle},

	{"block title above document title gets carried over to first block in first section if no preamble", "asciidoctor/blocks_test_block_title_above_document_title_gets_carried_over_to_first_block_in_first_section_if_no_preamble.adoc", blocksTestBlockTitleAboveDocumentTitleGetsCarriedOverToFirstBlockInFirstSectionIfNoPreamble},

	{"should apply substitutions to a block title in normal order", "asciidoctor/blocks_test_should_apply_substitutions_to_a_block_title_in_normal_order.adoc", blocksTestShouldApplySubstitutionsToABlockTitleInNormalOrder},

	{"empty attribute list should not appear in output", "asciidoctor/blocks_test_empty_attribute_list_should_not_appear_in_output.adoc", blocksTestEmptyAttributeListShouldNotAppearInOutput},

	{"empty block anchor should not appear in output", "asciidoctor/blocks_test_empty_block_anchor_should_not_appear_in_output.adoc", blocksTestEmptyBlockAnchorShouldNotAppearInOutput},

	{"can convert block image with alt text defined in macro", "asciidoctor/blocks_test_can_convert_block_image_with_alt_text_defined_in_macro.adoc", blocksTestCanConvertBlockImageWithAltTextDefinedInMacro},

	{"converts SVG image with alt text using img element when safe mode is secure", "asciidoctor/blocks_test_converts_svg_image_with_alt_text_using_img_element_when_safe_mode_is_secure.adoc", blocksTestConvertsSvgImageWithAltTextUsingImgElementWhenSafeModeIsSecure},

	{"inserts fallback image for SVG inside object element using same dimensions", "asciidoctor/blocks_test_inserts_fallback_image_for_svg_inside_object_element_using_same_dimensions.adoc", blocksTestInsertsFallbackImageForSvgInsideObjectElementUsingSameDimensions},

	{"detects SVG image URI that contains a query string", "asciidoctor/blocks_test_detects_svg_image_uri_that_contains_a_query_string.adoc", blocksTestDetectsSvgImageUriThatContainsAQueryString},

	{"detects SVG image when format attribute is svg", "asciidoctor/blocks_test_detects_svg_image_when_format_attribute_is_svg.adoc", blocksTestDetectsSvgImageWhenFormatAttributeIsSvg},

	{"converts to inline SVG image when inline option is set on block", "asciidoctor/blocks_test_converts_to_inline_svg_image_when_inline_option_is_set_on_block.adoc", blocksTestConvertsToInlineSvgImageWhenInlineOptionIsSetOnBlock},

	{"should ignore link attribute if value is self and image target is inline SVG", "asciidoctor/blocks_test_should_ignore_link_attribute_if_value_is_self_and_image_target_is_inline_svg.adoc", blocksTestShouldIgnoreLinkAttributeIfValueIsSelfAndImageTargetIsInlineSvg},

	{"should honor percentage width for SVG image with inline option", "asciidoctor/blocks_test_should_honor_percentage_width_for_svg_image_with_inline_option.adoc", blocksTestShouldHonorPercentageWidthForSvgImageWithInlineOption},

	{"should not crash if explicit width on SVG image block is an integer", "asciidoctor/blocks_test_should_not_crash_if_explicit_width_on_svg_image_block_is_an_integer.adoc", blocksTestShouldNotCrashIfExplicitWidthOnSvgImageBlockIsAnInteger},

	{"converts to inline SVG image when inline option is set on block and data-uri is set on document", "asciidoctor/blocks_test_converts_to_inline_svg_image_when_inline_option_is_set_on_block_and_data_uri_is_set_on_document.adoc", blocksTestConvertsToInlineSvgImageWhenInlineOptionIsSetOnBlockAndDataUriIsSetOnDocument},

	{"should not throw exception if SVG to inline is empty", "asciidoctor/blocks_test_should_not_throw_exception_if_svg_to_inline_is_empty.adoc", blocksTestShouldNotThrowExceptionIfSvgToInlineIsEmpty},

	{"can convert block image with alt text defined in macro containing square bracket", "asciidoctor/blocks_test_can_convert_block_image_with_alt_text_defined_in_macro_containing_square_bracket.adoc", blocksTestCanConvertBlockImageWithAltTextDefinedInMacroContainingSquareBracket},

	{"alt text in macro overrides alt text above macro", "asciidoctor/blocks_test_alt_text_in_macro_overrides_alt_text_above_macro.adoc", blocksTestAltTextInMacroOverridesAltTextAboveMacro},

	{"should substitute attribute references in alt text defined in image block macro", "asciidoctor/blocks_test_should_substitute_attribute_references_in_alt_text_defined_in_image_block_macro.adoc", blocksTestShouldSubstituteAttributeReferencesInAltTextDefinedInImageBlockMacro},

	{"should set direction CSS class on image if float attribute is set", "asciidoctor/blocks_test_should_set_direction_css_class_on_image_if_float_attribute_is_set.adoc", blocksTestShouldSetDirectionCssClassOnImageIfFloatAttributeIsSet},

	{"should set text alignment CSS class on image if align attribute is set", "asciidoctor/blocks_test_should_set_text_alignment_css_class_on_image_if_align_attribute_is_set.adoc", blocksTestShouldSetTextAlignmentCssClassOnImageIfAlignAttributeIsSet},

	{"style attribute is dropped from image macro", "asciidoctor/blocks_test_style_attribute_is_dropped_from_image_macro.adoc", blocksTestStyleAttributeIsDroppedFromImageMacro},

	{"should auto-generate alt text for block image if alt text is not specified", "asciidoctor/blocks_test_should_auto_generate_alt_text_for_block_image_if_alt_text_is_not_specified.adoc", blocksTestShouldAutoGenerateAltTextForBlockImageIfAltTextIsNotSpecified},

	{"can convert block image with link to self", "asciidoctor/blocks_test_can_convert_block_image_with_link_to_self.adoc", blocksTestCanConvertBlockImageWithLinkToSelf},

	{"adds rel=noopener attribute to block image with link that targets _blank window", "asciidoctor/blocks_test_adds_rel=noopener_attribute_to_block_image_with_link_that_targets__blank_window.adoc", blocksTestAddsRelnoopenerAttributeToBlockImageWithLinkThatTargetsBlankWindow},

	{"can convert block image with explicit caption", "asciidoctor/blocks_test_can_convert_block_image_with_explicit_caption.adoc", blocksTestCanConvertBlockImageWithExplicitCaption},

	{"can align image in DocBook backend", "asciidoctor/blocks_test_can_align_image_in_doc_book_backend.adoc", blocksTestCanAlignImageInDocBookBackend},

	{"should not drop line if image target is missing attribute reference and attribute-missing is drop", "asciidoctor/blocks_test_should_not_drop_line_if_image_target_is_missing_attribute_reference_and_attribute_missing_is_drop.adoc", blocksTestShouldNotDropLineIfImageTargetIsMissingAttributeReferenceAndAttributeMissingIsDrop},

	{"drops line if image target is missing attribute reference and attribute-missing is drop-line", "asciidoctor/blocks_test_drops_line_if_image_target_is_missing_attribute_reference_and_attribute_missing_is_drop_line.adoc", blocksTestDropsLineIfImageTargetIsMissingAttributeReferenceAndAttributeMissingIsDropLine},

	{"should not drop line if image target resolves to blank and attribute-missing is drop-line", "asciidoctor/blocks_test_should_not_drop_line_if_image_target_resolves_to_blank_and_attribute_missing_is_drop_line.adoc", blocksTestShouldNotDropLineIfImageTargetResolvesToBlankAndAttributeMissingIsDropLine},

	{"dropped image does not break processing of following section and attribute-missing is drop-line", "asciidoctor/blocks_test_dropped_image_does_not_break_processing_of_following_section_and_attribute_missing_is_drop_line.adoc", blocksTestDroppedImageDoesNotBreakProcessingOfFollowingSectionAndAttributeMissingIsDropLine},

	{"should pass through image that references uri", "asciidoctor/blocks_test_should_pass_through_image_that_references_uri.adoc", blocksTestShouldPassThroughImageThatReferencesUri},

	{"should encode spaces in image target if value is a URI", "asciidoctor/blocks_test_should_encode_spaces_in_image_target_if_value_is_a_uri.adoc", blocksTestShouldEncodeSpacesInImageTargetIfValueIsAUri},

	{"embeds base64-encoded data uri for image when data-uri attribute is set", "asciidoctor/blocks_test_embeds_base_64_encoded_data_uri_for_image_when_data_uri_attribute_is_set.adoc", blocksTestEmbedsBase64EncodedDataUriForImageWhenDataUriAttributeIsSet},

	{"embeds SVG image with image/svg+xml mimetype when file extension is .svg", "asciidoctor/blocks_test_embeds_svg_image_with_image_svg+xml_mimetype_when_file_extension_is__svg.adoc", blocksTestEmbedsSvgImageWithImageSvgxmlMimetypeWhenFileExtensionIsSvg},

	{"should link to data URI if value of link attribute is self and image is embedded", "asciidoctor/blocks_test_should_link_to_data_uri_if_value_of_link_attribute_is_self_and_image_is_embedded.adoc", blocksTestShouldLinkToDataUriIfValueOfLinkAttributeIsSelfAndImageIsEmbedded},

	{"embeds empty base64-encoded data uri for unreadable image when data-uri attribute is set", "asciidoctor/blocks_test_embeds_empty_base_64_encoded_data_uri_for_unreadable_image_when_data_uri_attribute_is_set.adoc", blocksTestEmbedsEmptyBase64EncodedDataUriForUnreadableImageWhenDataUriAttributeIsSet},

	{"embeds base64-encoded data uri with application/octet-stream mimetype when file extension is missing", "asciidoctor/blocks_test_embeds_base_64_encoded_data_uri_with_application_octet_stream_mimetype_when_file_extension_is_missing.adoc", blocksTestEmbedsBase64EncodedDataUriWithApplicationOctetStreamMimetypeWhenFileExtensionIsMissing},

	{"can handle embedded data uri images", "asciidoctor/blocks_test_can_handle_embedded_data_uri_images.adoc", blocksTestCanHandleEmbeddedDataUriImages},

	{"cleans reference to ancestor directories in imagesdir before reading image if safe mode level is at least SAFE", "asciidoctor/blocks_test_cleans_reference_to_ancestor_directories_in_imagesdir_before_reading_image_if_safe_mode_level_is_at_least_safe.adoc", blocksTestCleansReferenceToAncestorDirectoriesInImagesdirBeforeReadingImageIfSafeModeLevelIsAtLeastSafe},

	{"cleans reference to ancestor directories in target before reading image if safe mode level is at least SAFE", "asciidoctor/blocks_test_cleans_reference_to_ancestor_directories_in_target_before_reading_image_if_safe_mode_level_is_at_least_safe.adoc", blocksTestCleansReferenceToAncestorDirectoriesInTargetBeforeReadingImageIfSafeModeLevelIsAtLeastSafe},

	{"should detect and convert video macro", "asciidoctor/blocks_test_should_detect_and_convert_video_macro.adoc", blocksTestShouldDetectAndConvertVideoMacro},

	{"video macro should not use imagesdir attribute to resolve target if target is a URL", "asciidoctor/blocks_test_video_macro_should_not_use_imagesdir_attribute_to_resolve_target_if_target_is_a_url.adoc", blocksTestVideoMacroShouldNotUseImagesdirAttributeToResolveTargetIfTargetIsAUrl},

	{"video macro should output custom HTML with iframe for vimeo service", "asciidoctor/blocks_test_video_macro_should_output_custom_html_with_iframe_for_vimeo_service.adoc", blocksTestVideoMacroShouldOutputCustomHtmlWithIframeForVimeoService},

	{"audio macro should not use imagesdir attribute to resolve target if target is a URL", "asciidoctor/blocks_test_audio_macro_should_not_use_imagesdir_attribute_to_resolve_target_if_target_is_a_url.adoc", blocksTestAudioMacroShouldNotUseImagesdirAttributeToResolveTargetIfTargetIsAUrl},

	{"audio macro should honor all options", "asciidoctor/blocks_test_audio_macro_should_honor_all_options.adoc", blocksTestAudioMacroShouldHonorAllOptions},

	{"can resolve icon relative to custom iconsdir", "asciidoctor/blocks_test_can_resolve_icon_relative_to_custom_iconsdir.adoc", blocksTestCanResolveIconRelativeToCustomIconsdir},

	{"should add file extension to custom icon if not specified", "asciidoctor/blocks_test_should_add_file_extension_to_custom_icon_if_not_specified.adoc", blocksTestShouldAddFileExtensionToCustomIconIfNotSpecified},

	{"should allow icontype to be specified when using built-in admonition icon", "asciidoctor/blocks_test_should_allow_icontype_to_be_specified_when_using_built_in_admonition_icon.adoc", blocksTestShouldAllowIcontypeToBeSpecifiedWhenUsingBuiltInAdmonitionIcon},

	{"embeds base64-encoded data uri of icon when data-uri attribute is set and safe mode level is less than SECURE", "asciidoctor/blocks_test_embeds_base_64_encoded_data_uri_of_icon_when_data_uri_attribute_is_set_and_safe_mode_level_is_less_than_secure.adoc", blocksTestEmbedsBase64EncodedDataUriOfIconWhenDataUriAttributeIsSetAndSafeModeLevelIsLessThanSecure},

	{"should embed base64-encoded data uri of custom icon when data-uri attribute is set", "asciidoctor/blocks_test_should_embed_base_64_encoded_data_uri_of_custom_icon_when_data_uri_attribute_is_set.adoc", blocksTestShouldEmbedBase64EncodedDataUriOfCustomIconWhenDataUriAttributeIsSet},

	{"does not embed base64-encoded data uri of icon when safe mode level is SECURE or greater", "asciidoctor/blocks_test_does_not_embed_base_64_encoded_data_uri_of_icon_when_safe_mode_level_is_secure_or_greater.adoc", blocksTestDoesNotEmbedBase64EncodedDataUriOfIconWhenSafeModeLevelIsSecureOrGreater},

	{"cleans reference to ancestor directories before reading icon if safe mode level is at least SAFE", "asciidoctor/blocks_test_cleans_reference_to_ancestor_directories_before_reading_icon_if_safe_mode_level_is_at_least_safe.adoc", blocksTestCleansReferenceToAncestorDirectoriesBeforeReadingIconIfSafeModeLevelIsAtLeastSafe},

	{"should import Font Awesome and use font-based icons when value of icons attribute is font", "asciidoctor/blocks_test_should_import_font_awesome_and_use_font_based_icons_when_value_of_icons_attribute_is_font.adoc", blocksTestShouldImportFontAwesomeAndUseFontBasedIconsWhenValueOfIconsAttributeIsFont},

	{"font-based icon should not override icon specified on admonition", "asciidoctor/blocks_test_font_based_icon_should_not_override_icon_specified_on_admonition.adoc", blocksTestFontBasedIconShouldNotOverrideIconSpecifiedOnAdmonition},

	{"should use http uri scheme for assets when asset-uri-scheme is http", "asciidoctor/blocks_test_should_use_http_uri_scheme_for_assets_when_asset_uri_scheme_is_http.adoc", blocksTestShouldUseHttpUriSchemeForAssetsWhenAssetUriSchemeIsHttp},

	{"should use no uri scheme for assets when asset-uri-scheme is blank", "asciidoctor/blocks_test_should_use_no_uri_scheme_for_assets_when_asset_uri_scheme_is_blank.adoc", blocksTestShouldUseNoUriSchemeForAssetsWhenAssetUriSchemeIsBlank},

	{"restricts access to ancestor directories when safe mode level is at least SAFE", "asciidoctor/blocks_test_restricts_access_to_ancestor_directories_when_safe_mode_level_is_at_least_safe.adoc", blocksTestRestrictsAccessToAncestorDirectoriesWhenSafeModeLevelIsAtLeastSafe},

	{"should not recognize fenced code blocks with more than three delimiters", "asciidoctor/blocks_test_should_not_recognize_fenced_code_blocks_with_more_than_three_delimiters.adoc", blocksTestShouldNotRecognizeFencedCodeBlocksWithMoreThanThreeDelimiters},

	{"should support fenced code blocks with languages", "asciidoctor/blocks_test_should_support_fenced_code_blocks_with_languages.adoc", blocksTestShouldSupportFencedCodeBlocksWithLanguages},

	{"should support fenced code blocks with languages and numbering", "asciidoctor/blocks_test_should_support_fenced_code_blocks_with_languages_and_numbering.adoc", blocksTestShouldSupportFencedCodeBlocksWithLanguagesAndNumbering},

	{"should allow source style to be specified on literal block", "asciidoctor/blocks_test_should_allow_source_style_to_be_specified_on_literal_block.adoc", blocksTestShouldAllowSourceStyleToBeSpecifiedOnLiteralBlock},

	{"should allow source style and language to be specified on literal block", "asciidoctor/blocks_test_should_allow_source_style_and_language_to_be_specified_on_literal_block.adoc", blocksTestShouldAllowSourceStyleAndLanguageToBeSpecifiedOnLiteralBlock},

	{"should make abstract on open block without title a quote block for article", "asciidoctor/blocks_test_should_make_abstract_on_open_block_without_title_a_quote_block_for_article.adoc", blocksTestShouldMakeAbstractOnOpenBlockWithoutTitleAQuoteBlockForArticle},

	{"should make abstract on open block with title a quote block with title for article", "asciidoctor/blocks_test_should_make_abstract_on_open_block_with_title_a_quote_block_with_title_for_article.adoc", blocksTestShouldMakeAbstractOnOpenBlockWithTitleAQuoteBlockWithTitleForArticle},

	{"should allow abstract in document with title if doctype is book", "asciidoctor/blocks_test_should_allow_abstract_in_document_with_title_if_doctype_is_book.adoc", blocksTestShouldAllowAbstractInDocumentWithTitleIfDoctypeIsBook},

	{"should not allow abstract as direct child of document if doctype is book", "asciidoctor/blocks_test_should_not_allow_abstract_as_direct_child_of_document_if_doctype_is_book.adoc", blocksTestShouldNotAllowAbstractAsDirectChildOfDocumentIfDoctypeIsBook},

	{"should make abstract on open block without title converted to DocBook", "asciidoctor/blocks_test_should_make_abstract_on_open_block_without_title_converted_to_doc_book.adoc", blocksTestShouldMakeAbstractOnOpenBlockWithoutTitleConvertedToDocBook},

	{"should make abstract on open block with title converted to DocBook", "asciidoctor/blocks_test_should_make_abstract_on_open_block_with_title_converted_to_doc_book.adoc", blocksTestShouldMakeAbstractOnOpenBlockWithTitleConvertedToDocBook},

	{"should allow abstract in document with title if doctype is book converted to DocBook", "asciidoctor/blocks_test_should_allow_abstract_in_document_with_title_if_doctype_is_book_converted_to_doc_book.adoc", blocksTestShouldAllowAbstractInDocumentWithTitleIfDoctypeIsBookConvertedToDocBook},

	{"should not allow abstract as direct child of document if doctype is book converted to DocBook", "asciidoctor/blocks_test_should_not_allow_abstract_as_direct_child_of_document_if_doctype_is_book_converted_to_doc_book.adoc", blocksTestShouldNotAllowAbstractAsDirectChildOfDocumentIfDoctypeIsBookConvertedToDocBook},

	{"should accept partintro on open block without title", "asciidoctor/blocks_test_should_accept_partintro_on_open_block_without_title.adoc", blocksTestShouldAcceptPartintroOnOpenBlockWithoutTitle},

	{"should accept partintro on open block with title", "asciidoctor/blocks_test_should_accept_partintro_on_open_block_with_title.adoc", blocksTestShouldAcceptPartintroOnOpenBlockWithTitle},

	{"should exclude partintro if not a child of part", "asciidoctor/blocks_test_should_exclude_partintro_if_not_a_child_of_part.adoc", blocksTestShouldExcludePartintroIfNotAChildOfPart},

	{"should not allow partintro unless doctype is book", "asciidoctor/blocks_test_should_not_allow_partintro_unless_doctype_is_book.adoc", blocksTestShouldNotAllowPartintroUnlessDoctypeIsBook},

	{"should accept partintro on open block without title converted to DocBook", "asciidoctor/blocks_test_should_accept_partintro_on_open_block_without_title_converted_to_doc_book.adoc", blocksTestShouldAcceptPartintroOnOpenBlockWithoutTitleConvertedToDocBook},

	{"should accept partintro on open block with title converted to DocBook", "asciidoctor/blocks_test_should_accept_partintro_on_open_block_with_title_converted_to_doc_book.adoc", blocksTestShouldAcceptPartintroOnOpenBlockWithTitleConvertedToDocBook},

	{"should exclude partintro if not a child of part converted to DocBook", "asciidoctor/blocks_test_should_exclude_partintro_if_not_a_child_of_part_converted_to_doc_book.adoc", blocksTestShouldExcludePartintroIfNotAChildOfPartConvertedToDocBook},

	{"should not allow partintro unless doctype is book converted to DocBook", "asciidoctor/blocks_test_should_not_allow_partintro_unless_doctype_is_book_converted_to_doc_book.adoc", blocksTestShouldNotAllowPartintroUnlessDoctypeIsBookConvertedToDocBook},

	{"processor should not crash if subs are empty", "asciidoctor/blocks_test_processor_should_not_crash_if_subs_are_empty.adoc", blocksTestProcessorShouldNotCrashIfSubsAreEmpty},

	{"should be able to append subs to default block substitution list", "asciidoctor/blocks_test_should_be_able_to_append_subs_to_default_block_substitution_list.adoc", blocksTestShouldBeAbleToAppendSubsToDefaultBlockSubstitutionList},

	{"should be able to prepend subs to default block substitution list", "asciidoctor/blocks_test_should_be_able_to_prepend_subs_to_default_block_substitution_list.adoc", blocksTestShouldBeAbleToPrependSubsToDefaultBlockSubstitutionList},

	{"should be able to remove subs to default block substitution list", "asciidoctor/blocks_test_should_be_able_to_remove_subs_to_default_block_substitution_list.adoc", blocksTestShouldBeAbleToRemoveSubsToDefaultBlockSubstitutionList},

	{"should be able to prepend, append and remove subs from default block substitution list", "asciidoctor/blocks_test_should_be_able_to_prepend_append_and_remove_subs_from_default_block_substitution_list.adoc", blocksTestShouldBeAbleToPrependAppendAndRemoveSubsFromDefaultBlockSubstitutionList},

	{"should be able to set subs then modify them", "asciidoctor/blocks_test_should_be_able_to_set_subs_then_modify_them.adoc", blocksTestShouldBeAbleToSetSubsThenModifyThem},

	{"should not recognize block anchor with illegal id characters", "asciidoctor/blocks_test_should_not_recognize_block_anchor_with_illegal_id_characters.adoc", blocksTestShouldNotRecognizeBlockAnchorWithIllegalIdCharacters},

	{"should not recognize block anchor that starts with digit", "asciidoctor/blocks_test_should_not_recognize_block_anchor_that_starts_with_digit.adoc", blocksTestShouldNotRecognizeBlockAnchorThatStartsWithDigit},

	{"should recognize block anchor that starts with colon", "asciidoctor/blocks_test_should_recognize_block_anchor_that_starts_with_colon.adoc", blocksTestShouldRecognizeBlockAnchorThatStartsWithColon},

	{"should use specified id and reftext when registering block reference", "asciidoctor/blocks_test_should_use_specified_id_and_reftext_when_registering_block_reference.adoc", blocksTestShouldUseSpecifiedIdAndReftextWhenRegisteringBlockReference},

	{"should allow square brackets in block reference text", "asciidoctor/blocks_test_should_allow_square_brackets_in_block_reference_text.adoc", blocksTestShouldAllowSquareBracketsInBlockReferenceText},

	{"should allow comma in block reference text", "asciidoctor/blocks_test_should_allow_comma_in_block_reference_text.adoc", blocksTestShouldAllowCommaInBlockReferenceText},

	{"should resolve attribute reference in title using attribute defined at location of block", "asciidoctor/blocks_test_should_resolve_attribute_reference_in_title_using_attribute_defined_at_location_of_block.adoc", blocksTestShouldResolveAttributeReferenceInTitleUsingAttributeDefinedAtLocationOfBlock},

	{"should substitute attribute references in reftext when registering block reference", "asciidoctor/blocks_test_should_substitute_attribute_references_in_reftext_when_registering_block_reference.adoc", blocksTestShouldSubstituteAttributeReferencesInReftextWhenRegisteringBlockReference},

	{"should use specified reftext when registering block reference", "asciidoctor/blocks_test_should_use_specified_reftext_when_registering_block_reference.adoc", blocksTestShouldUseSpecifiedReftextWhenRegisteringBlockReference},
}

var blocksTestHorizontalRuleBetweenBlocks = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.ThematicBreak{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: nil,
					ID:    nil,
					Roles: []*asciidoc.ShorthandRole{
						&asciidoc.ShorthandRole{
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "fancy",
								},
							},
						},
					},
					Options: nil,
				},
			},
		},
	},
}

var blocksTestLineCommentBetweenParagraphsOffsetByBlankLines = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "first paragraph",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.SingleLineComment{
			Value: " line comment",
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "second paragraph",
		},
		&asciidoc.NewLine{},
	},
}

var blocksTestAdjacentLineCommentBetweenParagraphs = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "first line",
		},
		&asciidoc.NewLine{},
		&asciidoc.SingleLineComment{
			Value: " line comment",
		},
		&asciidoc.String{
			Value: "second line",
		},
		&asciidoc.NewLine{},
	},
}

var blocksTestCommentBlockBetweenParagraphsOffsetByBlankLines = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "first paragraph",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.MultiLineComment{
			Delimiter: asciidoc.Delimiter{
				Type:   2,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"block comment",
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "second paragraph",
		},
		&asciidoc.NewLine{},
	},
}

var blocksTestCommentBlockBetweenParagraphsOffsetByBlankLinesInsideDelimitedBlock = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.ExampleBlock{
			Delimiter: asciidoc.Delimiter{
				Type:   3,
				Length: 4,
			},
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "first paragraph",
				},
				&asciidoc.NewLine{},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.MultiLineComment{
					Delimiter: asciidoc.Delimiter{
						Type:   2,
						Length: 4,
					},
					LineList: asciidoc.LineList{
						"block comment",
					},
				},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "second paragraph",
				},
				&asciidoc.NewLine{},
			},
		},
	},
}

var blocksTestAdjacentCommentBlockBetweenParagraphs = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "first paragraph",
		},
		&asciidoc.NewLine{},
		&asciidoc.MultiLineComment{
			Delimiter: asciidoc.Delimiter{
				Type:   2,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"block comment",
			},
		},
		&asciidoc.String{
			Value: "second paragraph",
		},
		&asciidoc.NewLine{},
	},
}

var blocksTestCanConvertWithBlockCommentAtEndOfDocumentWithTrailingNewlines = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "paragraph",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.MultiLineComment{
			Delimiter: asciidoc.Delimiter{
				Type:   2,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"block comment",
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
	},
}

var blocksTestTrailingNewlinesAfterBlockCommentAtEndOfDocumentDoesNotCreateParagraph = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "paragraph",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.MultiLineComment{
			Delimiter: asciidoc.Delimiter{
				Type:   2,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"block comment",
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
	},
}

var blocksTestLineStartingWithThreeSlashesShouldNotBeLineComment = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Paragraph{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.TitleAttribute{
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "sample title",
						},
					},
				},
			},
			Elements:   asciidoc.Elements{},
			Admonition: 0,
		},
		&asciidoc.MultiLineComment{
			Delimiter: asciidoc.Delimiter{
				Type:   2,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"ifdef::asciidoctor[////]",
			},
		},
		&asciidoc.String{
			Value: "line should be shown",
		},
		&asciidoc.NewLine{},
	},
}

var blocksTestPreprocessorDirectivesShouldNotBeProcessedWithinCommentBlock = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "dummy line",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.MultiLineComment{
			Delimiter: asciidoc.Delimiter{
				Type:   2,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"ifdef::asciidoctor[////]",
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "line should be shown",
		},
		&asciidoc.NewLine{},
	},
}

var blocksTestShouldWarnIfUnterminatedCommentBlockIsDetectedInBody = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "before comment block",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "////",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "content that has been disabled",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "supposed to be after comment block, except it got swallowed by block comment",
		},
		&asciidoc.NewLine{},
	},
}

var blocksTestShouldWarnIfUnterminatedCommentBlockIsDetectedInsideAnotherBlock = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "before sidebar block",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.SidebarBlock{
			Delimiter: asciidoc.Delimiter{
				Type:   8,
				Length: 4,
			},
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "////",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "content that has been disabled",
				},
				&asciidoc.NewLine{},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "supposed to be after sidebar block, except it got swallowed by block comment",
		},
		&asciidoc.NewLine{},
	},
}

var blocksTestPreprocessorDirectivesShouldNotBeProcessedWithinCommentOpenBlock = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OpenBlock{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "comment",
							},
						},
					},
					ID:      nil,
					Roles:   nil,
					Options: nil,
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   7,
				Length: 2,
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "first line of comment",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "ifdef::asciidoctor[--]",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "line should not be shown",
				},
				&asciidoc.NewLine{},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
	},
}

var blocksTestPreprocessorDirectivesShouldNotBeProcessedOnSubsequentLinesOfACommentParagraph = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Paragraph{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "comment",
							},
						},
					},
					ID:      nil,
					Roles:   nil,
					Options: nil,
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "first line of content",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "ifdef::asciidoctor[////]",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 0,
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "this line should be shown",
		},
		&asciidoc.NewLine{},
	},
}

var blocksTestCommentStyleOnOpenBlockShouldOnlySkipBlock = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OpenBlock{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "comment",
							},
						},
					},
					ID:      nil,
					Roles:   nil,
					Options: nil,
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   7,
				Length: 2,
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "skip",
				},
				&asciidoc.NewLine{},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "this block",
				},
				&asciidoc.NewLine{},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "not this text",
		},
		&asciidoc.NewLine{},
	},
}

var blocksTestCommentStyleOnParagraphShouldOnlySkipParagraph = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Paragraph{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "comment",
							},
						},
					},
					ID:      nil,
					Roles:   nil,
					Options: nil,
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "skip",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "this paragraph",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 0,
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "not this text",
		},
		&asciidoc.NewLine{},
	},
}

var blocksTestCommentStyleOnParagraphShouldNotCauseAdjacentBlockToBeSkipped = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Paragraph{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "comment",
							},
						},
					},
					ID:      nil,
					Roles:   nil,
					Options: nil,
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "skip",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "this paragraph",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 0,
		},
		&asciidoc.Paragraph{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "example",
							},
						},
					},
					ID:      nil,
					Roles:   nil,
					Options: nil,
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "not this text",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 0,
		},
	},
}

var blocksTestShouldNotDropContentThatFollowsSkippedContentInsideADelimitedBlock = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.ExampleBlock{
			Delimiter: asciidoc.Delimiter{
				Type:   3,
				Length: 4,
			},
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "paragraph",
				},
				&asciidoc.NewLine{},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.Paragraph{
					AttributeList: asciidoc.AttributeList{
						&asciidoc.ShorthandAttribute{
							Style: &asciidoc.ShorthandStyle{
								Elements: asciidoc.Elements{
									&asciidoc.String{
										Value: "comment",
									},
								},
							},
							ID: &asciidoc.ShorthandID{
								Elements: asciidoc.Elements{
									&asciidoc.String{
										Value: "idname",
									},
								},
							},
							Roles:   nil,
							Options: nil,
						},
					},
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "skip",
						},
						&asciidoc.NewLine{},
					},
					Admonition: 0,
				},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "paragraph",
				},
				&asciidoc.NewLine{},
			},
		},
	},
}

var blocksTestShouldParseSidebarBlock = &asciidoc.Document{
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
				&asciidoc.SidebarBlock{
					Delimiter: asciidoc.Delimiter{
						Type:   8,
						Length: 4,
					},
					AttributeList: asciidoc.AttributeList{
						&asciidoc.TitleAttribute{
							Val: asciidoc.Elements{
								&asciidoc.String{
									Value: "Sidebar",
								},
							},
						},
					},
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "Content goes here",
						},
						&asciidoc.NewLine{},
					},
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Section",
				},
			},
			Level: 1,
		},
	},
}

var blocksTestQuoteBlockWithNoAttribution = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.QuoteBlock{
			Delimiter: asciidoc.Delimiter{
				Type:   11,
				Length: 4,
			},
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "A famous quote.",
				},
				&asciidoc.NewLine{},
			},
		},
	},
}

var blocksTestQuoteBlockWithAttribution = &asciidoc.Document{
	Elements: asciidoc.Elements{
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
							Value: "Famous Person",
						},
					},
				},
				&asciidoc.PositionalAttribute{
					Offset:      2,
					ImpliedName: "",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Famous Book (1999)",
						},
					},
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "A famous quote.",
				},
				&asciidoc.NewLine{},
			},
		},
	},
}

var blocksTestQuoteBlockWithAttributeAndIdAndRoleShorthand = &asciidoc.Document{
	Elements: asciidoc.Elements{
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
					ID: &asciidoc.ShorthandID{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "justice-to-all",
							},
						},
					},
					Roles: []*asciidoc.ShorthandRole{
						&asciidoc.ShorthandRole{
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "solidarity",
								},
							},
						},
					},
					Options: nil,
				},
				&asciidoc.PositionalAttribute{
					Offset:      1,
					ImpliedName: "",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Martin Luther King",
						},
					},
				},
				&asciidoc.PositionalAttribute{
					Offset:      2,
					ImpliedName: "",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Jr.",
						},
					},
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Injustice anywhere is a threat to justice everywhere.",
				},
				&asciidoc.NewLine{},
			},
		},
	},
}

var blocksTestSettingIdUsingStyleShorthandShouldNotResetBlockStyle = &asciidoc.Document{
	Elements: asciidoc.Elements{
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
							Value: "#justice-to-all.solidarity",
						},
					},
				},
				&asciidoc.PositionalAttribute{
					Offset:      2,
					ImpliedName: "",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Martin Luther King",
						},
					},
				},
				&asciidoc.PositionalAttribute{
					Offset:      3,
					ImpliedName: "",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Jr.",
						},
					},
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Injustice anywhere is a threat to justice everywhere.",
				},
				&asciidoc.NewLine{},
			},
		},
	},
}

var blocksTestQuoteBlockWithComplexContent = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.QuoteBlock{
			Delimiter: asciidoc.Delimiter{
				Type:   11,
				Length: 4,
			},
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "A famous quote.",
				},
				&asciidoc.NewLine{},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.Paragraph{
					AttributeList: nil,
					Elements: asciidoc.Elements{
						&asciidoc.Italic{
							AttributeList: nil,
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "That",
								},
							},
						},
						&asciidoc.String{
							Value: " was inspiring.",
						},
						&asciidoc.NewLine{},
					},
					Admonition: 1,
				},
			},
		},
	},
}

var blocksTestQuoteBlockWithAttributionConvertedToDocBook = &asciidoc.Document{
	Elements: asciidoc.Elements{
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
							Value: "Famous Person",
						},
					},
				},
				&asciidoc.PositionalAttribute{
					Offset:      2,
					ImpliedName: "",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Famous Book (1999)",
						},
					},
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "A famous quote.",
				},
				&asciidoc.NewLine{},
			},
		},
	},
}

var blocksTestEpigraphQuoteBlockWithAttributionConvertedToDocBook = &asciidoc.Document{
	Elements: asciidoc.Elements{
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
					Style: nil,
					ID:    nil,
					Roles: []*asciidoc.ShorthandRole{
						&asciidoc.ShorthandRole{
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "epigraph",
								},
							},
						},
					},
					Options: nil,
				},
				&asciidoc.PositionalAttribute{
					Offset:      1,
					ImpliedName: "",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Famous Person",
						},
					},
				},
				&asciidoc.PositionalAttribute{
					Offset:      2,
					ImpliedName: "",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Famous Book (1999)",
						},
					},
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "A famous quote.",
				},
				&asciidoc.NewLine{},
			},
		},
	},
}

var blocksTestMarkdownStyleQuoteBlockWithSingleParagraphAndNoAttribution = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "> A famous quote.",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "> Some more inspiring words.",
		},
		&asciidoc.NewLine{},
	},
}

var blocksTestLazyMarkdownStyleQuoteBlockWithSingleParagraphAndNoAttribution = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "> A famous quote.",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "Some more inspiring words.",
		},
		&asciidoc.NewLine{},
	},
}

var blocksTestMarkdownStyleQuoteBlockWithMultipleParagraphsAndNoAttribution = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "> A famous quote.",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: ">",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "> Some more inspiring words.",
		},
		&asciidoc.NewLine{},
	},
}

var blocksTestMarkdownStyleQuoteBlockWithMultipleBlocksAndNoAttribution = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "> A famous quote.",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: ">",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "> NOTE: Some more inspiring words.",
		},
		&asciidoc.NewLine{},
	},
}

var blocksTestMarkdownStyleQuoteBlockWithSingleParagraphAndAttribution = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "> A famous quote.",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "> Some more inspiring words.",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "> -- Famous Person, Famous Source, Volume 1 (1999)",
		},
		&asciidoc.NewLine{},
	},
}

var blocksTestMarkdownStyleQuoteBlockWithOnlyAttribution = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "> I hold it that a little rebellion now and then is a good thing,",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "> and as necessary in the political world as storms in the physical.",
		},
		&asciidoc.NewLine{},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Thomas Jefferson, ",
				},
				&asciidoc.Link{
					AttributeList: asciidoc.AttributeList{
						&asciidoc.PositionalAttribute{
							Offset:      0,
							ImpliedName: "alt",
							Val: asciidoc.Elements{
								&asciidoc.String{
									Value: "The Papers of Thomas Jefferson",
								},
							},
						},
						&asciidoc.PositionalAttribute{
							Offset:      1,
							ImpliedName: "",
							Val: asciidoc.Elements{
								&asciidoc.String{
									Value: "Volume 11",
								},
							},
						},
					},
					URL: asciidoc.URL{
						Scheme: "https://",
						Path: asciidoc.Elements{
							&asciidoc.String{
								Value: "jeffersonpapers.princeton.edu/selected-documents/james-madison-1",
							},
						},
					},
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "--",
			Checklist:     0,
		},
	},
}

var blocksTestQuotedParagraphStyleQuoteBlockWithAttribution = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "\"A famous quote.",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "Some more inspiring words.\"",
		},
		&asciidoc.NewLine{},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Famous Person, Famous Source, Volume 1 (1999)",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "--",
			Checklist:     0,
		},
	},
}

var blocksTestShouldParseCreditLineInQuotedParagraphStyleQuoteBlockLikePositionalBlockAttributes = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "\"I hold it that a little rebellion now and then is a good thing,",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "and as necessary in the political world as storms in the physical.\"",
		},
		&asciidoc.NewLine{},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Thomas Jefferson, ",
				},
				&asciidoc.Link{
					AttributeList: asciidoc.AttributeList{
						&asciidoc.PositionalAttribute{
							Offset:      0,
							ImpliedName: "alt",
							Val: asciidoc.Elements{
								&asciidoc.String{
									Value: "The Papers of Thomas Jefferson",
								},
							},
						},
						&asciidoc.PositionalAttribute{
							Offset:      1,
							ImpliedName: "",
							Val: asciidoc.Elements{
								&asciidoc.String{
									Value: "Volume 11",
								},
							},
						},
					},
					URL: asciidoc.URL{
						Scheme: "https://",
						Path: asciidoc.Elements{
							&asciidoc.String{
								Value: "jeffersonpapers.princeton.edu/selected-documents/james-madison-1",
							},
						},
					},
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "--",
			Checklist:     0,
		},
	},
}

var blocksTestSingleLineVerseBlockWithoutAttribution = &asciidoc.Document{
	Elements: asciidoc.Elements{
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
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "A famous verse.",
				},
				&asciidoc.NewLine{},
			},
		},
	},
}

var blocksTestSingleLineVerseBlockWithAttribution = &asciidoc.Document{
	Elements: asciidoc.Elements{
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
							Value: "Famous Poet",
						},
					},
				},
				&asciidoc.PositionalAttribute{
					Offset:      2,
					ImpliedName: "",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Famous Poem",
						},
					},
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "A famous verse.",
				},
				&asciidoc.NewLine{},
			},
		},
	},
}

var blocksTestSingleLineVerseBlockWithAttributionConvertedToDocBook = &asciidoc.Document{
	Elements: asciidoc.Elements{
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
							Value: "Famous Poet",
						},
					},
				},
				&asciidoc.PositionalAttribute{
					Offset:      2,
					ImpliedName: "",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Famous Poem",
						},
					},
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "A famous verse.",
				},
				&asciidoc.NewLine{},
			},
		},
	},
}

var blocksTestSingleLineEpigraphVerseBlockWithAttributionConvertedToDocBook = &asciidoc.Document{
	Elements: asciidoc.Elements{
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
					ID: nil,
					Roles: []*asciidoc.ShorthandRole{
						&asciidoc.ShorthandRole{
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "epigraph",
								},
							},
						},
					},
					Options: nil,
				},
				&asciidoc.PositionalAttribute{
					Offset:      1,
					ImpliedName: "",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Famous Poet",
						},
					},
				},
				&asciidoc.PositionalAttribute{
					Offset:      2,
					ImpliedName: "",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Famous Poem",
						},
					},
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "A famous verse.",
				},
				&asciidoc.NewLine{},
			},
		},
	},
}

var blocksTestMultiStanzaVerseBlock = &asciidoc.Document{
	Elements: asciidoc.Elements{
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
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "A famous verse.",
				},
				&asciidoc.NewLine{},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "Stanza two.",
				},
				&asciidoc.NewLine{},
			},
		},
	},
}

var blocksTestVerseBlockDoesNotContainBlockElements = &asciidoc.Document{
	Elements: asciidoc.Elements{
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
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "A famous verse.",
				},
				&asciidoc.NewLine{},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.LiteralBlock{
					AttributeList: nil,
					Delimiter: asciidoc.Delimiter{
						Type:   6,
						Length: 4,
					},
					LineList: asciidoc.LineList{
						"not a literal",
					},
				},
			},
		},
	},
}

var blocksTestVerseShouldHaveNormalSubs = &asciidoc.Document{
	Elements: asciidoc.Elements{
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
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "A famous verse",
				},
				&asciidoc.NewLine{},
			},
		},
	},
}

var blocksTestShouldNotRecognizeCalloutsInAVerse = &asciidoc.Document{
	Elements: asciidoc.Elements{
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
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "La la la <1>",
				},
				&asciidoc.NewLine{},
			},
		},
		&asciidoc.String{
			Value: "<1> Not pointing to a callout",
		},
		&asciidoc.NewLine{},
	},
}

var blocksTestShouldPerformNormalSubsOnAVerseBlock = &asciidoc.Document{
	Elements: asciidoc.Elements{
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
			},
			Elements: asciidoc.Elements{
				&asciidoc.Italic{
					AttributeList: nil,
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "GET /groups/link:#group-id[\\{group-id\\}]",
						},
					},
				},
				&asciidoc.NewLine{},
			},
		},
	},
}

var blocksTestCanConvertExampleBlock = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.ExampleBlock{
			Delimiter: asciidoc.Delimiter{
				Type:   3,
				Length: 4,
			},
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "This is an example of an example block.",
				},
				&asciidoc.NewLine{},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "How crazy is that?",
				},
				&asciidoc.NewLine{},
			},
		},
	},
}

var blocksTestAssignsSequentialNumberedCaptionToExampleBlockWithTitle = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.ExampleBlock{
			Delimiter: asciidoc.Delimiter{
				Type:   3,
				Length: 4,
			},
			AttributeList: asciidoc.AttributeList{
				&asciidoc.TitleAttribute{
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Writing Docs with AsciiDoc",
						},
					},
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Here's how you write AsciiDoc.",
				},
				&asciidoc.NewLine{},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "You just write.",
				},
				&asciidoc.NewLine{},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.ExampleBlock{
			Delimiter: asciidoc.Delimiter{
				Type:   3,
				Length: 4,
			},
			AttributeList: asciidoc.AttributeList{
				&asciidoc.TitleAttribute{
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Writing Docs with DocBook",
						},
					},
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Here's how you write DocBook.",
				},
				&asciidoc.NewLine{},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "You futz with XML.",
				},
				&asciidoc.NewLine{},
			},
		},
	},
}

var blocksTestAssignsSequentialCharacterCaptionToExampleBlockWithTitle = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "example-number",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "@",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.ExampleBlock{
			Delimiter: asciidoc.Delimiter{
				Type:   3,
				Length: 4,
			},
			AttributeList: asciidoc.AttributeList{
				&asciidoc.TitleAttribute{
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Writing Docs with AsciiDoc",
						},
					},
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Here's how you write AsciiDoc.",
				},
				&asciidoc.NewLine{},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "You just write.",
				},
				&asciidoc.NewLine{},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.ExampleBlock{
			Delimiter: asciidoc.Delimiter{
				Type:   3,
				Length: 4,
			},
			AttributeList: asciidoc.AttributeList{
				&asciidoc.TitleAttribute{
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Writing Docs with DocBook",
						},
					},
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Here's how you write DocBook.",
				},
				&asciidoc.NewLine{},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "You futz with XML.",
				},
				&asciidoc.NewLine{},
			},
		},
	},
}

var blocksTestShouldIncrementCounterForExampleEvenWhenExampleNumberIsLockedByTheApi = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.ExampleBlock{
			Delimiter: asciidoc.Delimiter{
				Type:   3,
				Length: 4,
			},
			AttributeList: asciidoc.AttributeList{
				&asciidoc.TitleAttribute{
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Writing Docs with AsciiDoc",
						},
					},
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Here's how you write AsciiDoc.",
				},
				&asciidoc.NewLine{},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "You just write.",
				},
				&asciidoc.NewLine{},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.ExampleBlock{
			Delimiter: asciidoc.Delimiter{
				Type:   3,
				Length: 4,
			},
			AttributeList: asciidoc.AttributeList{
				&asciidoc.TitleAttribute{
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Writing Docs with DocBook",
						},
					},
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Here's how you write DocBook.",
				},
				&asciidoc.NewLine{},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "You futz with XML.",
				},
				&asciidoc.NewLine{},
			},
		},
	},
}

var blocksTestShouldUseExplicitCaptionIfSpecified = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.ExampleBlock{
			Delimiter: asciidoc.Delimiter{
				Type:   3,
				Length: 4,
			},
			AttributeList: asciidoc.AttributeList{
				&asciidoc.NamedAttribute{
					Name: "caption",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Look! ",
						},
					},
					Quote: 2,
				},
				&asciidoc.TitleAttribute{
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Writing Docs with AsciiDoc",
						},
					},
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Here's how you write AsciiDoc.",
				},
				&asciidoc.NewLine{},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "You just write.",
				},
				&asciidoc.NewLine{},
			},
		},
	},
}

var blocksTestAutomaticCaptionCanBeTurnedOffAndOnAndModified = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.ExampleBlock{
			Delimiter: asciidoc.Delimiter{
				Type:   3,
				Length: 4,
			},
			AttributeList: asciidoc.AttributeList{
				&asciidoc.TitleAttribute{
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "first example",
						},
					},
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "an example",
				},
				&asciidoc.NewLine{},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name:     "caption",
			Elements: nil,
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.ExampleBlock{
			Delimiter: asciidoc.Delimiter{
				Type:   3,
				Length: 4,
			},
			AttributeList: asciidoc.AttributeList{
				&asciidoc.TitleAttribute{
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "second example",
						},
					},
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "another example",
				},
				&asciidoc.NewLine{},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeReset{
			Name: "caption",
		},
		&asciidoc.AttributeEntry{
			Name: "example-caption",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Exhibit",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.ExampleBlock{
			Delimiter: asciidoc.Delimiter{
				Type:   3,
				Length: 4,
			},
			AttributeList: asciidoc.AttributeList{
				&asciidoc.TitleAttribute{
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "third example",
						},
					},
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "yet another example",
				},
				&asciidoc.NewLine{},
			},
		},
	},
}

var blocksTestShouldUseExplicitCaptionIfSpecifiedEvenIfBlockSpecificGlobalCaptionIsDisabled = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeReset{
			Name: "example-caption",
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.ExampleBlock{
			Delimiter: asciidoc.Delimiter{
				Type:   3,
				Length: 4,
			},
			AttributeList: asciidoc.AttributeList{
				&asciidoc.NamedAttribute{
					Name: "caption",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Look! ",
						},
					},
					Quote: 2,
				},
				&asciidoc.TitleAttribute{
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Writing Docs with AsciiDoc",
						},
					},
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Here's how you write AsciiDoc.",
				},
				&asciidoc.NewLine{},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "You just write.",
				},
				&asciidoc.NewLine{},
			},
		},
	},
}

var blocksTestShouldUseGlobalCaptionIfSpecifiedEvenIfBlockSpecificGlobalCaptionIsDisabled = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeReset{
			Name: "example-caption",
		},
		&asciidoc.AttributeEntry{
			Name: "caption",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Look!{sp}",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.ExampleBlock{
			Delimiter: asciidoc.Delimiter{
				Type:   3,
				Length: 4,
			},
			AttributeList: asciidoc.AttributeList{
				&asciidoc.TitleAttribute{
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Writing Docs with AsciiDoc",
						},
					},
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Here's how you write AsciiDoc.",
				},
				&asciidoc.NewLine{},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "You just write.",
				},
				&asciidoc.NewLine{},
			},
		},
	},
}

var blocksTestShouldNotProcessCaptionAttributeOnBlockThatDoesNotSupportACaption = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OpenBlock{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.NamedAttribute{
					Name: "caption",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Look! ",
						},
					},
					Quote: 2,
				},
				&asciidoc.TitleAttribute{
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "No caption here",
						},
					},
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   7,
				Length: 2,
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "content",
				},
				&asciidoc.NewLine{},
			},
		},
	},
}

var blocksTestShouldCreateDetailsSummarySetIfCollapsibleOptionIsSet = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.ExampleBlock{
			Delimiter: asciidoc.Delimiter{
				Type:   3,
				Length: 4,
			},
			AttributeList: asciidoc.AttributeList{
				&asciidoc.TitleAttribute{
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Toggle Me",
						},
					},
				},
				&asciidoc.ShorthandAttribute{
					Style: nil,
					ID:    nil,
					Roles: nil,
					Options: []*asciidoc.ShorthandOption{
						&asciidoc.ShorthandOption{
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "collapsible",
								},
							},
						},
					},
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "This content is revealed when the user clicks the words \"Toggle Me\".",
				},
				&asciidoc.NewLine{},
			},
		},
	},
}

var blocksTestShouldOpenDetailsSummarySetIfCollapsibleAndOpenOptionsAreSet = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.ExampleBlock{
			Delimiter: asciidoc.Delimiter{
				Type:   3,
				Length: 4,
			},
			AttributeList: asciidoc.AttributeList{
				&asciidoc.TitleAttribute{
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Toggle Me",
						},
					},
				},
				&asciidoc.ShorthandAttribute{
					Style: nil,
					ID:    nil,
					Roles: nil,
					Options: []*asciidoc.ShorthandOption{
						&asciidoc.ShorthandOption{
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "collapsible",
								},
							},
						},
						&asciidoc.ShorthandOption{
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "open",
								},
							},
						},
					},
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "This content is revealed when the user clicks the words \"Toggle Me\".",
				},
				&asciidoc.NewLine{},
			},
		},
	},
}

var blocksTestShouldAddDefaultSummaryElementIfCollapsibleOptionIsSetAndTitleIsNotSpecifed = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.ExampleBlock{
			Delimiter: asciidoc.Delimiter{
				Type:   3,
				Length: 4,
			},
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: nil,
					ID:    nil,
					Roles: nil,
					Options: []*asciidoc.ShorthandOption{
						&asciidoc.ShorthandOption{
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "collapsible",
								},
							},
						},
					},
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "This content is revealed when the user clicks the words \"Details\".",
				},
				&asciidoc.NewLine{},
			},
		},
	},
}

var blocksTestShouldNotAllowCollapsibleBlockToIncrementExampleNumber = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.ExampleBlock{
			Delimiter: asciidoc.Delimiter{
				Type:   3,
				Length: 4,
			},
			AttributeList: asciidoc.AttributeList{
				&asciidoc.TitleAttribute{
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Before",
						},
					},
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "before",
				},
				&asciidoc.NewLine{},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.ExampleBlock{
			Delimiter: asciidoc.Delimiter{
				Type:   3,
				Length: 4,
			},
			AttributeList: asciidoc.AttributeList{
				&asciidoc.TitleAttribute{
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Show Me The Goods",
						},
					},
				},
				&asciidoc.ShorthandAttribute{
					Style: nil,
					ID:    nil,
					Roles: nil,
					Options: []*asciidoc.ShorthandOption{
						&asciidoc.ShorthandOption{
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "collapsible",
								},
							},
						},
					},
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "This content is revealed when the user clicks the words \"Show Me The Goods\".",
				},
				&asciidoc.NewLine{},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.ExampleBlock{
			Delimiter: asciidoc.Delimiter{
				Type:   3,
				Length: 4,
			},
			AttributeList: asciidoc.AttributeList{
				&asciidoc.TitleAttribute{
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "After",
						},
					},
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "after",
				},
				&asciidoc.NewLine{},
			},
		},
	},
}

var blocksTestShouldWarnIfExampleBlockIsNotTerminated = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "outside",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "====",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "inside",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "still inside",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "eof",
		},
		&asciidoc.NewLine{},
	},
}

var blocksTestCaptionBlockLevelAttributeShouldBeUsedAsCaption = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "tip-caption",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Pro Tip",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Paragraph{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.NamedAttribute{
					Name: "caption",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Pro Tip",
						},
					},
					Quote: 2,
				},
			},
			Elements:   asciidoc.Elements{},
			Admonition: 0,
		},
		&asciidoc.Paragraph{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Override the caption of an admonition block using an attribute entry",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 2,
		},
	},
}

var blocksTestCanOverrideCaptionOfAdmonitionBlockUsingDocumentAttribute = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "tip-caption",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Pro Tip",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Paragraph{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Override the caption of an admonition block using an attribute entry",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 2,
		},
	},
}

var blocksTestBlankCaptionDocumentAttributeShouldNotBlankAdmonitionBlockCaption = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name:     "caption",
			Elements: nil,
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Paragraph{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Override the caption of an admonition block using an attribute entry",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 2,
		},
	},
}

var blocksTestShouldSeparateAdjacentParagraphsAndListingIntoBlocks = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "paragraph 1",
		},
		&asciidoc.NewLine{},
		&asciidoc.Listing{
			AttributeList: nil,
			Delimiter: asciidoc.Delimiter{
				Type:   5,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"listing content",
			},
		},
		&asciidoc.String{
			Value: "paragraph 2",
		},
		&asciidoc.NewLine{},
	},
}

var blocksTestShouldWarnIfListingBlockIsNotTerminated = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "outside",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "----",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "inside",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "still inside",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "eof",
		},
		&asciidoc.NewLine{},
	},
}

var blocksTestShouldNotCrashWhenConvertingVerbatimBlockThatHasNoLines = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.LiteralBlock{
			AttributeList: nil,
			Delimiter: asciidoc.Delimiter{
				Type:   6,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"line one",
				"",
				"line two",
				"",
				"line three",
			},
		},
	},
}

var blocksTestShouldPreserveNewlinesInListingBlock = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Listing{
			AttributeList: nil,
			Delimiter: asciidoc.Delimiter{
				Type:   5,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"line one",
				"",
				"line two",
				"",
				"line three",
			},
		},
	},
}

var blocksTestShouldPreserveNewlinesInVerseBlock = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OpenBlock{
			AttributeList: nil,
			Delimiter: asciidoc.Delimiter{
				Type:   7,
				Length: 2,
			},
			Elements: asciidoc.Elements{
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
					},
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "line one",
						},
						&asciidoc.NewLine{},
						&asciidoc.EmptyLine{
							Text: "",
						},
						&asciidoc.String{
							Value: "line two",
						},
						&asciidoc.NewLine{},
						&asciidoc.EmptyLine{
							Text: "",
						},
						&asciidoc.String{
							Value: "line three",
						},
						&asciidoc.NewLine{},
					},
				},
			},
		},
	},
}

var blocksTestShouldStripLeadingAndTrailingBlankLinesWhenConvertingVerbatimBlock = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.LiteralBlock{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.NamedAttribute{
					Name: "subs",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "attributes+",
						},
					},
					Quote: 0,
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   6,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"",
				"",
				"  first line",
				"",
				"last line",
				"",
				"{empty}",
				"",
			},
		},
	},
}

var blocksTestShouldRemoveBlockIndentIfIndentAttributeIs0 = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Listing{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.NamedAttribute{
					Name: "indent",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "0",
						},
					},
					Quote: 2,
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   5,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"    def names",
				"",
				"      @names.split",
				"",
				"    end",
			},
		},
	},
}

var blocksTestShouldNotRemoveBlockIndentIfIndentAttributeIs1 = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Listing{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.NamedAttribute{
					Name: "indent",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "-1",
						},
					},
					Quote: 2,
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   5,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"    def names",
				"",
				"      @names.split",
				"",
				"    end",
			},
		},
	},
}

var blocksTestShouldSetBlockIndentToValueSpecifiedByIndentAttribute = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Listing{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.NamedAttribute{
					Name: "indent",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "1",
						},
					},
					Quote: 2,
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   5,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"    def names",
				"",
				"      @names.split",
				"",
				"    end",
			},
		},
	},
}

var blocksTestShouldSetBlockIndentToValueSpecifiedByIndentDocumentAttribute = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "source-indent",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "1",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Listing{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "source",
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
							Value: "ruby",
						},
					},
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   5,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"    def names",
				"",
				"      @names.split",
				"",
				"    end",
			},
		},
	},
}

var blocksTestLiteralBlockShouldHonorNowrapOption = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Listing{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.NamedAttribute{
					Name: "options",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "nowrap",
						},
					},
					Quote: 2,
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   5,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"Do not wrap me if I get too long.",
			},
		},
	},
}

var blocksTestLiteralBlockShouldSetNowrapClassIfPrewrapDocumentAttributeIsDisabled = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeReset{
			Name: "prewrap",
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Listing{
			AttributeList: nil,
			Delimiter: asciidoc.Delimiter{
				Type:   5,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"Do not wrap me if I get too long.",
			},
		},
	},
}

var blocksTestShouldPreserveGuardInFrontOfCalloutIfIconsAreNotEnabled = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Listing{
			AttributeList: nil,
			Delimiter: asciidoc.Delimiter{
				Type:   5,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"puts 'Hello, World!' # <1>",
				"puts 'Goodbye, World ;(' # <2>",
			},
		},
	},
}

var blocksTestShouldPreserveGuardAroundCalloutIfIconsAreNotEnabled = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Listing{
			AttributeList: nil,
			Delimiter: asciidoc.Delimiter{
				Type:   5,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"<parent> <!--1-->",
				"  <child/> <!--2-->",
				"</parent>",
			},
		},
	},
}

var blocksTestLiteralBlockShouldHonorExplicitSubsList = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Listing{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.NamedAttribute{
					Name: "subs",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "verbatim,quotes",
						},
					},
					Quote: 2,
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   5,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"Map<String, String> *attributes*; //<1>",
			},
		},
	},
}

var blocksTestShouldBeAbleToDisableCalloutsForLiteralBlock = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Listing{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.NamedAttribute{
					Name: "subs",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "specialcharacters",
						},
					},
					Quote: 2,
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   5,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"No callout here <1>",
			},
		},
	},
}

var blocksTestListingBlockShouldHonorExplicitSubsList = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Listing{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.NamedAttribute{
					Name: "subs",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "specialcharacters,quotes",
						},
					},
					Quote: 2,
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   5,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"$ *python functional_tests.py*",
				"Traceback (most recent call last):",
				"  File \"functional_tests.py\", line 4, in <module>",
				"    assert 'Django' in browser.title",
				"AssertionError",
			},
		},
	},
}

var blocksTestShouldNotMangleArrayThatContainsFormattedTextWithRoleInListingBlockWithQuotesSubEnabled = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "[,ruby,subs=+quotes]",
		},
		&asciidoc.NewLine{},
		&asciidoc.Listing{
			AttributeList: nil,
			Delimiter: asciidoc.Delimiter{
				Type:   5,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"nums = [1, 2, 3, [.added]#4#]",
			},
		},
	},
}

var blocksTestFirstCharacterOfBlockTitleMayBeAPeriodIfNotFollowedBySpace = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "..gitignore",
		},
		&asciidoc.NewLine{},
		&asciidoc.Listing{
			AttributeList: nil,
			Delimiter: asciidoc.Delimiter{
				Type:   5,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"/.bundle/",
				"/build/",
				"/Gemfile.lock",
			},
		},
	},
}

var blocksTestListingBlockWithoutTitleShouldGenerateScreenElementInDocbook = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Listing{
			AttributeList: nil,
			Delimiter: asciidoc.Delimiter{
				Type:   5,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"listing block",
			},
		},
	},
}

var blocksTestListingBlockWithTitleShouldGenerateScreenElementInsideFormalparaElementInDocbook = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Listing{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.TitleAttribute{
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "title",
						},
					},
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   5,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"listing block",
			},
		},
	},
}

var blocksTestShouldNotPrependCaptionToTitleOfListingBlockWithTitleIfListingCaptionAttributeIsNotSet = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Listing{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.TitleAttribute{
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "title",
						},
					},
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   5,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"listing block content",
			},
		},
	},
}

var blocksTestShouldPrependCaptionSpecifiedByListingCaptionAttributeAndNumberToTitleOfListingBlockWithTitle = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "listing-caption",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Listing",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Listing{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.TitleAttribute{
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "title",
						},
					},
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   5,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"listing block content",
			},
		},
	},
}

var blocksTestShouldPrependCaptionSpecifiedByCaptionAttributeOnListingBlockEvenIfListingCaptionAttributeIsNotSet = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Listing{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.NamedAttribute{
					Name: "caption",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Listing ",
						},
						&asciidoc.Counter{
							Name:         "listing-number",
							InitialValue: "",
							Display:      asciidoc.CounterVisibilityVisible,
						},
						&asciidoc.String{
							Value: ". ",
						},
					},
					Quote: 2,
				},
				&asciidoc.TitleAttribute{
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Behold!",
						},
					},
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   5,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"listing block content",
			},
		},
	},
}

var blocksTestListingBlockWithoutAnExplicitStyleAndWithASecondPositionalArgumentShouldBePromotedToASourceBlock = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "[,ruby]",
		},
		&asciidoc.NewLine{},
		&asciidoc.Listing{
			AttributeList: nil,
			Delimiter: asciidoc.Delimiter{
				Type:   5,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"puts 'Hello, Ruby!'",
			},
		},
	},
}

var blocksTestListingBlockWithoutAnExplicitStyleShouldBePromotedToASourceBlockIfSourceLanguageIsSet = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "source-language",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "ruby",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Listing{
			AttributeList: nil,
			Delimiter: asciidoc.Delimiter{
				Type:   5,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"puts 'Hello, Ruby!'",
			},
		},
	},
}

var blocksTestListingBlockWithAnExplicitStyleAndASecondPositionalArgumentShouldNotBePromotedToASourceBlock = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Listing{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "listing",
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
							Value: "ruby",
						},
					},
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   5,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"puts 'Hello, Ruby!'",
			},
		},
	},
}

var blocksTestListingBlockWithAnExplicitStyleShouldNotBePromotedToASourceBlockIfSourceLanguageIsSet = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "source-language",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "ruby",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Listing{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "listing",
							},
						},
					},
					ID:      nil,
					Roles:   nil,
					Options: nil,
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   5,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"puts 'Hello, Ruby!'",
			},
		},
	},
}

var blocksTestSourceBlockWithNoTitleOrLanguageShouldGenerateScreenElementInDocbook = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Listing{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "source",
							},
						},
					},
					ID:      nil,
					Roles:   nil,
					Options: nil,
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   5,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"source block",
			},
		},
	},
}

var blocksTestSourceBlockWithTitleAndNoLanguageShouldGenerateScreenElementInsideFormalparaElementForDocbook = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Listing{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "source",
							},
						},
					},
					ID:      nil,
					Roles:   nil,
					Options: nil,
				},
				&asciidoc.TitleAttribute{
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "title",
						},
					},
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   5,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"source block",
			},
		},
	},
}

var blocksTestCanConvertOpenBlock = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OpenBlock{
			AttributeList: nil,
			Delimiter: asciidoc.Delimiter{
				Type:   7,
				Length: 2,
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "This is an open block.",
				},
				&asciidoc.NewLine{},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "It can span multiple lines.",
				},
				&asciidoc.NewLine{},
			},
		},
	},
}

var blocksTestOpenBlockCanContainAnotherBlock = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OpenBlock{
			AttributeList: nil,
			Delimiter: asciidoc.Delimiter{
				Type:   7,
				Length: 2,
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "This is an open block.",
				},
				&asciidoc.NewLine{},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "It can span multiple lines.",
				},
				&asciidoc.NewLine{},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.QuoteBlock{
					Delimiter: asciidoc.Delimiter{
						Type:   11,
						Length: 4,
					},
					AttributeList: nil,
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "It can hold great quotes like this one.",
						},
						&asciidoc.NewLine{},
					},
				},
			},
		},
	},
}

var blocksTestShouldTransferIdAndReftextOnOpenBlockToDocBookOutput = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "Check out that ",
		},
		&asciidoc.CrossReference{
			AttributeList: nil,
			Elements:      nil,
			ID: asciidoc.Elements{
				&asciidoc.String{
					Value: "open",
				},
			},
			Format: 0,
		},
		&asciidoc.String{
			Value: "!",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OpenBlock{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.AnchorAttribute{
					ID: asciidoc.Elements{
						&asciidoc.String{
							Value: "open",
						},
					},
					Label: asciidoc.Elements{
						&asciidoc.String{
							Value: "Open Block",
						},
					},
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   7,
				Length: 2,
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "This is an open block.",
				},
				&asciidoc.NewLine{},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.Paragraph{
					AttributeList: nil,
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "An open block can have other blocks inside of it.",
						},
						&asciidoc.NewLine{},
					},
					Admonition: 2,
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "Back to our regularly scheduled programming.",
		},
		&asciidoc.NewLine{},
	},
}

var blocksTestShouldTransferIdAndReftextOnOpenParagraphToDocBookOutput = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Paragraph{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "open",
							},
						},
					},
					ID: &asciidoc.ShorthandID{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "openpara",
							},
						},
					},
					Roles:   nil,
					Options: nil,
				},
				&asciidoc.NamedAttribute{
					Name: "reftext",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Open Paragraph",
						},
					},
					Quote: 2,
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "This is an open paragraph.",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 0,
		},
	},
}

var blocksTestShouldTransferTitleOnOpenBlockToDocBookOutput = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OpenBlock{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.TitleAttribute{
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Behold the open",
						},
					},
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   7,
				Length: 2,
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "This is an open block with a title.",
				},
				&asciidoc.NewLine{},
			},
		},
	},
}

var blocksTestShouldTransferTitleOnOpenParagraphToDocBookOutput = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Paragraph{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.TitleAttribute{
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Behold the open",
						},
					},
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "This is an open paragraph with a title.",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 0,
		},
	},
}

var blocksTestShouldTransferRoleOnOpenBlockToDocBookOutput = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OpenBlock{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: nil,
					ID:    nil,
					Roles: []*asciidoc.ShorthandRole{
						&asciidoc.ShorthandRole{
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "container",
								},
							},
						},
					},
					Options: nil,
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   7,
				Length: 2,
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "This is an open block.",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "It holds stuff.",
				},
				&asciidoc.NewLine{},
			},
		},
	},
}

var blocksTestShouldTransferRoleOnOpenParagraphToDocBookOutput = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Paragraph{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: nil,
					ID:    nil,
					Roles: []*asciidoc.ShorthandRole{
						&asciidoc.ShorthandRole{
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "container",
								},
							},
						},
					},
					Options: nil,
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "This is an open block.",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "It holds stuff.",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 0,
		},
	},
}

var blocksTestCanParseAPassthroughBlock = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.StemBlock{
			AttributeList: nil,
			Delimiter: asciidoc.Delimiter{
				Type:   9,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"This is a passthrough block.",
			},
		},
	},
}

var blocksTestDoesNotPerformSubsOnAPassthroughBlockByDefault = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "type",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "passthrough",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.StemBlock{
			AttributeList: nil,
			Delimiter: asciidoc.Delimiter{
				Type:   9,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"This is a '{type}' block.",
				"http://asciidoc.org",
				"image:tiger.png[]",
			},
		},
	},
}

var blocksTestDoesNotPerformSubsOnAPassthroughBlockWithPassStyleByDefault = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "type",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "passthrough",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.StemBlock{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "pass",
							},
						},
					},
					ID:      nil,
					Roles:   nil,
					Options: nil,
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   9,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"This is a '{type}' block.",
				"http://asciidoc.org",
				"image:tiger.png[]",
			},
		},
	},
}

var blocksTestPassthroughBlockHonorsExplicitSubsList = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "type",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "passthrough",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.StemBlock{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.NamedAttribute{
					Name: "subs",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "attributes,quotes,macros",
						},
					},
					Quote: 2,
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   9,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"This is a _{type}_ block.",
				"http://asciidoc.org",
			},
		},
	},
}

var blocksTestShouldStripLeadingAndTrailingBlankLinesWhenConvertingRawBlock = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.StemBlock{
			AttributeList: nil,
			Delimiter: asciidoc.Delimiter{
				Type:   9,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"line above",
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.StemBlock{
			AttributeList: nil,
			Delimiter: asciidoc.Delimiter{
				Type:   9,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"",
				"",
				"  first line",
				"",
				"last line",
				"",
				"",
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.StemBlock{
			AttributeList: nil,
			Delimiter: asciidoc.Delimiter{
				Type:   9,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"line below",
			},
		},
	},
}

var blocksTestShouldNotCrashWhenConvertingStemBlockThatHasNoLines = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.StemBlock{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "stem",
							},
						},
					},
					ID:      nil,
					Roles:   nil,
					Options: nil,
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   9,
				Length: 4,
			},
			LineList: asciidoc.LineList{},
		},
	},
}

var blocksTestShouldReturnContentAsEmptyStringForStemOrPassBlockThatHasNoLines = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.StemBlock{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "latexmath",
							},
						},
					},
					ID:      nil,
					Roles:   nil,
					Options: nil,
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   9,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"\\sqrt{3x-1}+(1+x)^2 < y",
			},
		},
	},
}

var blocksTestShouldNotAddLaTeXMathDelimitersAroundLatexmathBlockContentIfAlreadyPresent = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.StemBlock{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "latexmath",
							},
						},
					},
					ID:      nil,
					Roles:   nil,
					Options: nil,
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   9,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"\\[\\sqrt{3x-1}+(1+x)^2 < y\\]",
			},
		},
	},
}

var blocksTestShouldDisplayLatexmathBlockInAltOfEquationInDocBookBackend = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.StemBlock{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "latexmath",
							},
						},
					},
					ID:      nil,
					Roles:   nil,
					Options: nil,
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   9,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"\\sqrt{3x-1}+(1+x)^2 < y",
			},
		},
	},
}

var blocksTestShouldSetAutoNumberOptionForLatexmathToNoneByDefault = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "stem",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "latexmath",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.StemBlock{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "stem",
							},
						},
					},
					ID:      nil,
					Roles:   nil,
					Options: nil,
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   9,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"y = x^2",
			},
		},
	},
}

var blocksTestShouldSetAutoNumberOptionForLatexmathToNoneIfEqnumsIsSetToNone = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "stem",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "latexmath",
				},
			},
		},
		&asciidoc.AttributeEntry{
			Name: "eqnums",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "none",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.StemBlock{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "stem",
							},
						},
					},
					ID:      nil,
					Roles:   nil,
					Options: nil,
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   9,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"y = x^2",
			},
		},
	},
}

var blocksTestShouldSetAutoNumberOptionForLatexmathToAmsIfEqnumsIsSet = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "stem",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "latexmath",
				},
			},
		},
		&asciidoc.AttributeEntry{
			Name:     "eqnums",
			Elements: nil,
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.StemBlock{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "stem",
							},
						},
					},
					ID:      nil,
					Roles:   nil,
					Options: nil,
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   9,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"\\begin{equation}",
				"y = x^2",
				"\\end{equation}",
			},
		},
	},
}

var blocksTestShouldSetAutoNumberOptionForLatexmathToAllIfEqnumsIsSetToAll = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "stem",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "latexmath",
				},
			},
		},
		&asciidoc.AttributeEntry{
			Name: "eqnums",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "all",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.StemBlock{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "stem",
							},
						},
					},
					ID:      nil,
					Roles:   nil,
					Options: nil,
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   9,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"y = x^2",
			},
		},
	},
}

var blocksTestShouldNotSplitEquationInAsciiMathBlockAtSingleNewline = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.StemBlock{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "asciimath",
							},
						},
					},
					ID:      nil,
					Roles:   nil,
					Options: nil,
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   9,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"f: bbb\"N\" -> bbb\"N\"",
				"f: x |-> x + 1",
			},
		},
	},
}

var blocksTestShouldSplitEquationInAsciiMathBlockAtEscapedNewline = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.StemBlock{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "asciimath",
							},
						},
					},
					ID:      nil,
					Roles:   nil,
					Options: nil,
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   9,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"f: bbb\"N\" -> bbb\"N\" \\",
				"f: x |-> x + 1",
			},
		},
	},
}

var blocksTestShouldSplitEquationInAsciiMathBlockAtSequenceOfEscapedNewlines = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.StemBlock{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "asciimath",
							},
						},
					},
					ID:      nil,
					Roles:   nil,
					Options: nil,
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   9,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"f: bbb\"N\" -> bbb\"N\" \\",
				"\\",
				"f: x |-> x + 1",
			},
		},
	},
}

var blocksTestShouldSplitEquationInAsciiMathBlockAtNewlineSequenceAndPreserveBreaks = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.StemBlock{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "asciimath",
							},
						},
					},
					ID:      nil,
					Roles:   nil,
					Options: nil,
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   9,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"f: bbb\"N\" -> bbb\"N\"",
				"",
				"",
				"f: x |-> x + 1",
			},
		},
	},
}

var blocksTestShouldAddAsciiMathDelimitersAroundAsciimathBlockContent = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.StemBlock{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "asciimath",
							},
						},
					},
					ID:      nil,
					Roles:   nil,
					Options: nil,
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   9,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"sqrt(3x-1)+(1+x)^2 < y",
			},
		},
	},
}

var blocksTestShouldNotAddAsciiMathDelimitersAroundAsciimathBlockContentIfAlreadyPresent = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.StemBlock{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "asciimath",
							},
						},
					},
					ID:      nil,
					Roles:   nil,
					Options: nil,
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   9,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"\\$sqrt(3x-1)+(1+x)^2 < y\\$",
			},
		},
	},
}

var blocksTestShouldConvertContentsOfAsciimathBlockToMathMlInDocBookOutputIfAsciimathGemIsAvailable = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.StemBlock{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "asciimath",
							},
						},
					},
					ID:      nil,
					Roles:   nil,
					Options: nil,
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   9,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"x+b/(2a)<+-sqrt((b^2)/(4a^2)-c/a)",
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.StemBlock{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "asciimath",
							},
						},
					},
					ID:      nil,
					Roles:   nil,
					Options: nil,
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   9,
				Length: 4,
			},
			LineList: asciidoc.LineList{},
		},
	},
}

var blocksTestShouldOutputTitleForLatexmathBlockIfDefined = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.StemBlock{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.TitleAttribute{
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "The Lorenz Equations",
						},
					},
				},
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "latexmath",
							},
						},
					},
					ID:      nil,
					Roles:   nil,
					Options: nil,
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   9,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"\\begin{aligned}",
				"\\dot{x} & = \\sigma(y-x) \\\\",
				"\\dot{y} & = \\rho x - y - xz \\\\",
				"\\dot{z} & = -\\beta z + xy",
				"\\end{aligned}",
			},
		},
	},
}

var blocksTestShouldOutputTitleForAsciimathBlockIfDefined = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.StemBlock{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.TitleAttribute{
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Simple fraction",
						},
					},
				},
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "asciimath",
							},
						},
					},
					ID:      nil,
					Roles:   nil,
					Options: nil,
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   9,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"a//b",
			},
		},
	},
}

var blocksTestShouldAddAsciiMathDelimitersAroundStemBlockContentIfStemAttributeIsAsciimathEmptyOrNotSet = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.StemBlock{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "stem",
							},
						},
					},
					ID:      nil,
					Roles:   nil,
					Options: nil,
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   9,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"sqrt(3x-1)+(1+x)^2 < y",
			},
		},
	},
}

var blocksTestShouldAddLaTeXMathDelimitersAroundStemBlockContentIfStemAttributeIsLatexmathLatexOrTex = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.StemBlock{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "stem",
							},
						},
					},
					ID:      nil,
					Roles:   nil,
					Options: nil,
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   9,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"\\sqrt{3x-1}+(1+x)^2 < y",
			},
		},
	},
}

var blocksTestShouldAllowStemStyleToBeSetUsingSecondPositionalArgumentOfBlockAttributes = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "stem",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "latexmath",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.StemBlock{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "stem",
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
							Value: "asciimath",
						},
					},
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   9,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"sqrt(3x-1)+(1+x)^2 < y",
			},
		},
	},
}

var blocksTestShouldNotWarnIfBlockStyleIsUnknown = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OpenBlock{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "foo",
							},
						},
					},
					ID:      nil,
					Roles:   nil,
					Options: nil,
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   7,
				Length: 2,
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "bar",
				},
				&asciidoc.NewLine{},
			},
		},
	},
}

var blocksTestShouldLogDebugMessageIfBlockStyleIsUnknownAndDebugLevelIsEnabled = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OpenBlock{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "foo",
							},
						},
					},
					ID:      nil,
					Roles:   nil,
					Options: nil,
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   7,
				Length: 2,
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "bar",
				},
				&asciidoc.NewLine{},
			},
		},
	},
}

var blocksTestBlockTitleAboveSectionGetsCarriedOverToFirstBlockInSection = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.TitleAttribute{
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Title",
						},
					},
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "paragraph",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Section",
				},
			},
			Level: 1,
		},
	},
}

var blocksTestBlockTitleAboveDocumentTitleDemotesDocumentTitleToASectionTitle = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
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
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "section paragraph",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Section Title",
				},
			},
			Level: 0,
		},
	},
}

var blocksTestBlockTitleAboveDocumentTitleGetsCarriedOverToFirstBlockInFirstSectionIfNoPreamble = &asciidoc.Document{
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
		&asciidoc.Section{
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
							Value: "paragraph",
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
					Value: "Document Title",
				},
			},
			Level: 0,
		},
	},
}

var blocksTestShouldApplySubstitutionsToABlockTitleInNormalOrder = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Paragraph{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.TitleAttribute{
					Val: asciidoc.Elements{
						&asciidoc.UserAttributeReference{
							Value: "link-url",
						},
						&asciidoc.String{
							Value: "[",
						},
						&asciidoc.UserAttributeReference{
							Value: "link-text",
						},
						&asciidoc.String{
							Value: "]",
						},
						&asciidoc.UserAttributeReference{
							Value: "tm",
						},
					},
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "The one and only!",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 0,
		},
	},
}

var blocksTestEmptyAttributeListShouldNotAppearInOutput = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OpenBlock{
			AttributeList: nil,
			Delimiter: asciidoc.Delimiter{
				Type:   7,
				Length: 2,
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Block content",
				},
				&asciidoc.NewLine{},
			},
		},
	},
}

var blocksTestEmptyBlockAnchorShouldNotAppearInOutput = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "[[]]",
		},
		&asciidoc.NewLine{},
		&asciidoc.OpenBlock{
			AttributeList: nil,
			Delimiter: asciidoc.Delimiter{
				Type:   7,
				Length: 2,
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Block content",
				},
				&asciidoc.NewLine{},
			},
		},
	},
}

var blocksTestCanConvertBlockImageWithAltTextDefinedInMacro = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "imagesdir",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "images",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.BlockImage{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.PositionalAttribute{
					Offset:      0,
					ImpliedName: "alt",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Tiger",
						},
					},
				},
				&asciidoc.PositionalAttribute{
					Offset:      1,
					ImpliedName: "width",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "100",
						},
					},
				},
				&asciidoc.PositionalAttribute{
					Offset:      0,
					ImpliedName: "alt",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "%interactive",
						},
					},
				},
			},
			ImagePath: asciidoc.Elements{
				&asciidoc.String{
					Value: "tiger.svg",
				},
			},
		},
	},
}

var blocksTestConvertsSvgImageWithAltTextUsingImgElementWhenSafeModeIsSecure = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.BlockImage{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.PositionalAttribute{
					Offset:      0,
					ImpliedName: "alt",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Tiger",
						},
					},
				},
				&asciidoc.PositionalAttribute{
					Offset:      1,
					ImpliedName: "width",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "100",
						},
					},
				},
				&asciidoc.PositionalAttribute{
					Offset:      0,
					ImpliedName: "alt",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "%interactive",
						},
					},
				},
			},
			ImagePath: asciidoc.Elements{
				&asciidoc.String{
					Value: "images/tiger.svg",
				},
			},
		},
	},
}

var blocksTestInsertsFallbackImageForSvgInsideObjectElementUsingSameDimensions = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "imagesdir",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "images",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.BlockImage{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.PositionalAttribute{
					Offset:      0,
					ImpliedName: "alt",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Tiger",
						},
					},
				},
				&asciidoc.PositionalAttribute{
					Offset:      1,
					ImpliedName: "width",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "100",
						},
					},
				},
				&asciidoc.NamedAttribute{
					Name: "fallback",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "tiger.png",
						},
					},
					Quote: 0,
				},
				&asciidoc.PositionalAttribute{
					Offset:      0,
					ImpliedName: "alt",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "%interactive",
						},
					},
				},
			},
			ImagePath: asciidoc.Elements{
				&asciidoc.String{
					Value: "tiger.svg",
				},
			},
		},
	},
}

var blocksTestDetectsSvgImageUriThatContainsAQueryString = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "imagesdir",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "images",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.BlockImage{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.PositionalAttribute{
					Offset:      0,
					ImpliedName: "alt",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Tiger",
						},
					},
				},
				&asciidoc.PositionalAttribute{
					Offset:      1,
					ImpliedName: "width",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "100",
						},
					},
				},
				&asciidoc.PositionalAttribute{
					Offset:      0,
					ImpliedName: "alt",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "%interactive",
						},
					},
				},
			},
			ImagePath: asciidoc.Elements{
				&asciidoc.String{
					Value: "http://example.org/tiger.svg?foo=bar",
				},
			},
		},
	},
}

var blocksTestDetectsSvgImageWhenFormatAttributeIsSvg = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "imagesdir",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "images",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.BlockImage{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.PositionalAttribute{
					Offset:      0,
					ImpliedName: "alt",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Tiger",
						},
					},
				},
				&asciidoc.PositionalAttribute{
					Offset:      1,
					ImpliedName: "width",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "100",
						},
					},
				},
				&asciidoc.NamedAttribute{
					Name: "format",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "svg",
						},
					},
					Quote: 0,
				},
				&asciidoc.PositionalAttribute{
					Offset:      0,
					ImpliedName: "alt",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "%interactive",
						},
					},
				},
			},
			ImagePath: asciidoc.Elements{
				&asciidoc.String{
					Value: "http://example.org/tiger-svg",
				},
			},
		},
	},
}

var blocksTestConvertsToInlineSvgImageWhenInlineOptionIsSetOnBlock = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "imagesdir",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "fixtures",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.BlockImage{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.PositionalAttribute{
					Offset:      0,
					ImpliedName: "alt",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Tiger",
						},
					},
				},
				&asciidoc.PositionalAttribute{
					Offset:      1,
					ImpliedName: "width",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "100",
						},
					},
				},
				&asciidoc.PositionalAttribute{
					Offset:      0,
					ImpliedName: "alt",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "%inline",
						},
					},
				},
			},
			ImagePath: asciidoc.Elements{
				&asciidoc.String{
					Value: "circle.svg",
				},
			},
		},
	},
}

var blocksTestShouldIgnoreLinkAttributeIfValueIsSelfAndImageTargetIsInlineSvg = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "imagesdir",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "fixtures",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.BlockImage{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.PositionalAttribute{
					Offset:      0,
					ImpliedName: "alt",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Tiger",
						},
					},
				},
				&asciidoc.PositionalAttribute{
					Offset:      1,
					ImpliedName: "width",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "100",
						},
					},
				},
				&asciidoc.NamedAttribute{
					Name: "link",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "self",
						},
					},
					Quote: 0,
				},
				&asciidoc.PositionalAttribute{
					Offset:      0,
					ImpliedName: "alt",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "%inline",
						},
					},
				},
			},
			ImagePath: asciidoc.Elements{
				&asciidoc.String{
					Value: "circle.svg",
				},
			},
		},
	},
}

var blocksTestShouldHonorPercentageWidthForSvgImageWithInlineOption = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "imagesdir",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "fixtures",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.BlockImage{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.PositionalAttribute{
					Offset:      0,
					ImpliedName: "alt",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Circle",
						},
					},
				},
				&asciidoc.PositionalAttribute{
					Offset:      1,
					ImpliedName: "width",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "50%",
						},
					},
				},
				&asciidoc.NamedAttribute{
					Name: "opts",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "inline",
						},
					},
					Quote: 0,
				},
			},
			ImagePath: asciidoc.Elements{
				&asciidoc.String{
					Value: "circle.svg",
				},
			},
		},
	},
}

var blocksTestShouldNotCrashIfExplicitWidthOnSvgImageBlockIsAnInteger = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "imagesdir",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "fixtures",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.BlockImage{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.PositionalAttribute{
					Offset:      0,
					ImpliedName: "alt",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Circle",
						},
					},
				},
				&asciidoc.NamedAttribute{
					Name: "opts",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "inline",
						},
					},
					Quote: 0,
				},
			},
			ImagePath: asciidoc.Elements{
				&asciidoc.String{
					Value: "circle.svg",
				},
			},
		},
	},
}

var blocksTestConvertsToInlineSvgImageWhenInlineOptionIsSetOnBlockAndDataUriIsSetOnDocument = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "imagesdir",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "fixtures",
				},
			},
		},
		&asciidoc.AttributeEntry{
			Name:     "data-uri",
			Elements: nil,
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.BlockImage{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.PositionalAttribute{
					Offset:      0,
					ImpliedName: "alt",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Tiger",
						},
					},
				},
				&asciidoc.PositionalAttribute{
					Offset:      1,
					ImpliedName: "width",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "100",
						},
					},
				},
				&asciidoc.PositionalAttribute{
					Offset:      0,
					ImpliedName: "alt",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "%inline",
						},
					},
				},
			},
			ImagePath: asciidoc.Elements{
				&asciidoc.String{
					Value: "circle.svg",
				},
			},
		},
	},
}

var blocksTestShouldNotThrowExceptionIfSvgToInlineIsEmpty = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.BlockImage{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.PositionalAttribute{
					Offset:      0,
					ImpliedName: "alt",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Alt Text",
						},
					},
				},
				&asciidoc.PositionalAttribute{
					Offset:      0,
					ImpliedName: "alt",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "%inline",
						},
					},
				},
			},
			ImagePath: asciidoc.Elements{
				&asciidoc.String{
					Value: "no-such-image.svg",
				},
			},
		},
	},
}

var blocksTestCanConvertBlockImageWithAltTextDefinedInMacroContainingSquareBracket = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.BlockImage{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.PositionalAttribute{
					Offset:      0,
					ImpliedName: "alt",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Tiger",
						},
					},
				},
			},
			ImagePath: asciidoc.Elements{
				&asciidoc.String{
					Value: "images/tiger.png",
				},
			},
		},
	},
}

var blocksTestAltTextInMacroOverridesAltTextAboveMacro = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.BlockImage{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.PositionalAttribute{
					Offset:      0,
					ImpliedName: "alt",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Tiger",
						},
					},
				},
				&asciidoc.PositionalAttribute{
					Offset:      0,
					ImpliedName: "alt",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Alt Text",
						},
					},
				},
			},
			ImagePath: asciidoc.Elements{
				&asciidoc.String{
					Value: "images/tiger.png",
				},
			},
		},
	},
}

var blocksTestShouldSubstituteAttributeReferencesInAltTextDefinedInImageBlockMacro = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "alt-text",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Tiger",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.BlockImage{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.PositionalAttribute{
					Offset:      0,
					ImpliedName: "alt",
					Val: asciidoc.Elements{
						&asciidoc.UserAttributeReference{
							Value: "alt-text",
						},
					},
				},
			},
			ImagePath: asciidoc.Elements{
				&asciidoc.String{
					Value: "images/tiger.png",
				},
			},
		},
	},
}

var blocksTestShouldSetDirectionCssClassOnImageIfFloatAttributeIsSet = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.BlockImage{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.PositionalAttribute{
					Offset:      0,
					ImpliedName: "alt",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Tiger",
						},
					},
				},
				&asciidoc.NamedAttribute{
					Name: "float",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "left",
						},
					},
					Quote: 0,
				},
			},
			ImagePath: asciidoc.Elements{
				&asciidoc.String{
					Value: "images/tiger.png",
				},
			},
		},
	},
}

var blocksTestShouldSetTextAlignmentCssClassOnImageIfAlignAttributeIsSet = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.BlockImage{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.PositionalAttribute{
					Offset:      0,
					ImpliedName: "alt",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Tiger",
						},
					},
				},
				&asciidoc.NamedAttribute{
					Name: "align",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "center",
						},
					},
					Quote: 0,
				},
			},
			ImagePath: asciidoc.Elements{
				&asciidoc.String{
					Value: "images/tiger.png",
				},
			},
		},
	},
}

var blocksTestStyleAttributeIsDroppedFromImageMacro = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.BlockImage{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.PositionalAttribute{
					Offset:      0,
					ImpliedName: "alt",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Tiger",
						},
					},
				},
				&asciidoc.NamedAttribute{
					Name: "style",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "value",
						},
					},
					Quote: 0,
				},
			},
			ImagePath: asciidoc.Elements{
				&asciidoc.String{
					Value: "images/tiger.png",
				},
			},
		},
	},
}

var blocksTestShouldAutoGenerateAltTextForBlockImageIfAltTextIsNotSpecified = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.BlockImage{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.PositionalAttribute{
					Offset:      0,
					ImpliedName: "alt",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Tiger",
						},
					},
				},
				&asciidoc.NamedAttribute{
					Name: "link",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "http://en.wikipedia.org/wiki/Tiger",
						},
					},
					Quote: 1,
				},
			},
			ImagePath: asciidoc.Elements{
				&asciidoc.String{
					Value: "images/tiger.png",
				},
			},
		},
	},
}

var blocksTestCanConvertBlockImageWithLinkToSelf = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "imagesdir",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "img",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.BlockImage{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.PositionalAttribute{
					Offset:      0,
					ImpliedName: "alt",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Tiger",
						},
					},
				},
				&asciidoc.NamedAttribute{
					Name: "link",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "self",
						},
					},
					Quote: 0,
				},
			},
			ImagePath: asciidoc.Elements{
				&asciidoc.String{
					Value: "tiger.png",
				},
			},
		},
	},
}

var blocksTestAddsRelnoopenerAttributeToBlockImageWithLinkThatTargetsBlankWindow = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.BlockImage{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.PositionalAttribute{
					Offset:      0,
					ImpliedName: "alt",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Tiger",
						},
					},
				},
				&asciidoc.TitleAttribute{
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "The AsciiDoc Tiger",
						},
					},
				},
			},
			ImagePath: asciidoc.Elements{
				&asciidoc.String{
					Value: "images/tiger.png",
				},
			},
		},
	},
}

var blocksTestCanConvertBlockImageWithExplicitCaption = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.BlockImage{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.PositionalAttribute{
					Offset:      0,
					ImpliedName: "alt",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Tiger",
						},
					},
				},
				&asciidoc.NamedAttribute{
					Name: "caption",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Voila! ",
						},
					},
					Quote: 2,
				},
				&asciidoc.TitleAttribute{
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "The AsciiDoc Tiger",
						},
					},
				},
			},
			ImagePath: asciidoc.Elements{
				&asciidoc.String{
					Value: "images/tiger.png",
				},
			},
		},
	},
}

var blocksTestCanAlignImageInDocBookBackend = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "attribute-missing",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "skip",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.BlockImage{
			AttributeList: nil,
			ImagePath: asciidoc.Elements{
				&asciidoc.UserAttributeReference{
					Value: "bogus",
				},
			},
		},
	},
}

var blocksTestShouldNotDropLineIfImageTargetIsMissingAttributeReferenceAndAttributeMissingIsDrop = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "attribute-missing",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "drop",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.BlockImage{
			AttributeList: nil,
			ImagePath: asciidoc.Elements{
				&asciidoc.UserAttributeReference{
					Value: "bogus",
				},
				&asciidoc.String{
					Value: "/photo.jpg",
				},
			},
		},
	},
}

var blocksTestDropsLineIfImageTargetIsMissingAttributeReferenceAndAttributeMissingIsDropLine = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "attribute-missing",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "drop-line",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.BlockImage{
			AttributeList: nil,
			ImagePath: asciidoc.Elements{
				&asciidoc.UserAttributeReference{
					Value: "bogus",
				},
			},
		},
	},
}

var blocksTestShouldNotDropLineIfImageTargetResolvesToBlankAndAttributeMissingIsDropLine = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "attribute-missing",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "drop-line",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.BlockImage{
			AttributeList: nil,
			ImagePath: asciidoc.Elements{
				&asciidoc.CharacterReplacementReference{
					Value: "blank",
				},
			},
		},
	},
}

var blocksTestDroppedImageDoesNotBreakProcessingOfFollowingSectionAndAttributeMissingIsDropLine = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "attribute-missing",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "drop-line",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.BlockImage{
			AttributeList: nil,
			ImagePath: asciidoc.Elements{
				&asciidoc.UserAttributeReference{
					Value: "bogus",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Elements:      nil,
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Section Title",
				},
			},
			Level: 1,
		},
	},
}

var blocksTestShouldPassThroughImageThatReferencesUri = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "imagesdir",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "images",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.BlockImage{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.PositionalAttribute{
					Offset:      0,
					ImpliedName: "alt",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Tiger",
						},
					},
				},
			},
			ImagePath: asciidoc.Elements{
				&asciidoc.String{
					Value: "http://asciidoc.org/images/tiger.png",
				},
			},
		},
	},
}

var blocksTestShouldEncodeSpacesInImageTargetIfValueIsAUri = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "imagesdir",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "images",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.BlockImage{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.PositionalAttribute{
					Offset:      0,
					ImpliedName: "alt",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Tiger",
						},
					},
				},
			},
			ImagePath: asciidoc.Elements{
				&asciidoc.String{
					Value: "tiger.png",
				},
			},
		},
	},
}

var blocksTestEmbedsBase64EncodedDataUriForImageWhenDataUriAttributeIsSet = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name:     "data-uri",
			Elements: nil,
		},
		&asciidoc.AttributeEntry{
			Name: "imagesdir",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "fixtures",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.BlockImage{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.PositionalAttribute{
					Offset:      0,
					ImpliedName: "alt",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Dot",
						},
					},
				},
			},
			ImagePath: asciidoc.Elements{
				&asciidoc.String{
					Value: "dot.gif",
				},
			},
		},
	},
}

var blocksTestEmbedsSvgImageWithImageSvgxmlMimetypeWhenFileExtensionIsSvg = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "imagesdir",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "fixtures",
				},
			},
		},
		&asciidoc.AttributeEntry{
			Name:     "data-uri",
			Elements: nil,
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.BlockImage{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.PositionalAttribute{
					Offset:      0,
					ImpliedName: "alt",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Tiger",
						},
					},
				},
				&asciidoc.PositionalAttribute{
					Offset:      1,
					ImpliedName: "width",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "100",
						},
					},
				},
			},
			ImagePath: asciidoc.Elements{
				&asciidoc.String{
					Value: "circle.svg",
				},
			},
		},
	},
}

var blocksTestShouldLinkToDataUriIfValueOfLinkAttributeIsSelfAndImageIsEmbedded = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "imagesdir",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "fixtures",
				},
			},
		},
		&asciidoc.AttributeEntry{
			Name:     "data-uri",
			Elements: nil,
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.BlockImage{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.PositionalAttribute{
					Offset:      0,
					ImpliedName: "alt",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Tiger",
						},
					},
				},
				&asciidoc.PositionalAttribute{
					Offset:      1,
					ImpliedName: "width",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "100",
						},
					},
				},
				&asciidoc.NamedAttribute{
					Name: "link",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "self",
						},
					},
					Quote: 0,
				},
			},
			ImagePath: asciidoc.Elements{
				&asciidoc.String{
					Value: "circle.svg",
				},
			},
		},
	},
}

var blocksTestEmbedsEmptyBase64EncodedDataUriForUnreadableImageWhenDataUriAttributeIsSet = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name:     "data-uri",
			Elements: nil,
		},
		&asciidoc.AttributeEntry{
			Name: "imagesdir",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "fixtures",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.BlockImage{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.PositionalAttribute{
					Offset:      0,
					ImpliedName: "alt",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Dot",
						},
					},
				},
			},
			ImagePath: asciidoc.Elements{
				&asciidoc.String{
					Value: "unreadable.gif",
				},
			},
		},
	},
}

var blocksTestEmbedsBase64EncodedDataUriWithApplicationOctetStreamMimetypeWhenFileExtensionIsMissing = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name:     "data-uri",
			Elements: nil,
		},
		&asciidoc.AttributeEntry{
			Name: "imagesdir",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "fixtures",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.BlockImage{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.PositionalAttribute{
					Offset:      0,
					ImpliedName: "alt",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Dot",
						},
					},
				},
			},
			ImagePath: asciidoc.Elements{
				&asciidoc.String{
					Value: "dot",
				},
			},
		},
	},
}

var blocksTestCanHandleEmbeddedDataUriImages = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name:     "data-uri",
			Elements: nil,
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.BlockImage{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.PositionalAttribute{
					Offset:      0,
					ImpliedName: "alt",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Dot",
						},
					},
				},
			},
			ImagePath: asciidoc.Elements{
				&asciidoc.String{
					Value: "data:image/gif;base64,R0lGODlhAQABAIAAAAUEBAAAACwAAAAAAQABAAACAkQBADs=",
				},
			},
		},
	},
}

var blocksTestCleansReferenceToAncestorDirectoriesInImagesdirBeforeReadingImageIfSafeModeLevelIsAtLeastSafe = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name:     "data-uri",
			Elements: nil,
		},
		&asciidoc.AttributeEntry{
			Name: "imagesdir",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "../..//fixtures/./../../fixtures",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.BlockImage{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.PositionalAttribute{
					Offset:      0,
					ImpliedName: "alt",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Dot",
						},
					},
				},
			},
			ImagePath: asciidoc.Elements{
				&asciidoc.String{
					Value: "dot.gif",
				},
			},
		},
	},
}

var blocksTestCleansReferenceToAncestorDirectoriesInTargetBeforeReadingImageIfSafeModeLevelIsAtLeastSafe = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name:     "data-uri",
			Elements: nil,
		},
		&asciidoc.AttributeEntry{
			Name: "imagesdir",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "./",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.BlockImage{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.PositionalAttribute{
					Offset:      0,
					ImpliedName: "alt",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Dot",
						},
					},
				},
			},
			ImagePath: asciidoc.Elements{
				&asciidoc.String{
					Value: "../..//fixtures/./../../fixtures/dot.gif",
				},
			},
		},
	},
}

var blocksTestShouldDetectAndConvertVideoMacro = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "imagesdir",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "assets",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "video::cats-vs-dogs.avi[cats-and-dogs.png, 200, 300]",
		},
		&asciidoc.NewLine{},
	},
}

var blocksTestVideoMacroShouldNotUseImagesdirAttributeToResolveTargetIfTargetIsAUrl = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "imagesdir",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "assets",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "video::",
		},
		&asciidoc.Link{
			AttributeList: nil,
			URL: asciidoc.URL{
				Scheme: "http://",
				Path: asciidoc.Elements{
					&asciidoc.String{
						Value: "example.org/videos/cats-vs-dogs.avi",
					},
				},
			},
		},
		&asciidoc.NewLine{},
	},
}

var blocksTestVideoMacroShouldOutputCustomHtmlWithIframeForVimeoService = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "imagesdir",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "assets",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "audio::podcast.mp3[]",
		},
		&asciidoc.NewLine{},
	},
}

var blocksTestAudioMacroShouldNotUseImagesdirAttributeToResolveTargetIfTargetIsAUrl = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "imagesdir",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "assets",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "video::",
		},
		&asciidoc.Link{
			AttributeList: nil,
			URL: asciidoc.URL{
				Scheme: "http://",
				Path: asciidoc.Elements{
					&asciidoc.String{
						Value: "example.org/podcast.mp3",
					},
				},
			},
		},
		&asciidoc.NewLine{},
	},
}

var blocksTestAudioMacroShouldHonorAllOptions = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name:     "icons",
			Elements: nil,
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
								Value: "TIP",
							},
						},
					},
					ID:      nil,
					Roles:   nil,
					Options: nil,
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "You can use icons for admonitions by setting the 'icons' attribute.",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 0,
		},
	},
}

var blocksTestCanResolveIconRelativeToCustomIconsdir = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name:     "icons",
			Elements: nil,
		},
		&asciidoc.AttributeEntry{
			Name: "iconsdir",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "icons",
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
								Value: "TIP",
							},
						},
					},
					ID:      nil,
					Roles:   nil,
					Options: nil,
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "You can use icons for admonitions by setting the 'icons' attribute.",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 0,
		},
	},
}

var blocksTestShouldAddFileExtensionToCustomIconIfNotSpecified = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "icons",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "font",
				},
			},
		},
		&asciidoc.AttributeEntry{
			Name: "iconsdir",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "images/icons",
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
								Value: "TIP",
							},
						},
					},
					ID:      nil,
					Roles:   nil,
					Options: nil,
				},
				&asciidoc.NamedAttribute{
					Name: "icon",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "a",
						},
					},
					Quote: 0,
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Override the icon of an admonition block using an attribute",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 0,
		},
	},
}

var blocksTestShouldAllowIcontypeToBeSpecifiedWhenUsingBuiltInAdmonitionIcon = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Paragraph{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "TIP",
							},
						},
					},
					ID:      nil,
					Roles:   nil,
					Options: nil,
				},
				&asciidoc.NamedAttribute{
					Name: "icon",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "hint",
						},
					},
					Quote: 0,
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Set the icontype using either the icontype attribute on the icons attribute.",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 0,
		},
	},
}

var blocksTestEmbedsBase64EncodedDataUriOfIconWhenDataUriAttributeIsSetAndSafeModeLevelIsLessThanSecure = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name:     "icons",
			Elements: nil,
		},
		&asciidoc.AttributeEntry{
			Name: "iconsdir",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "fixtures",
				},
			},
		},
		&asciidoc.AttributeEntry{
			Name: "icontype",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "gif",
				},
			},
		},
		&asciidoc.AttributeEntry{
			Name:     "data-uri",
			Elements: nil,
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
								Value: "TIP",
							},
						},
					},
					ID:      nil,
					Roles:   nil,
					Options: nil,
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "You can use icons for admonitions by setting the 'icons' attribute.",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 0,
		},
	},
}

var blocksTestShouldEmbedBase64EncodedDataUriOfCustomIconWhenDataUriAttributeIsSet = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name:     "icons",
			Elements: nil,
		},
		&asciidoc.AttributeEntry{
			Name: "iconsdir",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "fixtures",
				},
			},
		},
		&asciidoc.AttributeEntry{
			Name: "icontype",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "gif",
				},
			},
		},
		&asciidoc.AttributeEntry{
			Name:     "data-uri",
			Elements: nil,
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
								Value: "TIP",
							},
						},
					},
					ID:      nil,
					Roles:   nil,
					Options: nil,
				},
				&asciidoc.NamedAttribute{
					Name: "icon",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "tip",
						},
					},
					Quote: 0,
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "You can set a custom icon using the icon attribute on the block.",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 0,
		},
	},
}

var blocksTestDoesNotEmbedBase64EncodedDataUriOfIconWhenSafeModeLevelIsSecureOrGreater = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name:     "icons",
			Elements: nil,
		},
		&asciidoc.AttributeEntry{
			Name: "iconsdir",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "fixtures",
				},
			},
		},
		&asciidoc.AttributeEntry{
			Name: "icontype",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "gif",
				},
			},
		},
		&asciidoc.AttributeEntry{
			Name:     "data-uri",
			Elements: nil,
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
								Value: "TIP",
							},
						},
					},
					ID:      nil,
					Roles:   nil,
					Options: nil,
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "You can use icons for admonitions by setting the 'icons' attribute.",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 0,
		},
	},
}

var blocksTestCleansReferenceToAncestorDirectoriesBeforeReadingIconIfSafeModeLevelIsAtLeastSafe = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name:     "icons",
			Elements: nil,
		},
		&asciidoc.AttributeEntry{
			Name: "iconsdir",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "../fixtures",
				},
			},
		},
		&asciidoc.AttributeEntry{
			Name: "icontype",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "gif",
				},
			},
		},
		&asciidoc.AttributeEntry{
			Name:     "data-uri",
			Elements: nil,
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
								Value: "TIP",
							},
						},
					},
					ID:      nil,
					Roles:   nil,
					Options: nil,
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "You can use icons for admonitions by setting the 'icons' attribute.",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 0,
		},
	},
}

var blocksTestShouldImportFontAwesomeAndUseFontBasedIconsWhenValueOfIconsAttributeIsFont = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "icons",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "font",
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
								Value: "TIP",
							},
						},
					},
					ID:      nil,
					Roles:   nil,
					Options: nil,
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "You can use icons for admonitions by setting the 'icons' attribute.",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 0,
		},
	},
}

var blocksTestFontBasedIconShouldNotOverrideIconSpecifiedOnAdmonition = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "icons",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "font",
				},
			},
		},
		&asciidoc.AttributeEntry{
			Name: "iconsdir",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "images/icons",
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
								Value: "TIP",
							},
						},
					},
					ID:      nil,
					Roles:   nil,
					Options: nil,
				},
				&asciidoc.NamedAttribute{
					Name: "icon",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "a.png",
						},
					},
					Quote: 0,
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Override the icon of an admonition block using an attribute",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 0,
		},
	},
}

var blocksTestShouldUseHttpUriSchemeForAssetsWhenAssetUriSchemeIsHttp = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "asset-uri-scheme",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "http",
				},
			},
		},
		&asciidoc.AttributeEntry{
			Name: "icons",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "font",
				},
			},
		},
		&asciidoc.AttributeEntry{
			Name: "source-highlighter",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "highlightjs",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Paragraph{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "You can control the URI scheme used for assets with the asset-uri-scheme attribute",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 2,
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
								Value: "source",
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
							Value: "ruby",
						},
					},
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "puts \"AsciiDoc, FTW!\"",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 0,
		},
	},
}

var blocksTestShouldUseNoUriSchemeForAssetsWhenAssetUriSchemeIsBlank = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name:     "asset-uri-scheme",
			Elements: nil,
		},
		&asciidoc.AttributeEntry{
			Name: "icons",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "font",
				},
			},
		},
		&asciidoc.AttributeEntry{
			Name: "source-highlighter",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "highlightjs",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Paragraph{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "You can control the URI scheme used for assets with the asset-uri-scheme attribute",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 2,
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
								Value: "source",
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
							Value: "ruby",
						},
					},
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "puts \"AsciiDoc, FTW!\"",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 0,
		},
	},
}

var blocksTestRestrictsAccessToAncestorDirectoriesWhenSafeModeLevelIsAtLeastSafe = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.FencedBlock{
			Delimiter: asciidoc.FencedDelimiter{
				Delimiter: asciidoc.Delimiter{
					Type:   4,
					Length: 3,
				},
				Language: nil,
			},
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "puts \"Hello, World!\"",
				},
				&asciidoc.NewLine{},
			},
		},
	},
}

var blocksTestShouldNotRecognizeFencedCodeBlocksWithMoreThanThreeDelimiters = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "````ruby",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "puts \"Hello, World!\"",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "````",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "~~~~ javascript",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "alert(\"Hello, World!\")",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "~~~~",
		},
		&asciidoc.NewLine{},
	},
}

var blocksTestShouldSupportFencedCodeBlocksWithLanguages = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.FencedBlock{
			Delimiter: asciidoc.FencedDelimiter{
				Delimiter: asciidoc.Delimiter{
					Type:   4,
					Length: 3,
				},
				Language: asciidoc.Elements{
					&asciidoc.String{
						Value: "ruby",
					},
				},
			},
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "puts \"Hello, World!\"",
				},
				&asciidoc.NewLine{},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.FencedBlock{
			Delimiter: asciidoc.FencedDelimiter{
				Delimiter: asciidoc.Delimiter{
					Type:   4,
					Length: 3,
				},
				Language: asciidoc.Elements{
					&asciidoc.String{
						Value: " javascript",
					},
				},
			},
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "alert(\"Hello, World!\")",
				},
				&asciidoc.NewLine{},
			},
		},
	},
}

var blocksTestShouldSupportFencedCodeBlocksWithLanguagesAndNumbering = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.FencedBlock{
			Delimiter: asciidoc.FencedDelimiter{
				Delimiter: asciidoc.Delimiter{
					Type:   4,
					Length: 3,
				},
				Language: asciidoc.Elements{
					&asciidoc.String{
						Value: "ruby,numbered",
					},
				},
			},
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "puts \"Hello, World!\"",
				},
				&asciidoc.NewLine{},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.FencedBlock{
			Delimiter: asciidoc.FencedDelimiter{
				Delimiter: asciidoc.Delimiter{
					Type:   4,
					Length: 3,
				},
				Language: asciidoc.Elements{
					&asciidoc.String{
						Value: " javascript, numbered",
					},
				},
			},
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "alert(\"Hello, World!\")",
				},
				&asciidoc.NewLine{},
			},
		},
	},
}

var blocksTestShouldAllowSourceStyleToBeSpecifiedOnLiteralBlock = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.LiteralBlock{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "source",
							},
						},
					},
					ID:      nil,
					Roles:   nil,
					Options: nil,
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   6,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"console.log('Hello, World!')",
			},
		},
	},
}

var blocksTestShouldAllowSourceStyleAndLanguageToBeSpecifiedOnLiteralBlock = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.LiteralBlock{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "source",
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
							Value: "js",
						},
					},
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   6,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"console.log('Hello, World!')",
			},
		},
	},
}

var blocksTestShouldMakeAbstractOnOpenBlockWithoutTitleAQuoteBlockForArticle = &asciidoc.Document{
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
				&asciidoc.OpenBlock{
					AttributeList: asciidoc.AttributeList{
						&asciidoc.ShorthandAttribute{
							Style: &asciidoc.ShorthandStyle{
								Elements: asciidoc.Elements{
									&asciidoc.String{
										Value: "abstract",
									},
								},
							},
							ID:      nil,
							Roles:   nil,
							Options: nil,
						},
					},
					Delimiter: asciidoc.Delimiter{
						Type:   7,
						Length: 2,
					},
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "This article is about stuff.",
						},
						&asciidoc.NewLine{},
						&asciidoc.EmptyLine{
							Text: "",
						},
						&asciidoc.String{
							Value: "And other stuff.",
						},
						&asciidoc.NewLine{},
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
							Value: "content",
						},
						&asciidoc.NewLine{},
					},
					Title: asciidoc.Elements{
						&asciidoc.String{
							Value: "Section One",
						},
					},
					Level: 1,
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Article",
				},
			},
			Level: 0,
		},
	},
}

var blocksTestShouldMakeAbstractOnOpenBlockWithTitleAQuoteBlockWithTitleForArticle = &asciidoc.Document{
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
				&asciidoc.OpenBlock{
					AttributeList: asciidoc.AttributeList{
						&asciidoc.TitleAttribute{
							Val: asciidoc.Elements{
								&asciidoc.String{
									Value: "My abstract",
								},
							},
						},
						&asciidoc.ShorthandAttribute{
							Style: &asciidoc.ShorthandStyle{
								Elements: asciidoc.Elements{
									&asciidoc.String{
										Value: "abstract",
									},
								},
							},
							ID:      nil,
							Roles:   nil,
							Options: nil,
						},
					},
					Delimiter: asciidoc.Delimiter{
						Type:   7,
						Length: 2,
					},
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "This article is about stuff.",
						},
						&asciidoc.NewLine{},
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
							Value: "content",
						},
						&asciidoc.NewLine{},
					},
					Title: asciidoc.Elements{
						&asciidoc.String{
							Value: "Section One",
						},
					},
					Level: 1,
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Article",
				},
			},
			Level: 0,
		},
	},
}

var blocksTestShouldAllowAbstractInDocumentWithTitleIfDoctypeIsBook = &asciidoc.Document{
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
				&asciidoc.Paragraph{
					AttributeList: asciidoc.AttributeList{
						&asciidoc.ShorthandAttribute{
							Style: &asciidoc.ShorthandStyle{
								Elements: asciidoc.Elements{
									&asciidoc.String{
										Value: "abstract",
									},
								},
							},
							ID:      nil,
							Roles:   nil,
							Options: nil,
						},
					},
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "Abstract for book with title is valid",
						},
						&asciidoc.NewLine{},
					},
					Admonition: 0,
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Book",
				},
			},
			Level: 0,
		},
	},
}

var blocksTestShouldNotAllowAbstractAsDirectChildOfDocumentIfDoctypeIsBook = &asciidoc.Document{
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
		&asciidoc.Paragraph{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "abstract",
							},
						},
					},
					ID:      nil,
					Roles:   nil,
					Options: nil,
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Abstract for book without title is invalid.",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 0,
		},
	},
}

var blocksTestShouldMakeAbstractOnOpenBlockWithoutTitleConvertedToDocBook = &asciidoc.Document{
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
				&asciidoc.OpenBlock{
					AttributeList: asciidoc.AttributeList{
						&asciidoc.ShorthandAttribute{
							Style: &asciidoc.ShorthandStyle{
								Elements: asciidoc.Elements{
									&asciidoc.String{
										Value: "abstract",
									},
								},
							},
							ID:      nil,
							Roles:   nil,
							Options: nil,
						},
					},
					Delimiter: asciidoc.Delimiter{
						Type:   7,
						Length: 2,
					},
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "This article is about stuff.",
						},
						&asciidoc.NewLine{},
						&asciidoc.EmptyLine{
							Text: "",
						},
						&asciidoc.String{
							Value: "And other stuff.",
						},
						&asciidoc.NewLine{},
					},
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Article",
				},
			},
			Level: 0,
		},
	},
}

var blocksTestShouldMakeAbstractOnOpenBlockWithTitleConvertedToDocBook = &asciidoc.Document{
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
				&asciidoc.OpenBlock{
					AttributeList: asciidoc.AttributeList{
						&asciidoc.TitleAttribute{
							Val: asciidoc.Elements{
								&asciidoc.String{
									Value: "My abstract",
								},
							},
						},
						&asciidoc.ShorthandAttribute{
							Style: &asciidoc.ShorthandStyle{
								Elements: asciidoc.Elements{
									&asciidoc.String{
										Value: "abstract",
									},
								},
							},
							ID:      nil,
							Roles:   nil,
							Options: nil,
						},
					},
					Delimiter: asciidoc.Delimiter{
						Type:   7,
						Length: 2,
					},
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "This article is about stuff.",
						},
						&asciidoc.NewLine{},
					},
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Article",
				},
			},
			Level: 0,
		},
	},
}

var blocksTestShouldAllowAbstractInDocumentWithTitleIfDoctypeIsBookConvertedToDocBook = &asciidoc.Document{
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
				&asciidoc.Paragraph{
					AttributeList: asciidoc.AttributeList{
						&asciidoc.ShorthandAttribute{
							Style: &asciidoc.ShorthandStyle{
								Elements: asciidoc.Elements{
									&asciidoc.String{
										Value: "abstract",
									},
								},
							},
							ID:      nil,
							Roles:   nil,
							Options: nil,
						},
					},
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "Abstract for book with title is valid",
						},
						&asciidoc.NewLine{},
					},
					Admonition: 0,
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Book",
				},
			},
			Level: 0,
		},
	},
}

var blocksTestShouldNotAllowAbstractAsDirectChildOfDocumentIfDoctypeIsBookConvertedToDocBook = &asciidoc.Document{
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
		&asciidoc.Paragraph{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "abstract",
							},
						},
					},
					ID:      nil,
					Roles:   nil,
					Options: nil,
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Abstract for book is invalid.",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 0,
		},
	},
}

var blocksTestShouldAcceptPartintroOnOpenBlockWithoutTitle = &asciidoc.Document{
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
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Book",
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
				&asciidoc.OpenBlock{
					AttributeList: asciidoc.AttributeList{
						&asciidoc.ShorthandAttribute{
							Style: &asciidoc.ShorthandStyle{
								Elements: asciidoc.Elements{
									&asciidoc.String{
										Value: "partintro",
									},
								},
							},
							ID:      nil,
							Roles:   nil,
							Options: nil,
						},
					},
					Delimiter: asciidoc.Delimiter{
						Type:   7,
						Length: 2,
					},
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "This is a part intro.",
						},
						&asciidoc.NewLine{},
						&asciidoc.EmptyLine{
							Text: "",
						},
						&asciidoc.String{
							Value: "It can have multiple paragraphs.",
						},
						&asciidoc.NewLine{},
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
							Value: "content",
						},
						&asciidoc.NewLine{},
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
					Value: "Part 1",
				},
			},
			Level: 0,
		},
	},
}

var blocksTestShouldAcceptPartintroOnOpenBlockWithTitle = &asciidoc.Document{
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
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Book",
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
				&asciidoc.OpenBlock{
					AttributeList: asciidoc.AttributeList{
						&asciidoc.TitleAttribute{
							Val: asciidoc.Elements{
								&asciidoc.String{
									Value: "Intro title",
								},
							},
						},
						&asciidoc.ShorthandAttribute{
							Style: &asciidoc.ShorthandStyle{
								Elements: asciidoc.Elements{
									&asciidoc.String{
										Value: "partintro",
									},
								},
							},
							ID:      nil,
							Roles:   nil,
							Options: nil,
						},
					},
					Delimiter: asciidoc.Delimiter{
						Type:   7,
						Length: 2,
					},
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "This is a part intro with a title.",
						},
						&asciidoc.NewLine{},
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
							Value: "content",
						},
						&asciidoc.NewLine{},
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
					Value: "Part 1",
				},
			},
			Level: 0,
		},
	},
}

var blocksTestShouldExcludePartintroIfNotAChildOfPart = &asciidoc.Document{
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
				&asciidoc.Paragraph{
					AttributeList: asciidoc.AttributeList{
						&asciidoc.ShorthandAttribute{
							Style: &asciidoc.ShorthandStyle{
								Elements: asciidoc.Elements{
									&asciidoc.String{
										Value: "partintro",
									},
								},
							},
							ID:      nil,
							Roles:   nil,
							Options: nil,
						},
					},
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "part intro paragraph",
						},
						&asciidoc.NewLine{},
					},
					Admonition: 0,
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Book",
				},
			},
			Level: 0,
		},
	},
}

var blocksTestShouldNotAllowPartintroUnlessDoctypeIsBook = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Paragraph{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "partintro",
							},
						},
					},
					ID:      nil,
					Roles:   nil,
					Options: nil,
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "part intro paragraph",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 0,
		},
	},
}

var blocksTestShouldAcceptPartintroOnOpenBlockWithoutTitleConvertedToDocBook = &asciidoc.Document{
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
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Book",
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
				&asciidoc.OpenBlock{
					AttributeList: asciidoc.AttributeList{
						&asciidoc.ShorthandAttribute{
							Style: &asciidoc.ShorthandStyle{
								Elements: asciidoc.Elements{
									&asciidoc.String{
										Value: "partintro",
									},
								},
							},
							ID:      nil,
							Roles:   nil,
							Options: nil,
						},
					},
					Delimiter: asciidoc.Delimiter{
						Type:   7,
						Length: 2,
					},
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "This is a part intro.",
						},
						&asciidoc.NewLine{},
						&asciidoc.EmptyLine{
							Text: "",
						},
						&asciidoc.String{
							Value: "It can have multiple paragraphs.",
						},
						&asciidoc.NewLine{},
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
							Value: "content",
						},
						&asciidoc.NewLine{},
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
					Value: "Part 1",
				},
			},
			Level: 0,
		},
	},
}

var blocksTestShouldAcceptPartintroOnOpenBlockWithTitleConvertedToDocBook = &asciidoc.Document{
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
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Book",
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
				&asciidoc.OpenBlock{
					AttributeList: asciidoc.AttributeList{
						&asciidoc.TitleAttribute{
							Val: asciidoc.Elements{
								&asciidoc.String{
									Value: "Intro title",
								},
							},
						},
						&asciidoc.ShorthandAttribute{
							Style: &asciidoc.ShorthandStyle{
								Elements: asciidoc.Elements{
									&asciidoc.String{
										Value: "partintro",
									},
								},
							},
							ID:      nil,
							Roles:   nil,
							Options: nil,
						},
					},
					Delimiter: asciidoc.Delimiter{
						Type:   7,
						Length: 2,
					},
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "This is a part intro with a title.",
						},
						&asciidoc.NewLine{},
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
							Value: "content",
						},
						&asciidoc.NewLine{},
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
					Value: "Part 1",
				},
			},
			Level: 0,
		},
	},
}

var blocksTestShouldExcludePartintroIfNotAChildOfPartConvertedToDocBook = &asciidoc.Document{
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
				&asciidoc.Paragraph{
					AttributeList: asciidoc.AttributeList{
						&asciidoc.ShorthandAttribute{
							Style: &asciidoc.ShorthandStyle{
								Elements: asciidoc.Elements{
									&asciidoc.String{
										Value: "partintro",
									},
								},
							},
							ID:      nil,
							Roles:   nil,
							Options: nil,
						},
					},
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "part intro paragraph",
						},
						&asciidoc.NewLine{},
					},
					Admonition: 0,
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Book",
				},
			},
			Level: 0,
		},
	},
}

var blocksTestShouldNotAllowPartintroUnlessDoctypeIsBookConvertedToDocBook = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Paragraph{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "partintro",
							},
						},
					},
					ID:      nil,
					Roles:   nil,
					Options: nil,
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "part intro paragraph",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 0,
		},
	},
}

var blocksTestProcessorShouldNotCrashIfSubsAreEmpty = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.LiteralBlock{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.NamedAttribute{
					Name: "subs",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: ",",
						},
					},
					Quote: 2,
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   6,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"content",
			},
		},
	},
}

var blocksTestShouldBeAbleToAppendSubsToDefaultBlockSubstitutionList = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "application",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Asciidoctor",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.LiteralBlock{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.NamedAttribute{
					Name: "subs",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "+attributes,+macros",
						},
					},
					Quote: 2,
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   6,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"{application}",
			},
		},
	},
}

var blocksTestShouldBeAbleToPrependSubsToDefaultBlockSubstitutionList = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "application",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Asciidoctor",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.LiteralBlock{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.NamedAttribute{
					Name: "subs",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "attributes+",
						},
					},
					Quote: 2,
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   6,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"{application}",
			},
		},
	},
}

var blocksTestShouldBeAbleToRemoveSubsToDefaultBlockSubstitutionList = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Paragraph{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.NamedAttribute{
					Name: "subs",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "-quotes,-replacements",
						},
					},
					Quote: 2,
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "content",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 0,
		},
	},
}

var blocksTestShouldBeAbleToPrependAppendAndRemoveSubsFromDefaultBlockSubstitutionList = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "application",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "asciidoctor",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.LiteralBlock{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.NamedAttribute{
					Name: "subs",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "attributes+,-verbatim,+specialcharacters,+macros",
						},
					},
					Quote: 2,
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   6,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"https://{application}.org[{gt}{gt}] <1>",
			},
		},
	},
}

var blocksTestShouldBeAbleToSetSubsThenModifyThem = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Paragraph{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.NamedAttribute{
					Name: "subs",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "verbatim,-callouts",
						},
					},
					Quote: 2,
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.Italic{
					AttributeList: nil,
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "hey now",
						},
					},
				},
				&asciidoc.String{
					Value: " <1>",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 0,
		},
	},
}

var blocksTestShouldNotRecognizeBlockAnchorWithIllegalIdCharacters = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "[[illegal$id,Reference Text]]",
		},
		&asciidoc.NewLine{},
		&asciidoc.Listing{
			AttributeList: nil,
			Delimiter: asciidoc.Delimiter{
				Type:   5,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"content",
			},
		},
	},
}

var blocksTestShouldNotRecognizeBlockAnchorThatStartsWithDigit = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "[[3-blind-mice]]",
		},
		&asciidoc.NewLine{},
		&asciidoc.OpenBlock{
			AttributeList: nil,
			Delimiter: asciidoc.Delimiter{
				Type:   7,
				Length: 2,
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "see how they run",
				},
				&asciidoc.NewLine{},
			},
		},
	},
}

var blocksTestShouldRecognizeBlockAnchorThatStartsWithColon = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OpenBlock{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.AnchorAttribute{
					ID: asciidoc.Elements{
						&asciidoc.String{
							Value: ":idname",
						},
					},
					Label: nil,
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   7,
				Length: 2,
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "content",
				},
				&asciidoc.NewLine{},
			},
		},
	},
}

var blocksTestShouldUseSpecifiedIdAndReftextWhenRegisteringBlockReference = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Listing{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.AnchorAttribute{
					ID: asciidoc.Elements{
						&asciidoc.String{
							Value: "debian",
						},
					},
					Label: asciidoc.Elements{
						&asciidoc.String{
							Value: "Debian Install",
						},
					},
				},
				&asciidoc.TitleAttribute{
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Installation on Debian",
						},
					},
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   5,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"$ apt-get install asciidoctor",
			},
		},
	},
}

var blocksTestShouldAllowSquareBracketsInBlockReferenceText = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "[[debian,[Debian] Install]]",
		},
		&asciidoc.NewLine{},
		&asciidoc.Listing{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.TitleAttribute{
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Installation on Debian",
						},
					},
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   5,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"$ apt-get install asciidoctor",
			},
		},
	},
}

var blocksTestShouldAllowCommaInBlockReferenceText = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Listing{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.AnchorAttribute{
					ID: asciidoc.Elements{
						&asciidoc.String{
							Value: "debian",
						},
					},
					Label: asciidoc.Elements{
						&asciidoc.String{
							Value: " Debian, Ubuntu",
						},
					},
				},
				&asciidoc.TitleAttribute{
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Installation on Debian",
						},
					},
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   5,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"$ apt-get install asciidoctor",
			},
		},
	},
}

var blocksTestShouldResolveAttributeReferenceInTitleUsingAttributeDefinedAtLocationOfBlock = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.AttributeEntry{
					Name: "foo",
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "baz",
						},
					},
				},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "intro paragraph. see ",
				},
				&asciidoc.CrossReference{
					AttributeList: nil,
					Elements:      nil,
					ID: asciidoc.Elements{
						&asciidoc.String{
							Value: "free-standing",
						},
					},
					Format: 0,
				},
				&asciidoc.String{
					Value: ".",
				},
				&asciidoc.NewLine{},
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
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.Paragraph{
					AttributeList: asciidoc.AttributeList{
						&asciidoc.TitleAttribute{
							Val: asciidoc.Elements{
								&asciidoc.String{
									Value: "foo is ",
								},
								&asciidoc.UserAttributeReference{
									Value: "foo",
								},
							},
						},
						&asciidoc.ShorthandAttribute{
							Style: nil,
							ID: &asciidoc.ShorthandID{
								Elements: asciidoc.Elements{
									&asciidoc.String{
										Value: "formal-para",
									},
								},
							},
							Roles:   nil,
							Options: nil,
						},
					},
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "paragraph with title",
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
							Style: &asciidoc.ShorthandStyle{
								Elements: asciidoc.Elements{
									&asciidoc.String{
										Value: "discrete",
									},
								},
							},
							ID: &asciidoc.ShorthandID{
								Elements: asciidoc.Elements{
									&asciidoc.String{
										Value: "free-standing",
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
							Value: "foo is still {foo}",
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

var blocksTestShouldSubstituteAttributeReferencesInReftextWhenRegisteringBlockReference = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "label-tiger",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Tiger",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.SidebarBlock{
			Delimiter: asciidoc.Delimiter{
				Type:   8,
				Length: 4,
			},
			AttributeList: asciidoc.AttributeList{
				&asciidoc.AnchorAttribute{
					ID: asciidoc.Elements{
						&asciidoc.String{
							Value: "tiger-evolution",
						},
					},
					Label: asciidoc.Elements{
						&asciidoc.String{
							Value: "Evolution of the ",
						},
						&asciidoc.UserAttributeReference{
							Value: "label-tiger",
						},
					},
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Information about the evolution of the tiger.",
				},
				&asciidoc.NewLine{},
			},
		},
	},
}

var blocksTestShouldUseSpecifiedReftextWhenRegisteringBlockReference = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Listing{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.AnchorAttribute{
					ID: asciidoc.Elements{
						&asciidoc.String{
							Value: "debian",
						},
					},
					Label: nil,
				},
				&asciidoc.NamedAttribute{
					Name: "reftext",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Debian Install",
						},
					},
					Quote: 2,
				},
				&asciidoc.TitleAttribute{
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Installation on Debian",
						},
					},
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   5,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"$ apt-get install asciidoctor",
			},
		},
	},
}
