package tests

import (
	"testing"

	"github.com/project-chip/alchemy/asciidoc"
)

func TestApi(t *testing.T) {
	apiTests.run(t)
}

var apiTests = parseTests{

	{"should load input file", "asciidoctor/api_test_should_load_input_file.adoc", apiTestShouldLoadInputFile},

	{"should load input string", "asciidoctor/api_test_should_load_input_string.adoc", apiTestShouldLoadInputString},

	{"should load input string array", "asciidoctor/api_test_should_load_input_string_array.adoc", apiTestShouldLoadInputStringArray},

	{"should load nil input", "asciidoctor/api_test_should_load_nil_input.adoc", apiTestShouldLoadNilInput},

	{"render method on node is aliased to convert method", "asciidoctor/api_test_render_method_on_node_is_aliased_to_convert_method.adoc", apiTestRenderMethodOnNodeIsAliasedToConvertMethod},

	{"content method on Inline node is aliased to text method", "asciidoctor/api_test_content_method_on_inline_node_is_aliased_to_text_method.adoc", apiTestContentMethodOnInlineNodeIsAliasedToTextMethod},

	{"should track file and line information with blocks if sourcemap option is set", "asciidoctor/api_test_should_track_file_and_line_information_with_blocks_if_sourcemap_option_is_set.adoc", apiTestShouldTrackFileAndLineInformationWithBlocksIfSourcemapOptionIsSet},

	{"should assign correct lineno for multi-line paragraph inside a conditional preprocessor directive", "asciidoctor/api_test_should_assign_correct_lineno_for_multi_line_paragraph_inside_a_conditional_preprocessor_directive.adoc", apiTestShouldAssignCorrectLinenoForMultiLineParagraphInsideAConditionalPreprocessorDirective},

	{"should assign correct source location to blocks that follow a detached list continuation", "asciidoctor/api_test_should_assign_correct_source_location_to_blocks_that_follow_a_detached_list_continuation.adoc", apiTestShouldAssignCorrectSourceLocationToBlocksThatFollowADetachedListContinuation},

	{"should assign correct source location if section occurs on last line of input", "asciidoctor/api_test_should_assign_correct_source_location_if_section_occurs_on_last_line_of_input.adoc", apiTestShouldAssignCorrectSourceLocationIfSectionOccursOnLastLineOfInput},

	{"should allow sourcemap option on document to be modified before document is parsed", "asciidoctor/api_test_should_allow_sourcemap_option_on_document_to_be_modified_before_document_is_parsed.adoc", apiTestShouldAllowSourcemapOptionOnDocumentToBeModifiedBeforeDocumentIsParsed},

	{"find_by should return an empty Array if no matches are found", "asciidoctor/api_test_find_by_should_return_an_empty_array_if_no_matches_are_found.adoc", apiTestFindByShouldReturnAnEmptyArrayIfNoMatchesAreFound},

	{"find_by should discover blocks inside AsciiDoc table cells if traverse_documents selector option is true", "asciidoctor/api_test_find_by_should_discover_blocks_inside_ascii_doc_table_cells_if_traverse_documents_selector_option_is_true.adoc", apiTestFindByShouldDiscoverBlocksInsideAsciiDocTableCellsIfTraverseDocumentsSelectorOptionIsTrue},

	{"find_by should return inner document of AsciiDoc table cell if traverse_documents selector option is true", "asciidoctor/api_test_find_by_should_return_inner_document_of_ascii_doc_table_cell_if_traverse_documents_selector_option_is_true.adoc", apiTestFindByShouldReturnInnerDocumentOfAsciiDocTableCellIfTraverseDocumentsSelectorOptionIsTrue},

	{"find_by should match table cells", "asciidoctor/api_test_find_by_should_match_table_cells.adoc", apiTestFindByShouldMatchTableCells},

	{"find_by should return Array of blocks that match style criteria", "asciidoctor/api_test_find_by_should_return_array_of_blocks_that_match_style_criteria.adoc", apiTestFindByShouldReturnArrayOfBlocksThatMatchStyleCriteria},

	{"find_by should return Array of blocks that match role criteria", "asciidoctor/api_test_find_by_should_return_array_of_blocks_that_match_role_criteria.adoc", apiTestFindByShouldReturnArrayOfBlocksThatMatchRoleCriteria},

	{"find_by should return the document title section if context selector is :section", "asciidoctor/api_test_find_by_should_return_the_document_title_section_if_context_selector_is_section.adoc", apiTestFindByShouldReturnTheDocumentTitleSectionIfContextSelectorIssection},

	{"find_by should only return results for which the block argument yields true", "asciidoctor/api_test_find_by_should_only_return_results_for_which_the_block_argument_yields_true.adoc", apiTestFindByShouldOnlyReturnResultsForWhichTheBlockArgumentYieldsTrue},

	{"find_by should reject node and its children if block returns :reject", "asciidoctor/api_test_find_by_should_reject_node_and_its_children_if_block_returns_reject.adoc", apiTestFindByShouldRejectNodeAndItsChildrenIfBlockReturnsreject},

	{"find_by should reject node matched by ID selector if block returns :reject", "asciidoctor/api_test_find_by_should_reject_node_matched_by_id_selector_if_block_returns_reject.adoc", apiTestFindByShouldRejectNodeMatchedByIdSelectorIfBlockReturnsreject},

	{"find_by should accept node matched by ID selector if block returns :prune", "asciidoctor/api_test_find_by_should_accept_node_matched_by_id_selector_if_block_returns_prune.adoc", apiTestFindByShouldAcceptNodeMatchedByIdSelectorIfBlockReturnsprune},

	{"find_by should accept node but reject its children if block returns :prune", "asciidoctor/api_test_find_by_should_accept_node_but_reject_its_children_if_block_returns_prune.adoc", apiTestFindByShouldAcceptNodeButRejectItsChildrenIfBlockReturnsprune},

	{"find_by should stop looking for blocks when StopIteration is raised", "asciidoctor/api_test_find_by_should_stop_looking_for_blocks_when_stop_iteration_is_raised.adoc", apiTestFindByShouldStopLookingForBlocksWhenStopIterationIsRaised},

	{"find_by should stop looking for blocks when filter block returns :stop directive", "asciidoctor/api_test_find_by_should_stop_looking_for_blocks_when_filter_block_returns_stop_directive.adoc", apiTestFindByShouldStopLookingForBlocksWhenFilterBlockReturnsstopDirective},

	{"find_by should only return one result when matching by id", "asciidoctor/api_test_find_by_should_only_return_one_result_when_matching_by_id.adoc", apiTestFindByShouldOnlyReturnOneResultWhenMatchingById},

	{"find_by should stop seeking once match is found", "asciidoctor/api_test_find_by_should_stop_seeking_once_match_is_found.adoc", apiTestFindByShouldStopSeekingOnceMatchIsFound},

	{"find_by should return an empty Array if the id criteria matches but the block argument yields false", "asciidoctor/api_test_find_by_should_return_an_empty_array_if_the_id_criteria_matches_but_the_block_argument_yields_false.adoc", apiTestFindByShouldReturnAnEmptyArrayIfTheIdCriteriaMatchesButTheBlockArgumentYieldsFalse},

	{"find_by should not crash if dlist entry does not have description", "asciidoctor/api_test_find_by_should_not_crash_if_dlist_entry_does_not_have_description.adoc", apiTestFindByShouldNotCrashIfDlistEntryDoesNotHaveDescription},

	{"can substitute an extended syntax highlighter factory implementation using the :syntax_highlighters option", "asciidoctor/api_test_can_substitute_an_extended_syntax_highlighter_factory_implementation_using_the_syntax_highlighters_option.adoc", apiTestCanSubstituteAnExtendedSyntaxHighlighterFactoryImplementationUsingThesyntaxHighlightersOption},

	{"render_file is aliased to convert_file", "asciidoctor/api_test_render_file_is_aliased_to_convert_file.adoc", apiTestRenderFileIsAliasedToConvertFile},

	{"should embed remote stylesheet by default if SafeMode is less than SECURE and allow-uri-read is set", "asciidoctor/api_test_should_embed_remote_stylesheet_by_default_if_safe_mode_is_less_than_secure_and_allow_uri_read_is_set.adoc", apiTestShouldEmbedRemoteStylesheetByDefaultIfSafeModeIsLessThanSecureAndAllowUriReadIsSet},

	{"should not allow linkcss be unset from document if SafeMode is SECURE or greater", "asciidoctor/api_test_should_not_allow_linkcss_be_unset_from_document_if_safe_mode_is_secure_or_greater.adoc", apiTestShouldNotAllowLinkcssBeUnsetFromDocumentIfSafeModeIsSecureOrGreater},

	{"should embed default stylesheet if linkcss is unset from API and SafeMode is SECURE or greater", "asciidoctor/api_test_should_embed_default_stylesheet_if_linkcss_is_unset_from_api_and_safe_mode_is_secure_or_greater.adoc", apiTestShouldEmbedDefaultStylesheetIfLinkcssIsUnsetFromApiAndSafeModeIsSecureOrGreater},

	{"should embed default stylesheet if safe mode is less than SECURE and linkcss is unset from API", "asciidoctor/api_test_should_embed_default_stylesheet_if_safe_mode_is_less_than_secure_and_linkcss_is_unset_from_api.adoc", apiTestShouldEmbedDefaultStylesheetIfSafeModeIsLessThanSecureAndLinkcssIsUnsetFromApi},

	{"should link to custom stylesheet if specified in stylesheet attribute", "asciidoctor/api_test_should_link_to_custom_stylesheet_if_specified_in_stylesheet_attribute.adoc", apiTestShouldLinkToCustomStylesheetIfSpecifiedInStylesheetAttribute},

	{"should resolve custom stylesheet relative to stylesdir", "asciidoctor/api_test_should_resolve_custom_stylesheet_relative_to_stylesdir.adoc", apiTestShouldResolveCustomStylesheetRelativeToStylesdir},

	{"should resolve custom stylesheet to embed relative to stylesdir", "asciidoctor/api_test_should_resolve_custom_stylesheet_to_embed_relative_to_stylesdir.adoc", apiTestShouldResolveCustomStylesheetToEmbedRelativeToStylesdir},

	{"should embed custom stylesheet in remote stylesdir if SafeMode is less than SECURE and allow-uri-read is set", "asciidoctor/api_test_should_embed_custom_stylesheet_in_remote_stylesdir_if_safe_mode_is_less_than_secure_and_allow_uri_read_is_set.adoc", apiTestShouldEmbedCustomStylesheetInRemoteStylesdirIfSafeModeIsLessThanSecureAndAllowUriReadIsSet},

	{"should respect outfilesuffix soft set from API", "asciidoctor/api_test_should_respect_outfilesuffix_soft_set_from_api.adoc", apiTestShouldRespectOutfilesuffixSoftSetFromApi},

	{"with no author", "asciidoctor/api_test_with_no_author.adoc", apiTestWithNoAuthor},

	{"with one author", "asciidoctor/api_test_with_one_author.adoc", apiTestWithOneAuthor},

	{"with two authors", "asciidoctor/api_test_with_two_authors.adoc", apiTestWithTwoAuthors},

	{"with authors as attributes", "asciidoctor/api_test_with_authors_as_attributes.adoc", apiTestWithAuthorsAsAttributes},

	{"should not crash if nil cell text is passed to Cell constructor", "asciidoctor/api_test_should_not_crash_if_nil_cell_text_is_passed_to_cell_constructor.adoc", apiTestShouldNotCrashIfNilCellTextIsPassedToCellConstructor},

	{"should set option on node when set_option is called", "asciidoctor/api_test_should_set_option_on_node_when_set_option_is_called.adoc", apiTestShouldSetOptionOnNodeWhenSetOptionIsCalled},

	{"enabled_options should return all options which are set", "asciidoctor/api_test_enabled_options_should_return_all_options_which_are_set.adoc", apiTestEnabledOptionsShouldReturnAllOptionsWhichAreSet},

	{"should append option to existing options", "asciidoctor/api_test_should_append_option_to_existing_options.adoc", apiTestShouldAppendOptionToExistingOptions},

	{"should not append option if option is already set", "asciidoctor/api_test_should_not_append_option_if_option_is_already_set.adoc", apiTestShouldNotAppendOptionIfOptionIsAlreadySet},

	{"should return set of option names", "asciidoctor/api_test_should_return_set_of_option_names.adoc", apiTestShouldReturnSetOfOptionNames},

	{"should set linenums option if linenums enabled on source block", "asciidoctor/api_test_should_set_linenums_option_if_linenums_enabled_on_source_block.adoc", apiTestShouldSetLinenumsOptionIfLinenumsEnabledOnSourceBlock},

	{"should set linenums option if linenums enabled on fenced code block", "asciidoctor/api_test_should_set_linenums_option_if_linenums_enabled_on_fenced_code_block.adoc", apiTestShouldSetLinenumsOptionIfLinenumsEnabledOnFencedCodeBlock},

	{"should not set linenums attribute if linenums option is enabled on source block", "asciidoctor/api_test_should_not_set_linenums_attribute_if_linenums_option_is_enabled_on_source_block.adoc", apiTestShouldNotSetLinenumsAttributeIfLinenumsOptionIsEnabledOnSourceBlock},

	{"should not set linenums attribute if linenums option is enabled on fenced code block", "asciidoctor/api_test_should_not_set_linenums_attribute_if_linenums_option_is_enabled_on_fenced_code_block.adoc", apiTestShouldNotSetLinenumsAttributeIfLinenumsOptionIsEnabledOnFencedCodeBlock},

	{"table column should not be a block or inline", "asciidoctor/api_test_table_column_should_not_be_a_block_or_inline.adoc", apiTestTableColumnShouldNotBeABlockOrInline},

	{"table cell should be a block", "asciidoctor/api_test_table_cell_should_be_a_block.adoc", apiTestTableCellShouldBeABlock},

	{"next_adjacent_block should return next block", "asciidoctor/api_test_next_adjacent_block_should_return_next_block.adoc", apiTestNextAdjacentBlockShouldReturnNextBlock},

	{"next_adjacent_block should return next sibling of parent if called on last sibling", "asciidoctor/api_test_next_adjacent_block_should_return_next_sibling_of_parent_if_called_on_last_sibling.adoc", apiTestNextAdjacentBlockShouldReturnNextSiblingOfParentIfCalledOnLastSibling},

	{"next_adjacent_block should return next sibling of list if called on last item", "asciidoctor/api_test_next_adjacent_block_should_return_next_sibling_of_list_if_called_on_last_item.adoc", apiTestNextAdjacentBlockShouldReturnNextSiblingOfListIfCalledOnLastItem},

	{"next_adjacent_block should return next item in dlist if called on last block of list item", "asciidoctor/api_test_next_adjacent_block_should_return_next_item_in_dlist_if_called_on_last_block_of_list_item.adoc", apiTestNextAdjacentBlockShouldReturnNextItemInDlistIfCalledOnLastBlockOfListItem},

	{"should return true when sections? is called on a document or section that has sections", "asciidoctor/api_test_should_return_true_when_sections_is_called_on_a_document_or_section_that_has_sections.adoc", apiTestShouldReturnTrueWhenSectionsIsCalledOnADocumentOrSectionThatHasSections},

	{"should return false when sections? is called on a document with no sections", "asciidoctor/api_test_should_return_false_when_sections_is_called_on_a_document_with_no_sections.adoc", apiTestShouldReturnFalseWhenSectionsIsCalledOnADocumentWithNoSections},

	{"should return false when sections? is called on a section with no sections", "asciidoctor/api_test_should_return_false_when_sections_is_called_on_a_section_with_no_sections.adoc", apiTestShouldReturnFalseWhenSectionsIsCalledOnASectionWithNoSections},

	{"should return false when sections? is called on anything that is not a section", "asciidoctor/api_test_should_return_false_when_sections_is_called_on_anything_that_is_not_a_section.adoc", apiTestShouldReturnFalseWhenSectionsIsCalledOnAnythingThatIsNotASection},
}

