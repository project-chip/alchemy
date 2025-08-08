package tests

import (
	"testing"

	"github.com/project-chip/alchemy/asciidoc"
)

func TestSyntaxHighlighter(t *testing.T) {
	syntaxHighlighterTests.run(t)
}

var syntaxHighlighterTests = parseTests{

	{"should set syntax_highlighter property on document if source highlighter is set and basebackend is html", "asciidoctor/syntax_highlighter_test_should_set_syntax_highlighter_property_on_document_if_source_highlighter_is_set_and_basebackend_is_html.adoc", syntaxHighlighterTestShouldSetSyntaxHighlighterPropertyOnDocumentIfSourceHighlighterIsSetAndBasebackendIsHtml, nil},

	{"should not set syntax_highlighter property on document if source highlighter is set and basebackend is not html", "asciidoctor/syntax_highlighter_test_should_not_set_syntax_highlighter_property_on_document_if_source_highlighter_is_set_and_basebackend_is_not_html.adoc", syntaxHighlighterTestShouldNotSetSyntaxHighlighterPropertyOnDocumentIfSourceHighlighterIsSetAndBasebackendIsNotHtml, nil},

	{"should not set syntax_highlighter property on document if source highlighter is not set", "asciidoctor/syntax_highlighter_test_should_not_set_syntax_highlighter_property_on_document_if_source_highlighter_is_not_set.adoc", syntaxHighlighterTestShouldNotSetSyntaxHighlighterPropertyOnDocumentIfSourceHighlighterIsNotSet, nil},

	{"should not set syntax_highlighter property on document if syntax highlighter cannot be resolved", "asciidoctor/syntax_highlighter_test_should_not_set_syntax_highlighter_property_on_document_if_syntax_highlighter_cannot_be_resolved.adoc", syntaxHighlighterTestShouldNotSetSyntaxHighlighterPropertyOnDocumentIfSyntaxHighlighterCannotBeResolved, nil},

	{"should set language on output of source block when source-highlighter attribute is not set", "asciidoctor/syntax_highlighter_test_should_set_language_on_output_of_source_block_when_source_highlighter_attribute_is_not_set.adoc", syntaxHighlighterTestShouldSetLanguageOnOutputOfSourceBlockWhenSourceHighlighterAttributeIsNotSet, nil},

	{"should set language on output of source block when source-highlighter attribute is not recognized", "asciidoctor/syntax_highlighter_test_should_set_language_on_output_of_source_block_when_source_highlighter_attribute_is_not_recognized.adoc", syntaxHighlighterTestShouldSetLanguageOnOutputOfSourceBlockWhenSourceHighlighterAttributeIsNotRecognized, nil},

	{"should highlight source if source highlighter is set even if language is not set", "asciidoctor/syntax_highlighter_test_should_highlight_source_if_source_highlighter_is_set_even_if_language_is_not_set.adoc", syntaxHighlighterTestShouldHighlightSourceIfSourceHighlighterIsSetEvenIfLanguageIsNotSet, nil},

	{"should not crash if source block has no lines and source highlighter is set", "asciidoctor/syntax_highlighter_test_should_not_crash_if_source_block_has_no_lines_and_source_highlighter_is_set.adoc", syntaxHighlighterTestShouldNotCrashIfSourceBlockHasNoLinesAndSourceHighlighterIsSet, nil},

	{"should highlight source inside AsciiDoc table cell if source-highlighter attribute is set", "asciidoctor/syntax_highlighter_test_should_highlight_source_inside_ascii_doc_table_cell_if_source_highlighter_attribute_is_set.adoc", syntaxHighlighterTestShouldHighlightSourceInsideAsciiDocTableCellIfSourceHighlighterAttributeIsSet, nil},

	{"should set starting line number in DocBook output if linenums option is enabled and start attribute is set", "asciidoctor/syntax_highlighter_test_should_set_starting_line_number_in_doc_book_output_if_linenums_option_is_enabled_and_start_attribute_is_set.adoc", syntaxHighlighterTestShouldSetStartingLineNumberInDocBookOutputIfLinenumsOptionIsEnabledAndStartAttributeIsSet, nil},

	{"should read source language from source-language document attribute if not specified on source block", "asciidoctor/syntax_highlighter_test_should_read_source_language_from_source_language_document_attribute_if_not_specified_on_source_block.adoc", syntaxHighlighterTestShouldReadSourceLanguageFromSourceLanguageDocumentAttributeIfNotSpecifiedOnSourceBlock, nil},

	{"should rename document attribute named language to source-language when compat-mode is enabled", "asciidoctor/syntax_highlighter_test_should_rename_document_attribute_named_language_to_source_language_when_compat_mode_is_enabled.adoc", syntaxHighlighterTestShouldRenameDocumentAttributeNamedLanguageToSourceLanguageWhenCompatModeIsEnabled, nil},

	{"should highlight source if source-highlighter attribute is set", "asciidoctor/syntax_highlighter_test_should_highlight_source_if_source_highlighter_attribute_is_set.adoc", syntaxHighlighterTestShouldHighlightSourceIfSourceHighlighterAttributeIsSet, nil},

	{"should not fail if source language is invalid", "asciidoctor/syntax_highlighter_test_should_not_fail_if_source_language_is_invalid.adoc", syntaxHighlighterTestShouldNotFailIfSourceLanguageIsInvalid, nil},

	{"should number lines if third positional attribute is set", "asciidoctor/syntax_highlighter_test_should_number_lines_if_third_positional_attribute_is_set.adoc", syntaxHighlighterTestShouldNumberLinesIfThirdPositionalAttributeIsSet, nil},

	{"should number lines if linenums option is set on source block", "asciidoctor/syntax_highlighter_test_should_number_lines_if_linenums_option_is_set_on_source_block.adoc", syntaxHighlighterTestShouldNumberLinesIfLinenumsOptionIsSetOnSourceBlock, nil},

	{"should number lines of source block if source-linenums-option document attribute is set", "asciidoctor/syntax_highlighter_test_should_number_lines_of_source_block_if_source_linenums_option_document_attribute_is_set.adoc", syntaxHighlighterTestShouldNumberLinesOfSourceBlockIfSourceLinenumsOptionDocumentAttributeIsSet, nil},

	{"should set starting line number in HTML output if linenums option is enabled and start attribute is set", "asciidoctor/syntax_highlighter_test_should_set_starting_line_number_in_html_output_if_linenums_option_is_enabled_and_start_attribute_is_set.adoc", syntaxHighlighterTestShouldSetStartingLineNumberInHtmlOutputIfLinenumsOptionIsEnabledAndStartAttributeIsSet, nil},

	{"should replace callout marks but not highlight them if source-highlighter attribute is coderay", "asciidoctor/syntax_highlighter_test_should_replace_callout_marks_but_not_highlight_them_if_source_highlighter_attribute_is_coderay.adoc", syntaxHighlighterTestShouldReplaceCalloutMarksButNotHighlightThemIfSourceHighlighterAttributeIsCoderay, nil},

	{"should support autonumbered callout marks if source-highlighter attribute is coderay", "asciidoctor/syntax_highlighter_test_should_support_autonumbered_callout_marks_if_source_highlighter_attribute_is_coderay.adoc", syntaxHighlighterTestShouldSupportAutonumberedCalloutMarksIfSourceHighlighterAttributeIsCoderay, nil},

	{"should restore callout marks to correct lines if source highlighter is coderay and table line numbering is enabled", "asciidoctor/syntax_highlighter_test_should_restore_callout_marks_to_correct_lines_if_source_highlighter_is_coderay_and_table_line_numbering_is_enabled.adoc", syntaxHighlighterTestShouldRestoreCalloutMarksToCorrectLinesIfSourceHighlighterIsCoderayAndTableLineNumberingIsEnabled, nil},

	{"should restore isolated callout mark on last line of source when source highlighter is coderay", "asciidoctor/syntax_highlighter_test_should_restore_isolated_callout_mark_on_last_line_of_source_when_source_highlighter_is_coderay.adoc", syntaxHighlighterTestShouldRestoreIsolatedCalloutMarkOnLastLineOfSourceWhenSourceHighlighterIsCoderay, nil},

	{"should preserve passthrough placeholders when highlighting source using coderay", "asciidoctor/syntax_highlighter_test_should_preserve_passthrough_placeholders_when_highlighting_source_using_coderay.adoc", syntaxHighlighterTestShouldPreservePassthroughPlaceholdersWhenHighlightingSourceUsingCoderay, nil},

	{"should link to CodeRay stylesheet if source-highlighter is coderay and linkcss is set", "asciidoctor/syntax_highlighter_test_should_link_to_code_ray_stylesheet_if_source_highlighter_is_coderay_and_linkcss_is_set.adoc", syntaxHighlighterTestShouldLinkToCodeRayStylesheetIfSourceHighlighterIsCoderayAndLinkcssIsSet, nil},

	{"should highlight source inline if source-highlighter attribute is coderay and coderay-css is style", "asciidoctor/syntax_highlighter_test_should_highlight_source_inline_if_source_highlighter_attribute_is_coderay_and_coderay_css_is_style.adoc", syntaxHighlighterTestShouldHighlightSourceInlineIfSourceHighlighterAttributeIsCoderayAndCoderayCssIsStyle, nil},

	{"should read stylesheet", "asciidoctor/syntax_highlighter_test_should_read_stylesheet.adoc", syntaxHighlighterTestShouldReadStylesheet, nil},

	{"should include remote highlight.js assets if source-highlighter attribute is highlight.js", "asciidoctor/syntax_highlighter_test_should_include_remote_highlight_js_assets_if_source_highlighter_attribute_is_highlight_js.adoc", syntaxHighlighterTestShouldIncludeRemoteHighlightJsAssetsIfSourceHighlighterAttributeIsHighlightJs, nil},

	{"should add language-none class to source block when source-highlighter is highlight.js and language is not set", "asciidoctor/syntax_highlighter_test_should_add_language_none_class_to_source_block_when_source_highlighter_is_highlight_js_and_language_is_not_set.adoc", syntaxHighlighterTestShouldAddLanguageNoneClassToSourceBlockWhenSourceHighlighterIsHighlightJsAndLanguageIsNotSet, nil},

	{"should load additional languages specified by highlightjs-languages", "asciidoctor/syntax_highlighter_test_should_load_additional_languages_specified_by_highlightjs_languages.adoc", syntaxHighlighterTestShouldLoadAdditionalLanguagesSpecifiedByHighlightjsLanguages, nil},

	{"should add language classes to child code element when source-highlighter is prettify", "asciidoctor/syntax_highlighter_test_should_add_language_classes_to_child_code_element_when_source_highlighter_is_prettify.adoc", syntaxHighlighterTestShouldAddLanguageClassesToChildCodeElementWhenSourceHighlighterIsPrettify, nil},

	{"should set linenums start if linenums are enabled and start attribute is set when source-highlighter is prettify", "asciidoctor/syntax_highlighter_test_should_set_linenums_start_if_linenums_are_enabled_and_start_attribute_is_set_when_source_highlighter_is_prettify.adoc", syntaxHighlighterTestShouldSetLinenumsStartIfLinenumsAreEnabledAndStartAttributeIsSetWhenSourceHighlighterIsPrettify, nil},

	{"should set lang attribute on pre when source-highlighter is html-pipeline", "asciidoctor/syntax_highlighter_test_should_set_lang_attribute_on_pre_when_source_highlighter_is_html_pipeline.adoc", syntaxHighlighterTestShouldSetLangAttributeOnPreWhenSourceHighlighterIsHtmlPipeline, nil},

	{"should syntax highlight source if source-highlighter attribute is set", "asciidoctor/syntax_highlighter_test_should_syntax_highlight_source_if_source_highlighter_attribute_is_set.adoc", syntaxHighlighterTestShouldSyntaxHighlightSourceIfSourceHighlighterAttributeIsSet, nil},

	{"should highlight source using a mixed lexer (HTML + JavaScript)", "asciidoctor/syntax_highlighter_test_should_highlight_source_using_a_mixed_lexer_(html_+_java_script).adoc", syntaxHighlighterTestShouldHighlightSourceUsingAMixedLexerHtmlJavaScript, nil},

	{"should enable start_inline for PHP by default", "asciidoctor/syntax_highlighter_test_should_enable_start_inline_for_php_by_default.adoc", syntaxHighlighterTestShouldEnableStartInlineForPhpByDefault, nil},

	{"should not enable start_inline for PHP if disabled using cgi-style option on language", "asciidoctor/syntax_highlighter_test_should_not_enable_start_inline_for_php_if_disabled_using_cgi_style_option_on_language.adoc", syntaxHighlighterTestShouldNotEnableStartInlineForPhpIfDisabledUsingCgiStyleOptionOnLanguage, nil},

	{"should not enable start_inline for PHP if mixed option is set", "asciidoctor/syntax_highlighter_test_should_not_enable_start_inline_for_php_if_mixed_option_is_set.adoc", syntaxHighlighterTestShouldNotEnableStartInlineForPhpIfMixedOptionIsSet, nil},

	{"should not crash if source-highlighter attribute is set and source block does not define a language", "asciidoctor/syntax_highlighter_test_should_not_crash_if_source_highlighter_attribute_is_set_and_source_block_does_not_define_a_language.adoc", syntaxHighlighterTestShouldNotCrashIfSourceHighlighterAttributeIsSetAndSourceBlockDoesNotDefineALanguage, nil},

	{"should default to plain text lexer if lexer cannot be resolved for language", "asciidoctor/syntax_highlighter_test_should_default_to_plain_text_lexer_if_lexer_cannot_be_resolved_for_language.adoc", syntaxHighlighterTestShouldDefaultToPlainTextLexerIfLexerCannotBeResolvedForLanguage, nil},

	{"should number lines using table layout if linenums option is enabled and linenums mode is not set", "asciidoctor/syntax_highlighter_test_should_number_lines_using_table_layout_if_linenums_option_is_enabled_and_linenums_mode_is_not_set.adoc", syntaxHighlighterTestShouldNumberLinesUsingTableLayoutIfLinenumsOptionIsEnabledAndLinenumsModeIsNotSet, nil},

	{"should number lines using inline element if linenums option is enabled and linenums mode is inline", "asciidoctor/syntax_highlighter_test_should_number_lines_using_inline_element_if_linenums_option_is_enabled_and_linenums_mode_is_inline.adoc", syntaxHighlighterTestShouldNumberLinesUsingInlineElementIfLinenumsOptionIsEnabledAndLinenumsModeIsInline, nil},

	{"should gracefully fallback to default style if specified style not recognized", "asciidoctor/syntax_highlighter_test_should_gracefully_fallback_to_default_style_if_specified_style_not_recognized.adoc", syntaxHighlighterTestShouldGracefullyFallbackToDefaultStyleIfSpecifiedStyleNotRecognized, nil},

	{"should restore isolated callout mark on last line of source", "asciidoctor/syntax_highlighter_test_should_restore_isolated_callout_mark_on_last_line_of_source.adoc", syntaxHighlighterTestShouldRestoreIsolatedCalloutMarkOnLastLineOfSource, nil},

	{"should number all lines when isolated callout mark is on last line of source and starting line number is set", "asciidoctor/syntax_highlighter_test_should_number_all_lines_when_isolated_callout_mark_is_on_last_line_of_source_and_starting_line_number_is_set.adoc", syntaxHighlighterTestShouldNumberAllLinesWhenIsolatedCalloutMarkIsOnLastLineOfSourceAndStartingLineNumberIsSet, nil},

	{"should preserve guard in front of callout if icons are not enabled", "asciidoctor/syntax_highlighter_test_should_preserve_guard_in_front_of_callout_if_icons_are_not_enabled.adoc", syntaxHighlighterTestShouldPreserveGuardInFrontOfCalloutIfIconsAreNotEnabled, nil},

	{"should preserve guard around callout if icons are not enabled", "asciidoctor/syntax_highlighter_test_should_preserve_guard_around_callout_if_icons_are_not_enabled.adoc", syntaxHighlighterTestShouldPreserveGuardAroundCalloutIfIconsAreNotEnabled, nil},

	{"should read stylesheet for specified style", "asciidoctor/syntax_highlighter_test_should_read_stylesheet_for_specified_style.adoc", syntaxHighlighterTestShouldReadStylesheetForSpecifiedStyle, nil},

	{"should embed stylesheet for pygments style", "asciidoctor/syntax_highlighter_test_should_embed_stylesheet_for_pygments_style.adoc", syntaxHighlighterTestShouldEmbedStylesheetForPygmentsStyle, nil},

	{"should restore callout marks to correct lines if table line numbering is enabled", "asciidoctor/syntax_highlighter_test_should_restore_callout_marks_to_correct_lines_if_table_line_numbering_is_enabled.adoc", syntaxHighlighterTestShouldRestoreCalloutMarksToCorrectLinesIfTableLineNumberingIsEnabled, nil},

	{"should not hardcode inline styles on lineno div and pre elements when linenums are enabled in table mode", "asciidoctor/syntax_highlighter_test_should_not_hardcode_inline_styles_on_lineno_div_and_pre_elements_when_linenums_are_enabled_in_table_mode.adoc", syntaxHighlighterTestShouldNotHardcodeInlineStylesOnLinenoDivAndPreElementsWhenLinenumsAreEnabledInTableMode, nil},

	{"should line highlight specified lines", "asciidoctor/syntax_highlighter_test_should_line_highlight_specified_lines.adoc", syntaxHighlighterTestShouldLineHighlightSpecifiedLines, nil},
}

