package tests

import (
	"testing"

	"github.com/hasty/alchemy/asciidoc"
)

func TestTables(t *testing.T) {
	tablesTests.run(t)
}

var tablesTests = parseTests{
	
	{ "converts simple psv table", "asciidoctor/tables_test_converts_simple_psv_table.adoc", convertsSimplePsvTable },
	
	{ "should add direction CSS class if float attribute is set on table", "asciidoctor/tables_test_should_add_direction_css_class_if_float_attribute_is_set_on_table.adoc", shouldAddDirectionCssClassIfFloatAttributeIsSetOnTable },
	
	{ "should set stripes class if stripes option is set", "asciidoctor/tables_test_should_set_stripes_class_if_stripes_option_is_set.adoc", shouldSetStripesClassIfStripesOptionIsSet },
	
	{ "outputs a caption on simple psv table", "asciidoctor/tables_test_outputs_a_caption_on_simple_psv_table.adoc", outputsACaptionOnSimplePsvTable },
	
	{ "only increments table counter for tables that have a title", "asciidoctor/tables_test_only_increments_table_counter_for_tables_that_have_a_title.adoc", onlyIncrementsTableCounterForTablesThatHaveATitle },
	
	{ "uses explicit caption in front of title in place of default caption and number", "asciidoctor/tables_test_uses_explicit_caption_in_front_of_title_in_place_of_default_caption_and_number.adoc", usesExplicitCaptionInFrontOfTitleInPlaceOfDefaultCaptionAndNumber },
	
	{ "disables caption when caption attribute on table is empty", "asciidoctor/tables_test_disables_caption_when_caption_attribute_on_table_is_empty.adoc", disablesCaptionWhenCaptionAttributeOnTableIsEmpty },
	
	{ "disables caption when caption attribute on table is empty string", "asciidoctor/tables_test_disables_caption_when_caption_attribute_on_table_is_empty_string.adoc", disablesCaptionWhenCaptionAttributeOnTableIsEmptyString },
	
	{ "disables caption on table when table-caption document attribute is unset", "asciidoctor/tables_test_disables_caption_on_table_when_table_caption_document_attribute_is_unset.adoc", disablesCaptionOnTableWhenTableCaptionDocumentAttributeIsUnset },
	
	{ "ignores escaped separators", "asciidoctor/tables_test_ignores_escaped_separators.adoc", ignoresEscapedSeparators },
	
	{ "preserves escaped delimiters at the end of the line", "asciidoctor/tables_test_preserves_escaped_delimiters_at_the_end_of_the_line.adoc", preservesEscapedDelimitersAtTheEndOfTheLine },
	
	{ "should treat trailing pipe as an empty cell", "asciidoctor/tables_test_should_treat_trailing_pipe_as_an_empty_cell.adoc", shouldTreatTrailingPipeAsAnEmptyCell },
	
	{ "should auto recover with warning if missing leading separator on first cell", "asciidoctor/tables_test_should_auto_recover_with_warning_if_missing_leading_separator_on_first_cell.adoc", shouldAutoRecoverWithWarningIfMissingLeadingSeparatorOnFirstCell },
	
	{ "performs normal substitutions on cell content", "asciidoctor/tables_test_performs_normal_substitutions_on_cell_content.adoc", performsNormalSubstitutionsOnCellContent },
	
	{ "should only substitute specialchars for literal table cells", "asciidoctor/tables_test_should_only_substitute_specialchars_for_literal_table_cells.adoc", shouldOnlySubstituteSpecialcharsForLiteralTableCells },
	
	{ "should preserving leading spaces but not leading newlines or trailing spaces in literal table cells", "asciidoctor/tables_test_should_preserving_leading_spaces_but_not_leading_newlines_or_trailing_spaces_in_literal_table_cells.adoc", shouldPreservingLeadingSpacesButNotLeadingNewlinesOrTrailingSpacesInLiteralTableCells },
	
	{ "should ignore v table cell style", "asciidoctor/tables_test_should_ignore_v_table_cell_style.adoc", shouldIgnoreVTableCellStyle },
	
	{ "table and column width not assigned when autowidth option is specified", "asciidoctor/tables_test_table_and_column_width_not_assigned_when_autowidth_option_is_specified.adoc", tableAndColumnWidthNotAssignedWhenAutowidthOptionIsSpecified },
	
	{ "does not assign column width for autowidth columns in HTML output", "asciidoctor/tables_test_does_not_assign_column_width_for_autowidth_columns_in_html_output.adoc", doesNotAssignColumnWidthForAutowidthColumnsInHtmlOutput },
	
	{ "can assign autowidth to all columns even when table has a width", "asciidoctor/tables_test_can_assign_autowidth_to_all_columns_even_when_table_has_a_width.adoc", canAssignAutowidthToAllColumnsEvenWhenTableHasAWidth },
	
	{ "equally distributes remaining column width to autowidth columns in DocBook output", "asciidoctor/tables_test_equally_distributes_remaining_column_width_to_autowidth_columns_in_doc_book_output.adoc", equallyDistributesRemainingColumnWidthToAutowidthColumnsInDocBookOutput },
	
	{ "should compute column widths based on pagewidth when width is set on table in DocBook output", "asciidoctor/tables_test_should_compute_column_widths_based_on_pagewidth_when_width_is_set_on_table_in_doc_book_output.adoc", shouldComputeColumnWidthsBasedOnPagewidthWhenWidthIsSetOnTableInDocBookOutput },
	
	{ "explicit table width is used even when autowidth option is specified", "asciidoctor/tables_test_explicit_table_width_is_used_even_when_autowidth_option_is_specified.adoc", explicitTableWidthIsUsedEvenWhenAutowidthOptionIsSpecified },
	
	{ "first row sets number of columns when not specified", "asciidoctor/tables_test_first_row_sets_number_of_columns_when_not_specified.adoc", firstRowSetsNumberOfColumnsWhenNotSpecified },
	
	{ "colspec attribute using asterisk syntax sets number of columns", "asciidoctor/tables_test_colspec_attribute_using_asterisk_syntax_sets_number_of_columns.adoc", colspecAttributeUsingAsteriskSyntaxSetsNumberOfColumns },
	
	{ "table with explicit column count can have multiple rows on a single line", "asciidoctor/tables_test_table_with_explicit_column_count_can_have_multiple_rows_on_a_single_line.adoc", tableWithExplicitColumnCountCanHaveMultipleRowsOnASingleLine },
	
	{ "table with explicit deprecated colspec syntax can have multiple rows on a single line", "asciidoctor/tables_test_table_with_explicit_deprecated_colspec_syntax_can_have_multiple_rows_on_a_single_line.adoc", tableWithExplicitDeprecatedColspecSyntaxCanHaveMultipleRowsOnASingleLine },
	
	{ "columns are added for empty records in colspec attribute", "asciidoctor/tables_test_columns_are_added_for_empty_records_in_colspec_attribute.adoc", columnsAreAddedForEmptyRecordsInColspecAttribute },
	
	{ "cols may be separated by semi-colon instead of comma", "asciidoctor/tables_test_cols_may_be_separated_by_semi_colon_instead_of_comma.adoc", colsMayBeSeparatedBySemiColonInsteadOfComma },
	
	{ "cols attribute may include spaces", "asciidoctor/tables_test_cols_attribute_may_include_spaces.adoc", colsAttributeMayIncludeSpaces },
	
	{ "blank cols attribute should be ignored", "asciidoctor/tables_test_blank_cols_attribute_should_be_ignored.adoc", blankColsAttributeShouldBeIgnored },
	
	{ "empty cols attribute should be ignored", "asciidoctor/tables_test_empty_cols_attribute_should_be_ignored.adoc", emptyColsAttributeShouldBeIgnored },
	
	{ "table with header and footer", "asciidoctor/tables_test_table_with_header_and_footer.adoc", tableWithHeaderAndFooter },
	
	{ "table with header and footer docbook", "asciidoctor/tables_test_table_with_header_and_footer_docbook.adoc", tableWithHeaderAndFooterDocbook },
	
	{ "should set horizontal and vertical alignment when converting to DocBook", "asciidoctor/tables_test_should_set_horizontal_and_vertical_alignment_when_converting_to_doc_book.adoc", shouldSetHorizontalAndVerticalAlignmentWhenConvertingToDocBook },
	
	{ "should preserve frame value ends when converting to HTML", "asciidoctor/tables_test_should_preserve_frame_value_ends_when_converting_to_html.adoc", shouldPreserveFrameValueEndsWhenConvertingToHtml },
	
	{ "should normalize frame value topbot as ends when converting to HTML", "asciidoctor/tables_test_should_normalize_frame_value_topbot_as_ends_when_converting_to_html.adoc", shouldNormalizeFrameValueTopbotAsEndsWhenConvertingToHtml },
	
	{ "should preserve frame value topbot when converting to DocBook", "asciidoctor/tables_test_should_preserve_frame_value_topbot_when_converting_to_doc_book.adoc", shouldPreserveFrameValueTopbotWhenConvertingToDocBook },
	
	{ "should convert frame value ends to topbot when converting to DocBook", "asciidoctor/tables_test_should_convert_frame_value_ends_to_topbot_when_converting_to_doc_book.adoc", shouldConvertFrameValueEndsToTopbotWhenConvertingToDocBook },
	
	{ "table with implicit header row", "asciidoctor/tables_test_table_with_implicit_header_row.adoc", tableWithImplicitHeaderRow },
	
	{ "table with implicit header row only", "asciidoctor/tables_test_table_with_implicit_header_row_only.adoc", tableWithImplicitHeaderRowOnly },
	
	{ "table with implicit header row when other options set", "asciidoctor/tables_test_table_with_implicit_header_row_when_other_options_set.adoc", tableWithImplicitHeaderRowWhenOtherOptionsSet },
	
	{ "no implicit header row if second line not blank", "asciidoctor/tables_test_no_implicit_header_row_if_second_line_not_blank.adoc", noImplicitHeaderRowIfSecondLineNotBlank },
	
	{ "no implicit header row if cell in first line spans multiple lines", "asciidoctor/tables_test_no_implicit_header_row_if_cell_in_first_line_spans_multiple_lines.adoc", noImplicitHeaderRowIfCellInFirstLineSpansMultipleLines },
	
	{ "should format first cell as literal if there is no implicit header row and column has l style", "asciidoctor/tables_test_should_format_first_cell_as_literal_if_there_is_no_implicit_header_row_and_column_has_l_style.adoc", shouldFormatFirstCellAsLiteralIfThereIsNoImplicitHeaderRowAndColumnHasLStyle },
	
	{ "should format first cell as AsciiDoc if there is no implicit header row and column has a style", "asciidoctor/tables_test_should_format_first_cell_as_ascii_doc_if_there_is_no_implicit_header_row_and_column_has_a_style.adoc", shouldFormatFirstCellAsAsciiDocIfThereIsNoImplicitHeaderRowAndColumnHasAStyle },
	
	{ "should interpret leading indent if first cell is AsciiDoc and there is no implicit header row", "asciidoctor/tables_test_should_interpret_leading_indent_if_first_cell_is_ascii_doc_and_there_is_no_implicit_header_row.adoc", shouldInterpretLeadingIndentIfFirstCellIsAsciiDocAndThereIsNoImplicitHeaderRow },
	
	{ "should format first cell as AsciiDoc if there is no implicit header row and cell has a style", "asciidoctor/tables_test_should_format_first_cell_as_ascii_doc_if_there_is_no_implicit_header_row_and_cell_has_a_style.adoc", shouldFormatFirstCellAsAsciiDocIfThereIsNoImplicitHeaderRowAndCellHasAStyle },
	
	{ "no implicit header row if AsciiDoc cell in first line spans multiple lines", "asciidoctor/tables_test_no_implicit_header_row_if_ascii_doc_cell_in_first_line_spans_multiple_lines.adoc", noImplicitHeaderRowIfAsciiDocCellInFirstLineSpansMultipleLines },
	
	{ "no implicit header row if first line blank", "asciidoctor/tables_test_no_implicit_header_row_if_first_line_blank.adoc", noImplicitHeaderRowIfFirstLineBlank },
	
	{ "no implicit header row if noheader option is specified", "asciidoctor/tables_test_no_implicit_header_row_if_noheader_option_is_specified.adoc", noImplicitHeaderRowIfNoheaderOptionIsSpecified },
	
	{ "styles not applied to header cells", "asciidoctor/tables_test_styles_not_applied_to_header_cells.adoc", stylesNotAppliedToHeaderCells },
	
	{ "should apply text formatting to cells in implicit header row when column has a style", "asciidoctor/tables_test_should_apply_text_formatting_to_cells_in_implicit_header_row_when_column_has_a_style.adoc", shouldApplyTextFormattingToCellsInImplicitHeaderRowWhenColumnHasAStyle },
	
	{ "should apply style and text formatting to cells in first row if no implicit header", "asciidoctor/tables_test_should_apply_style_and_text_formatting_to_cells_in_first_row_if_no_implicit_header.adoc", shouldApplyStyleAndTextFormattingToCellsInFirstRowIfNoImplicitHeader },
	
	{ "vertical table headers use th element instead of header class", "asciidoctor/tables_test_vertical_table_headers_use_th_element_instead_of_header_class.adoc", verticalTableHeadersUseThElementInsteadOfHeaderClass },
	
	{ "supports horizontal and vertical source data with blank lines and table header", "asciidoctor/tables_test_supports_horizontal_and_vertical_source_data_with_blank_lines_and_table_header.adoc", supportsHorizontalAndVerticalSourceDataWithBlankLinesAndTableHeader },
	
	{ "percentages as column widths", "asciidoctor/tables_test_percentages_as_column_widths.adoc", percentagesAsColumnWidths },
	
	{ "spans, alignments and styles", "asciidoctor/tables_test_spans_alignments_and_styles.adoc", spansAlignmentsAndStyles },
	
	{ "sets up columns correctly if first row has cell that spans columns", "asciidoctor/tables_test_sets_up_columns_correctly_if_first_row_has_cell_that_spans_columns.adoc", setsUpColumnsCorrectlyIfFirstRowHasCellThatSpansColumns },
	
	{ "supports repeating cells", "asciidoctor/tables_test_supports_repeating_cells.adoc", supportsRepeatingCells },
	
	{ "calculates colnames correctly when using implicit column count and single cell with colspan", "asciidoctor/tables_test_calculates_colnames_correctly_when_using_implicit_column_count_and_single_cell_with_colspan.adoc", calculatesColnamesCorrectlyWhenUsingImplicitColumnCountAndSingleCellWithColspan },
	
	{ "calculates colnames correctly when using implicit column count and cells with mixed colspans", "asciidoctor/tables_test_calculates_colnames_correctly_when_using_implicit_column_count_and_cells_with_mixed_colspans.adoc", calculatesColnamesCorrectlyWhenUsingImplicitColumnCountAndCellsWithMixedColspans },
	
	{ "assigns unique column names for table with implicit column count and colspans in first row", "asciidoctor/tables_test_assigns_unique_column_names_for_table_with_implicit_column_count_and_colspans_in_first_row.adoc", assignsUniqueColumnNamesForTableWithImplicitColumnCountAndColspansInFirstRow },
	
	{ "ignores cell with colspan that exceeds colspec", "asciidoctor/tables_test_ignores_cell_with_colspan_that_exceeds_colspec.adoc", ignoresCellWithColspanThatExceedsColspec },
	
	{ "paragraph and literal repeated content", "asciidoctor/tables_test_paragraph_and_literal_repeated_content.adoc", paragraphAndLiteralRepeatedContent },
	
	{ "should not split paragraph at line containing only {blank} that is directly adjacent to non-blank lines", "asciidoctor/tables_test_should_not_split_paragraph_at_line_containing_only_{blank}_that_is_directly_adjacent_to_non_blank_lines.adoc", shouldNotSplitParagraphAtLineContainingOnlyblankThatIsDirectlyAdjacentToNonBlankLines },
	
	{ "should strip trailing newlines when splitting paragraphs", "asciidoctor/tables_test_should_strip_trailing_newlines_when_splitting_paragraphs.adoc", shouldStripTrailingNewlinesWhenSplittingParagraphs },
	
	{ "basic AsciiDoc cell", "asciidoctor/tables_test_basic_ascii_doc_cell.adoc", basicAsciiDocCell },
	
	{ "AsciiDoc table cell should be wrapped in div with class \"content\"", "asciidoctor/tables_test_ascii_doc_table_cell_should_be_wrapped_in_div_with_class__content.adoc", asciiDocTableCellShouldBeWrappedInDivWithClassContent },
	
	{ "doctype can be set in AsciiDoc table cell", "asciidoctor/tables_test_doctype_can_be_set_in_ascii_doc_table_cell.adoc", doctypeCanBeSetInAsciiDocTableCell },
	
	{ "should reset doctype to default in AsciiDoc table cell", "asciidoctor/tables_test_should_reset_doctype_to_default_in_ascii_doc_table_cell.adoc", shouldResetDoctypeToDefaultInAsciiDocTableCell },
	
	{ "should update doctype-related attributes in AsciiDoc table cell when doctype is set", "asciidoctor/tables_test_should_update_doctype_related_attributes_in_ascii_doc_table_cell_when_doctype_is_set.adoc", shouldUpdateDoctypeRelatedAttributesInAsciiDocTableCellWhenDoctypeIsSet },
	
	{ "should not allow AsciiDoc table cell to set a document attribute that was hard set by the API", "asciidoctor/tables_test_should_not_allow_ascii_doc_table_cell_to_set_a_document_attribute_that_was_hard_set_by_the_api.adoc", shouldNotAllowAsciiDocTableCellToSetADocumentAttributeThatWasHardSetByTheApi },
	
	{ "should not allow AsciiDoc table cell to set a document attribute that was hard unset by the API", "asciidoctor/tables_test_should_not_allow_ascii_doc_table_cell_to_set_a_document_attribute_that_was_hard_unset_by_the_api.adoc", shouldNotAllowAsciiDocTableCellToSetADocumentAttributeThatWasHardUnsetByTheApi },
	
	{ "should keep attribute unset in AsciiDoc table cell if unset in parent document", "asciidoctor/tables_test_should_keep_attribute_unset_in_ascii_doc_table_cell_if_unset_in_parent_document.adoc", shouldKeepAttributeUnsetInAsciiDocTableCellIfUnsetInParentDocument },
	
	{ "should allow attribute unset in parent document to be set in AsciiDoc table cell", "asciidoctor/tables_test_should_allow_attribute_unset_in_parent_document_to_be_set_in_ascii_doc_table_cell.adoc", shouldAllowAttributeUnsetInParentDocumentToBeSetInAsciiDocTableCell },
	
	{ "should not allow locked attribute unset in parent document to be set in AsciiDoc table cell", "asciidoctor/tables_test_should_not_allow_locked_attribute_unset_in_parent_document_to_be_set_in_ascii_doc_table_cell.adoc", shouldNotAllowLockedAttributeUnsetInParentDocumentToBeSetInAsciiDocTableCell },
	
	{ "AsciiDoc content", "asciidoctor/tables_test_ascii_doc_content.adoc", asciiDocContent },
	
	{ "should preserve leading indentation in contents of AsciiDoc table cell if contents starts with newline", "asciidoctor/tables_test_should_preserve_leading_indentation_in_contents_of_ascii_doc_table_cell_if_contents_starts_with_newline.adoc", shouldPreserveLeadingIndentationInContentsOfAsciiDocTableCellIfContentsStartsWithNewline },
	
	{ "preprocessor directive on first line of an AsciiDoc table cell should be processed", "asciidoctor/tables_test_preprocessor_directive_on_first_line_of_an_ascii_doc_table_cell_should_be_processed.adoc", preprocessorDirectiveOnFirstLineOfAnAsciiDocTableCellShouldBeProcessed },
	
	{ "cross reference link in an AsciiDoc table cell should resolve to reference in main document", "asciidoctor/tables_test_cross_reference_link_in_an_ascii_doc_table_cell_should_resolve_to_reference_in_main_document.adoc", crossReferenceLinkInAnAsciiDocTableCellShouldResolveToReferenceInMainDocument },
	
	{ "should discover anchor at start of cell and register it as a reference", "asciidoctor/tables_test_should_discover_anchor_at_start_of_cell_and_register_it_as_a_reference.adoc", shouldDiscoverAnchorAtStartOfCellAndRegisterItAsAReference },
	
	{ "should catalog anchor at start of cell in implicit header row when column has a style", "asciidoctor/tables_test_should_catalog_anchor_at_start_of_cell_in_implicit_header_row_when_column_has_a_style.adoc", shouldCatalogAnchorAtStartOfCellInImplicitHeaderRowWhenColumnHasAStyle },
	
	{ "should catalog anchor at start of cell in explicit header row when column has a style", "asciidoctor/tables_test_should_catalog_anchor_at_start_of_cell_in_explicit_header_row_when_column_has_a_style.adoc", shouldCatalogAnchorAtStartOfCellInExplicitHeaderRowWhenColumnHasAStyle },
	
	{ "should catalog anchor at start of cell in first row", "asciidoctor/tables_test_should_catalog_anchor_at_start_of_cell_in_first_row.adoc", shouldCatalogAnchorAtStartOfCellInFirstRow },
	
	{ "footnotes should not be shared between an AsciiDoc table cell and the main document", "asciidoctor/tables_test_footnotes_should_not_be_shared_between_an_ascii_doc_table_cell_and_the_main_document.adoc", footnotesShouldNotBeSharedBetweenAnAsciiDocTableCellAndTheMainDocument },
	
	{ "callout numbers should be globally unique, including AsciiDoc table cells", "asciidoctor/tables_test_callout_numbers_should_be_globally_unique_including_ascii_doc_table_cells.adoc", calloutNumbersShouldBeGloballyUniqueIncludingAsciiDocTableCells },
	
	{ "compat mode can be activated in AsciiDoc table cell", "asciidoctor/tables_test_compat_mode_can_be_activated_in_ascii_doc_table_cell.adoc", compatModeCanBeActivatedInAsciiDocTableCell },
	
	{ "compat mode in AsciiDoc table cell inherits from parent document", "asciidoctor/tables_test_compat_mode_in_ascii_doc_table_cell_inherits_from_parent_document.adoc", compatModeInAsciiDocTableCellInheritsFromParentDocument },
	
	{ "compat mode in AsciiDoc table cell can be unset if set in parent document", "asciidoctor/tables_test_compat_mode_in_ascii_doc_table_cell_can_be_unset_if_set_in_parent_document.adoc", compatModeInAsciiDocTableCellCanBeUnsetIfSetInParentDocument },
	
	{ "nested table", "asciidoctor/tables_test_nested_table.adoc", nestedTable },
	
	{ "can set format of nested table to psv", "asciidoctor/tables_test_can_set_format_of_nested_table_to_psv.adoc", canSetFormatOfNestedTableToPsv },
	
	{ "AsciiDoc table cell should inherit to_dir option from parent document", "asciidoctor/tables_test_ascii_doc_table_cell_should_inherit_to_dir_option_from_parent_document.adoc", asciiDocTableCellShouldInheritToDirOptionFromParentDocument },
	
	{ "AsciiDoc table cell should not inherit toc setting from parent document", "asciidoctor/tables_test_ascii_doc_table_cell_should_not_inherit_toc_setting_from_parent_document.adoc", asciiDocTableCellShouldNotInheritTocSettingFromParentDocument },
	
	{ "should be able to enable toc in an AsciiDoc table cell", "asciidoctor/tables_test_should_be_able_to_enable_toc_in_an_ascii_doc_table_cell.adoc", shouldBeAbleToEnableTocInAnAsciiDocTableCell },
	
	{ "should be able to enable toc in an AsciiDoc table cell even if hard unset by API", "asciidoctor/tables_test_should_be_able_to_enable_toc_in_an_ascii_doc_table_cell_even_if_hard_unset_by_api.adoc", shouldBeAbleToEnableTocInAnAsciiDocTableCellEvenIfHardUnsetByApi },
	
	{ "should be able to enable toc in both outer document and in an AsciiDoc table cell", "asciidoctor/tables_test_should_be_able_to_enable_toc_in_both_outer_document_and_in_an_ascii_doc_table_cell.adoc", shouldBeAbleToEnableTocInBothOuterDocumentAndInAnAsciiDocTableCell },
	
	{ "document in an AsciiDoc table cell should not see doctitle of parent", "asciidoctor/tables_test_document_in_an_ascii_doc_table_cell_should_not_see_doctitle_of_parent.adoc", documentInAnAsciiDocTableCellShouldNotSeeDoctitleOfParent },
	
	{ "cell background color", "asciidoctor/tables_test_cell_background_color.adoc", cellBackgroundColor },
	
	{ "should warn if table block is not terminated", "asciidoctor/tables_test_should_warn_if_table_block_is_not_terminated.adoc", shouldWarnIfTableBlockIsNotTerminated },
	
	{ "should show correct line number in warning about unterminated block inside AsciiDoc table cell", "asciidoctor/tables_test_should_show_correct_line_number_in_warning_about_unterminated_block_inside_ascii_doc_table_cell.adoc", shouldShowCorrectLineNumberInWarningAboutUnterminatedBlockInsideAsciiDocTableCell },
	
	{ "custom separator for an AsciiDoc table cell", "asciidoctor/tables_test_custom_separator_for_an_ascii_doc_table_cell.adoc", customSeparatorForAnAsciiDocTableCell },
	
	{ "table with breakable option docbook 5", "asciidoctor/tables_test_table_with_breakable_option_docbook_5.adoc", tableWithBreakableOptionDocbook5 },
	
	{ "table with unbreakable option docbook 5", "asciidoctor/tables_test_table_with_unbreakable_option_docbook_5.adoc", tableWithUnbreakableOptionDocbook5 },
	
	{ "no implicit header row if cell in first line is quoted and spans multiple lines", "asciidoctor/tables_test_no_implicit_header_row_if_cell_in_first_line_is_quoted_and_spans_multiple_lines.adoc", noImplicitHeaderRowIfCellInFirstLineIsQuotedAndSpansMultipleLines },
	
	{ "converts simple dsv table", "asciidoctor/tables_test_converts_simple_dsv_table.adoc", convertsSimpleDsvTable },
	
	{ "dsv format shorthand", "asciidoctor/tables_test_dsv_format_shorthand.adoc", dsvFormatShorthand },
	
	{ "single cell in DSV table should only produce single row", "asciidoctor/tables_test_single_cell_in_dsv_table_should_only_produce_single_row.adoc", singleCellInDsvTableShouldOnlyProduceSingleRow },
	
	{ "should treat trailing colon as an empty cell", "asciidoctor/tables_test_should_treat_trailing_colon_as_an_empty_cell.adoc", shouldTreatTrailingColonAsAnEmptyCell },
	
	{ "should treat trailing comma as an empty cell", "asciidoctor/tables_test_should_treat_trailing_comma_as_an_empty_cell.adoc", shouldTreatTrailingCommaAsAnEmptyCell },
	
	{ "should log error but not crash if cell data has unclosed quote", "asciidoctor/tables_test_should_log_error_but_not_crash_if_cell_data_has_unclosed_quote.adoc", shouldLogErrorButNotCrashIfCellDataHasUnclosedQuote },
	
	{ "should preserve newlines in quoted CSV values", "asciidoctor/tables_test_should_preserve_newlines_in_quoted_csv_values.adoc", shouldPreserveNewlinesInQuotedCsvValues },
	
	{ "should not drop trailing empty cell in TSV data when loaded from an include file", "asciidoctor/tables_test_should_not_drop_trailing_empty_cell_in_tsv_data_when_loaded_from_an_include_file.adoc", shouldNotDropTrailingEmptyCellInTsvDataWhenLoadedFromAnIncludeFile },
	
	{ "mixed unquoted records and quoted records with escaped quotes, commas, and wrapped lines", "asciidoctor/tables_test_mixed_unquoted_records_and_quoted_records_with_escaped_quotes_commas_and_wrapped_lines.adoc", mixedUnquotedRecordsAndQuotedRecordsWithEscapedQuotesCommasAndWrappedLines },
	
	{ "should allow quotes around a CSV value to be on their own lines", "asciidoctor/tables_test_should_allow_quotes_around_a_csv_value_to_be_on_their_own_lines.adoc", shouldAllowQuotesAroundACsvValueToBeOnTheirOwnLines },
	
	{ "csv format shorthand", "asciidoctor/tables_test_csv_format_shorthand.adoc", csvFormatShorthand },
	
	{ "custom csv separator", "asciidoctor/tables_test_custom_csv_separator.adoc", customCsvSeparator },
	
	{ "single cell in CSV table should only produce single row", "asciidoctor/tables_test_single_cell_in_csv_table_should_only_produce_single_row.adoc", singleCellInCsvTableShouldOnlyProduceSingleRow },
	
	{ "cell formatted with AsciiDoc style", "asciidoctor/tables_test_cell_formatted_with_ascii_doc_style.adoc", cellFormattedWithAsciiDocStyle },
	
	{ "should strip whitespace around contents of AsciiDoc cell", "asciidoctor/tables_test_should_strip_whitespace_around_contents_of_ascii_doc_cell.adoc", shouldStripWhitespaceAroundContentsOfAsciiDocCell },
	
}


