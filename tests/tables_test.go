package tests

import (
	"testing"

	"github.com/project-chip/alchemy/asciidoc"
)

func TestTables(t *testing.T) {
	tablesTests.run(t)
}

var tablesTests = parseTests{

	{"converts simple psv table", "asciidoctor/tables_test_converts_simple_psv_table.adoc", tablesTestConvertsSimplePsvTable, nil},

	{"should add direction CSS class if float attribute is set on table", "asciidoctor/tables_test_should_add_direction_css_class_if_float_attribute_is_set_on_table.adoc", tablesTestShouldAddDirectionCssClassIfFloatAttributeIsSetOnTable, nil},

	{"should set stripes class if stripes option is set", "asciidoctor/tables_test_should_set_stripes_class_if_stripes_option_is_set.adoc", tablesTestShouldSetStripesClassIfStripesOptionIsSet, nil},

	{"outputs a caption on simple psv table", "asciidoctor/tables_test_outputs_a_caption_on_simple_psv_table.adoc", tablesTestOutputsACaptionOnSimplePsvTable, nil},

	{"only increments table counter for tables that have a title", "asciidoctor/tables_test_only_increments_table_counter_for_tables_that_have_a_title.adoc", tablesTestOnlyIncrementsTableCounterForTablesThatHaveATitle, nil},

	{"uses explicit caption in front of title in place of default caption and number", "asciidoctor/tables_test_uses_explicit_caption_in_front_of_title_in_place_of_default_caption_and_number.adoc", tablesTestUsesExplicitCaptionInFrontOfTitleInPlaceOfDefaultCaptionAndNumber, nil},

	{"disables caption when caption attribute on table is empty", "asciidoctor/tables_test_disables_caption_when_caption_attribute_on_table_is_empty.adoc", tablesTestDisablesCaptionWhenCaptionAttributeOnTableIsEmpty, nil},

	{"disables caption when caption attribute on table is empty string", "asciidoctor/tables_test_disables_caption_when_caption_attribute_on_table_is_empty_string.adoc", tablesTestDisablesCaptionWhenCaptionAttributeOnTableIsEmptyString, nil},

	{"disables caption on table when table-caption document attribute is unset", "asciidoctor/tables_test_disables_caption_on_table_when_table_caption_document_attribute_is_unset.adoc", tablesTestDisablesCaptionOnTableWhenTableCaptionDocumentAttributeIsUnset, nil},

	{"ignores escaped separators", "asciidoctor/tables_test_ignores_escaped_separators.adoc", tablesTestIgnoresEscapedSeparators, nil},

	{"preserves escaped delimiters at the end of the line", "asciidoctor/tables_test_preserves_escaped_delimiters_at_the_end_of_the_line.adoc", tablesTestPreservesEscapedDelimitersAtTheEndOfTheLine, nil},

	{"should treat trailing pipe as an empty cell", "asciidoctor/tables_test_should_treat_trailing_pipe_as_an_empty_cell.adoc", tablesTestShouldTreatTrailingPipeAsAnEmptyCell, nil},

	{"should auto recover with warning if missing leading separator on first cell", "asciidoctor/tables_test_should_auto_recover_with_warning_if_missing_leading_separator_on_first_cell.adoc", tablesTestShouldAutoRecoverWithWarningIfMissingLeadingSeparatorOnFirstCell, nil},

	{"performs normal substitutions on cell content", "asciidoctor/tables_test_performs_normal_substitutions_on_cell_content.adoc", tablesTestPerformsNormalSubstitutionsOnCellContent, nil},

	{"should only substitute specialchars for literal table cells", "asciidoctor/tables_test_should_only_substitute_specialchars_for_literal_table_cells.adoc", tablesTestShouldOnlySubstituteSpecialcharsForLiteralTableCells, nil},

	{"should preserving leading spaces but not leading newlines or trailing spaces in literal table cells", "asciidoctor/tables_test_should_preserving_leading_spaces_but_not_leading_newlines_or_trailing_spaces_in_literal_table_cells.adoc", tablesTestShouldPreservingLeadingSpacesButNotLeadingNewlinesOrTrailingSpacesInLiteralTableCells, nil},

	{"should ignore v table cell style", "asciidoctor/tables_test_should_ignore_v_table_cell_style.adoc", tablesTestShouldIgnoreVTableCellStyle, nil},

	{"table and column width not assigned when autowidth option is specified", "asciidoctor/tables_test_table_and_column_width_not_assigned_when_autowidth_option_is_specified.adoc", tablesTestTableAndColumnWidthNotAssignedWhenAutowidthOptionIsSpecified, nil},

	{"does not assign column width for autowidth columns in HTML output", "asciidoctor/tables_test_does_not_assign_column_width_for_autowidth_columns_in_html_output.adoc", tablesTestDoesNotAssignColumnWidthForAutowidthColumnsInHtmlOutput, nil},

	{"can assign autowidth to all columns even when table has a width", "asciidoctor/tables_test_can_assign_autowidth_to_all_columns_even_when_table_has_a_width.adoc", tablesTestCanAssignAutowidthToAllColumnsEvenWhenTableHasAWidth, nil},

	{"equally distributes remaining column width to autowidth columns in DocBook output", "asciidoctor/tables_test_equally_distributes_remaining_column_width_to_autowidth_columns_in_doc_book_output.adoc", tablesTestEquallyDistributesRemainingColumnWidthToAutowidthColumnsInDocBookOutput, nil},

	{"should compute column widths based on pagewidth when width is set on table in DocBook output", "asciidoctor/tables_test_should_compute_column_widths_based_on_pagewidth_when_width_is_set_on_table_in_doc_book_output.adoc", tablesTestShouldComputeColumnWidthsBasedOnPagewidthWhenWidthIsSetOnTableInDocBookOutput, nil},

	{"explicit table width is used even when autowidth option is specified", "asciidoctor/tables_test_explicit_table_width_is_used_even_when_autowidth_option_is_specified.adoc", tablesTestExplicitTableWidthIsUsedEvenWhenAutowidthOptionIsSpecified, nil},

	{"first row sets number of columns when not specified", "asciidoctor/tables_test_first_row_sets_number_of_columns_when_not_specified.adoc", tablesTestFirstRowSetsNumberOfColumnsWhenNotSpecified, nil},

	{"colspec attribute using asterisk syntax sets number of columns", "asciidoctor/tables_test_colspec_attribute_using_asterisk_syntax_sets_number_of_columns.adoc", tablesTestColspecAttributeUsingAsteriskSyntaxSetsNumberOfColumns, nil},

	{"table with explicit column count can have multiple rows on a single line", "asciidoctor/tables_test_table_with_explicit_column_count_can_have_multiple_rows_on_a_single_line.adoc", tablesTestTableWithExplicitColumnCountCanHaveMultipleRowsOnASingleLine, nil},

	{"table with explicit deprecated colspec syntax can have multiple rows on a single line", "asciidoctor/tables_test_table_with_explicit_deprecated_colspec_syntax_can_have_multiple_rows_on_a_single_line.adoc", tablesTestTableWithExplicitDeprecatedColspecSyntaxCanHaveMultipleRowsOnASingleLine, nil},

	{"columns are added for empty records in colspec attribute", "asciidoctor/tables_test_columns_are_added_for_empty_records_in_colspec_attribute.adoc", tablesTestColumnsAreAddedForEmptyRecordsInColspecAttribute, nil},

	{"cols may be separated by semi-colon instead of comma", "asciidoctor/tables_test_cols_may_be_separated_by_semi_colon_instead_of_comma.adoc", tablesTestColsMayBeSeparatedBySemiColonInsteadOfComma, nil},

	{"cols attribute may include spaces", "asciidoctor/tables_test_cols_attribute_may_include_spaces.adoc", tablesTestColsAttributeMayIncludeSpaces, nil},

	{"blank cols attribute should be ignored", "asciidoctor/tables_test_blank_cols_attribute_should_be_ignored.adoc", tablesTestBlankColsAttributeShouldBeIgnored, nil},

	{"empty cols attribute should be ignored", "asciidoctor/tables_test_empty_cols_attribute_should_be_ignored.adoc", tablesTestEmptyColsAttributeShouldBeIgnored, nil},

	{"table with header and footer", "asciidoctor/tables_test_table_with_header_and_footer.adoc", tablesTestTableWithHeaderAndFooter, nil},

	{"table with header and footer docbook", "asciidoctor/tables_test_table_with_header_and_footer_docbook.adoc", tablesTestTableWithHeaderAndFooterDocbook, nil},

	{"should set horizontal and vertical alignment when converting to DocBook", "asciidoctor/tables_test_should_set_horizontal_and_vertical_alignment_when_converting_to_doc_book.adoc", tablesTestShouldSetHorizontalAndVerticalAlignmentWhenConvertingToDocBook, nil},

	{"should preserve frame value ends when converting to HTML", "asciidoctor/tables_test_should_preserve_frame_value_ends_when_converting_to_html.adoc", tablesTestShouldPreserveFrameValueEndsWhenConvertingToHtml, nil},

	{"should normalize frame value topbot as ends when converting to HTML", "asciidoctor/tables_test_should_normalize_frame_value_topbot_as_ends_when_converting_to_html.adoc", tablesTestShouldNormalizeFrameValueTopbotAsEndsWhenConvertingToHtml, nil},

	{"should preserve frame value topbot when converting to DocBook", "asciidoctor/tables_test_should_preserve_frame_value_topbot_when_converting_to_doc_book.adoc", tablesTestShouldPreserveFrameValueTopbotWhenConvertingToDocBook, nil},

	{"should convert frame value ends to topbot when converting to DocBook", "asciidoctor/tables_test_should_convert_frame_value_ends_to_topbot_when_converting_to_doc_book.adoc", tablesTestShouldConvertFrameValueEndsToTopbotWhenConvertingToDocBook, nil},

	{"table with implicit header row", "asciidoctor/tables_test_table_with_implicit_header_row.adoc", tablesTestTableWithImplicitHeaderRow, nil},

	{"table with implicit header row only", "asciidoctor/tables_test_table_with_implicit_header_row_only.adoc", tablesTestTableWithImplicitHeaderRowOnly, nil},

	{"table with implicit header row when other options set", "asciidoctor/tables_test_table_with_implicit_header_row_when_other_options_set.adoc", tablesTestTableWithImplicitHeaderRowWhenOtherOptionsSet, nil},

	{"no implicit header row if second line not blank", "asciidoctor/tables_test_no_implicit_header_row_if_second_line_not_blank.adoc", tablesTestNoImplicitHeaderRowIfSecondLineNotBlank, nil},

	{"no implicit header row if cell in first line spans multiple lines", "asciidoctor/tables_test_no_implicit_header_row_if_cell_in_first_line_spans_multiple_lines.adoc", tablesTestNoImplicitHeaderRowIfCellInFirstLineSpansMultipleLines, nil},

	{"should format first cell as literal if there is no implicit header row and column has l style", "asciidoctor/tables_test_should_format_first_cell_as_literal_if_there_is_no_implicit_header_row_and_column_has_l_style.adoc", tablesTestShouldFormatFirstCellAsLiteralIfThereIsNoImplicitHeaderRowAndColumnHasLStyle, nil},

	{"should format first cell as AsciiDoc if there is no implicit header row and column has a style", "asciidoctor/tables_test_should_format_first_cell_as_ascii_doc_if_there_is_no_implicit_header_row_and_column_has_a_style.adoc", tablesTestShouldFormatFirstCellAsAsciiDocIfThereIsNoImplicitHeaderRowAndColumnHasAStyle, nil},

	{"should interpret leading indent if first cell is AsciiDoc and there is no implicit header row", "asciidoctor/tables_test_should_interpret_leading_indent_if_first_cell_is_ascii_doc_and_there_is_no_implicit_header_row.adoc", tablesTestShouldInterpretLeadingIndentIfFirstCellIsAsciiDocAndThereIsNoImplicitHeaderRow, nil},

	{"should format first cell as AsciiDoc if there is no implicit header row and cell has a style", "asciidoctor/tables_test_should_format_first_cell_as_ascii_doc_if_there_is_no_implicit_header_row_and_cell_has_a_style.adoc", tablesTestShouldFormatFirstCellAsAsciiDocIfThereIsNoImplicitHeaderRowAndCellHasAStyle, nil},

	{"no implicit header row if AsciiDoc cell in first line spans multiple lines", "asciidoctor/tables_test_no_implicit_header_row_if_ascii_doc_cell_in_first_line_spans_multiple_lines.adoc", tablesTestNoImplicitHeaderRowIfAsciiDocCellInFirstLineSpansMultipleLines, nil},

	{"no implicit header row if first line blank", "asciidoctor/tables_test_no_implicit_header_row_if_first_line_blank.adoc", tablesTestNoImplicitHeaderRowIfFirstLineBlank, nil},

	{"no implicit header row if noheader option is specified", "asciidoctor/tables_test_no_implicit_header_row_if_noheader_option_is_specified.adoc", tablesTestNoImplicitHeaderRowIfNoheaderOptionIsSpecified, nil},

	{"styles not applied to header cells", "asciidoctor/tables_test_styles_not_applied_to_header_cells.adoc", tablesTestStylesNotAppliedToHeaderCells, nil},

	{"should apply text formatting to cells in implicit header row when column has a style", "asciidoctor/tables_test_should_apply_text_formatting_to_cells_in_implicit_header_row_when_column_has_a_style.adoc", tablesTestShouldApplyTextFormattingToCellsInImplicitHeaderRowWhenColumnHasAStyle, nil},

	{"should apply style and text formatting to cells in first row if no implicit header", "asciidoctor/tables_test_should_apply_style_and_text_formatting_to_cells_in_first_row_if_no_implicit_header.adoc", tablesTestShouldApplyStyleAndTextFormattingToCellsInFirstRowIfNoImplicitHeader, nil},

	{"vertical table headers use th element instead of header class", "asciidoctor/tables_test_vertical_table_headers_use_th_element_instead_of_header_class.adoc", tablesTestVerticalTableHeadersUseThElementInsteadOfHeaderClass, nil},

	{"supports horizontal and vertical source data with blank lines and table header", "asciidoctor/tables_test_supports_horizontal_and_vertical_source_data_with_blank_lines_and_table_header.adoc", tablesTestSupportsHorizontalAndVerticalSourceDataWithBlankLinesAndTableHeader, nil},

	{"percentages as column widths", "asciidoctor/tables_test_percentages_as_column_widths.adoc", tablesTestPercentagesAsColumnWidths, nil},

	{"spans, alignments and styles", "asciidoctor/tables_test_spans_alignments_and_styles.adoc", tablesTestSpansAlignmentsAndStyles, nil},

	{"sets up columns correctly if first row has cell that spans columns", "asciidoctor/tables_test_sets_up_columns_correctly_if_first_row_has_cell_that_spans_columns.adoc", tablesTestSetsUpColumnsCorrectlyIfFirstRowHasCellThatSpansColumns, nil},

	{"supports repeating cells", "asciidoctor/tables_test_supports_repeating_cells.adoc", tablesTestSupportsRepeatingCells, nil},

	{"calculates colnames correctly when using implicit column count and single cell with colspan", "asciidoctor/tables_test_calculates_colnames_correctly_when_using_implicit_column_count_and_single_cell_with_colspan.adoc", tablesTestCalculatesColnamesCorrectlyWhenUsingImplicitColumnCountAndSingleCellWithColspan, nil},

	{"calculates colnames correctly when using implicit column count and cells with mixed colspans", "asciidoctor/tables_test_calculates_colnames_correctly_when_using_implicit_column_count_and_cells_with_mixed_colspans.adoc", tablesTestCalculatesColnamesCorrectlyWhenUsingImplicitColumnCountAndCellsWithMixedColspans, nil},

	{"assigns unique column names for table with implicit column count and colspans in first row", "asciidoctor/tables_test_assigns_unique_column_names_for_table_with_implicit_column_count_and_colspans_in_first_row.adoc", tablesTestAssignsUniqueColumnNamesForTableWithImplicitColumnCountAndColspansInFirstRow, nil},

	{"should drop row but preserve remaining rows after cell with colspan exceeds number of columns", "asciidoctor/tables_test_should_drop_row_but_preserve_remaining_rows_after_cell_with_colspan_exceeds_number_of_columns.adoc", tablesTestShouldDropRowButPreserveRemainingRowsAfterCellWithColspanExceedsNumberOfColumns, nil},

	{"should drop last row if last cell in table has colspan that exceeds specified number of columns", "asciidoctor/tables_test_should_drop_last_row_if_last_cell_in_table_has_colspan_that_exceeds_specified_number_of_columns.adoc", tablesTestShouldDropLastRowIfLastCellInTableHasColspanThatExceedsSpecifiedNumberOfColumns, nil},

	{"should drop last row if last cell in table has colspan that exceeds implicit number of columns", "asciidoctor/tables_test_should_drop_last_row_if_last_cell_in_table_has_colspan_that_exceeds_implicit_number_of_columns.adoc", tablesTestShouldDropLastRowIfLastCellInTableHasColspanThatExceedsImplicitNumberOfColumns, nil},

	{"should take colspan into account when taking cells for row", "asciidoctor/tables_test_should_take_colspan_into_account_when_taking_cells_for_row.adoc", tablesTestShouldTakeColspanIntoAccountWhenTakingCellsForRow, nil},

	{"should drop incomplete row at end of table and log an error", "asciidoctor/tables_test_should_drop_incomplete_row_at_end_of_table_and_log_an_error.adoc", tablesTestShouldDropIncompleteRowAtEndOfTableAndLogAnError, nil},

	{"should apply cell style for column to repeated content", "asciidoctor/tables_test_should_apply_cell_style_for_column_to_repeated_content.adoc", tablesTestShouldApplyCellStyleForColumnToRepeatedContent, nil},

	{"should not split paragraph at line containing only {blank} that is directly adjacent to non-blank lines", "asciidoctor/tables_test_should_not_split_paragraph_at_line_containing_only_{blank}_that_is_directly_adjacent_to_non_blank_lines.adoc", tablesTestShouldNotSplitParagraphAtLineContainingOnlyblankThatIsDirectlyAdjacentToNonBlankLines, nil},

	{"should strip trailing newlines when splitting paragraphs", "asciidoctor/tables_test_should_strip_trailing_newlines_when_splitting_paragraphs.adoc", tablesTestShouldStripTrailingNewlinesWhenSplittingParagraphs, nil},

	{"basic AsciiDoc cell", "asciidoctor/tables_test_basic_ascii_doc_cell.adoc", tablesTestBasicAsciiDocCell, nil},

	{"AsciiDoc table cell should be wrapped in div with class \"content\"", "asciidoctor/tables_test_ascii_doc_table_cell_should_be_wrapped_in_div_with_class__content.adoc", tablesTestAsciiDocTableCellShouldBeWrappedInDivWithClassContent, nil},

	{"doctype can be set in AsciiDoc table cell", "asciidoctor/tables_test_doctype_can_be_set_in_ascii_doc_table_cell.adoc", tablesTestDoctypeCanBeSetInAsciiDocTableCell, nil},

	{"should reset doctype to default in AsciiDoc table cell", "asciidoctor/tables_test_should_reset_doctype_to_default_in_ascii_doc_table_cell.adoc", tablesTestShouldResetDoctypeToDefaultInAsciiDocTableCell, nil},

	{"should update doctype-related attributes in AsciiDoc table cell when doctype is set", "asciidoctor/tables_test_should_update_doctype_related_attributes_in_ascii_doc_table_cell_when_doctype_is_set.adoc", tablesTestShouldUpdateDoctypeRelatedAttributesInAsciiDocTableCellWhenDoctypeIsSet, nil},

	{"should not allow AsciiDoc table cell to set a document attribute that was hard set by the API", "asciidoctor/tables_test_should_not_allow_ascii_doc_table_cell_to_set_a_document_attribute_that_was_hard_set_by_the_api.adoc", tablesTestShouldNotAllowAsciiDocTableCellToSetADocumentAttributeThatWasHardSetByTheApi, nil},

	{"should not allow AsciiDoc table cell to set a document attribute that was hard unset by the API", "asciidoctor/tables_test_should_not_allow_ascii_doc_table_cell_to_set_a_document_attribute_that_was_hard_unset_by_the_api.adoc", tablesTestShouldNotAllowAsciiDocTableCellToSetADocumentAttributeThatWasHardUnsetByTheApi, nil},

	{"should keep attribute unset in AsciiDoc table cell if unset in parent document", "asciidoctor/tables_test_should_keep_attribute_unset_in_ascii_doc_table_cell_if_unset_in_parent_document.adoc", tablesTestShouldKeepAttributeUnsetInAsciiDocTableCellIfUnsetInParentDocument, nil},

	{"should allow attribute unset in parent document to be set in AsciiDoc table cell", "asciidoctor/tables_test_should_allow_attribute_unset_in_parent_document_to_be_set_in_ascii_doc_table_cell.adoc", tablesTestShouldAllowAttributeUnsetInParentDocumentToBeSetInAsciiDocTableCell, nil},

	{"should not allow locked attribute unset in parent document to be set in AsciiDoc table cell", "asciidoctor/tables_test_should_not_allow_locked_attribute_unset_in_parent_document_to_be_set_in_ascii_doc_table_cell.adoc", tablesTestShouldNotAllowLockedAttributeUnsetInParentDocumentToBeSetInAsciiDocTableCell, nil},

	{"AsciiDoc content", "asciidoctor/tables_test_ascii_doc_content.adoc", tablesTestAsciiDocContent, nil},

	{"should preserve leading indentation in contents of AsciiDoc table cell if contents starts with newline", "asciidoctor/tables_test_should_preserve_leading_indentation_in_contents_of_ascii_doc_table_cell_if_contents_starts_with_newline.adoc", tablesTestShouldPreserveLeadingIndentationInContentsOfAsciiDocTableCellIfContentsStartsWithNewline, nil},

	{"preprocessor directive on first line of an AsciiDoc table cell should be processed", "asciidoctor/tables_test_preprocessor_directive_on_first_line_of_an_ascii_doc_table_cell_should_be_processed.adoc", tablesTestPreprocessorDirectiveOnFirstLineOfAnAsciiDocTableCellShouldBeProcessed, nil},

	{"error about unresolved preprocessor directive on first line of an AsciiDoc table cell should have correct cursor", "asciidoctor/tables_test_error_about_unresolved_preprocessor_directive_on_first_line_of_an_ascii_doc_table_cell_should_have_correct_cursor.adoc", tablesTestErrorAboutUnresolvedPreprocessorDirectiveOnFirstLineOfAnAsciiDocTableCellShouldHaveCorrectCursor, nil},

	{"cross reference link in an AsciiDoc table cell should resolve to reference in main document", "asciidoctor/tables_test_cross_reference_link_in_an_ascii_doc_table_cell_should_resolve_to_reference_in_main_document.adoc", tablesTestCrossReferenceLinkInAnAsciiDocTableCellShouldResolveToReferenceInMainDocument, nil},

	{"should discover anchor at start of cell and register it as a reference", "asciidoctor/tables_test_should_discover_anchor_at_start_of_cell_and_register_it_as_a_reference.adoc", tablesTestShouldDiscoverAnchorAtStartOfCellAndRegisterItAsAReference, nil},

	{"should catalog anchor at start of cell in implicit header row when column has a style", "asciidoctor/tables_test_should_catalog_anchor_at_start_of_cell_in_implicit_header_row_when_column_has_a_style.adoc", tablesTestShouldCatalogAnchorAtStartOfCellInImplicitHeaderRowWhenColumnHasAStyle, nil},

	{"should catalog anchor at start of cell in explicit header row when column has a style", "asciidoctor/tables_test_should_catalog_anchor_at_start_of_cell_in_explicit_header_row_when_column_has_a_style.adoc", tablesTestShouldCatalogAnchorAtStartOfCellInExplicitHeaderRowWhenColumnHasAStyle, nil},

	{"should catalog anchor at start of cell in first row", "asciidoctor/tables_test_should_catalog_anchor_at_start_of_cell_in_first_row.adoc", tablesTestShouldCatalogAnchorAtStartOfCellInFirstRow, nil},

	{"footnotes should not be shared between an AsciiDoc table cell and the main document", "asciidoctor/tables_test_footnotes_should_not_be_shared_between_an_ascii_doc_table_cell_and_the_main_document.adoc", tablesTestFootnotesShouldNotBeSharedBetweenAnAsciiDocTableCellAndTheMainDocument, nil},

	{"callout numbers should be globally unique, including AsciiDoc table cells", "asciidoctor/tables_test_callout_numbers_should_be_globally_unique_including_ascii_doc_table_cells.adoc", tablesTestCalloutNumbersShouldBeGloballyUniqueIncludingAsciiDocTableCells, nil},

	{"compat mode can be activated in AsciiDoc table cell", "asciidoctor/tables_test_compat_mode_can_be_activated_in_ascii_doc_table_cell.adoc", tablesTestCompatModeCanBeActivatedInAsciiDocTableCell, nil},

	{"compat mode in AsciiDoc table cell inherits from parent document", "asciidoctor/tables_test_compat_mode_in_ascii_doc_table_cell_inherits_from_parent_document.adoc", tablesTestCompatModeInAsciiDocTableCellInheritsFromParentDocument, nil},

	{"compat mode in AsciiDoc table cell can be unset if set in parent document", "asciidoctor/tables_test_compat_mode_in_ascii_doc_table_cell_can_be_unset_if_set_in_parent_document.adoc", tablesTestCompatModeInAsciiDocTableCellCanBeUnsetIfSetInParentDocument, nil},

	{"nested table", "asciidoctor/tables_test_nested_table.adoc", tablesTestNestedTable, nil},

	{"can set format of nested table to psv", "asciidoctor/tables_test_can_set_format_of_nested_table_to_psv.adoc", tablesTestCanSetFormatOfNestedTableToPsv, nil},

	{"AsciiDoc table cell should inherit to_dir option from parent document", "asciidoctor/tables_test_ascii_doc_table_cell_should_inherit_to_dir_option_from_parent_document.adoc", tablesTestAsciiDocTableCellShouldInheritToDirOptionFromParentDocument, nil},

	{"AsciiDoc table cell should not inherit toc setting from parent document", "asciidoctor/tables_test_ascii_doc_table_cell_should_not_inherit_toc_setting_from_parent_document.adoc", tablesTestAsciiDocTableCellShouldNotInheritTocSettingFromParentDocument, nil},

	{"should be able to enable toc in an AsciiDoc table cell", "asciidoctor/tables_test_should_be_able_to_enable_toc_in_an_ascii_doc_table_cell.adoc", tablesTestShouldBeAbleToEnableTocInAnAsciiDocTableCell, nil},

	{"should be able to enable toc in an AsciiDoc table cell even if hard unset by API", "asciidoctor/tables_test_should_be_able_to_enable_toc_in_an_ascii_doc_table_cell_even_if_hard_unset_by_api.adoc", tablesTestShouldBeAbleToEnableTocInAnAsciiDocTableCellEvenIfHardUnsetByApi, nil},

	{"should be able to enable toc in both outer document and in an AsciiDoc table cell", "asciidoctor/tables_test_should_be_able_to_enable_toc_in_both_outer_document_and_in_an_ascii_doc_table_cell.adoc", tablesTestShouldBeAbleToEnableTocInBothOuterDocumentAndInAnAsciiDocTableCell, nil},

	{"document in an AsciiDoc table cell should not see doctitle of parent", "asciidoctor/tables_test_document_in_an_ascii_doc_table_cell_should_not_see_doctitle_of_parent.adoc", tablesTestDocumentInAnAsciiDocTableCellShouldNotSeeDoctitleOfParent, nil},

	{"cell background color", "asciidoctor/tables_test_cell_background_color.adoc", tablesTestCellBackgroundColor, nil},

	{"should warn if table block is not terminated", "asciidoctor/tables_test_should_warn_if_table_block_is_not_terminated.adoc", tablesTestShouldWarnIfTableBlockIsNotTerminated, nil},

	{"should show correct line number in warning about unterminated block inside AsciiDoc table cell", "asciidoctor/tables_test_should_show_correct_line_number_in_warning_about_unterminated_block_inside_ascii_doc_table_cell.adoc", tablesTestShouldShowCorrectLineNumberInWarningAboutUnterminatedBlockInsideAsciiDocTableCell, nil},

	{"custom separator for an AsciiDoc table cell", "asciidoctor/tables_test_custom_separator_for_an_ascii_doc_table_cell.adoc", tablesTestCustomSeparatorForAnAsciiDocTableCell, nil},

	{"table with breakable option docbook 5", "asciidoctor/tables_test_table_with_breakable_option_docbook_5.adoc", tablesTestTableWithBreakableOptionDocbook5, nil},

	{"table with unbreakable option docbook 5", "asciidoctor/tables_test_table_with_unbreakable_option_docbook_5.adoc", tablesTestTableWithUnbreakableOptionDocbook5, nil},

	{"no implicit header row if cell in first line is quoted and spans multiple lines", "asciidoctor/tables_test_no_implicit_header_row_if_cell_in_first_line_is_quoted_and_spans_multiple_lines.adoc", tablesTestNoImplicitHeaderRowIfCellInFirstLineIsQuotedAndSpansMultipleLines, nil},

	{"converts simple dsv table", "asciidoctor/tables_test_converts_simple_dsv_table.adoc", tablesTestConvertsSimpleDsvTable, nil},

	{"dsv format shorthand", "asciidoctor/tables_test_dsv_format_shorthand.adoc", tablesTestDsvFormatShorthand, nil},

	{"single cell in DSV table should only produce single row", "asciidoctor/tables_test_single_cell_in_dsv_table_should_only_produce_single_row.adoc", tablesTestSingleCellInDsvTableShouldOnlyProduceSingleRow, nil},

	{"should treat trailing colon as an empty cell", "asciidoctor/tables_test_should_treat_trailing_colon_as_an_empty_cell.adoc", tablesTestShouldTreatTrailingColonAsAnEmptyCell, nil},

	{"should treat trailing comma as an empty cell", "asciidoctor/tables_test_should_treat_trailing_comma_as_an_empty_cell.adoc", tablesTestShouldTreatTrailingCommaAsAnEmptyCell, nil},

	{"should log error but not crash if cell data has unclosed quote", "asciidoctor/tables_test_should_log_error_but_not_crash_if_cell_data_has_unclosed_quote.adoc", tablesTestShouldLogErrorButNotCrashIfCellDataHasUnclosedQuote, nil},

	{"should preserve newlines in quoted CSV values", "asciidoctor/tables_test_should_preserve_newlines_in_quoted_csv_values.adoc", tablesTestShouldPreserveNewlinesInQuotedCsvValues, nil},

	{"mixed unquoted records and quoted records with escaped quotes, commas, and wrapped lines", "asciidoctor/tables_test_mixed_unquoted_records_and_quoted_records_with_escaped_quotes_commas_and_wrapped_lines.adoc", tablesTestMixedUnquotedRecordsAndQuotedRecordsWithEscapedQuotesCommasAndWrappedLines, nil},

	{"should allow quotes around a CSV value to be on their own lines", "asciidoctor/tables_test_should_allow_quotes_around_a_csv_value_to_be_on_their_own_lines.adoc", tablesTestShouldAllowQuotesAroundACsvValueToBeOnTheirOwnLines, nil},

	{"csv format shorthand", "asciidoctor/tables_test_csv_format_shorthand.adoc", tablesTestCsvFormatShorthand, nil},

	{"custom csv separator", "asciidoctor/tables_test_custom_csv_separator.adoc", tablesTestCustomCsvSeparator, nil},

	{"single cell in CSV table should only produce single row", "asciidoctor/tables_test_single_cell_in_csv_table_should_only_produce_single_row.adoc", tablesTestSingleCellInCsvTableShouldOnlyProduceSingleRow, nil},

	{"cell formatted with AsciiDoc style", "asciidoctor/tables_test_cell_formatted_with_ascii_doc_style.adoc", tablesTestCellFormattedWithAsciiDocStyle, nil},

	{"should strip whitespace around contents of AsciiDoc cell", "asciidoctor/tables_test_should_strip_whitespace_around_contents_of_ascii_doc_cell.adoc", tablesTestShouldStripWhitespaceAroundContentsOfAsciiDocCell, nil},
}