var syntaxHighlighterTestShouldSetSyntaxHighlighterPropertyOnDocumentIfSourceHighlighterIsSetAndBasebackendIsHtml = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "source-highlighter",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "coderay",
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
				"puts 'Hello, World!'",
			},
		},
	},
}

var syntaxHighlighterTestShouldNotSetSyntaxHighlighterPropertyOnDocumentIfSourceHighlighterIsSetAndBasebackendIsNotHtml = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "source-highlighter",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "coderay",
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
				"puts 'Hello, World!'",
			},
		},
	},
}

var syntaxHighlighterTestShouldNotSetSyntaxHighlighterPropertyOnDocumentIfSourceHighlighterIsNotSet = &asciidoc.Document{
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

var syntaxHighlighterTestShouldNotSetSyntaxHighlighterPropertyOnDocumentIfSyntaxHighlighterCannotBeResolved = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "source-highlighter",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "unknown",
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
				"puts 'Hello, World!'",
			},
		},
	},
}

var syntaxHighlighterTestShouldSetLanguageOnOutputOfSourceBlockWhenSourceHighlighterAttributeIsNotSet = &asciidoc.Document{
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

var syntaxHighlighterTestShouldSetLanguageOnOutputOfSourceBlockWhenSourceHighlighterAttributeIsNotRecognized = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "source-highlighter",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "unknown",
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
				"puts 'Hello, World!'",
			},
		},
	},
}