var convertsSimplePsvTable = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: nil,
      ColumnCount: 3,
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
                  Value: "C",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "a",
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
                  Value: "b",
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
                  Value: "c",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "1",
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
                  Value: "2",
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
                  Value: "3",
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

var shouldAddDirectionCssClassIfFloatAttributeIsSetOnTable = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: asciidoc.AttributeList{
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
      ColumnCount: 3,
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
                  Value: "C",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "a",
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
                  Value: "b",
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
                  Value: "c",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "1",
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
                  Value: "2",
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
                  Value: "3",
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

var shouldSetStripesClassIfStripesOptionIsSet = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.NamedAttribute{
          Name: "stripes",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "odd",
            },
          },
          Quote: 0,
        },
      },
      ColumnCount: 3,
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
                  Value: "C",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "a",
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
                  Value: "b",
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
                  Value: "c",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "1",
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
                  Value: "2",
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
                  Value: "3",
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

var outputsACaptionOnSimplePsvTable = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TitleAttribute{
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Simple psv table",
            },
          },
        },
      },
      ColumnCount: 3,
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
                  Value: "C",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "a",
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
                  Value: "b",
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
                  Value: "c",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "1",
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
                  Value: "2",
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
                  Value: "3",
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

var onlyIncrementsTableCounterForTablesThatHaveATitle = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TitleAttribute{
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "First numbered table",
            },
          },
        },
      },
      ColumnCount: 3,
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
                  Value: "1",
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
                  Value: "2",
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
                  Value: "3",
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
    &asciidoc.Table{
      AttributeList: nil,
      ColumnCount: 3,
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
                  Value: "4",
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
                  Value: "5",
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
                  Value: "6",
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
    &asciidoc.Table{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TitleAttribute{
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Second numbered table",
            },
          },
        },
      },
      ColumnCount: 3,
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
                  Value: "7",
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
                  Value: "8",
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
                  Value: "9",
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

