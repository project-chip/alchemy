package tests

import (
	"testing"

	"github.com/hasty/alchemy/asciidoc"
)

func TestBlocks(t *testing.T) {
	blocksTests.run(t)
}

var blocksTests = parseTests{
	
	{"horizontal rule between blocks", "asciidoctor/blocks_test_horizontal_rule_between_blocks.adoc", horizontalRuleBetweenBlocks},
	
	{"line comment between paragraphs offset by blank lines", "asciidoctor/blocks_test_line_comment_between_paragraphs_offset_by_blank_lines.adoc", lineCommentBetweenParagraphsOffsetByBlankLines},
	
	{"adjacent line comment between paragraphs", "asciidoctor/blocks_test_adjacent_line_comment_between_paragraphs.adoc", adjacentLineCommentBetweenParagraphs},
	
	{"comment block between paragraphs offset by blank lines", "asciidoctor/blocks_test_comment_block_between_paragraphs_offset_by_blank_lines.adoc", commentBlockBetweenParagraphsOffsetByBlankLines},
	
	{"comment block between paragraphs offset by blank lines inside delimited block", "asciidoctor/blocks_test_comment_block_between_paragraphs_offset_by_blank_lines_inside_delimited_block.adoc", commentBlockBetweenParagraphsOffsetByBlankLinesInsideDelimitedBlock},
	
	{"adjacent comment block between paragraphs", "asciidoctor/blocks_test_adjacent_comment_block_between_paragraphs.adoc", adjacentCommentBlockBetweenParagraphs},
	
	{"can convert with block comment at end of document with trailing newlines", "asciidoctor/blocks_test_can_convert_with_block_comment_at_end_of_document_with_trailing_newlines.adoc", canConvertWithBlockCommentAtEndOfDocumentWithTrailingNewlines},
	
	{"trailing newlines after block comment at end of document does not create paragraph", "asciidoctor/blocks_test_trailing_newlines_after_block_comment_at_end_of_document_does_not_create_paragraph.adoc", trailingNewlinesAfterBlockCommentAtEndOfDocumentDoesNotCreateParagraph},
	
	{"line starting with three slashes should not be line comment", "asciidoctor/blocks_test_line_starting_with_three_slashes_should_not_be_line_comment.adoc", lineStartingWithThreeSlashesShouldNotBeLineComment},
	
	{"preprocessor directives should not be processed within comment block", "asciidoctor/blocks_test_preprocessor_directives_should_not_be_processed_within_comment_block.adoc", preprocessorDirectivesShouldNotBeProcessedWithinCommentBlock},
	
	{"should warn if unterminated comment block is detected in body", "asciidoctor/blocks_test_should_warn_if_unterminated_comment_block_is_detected_in_body.adoc", shouldWarnIfUnterminatedCommentBlockIsDetectedInBody},
	
	{"should warn if unterminated comment block is detected inside another block", "asciidoctor/blocks_test_should_warn_if_unterminated_comment_block_is_detected_inside_another_block.adoc", shouldWarnIfUnterminatedCommentBlockIsDetectedInsideAnotherBlock},
	
	{"preprocessor directives should not be processed within comment open block", "asciidoctor/blocks_test_preprocessor_directives_should_not_be_processed_within_comment_open_block.adoc", preprocessorDirectivesShouldNotBeProcessedWithinCommentOpenBlock},
	
	{"preprocessor directives should not be processed on subsequent lines of a comment paragraph", "asciidoctor/blocks_test_preprocessor_directives_should_not_be_processed_on_subsequent_lines_of_a_comment_paragraph.adoc", preprocessorDirectivesShouldNotBeProcessedOnSubsequentLinesOfACommentParagraph},
	
	{"comment style on open block should only skip block", "asciidoctor/blocks_test_comment_style_on_open_block_should_only_skip_block.adoc", commentStyleOnOpenBlockShouldOnlySkipBlock},
	
	{"comment style on paragraph should only skip paragraph", "asciidoctor/blocks_test_comment_style_on_paragraph_should_only_skip_paragraph.adoc", commentStyleOnParagraphShouldOnlySkipParagraph},
	
	{"comment style on paragraph should not cause adjacent block to be skipped", "asciidoctor/blocks_test_comment_style_on_paragraph_should_not_cause_adjacent_block_to_be_skipped.adoc", commentStyleOnParagraphShouldNotCauseAdjacentBlockToBeSkipped},
	
	{"should not drop content that follows skipped content inside a delimited block", "asciidoctor/blocks_test_should_not_drop_content_that_follows_skipped_content_inside_a_delimited_block.adoc", shouldNotDropContentThatFollowsSkippedContentInsideADelimitedBlock},
	
	{"should parse sidebar block", "asciidoctor/blocks_test_should_parse_sidebar_block.adoc", shouldParseSidebarBlock},
	
	{"quote block with no attribution", "asciidoctor/blocks_test_quote_block_with_no_attribution.adoc", quoteBlockWithNoAttribution},
	
	{"quote block with attribution", "asciidoctor/blocks_test_quote_block_with_attribution.adoc", quoteBlockWithAttribution},
	
	{"quote block with attribute and id and role shorthand", "asciidoctor/blocks_test_quote_block_with_attribute_and_id_and_role_shorthand.adoc", quoteBlockWithAttributeAndIdAndRoleShorthand},
	
	{"setting ID using style shorthand should not reset block style", "asciidoctor/blocks_test_setting_id_using_style_shorthand_should_not_reset_block_style.adoc", settingIdUsingStyleShorthandShouldNotResetBlockStyle},
	
	{"quote block with complex content", "asciidoctor/blocks_test_quote_block_with_complex_content.adoc", quoteBlockWithComplexContent},
	
	{"quote block with attribution converted to DocBook", "asciidoctor/blocks_test_quote_block_with_attribution_converted_to_doc_book.adoc", quoteBlockWithAttributionConvertedToDocBook},
	
	{"epigraph quote block with attribution converted to DocBook", "asciidoctor/blocks_test_epigraph_quote_block_with_attribution_converted_to_doc_book.adoc", epigraphQuoteBlockWithAttributionConvertedToDocBook},
	
	{"markdown-style quote block with single paragraph and no attribution", "asciidoctor/blocks_test_markdown_style_quote_block_with_single_paragraph_and_no_attribution.adoc", markdownStyleQuoteBlockWithSingleParagraphAndNoAttribution},
	
	{"lazy markdown-style quote block with single paragraph and no attribution", "asciidoctor/blocks_test_lazy_markdown_style_quote_block_with_single_paragraph_and_no_attribution.adoc", lazyMarkdownStyleQuoteBlockWithSingleParagraphAndNoAttribution},
	
	{"markdown-style quote block with multiple paragraphs and no attribution", "asciidoctor/blocks_test_markdown_style_quote_block_with_multiple_paragraphs_and_no_attribution.adoc", markdownStyleQuoteBlockWithMultipleParagraphsAndNoAttribution},
	
	{"markdown-style quote block with multiple blocks and no attribution", "asciidoctor/blocks_test_markdown_style_quote_block_with_multiple_blocks_and_no_attribution.adoc", markdownStyleQuoteBlockWithMultipleBlocksAndNoAttribution},
	
	{"markdown-style quote block with single paragraph and attribution", "asciidoctor/blocks_test_markdown_style_quote_block_with_single_paragraph_and_attribution.adoc", markdownStyleQuoteBlockWithSingleParagraphAndAttribution},
	
	{"markdown-style quote block with only attribution", "asciidoctor/blocks_test_markdown_style_quote_block_with_only_attribution.adoc", markdownStyleQuoteBlockWithOnlyAttribution},
	
	{"quoted paragraph-style quote block with attribution", "asciidoctor/blocks_test_quoted_paragraph_style_quote_block_with_attribution.adoc", quotedParagraphStyleQuoteBlockWithAttribution},
	
	{"should parse credit line in quoted paragraph-style quote block like positional block attributes", "asciidoctor/blocks_test_should_parse_credit_line_in_quoted_paragraph_style_quote_block_like_positional_block_attributes.adoc", shouldParseCreditLineInQuotedParagraphStyleQuoteBlockLikePositionalBlockAttributes},
	
	{"single-line verse block without attribution", "asciidoctor/blocks_test_single_line_verse_block_without_attribution.adoc", singleLineVerseBlockWithoutAttribution},
	
	{"single-line verse block with attribution", "asciidoctor/blocks_test_single_line_verse_block_with_attribution.adoc", singleLineVerseBlockWithAttribution},
	
	{"single-line verse block with attribution converted to DocBook", "asciidoctor/blocks_test_single_line_verse_block_with_attribution_converted_to_doc_book.adoc", singleLineVerseBlockWithAttributionConvertedToDocBook},
	
	{"single-line epigraph verse block with attribution converted to DocBook", "asciidoctor/blocks_test_single_line_epigraph_verse_block_with_attribution_converted_to_doc_book.adoc", singleLineEpigraphVerseBlockWithAttributionConvertedToDocBook},
	
	{"multi-stanza verse block", "asciidoctor/blocks_test_multi_stanza_verse_block.adoc", multiStanzaVerseBlock},
	
	{"verse block does not contain block elements", "asciidoctor/blocks_test_verse_block_does_not_contain_block_elements.adoc", verseBlockDoesNotContainBlockElements},
	
	{"verse should have normal subs", "asciidoctor/blocks_test_verse_should_have_normal_subs.adoc", verseShouldHaveNormalSubs},
	
	{"should not recognize callouts in a verse", "asciidoctor/blocks_test_should_not_recognize_callouts_in_a_verse.adoc", shouldNotRecognizeCalloutsInAVerse},
	
	{"should perform normal subs on a verse block", "asciidoctor/blocks_test_should_perform_normal_subs_on_a_verse_block.adoc", shouldPerformNormalSubsOnAVerseBlock},
	
	{"can convert example block", "asciidoctor/blocks_test_can_convert_example_block.adoc", canConvertExampleBlock},
	
	{"assigns sequential numbered caption to example block with title", "asciidoctor/blocks_test_assigns_sequential_numbered_caption_to_example_block_with_title.adoc", assignsSequentialNumberedCaptionToExampleBlockWithTitle},
	
	{"assigns sequential character caption to example block with title", "asciidoctor/blocks_test_assigns_sequential_character_caption_to_example_block_with_title.adoc", assignsSequentialCharacterCaptionToExampleBlockWithTitle},
	
	{"should increment counter for example even when example-number is locked by the API", "asciidoctor/blocks_test_should_increment_counter_for_example_even_when_example_number_is_locked_by_the_api.adoc", shouldIncrementCounterForExampleEvenWhenExampleNumberIsLockedByTheApi},
	
	{"should use explicit caption if specified", "asciidoctor/blocks_test_should_use_explicit_caption_if_specified.adoc", shouldUseExplicitCaptionIfSpecified},
	
	{"automatic caption can be turned off and on and modified", "asciidoctor/blocks_test_automatic_caption_can_be_turned_off_and_on_and_modified.adoc", automaticCaptionCanBeTurnedOffAndOnAndModified},
	
	{"should use explicit caption if specified even if block-specific global caption is disabled", "asciidoctor/blocks_test_should_use_explicit_caption_if_specified_even_if_block_specific_global_caption_is_disabled.adoc", shouldUseExplicitCaptionIfSpecifiedEvenIfBlockSpecificGlobalCaptionIsDisabled},
	
	{"should use global caption if specified even if block-specific global caption is disabled", "asciidoctor/blocks_test_should_use_global_caption_if_specified_even_if_block_specific_global_caption_is_disabled.adoc", shouldUseGlobalCaptionIfSpecifiedEvenIfBlockSpecificGlobalCaptionIsDisabled},
	
	{"should not process caption attribute on block that does not support a caption", "asciidoctor/blocks_test_should_not_process_caption_attribute_on_block_that_does_not_support_a_caption.adoc", shouldNotProcessCaptionAttributeOnBlockThatDoesNotSupportACaption},
	
	{"should create details/summary set if collapsible option is set", "asciidoctor/blocks_test_should_create_details_summary_set_if_collapsible_option_is_set.adoc", shouldCreateDetailsSummarySetIfCollapsibleOptionIsSet},
	
	{"should open details/summary set if collapsible and open options are set", "asciidoctor/blocks_test_should_open_details_summary_set_if_collapsible_and_open_options_are_set.adoc", shouldOpenDetailsSummarySetIfCollapsibleAndOpenOptionsAreSet},
	
	{"should add default summary element if collapsible option is set and title is not specifed", "asciidoctor/blocks_test_should_add_default_summary_element_if_collapsible_option_is_set_and_title_is_not_specifed.adoc", shouldAddDefaultSummaryElementIfCollapsibleOptionIsSetAndTitleIsNotSpecifed},
	
	{"should not allow collapsible block to increment example number", "asciidoctor/blocks_test_should_not_allow_collapsible_block_to_increment_example_number.adoc", shouldNotAllowCollapsibleBlockToIncrementExampleNumber},
	
	{"should warn if example block is not terminated", "asciidoctor/blocks_test_should_warn_if_example_block_is_not_terminated.adoc", shouldWarnIfExampleBlockIsNotTerminated},
	
	{"caption block-level attribute should be used as caption", "asciidoctor/blocks_test_caption_block_level_attribute_should_be_used_as_caption.adoc", captionBlockLevelAttributeShouldBeUsedAsCaption},
	
	{"can override caption of admonition block using document attribute", "asciidoctor/blocks_test_can_override_caption_of_admonition_block_using_document_attribute.adoc", canOverrideCaptionOfAdmonitionBlockUsingDocumentAttribute},
	
	{"blank caption document attribute should not blank admonition block caption", "asciidoctor/blocks_test_blank_caption_document_attribute_should_not_blank_admonition_block_caption.adoc", blankCaptionDocumentAttributeShouldNotBlankAdmonitionBlockCaption},
	
	{"should separate adjacent paragraphs and listing into blocks", "asciidoctor/blocks_test_should_separate_adjacent_paragraphs_and_listing_into_blocks.adoc", shouldSeparateAdjacentParagraphsAndListingIntoBlocks},
	
	{"should warn if listing block is not terminated", "asciidoctor/blocks_test_should_warn_if_listing_block_is_not_terminated.adoc", shouldWarnIfListingBlockIsNotTerminated},
	
	{"should not crash when converting verbatim block that has no lines", "asciidoctor/blocks_test_should_not_crash_when_converting_verbatim_block_that_has_no_lines.adoc", shouldNotCrashWhenConvertingVerbatimBlockThatHasNoLines},
	
	{"should preserve newlines in listing block", "asciidoctor/blocks_test_should_preserve_newlines_in_listing_block.adoc", shouldPreserveNewlinesInListingBlock},
	
	{"should preserve newlines in verse block", "asciidoctor/blocks_test_should_preserve_newlines_in_verse_block.adoc", shouldPreserveNewlinesInVerseBlock},
	
	{"should strip leading and trailing blank lines when converting verbatim block", "asciidoctor/blocks_test_should_strip_leading_and_trailing_blank_lines_when_converting_verbatim_block.adoc", shouldStripLeadingAndTrailingBlankLinesWhenConvertingVerbatimBlock},
	
	{"should remove block indent if indent attribute is 0", "asciidoctor/blocks_test_should_remove_block_indent_if_indent_attribute_is_0.adoc", shouldRemoveBlockIndentIfIndentAttributeIs0},
	
	{"should not remove block indent if indent attribute is -1", "asciidoctor/blocks_test_should_not_remove_block_indent_if_indent_attribute_is__1.adoc", shouldNotRemoveBlockIndentIfIndentAttributeIs1},
	
	{"should set block indent to value specified by indent attribute", "asciidoctor/blocks_test_should_set_block_indent_to_value_specified_by_indent_attribute.adoc", shouldSetBlockIndentToValueSpecifiedByIndentAttribute},
	
	{"should set block indent to value specified by indent document attribute", "asciidoctor/blocks_test_should_set_block_indent_to_value_specified_by_indent_document_attribute.adoc", shouldSetBlockIndentToValueSpecifiedByIndentDocumentAttribute},
	
	{"literal block should honor nowrap option", "asciidoctor/blocks_test_literal_block_should_honor_nowrap_option.adoc", literalBlockShouldHonorNowrapOption},
	
	{"literal block should set nowrap class if prewrap document attribute is disabled", "asciidoctor/blocks_test_literal_block_should_set_nowrap_class_if_prewrap_document_attribute_is_disabled.adoc", literalBlockShouldSetNowrapClassIfPrewrapDocumentAttributeIsDisabled},
	
	{"should preserve guard in front of callout if icons are not enabled", "asciidoctor/blocks_test_should_preserve_guard_in_front_of_callout_if_icons_are_not_enabled.adoc", shouldPreserveGuardInFrontOfCalloutIfIconsAreNotEnabled},
	
	{"should preserve guard around callout if icons are not enabled", "asciidoctor/blocks_test_should_preserve_guard_around_callout_if_icons_are_not_enabled.adoc", shouldPreserveGuardAroundCalloutIfIconsAreNotEnabled},
	
	{"literal block should honor explicit subs list", "asciidoctor/blocks_test_literal_block_should_honor_explicit_subs_list.adoc", literalBlockShouldHonorExplicitSubsList},
	
	{"should be able to disable callouts for literal block", "asciidoctor/blocks_test_should_be_able_to_disable_callouts_for_literal_block.adoc", shouldBeAbleToDisableCalloutsForLiteralBlock},
	
	{"listing block should honor explicit subs list", "asciidoctor/blocks_test_listing_block_should_honor_explicit_subs_list.adoc", listingBlockShouldHonorExplicitSubsList},
	
	{"should not mangle array that contains formatted text with role in listing block with quotes sub enabled", "asciidoctor/blocks_test_should_not_mangle_array_that_contains_formatted_text_with_role_in_listing_block_with_quotes_sub_enabled.adoc", shouldNotMangleArrayThatContainsFormattedTextWithRoleInListingBlockWithQuotesSubEnabled},
	
	{"first character of block title may be a period if not followed by space", "asciidoctor/blocks_test_first_character_of_block_title_may_be_a_period_if_not_followed_by_space.adoc", firstCharacterOfBlockTitleMayBeAPeriodIfNotFollowedBySpace},
	
	{"listing block without title should generate screen element in docbook", "asciidoctor/blocks_test_listing_block_without_title_should_generate_screen_element_in_docbook.adoc", listingBlockWithoutTitleShouldGenerateScreenElementInDocbook},
	
	{"listing block with title should generate screen element inside formalpara element in docbook", "asciidoctor/blocks_test_listing_block_with_title_should_generate_screen_element_inside_formalpara_element_in_docbook.adoc", listingBlockWithTitleShouldGenerateScreenElementInsideFormalparaElementInDocbook},
	
	{"should not prepend caption to title of listing block with title if listing-caption attribute is not set", "asciidoctor/blocks_test_should_not_prepend_caption_to_title_of_listing_block_with_title_if_listing_caption_attribute_is_not_set.adoc", shouldNotPrependCaptionToTitleOfListingBlockWithTitleIfListingCaptionAttributeIsNotSet},
	
	{"should prepend caption specified by listing-caption attribute and number to title of listing block with title", "asciidoctor/blocks_test_should_prepend_caption_specified_by_listing_caption_attribute_and_number_to_title_of_listing_block_with_title.adoc", shouldPrependCaptionSpecifiedByListingCaptionAttributeAndNumberToTitleOfListingBlockWithTitle},
	
	{"should prepend caption specified by caption attribute on listing block even if listing-caption attribute is not set", "asciidoctor/blocks_test_should_prepend_caption_specified_by_caption_attribute_on_listing_block_even_if_listing_caption_attribute_is_not_set.adoc", shouldPrependCaptionSpecifiedByCaptionAttributeOnListingBlockEvenIfListingCaptionAttributeIsNotSet},
	
	{"listing block without an explicit style and with a second positional argument should be promoted to a source block", "asciidoctor/blocks_test_listing_block_without_an_explicit_style_and_with_a_second_positional_argument_should_be_promoted_to_a_source_block.adoc", listingBlockWithoutAnExplicitStyleAndWithASecondPositionalArgumentShouldBePromotedToASourceBlock},
	
	{"listing block without an explicit style should be promoted to a source block if source-language is set", "asciidoctor/blocks_test_listing_block_without_an_explicit_style_should_be_promoted_to_a_source_block_if_source_language_is_set.adoc", listingBlockWithoutAnExplicitStyleShouldBePromotedToASourceBlockIfSourceLanguageIsSet},
	
	{"listing block with an explicit style and a second positional argument should not be promoted to a source block", "asciidoctor/blocks_test_listing_block_with_an_explicit_style_and_a_second_positional_argument_should_not_be_promoted_to_a_source_block.adoc", listingBlockWithAnExplicitStyleAndASecondPositionalArgumentShouldNotBePromotedToASourceBlock},
	
	{"listing block with an explicit style should not be promoted to a source block if source-language is set", "asciidoctor/blocks_test_listing_block_with_an_explicit_style_should_not_be_promoted_to_a_source_block_if_source_language_is_set.adoc", listingBlockWithAnExplicitStyleShouldNotBePromotedToASourceBlockIfSourceLanguageIsSet},
	
	{"source block with no title or language should generate screen element in docbook", "asciidoctor/blocks_test_source_block_with_no_title_or_language_should_generate_screen_element_in_docbook.adoc", sourceBlockWithNoTitleOrLanguageShouldGenerateScreenElementInDocbook},
	
	{"source block with title and no language should generate screen element inside formalpara element for docbook", "asciidoctor/blocks_test_source_block_with_title_and_no_language_should_generate_screen_element_inside_formalpara_element_for_docbook.adoc", sourceBlockWithTitleAndNoLanguageShouldGenerateScreenElementInsideFormalparaElementForDocbook},
	
	{"can convert open block", "asciidoctor/blocks_test_can_convert_open_block.adoc", canConvertOpenBlock},
	
	{"open block can contain another block", "asciidoctor/blocks_test_open_block_can_contain_another_block.adoc", openBlockCanContainAnotherBlock},
	
	{"should transfer id and reftext on open block to DocBook output", "asciidoctor/blocks_test_should_transfer_id_and_reftext_on_open_block_to_doc_book_output.adoc", shouldTransferIdAndReftextOnOpenBlockToDocBookOutput},
	
	{"should transfer id and reftext on open paragraph to DocBook output", "asciidoctor/blocks_test_should_transfer_id_and_reftext_on_open_paragraph_to_doc_book_output.adoc", shouldTransferIdAndReftextOnOpenParagraphToDocBookOutput},
	
	{"should transfer title on open block to DocBook output", "asciidoctor/blocks_test_should_transfer_title_on_open_block_to_doc_book_output.adoc", shouldTransferTitleOnOpenBlockToDocBookOutput},
	
	{"should transfer title on open paragraph to DocBook output", "asciidoctor/blocks_test_should_transfer_title_on_open_paragraph_to_doc_book_output.adoc", shouldTransferTitleOnOpenParagraphToDocBookOutput},
	
	{"should transfer role on open block to DocBook output", "asciidoctor/blocks_test_should_transfer_role_on_open_block_to_doc_book_output.adoc", shouldTransferRoleOnOpenBlockToDocBookOutput},
	
	{"should transfer role on open paragraph to DocBook output", "asciidoctor/blocks_test_should_transfer_role_on_open_paragraph_to_doc_book_output.adoc", shouldTransferRoleOnOpenParagraphToDocBookOutput},
	
	{"can parse a passthrough block", "asciidoctor/blocks_test_can_parse_a_passthrough_block.adoc", canParseAPassthroughBlock},
	
	{"does not perform subs on a passthrough block by default", "asciidoctor/blocks_test_does_not_perform_subs_on_a_passthrough_block_by_default.adoc", doesNotPerformSubsOnAPassthroughBlockByDefault},
	
	{"does not perform subs on a passthrough block with pass style by default", "asciidoctor/blocks_test_does_not_perform_subs_on_a_passthrough_block_with_pass_style_by_default.adoc", doesNotPerformSubsOnAPassthroughBlockWithPassStyleByDefault},
	
	{"passthrough block honors explicit subs list", "asciidoctor/blocks_test_passthrough_block_honors_explicit_subs_list.adoc", passthroughBlockHonorsExplicitSubsList},
	
	{"should strip leading and trailing blank lines when converting raw block", "asciidoctor/blocks_test_should_strip_leading_and_trailing_blank_lines_when_converting_raw_block.adoc", shouldStripLeadingAndTrailingBlankLinesWhenConvertingRawBlock},
	
	{"should not crash when converting stem block that has no lines", "asciidoctor/blocks_test_should_not_crash_when_converting_stem_block_that_has_no_lines.adoc", shouldNotCrashWhenConvertingStemBlockThatHasNoLines},
	
	{"should return content as empty string for stem or pass block that has no lines", "asciidoctor/blocks_test_should_return_content_as_empty_string_for_stem_or_pass_block_that_has_no_lines.adoc", shouldReturnContentAsEmptyStringForStemOrPassBlockThatHasNoLines},
	
	{"should not add LaTeX math delimiters around latexmath block content if already present", "asciidoctor/blocks_test_should_not_add_la_te_x_math_delimiters_around_latexmath_block_content_if_already_present.adoc", shouldNotAddLaTeXMathDelimitersAroundLatexmathBlockContentIfAlreadyPresent},
	
	{"should display latexmath block in alt of equation in DocBook backend", "asciidoctor/blocks_test_should_display_latexmath_block_in_alt_of_equation_in_doc_book_backend.adoc", shouldDisplayLatexmathBlockInAltOfEquationInDocBookBackend},
	
	{"should set autoNumber option for latexmath to none by default", "asciidoctor/blocks_test_should_set_auto_number_option_for_latexmath_to_none_by_default.adoc", shouldSetAutoNumberOptionForLatexmathToNoneByDefault},
	
	{"should set autoNumber option for latexmath to none if eqnums is set to none", "asciidoctor/blocks_test_should_set_auto_number_option_for_latexmath_to_none_if_eqnums_is_set_to_none.adoc", shouldSetAutoNumberOptionForLatexmathToNoneIfEqnumsIsSetToNone},
	
	{"should set autoNumber option for latexmath to AMS if eqnums is set", "asciidoctor/blocks_test_should_set_auto_number_option_for_latexmath_to_ams_if_eqnums_is_set.adoc", shouldSetAutoNumberOptionForLatexmathToAmsIfEqnumsIsSet},
	
	{"should set autoNumber option for latexmath to all if eqnums is set to all", "asciidoctor/blocks_test_should_set_auto_number_option_for_latexmath_to_all_if_eqnums_is_set_to_all.adoc", shouldSetAutoNumberOptionForLatexmathToAllIfEqnumsIsSetToAll},
	
	{"should not split equation in AsciiMath block at single newline", "asciidoctor/blocks_test_should_not_split_equation_in_ascii_math_block_at_single_newline.adoc", shouldNotSplitEquationInAsciiMathBlockAtSingleNewline},
	
	{"should split equation in AsciiMath block at escaped newline", "asciidoctor/blocks_test_should_split_equation_in_ascii_math_block_at_escaped_newline.adoc", shouldSplitEquationInAsciiMathBlockAtEscapedNewline},
	
	{"should split equation in AsciiMath block at sequence of escaped newlines", "asciidoctor/blocks_test_should_split_equation_in_ascii_math_block_at_sequence_of_escaped_newlines.adoc", shouldSplitEquationInAsciiMathBlockAtSequenceOfEscapedNewlines},
	
	{"should split equation in AsciiMath block at newline sequence and preserve breaks", "asciidoctor/blocks_test_should_split_equation_in_ascii_math_block_at_newline_sequence_and_preserve_breaks.adoc", shouldSplitEquationInAsciiMathBlockAtNewlineSequenceAndPreserveBreaks},
	
	{"should add AsciiMath delimiters around asciimath block content", "asciidoctor/blocks_test_should_add_ascii_math_delimiters_around_asciimath_block_content.adoc", shouldAddAsciiMathDelimitersAroundAsciimathBlockContent},
	
	{"should not add AsciiMath delimiters around asciimath block content if already present", "asciidoctor/blocks_test_should_not_add_ascii_math_delimiters_around_asciimath_block_content_if_already_present.adoc", shouldNotAddAsciiMathDelimitersAroundAsciimathBlockContentIfAlreadyPresent},
	
	{"should convert contents of asciimath block to MathML in DocBook output if asciimath gem is available", "asciidoctor/blocks_test_should_convert_contents_of_asciimath_block_to_math_ml_in_doc_book_output_if_asciimath_gem_is_available.adoc", shouldConvertContentsOfAsciimathBlockToMathMlInDocBookOutputIfAsciimathGemIsAvailable},
	
	{"should output title for latexmath block if defined", "asciidoctor/blocks_test_should_output_title_for_latexmath_block_if_defined.adoc", shouldOutputTitleForLatexmathBlockIfDefined},
	
	{"should output title for asciimath block if defined", "asciidoctor/blocks_test_should_output_title_for_asciimath_block_if_defined.adoc", shouldOutputTitleForAsciimathBlockIfDefined},
	
	{"should add AsciiMath delimiters around stem block content if stem attribute is asciimath, empty, or not set", "asciidoctor/blocks_test_should_add_ascii_math_delimiters_around_stem_block_content_if_stem_attribute_is_asciimath,_empty,_or_not_set.adoc", shouldAddAsciiMathDelimitersAroundStemBlockContentIfStemAttributeIsAsciimathEmptyOrNotSet},
	
	{"should add LaTeX math delimiters around stem block content if stem attribute is latexmath, latex, or tex", "asciidoctor/blocks_test_should_add_la_te_x_math_delimiters_around_stem_block_content_if_stem_attribute_is_latexmath,_latex,_or_tex.adoc", shouldAddLaTeXMathDelimitersAroundStemBlockContentIfStemAttributeIsLatexmathLatexOrTex},
	
	{"should allow stem style to be set using second positional argument of block attributes", "asciidoctor/blocks_test_should_allow_stem_style_to_be_set_using_second_positional_argument_of_block_attributes.adoc", shouldAllowStemStyleToBeSetUsingSecondPositionalArgumentOfBlockAttributes},
	
	{"should not warn if block style is unknown", "asciidoctor/blocks_test_should_not_warn_if_block_style_is_unknown.adoc", shouldNotWarnIfBlockStyleIsUnknown},
	
	{"should log debug message if block style is unknown and debug level is enabled", "asciidoctor/blocks_test_should_log_debug_message_if_block_style_is_unknown_and_debug_level_is_enabled.adoc", shouldLogDebugMessageIfBlockStyleIsUnknownAndDebugLevelIsEnabled},
	
	{"block title above section gets carried over to first block in section", "asciidoctor/blocks_test_block_title_above_section_gets_carried_over_to_first_block_in_section.adoc", blockTitleAboveSectionGetsCarriedOverToFirstBlockInSection},
	
	{"block title above document title demotes document title to a section title", "asciidoctor/blocks_test_block_title_above_document_title_demotes_document_title_to_a_section_title.adoc", blockTitleAboveDocumentTitleDemotesDocumentTitleToASectionTitle},
	
	{"block title above document title gets carried over to first block in first section if no preamble", "asciidoctor/blocks_test_block_title_above_document_title_gets_carried_over_to_first_block_in_first_section_if_no_preamble.adoc", blockTitleAboveDocumentTitleGetsCarriedOverToFirstBlockInFirstSectionIfNoPreamble},
	
	{"should apply substitutions to a block title in normal order", "asciidoctor/blocks_test_should_apply_substitutions_to_a_block_title_in_normal_order.adoc", shouldApplySubstitutionsToABlockTitleInNormalOrder},
	
	{"empty attribute list should not appear in output", "asciidoctor/blocks_test_empty_attribute_list_should_not_appear_in_output.adoc", emptyAttributeListShouldNotAppearInOutput},
	
	{"empty block anchor should not appear in output", "asciidoctor/blocks_test_empty_block_anchor_should_not_appear_in_output.adoc", emptyBlockAnchorShouldNotAppearInOutput},
	
	{"can convert block image with alt text defined in macro", "asciidoctor/blocks_test_can_convert_block_image_with_alt_text_defined_in_macro.adoc", canConvertBlockImageWithAltTextDefinedInMacro},
	
	{"converts SVG image with alt text using img element when safe mode is secure", "asciidoctor/blocks_test_converts_svg_image_with_alt_text_using_img_element_when_safe_mode_is_secure.adoc", convertsSvgImageWithAltTextUsingImgElementWhenSafeModeIsSecure},
	
	{"inserts fallback image for SVG inside object element using same dimensions", "asciidoctor/blocks_test_inserts_fallback_image_for_svg_inside_object_element_using_same_dimensions.adoc", insertsFallbackImageForSvgInsideObjectElementUsingSameDimensions},
	
	{"detects SVG image URI that contains a query string", "asciidoctor/blocks_test_detects_svg_image_uri_that_contains_a_query_string.adoc", detectsSvgImageUriThatContainsAQueryString},
	
	{"detects SVG image when format attribute is svg", "asciidoctor/blocks_test_detects_svg_image_when_format_attribute_is_svg.adoc", detectsSvgImageWhenFormatAttributeIsSvg},
	
	{"converts to inline SVG image when inline option is set on block", "asciidoctor/blocks_test_converts_to_inline_svg_image_when_inline_option_is_set_on_block.adoc", convertsToInlineSvgImageWhenInlineOptionIsSetOnBlock},
	
	{"should ignore link attribute if value is self and image target is inline SVG", "asciidoctor/blocks_test_should_ignore_link_attribute_if_value_is_self_and_image_target_is_inline_svg.adoc", shouldIgnoreLinkAttributeIfValueIsSelfAndImageTargetIsInlineSvg},
	
	{"should honor percentage width for SVG image with inline option", "asciidoctor/blocks_test_should_honor_percentage_width_for_svg_image_with_inline_option.adoc", shouldHonorPercentageWidthForSvgImageWithInlineOption},
	
	{"should not crash if explicit width on SVG image block is an integer", "asciidoctor/blocks_test_should_not_crash_if_explicit_width_on_svg_image_block_is_an_integer.adoc", shouldNotCrashIfExplicitWidthOnSvgImageBlockIsAnInteger},
	
	{"converts to inline SVG image when inline option is set on block and data-uri is set on document", "asciidoctor/blocks_test_converts_to_inline_svg_image_when_inline_option_is_set_on_block_and_data_uri_is_set_on_document.adoc", convertsToInlineSvgImageWhenInlineOptionIsSetOnBlockAndDataUriIsSetOnDocument},
	
	{"should not throw exception if SVG to inline is empty", "asciidoctor/blocks_test_should_not_throw_exception_if_svg_to_inline_is_empty.adoc", shouldNotThrowExceptionIfSvgToInlineIsEmpty},
	
	{"can convert block image with alt text defined in macro containing square bracket", "asciidoctor/blocks_test_can_convert_block_image_with_alt_text_defined_in_macro_containing_square_bracket.adoc", canConvertBlockImageWithAltTextDefinedInMacroContainingSquareBracket},
	
	{"alt text in macro overrides alt text above macro", "asciidoctor/blocks_test_alt_text_in_macro_overrides_alt_text_above_macro.adoc", altTextInMacroOverridesAltTextAboveMacro},
	
	{"should substitute attribute references in alt text defined in image block macro", "asciidoctor/blocks_test_should_substitute_attribute_references_in_alt_text_defined_in_image_block_macro.adoc", shouldSubstituteAttributeReferencesInAltTextDefinedInImageBlockMacro},
	
	{"should set direction CSS class on image if float attribute is set", "asciidoctor/blocks_test_should_set_direction_css_class_on_image_if_float_attribute_is_set.adoc", shouldSetDirectionCssClassOnImageIfFloatAttributeIsSet},
	
	{"should set text alignment CSS class on image if align attribute is set", "asciidoctor/blocks_test_should_set_text_alignment_css_class_on_image_if_align_attribute_is_set.adoc", shouldSetTextAlignmentCssClassOnImageIfAlignAttributeIsSet},
	
	{"style attribute is dropped from image macro", "asciidoctor/blocks_test_style_attribute_is_dropped_from_image_macro.adoc", styleAttributeIsDroppedFromImageMacro},
	
	{"should auto-generate alt text for block image if alt text is not specified", "asciidoctor/blocks_test_should_auto_generate_alt_text_for_block_image_if_alt_text_is_not_specified.adoc", shouldAutoGenerateAltTextForBlockImageIfAltTextIsNotSpecified},
	
	{"can convert block image with link to self", "asciidoctor/blocks_test_can_convert_block_image_with_link_to_self.adoc", canConvertBlockImageWithLinkToSelf},
	
	{"adds rel=noopener attribute to block image with link that targets _blank window", "asciidoctor/blocks_test_adds_rel=noopener_attribute_to_block_image_with_link_that_targets__blank_window.adoc", addsRelnoopenerAttributeToBlockImageWithLinkThatTargetsBlankWindow},
	
	{"can convert block image with explicit caption", "asciidoctor/blocks_test_can_convert_block_image_with_explicit_caption.adoc", canConvertBlockImageWithExplicitCaption},
	
	{"can align image in DocBook backend", "asciidoctor/blocks_test_can_align_image_in_doc_book_backend.adoc", canAlignImageInDocBookBackend},
	
	{"should not drop line if image target is missing attribute reference and attribute-missing is drop", "asciidoctor/blocks_test_should_not_drop_line_if_image_target_is_missing_attribute_reference_and_attribute_missing_is_drop.adoc", shouldNotDropLineIfImageTargetIsMissingAttributeReferenceAndAttributeMissingIsDrop},
	
	{"drops line if image target is missing attribute reference and attribute-missing is drop-line", "asciidoctor/blocks_test_drops_line_if_image_target_is_missing_attribute_reference_and_attribute_missing_is_drop_line.adoc", dropsLineIfImageTargetIsMissingAttributeReferenceAndAttributeMissingIsDropLine},
	
	{"should not drop line if image target resolves to blank and attribute-missing is drop-line", "asciidoctor/blocks_test_should_not_drop_line_if_image_target_resolves_to_blank_and_attribute_missing_is_drop_line.adoc", shouldNotDropLineIfImageTargetResolvesToBlankAndAttributeMissingIsDropLine},
	
	{"dropped image does not break processing of following section and attribute-missing is drop-line", "asciidoctor/blocks_test_dropped_image_does_not_break_processing_of_following_section_and_attribute_missing_is_drop_line.adoc", droppedImageDoesNotBreakProcessingOfFollowingSectionAndAttributeMissingIsDropLine},
	
	{"should pass through image that references uri", "asciidoctor/blocks_test_should_pass_through_image_that_references_uri.adoc", shouldPassThroughImageThatReferencesUri},
	
	{"should encode spaces in image target if value is a URI", "asciidoctor/blocks_test_should_encode_spaces_in_image_target_if_value_is_a_uri.adoc", shouldEncodeSpacesInImageTargetIfValueIsAUri},
	
	{"embeds base64-encoded data uri for image when data-uri attribute is set", "asciidoctor/blocks_test_embeds_base_64_encoded_data_uri_for_image_when_data_uri_attribute_is_set.adoc", embedsBase64EncodedDataUriForImageWhenDataUriAttributeIsSet},
	
	{"embeds SVG image with image/svg+xml mimetype when file extension is .svg", "asciidoctor/blocks_test_embeds_svg_image_with_image_svg+xml_mimetype_when_file_extension_is__svg.adoc", embedsSvgImageWithImageSvgxmlMimetypeWhenFileExtensionIsSvg},
	
	{"should link to data URI if value of link attribute is self and image is embedded", "asciidoctor/blocks_test_should_link_to_data_uri_if_value_of_link_attribute_is_self_and_image_is_embedded.adoc", shouldLinkToDataUriIfValueOfLinkAttributeIsSelfAndImageIsEmbedded},
	
	{"embeds empty base64-encoded data uri for unreadable image when data-uri attribute is set", "asciidoctor/blocks_test_embeds_empty_base_64_encoded_data_uri_for_unreadable_image_when_data_uri_attribute_is_set.adoc", embedsEmptyBase64EncodedDataUriForUnreadableImageWhenDataUriAttributeIsSet},
	
	{"embeds base64-encoded data uri with application/octet-stream mimetype when file extension is missing", "asciidoctor/blocks_test_embeds_base_64_encoded_data_uri_with_application_octet_stream_mimetype_when_file_extension_is_missing.adoc", embedsBase64EncodedDataUriWithApplicationOctetStreamMimetypeWhenFileExtensionIsMissing},
	
	{"can handle embedded data uri images", "asciidoctor/blocks_test_can_handle_embedded_data_uri_images.adoc", canHandleEmbeddedDataUriImages},
	
	{"cleans reference to ancestor directories in imagesdir before reading image if safe mode level is at least SAFE", "asciidoctor/blocks_test_cleans_reference_to_ancestor_directories_in_imagesdir_before_reading_image_if_safe_mode_level_is_at_least_safe.adoc", cleansReferenceToAncestorDirectoriesInImagesdirBeforeReadingImageIfSafeModeLevelIsAtLeastSafe},
	
	{"cleans reference to ancestor directories in target before reading image if safe mode level is at least SAFE", "asciidoctor/blocks_test_cleans_reference_to_ancestor_directories_in_target_before_reading_image_if_safe_mode_level_is_at_least_safe.adoc", cleansReferenceToAncestorDirectoriesInTargetBeforeReadingImageIfSafeModeLevelIsAtLeastSafe},
	
	{"should detect and convert video macro", "asciidoctor/blocks_test_should_detect_and_convert_video_macro.adoc", shouldDetectAndConvertVideoMacro},
	
	{"video macro should not use imagesdir attribute to resolve target if target is a URL", "asciidoctor/blocks_test_video_macro_should_not_use_imagesdir_attribute_to_resolve_target_if_target_is_a_url.adoc", videoMacroShouldNotUseImagesdirAttributeToResolveTargetIfTargetIsAUrl},
	
	{"video macro should output custom HTML with iframe for vimeo service", "asciidoctor/blocks_test_video_macro_should_output_custom_html_with_iframe_for_vimeo_service.adoc", videoMacroShouldOutputCustomHtmlWithIframeForVimeoService},
	
	{"audio macro should not use imagesdir attribute to resolve target if target is a URL", "asciidoctor/blocks_test_audio_macro_should_not_use_imagesdir_attribute_to_resolve_target_if_target_is_a_url.adoc", audioMacroShouldNotUseImagesdirAttributeToResolveTargetIfTargetIsAUrl},
	
	{"audio macro should honor all options", "asciidoctor/blocks_test_audio_macro_should_honor_all_options.adoc", audioMacroShouldHonorAllOptions},
	
	{"can resolve icon relative to custom iconsdir", "asciidoctor/blocks_test_can_resolve_icon_relative_to_custom_iconsdir.adoc", canResolveIconRelativeToCustomIconsdir},
	
	{"should add file extension to custom icon if not specified", "asciidoctor/blocks_test_should_add_file_extension_to_custom_icon_if_not_specified.adoc", shouldAddFileExtensionToCustomIconIfNotSpecified},
	
	{"should allow icontype to be specified when using built-in admonition icon", "asciidoctor/blocks_test_should_allow_icontype_to_be_specified_when_using_built_in_admonition_icon.adoc", shouldAllowIcontypeToBeSpecifiedWhenUsingBuiltInAdmonitionIcon},
	
	{"embeds base64-encoded data uri of icon when data-uri attribute is set and safe mode level is less than SECURE", "asciidoctor/blocks_test_embeds_base_64_encoded_data_uri_of_icon_when_data_uri_attribute_is_set_and_safe_mode_level_is_less_than_secure.adoc", embedsBase64EncodedDataUriOfIconWhenDataUriAttributeIsSetAndSafeModeLevelIsLessThanSecure},
	
	{"should embed base64-encoded data uri of custom icon when data-uri attribute is set", "asciidoctor/blocks_test_should_embed_base_64_encoded_data_uri_of_custom_icon_when_data_uri_attribute_is_set.adoc", shouldEmbedBase64EncodedDataUriOfCustomIconWhenDataUriAttributeIsSet},
	
	{"does not embed base64-encoded data uri of icon when safe mode level is SECURE or greater", "asciidoctor/blocks_test_does_not_embed_base_64_encoded_data_uri_of_icon_when_safe_mode_level_is_secure_or_greater.adoc", doesNotEmbedBase64EncodedDataUriOfIconWhenSafeModeLevelIsSecureOrGreater},
	
	{"cleans reference to ancestor directories before reading icon if safe mode level is at least SAFE", "asciidoctor/blocks_test_cleans_reference_to_ancestor_directories_before_reading_icon_if_safe_mode_level_is_at_least_safe.adoc", cleansReferenceToAncestorDirectoriesBeforeReadingIconIfSafeModeLevelIsAtLeastSafe},
	
	{"should import Font Awesome and use font-based icons when value of icons attribute is font", "asciidoctor/blocks_test_should_import_font_awesome_and_use_font_based_icons_when_value_of_icons_attribute_is_font.adoc", shouldImportFontAwesomeAndUseFontBasedIconsWhenValueOfIconsAttributeIsFont},
	
	{"font-based icon should not override icon specified on admonition", "asciidoctor/blocks_test_font_based_icon_should_not_override_icon_specified_on_admonition.adoc", fontBasedIconShouldNotOverrideIconSpecifiedOnAdmonition},
	
	{"should use http uri scheme for assets when asset-uri-scheme is http", "asciidoctor/blocks_test_should_use_http_uri_scheme_for_assets_when_asset_uri_scheme_is_http.adoc", shouldUseHttpUriSchemeForAssetsWhenAssetUriSchemeIsHttp},
	
	{"should use no uri scheme for assets when asset-uri-scheme is blank", "asciidoctor/blocks_test_should_use_no_uri_scheme_for_assets_when_asset_uri_scheme_is_blank.adoc", shouldUseNoUriSchemeForAssetsWhenAssetUriSchemeIsBlank},
	
	{"restricts access to ancestor directories when safe mode level is at least SAFE", "asciidoctor/blocks_test_restricts_access_to_ancestor_directories_when_safe_mode_level_is_at_least_safe.adoc", restrictsAccessToAncestorDirectoriesWhenSafeModeLevelIsAtLeastSafe},
	
	{"should not recognize fenced code blocks with more than three delimiters", "asciidoctor/blocks_test_should_not_recognize_fenced_code_blocks_with_more_than_three_delimiters.adoc", shouldNotRecognizeFencedCodeBlocksWithMoreThanThreeDelimiters},
	
	{"should support fenced code blocks with languages", "asciidoctor/blocks_test_should_support_fenced_code_blocks_with_languages.adoc", shouldSupportFencedCodeBlocksWithLanguages},
	
	{"should support fenced code blocks with languages and numbering", "asciidoctor/blocks_test_should_support_fenced_code_blocks_with_languages_and_numbering.adoc", shouldSupportFencedCodeBlocksWithLanguagesAndNumbering},
	
	{"should allow source style to be specified on literal block", "asciidoctor/blocks_test_should_allow_source_style_to_be_specified_on_literal_block.adoc", shouldAllowSourceStyleToBeSpecifiedOnLiteralBlock},
	
	{"should allow source style and language to be specified on literal block", "asciidoctor/blocks_test_should_allow_source_style_and_language_to_be_specified_on_literal_block.adoc", shouldAllowSourceStyleAndLanguageToBeSpecifiedOnLiteralBlock},
	
	{"should make abstract on open block without title a quote block for article", "asciidoctor/blocks_test_should_make_abstract_on_open_block_without_title_a_quote_block_for_article.adoc", shouldMakeAbstractOnOpenBlockWithoutTitleAQuoteBlockForArticle},
	
	{"should make abstract on open block with title a quote block with title for article", "asciidoctor/blocks_test_should_make_abstract_on_open_block_with_title_a_quote_block_with_title_for_article.adoc", shouldMakeAbstractOnOpenBlockWithTitleAQuoteBlockWithTitleForArticle},
	
	{"should allow abstract in document with title if doctype is book", "asciidoctor/blocks_test_should_allow_abstract_in_document_with_title_if_doctype_is_book.adoc", shouldAllowAbstractInDocumentWithTitleIfDoctypeIsBook},
	
	{"should not allow abstract as direct child of document if doctype is book", "asciidoctor/blocks_test_should_not_allow_abstract_as_direct_child_of_document_if_doctype_is_book.adoc", shouldNotAllowAbstractAsDirectChildOfDocumentIfDoctypeIsBook},
	
	{"should make abstract on open block without title converted to DocBook", "asciidoctor/blocks_test_should_make_abstract_on_open_block_without_title_converted_to_doc_book.adoc", shouldMakeAbstractOnOpenBlockWithoutTitleConvertedToDocBook},
	
	{"should make abstract on open block with title converted to DocBook", "asciidoctor/blocks_test_should_make_abstract_on_open_block_with_title_converted_to_doc_book.adoc", shouldMakeAbstractOnOpenBlockWithTitleConvertedToDocBook},
	
	{"should allow abstract in document with title if doctype is book converted to DocBook", "asciidoctor/blocks_test_should_allow_abstract_in_document_with_title_if_doctype_is_book_converted_to_doc_book.adoc", shouldAllowAbstractInDocumentWithTitleIfDoctypeIsBookConvertedToDocBook},
	
	{"should not allow abstract as direct child of document if doctype is book converted to DocBook", "asciidoctor/blocks_test_should_not_allow_abstract_as_direct_child_of_document_if_doctype_is_book_converted_to_doc_book.adoc", shouldNotAllowAbstractAsDirectChildOfDocumentIfDoctypeIsBookConvertedToDocBook},
	
	{"should accept partintro on open block without title", "asciidoctor/blocks_test_should_accept_partintro_on_open_block_without_title.adoc", shouldAcceptPartintroOnOpenBlockWithoutTitle},
	
	{"should accept partintro on open block with title", "asciidoctor/blocks_test_should_accept_partintro_on_open_block_with_title.adoc", shouldAcceptPartintroOnOpenBlockWithTitle},
	
	{"should exclude partintro if not a child of part", "asciidoctor/blocks_test_should_exclude_partintro_if_not_a_child_of_part.adoc", shouldExcludePartintroIfNotAChildOfPart},
	
	{"should not allow partintro unless doctype is book", "asciidoctor/blocks_test_should_not_allow_partintro_unless_doctype_is_book.adoc", shouldNotAllowPartintroUnlessDoctypeIsBook},
	
	{"should accept partintro on open block without title converted to DocBook", "asciidoctor/blocks_test_should_accept_partintro_on_open_block_without_title_converted_to_doc_book.adoc", shouldAcceptPartintroOnOpenBlockWithoutTitleConvertedToDocBook},
	
	{"should accept partintro on open block with title converted to DocBook", "asciidoctor/blocks_test_should_accept_partintro_on_open_block_with_title_converted_to_doc_book.adoc", shouldAcceptPartintroOnOpenBlockWithTitleConvertedToDocBook},
	
	{"should exclude partintro if not a child of part converted to DocBook", "asciidoctor/blocks_test_should_exclude_partintro_if_not_a_child_of_part_converted_to_doc_book.adoc", shouldExcludePartintroIfNotAChildOfPartConvertedToDocBook},
	
	{"should not allow partintro unless doctype is book converted to DocBook", "asciidoctor/blocks_test_should_not_allow_partintro_unless_doctype_is_book_converted_to_doc_book.adoc", shouldNotAllowPartintroUnlessDoctypeIsBookConvertedToDocBook},
	
	{"processor should not crash if subs are empty", "asciidoctor/blocks_test_processor_should_not_crash_if_subs_are_empty.adoc", processorShouldNotCrashIfSubsAreEmpty},
	
	{"should be able to append subs to default block substitution list", "asciidoctor/blocks_test_should_be_able_to_append_subs_to_default_block_substitution_list.adoc", shouldBeAbleToAppendSubsToDefaultBlockSubstitutionList},
	
	{"should be able to prepend subs to default block substitution list", "asciidoctor/blocks_test_should_be_able_to_prepend_subs_to_default_block_substitution_list.adoc", shouldBeAbleToPrependSubsToDefaultBlockSubstitutionList},
	
	{"should be able to remove subs to default block substitution list", "asciidoctor/blocks_test_should_be_able_to_remove_subs_to_default_block_substitution_list.adoc", shouldBeAbleToRemoveSubsToDefaultBlockSubstitutionList},
	
	{"should be able to prepend, append and remove subs from default block substitution list", "asciidoctor/blocks_test_should_be_able_to_prepend,_append_and_remove_subs_from_default_block_substitution_list.adoc", shouldBeAbleToPrependAppendAndRemoveSubsFromDefaultBlockSubstitutionList},
	
	{"should be able to set subs then modify them", "asciidoctor/blocks_test_should_be_able_to_set_subs_then_modify_them.adoc", shouldBeAbleToSetSubsThenModifyThem},
	
	{"should not recognize block anchor with illegal id characters", "asciidoctor/blocks_test_should_not_recognize_block_anchor_with_illegal_id_characters.adoc", shouldNotRecognizeBlockAnchorWithIllegalIdCharacters},
	
	{"should not recognize block anchor that starts with digit", "asciidoctor/blocks_test_should_not_recognize_block_anchor_that_starts_with_digit.adoc", shouldNotRecognizeBlockAnchorThatStartsWithDigit},
	
	{"should recognize block anchor that starts with colon", "asciidoctor/blocks_test_should_recognize_block_anchor_that_starts_with_colon.adoc", shouldRecognizeBlockAnchorThatStartsWithColon},
	
	{"should use specified id and reftext when registering block reference", "asciidoctor/blocks_test_should_use_specified_id_and_reftext_when_registering_block_reference.adoc", shouldUseSpecifiedIdAndReftextWhenRegisteringBlockReference},
	
	{"should allow square brackets in block reference text", "asciidoctor/blocks_test_should_allow_square_brackets_in_block_reference_text.adoc", shouldAllowSquareBracketsInBlockReferenceText},
	
	{"should allow comma in block reference text", "asciidoctor/blocks_test_should_allow_comma_in_block_reference_text.adoc", shouldAllowCommaInBlockReferenceText},
	
	{"should resolve attribute reference in title using attribute defined at location of block", "asciidoctor/blocks_test_should_resolve_attribute_reference_in_title_using_attribute_defined_at_location_of_block.adoc", shouldResolveAttributeReferenceInTitleUsingAttributeDefinedAtLocationOfBlock},
	
	{"should substitute attribute references in reftext when registering block reference", "asciidoctor/blocks_test_should_substitute_attribute_references_in_reftext_when_registering_block_reference.adoc", shouldSubstituteAttributeReferencesInReftextWhenRegisteringBlockReference},
	
	{"should use specified reftext when registering block reference", "asciidoctor/blocks_test_should_use_specified_reftext_when_registering_block_reference.adoc", shouldUseSpecifiedReftextWhenRegisteringBlockReference},
	
}


