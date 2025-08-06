package tests

import (
	"testing"

	"github.com/project-chip/alchemy/asciidoc"
)

func TestExtensions(t *testing.T) {
	extensionsTests.run(t)
}

var extensionsTests = parseTests{

	{"should not activate registry if no extension groups are registered", "asciidoctor/extensions_test_should_not_activate_registry_if_no_extension_groups_are_registered.adoc", extensionsTestShouldNotActivateRegistryIfNoExtensionGroupsAreRegistered, nil},

	{"should invoke include processor to process include directive", "asciidoctor/extensions_test_should_invoke_include_processor_to_process_include_directive.adoc", extensionsTestShouldInvokeIncludeProcessorToProcessIncludeDirective, nil},

	{"should invoke include processor if it offers to handle include directive", "asciidoctor/extensions_test_should_invoke_include_processor_if_it_offers_to_handle_include_directive.adoc", extensionsTestShouldInvokeIncludeProcessorIfItOffersToHandleIncludeDirective, nil},

	{"should invoke tree processors after parsing document", "asciidoctor/extensions_test_should_invoke_tree_processors_after_parsing_document.adoc", extensionsTestShouldInvokeTreeProcessorsAfterParsingDocument, nil},

	{"should allow tree processor to replace tree", "asciidoctor/extensions_test_should_allow_tree_processor_to_replace_tree.adoc", extensionsTestShouldAllowTreeProcessorToReplaceTree, nil},

	{"should honor block title assigned in tree processor", "asciidoctor/extensions_test_should_honor_block_title_assigned_in_tree_processor.adoc", extensionsTestShouldHonorBlockTitleAssignedInTreeProcessor, nil},

	{"should invoke postprocessors after converting document", "asciidoctor/extensions_test_should_invoke_postprocessors_after_converting_document.adoc", extensionsTestShouldInvokePostprocessorsAfterConvertingDocument, nil},

	{"should yield to document processor block if block has non-zero arity", "asciidoctor/extensions_test_should_yield_to_document_processor_block_if_block_has_non_zero_arity.adoc", extensionsTestShouldYieldToDocumentProcessorBlockIfBlockHasNonZeroArity, nil},

	{"should invoke processor for custom block", "asciidoctor/extensions_test_should_invoke_processor_for_custom_block.adoc", extensionsTestShouldInvokeProcessorForCustomBlock, nil},

	{"should invoke processor for custom block in an AsciiDoc table cell", "asciidoctor/extensions_test_should_invoke_processor_for_custom_block_in_an_ascii_doc_table_cell.adoc", extensionsTestShouldInvokeProcessorForCustomBlockInAnAsciiDocTableCell, nil},

	{"should yield to syntax processor block if block has non-zero arity", "asciidoctor/extensions_test_should_yield_to_syntax_processor_block_if_block_has_non_zero_arity.adoc", extensionsTestShouldYieldToSyntaxProcessorBlockIfBlockHasNonZeroArity, nil},

	{"should pass cloaked context in attributes passed to process method of custom block", "asciidoctor/extensions_test_should_pass_cloaked_context_in_attributes_passed_to_process_method_of_custom_block.adoc", extensionsTestShouldPassCloakedContextInAttributesPassedToProcessMethodOfCustomBlock, nil},

	{"should allow extension to promote paragraph to compound block", "asciidoctor/extensions_test_should_allow_extension_to_promote_paragraph_to_compound_block.adoc", extensionsTestShouldAllowExtensionToPromoteParagraphToCompoundBlock, nil},

	{"should drop block macro line if target references missing attribute and attribute-missing is drop-line", "asciidoctor/extensions_test_should_drop_block_macro_line_if_target_references_missing_attribute_and_attribute_missing_is_drop_line.adoc", extensionsTestShouldDropBlockMacroLineIfTargetReferencesMissingAttributeAndAttributeMissingIsDropLine, nil},

	{"should invoke processor for custom block macro in an AsciiDoc table cell", "asciidoctor/extensions_test_should_invoke_processor_for_custom_block_macro_in_an_ascii_doc_table_cell.adoc", extensionsTestShouldInvokeProcessorForCustomBlockMacroInAnAsciiDocTableCell, nil},

	{"should fail to convert if name of block macro is illegal", "asciidoctor/extensions_test_should_fail_to_convert_if_name_of_block_macro_is_illegal.adoc", extensionsTestShouldFailToConvertIfNameOfBlockMacroIsIllegal, nil},

	{"should parse text in square brackets as attrlist by default", "asciidoctor/extensions_test_should_parse_text_in_square_brackets_as_attrlist_by_default.adoc", extensionsTestShouldParseTextInSquareBracketsAsAttrlistByDefault, nil},

	{"should prefer attributes parsed from inline macro over default attributes", "asciidoctor/extensions_test_should_prefer_attributes_parsed_from_inline_macro_over_default_attributes.adoc", extensionsTestShouldPreferAttributesParsedFromInlineMacroOverDefaultAttributes, nil},

	{"should not invoke process method or carry over attributes if block processor declares skip content model", "asciidoctor/extensions_test_should_not_invoke_process_method_or_carry_over_attributes_if_block_processor_declares_skip_content_model.adoc", extensionsTestShouldNotInvokeProcessMethodOrCarryOverAttributesIfBlockProcessorDeclaresSkipContentModel, nil},

	{"should pass attributes by value to block processor", "asciidoctor/extensions_test_should_pass_attributes_by_value_to_block_processor.adoc", extensionsTestShouldPassAttributesByValueToBlockProcessor, nil},

	{"should allow extension to replace custom block with a section", "asciidoctor/extensions_test_should_allow_extension_to_replace_custom_block_with_a_section.adoc", extensionsTestShouldAllowExtensionToReplaceCustomBlockWithASection, nil},

	{"can use parse_content to append blocks to current parent", "asciidoctor/extensions_test_can_use_parse_content_to_append_blocks_to_current_parent.adoc", extensionsTestCanUseParseContentToAppendBlocksToCurrentParent, nil},

	{"should ignore return value of custom block if value is parent", "asciidoctor/extensions_test_should_ignore_return_value_of_custom_block_if_value_is_parent.adoc", extensionsTestShouldIgnoreReturnValueOfCustomBlockIfValueIsParent, nil},

	{"should ignore return value of custom block macro if value is parent", "asciidoctor/extensions_test_should_ignore_return_value_of_custom_block_macro_if_value_is_parent.adoc", extensionsTestShouldIgnoreReturnValueOfCustomBlockMacroIfValueIsParent, nil},

	{"parse_content should not share attributes between parsed blocks", "asciidoctor/extensions_test_parse_content_should_not_share_attributes_between_parsed_blocks.adoc", extensionsTestParseContentShouldNotShareAttributesBetweenParsedBlocks, nil},

	{"can use parse_attributes to parse attrlist", "asciidoctor/extensions_test_can_use_parse_attributes_to_parse_attrlist.adoc", extensionsTestCanUseParseAttributesToParseAttrlist, nil},

	{"create_section should set up all section properties", "asciidoctor/extensions_test_create_section_should_set_up_all_section_properties.adoc", extensionsTestCreateSectionShouldSetUpAllSectionProperties, nil},

	{"should add docinfo to document", "asciidoctor/extensions_test_should_add_docinfo_to_document.adoc", extensionsTestShouldAddDocinfoToDocument, nil},

	{"should add multiple docinfo to document", "asciidoctor/extensions_test_should_add_multiple_docinfo_to_document.adoc", extensionsTestShouldAddMultipleDocinfoToDocument, nil},

	{"should not assign caption on image block if title is not set on custom block macro", "asciidoctor/extensions_test_should_not_assign_caption_on_image_block_if_title_is_not_set_on_custom_block_macro.adoc", extensionsTestShouldNotAssignCaptionOnImageBlockIfTitleIsNotSetOnCustomBlockMacro, nil},
}