var usesExplicitCaptionInFrontOfTitleInPlaceOfDefaultCaptionAndNumber = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.NamedAttribute{
          Name: "caption",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "All the Data. ",
            },
          },
          Quote: 2,
        },
        &asciidoc.TitleAttribute{
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Simple psv table",
            },
          },
        },
      },
      ColumnCount: 3,
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
                  Value: "C",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "a",
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
                  Value: "b",
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
                  Value: "c",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "1",
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
                  Value: "2",
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
                  Value: "3",
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

var disablesCaptionWhenCaptionAttributeOnTableIsEmpty = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "[caption=]",
    },
    &asciidoc.NewLine{},
    &asciidoc.Table{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TitleAttribute{
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Simple psv table",
            },
          },
        },
      },
      ColumnCount: 3,
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
                  Value: "C",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "a",
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
                  Value: "b",
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
                  Value: "c",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "1",
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
                  Value: "2",
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
                  Value: "3",
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

var disablesCaptionWhenCaptionAttributeOnTableIsEmptyString = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.NamedAttribute{
          Name: "caption",
          Val: nil,
          Quote: 2,
        },
        &asciidoc.TitleAttribute{
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Simple psv table",
            },
          },
        },
      },
      ColumnCount: 3,
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
                  Value: "C",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "a",
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
                  Value: "b",
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
                  Value: "c",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "1",
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
                  Value: "2",
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
                  Value: "3",
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

var disablesCaptionOnTableWhenTableCaptionDocumentAttributeIsUnset = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeReset{
      Name: "table-caption",
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TitleAttribute{
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Simple psv table",
            },
          },
        },
      },
      ColumnCount: 3,
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
                  Value: "C",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "a",
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
                  Value: "b",
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
                  Value: "c",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "1",
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
                  Value: "2",
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
                  Value: "3",
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

var ignoresEscapedSeparators = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: nil,
      ColumnCount: 2,
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
                  Value: "A \\| here",
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
                  Value: "a \\| there",
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

var preservesEscapedDelimitersAtTheEndOfTheLine = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "%header",
            },
          },
        },
        &asciidoc.TableColumnsAttribute{
          Columns: []*asciidoc.TableColumn{
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 1,
                IsSet: false,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 0,
                IsSet: false,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 1,
                IsSet: true,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 0,
                IsSet: false,
              },
            },
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 1,
                IsSet: false,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 0,
                IsSet: false,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 1,
                IsSet: true,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 0,
                IsSet: false,
              },
            },
          },
        },
      },
      ColumnCount: 2,
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
                  Value: "B\\|",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "A1",
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
                  Value: "B1\\|",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "A2",
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
                  Value: "B2\\|",
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

var shouldTreatTrailingPipeAsAnEmptyCell = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: nil,
      ColumnCount: 2,
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
                  Value: "A1",
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
                  Value: "",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "B1",
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
                  Value: "B2",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "C1",
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
                  Value: "C2",
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

var shouldAutoRecoverWithWarningIfMissingLeadingSeparatorOnFirstCell = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "|===",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "A | here| a | there",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "| x",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "| y",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "| z",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "| end",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "|===",
    },
    &asciidoc.NewLine{},
  },
}

var performsNormalSubstitutionsOnCellContent = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "show_title",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "Cool new show",
        },
      },
    },
    &asciidoc.Table{
      AttributeList: nil,
      ColumnCount: 2,
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
                &asciidoc.UserAttributeReference{
                  Value: "show_title",
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
                  Value: "Coming soon...",
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

var shouldOnlySubstituteSpecialcharsForLiteralTableCells = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: nil,
      ColumnCount: 1,
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
                  Value: 4,
                  IsSet: true,
                },
              },
              Set: asciidoc.Set{
                &asciidoc.String{
                  Value: "one",
                },
                &asciidoc.NewLine{},
                &asciidoc.String{
                  Value: "*two*",
                },
                &asciidoc.NewLine{},
                &asciidoc.String{
                  Value: "three",
                },
                &asciidoc.NewLine{},
                &asciidoc.String{
                  Value: "<four>",
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

var shouldPreservingLeadingSpacesButNotLeadingNewlinesOrTrailingSpacesInLiteralTableCells = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TableColumnsAttribute{
          Columns: []*asciidoc.TableColumn{
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 2,
                IsSet: true,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 0,
                IsSet: false,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 1,
                IsSet: false,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 0,
                IsSet: false,
              },
            },
          },
        },
      },
      ColumnCount: 2,
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
                  Value: 4,
                  IsSet: true,
                },
              },
              Set: asciidoc.Set{
                &asciidoc.NewLine{},
                &asciidoc.String{
                  Value: "one",
                },
                &asciidoc.NewLine{},
                &asciidoc.String{
                  Value: "two",
                },
                &asciidoc.NewLine{},
                &asciidoc.String{
                  Value: "three",
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
                  Value: "normal",
                },
              },
              Blank: false,
            },
          },
        },
        asciidoc.EmptyLine{
          Text: "\n",
        },
      },
    },
  },
}

var shouldIgnoreVTableCellStyle = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Paragraph{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TableColumnsAttribute{
          Columns: []*asciidoc.TableColumn{
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 2,
                IsSet: true,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 0,
                IsSet: false,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 1,
                IsSet: false,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 0,
                IsSet: false,
              },
            },
          },
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "|===",
        },
        &asciidoc.NewLine{},
        &asciidoc.String{
          Value: "v|",
        },
        &asciidoc.NewLine{},
        &asciidoc.String{
          Value: "  one",
        },
        &asciidoc.NewLine{},
        &asciidoc.String{
          Value: "  two",
        },
        &asciidoc.NewLine{},
        &asciidoc.String{
          Value: "three",
        },
        &asciidoc.NewLine{},
      },
      Admonition: 0,
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "  | normal",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "|===",
    },
    &asciidoc.NewLine{},
  },
}

var tableAndColumnWidthNotAssignedWhenAutowidthOptionIsSpecified = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.NamedAttribute{
          Name: "options",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "autowidth",
            },
          },
          Quote: 2,
        },
      },
      ColumnCount: 3,
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
                  Value: "C",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "a",
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
                  Value: "b",
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
                  Value: "c",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "1",
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
                  Value: "2",
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
                  Value: "3",
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

var doesNotAssignColumnWidthForAutowidthColumnsInHtmlOutput = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TableColumnsAttribute{
          Columns: []*asciidoc.TableColumn{
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 1,
                IsSet: false,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 0,
                IsSet: false,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 1,
                IsSet: false,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 0,
                IsSet: false,
              },
            },
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 3,
                IsSet: true,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 0,
                IsSet: false,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: -1,
                IsSet: true,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 0,
                IsSet: false,
              },
            },
          },
        },
      },
      ColumnCount: 4,
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
                  Value: "C",
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
                  Value: "D",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "a",
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
                  Value: "b",
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
                  Value: "c",
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
                  Value: "d",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "1",
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
                  Value: "2",
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
                  Value: "3",
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
                  Value: "4",
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

var canAssignAutowidthToAllColumnsEvenWhenTableHasAWidth = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TableColumnsAttribute{
          Columns: []*asciidoc.TableColumn{
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 4,
                IsSet: true,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 0,
                IsSet: false,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: -1,
                IsSet: true,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 0,
                IsSet: false,
              },
            },
          },
        },
        &asciidoc.NamedAttribute{
          Name: "width",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "50%",
            },
          },
          Quote: 0,
        },
      },
      ColumnCount: 4,
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
                  Value: "C",
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
                  Value: "D",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "a",
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
                  Value: "b",
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
                  Value: "c",
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
                  Value: "d",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "1",
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
                  Value: "2",
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
                  Value: "3",
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
                  Value: "4",
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

var equallyDistributesRemainingColumnWidthToAutowidthColumnsInDocBookOutput = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TableColumnsAttribute{
          Columns: []*asciidoc.TableColumn{
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 1,
                IsSet: false,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 0,
                IsSet: false,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 1,
                IsSet: false,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 0,
                IsSet: false,
              },
            },
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 3,
                IsSet: true,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 0,
                IsSet: false,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: -1,
                IsSet: true,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 0,
                IsSet: false,
              },
            },
          },
        },
      },
      ColumnCount: 4,
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
                  Value: "C",
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
                  Value: "D",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "a",
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
                  Value: "b",
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
                  Value: "c",
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
                  Value: "d",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "1",
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
                  Value: "2",
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
                  Value: "3",
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
                  Value: "4",
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

var shouldComputeColumnWidthsBasedOnPagewidthWhenWidthIsSetOnTableInDocBookOutput = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "pagewidth",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "500",
        },
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.NamedAttribute{
          Name: "width",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "50%",
            },
          },
          Quote: 0,
        },
      },
      ColumnCount: 4,
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
                  Value: "C",
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
                  Value: "D",
                },
              },
              Blank: false,
            },
          },
        },
        asciidoc.EmptyLine{
          Text: "\n",
        },
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
                  Value: "a",
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
                  Value: "b",
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
                  Value: "c",
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
                  Value: "d",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "1",
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
                  Value: "2",
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
                  Value: "3",
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
                  Value: "4",
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

var explicitTableWidthIsUsedEvenWhenAutowidthOptionIsSpecified = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "%autowidth",
            },
          },
        },
        &asciidoc.NamedAttribute{
          Name: "width",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "75%",
            },
          },
          Quote: 0,
        },
      },
      ColumnCount: 3,
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
                  Value: "C",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "a",
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
                  Value: "b",
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
                  Value: "c",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "1",
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
                  Value: "2",
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
                  Value: "3",
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

var firstRowSetsNumberOfColumnsWhenNotSpecified = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: nil,
      ColumnCount: 4,
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
                  Value: "first",
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
                  Value: "second",
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
                  Value: "third",
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
                  Value: "fourth",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "1",
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
                  Value: "2",
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
                  Value: "3",
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
                  Value: "4",
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

var colspecAttributeUsingAsteriskSyntaxSetsNumberOfColumns = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TableColumnsAttribute{
          Columns: []*asciidoc.TableColumn{
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 3,
                IsSet: true,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 0,
                IsSet: false,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 1,
                IsSet: false,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 0,
                IsSet: false,
              },
            },
          },
        },
      },
      ColumnCount: 3,
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
                  Value: "C",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "a",
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
                  Value: "b",
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
                  Value: "c",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "1",
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
                  Value: "2",
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
                  Value: "3",
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

var tableWithExplicitColumnCountCanHaveMultipleRowsOnASingleLine = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TableColumnsAttribute{
          Columns: []*asciidoc.TableColumn{
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 3,
                IsSet: true,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 0,
                IsSet: false,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 1,
                IsSet: false,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 0,
                IsSet: false,
              },
            },
          },
        },
      },
      ColumnCount: 3,
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
                  Value: "one",
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
                  Value: "two",
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
                  Value: "1",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "2",
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
                  Value: "a",
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
                  Value: "b",
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

var tableWithExplicitDeprecatedColspecSyntaxCanHaveMultipleRowsOnASingleLine = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TableColumnsAttribute{
          Columns: []*asciidoc.TableColumn{
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 1,
                IsSet: false,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 0,
                IsSet: false,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 1,
                IsSet: false,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 0,
                IsSet: false,
              },
            },
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 1,
                IsSet: false,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 0,
                IsSet: false,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 1,
                IsSet: false,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 0,
                IsSet: false,
              },
            },
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 1,
                IsSet: false,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 0,
                IsSet: false,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 1,
                IsSet: false,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 0,
                IsSet: false,
              },
            },
          },
        },
      },
      ColumnCount: 3,
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
                  Value: "one",
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
                  Value: "two",
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
                  Value: "1",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "2",
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
                  Value: "a",
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
                  Value: "b",
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

var columnsAreAddedForEmptyRecordsInColspecAttribute = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TableColumnsAttribute{
          Columns: []*asciidoc.TableColumn{
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 1,
                IsSet: false,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 0,
                IsSet: true,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 1,
                IsSet: false,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 0,
                IsSet: false,
              },
            },
          },
        },
      },
      ColumnCount: 1,
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
                  Value: "one",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "two",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "1",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "2",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "a",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "b",
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

var colsMayBeSeparatedBySemiColonInsteadOfComma = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TableColumnsAttribute{
          Columns: []*asciidoc.TableColumn{
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 1,
                IsSet: false,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 0,
                IsSet: false,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 1,
                IsSet: true,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 6,
                IsSet: true,
              },
            },
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 1,
                IsSet: false,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 0,
                IsSet: false,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 3,
                IsSet: true,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 5,
                IsSet: true,
              },
            },
          },
        },
      },
      ColumnCount: 2,
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
                  Value: "strong",
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
                  Value: "mono",
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

var colsAttributeMayIncludeSpaces = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TableColumnsAttribute{
          Columns: []*asciidoc.TableColumn{
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 1,
                IsSet: false,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 0,
                IsSet: false,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 1,
                IsSet: true,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 0,
                IsSet: false,
              },
            },
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 1,
                IsSet: false,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 0,
                IsSet: false,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 1,
                IsSet: true,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 0,
                IsSet: false,
              },
            },
          },
        },
      },
      ColumnCount: 2,
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
                  Value: "one",
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
                  Value: "two",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "1",
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
                  Value: "2",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "a",
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
                  Value: "b",
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

var blankColsAttributeShouldBeIgnored = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TableColumnsAttribute{
          Columns: nil,
        },
      },
      ColumnCount: 2,
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
                  Value: "one",
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
                  Value: "two",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "1",
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
                  Value: "2",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "a",
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
                  Value: "b",
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

var emptyColsAttributeShouldBeIgnored = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TableColumnsAttribute{
          Columns: nil,
        },
      },
      ColumnCount: 2,
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
                  Value: "one",
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
                  Value: "two",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "1",
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
                  Value: "2",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "a",
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
                  Value: "b",
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

var tableWithHeaderAndFooter = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.NamedAttribute{
          Name: "options",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "header,footer",
            },
          },
          Quote: 2,
        },
      },
      ColumnCount: 2,
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
                  Value: "Item",
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
                  Value: "Quantity",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "Item 1",
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
                  Value: "1",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "Item 2",
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
                  Value: "2",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "Item 3",
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
                  Value: "3",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "Total",
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
                  Value: "6",
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

var tableWithHeaderAndFooterDocbook = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TitleAttribute{
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Table with header, body and footer",
            },
          },
        },
        &asciidoc.NamedAttribute{
          Name: "options",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "header,footer",
            },
          },
          Quote: 2,
        },
      },
      ColumnCount: 2,
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
                  Value: "Item",
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
                  Value: "Quantity",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "Item 1",
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
                  Value: "1",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "Item 2",
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
                  Value: "2",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "Item 3",
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
                  Value: "3",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "Total",
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
                  Value: "6",
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

var shouldSetHorizontalAndVerticalAlignmentWhenConvertingToDocBook = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: nil,
      ColumnCount: 3,
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
                  Value: 2,
                  IsSet: true,
                },
                VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                  Value: 2,
                  IsSet: true,
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
                  Value: 1,
                  IsSet: true,
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
                  Value: "C",
                },
              },
              Blank: false,
            },
          },
        },
        asciidoc.EmptyLine{
          Text: "\n",
        },
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
                  Value: "A1",
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
                  Value: 2,
                  IsSet: true,
                },
                VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                  Value: 2,
                  IsSet: true,
                },
                Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                  Value: 0,
                  IsSet: false,
                },
              },
              Set: asciidoc.Set{
                &asciidoc.String{
                  Value: "B1",
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
                  Value: 1,
                  IsSet: true,
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
                  Value: "C1",
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

var shouldPreserveFrameValueEndsWhenConvertingToHtml = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.NamedAttribute{
          Name: "frame",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "ends",
            },
          },
          Quote: 0,
        },
      },
      ColumnCount: 3,
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
                  Value: "C",
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