var tablesTestConvertsSimplePsvTable = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Table{
			AttributeList: nil,
			ColumnCount:   3,
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
						&asciidoc.TableCell{
							Format: &asciidoc.TableCellFormat{
								Multiplier: asciidoc.Optional[int]{
									Value: 1,
									IsSet: false,
								},
								Span: asciidoc.TableCellSpan{
									Column: asciidoc.Optional[int]{
										Value: 1,
										IsSet: false,
									},
									Row: asciidoc.Optional[int]{
										Value: 1,
										IsSet: false,
									},
								},
								HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
									Value: 0,
									IsSet: false,
								},
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
									Value: "C",
								},
							},
							Blank: false,
						},
					},
				},
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "c",
								},
							},
							Blank: false,
						},
					},
				},
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
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

var tablesTestShouldAddDirectionCssClassIfFloatAttributeIsSetOnTable = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Table{
			AttributeList: asciidoc.AttributeList{
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
			ColumnCount: 3,
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
						&asciidoc.TableCell{
							Format: &asciidoc.TableCellFormat{
								Multiplier: asciidoc.Optional[int]{
									Value: 1,
									IsSet: false,
								},
								Span: asciidoc.TableCellSpan{
									Column: asciidoc.Optional[int]{
										Value: 1,
										IsSet: false,
									},
									Row: asciidoc.Optional[int]{
										Value: 1,
										IsSet: false,
									},
								},
								HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
									Value: 0,
									IsSet: false,
								},
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
									Value: "C",
								},
							},
							Blank: false,
						},
					},
				},
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "c",
								},
							},
							Blank: false,
						},
					},
				},
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
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