var syntaxHighlighterTestShouldHighlightSourceIfSourceHighlighterIsSetEvenIfLanguageIsNotSet = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "source-highlighter",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "coderay",
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
			},
			Delimiter: asciidoc.Delimiter{
				Type:   5,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"[numbers]",
				"one",
				"two",
				"three",
			},
		},
	},
}

var syntaxHighlighterTestShouldNotCrashIfSourceBlockHasNoLinesAndSourceHighlighterIsSet = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "source-highlighter",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "coderay",
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
							Value: "text",
						},
					},
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   5,
				Length: 4,
			},
			LineList: asciidoc.LineList{},
		},
	},
}

var syntaxHighlighterTestShouldHighlightSourceInsideAsciiDocTableCellIfSourceHighlighterAttributeIsSet = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "source-highlighter",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "coderay",
				},
			},
		},
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
										"require 'coderay'",
										"",
										"html = CodeRay.scan(\"puts 'Hello, world!'\", :ruby).div line_numbers: :table",
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

var syntaxHighlighterTestShouldSetStartingLineNumberInDocBookOutputIfLinenumsOptionIsEnabledAndStartAttributeIsSet = &asciidoc.Document{
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
							Value: "java",
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
			Delimiter: asciidoc.Delimiter{
				Type:   5,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"public class HelloWorld {",
				"  public static void main(String[] args) {",
				"    out.println(\"Hello, World!\");",
				"  }",
				"}",
			},
		},
	},
}