var shouldNormalizeFrameValueTopbotAsEndsWhenConvertingToHtml = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.NamedAttribute{
          Name: "frame",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "topbot",
            },
          },
          Quote: 0,
        },
      },
      ColumnCount: 3,
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
                  Value: "C",
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

var shouldPreserveFrameValueTopbotWhenConvertingToDocBook = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.NamedAttribute{
          Name: "frame",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "topbot",
            },
          },
          Quote: 0,
        },
      },
      ColumnCount: 3,
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
                  Value: "C",
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

var shouldConvertFrameValueEndsToTopbotWhenConvertingToDocBook = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.NamedAttribute{
          Name: "frame",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "ends",
            },
          },
          Quote: 0,
        },
      },
      ColumnCount: 3,
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
                  Value: "C",
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

var tableWithImplicitHeaderRow = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: nil,
      ColumnCount: 2,
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
                  Value: "Column 1",
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
                  Value: "Column 2",
                },
              },
              Blank: false,
            },
          },
        },
        asciidoc.EmptyLine{
          Text: "\n",
        },
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
                  Value: "Data A1",
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
                  Value: "Data B1",
                },
              },
              Blank: false,
            },
          },
        },
        asciidoc.EmptyLine{
          Text: "\n",
        },
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
                  Value: "Data A2",
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
                  Value: "Data B2",
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

var tableWithImplicitHeaderRowOnly = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: nil,
      ColumnCount: 2,
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
                  Value: "Column 1",
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
                  Value: "Column 2",
                },
              },
              Blank: false,
            },
          },
        },
        asciidoc.EmptyLine{
          Text: "\n",
        },
      },
    },
  },
}

var tableWithImplicitHeaderRowWhenOtherOptionsSet = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "%autowidth",
            },
          },
        },
      },
      ColumnCount: 2,
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
                  Value: "Column 1",
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
                  Value: "Column 2",
                },
              },
              Blank: false,
            },
          },
        },
        asciidoc.EmptyLine{
          Text: "\n",
        },
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
                  Value: "Data A1",
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
                  Value: "Data B1",
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

var noImplicitHeaderRowIfSecondLineNotBlank = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: nil,
      ColumnCount: 2,
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
                  Value: "Column 1",
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
                  Value: "Column 2",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "Data A1",
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
                  Value: "Data B1",
                },
              },
              Blank: false,
            },
          },
        },
        asciidoc.EmptyLine{
          Text: "\n",
        },
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
                  Value: "Data A2",
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
                  Value: "Data B2",
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

var noImplicitHeaderRowIfCellInFirstLineSpansMultipleLines = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TableColumnsAttribute{
          Columns: []*asciidoc.TableColumn{
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 2,
                IsSet: true,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 0,
                IsSet: false,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 1,
                IsSet: false,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 0,
                IsSet: false,
              },
            },
          },
        },
      },
      ColumnCount: 2,
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
                  Value: "A1",
                },
                &asciidoc.NewLine{},
                &asciidoc.NewLine{},
                &asciidoc.NewLine{},
                &asciidoc.String{
                  Value: "A1 continued",
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
                  Value: "B1",
                },
              },
              Blank: false,
            },
          },
        },
        asciidoc.EmptyLine{
          Text: "\n",
        },
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
                  Value: "A2",
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
                  Value: "B2",
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

var shouldFormatFirstCellAsLiteralIfThereIsNoImplicitHeaderRowAndColumnHasLStyle = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TableColumnsAttribute{
          Columns: []*asciidoc.TableColumn{
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 1,
                IsSet: false,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 0,
                IsSet: false,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 1,
                IsSet: true,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 4,
                IsSet: true,
              },
            },
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 1,
                IsSet: false,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 0,
                IsSet: false,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 1,
                IsSet: true,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 0,
                IsSet: false,
              },
            },
          },
        },
      },
      ColumnCount: 2,
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
                  Value: "literal",
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
                  Value: "normal",
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

var shouldFormatFirstCellAsAsciiDocIfThereIsNoImplicitHeaderRowAndColumnHasAStyle = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TableColumnsAttribute{
          Columns: []*asciidoc.TableColumn{
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 1,
                IsSet: false,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 0,
                IsSet: false,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 1,
                IsSet: true,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 1,
                IsSet: true,
              },
            },
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 1,
                IsSet: false,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 0,
                IsSet: false,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 1,
                IsSet: true,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 0,
                IsSet: false,
              },
            },
          },
        },
      },
      ColumnCount: 2,
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
                &asciidoc.UnorderedListItem{
                  Set: asciidoc.Set{
                    &asciidoc.String{
                      Value: "list",
                    },
                  },
                  AttributeList: nil,
                  Indent: "",
                  Marker: "*",
                  Checklist: 0,
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
                  Value: "normal",
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

var shouldInterpretLeadingIndentIfFirstCellIsAsciiDocAndThereIsNoImplicitHeaderRow = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Paragraph{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TableColumnsAttribute{
          Columns: []*asciidoc.TableColumn{
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 1,
                IsSet: false,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 0,
                IsSet: false,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 1,
                IsSet: true,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 1,
                IsSet: true,
              },
            },
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 1,
                IsSet: false,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 0,
                IsSet: false,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 1,
                IsSet: true,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 0,
                IsSet: false,
              },
            },
          },
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "|===",
        },
        &asciidoc.NewLine{},
        &asciidoc.String{
          Value: "|",
        },
        &asciidoc.NewLine{},
        &asciidoc.String{
          Value: "  literal",
        },
        &asciidoc.NewLine{},
        &asciidoc.String{
          Value: "| normal",
        },
        &asciidoc.NewLine{},
        &asciidoc.String{
          Value: "|===",
        },
        &asciidoc.NewLine{},
      },
      Admonition: 0,
    },
  },
}

var shouldFormatFirstCellAsAsciiDocIfThereIsNoImplicitHeaderRowAndCellHasAStyle = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: nil,
      ColumnCount: 1,
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
                &asciidoc.UnorderedListItem{
                  Set: asciidoc.Set{
                    &asciidoc.String{
                      Value: "list",
                    },
                  },
                  AttributeList: nil,
                  Indent: "",
                  Marker: "*",
                  Checklist: 0,
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "normal",
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

var noImplicitHeaderRowIfAsciiDocCellInFirstLineSpansMultipleLines = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TableColumnsAttribute{
          Columns: []*asciidoc.TableColumn{
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 2,
                IsSet: true,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 0,
                IsSet: false,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 1,
                IsSet: false,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 0,
                IsSet: false,
              },
            },
          },
        },
      },
      ColumnCount: 2,
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
                &asciidoc.String{
                  Value: "contains AsciiDoc content",
                },
                &asciidoc.NewLine{},
                asciidoc.EmptyLine{
                  Text: "",
                },
                &asciidoc.UnorderedListItem{
                  Set: asciidoc.Set{
                    &asciidoc.String{
                      Value: "a",
                    },
                  },
                  AttributeList: nil,
                  Indent: "",
                  Marker: "*",
                  Checklist: 0,
                },
                &asciidoc.UnorderedListItem{
                  Set: asciidoc.Set{
                    &asciidoc.String{
                      Value: "b",
                    },
                  },
                  AttributeList: nil,
                  Indent: "",
                  Marker: "*",
                  Checklist: 0,
                },
                &asciidoc.UnorderedListItem{
                  Set: asciidoc.Set{
                    &asciidoc.String{
                      Value: "c",
                    },
                  },
                  AttributeList: nil,
                  Indent: "",
                  Marker: "*",
                  Checklist: 0,
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
                  Value: 1,
                  IsSet: true,
                },
              },
              Set: asciidoc.Set{
                &asciidoc.String{
                  Value: "contains no AsciiDoc content",
                },
                &asciidoc.NewLine{},
                asciidoc.EmptyLine{
                  Text: "",
                },
                &asciidoc.String{
                  Value: "just text",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "A2",
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
                  Value: "B2",
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

var noImplicitHeaderRowIfFirstLineBlank = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: nil,
      ColumnCount: 2,
      Set: asciidoc.Set{
        asciidoc.EmptyLine{
          Text: "\n",
        },
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
                  Value: "Column 1",
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
                  Value: "Column 2",
                },
              },
              Blank: false,
            },
          },
        },
        asciidoc.EmptyLine{
          Text: "\n",
        },
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
                  Value: "Data A1",
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
                  Value: "Data B1",
                },
              },
              Blank: false,
            },
          },
        },
        asciidoc.EmptyLine{
          Text: "\n",
        },
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
                  Value: "Data A2",
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
                  Value: "Data B2",
                },
              },
              Blank: false,
            },
          },
        },
        asciidoc.EmptyLine{
          Text: "\n",
        },
      },
    },
  },
}

var noImplicitHeaderRowIfNoheaderOptionIsSpecified = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "%noheader",
            },
          },
        },
      },
      ColumnCount: 2,
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
                  Value: "Column 1",
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
                  Value: "Column 2",
                },
              },
              Blank: false,
            },
          },
        },
        asciidoc.EmptyLine{
          Text: "\n",
        },
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
                  Value: "Data A1",
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
                  Value: "Data B1",
                },
              },
              Blank: false,
            },
          },
        },
        asciidoc.EmptyLine{
          Text: "\n",
        },
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
                  Value: "Data A2",
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
                  Value: "Data B2",
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

var stylesNotAppliedToHeaderCells = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TableColumnsAttribute{
          Columns: []*asciidoc.TableColumn{
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 1,
                IsSet: false,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 0,
                IsSet: false,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 1,
                IsSet: true,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 3,
                IsSet: true,
              },
            },
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 1,
                IsSet: false,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 0,
                IsSet: false,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 1,
                IsSet: true,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 6,
                IsSet: true,
              },
            },
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 1,
                IsSet: false,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 0,
                IsSet: false,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 1,
                IsSet: true,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 2,
                IsSet: true,
              },
            },
          },
        },
        &asciidoc.NamedAttribute{
          Name: "options",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "header,footer",
            },
          },
          Quote: 2,
        },
      },
      ColumnCount: 3,
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
                  Value: "Name",
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
                  Value: "Occupation",
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
                  Value: "Website",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "Octocat",
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
                  Value: "Social coding",
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
                &asciidoc.Link{
                  AttributeList: nil,
                  URL: asciidoc.URL{
                    Scheme: "https://",
                    Path: asciidoc.Set{
                      &asciidoc.String{
                        Value: "github.com",
                      },
                    },
                  },
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "Name",
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
                  Value: "Occupation",
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
                  Value: "Website",
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

var shouldApplyTextFormattingToCellsInImplicitHeaderRowWhenColumnHasAStyle = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TableColumnsAttribute{
          Columns: []*asciidoc.TableColumn{
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 2,
                IsSet: true,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 0,
                IsSet: false,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 1,
                IsSet: false,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 1,
                IsSet: true,
              },
            },
          },
        },
      },
      ColumnCount: 2,
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
                &asciidoc.Italic{
                  AttributeList: nil,
                  Set: asciidoc.Set{
                    &asciidoc.String{
                      Value: "foo",
                    },
                  },
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
                &asciidoc.Bold{
                  AttributeList: nil,
                  Set: asciidoc.Set{
                    &asciidoc.String{
                      Value: "bar",
                    },
                  },
                },
              },
              Blank: false,
            },
          },
        },
        asciidoc.EmptyLine{
          Text: "\n",
        },
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
                &asciidoc.UnorderedListItem{
                  Set: asciidoc.Set{
                    &asciidoc.String{
                      Value: "list item",
                    },
                  },
                  AttributeList: nil,
                  Indent: "",
                  Marker: "*",
                  Checklist: 0,
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
                  Value: "paragraph",
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

var shouldApplyStyleAndTextFormattingToCellsInFirstRowIfNoImplicitHeader = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TableColumnsAttribute{
          Columns: []*asciidoc.TableColumn{
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 1,
                IsSet: false,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 0,
                IsSet: false,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 1,
                IsSet: false,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 6,
                IsSet: true,
              },
            },
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 1,
                IsSet: false,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 0,
                IsSet: false,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 1,
                IsSet: false,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 2,
                IsSet: true,
              },
            },
          },
        },
      },
      ColumnCount: 2,
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
                &asciidoc.Italic{
                  AttributeList: nil,
                  Set: asciidoc.Set{
                    &asciidoc.String{
                      Value: "strong",
                    },
                  },
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
                &asciidoc.Bold{
                  AttributeList: nil,
                  Set: asciidoc.Set{
                    &asciidoc.String{
                      Value: "emphasis",
                    },
                  },
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "strong",
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
                  Value: "emphasis",
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

var verticalTableHeadersUseThElementInsteadOfHeaderClass = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TableColumnsAttribute{
          Columns: []*asciidoc.TableColumn{
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 1,
                IsSet: false,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 0,
                IsSet: false,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 1,
                IsSet: true,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 3,
                IsSet: true,
              },
            },
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 1,
                IsSet: false,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 0,
                IsSet: false,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 1,
                IsSet: true,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 6,
                IsSet: true,
              },
            },
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 1,
                IsSet: false,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 0,
                IsSet: false,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 1,
                IsSet: true,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 2,
                IsSet: true,
              },
            },
          },
        },
      },
      ColumnCount: 3,
      Set: asciidoc.Set{
        asciidoc.EmptyLine{
          Text: "\n",
        },
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
                  Value: "Name",
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
                  Value: "Occupation",
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
                  Value: "Website",
                },
              },
              Blank: false,
            },
          },
        },
        asciidoc.EmptyLine{
          Text: "\n",
        },
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
                  Value: "Octocat",
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
                  Value: "Social coding",
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
                &asciidoc.Link{
                  AttributeList: nil,
                  URL: asciidoc.URL{
                    Scheme: "https://",
                    Path: asciidoc.Set{
                      &asciidoc.String{
                        Value: "github.com",
                      },
                    },
                  },
                },
              },
              Blank: false,
            },
          },
        },
        asciidoc.EmptyLine{
          Text: "\n",
        },
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
                  Value: "Name",
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
                  Value: "Occupation",
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
                  Value: "Website",
                },
              },
              Blank: false,
            },
          },
        },
        asciidoc.EmptyLine{
          Text: "\n",
        },
      },
    },
  },
}

