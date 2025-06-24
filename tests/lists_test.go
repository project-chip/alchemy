package tests

import (
	"testing"

	"github.com/project-chip/alchemy/asciidoc"
)

func TestLists(t *testing.T) {
	listsTests.run(t)
}

var listsTests = parseTests{

	{"dash elements with no blank lines", "asciidoctor/lists_test_dash_elements_with_no_blank_lines.adoc", dashElementsWithNoBlankLines, nil},

	{"dash elements separated by blank lines should merge lists", "asciidoctor/lists_test_dash_elements_separated_by_blank_lines_should_merge_lists.adoc", dashElementsSeparatedByBlankLinesShouldMergeLists, nil},

	{"dash elements with interspersed line comments should be skipped and not break list", "asciidoctor/lists_test_dash_elements_with_interspersed_line_comments_should_be_skipped_and_not_break_list.adoc", dashElementsWithInterspersedLineCommentsShouldBeSkippedAndNotBreakList, nil},

	{"dash elements separated by a line comment offset by blank lines should not merge lists", "asciidoctor/lists_test_dash_elements_separated_by_a_line_comment_offset_by_blank_lines_should_not_merge_lists.adoc", dashElementsSeparatedByALineCommentOffsetByBlankLinesShouldNotMergeLists, nil},

	{"dash elements separated by a block title offset by a blank line should not merge lists", "asciidoctor/lists_test_dash_elements_separated_by_a_block_title_offset_by_a_blank_line_should_not_merge_lists.adoc", dashElementsSeparatedByABlockTitleOffsetByABlankLineShouldNotMergeLists, nil},

	{"dash elements separated by an attribute entry offset by a blank line should not merge lists", "asciidoctor/lists_test_dash_elements_separated_by_an_attribute_entry_offset_by_a_blank_line_should_not_merge_lists.adoc", dashElementsSeparatedByAnAttributeEntryOffsetByABlankLineShouldNotMergeLists, nil},

	{"a non-indented wrapped line is folded into text of list item", "asciidoctor/lists_test_a_non_indented_wrapped_line_is_folded_into_text_of_list_item.adoc", aNonIndentedWrappedLineIsFoldedIntoTextOfListItem, nil},

	{"a non-indented wrapped line that resembles a block title is folded into text of list item", "asciidoctor/lists_test_a_non_indented_wrapped_line_that_resembles_a_block_title_is_folded_into_text_of_list_item.adoc", aNonIndentedWrappedLineThatResemblesABlockTitleIsFoldedIntoTextOfListItem, nil},

	{"a non-indented wrapped line that resembles an attribute entry is folded into text of list item", "asciidoctor/lists_test_a_non_indented_wrapped_line_that_resembles_an_attribute_entry_is_folded_into_text_of_list_item.adoc", aNonIndentedWrappedLineThatResemblesAnAttributeEntryIsFoldedIntoTextOfListItem, nil},

	{"a list item with a nested marker terminates non-indented paragraph for text of list item", "asciidoctor/lists_test_a_list_item_with_a_nested_marker_terminates_non_indented_paragraph_for_text_of_list_item.adoc", aListItemWithANestedMarkerTerminatesNonIndentedParagraphForTextOfListItem, nil},

	{"a list item for a different list terminates non-indented paragraph for text of list item", "asciidoctor/lists_test_a_list_item_for_a_different_list_terminates_non_indented_paragraph_for_text_of_list_item.adoc", aListItemForADifferentListTerminatesNonIndentedParagraphForTextOfListItem, nil},

	{"an indented wrapped line is unindented and folded into text of list item", "asciidoctor/lists_test_an_indented_wrapped_line_is_unindented_and_folded_into_text_of_list_item.adoc", anIndentedWrappedLineIsUnindentedAndFoldedIntoTextOfListItem, nil},

	{"wrapped list item with hanging indent followed by non-indented line", "asciidoctor/lists_test_wrapped_list_item_with_hanging_indent_followed_by_non_indented_line.adoc", wrappedListItemWithHangingIndentFollowedByNonIndentedLine, nil},

	{"a list item with a nested marker terminates indented paragraph for text of list item", "asciidoctor/lists_test_a_list_item_with_a_nested_marker_terminates_indented_paragraph_for_text_of_list_item.adoc", aListItemWithANestedMarkerTerminatesIndentedParagraphForTextOfListItem, nil},

	{"a list item for a different list terminates indented paragraph for text of list item", "asciidoctor/lists_test_a_list_item_for_a_different_list_terminates_indented_paragraph_for_text_of_list_item.adoc", aListItemForADifferentListTerminatesIndentedParagraphForTextOfListItem, nil},

	{"a literal paragraph offset by blank lines in list content is appended as a literal block", "asciidoctor/lists_test_a_literal_paragraph_offset_by_blank_lines_in_list_content_is_appended_as_a_literal_block.adoc", aLiteralParagraphOffsetByBlankLinesInListContentIsAppendedAsALiteralBlock, nil},

	{"should escape special characters in all literal paragraphs attached to list item", "asciidoctor/lists_test_should_escape_special_characters_in_all_literal_paragraphs_attached_to_list_item.adoc", shouldEscapeSpecialCharactersInAllLiteralParagraphsAttachedToListItem, nil},

	{"a literal paragraph offset by a blank line in list content followed by line with continuation is appended as two blocks", "asciidoctor/lists_test_a_literal_paragraph_offset_by_a_blank_line_in_list_content_followed_by_line_with_continuation_is_appended_as_two_blocks.adoc", aLiteralParagraphOffsetByABlankLineInListContentFollowedByLineWithContinuationIsAppendedAsTwoBlocks, nil},

	{"an admonition paragraph attached by a line continuation to a list item with wrapped text should produce admonition", "asciidoctor/lists_test_an_admonition_paragraph_attached_by_a_line_continuation_to_a_list_item_with_wrapped_text_should_produce_admonition.adoc", anAdmonitionParagraphAttachedByALineContinuationToAListItemWithWrappedTextShouldProduceAdmonition, nil},

	{"paragraph-like blocks attached to an ancestor list item by a list continuation should produce blocks", "asciidoctor/lists_test_paragraph_like_blocks_attached_to_an_ancestor_list_item_by_a_list_continuation_should_produce_blocks.adoc", paragraphLikeBlocksAttachedToAnAncestorListItemByAListContinuationShouldProduceBlocks, nil},

	{"should not inherit block attributes from previous block when block is attached using a list continuation", "asciidoctor/lists_test_should_not_inherit_block_attributes_from_previous_block_when_block_is_attached_using_a_list_continuation.adoc", shouldNotInheritBlockAttributesFromPreviousBlockWhenBlockIsAttachedUsingAListContinuation, nil},

	{"should continue to parse blocks attached by a list continuation after block is dropped", "asciidoctor/lists_test_should_continue_to_parse_blocks_attached_by_a_list_continuation_after_block_is_dropped.adoc", shouldContinueToParseBlocksAttachedByAListContinuationAfterBlockIsDropped, nil},

	{"appends line as paragraph if attached by continuation following line comment", "asciidoctor/lists_test_appends_line_as_paragraph_if_attached_by_continuation_following_line_comment.adoc", appendsLineAsParagraphIfAttachedByContinuationFollowingLineComment, nil},

	{"a literal paragraph with a line that appears as a list item that is followed by a continuation should create two blocks", "asciidoctor/lists_test_a_literal_paragraph_with_a_line_that_appears_as_a_list_item_that_is_followed_by_a_continuation_should_create_two_blocks.adoc", aLiteralParagraphWithALineThatAppearsAsAListItemThatIsFollowedByAContinuationShouldCreateTwoBlocks, nil},

	{"consecutive literal paragraph offset by blank lines in list content are appended as a literal blocks", "asciidoctor/lists_test_consecutive_literal_paragraph_offset_by_blank_lines_in_list_content_are_appended_as_a_literal_blocks.adoc", consecutiveLiteralParagraphOffsetByBlankLinesInListContentAreAppendedAsALiteralBlocks, nil},

	{"a literal paragraph without a trailing blank line consumes following list items", "asciidoctor/lists_test_a_literal_paragraph_without_a_trailing_blank_line_consumes_following_list_items.adoc", aLiteralParagraphWithoutATrailingBlankLineConsumesFollowingListItems, nil},

	{"asterisk elements with no blank lines", "asciidoctor/lists_test_asterisk_elements_with_no_blank_lines.adoc", asteriskElementsWithNoBlankLines, nil},

	{"asterisk elements separated by blank lines should merge lists", "asciidoctor/lists_test_asterisk_elements_separated_by_blank_lines_should_merge_lists.adoc", asteriskElementsSeparatedByBlankLinesShouldMergeLists, nil},

	{"asterisk elements with interspersed line comments should be skipped and not break list", "asciidoctor/lists_test_asterisk_elements_with_interspersed_line_comments_should_be_skipped_and_not_break_list.adoc", asteriskElementsWithInterspersedLineCommentsShouldBeSkippedAndNotBreakList, nil},

	{"asterisk elements separated by a line comment offset by blank lines should not merge lists", "asciidoctor/lists_test_asterisk_elements_separated_by_a_line_comment_offset_by_blank_lines_should_not_merge_lists.adoc", asteriskElementsSeparatedByALineCommentOffsetByBlankLinesShouldNotMergeLists, nil},

	{"asterisk elements separated by a block title offset by a blank line should not merge lists", "asciidoctor/lists_test_asterisk_elements_separated_by_a_block_title_offset_by_a_blank_line_should_not_merge_lists.adoc", asteriskElementsSeparatedByABlockTitleOffsetByABlankLineShouldNotMergeLists, nil},

	{"asterisk elements separated by an attribute entry offset by a blank line should not merge lists", "asciidoctor/lists_test_asterisk_elements_separated_by_an_attribute_entry_offset_by_a_blank_line_should_not_merge_lists.adoc", asteriskElementsSeparatedByAnAttributeEntryOffsetByABlankLineShouldNotMergeLists, nil},

	{"list should terminate before next lower section heading", "asciidoctor/lists_test_list_should_terminate_before_next_lower_section_heading.adoc", listShouldTerminateBeforeNextLowerSectionHeading, nil},

	{"list should terminate before next lower section heading with implicit id", "asciidoctor/lists_test_list_should_terminate_before_next_lower_section_heading_with_implicit_id.adoc", listShouldTerminateBeforeNextLowerSectionHeadingWithImplicitId, nil},

	{"should not find section title immediately below last list item", "asciidoctor/lists_test_should_not_find_section_title_immediately_below_last_list_item.adoc", shouldNotFindSectionTitleImmediatelyBelowLastListItem, nil},

	{"quoted text", "asciidoctor/lists_test_quoted_text.adoc", quotedText, nil},

	{"attribute substitutions", "asciidoctor/lists_test_attribute_substitutions.adoc", attributeSubstitutions, nil},

	{"leading dot is treated as text not block title", "asciidoctor/lists_test_leading_dot_is_treated_as_text_not_block_title.adoc", leadingDotIsTreatedAsTextNotBlockTitle, nil},

	{"word ending sentence on continuing line not treated as a list item", "asciidoctor/lists_test_word_ending_sentence_on_continuing_line_not_treated_as_a_list_item.adoc", wordEndingSentenceOnContinuingLineNotTreatedAsAListItem, nil},

	{"should discover anchor at start of unordered list item text and register it as a reference", "asciidoctor/lists_test_should_discover_anchor_at_start_of_unordered_list_item_text_and_register_it_as_a_reference.adoc", shouldDiscoverAnchorAtStartOfUnorderedListItemTextAndRegisterItAsAReference, nil},

	{"should discover anchor at start of ordered list item text and register it as a reference", "asciidoctor/lists_test_should_discover_anchor_at_start_of_ordered_list_item_text_and_register_it_as_a_reference.adoc", shouldDiscoverAnchorAtStartOfOrderedListItemTextAndRegisterItAsAReference, nil},

	{"should discover anchor at start of callout list item text and register it as a reference", "asciidoctor/lists_test_should_discover_anchor_at_start_of_callout_list_item_text_and_register_it_as_a_reference.adoc", shouldDiscoverAnchorAtStartOfCalloutListItemTextAndRegisterItAsAReference, nil},

	{"asterisk element mixed with dash elements should be nested", "asciidoctor/lists_test_asterisk_element_mixed_with_dash_elements_should_be_nested.adoc", asteriskElementMixedWithDashElementsShouldBeNested, nil},

	{"dash element mixed with asterisks elements should be nested", "asciidoctor/lists_test_dash_element_mixed_with_asterisks_elements_should_be_nested.adoc", dashElementMixedWithAsterisksElementsShouldBeNested, nil},

	{"lines prefixed with alternating list markers separated by blank lines should be nested", "asciidoctor/lists_test_lines_prefixed_with_alternating_list_markers_separated_by_blank_lines_should_be_nested.adoc", linesPrefixedWithAlternatingListMarkersSeparatedByBlankLinesShouldBeNested, nil},

	{"nested elements (2) with asterisks", "asciidoctor/lists_test_nested_elements_(2)_with_asterisks.adoc", nestedElements2WithAsterisks, nil},

	{"nested elements (3) with asterisks", "asciidoctor/lists_test_nested_elements_(3)_with_asterisks.adoc", nestedElements3WithAsterisks, nil},

	{"nested elements (4) with asterisks", "asciidoctor/lists_test_nested_elements_(4)_with_asterisks.adoc", nestedElements4WithAsterisks, nil},

	{"nested elements (5) with asterisks", "asciidoctor/lists_test_nested_elements_(5)_with_asterisks.adoc", nestedElements5WithAsterisks, nil},

	{"level of unordered list should match section level", "asciidoctor/lists_test_level_of_unordered_list_should_match_section_level.adoc", levelOfUnorderedListShouldMatchSectionLevel, nil},

	{"does not recognize lists with repeating unicode bullets", "asciidoctor/lists_test_does_not_recognize_lists_with_repeating_unicode_bullets.adoc", doesNotRecognizeListsWithRepeatingUnicodeBullets, nil},

	{"nested ordered elements (3)", "asciidoctor/lists_test_nested_ordered_elements_(3).adoc", nestedOrderedElements3, nil},

	{"level of ordered list should match section level", "asciidoctor/lists_test_level_of_ordered_list_should_match_section_level.adoc", levelOfOrderedListShouldMatchSectionLevel, nil},

	{"nested unordered inside ordered elements", "asciidoctor/lists_test_nested_unordered_inside_ordered_elements.adoc", nestedUnorderedInsideOrderedElements, nil},

	{"nested ordered inside unordered elements", "asciidoctor/lists_test_nested_ordered_inside_unordered_elements.adoc", nestedOrderedInsideUnorderedElements, nil},

	{"three levels of alternating unordered and ordered elements", "asciidoctor/lists_test_three_levels_of_alternating_unordered_and_ordered_elements.adoc", threeLevelsOfAlternatingUnorderedAndOrderedElements, nil},

	{"lines with alternating markers of unordered and ordered list types separated by blank lines should be nested", "asciidoctor/lists_test_lines_with_alternating_markers_of_unordered_and_ordered_list_types_separated_by_blank_lines_should_be_nested.adoc", linesWithAlternatingMarkersOfUnorderedAndOrderedListTypesSeparatedByBlankLinesShouldBeNested, nil},

	{"list item with literal content should not consume nested list of different type", "asciidoctor/lists_test_list_item_with_literal_content_should_not_consume_nested_list_of_different_type.adoc", listItemWithLiteralContentShouldNotConsumeNestedListOfDifferentType, nil},

	{"nested list item does not eat the title of the following detached block", "asciidoctor/lists_test_nested_list_item_does_not_eat_the_title_of_the_following_detached_block.adoc", nestedListItemDoesNotEatTheTitleOfTheFollowingDetachedBlock, nil},

	{"lines with alternating markers of bulleted and description list types separated by blank lines should be nested", "asciidoctor/lists_test_lines_with_alternating_markers_of_bulleted_and_description_list_types_separated_by_blank_lines_should_be_nested.adoc", linesWithAlternatingMarkersOfBulletedAndDescriptionListTypesSeparatedByBlankLinesShouldBeNested, nil},

	{"nested ordered with attribute inside unordered elements", "asciidoctor/lists_test_nested_ordered_with_attribute_inside_unordered_elements.adoc", nestedOrderedWithAttributeInsideUnorderedElements, nil},

	{"adjacent list continuation line attaches following paragraph", "asciidoctor/lists_test_adjacent_list_continuation_line_attaches_following_paragraph.adoc", adjacentListContinuationLineAttachesFollowingParagraph, nil},

	{"adjacent list continuation line attaches following block", "asciidoctor/lists_test_adjacent_list_continuation_line_attaches_following_block.adoc", adjacentListContinuationLineAttachesFollowingBlock, nil},

	{"adjacent list continuation line attaches following block with block attributes", "asciidoctor/lists_test_adjacent_list_continuation_line_attaches_following_block_with_block_attributes.adoc", adjacentListContinuationLineAttachesFollowingBlockWithBlockAttributes, nil},

	{"trailing block attribute line attached by continuation should not create block", "asciidoctor/lists_test_trailing_block_attribute_line_attached_by_continuation_should_not_create_block.adoc", trailingBlockAttributeLineAttachedByContinuationShouldNotCreateBlock, nil},

	{"trailing block title line attached by continuation should not create block", "asciidoctor/lists_test_trailing_block_title_line_attached_by_continuation_should_not_create_block.adoc", trailingBlockTitleLineAttachedByContinuationShouldNotCreateBlock, nil},

	{"consecutive blocks in list continuation attach to list item", "asciidoctor/lists_test_consecutive_blocks_in_list_continuation_attach_to_list_item.adoc", consecutiveBlocksInListContinuationAttachToListItem, nil},

	{"list item with hanging indent followed by block attached by list continuation", "asciidoctor/lists_test_list_item_with_hanging_indent_followed_by_block_attached_by_list_continuation.adoc", listItemWithHangingIndentFollowedByBlockAttachedByListContinuation, nil},

	{"list item paragraph in list item and nested list item", "asciidoctor/lists_test_list_item_paragraph_in_list_item_and_nested_list_item.adoc", listItemParagraphInListItemAndNestedListItem, nil},

	{"trailing list continuations should attach to list items at respective levels", "asciidoctor/lists_test_trailing_list_continuations_should_attach_to_list_items_at_respective_levels.adoc", trailingListContinuationsShouldAttachToListItemsAtRespectiveLevels, nil},

	{"trailing list continuations should attach to list items of different types at respective levels", "asciidoctor/lists_test_trailing_list_continuations_should_attach_to_list_items_of_different_types_at_respective_levels.adoc", trailingListContinuationsShouldAttachToListItemsOfDifferentTypesAtRespectiveLevels, nil},

	{"repeated list continuations should attach to list items at respective levels", "asciidoctor/lists_test_repeated_list_continuations_should_attach_to_list_items_at_respective_levels.adoc", repeatedListContinuationsShouldAttachToListItemsAtRespectiveLevels, nil},

	{"repeated list continuations attached directly to list item should attach to list items at respective levels", "asciidoctor/lists_test_repeated_list_continuations_attached_directly_to_list_item_should_attach_to_list_items_at_respective_levels.adoc", repeatedListContinuationsAttachedDirectlyToListItemShouldAttachToListItemsAtRespectiveLevels, nil},

	{"repeated list continuations should attach to list items at respective levels ignoring blank lines", "asciidoctor/lists_test_repeated_list_continuations_should_attach_to_list_items_at_respective_levels_ignoring_blank_lines.adoc", repeatedListContinuationsShouldAttachToListItemsAtRespectiveLevelsIgnoringBlankLines, nil},

	{"trailing list continuations should ignore preceding blank lines", "asciidoctor/lists_test_trailing_list_continuations_should_ignore_preceding_blank_lines.adoc", trailingListContinuationsShouldIgnorePrecedingBlankLines, nil},

	{"indented outline list item with different marker offset by a blank line should be recognized as a nested list", "asciidoctor/lists_test_indented_outline_list_item_with_different_marker_offset_by_a_blank_line_should_be_recognized_as_a_nested_list.adoc", indentedOutlineListItemWithDifferentMarkerOffsetByABlankLineShouldBeRecognizedAsANestedList, nil},

	{"indented description list item inside outline list item offset by a blank line should be recognized as a nested list", "asciidoctor/lists_test_indented_description_list_item_inside_outline_list_item_offset_by_a_blank_line_should_be_recognized_as_a_nested_list.adoc", indentedDescriptionListItemInsideOutlineListItemOffsetByABlankLineShouldBeRecognizedAsANestedList, nil},

	{"consecutive list continuation lines are folded", "asciidoctor/lists_test_consecutive_list_continuation_lines_are_folded.adoc", consecutiveListContinuationLinesAreFolded, nil},

	{"should warn if unterminated block is detected in list item", "asciidoctor/lists_test_should_warn_if_unterminated_block_is_detected_in_list_item.adoc", shouldWarnIfUnterminatedBlockIsDetectedInListItem, nil},

	{"dot elements with no blank lines", "asciidoctor/lists_test_dot_elements_with_no_blank_lines.adoc", dotElementsWithNoBlankLines, nil},

	{"should represent explicit role attribute as style class", "asciidoctor/lists_test_should_represent_explicit_role_attribute_as_style_class.adoc", shouldRepresentExplicitRoleAttributeAsStyleClass, nil},

	{"should base list style on marker length rather than list depth", "asciidoctor/lists_test_should_base_list_style_on_marker_length_rather_than_list_depth.adoc", shouldBaseListStyleOnMarkerLengthRatherThanListDepth, nil},

	{"should allow list style to be specified explicitly when using markers with implicit style", "asciidoctor/lists_test_should_allow_list_style_to_be_specified_explicitly_when_using_markers_with_implicit_style.adoc", shouldAllowListStyleToBeSpecifiedExplicitlyWhenUsingMarkersWithImplicitStyle, nil},

	{"should represent custom numbering and explicit role attribute as style classes", "asciidoctor/lists_test_should_represent_custom_numbering_and_explicit_role_attribute_as_style_classes.adoc", shouldRepresentCustomNumberingAndExplicitRoleAttributeAsStyleClasses, nil},

	{"should set reversed attribute on list if reversed option is set", "asciidoctor/lists_test_should_set_reversed_attribute_on_list_if_reversed_option_is_set.adoc", shouldSetReversedAttributeOnListIfReversedOptionIsSet, nil},

	{"should represent implicit role attribute as style class", "asciidoctor/lists_test_should_represent_implicit_role_attribute_as_style_class.adoc", shouldRepresentImplicitRoleAttributeAsStyleClass, nil},

	{"should represent custom numbering and implicit role attribute as style classes", "asciidoctor/lists_test_should_represent_custom_numbering_and_implicit_role_attribute_as_style_classes.adoc", shouldRepresentCustomNumberingAndImplicitRoleAttributeAsStyleClasses, nil},

	{"dot elements separated by blank lines should merge lists", "asciidoctor/lists_test_dot_elements_separated_by_blank_lines_should_merge_lists.adoc", dotElementsSeparatedByBlankLinesShouldMergeLists, nil},

	{"dot elements with interspersed line comments should be skipped and not break list", "asciidoctor/lists_test_dot_elements_with_interspersed_line_comments_should_be_skipped_and_not_break_list.adoc", dotElementsWithInterspersedLineCommentsShouldBeSkippedAndNotBreakList, nil},

	{"dot elements separated by line comment offset by blank lines should not merge lists", "asciidoctor/lists_test_dot_elements_separated_by_line_comment_offset_by_blank_lines_should_not_merge_lists.adoc", dotElementsSeparatedByLineCommentOffsetByBlankLinesShouldNotMergeLists, nil},

	{"dot elements separated by a block title offset by a blank line should not merge lists", "asciidoctor/lists_test_dot_elements_separated_by_a_block_title_offset_by_a_blank_line_should_not_merge_lists.adoc", dotElementsSeparatedByABlockTitleOffsetByABlankLineShouldNotMergeLists, nil},

	{"dot elements separated by an attribute entry offset by a blank line should not merge lists", "asciidoctor/lists_test_dot_elements_separated_by_an_attribute_entry_offset_by_a_blank_line_should_not_merge_lists.adoc", dotElementsSeparatedByAnAttributeEntryOffsetByABlankLineShouldNotMergeLists, nil},

	{"should use start number in docbook5 backend", "asciidoctor/lists_test_should_use_start_number_in_docbook_5_backend.adoc", shouldUseStartNumberInDocbook5Backend, nil},

	{"should warn if explicit uppercase roman numerals in list are out of sequence", "asciidoctor/lists_test_should_warn_if_explicit_uppercase_roman_numerals_in_list_are_out_of_sequence.adoc", shouldWarnIfExplicitUppercaseRomanNumeralsInListAreOutOfSequence, nil},

	{"should warn if explicit lowercase roman numerals in list are out of sequence", "asciidoctor/lists_test_should_warn_if_explicit_lowercase_roman_numerals_in_list_are_out_of_sequence.adoc", shouldWarnIfExplicitLowercaseRomanNumeralsInListAreOutOfSequence, nil},

	{"should not parse a bare dlist delimiter as a dlist", "asciidoctor/lists_test_should_not_parse_a_bare_dlist_delimiter_as_a_dlist.adoc", shouldNotParseABareDlistDelimiterAsADlist, nil},

	{"should parse sibling items using same rules", "asciidoctor/lists_test_should_parse_sibling_items_using_same_rules.adoc", shouldParseSiblingItemsUsingSameRules, nil},

	{"should allow term to end with a semicolon when using double semicolon delimiter", "asciidoctor/lists_test_should_allow_term_to_end_with_a_semicolon_when_using_double_semicolon_delimiter.adoc", shouldAllowTermToEndWithASemicolonWhenUsingDoubleSemicolonDelimiter, nil},

	{"single-line indented adjacent elements", "asciidoctor/lists_test_single_line_indented_adjacent_elements.adoc", singleLineIndentedAdjacentElements, nil},

	{"single-line elements separated by blank line should create a single list", "asciidoctor/lists_test_single_line_elements_separated_by_blank_line_should_create_a_single_list.adoc", singleLineElementsSeparatedByBlankLineShouldCreateASingleList, nil},

	{"a line comment between elements should divide them into separate lists", "asciidoctor/lists_test_a_line_comment_between_elements_should_divide_them_into_separate_lists.adoc", aLineCommentBetweenElementsShouldDivideThemIntoSeparateLists, nil},

	{"a ruler between elements should divide them into separate lists", "asciidoctor/lists_test_a_ruler_between_elements_should_divide_them_into_separate_lists.adoc", aRulerBetweenElementsShouldDivideThemIntoSeparateLists, nil},

	{"a block title between elements should divide them into separate lists", "asciidoctor/lists_test_a_block_title_between_elements_should_divide_them_into_separate_lists.adoc", aBlockTitleBetweenElementsShouldDivideThemIntoSeparateLists, nil},

	{"multi-line elements with paragraph content", "asciidoctor/lists_test_multi_line_elements_with_paragraph_content.adoc", multiLineElementsWithParagraphContent, nil},

	{"multi-line elements with indented paragraph content", "asciidoctor/lists_test_multi_line_elements_with_indented_paragraph_content.adoc", multiLineElementsWithIndentedParagraphContent, nil},

	{"multi-line elements with indented paragraph content that includes comment lines", "asciidoctor/lists_test_multi_line_elements_with_indented_paragraph_content_that_includes_comment_lines.adoc", multiLineElementsWithIndentedParagraphContentThatIncludesCommentLines, nil},

	{"should not strip comment line in literal paragraph block attached to list item", "asciidoctor/lists_test_should_not_strip_comment_line_in_literal_paragraph_block_attached_to_list_item.adoc", shouldNotStripCommentLineInLiteralParagraphBlockAttachedToListItem, nil},

	{"multi-line element with paragraph starting with multiple dashes should not be seen as list", "asciidoctor/lists_test_multi_line_element_with_paragraph_starting_with_multiple_dashes_should_not_be_seen_as_list.adoc", multiLineElementWithParagraphStartingWithMultipleDashesShouldNotBeSeenAsList, nil},

	{"multi-line element with multiple terms", "asciidoctor/lists_test_multi_line_element_with_multiple_terms.adoc", multiLineElementWithMultipleTerms, nil},

	{"consecutive terms share same varlistentry in docbook", "asciidoctor/lists_test_consecutive_terms_share_same_varlistentry_in_docbook.adoc", consecutiveTermsShareSameVarlistentryInDocbook, nil},

	{"multi-line elements with blank line before paragraph content", "asciidoctor/lists_test_multi_line_elements_with_blank_line_before_paragraph_content.adoc", multiLineElementsWithBlankLineBeforeParagraphContent, nil},

	{"multi-line elements with paragraph and literal content", "asciidoctor/lists_test_multi_line_elements_with_paragraph_and_literal_content.adoc", multiLineElementsWithParagraphAndLiteralContent, nil},

	{"mixed single and multi-line adjacent elements", "asciidoctor/lists_test_mixed_single_and_multi_line_adjacent_elements.adoc", mixedSingleAndMultiLineAdjacentElements, nil},

	{"should discover anchor at start of description term text and register it as a reference", "asciidoctor/lists_test_should_discover_anchor_at_start_of_description_term_text_and_register_it_as_a_reference.adoc", shouldDiscoverAnchorAtStartOfDescriptionTermTextAndRegisterItAsAReference, nil},

	{"missing space before term does not produce description list", "asciidoctor/lists_test_missing_space_before_term_does_not_produce_description_list.adoc", missingSpaceBeforeTermDoesNotProduceDescriptionList, nil},

	{"literal block inside description list", "asciidoctor/lists_test_literal_block_inside_description_list.adoc", literalBlockInsideDescriptionList, nil},

	{"literal block inside description list with trailing line continuation", "asciidoctor/lists_test_literal_block_inside_description_list_with_trailing_line_continuation.adoc", literalBlockInsideDescriptionListWithTrailingLineContinuation, nil},

	{"multiple listing blocks inside description list", "asciidoctor/lists_test_multiple_listing_blocks_inside_description_list.adoc", multipleListingBlocksInsideDescriptionList, nil},

	{"open block inside description list", "asciidoctor/lists_test_open_block_inside_description_list.adoc", openBlockInsideDescriptionList, nil},

	{"paragraph attached by a list continuation on either side in a description list", "asciidoctor/lists_test_paragraph_attached_by_a_list_continuation_on_either_side_in_a_description_list.adoc", paragraphAttachedByAListContinuationOnEitherSideInADescriptionList, nil},

	{"paragraph attached by a list continuation on either side to a multi-line element in a description list", "asciidoctor/lists_test_paragraph_attached_by_a_list_continuation_on_either_side_to_a_multi_line_element_in_a_description_list.adoc", paragraphAttachedByAListContinuationOnEitherSideToAMultiLineElementInADescriptionList, nil},

	{"should continue to parse subsequent blocks attached to list item after first block is dropped", "asciidoctor/lists_test_should_continue_to_parse_subsequent_blocks_attached_to_list_item_after_first_block_is_dropped.adoc", shouldContinueToParseSubsequentBlocksAttachedToListItemAfterFirstBlockIsDropped, nil},

	{"verse paragraph inside a description list", "asciidoctor/lists_test_verse_paragraph_inside_a_description_list.adoc", verseParagraphInsideADescriptionList, nil},

	{"list inside a description list", "asciidoctor/lists_test_list_inside_a_description_list.adoc", listInsideADescriptionList, nil},

	{"list inside a description list offset by blank lines", "asciidoctor/lists_test_list_inside_a_description_list_offset_by_blank_lines.adoc", listInsideADescriptionListOffsetByBlankLines, nil},

	{"should only grab one line following last item if item has no inline description", "asciidoctor/lists_test_should_only_grab_one_line_following_last_item_if_item_has_no_inline_description.adoc", shouldOnlyGrabOneLineFollowingLastItemIfItemHasNoInlineDescription, nil},

	{"should only grab one literal line following last item if item has no inline description", "asciidoctor/lists_test_should_only_grab_one_literal_line_following_last_item_if_item_has_no_inline_description.adoc", shouldOnlyGrabOneLiteralLineFollowingLastItemIfItemHasNoInlineDescription, nil},

	{"should append subsequent paragraph literals to list item as block content", "asciidoctor/lists_test_should_append_subsequent_paragraph_literals_to_list_item_as_block_content.adoc", shouldAppendSubsequentParagraphLiteralsToListItemAsBlockContent, nil},

	{"should not match comment line that looks like description list term", "asciidoctor/lists_test_should_not_match_comment_line_that_looks_like_description_list_term.adoc", shouldNotMatchCommentLineThatLooksLikeDescriptionListTerm, nil},

	{"should not match comment line following list that looks like description list term", "asciidoctor/lists_test_should_not_match_comment_line_following_list_that_looks_like_description_list_term.adoc", shouldNotMatchCommentLineFollowingListThatLooksLikeDescriptionListTerm, nil},

	{"should not match comment line that looks like sibling description list term", "asciidoctor/lists_test_should_not_match_comment_line_that_looks_like_sibling_description_list_term.adoc", shouldNotMatchCommentLineThatLooksLikeSiblingDescriptionListTerm, nil},

	{"should not hang on description list item in list that begins with ///", "asciidoctor/lists_test_should_not_hang_on_description_list_item_in_list_that_begins_with.adoc", shouldNotHangOnDescriptionListItemInListThatBeginsWith, nil},

	{"should not hang on sibling description list item that begins with ///", "asciidoctor/lists_test_should_not_hang_on_sibling_description_list_item_that_begins_with.adoc", shouldNotHangOnSiblingDescriptionListItemThatBeginsWith, nil},

	{"should skip dlist term that begins with // unless it begins with ///", "asciidoctor/lists_test_should_skip_dlist_term_that_begins_with____unless_it_begins_with.adoc", shouldSkipDlistTermThatBeginsWithUnlessItBeginsWith, nil},

	{"more than 4 consecutive colons should become part of description list term", "asciidoctor/lists_test_more_than_4_consecutive_colons_should_become_part_of_description_list_term.adoc", moreThan4ConsecutiveColonsShouldBecomePartOfDescriptionListTerm, nil},

	{"text method of dd node should return nil if dd node only contains blocks", "asciidoctor/lists_test_text_method_of_dd_node_should_return_nil_if_dd_node_only_contains_blocks.adoc", textMethodOfDdNodeShouldReturnNilIfDdNodeOnlyContainsBlocks, nil},

	{"should not parse a nested dlist delimiter without a term as a dlist", "asciidoctor/lists_test_should_not_parse_a_nested_dlist_delimiter_without_a_term_as_a_dlist.adoc", shouldNotParseANestedDlistDelimiterWithoutATermAsADlist, nil},

	{"should not parse a nested indented dlist delimiter without a term as a dlist", "asciidoctor/lists_test_should_not_parse_a_nested_indented_dlist_delimiter_without_a_term_as_a_dlist.adoc", shouldNotParseANestedIndentedDlistDelimiterWithoutATermAsADlist, nil},

	{"single-line adjacent nested elements", "asciidoctor/lists_test_single_line_adjacent_nested_elements.adoc", singleLineAdjacentNestedElements, nil},

	{"single-line adjacent maximum nested elements", "asciidoctor/lists_test_single_line_adjacent_maximum_nested_elements.adoc", singleLineAdjacentMaximumNestedElements, nil},

	{"single-line nested elements separated by blank line at top level", "asciidoctor/lists_test_single_line_nested_elements_separated_by_blank_line_at_top_level.adoc", singleLineNestedElementsSeparatedByBlankLineAtTopLevel, nil},

	{"single-line nested elements separated by blank line at nested level", "asciidoctor/lists_test_single_line_nested_elements_separated_by_blank_line_at_nested_level.adoc", singleLineNestedElementsSeparatedByBlankLineAtNestedLevel, nil},

	{"single-line adjacent nested elements with alternate delimiters", "asciidoctor/lists_test_single_line_adjacent_nested_elements_with_alternate_delimiters.adoc", singleLineAdjacentNestedElementsWithAlternateDelimiters, nil},

	{"multi-line adjacent nested elements", "asciidoctor/lists_test_multi_line_adjacent_nested_elements.adoc", multiLineAdjacentNestedElements, nil},

	{"multi-line nested elements separated by blank line at nested level repeated", "asciidoctor/lists_test_multi_line_nested_elements_separated_by_blank_line_at_nested_level_repeated.adoc", multiLineNestedElementsSeparatedByBlankLineAtNestedLevelRepeated, nil},

	{"multi-line element with indented nested element", "asciidoctor/lists_test_multi_line_element_with_indented_nested_element.adoc", multiLineElementWithIndentedNestedElement, nil},

	{"mixed single and multi-line elements with indented nested elements", "asciidoctor/lists_test_mixed_single_and_multi_line_elements_with_indented_nested_elements.adoc", mixedSingleAndMultiLineElementsWithIndentedNestedElements, nil},

	{"multi-line elements with first paragraph folded to text with adjacent nested element", "asciidoctor/lists_test_multi_line_elements_with_first_paragraph_folded_to_text_with_adjacent_nested_element.adoc", multiLineElementsWithFirstParagraphFoldedToTextWithAdjacentNestedElement, nil},

	{"nested dlist attached by list continuation should not consume detached paragraph", "asciidoctor/lists_test_nested_dlist_attached_by_list_continuation_should_not_consume_detached_paragraph.adoc", nestedDlistAttachedByListContinuationShouldNotConsumeDetachedParagraph, nil},

	{"nested dlist with attached block offset by empty line", "asciidoctor/lists_test_nested_dlist_with_attached_block_offset_by_empty_line.adoc", nestedDlistWithAttachedBlockOffsetByEmptyLine, nil},

	{"should convert glossary list with proper semantics", "asciidoctor/lists_test_should_convert_glossary_list_with_proper_semantics.adoc", shouldConvertGlossaryListWithProperSemantics, nil},

	{"consecutive glossary terms should share same glossentry element in docbook", "asciidoctor/lists_test_consecutive_glossary_terms_should_share_same_glossentry_element_in_docbook.adoc", consecutiveGlossaryTermsShouldShareSameGlossentryElementInDocbook, nil},

	{"should convert horizontal list with proper markup", "asciidoctor/lists_test_should_convert_horizontal_list_with_proper_markup.adoc", shouldConvertHorizontalListWithProperMarkup, nil},

	{"should set col widths of item and label if specified", "asciidoctor/lists_test_should_set_col_widths_of_item_and_label_if_specified.adoc", shouldSetColWidthsOfItemAndLabelIfSpecified, nil},

	{"should set col widths of item and label in docbook if specified", "asciidoctor/lists_test_should_set_col_widths_of_item_and_label_in_docbook_if_specified.adoc", shouldSetColWidthsOfItemAndLabelInDocbookIfSpecified, nil},

	{"should add strong class to label if strong option is set", "asciidoctor/lists_test_should_add_strong_class_to_label_if_strong_option_is_set.adoc", shouldAddStrongClassToLabelIfStrongOptionIsSet, nil},

	{"consecutive terms in horizontal list should share same cell", "asciidoctor/lists_test_consecutive_terms_in_horizontal_list_should_share_same_cell.adoc", consecutiveTermsInHorizontalListShouldShareSameCell, nil},

	{"consecutive terms in horizontal list should share same entry in docbook", "asciidoctor/lists_test_consecutive_terms_in_horizontal_list_should_share_same_entry_in_docbook.adoc", consecutiveTermsInHorizontalListShouldShareSameEntryInDocbook, nil},

	{"should convert horizontal list in docbook with proper markup", "asciidoctor/lists_test_should_convert_horizontal_list_in_docbook_with_proper_markup.adoc", shouldConvertHorizontalListInDocbookWithProperMarkup, nil},

	{"should convert qanda list in HTML with proper semantics", "asciidoctor/lists_test_should_convert_qanda_list_in_html_with_proper_semantics.adoc", shouldConvertQandaListInHtmlWithProperSemantics, nil},

	{"should convert qanda list in DocBook with proper semantics", "asciidoctor/lists_test_should_convert_qanda_list_in_doc_book_with_proper_semantics.adoc", shouldConvertQandaListInDocBookWithProperSemantics, nil},

	{"consecutive questions should share same question element in docbook", "asciidoctor/lists_test_consecutive_questions_should_share_same_question_element_in_docbook.adoc", consecutiveQuestionsShouldShareSameQuestionElementInDocbook, nil},

	{"should convert bibliography list with proper semantics", "asciidoctor/lists_test_should_convert_bibliography_list_with_proper_semantics.adoc", shouldConvertBibliographyListWithProperSemantics, nil},

	{"should convert bibliography list with proper semantics to DocBook", "asciidoctor/lists_test_should_convert_bibliography_list_with_proper_semantics_to_doc_book.adoc", shouldConvertBibliographyListWithProperSemanticsToDocBook, nil},

	{"should warn if a bibliography ID is already in use", "asciidoctor/lists_test_should_warn_if_a_bibliography_id_is_already_in_use.adoc", shouldWarnIfABibliographyIdIsAlreadyInUse, nil},

	{"should automatically add bibliography style to top-level lists in bibliography section", "asciidoctor/lists_test_should_automatically_add_bibliography_style_to_top_level_lists_in_bibliography_section.adoc", shouldAutomaticallyAddBibliographyStyleToTopLevelListsInBibliographySection, nil},

	{"should not recognize bibliography anchor that begins with a digit", "asciidoctor/lists_test_should_not_recognize_bibliography_anchor_that_begins_with_a_digit.adoc", shouldNotRecognizeBibliographyAnchorThatBeginsWithADigit, nil},

	{"should recognize bibliography anchor that contains a digit but does not start with one", "asciidoctor/lists_test_should_recognize_bibliography_anchor_that_contains_a_digit_but_does_not_start_with_one.adoc", shouldRecognizeBibliographyAnchorThatContainsADigitButDoesNotStartWithOne, nil},

	{"should catalog bibliography anchors in bibliography list", "asciidoctor/lists_test_should_catalog_bibliography_anchors_in_bibliography_list.adoc", shouldCatalogBibliographyAnchorsInBibliographyList, nil},

	{"should use reftext from bibliography anchor at xref and entry", "asciidoctor/lists_test_should_use_reftext_from_bibliography_anchor_at_xref_and_entry.adoc", shouldUseReftextFromBibliographyAnchorAtXrefAndEntry, nil},

	{"should assign reftext of bibliography anchor to xreflabel in DocBook backend", "asciidoctor/lists_test_should_assign_reftext_of_bibliography_anchor_to_xreflabel_in_doc_book_backend.adoc", shouldAssignReftextOfBibliographyAnchorToXreflabelInDocBookBackend, nil},

	{"folds text from subsequent line", "asciidoctor/lists_test_folds_text_from_subsequent_line.adoc", foldsTextFromSubsequentLine, nil},

	{"folds text from first line after blank lines", "asciidoctor/lists_test_folds_text_from_first_line_after_blank_lines.adoc", foldsTextFromFirstLineAfterBlankLines, nil},

	{"folds text from first line after blank line and immediately preceding next item", "asciidoctor/lists_test_folds_text_from_first_line_after_blank_line_and_immediately_preceding_next_item.adoc", foldsTextFromFirstLineAfterBlankLineAndImmediatelyPrecedingNextItem, nil},

	{"paragraph offset by blank lines does not break list if label does not have inline text", "asciidoctor/lists_test_paragraph_offset_by_blank_lines_does_not_break_list_if_label_does_not_have_inline_text.adoc", paragraphOffsetByBlankLinesDoesNotBreakListIfLabelDoesNotHaveInlineText, nil},

	{"folds text from first line after comment line", "asciidoctor/lists_test_folds_text_from_first_line_after_comment_line.adoc", foldsTextFromFirstLineAfterCommentLine, nil},

	{"folds text from line following comment line offset by blank line", "asciidoctor/lists_test_folds_text_from_line_following_comment_line_offset_by_blank_line.adoc", foldsTextFromLineFollowingCommentLineOffsetByBlankLine, nil},

	{"folds text from subsequent indented line", "asciidoctor/lists_test_folds_text_from_subsequent_indented_line.adoc", foldsTextFromSubsequentIndentedLine, nil},

	{"folds text from indented line after blank line", "asciidoctor/lists_test_folds_text_from_indented_line_after_blank_line.adoc", foldsTextFromIndentedLineAfterBlankLine, nil},

	{"folds text that looks like ruler offset by blank line", "asciidoctor/lists_test_folds_text_that_looks_like_ruler_offset_by_blank_line.adoc", foldsTextThatLooksLikeRulerOffsetByBlankLine, nil},

	{"folds text that looks like ruler offset by blank line and line comment", "asciidoctor/lists_test_folds_text_that_looks_like_ruler_offset_by_blank_line_and_line_comment.adoc", foldsTextThatLooksLikeRulerOffsetByBlankLineAndLineComment, nil},

	{"folds text that looks like ruler and the line following it offset by blank line", "asciidoctor/lists_test_folds_text_that_looks_like_ruler_and_the_line_following_it_offset_by_blank_line.adoc", foldsTextThatLooksLikeRulerAndTheLineFollowingItOffsetByBlankLine, nil},

	{"folds text that looks like title offset by blank line", "asciidoctor/lists_test_folds_text_that_looks_like_title_offset_by_blank_line.adoc", foldsTextThatLooksLikeTitleOffsetByBlankLine, nil},

	{"folds text that looks like title offset by blank line and line comment", "asciidoctor/lists_test_folds_text_that_looks_like_title_offset_by_blank_line_and_line_comment.adoc", foldsTextThatLooksLikeTitleOffsetByBlankLineAndLineComment, nil},

	{"folds text that looks like admonition offset by blank line", "asciidoctor/lists_test_folds_text_that_looks_like_admonition_offset_by_blank_line.adoc", foldsTextThatLooksLikeAdmonitionOffsetByBlankLine, nil},

	{"folds text that looks like section title offset by blank line", "asciidoctor/lists_test_folds_text_that_looks_like_section_title_offset_by_blank_line.adoc", foldsTextThatLooksLikeSectionTitleOffsetByBlankLine, nil},

	{"folds text of first literal line offset by blank line appends subsequent literals offset by blank line as blocks", "asciidoctor/lists_test_folds_text_of_first_literal_line_offset_by_blank_line_appends_subsequent_literals_offset_by_blank_line_as_blocks.adoc", foldsTextOfFirstLiteralLineOffsetByBlankLineAppendsSubsequentLiteralsOffsetByBlankLineAsBlocks, nil},

	{"folds text of subsequent line and appends following literal line offset by blank line as block if term has no inline description", "asciidoctor/lists_test_folds_text_of_subsequent_line_and_appends_following_literal_line_offset_by_blank_line_as_block_if_term_has_no_inline_description.adoc", foldsTextOfSubsequentLineAndAppendsFollowingLiteralLineOffsetByBlankLineAsBlockIfTermHasNoInlineDescription, nil},

	{"appends literal line attached by continuation as block if item has no inline description", "asciidoctor/lists_test_appends_literal_line_attached_by_continuation_as_block_if_item_has_no_inline_description.adoc", appendsLiteralLineAttachedByContinuationAsBlockIfItemHasNoInlineDescription, nil},

	{"appends literal line attached by continuation as block if item has no inline description followed by ruler", "asciidoctor/lists_test_appends_literal_line_attached_by_continuation_as_block_if_item_has_no_inline_description_followed_by_ruler.adoc", appendsLiteralLineAttachedByContinuationAsBlockIfItemHasNoInlineDescriptionFollowedByRuler, nil},

	{"appends line attached by continuation as block if item has no inline description followed by ruler", "asciidoctor/lists_test_appends_line_attached_by_continuation_as_block_if_item_has_no_inline_description_followed_by_ruler.adoc", appendsLineAttachedByContinuationAsBlockIfItemHasNoInlineDescriptionFollowedByRuler, nil},

	{"appends line attached by continuation as block if item has no inline description followed by block", "asciidoctor/lists_test_appends_line_attached_by_continuation_as_block_if_item_has_no_inline_description_followed_by_block.adoc", appendsLineAttachedByContinuationAsBlockIfItemHasNoInlineDescriptionFollowedByBlock, nil},

	{"appends block attached by continuation but not subsequent block not attached by continuation", "asciidoctor/lists_test_appends_block_attached_by_continuation_but_not_subsequent_block_not_attached_by_continuation.adoc", appendsBlockAttachedByContinuationButNotSubsequentBlockNotAttachedByContinuation, nil},

	{"appends list if item has no inline description", "asciidoctor/lists_test_appends_list_if_item_has_no_inline_description.adoc", appendsListIfItemHasNoInlineDescription, nil},

	{"appends list to first term when followed immediately by second term", "asciidoctor/lists_test_appends_list_to_first_term_when_followed_immediately_by_second_term.adoc", appendsListToFirstTermWhenFollowedImmediatelyBySecondTerm, nil},

	{"appends indented list to first term that is adjacent to second term", "asciidoctor/lists_test_appends_indented_list_to_first_term_that_is_adjacent_to_second_term.adoc", appendsIndentedListToFirstTermThatIsAdjacentToSecondTerm, nil},

	{"appends indented list to first term that is attached by a continuation and adjacent to second term", "asciidoctor/lists_test_appends_indented_list_to_first_term_that_is_attached_by_a_continuation_and_adjacent_to_second_term.adoc", appendsIndentedListToFirstTermThatIsAttachedByAContinuationAndAdjacentToSecondTerm, nil},

	{"appends list and paragraph block when line following list attached by continuation", "asciidoctor/lists_test_appends_list_and_paragraph_block_when_line_following_list_attached_by_continuation.adoc", appendsListAndParagraphBlockWhenLineFollowingListAttachedByContinuation, nil},

	{"first continued line associated with nested list item and second continued line associated with term", "asciidoctor/lists_test_first_continued_line_associated_with_nested_list_item_and_second_continued_line_associated_with_term.adoc", firstContinuedLineAssociatedWithNestedListItemAndSecondContinuedLineAssociatedWithTerm, nil},

	{"literal line attached by continuation swallows adjacent line that looks like term", "asciidoctor/lists_test_literal_line_attached_by_continuation_swallows_adjacent_line_that_looks_like_term.adoc", literalLineAttachedByContinuationSwallowsAdjacentLineThatLooksLikeTerm, nil},

	{"line attached by continuation is appended as paragraph if term has no inline description", "asciidoctor/lists_test_line_attached_by_continuation_is_appended_as_paragraph_if_term_has_no_inline_description.adoc", lineAttachedByContinuationIsAppendedAsParagraphIfTermHasNoInlineDescription, nil},

	{"attached paragraph does not break on adjacent nested description list term", "asciidoctor/lists_test_attached_paragraph_does_not_break_on_adjacent_nested_description_list_term.adoc", attachedParagraphDoesNotBreakOnAdjacentNestedDescriptionListTerm, nil},

	{"attached paragraph is terminated by adjacent sibling description list term", "asciidoctor/lists_test_attached_paragraph_is_terminated_by_adjacent_sibling_description_list_term.adoc", attachedParagraphIsTerminatedByAdjacentSiblingDescriptionListTerm, nil},

	{"attached styled paragraph does not break on adjacent nested description list term", "asciidoctor/lists_test_attached_styled_paragraph_does_not_break_on_adjacent_nested_description_list_term.adoc", attachedStyledParagraphDoesNotBreakOnAdjacentNestedDescriptionListTerm, nil},

	{"appends line as paragraph if attached by continuation following blank line and line comment when term has no inline description", "asciidoctor/lists_test_appends_line_as_paragraph_if_attached_by_continuation_following_blank_line_and_line_comment_when_term_has_no_inline_description.adoc", appendsLineAsParagraphIfAttachedByContinuationFollowingBlankLineAndLineCommentWhenTermHasNoInlineDescription, nil},

	{"line attached by continuation offset by blank line is appended as paragraph if term has no inline description", "asciidoctor/lists_test_line_attached_by_continuation_offset_by_blank_line_is_appended_as_paragraph_if_term_has_no_inline_description.adoc", lineAttachedByContinuationOffsetByBlankLineIsAppendedAsParagraphIfTermHasNoInlineDescription, nil},

	{"delimited block breaks list even when term has no inline description", "asciidoctor/lists_test_delimited_block_breaks_list_even_when_term_has_no_inline_description.adoc", delimitedBlockBreaksListEvenWhenTermHasNoInlineDescription, nil},

	{"block attribute line above delimited block that breaks a dlist is not duplicated", "asciidoctor/lists_test_block_attribute_line_above_delimited_block_that_breaks_a_dlist_is_not_duplicated.adoc", blockAttributeLineAboveDelimitedBlockThatBreaksADlistIsNotDuplicated, nil},

	{"block attribute line above paragraph breaks list even when term has no inline description", "asciidoctor/lists_test_block_attribute_line_above_paragraph_breaks_list_even_when_term_has_no_inline_description.adoc", blockAttributeLineAboveParagraphBreaksListEvenWhenTermHasNoInlineDescription, nil},

	{"block attribute line above paragraph that breaks a dlist is not duplicated", "asciidoctor/lists_test_block_attribute_line_above_paragraph_that_breaks_a_dlist_is_not_duplicated.adoc", blockAttributeLineAboveParagraphThatBreaksADlistIsNotDuplicated, nil},

	{"block anchor line breaks list even when term has no inline description", "asciidoctor/lists_test_block_anchor_line_breaks_list_even_when_term_has_no_inline_description.adoc", blockAnchorLineBreaksListEvenWhenTermHasNoInlineDescription, nil},

	{"block attribute lines above nested horizontal list does not break list", "asciidoctor/lists_test_block_attribute_lines_above_nested_horizontal_list_does_not_break_list.adoc", blockAttributeLinesAboveNestedHorizontalListDoesNotBreakList, nil},

	{"block attribute lines above nested list with style does not break list", "asciidoctor/lists_test_block_attribute_lines_above_nested_list_with_style_does_not_break_list.adoc", blockAttributeLinesAboveNestedListWithStyleDoesNotBreakList, nil},

	{"multiple block attribute lines above nested list does not break list", "asciidoctor/lists_test_multiple_block_attribute_lines_above_nested_list_does_not_break_list.adoc", multipleBlockAttributeLinesAboveNestedListDoesNotBreakList, nil},

	{"multiple block attribute lines separated by empty line above nested list does not break list", "asciidoctor/lists_test_multiple_block_attribute_lines_separated_by_empty_line_above_nested_list_does_not_break_list.adoc", multipleBlockAttributeLinesSeparatedByEmptyLineAboveNestedListDoesNotBreakList, nil},

	{"folds text from inline description and subsequent line", "asciidoctor/lists_test_folds_text_from_inline_description_and_subsequent_line.adoc", foldsTextFromInlineDescriptionAndSubsequentLine, nil},

	{"folds text from inline description and subsequent lines", "asciidoctor/lists_test_folds_text_from_inline_description_and_subsequent_lines.adoc", foldsTextFromInlineDescriptionAndSubsequentLines, nil},

	{"folds text from inline description and line following comment line", "asciidoctor/lists_test_folds_text_from_inline_description_and_line_following_comment_line.adoc", foldsTextFromInlineDescriptionAndLineFollowingCommentLine, nil},

	{"folds text from inline description and subsequent indented line", "asciidoctor/lists_test_folds_text_from_inline_description_and_subsequent_indented_line.adoc", foldsTextFromInlineDescriptionAndSubsequentIndentedLine, nil},

	{"appends literal line offset by blank line as block if item has inline description", "asciidoctor/lists_test_appends_literal_line_offset_by_blank_line_as_block_if_item_has_inline_description.adoc", appendsLiteralLineOffsetByBlankLineAsBlockIfItemHasInlineDescription, nil},

	{"appends literal line offset by blank line as block and appends line after continuation as block if item has inline description", "asciidoctor/lists_test_appends_literal_line_offset_by_blank_line_as_block_and_appends_line_after_continuation_as_block_if_item_has_inline_description.adoc", appendsLiteralLineOffsetByBlankLineAsBlockAndAppendsLineAfterContinuationAsBlockIfItemHasInlineDescription, nil},

	{"appends line after continuation as block and literal line offset by blank line as block if item has inline description", "asciidoctor/lists_test_appends_line_after_continuation_as_block_and_literal_line_offset_by_blank_line_as_block_if_item_has_inline_description.adoc", appendsLineAfterContinuationAsBlockAndLiteralLineOffsetByBlankLineAsBlockIfItemHasInlineDescription, nil},

	{"appends list if item has inline description", "asciidoctor/lists_test_appends_list_if_item_has_inline_description.adoc", appendsListIfItemHasInlineDescription, nil},

	{"appends literal line attached by continuation as block if item has inline description followed by ruler", "asciidoctor/lists_test_appends_literal_line_attached_by_continuation_as_block_if_item_has_inline_description_followed_by_ruler.adoc", appendsLiteralLineAttachedByContinuationAsBlockIfItemHasInlineDescriptionFollowedByRuler, nil},

	{"line offset by blank line breaks list if term has inline description", "asciidoctor/lists_test_line_offset_by_blank_line_breaks_list_if_term_has_inline_description.adoc", lineOffsetByBlankLineBreaksListIfTermHasInlineDescription, nil},

	{"nested term with description does not consume following heading", "asciidoctor/lists_test_nested_term_with_description_does_not_consume_following_heading.adoc", nestedTermWithDescriptionDoesNotConsumeFollowingHeading, nil},

	{"line attached by continuation is appended as paragraph if term has inline description followed by detached paragraph", "asciidoctor/lists_test_line_attached_by_continuation_is_appended_as_paragraph_if_term_has_inline_description_followed_by_detached_paragraph.adoc", lineAttachedByContinuationIsAppendedAsParagraphIfTermHasInlineDescriptionFollowedByDetachedParagraph, nil},

	{"line attached by continuation is appended as paragraph if term has inline description followed by detached block", "asciidoctor/lists_test_line_attached_by_continuation_is_appended_as_paragraph_if_term_has_inline_description_followed_by_detached_block.adoc", lineAttachedByContinuationIsAppendedAsParagraphIfTermHasInlineDescriptionFollowedByDetachedBlock, nil},

	{"line attached by continuation offset by line comment is appended as paragraph if term has inline description", "asciidoctor/lists_test_line_attached_by_continuation_offset_by_line_comment_is_appended_as_paragraph_if_term_has_inline_description.adoc", lineAttachedByContinuationOffsetByLineCommentIsAppendedAsParagraphIfTermHasInlineDescription, nil},

	{"line attached by continuation offset by blank line is appended as paragraph if term has inline description", "asciidoctor/lists_test_line_attached_by_continuation_offset_by_blank_line_is_appended_as_paragraph_if_term_has_inline_description.adoc", lineAttachedByContinuationOffsetByBlankLineIsAppendedAsParagraphIfTermHasInlineDescription, nil},

	{"line comment offset by blank line divides lists because item has text", "asciidoctor/lists_test_line_comment_offset_by_blank_line_divides_lists_because_item_has_text.adoc", lineCommentOffsetByBlankLineDividesListsBecauseItemHasText, nil},

	{"ruler offset by blank line divides lists because item has text", "asciidoctor/lists_test_ruler_offset_by_blank_line_divides_lists_because_item_has_text.adoc", rulerOffsetByBlankLineDividesListsBecauseItemHasText, nil},

	{"block title offset by blank line divides lists and becomes title of second list because item has text", "asciidoctor/lists_test_block_title_offset_by_blank_line_divides_lists_and_becomes_title_of_second_list_because_item_has_text.adoc", blockTitleOffsetByBlankLineDividesListsAndBecomesTitleOfSecondListBecauseItemHasText, nil},

	{"does not recognize callout list denoted by markers that only have a trailing bracket", "asciidoctor/lists_test_does_not_recognize_callout_list_denoted_by_markers_that_only_have_a_trailing_bracket.adoc", doesNotRecognizeCalloutListDenotedByMarkersThatOnlyHaveATrailingBracket, nil},

	{"should not hang if obsolete callout list is found inside list item", "asciidoctor/lists_test_should_not_hang_if_obsolete_callout_list_is_found_inside_list_item.adoc", shouldNotHangIfObsoleteCalloutListIsFoundInsideListItem, nil},

	{"should not hang if obsolete callout list is found inside dlist item", "asciidoctor/lists_test_should_not_hang_if_obsolete_callout_list_is_found_inside_dlist_item.adoc", shouldNotHangIfObsoleteCalloutListIsFoundInsideDlistItem, nil},

	{"should recognize auto-numberd callout list inside list", "asciidoctor/lists_test_should_recognize_auto_numberd_callout_list_inside_list.adoc", shouldRecognizeAutoNumberdCalloutListInsideList, nil},

	{"listing block with sequential callouts followed by adjacent callout list", "asciidoctor/lists_test_listing_block_with_sequential_callouts_followed_by_adjacent_callout_list.adoc", listingBlockWithSequentialCalloutsFollowedByAdjacentCalloutList, nil},

	{"listing block with sequential callouts followed by non-adjacent callout list", "asciidoctor/lists_test_listing_block_with_sequential_callouts_followed_by_non_adjacent_callout_list.adoc", listingBlockWithSequentialCalloutsFollowedByNonAdjacentCalloutList, nil},

	{"listing block with a callout that refers to two different lines", "asciidoctor/lists_test_listing_block_with_a_callout_that_refers_to_two_different_lines.adoc", listingBlockWithACalloutThatRefersToTwoDifferentLines, nil},

	{"source block with non-sequential callouts followed by adjacent callout list", "asciidoctor/lists_test_source_block_with_non_sequential_callouts_followed_by_adjacent_callout_list.adoc", sourceBlockWithNonSequentialCalloutsFollowedByAdjacentCalloutList, nil},

	{"two listing blocks can share the same callout list", "asciidoctor/lists_test_two_listing_blocks_can_share_the_same_callout_list.adoc", twoListingBlocksCanShareTheSameCalloutList, nil},

	{"two listing blocks each followed by an adjacent callout list", "asciidoctor/lists_test_two_listing_blocks_each_followed_by_an_adjacent_callout_list.adoc", twoListingBlocksEachFollowedByAnAdjacentCalloutList, nil},

	{"callout list retains block content", "asciidoctor/lists_test_callout_list_retains_block_content.adoc", calloutListRetainsBlockContent, nil},

	{"callout list retains block content when converted to DocBook", "asciidoctor/lists_test_callout_list_retains_block_content_when_converted_to_doc_book.adoc", calloutListRetainsBlockContentWhenConvertedToDocBook, nil},

	{"escaped callout should not be interpreted as a callout", "asciidoctor/lists_test_escaped_callout_should_not_be_interpreted_as_a_callout.adoc", escapedCalloutShouldNotBeInterpretedAsACallout, nil},

	{"should autonumber <.> callouts", "asciidoctor/lists_test_should_autonumber_<_>_callouts.adoc", shouldAutonumberCallouts, nil},

	{"should not recognize callouts in middle of line", "asciidoctor/lists_test_should_not_recognize_callouts_in_middle_of_line.adoc", shouldNotRecognizeCalloutsInMiddleOfLine, nil},

	{"should allow multiple callouts on the same line", "asciidoctor/lists_test_should_allow_multiple_callouts_on_the_same_line.adoc", shouldAllowMultipleCalloutsOnTheSameLine, nil},

	{"should allow XML comment-style callouts", "asciidoctor/lists_test_should_allow_xml_comment_style_callouts.adoc", shouldAllowXmlCommentStyleCallouts, nil},

	{"should not allow callouts with half an XML comment", "asciidoctor/lists_test_should_not_allow_callouts_with_half_an_xml_comment.adoc", shouldNotAllowCalloutsWithHalfAnXmlComment, nil},

	{"should not recognize callouts in an indented description list paragraph", "asciidoctor/lists_test_should_not_recognize_callouts_in_an_indented_description_list_paragraph.adoc", shouldNotRecognizeCalloutsInAnIndentedDescriptionListParagraph, nil},

	{"should not recognize callouts in an indented outline list paragraph", "asciidoctor/lists_test_should_not_recognize_callouts_in_an_indented_outline_list_paragraph.adoc", shouldNotRecognizeCalloutsInAnIndentedOutlineListParagraph, nil},

	{"should warn if numbers in callout list are out of sequence", "asciidoctor/lists_test_should_warn_if_numbers_in_callout_list_are_out_of_sequence.adoc", shouldWarnIfNumbersInCalloutListAreOutOfSequence, nil},

	{"should preserve line comment chars that precede callout number if icons is not set", "asciidoctor/lists_test_should_preserve_line_comment_chars_that_precede_callout_number_if_icons_is_not_set.adoc", shouldPreserveLineCommentCharsThatPrecedeCalloutNumberIfIconsIsNotSet, nil},

	{"should remove line comment chars that precede callout number if icons is font", "asciidoctor/lists_test_should_remove_line_comment_chars_that_precede_callout_number_if_icons_is_font.adoc", shouldRemoveLineCommentCharsThatPrecedeCalloutNumberIfIconsIsFont, nil},

	{"should allow line comment chars that precede callout number to be specified", "asciidoctor/lists_test_should_allow_line_comment_chars_that_precede_callout_number_to_be_specified.adoc", shouldAllowLineCommentCharsThatPrecedeCalloutNumberToBeSpecified, nil},

	{"should allow line comment chars preceding callout number to be configurable when source-highlighter is coderay", "asciidoctor/lists_test_should_allow_line_comment_chars_preceding_callout_number_to_be_configurable_when_source_highlighter_is_coderay.adoc", shouldAllowLineCommentCharsPrecedingCalloutNumberToBeConfigurableWhenSourceHighlighterIsCoderay, nil},

	{"should not eat whitespace before callout number if line-comment attribute is empty", "asciidoctor/lists_test_should_not_eat_whitespace_before_callout_number_if_line_comment_attribute_is_empty.adoc", shouldNotEatWhitespaceBeforeCalloutNumberIfLineCommentAttributeIsEmpty, nil},

	{"literal block with callouts", "asciidoctor/lists_test_literal_block_with_callouts.adoc", literalBlockWithCallouts, nil},

	{"callout list with icons enabled", "asciidoctor/lists_test_callout_list_with_icons_enabled.adoc", calloutListWithIconsEnabled, nil},

	{"callout list with font-based icons enabled", "asciidoctor/lists_test_callout_list_with_font_based_icons_enabled.adoc", calloutListWithFontBasedIconsEnabled, nil},

	{"should create checklist if at least one item has checkbox syntax", "asciidoctor/lists_test_should_create_checklist_if_at_least_one_item_has_checkbox_syntax.adoc", shouldCreateChecklistIfAtLeastOneItemHasCheckboxSyntax, nil},

	{"should create checklist with font icons if at least one item has checkbox syntax and icons attribute is font", "asciidoctor/lists_test_should_create_checklist_with_font_icons_if_at_least_one_item_has_checkbox_syntax_and_icons_attribute_is_font.adoc", shouldCreateChecklistWithFontIconsIfAtLeastOneItemHasCheckboxSyntaxAndIconsAttributeIsFont, nil},

	{"should create interactive checklist if interactive option is set even with icons attribute is font", "asciidoctor/lists_test_should_create_interactive_checklist_if_interactive_option_is_set_even_with_icons_attribute_is_font.adoc", shouldCreateInteractiveChecklistIfInteractiveOptionIsSetEvenWithIconsAttributeIsFont, nil},

	{"content should return items in list", "asciidoctor/lists_test_content_should_return_items_in_list.adoc", contentShouldReturnItemsInList, nil},

	{"list item should be the parent of block attached to a list item", "asciidoctor/lists_test_list_item_should_be_the_parent_of_block_attached_to_a_list_item.adoc", listItemShouldBeTheParentOfBlockAttachedToAListItem, nil},

	{"outline? should return true for unordered list", "asciidoctor/lists_test_outline_should_return_true_for_unordered_list.adoc", outlineShouldReturnTrueForUnorderedList, nil},

	{"outline? should return true for ordered list", "asciidoctor/lists_test_outline_should_return_true_for_ordered_list.adoc", outlineShouldReturnTrueForOrderedList, nil},

	{"outline? should return false for description list", "asciidoctor/lists_test_outline_should_return_false_for_description_list.adoc", outlineShouldReturnFalseForDescriptionList, nil},

	{"simple? should return true for list item with nested outline list", "asciidoctor/lists_test_simple_should_return_true_for_list_item_with_nested_outline_list.adoc", simpleShouldReturnTrueForListItemWithNestedOutlineList, nil},

	{"simple? should return false for list item with block content", "asciidoctor/lists_test_simple_should_return_false_for_list_item_with_block_content.adoc", simpleShouldReturnFalseForListItemWithBlockContent, nil},

	{"should allow text of ListItem to be assigned", "asciidoctor/lists_test_should_allow_text_of_list_item_to_be_assigned.adoc", shouldAllowTextOfListItemToBeAssigned, nil},

	{"id and role assigned to ulist item in model are transmitted to output", "asciidoctor/lists_test_id_and_role_assigned_to_ulist_item_in_model_are_transmitted_to_output.adoc", idAndRoleAssignedToUlistItemInModelAreTransmittedToOutput, nil},

	{"id and role assigned to olist item in model are transmitted to output", "asciidoctor/lists_test_id_and_role_assigned_to_olist_item_in_model_are_transmitted_to_output.adoc", idAndRoleAssignedToOlistItemInModelAreTransmittedToOutput, nil},

	{"should allow API control over substitutions applied to ListItem text", "asciidoctor/lists_test_should_allow_api_control_over_substitutions_applied_to_list_item_text.adoc", shouldAllowApiControlOverSubstitutionsAppliedToListItemText, nil},

	{"should set lineno to line number in source where list starts", "asciidoctor/lists_test_should_set_lineno_to_line_number_in_source_where_list_starts.adoc", shouldSetLinenoToLineNumberInSourceWhereListStarts, nil},
}

