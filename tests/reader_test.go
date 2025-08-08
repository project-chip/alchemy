package tests

import (
	"testing"

	"github.com/project-chip/alchemy/asciidoc"
)

func TestReader(t *testing.T) {
	readerTests.run(t)
}

var readerTests = parseTests{

	{"should prepare lines from Array data", "asciidoctor/reader_test_should_prepare_lines_from_array_data.adoc", readerTestShouldPrepareLinesFromArrayData, nil},

	{"Read lines until until blank line", "asciidoctor/reader_test_read_lines_until_until_blank_line.adoc", readerTestReadLinesUntilUntilBlankLine, nil},

	{"Read lines until until blank line preserving last line", "asciidoctor/reader_test_read_lines_until_until_blank_line_preserving_last_line.adoc", readerTestReadLinesUntilUntilBlankLinePreservingLastLine, nil},

	{"Read lines until until condition is true", "asciidoctor/reader_test_read_lines_until_until_condition_is_true.adoc", readerTestReadLinesUntilUntilConditionIsTrue, nil},

	{"Read lines until until condition is true, taking last line", "asciidoctor/reader_test_read_lines_until_until_condition_is_true_taking_last_line.adoc", readerTestReadLinesUntilUntilConditionIsTrueTakingLastLine, nil},

	{"Read lines until until condition is true, taking and preserving last line", "asciidoctor/reader_test_read_lines_until_until_condition_is_true_taking_and_preserving_last_line.adoc", readerTestReadLinesUntilUntilConditionIsTrueTakingAndPreservingLastLine, nil},

	{"read lines until terminator", "asciidoctor/reader_test_read_lines_until_terminator.adoc", readerTestReadLinesUntilTerminator, nil},

	{"should flag reader as unterminated if reader reaches end of source without finding terminator", "asciidoctor/reader_test_should_flag_reader_as_unterminated_if_reader_reaches_end_of_source_without_finding_terminator.adoc", readerTestShouldFlagReaderAsUnterminatedIfReaderReachesEndOfSourceWithoutFindingTerminator, nil},

	{"should not skip front matter by default", "asciidoctor/reader_test_should_not_skip_front_matter_by_default.adoc", readerTestShouldNotSkipFrontMatterByDefault, nil},

	{"should not skip front matter if ending delimiter is not found", "asciidoctor/reader_test_should_not_skip_front_matter_if_ending_delimiter_is_not_found.adoc", readerTestShouldNotSkipFrontMatterIfEndingDelimiterIsNotFound, nil},

	{"should skip front matter if specified by skip-front-matter attribute", "asciidoctor/reader_test_should_skip_front_matter_if_specified_by_skip_front_matter_attribute.adoc", readerTestShouldSkipFrontMatterIfSpecifiedBySkipFrontMatterAttribute, nil},

	{"should skip TOML front matter if specified by skip-front-matter attribute", "asciidoctor/reader_test_should_skip_toml_front_matter_if_specified_by_skip_front_matter_attribute.adoc", readerTestShouldSkipTomlFrontMatterIfSpecifiedBySkipFrontMatterAttribute, nil},

	{"should not track include in catalog for non-AsciiDoc include files", "asciidoctor/reader_test_should_not_track_include_in_catalog_for_non_ascii_doc_include_files.adoc", readerTestShouldNotTrackIncludeInCatalogForNonAsciiDocIncludeFiles, nil},

	{"include directive should resolve file with spaces in name", "asciidoctor/reader_test_include_directive_should_resolve_file_with_spaces_in_name.adoc", readerTestIncludeDirectiveShouldResolveFileWithSpacesInName, nil},

	{"include directive should resolve file relative to current include", "asciidoctor/reader_test_include_directive_should_resolve_file_relative_to_current_include.adoc", readerTestIncludeDirectiveShouldResolveFileRelativeToCurrentInclude, nil},

	{"should fail to read include file if not UTF-8 encoded and encoding is not specified", "asciidoctor/reader_test_should_fail_to_read_include_file_if_not_utf_8_encoded_and_encoding_is_not_specified.adoc", readerTestShouldFailToReadIncludeFileIfNotUtf8EncodedAndEncodingIsNotSpecified, nil},

	{"should ignore encoding attribute if value is not a valid encoding", "asciidoctor/reader_test_should_ignore_encoding_attribute_if_value_is_not_a_valid_encoding.adoc", readerTestShouldIgnoreEncodingAttributeIfValueIsNotAValidEncoding, nil},

	{"should use encoding specified by encoding attribute when reading include file", "asciidoctor/reader_test_should_use_encoding_specified_by_encoding_attribute_when_reading_include_file.adoc", readerTestShouldUseEncodingSpecifiedByEncodingAttributeWhenReadingIncludeFile, nil},

	{"unresolved target referenced by include directive is skipped when optional option is set", "asciidoctor/reader_test_unresolved_target_referenced_by_include_directive_is_skipped_when_optional_option_is_set.adoc", readerTestUnresolvedTargetReferencedByIncludeDirectiveIsSkippedWhenOptionalOptionIsSet, nil},

	{"should skip include directive that references missing file if optional option is set", "asciidoctor/reader_test_should_skip_include_directive_that_references_missing_file_if_optional_option_is_set.adoc", readerTestShouldSkipIncludeDirectiveThatReferencesMissingFileIfOptionalOptionIsSet, nil},

	{"should replace include directive that references missing file with message", "asciidoctor/reader_test_should_replace_include_directive_that_references_missing_file_with_message.adoc", readerTestShouldReplaceIncludeDirectiveThatReferencesMissingFileWithMessage, nil},

	{"nested include directives are resolved relative to current file", "asciidoctor/reader_test_nested_include_directives_are_resolved_relative_to_current_file.adoc", readerTestNestedIncludeDirectivesAreResolvedRelativeToCurrentFile, nil},

	{"include directive supports selecting lines by line number", "asciidoctor/reader_test_include_directive_supports_selecting_lines_by_line_number.adoc", readerTestIncludeDirectiveSupportsSelectingLinesByLineNumber, nil},

	{"include directive ignores lines attribute with invalid range", "asciidoctor/reader_test_include_directive_ignores_lines_attribute_with_invalid_range.adoc", readerTestIncludeDirectiveIgnoresLinesAttributeWithInvalidRange, nil},

	{"include directive supports selecting lines by tag in file that has CRLF line endings", "asciidoctor/reader_test_include_directive_supports_selecting_lines_by_tag_in_file_that_has_crlf_line_endings.adoc", readerTestIncludeDirectiveSupportsSelectingLinesByTagInFileThatHasCrlfLineEndings, nil},

	{"include directive skips lines inside tag which is negated", "asciidoctor/reader_test_include_directive_skips_lines_inside_tag_which_is_negated.adoc", readerTestIncludeDirectiveSkipsLinesInsideTagWhichIsNegated, nil},

	{"include directive selects all lines without a tag directive when value is double asterisk", "asciidoctor/reader_test_include_directive_selects_all_lines_without_a_tag_directive_when_value_is_double_asterisk.adoc", readerTestIncludeDirectiveSelectsAllLinesWithoutATagDirectiveWhenValueIsDoubleAsterisk, nil},

	{"include directive selects all lines except lines inside tag which is negated when value starts with double asterisk", "asciidoctor/reader_test_include_directive_selects_all_lines_except_lines_inside_tag_which_is_negated_when_value_starts_with_double_asterisk.adoc", readerTestIncludeDirectiveSelectsAllLinesExceptLinesInsideTagWhichIsNegatedWhenValueStartsWithDoubleAsterisk, nil},

	{"include directive selects all lines, including lines inside nested tags, except lines inside tag which is negated when value starts with double asterisk", "asciidoctor/reader_test_include_directive_selects_all_lines_including_lines_inside_nested_tags_except_lines_inside_tag_which_is_negated_when_value_starts_with_double_asterisk.adoc", readerTestIncludeDirectiveSelectsAllLinesIncludingLinesInsideNestedTagsExceptLinesInsideTagWhichIsNegatedWhenValueStartsWithDoubleAsterisk, nil},

	{"include directive selects all lines outside of tags when value is double asterisk followed by negated wildcard", "asciidoctor/reader_test_include_directive_selects_all_lines_outside_of_tags_when_value_is_double_asterisk_followed_by_negated_wildcard.adoc", readerTestIncludeDirectiveSelectsAllLinesOutsideOfTagsWhenValueIsDoubleAsteriskFollowedByNegatedWildcard, nil},

	{"include directive skips all tagged regions when value of tags attribute is negated wildcard", "asciidoctor/reader_test_include_directive_skips_all_tagged_regions_when_value_of_tags_attribute_is_negated_wildcard.adoc", readerTestIncludeDirectiveSkipsAllTaggedRegionsWhenValueOfTagsAttributeIsNegatedWildcard, nil},

	{"include directive selects all lines except for lines containing tag directive if value is double asterisk followed by nested tag names", "asciidoctor/reader_test_include_directive_selects_all_lines_except_for_lines_containing_tag_directive_if_value_is_double_asterisk_followed_by_nested_tag_names.adoc", readerTestIncludeDirectiveSelectsAllLinesExceptForLinesContainingTagDirectiveIfValueIsDoubleAsteriskFollowedByNestedTagNames, nil},

	{"include directive selects all lines except for lines containing tag directive when value is double asterisk followed by outer tag name", "asciidoctor/reader_test_include_directive_selects_all_lines_except_for_lines_containing_tag_directive_when_value_is_double_asterisk_followed_by_outer_tag_name.adoc", readerTestIncludeDirectiveSelectsAllLinesExceptForLinesContainingTagDirectiveWhenValueIsDoubleAsteriskFollowedByOuterTagName, nil},

	{"include directive selects all lines inside unspecified tags when value is negated double asterisk followed by negated tags", "asciidoctor/reader_test_include_directive_selects_all_lines_inside_unspecified_tags_when_value_is_negated_double_asterisk_followed_by_negated_tags.adoc", readerTestIncludeDirectiveSelectsAllLinesInsideUnspecifiedTagsWhenValueIsNegatedDoubleAsteriskFollowedByNegatedTags, nil},

	{"include directive selects all lines except tag which is negated when value only contains negated tag", "asciidoctor/reader_test_include_directive_selects_all_lines_except_tag_which_is_negated_when_value_only_contains_negated_tag.adoc", readerTestIncludeDirectiveSelectsAllLinesExceptTagWhichIsNegatedWhenValueOnlyContainsNegatedTag, nil},

	{"include directive selects all lines except tags which are negated when value only contains negated tags", "asciidoctor/reader_test_include_directive_selects_all_lines_except_tags_which_are_negated_when_value_only_contains_negated_tags.adoc", readerTestIncludeDirectiveSelectsAllLinesExceptTagsWhichAreNegatedWhenValueOnlyContainsNegatedTags, nil},

	{"should recognize tag wildcard if not at start of tags list", "asciidoctor/reader_test_should_recognize_tag_wildcard_if_not_at_start_of_tags_list.adoc", readerTestShouldRecognizeTagWildcardIfNotAtStartOfTagsList, nil},

	{"include directive selects lines between tags when value of tags attribute is wildcard", "asciidoctor/reader_test_include_directive_selects_lines_between_tags_when_value_of_tags_attribute_is_wildcard.adoc", readerTestIncludeDirectiveSelectsLinesBetweenTagsWhenValueOfTagsAttributeIsWildcard, nil},

	{"include directive selects lines inside tags when value of tags attribute is wildcard and tag surrounds content", "asciidoctor/reader_test_include_directive_selects_lines_inside_tags_when_value_of_tags_attribute_is_wildcard_and_tag_surrounds_content.adoc", readerTestIncludeDirectiveSelectsLinesInsideTagsWhenValueOfTagsAttributeIsWildcardAndTagSurroundsContent, nil},

	{"include directive selects lines inside all tags except tag which is negated when value of tags attribute is wildcard followed by negated tag", "asciidoctor/reader_test_include_directive_selects_lines_inside_all_tags_except_tag_which_is_negated_when_value_of_tags_attribute_is_wildcard_followed_by_negated_tag.adoc", readerTestIncludeDirectiveSelectsLinesInsideAllTagsExceptTagWhichIsNegatedWhenValueOfTagsAttributeIsWildcardFollowedByNegatedTag, nil},

	{"include directive includes regions outside tags and inside specified tags when value begins with negated wildcard", "asciidoctor/reader_test_include_directive_includes_regions_outside_tags_and_inside_specified_tags_when_value_begins_with_negated_wildcard.adoc", readerTestIncludeDirectiveIncludesRegionsOutsideTagsAndInsideSpecifiedTagsWhenValueBeginsWithNegatedWildcard, nil},

	{"include directive selects lines inside tag except for lines inside nested tags when tag is preceded by negated double asterisk and negated wildcard", "asciidoctor/reader_test_include_directive_selects_lines_inside_tag_except_for_lines_inside_nested_tags_when_tag_is_preceded_by_negated_double_asterisk_and_negated_wildcard.adoc", readerTestIncludeDirectiveSelectsLinesInsideTagExceptForLinesInsideNestedTagsWhenTagIsPrecededByNegatedDoubleAsteriskAndNegatedWildcard, nil},

	{"include directive does not select lines inside tag that has been included then excluded", "asciidoctor/reader_test_include_directive_does_not_select_lines_inside_tag_that_has_been_included_then_excluded.adoc", readerTestIncludeDirectiveDoesNotSelectLinesInsideTagThatHasBeenIncludedThenExcluded, nil},

	{"include directive selects lines inside specified tag and ignores lines inside a negated tag", "asciidoctor/reader_test_include_directive_selects_lines_inside_specified_tag_and_ignores_lines_inside_a_negated_tag.adoc", readerTestIncludeDirectiveSelectsLinesInsideSpecifiedTagAndIgnoresLinesInsideANegatedTag, nil},

	{"should not warn if specified negated tag is not found in include file", "asciidoctor/reader_test_should_not_warn_if_specified_negated_tag_is_not_found_in_include_file.adoc", readerTestShouldNotWarnIfSpecifiedNegatedTagIsNotFoundInIncludeFile, nil},

	{"should warn if specified tags are not found in include file", "asciidoctor/reader_test_should_warn_if_specified_tags_are_not_found_in_include_file.adoc", readerTestShouldWarnIfSpecifiedTagsAreNotFoundInIncludeFile, nil},

	{"should not warn if specified negated tags are not found in include file", "asciidoctor/reader_test_should_not_warn_if_specified_negated_tags_are_not_found_in_include_file.adoc", readerTestShouldNotWarnIfSpecifiedNegatedTagsAreNotFoundInIncludeFile, nil},

	{"should warn if specified tag in include file is not closed", "asciidoctor/reader_test_should_warn_if_specified_tag_in_include_file_is_not_closed.adoc", readerTestShouldWarnIfSpecifiedTagInIncludeFileIsNotClosed, nil},

	{"should warn if end tag in included file is mismatched", "asciidoctor/reader_test_should_warn_if_end_tag_in_included_file_is_mismatched.adoc", readerTestShouldWarnIfEndTagInIncludedFileIsMismatched, nil},

	{"should warn if unexpected end tag is found in included file", "asciidoctor/reader_test_should_warn_if_unexpected_end_tag_is_found_in_included_file.adoc", readerTestShouldWarnIfUnexpectedEndTagIsFoundInIncludedFile, nil},

	{"lines attribute takes precedence over tags attribute in include directive", "asciidoctor/reader_test_lines_attribute_takes_precedence_over_tags_attribute_in_include_directive.adoc", readerTestLinesAttributeTakesPrecedenceOverTagsAttributeInIncludeDirective, nil},

	{"should substitute attribute references in attrlist", "asciidoctor/reader_test_should_substitute_attribute_references_in_attrlist.adoc", readerTestShouldSubstituteAttributeReferencesInAttrlist, nil},

	{"should fall back to built-in include directive behavior when not handled by include processor", "asciidoctor/reader_test_should_fall_back_to_built_in_include_directive_behavior_when_not_handled_by_include_processor.adoc", readerTestShouldFallBackToBuiltInIncludeDirectiveBehaviorWhenNotHandledByIncludeProcessor, nil},

	{"attributes are substituted in target of include directive", "asciidoctor/reader_test_attributes_are_substituted_in_target_of_include_directive.adoc", readerTestAttributesAreSubstitutedInTargetOfIncludeDirective, nil},

	{"line following dropped include is not dropped", "asciidoctor/reader_test_line_following_dropped_include_is_not_dropped.adoc", readerTestLineFollowingDroppedIncludeIsNotDropped, nil},

	{"escaped include directive is left unprocessed", "asciidoctor/reader_test_escaped_include_directive_is_left_unprocessed.adoc", readerTestEscapedIncludeDirectiveIsLeftUnprocessed, nil},

	{"include directive not at start of line is ignored", "asciidoctor/reader_test_include_directive_not_at_start_of_line_is_ignored.adoc", readerTestIncludeDirectiveNotAtStartOfLineIsIgnored, nil},

	{"include directive should be disabled if max include depth has been exceeded", "asciidoctor/reader_test_include_directive_should_be_disabled_if_max_include_depth_has_been_exceeded.adoc", readerTestIncludeDirectiveShouldBeDisabledIfMaxIncludeDepthHasBeenExceeded, nil},

	{"skip_comment_lines should not process lines read", "asciidoctor/reader_test_skip_comment_lines_should_not_process_lines_read.adoc", readerTestSkipCommentLinesShouldNotProcessLinesRead, nil},

	{"process_line returns nil if cursor advanced", "asciidoctor/reader_test_process_line_returns_nil_if_cursor_advanced.adoc", readerTestProcessLineReturnsNilIfCursorAdvanced, nil},

	{"peek_line advances cursor to next conditional line of content", "asciidoctor/reader_test_peek_line_advances_cursor_to_next_conditional_line_of_content.adoc", readerTestPeekLineAdvancesCursorToNextConditionalLineOfContent, nil},

	{"peek_lines should preprocess lines if direct is false", "asciidoctor/reader_test_peek_lines_should_preprocess_lines_if_direct_is_false.adoc", readerTestPeekLinesShouldPreprocessLinesIfDirectIsFalse, nil},

	{"peek_lines should not preprocess lines if direct is true", "asciidoctor/reader_test_peek_lines_should_not_preprocess_lines_if_direct_is_true.adoc", readerTestPeekLinesShouldNotPreprocessLinesIfDirectIsTrue, nil},

	{"peek_lines should not prevent subsequent preprocessing of peeked lines", "asciidoctor/reader_test_peek_lines_should_not_prevent_subsequent_preprocessing_of_peeked_lines.adoc", readerTestPeekLinesShouldNotPreventSubsequentPreprocessingOfPeekedLines, nil},

	{"process_line returns line if cursor not advanced", "asciidoctor/reader_test_process_line_returns_line_if_cursor_not_advanced.adoc", readerTestProcessLineReturnsLineIfCursorNotAdvanced, nil},

	{"peek_line does not advance cursor when on a regular content line", "asciidoctor/reader_test_peek_line_does_not_advance_cursor_when_on_a_regular_content_line.adoc", readerTestPeekLineDoesNotAdvanceCursorWhenOnARegularContentLine, nil},

	{"peek_line returns nil if cursor advances past end of source", "asciidoctor/reader_test_peek_line_returns_nil_if_cursor_advances_past_end_of_source.adoc", readerTestPeekLineReturnsNilIfCursorAdvancesPastEndOfSource, nil},

	{"peek_line returns nil if contents of skipped conditional is empty line", "asciidoctor/reader_test_peek_line_returns_nil_if_contents_of_skipped_conditional_is_empty_line.adoc", readerTestPeekLineReturnsNilIfContentsOfSkippedConditionalIsEmptyLine, nil},

	{"ifdef with defined attribute includes content", "asciidoctor/reader_test_ifdef_with_defined_attribute_includes_content.adoc", readerTestIfdefWithDefinedAttributeIncludesContent, nil},

	{"ifdef with defined attribute includes text in brackets", "asciidoctor/reader_test_ifdef_with_defined_attribute_includes_text_in_brackets.adoc", readerTestIfdefWithDefinedAttributeIncludesTextInBrackets, nil},

	{"ifdef attribute name is not case sensitive", "asciidoctor/reader_test_ifdef_attribute_name_is_not_case_sensitive.adoc", readerTestIfdefAttributeNameIsNotCaseSensitive, nil},

	{"ifndef with defined attribute does not include text in brackets", "asciidoctor/reader_test_ifndef_with_defined_attribute_does_not_include_text_in_brackets.adoc", readerTestIfndefWithDefinedAttributeDoesNotIncludeTextInBrackets, nil},

	{"include with non-matching nested exclude", "asciidoctor/reader_test_include_with_non_matching_nested_exclude.adoc", readerTestIncludeWithNonMatchingNestedExclude, nil},

	{"nested excludes with same condition", "asciidoctor/reader_test_nested_excludes_with_same_condition.adoc", readerTestNestedExcludesWithSameCondition, nil},

	{"include with nested exclude of inverted condition", "asciidoctor/reader_test_include_with_nested_exclude_of_inverted_condition.adoc", readerTestIncludeWithNestedExcludeOfInvertedCondition, nil},

	{"exclude with matching nested exclude", "asciidoctor/reader_test_exclude_with_matching_nested_exclude.adoc", readerTestExcludeWithMatchingNestedExclude, nil},

	{"exclude with nested include using shorthand end", "asciidoctor/reader_test_exclude_with_nested_include_using_shorthand_end.adoc", readerTestExcludeWithNestedIncludeUsingShorthandEnd, nil},

	{"ifdef with one alternative attribute set includes content", "asciidoctor/reader_test_ifdef_with_one_alternative_attribute_set_includes_content.adoc", readerTestIfdefWithOneAlternativeAttributeSetIncludesContent, nil},

	{"ifdef with no alternative attributes set does not include content", "asciidoctor/reader_test_ifdef_with_no_alternative_attributes_set_does_not_include_content.adoc", readerTestIfdefWithNoAlternativeAttributesSetDoesNotIncludeContent, nil},

	{"ifdef with all required attributes set includes content", "asciidoctor/reader_test_ifdef_with_all_required_attributes_set_includes_content.adoc", readerTestIfdefWithAllRequiredAttributesSetIncludesContent, nil},

	{"ifdef with missing required attributes does not include content", "asciidoctor/reader_test_ifdef_with_missing_required_attributes_does_not_include_content.adoc", readerTestIfdefWithMissingRequiredAttributesDoesNotIncludeContent, nil},

	{"ifndef with undefined attribute includes block", "asciidoctor/reader_test_ifndef_with_undefined_attribute_includes_block.adoc", readerTestIfndefWithUndefinedAttributeIncludesBlock, nil},

	{"ifndef with one alternative attribute set does not include content", "asciidoctor/reader_test_ifndef_with_one_alternative_attribute_set_does_not_include_content.adoc", readerTestIfndefWithOneAlternativeAttributeSetDoesNotIncludeContent, nil},

	{"ifndef with both alternative attributes set does not include content", "asciidoctor/reader_test_ifndef_with_both_alternative_attributes_set_does_not_include_content.adoc", readerTestIfndefWithBothAlternativeAttributesSetDoesNotIncludeContent, nil},

	{"ifndef with no alternative attributes set includes content", "asciidoctor/reader_test_ifndef_with_no_alternative_attributes_set_includes_content.adoc", readerTestIfndefWithNoAlternativeAttributesSetIncludesContent, nil},

	{"ifndef with no required attributes set includes content", "asciidoctor/reader_test_ifndef_with_no_required_attributes_set_includes_content.adoc", readerTestIfndefWithNoRequiredAttributesSetIncludesContent, nil},

	{"ifndef with all required attributes set does not include content", "asciidoctor/reader_test_ifndef_with_all_required_attributes_set_does_not_include_content.adoc", readerTestIfndefWithAllRequiredAttributesSetDoesNotIncludeContent, nil},

	{"ifndef with at least one required attributes set does not include content", "asciidoctor/reader_test_ifndef_with_at_least_one_required_attributes_set_does_not_include_content.adoc", readerTestIfndefWithAtLeastOneRequiredAttributesSetDoesNotIncludeContent, nil},

	{"ifdef around empty line does not introduce extra line", "asciidoctor/reader_test_ifdef_around_empty_line_does_not_introduce_extra_line.adoc", readerTestIfdefAroundEmptyLineDoesNotIntroduceExtraLine, nil},

	{"should log warning if endif is mismatched", "asciidoctor/reader_test_should_log_warning_if_endif_is_mismatched.adoc", readerTestShouldLogWarningIfEndifIsMismatched, nil},

	{"should log warning if endif contains text", "asciidoctor/reader_test_should_log_warning_if_endif_contains_text.adoc", readerTestShouldLogWarningIfEndifContainsText, nil},

	{"escaped ifdef is unescaped and ignored", "asciidoctor/reader_test_escaped_ifdef_is_unescaped_and_ignored.adoc", readerTestEscapedIfdefIsUnescapedAndIgnored, nil},

	{"ifeval comparing missing attribute to nil includes content", "asciidoctor/reader_test_ifeval_comparing_missing_attribute_to_nil_includes_content.adoc", readerTestIfevalComparingMissingAttributeToNilIncludesContent, nil},

	{"ifeval comparing missing attribute to 0 drops content", "asciidoctor/reader_test_ifeval_comparing_missing_attribute_to_0_drops_content.adoc", readerTestIfevalComparingMissingAttributeTo0DropsContent, nil},

	{"ifeval comparing double-quoted attribute to matching string includes content", "asciidoctor/reader_test_ifeval_comparing_double_quoted_attribute_to_matching_string_includes_content.adoc", readerTestIfevalComparingDoubleQuotedAttributeToMatchingStringIncludesContent, nil},

	{"ifeval comparing single-quoted attribute to matching string includes content", "asciidoctor/reader_test_ifeval_comparing_single_quoted_attribute_to_matching_string_includes_content.adoc", readerTestIfevalComparingSingleQuotedAttributeToMatchingStringIncludesContent, nil},

	{"ifeval comparing quoted attribute to non-matching string drops content", "asciidoctor/reader_test_ifeval_comparing_quoted_attribute_to_non_matching_string_drops_content.adoc", readerTestIfevalComparingQuotedAttributeToNonMatchingStringDropsContent, nil},

	{"ifeval comparing attribute to self includes content", "asciidoctor/reader_test_ifeval_comparing_attribute_to_self_includes_content.adoc", readerTestIfevalComparingAttributeToSelfIncludesContent, nil},

	{"ifeval matching numeric equality includes content", "asciidoctor/reader_test_ifeval_matching_numeric_equality_includes_content.adoc", readerTestIfevalMatchingNumericEqualityIncludesContent, nil},

	{"should warn if ifeval has target", "asciidoctor/reader_test_should_warn_if_ifeval_has_target.adoc", readerTestShouldWarnIfIfevalHasTarget, nil},

	{"should warn if ifeval has invalid expression", "asciidoctor/reader_test_should_warn_if_ifeval_has_invalid_expression.adoc", readerTestShouldWarnIfIfevalHasInvalidExpression, nil},

	{"should warn if ifeval is missing expression", "asciidoctor/reader_test_should_warn_if_ifeval_is_missing_expression.adoc", readerTestShouldWarnIfIfevalIsMissingExpression, nil},

	{"ifdef with no target is ignored", "asciidoctor/reader_test_ifdef_with_no_target_is_ignored.adoc", readerTestIfdefWithNoTargetIsIgnored, nil},

	{"should not warn about invalid ifdef preprocessor directive if already skipping", "asciidoctor/reader_test_should_not_warn_about_invalid_ifdef_preprocessor_directive_if_already_skipping.adoc", readerTestShouldNotWarnAboutInvalidIfdefPreprocessorDirectiveIfAlreadySkipping, nil},

	{"should not warn about invalid ifeval preprocessor directive if already skipping", "asciidoctor/reader_test_should_not_warn_about_invalid_ifeval_preprocessor_directive_if_already_skipping.adoc", readerTestShouldNotWarnAboutInvalidIfevalPreprocessorDirectiveIfAlreadySkipping, nil},
}