var apiTestShouldLoadInputFile = &asciidoc.Document{
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
		&asciidoc.String{
			Value: "preamble",
		},
		&asciidoc.NewLine{},
	},
}

var apiTestShouldLoadInputString = &asciidoc.Document{
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
		&asciidoc.String{
			Value: "preamble",
		},
		&asciidoc.NewLine{},
	},
}

var apiTestShouldLoadInputStringArray = &asciidoc.Document{
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
		&asciidoc.String{
			Value: "preamble",
		},
		&asciidoc.NewLine{},
	},
}

var apiTestShouldLoadNilInput = &asciidoc.Document{
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
					Value: "      <div class=\"paragraph\">",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "      <p>paragraph text</p>",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "      </div>",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 0,
		},
	},
}

var apiTestRenderMethodOnNodeIsAliasedToConvertMethod = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "paragraph text",
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
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
	},
}

var apiTestContentMethodOnInlineNodeIsAliasedToTextMethod = &asciidoc.Document{
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
							Value: "bar",
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
				&asciidoc.EmptyLine{
					Text: "",
				},
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

var apiTestShouldTrackFileAndLineInformationWithBlocksIfSourcemapOptionIsSet = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name:     "conditional-attribute",
			Elements: nil,
		},
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
		&asciidoc.IfDef{
			Attributes: asciidoc.AttributeNames{
				"conditional-attribute",
			},
			Union:  0,
			Inline: false,
		},
		&asciidoc.String{
			Value: "subject",
		},
		&asciidoc.NewLine{},
		&asciidoc.EndIf{
			Attributes: nil,
			Union:      0,
			Open:       nil,
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

var apiTestShouldAssignCorrectLinenoForMultiLineParagraphInsideAConditionalPreprocessorDirective = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name:     "conditional-attribute",
			Elements: nil,
		},
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
		&asciidoc.IfDef{
			Attributes: asciidoc.AttributeNames{
				"conditional-attribute",
			},
			Union:  0,
			Inline: false,
		},
		&asciidoc.String{
			Value: "subject",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "subject",
		},
		&asciidoc.NewLine{},
		&asciidoc.EndIf{
			Attributes: nil,
			Union:      0,
			Open:       nil,
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

var apiTestShouldAssignCorrectSourceLocationToBlocksThatFollowADetachedListContinuation = &asciidoc.Document{
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
				&asciidoc.ListContinuation{
					ChildElement: &asciidoc.Paragraph{
						AttributeList: nil,
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "paragraph attached to parent",
							},
						},
						Admonition: 0,
					},
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
		&asciidoc.SidebarBlock{
			Delimiter: asciidoc.Delimiter{
				Type:   8,
				Length: 4,
			},
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "sidebar outside list",
				},
				&asciidoc.NewLine{},
			},
		},
	},
}

