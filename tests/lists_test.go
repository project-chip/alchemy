package tests

import (
	"testing"

	"github.com/project-chip/alchemy/asciidoc"
)

func TestLists(t *testing.T) {
	listsTests.run(t)
}

var listsTests = parseTests{

	{"dash elements with no blank lines", "asciidoctor/lists_test_dash_elements_with_no_blank_lines.adoc", listsTestDashElementsWithNoBlankLines},

	{"dash elements separated by blank lines should merge lists", "asciidoctor/lists_test_dash_elements_separated_by_blank_lines_should_merge_lists.adoc", listsTestDashElementsSeparatedByBlankLinesShouldMergeLists},

	{"dash elements with interspersed line comments should be skipped and not break list", "asciidoctor/lists_test_dash_elements_with_interspersed_line_comments_should_be_skipped_and_not_break_list.adoc", listsTestDashElementsWithInterspersedLineCommentsShouldBeSkippedAndNotBreakList},

	{"dash elements separated by a line comment offset by blank lines should not merge lists", "asciidoctor/lists_test_dash_elements_separated_by_a_line_comment_offset_by_blank_lines_should_not_merge_lists.adoc", listsTestDashElementsSeparatedByALineCommentOffsetByBlankLinesShouldNotMergeLists},

	{"dash elements separated by a block title offset by a blank line should not merge lists", "asciidoctor/lists_test_dash_elements_separated_by_a_block_title_offset_by_a_blank_line_should_not_merge_lists.adoc", listsTestDashElementsSeparatedByABlockTitleOffsetByABlankLineShouldNotMergeLists},

	{"dash elements separated by an attribute entry offset by a blank line should not merge lists", "asciidoctor/lists_test_dash_elements_separated_by_an_attribute_entry_offset_by_a_blank_line_should_not_merge_lists.adoc", listsTestDashElementsSeparatedByAnAttributeEntryOffsetByABlankLineShouldNotMergeLists},

	{"a non-indented wrapped line is folded into text of list item", "asciidoctor/lists_test_a_non_indented_wrapped_line_is_folded_into_text_of_list_item.adoc", listsTestANonIndentedWrappedLineIsFoldedIntoTextOfListItem},

	{"a non-indented wrapped line that resembles a block title is folded into text of list item", "asciidoctor/lists_test_a_non_indented_wrapped_line_that_resembles_a_block_title_is_folded_into_text_of_list_item.adoc", listsTestANonIndentedWrappedLineThatResemblesABlockTitleIsFoldedIntoTextOfListItem},

	{"a non-indented wrapped line that resembles an attribute entry is folded into text of list item", "asciidoctor/lists_test_a_non_indented_wrapped_line_that_resembles_an_attribute_entry_is_folded_into_text_of_list_item.adoc", listsTestANonIndentedWrappedLineThatResemblesAnAttributeEntryIsFoldedIntoTextOfListItem},

	{"a list item with a nested marker terminates non-indented paragraph for text of list item", "asciidoctor/lists_test_a_list_item_with_a_nested_marker_terminates_non_indented_paragraph_for_text_of_list_item.adoc", listsTestAListItemWithANestedMarkerTerminatesNonIndentedParagraphForTextOfListItem},

	{"a list item for a different list terminates non-indented paragraph for text of list item", "asciidoctor/lists_test_a_list_item_for_a_different_list_terminates_non_indented_paragraph_for_text_of_list_item.adoc", listsTestAListItemForADifferentListTerminatesNonIndentedParagraphForTextOfListItem},

	{"an indented wrapped line is unindented and folded into text of list item", "asciidoctor/lists_test_an_indented_wrapped_line_is_unindented_and_folded_into_text_of_list_item.adoc", listsTestAnIndentedWrappedLineIsUnindentedAndFoldedIntoTextOfListItem},

	{"wrapped list item with hanging indent followed by non-indented line", "asciidoctor/lists_test_wrapped_list_item_with_hanging_indent_followed_by_non_indented_line.adoc", listsTestWrappedListItemWithHangingIndentFollowedByNonIndentedLine},

	{"a list item with a nested marker terminates indented paragraph for text of list item", "asciidoctor/lists_test_a_list_item_with_a_nested_marker_terminates_indented_paragraph_for_text_of_list_item.adoc", listsTestAListItemWithANestedMarkerTerminatesIndentedParagraphForTextOfListItem},

	{"a list item for a different list terminates indented paragraph for text of list item", "asciidoctor/lists_test_a_list_item_for_a_different_list_terminates_indented_paragraph_for_text_of_list_item.adoc", listsTestAListItemForADifferentListTerminatesIndentedParagraphForTextOfListItem},

	{"a literal paragraph offset by blank lines in list content is appended as a literal block", "asciidoctor/lists_test_a_literal_paragraph_offset_by_blank_lines_in_list_content_is_appended_as_a_literal_block.adoc", listsTestALiteralParagraphOffsetByBlankLinesInListContentIsAppendedAsALiteralBlock},

	{"should escape special characters in all literal paragraphs attached to list item", "asciidoctor/lists_test_should_escape_special_characters_in_all_literal_paragraphs_attached_to_list_item.adoc", listsTestShouldEscapeSpecialCharactersInAllLiteralParagraphsAttachedToListItem},

	{"a literal paragraph offset by a blank line in list content followed by line with continuation is appended as two blocks", "asciidoctor/lists_test_a_literal_paragraph_offset_by_a_blank_line_in_list_content_followed_by_line_with_continuation_is_appended_as_two_blocks.adoc", listsTestALiteralParagraphOffsetByABlankLineInListContentFollowedByLineWithContinuationIsAppendedAsTwoBlocks},

	{"an admonition paragraph attached by a line continuation to a list item with wrapped text should produce admonition", "asciidoctor/lists_test_an_admonition_paragraph_attached_by_a_line_continuation_to_a_list_item_with_wrapped_text_should_produce_admonition.adoc", listsTestAnAdmonitionParagraphAttachedByALineContinuationToAListItemWithWrappedTextShouldProduceAdmonition},

	{"paragraph-like blocks attached to an ancestor list item by a list continuation should produce blocks", "asciidoctor/lists_test_paragraph_like_blocks_attached_to_an_ancestor_list_item_by_a_list_continuation_should_produce_blocks.adoc", listsTestParagraphLikeBlocksAttachedToAnAncestorListItemByAListContinuationShouldProduceBlocks},

	{"should not inherit block attributes from previous block when block is attached using a list continuation", "asciidoctor/lists_test_should_not_inherit_block_attributes_from_previous_block_when_block_is_attached_using_a_list_continuation.adoc", listsTestShouldNotInheritBlockAttributesFromPreviousBlockWhenBlockIsAttachedUsingAListContinuation},

	{"should continue to parse blocks attached by a list continuation after block is dropped", "asciidoctor/lists_test_should_continue_to_parse_blocks_attached_by_a_list_continuation_after_block_is_dropped.adoc", listsTestShouldContinueToParseBlocksAttachedByAListContinuationAfterBlockIsDropped},

	{"appends line as paragraph if attached by continuation following line comment", "asciidoctor/lists_test_appends_line_as_paragraph_if_attached_by_continuation_following_line_comment.adoc", listsTestAppendsLineAsParagraphIfAttachedByContinuationFollowingLineComment},

	{"a literal paragraph with a line that appears as a list item that is followed by a continuation should create two blocks", "asciidoctor/lists_test_a_literal_paragraph_with_a_line_that_appears_as_a_list_item_that_is_followed_by_a_continuation_should_create_two_blocks.adoc", listsTestALiteralParagraphWithALineThatAppearsAsAListItemThatIsFollowedByAContinuationShouldCreateTwoBlocks},

	{"consecutive literal paragraph offset by blank lines in list content are appended as a literal blocks", "asciidoctor/lists_test_consecutive_literal_paragraph_offset_by_blank_lines_in_list_content_are_appended_as_a_literal_blocks.adoc", listsTestConsecutiveLiteralParagraphOffsetByBlankLinesInListContentAreAppendedAsALiteralBlocks},

	{"a literal paragraph without a trailing blank line consumes following list items", "asciidoctor/lists_test_a_literal_paragraph_without_a_trailing_blank_line_consumes_following_list_items.adoc", listsTestALiteralParagraphWithoutATrailingBlankLineConsumesFollowingListItems},

	{"asterisk elements with no blank lines", "asciidoctor/lists_test_asterisk_elements_with_no_blank_lines.adoc", listsTestAsteriskElementsWithNoBlankLines},

	{"asterisk elements separated by blank lines should merge lists", "asciidoctor/lists_test_asterisk_elements_separated_by_blank_lines_should_merge_lists.adoc", listsTestAsteriskElementsSeparatedByBlankLinesShouldMergeLists},

	{"asterisk elements with interspersed line comments should be skipped and not break list", "asciidoctor/lists_test_asterisk_elements_with_interspersed_line_comments_should_be_skipped_and_not_break_list.adoc", listsTestAsteriskElementsWithInterspersedLineCommentsShouldBeSkippedAndNotBreakList},

	{"asterisk elements separated by a line comment offset by blank lines should not merge lists", "asciidoctor/lists_test_asterisk_elements_separated_by_a_line_comment_offset_by_blank_lines_should_not_merge_lists.adoc", listsTestAsteriskElementsSeparatedByALineCommentOffsetByBlankLinesShouldNotMergeLists},

	{"asterisk elements separated by a block title offset by a blank line should not merge lists", "asciidoctor/lists_test_asterisk_elements_separated_by_a_block_title_offset_by_a_blank_line_should_not_merge_lists.adoc", listsTestAsteriskElementsSeparatedByABlockTitleOffsetByABlankLineShouldNotMergeLists},

	{"asterisk elements separated by an attribute entry offset by a blank line should not merge lists", "asciidoctor/lists_test_asterisk_elements_separated_by_an_attribute_entry_offset_by_a_blank_line_should_not_merge_lists.adoc", listsTestAsteriskElementsSeparatedByAnAttributeEntryOffsetByABlankLineShouldNotMergeLists},

	{"list should terminate before next lower section heading", "asciidoctor/lists_test_list_should_terminate_before_next_lower_section_heading.adoc", listsTestListShouldTerminateBeforeNextLowerSectionHeading},

	{"list should terminate before next lower section heading with implicit id", "asciidoctor/lists_test_list_should_terminate_before_next_lower_section_heading_with_implicit_id.adoc", listsTestListShouldTerminateBeforeNextLowerSectionHeadingWithImplicitId},

	{"should not find section title immediately below last list item", "asciidoctor/lists_test_should_not_find_section_title_immediately_below_last_list_item.adoc", listsTestShouldNotFindSectionTitleImmediatelyBelowLastListItem},

	{"quoted text", "asciidoctor/lists_test_quoted_text.adoc", listsTestQuotedText},

	{"attribute substitutions", "asciidoctor/lists_test_attribute_substitutions.adoc", listsTestAttributeSubstitutions},

	{"leading dot is treated as text not block title", "asciidoctor/lists_test_leading_dot_is_treated_as_text_not_block_title.adoc", listsTestLeadingDotIsTreatedAsTextNotBlockTitle},

	{"word ending sentence on continuing line not treated as a list item", "asciidoctor/lists_test_word_ending_sentence_on_continuing_line_not_treated_as_a_list_item.adoc", listsTestWordEndingSentenceOnContinuingLineNotTreatedAsAListItem},

	{"should discover anchor at start of unordered list item text and register it as a reference", "asciidoctor/lists_test_should_discover_anchor_at_start_of_unordered_list_item_text_and_register_it_as_a_reference.adoc", listsTestShouldDiscoverAnchorAtStartOfUnorderedListItemTextAndRegisterItAsAReference},

	{"should discover anchor at start of ordered list item text and register it as a reference", "asciidoctor/lists_test_should_discover_anchor_at_start_of_ordered_list_item_text_and_register_it_as_a_reference.adoc", listsTestShouldDiscoverAnchorAtStartOfOrderedListItemTextAndRegisterItAsAReference},

	{"should discover anchor at start of callout list item text and register it as a reference", "asciidoctor/lists_test_should_discover_anchor_at_start_of_callout_list_item_text_and_register_it_as_a_reference.adoc", listsTestShouldDiscoverAnchorAtStartOfCalloutListItemTextAndRegisterItAsAReference},

	{"asterisk element mixed with dash elements should be nested", "asciidoctor/lists_test_asterisk_element_mixed_with_dash_elements_should_be_nested.adoc", listsTestAsteriskElementMixedWithDashElementsShouldBeNested},

	{"dash element mixed with asterisks elements should be nested", "asciidoctor/lists_test_dash_element_mixed_with_asterisks_elements_should_be_nested.adoc", listsTestDashElementMixedWithAsterisksElementsShouldBeNested},

	{"lines prefixed with alternating list markers separated by blank lines should be nested", "asciidoctor/lists_test_lines_prefixed_with_alternating_list_markers_separated_by_blank_lines_should_be_nested.adoc", listsTestLinesPrefixedWithAlternatingListMarkersSeparatedByBlankLinesShouldBeNested},

	{"nested elements (2) with asterisks", "asciidoctor/lists_test_nested_elements_(2)_with_asterisks.adoc", listsTestNestedElements2WithAsterisks},

	{"nested elements (3) with asterisks", "asciidoctor/lists_test_nested_elements_(3)_with_asterisks.adoc", listsTestNestedElements3WithAsterisks},

	{"nested elements (4) with asterisks", "asciidoctor/lists_test_nested_elements_(4)_with_asterisks.adoc", listsTestNestedElements4WithAsterisks},

	{"nested elements (5) with asterisks", "asciidoctor/lists_test_nested_elements_(5)_with_asterisks.adoc", listsTestNestedElements5WithAsterisks},

	{"level of unordered list should match section level", "asciidoctor/lists_test_level_of_unordered_list_should_match_section_level.adoc", listsTestLevelOfUnorderedListShouldMatchSectionLevel},

	{"does not recognize lists with repeating unicode bullets", "asciidoctor/lists_test_does_not_recognize_lists_with_repeating_unicode_bullets.adoc", listsTestDoesNotRecognizeListsWithRepeatingUnicodeBullets},

	{"nested ordered elements (3)", "asciidoctor/lists_test_nested_ordered_elements_(3).adoc", listsTestNestedOrderedElements3},

	{"level of ordered list should match section level", "asciidoctor/lists_test_level_of_ordered_list_should_match_section_level.adoc", listsTestLevelOfOrderedListShouldMatchSectionLevel},

	{"nested unordered inside ordered elements", "asciidoctor/lists_test_nested_unordered_inside_ordered_elements.adoc", listsTestNestedUnorderedInsideOrderedElements},

	{"nested ordered inside unordered elements", "asciidoctor/lists_test_nested_ordered_inside_unordered_elements.adoc", listsTestNestedOrderedInsideUnorderedElements},

	{"three levels of alternating unordered and ordered elements", "asciidoctor/lists_test_three_levels_of_alternating_unordered_and_ordered_elements.adoc", listsTestThreeLevelsOfAlternatingUnorderedAndOrderedElements},

	{"lines with alternating markers of unordered and ordered list types separated by blank lines should be nested", "asciidoctor/lists_test_lines_with_alternating_markers_of_unordered_and_ordered_list_types_separated_by_blank_lines_should_be_nested.adoc", listsTestLinesWithAlternatingMarkersOfUnorderedAndOrderedListTypesSeparatedByBlankLinesShouldBeNested},

	{"list item with literal content should not consume nested list of different type", "asciidoctor/lists_test_list_item_with_literal_content_should_not_consume_nested_list_of_different_type.adoc", listsTestListItemWithLiteralContentShouldNotConsumeNestedListOfDifferentType},

	{"nested list item does not eat the title of the following detached block", "asciidoctor/lists_test_nested_list_item_does_not_eat_the_title_of_the_following_detached_block.adoc", listsTestNestedListItemDoesNotEatTheTitleOfTheFollowingDetachedBlock},

	{"lines with alternating markers of bulleted and description list types separated by blank lines should be nested", "asciidoctor/lists_test_lines_with_alternating_markers_of_bulleted_and_description_list_types_separated_by_blank_lines_should_be_nested.adoc", listsTestLinesWithAlternatingMarkersOfBulletedAndDescriptionListTypesSeparatedByBlankLinesShouldBeNested},

	{"nested ordered with attribute inside unordered elements", "asciidoctor/lists_test_nested_ordered_with_attribute_inside_unordered_elements.adoc", listsTestNestedOrderedWithAttributeInsideUnorderedElements},

	{"adjacent list continuation line attaches following paragraph", "asciidoctor/lists_test_adjacent_list_continuation_line_attaches_following_paragraph.adoc", listsTestAdjacentListContinuationLineAttachesFollowingParagraph},

	{"adjacent list continuation line attaches following block", "asciidoctor/lists_test_adjacent_list_continuation_line_attaches_following_block.adoc", listsTestAdjacentListContinuationLineAttachesFollowingBlock},

	{"adjacent list continuation line attaches following block with block attributes", "asciidoctor/lists_test_adjacent_list_continuation_line_attaches_following_block_with_block_attributes.adoc", listsTestAdjacentListContinuationLineAttachesFollowingBlockWithBlockAttributes},

	{"trailing block attribute line attached by continuation should not create block", "asciidoctor/lists_test_trailing_block_attribute_line_attached_by_continuation_should_not_create_block.adoc", listsTestTrailingBlockAttributeLineAttachedByContinuationShouldNotCreateBlock},

	{"trailing block title line attached by continuation should not create block", "asciidoctor/lists_test_trailing_block_title_line_attached_by_continuation_should_not_create_block.adoc", listsTestTrailingBlockTitleLineAttachedByContinuationShouldNotCreateBlock},

	{"consecutive blocks in list continuation attach to list item", "asciidoctor/lists_test_consecutive_blocks_in_list_continuation_attach_to_list_item.adoc", listsTestConsecutiveBlocksInListContinuationAttachToListItem},

	{"list item with hanging indent followed by block attached by list continuation", "asciidoctor/lists_test_list_item_with_hanging_indent_followed_by_block_attached_by_list_continuation.adoc", listsTestListItemWithHangingIndentFollowedByBlockAttachedByListContinuation},

	{"list item paragraph in list item and nested list item", "asciidoctor/lists_test_list_item_paragraph_in_list_item_and_nested_list_item.adoc", listsTestListItemParagraphInListItemAndNestedListItem},

	{"trailing list continuations should attach to list items at respective levels", "asciidoctor/lists_test_trailing_list_continuations_should_attach_to_list_items_at_respective_levels.adoc", listsTestTrailingListContinuationsShouldAttachToListItemsAtRespectiveLevels},

	{"trailing list continuations should attach to list items of different types at respective levels", "asciidoctor/lists_test_trailing_list_continuations_should_attach_to_list_items_of_different_types_at_respective_levels.adoc", listsTestTrailingListContinuationsShouldAttachToListItemsOfDifferentTypesAtRespectiveLevels},

	{"repeated list continuations should attach to list items at respective levels", "asciidoctor/lists_test_repeated_list_continuations_should_attach_to_list_items_at_respective_levels.adoc", listsTestRepeatedListContinuationsShouldAttachToListItemsAtRespectiveLevels},

	{"repeated list continuations attached directly to list item should attach to list items at respective levels", "asciidoctor/lists_test_repeated_list_continuations_attached_directly_to_list_item_should_attach_to_list_items_at_respective_levels.adoc", listsTestRepeatedListContinuationsAttachedDirectlyToListItemShouldAttachToListItemsAtRespectiveLevels},

	{"repeated list continuations should attach to list items at respective levels ignoring blank lines", "asciidoctor/lists_test_repeated_list_continuations_should_attach_to_list_items_at_respective_levels_ignoring_blank_lines.adoc", listsTestRepeatedListContinuationsShouldAttachToListItemsAtRespectiveLevelsIgnoringBlankLines},

	{"trailing list continuations should ignore preceding blank lines", "asciidoctor/lists_test_trailing_list_continuations_should_ignore_preceding_blank_lines.adoc", listsTestTrailingListContinuationsShouldIgnorePrecedingBlankLines},

	{"indented outline list item with different marker offset by a blank line should be recognized as a nested list", "asciidoctor/lists_test_indented_outline_list_item_with_different_marker_offset_by_a_blank_line_should_be_recognized_as_a_nested_list.adoc", listsTestIndentedOutlineListItemWithDifferentMarkerOffsetByABlankLineShouldBeRecognizedAsANestedList},

	{"indented description list item inside outline list item offset by a blank line should be recognized as a nested list", "asciidoctor/lists_test_indented_description_list_item_inside_outline_list_item_offset_by_a_blank_line_should_be_recognized_as_a_nested_list.adoc", listsTestIndentedDescriptionListItemInsideOutlineListItemOffsetByABlankLineShouldBeRecognizedAsANestedList},

	{"consecutive list continuation lines are folded", "asciidoctor/lists_test_consecutive_list_continuation_lines_are_folded.adoc", listsTestConsecutiveListContinuationLinesAreFolded},

	{"should warn if unterminated block is detected in list item", "asciidoctor/lists_test_should_warn_if_unterminated_block_is_detected_in_list_item.adoc", listsTestShouldWarnIfUnterminatedBlockIsDetectedInListItem},

	{"dot elements with no blank lines", "asciidoctor/lists_test_dot_elements_with_no_blank_lines.adoc", listsTestDotElementsWithNoBlankLines},

	{"should represent explicit role attribute as style class", "asciidoctor/lists_test_should_represent_explicit_role_attribute_as_style_class.adoc", listsTestShouldRepresentExplicitRoleAttributeAsStyleClass},

	{"should base list style on marker length rather than list depth", "asciidoctor/lists_test_should_base_list_style_on_marker_length_rather_than_list_depth.adoc", listsTestShouldBaseListStyleOnMarkerLengthRatherThanListDepth},

	{"should allow list style to be specified explicitly when using markers with implicit style", "asciidoctor/lists_test_should_allow_list_style_to_be_specified_explicitly_when_using_markers_with_implicit_style.adoc", listsTestShouldAllowListStyleToBeSpecifiedExplicitlyWhenUsingMarkersWithImplicitStyle},

	{"should represent custom numbering and explicit role attribute as style classes", "asciidoctor/lists_test_should_represent_custom_numbering_and_explicit_role_attribute_as_style_classes.adoc", listsTestShouldRepresentCustomNumberingAndExplicitRoleAttributeAsStyleClasses},

	{"should set reversed attribute on list if reversed option is set", "asciidoctor/lists_test_should_set_reversed_attribute_on_list_if_reversed_option_is_set.adoc", listsTestShouldSetReversedAttributeOnListIfReversedOptionIsSet},

	{"should represent implicit role attribute as style class", "asciidoctor/lists_test_should_represent_implicit_role_attribute_as_style_class.adoc", listsTestShouldRepresentImplicitRoleAttributeAsStyleClass},

	{"should represent custom numbering and implicit role attribute as style classes", "asciidoctor/lists_test_should_represent_custom_numbering_and_implicit_role_attribute_as_style_classes.adoc", listsTestShouldRepresentCustomNumberingAndImplicitRoleAttributeAsStyleClasses},

	{"dot elements separated by blank lines should merge lists", "asciidoctor/lists_test_dot_elements_separated_by_blank_lines_should_merge_lists.adoc", listsTestDotElementsSeparatedByBlankLinesShouldMergeLists},

	{"dot elements with interspersed line comments should be skipped and not break list", "asciidoctor/lists_test_dot_elements_with_interspersed_line_comments_should_be_skipped_and_not_break_list.adoc", listsTestDotElementsWithInterspersedLineCommentsShouldBeSkippedAndNotBreakList},

	{"dot elements separated by line comment offset by blank lines should not merge lists", "asciidoctor/lists_test_dot_elements_separated_by_line_comment_offset_by_blank_lines_should_not_merge_lists.adoc", listsTestDotElementsSeparatedByLineCommentOffsetByBlankLinesShouldNotMergeLists},

	{"dot elements separated by a block title offset by a blank line should not merge lists", "asciidoctor/lists_test_dot_elements_separated_by_a_block_title_offset_by_a_blank_line_should_not_merge_lists.adoc", listsTestDotElementsSeparatedByABlockTitleOffsetByABlankLineShouldNotMergeLists},

	{"dot elements separated by an attribute entry offset by a blank line should not merge lists", "asciidoctor/lists_test_dot_elements_separated_by_an_attribute_entry_offset_by_a_blank_line_should_not_merge_lists.adoc", listsTestDotElementsSeparatedByAnAttributeEntryOffsetByABlankLineShouldNotMergeLists},

	{"should use start number in docbook5 backend", "asciidoctor/lists_test_should_use_start_number_in_docbook_5_backend.adoc", listsTestShouldUseStartNumberInDocbook5Backend},

	{"should warn if explicit uppercase roman numerals in list are out of sequence", "asciidoctor/lists_test_should_warn_if_explicit_uppercase_roman_numerals_in_list_are_out_of_sequence.adoc", listsTestShouldWarnIfExplicitUppercaseRomanNumeralsInListAreOutOfSequence},

	{"should warn if explicit lowercase roman numerals in list are out of sequence", "asciidoctor/lists_test_should_warn_if_explicit_lowercase_roman_numerals_in_list_are_out_of_sequence.adoc", listsTestShouldWarnIfExplicitLowercaseRomanNumeralsInListAreOutOfSequence},

	{"should not parse a bare dlist delimiter as a dlist", "asciidoctor/lists_test_should_not_parse_a_bare_dlist_delimiter_as_a_dlist.adoc", listsTestShouldNotParseABareDlistDelimiterAsADlist},

	{"should parse sibling items using same rules", "asciidoctor/lists_test_should_parse_sibling_items_using_same_rules.adoc", listsTestShouldParseSiblingItemsUsingSameRules},

	{"should allow term to end with a semicolon when using double semicolon delimiter", "asciidoctor/lists_test_should_allow_term_to_end_with_a_semicolon_when_using_double_semicolon_delimiter.adoc", listsTestShouldAllowTermToEndWithASemicolonWhenUsingDoubleSemicolonDelimiter},

	{"single-line indented adjacent elements", "asciidoctor/lists_test_single_line_indented_adjacent_elements.adoc", listsTestSingleLineIndentedAdjacentElements},

	{"single-line elements separated by blank line should create a single list", "asciidoctor/lists_test_single_line_elements_separated_by_blank_line_should_create_a_single_list.adoc", listsTestSingleLineElementsSeparatedByBlankLineShouldCreateASingleList},

	{"a line comment between elements should divide them into separate lists", "asciidoctor/lists_test_a_line_comment_between_elements_should_divide_them_into_separate_lists.adoc", listsTestALineCommentBetweenElementsShouldDivideThemIntoSeparateLists},

	{"a ruler between elements should divide them into separate lists", "asciidoctor/lists_test_a_ruler_between_elements_should_divide_them_into_separate_lists.adoc", listsTestARulerBetweenElementsShouldDivideThemIntoSeparateLists},

	{"a block title between elements should divide them into separate lists", "asciidoctor/lists_test_a_block_title_between_elements_should_divide_them_into_separate_lists.adoc", listsTestABlockTitleBetweenElementsShouldDivideThemIntoSeparateLists},

	{"multi-line elements with paragraph content", "asciidoctor/lists_test_multi_line_elements_with_paragraph_content.adoc", listsTestMultiLineElementsWithParagraphContent},

	{"multi-line elements with indented paragraph content", "asciidoctor/lists_test_multi_line_elements_with_indented_paragraph_content.adoc", listsTestMultiLineElementsWithIndentedParagraphContent},

	{"multi-line elements with indented paragraph content that includes comment lines", "asciidoctor/lists_test_multi_line_elements_with_indented_paragraph_content_that_includes_comment_lines.adoc", listsTestMultiLineElementsWithIndentedParagraphContentThatIncludesCommentLines},

	{"should not strip comment line in literal paragraph block attached to list item", "asciidoctor/lists_test_should_not_strip_comment_line_in_literal_paragraph_block_attached_to_list_item.adoc", listsTestShouldNotStripCommentLineInLiteralParagraphBlockAttachedToListItem},

	{"multi-line element with paragraph starting with multiple dashes should not be seen as list", "asciidoctor/lists_test_multi_line_element_with_paragraph_starting_with_multiple_dashes_should_not_be_seen_as_list.adoc", listsTestMultiLineElementWithParagraphStartingWithMultipleDashesShouldNotBeSeenAsList},

	{"multi-line element with multiple terms", "asciidoctor/lists_test_multi_line_element_with_multiple_terms.adoc", listsTestMultiLineElementWithMultipleTerms},

	{"consecutive terms share same varlistentry in docbook", "asciidoctor/lists_test_consecutive_terms_share_same_varlistentry_in_docbook.adoc", listsTestConsecutiveTermsShareSameVarlistentryInDocbook},

	{"multi-line elements with blank line before paragraph content", "asciidoctor/lists_test_multi_line_elements_with_blank_line_before_paragraph_content.adoc", listsTestMultiLineElementsWithBlankLineBeforeParagraphContent},

	{"multi-line elements with paragraph and literal content", "asciidoctor/lists_test_multi_line_elements_with_paragraph_and_literal_content.adoc", listsTestMultiLineElementsWithParagraphAndLiteralContent},

	{"mixed single and multi-line adjacent elements", "asciidoctor/lists_test_mixed_single_and_multi_line_adjacent_elements.adoc", listsTestMixedSingleAndMultiLineAdjacentElements},

	{"should discover anchor at start of description term text and register it as a reference", "asciidoctor/lists_test_should_discover_anchor_at_start_of_description_term_text_and_register_it_as_a_reference.adoc", listsTestShouldDiscoverAnchorAtStartOfDescriptionTermTextAndRegisterItAsAReference},

	{"missing space before term does not produce description list", "asciidoctor/lists_test_missing_space_before_term_does_not_produce_description_list.adoc", listsTestMissingSpaceBeforeTermDoesNotProduceDescriptionList},

	{"literal block inside description list", "asciidoctor/lists_test_literal_block_inside_description_list.adoc", listsTestLiteralBlockInsideDescriptionList},

	{"literal block inside description list with trailing line continuation", "asciidoctor/lists_test_literal_block_inside_description_list_with_trailing_line_continuation.adoc", listsTestLiteralBlockInsideDescriptionListWithTrailingLineContinuation},

	{"multiple listing blocks inside description list", "asciidoctor/lists_test_multiple_listing_blocks_inside_description_list.adoc", listsTestMultipleListingBlocksInsideDescriptionList},

	{"open block inside description list", "asciidoctor/lists_test_open_block_inside_description_list.adoc", listsTestOpenBlockInsideDescriptionList},

	{"paragraph attached by a list continuation on either side in a description list", "asciidoctor/lists_test_paragraph_attached_by_a_list_continuation_on_either_side_in_a_description_list.adoc", listsTestParagraphAttachedByAListContinuationOnEitherSideInADescriptionList},

	{"paragraph attached by a list continuation on either side to a multi-line element in a description list", "asciidoctor/lists_test_paragraph_attached_by_a_list_continuation_on_either_side_to_a_multi_line_element_in_a_description_list.adoc", listsTestParagraphAttachedByAListContinuationOnEitherSideToAMultiLineElementInADescriptionList},

	{"should continue to parse subsequent blocks attached to list item after first block is dropped", "asciidoctor/lists_test_should_continue_to_parse_subsequent_blocks_attached_to_list_item_after_first_block_is_dropped.adoc", listsTestShouldContinueToParseSubsequentBlocksAttachedToListItemAfterFirstBlockIsDropped},

	{"verse paragraph inside a description list", "asciidoctor/lists_test_verse_paragraph_inside_a_description_list.adoc", listsTestVerseParagraphInsideADescriptionList},

	{"list inside a description list", "asciidoctor/lists_test_list_inside_a_description_list.adoc", listsTestListInsideADescriptionList},

	{"list inside a description list offset by blank lines", "asciidoctor/lists_test_list_inside_a_description_list_offset_by_blank_lines.adoc", listsTestListInsideADescriptionListOffsetByBlankLines},

	{"should only grab one line following last item if item has no inline description", "asciidoctor/lists_test_should_only_grab_one_line_following_last_item_if_item_has_no_inline_description.adoc", listsTestShouldOnlyGrabOneLineFollowingLastItemIfItemHasNoInlineDescription},

	{"should only grab one literal line following last item if item has no inline description", "asciidoctor/lists_test_should_only_grab_one_literal_line_following_last_item_if_item_has_no_inline_description.adoc", listsTestShouldOnlyGrabOneLiteralLineFollowingLastItemIfItemHasNoInlineDescription},

	{"should append subsequent paragraph literals to list item as block content", "asciidoctor/lists_test_should_append_subsequent_paragraph_literals_to_list_item_as_block_content.adoc", listsTestShouldAppendSubsequentParagraphLiteralsToListItemAsBlockContent},

	{"should not match comment line that looks like description list term", "asciidoctor/lists_test_should_not_match_comment_line_that_looks_like_description_list_term.adoc", listsTestShouldNotMatchCommentLineThatLooksLikeDescriptionListTerm},

	{"should not match comment line following list that looks like description list term", "asciidoctor/lists_test_should_not_match_comment_line_following_list_that_looks_like_description_list_term.adoc", listsTestShouldNotMatchCommentLineFollowingListThatLooksLikeDescriptionListTerm},

	{"should not match comment line that looks like sibling description list term", "asciidoctor/lists_test_should_not_match_comment_line_that_looks_like_sibling_description_list_term.adoc", listsTestShouldNotMatchCommentLineThatLooksLikeSiblingDescriptionListTerm},

	{"should not hang on description list item in list that begins with ///", "asciidoctor/lists_test_should_not_hang_on_description_list_item_in_list_that_begins_with.adoc", listsTestShouldNotHangOnDescriptionListItemInListThatBeginsWith},

	{"should not hang on sibling description list item that begins with ///", "asciidoctor/lists_test_should_not_hang_on_sibling_description_list_item_that_begins_with.adoc", listsTestShouldNotHangOnSiblingDescriptionListItemThatBeginsWith},

	{"should skip dlist term that begins with // unless it begins with ///", "asciidoctor/lists_test_should_skip_dlist_term_that_begins_with____unless_it_begins_with.adoc", listsTestShouldSkipDlistTermThatBeginsWithUnlessItBeginsWith},

	{"more than 4 consecutive colons should become part of description list term", "asciidoctor/lists_test_more_than_4_consecutive_colons_should_become_part_of_description_list_term.adoc", listsTestMoreThan4ConsecutiveColonsShouldBecomePartOfDescriptionListTerm},

	{"text method of dd node should return nil if dd node only contains blocks", "asciidoctor/lists_test_text_method_of_dd_node_should_return_nil_if_dd_node_only_contains_blocks.adoc", listsTestTextMethodOfDdNodeShouldReturnNilIfDdNodeOnlyContainsBlocks},

	{"should not parse a nested dlist delimiter without a term as a dlist", "asciidoctor/lists_test_should_not_parse_a_nested_dlist_delimiter_without_a_term_as_a_dlist.adoc", listsTestShouldNotParseANestedDlistDelimiterWithoutATermAsADlist},

	{"should not parse a nested indented dlist delimiter without a term as a dlist", "asciidoctor/lists_test_should_not_parse_a_nested_indented_dlist_delimiter_without_a_term_as_a_dlist.adoc", listsTestShouldNotParseANestedIndentedDlistDelimiterWithoutATermAsADlist},

	{"single-line adjacent nested elements", "asciidoctor/lists_test_single_line_adjacent_nested_elements.adoc", listsTestSingleLineAdjacentNestedElements},

	{"single-line adjacent maximum nested elements", "asciidoctor/lists_test_single_line_adjacent_maximum_nested_elements.adoc", listsTestSingleLineAdjacentMaximumNestedElements},

	{"single-line nested elements separated by blank line at top level", "asciidoctor/lists_test_single_line_nested_elements_separated_by_blank_line_at_top_level.adoc", listsTestSingleLineNestedElementsSeparatedByBlankLineAtTopLevel},

	{"single-line nested elements separated by blank line at nested level", "asciidoctor/lists_test_single_line_nested_elements_separated_by_blank_line_at_nested_level.adoc", listsTestSingleLineNestedElementsSeparatedByBlankLineAtNestedLevel},

	{"single-line adjacent nested elements with alternate delimiters", "asciidoctor/lists_test_single_line_adjacent_nested_elements_with_alternate_delimiters.adoc", listsTestSingleLineAdjacentNestedElementsWithAlternateDelimiters},

	{"multi-line adjacent nested elements", "asciidoctor/lists_test_multi_line_adjacent_nested_elements.adoc", listsTestMultiLineAdjacentNestedElements},

	{"multi-line nested elements separated by blank line at nested level repeated", "asciidoctor/lists_test_multi_line_nested_elements_separated_by_blank_line_at_nested_level_repeated.adoc", listsTestMultiLineNestedElementsSeparatedByBlankLineAtNestedLevelRepeated},

	{"multi-line element with indented nested element", "asciidoctor/lists_test_multi_line_element_with_indented_nested_element.adoc", listsTestMultiLineElementWithIndentedNestedElement},

	{"mixed single and multi-line elements with indented nested elements", "asciidoctor/lists_test_mixed_single_and_multi_line_elements_with_indented_nested_elements.adoc", listsTestMixedSingleAndMultiLineElementsWithIndentedNestedElements},

	{"multi-line elements with first paragraph folded to text with adjacent nested element", "asciidoctor/lists_test_multi_line_elements_with_first_paragraph_folded_to_text_with_adjacent_nested_element.adoc", listsTestMultiLineElementsWithFirstParagraphFoldedToTextWithAdjacentNestedElement},

	{"nested dlist attached by list continuation should not consume detached paragraph", "asciidoctor/lists_test_nested_dlist_attached_by_list_continuation_should_not_consume_detached_paragraph.adoc", listsTestNestedDlistAttachedByListContinuationShouldNotConsumeDetachedParagraph},

	{"nested dlist with attached block offset by empty line", "asciidoctor/lists_test_nested_dlist_with_attached_block_offset_by_empty_line.adoc", listsTestNestedDlistWithAttachedBlockOffsetByEmptyLine},

	{"should convert glossary list with proper semantics", "asciidoctor/lists_test_should_convert_glossary_list_with_proper_semantics.adoc", listsTestShouldConvertGlossaryListWithProperSemantics},

	{"consecutive glossary terms should share same glossentry element in docbook", "asciidoctor/lists_test_consecutive_glossary_terms_should_share_same_glossentry_element_in_docbook.adoc", listsTestConsecutiveGlossaryTermsShouldShareSameGlossentryElementInDocbook},

	{"should convert horizontal list with proper markup", "asciidoctor/lists_test_should_convert_horizontal_list_with_proper_markup.adoc", listsTestShouldConvertHorizontalListWithProperMarkup},

	{"should set col widths of item and label if specified", "asciidoctor/lists_test_should_set_col_widths_of_item_and_label_if_specified.adoc", listsTestShouldSetColWidthsOfItemAndLabelIfSpecified},

	{"should set col widths of item and label in docbook if specified", "asciidoctor/lists_test_should_set_col_widths_of_item_and_label_in_docbook_if_specified.adoc", listsTestShouldSetColWidthsOfItemAndLabelInDocbookIfSpecified},

	{"should add strong class to label if strong option is set", "asciidoctor/lists_test_should_add_strong_class_to_label_if_strong_option_is_set.adoc", listsTestShouldAddStrongClassToLabelIfStrongOptionIsSet},

	{"consecutive terms in horizontal list should share same cell", "asciidoctor/lists_test_consecutive_terms_in_horizontal_list_should_share_same_cell.adoc", listsTestConsecutiveTermsInHorizontalListShouldShareSameCell},

	{"consecutive terms in horizontal list should share same entry in docbook", "asciidoctor/lists_test_consecutive_terms_in_horizontal_list_should_share_same_entry_in_docbook.adoc", listsTestConsecutiveTermsInHorizontalListShouldShareSameEntryInDocbook},

	{"should convert horizontal list in docbook with proper markup", "asciidoctor/lists_test_should_convert_horizontal_list_in_docbook_with_proper_markup.adoc", listsTestShouldConvertHorizontalListInDocbookWithProperMarkup},

	{"should convert qanda list in HTML with proper semantics", "asciidoctor/lists_test_should_convert_qanda_list_in_html_with_proper_semantics.adoc", listsTestShouldConvertQandaListInHtmlWithProperSemantics},

	{"should convert qanda list in DocBook with proper semantics", "asciidoctor/lists_test_should_convert_qanda_list_in_doc_book_with_proper_semantics.adoc", listsTestShouldConvertQandaListInDocBookWithProperSemantics},

	{"consecutive questions should share same question element in docbook", "asciidoctor/lists_test_consecutive_questions_should_share_same_question_element_in_docbook.adoc", listsTestConsecutiveQuestionsShouldShareSameQuestionElementInDocbook},

	{"should convert bibliography list with proper semantics", "asciidoctor/lists_test_should_convert_bibliography_list_with_proper_semantics.adoc", listsTestShouldConvertBibliographyListWithProperSemantics},

	{"should convert bibliography list with proper semantics to DocBook", "asciidoctor/lists_test_should_convert_bibliography_list_with_proper_semantics_to_doc_book.adoc", listsTestShouldConvertBibliographyListWithProperSemanticsToDocBook},

	{"should warn if a bibliography ID is already in use", "asciidoctor/lists_test_should_warn_if_a_bibliography_id_is_already_in_use.adoc", listsTestShouldWarnIfABibliographyIdIsAlreadyInUse},

	{"should automatically add bibliography style to top-level lists in bibliography section", "asciidoctor/lists_test_should_automatically_add_bibliography_style_to_top_level_lists_in_bibliography_section.adoc", listsTestShouldAutomaticallyAddBibliographyStyleToTopLevelListsInBibliographySection},

	{"should not recognize bibliography anchor that begins with a digit", "asciidoctor/lists_test_should_not_recognize_bibliography_anchor_that_begins_with_a_digit.adoc", listsTestShouldNotRecognizeBibliographyAnchorThatBeginsWithADigit},

	{"should recognize bibliography anchor that contains a digit but does not start with one", "asciidoctor/lists_test_should_recognize_bibliography_anchor_that_contains_a_digit_but_does_not_start_with_one.adoc", listsTestShouldRecognizeBibliographyAnchorThatContainsADigitButDoesNotStartWithOne},

	{"should catalog bibliography anchors in bibliography list", "asciidoctor/lists_test_should_catalog_bibliography_anchors_in_bibliography_list.adoc", listsTestShouldCatalogBibliographyAnchorsInBibliographyList},

	{"should use reftext from bibliography anchor at xref and entry", "asciidoctor/lists_test_should_use_reftext_from_bibliography_anchor_at_xref_and_entry.adoc", listsTestShouldUseReftextFromBibliographyAnchorAtXrefAndEntry},

	{"should assign reftext of bibliography anchor to xreflabel in DocBook backend", "asciidoctor/lists_test_should_assign_reftext_of_bibliography_anchor_to_xreflabel_in_doc_book_backend.adoc", listsTestShouldAssignReftextOfBibliographyAnchorToXreflabelInDocBookBackend},

	{"folds text from subsequent line", "asciidoctor/lists_test_folds_text_from_subsequent_line.adoc", listsTestFoldsTextFromSubsequentLine},

	{"folds text from first line after blank lines", "asciidoctor/lists_test_folds_text_from_first_line_after_blank_lines.adoc", listsTestFoldsTextFromFirstLineAfterBlankLines},

	{"folds text from first line after blank line and immediately preceding next item", "asciidoctor/lists_test_folds_text_from_first_line_after_blank_line_and_immediately_preceding_next_item.adoc", listsTestFoldsTextFromFirstLineAfterBlankLineAndImmediatelyPrecedingNextItem},

	{"paragraph offset by blank lines does not break list if label does not have inline text", "asciidoctor/lists_test_paragraph_offset_by_blank_lines_does_not_break_list_if_label_does_not_have_inline_text.adoc", listsTestParagraphOffsetByBlankLinesDoesNotBreakListIfLabelDoesNotHaveInlineText},

	{"folds text from first line after comment line", "asciidoctor/lists_test_folds_text_from_first_line_after_comment_line.adoc", listsTestFoldsTextFromFirstLineAfterCommentLine},

	{"folds text from line following comment line offset by blank line", "asciidoctor/lists_test_folds_text_from_line_following_comment_line_offset_by_blank_line.adoc", listsTestFoldsTextFromLineFollowingCommentLineOffsetByBlankLine},

	{"folds text from subsequent indented line", "asciidoctor/lists_test_folds_text_from_subsequent_indented_line.adoc", listsTestFoldsTextFromSubsequentIndentedLine},

	{"folds text from indented line after blank line", "asciidoctor/lists_test_folds_text_from_indented_line_after_blank_line.adoc", listsTestFoldsTextFromIndentedLineAfterBlankLine},

	{"folds text that looks like ruler offset by blank line", "asciidoctor/lists_test_folds_text_that_looks_like_ruler_offset_by_blank_line.adoc", listsTestFoldsTextThatLooksLikeRulerOffsetByBlankLine},

	{"folds text that looks like ruler offset by blank line and line comment", "asciidoctor/lists_test_folds_text_that_looks_like_ruler_offset_by_blank_line_and_line_comment.adoc", listsTestFoldsTextThatLooksLikeRulerOffsetByBlankLineAndLineComment},

	{"folds text that looks like ruler and the line following it offset by blank line", "asciidoctor/lists_test_folds_text_that_looks_like_ruler_and_the_line_following_it_offset_by_blank_line.adoc", listsTestFoldsTextThatLooksLikeRulerAndTheLineFollowingItOffsetByBlankLine},

	{"folds text that looks like title offset by blank line", "asciidoctor/lists_test_folds_text_that_looks_like_title_offset_by_blank_line.adoc", listsTestFoldsTextThatLooksLikeTitleOffsetByBlankLine},

	{"folds text that looks like title offset by blank line and line comment", "asciidoctor/lists_test_folds_text_that_looks_like_title_offset_by_blank_line_and_line_comment.adoc", listsTestFoldsTextThatLooksLikeTitleOffsetByBlankLineAndLineComment},

	{"folds text that looks like admonition offset by blank line", "asciidoctor/lists_test_folds_text_that_looks_like_admonition_offset_by_blank_line.adoc", listsTestFoldsTextThatLooksLikeAdmonitionOffsetByBlankLine},

	{"folds text that looks like section title offset by blank line", "asciidoctor/lists_test_folds_text_that_looks_like_section_title_offset_by_blank_line.adoc", listsTestFoldsTextThatLooksLikeSectionTitleOffsetByBlankLine},

	{"folds text of first literal line offset by blank line appends subsequent literals offset by blank line as blocks", "asciidoctor/lists_test_folds_text_of_first_literal_line_offset_by_blank_line_appends_subsequent_literals_offset_by_blank_line_as_blocks.adoc", listsTestFoldsTextOfFirstLiteralLineOffsetByBlankLineAppendsSubsequentLiteralsOffsetByBlankLineAsBlocks},

	{"folds text of subsequent line and appends following literal line offset by blank line as block if term has no inline description", "asciidoctor/lists_test_folds_text_of_subsequent_line_and_appends_following_literal_line_offset_by_blank_line_as_block_if_term_has_no_inline_description.adoc", listsTestFoldsTextOfSubsequentLineAndAppendsFollowingLiteralLineOffsetByBlankLineAsBlockIfTermHasNoInlineDescription},

	{"appends literal line attached by continuation as block if item has no inline description", "asciidoctor/lists_test_appends_literal_line_attached_by_continuation_as_block_if_item_has_no_inline_description.adoc", listsTestAppendsLiteralLineAttachedByContinuationAsBlockIfItemHasNoInlineDescription},

	{"appends literal line attached by continuation as block if item has no inline description followed by ruler", "asciidoctor/lists_test_appends_literal_line_attached_by_continuation_as_block_if_item_has_no_inline_description_followed_by_ruler.adoc", listsTestAppendsLiteralLineAttachedByContinuationAsBlockIfItemHasNoInlineDescriptionFollowedByRuler},

	{"appends line attached by continuation as block if item has no inline description followed by ruler", "asciidoctor/lists_test_appends_line_attached_by_continuation_as_block_if_item_has_no_inline_description_followed_by_ruler.adoc", listsTestAppendsLineAttachedByContinuationAsBlockIfItemHasNoInlineDescriptionFollowedByRuler},

	{"appends line attached by continuation as block if item has no inline description followed by block", "asciidoctor/lists_test_appends_line_attached_by_continuation_as_block_if_item_has_no_inline_description_followed_by_block.adoc", listsTestAppendsLineAttachedByContinuationAsBlockIfItemHasNoInlineDescriptionFollowedByBlock},

	{"appends block attached by continuation but not subsequent block not attached by continuation", "asciidoctor/lists_test_appends_block_attached_by_continuation_but_not_subsequent_block_not_attached_by_continuation.adoc", listsTestAppendsBlockAttachedByContinuationButNotSubsequentBlockNotAttachedByContinuation},

	{"appends list if item has no inline description", "asciidoctor/lists_test_appends_list_if_item_has_no_inline_description.adoc", listsTestAppendsListIfItemHasNoInlineDescription},

	{"appends list to first term when followed immediately by second term", "asciidoctor/lists_test_appends_list_to_first_term_when_followed_immediately_by_second_term.adoc", listsTestAppendsListToFirstTermWhenFollowedImmediatelyBySecondTerm},

	{"appends indented list to first term that is adjacent to second term", "asciidoctor/lists_test_appends_indented_list_to_first_term_that_is_adjacent_to_second_term.adoc", listsTestAppendsIndentedListToFirstTermThatIsAdjacentToSecondTerm},

	{"appends indented list to first term that is attached by a continuation and adjacent to second term", "asciidoctor/lists_test_appends_indented_list_to_first_term_that_is_attached_by_a_continuation_and_adjacent_to_second_term.adoc", listsTestAppendsIndentedListToFirstTermThatIsAttachedByAContinuationAndAdjacentToSecondTerm},

	{"appends list and paragraph block when line following list attached by continuation", "asciidoctor/lists_test_appends_list_and_paragraph_block_when_line_following_list_attached_by_continuation.adoc", listsTestAppendsListAndParagraphBlockWhenLineFollowingListAttachedByContinuation},

	{"first continued line associated with nested list item and second continued line associated with term", "asciidoctor/lists_test_first_continued_line_associated_with_nested_list_item_and_second_continued_line_associated_with_term.adoc", listsTestFirstContinuedLineAssociatedWithNestedListItemAndSecondContinuedLineAssociatedWithTerm},

	{"literal line attached by continuation swallows adjacent line that looks like term", "asciidoctor/lists_test_literal_line_attached_by_continuation_swallows_adjacent_line_that_looks_like_term.adoc", listsTestLiteralLineAttachedByContinuationSwallowsAdjacentLineThatLooksLikeTerm},

	{"line attached by continuation is appended as paragraph if term has no inline description", "asciidoctor/lists_test_line_attached_by_continuation_is_appended_as_paragraph_if_term_has_no_inline_description.adoc", listsTestLineAttachedByContinuationIsAppendedAsParagraphIfTermHasNoInlineDescription},

	{"attached paragraph does not break on adjacent nested description list term", "asciidoctor/lists_test_attached_paragraph_does_not_break_on_adjacent_nested_description_list_term.adoc", listsTestAttachedParagraphDoesNotBreakOnAdjacentNestedDescriptionListTerm},

	{"attached paragraph is terminated by adjacent sibling description list term", "asciidoctor/lists_test_attached_paragraph_is_terminated_by_adjacent_sibling_description_list_term.adoc", listsTestAttachedParagraphIsTerminatedByAdjacentSiblingDescriptionListTerm},

	{"attached styled paragraph does not break on adjacent nested description list term", "asciidoctor/lists_test_attached_styled_paragraph_does_not_break_on_adjacent_nested_description_list_term.adoc", listsTestAttachedStyledParagraphDoesNotBreakOnAdjacentNestedDescriptionListTerm},

	{"appends line as paragraph if attached by continuation following blank line and line comment when term has no inline description", "asciidoctor/lists_test_appends_line_as_paragraph_if_attached_by_continuation_following_blank_line_and_line_comment_when_term_has_no_inline_description.adoc", listsTestAppendsLineAsParagraphIfAttachedByContinuationFollowingBlankLineAndLineCommentWhenTermHasNoInlineDescription},

	{"line attached by continuation offset by blank line is appended as paragraph if term has no inline description", "asciidoctor/lists_test_line_attached_by_continuation_offset_by_blank_line_is_appended_as_paragraph_if_term_has_no_inline_description.adoc", listsTestLineAttachedByContinuationOffsetByBlankLineIsAppendedAsParagraphIfTermHasNoInlineDescription},

	{"delimited block breaks list even when term has no inline description", "asciidoctor/lists_test_delimited_block_breaks_list_even_when_term_has_no_inline_description.adoc", listsTestDelimitedBlockBreaksListEvenWhenTermHasNoInlineDescription},

	{"block attribute line above delimited block that breaks a dlist is not duplicated", "asciidoctor/lists_test_block_attribute_line_above_delimited_block_that_breaks_a_dlist_is_not_duplicated.adoc", listsTestBlockAttributeLineAboveDelimitedBlockThatBreaksADlistIsNotDuplicated},

	{"block attribute line above paragraph breaks list even when term has no inline description", "asciidoctor/lists_test_block_attribute_line_above_paragraph_breaks_list_even_when_term_has_no_inline_description.adoc", listsTestBlockAttributeLineAboveParagraphBreaksListEvenWhenTermHasNoInlineDescription},

	{"block attribute line above paragraph that breaks a dlist is not duplicated", "asciidoctor/lists_test_block_attribute_line_above_paragraph_that_breaks_a_dlist_is_not_duplicated.adoc", listsTestBlockAttributeLineAboveParagraphThatBreaksADlistIsNotDuplicated},

	{"block anchor line breaks list even when term has no inline description", "asciidoctor/lists_test_block_anchor_line_breaks_list_even_when_term_has_no_inline_description.adoc", listsTestBlockAnchorLineBreaksListEvenWhenTermHasNoInlineDescription},

	{"block attribute lines above nested horizontal list does not break list", "asciidoctor/lists_test_block_attribute_lines_above_nested_horizontal_list_does_not_break_list.adoc", listsTestBlockAttributeLinesAboveNestedHorizontalListDoesNotBreakList},

	{"block attribute lines above nested list with style does not break list", "asciidoctor/lists_test_block_attribute_lines_above_nested_list_with_style_does_not_break_list.adoc", listsTestBlockAttributeLinesAboveNestedListWithStyleDoesNotBreakList},

	{"multiple block attribute lines above nested list does not break list", "asciidoctor/lists_test_multiple_block_attribute_lines_above_nested_list_does_not_break_list.adoc", listsTestMultipleBlockAttributeLinesAboveNestedListDoesNotBreakList},

	{"multiple block attribute lines separated by empty line above nested list does not break list", "asciidoctor/lists_test_multiple_block_attribute_lines_separated_by_empty_line_above_nested_list_does_not_break_list.adoc", listsTestMultipleBlockAttributeLinesSeparatedByEmptyLineAboveNestedListDoesNotBreakList},

	{"folds text from inline description and subsequent line", "asciidoctor/lists_test_folds_text_from_inline_description_and_subsequent_line.adoc", listsTestFoldsTextFromInlineDescriptionAndSubsequentLine},

	{"folds text from inline description and subsequent lines", "asciidoctor/lists_test_folds_text_from_inline_description_and_subsequent_lines.adoc", listsTestFoldsTextFromInlineDescriptionAndSubsequentLines},

	{"folds text from inline description and line following comment line", "asciidoctor/lists_test_folds_text_from_inline_description_and_line_following_comment_line.adoc", listsTestFoldsTextFromInlineDescriptionAndLineFollowingCommentLine},

	{"folds text from inline description and subsequent indented line", "asciidoctor/lists_test_folds_text_from_inline_description_and_subsequent_indented_line.adoc", listsTestFoldsTextFromInlineDescriptionAndSubsequentIndentedLine},

	{"appends literal line offset by blank line as block if item has inline description", "asciidoctor/lists_test_appends_literal_line_offset_by_blank_line_as_block_if_item_has_inline_description.adoc", listsTestAppendsLiteralLineOffsetByBlankLineAsBlockIfItemHasInlineDescription},

	{"appends literal line offset by blank line as block and appends line after continuation as block if item has inline description", "asciidoctor/lists_test_appends_literal_line_offset_by_blank_line_as_block_and_appends_line_after_continuation_as_block_if_item_has_inline_description.adoc", listsTestAppendsLiteralLineOffsetByBlankLineAsBlockAndAppendsLineAfterContinuationAsBlockIfItemHasInlineDescription},

	{"appends line after continuation as block and literal line offset by blank line as block if item has inline description", "asciidoctor/lists_test_appends_line_after_continuation_as_block_and_literal_line_offset_by_blank_line_as_block_if_item_has_inline_description.adoc", listsTestAppendsLineAfterContinuationAsBlockAndLiteralLineOffsetByBlankLineAsBlockIfItemHasInlineDescription},

	{"appends list if item has inline description", "asciidoctor/lists_test_appends_list_if_item_has_inline_description.adoc", listsTestAppendsListIfItemHasInlineDescription},

	{"appends literal line attached by continuation as block if item has inline description followed by ruler", "asciidoctor/lists_test_appends_literal_line_attached_by_continuation_as_block_if_item_has_inline_description_followed_by_ruler.adoc", listsTestAppendsLiteralLineAttachedByContinuationAsBlockIfItemHasInlineDescriptionFollowedByRuler},

	{"line offset by blank line breaks list if term has inline description", "asciidoctor/lists_test_line_offset_by_blank_line_breaks_list_if_term_has_inline_description.adoc", listsTestLineOffsetByBlankLineBreaksListIfTermHasInlineDescription},

	{"nested term with description does not consume following heading", "asciidoctor/lists_test_nested_term_with_description_does_not_consume_following_heading.adoc", listsTestNestedTermWithDescriptionDoesNotConsumeFollowingHeading},

	{"line attached by continuation is appended as paragraph if term has inline description followed by detached paragraph", "asciidoctor/lists_test_line_attached_by_continuation_is_appended_as_paragraph_if_term_has_inline_description_followed_by_detached_paragraph.adoc", listsTestLineAttachedByContinuationIsAppendedAsParagraphIfTermHasInlineDescriptionFollowedByDetachedParagraph},

	{"line attached by continuation is appended as paragraph if term has inline description followed by detached block", "asciidoctor/lists_test_line_attached_by_continuation_is_appended_as_paragraph_if_term_has_inline_description_followed_by_detached_block.adoc", listsTestLineAttachedByContinuationIsAppendedAsParagraphIfTermHasInlineDescriptionFollowedByDetachedBlock},

	{"line attached by continuation offset by line comment is appended as paragraph if term has inline description", "asciidoctor/lists_test_line_attached_by_continuation_offset_by_line_comment_is_appended_as_paragraph_if_term_has_inline_description.adoc", listsTestLineAttachedByContinuationOffsetByLineCommentIsAppendedAsParagraphIfTermHasInlineDescription},

	{"line attached by continuation offset by blank line is appended as paragraph if term has inline description", "asciidoctor/lists_test_line_attached_by_continuation_offset_by_blank_line_is_appended_as_paragraph_if_term_has_inline_description.adoc", listsTestLineAttachedByContinuationOffsetByBlankLineIsAppendedAsParagraphIfTermHasInlineDescription},

	{"line comment offset by blank line divides lists because item has text", "asciidoctor/lists_test_line_comment_offset_by_blank_line_divides_lists_because_item_has_text.adoc", listsTestLineCommentOffsetByBlankLineDividesListsBecauseItemHasText},

	{"ruler offset by blank line divides lists because item has text", "asciidoctor/lists_test_ruler_offset_by_blank_line_divides_lists_because_item_has_text.adoc", listsTestRulerOffsetByBlankLineDividesListsBecauseItemHasText},

	{"block title offset by blank line divides lists and becomes title of second list because item has text", "asciidoctor/lists_test_block_title_offset_by_blank_line_divides_lists_and_becomes_title_of_second_list_because_item_has_text.adoc", listsTestBlockTitleOffsetByBlankLineDividesListsAndBecomesTitleOfSecondListBecauseItemHasText},

	{"does not recognize callout list denoted by markers that only have a trailing bracket", "asciidoctor/lists_test_does_not_recognize_callout_list_denoted_by_markers_that_only_have_a_trailing_bracket.adoc", listsTestDoesNotRecognizeCalloutListDenotedByMarkersThatOnlyHaveATrailingBracket},

	{"should not hang if obsolete callout list is found inside list item", "asciidoctor/lists_test_should_not_hang_if_obsolete_callout_list_is_found_inside_list_item.adoc", listsTestShouldNotHangIfObsoleteCalloutListIsFoundInsideListItem},

	{"should not hang if obsolete callout list is found inside dlist item", "asciidoctor/lists_test_should_not_hang_if_obsolete_callout_list_is_found_inside_dlist_item.adoc", listsTestShouldNotHangIfObsoleteCalloutListIsFoundInsideDlistItem},

	{"should recognize auto-numberd callout list inside list", "asciidoctor/lists_test_should_recognize_auto_numberd_callout_list_inside_list.adoc", listsTestShouldRecognizeAutoNumberdCalloutListInsideList},

	{"listing block with sequential callouts followed by adjacent callout list", "asciidoctor/lists_test_listing_block_with_sequential_callouts_followed_by_adjacent_callout_list.adoc", listsTestListingBlockWithSequentialCalloutsFollowedByAdjacentCalloutList},

	{"listing block with sequential callouts followed by non-adjacent callout list", "asciidoctor/lists_test_listing_block_with_sequential_callouts_followed_by_non_adjacent_callout_list.adoc", listsTestListingBlockWithSequentialCalloutsFollowedByNonAdjacentCalloutList},

	{"listing block with a callout that refers to two different lines", "asciidoctor/lists_test_listing_block_with_a_callout_that_refers_to_two_different_lines.adoc", listsTestListingBlockWithACalloutThatRefersToTwoDifferentLines},

	{"source block with non-sequential callouts followed by adjacent callout list", "asciidoctor/lists_test_source_block_with_non_sequential_callouts_followed_by_adjacent_callout_list.adoc", listsTestSourceBlockWithNonSequentialCalloutsFollowedByAdjacentCalloutList},

	{"two listing blocks can share the same callout list", "asciidoctor/lists_test_two_listing_blocks_can_share_the_same_callout_list.adoc", listsTestTwoListingBlocksCanShareTheSameCalloutList},

	{"two listing blocks each followed by an adjacent callout list", "asciidoctor/lists_test_two_listing_blocks_each_followed_by_an_adjacent_callout_list.adoc", listsTestTwoListingBlocksEachFollowedByAnAdjacentCalloutList},

	{"callout list retains block content", "asciidoctor/lists_test_callout_list_retains_block_content.adoc", listsTestCalloutListRetainsBlockContent},

	{"callout list retains block content when converted to DocBook", "asciidoctor/lists_test_callout_list_retains_block_content_when_converted_to_doc_book.adoc", listsTestCalloutListRetainsBlockContentWhenConvertedToDocBook},

	{"escaped callout should not be interpreted as a callout", "asciidoctor/lists_test_escaped_callout_should_not_be_interpreted_as_a_callout.adoc", listsTestEscapedCalloutShouldNotBeInterpretedAsACallout},

	{"should autonumber <.> callouts", "asciidoctor/lists_test_should_autonumber___callouts.adoc", listsTestShouldAutonumberCallouts},

	{"should not recognize callouts in middle of line", "asciidoctor/lists_test_should_not_recognize_callouts_in_middle_of_line.adoc", listsTestShouldNotRecognizeCalloutsInMiddleOfLine},

	{"should allow multiple callouts on the same line", "asciidoctor/lists_test_should_allow_multiple_callouts_on_the_same_line.adoc", listsTestShouldAllowMultipleCalloutsOnTheSameLine},

	{"should allow XML comment-style callouts", "asciidoctor/lists_test_should_allow_xml_comment_style_callouts.adoc", listsTestShouldAllowXmlCommentStyleCallouts},

	{"should not allow callouts with half an XML comment", "asciidoctor/lists_test_should_not_allow_callouts_with_half_an_xml_comment.adoc", listsTestShouldNotAllowCalloutsWithHalfAnXmlComment},

	{"should not recognize callouts in an indented description list paragraph", "asciidoctor/lists_test_should_not_recognize_callouts_in_an_indented_description_list_paragraph.adoc", listsTestShouldNotRecognizeCalloutsInAnIndentedDescriptionListParagraph},

	{"should not recognize callouts in an indented outline list paragraph", "asciidoctor/lists_test_should_not_recognize_callouts_in_an_indented_outline_list_paragraph.adoc", listsTestShouldNotRecognizeCalloutsInAnIndentedOutlineListParagraph},

	{"should warn if numbers in callout list are out of sequence", "asciidoctor/lists_test_should_warn_if_numbers_in_callout_list_are_out_of_sequence.adoc", listsTestShouldWarnIfNumbersInCalloutListAreOutOfSequence},

	{"should preserve line comment chars that precede callout number if icons is not set", "asciidoctor/lists_test_should_preserve_line_comment_chars_that_precede_callout_number_if_icons_is_not_set.adoc", listsTestShouldPreserveLineCommentCharsThatPrecedeCalloutNumberIfIconsIsNotSet},

	{"should remove line comment chars that precede callout number if icons is font", "asciidoctor/lists_test_should_remove_line_comment_chars_that_precede_callout_number_if_icons_is_font.adoc", listsTestShouldRemoveLineCommentCharsThatPrecedeCalloutNumberIfIconsIsFont},

	{"should allow line comment chars that precede callout number to be specified", "asciidoctor/lists_test_should_allow_line_comment_chars_that_precede_callout_number_to_be_specified.adoc", listsTestShouldAllowLineCommentCharsThatPrecedeCalloutNumberToBeSpecified},

	{"should allow line comment chars preceding callout number to be configurable when source-highlighter is coderay", "asciidoctor/lists_test_should_allow_line_comment_chars_preceding_callout_number_to_be_configurable_when_source_highlighter_is_coderay.adoc", listsTestShouldAllowLineCommentCharsPrecedingCalloutNumberToBeConfigurableWhenSourceHighlighterIsCoderay},

	{"should not eat whitespace before callout number if line-comment attribute is empty", "asciidoctor/lists_test_should_not_eat_whitespace_before_callout_number_if_line_comment_attribute_is_empty.adoc", listsTestShouldNotEatWhitespaceBeforeCalloutNumberIfLineCommentAttributeIsEmpty},

	{"literal block with callouts", "asciidoctor/lists_test_literal_block_with_callouts.adoc", listsTestLiteralBlockWithCallouts},

	{"callout list with icons enabled", "asciidoctor/lists_test_callout_list_with_icons_enabled.adoc", listsTestCalloutListWithIconsEnabled},

	{"callout list with font-based icons enabled", "asciidoctor/lists_test_callout_list_with_font_based_icons_enabled.adoc", listsTestCalloutListWithFontBasedIconsEnabled},

	{"should create checklist if at least one item has checkbox syntax", "asciidoctor/lists_test_should_create_checklist_if_at_least_one_item_has_checkbox_syntax.adoc", listsTestShouldCreateChecklistIfAtLeastOneItemHasCheckboxSyntax},

	{"should create checklist with font icons if at least one item has checkbox syntax and icons attribute is font", "asciidoctor/lists_test_should_create_checklist_with_font_icons_if_at_least_one_item_has_checkbox_syntax_and_icons_attribute_is_font.adoc", listsTestShouldCreateChecklistWithFontIconsIfAtLeastOneItemHasCheckboxSyntaxAndIconsAttributeIsFont},

	{"should create interactive checklist if interactive option is set even with icons attribute is font", "asciidoctor/lists_test_should_create_interactive_checklist_if_interactive_option_is_set_even_with_icons_attribute_is_font.adoc", listsTestShouldCreateInteractiveChecklistIfInteractiveOptionIsSetEvenWithIconsAttributeIsFont},

	{"content should return items in list", "asciidoctor/lists_test_content_should_return_items_in_list.adoc", listsTestContentShouldReturnItemsInList},

	{"list item should be the parent of block attached to a list item", "asciidoctor/lists_test_list_item_should_be_the_parent_of_block_attached_to_a_list_item.adoc", listsTestListItemShouldBeTheParentOfBlockAttachedToAListItem},

	{"outline? should return true for unordered list", "asciidoctor/lists_test_outline_should_return_true_for_unordered_list.adoc", listsTestOutlineShouldReturnTrueForUnorderedList},

	{"outline? should return true for ordered list", "asciidoctor/lists_test_outline_should_return_true_for_ordered_list.adoc", listsTestOutlineShouldReturnTrueForOrderedList},

	{"outline? should return false for description list", "asciidoctor/lists_test_outline_should_return_false_for_description_list.adoc", listsTestOutlineShouldReturnFalseForDescriptionList},

	{"simple? should return true for list item with nested outline list", "asciidoctor/lists_test_simple_should_return_true_for_list_item_with_nested_outline_list.adoc", listsTestSimpleShouldReturnTrueForListItemWithNestedOutlineList},

	{"simple? should return false for list item with block content", "asciidoctor/lists_test_simple_should_return_false_for_list_item_with_block_content.adoc", listsTestSimpleShouldReturnFalseForListItemWithBlockContent},

	{"should allow text of ListItem to be assigned", "asciidoctor/lists_test_should_allow_text_of_list_item_to_be_assigned.adoc", listsTestShouldAllowTextOfListItemToBeAssigned},

	{"id and role assigned to ulist item in model are transmitted to output", "asciidoctor/lists_test_id_and_role_assigned_to_ulist_item_in_model_are_transmitted_to_output.adoc", listsTestIdAndRoleAssignedToUlistItemInModelAreTransmittedToOutput},

	{"id and role assigned to olist item in model are transmitted to output", "asciidoctor/lists_test_id_and_role_assigned_to_olist_item_in_model_are_transmitted_to_output.adoc", listsTestIdAndRoleAssignedToOlistItemInModelAreTransmittedToOutput},

	{"should allow API control over substitutions applied to ListItem text", "asciidoctor/lists_test_should_allow_api_control_over_substitutions_applied_to_list_item_text.adoc", listsTestShouldAllowApiControlOverSubstitutionsAppliedToListItemText},

	{"should set lineno to line number in source where list starts", "asciidoctor/lists_test_should_set_lineno_to_line_number_in_source_where_list_starts.adoc", listsTestShouldSetLinenoToLineNumberInSourceWhereListStarts},
}