var readerTestShouldPrepareLinesFromArrayData = &asciidoc.Document{
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
					Value: "        This is one paragraph.",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 0,
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "        This is another paragraph.",
		},
		&asciidoc.NewLine{},
	},
}

var readerTestReadLinesUntilUntilBlankLine = &asciidoc.Document{
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
					Value: "        This is one paragraph.",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 0,
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "        This is another paragraph.",
		},
		&asciidoc.NewLine{},
	},
}

var readerTestReadLinesUntilUntilBlankLinePreservingLastLine = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.String{
			Value: ".split Asciidoctor::LF",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "        This is one paragraph.",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "        This is another paragraph.",
		},
		&asciidoc.NewLine{},
	},
}

var readerTestReadLinesUntilUntilConditionIsTrue = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.String{
			Value: ".split Asciidoctor::LF",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "        --",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "        This is one paragraph inside the block.",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "        This is another paragraph inside the block.",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "        --",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "        This is a paragraph outside the block.",
		},
		&asciidoc.NewLine{},
	},
}

var readerTestReadLinesUntilUntilConditionIsTrueTakingLastLine = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.String{
			Value: ".split Asciidoctor::LF",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "        --",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "        This is one paragraph inside the block.",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "        This is another paragraph inside the block.",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "        --",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "        This is a paragraph outside the block.",
		},
		&asciidoc.NewLine{},
	},
}