var horizontalRuleBetweenBlocks = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.ThematicBreak{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: ".fancy",
            },
          },
        },
      },
    },
  },
}

var lineCommentBetweenParagraphsOffsetByBlankLines = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "first paragraph",
    },
    &asciidoc.NewLine{},
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.SingleLineComment{
      Value: " line comment",
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "second paragraph",
    },
    &asciidoc.NewLine{},
  },
}

var adjacentLineCommentBetweenParagraphs = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
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

var commentBlockBetweenParagraphsOffsetByBlankLines = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "first paragraph",
    },
    &asciidoc.NewLine{},
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.MultiLineComment{
      Delimiter: asciidoc.Delimiter{
        Type: 2,
        Length: 4,
      },
      LineList: asciidoc.LineList{
        "block comment",
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "second paragraph",
    },
    &asciidoc.NewLine{},
  },
}

var commentBlockBetweenParagraphsOffsetByBlankLinesInsideDelimitedBlock = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.ExampleBlock{
      Delimiter: asciidoc.Delimiter{
        Type: 3,
        Length: 4,
      },
      AttributeList: nil,
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "first paragraph",
        },
        &asciidoc.NewLine{},
        asciidoc.EmptyLine{
          Text: "",
        },
        &asciidoc.MultiLineComment{
          Delimiter: asciidoc.Delimiter{
            Type: 2,
            Length: 4,
          },
          LineList: asciidoc.LineList{
            "block comment",
          },
        },
        asciidoc.EmptyLine{
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

var adjacentCommentBlockBetweenParagraphs = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "first paragraph",
    },
    &asciidoc.NewLine{},
    &asciidoc.MultiLineComment{
      Delimiter: asciidoc.Delimiter{
        Type: 2,
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

var canConvertWithBlockCommentAtEndOfDocumentWithTrailingNewlines = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "paragraph",
    },
    &asciidoc.NewLine{},
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.MultiLineComment{
      Delimiter: asciidoc.Delimiter{
        Type: 2,
        Length: 4,
      },
      LineList: asciidoc.LineList{
        "block comment",
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    asciidoc.EmptyLine{
      Text: "",
    },
  },
}

var trailingNewlinesAfterBlockCommentAtEndOfDocumentDoesNotCreateParagraph = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "paragraph",
    },
    &asciidoc.NewLine{},
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.MultiLineComment{
      Delimiter: asciidoc.Delimiter{
        Type: 2,
        Length: 4,
      },
      LineList: asciidoc.LineList{
        "block comment",
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    asciidoc.EmptyLine{
      Text: "",
    },
  },
}