var apiTestShouldAssignCorrectSourceLocationIfSectionOccursOnLastLineOfInput = &asciidoc.Document{
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
							Value: "content",
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
					Elements:      nil,
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

var apiTestShouldAllowSourcemapOptionOnDocumentToBeModifiedBeforeDocumentIsParsed = &asciidoc.Document{
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
							Value: "paragraph",
						},
						&asciidoc.NewLine{},
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
								&asciidoc.DescriptionListItem{
									Elements: asciidoc.Elements{
										&asciidoc.NewLine{},
										&asciidoc.LineBreak{},
									},
									AttributeList: nil,
									Marker:        "::",
									Term: asciidoc.Elements{
										&asciidoc.String{
											Value: "Exhibit A",
										},
									},
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
													Value: "#tiger.animal",
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
											Value: "Shoe",
										},
									},
								},
							},
							ImagePath: asciidoc.Elements{
								&asciidoc.String{
									Value: "shoe.png",
								},
							},
						},
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
							Value: "paragraph",
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

var apiTestFindByShouldReturnAnEmptyArrayIfNoMatchesAreFound = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "foo",
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
					Value: "yin",
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
					Value: "zen",
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
					Value: "yang",
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
					Value: "bar",
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
					Value: "baz",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
	},
}

var apiTestFindByShouldDiscoverBlocksInsideAsciiDocTableCellsIfTraverseDocumentsSelectorOptionIsTrue = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "paragraph in parent document (before)",
		},
		&asciidoc.NewLine{},
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
									Value: "footer",
								},
							},
						},
					},
				},
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
								&asciidoc.NewLine{},
								&asciidoc.String{
									Value: "paragraph in nested document (body)",
								},
							},
							Blank: false,
						},
						&asciidoc.TableCell{
							Format: &asciidoc.TableCellFormat{
								Multiplier: asciidoc.Optional[int]{
									Value: 1,
									IsSet: false,
								},
								Span: asciidoc.TableCellSpan{
									Column: asciidoc.Optional[int]{
										Value: 1,
										IsSet: false,
									},
									Row: asciidoc.Optional[int]{
										Value: 1,
										IsSet: false,
									},
								},
								HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
									Value: 0,
									IsSet: false,
								},
								VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
									Value: 0,
									IsSet: false,
								},
								Style: asciidoc.Optional[asciidoc.TableCellStyle]{
									Value: 0,
									IsSet: false,
								},
							},
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "normal table cell",
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
									Value: 1,
									IsSet: true,
								},
							},
							Elements: asciidoc.Elements{
								&asciidoc.NewLine{},
								&asciidoc.String{
									Value: "paragraph in nested document (foot)",
								},
							},
							Blank: false,
						},
						&asciidoc.TableCell{
							Format: &asciidoc.TableCellFormat{
								Multiplier: asciidoc.Optional[int]{
									Value: 1,
									IsSet: false,
								},
								Span: asciidoc.TableCellSpan{
									Column: asciidoc.Optional[int]{
										Value: 1,
										IsSet: false,
									},
									Row: asciidoc.Optional[int]{
										Value: 1,
										IsSet: false,
									},
								},
								HorizontalAlign: asciidoc.Optional[asciidoc.TableCellHorizontalAlign]{
									Value: 0,
									IsSet: false,
								},
								VerticalAlign: asciidoc.Optional[asciidoc.TableCellVerticalAlign]{
									Value: 0,
									IsSet: false,
								},
								Style: asciidoc.Optional[asciidoc.TableCellStyle]{
									Value: 0,
									IsSet: false,
								},
							},
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "normal table cell",
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
			Value: "paragraph in parent document (after)",
		},
		&asciidoc.NewLine{},
	},
}