var dashElementsWithNoBlankLines = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
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

var dashElementsSeparatedByBlankLinesShouldMergeLists = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Foo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "-",
			Checklist:     0,
		},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Boo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "-",
			Checklist:     0,
		},
		asciidoc.EmptyLine{
			Text: "",
		},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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

var dashElementsWithInterspersedLineCommentsShouldBeSkippedAndNotBreakList = &asciidoc.Document{
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
				&asciidoc.UnorderedListItem{
					Set: asciidoc.Set{
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
					Set: asciidoc.Set{
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
					Set: asciidoc.Set{
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
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "List",
				},
			},
			Level: 1,
		},
	},
}

var dashElementsSeparatedByALineCommentOffsetByBlankLinesShouldNotMergeLists = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Boo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "-",
			Checklist:     0,
		},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.SingleLineComment{
			Value: "",
		},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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

var dashElementsSeparatedByABlockTitleOffsetByABlankLineShouldNotMergeLists = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Boo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "-",
			Checklist:     0,
		},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Blech",
				},
			},
			AttributeList: asciidoc.AttributeList{
				&asciidoc.TitleAttribute{
					Val: asciidoc.Set{
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

var dashElementsSeparatedByAnAttributeEntryOffsetByABlankLineShouldNotMergeLists = &asciidoc.Document{
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
				&asciidoc.UnorderedListItem{
					Set: asciidoc.Set{
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
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "Boo",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        "-",
					Checklist:     0,
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
				&asciidoc.UnorderedListItem{
					Set: asciidoc.Set{
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
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "List",
				},
			},
			Level: 1,
		},
	},
}

var aNonIndentedWrappedLineIsFoldedIntoTextOfListItem = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
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

var aNonIndentedWrappedLineThatResemblesABlockTitleIsFoldedIntoTextOfListItem = &asciidoc.Document{
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
				&asciidoc.UnorderedListItem{
					Set: asciidoc.Set{
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
					Set: asciidoc.Set{
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
					Set: asciidoc.Set{
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
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "List",
				},
			},
			Level: 1,
		},
	},
}

var aNonIndentedWrappedLineThatResemblesAnAttributeEntryIsFoldedIntoTextOfListItem = &asciidoc.Document{
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
				&asciidoc.UnorderedListItem{
					Set: asciidoc.Set{
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
					Set: asciidoc.Set{
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
					Set: asciidoc.Set{
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
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "List",
				},
			},
			Level: 1,
		},
	},
}

var aListItemWithANestedMarkerTerminatesNonIndentedParagraphForTextOfListItem = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
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

var aListItemForADifferentListTerminatesNonIndentedParagraphForTextOfListItem = &asciidoc.Document{
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
				&asciidoc.UnorderedListItem{
					Set: asciidoc.Set{
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
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "Foo",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        ".",
				},
				asciidoc.EmptyLine{
					Text: "",
				},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Example 1",
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
				&asciidoc.UnorderedListItem{
					Set: asciidoc.Set{
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
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Example 2",
				},
			},
			Level: 1,
		},
	},
}

var anIndentedWrappedLineIsUnindentedAndFoldedIntoTextOfListItem = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
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

var wrappedListItemWithHangingIndentFollowedByNonIndentedLine = &asciidoc.Document{
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
				&asciidoc.UnorderedListItem{
					Set: asciidoc.Set{
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
					Set: asciidoc.Set{
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
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var aListItemWithANestedMarkerTerminatesIndentedParagraphForTextOfListItem = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
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

var aListItemForADifferentListTerminatesIndentedParagraphForTextOfListItem = &asciidoc.Document{
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
				&asciidoc.UnorderedListItem{
					Set: asciidoc.Set{
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
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "Foo",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        ".",
				},
				asciidoc.EmptyLine{
					Text: "",
				},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Example 1",
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
				&asciidoc.UnorderedListItem{
					Set: asciidoc.Set{
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
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Example 2",
				},
			},
			Level: 1,
		},
	},
}

var aLiteralParagraphOffsetByBlankLinesInListContentIsAppendedAsALiteralBlock = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Foo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "-",
			Checklist:     0,
		},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "  literal",
		},
		&asciidoc.NewLine{},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
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

var shouldEscapeSpecialCharactersInAllLiteralParagraphsAttachedToListItem = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "first item",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "  <code>text</code>",
		},
		&asciidoc.NewLine{},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "  more <code>text</code>",
		},
		&asciidoc.NewLine{},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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

var aLiteralParagraphOffsetByABlankLineInListContentFollowedByLineWithContinuationIsAppendedAsTwoBlocks = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Foo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "-",
			Checklist:     0,
		},
		asciidoc.EmptyLine{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
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

var anAdmonitionParagraphAttachedByALineContinuationToAListItemWithWrappedTextShouldProduceAdmonition = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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
					Set: asciidoc.Set{
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

var paragraphLikeBlocksAttachedToAnAncestorListItemByAListContinuationShouldProduceBlocks = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "child",
				},
			},
			AttributeList: nil,
			Indent:        " ",
			Marker:        "**",
			Checklist:     0,
		},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.LineBreak{},
		&asciidoc.NewLine{},
		&asciidoc.Paragraph{
			AttributeList: nil,
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "This is a note.",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 1,
		},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "another child",
				},
			},
			AttributeList: nil,
			Indent:        " ",
			Marker:        "**",
			Checklist:     0,
		},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.ListContinuation{
			ChildElement: &asciidoc.ThematicBreak{
				AttributeList: nil,
			},
		},
	},
}