var supportsHorizontalAndVerticalSourceDataWithBlankLinesAndTableHeader = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Paragraph{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TitleAttribute{
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Horizontal and vertical source data",
            },
          },
        },
        &asciidoc.NamedAttribute{
          Name: "width",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "80%",
            },
          },
          Quote: 2,
        },
        &asciidoc.TableColumnsAttribute{
          Columns: []*asciidoc.TableColumn{
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 1,
                IsSet: false,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 0,
                IsSet: false,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 3,
                IsSet: true,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 0,
                IsSet: false,
              },
            },
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 1,
                IsSet: false,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 2,
                IsSet: true,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 2,
                IsSet: true,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 0,
                IsSet: false,
              },
            },
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 1,
                IsSet: false,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 2,
                IsSet: true,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 2,
                IsSet: true,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 0,
                IsSet: false,
              },
            },
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 1,
                IsSet: false,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 0,
                IsSet: false,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 10,
                IsSet: true,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 0,
                IsSet: false,
              },
            },
          },
        },
        &asciidoc.NamedAttribute{
          Name: "options",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "header",
            },
          },
          Quote: 2,
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "|===",
        },
        &asciidoc.NewLine{},
        &asciidoc.String{
          Value: "|Date |Duration |Avg HR |Notes",
        },
        &asciidoc.NewLine{},
      },
      Admonition: 0,
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "|22-Aug-08 |10:24 | 157 |",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "Worked out MSHR (max sustainable heart rate) by going hard",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "for this interval.",
    },
    &asciidoc.NewLine{},
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "|22-Aug-08 |23:03 | 152 |",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "Back-to-back with previous interval.",
    },
    &asciidoc.NewLine{},
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "|24-Aug-08 |40:00 | 145 |",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "Moderately hard interspersed with 3x 3min intervals (2 min",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "hard + 1 min really hard taking the HR up to 160).",
    },
    &asciidoc.NewLine{},
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "I am getting in shape!",
    },
    &asciidoc.NewLine{},
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "|===",
    },
    &asciidoc.NewLine{},
  },
}

var percentagesAsColumnWidths = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TableColumnsAttribute{
          Columns: []*asciidoc.TableColumn{
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 1,
                IsSet: false,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 0,
                IsSet: true,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 2,
                IsSet: true,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 1,
                IsSet: false,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 0,
                IsSet: false,
              },
            },
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 1,
                IsSet: false,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 0,
                IsSet: true,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 1,
                IsSet: false,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 0,
                IsSet: false,
              },
            },
          },
        },
      },
      ColumnCount: 2,
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
                  Value: "column A",
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
                  Value: "column B",
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

var spansAlignmentsAndStyles = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TableColumnsAttribute{
          Columns: []*asciidoc.TableColumn{
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 1,
                IsSet: false,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 0,
                IsSet: false,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 1,
                IsSet: false,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 2,
                IsSet: true,
              },
            },
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 1,
                IsSet: false,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 0,
                IsSet: false,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 1,
                IsSet: false,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 5,
                IsSet: true,
              },
            },
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 1,
                IsSet: false,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 2,
                IsSet: true,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 1,
                IsSet: false,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 0,
                IsSet: false,
              },
            },
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 1,
                IsSet: false,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 1,
                IsSet: true,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 1,
                IsSet: false,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 6,
                IsSet: true,
              },
            },
          },
        },
        &asciidoc.NamedAttribute{
          Name: "width",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "25%",
            },
          },
          Quote: 2,
        },
      },
      ColumnCount: 4,
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
                  Value: "1",
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
                  Value: 1,
                  IsSet: true,
                },
                VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                  Value: 0,
                  IsSet: false,
                },
                Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                  Value: 6,
                  IsSet: true,
                },
              },
              Set: asciidoc.Set{
                &asciidoc.String{
                  Value: "2",
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
                  Value: "3",
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
                  Value: "4",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: 2,
                  IsSet: true,
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
                  Value: "5",
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
                    Value: 2,
                    IsSet: true,
                  },
                  Row: asciidoc.Optional[int]{
                    Value: 2,
                    IsSet: true,
                  },
                },
                HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                  Value: 2,
                  IsSet: true,
                },
                VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                  Value: 2,
                  IsSet: true,
                },
                Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                  Value: 0,
                  IsSet: false,
                },
              },
              Set: asciidoc.Set{
                &asciidoc.String{
                  Value: "6",
                },
              },
              Blank: false,
            },
            &asciidoc.TableCell{
              Format: nil,
              Set: nil,
              Blank: true,
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
                    Value: 3,
                    IsSet: true,
                  },
                },
                HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                  Value: 0,
                  IsSet: true,
                },
                VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                  Value: 1,
                  IsSet: true,
                },
                Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                  Value: 5,
                  IsSet: true,
                },
              },
              Set: asciidoc.Set{
                &asciidoc.String{
                  Value: "7",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: 2,
                  IsSet: true,
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
                  Value: "8",
                },
              },
              Blank: false,
            },
            &asciidoc.TableCell{
              Format: nil,
              Set: nil,
              Blank: true,
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
                  IsSet: true,
                },
              },
              Set: asciidoc.Set{
                &asciidoc.String{
                  Value: "9",
                },
              },
              Blank: false,
            },
            &asciidoc.TableCell{
              Format: nil,
              Set: nil,
              Blank: true,
            },
          },
        },
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
                    Value: 2,
                    IsSet: true,
                  },
                  Row: asciidoc.Optional[int]{
                    Value: 1,
                    IsSet: false,
                  },
                },
                HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                  Value: 1,
                  IsSet: true,
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
                  Value: "10",
                },
              },
              Blank: false,
            },
            &asciidoc.TableCell{
              Format: nil,
              Set: nil,
              Blank: true,
            },
            &asciidoc.TableCell{
              Format: nil,
              Set: nil,
              Blank: true,
            },
            &asciidoc.TableCell{
              Format: nil,
              Set: nil,
              Blank: true,
            },
          },
        },
      },
    },
  },
}

var setsUpColumnsCorrectlyIfFirstRowHasCellThatSpansColumns = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: nil,
      ColumnCount: 3,
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
                    Value: 2,
                    IsSet: true,
                  },
                  Row: asciidoc.Optional[int]{
                    Value: 1,
                    IsSet: false,
                  },
                },
                HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                  Value: 2,
                  IsSet: true,
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
                  Value: "AAA",
                },
              },
              Blank: false,
            },
            &asciidoc.TableCell{
              Format: nil,
              Set: nil,
              Blank: true,
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
                  Value: "CCC",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "AAA",
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
                  Value: "BBB",
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
                  Value: "CCC",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "AAA",
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
                  Value: "BBB",
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
                  Value: "CCC",
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

var supportsRepeatingCells = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: nil,
      ColumnCount: 1,
      Set: asciidoc.Set{
        &asciidoc.TableRow{
          Set: asciidoc.Set{
            &asciidoc.TableCell{
              Format: &asciidoc.TableCellFormat{
                Multiplier: asciidoc.Optional[int]{
                  Value: 3,
                  IsSet: true,
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
          },
        },
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
                  Value: "1",
                },
              },
              Blank: false,
            },
          },
        },
        &asciidoc.TableRow{
          Set: asciidoc.Set{
            &asciidoc.TableCell{
              Format: &asciidoc.TableCellFormat{
                Multiplier: asciidoc.Optional[int]{
                  Value: 3,
                  IsSet: true,
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
                  Value: "2",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "b",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "c",
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

var calculatesColnamesCorrectlyWhenUsingImplicitColumnCountAndSingleCellWithColspan = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: nil,
      ColumnCount: 2,
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
                    Value: 2,
                    IsSet: true,
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
                  Value: "Two Columns",
                },
              },
              Blank: false,
            },
            &asciidoc.TableCell{
              Format: nil,
              Set: nil,
              Blank: true,
            },
          },
        },
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
                  Value: "One Column",
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
                  Value: "One Column",
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

var calculatesColnamesCorrectlyWhenUsingImplicitColumnCountAndCellsWithMixedColspans = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: nil,
      ColumnCount: 3,
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
                    Value: 2,
                    IsSet: true,
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
                  Value: "Two Columns",
                },
              },
              Blank: false,
            },
            &asciidoc.TableCell{
              Format: nil,
              Set: nil,
              Blank: true,
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
                  Value: "One Column",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "One Column",
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
                  Value: "One Column",
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
                  Value: "One Column",
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

var assignsUniqueColumnNamesForTableWithImplicitColumnCountAndColspansInFirstRow = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: nil,
      ColumnCount: 5,
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
                  Value: "",
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
                    Value: 2,
                    IsSet: true,
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
                  Value: "Node 0",
                },
              },
              Blank: false,
            },
            &asciidoc.TableCell{
              Format: nil,
              Set: nil,
              Blank: true,
            },
            &asciidoc.TableCell{
              Format: &asciidoc.TableCellFormat{
                Multiplier: asciidoc.Optional[int]{
                  Value: 1,
                  IsSet: false,
                },
                Span: asciidoc.TableCellSpan{
                  Column: asciidoc.Optional[int]{
                    Value: 2,
                    IsSet: true,
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
                  Value: "Node 1",
                },
              },
              Blank: false,
            },
            &asciidoc.TableCell{
              Format: nil,
              Set: nil,
              Blank: true,
            },
          },
        },
        asciidoc.EmptyLine{
          Text: "\n",
        },
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
                  Value: "Host processes",
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
                  Value: "Core 0",
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
                  Value: "Core 1",
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
                  Value: "Core 4",
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
                  Value: "Core 5",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "Guest processes",
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
                  Value: "Core 2",
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
                  Value: "Core 3",
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
                  Value: "Core 6",
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
                  Value: "Core 7",
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

var ignoresCellWithColspanThatExceedsColspec = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TableColumnsAttribute{
          Columns: []*asciidoc.TableColumn{
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 2,
                IsSet: true,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 0,
                IsSet: false,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 1,
                IsSet: false,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 0,
                IsSet: false,
              },
            },
          },
        },
      },
      ColumnCount: 2,
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
                    Value: 3,
                    IsSet: true,
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
              Format: nil,
              Set: nil,
              Blank: true,
            },
            &asciidoc.TableCell{
              Format: nil,
              Set: nil,
              Blank: true,
            },
          },
        },
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
                  Value: "B",
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
                  Value: 1,
                  IsSet: true,
                },
              },
              Set: asciidoc.Set{
                &asciidoc.String{
                  Value: "C",
                },
                &asciidoc.NewLine{},
                asciidoc.EmptyLine{
                  Text: "",
                },
                &asciidoc.String{
                  Value: "more C",
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

var paragraphAndLiteralRepeatedContent = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TableColumnsAttribute{
          Columns: []*asciidoc.TableColumn{
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 1,
                IsSet: false,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 2,
                IsSet: true,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 1,
                IsSet: false,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 4,
                IsSet: true,
              },
            },
          },
        },
      },
      ColumnCount: 1,
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
                  Value: "Paragraphs",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "Literal",
                },
              },
              Blank: false,
            },
          },
        },
        asciidoc.EmptyLine{
          Text: "\n",
        },
        &asciidoc.TableRow{
          Set: asciidoc.Set{
            &asciidoc.TableCell{
              Format: &asciidoc.TableCellFormat{
                Multiplier: asciidoc.Optional[int]{
                  Value: 3,
                  IsSet: true,
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
                  Value: "The discussion about what is good,",
                },
                &asciidoc.NewLine{},
                &asciidoc.String{
                  Value: "what is beautiful, what is noble,",
                },
                &asciidoc.NewLine{},
                &asciidoc.String{
                  Value: "what is pure, and what is true",
                },
                &asciidoc.NewLine{},
                &asciidoc.String{
                  Value: "could always go on.",
                },
                &asciidoc.NewLine{},
                &asciidoc.NewLine{},
                &asciidoc.String{
                  Value: "Why is that important?",
                },
                &asciidoc.NewLine{},
                &asciidoc.String{
                  Value: "Why would I like to do that?",
                },
                &asciidoc.NewLine{},
                &asciidoc.NewLine{},
                &asciidoc.String{
                  Value: "Because that's the only conversation worth having.",
                },
                &asciidoc.NewLine{},
                &asciidoc.NewLine{},
                &asciidoc.String{
                  Value: "And whether it goes on or not after I die, I don't know.",
                },
                &asciidoc.NewLine{},
                &asciidoc.String{
                  Value: "But, I do know that it is the conversation I want to have while I am still alive.",
                },
                &asciidoc.NewLine{},
                &asciidoc.NewLine{},
                &asciidoc.String{
                  Value: "Which means that to me the offer of certainty,",
                },
                &asciidoc.NewLine{},
                &asciidoc.String{
                  Value: "the offer of complete security,",
                },
                &asciidoc.NewLine{},
                &asciidoc.String{
                  Value: "the offer of an impermeable faith that can't give way",
                },
                &asciidoc.NewLine{},
                &asciidoc.String{
                  Value: "is an offer of something not worth having.",
                },
                &asciidoc.NewLine{},
                &asciidoc.NewLine{},
                &asciidoc.String{
                  Value: "I want to live my life taking the risk all the time",
                },
                &asciidoc.NewLine{},
                &asciidoc.String{
                  Value: "that I don't know anything like enough yet...",
                },
                &asciidoc.NewLine{},
                &asciidoc.String{
                  Value: "that I haven't understood enough...",
                },
                &asciidoc.NewLine{},
                &asciidoc.String{
                  Value: "that I can't know enough...",
                },
                &asciidoc.NewLine{},
                &asciidoc.String{
                  Value: "that I am always hungrily operating on the margins",
                },
                &asciidoc.NewLine{},
                &asciidoc.String{
                  Value: "of a potentially great harvest of future knowledge and wisdom.",
                },
                &asciidoc.NewLine{},
                &asciidoc.NewLine{},
                &asciidoc.String{
                  Value: "I wouldn't have it any other way.",
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

var shouldNotSplitParagraphAtLineContainingOnlyblankThatIsDirectlyAdjacentToNonBlankLines = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: nil,
      ColumnCount: 1,
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
                  Value: "paragraph",
                },
                &asciidoc.NewLine{},
                &asciidoc.CharacterReplacementReference{
                  Value: "blank",
                },
                &asciidoc.NewLine{},
                &asciidoc.String{
                  Value: "still one paragraph",
                },
                &asciidoc.NewLine{},
                &asciidoc.CharacterReplacementReference{
                  Value: "blank",
                },
                &asciidoc.NewLine{},
                &asciidoc.String{
                  Value: "still one paragraph",
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

var shouldStripTrailingNewlinesWhenSplittingParagraphs = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: nil,
      ColumnCount: 1,
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
                  Value: "first wrapped",
                },
                &asciidoc.NewLine{},
                &asciidoc.String{
                  Value: "paragraph",
                },
                &asciidoc.NewLine{},
                &asciidoc.NewLine{},
                &asciidoc.String{
                  Value: "second paragraph",
                },
                &asciidoc.NewLine{},
                &asciidoc.NewLine{},
                &asciidoc.String{
                  Value: "third paragraph",
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

var basicAsciiDocCell = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: nil,
      ColumnCount: 1,
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
                &asciidoc.OpenBlock{
                  AttributeList: nil,
                  Delimiter: asciidoc.Delimiter{
                    Type: 7,
                    Length: 2,
                  },
                  Set: asciidoc.Set{
                    &asciidoc.Admonition{
                      AdmonitionType: 1,
                      AttributeList: nil,
                    },
                    &asciidoc.String{
                      Value: "content",
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

var asciiDocTableCellShouldBeWrappedInDivWithClassContent = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: nil,
      ColumnCount: 1,
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
                &asciidoc.String{
                  Value: "AsciiDoc table cell",
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

var doctypeCanBeSetInAsciiDocTableCell = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: nil,
      ColumnCount: 1,
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
                &asciidoc.AttributeEntry{
                  Name: "doctype",
                  Set: asciidoc.Set{
                    &asciidoc.String{
                      Value: "inline",
                    },
                  },
                },
                asciidoc.EmptyLine{
                  Text: "",
                },
                &asciidoc.String{
                  Value: "content",
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

var shouldResetDoctypeToDefaultInAsciiDocTableCell = &asciidoc.Document{
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
        asciidoc.EmptyLine{
          Text: "",
        },
        &asciidoc.Section{
          AttributeList: nil,
          Set: asciidoc.Set{
            asciidoc.EmptyLine{
              Text: "",
            },
            &asciidoc.Table{
              AttributeList: nil,
              ColumnCount: 1,
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
                        &asciidoc.Section{
                          AttributeList: nil,
                          Set: nil,
                          Title: asciidoc.Set{
                            &asciidoc.String{
                              Value: "AsciiDoc Table Cell",
                            },
                          },
                          Level: 0,
                        },
                        asciidoc.EmptyLine{
                          Text: "",
                        },
                        &asciidoc.String{
                          Value: "doctype=",
                        },
                        &asciidoc.UserAttributeReference{
                          Value: "doctype",
                        },
                        &asciidoc.NewLine{},
                        &asciidoc.UserAttributeReference{
                          Value: "backend-html5-doctype-article",
                        },
                        &asciidoc.NewLine{},
                        &asciidoc.UserAttributeReference{
                          Value: "backend-html5-doctype-book",
                        },
                      },
                      Blank: false,
                    },
                  },
                },
              },
            },
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
          Value: "Book Title",
        },
      },
      Level: 0,
    },
  },
}

var shouldUpdateDoctypeRelatedAttributesInAsciiDocTableCellWhenDoctypeIsSet = &asciidoc.Document{
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
              Value: "article",
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
            &asciidoc.Table{
              AttributeList: nil,
              ColumnCount: 1,
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
                        &asciidoc.Section{
                          AttributeList: nil,
                          Set: nil,
                          Title: asciidoc.Set{
                            &asciidoc.String{
                              Value: "AsciiDoc Table Cell",
                            },
                          },
                          Level: 0,
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
                        &asciidoc.String{
                          Value: "doctype=",
                        },
                        &asciidoc.UserAttributeReference{
                          Value: "doctype",
                        },
                        &asciidoc.NewLine{},
                        &asciidoc.UserAttributeReference{
                          Value: "backend-html5-doctype-book",
                        },
                        &asciidoc.NewLine{},
                        &asciidoc.UserAttributeReference{
                          Value: "backend-html5-doctype-article",
                        },
                      },
                      Blank: false,
                    },
                  },
                },
              },
            },
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
          Value: "Document Title",
        },
      },
      Level: 0,
    },
  },
}