var listsTestDashElementsWithNoBlankLines = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "List",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "====",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Foo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "-",
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Boo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "-",
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Blech",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "-",
			Checklist:     0,
		},
	},
}

var listsTestDashElementsSeparatedByBlankLinesShouldMergeLists = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "List",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "====",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Foo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "-",
			Checklist:     0,
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Boo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "-",
			Checklist:     0,
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Blech",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "-",
			Checklist:     0,
		},
	},
}

var listsTestDashElementsWithInterspersedLineCommentsShouldBeSkippedAndNotBreakList = &asciidoc.Document{
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
				&asciidoc.UnorderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "Foo",
						},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "// line comment",
						},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "// another line comment",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        "-",
					Checklist:     0,
				},
				&asciidoc.UnorderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "Boo",
						},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "// line comment",
						},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "more text",
						},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "// another line comment",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        "-",
					Checklist:     0,
				},
				&asciidoc.UnorderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "Blech",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        "-",
					Checklist:     0,
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "List",
				},
			},
			Level: 1,
		},
	},
}

var listsTestDashElementsSeparatedByALineCommentOffsetByBlankLinesShouldNotMergeLists = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "List",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "====",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Foo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "-",
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Boo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "-",
			Checklist:     0,
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.SingleLineComment{
			Value: "",
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Blech",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "-",
			Checklist:     0,
		},
	},
}