var shouldNotInheritBlockAttributesFromPreviousBlockWhenBlockIsAttachedUsingAListContinuation = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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

var shouldContinueToParseBlocksAttachedByAListContinuationAfterBlockIsDropped = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
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

var appendsLineAsParagraphIfAttachedByContinuationFollowingLineComment = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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

var aLiteralParagraphWithALineThatAppearsAsAListItemThatIsFollowedByAContinuationShouldCreateTwoBlocks = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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

var consecutiveLiteralParagraphOffsetByBlankLinesInListContentAreAppendedAsALiteralBlocks = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Foo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "-",
			Checklist:     0,
		},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "  literal",
		},
		&asciidoc.NewLine{},
		asciidoc.EmptyLine{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
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

var aLiteralParagraphWithoutATrailingBlankLineConsumesFollowingListItems = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Foo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "-",
			Checklist:     0,
		},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "  literal",
		},
		&asciidoc.NewLine{},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
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

var asteriskElementsWithNoBlankLines = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
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

var asteriskElementsSeparatedByBlankLinesShouldMergeLists = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Foo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Boo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		asciidoc.EmptyLine{
			Text: "",
		},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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

var asteriskElementsWithInterspersedLineCommentsShouldBeSkippedAndNotBreakList = &asciidoc.Document{
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
				&asciidoc.UnorderedListItem{
					Set: asciidoc.Set{
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
					Set: asciidoc.Set{
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
					Set: asciidoc.Set{
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
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "List",
				},
			},
			Level: 1,
		},
	},
}