var apiTestFindByShouldReturnInnerDocumentOfAsciiDocTableCellIfTraverseDocumentsSelectorOptionIsTrue = &asciidoc.Document{
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
									Value: "paragraph in nested document",
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

var apiTestFindByShouldMatchTableCells = &asciidoc.Document{
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
									Value: "1",
								},
								&asciidoc.NewLine{},
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
									Value: 1,
									IsSet: true,
								},
							},
							Elements: asciidoc.Elements{
								&asciidoc.Paragraph{
									AttributeList: nil,
									Elements: asciidoc.Elements{
										&asciidoc.String{
											Value: "2, as it goes.",
										},
									},
									Admonition: 1,
								},
							},
							Blank: false,
						},
						&asciidoc.TableCell{
							Format: &asciidoc.TableCellFormat{
								Multiplier: asciidoc.Optional[int]{
									Value: 1,
									IsSet: false,
								},
								Span: asciidoc.TableCellSpan{
									Column: asciidoc.Optional[int]{
										Value: 1,
										IsSet: false,
									},
									Row: asciidoc.Optional[int]{
										Value: 1,
										IsSet: false,
									},
								},
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
									Value: "3",
								},
								&asciidoc.NewLine{},
								&asciidoc.String{
									Value: " you",
								},
								&asciidoc.NewLine{},
								&asciidoc.String{
									Value: "me",
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

var apiTestFindByShouldReturnArrayOfBlocksThatMatchStyleCriteria = &asciidoc.Document{
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
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "square",
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
		&asciidoc.ThematicBreak{
			AttributeList: nil,
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "apples",
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
					Value: "bananas",
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
					Value: "pears",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     0,
		},
	},
}

var apiTestFindByShouldReturnArrayOfBlocksThatMatchRoleCriteria = &asciidoc.Document{
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
							Value: "#tiger.animal",
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
							Value: "Shoe",
						},
					},
				},
			},
			ImagePath: asciidoc.Elements{
				&asciidoc.String{
					Value: "shoe.png",
				},
			},
		},
	},
}