var listsTestDashElementsSeparatedByABlockTitleOffsetByABlankLineShouldNotMergeLists = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "List",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "====",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Foo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "-",
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Boo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "-",
			Checklist:     0,
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Blech",
				},
			},
			AttributeList: asciidoc.AttributeList{
				&asciidoc.TitleAttribute{
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Also",
						},
					},
				},
			},
			Indent:    "",
			Marker:    "-",
			Checklist: 0,
		},
	},
}

var listsTestDashElementsSeparatedByAnAttributeEntryOffsetByABlankLineShouldNotMergeLists = &asciidoc.Document{
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
				&asciidoc.UnorderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "Foo",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        "-",
					Checklist:     0,
				},
				&asciidoc.UnorderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "Boo",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        "-",
					Checklist:     0,
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
				&asciidoc.UnorderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "Blech",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        "-",
					Checklist:     0,
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "List",
				},
			},
			Level: 1,
		},
	},
}

var listsTestANonIndentedWrappedLineIsFoldedIntoTextOfListItem = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "List",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "====",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Foo",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "wrapped content",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "-",
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Boo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "-",
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Blech",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "-",
			Checklist:     0,
		},
	},
}

var listsTestANonIndentedWrappedLineThatResemblesABlockTitleIsFoldedIntoTextOfListItem = &asciidoc.Document{
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
				&asciidoc.UnorderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "Foo",
						},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: ".wrapped content",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        "-",
					Checklist:     0,
				},
				&asciidoc.UnorderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "Boo",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        "-",
					Checklist:     0,
				},
				&asciidoc.UnorderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "Blech",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        "-",
					Checklist:     0,
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "List",
				},
			},
			Level: 1,
		},
	},
}