var readerTestReadLinesUntilUntilConditionIsTrueTakingAndPreservingLastLine = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.String{
			Value: ".split Asciidoctor::LF",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "        --",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "        This is one paragraph inside the block.",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "        This is another paragraph inside the block.",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "        --",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "        This is a paragraph outside the block.",
		},
		&asciidoc.NewLine{},
	},
}

var readerTestReadLinesUntilTerminator = &asciidoc.Document{
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
					Value: "        ****",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "        captured",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 0,
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "        also captured",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "        ****",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "        not captured",
		},
		&asciidoc.NewLine{},
	},
}

var readerTestShouldFlagReaderAsUnterminatedIfReaderReachesEndOfSourceWithoutFindingTerminator = &asciidoc.Document{
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
					Value: "        ****",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "        captured",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 0,
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "        also captured",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "        captured yet again",
		},
		&asciidoc.NewLine{},
	},
}

var readerTestShouldNotSkipFrontMatterByDefault = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.ThematicBreak{
			AttributeList: nil,
		},
		&asciidoc.String{
			Value: "layout: post",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "title: Document Title",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "author: username",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "tags: [ first, second ]",
		},
		&asciidoc.NewLine{},
		&asciidoc.ThematicBreak{
			AttributeList: nil,
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

var readerTestShouldNotSkipFrontMatterIfEndingDelimiterIsNotFound = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.ThematicBreak{
			AttributeList: nil,
		},
		&asciidoc.String{
			Value: "title: Document Title",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "tags: [ first, second ]",
		},
		&asciidoc.NewLine{},
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

var readerTestShouldSkipFrontMatterIfSpecifiedBySkipFrontMatterAttribute = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.Paragraph{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.TitleAttribute{
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "chop",
						},
					},
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "        layout: post",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "        title: Document Title",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "        author: username",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "        tags: [ first, second ]",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 0,
		},
	},
}