var tablesTestShouldSetStripesClassIfStripesOptionIsSet = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Table{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.NamedAttribute{
					Name: "stripes",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "odd",
						},
					},
					Quote: 0,
				},
			},
			ColumnCount: 3,
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
						&asciidoc.TableCell{
							Format: &asciidoc.TableCellFormat{
								Multiplier: asciidoc.Optional[int]{
									Value: 1,
									IsSet: false,
								},
								Span: asciidoc.TableCellSpan{
									Column: asciidoc.Optional[int]{
										Value: 1,
										IsSet: false,
									},
									Row: asciidoc.Optional[int]{
										Value: 1,
										IsSet: false,
									},
								},
								HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
									Value: 0,
									IsSet: false,
								},
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
									Value: "C",
								},
							},
							Blank: false,
						},
					},
				},
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "c",
								},
							},
							Blank: false,
						},
					},
				},
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
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

var tablesTestOutputsACaptionOnSimplePsvTable = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Table{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.TitleAttribute{
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Simple psv table",
						},
					},
				},
			},
			ColumnCount: 3,
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
						&asciidoc.TableCell{
							Format: &asciidoc.TableCellFormat{
								Multiplier: asciidoc.Optional[int]{
									Value: 1,
									IsSet: false,
								},
								Span: asciidoc.TableCellSpan{
									Column: asciidoc.Optional[int]{
										Value: 1,
										IsSet: false,
									},
									Row: asciidoc.Optional[int]{
										Value: 1,
										IsSet: false,
									},
								},
								HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
									Value: 0,
									IsSet: false,
								},
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
									Value: "C",
								},
							},
							Blank: false,
						},
					},
				},
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "c",
								},
							},
							Blank: false,
						},
					},
				},
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
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