var listsTestANonIndentedWrappedLineThatResemblesAnAttributeEntryIsFoldedIntoTextOfListItem = &asciidoc.Document{
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
				&asciidoc.UnorderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "Foo",
						},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: ":foo: bar",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        "-",
					Checklist:     0,
				},
				&asciidoc.UnorderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "Boo",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        "-",
					Checklist:     0,
				},
				&asciidoc.UnorderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "Blech",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        "-",
					Checklist:     0,
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "List",
				},
			},
			Level: 1,
		},
	},
}

var listsTestAListItemWithANestedMarkerTerminatesNonIndentedParagraphForTextOfListItem = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Foo",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "Bar",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "-",
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Foo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
	},
}

var listsTestAListItemForADifferentListTerminatesNonIndentedParagraphForTextOfListItem = &asciidoc.Document{
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
				&asciidoc.UnorderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "Foo",
						},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "Bar",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        "-",
					Checklist:     0,
				},
				&asciidoc.OrderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "Foo",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        ".",
				},
				&asciidoc.EmptyLine{
					Text: "",
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Example 1",
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
				&asciidoc.UnorderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "Item",
						},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "text",
						},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "term:: def",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        "*",
					Checklist:     0,
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Example 2",
				},
			},
			Level: 1,
		},
	},
}

var listsTestAnIndentedWrappedLineIsUnindentedAndFoldedIntoTextOfListItem = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "List",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "====",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Foo",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "  wrapped content",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "-",
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Boo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "-",
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Blech",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "-",
			Checklist:     0,
		},
	},
}

var listsTestWrappedListItemWithHangingIndentFollowedByNonIndentedLine = &asciidoc.Document{
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
				&asciidoc.UnorderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "list item 1",
						},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "  // not line comment",
						},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "second wrapped line",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        "-",
					Checklist:     0,
				},
				&asciidoc.UnorderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "list item 2",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        "-",
					Checklist:     0,
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var listsTestAListItemWithANestedMarkerTerminatesIndentedParagraphForTextOfListItem = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Foo",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "  Bar",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "-",
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Foo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
	},
}

var listsTestAListItemForADifferentListTerminatesIndentedParagraphForTextOfListItem = &asciidoc.Document{
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
				&asciidoc.UnorderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "Foo",
						},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "  Bar",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        "-",
					Checklist:     0,
				},
				&asciidoc.OrderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "Foo",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        ".",
				},
				&asciidoc.EmptyLine{
					Text: "",
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Example 1",
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
				&asciidoc.UnorderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "Item",
						},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "  text",
						},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "term:: def",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        "*",
					Checklist:     0,
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Example 2",
				},
			},
			Level: 1,
		},
	},
}

var listsTestALiteralParagraphOffsetByBlankLinesInListContentIsAppendedAsALiteralBlock = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "List",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "====",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Foo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "-",
			Checklist:     0,
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "  literal",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Boo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "-",
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Blech",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "-",
			Checklist:     0,
		},
	},
}

var listsTestShouldEscapeSpecialCharactersInAllLiteralParagraphsAttachedToListItem = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "first item",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "  <code>text</code>",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "  more <code>text</code>",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "second item",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
	},
}

var listsTestALiteralParagraphOffsetByABlankLineInListContentFollowedByLineWithContinuationIsAppendedAsTwoBlocks = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "List",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "====",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Foo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "-",
			Checklist:     0,
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "  literal",
		},
		&asciidoc.NewLine{},
		&asciidoc.LineBreak{},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "para",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Boo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "-",
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Blech",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "-",
			Checklist:     0,
		},
	},
}

var listsTestAnAdmonitionParagraphAttachedByALineContinuationToAListItemWithWrappedTextShouldProduceAdmonition = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "first-line text",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "  wrapped text",
				},
				&asciidoc.NewLine{},
				&asciidoc.LineBreak{},
				&asciidoc.NewLine{},
				&asciidoc.Paragraph{
					AttributeList: nil,
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "This is a note.",
						},
					},
					Admonition: 1,
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "-",
			Checklist:     0,
		},
	},
}

var listsTestParagraphLikeBlocksAttachedToAnAncestorListItemByAListContinuationShouldProduceBlocks = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "parent",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "child",
				},
			},
			AttributeList: nil,
			Indent:        " ",
			Marker:        "**",
			Checklist:     0,
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.LineBreak{},
		&asciidoc.NewLine{},
		&asciidoc.Paragraph{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "This is a note.",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 1,
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "another parent",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "another child",
				},
			},
			AttributeList: nil,
			Indent:        " ",
			Marker:        "**",
			Checklist:     0,
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.ListContinuation{
			ChildElement: &asciidoc.ThematicBreak{
				AttributeList: nil,
			},
		},
	},
}

var listsTestShouldNotInheritBlockAttributesFromPreviousBlockWhenBlockIsAttachedUsingAListContinuation = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "complex list item",
				},
				&asciidoc.NewLine{},
				&asciidoc.LineBreak{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "[source,xml]",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		&asciidoc.Listing{
			AttributeList: nil,
			Delimiter: asciidoc.Delimiter{
				Type:   5,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"<name>value</name> <!--1-->",
			},
		},
		&asciidoc.String{
			Value: "<1> a configuration value",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestShouldContinueToParseBlocksAttachedByAListContinuationAfterBlockIsDropped = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "item",
				},
				&asciidoc.NewLine{},
				&asciidoc.LineBreak{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "paragraph",
				},
				&asciidoc.NewLine{},
				&asciidoc.LineBreak{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "[comment]",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "comment",
				},
				&asciidoc.NewLine{},
				&asciidoc.LineBreak{},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		&asciidoc.ExampleBlock{
			Delimiter: asciidoc.Delimiter{
				Type:   3,
				Length: 4,
			},
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "example",
				},
				&asciidoc.NewLine{},
			},
		},
		&asciidoc.ThematicBreak{
			AttributeList: nil,
		},
	},
}

var listsTestAppendsLineAsParagraphIfAttachedByContinuationFollowingLineComment = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "list item 1",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "// line comment",
				},
				&asciidoc.NewLine{},
				&asciidoc.LineBreak{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "paragraph in list item 1",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "-",
			Checklist:     0,
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "list item 2",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "-",
			Checklist:     0,
		},
	},
}

var listsTestALiteralParagraphWithALineThatAppearsAsAListItemThatIsFollowedByAContinuationShouldCreateTwoBlocks = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Foo",
				},
				&asciidoc.NewLine{},
				&asciidoc.LineBreak{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "  literal",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "still literal",
				},
				&asciidoc.NewLine{},
				&asciidoc.LineBreak{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "para",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Bar",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
	},
}

var listsTestConsecutiveLiteralParagraphOffsetByBlankLinesInListContentAreAppendedAsALiteralBlocks = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "List",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "====",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Foo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "-",
			Checklist:     0,
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "  literal",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "  more",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "  literal",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Boo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "-",
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Blech",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "-",
			Checklist:     0,
		},
	},
}

var listsTestALiteralParagraphWithoutATrailingBlankLineConsumesFollowingListItems = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "List",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "====",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Foo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "-",
			Checklist:     0,
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "  literal",
		},
		&asciidoc.NewLine{},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Boo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "-",
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Blech",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "-",
			Checklist:     0,
		},
	},
}

var listsTestAsteriskElementsWithNoBlankLines = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "List",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "====",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Foo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Boo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Blech",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
	},
}

var listsTestAsteriskElementsSeparatedByBlankLinesShouldMergeLists = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "List",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "====",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Foo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Boo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Blech",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
	},
}

var listsTestAsteriskElementsWithInterspersedLineCommentsShouldBeSkippedAndNotBreakList = &asciidoc.Document{
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
				&asciidoc.UnorderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "Foo",
						},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "// line comment",
						},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "// another line comment",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        "*",
					Checklist:     0,
				},
				&asciidoc.UnorderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "Boo",
						},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "// line comment",
						},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "more text",
						},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "// another line comment",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        "*",
					Checklist:     0,
				},
				&asciidoc.UnorderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "Blech",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        "*",
					Checklist:     0,
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "List",
				},
			},
			Level: 1,
		},
	},
}

var listsTestAsteriskElementsSeparatedByALineCommentOffsetByBlankLinesShouldNotMergeLists = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "List",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "====",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Foo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Boo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.SingleLineComment{
			Value: "",
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Blech",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
	},
}

var listsTestAsteriskElementsSeparatedByABlockTitleOffsetByABlankLineShouldNotMergeLists = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "List",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "====",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Foo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Boo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Blech",
				},
			},
			AttributeList: asciidoc.AttributeList{
				&asciidoc.TitleAttribute{
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Also",
						},
					},
				},
			},
			Indent:    "",
			Marker:    "*",
			Checklist: 0,
		},
	},
}

var listsTestAsteriskElementsSeparatedByAnAttributeEntryOffsetByABlankLineShouldNotMergeLists = &asciidoc.Document{
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
				&asciidoc.UnorderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "Foo",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        "*",
					Checklist:     0,
				},
				&asciidoc.UnorderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "Boo",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        "*",
					Checklist:     0,
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
				&asciidoc.UnorderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "Blech",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        "*",
					Checklist:     0,
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "List",
				},
			},
			Level: 1,
		},
	},
}

var listsTestListShouldTerminateBeforeNextLowerSectionHeading = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "List",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "====",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "first",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "item",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "second",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "item",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Elements:      nil,
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Section",
				},
			},
			Level: 1,
		},
	},
}

var listsTestListShouldTerminateBeforeNextLowerSectionHeadingWithImplicitId = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "List",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "====",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "first",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "item",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "second",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "item",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.AnchorAttribute{
					ID: &asciidoc.String{
						Value: "sec",
					},
					Label: nil,
				},
			},
			Elements: nil,
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Section",
				},
			},
			Level: 1,
		},
	},
}

var listsTestShouldNotFindSectionTitleImmediatelyBelowLastListItem = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "first",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "second",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "== Not a section",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
	},
}

var listsTestQuotedText = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "List",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "====",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "I am ",
				},
				&asciidoc.Bold{
					AttributeList: nil,
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "strong",
						},
					},
				},
				&asciidoc.String{
					Value: ".",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "-",
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "I am ",
				},
				&asciidoc.Italic{
					AttributeList: nil,
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "stressed",
						},
					},
				},
				&asciidoc.String{
					Value: ".",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "-",
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "I am ",
				},
				&asciidoc.Monospace{
					AttributeList: nil,
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "flexible",
						},
					},
				},
				&asciidoc.String{
					Value: ".",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "-",
			Checklist:     0,
		},
	},
}

var listsTestAttributeSubstitutions = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "List",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "====",
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
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "side a ",
				},
				&asciidoc.CharacterReplacementReference{
					Value: "vbar",
				},
				&asciidoc.String{
					Value: " side b",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "-",
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Take me to a ",
				},
				&asciidoc.UserAttributeReference{
					Value: "foo",
				},
				&asciidoc.String{
					Value: ".",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "-",
			Checklist:     0,
		},
	},
}

var listsTestLeadingDotIsTreatedAsTextNotBlockTitle = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: ".first",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: ".second",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: ".third",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
	},
}

var listsTestWordEndingSentenceOnContinuingLineNotTreatedAsAListItem = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "This is the story about",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "   AsciiDoc. It begins here.",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "A.",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "And it ends here.",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "B.",
		},
	},
}

var listsTestShouldDiscoverAnchorAtStartOfUnorderedListItemTextAndRegisterItAsAReference = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "The highest peak in the Front Range is ",
		},
		&asciidoc.CrossReference{
			AttributeList: nil,
			Elements:      nil,
			ID:            "grays-peak",
			Format:        0,
		},
		&asciidoc.String{
			Value: ", which tops ",
		},
		&asciidoc.CrossReference{
			AttributeList: nil,
			Elements:      nil,
			ID:            "mount-evans",
			Format:        0,
		},
		&asciidoc.String{
			Value: " by just a few feet.",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.Anchor{
					ID: "mount-evans",
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "Mount Evans",
						},
					},
				},
				&asciidoc.String{
					Value: "At 14,271 feet, Mount Evans is the highest summit of the Chicago Peaks in the Front Range of the Rocky Mountains.",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.Anchor{
					ID: "grays-peak",
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "Grays Peak",
						},
					},
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "Grays Peak rises to 14,278 feet, making it the highest summit in the Front Range of the Rocky Mountains.",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Longs Peak is a 14,259-foot high, prominent mountain summit in the northern Front Range of the Rocky Mountains.",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Pikes Peak is the highest summit of the southern Front Range of the Rocky Mountains at 14,115 feet.",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
	},
}

var listsTestShouldDiscoverAnchorAtStartOfOrderedListItemTextAndRegisterItAsAReference = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "This is a cross-reference to ",
		},
		&asciidoc.CrossReference{
			AttributeList: nil,
			Elements:      nil,
			ID:            "step-2",
			Format:        0,
		},
		&asciidoc.String{
			Value: ".",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "This is a cross-reference to ",
		},
		&asciidoc.CrossReference{
			AttributeList: nil,
			Elements:      nil,
			ID:            "step-4",
			Format:        0,
		},
		&asciidoc.String{
			Value: ".",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Ordered list, item 1, without anchor",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.Anchor{
					ID: "step-2",
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "Step 2",
						},
					},
				},
				&asciidoc.String{
					Value: "Ordered list, item 2, with anchor",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Ordered list, item 3, without anchor",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.Anchor{
					ID: "step-4",
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "Step 4",
						},
					},
				},
				&asciidoc.String{
					Value: "Ordered list, item 4, with anchor",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
	},
}

var listsTestShouldDiscoverAnchorAtStartOfCalloutListItemTextAndRegisterItAsAReference = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "This is a cross-reference to ",
		},
		&asciidoc.CrossReference{
			AttributeList: nil,
			Elements:      nil,
			ID:            "url-mapping",
			Format:        0,
		},
		&asciidoc.String{
			Value: ".",
		},
		&asciidoc.NewLine{},
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
				"require 'sinatra' <1>",
				"",
				"get '/hi' do <2> <3>",
				"  \"Hello World!\"",
				"end",
			},
		},
		&asciidoc.String{
			Value: "<1> Library import",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "<2> ",
		},
		&asciidoc.Anchor{
			ID: "url-mapping",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "url mapping",
				},
			},
		},
		&asciidoc.String{
			Value: "URL mapping",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "<3> Response block",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestAsteriskElementMixedWithDashElementsShouldBeNested = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "List",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "====",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Foo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "-",
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Boo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Blech",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "-",
			Checklist:     0,
		},
	},
}

var listsTestDashElementMixedWithAsterisksElementsShouldBeNested = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "List",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "====",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Foo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Boo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "-",
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Blech",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
	},
}

var listsTestLinesPrefixedWithAlternatingListMarkersSeparatedByBlankLinesShouldBeNested = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "List",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "====",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Foo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "-",
			Checklist:     0,
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Boo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Blech",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "-",
			Checklist:     0,
		},
	},
}

var listsTestNestedElements2WithAsterisks = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "List",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "====",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Foo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Boo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "**",
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Blech",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
	},
}

var listsTestNestedElements3WithAsterisks = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "List",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "====",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Foo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Boo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "**",
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Snoo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "***",
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Blech",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
	},
}

var listsTestNestedElements4WithAsterisks = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "List",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "====",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Foo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Boo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "**",
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Snoo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "***",
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Froo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "****",
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Blech",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
	},
}

var listsTestNestedElements5WithAsterisks = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "List",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "====",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Foo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Boo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "**",
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Snoo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "***",
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Froo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "****",
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Groo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*****",
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Blech",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
	},
}

var listsTestLevelOfUnorderedListShouldMatchSectionLevel = &asciidoc.Document{
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
				&asciidoc.UnorderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "item 1.1",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        "*",
					Checklist:     0,
				},
				&asciidoc.UnorderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "item 2.1",
						},
					},
					AttributeList: nil,
					Indent:        " ",
					Marker:        "**",
					Checklist:     0,
				},
				&asciidoc.UnorderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "item 3.1",
						},
					},
					AttributeList: nil,
					Indent:        "  ",
					Marker:        "***",
					Checklist:     0,
				},
				&asciidoc.UnorderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "item 2.2",
						},
					},
					AttributeList: nil,
					Indent:        " ",
					Marker:        "**",
					Checklist:     0,
				},
				&asciidoc.UnorderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "item 1.2",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        "*",
					Checklist:     0,
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
						&asciidoc.UnorderedListItem{
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "item 1.1",
								},
							},
							AttributeList: nil,
							Indent:        "",
							Marker:        "*",
							Checklist:     0,
						},
					},
					Title: asciidoc.Elements{
						&asciidoc.String{
							Value: "Nested Section",
						},
					},
					Level: 2,
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Parent Section",
				},
			},
			Level: 1,
		},
	},
}

var listsTestDoesNotRecognizeListsWithRepeatingUnicodeBullets = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "List",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "====",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Foo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Boo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "..",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Blech",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
	},
}

var listsTestNestedOrderedElements3 = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "List",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "====",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Foo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Boo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "..",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Snoo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "...",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Blech",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
	},
}

var listsTestLevelOfOrderedListShouldMatchSectionLevel = &asciidoc.Document{
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
				&asciidoc.OrderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "item 1.1",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        ".",
				},
				&asciidoc.OrderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "item 2.1",
						},
					},
					AttributeList: nil,
					Indent:        " ",
					Marker:        "..",
				},
				&asciidoc.OrderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "item 3.1",
						},
					},
					AttributeList: nil,
					Indent:        "  ",
					Marker:        "...",
				},
				&asciidoc.OrderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "item 2.2",
						},
					},
					AttributeList: nil,
					Indent:        " ",
					Marker:        "..",
				},
				&asciidoc.OrderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "item 1.2",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        ".",
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
						&asciidoc.OrderedListItem{
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "item 1.1",
								},
							},
							AttributeList: nil,
							Indent:        "",
							Marker:        ".",
						},
					},
					Title: asciidoc.Elements{
						&asciidoc.String{
							Value: "Nested Section",
						},
					},
					Level: 2,
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Parent Section",
				},
			},
			Level: 1,
		},
	},
}

var listsTestNestedUnorderedInsideOrderedElements = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "List",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "====",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Foo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Boo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Blech",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
	},
}

var listsTestNestedOrderedInsideUnorderedElements = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "List",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "====",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Foo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Boo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Blech",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
	},
}

var listsTestThreeLevelsOfAlternatingUnorderedAndOrderedElements = &asciidoc.Document{
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
				&asciidoc.UnorderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "bullet 1",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        "*",
					Checklist:     0,
				},
				&asciidoc.OrderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "numbered 1.1",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        ".",
				},
				&asciidoc.UnorderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "bullet 1.1.1",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        "**",
					Checklist:     0,
				},
				&asciidoc.UnorderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "bullet 2",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        "*",
					Checklist:     0,
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var listsTestLinesWithAlternatingMarkersOfUnorderedAndOrderedListTypesSeparatedByBlankLinesShouldBeNested = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "List",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "====",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Foo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Boo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Blech",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
	},
}

var listsTestListItemWithLiteralContentShouldNotConsumeNestedListOfDifferentType = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "List",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "====",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "bullet",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "-",
			Checklist:     0,
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "  literal",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "  but not",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "  hungry",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "numbered",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
	},
}

var listsTestNestedListItemDoesNotEatTheTitleOfTheFollowingDetachedBlock = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "List",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "====",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "bullet",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "-",
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "nested bullet 1",
				},
			},
			AttributeList: nil,
			Indent:        "  ",
			Marker:        "*",
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "nested bullet 2",
				},
			},
			AttributeList: nil,
			Indent:        "  ",
			Marker:        "*",
			Checklist:     0,
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.LiteralBlock{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.TitleAttribute{
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Title",
						},
					},
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   6,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"literal",
			},
		},
	},
}

var listsTestLinesWithAlternatingMarkersOfBulletedAndDescriptionListTypesSeparatedByBlankLinesShouldBeNested = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "List",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "====",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Foo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "term1:: def1",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Blech",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
	},
}

var listsTestNestedOrderedWithAttributeInsideUnorderedElements = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "Blah",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "====",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Foo",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "[start=2]",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Boo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Blech",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
	},
}

var listsTestAdjacentListContinuationLineAttachesFollowingParagraph = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "Lists",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "=====",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Item one, paragraph one",
				},
				&asciidoc.NewLine{},
				&asciidoc.LineBreak{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "Item one, paragraph two",
				},
				&asciidoc.NewLine{},
				&asciidoc.LineBreak{},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Item two",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
	},
}