var readerTestShouldSkipTomlFrontMatterIfSpecifiedBySkipFrontMatterAttribute = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.Paragraph{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.TitleAttribute{
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "chop",
						},
					},
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "        layout = 'post'",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "        title = 'Document Title'",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "        author = 'username'",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "        tags = ['first', 'second']",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 0,
		},
	},
}

var readerTestShouldNotTrackIncludeInCatalogForNonAsciiDocIncludeFiles = &asciidoc.Document{
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
				"include::fixtures/circle.svg[]",
			},
		},
	},
}

var readerTestIncludeDirectiveShouldResolveFileWithSpacesInName = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "target",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "not-a-file.adoc + \\",
				},
			},
		},
		&asciidoc.Link{
			AttributeList: nil,
			URL: asciidoc.URL{
				Scheme: "http://",
				Path: asciidoc.Elements{
					&asciidoc.String{
						Value: "example.org/team.adoc",
					},
				},
			},
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.FileInclude{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.UserAttributeReference{
					Value: "target",
				},
			},
		},
	},
}

var readerTestIncludeDirectiveShouldResolveFileRelativeToCurrentInclude = &asciidoc.Document{
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
				"include::fixtures/data.tsv[]",
			},
		},
	},
}

var readerTestShouldFailToReadIncludeFileIfNotUtf8EncodedAndEncodingIsNotSpecified = &asciidoc.Document{
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
				"include::fixtures/iso-8859-1.txt[]",
			},
		},
	},
}

var readerTestShouldIgnoreEncodingAttributeIfValueIsNotAValidEncoding = &asciidoc.Document{
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
				"include::fixtures/encoding.adoc[tag=romé,encoding=iso-1000-1]",
			},
		},
	},
}

var readerTestShouldUseEncodingSpecifiedByEncodingAttributeWhenReadingIncludeFile = &asciidoc.Document{
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
				"include::fixtures/iso-8859-1.txt[encoding=iso-8859-1]",
			},
		},
	},
}