var asteriskElementsSeparatedByALineCommentOffsetByBlankLinesShouldNotMergeLists = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Boo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.SingleLineComment{
			Value: "",
		},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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

var asteriskElementsSeparatedByABlockTitleOffsetByABlankLineShouldNotMergeLists = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Boo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Blech",
				},
			},
			AttributeList: asciidoc.AttributeList{
				&asciidoc.TitleAttribute{
					Val: asciidoc.Set{
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

var asteriskElementsSeparatedByAnAttributeEntryOffsetByABlankLineShouldNotMergeLists = &asciidoc.Document{
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
				&asciidoc.UnorderedListItem{
					Set: asciidoc.Set{
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
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "Boo",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        "*",
					Checklist:     0,
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
				&asciidoc.UnorderedListItem{
					Set: asciidoc.Set{
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
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "List",
				},
			},
			Level: 1,
		},
	},
}

var listShouldTerminateBeforeNextLowerSectionHeading = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Set:           nil,
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Section",
				},
			},
			Level: 1,
		},
	},
}

var listShouldTerminateBeforeNextLowerSectionHeadingWithImplicitId = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
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
		asciidoc.EmptyLine{
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
			Set: nil,
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Section",
				},
			},
			Level: 1,
		},
	},
}

var shouldNotFindSectionTitleImmediatelyBelowLastListItem = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
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