var tablesTestOnlyIncrementsTableCounterForTablesThatHaveATitle = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Table{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.TitleAttribute{
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "First numbered table",
						},
					},
				},
			},
			ColumnCount: 3,
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
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
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Table{
			AttributeList: nil,
			ColumnCount:   3,
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
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
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Table{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.TitleAttribute{
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Second numbered table",
						},
					},
				},
			},
			ColumnCount: 3,
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
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

var tablesTestUsesExplicitCaptionInFrontOfTitleInPlaceOfDefaultCaptionAndNumber = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Table{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.NamedAttribute{
					Name: "caption",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "All the Data. ",
						},
					},
					Quote: 2,
				},
				&asciidoc.TitleAttribute{
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Simple psv table",
						},
					},
				},
			},
			ColumnCount: 3,
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
						&asciidoc.TableCell{
							Format: &asciidoc.TableCellFormat{
								Multiplier: asciidoc.Optional[int]{
									Value: 1,
									IsSet: false,
								},
								Span: asciidoc.TableCellSpan{
									Column: asciidoc.Optional[int]{
										Value: 1,
										IsSet: false,
									},
									Row: asciidoc.Optional[int]{
										Value: 1,
										IsSet: false,
									},
								},
								HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
									Value: 0,
									IsSet: false,
								},
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
									Value: "C",
								},
							},
							Blank: false,
						},
					},
				},
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "c",
								},
							},
							Blank: false,
						},
					},
				},
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
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

var tablesTestDisablesCaptionWhenCaptionAttributeOnTableIsEmpty = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "[caption=]",
		},
		&asciidoc.NewLine{},
		&asciidoc.Table{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.TitleAttribute{
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Simple psv table",
						},
					},
				},
			},
			ColumnCount: 3,
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
						&asciidoc.TableCell{
							Format: &asciidoc.TableCellFormat{
								Multiplier: asciidoc.Optional[int]{
									Value: 1,
									IsSet: false,
								},
								Span: asciidoc.TableCellSpan{
									Column: asciidoc.Optional[int]{
										Value: 1,
										IsSet: false,
									},
									Row: asciidoc.Optional[int]{
										Value: 1,
										IsSet: false,
									},
								},
								HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
									Value: 0,
									IsSet: false,
								},
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
									Value: "C",
								},
							},
							Blank: false,
						},
					},
				},
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "c",
								},
							},
							Blank: false,
						},
					},
				},
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
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

var tablesTestDisablesCaptionWhenCaptionAttributeOnTableIsEmptyString = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Table{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.NamedAttribute{
					Name:  "caption",
					Val:   nil,
					Quote: 2,
				},
				&asciidoc.TitleAttribute{
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Simple psv table",
						},
					},
				},
			},
			ColumnCount: 3,
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
						&asciidoc.TableCell{
							Format: &asciidoc.TableCellFormat{
								Multiplier: asciidoc.Optional[int]{
									Value: 1,
									IsSet: false,
								},
								Span: asciidoc.TableCellSpan{
									Column: asciidoc.Optional[int]{
										Value: 1,
										IsSet: false,
									},
									Row: asciidoc.Optional[int]{
										Value: 1,
										IsSet: false,
									},
								},
								HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
									Value: 0,
									IsSet: false,
								},
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
									Value: "C",
								},
							},
							Blank: false,
						},
					},
				},
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "c",
								},
							},
							Blank: false,
						},
					},
				},
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
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

var tablesTestDisablesCaptionOnTableWhenTableCaptionDocumentAttributeIsUnset = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeReset{
			Name: "table-caption",
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Table{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.TitleAttribute{
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Simple psv table",
						},
					},
				},
			},
			ColumnCount: 3,
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
						&asciidoc.TableCell{
							Format: &asciidoc.TableCellFormat{
								Multiplier: asciidoc.Optional[int]{
									Value: 1,
									IsSet: false,
								},
								Span: asciidoc.TableCellSpan{
									Column: asciidoc.Optional[int]{
										Value: 1,
										IsSet: false,
									},
									Row: asciidoc.Optional[int]{
										Value: 1,
										IsSet: false,
									},
								},
								HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
									Value: 0,
									IsSet: false,
								},
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
									Value: "C",
								},
							},
							Blank: false,
						},
					},
				},
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "c",
								},
							},
							Blank: false,
						},
					},
				},
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
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

var tablesTestIgnoresEscapedSeparators = &asciidoc.Document{
	Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
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

var tablesTestPreservesEscapedDelimitersAtTheEndOfTheLine = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Table{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: nil,
					ID:    nil,
					Roles: nil,
					Options: []*asciidoc.ShorthandOption{
						&asciidoc.ShorthandOption{
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "header",
								},
							},
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
									Value: "B\\|",
								},
							},
							Blank: false,
						},
					},
				},
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "B1\\|",
								},
							},
							Blank: false,
						},
					},
				},
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
							Elements: asciidoc.Elements{
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

var tablesTestShouldTreatTrailingPipeAsAnEmptyCell = &asciidoc.Document{
	Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{},
							Blank:    false,
						},
					},
				},
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "B2",
								},
							},
							Blank: false,
						},
					},
				},
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
							Elements: asciidoc.Elements{
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

var tablesTestShouldAutoRecoverWithWarningIfMissingLeadingSeparatorOnFirstCell = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
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

var tablesTestPerformsNormalSubstitutionsOnCellContent = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "show_title",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Cool new show",
				},
			},
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
							Elements: asciidoc.Elements{
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

var tablesTestShouldOnlySubstituteSpecialcharsForLiteralTableCells = &asciidoc.Document{
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
									Value: 4,
									IsSet: true,
								},
							},
							Elements: asciidoc.Elements{
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

var tablesTestShouldPreservingLeadingSpacesButNotLeadingNewlinesOrTrailingSpacesInLiteralTableCells = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
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
									Value: 4,
									IsSet: true,
								},
							},
							Elements: asciidoc.Elements{
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
									Value: "normal",
								},
							},
							Blank: false,
						},
					},
				},
				&asciidoc.EmptyLine{
					Text: "\n",
				},
			},
		},
	},
}

var tablesTestShouldIgnoreVTableCellStyle = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
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
			Elements: asciidoc.Elements{
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
		&asciidoc.EmptyLine{
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

var tablesTestTableAndColumnWidthNotAssignedWhenAutowidthOptionIsSpecified = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Table{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.NamedAttribute{
					Name: "options",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "autowidth",
						},
					},
					Quote: 2,
				},
			},
			ColumnCount: 3,
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
						&asciidoc.TableCell{
							Format: &asciidoc.TableCellFormat{
								Multiplier: asciidoc.Optional[int]{
									Value: 1,
									IsSet: false,
								},
								Span: asciidoc.TableCellSpan{
									Column: asciidoc.Optional[int]{
										Value: 1,
										IsSet: false,
									},
									Row: asciidoc.Optional[int]{
										Value: 1,
										IsSet: false,
									},
								},
								HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
									Value: 0,
									IsSet: false,
								},
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
									Value: "C",
								},
							},
							Blank: false,
						},
					},
				},
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "c",
								},
							},
							Blank: false,
						},
					},
				},
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
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

var tablesTestDoesNotAssignColumnWidthForAutowidthColumnsInHtmlOutput = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
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
								Value: 15,
								IsSet: true,
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
						&asciidoc.TableCell{
							Format: &asciidoc.TableCellFormat{
								Multiplier: asciidoc.Optional[int]{
									Value: 1,
									IsSet: false,
								},
								Span: asciidoc.TableCellSpan{
									Column: asciidoc.Optional[int]{
										Value: 1,
										IsSet: false,
									},
									Row: asciidoc.Optional[int]{
										Value: 1,
										IsSet: false,
									},
								},
								HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
									Value: 0,
									IsSet: false,
								},
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "D",
								},
							},
							Blank: false,
						},
					},
				},
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "d",
								},
							},
							Blank: false,
						},
					},
				},
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
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

var tablesTestCanAssignAutowidthToAllColumnsEvenWhenTableHasAWidth = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
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
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "50%",
						},
					},
					Quote: 0,
				},
			},
			ColumnCount: 4,
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
						&asciidoc.TableCell{
							Format: &asciidoc.TableCellFormat{
								Multiplier: asciidoc.Optional[int]{
									Value: 1,
									IsSet: false,
								},
								Span: asciidoc.TableCellSpan{
									Column: asciidoc.Optional[int]{
										Value: 1,
										IsSet: false,
									},
									Row: asciidoc.Optional[int]{
										Value: 1,
										IsSet: false,
									},
								},
								HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
									Value: 0,
									IsSet: false,
								},
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "D",
								},
							},
							Blank: false,
						},
					},
				},
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "d",
								},
							},
							Blank: false,
						},
					},
				},
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
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

var tablesTestEquallyDistributesRemainingColumnWidthToAutowidthColumnsInDocBookOutput = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
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
								Value: 15,
								IsSet: true,
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
						&asciidoc.TableCell{
							Format: &asciidoc.TableCellFormat{
								Multiplier: asciidoc.Optional[int]{
									Value: 1,
									IsSet: false,
								},
								Span: asciidoc.TableCellSpan{
									Column: asciidoc.Optional[int]{
										Value: 1,
										IsSet: false,
									},
									Row: asciidoc.Optional[int]{
										Value: 1,
										IsSet: false,
									},
								},
								HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
									Value: 0,
									IsSet: false,
								},
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "D",
								},
							},
							Blank: false,
						},
					},
				},
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "d",
								},
							},
							Blank: false,
						},
					},
				},
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
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

var tablesTestShouldComputeColumnWidthsBasedOnPagewidthWhenWidthIsSetOnTableInDocBookOutput = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "pagewidth",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "500",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Table{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.NamedAttribute{
					Name: "width",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "50%",
						},
					},
					Quote: 0,
				},
			},
			ColumnCount: 4,
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
						&asciidoc.TableCell{
							Format: &asciidoc.TableCellFormat{
								Multiplier: asciidoc.Optional[int]{
									Value: 1,
									IsSet: false,
								},
								Span: asciidoc.TableCellSpan{
									Column: asciidoc.Optional[int]{
										Value: 1,
										IsSet: false,
									},
									Row: asciidoc.Optional[int]{
										Value: 1,
										IsSet: false,
									},
								},
								HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
									Value: 0,
									IsSet: false,
								},
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "D",
								},
							},
							Blank: false,
						},
					},
				},
				&asciidoc.EmptyLine{
					Text: "\n",
				},
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "d",
								},
							},
							Blank: false,
						},
					},
				},
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
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