var readerTestUnresolvedTargetReferencedByIncludeDirectiveIsSkippedWhenOptionalOptionIsSet = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.FileInclude{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.NamedAttribute{
					Name: "opts",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "optional",
						},
					},
					Quote: 0,
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "fixtures/",
				},
				&asciidoc.UserAttributeReference{
					Value: "no-such-file",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "trailing content",
		},
		&asciidoc.NewLine{},
	},
}

var readerTestShouldSkipIncludeDirectiveThatReferencesMissingFileIfOptionalOptionIsSet = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.FileInclude{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.NamedAttribute{
					Name: "opts",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "optional",
						},
					},
					Quote: 0,
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "fixtures/no-such-file.adoc",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "trailing content",
		},
		&asciidoc.NewLine{},
	},
}

var readerTestShouldReplaceIncludeDirectiveThatReferencesMissingFileWithMessage = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.FileInclude{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "fixtures/no-such-file.adoc",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "trailing content",
		},
		&asciidoc.NewLine{},
	},
}

var readerTestNestedIncludeDirectivesAreResolvedRelativeToCurrentFile = &asciidoc.Document{
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
				"include::fixtures/outer-include.adoc[]",
			},
		},
	},
}

var readerTestIncludeDirectiveSupportsSelectingLinesByLineNumber = &asciidoc.Document{
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
				"include::fixtures/include-file.adoc[lines=]",
			},
		},
	},
}

var readerTestIncludeDirectiveIgnoresLinesAttributeWithInvalidRange = &asciidoc.Document{
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
				"include::fixtures/include-file.adoc[lines=10..5]",
			},
		},
	},
}

var readerTestIncludeDirectiveSupportsSelectingLinesByTagInFileThatHasCrlfLineEndings = &asciidoc.Document{
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
				"include::fixtures/include-file.adoc[tags=snippet]",
			},
		},
	},
}

var readerTestIncludeDirectiveSkipsLinesInsideTagWhichIsNegated = &asciidoc.Document{
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
				"include::fixtures/tagged-class-enclosed.rb[tags=all;!bark]",
			},
		},
	},
}

var readerTestIncludeDirectiveSelectsAllLinesWithoutATagDirectiveWhenValueIsDoubleAsterisk = &asciidoc.Document{
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
				"include::fixtures/tagged-class.rb[tags=**]",
			},
		},
	},
}

var readerTestIncludeDirectiveSelectsAllLinesExceptLinesInsideTagWhichIsNegatedWhenValueStartsWithDoubleAsterisk = &asciidoc.Document{
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
				"include::fixtures/tagged-class.rb[tags=**;!bark]",
			},
		},
	},
}

var readerTestIncludeDirectiveSelectsAllLinesIncludingLinesInsideNestedTagsExceptLinesInsideTagWhichIsNegatedWhenValueStartsWithDoubleAsterisk = &asciidoc.Document{
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
				"include::fixtures/tagged-class.rb[tags=**;!init]",
			},
		},
	},
}

var readerTestIncludeDirectiveSelectsAllLinesOutsideOfTagsWhenValueIsDoubleAsteriskFollowedByNegatedWildcard = &asciidoc.Document{
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
				"include::fixtures/tagged-class.rb[tags=**;!*]",
			},
		},
	},
}

var readerTestIncludeDirectiveSkipsAllTaggedRegionsWhenValueOfTagsAttributeIsNegatedWildcard = &asciidoc.Document{
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
				"include::fixtures/tagged-class.rb[tags=!*]",
			},
		},
	},
}

var readerTestIncludeDirectiveSelectsAllLinesExceptForLinesContainingTagDirectiveIfValueIsDoubleAsteriskFollowedByNestedTagNames = &asciidoc.Document{
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
				"include::fixtures/tagged-class.rb[tags=**;bark-beagle;bark-all]",
			},
		},
	},
}

var readerTestIncludeDirectiveSelectsAllLinesExceptForLinesContainingTagDirectiveWhenValueIsDoubleAsteriskFollowedByOuterTagName = &asciidoc.Document{
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
				"include::fixtures/tagged-class.rb[tags=**;bark]",
			},
		},
	},
}

var readerTestIncludeDirectiveSelectsAllLinesInsideUnspecifiedTagsWhenValueIsNegatedDoubleAsteriskFollowedByNegatedTags = &asciidoc.Document{
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
				"include::fixtures/tagged-class.rb[tags=!**;!init]",
			},
		},
	},
}

var readerTestIncludeDirectiveSelectsAllLinesExceptTagWhichIsNegatedWhenValueOnlyContainsNegatedTag = &asciidoc.Document{
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
				"include::fixtures/tagged-class.rb[tag=!bark]",
			},
		},
	},
}

var readerTestIncludeDirectiveSelectsAllLinesExceptTagsWhichAreNegatedWhenValueOnlyContainsNegatedTags = &asciidoc.Document{
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
				"include::fixtures/tagged-class.rb[tags=!bark;!init]",
			},
		},
	},
}

var readerTestShouldRecognizeTagWildcardIfNotAtStartOfTagsList = &asciidoc.Document{
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
				"include::fixtures/tagged-class.rb[tags=init;**;*;!bark-other]",
			},
		},
	},
}

var readerTestIncludeDirectiveSelectsLinesBetweenTagsWhenValueOfTagsAttributeIsWildcard = &asciidoc.Document{
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
				"include::fixtures/tagged-class.rb[tags=*]",
			},
		},
	},
}

var readerTestIncludeDirectiveSelectsLinesInsideTagsWhenValueOfTagsAttributeIsWildcardAndTagSurroundsContent = &asciidoc.Document{
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
				"include::fixtures/tagged-class-enclosed.rb[tags=*]",
			},
		},
	},
}

var readerTestIncludeDirectiveSelectsLinesInsideAllTagsExceptTagWhichIsNegatedWhenValueOfTagsAttributeIsWildcardFollowedByNegatedTag = &asciidoc.Document{
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
				"include::fixtures/tagged-class-enclosed.rb[tags=*;!init]",
			},
		},
	},
}

var readerTestIncludeDirectiveIncludesRegionsOutsideTagsAndInsideSpecifiedTagsWhenValueBeginsWithNegatedWildcard = &asciidoc.Document{
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
				"include::fixtures/tagged-class.rb[tags=!*;bark]",
			},
		},
	},
}

var readerTestIncludeDirectiveSelectsLinesInsideTagExceptForLinesInsideNestedTagsWhenTagIsPrecededByNegatedDoubleAsteriskAndNegatedWildcard = &asciidoc.Document{
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
				"include::fixtures/tagged-class.rb[tags=!**;!*;bark]",
			},
		},
	},
}

var readerTestIncludeDirectiveDoesNotSelectLinesInsideTagThatHasBeenIncludedThenExcluded = &asciidoc.Document{
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
				"include::fixtures/tagged-class.rb[tags=!*;init;!init]",
			},
		},
	},
}

var readerTestIncludeDirectiveSelectsLinesInsideSpecifiedTagAndIgnoresLinesInsideANegatedTag = &asciidoc.Document{
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
					Quote: 0,
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   5,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"include::fixtures/tagged-class.rb[tags=bark;!bark-other]",
			},
		},
	},
}

var readerTestShouldNotWarnIfSpecifiedNegatedTagIsNotFoundInIncludeFile = &asciidoc.Document{
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
				"include::fixtures/tagged-class-enclosed.rb[tag=!no-such-tag]",
			},
		},
	},
}

var readerTestShouldWarnIfSpecifiedTagsAreNotFoundInIncludeFile = &asciidoc.Document{
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
				"include::fixtures/include-file.adoc[tags=no-such-tag-b;no-such-tag-a]",
			},
		},
	},
}

var readerTestShouldNotWarnIfSpecifiedNegatedTagsAreNotFoundInIncludeFile = &asciidoc.Document{
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
				"include::fixtures/tagged-class-enclosed.rb[tags=all;!no-such-tag;!unknown-tag]",
			},
		},
	},
}

var readerTestShouldWarnIfSpecifiedTagInIncludeFileIsNotClosed = &asciidoc.Document{
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
				"include::fixtures/unclosed-tag.adoc[tag=a]",
			},
		},
	},
}

var readerTestShouldWarnIfEndTagInIncludedFileIsMismatched = &asciidoc.Document{
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
				"include::fixtures/mismatched-end-tag.adoc[tags=a;b]",
			},
		},
	},
}

var readerTestShouldWarnIfUnexpectedEndTagIsFoundInIncludedFile = &asciidoc.Document{
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
				"include::fixtures/unexpected-end-tag.adoc[tags=a]",
			},
		},
	},
}