var lineStartingWithThreeSlashesShouldNotBeLineComment = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Paragraph{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TitleAttribute{
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "sample title",
            },
          },
        },
      },
      Set: asciidoc.Set{},
      Admonition: 0,
    },
    &asciidoc.MultiLineComment{
      Delimiter: asciidoc.Delimiter{
        Type: 2,
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

var preprocessorDirectivesShouldNotBeProcessedWithinCommentBlock = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "dummy line",
    },
    &asciidoc.NewLine{},
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.MultiLineComment{
      Delimiter: asciidoc.Delimiter{
        Type: 2,
        Length: 4,
      },
      LineList: asciidoc.LineList{
        "ifdef::asciidoctor[////]",
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "line should be shown",
    },
    &asciidoc.NewLine{},
  },
}

var shouldWarnIfUnterminatedCommentBlockIsDetectedInBody = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "before comment block",
    },
    &asciidoc.NewLine{},
    asciidoc.EmptyLine{
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
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "supposed to be after comment block, except it got swallowed by block comment",
    },
    &asciidoc.NewLine{},
  },
}

var shouldWarnIfUnterminatedCommentBlockIsDetectedInsideAnotherBlock = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "before sidebar block",
    },
    &asciidoc.NewLine{},
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.SidebarBlock{
      Delimiter: asciidoc.Delimiter{
        Type: 8,
        Length: 4,
      },
      AttributeList: nil,
      Set: asciidoc.Set{
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
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "supposed to be after sidebar block, except it got swallowed by block comment",
    },
    &asciidoc.NewLine{},
  },
}