var syntaxHighlighterTestShouldReadSourceLanguageFromSourceLanguageDocumentAttributeIfNotSpecifiedOnSourceBlock = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "source-highlighter",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "coderay",
				},
			},
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
				"require 'coderay'",
				"",
				"html = CodeRay.scan(\"puts 'Hello, world!'\", :ruby).div line_numbers: :table",
			},
		},
	},
}

var syntaxHighlighterTestShouldRenameDocumentAttributeNamedLanguageToSourceLanguageWhenCompatModeIsEnabled = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "language",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "ruby",
				},
			},
		},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.UserAttributeReference{
			Value: "source-language",
		},
		&asciidoc.NewLine{},
	},
}

var syntaxHighlighterTestShouldHighlightSourceIfSourceHighlighterAttributeIsSet = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "source-highlighter",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "coderay",
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
				"require 'coderay'",
				"",
				"html = CodeRay.scan(\"puts 'Hello, world!'\", :ruby).div line_numbers: :table",
			},
		},
	},
}

var syntaxHighlighterTestShouldNotFailIfSourceLanguageIsInvalid = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "source-highlighter",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "coderay",
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
							Value: "n/a",
						},
					},
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   5,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"PRINT 'yo'",
			},
		},
	},
}

var syntaxHighlighterTestShouldNumberLinesIfThirdPositionalAttributeIsSet = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "source-highlighter",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "coderay",
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
				&asciidoc.PositionalAttribute{
					Offset:      2,
					ImpliedName: "",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "linenums",
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

var syntaxHighlighterTestShouldNumberLinesIfLinenumsOptionIsSetOnSourceBlock = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "source-highlighter",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "coderay",
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
				"puts 'Hello, World!'",
			},
		},
	},
}