var readerTestLinesAttributeTakesPrecedenceOverTagsAttributeInIncludeDirective = &asciidoc.Document{
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
				"include::fixtures/basic-docinfo.xml[lines=2..3, indent=0]",
			},
		},
	},
}

var readerTestShouldSubstituteAttributeReferencesInAttrlist = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "name-of-tag",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "snippetA",
				},
			},
		},
		&asciidoc.FileInclude{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.NamedAttribute{
					Name: "tag",
					Val: asciidoc.Elements{
						&asciidoc.UserAttributeReference{
							Value: "name-of-tag",
						},
					},
					Quote: 0,
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "fixtures/include-file.adoc",
				},
			},
		},
	},
}

var readerTestShouldFallBackToBuiltInIncludeDirectiveBehaviorWhenNotHandledByIncludeProcessor = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.String{
			Value: ".split Asciidoctor::LF",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "        = Main Document",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "        preamble",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "        :leveloffset: +1",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "        = Chapter A",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "        content",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "        :leveloffset!:",
		},
		&asciidoc.NewLine{},
	},
}

var readerTestAttributesAreSubstitutedInTargetOfIncludeDirective = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "fixturesdir",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "fixtures",
				},
			},
		},
		&asciidoc.AttributeEntry{
			Name: "ext",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "adoc",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.FileInclude{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.UserAttributeReference{
					Value: "fixturesdir",
				},
				&asciidoc.String{
					Value: "/include-file.",
				},
				&asciidoc.UserAttributeReference{
					Value: "ext",
				},
			},
		},
	},
}

var readerTestLineFollowingDroppedIncludeIsNotDropped = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.FileInclude{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.UserAttributeReference{
					Value: "foodir",
				},
				&asciidoc.String{
					Value: "/include-file.adoc",
				},
			},
		},
		&asciidoc.String{
			Value: "yo",
		},
		&asciidoc.NewLine{},
	},
}

var readerTestEscapedIncludeDirectiveIsLeftUnprocessed = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "\\include::fixtures/include-file.adoc[]",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "\\escape preserved here",
		},
		&asciidoc.NewLine{},
	},
}

var readerTestIncludeDirectiveNotAtStartOfLineIsIgnored = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "max-include-depth",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "1",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.FileInclude{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "include-file.adoc",
				},
			},
		},
	},
}

var readerTestIncludeDirectiveShouldBeDisabledIfMaxIncludeDepthHasBeenExceeded = &asciidoc.Document{
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
					Value: "        ////",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "        include::fixtures/no-such-file.adoc[]",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "        ////",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 0,
		},
	},
}

var readerTestSkipCommentLinesShouldNotProcessLinesRead = &asciidoc.Document{
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
					Value: "        ////",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "        include::fixtures/no-such-file.adoc[]",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "        ////",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 0,
		},
	},
}

var readerTestProcessLineReturnsNilIfCursorAdvanced = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.IfDef{
			Attributes: asciidoc.AttributeNames{
				"asciidoctor",
			},
			Union:  0,
			Inline: false,
		},
		&asciidoc.String{
			Value: "Asciidoctor!",
		},
		&asciidoc.NewLine{},
		&asciidoc.EndIf{
			Attributes: asciidoc.AttributeNames{
				"asciidoctor",
			},
			Union: 0,
		},
	},
}

var readerTestPeekLineAdvancesCursorToNextConditionalLineOfContent = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.IfDef{
			Attributes: asciidoc.AttributeNames{
				"asciidoctor",
			},
			Union:  0,
			Inline: false,
		},
		&asciidoc.String{
			Value: "Asciidoctor!",
		},
		&asciidoc.NewLine{},
		&asciidoc.EndIf{
			Attributes: asciidoc.AttributeNames{
				"asciidoctor",
			},
			Union: 0,
		},
	},
}

var readerTestPeekLinesShouldPreprocessLinesIfDirectIsFalse = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "The Asciidoctor",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "ifdef::asciidoctor[is in.]",
		},
		&asciidoc.NewLine{},
	},
}

var readerTestPeekLinesShouldNotPreprocessLinesIfDirectIsTrue = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "The Asciidoctor",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "ifdef::asciidoctor[is in.]",
		},
		&asciidoc.NewLine{},
	},
}

var readerTestPeekLinesShouldNotPreventSubsequentPreprocessingOfPeekedLines = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "The Asciidoctor",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "ifdef::asciidoctor[is in.]",
		},
		&asciidoc.NewLine{},
	},
}

var readerTestProcessLineReturnsLineIfCursorNotAdvanced = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "content",
		},
		&asciidoc.NewLine{},
		&asciidoc.IfDef{
			Attributes: asciidoc.AttributeNames{
				"asciidoctor",
			},
			Union:  0,
			Inline: false,
		},
		&asciidoc.String{
			Value: "Asciidoctor!",
		},
		&asciidoc.NewLine{},
		&asciidoc.EndIf{
			Attributes: asciidoc.AttributeNames{
				"asciidoctor",
			},
			Union: 0,
		},
	},
}

var readerTestPeekLineDoesNotAdvanceCursorWhenOnARegularContentLine = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "content",
		},
		&asciidoc.NewLine{},
		&asciidoc.IfDef{
			Attributes: asciidoc.AttributeNames{
				"asciidoctor",
			},
			Union:  0,
			Inline: false,
		},
		&asciidoc.String{
			Value: "Asciidoctor!",
		},
		&asciidoc.NewLine{},
		&asciidoc.EndIf{
			Attributes: asciidoc.AttributeNames{
				"asciidoctor",
			},
			Union: 0,
		},
	},
}

var readerTestPeekLineReturnsNilIfCursorAdvancesPastEndOfSource = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.IfDef{
			Attributes: asciidoc.AttributeNames{
				"foobar",
			},
			Union:  0,
			Inline: false,
		},
		&asciidoc.String{
			Value: "swallowed content",
		},
		&asciidoc.NewLine{},
		&asciidoc.EndIf{
			Attributes: asciidoc.AttributeNames{
				"foobar",
			},
			Union: 0,
		},
	},
}

var readerTestPeekLineReturnsNilIfContentsOfSkippedConditionalIsEmptyLine = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.IfDef{
			Attributes: asciidoc.AttributeNames{
				"foobar",
			},
			Union:  0,
			Inline: false,
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.EndIf{
			Attributes: asciidoc.AttributeNames{
				"foobar",
			},
			Union: 0,
		},
	},
}

var readerTestIfdefWithDefinedAttributeIncludesContent = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.IfDef{
			Attributes: asciidoc.AttributeNames{
				"holygrail",
			},
			Union:  0,
			Inline: false,
		},
		&asciidoc.String{
			Value: "There is a holy grail!",
		},
		&asciidoc.NewLine{},
		&asciidoc.EndIf{
			Attributes: asciidoc.AttributeNames{
				"holygrail",
			},
			Union: 0,
		},
	},
}

var readerTestIfdefWithDefinedAttributeIncludesTextInBrackets = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "On our quest we go...",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "ifdef::holygrail[There is a holy grail!]",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "There was much rejoicing.",
		},
		&asciidoc.NewLine{},
	},
}

var readerTestIfdefAttributeNameIsNotCaseSensitive = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.IfDef{
			Attributes: asciidoc.AttributeNames{
				"showScript",
			},
			Union:  0,
			Inline: false,
		},
		&asciidoc.String{
			Value: "The script is shown!",
		},
		&asciidoc.NewLine{},
		&asciidoc.EndIf{
			Attributes: asciidoc.AttributeNames{
				"showScript",
			},
			Union: 0,
		},
	},
}

var readerTestIfndefWithDefinedAttributeDoesNotIncludeTextInBrackets = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "On our quest we go...",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "ifndef::hardships[There is a holy grail!]",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "There was no rejoicing.",
		},
		&asciidoc.NewLine{},
	},
}

var readerTestIncludeWithNonMatchingNestedExclude = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.IfDef{
			Attributes: asciidoc.AttributeNames{
				"grail",
			},
			Union:  0,
			Inline: false,
		},
		&asciidoc.String{
			Value: "holy",
		},
		&asciidoc.NewLine{},
		&asciidoc.IfDef{
			Attributes: asciidoc.AttributeNames{
				"swallow",
			},
			Union:  0,
			Inline: false,
		},
		&asciidoc.String{
			Value: "swallow",
		},
		&asciidoc.NewLine{},
		&asciidoc.EndIf{
			Attributes: asciidoc.AttributeNames{
				"swallow",
			},
			Union: 0,
		},
		&asciidoc.String{
			Value: "grail",
		},
		&asciidoc.NewLine{},
		&asciidoc.EndIf{
			Attributes: asciidoc.AttributeNames{
				"grail",
			},
			Union: 0,
		},
	},
}