var listsTestAdjacentListContinuationLineAttachesFollowingBlock = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "Lists",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "=====",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Item one, paragraph one",
				},
				&asciidoc.NewLine{},
				&asciidoc.LineBreak{},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		&asciidoc.LiteralBlock{
			AttributeList: nil,
			Delimiter: asciidoc.Delimiter{
				Type:   6,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"Item one, literal block",
			},
		},
		&asciidoc.ListContinuation{
			ChildElement: &asciidoc.UnorderedListItem{
				Elements: asciidoc.Elements{
					&asciidoc.String{
						Value: "Item two",
					},
				},
				AttributeList: nil,
				Indent:        "",
				Marker:        "*",
				Checklist:     0,
			},
		},
	},
}

var listsTestAdjacentListContinuationLineAttachesFollowingBlockWithBlockAttributes = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "Lists",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "=====",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Item one, paragraph one",
				},
				&asciidoc.NewLine{},
				&asciidoc.LineBreak{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: ":foo: bar",
				},
				&asciidoc.NewLine{},
				&asciidoc.Anchor{
					ID:       "beck",
					Elements: nil,
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: ".Read the following aloud to yourself",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "[source, ruby]",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		&asciidoc.Listing{
			AttributeList: nil,
			Delimiter: asciidoc.Delimiter{
				Type:   5,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"5.times { print \"Odelay!\" }",
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Item two",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
	},
}

var listsTestTrailingBlockAttributeLineAttachedByContinuationShouldNotCreateBlock = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "Lists",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "=====",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Item one, paragraph one",
				},
				&asciidoc.NewLine{},
				&asciidoc.LineBreak{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "[source]",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Item two",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
	},
}

var listsTestTrailingBlockTitleLineAttachedByContinuationShouldNotCreateBlock = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "Lists",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "=====",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Item one, paragraph one",
				},
				&asciidoc.NewLine{},
				&asciidoc.LineBreak{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: ".Disappears into the ether",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Item two",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
	},
}

var listsTestConsecutiveBlocksInListContinuationAttachToListItem = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "Lists",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "=====",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Item one, paragraph one",
				},
				&asciidoc.NewLine{},
				&asciidoc.LineBreak{},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		&asciidoc.LiteralBlock{
			AttributeList: nil,
			Delimiter: asciidoc.Delimiter{
				Type:   6,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"Item one, literal block",
			},
		},
		&asciidoc.LineBreak{},
		&asciidoc.NewLine{},
		&asciidoc.QuoteBlock{
			Delimiter: asciidoc.Delimiter{
				Type:   11,
				Length: 4,
			},
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Item one, quote block",
				},
				&asciidoc.NewLine{},
			},
		},
		&asciidoc.ListContinuation{
			ChildElement: &asciidoc.UnorderedListItem{
				Elements: asciidoc.Elements{
					&asciidoc.String{
						Value: "Item two",
					},
				},
				AttributeList: nil,
				Indent:        "",
				Marker:        "*",
				Checklist:     0,
			},
		},
	},
}

var listsTestListItemWithHangingIndentFollowedByBlockAttachedByListContinuation = &asciidoc.Document{
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
				&asciidoc.OrderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "list item 1",
						},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "  continued",
						},
						&asciidoc.NewLine{},
						&asciidoc.LineBreak{},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "--",
						},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "open block in list item 1",
						},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "--",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        ".",
				},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.OrderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "list item 2",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        ".",
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var listsTestListItemParagraphInListItemAndNestedListItem = &asciidoc.Document{
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
				&asciidoc.OrderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "list item 1",
						},
						&asciidoc.NewLine{},
						&asciidoc.LineBreak{},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "list item 1 paragraph",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        ".",
				},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.UnorderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "nested list item",
						},
						&asciidoc.NewLine{},
						&asciidoc.LineBreak{},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "nested list item paragraph",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        "*",
					Checklist:     0,
				},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.OrderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "list item 2",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        ".",
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var listsTestTrailingListContinuationsShouldAttachToListItemsAtRespectiveLevels = &asciidoc.Document{
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
				&asciidoc.OrderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "list item 1",
						},
						&asciidoc.NewLine{},
						&asciidoc.LineBreak{},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        ".",
				},
				&asciidoc.UnorderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "nested list item 1",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        "*",
					Checklist:     0,
				},
				&asciidoc.UnorderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "nested list item 2",
						},
						&asciidoc.NewLine{},
						&asciidoc.LineBreak{},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "paragraph for nested list item 2",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        "*",
					Checklist:     0,
				},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.LineBreak{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "paragraph for list item 1",
				},
				&asciidoc.NewLine{},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.OrderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "list item 2",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        ".",
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var listsTestTrailingListContinuationsShouldAttachToListItemsOfDifferentTypesAtRespectiveLevels = &asciidoc.Document{
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
				&asciidoc.UnorderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "bullet 1",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        "*",
					Checklist:     0,
				},
				&asciidoc.OrderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "numbered 1.1",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        ".",
				},
				&asciidoc.UnorderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "bullet 1.1.1",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        "**",
					Checklist:     0,
				},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.LineBreak{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "numbered 1.1 paragraph",
				},
				&asciidoc.NewLine{},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.LineBreak{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "bullet 1 paragraph",
				},
				&asciidoc.NewLine{},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.UnorderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "bullet 2",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        "*",
					Checklist:     0,
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var listsTestRepeatedListContinuationsShouldAttachToListItemsAtRespectiveLevels = &asciidoc.Document{
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
				&asciidoc.OrderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "list item 1",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        ".",
				},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.UnorderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "nested list item 1",
						},
						&asciidoc.NewLine{},
						&asciidoc.LineBreak{},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "--",
						},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "open block for nested list item 1",
						},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "--",
						},
						&asciidoc.NewLine{},
						&asciidoc.LineBreak{},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        "*",
					Checklist:     0,
				},
				&asciidoc.UnorderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "nested list item 2",
						},
						&asciidoc.NewLine{},
						&asciidoc.LineBreak{},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "paragraph for nested list item 2",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        "*",
					Checklist:     0,
				},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.LineBreak{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "paragraph for list item 1",
				},
				&asciidoc.NewLine{},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.OrderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "list item 2",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        ".",
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var listsTestRepeatedListContinuationsAttachedDirectlyToListItemShouldAttachToListItemsAtRespectiveLevels = &asciidoc.Document{
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
				&asciidoc.OrderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "list item 1",
						},
						&asciidoc.NewLine{},
						&asciidoc.LineBreak{},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        ".",
				},
				&asciidoc.UnorderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "nested list item 1",
						},
						&asciidoc.NewLine{},
						&asciidoc.LineBreak{},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "--",
						},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "open block for nested list item 1",
						},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "--",
						},
						&asciidoc.NewLine{},
						&asciidoc.LineBreak{},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        "*",
					Checklist:     0,
				},
				&asciidoc.UnorderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "nested list item 2",
						},
						&asciidoc.NewLine{},
						&asciidoc.LineBreak{},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "paragraph for nested list item 2",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        "*",
					Checklist:     0,
				},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.LineBreak{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "paragraph for list item 1",
				},
				&asciidoc.NewLine{},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.OrderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "list item 2",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        ".",
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var listsTestRepeatedListContinuationsShouldAttachToListItemsAtRespectiveLevelsIgnoringBlankLines = &asciidoc.Document{
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
				&asciidoc.OrderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "list item 1",
						},
						&asciidoc.NewLine{},
						&asciidoc.LineBreak{},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        ".",
				},
				&asciidoc.UnorderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "nested list item 1",
						},
						&asciidoc.NewLine{},
						&asciidoc.LineBreak{},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "--",
						},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "open block for nested list item 1",
						},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "--",
						},
						&asciidoc.NewLine{},
						&asciidoc.LineBreak{},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        "*",
					Checklist:     0,
				},
				&asciidoc.UnorderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "nested list item 2",
						},
						&asciidoc.NewLine{},
						&asciidoc.LineBreak{},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "paragraph for nested list item 2",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        "*",
					Checklist:     0,
				},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.LineBreak{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "paragraph for list item 1",
				},
				&asciidoc.NewLine{},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.OrderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "list item 2",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        ".",
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var listsTestTrailingListContinuationsShouldIgnorePrecedingBlankLines = &asciidoc.Document{
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
				&asciidoc.UnorderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "bullet 1",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        "*",
					Checklist:     0,
				},
				&asciidoc.UnorderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "bullet 1.1",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        "**",
					Checklist:     0,
				},
				&asciidoc.UnorderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "bullet 1.1.1",
						},
						&asciidoc.NewLine{},
						&asciidoc.LineBreak{},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "--",
						},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "open block",
						},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "--",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        "***",
					Checklist:     0,
				},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.LineBreak{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "bullet 1.1 paragraph",
				},
				&asciidoc.NewLine{},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.LineBreak{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "bullet 1 paragraph",
				},
				&asciidoc.NewLine{},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.UnorderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "bullet 2",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        "*",
					Checklist:     0,
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var listsTestIndentedOutlineListItemWithDifferentMarkerOffsetByABlankLineShouldBeRecognizedAsANestedList = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "item 1",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "item 1.1",
				},
				&asciidoc.NewLine{},
				&asciidoc.LineBreak{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "attached paragraph",
				},
			},
			AttributeList: nil,
			Indent:        "  ",
			Marker:        ".",
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "item 1.2",
				},
				&asciidoc.NewLine{},
				&asciidoc.LineBreak{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "attached paragraph",
				},
			},
			AttributeList: nil,
			Indent:        "  ",
			Marker:        ".",
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "item 2",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
	},
}

var listsTestIndentedDescriptionListItemInsideOutlineListItemOffsetByABlankLineShouldBeRecognizedAsANestedList = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "item 1",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "  term a:: description a",
		},
		&asciidoc.NewLine{},
		&asciidoc.LineBreak{},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "attached paragraph",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "  term b:: description b",
		},
		&asciidoc.NewLine{},
		&asciidoc.LineBreak{},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "attached paragraph",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "item 2",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
	},
}

var listsTestConsecutiveListContinuationLinesAreFolded = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "Lists",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "=====",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Item one, paragraph one",
				},
				&asciidoc.NewLine{},
				&asciidoc.LineBreak{},
				&asciidoc.NewLine{},
				&asciidoc.LineBreak{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "Item one, paragraph two",
				},
				&asciidoc.NewLine{},
				&asciidoc.LineBreak{},
				&asciidoc.NewLine{},
				&asciidoc.LineBreak{},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Item two",
				},
				&asciidoc.NewLine{},
				&asciidoc.LineBreak{},
				&asciidoc.NewLine{},
				&asciidoc.LineBreak{},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
	},
}

var listsTestShouldWarnIfUnterminatedBlockIsDetectedInListItem = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "item",
				},
				&asciidoc.NewLine{},
				&asciidoc.LineBreak{},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		&asciidoc.String{
			Value: "====",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "example",
		},
		&asciidoc.NewLine{},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "swallowed item",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
	},
}

var listsTestDotElementsWithNoBlankLines = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "List",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "====",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Foo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Boo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Blech",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
	},
}

var listsTestShouldRepresentExplicitRoleAttributeAsStyleClass = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Once",
				},
			},
			AttributeList: asciidoc.AttributeList{
				&asciidoc.NamedAttribute{
					Name: "role",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "dry",
						},
					},
					Quote: 2,
				},
			},
			Indent: "",
			Marker: ".",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Again",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Refactor!",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
	},
}

var listsTestShouldBaseListStyleOnMarkerLengthRatherThanListDepth = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "parent",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "...",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "child",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "..",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "grandchild",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
	},
}

var listsTestShouldAllowListStyleToBeSpecifiedExplicitlyWhenUsingMarkersWithImplicitStyle = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "1",
				},
			},
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "loweralpha",
							},
						},
					},
					ID:      nil,
					Roles:   nil,
					Options: nil,
				},
			},
			Indent: "",
			Marker: "i)",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "2",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "ii)",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "3",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "iii)",
		},
	},
}

var listsTestShouldRepresentCustomNumberingAndExplicitRoleAttributeAsStyleClasses = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Once",
				},
			},
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "loweralpha",
							},
						},
					},
					ID:      nil,
					Roles:   nil,
					Options: nil,
				},
				&asciidoc.NamedAttribute{
					Name: "role",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "dry",
						},
					},
					Quote: 2,
				},
			},
			Indent: "",
			Marker: ".",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Again",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Refactor!",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
	},
}

var listsTestShouldSetReversedAttributeOnListIfReversedOptionIsSet = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "three",
				},
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
									Value: "reversed",
								},
							},
						},
					},
				},
				&asciidoc.NamedAttribute{
					Name: "start",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "3",
						},
					},
					Quote: 0,
				},
			},
			Indent: "",
			Marker: ".",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "two",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "one",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "blast off!",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
	},
}

var listsTestShouldRepresentImplicitRoleAttributeAsStyleClass = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Once",
				},
			},
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: nil,
					ID:    nil,
					Roles: []*asciidoc.ShorthandRole{
						&asciidoc.ShorthandRole{
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "dry",
								},
							},
						},
					},
					Options: nil,
				},
			},
			Indent: "",
			Marker: ".",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Again",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Refactor!",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
	},
}

var listsTestShouldRepresentCustomNumberingAndImplicitRoleAttributeAsStyleClasses = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Once",
				},
			},
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "loweralpha",
							},
						},
					},
					ID: nil,
					Roles: []*asciidoc.ShorthandRole{
						&asciidoc.ShorthandRole{
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "dry",
								},
							},
						},
					},
					Options: nil,
				},
			},
			Indent: "",
			Marker: ".",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Again",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Refactor!",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
	},
}

var listsTestDotElementsSeparatedByBlankLinesShouldMergeLists = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "List",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "====",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Foo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Boo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Blech",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
	},
}

var listsTestDotElementsWithInterspersedLineCommentsShouldBeSkippedAndNotBreakList = &asciidoc.Document{
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
				&asciidoc.OrderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "Foo",
						},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "// line comment",
						},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "// another line comment",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        ".",
				},
				&asciidoc.OrderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "Boo",
						},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "// line comment",
						},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "more text",
						},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "// another line comment",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        ".",
				},
				&asciidoc.OrderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "Blech",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        ".",
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "List",
				},
			},
			Level: 1,
		},
	},
}

var listsTestDotElementsSeparatedByLineCommentOffsetByBlankLinesShouldNotMergeLists = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "List",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "====",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Foo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Boo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.SingleLineComment{
			Value: "",
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Blech",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
	},
}

var listsTestDotElementsSeparatedByABlockTitleOffsetByABlankLineShouldNotMergeLists = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "List",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "====",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Foo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Boo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Blech",
				},
			},
			AttributeList: asciidoc.AttributeList{
				&asciidoc.TitleAttribute{
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Also",
						},
					},
				},
			},
			Indent: "",
			Marker: ".",
		},
	},
}

var listsTestDotElementsSeparatedByAnAttributeEntryOffsetByABlankLineShouldNotMergeLists = &asciidoc.Document{
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
				&asciidoc.OrderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "Foo",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        ".",
				},
				&asciidoc.OrderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "Boo",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        ".",
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
				&asciidoc.OrderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "Blech",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        ".",
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "List",
				},
			},
			Level: 1,
		},
	},
}

var listsTestShouldUseStartNumberInDocbook5Backend = &asciidoc.Document{
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
				&asciidoc.OrderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "item 7",
						},
					},
					AttributeList: asciidoc.AttributeList{
						&asciidoc.NamedAttribute{
							Name: "start",
							Val: asciidoc.Elements{
								&asciidoc.String{
									Value: "7",
								},
							},
							Quote: 0,
						},
					},
					Indent: "",
					Marker: ".",
				},
				&asciidoc.OrderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "item 8",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        ".",
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "List",
				},
			},
			Level: 1,
		},
	},
}

var listsTestShouldWarnIfExplicitUppercaseRomanNumeralsInListAreOutOfSequence = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "one",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "I)",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "three",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "III)",
		},
	},
}

var listsTestShouldWarnIfExplicitLowercaseRomanNumeralsInListAreOutOfSequence = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "one",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "i)",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "three",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "iii)",
		},
	},
}

var listsTestShouldNotParseABareDlistDelimiterAsADlist = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "term1:: def1",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "term2:: def2",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestShouldParseSiblingItemsUsingSameRules = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "term1;; ;; def1",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "term2;; ;; def2",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestShouldAllowTermToEndWithASemicolonWhenUsingDoubleSemicolonDelimiter = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "term;;; def",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestSingleLineIndentedAdjacentElements = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "term1:: def1",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: " term2:: def2",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestSingleLineElementsSeparatedByBlankLineShouldCreateASingleList = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "term1:: def1",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "term2:: def2",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestALineCommentBetweenElementsShouldDivideThemIntoSeparateLists = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "term1:: def1",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.SingleLineComment{
			Value: "",
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "term2:: def2",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestARulerBetweenElementsShouldDivideThemIntoSeparateLists = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "term1:: def1",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.ThematicBreak{
			AttributeList: nil,
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "term2:: def2",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestABlockTitleBetweenElementsShouldDivideThemIntoSeparateLists = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "term1:: def1",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Paragraph{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.TitleAttribute{
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Some more",
						},
					},
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "term2:: def2",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 0,
		},
	},
}

var listsTestMultiLineElementsWithParagraphContent = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Elements: asciidoc.Elements{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "def1",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Elements{
				&asciidoc.String{
					Value: "term",
				},
				&asciidoc.String{
					Value: "1",
				},
			},
		},
		&asciidoc.DescriptionListItem{
			Elements: asciidoc.Elements{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "def2",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Elements{
				&asciidoc.String{
					Value: "term",
				},
				&asciidoc.String{
					Value: "2",
				},
			},
		},
	},
}

var listsTestMultiLineElementsWithIndentedParagraphContent = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Elements: asciidoc.Elements{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: " def1",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Elements{
				&asciidoc.String{
					Value: "term",
				},
				&asciidoc.String{
					Value: "1",
				},
			},
		},
		&asciidoc.DescriptionListItem{
			Elements: asciidoc.Elements{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "  def2",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Elements{
				&asciidoc.String{
					Value: "term",
				},
				&asciidoc.String{
					Value: "2",
				},
			},
		},
	},
}

var listsTestMultiLineElementsWithIndentedParagraphContentThatIncludesCommentLines = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Elements: asciidoc.Elements{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: " def1",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Elements{
				&asciidoc.String{
					Value: "term",
				},
				&asciidoc.String{
					Value: "1",
				},
			},
		},
		&asciidoc.SingleLineComment{
			Value: " comment",
		},
		&asciidoc.DescriptionListItem{
			Elements: asciidoc.Elements{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "  def2",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Elements{
				&asciidoc.String{
					Value: "term",
				},
				&asciidoc.String{
					Value: "2",
				},
			},
		},
		&asciidoc.SingleLineComment{
			Value: " comment",
		},
		&asciidoc.String{
			Value: "  def2 continued",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestShouldNotStripCommentLineInLiteralParagraphBlockAttachedToListItem = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Elements: asciidoc.Elements{
				&asciidoc.NewLine{},
				&asciidoc.LineBreak{},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Elements{
				&asciidoc.String{
					Value: "term",
				},
				&asciidoc.String{
					Value: "1",
				},
			},
		},
		&asciidoc.String{
			Value: " line 1",
		},
		&asciidoc.NewLine{},
		&asciidoc.SingleLineComment{
			Value: " not a comment",
		},
		&asciidoc.String{
			Value: " line 3",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestMultiLineElementWithParagraphStartingWithMultipleDashesShouldNotBeSeenAsList = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Elements: asciidoc.Elements{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "  def1",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Elements{
				&asciidoc.String{
					Value: "term",
				},
				&asciidoc.String{
					Value: "1",
				},
			},
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "and a note",
				},
			},
			AttributeList: nil,
			Indent:        "  ",
			Marker:        "--",
			Checklist:     0,
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Elements: asciidoc.Elements{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "  def2",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Elements{
				&asciidoc.String{
					Value: "term",
				},
				&asciidoc.String{
					Value: "2",
				},
			},
		},
	},
}