var extensionsTestShouldNotActivateRegistryIfNoExtensionGroupsAreRegistered = &asciidoc.Document{
	Set: asciidoc.Set{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "junk line",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Set: asciidoc.Set{
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "sample content",
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

var extensionsTestShouldInvokeIncludeProcessorToProcessIncludeDirective = &asciidoc.Document{
	Set: asciidoc.Set{
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
		&asciidoc.FileInclude{
			AttributeList: nil,
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "lorem-ipsum.txt",
				},
			},
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

var extensionsTestShouldInvokeIncludeProcessorIfItOffersToHandleIncludeDirective = &asciidoc.Document{
	Set: asciidoc.Set{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.FileInclude{
			AttributeList: nil,
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "skip-me.adoc",
				},
			},
		},
		&asciidoc.String{
			Value: "line after skip",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.FileInclude{
			AttributeList: nil,
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "include-file.adoc",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.FileInclude{
			AttributeList: nil,
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "fixtures/grandchild-include.adoc",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "last line",
		},
		&asciidoc.NewLine{},
	},
}

var extensionsTestShouldInvokeTreeProcessorsAfterParsingDocument = &asciidoc.Document{
	Set: asciidoc.Set{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Doc Writer",
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
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Document Title",
				},
			},
			Level: 0,
		},
	},
}

var extensionsTestShouldAllowTreeProcessorToReplaceTree = &asciidoc.Document{
	Set: asciidoc.Set{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Doc Writer",
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
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Original Document",
				},
			},
			Level: 0,
		},
	},
}