var tablesTestExplicitTableWidthIsUsedEvenWhenAutowidthOptionIsSpecified = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Table{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: nil,
					ID:    nil,
					Roles: nil,
					Options: []*asciidoc.ShorthandOption{
						&asciidoc.ShorthandOption{
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "autowidth",
								},
							},
						},
					},
				},
				&asciidoc.NamedAttribute{
					Name: "width",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "75%",
						},
					},
					Quote: 0,
				},
			},
			ColumnCount: 3,
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
						&asciidoc.TableCell{
							Format: &asciidoc.TableCellFormat{
								Multiplier: asciidoc.Optional[int]{
									Value: 1,
									IsSet: false,
								},
								Span: asciidoc.TableCellSpan{
									Column: asciidoc.Optional[int]{
										Value: 1,
										IsSet: false,
									},
									Row: asciidoc.Optional[int]{
										Value: 1,
										IsSet: false,
									},
								},
								HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
									Value: 0,
									IsSet: false,
								},
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
									Value: "C",
								},
							},
							Blank: false,
						},
					},
				},
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "c",
								},
							},
							Blank: false,
						},
					},
				},
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
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

var tablesTestFirstRowSetsNumberOfColumnsWhenNotSpecified = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Table{
			AttributeList: nil,
			ColumnCount:   4,
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "fourth",
								},
							},
							Blank: false,
						},
					},
				},
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
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

var tablesTestColspecAttributeUsingAsteriskSyntaxSetsNumberOfColumns = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
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
						&asciidoc.TableCell{
							Format: &asciidoc.TableCellFormat{
								Multiplier: asciidoc.Optional[int]{
									Value: 1,
									IsSet: false,
								},
								Span: asciidoc.TableCellSpan{
									Column: asciidoc.Optional[int]{
										Value: 1,
										IsSet: false,
									},
									Row: asciidoc.Optional[int]{
										Value: 1,
										IsSet: false,
									},
								},
								HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
									Value: 0,
									IsSet: false,
								},
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
									Value: "C",
								},
							},
							Blank: false,
						},
					},
				},
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "c",
								},
							},
							Blank: false,
						},
					},
				},
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
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

var tablesTestTableWithExplicitColumnCountCanHaveMultipleRowsOnASingleLine = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "1",
								},
							},
							Blank: false,
						},
					},
				},
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
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

var tablesTestTableWithExplicitDeprecatedColspecSyntaxCanHaveMultipleRowsOnASingleLine = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "1",
								},
							},
							Blank: false,
						},
					},
				},
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
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

var tablesTestColumnsAreAddedForEmptyRecordsInColspecAttribute = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
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
			ColumnCount: 2,
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "two",
								},
							},
							Blank: false,
						},
					},
				},
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "2",
								},
							},
							Blank: false,
						},
					},
				},
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
							Elements: asciidoc.Elements{
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

var tablesTestColsMayBeSeparatedBySemiColonInsteadOfComma = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
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
							Elements: asciidoc.Elements{
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

var tablesTestColsAttributeMayIncludeSpaces = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "two",
								},
							},
							Blank: false,
						},
					},
				},
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "2",
								},
							},
							Blank: false,
						},
					},
				},
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
							Elements: asciidoc.Elements{
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

var tablesTestBlankColsAttributeShouldBeIgnored = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Table{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.TableColumnsAttribute{
					Columns: nil,
				},
			},
			ColumnCount: 2,
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "two",
								},
							},
							Blank: false,
						},
					},
				},
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "2",
								},
							},
							Blank: false,
						},
					},
				},
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
							Elements: asciidoc.Elements{
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

var tablesTestEmptyColsAttributeShouldBeIgnored = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Table{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.TableColumnsAttribute{
					Columns: nil,
				},
			},
			ColumnCount: 2,
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "two",
								},
							},
							Blank: false,
						},
					},
				},
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "2",
								},
							},
							Blank: false,
						},
					},
				},
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
							Elements: asciidoc.Elements{
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

var tablesTestTableWithHeaderAndFooter = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Table{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.NamedAttribute{
					Name: "options",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "header,footer",
						},
					},
					Quote: 2,
				},
			},
			ColumnCount: 2,
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "Quantity",
								},
							},
							Blank: false,
						},
					},
				},
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "1",
								},
							},
							Blank: false,
						},
					},
				},
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "2",
								},
							},
							Blank: false,
						},
					},
				},
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "3",
								},
							},
							Blank: false,
						},
					},
				},
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
							Elements: asciidoc.Elements{
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

var tablesTestTableWithHeaderAndFooterDocbook = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Table{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.TitleAttribute{
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Table with header, body and footer",
						},
					},
				},
				&asciidoc.NamedAttribute{
					Name: "options",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "header,footer",
						},
					},
					Quote: 2,
				},
			},
			ColumnCount: 2,
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "Quantity",
								},
							},
							Blank: false,
						},
					},
				},
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "1",
								},
							},
							Blank: false,
						},
					},
				},
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "2",
								},
							},
							Blank: false,
						},
					},
				},
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "3",
								},
							},
							Blank: false,
						},
					},
				},
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
							Elements: asciidoc.Elements{
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

var tablesTestShouldSetHorizontalAndVerticalAlignmentWhenConvertingToDocBook = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Table{
			AttributeList: nil,
			ColumnCount:   3,
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "C",
								},
							},
							Blank: false,
						},
					},
				},
				&asciidoc.EmptyLine{
					Text: "\n",
				},
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
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

var tablesTestShouldPreserveFrameValueEndsWhenConvertingToHtml = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Table{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.NamedAttribute{
					Name: "frame",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "ends",
						},
					},
					Quote: 0,
				},
			},
			ColumnCount: 3,
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
						&asciidoc.TableCell{
							Format: &asciidoc.TableCellFormat{
								Multiplier: asciidoc.Optional[int]{
									Value: 1,
									IsSet: false,
								},
								Span: asciidoc.TableCellSpan{
									Column: asciidoc.Optional[int]{
										Value: 1,
										IsSet: false,
									},
									Row: asciidoc.Optional[int]{
										Value: 1,
										IsSet: false,
									},
								},
								HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
									Value: 0,
									IsSet: false,
								},
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

var tablesTestShouldNormalizeFrameValueTopbotAsEndsWhenConvertingToHtml = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Table{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.NamedAttribute{
					Name: "frame",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "topbot",
						},
					},
					Quote: 0,
				},
			},
			ColumnCount: 3,
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
						&asciidoc.TableCell{
							Format: &asciidoc.TableCellFormat{
								Multiplier: asciidoc.Optional[int]{
									Value: 1,
									IsSet: false,
								},
								Span: asciidoc.TableCellSpan{
									Column: asciidoc.Optional[int]{
										Value: 1,
										IsSet: false,
									},
									Row: asciidoc.Optional[int]{
										Value: 1,
										IsSet: false,
									},
								},
								HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
									Value: 0,
									IsSet: false,
								},
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

var tablesTestShouldPreserveFrameValueTopbotWhenConvertingToDocBook = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Table{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.NamedAttribute{
					Name: "frame",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "topbot",
						},
					},
					Quote: 0,
				},
			},
			ColumnCount: 3,
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
						&asciidoc.TableCell{
							Format: &asciidoc.TableCellFormat{
								Multiplier: asciidoc.Optional[int]{
									Value: 1,
									IsSet: false,
								},
								Span: asciidoc.TableCellSpan{
									Column: asciidoc.Optional[int]{
										Value: 1,
										IsSet: false,
									},
									Row: asciidoc.Optional[int]{
										Value: 1,
										IsSet: false,
									},
								},
								HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
									Value: 0,
									IsSet: false,
								},
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

var tablesTestShouldConvertFrameValueEndsToTopbotWhenConvertingToDocBook = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Table{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.NamedAttribute{
					Name: "frame",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "ends",
						},
					},
					Quote: 0,
				},
			},
			ColumnCount: 3,
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
						&asciidoc.TableCell{
							Format: &asciidoc.TableCellFormat{
								Multiplier: asciidoc.Optional[int]{
									Value: 1,
									IsSet: false,
								},
								Span: asciidoc.TableCellSpan{
									Column: asciidoc.Optional[int]{
										Value: 1,
										IsSet: false,
									},
									Row: asciidoc.Optional[int]{
										Value: 1,
										IsSet: false,
									},
								},
								HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
									Value: 0,
									IsSet: false,
								},
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

var tablesTestTableWithImplicitHeaderRow = &asciidoc.Document{
	Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "Column 2",
								},
							},
							Blank: false,
						},
					},
				},
				&asciidoc.EmptyLine{
					Text: "\n",
				},
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "Data B1",
								},
							},
							Blank: false,
						},
					},
				},
				&asciidoc.EmptyLine{
					Text: "\n",
				},
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
							Elements: asciidoc.Elements{
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

var tablesTestTableWithImplicitHeaderRowOnly = &asciidoc.Document{
	Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "Column 2",
								},
							},
							Blank: false,
						},
					},
				},
				&asciidoc.EmptyLine{
					Text: "\n",
				},
			},
		},
	},
}

var tablesTestTableWithImplicitHeaderRowWhenOtherOptionsSet = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Table{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: nil,
					ID:    nil,
					Roles: nil,
					Options: []*asciidoc.ShorthandOption{
						&asciidoc.ShorthandOption{
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "autowidth",
								},
							},
						},
					},
				},
			},
			ColumnCount: 2,
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "Column 2",
								},
							},
							Blank: false,
						},
					},
				},
				&asciidoc.EmptyLine{
					Text: "\n",
				},
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
							Elements: asciidoc.Elements{
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

var tablesTestNoImplicitHeaderRowIfSecondLineNotBlank = &asciidoc.Document{
	Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "Column 2",
								},
							},
							Blank: false,
						},
					},
				},
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "Data B1",
								},
							},
							Blank: false,
						},
					},
				},
				&asciidoc.EmptyLine{
					Text: "\n",
				},
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
							Elements: asciidoc.Elements{
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

var tablesTestNoImplicitHeaderRowIfCellInFirstLineSpansMultipleLines = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "B1",
								},
							},
							Blank: false,
						},
					},
				},
				&asciidoc.EmptyLine{
					Text: "\n",
				},
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
							Elements: asciidoc.Elements{
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

var tablesTestShouldFormatFirstCellAsLiteralIfThereIsNoImplicitHeaderRowAndColumnHasLStyle = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
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
							Elements: asciidoc.Elements{
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

var tablesTestShouldFormatFirstCellAsAsciiDocIfThereIsNoImplicitHeaderRowAndColumnHasAStyle = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
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
								&asciidoc.UnorderedListItem{
									Elements: asciidoc.Elements{
										&asciidoc.String{
											Value: "list",
										},
									},
									AttributeList: nil,
									Indent:        " ",
									Marker:        "*",
									Checklist:     0,
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

var tablesTestShouldInterpretLeadingIndentIfFirstCellIsAsciiDocAndThereIsNoImplicitHeaderRow = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
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
								&asciidoc.NewLine{},
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
							Elements: asciidoc.Elements{
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

var tablesTestShouldFormatFirstCellAsAsciiDocIfThereIsNoImplicitHeaderRowAndCellHasAStyle = &asciidoc.Document{
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
								&asciidoc.UnorderedListItem{
									Elements: asciidoc.Elements{
										&asciidoc.String{
											Value: "list",
										},
									},
									AttributeList: nil,
									Indent:        " ",
									Marker:        "*",
									Checklist:     0,
								},
							},
							Blank: false,
						},
					},
				},
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

var tablesTestNoImplicitHeaderRowIfAsciiDocCellInFirstLineSpansMultipleLines = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
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
									Value: "contains AsciiDoc content",
								},
								&asciidoc.NewLine{},
								&asciidoc.EmptyLine{
									Text: "",
								},
								&asciidoc.UnorderedListItem{
									Elements: asciidoc.Elements{
										&asciidoc.String{
											Value: "a",
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
											Value: "b",
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
											Value: "c",
										},
									},
									AttributeList: nil,
									Indent:        "",
									Marker:        "*",
									Checklist:     0,
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "contains no AsciiDoc content",
								},
								&asciidoc.NewLine{},
								&asciidoc.EmptyLine{
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
							Elements: asciidoc.Elements{
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

var tablesTestNoImplicitHeaderRowIfFirstLineBlank = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Table{
			AttributeList: nil,
			ColumnCount:   2,
			Elements: asciidoc.Elements{
				&asciidoc.EmptyLine{
					Text: "\n",
				},
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "Column 2",
								},
							},
							Blank: false,
						},
					},
				},
				&asciidoc.EmptyLine{
					Text: "\n",
				},
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "Data B1",
								},
							},
							Blank: false,
						},
					},
				},
				&asciidoc.EmptyLine{
					Text: "\n",
				},
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "Data B2",
								},
							},
							Blank: false,
						},
					},
				},
				&asciidoc.EmptyLine{
					Text: "\n",
				},
			},
		},
	},
}