var shouldNotAllowAsciiDocTableCellToSetADocumentAttributeThatWasHardSetByTheApi = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: nil,
      ColumnCount: 1,
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
                &asciidoc.AttributeEntry{
                  Name: "icons",
                  Set: nil,
                },
                asciidoc.EmptyLine{
                  Text: "",
                },
                &asciidoc.Paragraph{
                  AttributeList: nil,
                  Set: asciidoc.Set{
                    &asciidoc.String{
                      Value: "This admonition does not have a font-based icon.",
                    },
                  },
                  Admonition: 1,
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

var shouldNotAllowAsciiDocTableCellToSetADocumentAttributeThatWasHardUnsetByTheApi = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: nil,
      ColumnCount: 1,
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
                  AttributeList: nil,
                  Set: asciidoc.Set{
                    &asciidoc.String{
                      Value: "This admonition does not have a font-based icon.",
                    },
                  },
                  Admonition: 1,
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

var shouldKeepAttributeUnsetInAsciiDocTableCellIfUnsetInParentDocument = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeReset{
      Name: "sectids",
    },
    &asciidoc.AttributeReset{
      Name: "table-caption",
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
        &asciidoc.Table{
          AttributeList: asciidoc.AttributeList{
            &asciidoc.TitleAttribute{
              Val: asciidoc.Set{
                &asciidoc.String{
                  Value: "Outer Table",
                },
              },
            },
          },
          ColumnCount: 1,
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
                    asciidoc.EmptyLine{
                      Text: "",
                    },
                    &asciidoc.Section{
                      AttributeList: nil,
                      Set: nil,
                      Title: asciidoc.Set{
                        &asciidoc.String{
                          Value: "Inner Heading",
                        },
                      },
                      Level: 1,
                    },
                    asciidoc.EmptyLine{
                      Text: "",
                    },
                    &asciidoc.Paragraph{
                      AttributeList: asciidoc.AttributeList{
                        &asciidoc.TitleAttribute{
                          Val: asciidoc.Set{
                            &asciidoc.String{
                              Value: "Inner Table",
                            },
                          },
                        },
                      },
                      Set: asciidoc.Set{
                        &asciidoc.String{
                          Value: "!===",
                        },
                        &asciidoc.NewLine{},
                        &asciidoc.String{
                          Value: "! table cell",
                        },
                        &asciidoc.NewLine{},
                        &asciidoc.String{
                          Value: "!===",
                        },
                      },
                      Admonition: 0,
                    },
                  },
                  Blank: false,
                },
              },
            },
          },
        },
      },
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "Outer Heading",
        },
      },
      Level: 1,
    },
  },
}

var shouldAllowAttributeUnsetInParentDocumentToBeSetInAsciiDocTableCell = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeReset{
      Name: "sectids",
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
        &asciidoc.Table{
          AttributeList: nil,
          ColumnCount: 1,
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
                    asciidoc.EmptyLine{
                      Text: "",
                    },
                    &asciidoc.Section{
                      AttributeList: nil,
                      Set: nil,
                      Title: asciidoc.Set{
                        &asciidoc.String{
                          Value: "No ID",
                        },
                      },
                      Level: 1,
                    },
                    asciidoc.EmptyLine{
                      Text: "",
                    },
                    &asciidoc.AttributeEntry{
                      Name: "sectids",
                      Set: nil,
                    },
                    asciidoc.EmptyLine{
                      Text: "",
                    },
                    &asciidoc.Section{
                      AttributeList: nil,
                      Set: nil,
                      Title: asciidoc.Set{
                        &asciidoc.String{
                          Value: "Has ID",
                        },
                      },
                      Level: 1,
                    },
                  },
                  Blank: false,
                },
              },
            },
          },
        },
      },
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "No ID",
        },
      },
      Level: 1,
    },
  },
}

var shouldNotAllowLockedAttributeUnsetInParentDocumentToBeSetInAsciiDocTableCell = &asciidoc.Document{
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
        &asciidoc.Table{
          AttributeList: nil,
          ColumnCount: 1,
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
                    asciidoc.EmptyLine{
                      Text: "",
                    },
                    &asciidoc.Section{
                      AttributeList: nil,
                      Set: nil,
                      Title: asciidoc.Set{
                        &asciidoc.String{
                          Value: "No ID",
                        },
                      },
                      Level: 1,
                    },
                    asciidoc.EmptyLine{
                      Text: "",
                    },
                    &asciidoc.AttributeEntry{
                      Name: "sectids",
                      Set: nil,
                    },
                    asciidoc.EmptyLine{
                      Text: "",
                    },
                    &asciidoc.Section{
                      AttributeList: nil,
                      Set: nil,
                      Title: asciidoc.Set{
                        &asciidoc.String{
                          Value: "Has ID",
                        },
                      },
                      Level: 1,
                    },
                  },
                  Blank: false,
                },
              },
            },
          },
        },
      },
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "No ID",
        },
      },
      Level: 1,
    },
  },
}

var asciiDocContent = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TableColumnsAttribute{
          Columns: []*asciidoc.TableColumn{
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 1,
                IsSet: false,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 0,
                IsSet: false,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 1,
                IsSet: true,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 2,
                IsSet: true,
              },
            },
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 1,
                IsSet: false,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 0,
                IsSet: false,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 1,
                IsSet: true,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 0,
                IsSet: false,
              },
            },
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 1,
                IsSet: false,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 0,
                IsSet: false,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 5,
                IsSet: true,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 1,
                IsSet: true,
              },
            },
          },
        },
      },
      ColumnCount: 3,
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
                  Value: "Name",
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
                  Value: "Backends",
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
                  Value: "Description",
                },
              },
              Blank: false,
            },
          },
        },
        asciidoc.EmptyLine{
          Text: "\n",
        },
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
                  Value: "badges",
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
                  Value: "xhtml11, html5",
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
                &asciidoc.NewLine{},
                &asciidoc.String{
                  Value: "Link badges ('XHTML 1.1' and 'CSS') in document footers.",
                },
                &asciidoc.NewLine{},
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
                          Value: "NOTE",
                        },
                      },
                    },
                  },
                  Set: asciidoc.Set{
                    &asciidoc.String{
                      Value: "The path names of images, icons and scripts are relative path",
                    },
                    &asciidoc.NewLine{},
                    &asciidoc.String{
                      Value: "names to the output document not the source document.",
                    },
                    &asciidoc.NewLine{},
                  },
                },
              },
              Blank: false,
            },
          },
        },
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
                &asciidoc.Anchor{
                  ID: "X97",
                  Set: nil,
                },
                &asciidoc.String{
                  Value: " docinfo, docinfo1, docinfo2",
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
                  Value: "All backends",
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
                &asciidoc.NewLine{},
                &asciidoc.String{
                  Value: "These three attributes control which document information",
                },
                &asciidoc.NewLine{},
                &asciidoc.String{
                  Value: "files will be included in the the header of the output file:",
                },
                &asciidoc.NewLine{},
                asciidoc.EmptyLine{
                  Text: "",
                },
                &asciidoc.String{
                  Value: "docinfo:: Include ",
                },
                &asciidoc.Monospace{
                  AttributeList: nil,
                  Set: asciidoc.Set{
                    asciidoc.SpecialCharacter{
                      Character: "<",
                    },
                    &asciidoc.String{
                      Value: "filename",
                    },
                    asciidoc.SpecialCharacter{
                      Character: ">",
                    },
                    &asciidoc.String{
                      Value: "-docinfo.",
                    },
                    asciidoc.SpecialCharacter{
                      Character: "<",
                    },
                    &asciidoc.String{
                      Value: "ext",
                    },
                    asciidoc.SpecialCharacter{
                      Character: ">",
                    },
                  },
                },
                &asciidoc.NewLine{},
                &asciidoc.String{
                  Value: "docinfo1:: Include ",
                },
                &asciidoc.Monospace{
                  AttributeList: nil,
                  Set: asciidoc.Set{
                    &asciidoc.String{
                      Value: "docinfo.",
                    },
                    asciidoc.SpecialCharacter{
                      Character: "<",
                    },
                    &asciidoc.String{
                      Value: "ext",
                    },
                    asciidoc.SpecialCharacter{
                      Character: ">",
                    },
                  },
                },
                &asciidoc.NewLine{},
                &asciidoc.String{
                  Value: "docinfo2:: Include ",
                },
                &asciidoc.Monospace{
                  AttributeList: nil,
                  Set: asciidoc.Set{
                    &asciidoc.String{
                      Value: "docinfo.",
                    },
                    asciidoc.SpecialCharacter{
                      Character: "<",
                    },
                    &asciidoc.String{
                      Value: "ext",
                    },
                    asciidoc.SpecialCharacter{
                      Character: ">",
                    },
                  },
                },
                &asciidoc.String{
                  Value: " and ",
                },
                &asciidoc.Monospace{
                  AttributeList: nil,
                  Set: asciidoc.Set{
                    asciidoc.SpecialCharacter{
                      Character: "<",
                    },
                    &asciidoc.String{
                      Value: "filename",
                    },
                    asciidoc.SpecialCharacter{
                      Character: ">",
                    },
                    &asciidoc.String{
                      Value: "-docinfo.",
                    },
                    asciidoc.SpecialCharacter{
                      Character: "<",
                    },
                    &asciidoc.String{
                      Value: "ext",
                    },
                    asciidoc.SpecialCharacter{
                      Character: ">",
                    },
                  },
                },
                &asciidoc.NewLine{},
                asciidoc.EmptyLine{
                  Text: "",
                },
                &asciidoc.String{
                  Value: "Where ",
                },
                &asciidoc.Monospace{
                  AttributeList: nil,
                  Set: asciidoc.Set{
                    asciidoc.SpecialCharacter{
                      Character: "<",
                    },
                    &asciidoc.String{
                      Value: "filename",
                    },
                    asciidoc.SpecialCharacter{
                      Character: ">",
                    },
                  },
                },
                &asciidoc.String{
                  Value: " is the file name (sans extension) of the AsciiDoc",
                },
                &asciidoc.NewLine{},
                &asciidoc.String{
                  Value: "input file and ",
                },
                &asciidoc.Monospace{
                  AttributeList: nil,
                  Set: asciidoc.Set{
                    asciidoc.SpecialCharacter{
                      Character: "<",
                    },
                    &asciidoc.String{
                      Value: "ext",
                    },
                    asciidoc.SpecialCharacter{
                      Character: ">",
                    },
                  },
                },
                &asciidoc.String{
                  Value: " is ",
                },
                &asciidoc.Monospace{
                  AttributeList: nil,
                  Set: asciidoc.Set{
                    &asciidoc.String{
                      Value: ".html",
                    },
                  },
                },
                &asciidoc.String{
                  Value: " for HTML outputs or ",
                },
                &asciidoc.Monospace{
                  AttributeList: nil,
                  Set: asciidoc.Set{
                    &asciidoc.String{
                      Value: ".xml",
                    },
                  },
                },
                &asciidoc.String{
                  Value: " for",
                },
                &asciidoc.NewLine{},
                &asciidoc.String{
                  Value: "DocBook outputs. If the input file is the standard input then the",
                },
                &asciidoc.NewLine{},
                &asciidoc.String{
                  Value: "output file name is used.",
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

var shouldPreserveLeadingIndentationInContentsOfAsciiDocTableCellIfContentsStartsWithNewline = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "|===",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "a|",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: " $ command",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "a| paragraph",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "|===",
    },
    &asciidoc.NewLine{},
  },
}

var preprocessorDirectiveOnFirstLineOfAnAsciiDocTableCellShouldBeProcessed = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: nil,
      ColumnCount: 1,
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
                &asciidoc.FileInclude{
                  AttributeList: nil,
                  Set: asciidoc.Set{
                    &asciidoc.String{
                      Value: "fixtures/include-file.adoc",
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

var crossReferenceLinkInAnAsciiDocTableCellShouldResolveToReferenceInMainDocument = &asciidoc.Document{
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
        &asciidoc.Table{
          AttributeList: nil,
          ColumnCount: 1,
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
                    &asciidoc.String{
                      Value: "See ",
                    },
                    &asciidoc.CrossReference{
                      Set: nil,
                      ID: "_more",
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
      },
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "Some",
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
          Value: "content",
        },
        &asciidoc.NewLine{},
      },
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "More",
        },
      },
      Level: 1,
    },
  },
}

var shouldDiscoverAnchorAtStartOfCellAndRegisterItAsAReference = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "The highest peak in the Front Range is ",
    },
    &asciidoc.CrossReference{
      Set: nil,
      ID: "grays-peak",
    },
    &asciidoc.String{
      Value: ", which tops ",
    },
    &asciidoc.CrossReference{
      Set: nil,
      ID: "mount-evans",
    },
    &asciidoc.String{
      Value: " by just a few feet.",
    },
    &asciidoc.NewLine{},
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TableColumnsAttribute{
          Columns: []*asciidoc.TableColumn{
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 1,
                IsSet: false,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 0,
                IsSet: false,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 1,
                IsSet: true,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 6,
                IsSet: true,
              },
            },
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 1,
                IsSet: false,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 0,
                IsSet: false,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 1,
                IsSet: true,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 0,
                IsSet: false,
              },
            },
          },
        },
      },
      ColumnCount: 2,
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
                &asciidoc.Anchor{
                  ID: "mount-evans",
                  Set: asciidoc.Set{
                    &asciidoc.String{
                      Value: "Mount Evans",
                    },
                  },
                },
                &asciidoc.String{
                  Value: "Mount Evans",
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
                  Value: "14,271 feet",
                },
              },
              Blank: false,
            },
          },
        },
        asciidoc.EmptyLine{
          Text: "\n",
        },
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
                  Value: 3,
                  IsSet: true,
                },
              },
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
                  Value: "Grays Peak",
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
                  Value: "14,278 feet",
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