var extensionsTestShouldHonorBlockTitleAssignedInTreeProcessor = &asciidoc.Document{
	Set: asciidoc.Set{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Set: asciidoc.Set{
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
						&asciidoc.TitleAttribute{
							Val: asciidoc.Set{
								&asciidoc.String{
									Value: "Old block title",
								},
							},
						},
					},
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "example block content",
						},
						&asciidoc.NewLine{},
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

var extensionsTestShouldInvokePostprocessorsAfterConvertingDocument = &asciidoc.Document{
	Set: asciidoc.Set{
		&asciidoc.EmptyLine{
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

var extensionsTestShouldYieldToDocumentProcessorBlockIfBlockHasNonZeroArity = &asciidoc.Document{
	Set: asciidoc.Set{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "hi!",
		},
		&asciidoc.NewLine{},
	},
}

var extensionsTestShouldInvokeProcessorForCustomBlock = &asciidoc.Document{
	Set: asciidoc.Set{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Paragraph{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Set: asciidoc.Set{
							&asciidoc.String{
								Value: "yell",
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
					Value: "Hi there!",
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
					Style: &asciidoc.ShorthandStyle{
						Set: asciidoc.Set{
							&asciidoc.String{
								Value: "yell",
							},
						},
					},
					ID:      nil,
					Roles:   nil,
					Options: nil,
				},
				&asciidoc.NamedAttribute{
					Name: "chars",
					Val: asciidoc.Set{
						&asciidoc.String{
							Value: "aeiou",
						},
					},
					Quote: 0,
				},
			},
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "Hi there!",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 0,
		},
	},
}

var extensionsTestShouldInvokeProcessorForCustomBlockInAnAsciiDocTableCell = &asciidoc.Document{
	Set: asciidoc.Set{
		&asciidoc.EmptyLine{
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
								&asciidoc.Paragraph{
									AttributeList: asciidoc.AttributeList{
										&asciidoc.ShorthandAttribute{
											Style: &asciidoc.ShorthandStyle{
												Set: asciidoc.Set{
													&asciidoc.String{
														Value: "yell",
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
											Value: "Hi there!",
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

var extensionsTestShouldYieldToSyntaxProcessorBlockIfBlockHasNonZeroArity = &asciidoc.Document{
	Set: asciidoc.Set{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.LiteralBlock{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Set: asciidoc.Set{
							&asciidoc.String{
								Value: "eval",
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
				"'yolo' * 5",
			},
		},
	},
}

var extensionsTestShouldPassCloakedContextInAttributesPassedToProcessMethodOfCustomBlock = &asciidoc.Document{
	Set: asciidoc.Set{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.SidebarBlock{
			Delimiter: asciidoc.Delimiter{
				Type:   8,
				Length: 4,
			},
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Set: asciidoc.Set{
							&asciidoc.String{
								Value: "custom",
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
					Value: "sidebar",
				},
				&asciidoc.NewLine{},
			},
		},
	},
}

var extensionsTestShouldAllowExtensionToPromoteParagraphToCompoundBlock = &asciidoc.Document{
	Set: asciidoc.Set{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Paragraph{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Set: asciidoc.Set{
							&asciidoc.String{
								Value: "ex",
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
					Value: "example",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 0,
		},
	},
}

var extensionsTestShouldDropBlockMacroLineIfTargetReferencesMissingAttributeAndAttributeMissingIsDropLine = &asciidoc.Document{
	Set: asciidoc.Set{
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
					Value: "snippet::",
				},
				&asciidoc.UserAttributeReference{
					Value: "gist-ns",
				},
				&asciidoc.String{
					Value: "12345[mode=edit]",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 0,
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "following paragraph",
		},
		&asciidoc.NewLine{},
	},
}

var extensionsTestShouldInvokeProcessorForCustomBlockMacroInAnAsciiDocTableCell = &asciidoc.Document{
	Set: asciidoc.Set{
		&asciidoc.EmptyLine{
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
								&asciidoc.String{
									Value: "message::hi[]",
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

var extensionsTestShouldFailToConvertIfNameOfBlockMacroIsIllegal = &asciidoc.Document{
	Set: asciidoc.Set{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "attribute::yin[yang]",
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "header_attribute::foo[bar]",
		},
		&asciidoc.NewLine{},
	},
}

var extensionsTestShouldParseTextInSquareBracketsAsAttrlistByDefault = &asciidoc.Document{
	Set: asciidoc.Set{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.StemBlock{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.NamedAttribute{
					Name: "subs",
					Val: asciidoc.Set{
						&asciidoc.String{
							Value: "normal",
						},
					},
					Quote: 0,
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   9,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"short_attributes:[]",
				"short_attributes:[value,key=val]",
				"short_text:[]",
				"short_text:[[text\\]]",
				"full-attributes:target[]",
				"full-attributes:target[value,key=val]",
				"full-text:target[]",
				"full-text:target[[text\\]]",
				"@target",
			},
		},
	},
}

var extensionsTestShouldPreferAttributesParsedFromInlineMacroOverDefaultAttributes = &asciidoc.Document{
	Set: asciidoc.Set{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Paragraph{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.TitleAttribute{
					Val: asciidoc.Set{
						&asciidoc.String{
							Value: "unused title",
						},
					},
				},
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Set: asciidoc.Set{
							&asciidoc.String{
								Value: "skip-me",
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
					Value: "not shown",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 0,
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OpenBlock{
			AttributeList: nil,
			Delimiter: asciidoc.Delimiter{
				Type:   7,
				Length: 2,
			},
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "shown",
				},
				&asciidoc.NewLine{},
			},
		},
	},
}

var extensionsTestShouldNotInvokeProcessMethodOrCarryOverAttributesIfBlockProcessorDeclaresSkipContentModel = &asciidoc.Document{
	Set: asciidoc.Set{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Paragraph{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.TitleAttribute{
					Val: asciidoc.Set{
						&asciidoc.String{
							Value: "unused title",
						},
					},
				},
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Set: asciidoc.Set{
							&asciidoc.String{
								Value: "ignore",
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
					Value: "not shown",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 0,
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OpenBlock{
			AttributeList: nil,
			Delimiter: asciidoc.Delimiter{
				Type:   7,
				Length: 2,
			},
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "shown",
				},
				&asciidoc.NewLine{},
			},
		},
	},
}

var extensionsTestShouldPassAttributesByValueToBlockProcessor = &asciidoc.Document{
	Set: asciidoc.Set{
		&asciidoc.EmptyLine{
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
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Set: asciidoc.Set{
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

var extensionsTestShouldAllowExtensionToReplaceCustomBlockWithASection = &asciidoc.Document{
	Set: asciidoc.Set{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OpenBlock{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.TitleAttribute{
					Val: asciidoc.Set{
						&asciidoc.String{
							Value: "Section Title",
						},
					},
				},
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Set: asciidoc.Set{
							&asciidoc.String{
								Value: "sect",
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
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "a",
				},
				&asciidoc.NewLine{},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "b",
				},
				&asciidoc.NewLine{},
			},
		},
	},
}

var extensionsTestCanUseParseContentToAppendBlocksToCurrentParent = &asciidoc.Document{
	Set: asciidoc.Set{
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
		&asciidoc.LiteralBlock{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Set: asciidoc.Set{
							&asciidoc.String{
								Value: "csv",
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
				"a,b,c",
			},
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

var extensionsTestShouldIgnoreReturnValueOfCustomBlockIfValueIsParent = &asciidoc.Document{
	Set: asciidoc.Set{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OpenBlock{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Set: asciidoc.Set{
							&asciidoc.String{
								Value: "unwrap",
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
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "a",
				},
				&asciidoc.NewLine{},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "b",
				},
				&asciidoc.NewLine{},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "c",
				},
				&asciidoc.NewLine{},
			},
		},
	},
}

var extensionsTestShouldIgnoreReturnValueOfCustomBlockMacroIfValueIsParent = &asciidoc.Document{
	Set: asciidoc.Set{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "para::text[]",
		},
		&asciidoc.NewLine{},
	},
}

var extensionsTestParseContentShouldNotShareAttributesBetweenParsedBlocks = &asciidoc.Document{
	Set: asciidoc.Set{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OpenBlock{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Set: asciidoc.Set{
							&asciidoc.String{
								Value: "wrap",
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
			Set: asciidoc.Set{
				&asciidoc.ExampleBlock{
					Delimiter: asciidoc.Delimiter{
						Type:   3,
						Length: 4,
					},
					AttributeList: asciidoc.AttributeList{
						&asciidoc.NamedAttribute{
							Name: "foo",
							Val: asciidoc.Set{
								&asciidoc.String{
									Value: "bar",
								},
							},
							Quote: 0,
						},
					},
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "content",
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
						&asciidoc.NamedAttribute{
							Name: "baz",
							Val: asciidoc.Set{
								&asciidoc.String{
									Value: "qux",
								},
							},
							Quote: 0,
						},
					},
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "content",
						},
						&asciidoc.NewLine{},
					},
				},
			},
		},
	},
}

var extensionsTestCanUseParseAttributesToParseAttrlist = &asciidoc.Document{
	Set: asciidoc.Set{
		&asciidoc.EmptyLine{
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
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OpenBlock{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.ShorthandAttribute{
					Style: &asciidoc.ShorthandStyle{
						Set: asciidoc.Set{
							&asciidoc.String{
								Value: "attrs",
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
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "a,b,c,key=val",
				},
				&asciidoc.NewLine{},
			},
		},
	},
}

var extensionsTestCreateSectionShouldSetUpAllSectionProperties = &asciidoc.Document{
	Set: asciidoc.Set{
		&asciidoc.EmptyLine{
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
				&asciidoc.AttributeEntry{
					Name: "sectnums",
					Set:  nil,
				},
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "sect::[%s]",
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

var extensionsTestShouldAddDocinfoToDocument = &asciidoc.Document{
	Set: asciidoc.Set{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Set: asciidoc.Set{
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "sample content",
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

var extensionsTestShouldAddMultipleDocinfoToDocument = &asciidoc.Document{
	Set: asciidoc.Set{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Set: asciidoc.Set{
				&asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "sample content",
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

var extensionsTestShouldNotAssignCaptionOnImageBlockIfTitleIsNotSetOnCustomBlockMacro = &asciidoc.Document{
	Set: asciidoc.Set{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Paragraph{
			AttributeList: asciidoc.AttributeList{
				&asciidoc.TitleAttribute{
					Val: asciidoc.Set{
						&asciidoc.String{
							Value: "Cat in Sink?",
						},
					},
				},
			},
			Set: asciidoc.Set{
				&asciidoc.String{
					Value: "cat_in_sink::30[]",
				},
				&asciidoc.NewLine{},
			},
			Admonition: 0,
		},
	},
}