var tablesTestNoImplicitHeaderRowIfNoheaderOptionIsSpecified = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Table{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: nil,
					ID:    nil,
					Roles: nil,
					Options: []*asciidoc.ShorthandOption{
						&asciidoc.ShorthandOption{
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "noheader",
								},
							},
						},
					},
				},
			},
			ColumnCount: 2,
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "Column 2",
								},
							},
							Blank: false,
						},
					},
				},
				&asciidoc.EmptyLine{
					Text: "\n",
				},
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "Data B1",
								},
							},
							Blank: false,
						},
					},
				},
				&asciidoc.EmptyLine{
					Text: "\n",
				},
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
							Elements: asciidoc.Elements{
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

var tablesTestStylesNotAppliedToHeaderCells = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
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
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "header,footer",
						},
					},
					Quote: 2,
				},
			},
			ColumnCount: 3,
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "Website",
								},
							},
							Blank: false,
						},
					},
				},
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
								&asciidoc.Link{
									AttributeList: nil,
									URL: asciidoc.URL{
										Scheme: "https://",
										Path: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
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

var tablesTestShouldApplyTextFormattingToCellsInImplicitHeaderRowWhenColumnHasAStyle = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
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
									Value: " ",
								},
								&asciidoc.Italic{
									AttributeList: nil,
									Elements: asciidoc.Elements{
										&asciidoc.String{
											Value: "foo",
										},
									},
								},
								&asciidoc.String{
									Value: " ",
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
								&asciidoc.Bold{
									AttributeList: nil,
									Elements: asciidoc.Elements{
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
				&asciidoc.EmptyLine{
					Text: "\n",
				},
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
								&asciidoc.UnorderedListItem{
									Elements: asciidoc.Elements{
										&asciidoc.String{
											Value: "list item",
										},
									},
									AttributeList: nil,
									Indent:        " ",
									Marker:        "*",
									Checklist:     0,
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

var tablesTestShouldApplyStyleAndTextFormattingToCellsInFirstRowIfNoImplicitHeader = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
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
								&asciidoc.Italic{
									AttributeList: nil,
									Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
								&asciidoc.Bold{
									AttributeList: nil,
									Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
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

var tablesTestVerticalTableHeadersUseThElementInsteadOfHeaderClass = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
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
			Elements: asciidoc.Elements{
				&asciidoc.EmptyLine{
					Text: "\n",
				},
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "Website",
								},
							},
							Blank: false,
						},
					},
				},
				&asciidoc.EmptyLine{
					Text: "\n",
				},
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
								&asciidoc.Link{
									AttributeList: nil,
									URL: asciidoc.URL{
										Scheme: "https://",
										Path: asciidoc.Elements{
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
				&asciidoc.EmptyLine{
					Text: "\n",
				},
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "Website",
								},
							},
							Blank: false,
						},
					},
				},
				&asciidoc.EmptyLine{
					Text: "\n",
				},
			},
		},
	},
}

var tablesTestSupportsHorizontalAndVerticalSourceDataWithBlankLinesAndTableHeader = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Table{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.TitleAttribute{
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Horizontal and vertical source data",
						},
					},
				},
				&asciidoc.NamedAttribute{
					Name: "width",
					Val: asciidoc.Elements{
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
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "header",
						},
					},
					Quote: 2,
				},
			},
			ColumnCount: 4,
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
									Value: "Date",
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
									Value: "Duration",
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
									Value: "Avg HR",
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
									Value: "Notes",
								},
							},
							Blank: false,
						},
					},
				},
				&asciidoc.EmptyLine{
					Text: "\n",
				},
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
									Value: "22-Aug-08",
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
									Value: "10:24",
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
									Value: "157",
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
								&asciidoc.NewLine{},
								&asciidoc.String{
									Value: "Worked out MSHR (max sustainable heart rate) by going hard",
								},
								&asciidoc.NewLine{},
								&asciidoc.String{
									Value: "for this interval.",
								},
							},
							Blank: false,
						},
					},
				},
				&asciidoc.EmptyLine{
					Text: "\n",
				},
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
									Value: "22-Aug-08",
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
									Value: "23:03",
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
									Value: "152",
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
								&asciidoc.NewLine{},
								&asciidoc.String{
									Value: "Back-to-back with previous interval.",
								},
							},
							Blank: false,
						},
					},
				},
				&asciidoc.EmptyLine{
					Text: "\n",
				},
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
									Value: "24-Aug-08",
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
									Value: "40:00",
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
									Value: "145",
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
								&asciidoc.NewLine{},
								&asciidoc.String{
									Value: "Moderately hard interspersed with 3x 3min intervals (2 min",
								},
								&asciidoc.NewLine{},
								&asciidoc.String{
									Value: "hard + 1 min really hard taking the HR up to 160).",
								},
								&asciidoc.NewLine{},
								&asciidoc.NewLine{},
								&asciidoc.String{
									Value: "I am getting in shape!",
								},
							},
							Blank: false,
						},
					},
				},
				&asciidoc.EmptyLine{
					Text: "\n",
				},
			},
		},
	},
}

var tablesTestPercentagesAsColumnWidths = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
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
								Value: 10,
								IsSet: true,
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
								Value: 90,
								IsSet: true,
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
							Elements: asciidoc.Elements{
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

var tablesTestSpansAlignmentsAndStyles = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
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
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "25%",
						},
					},
					Quote: 2,
				},
			},
			ColumnCount: 4,
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "4",
								},
							},
							Blank: false,
						},
					},
				},
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "6",
								},
							},
							Blank: false,
						},
						&asciidoc.TableCell{
							Format:   nil,
							Elements: nil,
							Blank:    true,
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "7",
								},
							},
							Blank: false,
						},
					},
				},
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "8",
								},
							},
							Blank: false,
						},
						&asciidoc.TableCell{
							Format:   nil,
							Elements: nil,
							Blank:    true,
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "9",
								},
							},
							Blank: false,
						},
						&asciidoc.TableCell{
							Format:   nil,
							Elements: nil,
							Blank:    true,
						},
					},
				},
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "10",
								},
							},
							Blank: false,
						},
						&asciidoc.TableCell{
							Format:   nil,
							Elements: nil,
							Blank:    true,
						},
						&asciidoc.TableCell{
							Format:   nil,
							Elements: nil,
							Blank:    true,
						},
						&asciidoc.TableCell{
							Format:   nil,
							Elements: nil,
							Blank:    true,
						},
					},
				},
			},
		},
	},
}

var tablesTestSetsUpColumnsCorrectlyIfFirstRowHasCellThatSpansColumns = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Table{
			AttributeList: nil,
			ColumnCount:   3,
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "AAA",
								},
							},
							Blank: false,
						},
						&asciidoc.TableCell{
							Format:   nil,
							Elements: nil,
							Blank:    true,
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
									Value: "CCC",
								},
							},
							Blank: false,
						},
					},
				},
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "CCC",
								},
							},
							Blank: false,
						},
					},
				},
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
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

var tablesTestSupportsRepeatingCells = &asciidoc.Document{
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "A",
								},
							},
							Blank: false,
						},
					},
				},
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
									Value: "1",
								},
							},
							Blank: false,
						},
					},
				},
				&asciidoc.TableRow{
					Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "2",
								},
							},
							Blank: false,
						},
					},
				},
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
									Value: "b",
								},
							},
							Blank: false,
						},
					},
				},
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

var tablesTestCalculatesColnamesCorrectlyWhenUsingImplicitColumnCountAndSingleCellWithColspan = &asciidoc.Document{
	Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "Two Columns",
								},
							},
							Blank: false,
						},
						&asciidoc.TableCell{
							Format:   nil,
							Elements: nil,
							Blank:    true,
						},
					},
				},
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
							Elements: asciidoc.Elements{
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

var tablesTestCalculatesColnamesCorrectlyWhenUsingImplicitColumnCountAndCellsWithMixedColspans = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Table{
			AttributeList: nil,
			ColumnCount:   3,
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "Two Columns",
								},
							},
							Blank: false,
						},
						&asciidoc.TableCell{
							Format:   nil,
							Elements: nil,
							Blank:    true,
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
									Value: "One Column",
								},
							},
							Blank: false,
						},
					},
				},
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
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

var tablesTestAssignsUniqueColumnNamesForTableWithImplicitColumnCountAndColspansInFirstRow = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Table{
			AttributeList: nil,
			ColumnCount:   5,
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
							Elements: asciidoc.Elements{},
							Blank:    false,
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "Node 0",
								},
							},
							Blank: false,
						},
						&asciidoc.TableCell{
							Format:   nil,
							Elements: nil,
							Blank:    true,
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "Node 1",
								},
							},
							Blank: false,
						},
						&asciidoc.TableCell{
							Format:   nil,
							Elements: nil,
							Blank:    true,
						},
					},
				},
				&asciidoc.EmptyLine{
					Text: "\n",
				},
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "Core 5",
								},
							},
							Blank: false,
						},
					},
				},
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
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

var tablesTestShouldDropRowButPreserveRemainingRowsAfterCellWithColspanExceedsNumberOfColumns = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "A",
								},
							},
							Blank: false,
						},
						&asciidoc.TableCell{
							Format:   nil,
							Elements: nil,
							Blank:    true,
						},
						&asciidoc.TableCell{
							Format:   nil,
							Elements: nil,
							Blank:    true,
						},
					},
				},
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "C",
								},
								&asciidoc.NewLine{},
								&asciidoc.EmptyLine{
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

var tablesTestShouldDropLastRowIfLastCellInTableHasColspanThatExceedsSpecifiedNumberOfColumns = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
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
							Elements: asciidoc.Elements{
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

var tablesTestShouldDropLastRowIfLastCellInTableHasColspanThatExceedsImplicitNumberOfColumns = &asciidoc.Document{
	Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "b",
								},
							},
							Blank: false,
						},
					},
				},
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "d",
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

var tablesTestShouldTakeColspanIntoAccountWhenTakingCellsForRow = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
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
			ColumnCount: 7,
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "a",
								},
							},
							Blank: false,
						},
						&asciidoc.TableCell{
							Format:   nil,
							Elements: nil,
							Blank:    true,
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "b",
								},
							},
							Blank: false,
						},
						&asciidoc.TableCell{
							Format:   nil,
							Elements: nil,
							Blank:    true,
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "c",
								},
							},
							Blank: false,
						},
						&asciidoc.TableCell{
							Format:   nil,
							Elements: nil,
							Blank:    true,
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "d",
								},
							},
							Blank: false,
						},
					},
				},
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
									Value: "e",
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
									Value: "f",
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
									Value: "g",
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
									Value: "h",
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
									Value: "i",
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
									Value: "j",
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
									Value: "k",
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