var readerTestNestedExcludesWithSameCondition = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.IfNDef{
			Attributes: asciidoc.AttributeNames{
				"grail",
			},
			Union:  0,
			Inline: false,
		},
		&asciidoc.IfNDef{
			Attributes: asciidoc.AttributeNames{
				"grail",
			},
			Union:  0,
			Inline: false,
		},
		&asciidoc.String{
			Value: "not here",
		},
		&asciidoc.NewLine{},
		&asciidoc.EndIf{
			Attributes: asciidoc.AttributeNames{
				"grail",
			},
			Union: 0,
		},
		&asciidoc.EndIf{
			Attributes: asciidoc.AttributeNames{
				"grail",
			},
			Union: 0,
		},
	},
}

var readerTestIncludeWithNestedExcludeOfInvertedCondition = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.IfDef{
			Attributes: asciidoc.AttributeNames{
				"grail",
			},
			Union:  0,
			Inline: false,
		},
		&asciidoc.String{
			Value: "holy",
		},
		&asciidoc.NewLine{},
		&asciidoc.IfNDef{
			Attributes: asciidoc.AttributeNames{
				"grail",
			},
			Union:  0,
			Inline: false,
		},
		&asciidoc.String{
			Value: "not here",
		},
		&asciidoc.NewLine{},
		&asciidoc.EndIf{
			Attributes: asciidoc.AttributeNames{
				"grail",
			},
			Union: 0,
		},
		&asciidoc.String{
			Value: "grail",
		},
		&asciidoc.NewLine{},
		&asciidoc.EndIf{
			Attributes: asciidoc.AttributeNames{
				"grail",
			},
			Union: 0,
		},
	},
}

var readerTestExcludeWithMatchingNestedExclude = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "poof",
		},
		&asciidoc.NewLine{},
		&asciidoc.IfDef{
			Attributes: asciidoc.AttributeNames{
				"swallow",
			},
			Union:  0,
			Inline: false,
		},
		&asciidoc.String{
			Value: "no",
		},
		&asciidoc.NewLine{},
		&asciidoc.IfDef{
			Attributes: asciidoc.AttributeNames{
				"swallow",
			},
			Union:  0,
			Inline: false,
		},
		&asciidoc.String{
			Value: "swallow",
		},
		&asciidoc.NewLine{},
		&asciidoc.EndIf{
			Attributes: asciidoc.AttributeNames{
				"swallow",
			},
			Union: 0,
		},
		&asciidoc.String{
			Value: "here",
		},
		&asciidoc.NewLine{},
		&asciidoc.EndIf{
			Attributes: asciidoc.AttributeNames{
				"swallow",
			},
			Union: 0,
		},
		&asciidoc.String{
			Value: "gone",
		},
		&asciidoc.NewLine{},
	},
}

var readerTestExcludeWithNestedIncludeUsingShorthandEnd = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "poof",
		},
		&asciidoc.NewLine{},
		&asciidoc.IfNDef{
			Attributes: asciidoc.AttributeNames{
				"grail",
			},
			Union:  0,
			Inline: false,
		},
		&asciidoc.String{
			Value: "no grail",
		},
		&asciidoc.NewLine{},
		&asciidoc.IfNDef{
			Attributes: asciidoc.AttributeNames{
				"swallow",
			},
			Union:  0,
			Inline: false,
		},
		&asciidoc.String{
			Value: "or swallow",
		},
		&asciidoc.NewLine{},
		&asciidoc.EndIf{
			Attributes: nil,
			Union:      0,
		},
		&asciidoc.String{
			Value: "in here",
		},
		&asciidoc.NewLine{},
		&asciidoc.EndIf{
			Attributes: nil,
			Union:      0,
		},
		&asciidoc.String{
			Value: "gone",
		},
		&asciidoc.NewLine{},
	},
}

var readerTestIfdefWithOneAlternativeAttributeSetIncludesContent = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.IfDef{
			Attributes: asciidoc.AttributeNames{
				"holygrail",
				"swallow",
			},
			Union:  0,
			Inline: false,
		},
		&asciidoc.String{
			Value: "Our quest is complete!",
		},
		&asciidoc.NewLine{},
		&asciidoc.EndIf{
			Attributes: asciidoc.AttributeNames{
				"holygrail",
				"swallow",
			},
			Union: 0,
		},
	},
}

var readerTestIfdefWithNoAlternativeAttributesSetDoesNotIncludeContent = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.IfDef{
			Attributes: asciidoc.AttributeNames{
				"holygrail",
				"swallow",
			},
			Union:  0,
			Inline: false,
		},
		&asciidoc.String{
			Value: "Our quest is complete!",
		},
		&asciidoc.NewLine{},
		&asciidoc.EndIf{
			Attributes: asciidoc.AttributeNames{
				"holygrail",
				"swallow",
			},
			Union: 0,
		},
	},
}

var readerTestIfdefWithAllRequiredAttributesSetIncludesContent = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "ifdef::holygrail+swallow[]",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "Our quest is complete!",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "endif::holygrail+swallow[]",
		},
		&asciidoc.NewLine{},
	},
}

var readerTestIfdefWithMissingRequiredAttributesDoesNotIncludeContent = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "ifdef::holygrail+swallow[]",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "Our quest is complete!",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "endif::holygrail+swallow[]",
		},
		&asciidoc.NewLine{},
	},
}

var readerTestIfndefWithUndefinedAttributeIncludesBlock = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.IfNDef{
			Attributes: asciidoc.AttributeNames{
				"holygrail",
			},
			Union:  0,
			Inline: false,
		},
		&asciidoc.String{
			Value: "Our quest continues to find the holy grail!",
		},
		&asciidoc.NewLine{},
		&asciidoc.EndIf{
			Attributes: asciidoc.AttributeNames{
				"holygrail",
			},
			Union: 0,
		},
	},
}

var readerTestIfndefWithOneAlternativeAttributeSetDoesNotIncludeContent = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.IfNDef{
			Attributes: asciidoc.AttributeNames{
				"holygrail",
				"swallow",
			},
			Union:  0,
			Inline: false,
		},
		&asciidoc.String{
			Value: "Our quest is complete!",
		},
		&asciidoc.NewLine{},
		&asciidoc.EndIf{
			Attributes: asciidoc.AttributeNames{
				"holygrail",
				"swallow",
			},
			Union: 0,
		},
	},
}

var readerTestIfndefWithBothAlternativeAttributesSetDoesNotIncludeContent = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.IfNDef{
			Attributes: asciidoc.AttributeNames{
				"holygrail",
				"swallow",
			},
			Union:  0,
			Inline: false,
		},
		&asciidoc.String{
			Value: "Our quest is complete!",
		},
		&asciidoc.NewLine{},
		&asciidoc.EndIf{
			Attributes: asciidoc.AttributeNames{
				"holygrail",
				"swallow",
			},
			Union: 0,
		},
	},
}

var readerTestIfndefWithNoAlternativeAttributesSetIncludesContent = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.IfNDef{
			Attributes: asciidoc.AttributeNames{
				"holygrail",
				"swallow",
			},
			Union:  0,
			Inline: false,
		},
		&asciidoc.String{
			Value: "Our quest is complete!",
		},
		&asciidoc.NewLine{},
		&asciidoc.EndIf{
			Attributes: asciidoc.AttributeNames{
				"holygrail",
				"swallow",
			},
			Union: 0,
		},
	},
}

var readerTestIfndefWithNoRequiredAttributesSetIncludesContent = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "ifndef::holygrail+swallow[]",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "Our quest is complete!",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "endif::holygrail+swallow[]",
		},
		&asciidoc.NewLine{},
	},
}

var readerTestIfndefWithAllRequiredAttributesSetDoesNotIncludeContent = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "ifndef::holygrail+swallow[]",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "Our quest is complete!",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "endif::holygrail+swallow[]",
		},
		&asciidoc.NewLine{},
	},
}

var readerTestIfndefWithAtLeastOneRequiredAttributesSetDoesNotIncludeContent = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "ifndef::holygrail+swallow[]",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "Our quest is complete!",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "endif::holygrail+swallow[]",
		},
		&asciidoc.NewLine{},
	},
}