var preprocessorDirectivesShouldNotBeProcessedWithinCommentOpenBlock = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.OpenBlock{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "comment",
            },
          },
        },
      },
      Delimiter: asciidoc.Delimiter{
        Type: 7,
        Length: 2,
      },
      Set: asciidoc.Set{
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
    asciidoc.EmptyLine{
      Text: "",
    },
  },
}

var preprocessorDirectivesShouldNotBeProcessedOnSubsequentLinesOfACommentParagraph = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Paragraph{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "comment",
            },
          },
        },
      },
      Set: asciidoc.Set{
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
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "this line should be shown",
    },
    &asciidoc.NewLine{},
  },
}

var commentStyleOnOpenBlockShouldOnlySkipBlock = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.OpenBlock{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "comment",
            },
          },
        },
      },
      Delimiter: asciidoc.Delimiter{
        Type: 7,
        Length: 2,
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "skip",
        },
        &asciidoc.NewLine{},
        asciidoc.EmptyLine{
          Text: "",
        },
        &asciidoc.String{
          Value: "this block",
        },
        &asciidoc.NewLine{},
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "not this text",
    },
    &asciidoc.NewLine{},
  },
}

var commentStyleOnParagraphShouldOnlySkipParagraph = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Paragraph{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "comment",
            },
          },
        },
      },
      Set: asciidoc.Set{
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
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "not this text",
    },
    &asciidoc.NewLine{},
  },
}

var commentStyleOnParagraphShouldNotCauseAdjacentBlockToBeSkipped = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Paragraph{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "comment",
            },
          },
        },
      },
      Set: asciidoc.Set{
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
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "example",
            },
          },
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "not this text",
        },
        &asciidoc.NewLine{},
      },
      Admonition: 0,
    },
  },
}

var shouldNotDropContentThatFollowsSkippedContentInsideADelimitedBlock = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.ExampleBlock{
      Delimiter: asciidoc.Delimiter{
        Type: 3,
        Length: 4,
      },
      AttributeList: nil,
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "paragraph",
        },
        &asciidoc.NewLine{},
        asciidoc.EmptyLine{
          Text: "",
        },
        &asciidoc.Paragraph{
          AttributeList: asciidoc.AttributeList{
            &asciidoc.PositionalAttribute{
              Offset: 0,
              ImpliedName: "",
              Val: asciidoc.Set{
                &asciidoc.String{
                  Value: "comment#idname",
                },
              },
            },
          },
          Set: asciidoc.Set{
            &asciidoc.String{
              Value: "skip",
            },
            &asciidoc.NewLine{},
          },
          Admonition: 0,
        },
        asciidoc.EmptyLine{
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

var shouldParseSidebarBlock = &asciidoc.Document{
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
        &asciidoc.SidebarBlock{
          Delimiter: asciidoc.Delimiter{
            Type: 8,
            Length: 4,
          },
          AttributeList: asciidoc.AttributeList{
            &asciidoc.TitleAttribute{
              Val: asciidoc.Set{
                &asciidoc.String{
                  Value: "Sidebar",
                },
              },
            },
          },
          Set: asciidoc.Set{
            &asciidoc.String{
              Value: "Content goes here",
            },
            &asciidoc.NewLine{},
          },
        },
      },
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "Section",
        },
      },
      Level: 1,
    },
  },
}

var quoteBlockWithNoAttribution = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.QuoteBlock{
      Delimiter: asciidoc.Delimiter{
        Type: 11,
        Length: 4,
      },
      AttributeList: nil,
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "A famous quote.",
        },
        &asciidoc.NewLine{},
      },
    },
  },
}

var quoteBlockWithAttribution = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.QuoteBlock{
      Delimiter: asciidoc.Delimiter{
        Type: 11,
        Length: 4,
      },
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "quote",
            },
          },
        },
        &asciidoc.PositionalAttribute{
          Offset: 1,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Famous Person",
            },
          },
        },
        &asciidoc.PositionalAttribute{
          Offset: 2,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Famous Book (1999)",
            },
          },
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "A famous quote.",
        },
        &asciidoc.NewLine{},
      },
    },
  },
}

var quoteBlockWithAttributeAndIdAndRoleShorthand = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.QuoteBlock{
      Delimiter: asciidoc.Delimiter{
        Type: 11,
        Length: 4,
      },
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "quote#justice-to-all.solidarity",
            },
          },
        },
        &asciidoc.PositionalAttribute{
          Offset: 1,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Martin Luther King",
            },
          },
        },
        &asciidoc.PositionalAttribute{
          Offset: 2,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Jr.",
            },
          },
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "Injustice anywhere is a threat to justice everywhere.",
        },
        &asciidoc.NewLine{},
      },
    },
  },
}

var settingIdUsingStyleShorthandShouldNotResetBlockStyle = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.QuoteBlock{
      Delimiter: asciidoc.Delimiter{
        Type: 11,
        Length: 4,
      },
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "quote",
            },
          },
        },
        &asciidoc.PositionalAttribute{
          Offset: 1,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "#justice-to-all.solidarity",
            },
          },
        },
        &asciidoc.PositionalAttribute{
          Offset: 2,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Martin Luther King",
            },
          },
        },
        &asciidoc.PositionalAttribute{
          Offset: 3,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Jr.",
            },
          },
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "Injustice anywhere is a threat to justice everywhere.",
        },
        &asciidoc.NewLine{},
      },
    },
  },
}

var quoteBlockWithComplexContent = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.QuoteBlock{
      Delimiter: asciidoc.Delimiter{
        Type: 11,
        Length: 4,
      },
      AttributeList: nil,
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "A famous quote.",
        },
        &asciidoc.NewLine{},
        asciidoc.EmptyLine{
          Text: "",
        },
        &asciidoc.Paragraph{
          AttributeList: nil,
          Set: asciidoc.Set{
            &asciidoc.Italic{
              AttributeList: nil,
              Set: asciidoc.Set{
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

var quoteBlockWithAttributionConvertedToDocBook = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.QuoteBlock{
      Delimiter: asciidoc.Delimiter{
        Type: 11,
        Length: 4,
      },
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "quote",
            },
          },
        },
        &asciidoc.PositionalAttribute{
          Offset: 1,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Famous Person",
            },
          },
        },
        &asciidoc.PositionalAttribute{
          Offset: 2,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Famous Book (1999)",
            },
          },
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "A famous quote.",
        },
        &asciidoc.NewLine{},
      },
    },
  },
}

var epigraphQuoteBlockWithAttributionConvertedToDocBook = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.QuoteBlock{
      Delimiter: asciidoc.Delimiter{
        Type: 11,
        Length: 4,
      },
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: ".epigraph",
            },
          },
        },
        &asciidoc.PositionalAttribute{
          Offset: 1,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Famous Person",
            },
          },
        },
        &asciidoc.PositionalAttribute{
          Offset: 2,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Famous Book (1999)",
            },
          },
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "A famous quote.",
        },
        &asciidoc.NewLine{},
      },
    },
  },
}

var markdownStyleQuoteBlockWithSingleParagraphAndNoAttribution = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
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

var lazyMarkdownStyleQuoteBlockWithSingleParagraphAndNoAttribution = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
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

var markdownStyleQuoteBlockWithMultipleParagraphsAndNoAttribution = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
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

var markdownStyleQuoteBlockWithMultipleBlocksAndNoAttribution = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
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

var markdownStyleQuoteBlockWithSingleParagraphAndAttribution = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
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

var markdownStyleQuoteBlockWithOnlyAttribution = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
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
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "Thomas Jefferson, ",
        },
        &asciidoc.Link{
          AttributeList: asciidoc.AttributeList{
            &asciidoc.PositionalAttribute{
              Offset: 0,
              ImpliedName: "alt",
              Val: asciidoc.Set{
                &asciidoc.String{
                  Value: "The Papers of Thomas Jefferson",
                },
              },
            },
            &asciidoc.PositionalAttribute{
              Offset: 1,
              ImpliedName: "",
              Val: asciidoc.Set{
                &asciidoc.String{
                  Value: "Volume 11",
                },
              },
            },
          },
          URL: asciidoc.URL{
            Scheme: "https://",
            Path: asciidoc.Set{
              &asciidoc.String{
                Value: "jeffersonpapers.princeton.edu/selected-documents/james-madison-1",
              },
            },
          },
        },
      },
      AttributeList: nil,
      Indent: "",
      Marker: "--",
      Checklist: 0,
    },
  },
}

var quotedParagraphStyleQuoteBlockWithAttribution = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
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
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "Famous Person, Famous Source, Volume 1 (1999)",
        },
      },
      AttributeList: nil,
      Indent: "",
      Marker: "--",
      Checklist: 0,
    },
  },
}

var shouldParseCreditLineInQuotedParagraphStyleQuoteBlockLikePositionalBlockAttributes = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
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
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "Thomas Jefferson, ",
        },
        &asciidoc.Link{
          AttributeList: asciidoc.AttributeList{
            &asciidoc.PositionalAttribute{
              Offset: 0,
              ImpliedName: "alt",
              Val: asciidoc.Set{
                &asciidoc.String{
                  Value: "The Papers of Thomas Jefferson",
                },
              },
            },
            &asciidoc.PositionalAttribute{
              Offset: 1,
              ImpliedName: "",
              Val: asciidoc.Set{
                &asciidoc.String{
                  Value: "Volume 11",
                },
              },
            },
          },
          URL: asciidoc.URL{
            Scheme: "https://",
            Path: asciidoc.Set{
              &asciidoc.String{
                Value: "jeffersonpapers.princeton.edu/selected-documents/james-madison-1",
              },
            },
          },
        },
      },
      AttributeList: nil,
      Indent: "",
      Marker: "--",
      Checklist: 0,
    },
  },
}

var singleLineVerseBlockWithoutAttribution = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.QuoteBlock{
      Delimiter: asciidoc.Delimiter{
        Type: 11,
        Length: 4,
      },
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "verse",
            },
          },
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "A famous verse.",
        },
        &asciidoc.NewLine{},
      },
    },
  },
}

var singleLineVerseBlockWithAttribution = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.QuoteBlock{
      Delimiter: asciidoc.Delimiter{
        Type: 11,
        Length: 4,
      },
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "verse",
            },
          },
        },
        &asciidoc.PositionalAttribute{
          Offset: 1,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Famous Poet",
            },
          },
        },
        &asciidoc.PositionalAttribute{
          Offset: 2,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Famous Poem",
            },
          },
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "A famous verse.",
        },
        &asciidoc.NewLine{},
      },
    },
  },
}

var singleLineVerseBlockWithAttributionConvertedToDocBook = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.QuoteBlock{
      Delimiter: asciidoc.Delimiter{
        Type: 11,
        Length: 4,
      },
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "verse",
            },
          },
        },
        &asciidoc.PositionalAttribute{
          Offset: 1,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Famous Poet",
            },
          },
        },
        &asciidoc.PositionalAttribute{
          Offset: 2,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Famous Poem",
            },
          },
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "A famous verse.",
        },
        &asciidoc.NewLine{},
      },
    },
  },
}

var singleLineEpigraphVerseBlockWithAttributionConvertedToDocBook = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.QuoteBlock{
      Delimiter: asciidoc.Delimiter{
        Type: 11,
        Length: 4,
      },
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "verse.epigraph",
            },
          },
        },
        &asciidoc.PositionalAttribute{
          Offset: 1,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Famous Poet",
            },
          },
        },
        &asciidoc.PositionalAttribute{
          Offset: 2,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Famous Poem",
            },
          },
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "A famous verse.",
        },
        &asciidoc.NewLine{},
      },
    },
  },
}

var multiStanzaVerseBlock = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.QuoteBlock{
      Delimiter: asciidoc.Delimiter{
        Type: 11,
        Length: 4,
      },
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "verse",
            },
          },
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "A famous verse.",
        },
        &asciidoc.NewLine{},
        asciidoc.EmptyLine{
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

var verseBlockDoesNotContainBlockElements = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.QuoteBlock{
      Delimiter: asciidoc.Delimiter{
        Type: 11,
        Length: 4,
      },
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "verse",
            },
          },
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "A famous verse.",
        },
        &asciidoc.NewLine{},
        asciidoc.EmptyLine{
          Text: "",
        },
        &asciidoc.LiteralBlock{
          AttributeList: nil,
          Delimiter: asciidoc.Delimiter{
            Type: 6,
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

var verseShouldHaveNormalSubs = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.QuoteBlock{
      Delimiter: asciidoc.Delimiter{
        Type: 11,
        Length: 4,
      },
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "verse",
            },
          },
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "A famous verse",
        },
        &asciidoc.NewLine{},
      },
    },
  },
}

var shouldNotRecognizeCalloutsInAVerse = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.QuoteBlock{
      Delimiter: asciidoc.Delimiter{
        Type: 11,
        Length: 4,
      },
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "verse",
            },
          },
        },
      },
      Set: asciidoc.Set{
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

var shouldPerformNormalSubsOnAVerseBlock = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.QuoteBlock{
      Delimiter: asciidoc.Delimiter{
        Type: 11,
        Length: 4,
      },
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "verse",
            },
          },
        },
      },
      Set: asciidoc.Set{
        &asciidoc.Italic{
          AttributeList: nil,
          Set: asciidoc.Set{
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

var canConvertExampleBlock = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.ExampleBlock{
      Delimiter: asciidoc.Delimiter{
        Type: 3,
        Length: 4,
      },
      AttributeList: nil,
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "This is an example of an example block.",
        },
        &asciidoc.NewLine{},
        asciidoc.EmptyLine{
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

var assignsSequentialNumberedCaptionToExampleBlockWithTitle = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.ExampleBlock{
      Delimiter: asciidoc.Delimiter{
        Type: 3,
        Length: 4,
      },
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TitleAttribute{
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Writing Docs with AsciiDoc",
            },
          },
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "Here's how you write AsciiDoc.",
        },
        &asciidoc.NewLine{},
        asciidoc.EmptyLine{
          Text: "",
        },
        &asciidoc.String{
          Value: "You just write.",
        },
        &asciidoc.NewLine{},
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.ExampleBlock{
      Delimiter: asciidoc.Delimiter{
        Type: 3,
        Length: 4,
      },
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TitleAttribute{
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Writing Docs with DocBook",
            },
          },
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "Here's how you write DocBook.",
        },
        &asciidoc.NewLine{},
        asciidoc.EmptyLine{
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

var assignsSequentialCharacterCaptionToExampleBlockWithTitle = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "example-number",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "@",
        },
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.ExampleBlock{
      Delimiter: asciidoc.Delimiter{
        Type: 3,
        Length: 4,
      },
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TitleAttribute{
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Writing Docs with AsciiDoc",
            },
          },
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "Here's how you write AsciiDoc.",
        },
        &asciidoc.NewLine{},
        asciidoc.EmptyLine{
          Text: "",
        },
        &asciidoc.String{
          Value: "You just write.",
        },
        &asciidoc.NewLine{},
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.ExampleBlock{
      Delimiter: asciidoc.Delimiter{
        Type: 3,
        Length: 4,
      },
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TitleAttribute{
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Writing Docs with DocBook",
            },
          },
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "Here's how you write DocBook.",
        },
        &asciidoc.NewLine{},
        asciidoc.EmptyLine{
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

var shouldIncrementCounterForExampleEvenWhenExampleNumberIsLockedByTheApi = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.ExampleBlock{
      Delimiter: asciidoc.Delimiter{
        Type: 3,
        Length: 4,
      },
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TitleAttribute{
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Writing Docs with AsciiDoc",
            },
          },
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "Here's how you write AsciiDoc.",
        },
        &asciidoc.NewLine{},
        asciidoc.EmptyLine{
          Text: "",
        },
        &asciidoc.String{
          Value: "You just write.",
        },
        &asciidoc.NewLine{},
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.ExampleBlock{
      Delimiter: asciidoc.Delimiter{
        Type: 3,
        Length: 4,
      },
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TitleAttribute{
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Writing Docs with DocBook",
            },
          },
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "Here's how you write DocBook.",
        },
        &asciidoc.NewLine{},
        asciidoc.EmptyLine{
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

var shouldUseExplicitCaptionIfSpecified = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.ExampleBlock{
      Delimiter: asciidoc.Delimiter{
        Type: 3,
        Length: 4,
      },
      AttributeList: asciidoc.AttributeList{
        &asciidoc.NamedAttribute{
          Name: "caption",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Look! ",
            },
          },
          Quote: 2,
        },
        &asciidoc.TitleAttribute{
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Writing Docs with AsciiDoc",
            },
          },
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "Here's how you write AsciiDoc.",
        },
        &asciidoc.NewLine{},
        asciidoc.EmptyLine{
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

var automaticCaptionCanBeTurnedOffAndOnAndModified = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.ExampleBlock{
      Delimiter: asciidoc.Delimiter{
        Type: 3,
        Length: 4,
      },
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TitleAttribute{
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "first example",
            },
          },
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "an example",
        },
        &asciidoc.NewLine{},
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "caption",
      Set: nil,
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.ExampleBlock{
      Delimiter: asciidoc.Delimiter{
        Type: 3,
        Length: 4,
      },
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TitleAttribute{
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "second example",
            },
          },
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "another example",
        },
        &asciidoc.NewLine{},
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeReset{
      Name: "caption",
    },
    &asciidoc.AttributeEntry{
      Name: "example-caption",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "Exhibit",
        },
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.ExampleBlock{
      Delimiter: asciidoc.Delimiter{
        Type: 3,
        Length: 4,
      },
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TitleAttribute{
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "third example",
            },
          },
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "yet another example",
        },
        &asciidoc.NewLine{},
      },
    },
  },
}

var shouldUseExplicitCaptionIfSpecifiedEvenIfBlockSpecificGlobalCaptionIsDisabled = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeReset{
      Name: "example-caption",
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.ExampleBlock{
      Delimiter: asciidoc.Delimiter{
        Type: 3,
        Length: 4,
      },
      AttributeList: asciidoc.AttributeList{
        &asciidoc.NamedAttribute{
          Name: "caption",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Look! ",
            },
          },
          Quote: 2,
        },
        &asciidoc.TitleAttribute{
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Writing Docs with AsciiDoc",
            },
          },
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "Here's how you write AsciiDoc.",
        },
        &asciidoc.NewLine{},
        asciidoc.EmptyLine{
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

var shouldUseGlobalCaptionIfSpecifiedEvenIfBlockSpecificGlobalCaptionIsDisabled = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeReset{
      Name: "example-caption",
    },
    &asciidoc.AttributeEntry{
      Name: "caption",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "Look!{sp}",
        },
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.ExampleBlock{
      Delimiter: asciidoc.Delimiter{
        Type: 3,
        Length: 4,
      },
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TitleAttribute{
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Writing Docs with AsciiDoc",
            },
          },
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "Here's how you write AsciiDoc.",
        },
        &asciidoc.NewLine{},
        asciidoc.EmptyLine{
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

var shouldNotProcessCaptionAttributeOnBlockThatDoesNotSupportACaption = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.OpenBlock{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.NamedAttribute{
          Name: "caption",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Look! ",
            },
          },
          Quote: 2,
        },
        &asciidoc.TitleAttribute{
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "No caption here",
            },
          },
        },
      },
      Delimiter: asciidoc.Delimiter{
        Type: 7,
        Length: 2,
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "content",
        },
        &asciidoc.NewLine{},
      },
    },
  },
}