var syntaxHighlighterTestShouldNumberLinesOfSourceBlockIfSourceLinenumsOptionDocumentAttributeIsSet = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "source-highlighter",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "coderay",
				},
			},
		},
		&asciidoc.AttributeEntry{
			Name:     "source-linenums-option",
			Elements: nil,
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
				"puts 'Hello, World!'",
			},
		},
	},
}

var syntaxHighlighterTestShouldSetStartingLineNumberInHtmlOutputIfLinenumsOptionIsEnabledAndStartAttributeIsSet = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "source-highlighter",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "coderay",
				},
			},
		},
		&asciidoc.AttributeEntry{
			Name: "coderay-linenums-mode",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "inline",
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
				&asciidoc.NamedAttribute{
					Name: "start",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "10",
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
				"puts 'Hello, World!'",
			},
		},
	},
}

var syntaxHighlighterTestShouldReplaceCalloutMarksButNotHighlightThemIfSourceHighlighterAttributeIsCoderay = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "source-highlighter",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "coderay",
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
				"require 'coderay' # <1>",
				"",
				"html = CodeRay.scan(\"puts 'Hello, world!'\", :ruby).div line_numbers: :table # <2>",
				"puts html # <3> <4>",
				"exit 0 # <5><6>",
			},
		},
		&asciidoc.String{
			Value: "<1> Load library",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "<2> Highlight source",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "<3> Print to stdout",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "<4> Redirect to a file to capture output",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "<5> Exit program",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "<6> Reports success",
		},
		&asciidoc.NewLine{},
	},
}