var shouldCatalogAnchorAtStartOfCellInImplicitHeaderRowWhenColumnHasAStyle = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TableColumnsAttribute{
          Columns: []*asciidoc.TableColumn{
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 1,
                IsSet: false,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 0,
                IsSet: false,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 1,
                IsSet: true,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 1,
                IsSet: true,
              },
            },
          },
        },
      },
      ColumnCount: 1,
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
                &asciidoc.Anchor{
                  ID: "foo",
                  Set: asciidoc.Set{
                    &asciidoc.String{
                      Value: "Foo",
                    },
                  },
                },
                &asciidoc.String{
                  Value: "* not AsciiDoc",
                },
              },
              Blank: false,
            },
          },
        },
        asciidoc.EmptyLine{
          Text: "\n",
        },
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
                  Value: "AsciiDoc",
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

var shouldCatalogAnchorAtStartOfCellInExplicitHeaderRowWhenColumnHasAStyle = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "%header",
            },
          },
        },
        &asciidoc.TableColumnsAttribute{
          Columns: []*asciidoc.TableColumn{
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 1,
                IsSet: false,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 0,
                IsSet: false,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 1,
                IsSet: true,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 1,
                IsSet: true,
              },
            },
          },
        },
      },
      ColumnCount: 1,
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
                &asciidoc.Anchor{
                  ID: "foo",
                  Set: asciidoc.Set{
                    &asciidoc.String{
                      Value: "Foo",
                    },
                  },
                },
                &asciidoc.String{
                  Value: "* not AsciiDoc",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "AsciiDoc",
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

var shouldCatalogAnchorAtStartOfCellInFirstRow = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: nil,
      ColumnCount: 1,
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
                &asciidoc.Anchor{
                  ID: "foo",
                  Set: asciidoc.Set{
                    &asciidoc.String{
                      Value: "Foo",
                    },
                  },
                },
                &asciidoc.String{
                  Value: "foo",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "bar",
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

var footnotesShouldNotBeSharedBetweenAnAsciiDocTableCellAndTheMainDocument = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: nil,
      ColumnCount: 1,
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
                &asciidoc.String{
                  Value: "AsciiDoc footnote:[A lightweight markup language.]",
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

var calloutNumbersShouldBeGloballyUniqueIncludingAsciiDocTableCells = &asciidoc.Document{
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
            &asciidoc.Table{
              AttributeList: nil,
              ColumnCount: 1,
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
                                  Value: "yaml",
                                },
                              },
                            },
                          },
                          Delimiter: asciidoc.Delimiter{
                            Type: 5,
                            Length: 4,
                          },
                          LineList: asciidoc.LineList{
                            "key: value <1>",
                          },
                        },
                        &asciidoc.String{
                          Value: "<1> First callout",
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
          },
          Title: asciidoc.Set{
            &asciidoc.String{
              Value: "Section 1",
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
            &asciidoc.Table{
              AttributeList: nil,
              ColumnCount: 1,
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
                                  Value: "yaml",
                                },
                              },
                            },
                          },
                          Delimiter: asciidoc.Delimiter{
                            Type: 5,
                            Length: 4,
                          },
                          LineList: asciidoc.LineList{
                            "key: value <1>",
                          },
                        },
                        &asciidoc.String{
                          Value: "<1> Second callout",
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
          },
          Title: asciidoc.Set{
            &asciidoc.String{
              Value: "Section 2",
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
                      Value: "yaml",
                    },
                  },
                },
              },
              Delimiter: asciidoc.Delimiter{
                Type: 5,
                Length: 4,
              },
              LineList: asciidoc.LineList{
                "key: value <1>",
              },
            },
            &asciidoc.String{
              Value: "<1> Third callout",
            },
            &asciidoc.NewLine{},
          },
          Title: asciidoc.Set{
            &asciidoc.String{
              Value: "Section 3",
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

var compatModeCanBeActivatedInAsciiDocTableCell = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: nil,
      ColumnCount: 1,
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
                &asciidoc.AttributeEntry{
                  Name: "compat-mode",
                  Set: nil,
                },
                asciidoc.EmptyLine{
                  Text: "",
                },
                &asciidoc.String{
                  Value: "The word 'italic' is emphasized.",
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

var compatModeInAsciiDocTableCellInheritsFromParentDocument = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "compat-mode",
      Set: nil,
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "The word 'italic' is emphasized.",
    },
    &asciidoc.NewLine{},
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Paragraph{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TableColumnsAttribute{
          Columns: []*asciidoc.TableColumn{
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 1,
                IsSet: true,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 0,
                IsSet: false,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 1,
                IsSet: false,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 0,
                IsSet: false,
              },
            },
          },
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "|===",
        },
        &asciidoc.NewLine{},
        &asciidoc.String{
          Value: "|The word 'oblique' is emphasized.",
        },
        &asciidoc.NewLine{},
        &asciidoc.String{
          Value: "a|",
        },
        &asciidoc.NewLine{},
        &asciidoc.String{
          Value: "The word 'slanted' is emphasized.",
        },
        &asciidoc.NewLine{},
        &asciidoc.String{
          Value: "|===",
        },
        &asciidoc.NewLine{},
      },
      Admonition: 0,
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "The word 'askew' is emphasized.",
    },
    &asciidoc.NewLine{},
  },
}

var compatModeInAsciiDocTableCellCanBeUnsetIfSetInParentDocument = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "compat-mode",
      Set: nil,
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "The word 'italic' is emphasized.",
    },
    &asciidoc.NewLine{},
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TableColumnsAttribute{
          Columns: []*asciidoc.TableColumn{
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 1,
                IsSet: true,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 0,
                IsSet: false,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 1,
                IsSet: false,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 0,
                IsSet: false,
              },
            },
          },
        },
      },
      ColumnCount: 1,
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
                  Value: "The word 'oblique' is emphasized.",
                },
              },
              Blank: false,
            },
          },
        },
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
                &asciidoc.AttributeReset{
                  Name: "compat-mode",
                },
                asciidoc.EmptyLine{
                  Text: "",
                },
                &asciidoc.String{
                  Value: "The word 'slanted' is not emphasized.",
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
    &asciidoc.String{
      Value: "The word 'askew' is emphasized.",
    },
    &asciidoc.NewLine{},
  },
}

var nestedTable = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TableColumnsAttribute{
          Columns: []*asciidoc.TableColumn{
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 1,
                IsSet: false,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 0,
                IsSet: false,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 1,
                IsSet: true,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 0,
                IsSet: false,
              },
            },
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 1,
                IsSet: false,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 0,
                IsSet: false,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 2,
                IsSet: true,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 1,
                IsSet: true,
              },
            },
          },
        },
      },
      ColumnCount: 2,
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
                  Value: "Normal cell",
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
                  Value: "Cell with nested table",
                },
                &asciidoc.NewLine{},
                &asciidoc.Paragraph{
                  AttributeList: asciidoc.AttributeList{
                    &asciidoc.TableColumnsAttribute{
                      Columns: []*asciidoc.TableColumn{
                        &asciidoc.TableColumn{
                          Multiplier: asciidoc.Optional[int]{
                            Value: 1,
                            IsSet: false,
                          },
                          HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                            Value: 0,
                            IsSet: false,
                          },
                          VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                            Value: 0,
                            IsSet: false,
                          },
                          Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                            Value: 2,
                            IsSet: true,
                          },
                          Percentage: asciidoc.Optional[int]{
                            Value: 0,
                            IsSet: false,
                          },
                          Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                            Value: 0,
                            IsSet: false,
                          },
                        },
                        &asciidoc.TableColumn{
                          Multiplier: asciidoc.Optional[int]{
                            Value: 1,
                            IsSet: false,
                          },
                          HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                            Value: 0,
                            IsSet: false,
                          },
                          VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                            Value: 0,
                            IsSet: false,
                          },
                          Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                            Value: 1,
                            IsSet: true,
                          },
                          Percentage: asciidoc.Optional[int]{
                            Value: 0,
                            IsSet: false,
                          },
                          Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                            Value: 0,
                            IsSet: false,
                          },
                        },
                      },
                    },
                  },
                  Set: asciidoc.Set{
                    &asciidoc.String{
                      Value: "!===",
                    },
                    &asciidoc.NewLine{},
                    &asciidoc.String{
                      Value: "!Nested table cell 1 !Nested table cell 2",
                    },
                    &asciidoc.NewLine{},
                    &asciidoc.String{
                      Value: "!===",
                    },
                  },
                  Admonition: 0,
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

var canSetFormatOfNestedTableToPsv = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TableColumnsAttribute{
          Columns: []*asciidoc.TableColumn{
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 2,
                IsSet: true,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 0,
                IsSet: false,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 1,
                IsSet: false,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 0,
                IsSet: false,
              },
            },
          },
        },
      },
      ColumnCount: 2,
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
                  Value: "normal cell",
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
                  Value: 1,
                  IsSet: true,
                },
              },
              Set: asciidoc.Set{
                &asciidoc.NewLine{},
                &asciidoc.Paragraph{
                  AttributeList: asciidoc.AttributeList{
                    &asciidoc.NamedAttribute{
                      Name: "format",
                      Val: asciidoc.Set{
                        &asciidoc.String{
                          Value: "psv",
                        },
                      },
                      Quote: 0,
                    },
                  },
                  Set: asciidoc.Set{
                    &asciidoc.String{
                      Value: "!===",
                    },
                    &asciidoc.NewLine{},
                    &asciidoc.String{
                      Value: "!nested cell",
                    },
                    &asciidoc.NewLine{},
                    &asciidoc.String{
                      Value: "!===",
                    },
                  },
                  Admonition: 0,
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

var asciiDocTableCellShouldInheritToDirOptionFromParentDocument = &asciidoc.Document{
  Set: asciidoc.Set{
    &asciidoc.String{
      Value: ", parse: true, to_dir: testdir",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "      |===",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "      a|",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "      AsciiDoc table cell",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "      |===",
    },
    &asciidoc.NewLine{},
  },
}

var asciiDocTableCellShouldNotInheritTocSettingFromParentDocument = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Section{ // p0
      AttributeList: nil,
      Set: asciidoc.Set{
        &asciidoc.AttributeEntry{
          Name: "toc",
          Set: nil,
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
            &asciidoc.Table{
              AttributeList: nil,
              ColumnCount: 1,
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
                        &asciidoc.Section{
                          AttributeList: nil,
                          Set: nil,
                          Title: asciidoc.Set{
                            &asciidoc.String{
                              Value: "Section in Nested Document",
                            },
                          },
                          Level: 1,
                        },
                        asciidoc.EmptyLine{
                          Text: "",
                        },
                        &asciidoc.String{
                          Value: "content",
                        },
                      },
                      Blank: false,
                    },
                  },
                },
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
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "Document Title",
        },
      },
      Level: 0,
    },
  },
}

var shouldBeAbleToEnableTocInAnAsciiDocTableCell = &asciidoc.Document{
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
            &asciidoc.Table{
              AttributeList: nil,
              ColumnCount: 1,
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
                        &asciidoc.Section{
                          AttributeList: nil,
                          Set: nil,
                          Title: asciidoc.Set{
                            &asciidoc.String{
                              Value: "Subdocument Title",
                            },
                          },
                          Level: 0,
                        },
                        &asciidoc.AttributeEntry{
                          Name: "toc",
                          Set: nil,
                        },
                        asciidoc.EmptyLine{
                          Text: "",
                        },
                        &asciidoc.Section{
                          AttributeList: nil,
                          Set: nil,
                          Title: asciidoc.Set{
                            &asciidoc.String{
                              Value: "Subdocument Section A",
                            },
                          },
                          Level: 1,
                        },
                        asciidoc.EmptyLine{
                          Text: "",
                        },
                        &asciidoc.String{
                          Value: "content",
                        },
                      },
                      Blank: false,
                    },
                  },
                },
              },
            },
          },
          Title: asciidoc.Set{
            &asciidoc.String{
              Value: "Section A",
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

var shouldBeAbleToEnableTocInAnAsciiDocTableCellEvenIfHardUnsetByApi = &asciidoc.Document{
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
            &asciidoc.Table{
              AttributeList: nil,
              ColumnCount: 1,
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
                        &asciidoc.Section{
                          AttributeList: nil,
                          Set: nil,
                          Title: asciidoc.Set{
                            &asciidoc.String{
                              Value: "Subdocument Title",
                            },
                          },
                          Level: 0,
                        },
                        &asciidoc.AttributeEntry{
                          Name: "toc",
                          Set: nil,
                        },
                        asciidoc.EmptyLine{
                          Text: "",
                        },
                        &asciidoc.Section{
                          AttributeList: nil,
                          Set: nil,
                          Title: asciidoc.Set{
                            &asciidoc.String{
                              Value: "Subdocument Section A",
                            },
                          },
                          Level: 1,
                        },
                        asciidoc.EmptyLine{
                          Text: "",
                        },
                        &asciidoc.String{
                          Value: "content",
                        },
                      },
                      Blank: false,
                    },
                  },
                },
              },
            },
          },
          Title: asciidoc.Set{
            &asciidoc.String{
              Value: "Section A",
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

var shouldBeAbleToEnableTocInBothOuterDocumentAndInAnAsciiDocTableCell = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Section{ // p0
      AttributeList: nil,
      Set: asciidoc.Set{
        &asciidoc.AttributeEntry{
          Name: "toc",
          Set: nil,
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
            &asciidoc.Table{
              AttributeList: nil,
              ColumnCount: 1,
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
                        &asciidoc.Section{
                          AttributeList: nil,
                          Set: nil,
                          Title: asciidoc.Set{
                            &asciidoc.String{
                              Value: "Subdocument Title",
                            },
                          },
                          Level: 0,
                        },
                        &asciidoc.AttributeEntry{
                          Name: "toc",
                          Set: asciidoc.Set{
                            &asciidoc.String{
                              Value: "macro",
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
                                  Value: "#table-cell-toc",
                                },
                              },
                            },
                          },
                          Set: asciidoc.Set{
                            &asciidoc.String{
                              Value: "toc::[]",
                            },
                            &asciidoc.NewLine{},
                          },
                          Admonition: 0,
                        },
                        asciidoc.EmptyLine{
                          Text: "",
                        },
                        &asciidoc.Section{
                          AttributeList: nil,
                          Set: nil,
                          Title: asciidoc.Set{
                            &asciidoc.String{
                              Value: "Subdocument Section A",
                            },
                          },
                          Level: 1,
                        },
                        asciidoc.EmptyLine{
                          Text: "",
                        },
                        &asciidoc.String{
                          Value: "content",
                        },
                      },
                      Blank: false,
                    },
                  },
                },
              },
            },
          },
          Title: asciidoc.Set{
            &asciidoc.String{
              Value: "Section A",
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

var documentInAnAsciiDocTableCellShouldNotSeeDoctitleOfParent = &asciidoc.Document{
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
        &asciidoc.Table{
          AttributeList: asciidoc.AttributeList{
            &asciidoc.TableColumnsAttribute{
              Columns: []*asciidoc.TableColumn{
                &asciidoc.TableColumn{
                  Multiplier: asciidoc.Optional[int]{
                    Value: 1,
                    IsSet: false,
                  },
                  HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                    Value: 0,
                    IsSet: false,
                  },
                  VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                    Value: 0,
                    IsSet: false,
                  },
                  Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                    Value: 1,
                    IsSet: true,
                  },
                  Percentage: asciidoc.Optional[int]{
                    Value: 0,
                    IsSet: false,
                  },
                  Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                    Value: 1,
                    IsSet: true,
                  },
                },
              },
            },
          },
          ColumnCount: 1,
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
                      Value: "AsciiDoc content",
                    },
                  },
                  Blank: false,
                },
              },
            },
          },
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