var shouldCreateDetailsSummarySetIfCollapsibleOptionIsSet = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.ExampleBlock{
      Delimiter: asciidoc.Delimiter{
        Type: 3,
        Length: 4,
      },
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TitleAttribute{
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Toggle Me",
            },
          },
        },
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "%collapsible",
            },
          },
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "This content is revealed when the user clicks the words \"Toggle Me\".",
        },
        &asciidoc.NewLine{},
      },
    },
  },
}

var shouldOpenDetailsSummarySetIfCollapsibleAndOpenOptionsAreSet = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.ExampleBlock{
      Delimiter: asciidoc.Delimiter{
        Type: 3,
        Length: 4,
      },
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TitleAttribute{
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Toggle Me",
            },
          },
        },
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "%collapsible%open",
            },
          },
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "This content is revealed when the user clicks the words \"Toggle Me\".",
        },
        &asciidoc.NewLine{},
      },
    },
  },
}

var shouldAddDefaultSummaryElementIfCollapsibleOptionIsSetAndTitleIsNotSpecifed = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.ExampleBlock{
      Delimiter: asciidoc.Delimiter{
        Type: 3,
        Length: 4,
      },
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "%collapsible",
            },
          },
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "This content is revealed when the user clicks the words \"Details\".",
        },
        &asciidoc.NewLine{},
      },
    },
  },
}

var shouldNotAllowCollapsibleBlockToIncrementExampleNumber = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.ExampleBlock{
      Delimiter: asciidoc.Delimiter{
        Type: 3,
        Length: 4,
      },
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TitleAttribute{
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Before",
            },
          },
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "before",
        },
        &asciidoc.NewLine{},
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.ExampleBlock{
      Delimiter: asciidoc.Delimiter{
        Type: 3,
        Length: 4,
      },
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TitleAttribute{
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Show Me The Goods",
            },
          },
        },
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "%collapsible",
            },
          },
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "This content is revealed when the user clicks the words \"Show Me The Goods\".",
        },
        &asciidoc.NewLine{},
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.ExampleBlock{
      Delimiter: asciidoc.Delimiter{
        Type: 3,
        Length: 4,
      },
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TitleAttribute{
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "After",
            },
          },
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "after",
        },
        &asciidoc.NewLine{},
      },
    },
  },
}

var shouldWarnIfExampleBlockIsNotTerminated = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "outside",
    },
    &asciidoc.NewLine{},
    asciidoc.EmptyLine{
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
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "still inside",
    },
    &asciidoc.NewLine{},
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "eof",
    },
    &asciidoc.NewLine{},
  },
}

var captionBlockLevelAttributeShouldBeUsedAsCaption = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "tip-caption",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "Pro Tip",
        },
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Paragraph{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.NamedAttribute{
          Name: "caption",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Pro Tip",
            },
          },
          Quote: 2,
        },
      },
      Set: asciidoc.Set{},
      Admonition: 0,
    },
    &asciidoc.Paragraph{
      AttributeList: nil,
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "Override the caption of an admonition block using an attribute entry",
        },
        &asciidoc.NewLine{},
      },
      Admonition: 2,
    },
  },
}

var canOverrideCaptionOfAdmonitionBlockUsingDocumentAttribute = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "tip-caption",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "Pro Tip",
        },
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Paragraph{
      AttributeList: nil,
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "Override the caption of an admonition block using an attribute entry",
        },
        &asciidoc.NewLine{},
      },
      Admonition: 2,
    },
  },
}

var blankCaptionDocumentAttributeShouldNotBlankAdmonitionBlockCaption = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "caption",
      Set: nil,
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Paragraph{
      AttributeList: nil,
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "Override the caption of an admonition block using an attribute entry",
        },
        &asciidoc.NewLine{},
      },
      Admonition: 2,
    },
  },
}

var shouldSeparateAdjacentParagraphsAndListingIntoBlocks = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "paragraph 1",
    },
    &asciidoc.NewLine{},
    &asciidoc.Listing{
      AttributeList: nil,
      Delimiter: asciidoc.Delimiter{
        Type: 5,
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

var shouldWarnIfListingBlockIsNotTerminated = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "outside",
    },
    &asciidoc.NewLine{},
    asciidoc.EmptyLine{
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
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "still inside",
    },
    &asciidoc.NewLine{},
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "eof",
    },
    &asciidoc.NewLine{},
  },
}

var shouldNotCrashWhenConvertingVerbatimBlockThatHasNoLines = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.LiteralBlock{
      AttributeList: nil,
      Delimiter: asciidoc.Delimiter{
        Type: 6,
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

var shouldPreserveNewlinesInListingBlock = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Listing{
      AttributeList: nil,
      Delimiter: asciidoc.Delimiter{
        Type: 5,
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

var shouldPreserveNewlinesInVerseBlock = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.OpenBlock{
      AttributeList: nil,
      Delimiter: asciidoc.Delimiter{
        Type: 7,
        Length: 2,
      },
      Set: asciidoc.Set{
        &asciidoc.QuoteBlock{
          Delimiter: asciidoc.Delimiter{
            Type: 11,
            Length: 4,
          },
          AttributeList: asciidoc.AttributeList{
            &asciidoc.PositionalAttribute{
              Offset: 0,
              ImpliedName: "",
              Val: asciidoc.Set{
                &asciidoc.String{
                  Value: "verse",
                },
              },
            },
          },
          Set: asciidoc.Set{
            &asciidoc.String{
              Value: "line one",
            },
            &asciidoc.NewLine{},
            asciidoc.EmptyLine{
              Text: "",
            },
            &asciidoc.String{
              Value: "line two",
            },
            &asciidoc.NewLine{},
            asciidoc.EmptyLine{
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

var shouldStripLeadingAndTrailingBlankLinesWhenConvertingVerbatimBlock = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.LiteralBlock{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.NamedAttribute{
          Name: "subs",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "attributes+",
            },
          },
          Quote: 0,
        },
      },
      Delimiter: asciidoc.Delimiter{
        Type: 6,
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

var shouldRemoveBlockIndentIfIndentAttributeIs0 = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Listing{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.NamedAttribute{
          Name: "indent",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "0",
            },
          },
          Quote: 2,
        },
      },
      Delimiter: asciidoc.Delimiter{
        Type: 5,
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

var shouldNotRemoveBlockIndentIfIndentAttributeIs1 = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Listing{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.NamedAttribute{
          Name: "indent",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "-1",
            },
          },
          Quote: 2,
        },
      },
      Delimiter: asciidoc.Delimiter{
        Type: 5,
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

var shouldSetBlockIndentToValueSpecifiedByIndentAttribute = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Listing{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.NamedAttribute{
          Name: "indent",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "1",
            },
          },
          Quote: 2,
        },
      },
      Delimiter: asciidoc.Delimiter{
        Type: 5,
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

var shouldSetBlockIndentToValueSpecifiedByIndentDocumentAttribute = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "source-indent",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "1",
        },
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Listing{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "source",
            },
          },
        },
        &asciidoc.PositionalAttribute{
          Offset: 1,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "ruby",
            },
          },
        },
      },
      Delimiter: asciidoc.Delimiter{
        Type: 5,
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

var literalBlockShouldHonorNowrapOption = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Listing{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.NamedAttribute{
          Name: "options",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "nowrap",
            },
          },
          Quote: 2,
        },
      },
      Delimiter: asciidoc.Delimiter{
        Type: 5,
        Length: 4,
      },
      LineList: asciidoc.LineList{
        "Do not wrap me if I get too long.",
      },
    },
  },
}

var literalBlockShouldSetNowrapClassIfPrewrapDocumentAttributeIsDisabled = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeReset{
      Name: "prewrap",
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Listing{
      AttributeList: nil,
      Delimiter: asciidoc.Delimiter{
        Type: 5,
        Length: 4,
      },
      LineList: asciidoc.LineList{
        "Do not wrap me if I get too long.",
      },
    },
  },
}

var shouldPreserveGuardInFrontOfCalloutIfIconsAreNotEnabled = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Listing{
      AttributeList: nil,
      Delimiter: asciidoc.Delimiter{
        Type: 5,
        Length: 4,
      },
      LineList: asciidoc.LineList{
        "puts 'Hello, World!' # <1>",
        "puts 'Goodbye, World ;(' # <2>",
      },
    },
  },
}

var shouldPreserveGuardAroundCalloutIfIconsAreNotEnabled = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Listing{
      AttributeList: nil,
      Delimiter: asciidoc.Delimiter{
        Type: 5,
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

var literalBlockShouldHonorExplicitSubsList = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Listing{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.NamedAttribute{
          Name: "subs",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "verbatim,quotes",
            },
          },
          Quote: 2,
        },
      },
      Delimiter: asciidoc.Delimiter{
        Type: 5,
        Length: 4,
      },
      LineList: asciidoc.LineList{
        "Map<String, String> *attributes*; //<1>",
      },
    },
  },
}

var shouldBeAbleToDisableCalloutsForLiteralBlock = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Listing{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.NamedAttribute{
          Name: "subs",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "specialcharacters",
            },
          },
          Quote: 2,
        },
      },
      Delimiter: asciidoc.Delimiter{
        Type: 5,
        Length: 4,
      },
      LineList: asciidoc.LineList{
        "No callout here <1>",
      },
    },
  },
}

var listingBlockShouldHonorExplicitSubsList = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Listing{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.NamedAttribute{
          Name: "subs",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "specialcharacters,quotes",
            },
          },
          Quote: 2,
        },
      },
      Delimiter: asciidoc.Delimiter{
        Type: 5,
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

var shouldNotMangleArrayThatContainsFormattedTextWithRoleInListingBlockWithQuotesSubEnabled = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "[,ruby,subs=+quotes]",
    },
    &asciidoc.NewLine{},
    &asciidoc.Listing{
      AttributeList: nil,
      Delimiter: asciidoc.Delimiter{
        Type: 5,
        Length: 4,
      },
      LineList: asciidoc.LineList{
        "nums = [1, 2, 3, [.added]#4#]",
      },
    },
  },
}

var firstCharacterOfBlockTitleMayBeAPeriodIfNotFollowedBySpace = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "..gitignore",
    },
    &asciidoc.NewLine{},
    &asciidoc.Listing{
      AttributeList: nil,
      Delimiter: asciidoc.Delimiter{
        Type: 5,
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

var listingBlockWithoutTitleShouldGenerateScreenElementInDocbook = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Listing{
      AttributeList: nil,
      Delimiter: asciidoc.Delimiter{
        Type: 5,
        Length: 4,
      },
      LineList: asciidoc.LineList{
        "listing block",
      },
    },
  },
}

var listingBlockWithTitleShouldGenerateScreenElementInsideFormalparaElementInDocbook = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Listing{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TitleAttribute{
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "title",
            },
          },
        },
      },
      Delimiter: asciidoc.Delimiter{
        Type: 5,
        Length: 4,
      },
      LineList: asciidoc.LineList{
        "listing block",
      },
    },
  },
}

var shouldNotPrependCaptionToTitleOfListingBlockWithTitleIfListingCaptionAttributeIsNotSet = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Listing{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TitleAttribute{
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "title",
            },
          },
        },
      },
      Delimiter: asciidoc.Delimiter{
        Type: 5,
        Length: 4,
      },
      LineList: asciidoc.LineList{
        "listing block content",
      },
    },
  },
}

var shouldPrependCaptionSpecifiedByListingCaptionAttributeAndNumberToTitleOfListingBlockWithTitle = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "listing-caption",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "Listing",
        },
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Listing{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TitleAttribute{
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "title",
            },
          },
        },
      },
      Delimiter: asciidoc.Delimiter{
        Type: 5,
        Length: 4,
      },
      LineList: asciidoc.LineList{
        "listing block content",
      },
    },
  },
}

var shouldPrependCaptionSpecifiedByCaptionAttributeOnListingBlockEvenIfListingCaptionAttributeIsNotSet = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Listing{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.NamedAttribute{
          Name: "caption",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Listing ",
            },
            &asciidoc.Counter{
              Name: "listing-number",
              InitialValue: "",
              Display: true,
            },
            &asciidoc.String{
              Value: ". ",
            },
          },
          Quote: 2,
        },
        &asciidoc.TitleAttribute{
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Behold!",
            },
          },
        },
      },
      Delimiter: asciidoc.Delimiter{
        Type: 5,
        Length: 4,
      },
      LineList: asciidoc.LineList{
        "listing block content",
      },
    },
  },
}

var listingBlockWithoutAnExplicitStyleAndWithASecondPositionalArgumentShouldBePromotedToASourceBlock = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "[,ruby]",
    },
    &asciidoc.NewLine{},
    &asciidoc.Listing{
      AttributeList: nil,
      Delimiter: asciidoc.Delimiter{
        Type: 5,
        Length: 4,
      },
      LineList: asciidoc.LineList{
        "puts 'Hello, Ruby!'",
      },
    },
  },
}

var listingBlockWithoutAnExplicitStyleShouldBePromotedToASourceBlockIfSourceLanguageIsSet = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "source-language",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "ruby",
        },
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Listing{
      AttributeList: nil,
      Delimiter: asciidoc.Delimiter{
        Type: 5,
        Length: 4,
      },
      LineList: asciidoc.LineList{
        "puts 'Hello, Ruby!'",
      },
    },
  },
}

var listingBlockWithAnExplicitStyleAndASecondPositionalArgumentShouldNotBePromotedToASourceBlock = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Listing{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "listing",
            },
          },
        },
        &asciidoc.PositionalAttribute{
          Offset: 1,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "ruby",
            },
          },
        },
      },
      Delimiter: asciidoc.Delimiter{
        Type: 5,
        Length: 4,
      },
      LineList: asciidoc.LineList{
        "puts 'Hello, Ruby!'",
      },
    },
  },
}

var listingBlockWithAnExplicitStyleShouldNotBePromotedToASourceBlockIfSourceLanguageIsSet = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "source-language",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "ruby",
        },
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Listing{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "listing",
            },
          },
        },
      },
      Delimiter: asciidoc.Delimiter{
        Type: 5,
        Length: 4,
      },
      LineList: asciidoc.LineList{
        "puts 'Hello, Ruby!'",
      },
    },
  },
}

var sourceBlockWithNoTitleOrLanguageShouldGenerateScreenElementInDocbook = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Listing{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "source",
            },
          },
        },
      },
      Delimiter: asciidoc.Delimiter{
        Type: 5,
        Length: 4,
      },
      LineList: asciidoc.LineList{
        "source block",
      },
    },
  },
}

var sourceBlockWithTitleAndNoLanguageShouldGenerateScreenElementInsideFormalparaElementForDocbook = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Listing{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "source",
            },
          },
        },
        &asciidoc.TitleAttribute{
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "title",
            },
          },
        },
      },
      Delimiter: asciidoc.Delimiter{
        Type: 5,
        Length: 4,
      },
      LineList: asciidoc.LineList{
        "source block",
      },
    },
  },
}

var canConvertOpenBlock = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.OpenBlock{
      AttributeList: nil,
      Delimiter: asciidoc.Delimiter{
        Type: 7,
        Length: 2,
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "This is an open block.",
        },
        &asciidoc.NewLine{},
        asciidoc.EmptyLine{
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

var openBlockCanContainAnotherBlock = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.OpenBlock{
      AttributeList: nil,
      Delimiter: asciidoc.Delimiter{
        Type: 7,
        Length: 2,
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "This is an open block.",
        },
        &asciidoc.NewLine{},
        asciidoc.EmptyLine{
          Text: "",
        },
        &asciidoc.String{
          Value: "It can span multiple lines.",
        },
        &asciidoc.NewLine{},
        asciidoc.EmptyLine{
          Text: "",
        },
        &asciidoc.QuoteBlock{
          Delimiter: asciidoc.Delimiter{
            Type: 11,
            Length: 4,
          },
          AttributeList: nil,
          Set: asciidoc.Set{
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

var shouldTransferIdAndReftextOnOpenBlockToDocBookOutput = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "Check out that ",
    },
    &asciidoc.CrossReference{
      Set: nil,
      ID: "open",
    },
    &asciidoc.String{
      Value: "!",
    },
    &asciidoc.NewLine{},
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.OpenBlock{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.AnchorAttribute{
          ID: &asciidoc.String{
            Value: "open",
          },
          Label: asciidoc.Set{
            &asciidoc.String{
              Value: "Open Block",
            },
          },
        },
      },
      Delimiter: asciidoc.Delimiter{
        Type: 7,
        Length: 2,
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "This is an open block.",
        },
        &asciidoc.NewLine{},
        asciidoc.EmptyLine{
          Text: "",
        },
        &asciidoc.Paragraph{
          AttributeList: nil,
          Set: asciidoc.Set{
            &asciidoc.String{
              Value: "An open block can have other blocks inside of it.",
            },
            &asciidoc.NewLine{},
          },
          Admonition: 2,
        },
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "Back to our regularly scheduled programming.",
    },
    &asciidoc.NewLine{},
  },
}

var shouldTransferIdAndReftextOnOpenParagraphToDocBookOutput = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Paragraph{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "open#openpara",
            },
          },
        },
        &asciidoc.NamedAttribute{
          Name: "reftext",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Open Paragraph",
            },
          },
          Quote: 2,
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "This is an open paragraph.",
        },
        &asciidoc.NewLine{},
      },
      Admonition: 0,
    },
  },
}

var shouldTransferTitleOnOpenBlockToDocBookOutput = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.OpenBlock{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TitleAttribute{
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Behold the open",
            },
          },
        },
      },
      Delimiter: asciidoc.Delimiter{
        Type: 7,
        Length: 2,
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "This is an open block with a title.",
        },
        &asciidoc.NewLine{},
      },
    },
  },
}

var shouldTransferTitleOnOpenParagraphToDocBookOutput = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Paragraph{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TitleAttribute{
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Behold the open",
            },
          },
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "This is an open paragraph with a title.",
        },
        &asciidoc.NewLine{},
      },
      Admonition: 0,
    },
  },
}

var shouldTransferRoleOnOpenBlockToDocBookOutput = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.OpenBlock{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: ".container",
            },
          },
        },
      },
      Delimiter: asciidoc.Delimiter{
        Type: 7,
        Length: 2,
      },
      Set: asciidoc.Set{
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

var shouldTransferRoleOnOpenParagraphToDocBookOutput = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Paragraph{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: ".container",
            },
          },
        },
      },
      Set: asciidoc.Set{
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

var canParseAPassthroughBlock = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.StemBlock{
      AttributeList: nil,
      Delimiter: asciidoc.Delimiter{
        Type: 9,
        Length: 4,
      },
      LineList: asciidoc.LineList{
        "This is a passthrough block.",
      },
    },
  },
}

var doesNotPerformSubsOnAPassthroughBlockByDefault = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "type",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "passthrough",
        },
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.StemBlock{
      AttributeList: nil,
      Delimiter: asciidoc.Delimiter{
        Type: 9,
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

var doesNotPerformSubsOnAPassthroughBlockWithPassStyleByDefault = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "type",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "passthrough",
        },
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.StemBlock{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "pass",
            },
          },
        },
      },
      Delimiter: asciidoc.Delimiter{
        Type: 9,
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

var passthroughBlockHonorsExplicitSubsList = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "type",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "passthrough",
        },
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.StemBlock{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.NamedAttribute{
          Name: "subs",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "attributes,quotes,macros",
            },
          },
          Quote: 2,
        },
      },
      Delimiter: asciidoc.Delimiter{
        Type: 9,
        Length: 4,
      },
      LineList: asciidoc.LineList{
        "This is a _{type}_ block.",
        "http://asciidoc.org",
      },
    },
  },
}