var syntaxHighlighterTestShouldSupportAutonumberedCalloutMarksIfSourceHighlighterAttributeIsCoderay = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "source-highlighter",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "coderay",
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
				"require 'coderay' # <.><.>",
				"",
				"html = CodeRay.scan(\"puts 'Hello, world!'\", :ruby).div line_numbers: :table # <.>",
				"puts html # <.>",
			},
		},
		&asciidoc.String{
			Value: "<.> Load library",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "<.> Gem must be installed",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "<.> Highlight source",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "<.> Print to stdout",
		},
		&asciidoc.NewLine{},
	},
}

var syntaxHighlighterTestShouldRestoreCalloutMarksToCorrectLinesIfSourceHighlighterIsCoderayAndTableLineNumberingIsEnabled = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "source-highlighter",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "coderay",
				},
			},
		},
		&asciidoc.AttributeEntry{
			Name: "coderay-linenums-mode",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "table",
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
				&asciidoc.PositionalAttribute{
					Offset:      2,
					ImpliedName: "",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "numbered",
						},
					},
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   5,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"require 'coderay' # <1>",
				"",
				"html = CodeRay.scan(\"puts 'Hello, world!'\", :ruby).div line_numbers: :table # <2>",
				"puts html # <3> <4>",
				"exit 0 # <5><6>",
			},
		},
		&asciidoc.String{
			Value: "<1> Load library",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "<2> Highlight source",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "<3> Print to stdout",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "<4> Redirect to a file to capture output",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "<5> Exit program",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "<6> Reports success",
		},
		&asciidoc.NewLine{},
	},
}

var syntaxHighlighterTestShouldRestoreIsolatedCalloutMarkOnLastLineOfSourceWhenSourceHighlighterIsCoderay = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "source-highlighter",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "coderay",
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
				&asciidoc.PositionalAttribute{
					Offset:      2,
					ImpliedName: "",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "linenums",
						},
					},
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   5,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"require 'app'",
				"",
				"launch_app",
				"# <1>",
			},
		},
		&asciidoc.String{
			Value: "<1> Profit.",
		},
		&asciidoc.NewLine{},
	},
}

var syntaxHighlighterTestShouldPreservePassthroughPlaceholdersWhenHighlightingSourceUsingCoderay = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "source-highlighter",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "coderay",
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
							Value: "java",
						},
					},
				},
				&asciidoc.NamedAttribute{
					Name: "subs",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "specialcharacters,macros,callouts",
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
				"public class Printer {",
				"  public static void main(String[] args) {",
				"    System.pass:quotes[_out_].println(\"*asterisks* make text pass:quotes[*bold*]\");",
				"  }",
				"}",
			},
		},
	},
}

var syntaxHighlighterTestShouldLinkToCodeRayStylesheetIfSourceHighlighterIsCoderayAndLinkcssIsSet = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "source-highlighter",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "coderay",
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
				"require 'coderay'",
				"",
				"html = CodeRay.scan(\"puts 'Hello, world!'\", :ruby).div line_numbers: :table",
			},
		},
	},
}

var syntaxHighlighterTestShouldHighlightSourceInlineIfSourceHighlighterAttributeIsCoderayAndCoderayCssIsStyle = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "source-highlighter",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "coderay",
				},
			},
		},
		&asciidoc.AttributeEntry{
			Name: "coderay-css",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "style",
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
				"require 'coderay'",
				"",
				"html = CodeRay.scan(\"puts 'Hello, world!'\", :ruby).div line_numbers: :table",
			},
		},
	},
}

var syntaxHighlighterTestShouldReadStylesheet = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "source-highlighter",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "highlight.js",
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
				"puts 'Hello, World!'",
			},
		},
	},
}

var syntaxHighlighterTestShouldIncludeRemoteHighlightJsAssetsIfSourceHighlighterAttributeIsHighlightJs = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "source-highlighter",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "highlight.js",
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
							Value: "html",
						},
					},
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   5,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"<p>Highlight me!</p>",
			},
		},
	},
}

var syntaxHighlighterTestShouldAddLanguageNoneClassToSourceBlockWhenSourceHighlighterIsHighlightJsAndLanguageIsNotSet = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "source-highlighter",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "highlight.js",
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
			},
			Delimiter: asciidoc.Delimiter{
				Type:   5,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"[numbers]",
				"one",
				"two",
				"three",
			},
		},
	},
}