var apiTestFindByShouldReturnTheDocumentTitleSectionIfContextSelectorIssection = &asciidoc.Document{
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
					Value: "Document Title",
				},
			},
			Level: 0,
		},
	},
}

var apiTestFindByShouldOnlyReturnResultsForWhichTheBlockArgumentYieldsTrue = &asciidoc.Document{
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
							Value: "Subsection",
						},
					},
					Level: 2,
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

var apiTestFindByShouldRejectNodeAndItsChildrenIfBlockReturnsreject = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "paragraph 1",
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
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "paragraph 2",
				},
				&asciidoc.NewLine{},
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
					Value: "paragraph 3",
				},
				&asciidoc.NewLine{},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "paragraph 4",
		},
		&asciidoc.NewLine{},
	},
}

var apiTestFindByShouldRejectNodeMatchedByIdSelectorIfBlockReturnsreject = &asciidoc.Document{
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
					Value: "paragraph 1",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 0,
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
								Value: "idname",
							},
						},
					},
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
					Value: "paragraph 2",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 0,
		},
	},
}

var apiTestFindByShouldAcceptNodeMatchedByIdSelectorIfBlockReturnsprune = &asciidoc.Document{
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
					Value: "paragraph 1",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 0,
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
				&asciidoc.ShorthandAttribute{
					Style: nil,
					ID: &asciidoc.ShorthandID{
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "idname",
							},
						},
					},
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
					Value: "paragraph 2",
				},
				&asciidoc.NewLine{},
			},
		},
	},
}