var listsTestMultiLineElementWithMultipleTerms = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Elements: asciidoc.Elements{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "term2::",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Elements{
				&asciidoc.String{
					Value: "term",
				},
				&asciidoc.String{
					Value: "1",
				},
			},
		},
		&asciidoc.String{
			Value: "def2",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestConsecutiveTermsShareSameVarlistentryInDocbook = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Elements: asciidoc.Elements{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "alt term::",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Elements{
				&asciidoc.String{
					Value: "term",
				},
			},
		},
		&asciidoc.String{
			Value: "description",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "last::",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestMultiLineElementsWithBlankLineBeforeParagraphContent = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Elements: asciidoc.Elements{
				&asciidoc.NewLine{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "def1",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Elements{
				&asciidoc.String{
					Value: "term",
				},
				&asciidoc.String{
					Value: "1",
				},
			},
		},
		&asciidoc.DescriptionListItem{
			Elements: asciidoc.Elements{
				&asciidoc.NewLine{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "def2",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Elements{
				&asciidoc.String{
					Value: "term",
				},
				&asciidoc.String{
					Value: "2",
				},
			},
		},
	},
}

var listsTestMultiLineElementsWithParagraphAndLiteralContent = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Elements: asciidoc.Elements{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "def1",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Elements{
				&asciidoc.String{
					Value: "term",
				},
				&asciidoc.String{
					Value: "1",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "  literal",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Elements: asciidoc.Elements{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "  def2",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Elements{
				&asciidoc.String{
					Value: "term",
				},
				&asciidoc.String{
					Value: "2",
				},
			},
		},
	},
}

var listsTestMixedSingleAndMultiLineAdjacentElements = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "term1:: def1",
		},
		&asciidoc.NewLine{},
		&asciidoc.DescriptionListItem{
			Elements: asciidoc.Elements{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "def2",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Elements{
				&asciidoc.String{
					Value: "term",
				},
				&asciidoc.String{
					Value: "2",
				},
			},
		},
	},
}

var listsTestShouldDiscoverAnchorAtStartOfDescriptionTermTextAndRegisterItAsAReference = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "The highest peak in the Front Range is ",
		},
		&asciidoc.CrossReference{
			AttributeList: nil,
			Elements:      nil,
			ID:            "grays-peak",
			Format:        0,
		},
		&asciidoc.String{
			Value: ", which tops ",
		},
		&asciidoc.CrossReference{
			AttributeList: nil,
			Elements:      nil,
			ID:            "mount-evans",
			Format:        0,
		},
		&asciidoc.String{
			Value: " by just a few feet.",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Anchor{
			ID: "mount-evans",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Mount Evans",
				},
			},
		},
		&asciidoc.String{
			Value: "Mount Evans:: 14,271 feet",
		},
		&asciidoc.NewLine{},
		&asciidoc.Anchor{
			ID:       "grays-peak",
			Elements: nil,
		},
		&asciidoc.String{
			Value: "Grays Peak:: 14,278 feet",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestMissingSpaceBeforeTermDoesNotProduceDescriptionList = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "term1::def1",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "term2::def2",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestLiteralBlockInsideDescriptionList = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Elements: asciidoc.Elements{
				&asciidoc.NewLine{},
				&asciidoc.LineBreak{},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Elements{
				&asciidoc.String{
					Value: "term",
				},
			},
		},
		&asciidoc.LiteralBlock{
			AttributeList: nil,
			Delimiter: asciidoc.Delimiter{
				Type:   6,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"literal, line 1",
				"",
				"literal, line 2",
			},
		},
		&asciidoc.String{
			Value: "anotherterm:: def",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestLiteralBlockInsideDescriptionListWithTrailingLineContinuation = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Elements: asciidoc.Elements{
				&asciidoc.NewLine{},
				&asciidoc.LineBreak{},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Elements{
				&asciidoc.String{
					Value: "term",
				},
			},
		},
		&asciidoc.LiteralBlock{
			AttributeList: nil,
			Delimiter: asciidoc.Delimiter{
				Type:   6,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"literal, line 1",
				"",
				"literal, line 2",
			},
		},
		&asciidoc.LineBreak{},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "anotherterm:: def",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestMultipleListingBlocksInsideDescriptionList = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Elements: asciidoc.Elements{
				&asciidoc.NewLine{},
				&asciidoc.LineBreak{},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Elements{
				&asciidoc.String{
					Value: "term",
				},
			},
		},
		&asciidoc.Listing{
			AttributeList: nil,
			Delimiter: asciidoc.Delimiter{
				Type:   5,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"listing, line 1",
				"",
				"listing, line 2",
			},
		},
		&asciidoc.LineBreak{},
		&asciidoc.NewLine{},
		&asciidoc.Listing{
			AttributeList: nil,
			Delimiter: asciidoc.Delimiter{
				Type:   5,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"listing, line 1",
				"",
				"listing, line 2",
			},
		},
		&asciidoc.String{
			Value: "anotherterm:: def",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestOpenBlockInsideDescriptionList = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Elements: asciidoc.Elements{
				&asciidoc.NewLine{},
				&asciidoc.LineBreak{},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Elements{
				&asciidoc.String{
					Value: "term",
				},
			},
		},
		&asciidoc.OpenBlock{
			AttributeList: nil,
			Delimiter: asciidoc.Delimiter{
				Type:   7,
				Length: 2,
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Open block as description of term.",
				},
				&asciidoc.NewLine{},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "And some more detail...",
				},
				&asciidoc.NewLine{},
			},
		},
		&asciidoc.String{
			Value: "anotherterm:: def",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestParagraphAttachedByAListContinuationOnEitherSideInADescriptionList = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "term1:: def1",
		},
		&asciidoc.NewLine{},
		&asciidoc.LineBreak{},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "more detail",
		},
		&asciidoc.NewLine{},
		&asciidoc.LineBreak{},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "term2:: def2",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestParagraphAttachedByAListContinuationOnEitherSideToAMultiLineElementInADescriptionList = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Elements: asciidoc.Elements{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "def1",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Elements{
				&asciidoc.String{
					Value: "term",
				},
				&asciidoc.String{
					Value: "1",
				},
			},
		},
		&asciidoc.LineBreak{},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "more detail",
		},
		&asciidoc.NewLine{},
		&asciidoc.LineBreak{},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "term2:: def2",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestShouldContinueToParseSubsequentBlocksAttachedToListItemAfterFirstBlockIsDropped = &asciidoc.Document{
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
		&asciidoc.DescriptionListItem{
			Elements: asciidoc.Elements{
				&asciidoc.NewLine{},
				&asciidoc.LineBreak{},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Elements{
				&asciidoc.String{
					Value: "term",
				},
			},
		},
		&asciidoc.BlockImage{
			AttributeList: nil,
			ImagePath: asciidoc.Elements{
				&asciidoc.UserAttributeReference{
					Value: "unresolved",
				},
			},
		},
		&asciidoc.LineBreak{},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "paragraph",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestVerseParagraphInsideADescriptionList = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "term1:: def",
		},
		&asciidoc.NewLine{},
		&asciidoc.LineBreak{},
		&asciidoc.NewLine{},
		&asciidoc.Paragraph{
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
					Value: "la la la",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 0,
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "term2:: def",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestListInsideADescriptionList = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Elements: asciidoc.Elements{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "* level 1",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Elements{
				&asciidoc.String{
					Value: "term",
				},
				&asciidoc.String{
					Value: "1",
				},
			},
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "level 2",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "**",
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "level 1",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "term2:: def",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
	},
}

var listsTestListInsideADescriptionListOffsetByBlankLines = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Elements: asciidoc.Elements{
				&asciidoc.NewLine{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "* level 1",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Elements{
				&asciidoc.String{
					Value: "term",
				},
				&asciidoc.String{
					Value: "1",
				},
			},
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "level 2",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "**",
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "level 1",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "term2:: def",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestShouldOnlyGrabOneLineFollowingLastItemIfItemHasNoInlineDescription = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Elements: asciidoc.Elements{
				&asciidoc.NewLine{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "def1",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Elements{
				&asciidoc.String{
					Value: "term",
				},
				&asciidoc.String{
					Value: "1",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Elements: asciidoc.Elements{
				&asciidoc.NewLine{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "def2",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Elements{
				&asciidoc.String{
					Value: "term",
				},
				&asciidoc.String{
					Value: "2",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "A new paragraph",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "Another new paragraph",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestShouldOnlyGrabOneLiteralLineFollowingLastItemIfItemHasNoInlineDescription = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Elements: asciidoc.Elements{
				&asciidoc.NewLine{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "def1",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Elements{
				&asciidoc.String{
					Value: "term",
				},
				&asciidoc.String{
					Value: "1",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Elements: asciidoc.Elements{
				&asciidoc.NewLine{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "  def2",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Elements{
				&asciidoc.String{
					Value: "term",
				},
				&asciidoc.String{
					Value: "2",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "A new paragraph",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "Another new paragraph",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestShouldAppendSubsequentParagraphLiteralsToListItemAsBlockContent = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Elements: asciidoc.Elements{
				&asciidoc.NewLine{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "def1",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Elements{
				&asciidoc.String{
					Value: "term",
				},
				&asciidoc.String{
					Value: "1",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Elements: asciidoc.Elements{
				&asciidoc.NewLine{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "  def2",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Elements{
				&asciidoc.String{
					Value: "term",
				},
				&asciidoc.String{
					Value: "2",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "  literal",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "A new paragraph.",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestShouldNotMatchCommentLineThatLooksLikeDescriptionListTerm = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "before",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.SingleLineComment{
			Value: "key:: val",
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "after",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestShouldNotMatchCommentLineFollowingListThatLooksLikeDescriptionListTerm = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "item",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.SingleLineComment{
			Value: "term:: desc",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "section text",
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

var listsTestShouldNotMatchCommentLineThatLooksLikeSiblingDescriptionListTerm = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "before",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "foo:: bar",
		},
		&asciidoc.NewLine{},
		&asciidoc.SingleLineComment{
			Value: "yin:: yang",
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "after",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestShouldNotHangOnDescriptionListItemInListThatBeginsWith = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "a",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "///b::",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "c",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
	},
}

var listsTestShouldNotHangOnSiblingDescriptionListItemThatBeginsWith = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Elements: asciidoc.Elements{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "///b::",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Elements{
				&asciidoc.String{
					Value: "a",
				},
			},
		},
		&asciidoc.String{
			Value: "c",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestShouldSkipDlistTermThatBeginsWithUnlessItBeginsWith = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Elements: asciidoc.Elements{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "//ignored term:: def",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Elements{
				&asciidoc.String{
					Value: "category a",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Elements: asciidoc.Elements{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "///term:: def",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Elements{
				&asciidoc.String{
					Value: "category b",
				},
			},
		},
	},
}

var listsTestMoreThan4ConsecutiveColonsShouldBecomePartOfDescriptionListTerm = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "A term::::: a description",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestTextMethodOfDdNodeShouldReturnNilIfDdNodeOnlyContainsBlocks = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Elements: asciidoc.Elements{
				&asciidoc.NewLine{},
				&asciidoc.LineBreak{},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Elements{
				&asciidoc.String{
					Value: "term",
				},
			},
		},
		&asciidoc.String{
			Value: "paragraph",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestShouldNotParseANestedDlistDelimiterWithoutATermAsADlist = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Elements: asciidoc.Elements{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: ";;",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Elements{
				&asciidoc.String{
					Value: "t",
				},
			},
		},
	},
}

var listsTestShouldNotParseANestedIndentedDlistDelimiterWithoutATermAsADlist = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Elements: asciidoc.Elements{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "desc",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Elements{
				&asciidoc.String{
					Value: "t",
				},
			},
		},
		&asciidoc.String{
			Value: "  ;;",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestSingleLineAdjacentNestedElements = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "term1:: def1",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "label1::: detail1",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "term2:: def2",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestSingleLineAdjacentMaximumNestedElements = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "term1:: def1",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "label1::: detail1",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "name1:::: value1",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "item1;; price1",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "term2:: def2",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestSingleLineNestedElementsSeparatedByBlankLineAtTopLevel = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "term1:: def1",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "label1::: detail1",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "term2:: def2",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestSingleLineNestedElementsSeparatedByBlankLineAtNestedLevel = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "term1:: def1",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "label1::: detail1",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "label2::: detail2",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "term2:: def2",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestSingleLineAdjacentNestedElementsWithAlternateDelimiters = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "term1:: def1",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "label1;; detail1",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "term2:: def2",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestMultiLineAdjacentNestedElements = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Elements: asciidoc.Elements{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "def1",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Elements{
				&asciidoc.String{
					Value: "term",
				},
				&asciidoc.String{
					Value: "1",
				},
			},
		},
		&asciidoc.DescriptionListItem{
			Elements: asciidoc.Elements{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "detail1",
				},
			},
			AttributeList: nil,
			Marker:        ":::",
			Term: asciidoc.Elements{
				&asciidoc.String{
					Value: "label",
				},
				&asciidoc.String{
					Value: "1",
				},
			},
		},
		&asciidoc.DescriptionListItem{
			Elements: asciidoc.Elements{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "def2",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Elements{
				&asciidoc.String{
					Value: "term",
				},
				&asciidoc.String{
					Value: "2",
				},
			},
		},
	},
}

var listsTestMultiLineNestedElementsSeparatedByBlankLineAtNestedLevelRepeated = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Elements: asciidoc.Elements{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "def1",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Elements{
				&asciidoc.String{
					Value: "term",
				},
				&asciidoc.String{
					Value: "1",
				},
			},
		},
		&asciidoc.DescriptionListItem{
			Elements: asciidoc.Elements{
				&asciidoc.NewLine{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "detail1",
				},
			},
			AttributeList: nil,
			Marker:        ":::",
			Term: asciidoc.Elements{
				&asciidoc.String{
					Value: "label",
				},
				&asciidoc.String{
					Value: "1",
				},
			},
		},
		&asciidoc.DescriptionListItem{
			Elements: asciidoc.Elements{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "detail2",
				},
			},
			AttributeList: nil,
			Marker:        ":::",
			Term: asciidoc.Elements{
				&asciidoc.String{
					Value: "label",
				},
				&asciidoc.String{
					Value: "2",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "term2:: def2",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestMultiLineElementWithIndentedNestedElement = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Elements: asciidoc.Elements{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "  def1",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Elements{
				&asciidoc.String{
					Value: "term",
				},
				&asciidoc.String{
					Value: "1",
				},
			},
		},
		&asciidoc.DescriptionListItem{
			Elements: asciidoc.Elements{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "   detail1",
				},
			},
			AttributeList: nil,
			Marker:        ";;",
			Term: asciidoc.Elements{
				&asciidoc.String{
					Value: "  label",
				},
				&asciidoc.String{
					Value: "1",
				},
			},
		},
		&asciidoc.DescriptionListItem{
			Elements: asciidoc.Elements{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "  def2",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Elements{
				&asciidoc.String{
					Value: "term",
				},
				&asciidoc.String{
					Value: "2",
				},
			},
		},
	},
}

var listsTestMixedSingleAndMultiLineElementsWithIndentedNestedElements = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "term1:: def1",
		},
		&asciidoc.NewLine{},
		&asciidoc.DescriptionListItem{
			Elements: asciidoc.Elements{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "   detail1",
				},
			},
			AttributeList: nil,
			Marker:        ":::",
			Term: asciidoc.Elements{
				&asciidoc.String{
					Value: "  label",
				},
				&asciidoc.String{
					Value: "1",
				},
			},
		},
		&asciidoc.String{
			Value: "term2:: def2",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestMultiLineElementsWithFirstParagraphFoldedToTextWithAdjacentNestedElement = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "term1:: def1",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "continued",
		},
		&asciidoc.NewLine{},
		&asciidoc.DescriptionListItem{
			Elements: asciidoc.Elements{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "detail1",
				},
			},
			AttributeList: nil,
			Marker:        ":::",
			Term: asciidoc.Elements{
				&asciidoc.String{
					Value: "label",
				},
				&asciidoc.String{
					Value: "1",
				},
			},
		},
	},
}

var listsTestNestedDlistAttachedByListContinuationShouldNotConsumeDetachedParagraph = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "term:: text",
		},
		&asciidoc.NewLine{},
		&asciidoc.LineBreak{},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "nested term::: text",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "paragraph",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestNestedDlistWithAttachedBlockOffsetByEmptyLine = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Elements: asciidoc.Elements{
				&asciidoc.NewLine{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "term 1:::",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Elements{
				&asciidoc.String{
					Value: "category",
				},
			},
		},
		&asciidoc.ListContinuation{
			ChildElement: &asciidoc.OpenBlock{
				AttributeList: nil,
				Delimiter: asciidoc.Delimiter{
					Type:   7,
					Length: 2,
				},
				Elements: asciidoc.Elements{
					&asciidoc.String{
						Value: "def 1",
					},
					&asciidoc.NewLine{},
				},
			},
		},
	},
}

var listsTestShouldConvertGlossaryListWithProperSemantics = &asciidoc.Document{
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
								Value: "glossary",
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
					Value: "term 1:: def 1",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "term 2:: def 2",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 0,
		},
	},
}

var listsTestConsecutiveGlossaryTermsShouldShareSameGlossentryElementInDocbook = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Elements: asciidoc.Elements{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "alt term::",
				},
			},
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "glossary",
							},
						},
					},
					ID:      nil,
					Roles:   nil,
					Options: nil,
				},
			},
			Marker: "::",
			Term: asciidoc.Elements{
				&asciidoc.String{
					Value: "term",
				},
			},
		},
		&asciidoc.String{
			Value: "description",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "last::",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestShouldConvertHorizontalListWithProperMarkup = &asciidoc.Document{
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
								Value: "horizontal",
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
					Value: "first term:: description",
				},
				&asciidoc.NewLine{},
				&asciidoc.LineBreak{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "more detail",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 0,
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "second term:: description",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestShouldSetColWidthsOfItemAndLabelIfSpecified = &asciidoc.Document{
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
							Value: "25",
						},
					},
					Quote: 2,
				},
				&asciidoc.NamedAttribute{
					Name: "itemwidth",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "75",
						},
					},
					Quote: 2,
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "term:: def",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 0,
		},
	},
}

var listsTestShouldSetColWidthsOfItemAndLabelInDocbookIfSpecified = &asciidoc.Document{
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
							Value: "25",
						},
					},
					Quote: 2,
				},
				&asciidoc.NamedAttribute{
					Name: "itemwidth",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "75",
						},
					},
					Quote: 2,
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "term:: def",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 0,
		},
	},
}

var listsTestShouldAddStrongClassToLabelIfStrongOptionIsSet = &asciidoc.Document{
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
								Value: "horizontal",
							},
						},
					},
					ID:      nil,
					Roles:   nil,
					Options: nil,
				},
				&asciidoc.NamedAttribute{
					Name: "options",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "strong",
						},
					},
					Quote: 2,
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "term:: def",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 0,
		},
	},
}

var listsTestConsecutiveTermsInHorizontalListShouldShareSameCell = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Elements: asciidoc.Elements{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "alt term::",
				},
			},
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
			},
			Marker: "::",
			Term: asciidoc.Elements{
				&asciidoc.String{
					Value: "term",
				},
			},
		},
		&asciidoc.String{
			Value: "description",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "last::",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestConsecutiveTermsInHorizontalListShouldShareSameEntryInDocbook = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Elements: asciidoc.Elements{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "alt term::",
				},
			},
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
			},
			Marker: "::",
			Term: asciidoc.Elements{
				&asciidoc.String{
					Value: "term",
				},
			},
		},
		&asciidoc.String{
			Value: "description",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "last::",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestShouldConvertHorizontalListInDocbookWithProperMarkup = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Paragraph{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.TitleAttribute{
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Terms",
						},
					},
				},
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
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "first term:: description",
				},
				&asciidoc.NewLine{},
				&asciidoc.LineBreak{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "more detail",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 0,
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "second term:: description",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestShouldConvertQandaListInHtmlWithProperSemantics = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Elements: asciidoc.Elements{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "        Answer 1.",
				},
			},
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "qanda",
							},
						},
					},
					ID:      nil,
					Roles:   nil,
					Options: nil,
				},
			},
			Marker: "::",
			Term: asciidoc.Elements{
				&asciidoc.String{
					Value: "Question ",
				},
				&asciidoc.String{
					Value: "1",
				},
			},
		},
		&asciidoc.DescriptionListItem{
			Elements: asciidoc.Elements{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "        Answer 2.",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Elements{
				&asciidoc.String{
					Value: "Question ",
				},
				&asciidoc.String{
					Value: "2",
				},
			},
		},
		&asciidoc.LineBreak{},
		&asciidoc.NewLine{},
		&asciidoc.Paragraph{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "A note about Answer 2.",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 1,
		},
	},
}

var listsTestShouldConvertQandaListInDocBookWithProperSemantics = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Elements: asciidoc.Elements{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "        Answer 1.",
				},
			},
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "qanda",
							},
						},
					},
					ID:      nil,
					Roles:   nil,
					Options: nil,
				},
			},
			Marker: "::",
			Term: asciidoc.Elements{
				&asciidoc.String{
					Value: "Question ",
				},
				&asciidoc.String{
					Value: "1",
				},
			},
		},
		&asciidoc.DescriptionListItem{
			Elements: asciidoc.Elements{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "        Answer 2.",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Elements{
				&asciidoc.String{
					Value: "Question ",
				},
				&asciidoc.String{
					Value: "2",
				},
			},
		},
		&asciidoc.LineBreak{},
		&asciidoc.NewLine{},
		&asciidoc.Paragraph{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "A note about Answer 2.",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 1,
		},
	},
}

var listsTestConsecutiveQuestionsShouldShareSameQuestionElementInDocbook = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Elements: asciidoc.Elements{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "follow-up question::",
				},
			},
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "qanda",
							},
						},
					},
					ID:      nil,
					Roles:   nil,
					Options: nil,
				},
			},
			Marker: "::",
			Term: asciidoc.Elements{
				&asciidoc.String{
					Value: "question",
				},
			},
		},
		&asciidoc.String{
			Value: "response",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "last question::",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestShouldConvertBibliographyListWithProperSemantics = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "[",
				},
				&asciidoc.Anchor{
					ID:       "taoup",
					Elements: nil,
				},
				&asciidoc.String{
					Value: "] Eric Steven Raymond. _The Art of Unix",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "  Programming_. Addison-Wesley. ISBN 0-13-142901-9.",
				},
			},
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "bibliography",
							},
						},
					},
					ID:      nil,
					Roles:   nil,
					Options: nil,
				},
			},
			Indent:    "",
			Marker:    "-",
			Checklist: 0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "[",
				},
				&asciidoc.Anchor{
					ID:       "walsh-muellner",
					Elements: nil,
				},
				&asciidoc.String{
					Value: "] Norman Walsh & Leonard Muellner.",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "  ",
				},
				&asciidoc.Italic{
					AttributeList: nil,
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "DocBook - The Definitive Guide",
						},
					},
				},
				&asciidoc.String{
					Value: ". O'Reilly & Associates. 1999.",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "  ISBN 1-56592-580-7.",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "-",
			Checklist:     0,
		},
	},
}

var listsTestShouldConvertBibliographyListWithProperSemanticsToDocBook = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "[",
				},
				&asciidoc.Anchor{
					ID:       "taoup",
					Elements: nil,
				},
				&asciidoc.String{
					Value: "] Eric Steven Raymond. _The Art of Unix",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "  Programming_. Addison-Wesley. ISBN 0-13-142901-9.",
				},
			},
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "bibliography",
							},
						},
					},
					ID:      nil,
					Roles:   nil,
					Options: nil,
				},
			},
			Indent:    "",
			Marker:    "-",
			Checklist: 0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "[",
				},
				&asciidoc.Anchor{
					ID:       "walsh-muellner",
					Elements: nil,
				},
				&asciidoc.String{
					Value: "] Norman Walsh & Leonard Muellner.",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "  ",
				},
				&asciidoc.Italic{
					AttributeList: nil,
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "DocBook - The Definitive Guide",
						},
					},
				},
				&asciidoc.String{
					Value: ". O'Reilly & Associates. 1999.",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "  ISBN 1-56592-580-7.",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "-",
			Checklist:     0,
		},
	},
}

var listsTestShouldWarnIfABibliographyIdIsAlreadyInUse = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "[",
				},
				&asciidoc.Anchor{
					ID:       "Fowler",
					Elements: nil,
				},
				&asciidoc.String{
					Value: "] Fowler M. ",
				},
				&asciidoc.Italic{
					AttributeList: nil,
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "Analysis Patterns: Reusable Object Models",
						},
					},
				},
				&asciidoc.String{
					Value: ".",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "Addison-Wesley. 1997.",
				},
			},
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "bibliography",
							},
						},
					},
					ID:      nil,
					Roles:   nil,
					Options: nil,
				},
			},
			Indent:    "",
			Marker:    "*",
			Checklist: 0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "[",
				},
				&asciidoc.Anchor{
					ID:       "Fowler",
					Elements: nil,
				},
				&asciidoc.String{
					Value: "] Fowler M. ",
				},
				&asciidoc.Italic{
					AttributeList: nil,
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "Analysis Patterns: Reusable Object Models",
						},
					},
				},
				&asciidoc.String{
					Value: ".",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "Addison-Wesley. 1997.",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
	},
}

var listsTestShouldAutomaticallyAddBibliographyStyleToTopLevelListsInBibliographySection = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "bibliography",
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
				&asciidoc.UnorderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "[",
						},
						&asciidoc.Anchor{
							ID:       "taoup",
							Elements: nil,
						},
						&asciidoc.String{
							Value: "] Eric Steven Raymond. _The Art of Unix",
						},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "  Programming_. Addison-Wesley. ISBN 0-13-142901-9.",
						},
					},
					AttributeList: asciidoc.AttributeList{
						&asciidoc.TitleAttribute{
							Val: asciidoc.Elements{
								&asciidoc.String{
									Value: "Books",
								},
							},
						},
					},
					Indent:    "",
					Marker:    "*",
					Checklist: 0,
				},
				&asciidoc.UnorderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "[",
						},
						&asciidoc.Anchor{
							ID:       "walsh-muellner",
							Elements: nil,
						},
						&asciidoc.String{
							Value: "] Norman Walsh & Leonard Muellner.",
						},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "  ",
						},
						&asciidoc.Italic{
							AttributeList: nil,
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "DocBook - The Definitive Guide",
								},
							},
						},
						&asciidoc.String{
							Value: ". O'Reilly & Associates. 1999.",
						},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "  ISBN 1-56592-580-7.",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        "*",
					Checklist:     0,
				},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.UnorderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "[",
						},
						&asciidoc.Anchor{
							ID:       "doc-writer",
							Elements: nil,
						},
						&asciidoc.String{
							Value: "] Doc Writer. ",
						},
						&asciidoc.Italic{
							AttributeList: nil,
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "Documentation As Code",
								},
							},
						},
						&asciidoc.String{
							Value: ". Static Times, 54. August 2016.",
						},
					},
					AttributeList: asciidoc.AttributeList{
						&asciidoc.TitleAttribute{
							Val: asciidoc.Elements{
								&asciidoc.String{
									Value: "Periodicals",
								},
							},
						},
					},
					Indent:    "",
					Marker:    "*",
					Checklist: 0,
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Bibliography",
				},
			},
			Level: 1,
		},
	},
}