var shouldStripLeadingAndTrailingBlankLinesWhenConvertingRawBlock = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.StemBlock{
      AttributeList: nil,
      Delimiter: asciidoc.Delimiter{
        Type: 9,
        Length: 4,
      },
      LineList: asciidoc.LineList{
        "line above",
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.StemBlock{
      AttributeList: nil,
      Delimiter: asciidoc.Delimiter{
        Type: 9,
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
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.StemBlock{
      AttributeList: nil,
      Delimiter: asciidoc.Delimiter{
        Type: 9,
        Length: 4,
      },
      LineList: asciidoc.LineList{
        "line below",
      },
    },
  },
}

var shouldNotCrashWhenConvertingStemBlockThatHasNoLines = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.StemBlock{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "stem",
            },
          },
        },
      },
      Delimiter: asciidoc.Delimiter{
        Type: 9,
        Length: 4,
      },
      LineList: asciidoc.LineList{},
    },
  },
}

var shouldReturnContentAsEmptyStringForStemOrPassBlockThatHasNoLines = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.StemBlock{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "latexmath",
            },
          },
        },
      },
      Delimiter: asciidoc.Delimiter{
        Type: 9,
        Length: 4,
      },
      LineList: asciidoc.LineList{
        "\\sqrt{3x-1}+(1+x)^2 < y",
      },
    },
  },
}

var shouldNotAddLaTeXMathDelimitersAroundLatexmathBlockContentIfAlreadyPresent = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.StemBlock{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "latexmath",
            },
          },
        },
      },
      Delimiter: asciidoc.Delimiter{
        Type: 9,
        Length: 4,
      },
      LineList: asciidoc.LineList{
        "\\[\\sqrt{3x-1}+(1+x)^2 < y\\]",
      },
    },
  },
}

var shouldDisplayLatexmathBlockInAltOfEquationInDocBookBackend = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.StemBlock{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "latexmath",
            },
          },
        },
      },
      Delimiter: asciidoc.Delimiter{
        Type: 9,
        Length: 4,
      },
      LineList: asciidoc.LineList{
        "\\sqrt{3x-1}+(1+x)^2 < y",
      },
    },
  },
}

var shouldSetAutoNumberOptionForLatexmathToNoneByDefault = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "stem",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "latexmath",
        },
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.StemBlock{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "stem",
            },
          },
        },
      },
      Delimiter: asciidoc.Delimiter{
        Type: 9,
        Length: 4,
      },
      LineList: asciidoc.LineList{
        "y = x^2",
      },
    },
  },
}

var shouldSetAutoNumberOptionForLatexmathToNoneIfEqnumsIsSetToNone = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "stem",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "latexmath",
        },
      },
    },
    &asciidoc.AttributeEntry{
      Name: "eqnums",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "none",
        },
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.StemBlock{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "stem",
            },
          },
        },
      },
      Delimiter: asciidoc.Delimiter{
        Type: 9,
        Length: 4,
      },
      LineList: asciidoc.LineList{
        "y = x^2",
      },
    },
  },
}

var shouldSetAutoNumberOptionForLatexmathToAmsIfEqnumsIsSet = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "stem",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "latexmath",
        },
      },
    },
    &asciidoc.AttributeEntry{
      Name: "eqnums",
      Set: nil,
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.StemBlock{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "stem",
            },
          },
        },
      },
      Delimiter: asciidoc.Delimiter{
        Type: 9,
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

var shouldSetAutoNumberOptionForLatexmathToAllIfEqnumsIsSetToAll = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "stem",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "latexmath",
        },
      },
    },
    &asciidoc.AttributeEntry{
      Name: "eqnums",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "all",
        },
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.StemBlock{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "stem",
            },
          },
        },
      },
      Delimiter: asciidoc.Delimiter{
        Type: 9,
        Length: 4,
      },
      LineList: asciidoc.LineList{
        "y = x^2",
      },
    },
  },
}

var shouldNotSplitEquationInAsciiMathBlockAtSingleNewline = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.StemBlock{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "asciimath",
            },
          },
        },
      },
      Delimiter: asciidoc.Delimiter{
        Type: 9,
        Length: 4,
      },
      LineList: asciidoc.LineList{
        "f: bbb\"N\" -> bbb\"N\"",
        "f: x |-> x + 1",
      },
    },
  },
}

var shouldSplitEquationInAsciiMathBlockAtEscapedNewline = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.StemBlock{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "asciimath",
            },
          },
        },
      },
      Delimiter: asciidoc.Delimiter{
        Type: 9,
        Length: 4,
      },
      LineList: asciidoc.LineList{
        "f: bbb\"N\" -> bbb\"N\" \\",
        "f: x |-> x + 1",
      },
    },
  },
}

var shouldSplitEquationInAsciiMathBlockAtSequenceOfEscapedNewlines = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.StemBlock{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "asciimath",
            },
          },
        },
      },
      Delimiter: asciidoc.Delimiter{
        Type: 9,
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

var shouldSplitEquationInAsciiMathBlockAtNewlineSequenceAndPreserveBreaks = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.StemBlock{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "asciimath",
            },
          },
        },
      },
      Delimiter: asciidoc.Delimiter{
        Type: 9,
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

var shouldAddAsciiMathDelimitersAroundAsciimathBlockContent = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.StemBlock{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "asciimath",
            },
          },
        },
      },
      Delimiter: asciidoc.Delimiter{
        Type: 9,
        Length: 4,
      },
      LineList: asciidoc.LineList{
        "sqrt(3x-1)+(1+x)^2 < y",
      },
    },
  },
}

var shouldNotAddAsciiMathDelimitersAroundAsciimathBlockContentIfAlreadyPresent = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.StemBlock{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "asciimath",
            },
          },
        },
      },
      Delimiter: asciidoc.Delimiter{
        Type: 9,
        Length: 4,
      },
      LineList: asciidoc.LineList{
        "\\$sqrt(3x-1)+(1+x)^2 < y\\$",
      },
    },
  },
}

var shouldConvertContentsOfAsciimathBlockToMathMlInDocBookOutputIfAsciimathGemIsAvailable = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.StemBlock{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "asciimath",
            },
          },
        },
      },
      Delimiter: asciidoc.Delimiter{
        Type: 9,
        Length: 4,
      },
      LineList: asciidoc.LineList{
        "x+b/(2a)<+-sqrt((b^2)/(4a^2)-c/a)",
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.StemBlock{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "asciimath",
            },
          },
        },
      },
      Delimiter: asciidoc.Delimiter{
        Type: 9,
        Length: 4,
      },
      LineList: asciidoc.LineList{},
    },
  },
}

var shouldOutputTitleForLatexmathBlockIfDefined = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.StemBlock{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TitleAttribute{
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "The Lorenz Equations",
            },
          },
        },
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "latexmath",
            },
          },
        },
      },
      Delimiter: asciidoc.Delimiter{
        Type: 9,
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

var shouldOutputTitleForAsciimathBlockIfDefined = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.StemBlock{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TitleAttribute{
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Simple fraction",
            },
          },
        },
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "asciimath",
            },
          },
        },
      },
      Delimiter: asciidoc.Delimiter{
        Type: 9,
        Length: 4,
      },
      LineList: asciidoc.LineList{
        "a//b",
      },
    },
  },
}

var shouldAddAsciiMathDelimitersAroundStemBlockContentIfStemAttributeIsAsciimathEmptyOrNotSet = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.StemBlock{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "stem",
            },
          },
        },
      },
      Delimiter: asciidoc.Delimiter{
        Type: 9,
        Length: 4,
      },
      LineList: asciidoc.LineList{
        "sqrt(3x-1)+(1+x)^2 < y",
      },
    },
  },
}

var shouldAddLaTeXMathDelimitersAroundStemBlockContentIfStemAttributeIsLatexmathLatexOrTex = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.StemBlock{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "stem",
            },
          },
        },
      },
      Delimiter: asciidoc.Delimiter{
        Type: 9,
        Length: 4,
      },
      LineList: asciidoc.LineList{
        "\\sqrt{3x-1}+(1+x)^2 < y",
      },
    },
  },
}

var shouldAllowStemStyleToBeSetUsingSecondPositionalArgumentOfBlockAttributes = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "stem",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "latexmath",
        },
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.StemBlock{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "stem",
            },
          },
        },
        &asciidoc.PositionalAttribute{
          Offset: 1,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "asciimath",
            },
          },
        },
      },
      Delimiter: asciidoc.Delimiter{
        Type: 9,
        Length: 4,
      },
      LineList: asciidoc.LineList{
        "sqrt(3x-1)+(1+x)^2 < y",
      },
    },
  },
}

var shouldNotWarnIfBlockStyleIsUnknown = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.OpenBlock{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "foo",
            },
          },
        },
      },
      Delimiter: asciidoc.Delimiter{
        Type: 7,
        Length: 2,
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "bar",
        },
        &asciidoc.NewLine{},
      },
    },
  },
}

var shouldLogDebugMessageIfBlockStyleIsUnknownAndDebugLevelIsEnabled = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.OpenBlock{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "foo",
            },
          },
        },
      },
      Delimiter: asciidoc.Delimiter{
        Type: 7,
        Length: 2,
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "bar",
        },
        &asciidoc.NewLine{},
      },
    },
  },
}

var blockTitleAboveSectionGetsCarriedOverToFirstBlockInSection = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Section{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TitleAttribute{
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Title",
            },
          },
        },
      },
      Set: asciidoc.Set{
        asciidoc.EmptyLine{
          Text: "",
        },
        &asciidoc.String{
          Value: "paragraph",
        },
        &asciidoc.NewLine{},
      },
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "Section",
        },
      },
      Level: 1,
    },
  },
}

var blockTitleAboveDocumentTitleDemotesDocumentTitleToASectionTitle = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Section{
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
        asciidoc.EmptyLine{
          Text: "",
        },
        &asciidoc.String{
          Value: "section paragraph",
        },
        &asciidoc.NewLine{},
      },
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "Section Title",
        },
      },
      Level: 0,
    },
  },
}

var blockTitleAboveDocumentTitleGetsCarriedOverToFirstBlockInFirstSectionIfNoPreamble = &asciidoc.Document{
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
    &asciidoc.Section{ // p0
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
              Value: "paragraph",
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
          Value: "Document Title",
        },
      },
      Level: 0,
    },
  },
}

var shouldApplySubstitutionsToABlockTitleInNormalOrder = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Paragraph{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TitleAttribute{
          Val: asciidoc.Set{
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
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "The one and only!",
        },
        &asciidoc.NewLine{},
      },
      Admonition: 0,
    },
  },
}

var emptyAttributeListShouldNotAppearInOutput = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.OpenBlock{
      AttributeList: nil,
      Delimiter: asciidoc.Delimiter{
        Type: 7,
        Length: 2,
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "Block content",
        },
        &asciidoc.NewLine{},
      },
    },
  },
}

var emptyBlockAnchorShouldNotAppearInOutput = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "[[]]",
    },
    &asciidoc.NewLine{},
    &asciidoc.OpenBlock{
      AttributeList: nil,
      Delimiter: asciidoc.Delimiter{
        Type: 7,
        Length: 2,
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "Block content",
        },
        &asciidoc.NewLine{},
      },
    },
  },
}

var canConvertBlockImageWithAltTextDefinedInMacro = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "imagesdir",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "images",
        },
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.BlockImage{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "alt",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Tiger",
            },
          },
        },
        &asciidoc.PositionalAttribute{
          Offset: 1,
          ImpliedName: "width",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "100",
            },
          },
        },
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "alt",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "%interactive",
            },
          },
        },
      },
      Path: asciidoc.Set{
        &asciidoc.String{
          Value: "tiger.svg",
        },
      },
    },
  },
}

var convertsSvgImageWithAltTextUsingImgElementWhenSafeModeIsSecure = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.BlockImage{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "alt",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Tiger",
            },
          },
        },
        &asciidoc.PositionalAttribute{
          Offset: 1,
          ImpliedName: "width",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "100",
            },
          },
        },
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "alt",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "%interactive",
            },
          },
        },
      },
      Path: asciidoc.Set{
        &asciidoc.String{
          Value: "images/tiger.svg",
        },
      },
    },
  },
}

var insertsFallbackImageForSvgInsideObjectElementUsingSameDimensions = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "imagesdir",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "images",
        },
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.BlockImage{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "alt",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Tiger",
            },
          },
        },
        &asciidoc.PositionalAttribute{
          Offset: 1,
          ImpliedName: "width",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "100",
            },
          },
        },
        &asciidoc.NamedAttribute{
          Name: "fallback",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "tiger.png",
            },
          },
          Quote: 0,
        },
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "alt",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "%interactive",
            },
          },
        },
      },
      Path: asciidoc.Set{
        &asciidoc.String{
          Value: "tiger.svg",
        },
      },
    },
  },
}

var detectsSvgImageUriThatContainsAQueryString = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "imagesdir",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "images",
        },
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.BlockImage{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "alt",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Tiger",
            },
          },
        },
        &asciidoc.PositionalAttribute{
          Offset: 1,
          ImpliedName: "width",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "100",
            },
          },
        },
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "alt",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "%interactive",
            },
          },
        },
      },
      Path: asciidoc.Set{
        &asciidoc.String{
          Value: "http://example.org/tiger.svg?foo=bar",
        },
      },
    },
  },
}

var detectsSvgImageWhenFormatAttributeIsSvg = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "imagesdir",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "images",
        },
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.BlockImage{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "alt",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Tiger",
            },
          },
        },
        &asciidoc.PositionalAttribute{
          Offset: 1,
          ImpliedName: "width",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "100",
            },
          },
        },
        &asciidoc.NamedAttribute{
          Name: "format",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "svg",
            },
          },
          Quote: 0,
        },
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "alt",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "%interactive",
            },
          },
        },
      },
      Path: asciidoc.Set{
        &asciidoc.String{
          Value: "http://example.org/tiger-svg",
        },
      },
    },
  },
}

var convertsToInlineSvgImageWhenInlineOptionIsSetOnBlock = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "imagesdir",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "fixtures",
        },
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.BlockImage{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "alt",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Tiger",
            },
          },
        },
        &asciidoc.PositionalAttribute{
          Offset: 1,
          ImpliedName: "width",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "100",
            },
          },
        },
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "alt",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "%inline",
            },
          },
        },
      },
      Path: asciidoc.Set{
        &asciidoc.String{
          Value: "circle.svg",
        },
      },
    },
  },
}

var shouldIgnoreLinkAttributeIfValueIsSelfAndImageTargetIsInlineSvg = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "imagesdir",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "fixtures",
        },
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.BlockImage{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "alt",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Tiger",
            },
          },
        },
        &asciidoc.PositionalAttribute{
          Offset: 1,
          ImpliedName: "width",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "100",
            },
          },
        },
        &asciidoc.NamedAttribute{
          Name: "link",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "self",
            },
          },
          Quote: 0,
        },
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "alt",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "%inline",
            },
          },
        },
      },
      Path: asciidoc.Set{
        &asciidoc.String{
          Value: "circle.svg",
        },
      },
    },
  },
}

var shouldHonorPercentageWidthForSvgImageWithInlineOption = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "imagesdir",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "fixtures",
        },
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.BlockImage{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "alt",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Circle",
            },
          },
        },
        &asciidoc.PositionalAttribute{
          Offset: 1,
          ImpliedName: "width",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "50%",
            },
          },
        },
        &asciidoc.NamedAttribute{
          Name: "opts",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "inline",
            },
          },
          Quote: 0,
        },
      },
      Path: asciidoc.Set{
        &asciidoc.String{
          Value: "circle.svg",
        },
      },
    },
  },
}

var shouldNotCrashIfExplicitWidthOnSvgImageBlockIsAnInteger = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "imagesdir",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "fixtures",
        },
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.BlockImage{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "alt",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Circle",
            },
          },
        },
        &asciidoc.NamedAttribute{
          Name: "opts",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "inline",
            },
          },
          Quote: 0,
        },
      },
      Path: asciidoc.Set{
        &asciidoc.String{
          Value: "circle.svg",
        },
      },
    },
  },
}

var convertsToInlineSvgImageWhenInlineOptionIsSetOnBlockAndDataUriIsSetOnDocument = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "imagesdir",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "fixtures",
        },
      },
    },
    &asciidoc.AttributeEntry{
      Name: "data-uri",
      Set: nil,
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.BlockImage{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "alt",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Tiger",
            },
          },
        },
        &asciidoc.PositionalAttribute{
          Offset: 1,
          ImpliedName: "width",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "100",
            },
          },
        },
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "alt",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "%inline",
            },
          },
        },
      },
      Path: asciidoc.Set{
        &asciidoc.String{
          Value: "circle.svg",
        },
      },
    },
  },
}

var shouldNotThrowExceptionIfSvgToInlineIsEmpty = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.BlockImage{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "alt",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Alt Text",
            },
          },
        },
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "alt",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "%inline",
            },
          },
        },
      },
      Path: asciidoc.Set{
        &asciidoc.String{
          Value: "no-such-image.svg",
        },
      },
    },
  },
}

var canConvertBlockImageWithAltTextDefinedInMacroContainingSquareBracket = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.BlockImage{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "alt",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Tiger",
            },
          },
        },
      },
      Path: asciidoc.Set{
        &asciidoc.String{
          Value: "images/tiger.png",
        },
      },
    },
  },
}

var altTextInMacroOverridesAltTextAboveMacro = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.BlockImage{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "alt",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Tiger",
            },
          },
        },
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "alt",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Alt Text",
            },
          },
        },
      },
      Path: asciidoc.Set{
        &asciidoc.String{
          Value: "images/tiger.png",
        },
      },
    },
  },
}

var shouldSubstituteAttributeReferencesInAltTextDefinedInImageBlockMacro = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "alt-text",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "Tiger",
        },
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.BlockImage{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "alt",
          Val: asciidoc.Set{
            &asciidoc.UserAttributeReference{
              Value: "alt-text",
            },
          },
        },
      },
      Path: asciidoc.Set{
        &asciidoc.String{
          Value: "images/tiger.png",
        },
      },
    },
  },
}

var shouldSetDirectionCssClassOnImageIfFloatAttributeIsSet = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.BlockImage{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "alt",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Tiger",
            },
          },
        },
        &asciidoc.NamedAttribute{
          Name: "float",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "left",
            },
          },
          Quote: 0,
        },
      },
      Path: asciidoc.Set{
        &asciidoc.String{
          Value: "images/tiger.png",
        },
      },
    },
  },
}

var shouldSetTextAlignmentCssClassOnImageIfAlignAttributeIsSet = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.BlockImage{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "alt",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Tiger",
            },
          },
        },
        &asciidoc.NamedAttribute{
          Name: "align",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "center",
            },
          },
          Quote: 0,
        },
      },
      Path: asciidoc.Set{
        &asciidoc.String{
          Value: "images/tiger.png",
        },
      },
    },
  },
}

var styleAttributeIsDroppedFromImageMacro = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.BlockImage{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "alt",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Tiger",
            },
          },
        },
        &asciidoc.NamedAttribute{
          Name: "style",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "value",
            },
          },
          Quote: 0,
        },
      },
      Path: asciidoc.Set{
        &asciidoc.String{
          Value: "images/tiger.png",
        },
      },
    },
  },
}