var apiTestFindByShouldAcceptNodeButRejectItsChildrenIfBlockReturnsprune = &asciidoc.Document{
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
					Value: "paragraph 2",
				},
				&asciidoc.NewLine{},
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
					Value: "paragraph 3",
				},
				&asciidoc.NewLine{},
			},
		},
	},
}

var apiTestFindByShouldStopLookingForBlocksWhenStopIterationIsRaised = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "paragraph 1",
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
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "paragraph 2",
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
							Value: "paragraph 3",
						},
						&asciidoc.NewLine{},
					},
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "paragraph 4",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "item",
				},
				&asciidoc.ListContinuation{
					ChildElement: &asciidoc.Paragraph{
						AttributeList: nil,
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "paragraph 5",
							},
						},
						Admonition: 0,
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

var apiTestFindByShouldStopLookingForBlocksWhenFilterBlockReturnsstopDirective = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "paragraph 1",
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
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "paragraph 2",
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
							Value: "paragraph 3",
						},
						&asciidoc.NewLine{},
					},
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "paragraph 4",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "item",
				},
				&asciidoc.ListContinuation{
					ChildElement: &asciidoc.Paragraph{
						AttributeList: nil,
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "paragraph 5",
							},
						},
						Admonition: 0,
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

var apiTestFindByShouldOnlyReturnOneResultWhenMatchingById = &asciidoc.Document{
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
										Value: "subsection",
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
							Value: "content",
						},
						&asciidoc.NewLine{},
					},
					Title: asciidoc.Elements{
						&asciidoc.String{
							Value: "Subsection",
						},
					},
					Level: 2,
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