var listsTestShouldNotRecognizeBibliographyAnchorThatBeginsWithADigit = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "[[[1984]]] George Orwell. ",
				},
				&asciidoc.Italic{
					AttributeList: nil,
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "1984",
						},
					},
				},
				&asciidoc.String{
					Value: ". New American Library. 1950.",
				},
			},
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "bibliography",
							},
						},
					},
					ID:      nil,
					Roles:   nil,
					Options: nil,
				},
			},
			Indent:    "",
			Marker:    "-",
			Checklist: 0,
		},
	},
}

var listsTestShouldRecognizeBibliographyAnchorThatContainsADigitButDoesNotStartWithOne = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "[",
				},
				&asciidoc.Anchor{
					ID:       "_1984",
					Elements: nil,
				},
				&asciidoc.String{
					Value: "] George Orwell. ",
				},
				&asciidoc.DoubleItalic{
					AttributeList: nil,
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "1984",
						},
					},
				},
				&asciidoc.String{
					Value: ". New American Library. 1950.",
				},
			},
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "bibliography",
							},
						},
					},
					ID:      nil,
					Roles:   nil,
					Options: nil,
				},
			},
			Indent:    "",
			Marker:    "-",
			Checklist: 0,
		},
	},
}

var listsTestShouldCatalogBibliographyAnchorsInBibliographyList = &asciidoc.Document{
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
					Value: "Please read ",
				},
				&asciidoc.CrossReference{
					AttributeList: nil,
					Elements:      nil,
					ID:            "Fowler_1997",
					Format:        0,
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
							Style: &asciidoc.ShorthandStyle{
								Elements: asciidoc.Elements{
									&asciidoc.String{
										Value: "bibliography",
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
						&asciidoc.UnorderedListItem{
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "[",
								},
								&asciidoc.Anchor{
									ID:       "Fowler_1997",
									Elements: nil,
								},
								&asciidoc.String{
									Value: "] Fowler M. ",
								},
								&asciidoc.Italic{
									AttributeList: nil,
									Elements: asciidoc.Elements{
										&asciidoc.String{
											Value: "Analysis Patterns: Reusable Object Models",
										},
									},
								},
								&asciidoc.String{
									Value: ". Addison-Wesley. 1997.",
								},
							},
							AttributeList: nil,
							Indent:        "",
							Marker:        "*",
							Checklist:     0,
						},
					},
					Title: asciidoc.Elements{
						&asciidoc.String{
							Value: "References",
						},
					},
					Level: 1,
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Article Title",
				},
			},
			Level: 0,
		},
	},
}

var listsTestShouldUseReftextFromBibliographyAnchorAtXrefAndEntry = &asciidoc.Document{
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
					Value: "Begin with ",
				},
				&asciidoc.CrossReference{
					AttributeList: nil,
					Elements:      nil,
					ID:            "TMMM",
					Format:        0,
				},
				&asciidoc.String{
					Value: ".",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "Then move on to ",
				},
				&asciidoc.CrossReference{
					AttributeList: nil,
					Elements:      nil,
					ID:            "Fowler_1997",
					Format:        0,
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
							Style: &asciidoc.ShorthandStyle{
								Elements: asciidoc.Elements{
									&asciidoc.String{
										Value: "bibliography",
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
						&asciidoc.UnorderedListItem{
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "[",
								},
								&asciidoc.Anchor{
									ID:       "TMMM",
									Elements: nil,
								},
								&asciidoc.String{
									Value: "] Brooks F. ",
								},
								&asciidoc.Italic{
									AttributeList: nil,
									Elements: asciidoc.Elements{
										&asciidoc.String{
											Value: "The Mythical Man-Month",
										},
									},
								},
								&asciidoc.String{
									Value: ". Addison-Wesley. 1975.",
								},
							},
							AttributeList: nil,
							Indent:        "",
							Marker:        "*",
							Checklist:     0,
						},
						&asciidoc.UnorderedListItem{
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "[",
								},
								&asciidoc.Anchor{
									ID: "Fowler_1997",
									Elements: asciidoc.Elements{
										&asciidoc.String{
											Value: "1",
										},
									},
								},
								&asciidoc.String{
									Value: "] Fowler M. ",
								},
								&asciidoc.Italic{
									AttributeList: nil,
									Elements: asciidoc.Elements{
										&asciidoc.String{
											Value: "Analysis Patterns: Reusable Object Models",
										},
									},
								},
								&asciidoc.String{
									Value: ". Addison-Wesley. 1997.",
								},
							},
							AttributeList: nil,
							Indent:        "",
							Marker:        "*",
							Checklist:     0,
						},
					},
					Title: asciidoc.Elements{
						&asciidoc.String{
							Value: "References",
						},
					},
					Level: 1,
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Article Title",
				},
			},
			Level: 0,
		},
	},
}

var listsTestShouldAssignReftextOfBibliographyAnchorToXreflabelInDocBookBackend = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "[",
				},
				&asciidoc.Anchor{
					ID: "Fowler_1997",
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "1",
						},
					},
				},
				&asciidoc.String{
					Value: "] Fowler M. ",
				},
				&asciidoc.Italic{
					AttributeList: nil,
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "Analysis Patterns: Reusable Object Models",
						},
					},
				},
				&asciidoc.String{
					Value: ". Addison-Wesley. 1997.",
				},
			},
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "bibliography",
							},
						},
					},
					ID:      nil,
					Roles:   nil,
					Options: nil,
				},
			},
			Indent:    "",
			Marker:    "*",
			Checklist: 0,
		},
	},
}

var listsTestFoldsTextFromSubsequentLine = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Elements: asciidoc.Elements{
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "def1",
						},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Elements{
						&asciidoc.String{
							Value: "term",
						},
						&asciidoc.String{
							Value: "1",
						},
					},
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var listsTestFoldsTextFromFirstLineAfterBlankLines = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Elements: asciidoc.Elements{
						&asciidoc.NewLine{},
						&asciidoc.NewLine{},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "def1",
						},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Elements{
						&asciidoc.String{
							Value: "term",
						},
						&asciidoc.String{
							Value: "1",
						},
					},
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var listsTestFoldsTextFromFirstLineAfterBlankLineAndImmediatelyPrecedingNextItem = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Elements: asciidoc.Elements{
						&asciidoc.NewLine{},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "def1",
						},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Elements{
						&asciidoc.String{
							Value: "term",
						},
						&asciidoc.String{
							Value: "1",
						},
					},
				},
				&asciidoc.String{
					Value: "term2:: def2",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var listsTestParagraphOffsetByBlankLinesDoesNotBreakListIfLabelDoesNotHaveInlineText = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Elements: asciidoc.Elements{
						&asciidoc.NewLine{},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "def1",
						},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Elements{
						&asciidoc.String{
							Value: "term",
						},
						&asciidoc.String{
							Value: "1",
						},
					},
				},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "term2:: def2",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var listsTestFoldsTextFromFirstLineAfterCommentLine = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Elements: asciidoc.Elements{
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "// comment",
						},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Elements{
						&asciidoc.String{
							Value: "term",
						},
						&asciidoc.String{
							Value: "1",
						},
					},
				},
				&asciidoc.String{
					Value: "def1",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var listsTestFoldsTextFromLineFollowingCommentLineOffsetByBlankLine = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Elements: asciidoc.Elements{
						&asciidoc.NewLine{},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "// comment",
						},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Elements{
						&asciidoc.String{
							Value: "term",
						},
						&asciidoc.String{
							Value: "1",
						},
					},
				},
				&asciidoc.String{
					Value: "def1",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var listsTestFoldsTextFromSubsequentIndentedLine = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Elements: asciidoc.Elements{
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "  def1",
						},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Elements{
						&asciidoc.String{
							Value: "term",
						},
						&asciidoc.String{
							Value: "1",
						},
					},
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var listsTestFoldsTextFromIndentedLineAfterBlankLine = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Elements: asciidoc.Elements{
						&asciidoc.NewLine{},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "  def1",
						},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Elements{
						&asciidoc.String{
							Value: "term",
						},
						&asciidoc.String{
							Value: "1",
						},
					},
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var listsTestFoldsTextThatLooksLikeRulerOffsetByBlankLine = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Elements: asciidoc.Elements{
						&asciidoc.NewLine{},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "'''",
						},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Elements{
						&asciidoc.String{
							Value: "term",
						},
						&asciidoc.String{
							Value: "1",
						},
					},
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var listsTestFoldsTextThatLooksLikeRulerOffsetByBlankLineAndLineComment = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Elements: asciidoc.Elements{
						&asciidoc.NewLine{},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "// comment",
						},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Elements{
						&asciidoc.String{
							Value: "term",
						},
						&asciidoc.String{
							Value: "1",
						},
					},
				},
				&asciidoc.ThematicBreak{
					AttributeList: nil,
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var listsTestFoldsTextThatLooksLikeRulerAndTheLineFollowingItOffsetByBlankLine = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Elements: asciidoc.Elements{
						&asciidoc.NewLine{},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "'''",
						},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Elements{
						&asciidoc.String{
							Value: "term",
						},
						&asciidoc.String{
							Value: "1",
						},
					},
				},
				&asciidoc.String{
					Value: "continued",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var listsTestFoldsTextThatLooksLikeTitleOffsetByBlankLine = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Elements: asciidoc.Elements{
						&asciidoc.NewLine{},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: ".def1",
						},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Elements{
						&asciidoc.String{
							Value: "term",
						},
						&asciidoc.String{
							Value: "1",
						},
					},
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var listsTestFoldsTextThatLooksLikeTitleOffsetByBlankLineAndLineComment = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Elements: asciidoc.Elements{
						&asciidoc.NewLine{},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "// comment",
						},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Elements{
						&asciidoc.String{
							Value: "term",
						},
						&asciidoc.String{
							Value: "1",
						},
					},
				},
				&asciidoc.Paragraph{
					AttributeList: asciidoc.AttributeList{
						&asciidoc.TitleAttribute{
							Val: asciidoc.Elements{
								&asciidoc.String{
									Value: "def1",
								},
							},
						},
					},
					Elements:   asciidoc.Elements{},
					Admonition: 0,
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var listsTestFoldsTextThatLooksLikeAdmonitionOffsetByBlankLine = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Elements: asciidoc.Elements{
						&asciidoc.NewLine{},
						&asciidoc.NewLine{},
						&asciidoc.Paragraph{
							AttributeList: nil,
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "def1",
								},
							},
							Admonition: 1,
						},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Elements{
						&asciidoc.String{
							Value: "term",
						},
						&asciidoc.String{
							Value: "1",
						},
					},
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var listsTestFoldsTextThatLooksLikeSectionTitleOffsetByBlankLine = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Elements: asciidoc.Elements{
						&asciidoc.NewLine{},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "== Another Section",
						},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Elements{
						&asciidoc.String{
							Value: "term",
						},
						&asciidoc.String{
							Value: "1",
						},
					},
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var listsTestFoldsTextOfFirstLiteralLineOffsetByBlankLineAppendsSubsequentLiteralsOffsetByBlankLineAsBlocks = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Elements: asciidoc.Elements{
						&asciidoc.NewLine{},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "  def1",
						},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Elements{
						&asciidoc.String{
							Value: "term",
						},
						&asciidoc.String{
							Value: "1",
						},
					},
				},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "  literal",
				},
				&asciidoc.NewLine{},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "  literal",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var listsTestFoldsTextOfSubsequentLineAndAppendsFollowingLiteralLineOffsetByBlankLineAsBlockIfTermHasNoInlineDescription = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Elements: asciidoc.Elements{
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "def1",
						},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Elements{
						&asciidoc.String{
							Value: "term",
						},
						&asciidoc.String{
							Value: "1",
						},
					},
				},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "  literal",
				},
				&asciidoc.NewLine{},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "term2:: def2",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var listsTestAppendsLiteralLineAttachedByContinuationAsBlockIfItemHasNoInlineDescription = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Elements: asciidoc.Elements{
						&asciidoc.NewLine{},
						&asciidoc.LineBreak{},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Elements{
						&asciidoc.String{
							Value: "term",
						},
						&asciidoc.String{
							Value: "1",
						},
					},
				},
				&asciidoc.String{
					Value: "  literal",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var listsTestAppendsLiteralLineAttachedByContinuationAsBlockIfItemHasNoInlineDescriptionFollowedByRuler = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Elements: asciidoc.Elements{
						&asciidoc.NewLine{},
						&asciidoc.LineBreak{},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Elements{
						&asciidoc.String{
							Value: "term",
						},
						&asciidoc.String{
							Value: "1",
						},
					},
				},
				&asciidoc.String{
					Value: "  literal",
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
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var listsTestAppendsLineAttachedByContinuationAsBlockIfItemHasNoInlineDescriptionFollowedByRuler = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Elements: asciidoc.Elements{
						&asciidoc.NewLine{},
						&asciidoc.LineBreak{},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Elements{
						&asciidoc.String{
							Value: "term",
						},
						&asciidoc.String{
							Value: "1",
						},
					},
				},
				&asciidoc.String{
					Value: "para",
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
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var listsTestAppendsLineAttachedByContinuationAsBlockIfItemHasNoInlineDescriptionFollowedByBlock = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Elements: asciidoc.Elements{
						&asciidoc.NewLine{},
						&asciidoc.LineBreak{},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Elements{
						&asciidoc.String{
							Value: "term",
						},
						&asciidoc.String{
							Value: "1",
						},
					},
				},
				&asciidoc.String{
					Value: "para",
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
						"literal",
					},
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var listsTestAppendsBlockAttachedByContinuationButNotSubsequentBlockNotAttachedByContinuation = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Elements: asciidoc.Elements{
						&asciidoc.NewLine{},
						&asciidoc.LineBreak{},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Elements{
						&asciidoc.String{
							Value: "term",
						},
						&asciidoc.String{
							Value: "1",
						},
					},
				},
				&asciidoc.LiteralBlock{
					AttributeList: nil,
					Delimiter: asciidoc.Delimiter{
						Type:   6,
						Length: 4,
					},
					LineList: asciidoc.LineList{
						"literal",
					},
				},
				&asciidoc.LiteralBlock{
					AttributeList: nil,
					Delimiter: asciidoc.Delimiter{
						Type:   6,
						Length: 4,
					},
					LineList: asciidoc.LineList{
						"detached",
					},
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var listsTestAppendsListIfItemHasNoInlineDescription = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Elements: asciidoc.Elements{
						&asciidoc.NewLine{},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "* one",
						},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Elements{
						&asciidoc.String{
							Value: "term",
						},
						&asciidoc.String{
							Value: "1",
						},
					},
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
					Checklist:     0,
				},
				&asciidoc.UnorderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "three",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        "*",
					Checklist:     0,
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var listsTestAppendsListToFirstTermWhenFollowedImmediatelyBySecondTerm = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Elements: asciidoc.Elements{
						&asciidoc.NewLine{},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "* one",
						},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Elements{
						&asciidoc.String{
							Value: "term",
						},
						&asciidoc.String{
							Value: "1",
						},
					},
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
					Checklist:     0,
				},
				&asciidoc.UnorderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "three",
						},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "term2:: def2",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        "*",
					Checklist:     0,
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var listsTestAppendsIndentedListToFirstTermThatIsAdjacentToSecondTerm = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Elements: asciidoc.Elements{
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "  description 1",
						},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Elements{
						&asciidoc.String{
							Value: "label ",
						},
						&asciidoc.String{
							Value: "1",
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
					Indent:        "  ",
					Marker:        "*",
					Checklist:     0,
				},
				&asciidoc.UnorderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "two",
						},
					},
					AttributeList: nil,
					Indent:        "  ",
					Marker:        "*",
					Checklist:     0,
				},
				&asciidoc.UnorderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "three",
						},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "label 2::",
						},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "  description 2",
						},
					},
					AttributeList: nil,
					Indent:        "  ",
					Marker:        "*",
					Checklist:     0,
				},
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
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var listsTestAppendsIndentedListToFirstTermThatIsAttachedByAContinuationAndAdjacentToSecondTerm = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Elements: asciidoc.Elements{
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "  description 1",
						},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Elements{
						&asciidoc.String{
							Value: "label ",
						},
						&asciidoc.String{
							Value: "1",
						},
					},
				},
				&asciidoc.LineBreak{},
				&asciidoc.NewLine{},
				&asciidoc.UnorderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "one",
						},
					},
					AttributeList: nil,
					Indent:        "  ",
					Marker:        "*",
					Checklist:     0,
				},
				&asciidoc.UnorderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "two",
						},
					},
					AttributeList: nil,
					Indent:        "  ",
					Marker:        "*",
					Checklist:     0,
				},
				&asciidoc.UnorderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "three",
						},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "label 2::",
						},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "  description 2",
						},
					},
					AttributeList: nil,
					Indent:        "  ",
					Marker:        "*",
					Checklist:     0,
				},
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
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var listsTestAppendsListAndParagraphBlockWhenLineFollowingListAttachedByContinuation = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Elements: asciidoc.Elements{
						&asciidoc.NewLine{},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "* one",
						},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Elements{
						&asciidoc.String{
							Value: "term",
						},
						&asciidoc.String{
							Value: "1",
						},
					},
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
					Checklist:     0,
				},
				&asciidoc.UnorderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "three",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        "*",
					Checklist:     0,
				},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.LineBreak{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "para",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var listsTestFirstContinuedLineAssociatedWithNestedListItemAndSecondContinuedLineAssociatedWithTerm = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Elements: asciidoc.Elements{
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "* one",
						},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Elements{
						&asciidoc.String{
							Value: "term",
						},
						&asciidoc.String{
							Value: "1",
						},
					},
				},
				&asciidoc.LineBreak{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "nested list para",
				},
				&asciidoc.NewLine{},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.LineBreak{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "term1 para",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var listsTestLiteralLineAttachedByContinuationSwallowsAdjacentLineThatLooksLikeTerm = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Elements: asciidoc.Elements{
						&asciidoc.NewLine{},
						&asciidoc.LineBreak{},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Elements{
						&asciidoc.String{
							Value: "term",
						},
						&asciidoc.String{
							Value: "1",
						},
					},
				},
				&asciidoc.String{
					Value: "  literal",
				},
				&asciidoc.NewLine{},
				&asciidoc.DescriptionListItem{
					Elements: asciidoc.Elements{
						&asciidoc.NewLine{},
						&asciidoc.LineBreak{},
					},
					AttributeList: nil,
					Marker:        ":::",
					Term: asciidoc.Elements{
						&asciidoc.String{
							Value: "notnestedterm",
						},
					},
				},
				&asciidoc.String{
					Value: "  literal",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "notnestedterm:::",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var listsTestLineAttachedByContinuationIsAppendedAsParagraphIfTermHasNoInlineDescription = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Elements: asciidoc.Elements{
						&asciidoc.NewLine{},
						&asciidoc.LineBreak{},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Elements{
						&asciidoc.String{
							Value: "term",
						},
						&asciidoc.String{
							Value: "1",
						},
					},
				},
				&asciidoc.String{
					Value: "para",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var listsTestAttachedParagraphDoesNotBreakOnAdjacentNestedDescriptionListTerm = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "term1:: def",
		},
		&asciidoc.NewLine{},
		&asciidoc.LineBreak{},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "more description",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "not a term::: def",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestAttachedParagraphIsTerminatedByAdjacentSiblingDescriptionListTerm = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "term1:: def",
		},
		&asciidoc.NewLine{},
		&asciidoc.LineBreak{},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "more description",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "not a term:: def",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestAttachedStyledParagraphDoesNotBreakOnAdjacentNestedDescriptionListTerm = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "term1:: def",
		},
		&asciidoc.NewLine{},
		&asciidoc.LineBreak{},
		&asciidoc.NewLine{},
		&asciidoc.Paragraph{
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
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "more description",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "not a term::: def",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 0,
		},
	},
}