var syntaxHighlighterTestShouldLoadAdditionalLanguagesSpecifiedByHighlightjsLanguages = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "source-highlighter",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "highlight.js",
				},
			},
		},
		&asciidoc.AttributeEntry{
			Name: "highlightjs-languages",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "yaml, scilab",
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
				"key: value",
			},
		},
	},
}

var syntaxHighlighterTestShouldAddLanguageClassesToChildCodeElementWhenSourceHighlighterIsPrettify = &asciidoc.Document{
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
				"puts \"foo\"",
			},
		},
	},
}

var syntaxHighlighterTestShouldSetLinenumsStartIfLinenumsAreEnabledAndStartAttributeIsSetWhenSourceHighlighterIsPrettify = &asciidoc.Document{
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
				&asciidoc.NamedAttribute{
					Name: "start",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "5",
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
				"puts \"foo\"",
			},
		},
	},
}

var syntaxHighlighterTestShouldSetLangAttributeOnPreWhenSourceHighlighterIsHtmlPipeline = &asciidoc.Document{
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
				"filters = [",
				"HTML::Pipeline::AsciiDocFilter,",
				"HTML::Pipeline::SanitizationFilter,",
				"HTML::Pipeline::SyntaxHighlightFilter",
				"]",
				"",
				"puts HTML::Pipeline.new(filters, {}).call(input)[:output]",
			},
		},
	},
}

var syntaxHighlighterTestShouldSyntaxHighlightSourceIfSourceHighlighterAttributeIsSet = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "source-highlighter",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "rouge",
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
				"require 'rouge'",
				"",
				"html = Rouge::Formatters::HTML.format(Rouge::Lexers::Ruby.lex('puts \"Hello, world!\"'))",
			},
		},
	},
}

var syntaxHighlighterTestShouldHighlightSourceUsingAMixedLexerHtmlJavaScript = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "[,html]",
		},
		&asciidoc.NewLine{},
		&asciidoc.Listing{
			AttributeList: nil,
			Delimiter: asciidoc.Delimiter{
				Type:   5,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"<meta name=\"description\" content=\"The dangerous and thrilling adventures of an open source documentation team.\">",
				"<script>alert(\"Do your worst!\")</script>",
			},
		},
	},
}

var syntaxHighlighterTestShouldEnableStartInlineForPhpByDefault = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "[,php]",
		},
		&asciidoc.NewLine{},
		&asciidoc.Listing{
			AttributeList: nil,
			Delimiter: asciidoc.Delimiter{
				Type:   5,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"echo \"<?php\";",
			},
		},
	},
}

var syntaxHighlighterTestShouldNotEnableStartInlineForPhpIfDisabledUsingCgiStyleOptionOnLanguage = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "[,php?start_inline=0]",
		},
		&asciidoc.NewLine{},
		&asciidoc.Listing{
			AttributeList: nil,
			Delimiter: asciidoc.Delimiter{
				Type:   5,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"echo \"<?php\";",
			},
		},
	},
}

var syntaxHighlighterTestShouldNotEnableStartInlineForPhpIfMixedOptionIsSet = &asciidoc.Document{
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
									Value: "mixed",
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
							Value: "php",
						},
					},
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   5,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"echo \"<?php\";",
			},
		},
	},
}

var syntaxHighlighterTestShouldNotCrashIfSourceHighlighterAttributeIsSetAndSourceBlockDoesNotDefineALanguage = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "source-highlighter",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "rouge",
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
			},
			Delimiter: asciidoc.Delimiter{
				Type:   5,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"require 'rouge'",
				"",
				"html = Rouge::Formatters::HTML.format(Rouge::Lexers::Ruby.lex('puts \"Hello, world!\"'))",
			},
		},
	},
}

var syntaxHighlighterTestShouldDefaultToPlainTextLexerIfLexerCannotBeResolvedForLanguage = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "source-highlighter",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "rouge",
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
							Value: "lolcode",
						},
					},
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   5,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"CAN HAS STDIO?",
				"PLZ OPEN FILE \"LOLCATS.TXT\"?",
				"KTHXBYE",
			},
		},
	},
}

var syntaxHighlighterTestShouldNumberLinesUsingTableLayoutIfLinenumsOptionIsEnabledAndLinenumsModeIsNotSet = &asciidoc.Document{
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
				"puts 'Hello, world!'",
				"puts 'Goodbye, world!'",
			},
		},
	},
}