var quotedText = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "I am ",
				},
				&asciidoc.Bold{
					AttributeList: nil,
					Set: asciidoc.Set{
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
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "I am ",
				},
				&asciidoc.Italic{
					AttributeList: nil,
					Set: asciidoc.Set{
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
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "I am ",
				},
				&asciidoc.Monospace{
					AttributeList: nil,
					Set: asciidoc.Set{
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

var attributeSubstitutions = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "bar",
				},
			},
		},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
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

var leadingDotIsTreatedAsTextNotBlockTitle = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
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

var wordEndingSentenceOnContinuingLineNotTreatedAsAListItem = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OrderedListItem{
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
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

var shouldDiscoverAnchorAtStartOfUnorderedListItemTextAndRegisterItAsAReference = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "The highest peak in the Front Range is ",
		},
		&asciidoc.CrossReference{
			Set: nil,
			ID:  "grays-peak",
		},
		&asciidoc.String{
			Value: ", which tops ",
		},
		&asciidoc.CrossReference{
			Set: nil,
			ID:  "mount-evans",
		},
		&asciidoc.String{
			Value: " by just a few feet.",
		},
		&asciidoc.NewLine{},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
				&asciidoc.Anchor{
					ID: "mount-evans",
					Set: asciidoc.Set{
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
			Set: asciidoc.Set{
				&asciidoc.Anchor{
					ID: "grays-peak",
					Set: asciidoc.Set{
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
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
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

var shouldDiscoverAnchorAtStartOfOrderedListItemTextAndRegisterItAsAReference = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "This is a cross-reference to ",
		},
		&asciidoc.CrossReference{
			Set: nil,
			ID:  "step-2",
		},
		&asciidoc.String{
			Value: ".",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "This is a cross-reference to ",
		},
		&asciidoc.CrossReference{
			Set: nil,
			ID:  "step-4",
		},
		&asciidoc.String{
			Value: ".",
		},
		&asciidoc.NewLine{},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OrderedListItem{
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Ordered list, item 1, without anchor",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
		&asciidoc.OrderedListItem{
			Set: asciidoc.Set{
				&asciidoc.Anchor{
					ID: "step-2",
					Set: asciidoc.Set{
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
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Ordered list, item 3, without anchor",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
		&asciidoc.OrderedListItem{
			Set: asciidoc.Set{
				&asciidoc.Anchor{
					ID: "step-4",
					Set: asciidoc.Set{
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

var shouldDiscoverAnchorAtStartOfCalloutListItemTextAndRegisterItAsAReference = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "This is a cross-reference to ",
		},
		&asciidoc.CrossReference{
			Set: nil,
			ID:  "url-mapping",
		},
		&asciidoc.String{
			Value: ".",
		},
		&asciidoc.NewLine{},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Listing{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Set: asciidoc.Set{
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
					Val: asciidoc.Set{
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
			Set: asciidoc.Set{
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

var asteriskElementMixedWithDashElementsShouldBeNested = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
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

var dashElementMixedWithAsterisksElementsShouldBeNested = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
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

var linesPrefixedWithAlternatingListMarkersSeparatedByBlankLinesShouldBeNested = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Foo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "-",
			Checklist:     0,
		},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Boo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		asciidoc.EmptyLine{
			Text: "",
		},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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

var nestedElements2WithAsterisks = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
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

var nestedElements3WithAsterisks = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
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

var nestedElements4WithAsterisks = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
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

var nestedElements5WithAsterisks = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
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

var levelOfUnorderedListShouldMatchSectionLevel = &asciidoc.Document{
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
				&asciidoc.UnorderedListItem{
					Set: asciidoc.Set{
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
					Set: asciidoc.Set{
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
					Set: asciidoc.Set{
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
					Set: asciidoc.Set{
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
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "item 1.2",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        "*",
					Checklist:     0,
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
						&asciidoc.UnorderedListItem{
							Set: asciidoc.Set{
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
					Title: asciidoc.Set{
						&asciidoc.String{
							Value: "Nested Section",
						},
					},
					Level: 2,
				},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Parent Section",
				},
			},
			Level: 1,
		},
	},
}

var doesNotRecognizeListsWithRepeatingUnicodeBullets = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OrderedListItem{
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Foo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
		&asciidoc.OrderedListItem{
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Boo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "..",
		},
		&asciidoc.OrderedListItem{
			Set: asciidoc.Set{
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

var nestedOrderedElements3 = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OrderedListItem{
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Foo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
		&asciidoc.OrderedListItem{
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Boo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "..",
		},
		&asciidoc.OrderedListItem{
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Snoo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "...",
		},
		&asciidoc.OrderedListItem{
			Set: asciidoc.Set{
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

var levelOfOrderedListShouldMatchSectionLevel = &asciidoc.Document{
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
				&asciidoc.OrderedListItem{
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "item 1.1",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        ".",
				},
				&asciidoc.OrderedListItem{
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "item 2.1",
						},
					},
					AttributeList: nil,
					Indent:        " ",
					Marker:        "..",
				},
				&asciidoc.OrderedListItem{
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "item 3.1",
						},
					},
					AttributeList: nil,
					Indent:        "  ",
					Marker:        "...",
				},
				&asciidoc.OrderedListItem{
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "item 2.2",
						},
					},
					AttributeList: nil,
					Indent:        " ",
					Marker:        "..",
				},
				&asciidoc.OrderedListItem{
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "item 1.2",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        ".",
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
						&asciidoc.OrderedListItem{
							Set: asciidoc.Set{
								&asciidoc.String{
									Value: "item 1.1",
								},
							},
							AttributeList: nil,
							Indent:        "",
							Marker:        ".",
						},
					},
					Title: asciidoc.Set{
						&asciidoc.String{
							Value: "Nested Section",
						},
					},
					Level: 2,
				},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Parent Section",
				},
			},
			Level: 1,
		},
	},
}

var nestedUnorderedInsideOrderedElements = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OrderedListItem{
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Foo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
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

var nestedOrderedInsideUnorderedElements = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Boo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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

var threeLevelsOfAlternatingUnorderedAndOrderedElements = &asciidoc.Document{
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
				&asciidoc.UnorderedListItem{
					Set: asciidoc.Set{
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
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "numbered 1.1",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        ".",
				},
				&asciidoc.UnorderedListItem{
					Set: asciidoc.Set{
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
					Set: asciidoc.Set{
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
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var linesWithAlternatingMarkersOfUnorderedAndOrderedListTypesSeparatedByBlankLinesShouldBeNested = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Foo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OrderedListItem{
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Boo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
		asciidoc.EmptyLine{
			Text: "",
		},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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

var listItemWithLiteralContentShouldNotConsumeNestedListOfDifferentType = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "bullet",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "-",
			Checklist:     0,
		},
		asciidoc.EmptyLine{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OrderedListItem{
			Set: asciidoc.Set{
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

var nestedListItemDoesNotEatTheTitleOfTheFollowingDetachedBlock = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "nested bullet 2",
				},
			},
			AttributeList: nil,
			Indent:        "  ",
			Marker:        "*",
			Checklist:     0,
		},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.LiteralBlock{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.TitleAttribute{
					Val: asciidoc.Set{
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

var linesWithAlternatingMarkersOfBulletedAndDescriptionListTypesSeparatedByBlankLinesShouldBeNested = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Foo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "term1:: def1",
		},
		&asciidoc.NewLine{},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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

var nestedOrderedWithAttributeInsideUnorderedElements = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Boo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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

var adjacentListContinuationLineAttachesFollowingParagraph = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
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

var adjacentListContinuationLineAttachesFollowingBlock = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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
				Set: asciidoc.Set{
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

var adjacentListContinuationLineAttachesFollowingBlockWithBlockAttributes = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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
					ID:  "beck",
					Set: nil,
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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

var trailingBlockAttributeLineAttachedByContinuationShouldNotCreateBlock = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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

var trailingBlockTitleLineAttachedByContinuationShouldNotCreateBlock = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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

var consecutiveBlocksInListContinuationAttachToListItem = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Item one, quote block",
				},
				&asciidoc.NewLine{},
			},
		},
		&asciidoc.ListContinuation{
			ChildElement: &asciidoc.UnorderedListItem{
				Set: asciidoc.Set{
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

var listItemWithHangingIndentFollowedByBlockAttachedByListContinuation = &asciidoc.Document{
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
				&asciidoc.OrderedListItem{
					Set: asciidoc.Set{
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
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.OrderedListItem{
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "list item 2",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        ".",
				},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var listItemParagraphInListItemAndNestedListItem = &asciidoc.Document{
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
				&asciidoc.OrderedListItem{
					Set: asciidoc.Set{
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
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.UnorderedListItem{
					Set: asciidoc.Set{
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
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.OrderedListItem{
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "list item 2",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        ".",
				},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var trailingListContinuationsShouldAttachToListItemsAtRespectiveLevels = &asciidoc.Document{
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
				&asciidoc.OrderedListItem{
					Set: asciidoc.Set{
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
					Set: asciidoc.Set{
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
					Set: asciidoc.Set{
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
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.LineBreak{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "paragraph for list item 1",
				},
				&asciidoc.NewLine{},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.OrderedListItem{
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "list item 2",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        ".",
				},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var trailingListContinuationsShouldAttachToListItemsOfDifferentTypesAtRespectiveLevels = &asciidoc.Document{
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
				&asciidoc.UnorderedListItem{
					Set: asciidoc.Set{
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
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "numbered 1.1",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        ".",
				},
				&asciidoc.UnorderedListItem{
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "bullet 1.1.1",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        "**",
					Checklist:     0,
				},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.LineBreak{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "numbered 1.1 paragraph",
				},
				&asciidoc.NewLine{},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.LineBreak{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "bullet 1 paragraph",
				},
				&asciidoc.NewLine{},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.UnorderedListItem{
					Set: asciidoc.Set{
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
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var repeatedListContinuationsShouldAttachToListItemsAtRespectiveLevels = &asciidoc.Document{
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
				&asciidoc.OrderedListItem{
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "list item 1",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        ".",
				},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.UnorderedListItem{
					Set: asciidoc.Set{
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
					Set: asciidoc.Set{
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
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.LineBreak{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "paragraph for list item 1",
				},
				&asciidoc.NewLine{},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.OrderedListItem{
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "list item 2",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        ".",
				},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var repeatedListContinuationsAttachedDirectlyToListItemShouldAttachToListItemsAtRespectiveLevels = &asciidoc.Document{
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
				&asciidoc.OrderedListItem{
					Set: asciidoc.Set{
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
					Set: asciidoc.Set{
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
					Set: asciidoc.Set{
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
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.LineBreak{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "paragraph for list item 1",
				},
				&asciidoc.NewLine{},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.OrderedListItem{
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "list item 2",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        ".",
				},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var repeatedListContinuationsShouldAttachToListItemsAtRespectiveLevelsIgnoringBlankLines = &asciidoc.Document{
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
				&asciidoc.OrderedListItem{
					Set: asciidoc.Set{
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
					Set: asciidoc.Set{
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
					Set: asciidoc.Set{
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
				asciidoc.EmptyLine{
					Text: "",
				},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.LineBreak{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "paragraph for list item 1",
				},
				&asciidoc.NewLine{},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.OrderedListItem{
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "list item 2",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        ".",
				},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var trailingListContinuationsShouldIgnorePrecedingBlankLines = &asciidoc.Document{
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
				&asciidoc.UnorderedListItem{
					Set: asciidoc.Set{
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
					Set: asciidoc.Set{
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
					Set: asciidoc.Set{
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
				asciidoc.EmptyLine{
					Text: "",
				},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.LineBreak{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "bullet 1.1 paragraph",
				},
				&asciidoc.NewLine{},
				asciidoc.EmptyLine{
					Text: "",
				},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.LineBreak{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "bullet 1 paragraph",
				},
				&asciidoc.NewLine{},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.UnorderedListItem{
					Set: asciidoc.Set{
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
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var indentedOutlineListItemWithDifferentMarkerOffsetByABlankLineShouldBeRecognizedAsANestedList = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "item 1",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OrderedListItem{
			Set: asciidoc.Set{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OrderedListItem{
			Set: asciidoc.Set{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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

var indentedDescriptionListItemInsideOutlineListItemOffsetByABlankLineShouldBeRecognizedAsANestedList = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "item 1",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		asciidoc.EmptyLine{
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
		asciidoc.EmptyLine{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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

var consecutiveListContinuationLinesAreFolded = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
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

var shouldWarnIfUnterminatedBlockIsDetectedInListItem = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
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

var dotElementsWithNoBlankLines = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OrderedListItem{
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Foo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
		&asciidoc.OrderedListItem{
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Boo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
		&asciidoc.OrderedListItem{
			Set: asciidoc.Set{
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

var shouldRepresentExplicitRoleAttributeAsStyleClass = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OrderedListItem{
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Once",
				},
			},
			AttributeList: asciidoc.AttributeList{
				&asciidoc.NamedAttribute{
					Name: "role",
					Val: asciidoc.Set{
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
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Again",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
		&asciidoc.OrderedListItem{
			Set: asciidoc.Set{
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

var shouldBaseListStyleOnMarkerLengthRatherThanListDepth = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OrderedListItem{
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "parent",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "...",
		},
		&asciidoc.OrderedListItem{
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "child",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "..",
		},
		&asciidoc.OrderedListItem{
			Set: asciidoc.Set{
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

var shouldAllowListStyleToBeSpecifiedExplicitlyWhenUsingMarkersWithImplicitStyle = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OrderedListItem{
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "1",
				},
			},
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Set: asciidoc.Set{
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
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "2",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "ii)",
		},
		&asciidoc.OrderedListItem{
			Set: asciidoc.Set{
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

var shouldRepresentCustomNumberingAndExplicitRoleAttributeAsStyleClasses = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OrderedListItem{
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Once",
				},
			},
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Set: asciidoc.Set{
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
					Val: asciidoc.Set{
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
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Again",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
		&asciidoc.OrderedListItem{
			Set: asciidoc.Set{
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

var shouldSetReversedAttributeOnListIfReversedOptionIsSet = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OrderedListItem{
			Set: asciidoc.Set{
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
							Set: asciidoc.Set{
								&asciidoc.String{
									Value: "reversed",
								},
							},
						},
					},
				},
				&asciidoc.NamedAttribute{
					Name: "start",
					Val: asciidoc.Set{
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
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "two",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
		&asciidoc.OrderedListItem{
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "one",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
		&asciidoc.OrderedListItem{
			Set: asciidoc.Set{
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

var shouldRepresentImplicitRoleAttributeAsStyleClass = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OrderedListItem{
			Set: asciidoc.Set{
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
							Set: asciidoc.Set{
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
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Again",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
		&asciidoc.OrderedListItem{
			Set: asciidoc.Set{
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

var shouldRepresentCustomNumberingAndImplicitRoleAttributeAsStyleClasses = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OrderedListItem{
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Once",
				},
			},
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Set: asciidoc.Set{
							&asciidoc.String{
								Value: "loweralpha",
							},
						},
					},
					ID: nil,
					Roles: []*asciidoc.ShorthandRole{
						&asciidoc.ShorthandRole{
							Set: asciidoc.Set{
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
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Again",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
		&asciidoc.OrderedListItem{
			Set: asciidoc.Set{
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

var dotElementsSeparatedByBlankLinesShouldMergeLists = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OrderedListItem{
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Foo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OrderedListItem{
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Boo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
		asciidoc.EmptyLine{
			Text: "",
		},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OrderedListItem{
			Set: asciidoc.Set{
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

var dotElementsWithInterspersedLineCommentsShouldBeSkippedAndNotBreakList = &asciidoc.Document{
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
				&asciidoc.OrderedListItem{
					Set: asciidoc.Set{
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
					Set: asciidoc.Set{
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
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "Blech",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        ".",
				},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "List",
				},
			},
			Level: 1,
		},
	},
}

var dotElementsSeparatedByLineCommentOffsetByBlankLinesShouldNotMergeLists = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OrderedListItem{
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Foo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
		&asciidoc.OrderedListItem{
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Boo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.SingleLineComment{
			Value: "",
		},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OrderedListItem{
			Set: asciidoc.Set{
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

var dotElementsSeparatedByABlockTitleOffsetByABlankLineShouldNotMergeLists = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OrderedListItem{
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Foo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
		&asciidoc.OrderedListItem{
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Boo",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OrderedListItem{
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Blech",
				},
			},
			AttributeList: asciidoc.AttributeList{
				&asciidoc.TitleAttribute{
					Val: asciidoc.Set{
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

var dotElementsSeparatedByAnAttributeEntryOffsetByABlankLineShouldNotMergeLists = &asciidoc.Document{
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
				&asciidoc.OrderedListItem{
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "Foo",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        ".",
				},
				&asciidoc.OrderedListItem{
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "Boo",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        ".",
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
				&asciidoc.OrderedListItem{
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "Blech",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        ".",
				},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "List",
				},
			},
			Level: 1,
		},
	},
}

var shouldUseStartNumberInDocbook5Backend = &asciidoc.Document{
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
				&asciidoc.OrderedListItem{
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "item 7",
						},
					},
					AttributeList: asciidoc.AttributeList{
						&asciidoc.NamedAttribute{
							Name: "start",
							Val: asciidoc.Set{
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
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "item 8",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        ".",
				},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "List",
				},
			},
			Level: 1,
		},
	},
}

var shouldWarnIfExplicitUppercaseRomanNumeralsInListAreOutOfSequence = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OrderedListItem{
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "one",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "I)",
		},
		&asciidoc.OrderedListItem{
			Set: asciidoc.Set{
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

var shouldWarnIfExplicitLowercaseRomanNumeralsInListAreOutOfSequence = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OrderedListItem{
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "one",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "i)",
		},
		&asciidoc.OrderedListItem{
			Set: asciidoc.Set{
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

var shouldNotParseABareDlistDelimiterAsADlist = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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

var shouldParseSiblingItemsUsingSameRules = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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

var shouldAllowTermToEndWithASemicolonWhenUsingDoubleSemicolonDelimiter = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "term;;; def",
		},
		&asciidoc.NewLine{},
	},
}

var singleLineIndentedAdjacentElements = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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

var singleLineElementsSeparatedByBlankLineShouldCreateASingleList = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "term1:: def1",
		},
		&asciidoc.NewLine{},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "term2:: def2",
		},
		&asciidoc.NewLine{},
	},
}

var aLineCommentBetweenElementsShouldDivideThemIntoSeparateLists = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "term1:: def1",
		},
		&asciidoc.NewLine{},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.SingleLineComment{
			Value: "",
		},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "term2:: def2",
		},
		&asciidoc.NewLine{},
	},
}

var aRulerBetweenElementsShouldDivideThemIntoSeparateLists = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "term1:: def1",
		},
		&asciidoc.NewLine{},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.ThematicBreak{
			AttributeList: nil,
		},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "term2:: def2",
		},
		&asciidoc.NewLine{},
	},
}

var aBlockTitleBetweenElementsShouldDivideThemIntoSeparateLists = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "term1:: def1",
		},
		&asciidoc.NewLine{},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Paragraph{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.TitleAttribute{
					Val: asciidoc.Set{
						&asciidoc.String{
							Value: "Some more",
						},
					},
				},
			},
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "term2:: def2",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 0,
		},
	},
}

var multiLineElementsWithParagraphContent = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Set: asciidoc.Set{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "def1",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Set{
				&asciidoc.String{
					Value: "term",
				},
				&asciidoc.String{
					Value: "1",
				},
			},
		},
		&asciidoc.DescriptionListItem{
			Set: asciidoc.Set{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "def2",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Set{
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

var multiLineElementsWithIndentedParagraphContent = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Set: asciidoc.Set{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: " def1",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Set{
				&asciidoc.String{
					Value: "term",
				},
				&asciidoc.String{
					Value: "1",
				},
			},
		},
		&asciidoc.DescriptionListItem{
			Set: asciidoc.Set{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "  def2",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Set{
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

var multiLineElementsWithIndentedParagraphContentThatIncludesCommentLines = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Set: asciidoc.Set{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: " def1",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Set{
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
			Set: asciidoc.Set{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "  def2",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Set{
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

var shouldNotStripCommentLineInLiteralParagraphBlockAttachedToListItem = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Set: asciidoc.Set{
				&asciidoc.NewLine{},
				&asciidoc.LineBreak{},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Set{
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

var multiLineElementWithParagraphStartingWithMultipleDashesShouldNotBeSeenAsList = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Set: asciidoc.Set{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "  def1",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Set{
				&asciidoc.String{
					Value: "term",
				},
				&asciidoc.String{
					Value: "1",
				},
			},
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "and a note",
				},
			},
			AttributeList: nil,
			Indent:        "  ",
			Marker:        "--",
			Checklist:     0,
		},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Set: asciidoc.Set{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "  def2",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Set{
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

var multiLineElementWithMultipleTerms = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Set: asciidoc.Set{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "term2::",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Set{
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

var consecutiveTermsShareSameVarlistentryInDocbook = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Set: asciidoc.Set{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "alt term::",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Set{
				&asciidoc.String{
					Value: "term",
				},
			},
		},
		&asciidoc.String{
			Value: "description",
		},
		&asciidoc.NewLine{},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "last::",
		},
		&asciidoc.NewLine{},
	},
}

var multiLineElementsWithBlankLineBeforeParagraphContent = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Set: asciidoc.Set{
				&asciidoc.NewLine{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "def1",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Set{
				&asciidoc.String{
					Value: "term",
				},
				&asciidoc.String{
					Value: "1",
				},
			},
		},
		&asciidoc.DescriptionListItem{
			Set: asciidoc.Set{
				&asciidoc.NewLine{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "def2",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Set{
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

var multiLineElementsWithParagraphAndLiteralContent = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Set: asciidoc.Set{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "def1",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Set{
				&asciidoc.String{
					Value: "term",
				},
				&asciidoc.String{
					Value: "1",
				},
			},
		},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "  literal",
		},
		&asciidoc.NewLine{},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Set: asciidoc.Set{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "  def2",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Set{
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

var mixedSingleAndMultiLineAdjacentElements = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "term1:: def1",
		},
		&asciidoc.NewLine{},
		&asciidoc.DescriptionListItem{
			Set: asciidoc.Set{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "def2",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Set{
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

var shouldDiscoverAnchorAtStartOfDescriptionTermTextAndRegisterItAsAReference = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "The highest peak in the Front Range is ",
		},
		&asciidoc.CrossReference{
			Set: nil,
			ID:  "grays-peak",
		},
		&asciidoc.String{
			Value: ", which tops ",
		},
		&asciidoc.CrossReference{
			Set: nil,
			ID:  "mount-evans",
		},
		&asciidoc.String{
			Value: " by just a few feet.",
		},
		&asciidoc.NewLine{},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Anchor{
			ID: "mount-evans",
			Set: asciidoc.Set{
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
			ID:  "grays-peak",
			Set: nil,
		},
		&asciidoc.String{
			Value: "Grays Peak:: 14,278 feet",
		},
		&asciidoc.NewLine{},
	},
}

var missingSpaceBeforeTermDoesNotProduceDescriptionList = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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

var literalBlockInsideDescriptionList = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Set: asciidoc.Set{
				&asciidoc.NewLine{},
				&asciidoc.LineBreak{},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Set{
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

var literalBlockInsideDescriptionListWithTrailingLineContinuation = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Set: asciidoc.Set{
				&asciidoc.NewLine{},
				&asciidoc.LineBreak{},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Set{
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

var multipleListingBlocksInsideDescriptionList = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Set: asciidoc.Set{
				&asciidoc.NewLine{},
				&asciidoc.LineBreak{},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Set{
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

var openBlockInsideDescriptionList = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Set: asciidoc.Set{
				&asciidoc.NewLine{},
				&asciidoc.LineBreak{},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Set{
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
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Open block as description of term.",
				},
				&asciidoc.NewLine{},
				asciidoc.EmptyLine{
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

var paragraphAttachedByAListContinuationOnEitherSideInADescriptionList = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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

var paragraphAttachedByAListContinuationOnEitherSideToAMultiLineElementInADescriptionList = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Set: asciidoc.Set{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "def1",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Set{
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

var shouldContinueToParseSubsequentBlocksAttachedToListItemAfterFirstBlockIsDropped = &asciidoc.Document{
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
		&asciidoc.DescriptionListItem{
			Set: asciidoc.Set{
				&asciidoc.NewLine{},
				&asciidoc.LineBreak{},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Set{
				&asciidoc.String{
					Value: "term",
				},
			},
		},
		&asciidoc.BlockImage{
			AttributeList: nil,
			ImagePath: asciidoc.Set{
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

var verseParagraphInsideADescriptionList = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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
			},
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "la la la",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 0,
		},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "term2:: def",
		},
		&asciidoc.NewLine{},
	},
}

var listInsideADescriptionList = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Set: asciidoc.Set{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "* level 1",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Set{
				&asciidoc.String{
					Value: "term",
				},
				&asciidoc.String{
					Value: "1",
				},
			},
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
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

var listInsideADescriptionListOffsetByBlankLines = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Set: asciidoc.Set{
				&asciidoc.NewLine{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "* level 1",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Set{
				&asciidoc.String{
					Value: "term",
				},
				&asciidoc.String{
					Value: "1",
				},
			},
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "level 1",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "term2:: def",
		},
		&asciidoc.NewLine{},
	},
}

var shouldOnlyGrabOneLineFollowingLastItemIfItemHasNoInlineDescription = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Set: asciidoc.Set{
				&asciidoc.NewLine{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "def1",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Set{
				&asciidoc.String{
					Value: "term",
				},
				&asciidoc.String{
					Value: "1",
				},
			},
		},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Set: asciidoc.Set{
				&asciidoc.NewLine{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "def2",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Set{
				&asciidoc.String{
					Value: "term",
				},
				&asciidoc.String{
					Value: "2",
				},
			},
		},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "A new paragraph",
		},
		&asciidoc.NewLine{},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "Another new paragraph",
		},
		&asciidoc.NewLine{},
	},
}

var shouldOnlyGrabOneLiteralLineFollowingLastItemIfItemHasNoInlineDescription = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Set: asciidoc.Set{
				&asciidoc.NewLine{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "def1",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Set{
				&asciidoc.String{
					Value: "term",
				},
				&asciidoc.String{
					Value: "1",
				},
			},
		},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Set: asciidoc.Set{
				&asciidoc.NewLine{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "  def2",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Set{
				&asciidoc.String{
					Value: "term",
				},
				&asciidoc.String{
					Value: "2",
				},
			},
		},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "A new paragraph",
		},
		&asciidoc.NewLine{},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "Another new paragraph",
		},
		&asciidoc.NewLine{},
	},
}

var shouldAppendSubsequentParagraphLiteralsToListItemAsBlockContent = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Set: asciidoc.Set{
				&asciidoc.NewLine{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "def1",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Set{
				&asciidoc.String{
					Value: "term",
				},
				&asciidoc.String{
					Value: "1",
				},
			},
		},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Set: asciidoc.Set{
				&asciidoc.NewLine{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "  def2",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Set{
				&asciidoc.String{
					Value: "term",
				},
				&asciidoc.String{
					Value: "2",
				},
			},
		},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "  literal",
		},
		&asciidoc.NewLine{},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "A new paragraph.",
		},
		&asciidoc.NewLine{},
	},
}

var shouldNotMatchCommentLineThatLooksLikeDescriptionListTerm = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "before",
		},
		&asciidoc.NewLine{},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.SingleLineComment{
			Value: "key:: val",
		},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "after",
		},
		&asciidoc.NewLine{},
	},
}

var shouldNotMatchCommentLineFollowingListThatLooksLikeDescriptionListTerm = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "item",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.SingleLineComment{
			Value: "term:: desc",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Set: asciidoc.Set{
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "section text",
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

var shouldNotMatchCommentLineThatLooksLikeSiblingDescriptionListTerm = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "before",
		},
		&asciidoc.NewLine{},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "foo:: bar",
		},
		&asciidoc.NewLine{},
		&asciidoc.SingleLineComment{
			Value: "yin:: yang",
		},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "after",
		},
		&asciidoc.NewLine{},
	},
}

var shouldNotHangOnDescriptionListItemInListThatBeginsWith = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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

var shouldNotHangOnSiblingDescriptionListItemThatBeginsWith = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Set: asciidoc.Set{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "///b::",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Set{
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

var shouldSkipDlistTermThatBeginsWithUnlessItBeginsWith = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Set: asciidoc.Set{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "//ignored term:: def",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Set{
				&asciidoc.String{
					Value: "category a",
				},
			},
		},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Set: asciidoc.Set{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "///term:: def",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Set{
				&asciidoc.String{
					Value: "category b",
				},
			},
		},
	},
}

var moreThan4ConsecutiveColonsShouldBecomePartOfDescriptionListTerm = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "A term::::: a description",
		},
		&asciidoc.NewLine{},
	},
}

var textMethodOfDdNodeShouldReturnNilIfDdNodeOnlyContainsBlocks = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Set: asciidoc.Set{
				&asciidoc.NewLine{},
				&asciidoc.LineBreak{},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Set{
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

var shouldNotParseANestedDlistDelimiterWithoutATermAsADlist = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Set: asciidoc.Set{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: ";;",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Set{
				&asciidoc.String{
					Value: "t",
				},
			},
		},
	},
}

var shouldNotParseANestedIndentedDlistDelimiterWithoutATermAsADlist = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Set: asciidoc.Set{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "desc",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Set{
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

var singleLineAdjacentNestedElements = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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

var singleLineAdjacentMaximumNestedElements = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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

var singleLineNestedElementsSeparatedByBlankLineAtTopLevel = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "term1:: def1",
		},
		&asciidoc.NewLine{},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "label1::: detail1",
		},
		&asciidoc.NewLine{},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "term2:: def2",
		},
		&asciidoc.NewLine{},
	},
}

var singleLineNestedElementsSeparatedByBlankLineAtNestedLevel = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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
		asciidoc.EmptyLine{
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

var singleLineAdjacentNestedElementsWithAlternateDelimiters = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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

var multiLineAdjacentNestedElements = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Set: asciidoc.Set{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "def1",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Set{
				&asciidoc.String{
					Value: "term",
				},
				&asciidoc.String{
					Value: "1",
				},
			},
		},
		&asciidoc.DescriptionListItem{
			Set: asciidoc.Set{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "detail1",
				},
			},
			AttributeList: nil,
			Marker:        ":::",
			Term: asciidoc.Set{
				&asciidoc.String{
					Value: "label",
				},
				&asciidoc.String{
					Value: "1",
				},
			},
		},
		&asciidoc.DescriptionListItem{
			Set: asciidoc.Set{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "def2",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Set{
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

var multiLineNestedElementsSeparatedByBlankLineAtNestedLevelRepeated = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Set: asciidoc.Set{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "def1",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Set{
				&asciidoc.String{
					Value: "term",
				},
				&asciidoc.String{
					Value: "1",
				},
			},
		},
		&asciidoc.DescriptionListItem{
			Set: asciidoc.Set{
				&asciidoc.NewLine{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "detail1",
				},
			},
			AttributeList: nil,
			Marker:        ":::",
			Term: asciidoc.Set{
				&asciidoc.String{
					Value: "label",
				},
				&asciidoc.String{
					Value: "1",
				},
			},
		},
		&asciidoc.DescriptionListItem{
			Set: asciidoc.Set{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "detail2",
				},
			},
			AttributeList: nil,
			Marker:        ":::",
			Term: asciidoc.Set{
				&asciidoc.String{
					Value: "label",
				},
				&asciidoc.String{
					Value: "2",
				},
			},
		},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "term2:: def2",
		},
		&asciidoc.NewLine{},
	},
}

var multiLineElementWithIndentedNestedElement = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Set: asciidoc.Set{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "  def1",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Set{
				&asciidoc.String{
					Value: "term",
				},
				&asciidoc.String{
					Value: "1",
				},
			},
		},
		&asciidoc.DescriptionListItem{
			Set: asciidoc.Set{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "   detail1",
				},
			},
			AttributeList: nil,
			Marker:        ";;",
			Term: asciidoc.Set{
				&asciidoc.String{
					Value: "  label",
				},
				&asciidoc.String{
					Value: "1",
				},
			},
		},
		&asciidoc.DescriptionListItem{
			Set: asciidoc.Set{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "  def2",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Set{
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

var mixedSingleAndMultiLineElementsWithIndentedNestedElements = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "term1:: def1",
		},
		&asciidoc.NewLine{},
		&asciidoc.DescriptionListItem{
			Set: asciidoc.Set{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "   detail1",
				},
			},
			AttributeList: nil,
			Marker:        ":::",
			Term: asciidoc.Set{
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

var multiLineElementsWithFirstParagraphFoldedToTextWithAdjacentNestedElement = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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
			Set: asciidoc.Set{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "detail1",
				},
			},
			AttributeList: nil,
			Marker:        ":::",
			Term: asciidoc.Set{
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

var nestedDlistAttachedByListContinuationShouldNotConsumeDetachedParagraph = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "paragraph",
		},
		&asciidoc.NewLine{},
	},
}

var nestedDlistWithAttachedBlockOffsetByEmptyLine = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Set: asciidoc.Set{
				&asciidoc.NewLine{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "term 1:::",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Set{
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
				Set: asciidoc.Set{
					&asciidoc.String{
						Value: "def 1",
					},
					&asciidoc.NewLine{},
				},
			},
		},
	},
}

var shouldConvertGlossaryListWithProperSemantics = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Paragraph{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Set: asciidoc.Set{
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
			Set: asciidoc.Set{
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

var consecutiveGlossaryTermsShouldShareSameGlossentryElementInDocbook = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Set: asciidoc.Set{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "alt term::",
				},
			},
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Set: asciidoc.Set{
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
			Term: asciidoc.Set{
				&asciidoc.String{
					Value: "term",
				},
			},
		},
		&asciidoc.String{
			Value: "description",
		},
		&asciidoc.NewLine{},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "last::",
		},
		&asciidoc.NewLine{},
	},
}

var shouldConvertHorizontalListWithProperMarkup = &asciidoc.Document{
	Set: asciidoc.Set{
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
			},
			Set: asciidoc.Set{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "second term:: description",
		},
		&asciidoc.NewLine{},
	},
}

var shouldSetColWidthsOfItemAndLabelIfSpecified = &asciidoc.Document{
	Set: asciidoc.Set{
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
							Value: "25",
						},
					},
					Quote: 2,
				},
				&asciidoc.NamedAttribute{
					Name: "itemwidth",
					Val: asciidoc.Set{
						&asciidoc.String{
							Value: "75",
						},
					},
					Quote: 2,
				},
			},
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "term:: def",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 0,
		},
	},
}

var shouldSetColWidthsOfItemAndLabelInDocbookIfSpecified = &asciidoc.Document{
	Set: asciidoc.Set{
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
							Value: "25",
						},
					},
					Quote: 2,
				},
				&asciidoc.NamedAttribute{
					Name: "itemwidth",
					Val: asciidoc.Set{
						&asciidoc.String{
							Value: "75",
						},
					},
					Quote: 2,
				},
			},
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "term:: def",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 0,
		},
	},
}

var shouldAddStrongClassToLabelIfStrongOptionIsSet = &asciidoc.Document{
	Set: asciidoc.Set{
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
					Name: "options",
					Val: asciidoc.Set{
						&asciidoc.String{
							Value: "strong",
						},
					},
					Quote: 2,
				},
			},
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "term:: def",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 0,
		},
	},
}

var consecutiveTermsInHorizontalListShouldShareSameCell = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Set: asciidoc.Set{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "alt term::",
				},
			},
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
			},
			Marker: "::",
			Term: asciidoc.Set{
				&asciidoc.String{
					Value: "term",
				},
			},
		},
		&asciidoc.String{
			Value: "description",
		},
		&asciidoc.NewLine{},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "last::",
		},
		&asciidoc.NewLine{},
	},
}

var consecutiveTermsInHorizontalListShouldShareSameEntryInDocbook = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Set: asciidoc.Set{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "alt term::",
				},
			},
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
			},
			Marker: "::",
			Term: asciidoc.Set{
				&asciidoc.String{
					Value: "term",
				},
			},
		},
		&asciidoc.String{
			Value: "description",
		},
		&asciidoc.NewLine{},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "last::",
		},
		&asciidoc.NewLine{},
	},
}

var shouldConvertHorizontalListInDocbookWithProperMarkup = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Paragraph{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.TitleAttribute{
					Val: asciidoc.Set{
						&asciidoc.String{
							Value: "Terms",
						},
					},
				},
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
			},
			Set: asciidoc.Set{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "second term:: description",
		},
		&asciidoc.NewLine{},
	},
}

var shouldConvertQandaListInHtmlWithProperSemantics = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Set: asciidoc.Set{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "        Answer 1.",
				},
			},
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Set: asciidoc.Set{
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
			Term: asciidoc.Set{
				&asciidoc.String{
					Value: "Question ",
				},
				&asciidoc.String{
					Value: "1",
				},
			},
		},
		&asciidoc.DescriptionListItem{
			Set: asciidoc.Set{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "        Answer 2.",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Set{
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
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "A note about Answer 2.",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 1,
		},
	},
}

var shouldConvertQandaListInDocBookWithProperSemantics = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Set: asciidoc.Set{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "        Answer 1.",
				},
			},
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Set: asciidoc.Set{
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
			Term: asciidoc.Set{
				&asciidoc.String{
					Value: "Question ",
				},
				&asciidoc.String{
					Value: "1",
				},
			},
		},
		&asciidoc.DescriptionListItem{
			Set: asciidoc.Set{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "        Answer 2.",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Set{
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
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "A note about Answer 2.",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 1,
		},
	},
}

var consecutiveQuestionsShouldShareSameQuestionElementInDocbook = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Set: asciidoc.Set{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "follow-up question::",
				},
			},
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Set: asciidoc.Set{
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
			Term: asciidoc.Set{
				&asciidoc.String{
					Value: "question",
				},
			},
		},
		&asciidoc.String{
			Value: "response",
		},
		&asciidoc.NewLine{},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "last question::",
		},
		&asciidoc.NewLine{},
	},
}

var shouldConvertBibliographyListWithProperSemantics = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "[",
				},
				&asciidoc.Anchor{
					ID:  "taoup",
					Set: nil,
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
						Set: asciidoc.Set{
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
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "[",
				},
				&asciidoc.Anchor{
					ID:  "walsh-muellner",
					Set: nil,
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
					Set: asciidoc.Set{
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

var shouldConvertBibliographyListWithProperSemanticsToDocBook = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "[",
				},
				&asciidoc.Anchor{
					ID:  "taoup",
					Set: nil,
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
						Set: asciidoc.Set{
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
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "[",
				},
				&asciidoc.Anchor{
					ID:  "walsh-muellner",
					Set: nil,
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
					Set: asciidoc.Set{
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

var shouldWarnIfABibliographyIdIsAlreadyInUse = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "[",
				},
				&asciidoc.Anchor{
					ID:  "Fowler",
					Set: nil,
				},
				&asciidoc.String{
					Value: "] Fowler M. ",
				},
				&asciidoc.Italic{
					AttributeList: nil,
					Set: asciidoc.Set{
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
						Set: asciidoc.Set{
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
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "[",
				},
				&asciidoc.Anchor{
					ID:  "Fowler",
					Set: nil,
				},
				&asciidoc.String{
					Value: "] Fowler M. ",
				},
				&asciidoc.Italic{
					AttributeList: nil,
					Set: asciidoc.Set{
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

var shouldAutomaticallyAddBibliographyStyleToTopLevelListsInBibliographySection = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Set: asciidoc.Set{
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
			Set: asciidoc.Set{
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.UnorderedListItem{
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "[",
						},
						&asciidoc.Anchor{
							ID:  "taoup",
							Set: nil,
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
							Val: asciidoc.Set{
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
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "[",
						},
						&asciidoc.Anchor{
							ID:  "walsh-muellner",
							Set: nil,
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
							Set: asciidoc.Set{
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
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.UnorderedListItem{
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "[",
						},
						&asciidoc.Anchor{
							ID:  "doc-writer",
							Set: nil,
						},
						&asciidoc.String{
							Value: "] Doc Writer. ",
						},
						&asciidoc.Italic{
							AttributeList: nil,
							Set: asciidoc.Set{
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
							Val: asciidoc.Set{
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
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Bibliography",
				},
			},
			Level: 1,
		},
	},
}

var shouldNotRecognizeBibliographyAnchorThatBeginsWithADigit = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "[[[1984]]] George Orwell. ",
				},
				&asciidoc.Italic{
					AttributeList: nil,
					Set: asciidoc.Set{
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
						Set: asciidoc.Set{
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

var shouldRecognizeBibliographyAnchorThatContainsADigitButDoesNotStartWithOne = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "[",
				},
				&asciidoc.Anchor{
					ID:  "_1984",
					Set: nil,
				},
				&asciidoc.String{
					Value: "] George Orwell. ",
				},
				&asciidoc.DoubleItalic{
					AttributeList: nil,
					Set: asciidoc.Set{
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
						Set: asciidoc.Set{
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

var shouldCatalogBibliographyAnchorsInBibliographyList = &asciidoc.Document{
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
				&asciidoc.String{
					Value: "Please read ",
				},
				&asciidoc.CrossReference{
					Set: nil,
					ID:  "Fowler_1997",
				},
				&asciidoc.String{
					Value: ".",
				},
				&asciidoc.NewLine{},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.Section{
					AttributeList: asciidoc.AttributeList{
						&asciidoc.ShorthandAttribute{
							Style: &asciidoc.ShorthandStyle{
								Set: asciidoc.Set{
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
					Set: asciidoc.Set{
						asciidoc.EmptyLine{
							Text: "",
						},
						&asciidoc.UnorderedListItem{
							Set: asciidoc.Set{
								&asciidoc.String{
									Value: "[",
								},
								&asciidoc.Anchor{
									ID:  "Fowler_1997",
									Set: nil,
								},
								&asciidoc.String{
									Value: "] Fowler M. ",
								},
								&asciidoc.Italic{
									AttributeList: nil,
									Set: asciidoc.Set{
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
					Title: asciidoc.Set{
						&asciidoc.String{
							Value: "References",
						},
					},
					Level: 1,
				},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Article Title",
				},
			},
			Level: 0,
		},
	},
}

var shouldUseReftextFromBibliographyAnchorAtXrefAndEntry = &asciidoc.Document{
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
				&asciidoc.String{
					Value: "Begin with ",
				},
				&asciidoc.CrossReference{
					Set: nil,
					ID:  "TMMM",
				},
				&asciidoc.String{
					Value: ".",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "Then move on to ",
				},
				&asciidoc.CrossReference{
					Set: nil,
					ID:  "Fowler_1997",
				},
				&asciidoc.String{
					Value: ".",
				},
				&asciidoc.NewLine{},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.Section{
					AttributeList: asciidoc.AttributeList{
						&asciidoc.ShorthandAttribute{
							Style: &asciidoc.ShorthandStyle{
								Set: asciidoc.Set{
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
					Set: asciidoc.Set{
						asciidoc.EmptyLine{
							Text: "",
						},
						&asciidoc.UnorderedListItem{
							Set: asciidoc.Set{
								&asciidoc.String{
									Value: "[",
								},
								&asciidoc.Anchor{
									ID:  "TMMM",
									Set: nil,
								},
								&asciidoc.String{
									Value: "] Brooks F. ",
								},
								&asciidoc.Italic{
									AttributeList: nil,
									Set: asciidoc.Set{
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
							Set: asciidoc.Set{
								&asciidoc.String{
									Value: "[",
								},
								&asciidoc.Anchor{
									ID: "Fowler_1997",
									Set: asciidoc.Set{
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
									Set: asciidoc.Set{
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
					Title: asciidoc.Set{
						&asciidoc.String{
							Value: "References",
						},
					},
					Level: 1,
				},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Article Title",
				},
			},
			Level: 0,
		},
	},
}

var shouldAssignReftextOfBibliographyAnchorToXreflabelInDocBookBackend = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "[",
				},
				&asciidoc.Anchor{
					ID: "Fowler_1997",
					Set: asciidoc.Set{
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
					Set: asciidoc.Set{
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
						Set: asciidoc.Set{
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

var foldsTextFromSubsequentLine = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Set: asciidoc.Set{
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "def1",
						},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Set{
						&asciidoc.String{
							Value: "term",
						},
						&asciidoc.String{
							Value: "1",
						},
					},
				},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var foldsTextFromFirstLineAfterBlankLines = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Set: asciidoc.Set{
						&asciidoc.NewLine{},
						&asciidoc.NewLine{},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "def1",
						},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Set{
						&asciidoc.String{
							Value: "term",
						},
						&asciidoc.String{
							Value: "1",
						},
					},
				},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var foldsTextFromFirstLineAfterBlankLineAndImmediatelyPrecedingNextItem = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Set: asciidoc.Set{
						&asciidoc.NewLine{},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "def1",
						},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Set{
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
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var paragraphOffsetByBlankLinesDoesNotBreakListIfLabelDoesNotHaveInlineText = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Set: asciidoc.Set{
						&asciidoc.NewLine{},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "def1",
						},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Set{
						&asciidoc.String{
							Value: "term",
						},
						&asciidoc.String{
							Value: "1",
						},
					},
				},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "term2:: def2",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var foldsTextFromFirstLineAfterCommentLine = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Set: asciidoc.Set{
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "// comment",
						},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Set{
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
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var foldsTextFromLineFollowingCommentLineOffsetByBlankLine = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Set: asciidoc.Set{
						&asciidoc.NewLine{},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "// comment",
						},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Set{
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
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var foldsTextFromSubsequentIndentedLine = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Set: asciidoc.Set{
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "  def1",
						},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Set{
						&asciidoc.String{
							Value: "term",
						},
						&asciidoc.String{
							Value: "1",
						},
					},
				},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var foldsTextFromIndentedLineAfterBlankLine = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Set: asciidoc.Set{
						&asciidoc.NewLine{},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "  def1",
						},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Set{
						&asciidoc.String{
							Value: "term",
						},
						&asciidoc.String{
							Value: "1",
						},
					},
				},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var foldsTextThatLooksLikeRulerOffsetByBlankLine = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Set: asciidoc.Set{
						&asciidoc.NewLine{},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "'''",
						},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Set{
						&asciidoc.String{
							Value: "term",
						},
						&asciidoc.String{
							Value: "1",
						},
					},
				},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var foldsTextThatLooksLikeRulerOffsetByBlankLineAndLineComment = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Set: asciidoc.Set{
						&asciidoc.NewLine{},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "// comment",
						},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Set{
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
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var foldsTextThatLooksLikeRulerAndTheLineFollowingItOffsetByBlankLine = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Set: asciidoc.Set{
						&asciidoc.NewLine{},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "'''",
						},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Set{
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
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var foldsTextThatLooksLikeTitleOffsetByBlankLine = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Set: asciidoc.Set{
						&asciidoc.NewLine{},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: ".def1",
						},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Set{
						&asciidoc.String{
							Value: "term",
						},
						&asciidoc.String{
							Value: "1",
						},
					},
				},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var foldsTextThatLooksLikeTitleOffsetByBlankLineAndLineComment = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Set: asciidoc.Set{
						&asciidoc.NewLine{},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "// comment",
						},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Set{
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
							Val: asciidoc.Set{
								&asciidoc.String{
									Value: "def1",
								},
							},
						},
					},
					Set:        asciidoc.Set{},
					Admonition: 0,
				},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var foldsTextThatLooksLikeAdmonitionOffsetByBlankLine = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Set: asciidoc.Set{
						&asciidoc.NewLine{},
						&asciidoc.NewLine{},
						&asciidoc.Paragraph{
							AttributeList: nil,
							Set: asciidoc.Set{
								&asciidoc.String{
									Value: "def1",
								},
							},
							Admonition: 1,
						},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Set{
						&asciidoc.String{
							Value: "term",
						},
						&asciidoc.String{
							Value: "1",
						},
					},
				},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var foldsTextThatLooksLikeSectionTitleOffsetByBlankLine = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Set: asciidoc.Set{
						&asciidoc.NewLine{},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "== Another Section",
						},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Set{
						&asciidoc.String{
							Value: "term",
						},
						&asciidoc.String{
							Value: "1",
						},
					},
				},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var foldsTextOfFirstLiteralLineOffsetByBlankLineAppendsSubsequentLiteralsOffsetByBlankLineAsBlocks = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Set: asciidoc.Set{
						&asciidoc.NewLine{},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "  def1",
						},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Set{
						&asciidoc.String{
							Value: "term",
						},
						&asciidoc.String{
							Value: "1",
						},
					},
				},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "  literal",
				},
				&asciidoc.NewLine{},
				asciidoc.EmptyLine{
					Text: "",
				},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "  literal",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var foldsTextOfSubsequentLineAndAppendsFollowingLiteralLineOffsetByBlankLineAsBlockIfTermHasNoInlineDescription = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Set: asciidoc.Set{
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "def1",
						},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Set{
						&asciidoc.String{
							Value: "term",
						},
						&asciidoc.String{
							Value: "1",
						},
					},
				},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "  literal",
				},
				&asciidoc.NewLine{},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "term2:: def2",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var appendsLiteralLineAttachedByContinuationAsBlockIfItemHasNoInlineDescription = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Set: asciidoc.Set{
						&asciidoc.NewLine{},
						&asciidoc.LineBreak{},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Set{
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
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var appendsLiteralLineAttachedByContinuationAsBlockIfItemHasNoInlineDescriptionFollowedByRuler = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Set: asciidoc.Set{
						&asciidoc.NewLine{},
						&asciidoc.LineBreak{},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Set{
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
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.ThematicBreak{
					AttributeList: nil,
				},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var appendsLineAttachedByContinuationAsBlockIfItemHasNoInlineDescriptionFollowedByRuler = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Set: asciidoc.Set{
						&asciidoc.NewLine{},
						&asciidoc.LineBreak{},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Set{
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
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.ThematicBreak{
					AttributeList: nil,
				},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var appendsLineAttachedByContinuationAsBlockIfItemHasNoInlineDescriptionFollowedByBlock = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Set: asciidoc.Set{
						&asciidoc.NewLine{},
						&asciidoc.LineBreak{},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Set{
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
				asciidoc.EmptyLine{
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
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var appendsBlockAttachedByContinuationButNotSubsequentBlockNotAttachedByContinuation = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Set: asciidoc.Set{
						&asciidoc.NewLine{},
						&asciidoc.LineBreak{},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Set{
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
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var appendsListIfItemHasNoInlineDescription = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Set: asciidoc.Set{
						&asciidoc.NewLine{},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "* one",
						},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Set{
						&asciidoc.String{
							Value: "term",
						},
						&asciidoc.String{
							Value: "1",
						},
					},
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
					Checklist:     0,
				},
				&asciidoc.UnorderedListItem{
					Set: asciidoc.Set{
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
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var appendsListToFirstTermWhenFollowedImmediatelyBySecondTerm = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Set: asciidoc.Set{
						&asciidoc.NewLine{},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "* one",
						},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Set{
						&asciidoc.String{
							Value: "term",
						},
						&asciidoc.String{
							Value: "1",
						},
					},
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
					Checklist:     0,
				},
				&asciidoc.UnorderedListItem{
					Set: asciidoc.Set{
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
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var appendsIndentedListToFirstTermThatIsAdjacentToSecondTerm = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Set: asciidoc.Set{
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "  description 1",
						},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Set{
						&asciidoc.String{
							Value: "label ",
						},
						&asciidoc.String{
							Value: "1",
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
					Indent:        "  ",
					Marker:        "*",
					Checklist:     0,
				},
				&asciidoc.UnorderedListItem{
					Set: asciidoc.Set{
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
					Set: asciidoc.Set{
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
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var appendsIndentedListToFirstTermThatIsAttachedByAContinuationAndAdjacentToSecondTerm = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Set: asciidoc.Set{
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "  description 1",
						},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Set{
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
					Set: asciidoc.Set{
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
					Set: asciidoc.Set{
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
					Set: asciidoc.Set{
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
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var appendsListAndParagraphBlockWhenLineFollowingListAttachedByContinuation = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Set: asciidoc.Set{
						&asciidoc.NewLine{},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "* one",
						},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Set{
						&asciidoc.String{
							Value: "term",
						},
						&asciidoc.String{
							Value: "1",
						},
					},
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
					Checklist:     0,
				},
				&asciidoc.UnorderedListItem{
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "three",
						},
					},
					AttributeList: nil,
					Indent:        "",
					Marker:        "*",
					Checklist:     0,
				},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.LineBreak{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "para",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var firstContinuedLineAssociatedWithNestedListItemAndSecondContinuedLineAssociatedWithTerm = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Set: asciidoc.Set{
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "* one",
						},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Set{
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
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.LineBreak{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "term1 para",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var literalLineAttachedByContinuationSwallowsAdjacentLineThatLooksLikeTerm = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Set: asciidoc.Set{
						&asciidoc.NewLine{},
						&asciidoc.LineBreak{},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Set{
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
					Set: asciidoc.Set{
						&asciidoc.NewLine{},
						&asciidoc.LineBreak{},
					},
					AttributeList: nil,
					Marker:        ":::",
					Term: asciidoc.Set{
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
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var lineAttachedByContinuationIsAppendedAsParagraphIfTermHasNoInlineDescription = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Set: asciidoc.Set{
						&asciidoc.NewLine{},
						&asciidoc.LineBreak{},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Set{
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
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var attachedParagraphDoesNotBreakOnAdjacentNestedDescriptionListTerm = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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

var attachedParagraphIsTerminatedByAdjacentSiblingDescriptionListTerm = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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

var attachedStyledParagraphDoesNotBreakOnAdjacentNestedDescriptionListTerm = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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
			},
			Set: asciidoc.Set{
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

var appendsLineAsParagraphIfAttachedByContinuationFollowingBlankLineAndLineCommentWhenTermHasNoInlineDescription = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Set: asciidoc.Set{
						&asciidoc.NewLine{},
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "// comment",
						},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Set{
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
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var lineAttachedByContinuationOffsetByBlankLineIsAppendedAsParagraphIfTermHasNoInlineDescription = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Set: asciidoc.Set{
						&asciidoc.NewLine{},
						&asciidoc.NewLine{},
						&asciidoc.LineBreak{},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Set{
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
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var delimitedBlockBreaksListEvenWhenTermHasNoInlineDescription = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Set: asciidoc.Set{
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "====",
						},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Set{
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
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var blockAttributeLineAboveDelimitedBlockThatBreaksADlistIsNotDuplicated = &asciidoc.Document{
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
									Set: asciidoc.Set{
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
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var blockAttributeLineAboveParagraphBreaksListEvenWhenTermHasNoInlineDescription = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Set: asciidoc.Set{
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "[verse]",
						},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Set{
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
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var blockAttributeLineAboveParagraphThatBreaksADlistIsNotDuplicated = &asciidoc.Document{
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
									Set: asciidoc.Set{
										&asciidoc.String{
											Value: "rolename",
										},
									},
								},
							},
							Options: nil,
						},
					},
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "detached",
						},
						&asciidoc.NewLine{},
					},
					Admonition: 0,
				},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var blockAnchorLineBreaksListEvenWhenTermHasNoInlineDescription = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Set: asciidoc.Set{
						&asciidoc.NewLine{},
						&asciidoc.Anchor{
							ID:  "id",
							Set: nil,
						},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Set{
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
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var blockAttributeLinesAboveNestedHorizontalListDoesNotBreakList = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Set: asciidoc.Set{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "[horizontal]",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Set{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Set: asciidoc.Set{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "  PaaS::: OpenShift",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Set{
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

var blockAttributeLinesAboveNestedListWithStyleDoesNotBreakList = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Set: asciidoc.Set{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "* get groceries",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Set{
				&asciidoc.String{
					Value: "TODO List",
				},
			},
		},
		&asciidoc.DescriptionListItem{
			Set: asciidoc.Set{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "[square]",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Set{
				&asciidoc.String{
					Value: "Grocery List",
				},
			},
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
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

var multipleBlockAttributeLinesAboveNestedListDoesNotBreakList = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Set: asciidoc.Set{
				&asciidoc.NewLine{},
				&asciidoc.Anchor{
					ID:  "variants",
					Set: nil,
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Set{
				&asciidoc.String{
					Value: "Operating Systems",
				},
			},
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
			},
			Set: asciidoc.Set{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Set: asciidoc.Set{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "  PaaS::: OpenShift",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Set{
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

var multipleBlockAttributeLinesSeparatedByEmptyLineAboveNestedListDoesNotBreakList = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Set: asciidoc.Set{
				&asciidoc.NewLine{},
				&asciidoc.Anchor{
					ID:  "variants",
					Set: nil,
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Set{
				&asciidoc.String{
					Value: "Operating Systems",
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
			},
			Set:        asciidoc.Set{},
			Admonition: 0,
		},
		asciidoc.EmptyLine{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Set: asciidoc.Set{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "  PaaS::: OpenShift",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Set{
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

var foldsTextFromInlineDescriptionAndSubsequentLine = &asciidoc.Document{
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
					Value: "term1:: def1",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "continued",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var foldsTextFromInlineDescriptionAndSubsequentLines = &asciidoc.Document{
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
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var foldsTextFromInlineDescriptionAndLineFollowingCommentLine = &asciidoc.Document{
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
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var foldsTextFromInlineDescriptionAndSubsequentIndentedLine = &asciidoc.Document{
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
					Value: "term1:: def1",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "  continued",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "List",
				},
			},
			Level: 1,
		},
	},
}

var appendsLiteralLineOffsetByBlankLineAsBlockIfItemHasInlineDescription = &asciidoc.Document{
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
					Value: "term1:: def1",
				},
				&asciidoc.NewLine{},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "  literal",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var appendsLiteralLineOffsetByBlankLineAsBlockAndAppendsLineAfterContinuationAsBlockIfItemHasInlineDescription = &asciidoc.Document{
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
					Value: "term1:: def1",
				},
				&asciidoc.NewLine{},
				asciidoc.EmptyLine{
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
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var appendsLineAfterContinuationAsBlockAndLiteralLineOffsetByBlankLineAsBlockIfItemHasInlineDescription = &asciidoc.Document{
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
					Value: "term1:: def1",
				},
				&asciidoc.NewLine{},
				&asciidoc.LineBreak{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "para",
				},
				&asciidoc.NewLine{},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "  literal",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var appendsListIfItemHasInlineDescription = &asciidoc.Document{
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
					Value: "term1:: def1",
				},
				&asciidoc.NewLine{},
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
					Checklist:     0,
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
					Checklist:     0,
				},
				&asciidoc.UnorderedListItem{
					Set: asciidoc.Set{
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
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var appendsLiteralLineAttachedByContinuationAsBlockIfItemHasInlineDescriptionFollowedByRuler = &asciidoc.Document{
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
					Value: "term1:: def1",
				},
				&asciidoc.NewLine{},
				&asciidoc.LineBreak{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "  literal",
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
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var lineOffsetByBlankLineBreaksListIfTermHasInlineDescription = &asciidoc.Document{
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
					Value: "term1:: def1",
				},
				&asciidoc.NewLine{},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "detached",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var nestedTermWithDescriptionDoesNotConsumeFollowingHeading = &asciidoc.Document{
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
				&asciidoc.DescriptionListItem{
					Set: asciidoc.Set{
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "  def",
						},
					},
					AttributeList: nil,
					Marker:        "::",
					Term: asciidoc.Set{
						&asciidoc.String{
							Value: "term",
						},
					},
				},
				&asciidoc.DescriptionListItem{
					Set: asciidoc.Set{
						&asciidoc.NewLine{},
						&asciidoc.String{
							Value: "    nesteddef",
						},
					},
					AttributeList: nil,
					Marker:        ";;",
					Term: asciidoc.Set{
						&asciidoc.String{
							Value: "  nestedterm",
						},
					},
				},
				asciidoc.EmptyLine{
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
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var lineAttachedByContinuationIsAppendedAsParagraphIfTermHasInlineDescriptionFollowedByDetachedParagraph = &asciidoc.Document{
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
					Value: "term1:: def1",
				},
				&asciidoc.NewLine{},
				&asciidoc.LineBreak{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "para",
				},
				&asciidoc.NewLine{},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "detached",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var lineAttachedByContinuationIsAppendedAsParagraphIfTermHasInlineDescriptionFollowedByDetachedBlock = &asciidoc.Document{
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
					Value: "term1:: def1",
				},
				&asciidoc.NewLine{},
				&asciidoc.LineBreak{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "para",
				},
				&asciidoc.NewLine{},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.SidebarBlock{
					Delimiter: asciidoc.Delimiter{
						Type:   8,
						Length: 4,
					},
					AttributeList: nil,
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "detached",
						},
						&asciidoc.NewLine{},
					},
				},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var lineAttachedByContinuationOffsetByLineCommentIsAppendedAsParagraphIfTermHasInlineDescription = &asciidoc.Document{
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
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var lineAttachedByContinuationOffsetByBlankLineIsAppendedAsParagraphIfTermHasInlineDescription = &asciidoc.Document{
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
					Value: "term1:: def1",
				},
				&asciidoc.NewLine{},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.LineBreak{},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "para",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var lineCommentOffsetByBlankLineDividesListsBecauseItemHasText = &asciidoc.Document{
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
					Value: "term1:: def1",
				},
				&asciidoc.NewLine{},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.SingleLineComment{
					Value: "",
				},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "term2:: def2",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var rulerOffsetByBlankLineDividesListsBecauseItemHasText = &asciidoc.Document{
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
					Value: "term1:: def1",
				},
				&asciidoc.NewLine{},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.ThematicBreak{
					AttributeList: nil,
				},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "term2:: def2",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var blockTitleOffsetByBlankLineDividesListsAndBecomesTitleOfSecondListBecauseItemHasText = &asciidoc.Document{
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
					Value: "term1:: def1",
				},
				&asciidoc.NewLine{},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.Paragraph{
					AttributeList: asciidoc.AttributeList{
						&asciidoc.TitleAttribute{
							Val: asciidoc.Set{
								&asciidoc.String{
									Value: "title",
								},
							},
						},
					},
					Set:        asciidoc.Set{},
					Admonition: 0,
				},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "term2:: def2",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Lists",
				},
			},
			Level: 1,
		},
	},
}

var doesNotRecognizeCalloutListDenotedByMarkersThatOnlyHaveATrailingBracket = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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

var shouldNotHangIfObsoleteCalloutListIsFoundInsideListItem = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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

var shouldNotHangIfObsoleteCalloutListIsFoundInsideDlistItem = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Set: asciidoc.Set{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "1> bar",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Set{
				&asciidoc.String{
					Value: "foo",
				},
			},
		},
	},
}

var shouldRecognizeAutoNumberdCalloutListInsideList = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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
			Set: asciidoc.Set{
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

var listingBlockWithSequentialCalloutsFollowedByAdjacentCalloutList = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Listing{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Set: asciidoc.Set{
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
					Val: asciidoc.Set{
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

var listingBlockWithSequentialCalloutsFollowedByNonAdjacentCalloutList = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Listing{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Set: asciidoc.Set{
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
					Val: asciidoc.Set{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "Paragraph.",
		},
		&asciidoc.NewLine{},
		asciidoc.EmptyLine{
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

var listingBlockWithACalloutThatRefersToTwoDifferentLines = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Listing{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Set: asciidoc.Set{
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
					Val: asciidoc.Set{
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

var sourceBlockWithNonSequentialCalloutsFollowedByAdjacentCalloutList = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Listing{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Set: asciidoc.Set{
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
					Val: asciidoc.Set{
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

var twoListingBlocksCanShareTheSameCalloutList = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Listing{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.TitleAttribute{
					Val: asciidoc.Set{
						&asciidoc.String{
							Value: "Import library",
						},
					},
				},
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Set: asciidoc.Set{
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
					Val: asciidoc.Set{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Listing{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.TitleAttribute{
					Val: asciidoc.Set{
						&asciidoc.String{
							Value: "Use library",
						},
					},
				},
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Set: asciidoc.Set{
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
					Val: asciidoc.Set{
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
		asciidoc.EmptyLine{
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

var twoListingBlocksEachFollowedByAnAdjacentCalloutList = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Listing{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.TitleAttribute{
					Val: asciidoc.Set{
						&asciidoc.String{
							Value: "Import library",
						},
					},
				},
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Set: asciidoc.Set{
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
					Val: asciidoc.Set{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Listing{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.TitleAttribute{
					Val: asciidoc.Set{
						&asciidoc.String{
							Value: "Use library",
						},
					},
				},
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Set: asciidoc.Set{
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
					Val: asciidoc.Set{
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

var calloutListRetainsBlockContent = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Listing{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Set: asciidoc.Set{
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
					Val: asciidoc.Set{
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
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
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

var calloutListRetainsBlockContentWhenConvertedToDocBook = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Listing{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Set: asciidoc.Set{
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
					Val: asciidoc.Set{
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
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
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

var escapedCalloutShouldNotBeInterpretedAsACallout = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Listing{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Set: asciidoc.Set{
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
					Val: asciidoc.Set{
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

var shouldAutonumberCallouts = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Listing{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Set: asciidoc.Set{
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
					Val: asciidoc.Set{
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

var shouldNotRecognizeCalloutsInMiddleOfLine = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Listing{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Set: asciidoc.Set{
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
					Val: asciidoc.Set{
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

var shouldAllowMultipleCalloutsOnTheSameLine = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Listing{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Set: asciidoc.Set{
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
					Val: asciidoc.Set{
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

var shouldAllowXmlCommentStyleCallouts = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Listing{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Set: asciidoc.Set{
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
					Val: asciidoc.Set{
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

var shouldNotAllowCalloutsWithHalfAnXmlComment = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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

var shouldNotRecognizeCalloutsInAnIndentedDescriptionListParagraph = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.DescriptionListItem{
			Set: asciidoc.Set{
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "  bar <1>",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Set{
				&asciidoc.String{
					Value: "foo",
				},
			},
		},
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "<1> Not pointing to a callout",
		},
		&asciidoc.NewLine{},
	},
}

var shouldNotRecognizeCalloutsInAnIndentedOutlineListParagraph = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "<1> Not pointing to a callout",
		},
		&asciidoc.NewLine{},
	},
}

var shouldWarnIfNumbersInCalloutListAreOutOfSequence = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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

var shouldPreserveLineCommentCharsThatPrecedeCalloutNumberIfIconsIsNotSet = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Listing{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Set: asciidoc.Set{
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
					Val: asciidoc.Set{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Listing{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Set: asciidoc.Set{
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
					Val: asciidoc.Set{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Listing{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Set: asciidoc.Set{
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
					Val: asciidoc.Set{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Listing{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Set: asciidoc.Set{
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
					Val: asciidoc.Set{
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

var shouldRemoveLineCommentCharsThatPrecedeCalloutNumberIfIconsIsFont = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Listing{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Set: asciidoc.Set{
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
					Val: asciidoc.Set{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Listing{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Set: asciidoc.Set{
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
					Val: asciidoc.Set{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Listing{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Set: asciidoc.Set{
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
					Val: asciidoc.Set{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Listing{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Set: asciidoc.Set{
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
					Val: asciidoc.Set{
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

var shouldAllowLineCommentCharsThatPrecedeCalloutNumberToBeSpecified = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Listing{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Set: asciidoc.Set{
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
					Val: asciidoc.Set{
						&asciidoc.String{
							Value: "erlang",
						},
					},
				},
				&asciidoc.NamedAttribute{
					Name: "line-comment",
					Val: asciidoc.Set{
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

var shouldAllowLineCommentCharsPrecedingCalloutNumberToBeConfigurableWhenSourceHighlighterIsCoderay = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Listing{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Set: asciidoc.Set{
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
					Val: asciidoc.Set{
						&asciidoc.String{
							Value: "html",
						},
					},
				},
				&asciidoc.NamedAttribute{
					Name: "line-comment",
					Val: asciidoc.Set{
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

var shouldNotEatWhitespaceBeforeCalloutNumberIfLineCommentAttributeIsEmpty = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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

var literalBlockWithCallouts = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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
		asciidoc.EmptyLine{
			Text: "",
		},
		asciidoc.EmptyLine{
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

var calloutListWithIconsEnabled = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Listing{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Set: asciidoc.Set{
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
					Val: asciidoc.Set{
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

var calloutListWithFontBasedIconsEnabled = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Listing{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Set: asciidoc.Set{
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

var shouldCreateChecklistIfAtLeastOneItemHasCheckboxSyntax = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
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

var shouldCreateChecklistWithFontIconsIfAtLeastOneItemHasCheckboxSyntaxAndIconsAttributeIsFont = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
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

var shouldCreateInteractiveChecklistIfInteractiveOptionIsSetEvenWithIconsAttributeIsFont = &asciidoc.Document{
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
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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
							Set: asciidoc.Set{
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
			Set: asciidoc.Set{
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

var contentShouldReturnItemsInList = &asciidoc.Document{
	Set: asciidoc.Set{
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
			Checklist:     0,
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
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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

var listItemShouldBeTheParentOfBlockAttachedToAListItem = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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

var outlineShouldReturnTrueForUnorderedList = &asciidoc.Document{
	Set: asciidoc.Set{
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
			Checklist:     0,
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
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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

var outlineShouldReturnTrueForOrderedList = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OrderedListItem{
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "one",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
		&asciidoc.OrderedListItem{
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "two",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
		&asciidoc.OrderedListItem{
			Set: asciidoc.Set{
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

var outlineShouldReturnFalseForDescriptionList = &asciidoc.Document{
	Set: asciidoc.Set{
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
			Checklist:     0,
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
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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

var simpleShouldReturnTrueForListItemWithNestedOutlineList = &asciidoc.Document{
	Set: asciidoc.Set{
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
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
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

var simpleShouldReturnFalseForListItemWithBlockContent = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
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

var shouldAllowTextOfListItemToBeAssigned = &asciidoc.Document{
	Set: asciidoc.Set{
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
			Checklist:     0,
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
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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

var idAndRoleAssignedToUlistItemInModelAreTransmittedToOutput = &asciidoc.Document{
	Set: asciidoc.Set{
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
			Checklist:     0,
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
			Checklist:     0,
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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

var idAndRoleAssignedToOlistItemInModelAreTransmittedToOutput = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OrderedListItem{
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "one",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
		&asciidoc.OrderedListItem{
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "two",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
		&asciidoc.OrderedListItem{
			Set: asciidoc.Set{
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

var shouldAllowApiControlOverSubstitutionsAppliedToListItemText = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
				&asciidoc.Bold{
					AttributeList: nil,
					Set: asciidoc.Set{
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
			Set: asciidoc.Set{
				&asciidoc.Italic{
					AttributeList: nil,
					Set: asciidoc.Set{
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
			Set: asciidoc.Set{
				&asciidoc.Monospace{
					AttributeList: nil,
					Set: asciidoc.Set{
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
			Set: asciidoc.Set{
				&asciidoc.Marked{
					AttributeList: nil,
					Set: asciidoc.Set{
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

var shouldSetLinenoToLineNumberInSourceWhereListStarts = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
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
			Set: asciidoc.Set{
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