var listsTestAppendsLineAsParagraphIfAttachedByContinuationFollowingBlankLineAndLineCommentWhenTermHasNoInlineDescription = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Elements: asciidoc.Elements{
						&asciidoc.NewLine{},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "// comment",
						},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Elements{
						&asciidoc.String{
							Value: "term",
						},
						&asciidoc.String{
							Value: "1",
						},
					},
				},
				&asciidoc.LineBreak{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "para",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var listsTestLineAttachedByContinuationOffsetByBlankLineIsAppendedAsParagraphIfTermHasNoInlineDescription = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Elements: asciidoc.Elements{
						&asciidoc.NewLine{},
						&asciidoc.NewLine{},
						&asciidoc.LineBreak{},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Elements{
						&asciidoc.String{
							Value: "term",
						},
						&asciidoc.String{
							Value: "1",
						},
					},
				},
				&asciidoc.String{
					Value: "para",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var listsTestDelimitedBlockBreaksListEvenWhenTermHasNoInlineDescription = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Elements: asciidoc.Elements{
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "====",
						},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Elements{
						&asciidoc.String{
							Value: "term",
						},
						&asciidoc.String{
							Value: "1",
						},
					},
				},
				&asciidoc.String{
					Value: "detached",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "====",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var listsTestBlockAttributeLineAboveDelimitedBlockThatBreaksADlistIsNotDuplicated = &asciidoc.Document{
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
					Value: "term:: desc",
				},
				&asciidoc.NewLine{},
				&asciidoc.Listing{
					AttributeList: asciidoc.AttributeList{
						&asciidoc.ShorthandAttribute{
							Style: nil,
							ID:    nil,
							Roles: []*asciidoc.ShorthandRole{
								&asciidoc.ShorthandRole{
									Elements: asciidoc.Elements{
										&asciidoc.String{
											Value: "rolename",
										},
									},
								},
							},
							Options: nil,
						},
					},
					Delimiter: asciidoc.Delimiter{
						Type:   5,
						Length: 4,
					},
					LineList: asciidoc.LineList{
						"detached",
					},
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var listsTestBlockAttributeLineAboveParagraphBreaksListEvenWhenTermHasNoInlineDescription = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Elements: asciidoc.Elements{
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "[verse]",
						},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Elements{
						&asciidoc.String{
							Value: "term",
						},
						&asciidoc.String{
							Value: "1",
						},
					},
				},
				&asciidoc.String{
					Value: "detached",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var listsTestBlockAttributeLineAboveParagraphThatBreaksADlistIsNotDuplicated = &asciidoc.Document{
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
					Value: "term:: desc",
				},
				&asciidoc.NewLine{},
				&asciidoc.Paragraph{
					AttributeList: asciidoc.AttributeList{
						&asciidoc.ShorthandAttribute{
							Style: nil,
							ID:    nil,
							Roles: []*asciidoc.ShorthandRole{
								&asciidoc.ShorthandRole{
									Elements: asciidoc.Elements{
										&asciidoc.String{
											Value: "rolename",
										},
									},
								},
							},
							Options: nil,
						},
					},
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "detached",
						},
						&asciidoc.NewLine{},
					},
					Admonition: 0,
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var listsTestBlockAnchorLineBreaksListEvenWhenTermHasNoInlineDescription = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Elements: asciidoc.Elements{
						&asciidoc.NewLine{},
						&asciidoc.Anchor{
							ID:       "id",
							Elements: nil,
						},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Elements{
						&asciidoc.String{
							Value: "term",
						},
						&asciidoc.String{
							Value: "1",
						},
					},
				},
				&asciidoc.String{
					Value: "detached",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var listsTestBlockAttributeLinesAboveNestedHorizontalListDoesNotBreakList = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Elements: asciidoc.Elements{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "[horizontal]",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Elements{
				&asciidoc.String{
					Value: "Operating Systems",
				},
			},
		},
		&asciidoc.String{
			Value: "  Linux::: Fedora",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "  BSD::: OpenBSD",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Elements: asciidoc.Elements{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "  PaaS::: OpenShift",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Elements{
				&asciidoc.String{
					Value: "Cloud Providers",
				},
			},
		},
		&asciidoc.String{
			Value: "  IaaS::: AWS",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestBlockAttributeLinesAboveNestedListWithStyleDoesNotBreakList = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Elements: asciidoc.Elements{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "* get groceries",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Elements{
				&asciidoc.String{
					Value: "TODO List",
				},
			},
		},
		&asciidoc.DescriptionListItem{
			Elements: asciidoc.Elements{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "[square]",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Elements{
				&asciidoc.String{
					Value: "Grocery List",
				},
			},
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "bread",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "milk",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "lettuce",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
	},
}

var listsTestMultipleBlockAttributeLinesAboveNestedListDoesNotBreakList = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Elements: asciidoc.Elements{
				&asciidoc.NewLine{},
				&asciidoc.Anchor{
					ID:       "variants",
					Elements: nil,
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Elements{
				&asciidoc.String{
					Value: "Operating Systems",
				},
			},
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
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "  Linux::: Fedora",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "  BSD::: OpenBSD",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 0,
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Elements: asciidoc.Elements{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "  PaaS::: OpenShift",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Elements{
				&asciidoc.String{
					Value: "Cloud Providers",
				},
			},
		},
		&asciidoc.String{
			Value: "  IaaS::: AWS",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestMultipleBlockAttributeLinesSeparatedByEmptyLineAboveNestedListDoesNotBreakList = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Elements: asciidoc.Elements{
				&asciidoc.NewLine{},
				&asciidoc.Anchor{
					ID:       "variants",
					Elements: nil,
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Elements{
				&asciidoc.String{
					Value: "Operating Systems",
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
			},
			Elements:   asciidoc.Elements{},
			Admonition: 0,
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "  Linux::: Fedora",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "  BSD::: OpenBSD",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Elements: asciidoc.Elements{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "  PaaS::: OpenShift",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Elements{
				&asciidoc.String{
					Value: "Cloud Providers",
				},
			},
		},
		&asciidoc.String{
			Value: "  IaaS::: AWS",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestFoldsTextFromInlineDescriptionAndSubsequentLine = &asciidoc.Document{
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
					Value: "term1:: def1",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "continued",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var listsTestFoldsTextFromInlineDescriptionAndSubsequentLines = &asciidoc.Document{
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
					Value: "term1:: def1",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "continued",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "continued",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var listsTestFoldsTextFromInlineDescriptionAndLineFollowingCommentLine = &asciidoc.Document{
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
					Value: "term1:: def1",
				},
				&asciidoc.NewLine{},
				&asciidoc.SingleLineComment{
					Value: " comment",
				},
				&asciidoc.String{
					Value: "continued",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var listsTestFoldsTextFromInlineDescriptionAndSubsequentIndentedLine = &asciidoc.Document{
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
					Value: "term1:: def1",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "  continued",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "List",
				},
			},
			Level: 1,
		},
	},
}

var listsTestAppendsLiteralLineOffsetByBlankLineAsBlockIfItemHasInlineDescription = &asciidoc.Document{
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
					Value: "term1:: def1",
				},
				&asciidoc.NewLine{},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "  literal",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var listsTestAppendsLiteralLineOffsetByBlankLineAsBlockAndAppendsLineAfterContinuationAsBlockIfItemHasInlineDescription = &asciidoc.Document{
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
					Value: "term1:: def1",
				},
				&asciidoc.NewLine{},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "  literal",
				},
				&asciidoc.NewLine{},
				&asciidoc.LineBreak{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "para",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var listsTestAppendsLineAfterContinuationAsBlockAndLiteralLineOffsetByBlankLineAsBlockIfItemHasInlineDescription = &asciidoc.Document{
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
					Value: "term1:: def1",
				},
				&asciidoc.NewLine{},
				&asciidoc.LineBreak{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "para",
				},
				&asciidoc.NewLine{},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "  literal",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var listsTestAppendsListIfItemHasInlineDescription = &asciidoc.Document{
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
					Value: "term1:: def1",
				},
				&asciidoc.NewLine{},
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
					Checklist:     0,
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
					Checklist:     0,
				},
				&asciidoc.UnorderedListItem{
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "three",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        "*",
					Checklist:     0,
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var listsTestAppendsLiteralLineAttachedByContinuationAsBlockIfItemHasInlineDescriptionFollowedByRuler = &asciidoc.Document{
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
					Value: "term1:: def1",
				},
				&asciidoc.NewLine{},
				&asciidoc.LineBreak{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "  literal",
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
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var listsTestLineOffsetByBlankLineBreaksListIfTermHasInlineDescription = &asciidoc.Document{
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
					Value: "term1:: def1",
				},
				&asciidoc.NewLine{},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "detached",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var listsTestNestedTermWithDescriptionDoesNotConsumeFollowingHeading = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Elements: asciidoc.Elements{
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "  def",
						},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Elements{
						&asciidoc.String{
							Value: "term",
						},
					},
				},
				&asciidoc.DescriptionListItem{
					Elements: asciidoc.Elements{
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "    nesteddef",
						},
					},
					AttributeList: nil,
					Marker:        ";;",
					Term: asciidoc.Elements{
						&asciidoc.String{
							Value: "  nestedterm",
						},
					},
				},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "Detached",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "~~~~~~~~",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var listsTestLineAttachedByContinuationIsAppendedAsParagraphIfTermHasInlineDescriptionFollowedByDetachedParagraph = &asciidoc.Document{
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
					Value: "term1:: def1",
				},
				&asciidoc.NewLine{},
				&asciidoc.LineBreak{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "para",
				},
				&asciidoc.NewLine{},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "detached",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var listsTestLineAttachedByContinuationIsAppendedAsParagraphIfTermHasInlineDescriptionFollowedByDetachedBlock = &asciidoc.Document{
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
					Value: "term1:: def1",
				},
				&asciidoc.NewLine{},
				&asciidoc.LineBreak{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "para",
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
							Value: "detached",
						},
						&asciidoc.NewLine{},
					},
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var listsTestLineAttachedByContinuationOffsetByLineCommentIsAppendedAsParagraphIfTermHasInlineDescription = &asciidoc.Document{
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
					Value: "term1:: def1",
				},
				&asciidoc.NewLine{},
				&asciidoc.SingleLineComment{
					Value: " comment",
				},
				&asciidoc.LineBreak{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "para",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var listsTestLineAttachedByContinuationOffsetByBlankLineIsAppendedAsParagraphIfTermHasInlineDescription = &asciidoc.Document{
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
					Value: "term1:: def1",
				},
				&asciidoc.NewLine{},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.LineBreak{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "para",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var listsTestLineCommentOffsetByBlankLineDividesListsBecauseItemHasText = &asciidoc.Document{
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
					Value: "term1:: def1",
				},
				&asciidoc.NewLine{},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.SingleLineComment{
					Value: "",
				},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "term2:: def2",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var listsTestRulerOffsetByBlankLineDividesListsBecauseItemHasText = &asciidoc.Document{
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
					Value: "term1:: def1",
				},
				&asciidoc.NewLine{},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.ThematicBreak{
					AttributeList: nil,
				},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "term2:: def2",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var listsTestBlockTitleOffsetByBlankLineDividesListsAndBecomesTitleOfSecondListBecauseItemHasText = &asciidoc.Document{
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
					Value: "term1:: def1",
				},
				&asciidoc.NewLine{},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.Paragraph{
					AttributeList: asciidoc.AttributeList{
						&asciidoc.TitleAttribute{
							Val: asciidoc.Elements{
								&asciidoc.String{
									Value: "title",
								},
							},
						},
					},
					Elements:   asciidoc.Elements{},
					Admonition: 0,
				},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "term2:: def2",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var listsTestDoesNotRecognizeCalloutListDenotedByMarkersThatOnlyHaveATrailingBracket = &asciidoc.Document{
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
				"require 'asciidoctor' # <1>",
			},
		},
		&asciidoc.String{
			Value: "1> Not a callout list item",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestShouldNotHangIfObsoleteCalloutListIsFoundInsideListItem = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "foo",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "1> bar",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
	},
}

var listsTestShouldNotHangIfObsoleteCalloutListIsFoundInsideDlistItem = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Elements: asciidoc.Elements{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "1> bar",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Elements{
				&asciidoc.String{
					Value: "foo",
				},
			},
		},
	},
}

var listsTestShouldRecognizeAutoNumberdCalloutListInsideList = &asciidoc.Document{
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
				"require 'asciidoctor' # <1>",
			},
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "foo",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "<.> bar",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
	},
}

var listsTestListingBlockWithSequentialCalloutsFollowedByAdjacentCalloutList = &asciidoc.Document{
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
				"require 'asciidoctor' # <1>",
				"doc = Asciidoctor::Document.new('Hello, World!') # <2>",
				"puts doc.convert # <3>",
			},
		},
		&asciidoc.String{
			Value: "<1> Describe the first line",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "<2> Describe the second line",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "<3> Describe the third line",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestListingBlockWithSequentialCalloutsFollowedByNonAdjacentCalloutList = &asciidoc.Document{
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
				"require 'asciidoctor' # <1>",
				"doc = Asciidoctor::Document.new('Hello, World!') # <2>",
				"puts doc.convert # <3>",
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "Paragraph.",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "<1> Describe the first line",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "<2> Describe the second line",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "<3> Describe the third line",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestListingBlockWithACalloutThatRefersToTwoDifferentLines = &asciidoc.Document{
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
				"require 'asciidoctor' # <1>",
				"doc = Asciidoctor::Document.new('Hello, World!') # <2>",
				"puts doc.convert # <2>",
			},
		},
		&asciidoc.String{
			Value: "<1> Import the library",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "<2> Where the magic happens",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestSourceBlockWithNonSequentialCalloutsFollowedByAdjacentCalloutList = &asciidoc.Document{
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
				"require 'asciidoctor' # <2>",
				"doc = Asciidoctor::Document.new('Hello, World!') # <3>",
				"puts doc.convert # <1>",
			},
		},
		&asciidoc.String{
			Value: "<1> Describe the first line",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "<2> Describe the second line",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "<3> Describe the third line",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestTwoListingBlocksCanShareTheSameCalloutList = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Listing{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.TitleAttribute{
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Import library",
						},
					},
				},
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
				"require 'asciidoctor' # <1>",
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
							Value: "Use library",
						},
					},
				},
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
				"doc = Asciidoctor::Document.new('Hello, World!') # <2>",
				"puts doc.convert # <3>",
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "<1> Describe the first line",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "<2> Describe the second line",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "<3> Describe the third line",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestTwoListingBlocksEachFollowedByAnAdjacentCalloutList = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Listing{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.TitleAttribute{
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Import library",
						},
					},
				},
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
				"require 'asciidoctor' # <1>",
			},
		},
		&asciidoc.String{
			Value: "<1> Describe the first line",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Listing{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.TitleAttribute{
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Use library",
						},
					},
				},
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
				"doc = Asciidoctor::Document.new('Hello, World!') # <1>",
				"puts doc.convert # <2>",
			},
		},
		&asciidoc.String{
			Value: "<1> Describe the second line",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "<2> Describe the third line",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestCalloutListRetainsBlockContent = &asciidoc.Document{
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
				"require 'asciidoctor' # <1>",
				"doc = Asciidoctor::Document.new('Hello, World!') # <2>",
				"puts doc.convert # <3>",
			},
		},
		&asciidoc.String{
			Value: "<1> Imports the library",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "as a RubyGem",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "<2> Creates a new document",
		},
		&asciidoc.NewLine{},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Scans the lines for known blocks",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Converts the lines into blocks",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "<3> Renders the document",
				},
				&asciidoc.NewLine{},
				&asciidoc.LineBreak{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "You can write this to file rather than printing to stdout.",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
	},
}

var listsTestCalloutListRetainsBlockContentWhenConvertedToDocBook = &asciidoc.Document{
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
				"require 'asciidoctor' # <1>",
				"doc = Asciidoctor::Document.new('Hello, World!') # <2>",
				"puts doc.convert # <3>",
			},
		},
		&asciidoc.String{
			Value: "<1> Imports the library",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "as a RubyGem",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "<2> Creates a new document",
		},
		&asciidoc.NewLine{},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Scans the lines for known blocks",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Converts the lines into blocks",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "<3> Renders the document",
				},
				&asciidoc.NewLine{},
				&asciidoc.LineBreak{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "You can write this to file rather than printing to stdout.",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
	},
}

var listsTestEscapedCalloutShouldNotBeInterpretedAsACallout = &asciidoc.Document{
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
				&asciidoc.PositionalAttribute{
					Offset:      1,
					ImpliedName: "",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "text",
						},
					},
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   5,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"require 'asciidoctor' # \\<1>",
				"Asciidoctor.convert 'convert me!' \\<2>",
			},
		},
	},
}

var listsTestShouldAutonumberCallouts = &asciidoc.Document{
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
				"require 'asciidoctor' # <.>",
				"doc = Asciidoctor::Document.new('Hello, World!') # <.>",
				"puts doc.convert # <.>",
			},
		},
		&asciidoc.String{
			Value: "<.> Describe the first line",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "<.> Describe the second line",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "<.> Describe the third line",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestShouldNotRecognizeCalloutsInMiddleOfLine = &asciidoc.Document{
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
				"puts \"The syntax <1> at the end of the line makes a code callout\"",
			},
		},
	},
}

var listsTestShouldAllowMultipleCalloutsOnTheSameLine = &asciidoc.Document{
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
				"require 'asciidoctor' <1>",
				"doc = Asciidoctor.load('Hello, World!') # <2> <3> <4>",
				"puts doc.convert <5><6>",
				"exit 0",
			},
		},
		&asciidoc.String{
			Value: "<1> Require library",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "<2> Load document from String",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "<3> Uses default backend and doctype",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "<4> One more for good luck",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "<5> Renders document to String",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "<6> Prints output to stdout",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestShouldAllowXmlCommentStyleCallouts = &asciidoc.Document{
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
				&asciidoc.PositionalAttribute{
					Offset:      1,
					ImpliedName: "",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "xml",
						},
					},
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   5,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"<section>",
				"  <title>Section Title</title> <!--1-->",
				"  <simpara>Just a paragraph</simpara> <!--2-->",
				"</section>",
			},
		},
		&asciidoc.String{
			Value: "<1> The title is required",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "<2> The content isn't",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestShouldNotAllowCalloutsWithHalfAnXmlComment = &asciidoc.Document{
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
				"First line <1-->",
				"Second line <2-->",
			},
		},
	},
}

var listsTestShouldNotRecognizeCalloutsInAnIndentedDescriptionListParagraph = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Elements: asciidoc.Elements{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "  bar <1>",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Elements{
				&asciidoc.String{
					Value: "foo",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "<1> Not pointing to a callout",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestShouldNotRecognizeCalloutsInAnIndentedOutlineListParagraph = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "foo",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "  bar <1>",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "<1> Not pointing to a callout",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestShouldWarnIfNumbersInCalloutListAreOutOfSequence = &asciidoc.Document{
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
				"<beans> <1>",
				"  <bean class=\"com.example.HelloWorld\"/>",
				"</beans>",
			},
		},
		&asciidoc.String{
			Value: "<1> Container of beans.",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "Beans are fun.",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "<3> An actual bean.",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestShouldPreserveLineCommentCharsThatPrecedeCalloutNumberIfIconsIsNotSet = &asciidoc.Document{
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
				"puts 'Hello, world!' # <1>",
			},
		},
		&asciidoc.String{
			Value: "<1> Ruby",
		},
		&asciidoc.NewLine{},
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
							Value: "groovy",
						},
					},
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   5,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"println 'Hello, world!' // <1>",
			},
		},
		&asciidoc.String{
			Value: "<1> Groovy",
		},
		&asciidoc.NewLine{},
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
							Value: "clojure",
						},
					},
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   5,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"(def hello (fn [] \"Hello, world!\")) ;; <1>",
				"(hello)",
			},
		},
		&asciidoc.String{
			Value: "<1> Clojure",
		},
		&asciidoc.NewLine{},
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
							Value: "haskell",
						},
					},
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   5,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"main = putStrLn \"Hello, World!\" -- <1>",
			},
		},
		&asciidoc.String{
			Value: "<1> Haskell",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestShouldRemoveLineCommentCharsThatPrecedeCalloutNumberIfIconsIsFont = &asciidoc.Document{
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
				"puts 'Hello, world!' # <1>",
			},
		},
		&asciidoc.String{
			Value: "<1> Ruby",
		},
		&asciidoc.NewLine{},
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
							Value: "groovy",
						},
					},
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   5,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"println 'Hello, world!' // <1>",
			},
		},
		&asciidoc.String{
			Value: "<1> Groovy",
		},
		&asciidoc.NewLine{},
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
							Value: "clojure",
						},
					},
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   5,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"(def hello (fn [] \"Hello, world!\")) ;; <1>",
				"(hello)",
			},
		},
		&asciidoc.String{
			Value: "<1> Clojure",
		},
		&asciidoc.NewLine{},
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
							Value: "haskell",
						},
					},
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   5,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"main = putStrLn \"Hello, World!\" -- <1>",
			},
		},
		&asciidoc.String{
			Value: "<1> Haskell",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestShouldAllowLineCommentCharsThatPrecedeCalloutNumberToBeSpecified = &asciidoc.Document{
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
				&asciidoc.PositionalAttribute{
					Offset:      1,
					ImpliedName: "",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "erlang",
						},
					},
				},
				&asciidoc.NamedAttribute{
					Name: "line-comment",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "%",
						},
					},
					Quote: 0,
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   5,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"hello_world() -> % <1>",
				"  io:fwrite(\"hello, world~n\"). %<2>",
			},
		},
		&asciidoc.String{
			Value: "<1> Erlang function clause head.",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "<2> ~n adds a new line to the output.",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestShouldAllowLineCommentCharsPrecedingCalloutNumberToBeConfigurableWhenSourceHighlighterIsCoderay = &asciidoc.Document{
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
				&asciidoc.PositionalAttribute{
					Offset:      1,
					ImpliedName: "",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "html",
						},
					},
				},
				&asciidoc.NamedAttribute{
					Name: "line-comment",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "-#",
						},
					},
					Quote: 0,
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   5,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"-# <1>",
				"%p Hello",
			},
		},
		&asciidoc.String{
			Value: "<1> Prints a paragraph with the text \"Hello\"",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestShouldNotEatWhitespaceBeforeCalloutNumberIfLineCommentAttributeIsEmpty = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "[source,asciidoc,line-comment=]",
		},
		&asciidoc.NewLine{},
		&asciidoc.Listing{
			AttributeList: nil,
			Delimiter: asciidoc.Delimiter{
				Type:   5,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"-- <1>",
			},
		},
		&asciidoc.String{
			Value: "<1> The start of an open block.",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestLiteralBlockWithCallouts = &asciidoc.Document{
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
				"Roses are red <1>",
				"Violets are blue <2>",
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "<1> And so is Ruby",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "<2> But violet is more like purple",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestCalloutListWithIconsEnabled = &asciidoc.Document{
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
				"require 'asciidoctor' # <1>",
				"doc = Asciidoctor::Document.new('Hello, World!') # <2>",
				"puts doc.convert # <3>",
			},
		},
		&asciidoc.String{
			Value: "<1> Describe the first line",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "<2> Describe the second line",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "<3> Describe the third line",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestCalloutListWithFontBasedIconsEnabled = &asciidoc.Document{
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
				"require 'asciidoctor' # <1>",
				"doc = Asciidoctor::Document.new('Hello, World!') #<2>",
				"puts doc.convert #<3>",
			},
		},
		&asciidoc.String{
			Value: "<1> Describe the first line",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "<2> Describe the second line",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "<3> Describe the third line",
		},
		&asciidoc.NewLine{},
	},
}

var listsTestShouldCreateChecklistIfAtLeastOneItemHasCheckboxSyntax = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "todo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "-",
			Checklist:     1,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "done",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "-",
			Checklist:     2,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "another todo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "-",
			Checklist:     1,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "another done",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "-",
			Checklist:     2,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "plain",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "-",
			Checklist:     0,
		},
	},
}

var listsTestShouldCreateChecklistWithFontIconsIfAtLeastOneItemHasCheckboxSyntaxAndIconsAttributeIsFont = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "todo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "-",
			Checklist:     1,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "done",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "-",
			Checklist:     2,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "plain",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "-",
			Checklist:     0,
		},
	},
}

var listsTestShouldCreateInteractiveChecklistIfInteractiveOptionIsSetEvenWithIconsAttributeIsFont = &asciidoc.Document{
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
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "todo",
				},
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
									Value: "interactive",
								},
							},
						},
					},
				},
			},
			Indent:    "",
			Marker:    "-",
			Checklist: 1,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "done",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "-",
			Checklist:     2,
		},
	},
}

var listsTestContentShouldReturnItemsInList = &asciidoc.Document{
	Elements: asciidoc.Elements{
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
			Checklist:     0,
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
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "three",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
	},
}

var listsTestListItemShouldBeTheParentOfBlockAttachedToAListItem = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "list item 1",
				},
				&asciidoc.NewLine{},
				&asciidoc.LineBreak{},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		&asciidoc.Listing{
			AttributeList: nil,
			Delimiter: asciidoc.Delimiter{
				Type:   5,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"listing block in list item 1",
			},
		},
	},
}

var listsTestOutlineShouldReturnTrueForUnorderedList = &asciidoc.Document{
	Elements: asciidoc.Elements{
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
			Checklist:     0,
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
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "three",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
	},
}

var listsTestOutlineShouldReturnTrueForOrderedList = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "one",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "two",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "three",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
	},
}

var listsTestOutlineShouldReturnFalseForDescriptionList = &asciidoc.Document{
	Elements: asciidoc.Elements{
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
			Checklist:     0,
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
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "three",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
	},
}

var listsTestSimpleShouldReturnTrueForListItemWithNestedOutlineList = &asciidoc.Document{
	Elements: asciidoc.Elements{
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
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "more about one",
				},
			},
			AttributeList: nil,
			Indent:        "  ",
			Marker:        "**",
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "and more",
				},
			},
			AttributeList: nil,
			Indent:        "  ",
			Marker:        "**",
			Checklist:     0,
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
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "three",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
	},
}

var listsTestSimpleShouldReturnFalseForListItemWithBlockContent = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "one",
				},
				&asciidoc.NewLine{},
				&asciidoc.LineBreak{},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		&asciidoc.Listing{
			AttributeList: nil,
			Delimiter: asciidoc.Delimiter{
				Type:   5,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"listing block in list item 1",
			},
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
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "three",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
	},
}

var listsTestShouldAllowTextOfListItemToBeAssigned = &asciidoc.Document{
	Elements: asciidoc.Elements{
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
			Checklist:     0,
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
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "three",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
	},
}

var listsTestIdAndRoleAssignedToUlistItemInModelAreTransmittedToOutput = &asciidoc.Document{
	Elements: asciidoc.Elements{
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
			Checklist:     0,
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
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "three",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
	},
}

var listsTestIdAndRoleAssignedToOlistItemInModelAreTransmittedToOutput = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "one",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "two",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "three",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
	},
}

var listsTestShouldAllowApiControlOverSubstitutionsAppliedToListItemText = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.Bold{
					AttributeList: nil,
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "one",
						},
					},
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.Italic{
					AttributeList: nil,
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "two",
						},
					},
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.Monospace{
					AttributeList: nil,
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "three",
						},
					},
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.Marked{
					AttributeList: nil,
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "four",
						},
					},
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
	},
}

var listsTestShouldSetLinenoToLineNumberInSourceWhereListStarts = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "bullet 1",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "bullet 1.1",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "**",
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "bullet 1.1.1",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "***",
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "bullet 2",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
	},
}