var syntaxHighlighterTestShouldNumberLinesUsingInlineElementIfLinenumsOptionIsEnabledAndLinenumsModeIsInline = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "rouge-linenums-mode",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "inline",
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
				"puts 'Hello, world!'",
				"puts 'Goodbye, world!'",
			},
		},
	},
}

var syntaxHighlighterTestShouldGracefullyFallbackToDefaultStyleIfSpecifiedStyleNotRecognized = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "source-highlighter",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "rouge",
				},
			},
		},
		&asciidoc.AttributeEntry{
			Name: "rouge-style",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "unknown",
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
				"puts 'Hello, world!'",
			},
		},
	},
}

var syntaxHighlighterTestShouldRestoreIsolatedCalloutMarkOnLastLineOfSource = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "source-highlighter",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "rouge",
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
				"require 'app'",
				"",
				"launch_app",
				"# <1>",
			},
		},
		&asciidoc.String{
			Value: "<1> Profit.",
		},
		&asciidoc.NewLine{},
	},
}

var syntaxHighlighterTestShouldNumberAllLinesWhenIsolatedCalloutMarkIsOnLastLineOfSourceAndStartingLineNumberIsSet = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "source-highlighter",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "rouge",
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
				&asciidoc.NamedAttribute{
					Name: "start",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "5",
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
				"require 'app'",
				"",
				"launch_app",
				"# <1>",
			},
		},
		&asciidoc.String{
			Value: "<1> Profit.",
		},
		&asciidoc.NewLine{},
	},
}

var syntaxHighlighterTestShouldPreserveGuardInFrontOfCalloutIfIconsAreNotEnabled = &asciidoc.Document{
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
				"puts 'Hello, World!' # <1>",
				"puts 'Goodbye, World ;(' # <2>",
			},
		},
	},
}

var syntaxHighlighterTestShouldPreserveGuardAroundCalloutIfIconsAreNotEnabled = &asciidoc.Document{
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

var syntaxHighlighterTestShouldReadStylesheetForSpecifiedStyle = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "source-highlighter",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "pygments",
				},
			},
		},
		&asciidoc.AttributeEntry{
			Name: "pygments-style",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "monokai",
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
							Value: "python",
						},
					},
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   5,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"from pygments import highlight",
				"from pygments.lexers import PythonLexer",
				"from pygments.formatters import HtmlFormatter",
				"",
				"source = 'print \"Hello World\"'",
				"print(highlight(source, PythonLexer(), HtmlFormatter()))",
			},
		},
	},
}

var syntaxHighlighterTestShouldEmbedStylesheetForPygmentsStyle = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "source-highlighter",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "pygments",
				},
			},
		},
		&asciidoc.AttributeEntry{
			Name: "pygments-style",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "monokai",
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
							Value: "python",
						},
					},
				},
			},
			Delimiter: asciidoc.Delimiter{
				Type:   5,
				Length: 4,
			},
			LineList: asciidoc.LineList{
				"from pygments import highlight",
				"from pygments.lexers import PythonLexer",
				"from pygments.formatters import HtmlFormatter",
				"",
				"source = 'print \"Hello World\"'",
				"print(highlight(source, PythonLexer(), HtmlFormatter()))",
			},
		},
	},
}

var syntaxHighlighterTestShouldRestoreCalloutMarksToCorrectLinesIfTableLineNumberingIsEnabled = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "source-highlighter",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "pygments",
				},
			},
		},
		&asciidoc.AttributeEntry{
			Name: "pygments-linenums-mode",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "table",
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
				"from pygments import highlight # <1>",
				"from pygments.lexers import PythonLexer",
				"from pygments.formatters import HtmlFormatter",
				"",
				"code = 'print \"Hello World\"'",
				"print(highlight(code, PythonLexer(), HtmlFormatter())) # <2><3>",
			},
		},
		&asciidoc.String{
			Value: "<1> Load library",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "<2> Highlight source",
		},
		&asciidoc.NewLine{},
		&asciidoc.String{
			Value: "<3> Print to stdout",
		},
		&asciidoc.NewLine{},
	},
}

var syntaxHighlighterTestShouldNotHardcodeInlineStylesOnLinenoDivAndPreElementsWhenLinenumsAreEnabledInTableMode = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "source-highlighter",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "pygments",
				},
			},
		},
		&asciidoc.AttributeEntry{
			Name: "pygments-css",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "inline",
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
				"puts 'Hello, World!'",
			},
		},
	},
}

var syntaxHighlighterTestShouldLineHighlightSpecifiedLines = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name: "source-highlighter",
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "pygments",
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
				&asciidoc.NamedAttribute{
					Name: "highlight",
					Val: asciidoc.Elements{
						&asciidoc.String{
							Value: "1..2",
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
				"puts 'Hello, world!'",
				"puts 'Goodbye, world!'",
			},
		},
	},
}