var shouldAutoGenerateAltTextForBlockImageIfAltTextIsNotSpecified = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.BlockImage{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "alt",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Tiger",
            },
          },
        },
        &asciidoc.NamedAttribute{
          Name: "link",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "http://en.wikipedia.org/wiki/Tiger",
            },
          },
          Quote: 1,
        },
      },
      Path: asciidoc.Set{
        &asciidoc.String{
          Value: "images/tiger.png",
        },
      },
    },
  },
}

var canConvertBlockImageWithLinkToSelf = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "imagesdir",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "img",
        },
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.BlockImage{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "alt",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Tiger",
            },
          },
        },
        &asciidoc.NamedAttribute{
          Name: "link",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "self",
            },
          },
          Quote: 0,
        },
      },
      Path: asciidoc.Set{
        &asciidoc.String{
          Value: "tiger.png",
        },
      },
    },
  },
}

var addsRelnoopenerAttributeToBlockImageWithLinkThatTargetsBlankWindow = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.BlockImage{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "alt",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Tiger",
            },
          },
        },
        &asciidoc.TitleAttribute{
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "The AsciiDoc Tiger",
            },
          },
        },
      },
      Path: asciidoc.Set{
        &asciidoc.String{
          Value: "images/tiger.png",
        },
      },
    },
  },
}

var canConvertBlockImageWithExplicitCaption = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.BlockImage{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "alt",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Tiger",
            },
          },
        },
        &asciidoc.NamedAttribute{
          Name: "caption",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Voila! ",
            },
          },
          Quote: 2,
        },
        &asciidoc.TitleAttribute{
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "The AsciiDoc Tiger",
            },
          },
        },
      },
      Path: asciidoc.Set{
        &asciidoc.String{
          Value: "images/tiger.png",
        },
      },
    },
  },
}

var canAlignImageInDocBookBackend = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "attribute-missing",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "skip",
        },
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.BlockImage{
      AttributeList: nil,
      Path: asciidoc.Set{
        &asciidoc.UserAttributeReference{
          Value: "bogus",
        },
      },
    },
  },
}

var shouldNotDropLineIfImageTargetIsMissingAttributeReferenceAndAttributeMissingIsDrop = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "attribute-missing",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "drop",
        },
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.BlockImage{
      AttributeList: nil,
      Path: asciidoc.Set{
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

var dropsLineIfImageTargetIsMissingAttributeReferenceAndAttributeMissingIsDropLine = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "attribute-missing",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "drop-line",
        },
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.BlockImage{
      AttributeList: nil,
      Path: asciidoc.Set{
        &asciidoc.UserAttributeReference{
          Value: "bogus",
        },
      },
    },
  },
}

var shouldNotDropLineIfImageTargetResolvesToBlankAndAttributeMissingIsDropLine = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "attribute-missing",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "drop-line",
        },
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.BlockImage{
      AttributeList: nil,
      Path: asciidoc.Set{
        &asciidoc.CharacterReplacementReference{
          Value: "blank",
        },
      },
    },
  },
}

var droppedImageDoesNotBreakProcessingOfFollowingSectionAndAttributeMissingIsDropLine = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "attribute-missing",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "drop-line",
        },
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.BlockImage{
      AttributeList: nil,
      Path: asciidoc.Set{
        &asciidoc.UserAttributeReference{
          Value: "bogus",
        },
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Section{
      AttributeList: nil,
      Set: nil,
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "Section Title",
        },
      },
      Level: 1,
    },
  },
}

var shouldPassThroughImageThatReferencesUri = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "imagesdir",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "images",
        },
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.BlockImage{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "alt",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Tiger",
            },
          },
        },
      },
      Path: asciidoc.Set{
        &asciidoc.String{
          Value: "http://asciidoc.org/images/tiger.png",
        },
      },
    },
  },
}

var shouldEncodeSpacesInImageTargetIfValueIsAUri = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "imagesdir",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "images",
        },
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.BlockImage{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "alt",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Tiger",
            },
          },
        },
      },
      Path: asciidoc.Set{
        &asciidoc.String{
          Value: "tiger.png",
        },
      },
    },
  },
}

var embedsBase64EncodedDataUriForImageWhenDataUriAttributeIsSet = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "data-uri",
      Set: nil,
    },
    &asciidoc.AttributeEntry{
      Name: "imagesdir",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "fixtures",
        },
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.BlockImage{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "alt",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Dot",
            },
          },
        },
      },
      Path: asciidoc.Set{
        &asciidoc.String{
          Value: "dot.gif",
        },
      },
    },
  },
}

var embedsSvgImageWithImageSvgxmlMimetypeWhenFileExtensionIsSvg = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "imagesdir",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "fixtures",
        },
      },
    },
    &asciidoc.AttributeEntry{
      Name: "data-uri",
      Set: nil,
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.BlockImage{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "alt",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Tiger",
            },
          },
        },
        &asciidoc.PositionalAttribute{
          Offset: 1,
          ImpliedName: "width",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "100",
            },
          },
        },
      },
      Path: asciidoc.Set{
        &asciidoc.String{
          Value: "circle.svg",
        },
      },
    },
  },
}

var shouldLinkToDataUriIfValueOfLinkAttributeIsSelfAndImageIsEmbedded = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "imagesdir",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "fixtures",
        },
      },
    },
    &asciidoc.AttributeEntry{
      Name: "data-uri",
      Set: nil,
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.BlockImage{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "alt",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Tiger",
            },
          },
        },
        &asciidoc.PositionalAttribute{
          Offset: 1,
          ImpliedName: "width",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "100",
            },
          },
        },
        &asciidoc.NamedAttribute{
          Name: "link",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "self",
            },
          },
          Quote: 0,
        },
      },
      Path: asciidoc.Set{
        &asciidoc.String{
          Value: "circle.svg",
        },
      },
    },
  },
}

var embedsEmptyBase64EncodedDataUriForUnreadableImageWhenDataUriAttributeIsSet = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "data-uri",
      Set: nil,
    },
    &asciidoc.AttributeEntry{
      Name: "imagesdir",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "fixtures",
        },
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.BlockImage{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "alt",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Dot",
            },
          },
        },
      },
      Path: asciidoc.Set{
        &asciidoc.String{
          Value: "unreadable.gif",
        },
      },
    },
  },
}

var embedsBase64EncodedDataUriWithApplicationOctetStreamMimetypeWhenFileExtensionIsMissing = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "data-uri",
      Set: nil,
    },
    &asciidoc.AttributeEntry{
      Name: "imagesdir",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "fixtures",
        },
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.BlockImage{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "alt",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Dot",
            },
          },
        },
      },
      Path: asciidoc.Set{
        &asciidoc.String{
          Value: "dot",
        },
      },
    },
  },
}

var canHandleEmbeddedDataUriImages = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "data-uri",
      Set: nil,
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.BlockImage{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "alt",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Dot",
            },
          },
        },
      },
      Path: asciidoc.Set{
        &asciidoc.String{
          Value: "data:image/gif;base64,R0lGODlhAQABAIAAAAUEBAAAACwAAAAAAQABAAACAkQBADs=",
        },
      },
    },
  },
}

var cleansReferenceToAncestorDirectoriesInImagesdirBeforeReadingImageIfSafeModeLevelIsAtLeastSafe = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "data-uri",
      Set: nil,
    },
    &asciidoc.AttributeEntry{
      Name: "imagesdir",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "../..//fixtures/./../../fixtures",
        },
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.BlockImage{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "alt",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Dot",
            },
          },
        },
      },
      Path: asciidoc.Set{
        &asciidoc.String{
          Value: "dot.gif",
        },
      },
    },
  },
}

var cleansReferenceToAncestorDirectoriesInTargetBeforeReadingImageIfSafeModeLevelIsAtLeastSafe = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "data-uri",
      Set: nil,
    },
    &asciidoc.AttributeEntry{
      Name: "imagesdir",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "./",
        },
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.BlockImage{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "alt",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Dot",
            },
          },
        },
      },
      Path: asciidoc.Set{
        &asciidoc.String{
          Value: "../..//fixtures/./../../fixtures/dot.gif",
        },
      },
    },
  },
}

var shouldDetectAndConvertVideoMacro = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "imagesdir",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "assets",
        },
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "video::cats-vs-dogs.avi[cats-and-dogs.png, 200, 300]",
    },
    &asciidoc.NewLine{},
  },
}

var videoMacroShouldNotUseImagesdirAttributeToResolveTargetIfTargetIsAUrl = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "imagesdir",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "assets",
        },
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "video::",
    },
    &asciidoc.Link{
      AttributeList: nil,
      URL: asciidoc.URL{
        Scheme: "http://",
        Path: asciidoc.Set{
          &asciidoc.String{
            Value: "example.org/videos/cats-vs-dogs.avi",
          },
        },
      },
    },
    &asciidoc.NewLine{},
  },
}

var videoMacroShouldOutputCustomHtmlWithIframeForVimeoService = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "imagesdir",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "assets",
        },
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "audio::podcast.mp3[]",
    },
    &asciidoc.NewLine{},
  },
}

var audioMacroShouldNotUseImagesdirAttributeToResolveTargetIfTargetIsAUrl = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "imagesdir",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "assets",
        },
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "video::",
    },
    &asciidoc.Link{
      AttributeList: nil,
      URL: asciidoc.URL{
        Scheme: "http://",
        Path: asciidoc.Set{
          &asciidoc.String{
            Value: "example.org/podcast.mp3",
          },
        },
      },
    },
    &asciidoc.NewLine{},
  },
}

var audioMacroShouldHonorAllOptions = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "icons",
      Set: nil,
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Paragraph{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "TIP",
            },
          },
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "You can use icons for admonitions by setting the 'icons' attribute.",
        },
        &asciidoc.NewLine{},
      },
      Admonition: 0,
    },
  },
}

var canResolveIconRelativeToCustomIconsdir = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "icons",
      Set: nil,
    },
    &asciidoc.AttributeEntry{
      Name: "iconsdir",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "icons",
        },
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Paragraph{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "TIP",
            },
          },
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "You can use icons for admonitions by setting the 'icons' attribute.",
        },
        &asciidoc.NewLine{},
      },
      Admonition: 0,
    },
  },
}

var shouldAddFileExtensionToCustomIconIfNotSpecified = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "icons",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "font",
        },
      },
    },
    &asciidoc.AttributeEntry{
      Name: "iconsdir",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "images/icons",
        },
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Paragraph{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "TIP",
            },
          },
        },
        &asciidoc.NamedAttribute{
          Name: "icon",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "a",
            },
          },
          Quote: 0,
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "Override the icon of an admonition block using an attribute",
        },
        &asciidoc.NewLine{},
      },
      Admonition: 0,
    },
  },
}

var shouldAllowIcontypeToBeSpecifiedWhenUsingBuiltInAdmonitionIcon = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Paragraph{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "TIP",
            },
          },
        },
        &asciidoc.NamedAttribute{
          Name: "icon",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "hint",
            },
          },
          Quote: 0,
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "Set the icontype using either the icontype attribute on the icons attribute.",
        },
        &asciidoc.NewLine{},
      },
      Admonition: 0,
    },
  },
}

var embedsBase64EncodedDataUriOfIconWhenDataUriAttributeIsSetAndSafeModeLevelIsLessThanSecure = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "icons",
      Set: nil,
    },
    &asciidoc.AttributeEntry{
      Name: "iconsdir",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "fixtures",
        },
      },
    },
    &asciidoc.AttributeEntry{
      Name: "icontype",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "gif",
        },
      },
    },
    &asciidoc.AttributeEntry{
      Name: "data-uri",
      Set: nil,
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Paragraph{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "TIP",
            },
          },
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "You can use icons for admonitions by setting the 'icons' attribute.",
        },
        &asciidoc.NewLine{},
      },
      Admonition: 0,
    },
  },
}

var shouldEmbedBase64EncodedDataUriOfCustomIconWhenDataUriAttributeIsSet = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "icons",
      Set: nil,
    },
    &asciidoc.AttributeEntry{
      Name: "iconsdir",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "fixtures",
        },
      },
    },
    &asciidoc.AttributeEntry{
      Name: "icontype",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "gif",
        },
      },
    },
    &asciidoc.AttributeEntry{
      Name: "data-uri",
      Set: nil,
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Paragraph{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "TIP",
            },
          },
        },
        &asciidoc.NamedAttribute{
          Name: "icon",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "tip",
            },
          },
          Quote: 0,
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "You can set a custom icon using the icon attribute on the block.",
        },
        &asciidoc.NewLine{},
      },
      Admonition: 0,
    },
  },
}

var doesNotEmbedBase64EncodedDataUriOfIconWhenSafeModeLevelIsSecureOrGreater = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "icons",
      Set: nil,
    },
    &asciidoc.AttributeEntry{
      Name: "iconsdir",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "fixtures",
        },
      },
    },
    &asciidoc.AttributeEntry{
      Name: "icontype",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "gif",
        },
      },
    },
    &asciidoc.AttributeEntry{
      Name: "data-uri",
      Set: nil,
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Paragraph{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "TIP",
            },
          },
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "You can use icons for admonitions by setting the 'icons' attribute.",
        },
        &asciidoc.NewLine{},
      },
      Admonition: 0,
    },
  },
}

var cleansReferenceToAncestorDirectoriesBeforeReadingIconIfSafeModeLevelIsAtLeastSafe = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "icons",
      Set: nil,
    },
    &asciidoc.AttributeEntry{
      Name: "iconsdir",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "../fixtures",
        },
      },
    },
    &asciidoc.AttributeEntry{
      Name: "icontype",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "gif",
        },
      },
    },
    &asciidoc.AttributeEntry{
      Name: "data-uri",
      Set: nil,
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Paragraph{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "TIP",
            },
          },
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "You can use icons for admonitions by setting the 'icons' attribute.",
        },
        &asciidoc.NewLine{},
      },
      Admonition: 0,
    },
  },
}

var shouldImportFontAwesomeAndUseFontBasedIconsWhenValueOfIconsAttributeIsFont = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "icons",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "font",
        },
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Paragraph{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "TIP",
            },
          },
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "You can use icons for admonitions by setting the 'icons' attribute.",
        },
        &asciidoc.NewLine{},
      },
      Admonition: 0,
    },
  },
}

var fontBasedIconShouldNotOverrideIconSpecifiedOnAdmonition = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "icons",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "font",
        },
      },
    },
    &asciidoc.AttributeEntry{
      Name: "iconsdir",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "images/icons",
        },
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Paragraph{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "TIP",
            },
          },
        },
        &asciidoc.NamedAttribute{
          Name: "icon",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "a.png",
            },
          },
          Quote: 0,
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "Override the icon of an admonition block using an attribute",
        },
        &asciidoc.NewLine{},
      },
      Admonition: 0,
    },
  },
}

var shouldUseHttpUriSchemeForAssetsWhenAssetUriSchemeIsHttp = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "asset-uri-scheme",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "http",
        },
      },
    },
    &asciidoc.AttributeEntry{
      Name: "icons",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "font",
        },
      },
    },
    &asciidoc.AttributeEntry{
      Name: "source-highlighter",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "highlightjs",
        },
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Paragraph{
      AttributeList: nil,
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "You can control the URI scheme used for assets with the asset-uri-scheme attribute",
        },
        &asciidoc.NewLine{},
      },
      Admonition: 2,
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Paragraph{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "source",
            },
          },
        },
        &asciidoc.PositionalAttribute{
          Offset: 1,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "ruby",
            },
          },
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "puts \"AsciiDoc, FTW!\"",
        },
        &asciidoc.NewLine{},
      },
      Admonition: 0,
    },
  },
}

var shouldUseNoUriSchemeForAssetsWhenAssetUriSchemeIsBlank = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "asset-uri-scheme",
      Set: nil,
    },
    &asciidoc.AttributeEntry{
      Name: "icons",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "font",
        },
      },
    },
    &asciidoc.AttributeEntry{
      Name: "source-highlighter",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "highlightjs",
        },
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Paragraph{
      AttributeList: nil,
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "You can control the URI scheme used for assets with the asset-uri-scheme attribute",
        },
        &asciidoc.NewLine{},
      },
      Admonition: 2,
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Paragraph{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "source",
            },
          },
        },
        &asciidoc.PositionalAttribute{
          Offset: 1,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "ruby",
            },
          },
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "puts \"AsciiDoc, FTW!\"",
        },
        &asciidoc.NewLine{},
      },
      Admonition: 0,
    },
  },
}

var restrictsAccessToAncestorDirectoriesWhenSafeModeLevelIsAtLeastSafe = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "```",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "puts \"Hello, World!\"",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "```",
    },
    &asciidoc.NewLine{},
  },
}

var shouldNotRecognizeFencedCodeBlocksWithMoreThanThreeDelimiters = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
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
    asciidoc.EmptyLine{
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

var shouldSupportFencedCodeBlocksWithLanguages = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "```ruby",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "puts \"Hello, World!\"",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "```",
    },
    &asciidoc.NewLine{},
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "``` javascript",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "alert(\"Hello, World!\")",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "```",
    },
    &asciidoc.NewLine{},
  },
}

var shouldSupportFencedCodeBlocksWithLanguagesAndNumbering = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "```ruby,numbered",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "puts \"Hello, World!\"",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "```",
    },
    &asciidoc.NewLine{},
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "``` javascript, numbered",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "alert(\"Hello, World!\")",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "```",
    },
    &asciidoc.NewLine{},
  },
}

var shouldAllowSourceStyleToBeSpecifiedOnLiteralBlock = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.LiteralBlock{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "source",
            },
          },
        },
      },
      Delimiter: asciidoc.Delimiter{
        Type: 6,
        Length: 4,
      },
      LineList: asciidoc.LineList{
        "console.log('Hello, World!')",
      },
    },
  },
}

var shouldAllowSourceStyleAndLanguageToBeSpecifiedOnLiteralBlock = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.LiteralBlock{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "source",
            },
          },
        },
        &asciidoc.PositionalAttribute{
          Offset: 1,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "js",
            },
          },
        },
      },
      Delimiter: asciidoc.Delimiter{
        Type: 6,
        Length: 4,
      },
      LineList: asciidoc.LineList{
        "console.log('Hello, World!')",
      },
    },
  },
}

var shouldMakeAbstractOnOpenBlockWithoutTitleAQuoteBlockForArticle = &asciidoc.Document{
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
        &asciidoc.OpenBlock{
          AttributeList: asciidoc.AttributeList{
            &asciidoc.PositionalAttribute{
              Offset: 0,
              ImpliedName: "",
              Val: asciidoc.Set{
                &asciidoc.String{
                  Value: "abstract",
                },
              },
            },
          },
          Delimiter: asciidoc.Delimiter{
            Type: 7,
            Length: 2,
          },
          Set: asciidoc.Set{
            &asciidoc.String{
              Value: "This article is about stuff.",
            },
            &asciidoc.NewLine{},
            asciidoc.EmptyLine{
              Text: "",
            },
            &asciidoc.String{
              Value: "And other stuff.",
            },
            &asciidoc.NewLine{},
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
              Value: "content",
            },
            &asciidoc.NewLine{},
          },
          Title: asciidoc.Set{
            &asciidoc.String{
              Value: "Section One",
            },
          },
          Level: 1,
        },
      },
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "Article",
        },
      },
      Level: 0,
    },
  },
}