var apiTestFindByShouldStopSeekingOnceMatchIsFound = &asciidoc.Document{
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
										Value: "subsection",
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
						&asciidoc.Paragraph{
							AttributeList: asciidoc.AttributeList{
								&asciidoc.ShorthandAttribute{
									Style: nil,
									ID: &asciidoc.ShorthandID{
										Elements: asciidoc.Elements{
											&asciidoc.String{
												Value: "last",
											},
										},
									},
									Roles:   nil,
									Options: nil,
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
					Title: asciidoc.Elements{
						&asciidoc.String{
							Value: "Subsection",
						},
					},
					Level: 2,
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

var apiTestFindByShouldReturnAnEmptyArrayIfTheIdCriteriaMatchesButTheBlockArgumentYieldsFalse = &asciidoc.Document{
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
										Value: "subsection",
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
							Value: "content",
						},
						&asciidoc.NewLine{},
					},
					Title: asciidoc.Elements{
						&asciidoc.String{
							Value: "Subsection",
						},
					},
					Level: 2,
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

var apiTestFindByShouldNotCrashIfDlistEntryDoesNotHaveDescription = &asciidoc.Document{
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
				"puts 'Hello, World!'",
			},
		},
	},
}

var apiTestCanSubstituteAnExtendedSyntaxHighlighterFactoryImplementationUsingThesyntaxHighlightersOption = &asciidoc.Document{
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
				"puts 'Hello, World!'",
			},
		},
	},
}

var apiTestRenderFileIsAliasedToConvertFile = &asciidoc.Document{
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

var apiTestShouldEmbedRemoteStylesheetByDefaultIfSafeModeIsLessThanSecureAndAllowUriReadIsSet = &asciidoc.Document{
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

var apiTestShouldNotAllowLinkcssBeUnsetFromDocumentIfSafeModeIsSecureOrGreater = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.AttributeReset{
					Name: "linkcss",
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
					Value: "Document Title",
				},
			},
			Level: 0,
		},
	},
}

var apiTestShouldEmbedDefaultStylesheetIfLinkcssIsUnsetFromApiAndSafeModeIsSecureOrGreater = &asciidoc.Document{
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

var apiTestShouldEmbedDefaultStylesheetIfSafeModeIsLessThanSecureAndLinkcssIsUnsetFromApi = &asciidoc.Document{
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

var apiTestShouldLinkToCustomStylesheetIfSpecifiedInStylesheetAttribute = &asciidoc.Document{
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

var apiTestShouldResolveCustomStylesheetRelativeToStylesdir = &asciidoc.Document{
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

var apiTestShouldResolveCustomStylesheetToEmbedRelativeToStylesdir = &asciidoc.Document{
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

var apiTestShouldEmbedCustomStylesheetInRemoteStylesdirIfSafeModeIsLessThanSecureAndAllowUriReadIsSet = &asciidoc.Document{
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

var apiTestShouldRespectOutfilesuffixSoftSetFromApi = &asciidoc.Document{
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
				"puts 'Hello, World!'",
			},
		},
	},
}

var apiTestWithNoAuthor = &asciidoc.Document{
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
					Value: "Getting Real details the business, design, programming, and marketing principles of 37signals.",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Getting Real: The Smarter, Faster, Easier Way to Build a Successful Web Application",
				},
			},
			Level: 0,
		},
	},
}

var apiTestWithOneAuthor = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "David Heinemeier Hansson <",
				},
				asciidoc.Email{
					Address: "david@37signals.com",
				},
				&asciidoc.String{
					Value: ">",
				},
				&asciidoc.NewLine{},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "Getting Real details the business, design, programming, and marketing principles of 37signals.",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Getting Real: The Smarter, Faster, Easier Way to Build a Successful Web Application",
				},
			},
			Level: 0,
		},
	},
}

var apiTestWithTwoAuthors = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "David Heinemeier Hansson <",
				},
				asciidoc.Email{
					Address: "david@37signals.com",
				},
				&asciidoc.String{
					Value: ">; Jason Fried <",
				},
				asciidoc.Email{
					Address: "jason@37signals.com",
				},
				&asciidoc.String{
					Value: ">",
				},
				&asciidoc.NewLine{},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "Getting Real details the business, design, programming, and marketing principles of 37signals.",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Getting Real: The Smarter, Faster, Easier Way to Build a Successful Web Application",
				},
			},
			Level: 0,
		},
	},
}

var apiTestWithAuthorsAsAttributes = &asciidoc.Document{
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
							Value: "David Heinemeier Hansson",
						},
					},
				},
				&asciidoc.AttributeEntry{
					Name: "email_1",
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "david@37signals.com",
						},
					},
				},
				&asciidoc.AttributeEntry{
					Name: "author_2",
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "Jason Fried",
						},
					},
				},
				&asciidoc.AttributeEntry{
					Name: "email_2",
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "jason@37signals.com",
						},
					},
				},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "Getting Real details the business, design, programming, and marketing principles of 37signals.",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "Getting Real: The Smarter, Faster, Easier Way to Build a Successful Web Application",
				},
			},
			Level: 0,
		},
	},
}