var readerTestIfdefAroundEmptyLineDoesNotIntroduceExtraLine = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "before",
		},
		&asciidoc.NewLine{},
		&asciidoc.IfDef{
			Attributes: asciidoc.AttributeNames{
				"no-such-attribute",
			},
			Union:  0,
			Inline: false,
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.EndIf{
			Attributes: nil,
			Union:      0,
		},
		&asciidoc.String{
			Value: "after",
		},
		&asciidoc.NewLine{},
	},
}

var readerTestShouldLogWarningIfEndifIsMismatched = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.IfDef{
			Attributes: asciidoc.AttributeNames{
				"on-quest",
			},
			Union:  0,
			Inline: false,
		},
		&asciidoc.String{
			Value: "Our quest is complete!",
		},
		&asciidoc.NewLine{},
		&asciidoc.EndIf{
			Attributes: asciidoc.AttributeNames{
				"on-journey",
			},
			Union: 0,
		},
	},
}

var readerTestShouldLogWarningIfEndifContainsText = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.IfDef{
			Attributes: asciidoc.AttributeNames{
				"on-quest",
			},
			Union:  0,
			Inline: false,
		},
		&asciidoc.String{
			Value: "Our quest is complete!",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "endif::on-quest[complete!]",
		},
		&asciidoc.NewLine{},
	},
}

var readerTestEscapedIfdefIsUnescapedAndIgnored = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "\\ifdef::holygrail[]",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "content",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "\\endif::holygrail[]",
		},
		&asciidoc.NewLine{},
	},
}

var readerTestIfevalComparingMissingAttributeToNilIncludesContent = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.IfEval{
			Left: asciidoc.IfEvalValue{
				Quote: 1,
				Value: asciidoc.Elements{
					&asciidoc.UserAttributeReference{
						Value: "foo",
					},
				},
			},
			Operator: 1,
			Right: asciidoc.IfEvalValue{
				Quote: 1,
				Value: nil,
			},
			Inline: false,
		},
		&asciidoc.String{
			Value: "No foo for you!",
		},
		&asciidoc.NewLine{},
		&asciidoc.EndIf{
			Attributes: nil,
			Union:      0,
		},
	},
}

var readerTestIfevalComparingMissingAttributeTo0DropsContent = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.IfEval{
			Left: asciidoc.IfEvalValue{
				Quote: 0,
				Value: asciidoc.Elements{
					&asciidoc.UserAttributeReference{
						Value: "leveloffset",
					},
					&asciidoc.String{
						Value: " ",
					},
				},
			},
			Operator: 1,
			Right: asciidoc.IfEvalValue{
				Quote: 0,
				Value: asciidoc.Elements{
					&asciidoc.String{
						Value: "0",
					},
				},
			},
			Inline: false,
		},
		&asciidoc.String{
			Value: "I didn't make the cut!",
		},
		&asciidoc.NewLine{},
		&asciidoc.EndIf{
			Attributes: nil,
			Union:      0,
		},
	},
}

var readerTestIfevalComparingDoubleQuotedAttributeToMatchingStringIncludesContent = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.IfEval{
			Left: asciidoc.IfEvalValue{
				Quote: 2,
				Value: asciidoc.Elements{
					&asciidoc.UserAttributeReference{
						Value: "gem",
					},
				},
			},
			Operator: 1,
			Right: asciidoc.IfEvalValue{
				Quote: 2,
				Value: asciidoc.Elements{
					&asciidoc.String{
						Value: "asciidoctor",
					},
				},
			},
			Inline: false,
		},
		&asciidoc.String{
			Value: "Asciidoctor it is!",
		},
		&asciidoc.NewLine{},
		&asciidoc.EndIf{
			Attributes: nil,
			Union:      0,
		},
	},
}

var readerTestIfevalComparingSingleQuotedAttributeToMatchingStringIncludesContent = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.IfEval{
			Left: asciidoc.IfEvalValue{
				Quote: 1,
				Value: asciidoc.Elements{
					&asciidoc.UserAttributeReference{
						Value: "gem",
					},
				},
			},
			Operator: 1,
			Right: asciidoc.IfEvalValue{
				Quote: 1,
				Value: asciidoc.Elements{
					&asciidoc.String{
						Value: "asciidoctor",
					},
				},
			},
			Inline: false,
		},
		&asciidoc.String{
			Value: "Asciidoctor it is!",
		},
		&asciidoc.NewLine{},
		&asciidoc.EndIf{
			Attributes: nil,
			Union:      0,
		},
	},
}

var readerTestIfevalComparingQuotedAttributeToNonMatchingStringDropsContent = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.IfEval{
			Left: asciidoc.IfEvalValue{
				Quote: 1,
				Value: asciidoc.Elements{
					&asciidoc.UserAttributeReference{
						Value: "gem",
					},
				},
			},
			Operator: 1,
			Right: asciidoc.IfEvalValue{
				Quote: 1,
				Value: asciidoc.Elements{
					&asciidoc.String{
						Value: "asciidoctor",
					},
				},
			},
			Inline: false,
		},
		&asciidoc.String{
			Value: "Asciidoctor it is!",
		},
		&asciidoc.NewLine{},
		&asciidoc.EndIf{
			Attributes: nil,
			Union:      0,
		},
	},
}

var readerTestIfevalComparingAttributeToSelfIncludesContent = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.IfEval{
			Left: asciidoc.IfEvalValue{
				Quote: 1,
				Value: asciidoc.Elements{
					&asciidoc.UserAttributeReference{
						Value: "asciidoctor-version",
					},
				},
			},
			Operator: 1,
			Right: asciidoc.IfEvalValue{
				Quote: 1,
				Value: asciidoc.Elements{
					&asciidoc.UserAttributeReference{
						Value: "asciidoctor-version",
					},
				},
			},
			Inline: false,
		},
		&asciidoc.String{
			Value: "Of course it's the same!",
		},
		&asciidoc.NewLine{},
		&asciidoc.EndIf{
			Attributes: nil,
			Union:      0,
		},
	},
}

var readerTestIfevalMatchingNumericEqualityIncludesContent = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.IfEval{
			Left: asciidoc.IfEvalValue{
				Quote: 0,
				Value: asciidoc.Elements{
					&asciidoc.UserAttributeReference{
						Value: "rings",
					},
					&asciidoc.String{
						Value: " ",
					},
				},
			},
			Operator: 1,
			Right: asciidoc.IfEvalValue{
				Quote: 0,
				Value: asciidoc.Elements{
					&asciidoc.String{
						Value: "1",
					},
				},
			},
			Inline: false,
		},
		&asciidoc.String{
			Value: "One ring to rule them all!",
		},
		&asciidoc.NewLine{},
		&asciidoc.EndIf{
			Attributes: nil,
			Union:      0,
		},
	},
}

var readerTestShouldWarnIfIfevalHasTarget = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "ifeval::target[1 == 1]",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "content",
		},
		&asciidoc.NewLine{},
	},
}

var readerTestShouldWarnIfIfevalHasInvalidExpression = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "ifeval::[1 | 2]",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "content",
		},
		&asciidoc.NewLine{},
	},
}

var readerTestShouldWarnIfIfevalIsMissingExpression = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "ifeval::[]",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "content",
		},
		&asciidoc.NewLine{},
	},
}

var readerTestIfdefWithNoTargetIsIgnored = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "ifdef::[]",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "content",
		},
		&asciidoc.NewLine{},
	},
}

var readerTestShouldNotWarnAboutInvalidIfdefPreprocessorDirectiveIfAlreadySkipping = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.IfDef{
			Attributes: asciidoc.AttributeNames{
				"attribute-not-set",
			},
			Union:  0,
			Inline: false,
		},
		&asciidoc.String{
			Value: "foo",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "ifdef::[]",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "bar",
		},
		&asciidoc.NewLine{},
		&asciidoc.EndIf{
			Attributes: nil,
			Union:      0,
		},
		&asciidoc.String{
			Value: "baz",
		},
		&asciidoc.NewLine{},
	},
}

var readerTestShouldNotWarnAboutInvalidIfevalPreprocessorDirectiveIfAlreadySkipping = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.IfDef{
			Attributes: asciidoc.AttributeNames{
				"attribute-not-set",
			},
			Union:  0,
			Inline: false,
		},
		&asciidoc.String{
			Value: "foo",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "ifeval::[]",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "bar",
		},
		&asciidoc.NewLine{},
		&asciidoc.EndIf{
			Attributes: nil,
			Union:      0,
		},
		&asciidoc.String{
			Value: "baz",
		},
		&asciidoc.NewLine{},
	},
}