var shouldMakeAbstractOnOpenBlockWithTitleAQuoteBlockWithTitleForArticle = &asciidoc.Document{
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
        &asciidoc.OpenBlock{
          AttributeList: asciidoc.AttributeList{
            &asciidoc.TitleAttribute{
              Val: asciidoc.Set{
                &asciidoc.String{
                  Value: "My abstract",
                },
              },
            },
            &asciidoc.PositionalAttribute{
              Offset: 0,
              ImpliedName: "",
              Val: asciidoc.Set{
                &asciidoc.String{
                  Value: "abstract",
                },
              },
            },
          },
          Delimiter: asciidoc.Delimiter{
            Type: 7,
            Length: 2,
          },
          Set: asciidoc.Set{
            &asciidoc.String{
              Value: "This article is about stuff.",
            },
            &asciidoc.NewLine{},
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
              Value: "content",
            },
            &asciidoc.NewLine{},
          },
          Title: asciidoc.Set{
            &asciidoc.String{
              Value: "Section One",
            },
          },
          Level: 1,
        },
      },
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "Article",
        },
      },
      Level: 0,
    },
  },
}

var shouldAllowAbstractInDocumentWithTitleIfDoctypeIsBook = &asciidoc.Document{
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
        &asciidoc.Paragraph{
          AttributeList: asciidoc.AttributeList{
            &asciidoc.PositionalAttribute{
              Offset: 0,
              ImpliedName: "",
              Val: asciidoc.Set{
                &asciidoc.String{
                  Value: "abstract",
                },
              },
            },
          },
          Set: asciidoc.Set{
            &asciidoc.String{
              Value: "Abstract for book with title is valid",
            },
            &asciidoc.NewLine{},
          },
          Admonition: 0,
        },
      },
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "Book",
        },
      },
      Level: 0,
    },
  },
}

var shouldNotAllowAbstractAsDirectChildOfDocumentIfDoctypeIsBook = &asciidoc.Document{
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
    &asciidoc.Paragraph{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "abstract",
            },
          },
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "Abstract for book without title is invalid.",
        },
        &asciidoc.NewLine{},
      },
      Admonition: 0,
    },
  },
}

var shouldMakeAbstractOnOpenBlockWithoutTitleConvertedToDocBook = &asciidoc.Document{
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
        &asciidoc.OpenBlock{
          AttributeList: asciidoc.AttributeList{
            &asciidoc.PositionalAttribute{
              Offset: 0,
              ImpliedName: "",
              Val: asciidoc.Set{
                &asciidoc.String{
                  Value: "abstract",
                },
              },
            },
          },
          Delimiter: asciidoc.Delimiter{
            Type: 7,
            Length: 2,
          },
          Set: asciidoc.Set{
            &asciidoc.String{
              Value: "This article is about stuff.",
            },
            &asciidoc.NewLine{},
            asciidoc.EmptyLine{
              Text: "",
            },
            &asciidoc.String{
              Value: "And other stuff.",
            },
            &asciidoc.NewLine{},
          },
        },
      },
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "Article",
        },
      },
      Level: 0,
    },
  },
}

var shouldMakeAbstractOnOpenBlockWithTitleConvertedToDocBook = &asciidoc.Document{
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
        &asciidoc.OpenBlock{
          AttributeList: asciidoc.AttributeList{
            &asciidoc.TitleAttribute{
              Val: asciidoc.Set{
                &asciidoc.String{
                  Value: "My abstract",
                },
              },
            },
            &asciidoc.PositionalAttribute{
              Offset: 0,
              ImpliedName: "",
              Val: asciidoc.Set{
                &asciidoc.String{
                  Value: "abstract",
                },
              },
            },
          },
          Delimiter: asciidoc.Delimiter{
            Type: 7,
            Length: 2,
          },
          Set: asciidoc.Set{
            &asciidoc.String{
              Value: "This article is about stuff.",
            },
            &asciidoc.NewLine{},
          },
        },
      },
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "Article",
        },
      },
      Level: 0,
    },
  },
}

var shouldAllowAbstractInDocumentWithTitleIfDoctypeIsBookConvertedToDocBook = &asciidoc.Document{
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
        &asciidoc.Paragraph{
          AttributeList: asciidoc.AttributeList{
            &asciidoc.PositionalAttribute{
              Offset: 0,
              ImpliedName: "",
              Val: asciidoc.Set{
                &asciidoc.String{
                  Value: "abstract",
                },
              },
            },
          },
          Set: asciidoc.Set{
            &asciidoc.String{
              Value: "Abstract for book with title is valid",
            },
            &asciidoc.NewLine{},
          },
          Admonition: 0,
        },
      },
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "Book",
        },
      },
      Level: 0,
    },
  },
}

var shouldNotAllowAbstractAsDirectChildOfDocumentIfDoctypeIsBookConvertedToDocBook = &asciidoc.Document{
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
    &asciidoc.Paragraph{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "abstract",
            },
          },
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "Abstract for book is invalid.",
        },
        &asciidoc.NewLine{},
      },
      Admonition: 0,
    },
  },
}

var shouldAcceptPartintroOnOpenBlockWithoutTitle = &asciidoc.Document{
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
      },
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "Book",
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
        &asciidoc.OpenBlock{
          AttributeList: asciidoc.AttributeList{
            &asciidoc.PositionalAttribute{
              Offset: 0,
              ImpliedName: "",
              Val: asciidoc.Set{
                &asciidoc.String{
                  Value: "partintro",
                },
              },
            },
          },
          Delimiter: asciidoc.Delimiter{
            Type: 7,
            Length: 2,
          },
          Set: asciidoc.Set{
            &asciidoc.String{
              Value: "This is a part intro.",
            },
            &asciidoc.NewLine{},
            asciidoc.EmptyLine{
              Text: "",
            },
            &asciidoc.String{
              Value: "It can have multiple paragraphs.",
            },
            &asciidoc.NewLine{},
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
              Value: "content",
            },
            &asciidoc.NewLine{},
          },
          Title: asciidoc.Set{
            &asciidoc.String{
              Value: "Chapter 1",
            },
          },
          Level: 1,
        },
      },
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "Part 1",
        },
      },
      Level: 0,
    },
  },
}

var shouldAcceptPartintroOnOpenBlockWithTitle = &asciidoc.Document{
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
      },
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "Book",
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
        &asciidoc.OpenBlock{
          AttributeList: asciidoc.AttributeList{
            &asciidoc.TitleAttribute{
              Val: asciidoc.Set{
                &asciidoc.String{
                  Value: "Intro title",
                },
              },
            },
            &asciidoc.PositionalAttribute{
              Offset: 0,
              ImpliedName: "",
              Val: asciidoc.Set{
                &asciidoc.String{
                  Value: "partintro",
                },
              },
            },
          },
          Delimiter: asciidoc.Delimiter{
            Type: 7,
            Length: 2,
          },
          Set: asciidoc.Set{
            &asciidoc.String{
              Value: "This is a part intro with a title.",
            },
            &asciidoc.NewLine{},
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
              Value: "content",
            },
            &asciidoc.NewLine{},
          },
          Title: asciidoc.Set{
            &asciidoc.String{
              Value: "Chapter 1",
            },
          },
          Level: 1,
        },
      },
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "Part 1",
        },
      },
      Level: 0,
    },
  },
}

var shouldExcludePartintroIfNotAChildOfPart = &asciidoc.Document{
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
        &asciidoc.Paragraph{
          AttributeList: asciidoc.AttributeList{
            &asciidoc.PositionalAttribute{
              Offset: 0,
              ImpliedName: "",
              Val: asciidoc.Set{
                &asciidoc.String{
                  Value: "partintro",
                },
              },
            },
          },
          Set: asciidoc.Set{
            &asciidoc.String{
              Value: "part intro paragraph",
            },
            &asciidoc.NewLine{},
          },
          Admonition: 0,
        },
      },
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "Book",
        },
      },
      Level: 0,
    },
  },
}

var shouldNotAllowPartintroUnlessDoctypeIsBook = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Paragraph{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "partintro",
            },
          },
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "part intro paragraph",
        },
        &asciidoc.NewLine{},
      },
      Admonition: 0,
    },
  },
}

var shouldAcceptPartintroOnOpenBlockWithoutTitleConvertedToDocBook = &asciidoc.Document{
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
      },
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "Book",
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
        &asciidoc.OpenBlock{
          AttributeList: asciidoc.AttributeList{
            &asciidoc.PositionalAttribute{
              Offset: 0,
              ImpliedName: "",
              Val: asciidoc.Set{
                &asciidoc.String{
                  Value: "partintro",
                },
              },
            },
          },
          Delimiter: asciidoc.Delimiter{
            Type: 7,
            Length: 2,
          },
          Set: asciidoc.Set{
            &asciidoc.String{
              Value: "This is a part intro.",
            },
            &asciidoc.NewLine{},
            asciidoc.EmptyLine{
              Text: "",
            },
            &asciidoc.String{
              Value: "It can have multiple paragraphs.",
            },
            &asciidoc.NewLine{},
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
              Value: "content",
            },
            &asciidoc.NewLine{},
          },
          Title: asciidoc.Set{
            &asciidoc.String{
              Value: "Chapter 1",
            },
          },
          Level: 1,
        },
      },
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "Part 1",
        },
      },
      Level: 0,
    },
  },
}

var shouldAcceptPartintroOnOpenBlockWithTitleConvertedToDocBook = &asciidoc.Document{
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
      },
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "Book",
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
        &asciidoc.OpenBlock{
          AttributeList: asciidoc.AttributeList{
            &asciidoc.TitleAttribute{
              Val: asciidoc.Set{
                &asciidoc.String{
                  Value: "Intro title",
                },
              },
            },
            &asciidoc.PositionalAttribute{
              Offset: 0,
              ImpliedName: "",
              Val: asciidoc.Set{
                &asciidoc.String{
                  Value: "partintro",
                },
              },
            },
          },
          Delimiter: asciidoc.Delimiter{
            Type: 7,
            Length: 2,
          },
          Set: asciidoc.Set{
            &asciidoc.String{
              Value: "This is a part intro with a title.",
            },
            &asciidoc.NewLine{},
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
              Value: "content",
            },
            &asciidoc.NewLine{},
          },
          Title: asciidoc.Set{
            &asciidoc.String{
              Value: "Chapter 1",
            },
          },
          Level: 1,
        },
      },
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "Part 1",
        },
      },
      Level: 0,
    },
  },
}

var shouldExcludePartintroIfNotAChildOfPartConvertedToDocBook = &asciidoc.Document{
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
        &asciidoc.Paragraph{
          AttributeList: asciidoc.AttributeList{
            &asciidoc.PositionalAttribute{
              Offset: 0,
              ImpliedName: "",
              Val: asciidoc.Set{
                &asciidoc.String{
                  Value: "partintro",
                },
              },
            },
          },
          Set: asciidoc.Set{
            &asciidoc.String{
              Value: "part intro paragraph",
            },
            &asciidoc.NewLine{},
          },
          Admonition: 0,
        },
      },
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "Book",
        },
      },
      Level: 0,
    },
  },
}

var shouldNotAllowPartintroUnlessDoctypeIsBookConvertedToDocBook = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Paragraph{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "partintro",
            },
          },
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "part intro paragraph",
        },
        &asciidoc.NewLine{},
      },
      Admonition: 0,
    },
  },
}

var processorShouldNotCrashIfSubsAreEmpty = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.LiteralBlock{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.NamedAttribute{
          Name: "subs",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: ",",
            },
          },
          Quote: 2,
        },
      },
      Delimiter: asciidoc.Delimiter{
        Type: 6,
        Length: 4,
      },
      LineList: asciidoc.LineList{
        "content",
      },
    },
  },
}

var shouldBeAbleToAppendSubsToDefaultBlockSubstitutionList = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "application",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "Asciidoctor",
        },
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.LiteralBlock{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.NamedAttribute{
          Name: "subs",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "+attributes,+macros",
            },
          },
          Quote: 2,
        },
      },
      Delimiter: asciidoc.Delimiter{
        Type: 6,
        Length: 4,
      },
      LineList: asciidoc.LineList{
        "{application}",
      },
    },
  },
}

var shouldBeAbleToPrependSubsToDefaultBlockSubstitutionList = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "application",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "Asciidoctor",
        },
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.LiteralBlock{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.NamedAttribute{
          Name: "subs",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "attributes+",
            },
          },
          Quote: 2,
        },
      },
      Delimiter: asciidoc.Delimiter{
        Type: 6,
        Length: 4,
      },
      LineList: asciidoc.LineList{
        "{application}",
      },
    },
  },
}

var shouldBeAbleToRemoveSubsToDefaultBlockSubstitutionList = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Paragraph{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.NamedAttribute{
          Name: "subs",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "-quotes,-replacements",
            },
          },
          Quote: 2,
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "content",
        },
        &asciidoc.NewLine{},
      },
      Admonition: 0,
    },
  },
}

var shouldBeAbleToPrependAppendAndRemoveSubsFromDefaultBlockSubstitutionList = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "application",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "asciidoctor",
        },
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.LiteralBlock{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.NamedAttribute{
          Name: "subs",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "attributes+,-verbatim,+specialcharacters,+macros",
            },
          },
          Quote: 2,
        },
      },
      Delimiter: asciidoc.Delimiter{
        Type: 6,
        Length: 4,
      },
      LineList: asciidoc.LineList{
        "https://{application}.org[{gt}{gt}] <1>",
      },
    },
  },
}

var shouldBeAbleToSetSubsThenModifyThem = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Paragraph{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.NamedAttribute{
          Name: "subs",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "verbatim,-callouts",
            },
          },
          Quote: 2,
        },
      },
      Set: asciidoc.Set{
        &asciidoc.Italic{
          AttributeList: nil,
          Set: asciidoc.Set{
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

var shouldNotRecognizeBlockAnchorWithIllegalIdCharacters = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "[[illegal$id,Reference Text]]",
    },
    &asciidoc.NewLine{},
    &asciidoc.Listing{
      AttributeList: nil,
      Delimiter: asciidoc.Delimiter{
        Type: 5,
        Length: 4,
      },
      LineList: asciidoc.LineList{
        "content",
      },
    },
  },
}

var shouldNotRecognizeBlockAnchorThatStartsWithDigit = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "[[3-blind-mice]]",
    },
    &asciidoc.NewLine{},
    &asciidoc.OpenBlock{
      AttributeList: nil,
      Delimiter: asciidoc.Delimiter{
        Type: 7,
        Length: 2,
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "see how they run",
        },
        &asciidoc.NewLine{},
      },
    },
  },
}

var shouldRecognizeBlockAnchorThatStartsWithColon = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.OpenBlock{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.AnchorAttribute{
          ID: &asciidoc.String{
            Value: ":idname",
          },
          Label: nil,
        },
      },
      Delimiter: asciidoc.Delimiter{
        Type: 7,
        Length: 2,
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "content",
        },
        &asciidoc.NewLine{},
      },
    },
  },
}

var shouldUseSpecifiedIdAndReftextWhenRegisteringBlockReference = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Listing{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.AnchorAttribute{
          ID: &asciidoc.String{
            Value: "debian",
          },
          Label: asciidoc.Set{
            &asciidoc.String{
              Value: "Debian Install",
            },
          },
        },
        &asciidoc.TitleAttribute{
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Installation on Debian",
            },
          },
        },
      },
      Delimiter: asciidoc.Delimiter{
        Type: 5,
        Length: 4,
      },
      LineList: asciidoc.LineList{
        "$ apt-get install asciidoctor",
      },
    },
  },
}

var shouldAllowSquareBracketsInBlockReferenceText = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "[[debian,[Debian] Install]]",
    },
    &asciidoc.NewLine{},
    &asciidoc.Listing{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TitleAttribute{
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Installation on Debian",
            },
          },
        },
      },
      Delimiter: asciidoc.Delimiter{
        Type: 5,
        Length: 4,
      },
      LineList: asciidoc.LineList{
        "$ apt-get install asciidoctor",
      },
    },
  },
}

var shouldAllowCommaInBlockReferenceText = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Listing{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.AnchorAttribute{
          ID: &asciidoc.String{
            Value: "debian",
          },
          Label: asciidoc.Set{
            &asciidoc.String{
              Value: " Debian, Ubuntu",
            },
          },
        },
        &asciidoc.TitleAttribute{
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Installation on Debian",
            },
          },
        },
      },
      Delimiter: asciidoc.Delimiter{
        Type: 5,
        Length: 4,
      },
      LineList: asciidoc.LineList{
        "$ apt-get install asciidoctor",
      },
    },
  },
}

var shouldResolveAttributeReferenceInTitleUsingAttributeDefinedAtLocationOfBlock = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Section{ // p0
      AttributeList: nil,
      Set: asciidoc.Set{
        &asciidoc.AttributeEntry{
          Name: "foo",
          Set: asciidoc.Set{
            &asciidoc.String{
              Value: "baz",
            },
          },
        },
        asciidoc.EmptyLine{
          Text: "",
        },
        &asciidoc.String{
          Value: "intro paragraph. see ",
        },
        &asciidoc.CrossReference{
          Set: nil,
          ID: "free-standing",
        },
        &asciidoc.String{
          Value: ".",
        },
        &asciidoc.NewLine{},
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
        asciidoc.EmptyLine{
          Text: "",
        },
        &asciidoc.Paragraph{
          AttributeList: asciidoc.AttributeList{
            &asciidoc.TitleAttribute{
              Val: asciidoc.Set{
                &asciidoc.String{
                  Value: "foo is ",
                },
                &asciidoc.UserAttributeReference{
                  Value: "foo",
                },
              },
            },
            &asciidoc.PositionalAttribute{
              Offset: 0,
              ImpliedName: "",
              Val: asciidoc.Set{
                &asciidoc.String{
                  Value: "#formal-para",
                },
              },
            },
          },
          Set: asciidoc.Set{
            &asciidoc.String{
              Value: "paragraph with title",
            },
            &asciidoc.NewLine{},
          },
          Admonition: 0,
        },
        asciidoc.EmptyLine{
          Text: "",
        },
        &asciidoc.Section{
          AttributeList: asciidoc.AttributeList{
            &asciidoc.PositionalAttribute{
              Offset: 0,
              ImpliedName: "",
              Val: asciidoc.Set{
                &asciidoc.String{
                  Value: "discrete#free-standing",
                },
              },
            },
          },
          Set: nil,
          Title: asciidoc.Set{
            &asciidoc.String{
              Value: "foo is still {foo}",
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

var shouldSubstituteAttributeReferencesInReftextWhenRegisteringBlockReference = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "label-tiger",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "Tiger",
        },
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.SidebarBlock{
      Delimiter: asciidoc.Delimiter{
        Type: 8,
        Length: 4,
      },
      AttributeList: asciidoc.AttributeList{
        &asciidoc.AnchorAttribute{
          ID: &asciidoc.String{
            Value: "tiger-evolution",
          },
          Label: asciidoc.Set{
            &asciidoc.String{
              Value: "Evolution of the ",
            },
            &asciidoc.UserAttributeReference{
              Value: "label-tiger",
            },
          },
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "Information about the evolution of the tiger.",
        },
        &asciidoc.NewLine{},
      },
    },
  },
}

var shouldUseSpecifiedReftextWhenRegisteringBlockReference = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Listing{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.AnchorAttribute{
          ID: &asciidoc.String{
            Value: "debian",
          },
          Label: nil,
        },
        &asciidoc.NamedAttribute{
          Name: "reftext",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Debian Install",
            },
          },
          Quote: 2,
        },
        &asciidoc.TitleAttribute{
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Installation on Debian",
            },
          },
        },
      },
      Delimiter: asciidoc.Delimiter{
        Type: 5,
        Length: 4,
      },
      LineList: asciidoc.LineList{
        "$ apt-get install asciidoctor",
      },
    },
  },
}