var apiTestShouldNotCrashIfNilCellTextIsPassedToCellConstructor = &asciidoc.Document{
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
									Value: "a",
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

var apiTestShouldSetOptionOnNodeWhenSetOptionIsCalled = &asciidoc.Document{
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
					Value: "one",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
	},
}

var apiTestEnabledOptionsShouldReturnAllOptionsWhichAreSet = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "code",
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
			Marker:    "*",
			Checklist: 2,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "test",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     1,
		},
		&asciidoc.UnorderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "profit",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "*",
			Checklist:     1,
		},
	},
}

var apiTestShouldAppendOptionToExistingOptions = &asciidoc.Document{
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
									Value: "fancy",
								},
							},
						},
					},
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
	},
}

var apiTestShouldNotAppendOptionIfOptionIsAlreadySet = &asciidoc.Document{
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
	},
}

var apiTestShouldReturnSetOfOptionNames = &asciidoc.Document{
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
									Value: "compact",
								},
							},
						},
						&asciidoc.ShorthandOption{
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "reversed",
								},
							},
						},
					},
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
	},
}

var apiTestShouldSetLinenumsOptionIfLinenumsEnabledOnSourceBlock = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "[,ruby,linenums]",
		},
		&asciidoc.NewLine{},
		&asciidoc.Listing{
			AttributeList: nil,
			Delimiter: asciidoc.Delimiter{
				Type:   5,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"puts \"Hello, World!\"",
			},
		},
	},
}

var apiTestShouldSetLinenumsOptionIfLinenumsEnabledOnFencedCodeBlock = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "```ruby,linenums",
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

var apiTestShouldNotSetLinenumsAttributeIfLinenumsOptionIsEnabledOnSourceBlock = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Listing{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: nil,
					ID:    nil,
					Roles: nil,
					Options: []*asciidoc.ShorthandOption{
						&asciidoc.ShorthandOption{
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "linenums",
								},
							},
						},
					},
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
				"puts \"Hello, World!\"",
			},
		},
	},
}

var apiTestShouldNotSetLinenumsAttributeIfLinenumsOptionIsEnabledOnFencedCodeBlock = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name:     "source-linenums-option",
			Elements: nil,
		},
		&asciidoc.EmptyLine{
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
	},
}

var apiTestTableColumnShouldNotBeABlockOrInline = &asciidoc.Document{
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
									Value: "a",
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

var apiTestTableCellShouldBeABlock = &asciidoc.Document{
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
									Value: "a",
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

var apiTestNextAdjacentBlockShouldReturnNextBlock = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "first",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "second",
		},
		&asciidoc.NewLine{},
	},
}

var apiTestNextAdjacentBlockShouldReturnNextSiblingOfParentIfCalledOnLastSibling = &asciidoc.Document{
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
					Value: "first",
				},
				&asciidoc.NewLine{},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "second",
		},
		&asciidoc.NewLine{},
	},
}

var apiTestNextAdjacentBlockShouldReturnNextSiblingOfListIfCalledOnLastItem = &asciidoc.Document{
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
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "second",
		},
		&asciidoc.NewLine{},
	},
}

var apiTestNextAdjacentBlockShouldReturnNextItemInDlistIfCalledOnLastBlockOfListItem = &asciidoc.Document{
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
				&asciidoc.ListContinuation{
					ChildElement: &asciidoc.Paragraph{
						AttributeList: nil,
						Elements: asciidoc.Elements{
							&asciidoc.String{
								Value: "more desc",
							},
						},
						Admonition: 0,
					},
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Elements{
				&asciidoc.String{
					Value: "first",
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
					Value: "desc",
				},
			},
			AttributeList: nil,
			Marker:        "::",
			Term: asciidoc.Elements{
				&asciidoc.String{
					Value: "second",
				},
			},
		},
	},
}

var apiTestShouldReturnTrueWhenSectionsIsCalledOnADocumentOrSectionThatHasSections = &asciidoc.Document{
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
									Value: "First subsection",
								},
							},
							Level: 2,
						},
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

var apiTestShouldReturnFalseWhenSectionsIsCalledOnADocumentWithNoSections = &asciidoc.Document{
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

var apiTestShouldReturnFalseWhenSectionsIsCalledOnASectionWithNoSections = &asciidoc.Document{
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

var apiTestShouldReturnFalseWhenSectionsIsCalledOnAnythingThatIsNotASection = &asciidoc.Document{
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
							Value: "Title",
						},
					},
				},
			},
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "I'm not section!",
				},
				&asciidoc.NewLine{},
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
					Value: "I'm not a section either!",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 0,
		},
	},
}