var cellBackgroundColor = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TableColumnsAttribute{
          Columns: []*asciidoc.TableColumn{
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 1,
                IsSet: false,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 0,
                IsSet: false,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 1,
                IsSet: true,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 2,
                IsSet: true,
              },
            },
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 1,
                IsSet: false,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 0,
                IsSet: false,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 1,
                IsSet: true,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 0,
                IsSet: false,
              },
            },
          },
        },
        &asciidoc.NamedAttribute{
          Name: "options",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "header",
            },
          },
          Quote: 2,
        },
      },
      ColumnCount: 2,
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
                  Value: "{set:cellbgcolor:green}green",
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
                  Value: "{set:cellbgcolor!}",
                },
                &asciidoc.NewLine{},
                &asciidoc.String{
                  Value: "plain",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "{set:cellbgcolor:red}red",
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
                  Value: "{set:cellbgcolor!}",
                },
                &asciidoc.NewLine{},
                &asciidoc.String{
                  Value: "plain",
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

var shouldWarnIfTableBlockIsNotTerminated = &asciidoc.Document{
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
      Value: "|===",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "|",
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

var shouldShowCorrectLineNumberInWarningAboutUnterminatedBlockInsideAsciiDocTableCell = &asciidoc.Document{
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
    &asciidoc.UnorderedListItem{
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "list item",
        },
        &asciidoc.NewLine{},
        &asciidoc.LineBreak{},
        &asciidoc.NewLine{},
        &asciidoc.String{
          Value: "|===",
        },
        &asciidoc.NewLine{},
        &asciidoc.String{
          Value: "|cell",
        },
        &asciidoc.NewLine{},
        &asciidoc.String{
          Value: "a|inside",
        },
      },
      AttributeList: nil,
      Indent: "",
      Marker: "*",
      Checklist: 0,
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "====",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "unterminated example block",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "|===",
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

var customSeparatorForAnAsciiDocTableCell = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Paragraph{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TableColumnsAttribute{
          Columns: []*asciidoc.TableColumn{
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 1,
                IsSet: false,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 0,
                IsSet: false,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 1,
                IsSet: false,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 0,
                IsSet: false,
              },
            },
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 1,
                IsSet: false,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 0,
                IsSet: false,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 1,
                IsSet: false,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 0,
                IsSet: false,
              },
            },
          },
        },
        &asciidoc.NamedAttribute{
          Name: "separator",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "!",
            },
          },
          Quote: 0,
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "|===",
        },
        &asciidoc.NewLine{},
        &asciidoc.String{
          Value: "!Pipe output to vim",
        },
        &asciidoc.NewLine{},
        &asciidoc.String{
          Value: "a!",
        },
        &asciidoc.NewLine{},
      },
      Admonition: 0,
    },
    &asciidoc.Listing{
      AttributeList: nil,
      Delimiter: asciidoc.Delimiter{
        Type: 5,
        Length: 4,
      },
      LineList: asciidoc.LineList{
        "asciidoctor -o - -s test.adoc | view -",
      },
    },
    &asciidoc.String{
      Value: "|===",
    },
    &asciidoc.NewLine{},
  },
}

var tableWithBreakableOptionDocbook5 = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TitleAttribute{
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Table with breakable",
            },
          },
        },
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "%breakable",
            },
          },
        },
      },
      ColumnCount: 2,
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
                  Value: "Item",
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
                  Value: "Quantity",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "Item 1",
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
                  Value: "1",
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

var tableWithUnbreakableOptionDocbook5 = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Table{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TitleAttribute{
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Table with unbreakable",
            },
          },
        },
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "%unbreakable",
            },
          },
        },
      },
      ColumnCount: 2,
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
                  Value: "Item",
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
                  Value: "Quantity",
                },
              },
              Blank: false,
            },
          },
        },
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
                  Value: "Item 1",
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
                  Value: "1",
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

var noImplicitHeaderRowIfCellInFirstLineIsQuotedAndSpansMultipleLines = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Paragraph{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TableColumnsAttribute{
          Columns: []*asciidoc.TableColumn{
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 2,
                IsSet: true,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 0,
                IsSet: false,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 1,
                IsSet: false,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 4,
                IsSet: true,
              },
            },
          },
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: ",===",
        },
        &asciidoc.NewLine{},
        &asciidoc.String{
          Value: "\"A1",
        },
        &asciidoc.NewLine{},
      },
      Admonition: 0,
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "A1 continued\",B1",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "A2,B2",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: ",===",
    },
    &asciidoc.NewLine{},
  },
}

var convertsSimpleDsvTable = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Paragraph{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.NamedAttribute{
          Name: "width",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "75%",
            },
          },
          Quote: 2,
        },
        &asciidoc.NamedAttribute{
          Name: "format",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "dsv",
            },
          },
          Quote: 2,
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "|===",
        },
        &asciidoc.NewLine{},
        &asciidoc.String{
          Value: "root:x:0:0:root:/root:/bin/bash",
        },
        &asciidoc.NewLine{},
        &asciidoc.String{
          Value: "bin:x:1:1:bin:/bin:/sbin/nologin",
        },
        &asciidoc.NewLine{},
        &asciidoc.String{
          Value: "mysql:x:27:27:MySQL\\:Server:/var/lib/mysql:/bin/bash",
        },
        &asciidoc.NewLine{},
        &asciidoc.String{
          Value: "gdm:x:42:42::/var/lib/gdm:/sbin/nologin",
        },
        &asciidoc.NewLine{},
        &asciidoc.String{
          Value: "sshd:x:74:74:Privilege-separated SSH:/var/empty/sshd:/sbin/nologin",
        },
        &asciidoc.NewLine{},
        &asciidoc.String{
          Value: "nobody:x:99:99:Nobody:/:/sbin/nologin",
        },
        &asciidoc.NewLine{},
        &asciidoc.String{
          Value: "|===",
        },
        &asciidoc.NewLine{},
      },
      Admonition: 0,
    },
  },
}

var dsvFormatShorthand = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: ":===",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "a:b:c",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "1:2:3",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: ":===",
    },
    &asciidoc.NewLine{},
  },
}

var singleCellInDsvTableShouldOnlyProduceSingleRow = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: ":===",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "single cell",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: ":===",
    },
    &asciidoc.NewLine{},
  },
}

var shouldTreatTrailingColonAsAnEmptyCell = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: ":===",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "A1:",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "B1:B2",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "C1:C2",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: ":===",
    },
    &asciidoc.NewLine{},
  },
}

var shouldTreatTrailingCommaAsAnEmptyCell = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: ",===",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "A1,",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "B1,B2",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "C1,C2",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: ",===",
    },
    &asciidoc.NewLine{},
  },
}

var shouldLogErrorButNotCrashIfCellDataHasUnclosedQuote = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: ",===",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "a,b",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "c,\"",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: ",===",
    },
    &asciidoc.NewLine{},
  },
}

var shouldPreserveNewlinesInQuotedCsvValues = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Paragraph{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TableColumnsAttribute{
          Columns: []*asciidoc.TableColumn{
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 1,
                IsSet: false,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 0,
                IsSet: false,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 1,
                IsSet: true,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 0,
                IsSet: false,
              },
            },
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 1,
                IsSet: false,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 0,
                IsSet: false,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 1,
                IsSet: true,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 0,
                IsSet: false,
              },
            },
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 1,
                IsSet: false,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 0,
                IsSet: false,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 1,
                IsSet: true,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 4,
                IsSet: true,
              },
            },
          },
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: ",===",
        },
        &asciidoc.NewLine{},
        &asciidoc.String{
          Value: "\"A",
        },
        &asciidoc.NewLine{},
        &asciidoc.String{
          Value: "B",
        },
        &asciidoc.NewLine{},
        &asciidoc.String{
          Value: "C\",\"one",
        },
        &asciidoc.NewLine{},
      },
      Admonition: 0,
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "two",
    },
    &asciidoc.NewLine{},
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "three\",\"do",
    },
    &asciidoc.NewLine{},
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "re",
    },
    &asciidoc.NewLine{},
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "me\"",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: ",===",
    },
    &asciidoc.NewLine{},
  },
}

var shouldNotDropTrailingEmptyCellInTsvDataWhenLoadedFromAnIncludeFile = &asciidoc.Document{
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
              Value: "%header",
            },
          },
        },
        &asciidoc.NamedAttribute{
          Name: "format",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "tsv",
            },
          },
          Quote: 0,
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "|===",
        },
        &asciidoc.NewLine{},
      },
      Admonition: 0,
    },
    &asciidoc.FileInclude{
      AttributeList: nil,
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "fixtures/data.tsv",
        },
      },
    },
    &asciidoc.String{
      Value: "|===",
    },
    &asciidoc.NewLine{},
  },
}

var mixedUnquotedRecordsAndQuotedRecordsWithEscapedQuotesCommasAndWrappedLines = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Paragraph{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.NamedAttribute{
          Name: "format",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "csv",
            },
          },
          Quote: 2,
        },
        &asciidoc.NamedAttribute{
          Name: "options",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "header",
            },
          },
          Quote: 2,
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "|===",
        },
        &asciidoc.NewLine{},
        &asciidoc.String{
          Value: "Year,Make,Model,Description,Price",
        },
        &asciidoc.NewLine{},
        &asciidoc.String{
          Value: "1997,Ford,E350,\"ac, abs, moon\",3000.00",
        },
        &asciidoc.NewLine{},
        &asciidoc.String{
          Value: "1999,Chevy,\"Venture \"\"Extended Edition\"\"\",\"\",4900.00",
        },
        &asciidoc.NewLine{},
        &asciidoc.String{
          Value: "1999,Chevy,\"Venture \"\"Extended Edition, Very Large\"\"\",,5000.00",
        },
        &asciidoc.NewLine{},
        &asciidoc.String{
          Value: "1996,Jeep,Grand Cherokee,\"MUST SELL!",
        },
        &asciidoc.NewLine{},
        &asciidoc.String{
          Value: "air, moon roof, loaded\",4799.00",
        },
        &asciidoc.NewLine{},
        &asciidoc.String{
          Value: "2000,Toyota,Tundra,\"\"\"This one's gonna to blow you're socks off,\"\" per the sticker\",10000.00",
        },
        &asciidoc.NewLine{},
        &asciidoc.String{
          Value: "2000,Toyota,Tundra,\"Check it, \"\"this one's gonna to blow you're socks off\"\", per the sticker\",10000.00",
        },
        &asciidoc.NewLine{},
        &asciidoc.String{
          Value: "|===",
        },
        &asciidoc.NewLine{},
      },
      Admonition: 0,
    },
  },
}

var shouldAllowQuotesAroundACsvValueToBeOnTheirOwnLines = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Paragraph{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TableColumnsAttribute{
          Columns: []*asciidoc.TableColumn{
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 2,
                IsSet: true,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 0,
                IsSet: false,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 1,
                IsSet: false,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 0,
                IsSet: false,
              },
            },
          },
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: ",===",
        },
        &asciidoc.NewLine{},
        &asciidoc.String{
          Value: "\"",
        },
        &asciidoc.NewLine{},
        &asciidoc.String{
          Value: "A",
        },
        &asciidoc.NewLine{},
        &asciidoc.String{
          Value: "\",\"",
        },
        &asciidoc.NewLine{},
        &asciidoc.String{
          Value: "B",
        },
        &asciidoc.NewLine{},
        &asciidoc.String{
          Value: "\"",
        },
        &asciidoc.NewLine{},
        &asciidoc.String{
          Value: ",===",
        },
        &asciidoc.NewLine{},
      },
      Admonition: 0,
    },
  },
}

var csvFormatShorthand = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: ",===",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "a,b,c",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "1,2,3",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: ",===",
    },
    &asciidoc.NewLine{},
  },
}

var customCsvSeparator = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Paragraph{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.NamedAttribute{
          Name: "format",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "csv",
            },
          },
          Quote: 0,
        },
        &asciidoc.NamedAttribute{
          Name: "separator",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: ";",
            },
          },
          Quote: 0,
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "|===",
        },
        &asciidoc.NewLine{},
        &asciidoc.String{
          Value: "a;b;c",
        },
        &asciidoc.NewLine{},
        &asciidoc.String{
          Value: "1;2;3",
        },
        &asciidoc.NewLine{},
        &asciidoc.String{
          Value: "|===",
        },
        &asciidoc.NewLine{},
      },
      Admonition: 0,
    },
  },
}

var singleCellInCsvTableShouldOnlyProduceSingleRow = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: ",===",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "single cell",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: ",===",
    },
    &asciidoc.NewLine{},
  },
}

var cellFormattedWithAsciiDocStyle = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Paragraph{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TableColumnsAttribute{
          Columns: []*asciidoc.TableColumn{
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 1,
                IsSet: false,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 0,
                IsSet: false,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 1,
                IsSet: true,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 0,
                IsSet: false,
              },
            },
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 1,
                IsSet: false,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 0,
                IsSet: false,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 1,
                IsSet: true,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 0,
                IsSet: false,
              },
            },
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 1,
                IsSet: false,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 0,
                IsSet: false,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 1,
                IsSet: true,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 1,
                IsSet: true,
              },
            },
          },
        },
        &asciidoc.NamedAttribute{
          Name: "separator",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: ";",
            },
          },
          Quote: 0,
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: ",===",
        },
        &asciidoc.NewLine{},
        &asciidoc.String{
          Value: "element;description;example",
        },
        &asciidoc.NewLine{},
      },
      Admonition: 0,
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "thematic break,a visible break; also known as a horizontal rule;---",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: ",===",
    },
    &asciidoc.NewLine{},
  },
}

var shouldStripWhitespaceAroundContentsOfAsciiDocCell = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Paragraph{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TableColumnsAttribute{
          Columns: []*asciidoc.TableColumn{
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 1,
                IsSet: false,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 0,
                IsSet: false,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 1,
                IsSet: true,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 0,
                IsSet: false,
              },
            },
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 1,
                IsSet: false,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 0,
                IsSet: false,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 1,
                IsSet: true,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 0,
                IsSet: false,
              },
            },
            &asciidoc.TableColumn{
              Multiplier: asciidoc.Optional[int]{
                Value: 1,
                IsSet: false,
              },
              HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
                Value: 0,
                IsSet: false,
              },
              VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
                Value: 0,
                IsSet: false,
              },
              Width: asciidoc.Optional[asciidoc.TableColumnWidth]{
                Value: 1,
                IsSet: true,
              },
              Percentage: asciidoc.Optional[int]{
                Value: 0,
                IsSet: false,
              },
              Style: asciidoc.Optional[asciidoc.TableCellStyle]{
                Value: 1,
                IsSet: true,
              },
            },
          },
        },
        &asciidoc.NamedAttribute{
          Name: "separator",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: ";",
            },
          },
          Quote: 0,
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: ",===",
        },
        &asciidoc.NewLine{},
        &asciidoc.String{
          Value: "element;description;example",
        },
        &asciidoc.NewLine{},
      },
      Admonition: 0,
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "paragraph;contiguous lines of words and phrases;\"",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "  one sentence, one line",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "  \"",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: ",===",
    },
    &asciidoc.NewLine{},
  },
}