var tablesTestShouldDropIncompleteRowAtEndOfTableAndLogAnError = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "b",
								},
							},
							Blank: false,
						},
					},
				},
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "d",
								},
							},
							Blank: false,
						},
					},
				},
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
									Value: "e",
								},
							},
							Blank: false,
						},
						&asciidoc.TableCell{
							Format:   nil,
							Elements: nil,
							Blank:    true,
						},
					},
				},
			},
		},
	},
}

var tablesTestShouldApplyCellStyleForColumnToRepeatedContent = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
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
			ColumnCount: 2,
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
									Value: "Paragraphs",
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
									Value: "Literal",
								},
							},
							Blank: false,
						},
					},
				},
				&asciidoc.EmptyLine{
					Text: "\n",
				},
				&asciidoc.TableRow{
					Elements: asciidoc.Elements{
						&asciidoc.TableCell{
							Format: &asciidoc.TableCellFormat{
								Multiplier: asciidoc.Optional[int]{
									Value: 2,
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
							Elements: asciidoc.Elements{
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
						&asciidoc.TableCell{
							Format:   nil,
							Elements: nil,
							Blank:    true,
						},
					},
				},
			},
		},
	},
}

var tablesTestShouldNotSplitParagraphAtLineContainingOnlyblankThatIsDirectlyAdjacentToNonBlankLines = &asciidoc.Document{
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
									Value: 0,
									IsSet: false,
								},
							},
							Elements: asciidoc.Elements{
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

var tablesTestShouldStripTrailingNewlinesWhenSplittingParagraphs = &asciidoc.Document{
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
									Value: 0,
									IsSet: false,
								},
							},
							Elements: asciidoc.Elements{
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

var tablesTestBasicAsciiDocCell = &asciidoc.Document{
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
								&asciidoc.OpenBlock{
									AttributeList: nil,
									Delimiter: asciidoc.Delimiter{
										Type:   7,
										Length: 2,
									},
									Elements: asciidoc.Elements{
										&asciidoc.Admonition{
											AdmonitionType: 1,
											AttributeList:  nil,
										},
										&asciidoc.String{
											Value: "content",
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

var tablesTestAsciiDocTableCellShouldBeWrappedInDivWithClassContent = &asciidoc.Document{
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

var tablesTestDoctypeCanBeSetInAsciiDocTableCell = &asciidoc.Document{
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
								&asciidoc.NewLine{},
								&asciidoc.AttributeEntry{
									Name: "doctype",
									Elements: asciidoc.Elements{
										&asciidoc.String{
											Value: "inline",
										},
									},
								},
								&asciidoc.EmptyLine{
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

var tablesTestShouldResetDoctypeToDefaultInAsciiDocTableCell = &asciidoc.Document{
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
												&asciidoc.NewLine{},
												&asciidoc.Section{
													AttributeList: nil,
													Elements:      nil,
													Title: asciidoc.Elements{
														&asciidoc.String{
															Value: "AsciiDoc Table Cell",
														},
													},
													Level: 0,
												},
												&asciidoc.EmptyLine{
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

var tablesTestShouldUpdateDoctypeRelatedAttributesInAsciiDocTableCellWhenDoctypeIsSet = &asciidoc.Document{
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
							Value: "article",
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
												&asciidoc.Section{
													AttributeList: nil,
													Elements:      nil,
													Title: asciidoc.Elements{
														&asciidoc.String{
															Value: "AsciiDoc Table Cell",
														},
													},
													Level: 0,
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
					Value: "Document Title",
				},
			},
			Level: 0,
		},
	},
}

var tablesTestShouldNotAllowAsciiDocTableCellToSetADocumentAttributeThatWasHardSetByTheApi = &asciidoc.Document{
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
								&asciidoc.NewLine{},
								&asciidoc.AttributeEntry{
									Name:     "icons",
									Elements: nil,
								},
								&asciidoc.EmptyLine{
									Text: "",
								},
								&asciidoc.Paragraph{
									AttributeList: nil,
									Elements: asciidoc.Elements{
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

var tablesTestShouldNotAllowAsciiDocTableCellToSetADocumentAttributeThatWasHardUnsetByTheApi = &asciidoc.Document{
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
								&asciidoc.NewLine{},
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
									AttributeList: nil,
									Elements: asciidoc.Elements{
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

var tablesTestShouldKeepAttributeUnsetInAsciiDocTableCellIfUnsetInParentDocument = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeReset{
			Name: "sectids",
		},
		&asciidoc.AttributeReset{
			Name: "table-caption",
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
				&asciidoc.Table{
					AttributeList: asciidoc.AttributeList{
						&asciidoc.TitleAttribute{
							Val: asciidoc.Elements{
								&asciidoc.String{
									Value: "Outer Table",
								},
							},
						},
					},
					ColumnCount: 1,
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
										&asciidoc.EmptyLine{
											Text: "",
										},
										&asciidoc.Section{
											AttributeList: nil,
											Elements:      nil,
											Title: asciidoc.Elements{
												&asciidoc.String{
													Value: "Inner Heading",
												},
											},
											Level: 1,
										},
										&asciidoc.EmptyLine{
											Text: "",
										},
										&asciidoc.Paragraph{
											AttributeList: asciidoc.AttributeList{
												&asciidoc.TitleAttribute{
													Val: asciidoc.Elements{
														&asciidoc.String{
															Value: "Inner Table",
														},
													},
												},
											},
											Elements: asciidoc.Elements{
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
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Outer Heading",
				},
			},
			Level: 1,
		},
	},
}

var tablesTestShouldAllowAttributeUnsetInParentDocumentToBeSetInAsciiDocTableCell = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeReset{
			Name: "sectids",
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
										&asciidoc.EmptyLine{
											Text: "",
										},
										&asciidoc.Section{
											AttributeList: nil,
											Elements:      nil,
											Title: asciidoc.Elements{
												&asciidoc.String{
													Value: "No ID",
												},
											},
											Level: 1,
										},
										&asciidoc.EmptyLine{
											Text: "",
										},
										&asciidoc.AttributeEntry{
											Name:     "sectids",
											Elements: nil,
										},
										&asciidoc.EmptyLine{
											Text: "",
										},
										&asciidoc.Section{
											AttributeList: nil,
											Elements:      nil,
											Title: asciidoc.Elements{
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
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "No ID",
				},
			},
			Level: 1,
		},
	},
}

var tablesTestShouldNotAllowLockedAttributeUnsetInParentDocumentToBeSetInAsciiDocTableCell = &asciidoc.Document{
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
										&asciidoc.NewLine{},
										&asciidoc.EmptyLine{
											Text: "",
										},
										&asciidoc.Section{
											AttributeList: nil,
											Elements:      nil,
											Title: asciidoc.Elements{
												&asciidoc.String{
													Value: "No ID",
												},
											},
											Level: 1,
										},
										&asciidoc.EmptyLine{
											Text: "",
										},
										&asciidoc.AttributeEntry{
											Name:     "sectids",
											Elements: nil,
										},
										&asciidoc.EmptyLine{
											Text: "",
										},
										&asciidoc.Section{
											AttributeList: nil,
											Elements:      nil,
											Title: asciidoc.Elements{
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
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "No ID",
				},
			},
			Level: 1,
		},
	},
}

var tablesTestAsciiDocContent = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "Description",
								},
							},
							Blank: false,
						},
					},
				},
				&asciidoc.EmptyLine{
					Text: "\n",
				},
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
								&asciidoc.NewLine{},
								&asciidoc.String{
									Value: "Link badges ('XHTML 1.1' and 'CSS') in document footers.",
								},
								&asciidoc.NewLine{},
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
											Style: &asciidoc.ShorthandStyle{
												Elements: asciidoc.Elements{
													&asciidoc.String{
														Value: "NOTE",
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
								&asciidoc.Anchor{
									ID:       "X97",
									Elements: nil,
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
								&asciidoc.NewLine{},
								&asciidoc.String{
									Value: "These three attributes control which document information",
								},
								&asciidoc.NewLine{},
								&asciidoc.String{
									Value: "files will be included in the the header of the output file:",
								},
								&asciidoc.NewLine{},
								&asciidoc.EmptyLine{
									Text: "",
								},
								&asciidoc.String{
									Value: "docinfo:: Include ",
								},
								&asciidoc.Monospace{
									AttributeList: nil,
									Elements: asciidoc.Elements{
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
									Elements: asciidoc.Elements{
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
									Elements: asciidoc.Elements{
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
									Elements: asciidoc.Elements{
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
								&asciidoc.EmptyLine{
									Text: "",
								},
								&asciidoc.String{
									Value: "Where ",
								},
								&asciidoc.Monospace{
									AttributeList: nil,
									Elements: asciidoc.Elements{
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
									Elements: asciidoc.Elements{
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
									Elements: asciidoc.Elements{
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
									Elements: asciidoc.Elements{
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

var tablesTestShouldPreserveLeadingIndentationInContentsOfAsciiDocTableCellIfContentsStartsWithNewline = &asciidoc.Document{
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
								&asciidoc.NewLine{},
								&asciidoc.String{
									Value: "$ command",
								},
							},
							Blank: false,
						},
					},
				},
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
									Value: " paragraph",
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

var tablesTestPreprocessorDirectiveOnFirstLineOfAnAsciiDocTableCellShouldBeProcessed = &asciidoc.Document{
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
								&asciidoc.FileInclude{
									AttributeList: nil,
									Elements: asciidoc.Elements{
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

var tablesTestErrorAboutUnresolvedPreprocessorDirectiveOnFirstLineOfAnAsciiDocTableCellShouldHaveCorrectCursor = &asciidoc.Document{
	Elements: asciidoc.Elements{
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
				&asciidoc.EmptyLine{
					Text: "\n",
				},
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
									Value: "text",
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
							Elements: asciidoc.Elements{
								&asciidoc.FileInclude{
									AttributeList: nil,
									Elements: asciidoc.Elements{
										&asciidoc.String{
											Value: "does-not-exist.adoc",
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

var tablesTestCrossReferenceLinkInAnAsciiDocTableCellShouldResolveToReferenceInMainDocument = &asciidoc.Document{
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
											ID:            "_more",
											Format:        0,
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
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Some",
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
					Value: "content",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "More",
				},
			},
			Level: 1,
		},
	},
}

var tablesTestShouldDiscoverAnchorAtStartOfCellAndRegisterItAsAReference = &asciidoc.Document{
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
								&asciidoc.Anchor{
									ID: "mount-evans",
									Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "14,271 feet",
								},
							},
							Blank: false,
						},
					},
				},
				&asciidoc.EmptyLine{
					Text: "\n",
				},
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
									Value: 3,
									IsSet: true,
								},
							},
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
							Elements: asciidoc.Elements{
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

var tablesTestShouldCatalogAnchorAtStartOfCellInImplicitHeaderRowWhenColumnHasAStyle = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
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
								&asciidoc.Anchor{
									ID: "foo",
									Elements: asciidoc.Elements{
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
				&asciidoc.EmptyLine{
					Text: "\n",
				},
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
									Value: " AsciiDoc",
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

var tablesTestShouldCatalogAnchorAtStartOfCellInExplicitHeaderRowWhenColumnHasAStyle = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Table{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: nil,
					ID:    nil,
					Roles: nil,
					Options: []*asciidoc.ShorthandOption{
						&asciidoc.ShorthandOption{
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "header",
								},
							},
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
								&asciidoc.Anchor{
									ID: "foo",
									Elements: asciidoc.Elements{
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
									Value: " AsciiDoc",
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

var tablesTestShouldCatalogAnchorAtStartOfCellInFirstRow = &asciidoc.Document{
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
									Value: 0,
									IsSet: false,
								},
							},
							Elements: asciidoc.Elements{
								&asciidoc.Anchor{
									ID: "foo",
									Elements: asciidoc.Elements{
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

var tablesTestFootnotesShouldNotBeSharedBetweenAnAsciiDocTableCellAndTheMainDocument = &asciidoc.Document{
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

var tablesTestCalloutNumbersShouldBeGloballyUniqueIncludingAsciiDocTableCells = &asciidoc.Document{
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
																	Value: "yaml",
																},
															},
														},
													},
													Delimiter: asciidoc.Delimiter{
														Type:   5,
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
						&asciidoc.EmptyLine{
							Text: "",
						},
					},
					Title: asciidoc.Elements{
						&asciidoc.String{
							Value: "Section 1",
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
																	Value: "yaml",
																},
															},
														},
													},
													Delimiter: asciidoc.Delimiter{
														Type:   5,
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
						&asciidoc.EmptyLine{
							Text: "",
						},
					},
					Title: asciidoc.Elements{
						&asciidoc.String{
							Value: "Section 2",
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
											Value: "yaml",
										},
									},
								},
							},
							Delimiter: asciidoc.Delimiter{
								Type:   5,
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
					Title: asciidoc.Elements{
						&asciidoc.String{
							Value: "Section 3",
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

var tablesTestCompatModeCanBeActivatedInAsciiDocTableCell = &asciidoc.Document{
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
								&asciidoc.NewLine{},
								&asciidoc.AttributeEntry{
									Name:     "compat-mode",
									Elements: nil,
								},
								&asciidoc.EmptyLine{
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

var tablesTestCompatModeInAsciiDocTableCellInheritsFromParentDocument = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name:     "compat-mode",
			Elements: nil,
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "The word 'italic' is emphasized.",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
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
									Value: "The word 'oblique' is emphasized.",
								},
							},
							Blank: false,
						},
					},
				},
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
								&asciidoc.String{
									Value: "The word 'slanted' is emphasized.",
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
		&asciidoc.String{
			Value: "The word 'askew' is emphasized.",
		},
		&asciidoc.NewLine{},
	},
}

var tablesTestCompatModeInAsciiDocTableCellCanBeUnsetIfSetInParentDocument = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name:     "compat-mode",
			Elements: nil,
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "The word 'italic' is emphasized.",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
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
									Value: "The word 'oblique' is emphasized.",
								},
							},
							Blank: false,
						},
					},
				},
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
								&asciidoc.AttributeReset{
									Name: "compat-mode",
								},
								&asciidoc.EmptyLine{
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
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "The word 'askew' is emphasized.",
		},
		&asciidoc.NewLine{},
	},
}

var tablesTestNestedTable = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
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
							Elements: asciidoc.Elements{
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
									Elements: asciidoc.Elements{
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

var tablesTestCanSetFormatOfNestedTableToPsv = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
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
							Elements: asciidoc.Elements{
								&asciidoc.NewLine{},
								&asciidoc.Paragraph{
									AttributeList: asciidoc.AttributeList{
										&asciidoc.NamedAttribute{
											Name: "format",
											Val: asciidoc.Elements{
												&asciidoc.String{
													Value: "psv",
												},
											},
											Quote: 0,
										},
									},
									Elements: asciidoc.Elements{
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

var tablesTestAsciiDocTableCellShouldInheritToDirOptionFromParentDocument = &asciidoc.Document{
	Elements: asciidoc.Elements{
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

var tablesTestAsciiDocTableCellShouldNotInheritTocSettingFromParentDocument = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.AttributeEntry{
					Name:     "toc",
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
												&asciidoc.Section{
													AttributeList: nil,
													Elements:      nil,
													Title: asciidoc.Elements{
														&asciidoc.String{
															Value: "Section in Nested Document",
														},
													},
													Level: 1,
												},
												&asciidoc.EmptyLine{
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
					Title: asciidoc.Elements{
						&asciidoc.String{
							Value: "Section",
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

var tablesTestShouldBeAbleToEnableTocInAnAsciiDocTableCell = &asciidoc.Document{
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
												&asciidoc.Section{
													AttributeList: nil,
													Elements:      nil,
													Title: asciidoc.Elements{
														&asciidoc.String{
															Value: "Subdocument Title",
														},
													},
													Level: 0,
												},
												&asciidoc.AttributeEntry{
													Name:     "toc",
													Elements: nil,
												},
												&asciidoc.EmptyLine{
													Text: "",
												},
												&asciidoc.Section{
													AttributeList: nil,
													Elements:      nil,
													Title: asciidoc.Elements{
														&asciidoc.String{
															Value: "Subdocument Section A",
														},
													},
													Level: 1,
												},
												&asciidoc.EmptyLine{
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
					Title: asciidoc.Elements{
						&asciidoc.String{
							Value: "Section A",
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

var tablesTestShouldBeAbleToEnableTocInAnAsciiDocTableCellEvenIfHardUnsetByApi = &asciidoc.Document{
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
												&asciidoc.Section{
													AttributeList: nil,
													Elements:      nil,
													Title: asciidoc.Elements{
														&asciidoc.String{
															Value: "Subdocument Title",
														},
													},
													Level: 0,
												},
												&asciidoc.AttributeEntry{
													Name:     "toc",
													Elements: nil,
												},
												&asciidoc.EmptyLine{
													Text: "",
												},
												&asciidoc.Section{
													AttributeList: nil,
													Elements:      nil,
													Title: asciidoc.Elements{
														&asciidoc.String{
															Value: "Subdocument Section A",
														},
													},
													Level: 1,
												},
												&asciidoc.EmptyLine{
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
					Title: asciidoc.Elements{
						&asciidoc.String{
							Value: "Section A",
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

var tablesTestShouldBeAbleToEnableTocInBothOuterDocumentAndInAnAsciiDocTableCell = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.AttributeEntry{
					Name:     "toc",
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
												&asciidoc.Section{
													AttributeList: nil,
													Elements:      nil,
													Title: asciidoc.Elements{
														&asciidoc.String{
															Value: "Subdocument Title",
														},
													},
													Level: 0,
												},
												&asciidoc.AttributeEntry{
													Name: "toc",
													Elements: asciidoc.Elements{
														&asciidoc.String{
															Value: "macro",
														},
													},
												},
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
																		Value: "table-cell-toc",
																	},
																},
															},
															Roles:   nil,
															Options: nil,
														},
													},
													Elements: asciidoc.Elements{
														&asciidoc.String{
															Value: "toc::[]",
														},
														&asciidoc.NewLine{},
													},
													Admonition: 0,
												},
												&asciidoc.EmptyLine{
													Text: "",
												},
												&asciidoc.Section{
													AttributeList: nil,
													Elements:      nil,
													Title: asciidoc.Elements{
														&asciidoc.String{
															Value: "Subdocument Section A",
														},
													},
													Level: 1,
												},
												&asciidoc.EmptyLine{
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
					Title: asciidoc.Elements{
						&asciidoc.String{
							Value: "Section A",
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

var tablesTestDocumentInAnAsciiDocTableCellShouldNotSeeDoctitleOfParent = &asciidoc.Document{
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
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Document Title",
				},
			},
			Level: 0,
		},
	},
}

var tablesTestCellBackgroundColor = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
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
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "header",
						},
					},
					Quote: 2,
				},
			},
			ColumnCount: 2,
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
							Elements: asciidoc.Elements{
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
							Elements: asciidoc.Elements{
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

var tablesTestShouldWarnIfTableBlockIsNotTerminated = &asciidoc.Document{
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

var tablesTestShouldShowCorrectLineNumberInWarningAboutUnterminatedBlockInsideAsciiDocTableCell = &asciidoc.Document{
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
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
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
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
		&asciidoc.EmptyLine{
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
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "eof",
		},
		&asciidoc.NewLine{},
	},
}

var tablesTestCustomSeparatorForAnAsciiDocTableCell = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
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
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "!",
						},
					},
					Quote: 0,
				},
			},
			Elements: asciidoc.Elements{
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
				Type:   5,
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

var tablesTestTableWithBreakableOptionDocbook5 = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Table{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.TitleAttribute{
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Table with breakable",
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
									Value: "breakable",
								},
							},
						},
					},
				},
			},
			ColumnCount: 2,
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "Quantity",
								},
							},
							Blank: false,
						},
					},
				},
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
							Elements: asciidoc.Elements{
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

var tablesTestTableWithUnbreakableOptionDocbook5 = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Table{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.TitleAttribute{
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "Table with unbreakable",
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
									Value: "unbreakable",
								},
							},
						},
					},
				},
			},
			ColumnCount: 2,
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
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "Quantity",
								},
							},
							Blank: false,
						},
					},
				},
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
							Elements: asciidoc.Elements{
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

var tablesTestNoImplicitHeaderRowIfCellInFirstLineIsQuotedAndSpansMultipleLines = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
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
			Elements: asciidoc.Elements{
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
		&asciidoc.EmptyLine{
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

var tablesTestConvertsSimpleDsvTable = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Paragraph{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.NamedAttribute{
					Name: "width",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "75%",
						},
					},
					Quote: 2,
				},
				&asciidoc.NamedAttribute{
					Name: "format",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "dsv",
						},
					},
					Quote: 2,
				},
			},
			Elements: asciidoc.Elements{
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

var tablesTestDsvFormatShorthand = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
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

var tablesTestSingleCellInDsvTableShouldOnlyProduceSingleRow = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
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

var tablesTestShouldTreatTrailingColonAsAnEmptyCell = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
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

var tablesTestShouldTreatTrailingCommaAsAnEmptyCell = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
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

var tablesTestShouldLogErrorButNotCrashIfCellDataHasUnclosedQuote = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
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

var tablesTestShouldPreserveNewlinesInQuotedCsvValues = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
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
			Elements: asciidoc.Elements{
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
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "two",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "three\",\"do",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "re",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
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

var tablesTestMixedUnquotedRecordsAndQuotedRecordsWithEscapedQuotesCommasAndWrappedLines = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Paragraph{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.NamedAttribute{
					Name: "format",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "csv",
						},
					},
					Quote: 2,
				},
				&asciidoc.NamedAttribute{
					Name: "options",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "header",
						},
					},
					Quote: 2,
				},
			},
			Elements: asciidoc.Elements{
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

var tablesTestShouldAllowQuotesAroundACsvValueToBeOnTheirOwnLines = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
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
			Elements: asciidoc.Elements{
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

var tablesTestCsvFormatShorthand = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
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

var tablesTestCustomCsvSeparator = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Paragraph{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.NamedAttribute{
					Name: "format",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "csv",
						},
					},
					Quote: 0,
				},
				&asciidoc.NamedAttribute{
					Name: "separator",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: ";",
						},
					},
					Quote: 0,
				},
			},
			Elements: asciidoc.Elements{
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

var tablesTestSingleCellInCsvTableShouldOnlyProduceSingleRow = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
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

var tablesTestCellFormattedWithAsciiDocStyle = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
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
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: ";",
						},
					},
					Quote: 0,
				},
			},
			Elements: asciidoc.Elements{
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
		&asciidoc.EmptyLine{
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

var tablesTestShouldStripWhitespaceAroundContentsOfAsciiDocCell = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
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
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: ";",
						},
					},
					Quote: 0,
				},
			},
			Elements: asciidoc.Elements{
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
		&asciidoc.EmptyLine{
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
