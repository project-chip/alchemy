package tests

import (
	"testing"

	"github.com/hasty/alchemy/asciidoc"
)

func TestLinks(t *testing.T) {
	linksTests.run(t)
}

var linksTests = parseTests{
	
	{ "unescapes square bracket in reftext of anchor macro", "asciidoctor/links_test_unescapes_square_bracket_in_reftext_of_anchor_macro.adoc", unescapesSquareBracketInReftextOfAnchorMacro },
	
	{ "xref using angled bracket syntax with label", "asciidoctor/links_test_xref_using_angled_bracket_syntax_with_label.adoc", xrefUsingAngledBracketSyntaxWithLabel },
	
	{ "xref should use title of target as link text when no explicit reftext is specified", "asciidoctor/links_test_xref_should_use_title_of_target_as_link_text_when_no_explicit_reftext_is_specified.adoc", xrefShouldUseTitleOfTargetAsLinkTextWhenNoExplicitReftextIsSpecified },
	
	{ "xref should use title of target as link text when explicit link text is empty", "asciidoctor/links_test_xref_should_use_title_of_target_as_link_text_when_explicit_link_text_is_empty.adoc", xrefShouldUseTitleOfTargetAsLinkTextWhenExplicitLinkTextIsEmpty },
	
	{ "xref using angled bracket syntax with quoted label", "asciidoctor/links_test_xref_using_angled_bracket_syntax_with_quoted_label.adoc", xrefUsingAngledBracketSyntaxWithQuotedLabel },
	
	{ "xref using angled bracket syntax inline with text", "asciidoctor/links_test_xref_using_angled_bracket_syntax_inline_with_text.adoc", xrefUsingAngledBracketSyntaxInlineWithText },
	
	{ "xref using angled bracket syntax with multi-line label inline with text", "asciidoctor/links_test_xref_using_angled_bracket_syntax_with_multi_line_label_inline_with_text.adoc", xrefUsingAngledBracketSyntaxWithMultiLineLabelInlineWithText },
	
	{ "xref with escaped text", "asciidoctor/links_test_xref_with_escaped_text.adoc", xrefWithEscapedText },
	
	{ "xref using macro syntax", "asciidoctor/links_test_xref_using_macro_syntax.adoc", xrefUsingMacroSyntax },
	
	{ "xref using macro syntax with explicit hash", "asciidoctor/links_test_xref_using_macro_syntax_with_explicit_hash.adoc", xrefUsingMacroSyntaxWithExplicitHash },
	
	{ "xref using macro syntax inline with text", "asciidoctor/links_test_xref_using_macro_syntax_inline_with_text.adoc", xrefUsingMacroSyntaxInlineWithText },
	
	{ "xref using macro syntax with multi-line label inline with text", "asciidoctor/links_test_xref_using_macro_syntax_with_multi_line_label_inline_with_text.adoc", xrefUsingMacroSyntaxWithMultiLineLabelInlineWithText },
	
	{ "xref using macro syntax with text that ends with an escaped closing bracket", "asciidoctor/links_test_xref_using_macro_syntax_with_text_that_ends_with_an_escaped_closing_bracket.adoc", xrefUsingMacroSyntaxWithTextThatEndsWithAnEscapedClosingBracket },
	
	{ "xref using macro syntax with text that contains an escaped closing bracket", "asciidoctor/links_test_xref_using_macro_syntax_with_text_that_contains_an_escaped_closing_bracket.adoc", xrefUsingMacroSyntaxWithTextThatContainsAnEscapedClosingBracket },
	
	{ "unescapes square bracket in reftext used by xref", "asciidoctor/links_test_unescapes_square_bracket_in_reftext_used_by_xref.adoc", unescapesSquareBracketInReftextUsedByXref },
	
	{ "xref using invalid macro syntax does not create link", "asciidoctor/links_test_xref_using_invalid_macro_syntax_does_not_create_link.adoc", xrefUsingInvalidMacroSyntaxDoesNotCreateLink },
	
	{ "should not warn if verbose flag is set and reference is found in compat mode", "asciidoctor/links_test_should_not_warn_if_verbose_flag_is_set_and_reference_is_found_in_compat_mode.adoc", shouldNotWarnIfVerboseFlagIsSetAndReferenceIsFoundInCompatMode },
	
	{ "should warn and create link if verbose flag is set and reference using # notation is not found", "asciidoctor/links_test_should_warn_and_create_link_if_verbose_flag_is_set_and_reference_using_#_notation_is_not_found.adoc", shouldWarnAndCreateLinkIfVerboseFlagIsSetAndReferenceUsingNotationIsNotFound },
	
	{ "should produce an internal anchor from an inter-document xref to file included into current file", "asciidoctor/links_test_should_produce_an_internal_anchor_from_an_inter_document_xref_to_file_included_into_current_file.adoc", shouldProduceAnInternalAnchorFromAnInterDocumentXrefToFileIncludedIntoCurrentFile },
	
	{ "should produce an internal anchor from an inter-document xref to file included entirely into current file using tags", "asciidoctor/links_test_should_produce_an_internal_anchor_from_an_inter_document_xref_to_file_included_entirely_into_current_file_using_tags.adoc", shouldProduceAnInternalAnchorFromAnInterDocumentXrefToFileIncludedEntirelyIntoCurrentFileUsingTags },
	
	{ "should not produce an internal anchor for inter-document xref to file partially included into current file", "asciidoctor/links_test_should_not_produce_an_internal_anchor_for_inter_document_xref_to_file_partially_included_into_current_file.adoc", shouldNotProduceAnInternalAnchorForInterDocumentXrefToFilePartiallyIncludedIntoCurrentFile },
	
	{ "should produce an internal anchor for inter-document xref to file included fully and partially", "asciidoctor/links_test_should_produce_an_internal_anchor_for_inter_document_xref_to_file_included_fully_and_partially.adoc", shouldProduceAnInternalAnchorForInterDocumentXrefToFileIncludedFullyAndPartially },
	
	{ "should warn and create link if debug mode is enabled, inter-document xref points to current doc, and reference not found", "asciidoctor/links_test_should_warn_and_create_link_if_debug_mode_is_enabled_inter_document_xref_points_to_current_doc_and_reference_not_found.adoc", shouldWarnAndCreateLinkIfDebugModeIsEnabledInterDocumentXrefPointsToCurrentDocAndReferenceNotFound },
	
	{ "should use doctitle as fallback link text if inter-document xref points to current doc and no link text is provided", "asciidoctor/links_test_should_use_doctitle_as_fallback_link_text_if_inter_document_xref_points_to_current_doc_and_no_link_text_is_provided.adoc", shouldUseDoctitleAsFallbackLinkTextIfInterDocumentXrefPointsToCurrentDocAndNoLinkTextIsProvided },
	
	{ "should use doctitle of root document as fallback link text for inter-document xref in AsciiDoc table cell that resolves to current doc", "asciidoctor/links_test_should_use_doctitle_of_root_document_as_fallback_link_text_for_inter_document_xref_in_ascii_doc_table_cell_that_resolves_to_current_doc.adoc", shouldUseDoctitleOfRootDocumentAsFallbackLinkTextForInterDocumentXrefInAsciiDocTableCellThatResolvesToCurrentDoc },
	
	{ "should use reftext on document as fallback link text if inter-document xref points to current doc and no link text is provided", "asciidoctor/links_test_should_use_reftext_on_document_as_fallback_link_text_if_inter_document_xref_points_to_current_doc_and_no_link_text_is_provided.adoc", shouldUseReftextOnDocumentAsFallbackLinkTextIfInterDocumentXrefPointsToCurrentDocAndNoLinkTextIsProvided },
	
	{ "should use reftext on document as fallback link text if xref points to empty fragment and no link text is provided", "asciidoctor/links_test_should_use_reftext_on_document_as_fallback_link_text_if_xref_points_to_empty_fragment_and_no_link_text_is_provided.adoc", shouldUseReftextOnDocumentAsFallbackLinkTextIfXrefPointsToEmptyFragmentAndNoLinkTextIsProvided },
	
	{ "should use fallback link text if inter-document xref points to current doc without header and no link text is provided", "asciidoctor/links_test_should_use_fallback_link_text_if_inter_document_xref_points_to_current_doc_without_header_and_no_link_text_is_provided.adoc", shouldUseFallbackLinkTextIfInterDocumentXrefPointsToCurrentDocWithoutHeaderAndNoLinkTextIsProvided },
	
	{ "should use fallback link text if fragment of internal xref is empty and no link text is provided", "asciidoctor/links_test_should_use_fallback_link_text_if_fragment_of_internal_xref_is_empty_and_no_link_text_is_provided.adoc", shouldUseFallbackLinkTextIfFragmentOfInternalXrefIsEmptyAndNoLinkTextIsProvided },
	
	{ "should use document id as linkend for self xref in DocBook backend", "asciidoctor/links_test_should_use_document_id_as_linkend_for_self_xref_in_doc_book_backend.adoc", shouldUseDocumentIdAsLinkendForSelfXrefInDocBookBackend },
	
	{ "should auto-generate document id to use as linkend for self xref in DocBook backend", "asciidoctor/links_test_should_auto_generate_document_id_to_use_as_linkend_for_self_xref_in_doc_book_backend.adoc", shouldAutoGenerateDocumentIdToUseAsLinkendForSelfXrefInDocBookBackend },
	
	{ "should produce an internal anchor for inter-document xref to file outside of base directory", "asciidoctor/links_test_should_produce_an_internal_anchor_for_inter_document_xref_to_file_outside_of_base_directory.adoc", shouldProduceAnInternalAnchorForInterDocumentXrefToFileOutsideOfBaseDirectory },
	
	{ "xref uses title of target as label for forward and backward references in html output", "asciidoctor/links_test_xref_uses_title_of_target_as_label_for_forward_and_backward_references_in_html_output.adoc", xrefUsesTitleOfTargetAsLabelForForwardAndBackwardReferencesInHtmlOutput },
	
	{ "should not fail to resolve broken xref in title of block with ID", "asciidoctor/links_test_should_not_fail_to_resolve_broken_xref_in_title_of_block_with_id.adoc", shouldNotFailToResolveBrokenXrefInTitleOfBlockWithId },
	
	{ "should resolve forward xref in title of block with ID", "asciidoctor/links_test_should_resolve_forward_xref_in_title_of_block_with_id.adoc", shouldResolveForwardXrefInTitleOfBlockWithId },
	
	{ "should not fail to resolve broken xref in section title", "asciidoctor/links_test_should_not_fail_to_resolve_broken_xref_in_section_title.adoc", shouldNotFailToResolveBrokenXrefInSectionTitle },
	
	{ "should break circular xref reference in section title", "asciidoctor/links_test_should_break_circular_xref_reference_in_section_title.adoc", shouldBreakCircularXrefReferenceInSectionTitle },
	
	{ "should drop nested anchor in xreftext", "asciidoctor/links_test_should_drop_nested_anchor_in_xreftext.adoc", shouldDropNestedAnchorInXreftext },
	
	{ "should not resolve forward xref evaluated during parsing", "asciidoctor/links_test_should_not_resolve_forward_xref_evaluated_during_parsing.adoc", shouldNotResolveForwardXrefEvaluatedDuringParsing },
	
	{ "should not resolve forward natural xref evaluated during parsing", "asciidoctor/links_test_should_not_resolve_forward_natural_xref_evaluated_during_parsing.adoc", shouldNotResolveForwardNaturalXrefEvaluatedDuringParsing },
	
	{ "should resolve first matching natural xref", "asciidoctor/links_test_should_resolve_first_matching_natural_xref.adoc", shouldResolveFirstMatchingNaturalXref },
	
	{ "should not match numeric character references while searching for fragment in xref target", "asciidoctor/links_test_should_not_match_numeric_character_references_while_searching_for_fragment_in_xref_target.adoc", shouldNotMatchNumericCharacterReferencesWhileSearchingForFragmentInXrefTarget },
	
	{ "should not match numeric character references in path of interdocument xref", "asciidoctor/links_test_should_not_match_numeric_character_references_in_path_of_interdocument_xref.adoc", shouldNotMatchNumericCharacterReferencesInPathOfInterdocumentXref },
	
}


var unescapesSquareBracketInReftextOfAnchorMacro = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "see ",
    },
    &asciidoc.CrossReference{
      Set: nil,
      ID: "foo",
    },
    &asciidoc.NewLine{},
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "anchor:foo[b[a\\]r]tex",
    },
    &asciidoc.NewLine{},
  },
}

var xrefUsingAngledBracketSyntaxWithLabel = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.CrossReference{
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "About Tigers",
        },
      },
      ID: "tigers",
    },
    &asciidoc.NewLine{},
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Section{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.ShorthandAttribute{
          Style: nil,
          ID: &asciidoc.ShorthandID{
            Set: asciidoc.Set{
              &asciidoc.String{
                Value: "tigers",
              },
            },
          },
          Roles: nil,
          Options: nil,
        },
      },
      Set: nil,
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "Tigers",
        },
      },
      Level: 1,
    },
  },
}

var xrefShouldUseTitleOfTargetAsLinkTextWhenNoExplicitReftextIsSpecified = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.CrossReference{
      Set: nil,
      ID: "tigers",
    },
    &asciidoc.NewLine{},
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Section{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.ShorthandAttribute{
          Style: nil,
          ID: &asciidoc.ShorthandID{
            Set: asciidoc.Set{
              &asciidoc.String{
                Value: "tigers",
              },
            },
          },
          Roles: nil,
          Options: nil,
        },
      },
      Set: nil,
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "Tigers",
        },
      },
      Level: 1,
    },
  },
}

var xrefShouldUseTitleOfTargetAsLinkTextWhenExplicitLinkTextIsEmpty = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "<<tigers,>>",
    },
    &asciidoc.NewLine{},
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Section{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.ShorthandAttribute{
          Style: nil,
          ID: &asciidoc.ShorthandID{
            Set: asciidoc.Set{
              &asciidoc.String{
                Value: "tigers",
              },
            },
          },
          Roles: nil,
          Options: nil,
        },
      },
      Set: nil,
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "Tigers",
        },
      },
      Level: 1,
    },
  },
}

var xrefUsingAngledBracketSyntaxWithQuotedLabel = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.CrossReference{
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "\"About Tigers\"",
        },
      },
      ID: "tigers",
    },
    &asciidoc.NewLine{},
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Section{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.ShorthandAttribute{
          Style: nil,
          ID: &asciidoc.ShorthandID{
            Set: asciidoc.Set{
              &asciidoc.String{
                Value: "tigers",
              },
            },
          },
          Roles: nil,
          Options: nil,
        },
      },
      Set: nil,
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "Tigers",
        },
      },
      Level: 1,
    },
  },
}

var xrefUsingAngledBracketSyntaxInlineWithText = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "Want to learn ",
    },
    &asciidoc.CrossReference{
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "about tigers",
        },
      },
      ID: "tigers",
    },
    &asciidoc.String{
      Value: "?",
    },
    &asciidoc.NewLine{},
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Section{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.ShorthandAttribute{
          Style: nil,
          ID: &asciidoc.ShorthandID{
            Set: asciidoc.Set{
              &asciidoc.String{
                Value: "tigers",
              },
            },
          },
          Roles: nil,
          Options: nil,
        },
      },
      Set: nil,
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "Tigers",
        },
      },
      Level: 1,
    },
  },
}

var xrefUsingAngledBracketSyntaxWithMultiLineLabelInlineWithText = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "Want to learn ",
    },
    &asciidoc.CrossReference{
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "about\ntigers",
        },
      },
      ID: "tigers",
    },
    &asciidoc.String{
      Value: "?",
    },
    &asciidoc.NewLine{},
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Section{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.ShorthandAttribute{
          Style: nil,
          ID: &asciidoc.ShorthandID{
            Set: asciidoc.Set{
              &asciidoc.String{
                Value: "tigers",
              },
            },
          },
          Roles: nil,
          Options: nil,
        },
      },
      Set: nil,
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "Tigers",
        },
      },
      Level: 1,
    },
  },
}

var xrefWithEscapedText = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "See the <<tigers, ",
    },
    &asciidoc.Monospace{
      AttributeList: nil,
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "+[tigers]+",
        },
      },
    },
    &asciidoc.String{
      Value: ">> section for details about tigers.",
    },
    &asciidoc.NewLine{},
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Section{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.ShorthandAttribute{
          Style: nil,
          ID: &asciidoc.ShorthandID{
            Set: asciidoc.Set{
              &asciidoc.String{
                Value: "tigers",
              },
            },
          },
          Roles: nil,
          Options: nil,
        },
      },
      Set: nil,
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "Tigers",
        },
      },
      Level: 1,
    },
  },
}

var xrefUsingMacroSyntax = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "This document has two sections, xref:sect-a[] and xref:sect-b[].",
    },
    &asciidoc.NewLine{},
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Section{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.ShorthandAttribute{
          Style: nil,
          ID: &asciidoc.ShorthandID{
            Set: asciidoc.Set{
              &asciidoc.String{
                Value: "sect-a",
              },
            },
          },
          Roles: nil,
          Options: nil,
        },
      },
      Set: asciidoc.Set{
        asciidoc.EmptyLine{
          Text: "",
        },
      },
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "Section A",
        },
      },
      Level: 1,
    },
    &asciidoc.Section{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.ShorthandAttribute{
          Style: nil,
          ID: &asciidoc.ShorthandID{
            Set: asciidoc.Set{
              &asciidoc.String{
                Value: "sect-b",
              },
            },
          },
          Roles: nil,
          Options: nil,
        },
      },
      Set: nil,
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "Section B",
        },
      },
      Level: 1,
    },
  },
}

var xrefUsingMacroSyntaxWithExplicitHash = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.DocumentCrossReference{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "alt",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "About Tigers",
            },
          },
        },
      },
      Path: asciidoc.Set{
        &asciidoc.String{
          Value: "tigers",
        },
      },
    },
    &asciidoc.NewLine{},
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Section{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.ShorthandAttribute{
          Style: nil,
          ID: &asciidoc.ShorthandID{
            Set: asciidoc.Set{
              &asciidoc.String{
                Value: "tigers",
              },
            },
          },
          Roles: nil,
          Options: nil,
        },
      },
      Set: nil,
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "Tigers",
        },
      },
      Level: 1,
    },
  },
}

var xrefUsingMacroSyntaxInlineWithText = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "Want to learn xref:tigers[about tigers]?",
    },
    &asciidoc.NewLine{},
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Section{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.ShorthandAttribute{
          Style: nil,
          ID: &asciidoc.ShorthandID{
            Set: asciidoc.Set{
              &asciidoc.String{
                Value: "tigers",
              },
            },
          },
          Roles: nil,
          Options: nil,
        },
      },
      Set: nil,
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "Tigers",
        },
      },
      Level: 1,
    },
  },
}

var xrefUsingMacroSyntaxWithMultiLineLabelInlineWithText = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "Want to learn xref:tigers[about",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "tigers]?",
    },
    &asciidoc.NewLine{},
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Section{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.ShorthandAttribute{
          Style: nil,
          ID: &asciidoc.ShorthandID{
            Set: asciidoc.Set{
              &asciidoc.String{
                Value: "tigers",
              },
            },
          },
          Roles: nil,
          Options: nil,
        },
      },
      Set: nil,
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "Tigers",
        },
      },
      Level: 1,
    },
  },
}

var xrefUsingMacroSyntaxWithTextThatEndsWithAnEscapedClosingBracket = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.DocumentCrossReference{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "alt",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "[tigers\\",
            },
          },
        },
      },
      Path: asciidoc.Set{
        &asciidoc.String{
          Value: "tigers",
        },
      },
    },
    &asciidoc.String{
      Value: "]",
    },
    &asciidoc.NewLine{},
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Section{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.ShorthandAttribute{
          Style: nil,
          ID: &asciidoc.ShorthandID{
            Set: asciidoc.Set{
              &asciidoc.String{
                Value: "tigers",
              },
            },
          },
          Roles: nil,
          Options: nil,
        },
      },
      Set: nil,
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "Tigers",
        },
      },
      Level: 1,
    },
  },
}

var xrefUsingMacroSyntaxWithTextThatContainsAnEscapedClosingBracket = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.DocumentCrossReference{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.PositionalAttribute{
          Offset: 0,
          ImpliedName: "alt",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "[tigers\\",
            },
          },
        },
      },
      Path: asciidoc.Set{
        &asciidoc.String{
          Value: "tigers",
        },
      },
    },
    &asciidoc.String{
      Value: " are cats]",
    },
    &asciidoc.NewLine{},
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Section{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.ShorthandAttribute{
          Style: nil,
          ID: &asciidoc.ShorthandID{
            Set: asciidoc.Set{
              &asciidoc.String{
                Value: "tigers",
              },
            },
          },
          Roles: nil,
          Options: nil,
        },
      },
      Set: nil,
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "Tigers",
        },
      },
      Level: 1,
    },
  },
}

var unescapesSquareBracketInReftextUsedByXref = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "anchor:foo[b[a\\]r]about",
    },
    &asciidoc.NewLine{},
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "see ",
    },
    &asciidoc.CrossReference{
      Set: nil,
      ID: "foo",
    },
    &asciidoc.NewLine{},
  },
}

var xrefUsingInvalidMacroSyntaxDoesNotCreateLink = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Section{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.ShorthandAttribute{
          Style: nil,
          ID: &asciidoc.ShorthandID{
            Set: asciidoc.Set{
              &asciidoc.String{
                Value: "foobar",
              },
            },
          },
          Roles: nil,
          Options: nil,
        },
      },
      Set: asciidoc.Set{
        asciidoc.EmptyLine{
          Text: "",
        },
      },
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "Foobar",
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
          Value: "See ",
        },
        &asciidoc.CrossReference{
          Set: nil,
          ID: "foobaz",
        },
        &asciidoc.String{
          Value: ".",
        },
        &asciidoc.NewLine{},
      },
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "Section B",
        },
      },
      Level: 1,
    },
  },
}

var shouldNotWarnIfVerboseFlagIsSetAndReferenceIsFoundInCompatMode = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Section{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.AnchorAttribute{
          ID: &asciidoc.String{
            Value: "foobar",
          },
          Label: nil,
        },
      },
      Set: asciidoc.Set{
        asciidoc.EmptyLine{
          Text: "",
        },
      },
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "Foobar",
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
          Value: "See ",
        },
        &asciidoc.CrossReference{
          Set: nil,
          ID: "foobar",
        },
        &asciidoc.String{
          Value: ".",
        },
        &asciidoc.NewLine{},
      },
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "Section B",
        },
      },
      Level: 1,
    },
  },
}

var shouldWarnAndCreateLinkIfVerboseFlagIsSetAndReferenceUsingNotationIsNotFound = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Section{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.ShorthandAttribute{
          Style: nil,
          ID: &asciidoc.ShorthandID{
            Set: asciidoc.Set{
              &asciidoc.String{
                Value: "foobar",
              },
            },
          },
          Roles: nil,
          Options: nil,
        },
      },
      Set: asciidoc.Set{
        asciidoc.EmptyLine{
          Text: "",
        },
      },
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "Foobar",
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
          Value: "See <<#foobaz>>.",
        },
        &asciidoc.NewLine{},
      },
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "Section B",
        },
      },
      Level: 1,
    },
  },
}

var shouldProduceAnInternalAnchorFromAnInterDocumentXrefToFileIncludedIntoCurrentFile = &asciidoc.Document{
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
          AttributeList: asciidoc.AttributeList{
            &asciidoc.ShorthandAttribute{
              Style: nil,
              ID: &asciidoc.ShorthandID{
                Set: asciidoc.Set{
                  &asciidoc.String{
                    Value: "ch1",
                  },
                },
              },
              Roles: nil,
              Options: nil,
            },
          },
          Set: asciidoc.Set{
            asciidoc.EmptyLine{
              Text: "",
            },
            &asciidoc.String{
              Value: "So it begins.",
            },
            &asciidoc.NewLine{},
            asciidoc.EmptyLine{
              Text: "",
            },
            &asciidoc.String{
              Value: "Read <<other-chapters.adoc#ch2>> to find out what happens next!",
            },
            &asciidoc.NewLine{},
            asciidoc.EmptyLine{
              Text: "",
            },
            &asciidoc.FileInclude{
              AttributeList: nil,
              Set: asciidoc.Set{
                &asciidoc.String{
                  Value: "other-chapters.adoc",
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

var shouldProduceAnInternalAnchorFromAnInterDocumentXrefToFileIncludedEntirelyIntoCurrentFileUsingTags = &asciidoc.Document{
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
          AttributeList: asciidoc.AttributeList{
            &asciidoc.ShorthandAttribute{
              Style: nil,
              ID: &asciidoc.ShorthandID{
                Set: asciidoc.Set{
                  &asciidoc.String{
                    Value: "ch1",
                  },
                },
              },
              Roles: nil,
              Options: nil,
            },
          },
          Set: asciidoc.Set{
            asciidoc.EmptyLine{
              Text: "",
            },
            &asciidoc.String{
              Value: "So it begins.",
            },
            &asciidoc.NewLine{},
            asciidoc.EmptyLine{
              Text: "",
            },
            &asciidoc.String{
              Value: "Read <<other-chapters.adoc#ch2>> to find out what happens next!",
            },
            &asciidoc.NewLine{},
            asciidoc.EmptyLine{
              Text: "",
            },
            &asciidoc.FileInclude{
              AttributeList: asciidoc.AttributeList{
                &asciidoc.NamedAttribute{
                  Name: "tags",
                  Val: asciidoc.Set{
                    &asciidoc.String{
                      Value: "**",
                    },
                  },
                  Quote: 0,
                },
              },
              Set: asciidoc.Set{
                &asciidoc.String{
                  Value: "other-chapters.adoc",
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

var shouldNotProduceAnInternalAnchorForInterDocumentXrefToFilePartiallyIncludedIntoCurrentFile = &asciidoc.Document{
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
          AttributeList: asciidoc.AttributeList{
            &asciidoc.ShorthandAttribute{
              Style: nil,
              ID: &asciidoc.ShorthandID{
                Set: asciidoc.Set{
                  &asciidoc.String{
                    Value: "ch1",
                  },
                },
              },
              Roles: nil,
              Options: nil,
            },
          },
          Set: asciidoc.Set{
            asciidoc.EmptyLine{
              Text: "",
            },
            &asciidoc.String{
              Value: "So it begins.",
            },
            &asciidoc.NewLine{},
            asciidoc.EmptyLine{
              Text: "",
            },
            &asciidoc.String{
              Value: "Read <<other-chapters.adoc#ch2,the next chapter>> to find out what happens next!",
            },
            &asciidoc.NewLine{},
            asciidoc.EmptyLine{
              Text: "",
            },
            &asciidoc.FileInclude{
              AttributeList: asciidoc.AttributeList{
                &asciidoc.NamedAttribute{
                  Name: "tags",
                  Val: asciidoc.Set{
                    &asciidoc.String{
                      Value: "ch2",
                    },
                  },
                  Quote: 0,
                },
              },
              Set: asciidoc.Set{
                &asciidoc.String{
                  Value: "other-chapters.adoc",
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

var shouldProduceAnInternalAnchorForInterDocumentXrefToFileIncludedFullyAndPartially = &asciidoc.Document{
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
          AttributeList: asciidoc.AttributeList{
            &asciidoc.ShorthandAttribute{
              Style: nil,
              ID: &asciidoc.ShorthandID{
                Set: asciidoc.Set{
                  &asciidoc.String{
                    Value: "ch1",
                  },
                },
              },
              Roles: nil,
              Options: nil,
            },
          },
          Set: asciidoc.Set{
            asciidoc.EmptyLine{
              Text: "",
            },
            &asciidoc.String{
              Value: "So it begins.",
            },
            &asciidoc.NewLine{},
            asciidoc.EmptyLine{
              Text: "",
            },
            &asciidoc.String{
              Value: "Read <<other-chapters.adoc#ch2,the next chapter>> to find out what happens next!",
            },
            &asciidoc.NewLine{},
            asciidoc.EmptyLine{
              Text: "",
            },
            &asciidoc.FileInclude{
              AttributeList: nil,
              Set: asciidoc.Set{
                &asciidoc.String{
                  Value: "other-chapters.adoc",
                },
              },
            },
            asciidoc.EmptyLine{
              Text: "",
            },
            &asciidoc.FileInclude{
              AttributeList: asciidoc.AttributeList{
                &asciidoc.NamedAttribute{
                  Name: "tag",
                  Val: asciidoc.Set{
                    &asciidoc.String{
                      Value: "ch2-noid",
                    },
                  },
                  Quote: 0,
                },
              },
              Set: asciidoc.Set{
                &asciidoc.String{
                  Value: "other-chapters.adoc",
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

var shouldWarnAndCreateLinkIfDebugModeIsEnabledInterDocumentXrefPointsToCurrentDocAndReferenceNotFound = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Section{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.ShorthandAttribute{
          Style: nil,
          ID: &asciidoc.ShorthandID{
            Set: asciidoc.Set{
              &asciidoc.String{
                Value: "foobar",
              },
            },
          },
          Roles: nil,
          Options: nil,
        },
      },
      Set: asciidoc.Set{
        asciidoc.EmptyLine{
          Text: "",
        },
      },
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "Foobar",
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
          Value: "See <<test.adoc#foobaz>>.",
        },
        &asciidoc.NewLine{},
      },
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "Section B",
        },
      },
      Level: 1,
    },
  },
}

var shouldUseDoctitleAsFallbackLinkTextIfInterDocumentXrefPointsToCurrentDocAndNoLinkTextIsProvided = &asciidoc.Document{
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
          Value: "See xref:test.adoc[]",
        },
        &asciidoc.NewLine{},
      },
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "Links & Stuff at https://example.org",
        },
      },
      Level: 0,
    },
  },
}

var shouldUseDoctitleOfRootDocumentAsFallbackLinkTextForInterDocumentXrefInAsciiDocTableCellThatResolvesToCurrentDoc = &asciidoc.Document{
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
                      Value: "See xref:test.adoc[]",
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

var shouldUseReftextOnDocumentAsFallbackLinkTextIfInterDocumentXrefPointsToCurrentDocAndNoLinkTextIsProvided = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Section{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.NamedAttribute{
          Name: "reftext",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Links and Stuff",
            },
          },
          Quote: 2,
        },
      },
      Set: asciidoc.Set{
        asciidoc.EmptyLine{
          Text: "",
        },
        &asciidoc.String{
          Value: "See xref:test.adoc[]",
        },
        &asciidoc.NewLine{},
      },
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "Links & Stuff",
        },
      },
      Level: 0,
    },
  },
}

var shouldUseReftextOnDocumentAsFallbackLinkTextIfXrefPointsToEmptyFragmentAndNoLinkTextIsProvided = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Section{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.NamedAttribute{
          Name: "reftext",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Links and Stuff",
            },
          },
          Quote: 2,
        },
      },
      Set: asciidoc.Set{
        asciidoc.EmptyLine{
          Text: "",
        },
        &asciidoc.String{
          Value: "See xref:#[]",
        },
        &asciidoc.NewLine{},
      },
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "Links & Stuff",
        },
      },
      Level: 0,
    },
  },
}

var shouldUseFallbackLinkTextIfInterDocumentXrefPointsToCurrentDocWithoutHeaderAndNoLinkTextIsProvided = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "See xref:test.adoc[]",
    },
    &asciidoc.NewLine{},
  },
}

var shouldUseFallbackLinkTextIfFragmentOfInternalXrefIsEmptyAndNoLinkTextIsProvided = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "See xref:#[]",
    },
    &asciidoc.NewLine{},
  },
}

var shouldUseDocumentIdAsLinkendForSelfXrefInDocBookBackend = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Section{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.ShorthandAttribute{
          Style: nil,
          ID: &asciidoc.ShorthandID{
            Set: asciidoc.Set{
              &asciidoc.String{
                Value: "docid",
              },
            },
          },
          Roles: nil,
          Options: nil,
        },
      },
      Set: asciidoc.Set{
        asciidoc.EmptyLine{
          Text: "",
        },
        &asciidoc.String{
          Value: "See xref:test.adoc[]",
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

var shouldAutoGenerateDocumentIdToUseAsLinkendForSelfXrefInDocBookBackend = &asciidoc.Document{
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
          Value: "See xref:test.adoc[]",
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

var shouldProduceAnInternalAnchorForInterDocumentXrefToFileOutsideOfBaseDirectory = &asciidoc.Document{
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
          Value: "See <<../section-a.adoc#section-a>>.",
        },
        &asciidoc.NewLine{},
        asciidoc.EmptyLine{
          Text: "",
        },
        &asciidoc.FileInclude{
          AttributeList: nil,
          Set: asciidoc.Set{
            &asciidoc.String{
              Value: "../section-a.adoc",
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

var xrefUsesTitleOfTargetAsLabelForForwardAndBackwardReferencesInHtmlOutput = &asciidoc.Document{
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
        &asciidoc.CrossReference{
          Set: nil,
          ID: "_section_b",
        },
        &asciidoc.NewLine{},
        asciidoc.EmptyLine{
          Text: "",
        },
      },
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "Section A",
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
        &asciidoc.CrossReference{
          Set: nil,
          ID: "_section_a",
        },
        &asciidoc.NewLine{},
      },
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "Section B",
        },
      },
      Level: 1,
    },
  },
}

var shouldNotFailToResolveBrokenXrefInTitleOfBlockWithId = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Paragraph{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.ShorthandAttribute{
          Style: nil,
          ID: &asciidoc.ShorthandID{
            Set: asciidoc.Set{
              &asciidoc.String{
                Value: "p1",
              },
            },
          },
          Roles: nil,
          Options: nil,
        },
        &asciidoc.TitleAttribute{
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "<<DNE>>",
            },
          },
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "paragraph text",
        },
        &asciidoc.NewLine{},
      },
      Admonition: 0,
    },
  },
}

var shouldResolveForwardXrefInTitleOfBlockWithId = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Paragraph{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.ShorthandAttribute{
          Style: nil,
          ID: &asciidoc.ShorthandID{
            Set: asciidoc.Set{
              &asciidoc.String{
                Value: "p1",
              },
            },
          },
          Roles: nil,
          Options: nil,
        },
        &asciidoc.TitleAttribute{
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "<<conclusion>>",
            },
          },
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "paragraph text",
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
        &asciidoc.ShorthandAttribute{
          Style: nil,
          ID: &asciidoc.ShorthandID{
            Set: asciidoc.Set{
              &asciidoc.String{
                Value: "conclusion",
              },
            },
          },
          Roles: nil,
          Options: nil,
        },
      },
      Set: nil,
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "Conclusion",
        },
      },
      Level: 1,
    },
  },
}

var shouldNotFailToResolveBrokenXrefInSectionTitle = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Section{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.ShorthandAttribute{
          Style: nil,
          ID: &asciidoc.ShorthandID{
            Set: asciidoc.Set{
              &asciidoc.String{
                Value: "s1",
              },
            },
          },
          Roles: nil,
          Options: nil,
        },
      },
      Set: asciidoc.Set{
        asciidoc.EmptyLine{
          Text: "",
        },
      },
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "<<DNE>>",
        },
      },
      Level: 1,
    },
    &asciidoc.Section{
      AttributeList: nil,
      Set: nil,
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "<<s1>>",
        },
      },
      Level: 1,
    },
  },
}

var shouldBreakCircularXrefReferenceInSectionTitle = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Section{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.ShorthandAttribute{
          Style: nil,
          ID: &asciidoc.ShorthandID{
            Set: asciidoc.Set{
              &asciidoc.String{
                Value: "a",
              },
            },
          },
          Roles: nil,
          Options: nil,
        },
      },
      Set: asciidoc.Set{
        asciidoc.EmptyLine{
          Text: "",
        },
      },
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "A <<b>>",
        },
      },
      Level: 1,
    },
    &asciidoc.Section{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.ShorthandAttribute{
          Style: nil,
          ID: &asciidoc.ShorthandID{
            Set: asciidoc.Set{
              &asciidoc.String{
                Value: "b",
              },
            },
          },
          Roles: nil,
          Options: nil,
        },
      },
      Set: nil,
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "B <<a>>",
        },
      },
      Level: 1,
    },
  },
}

var shouldDropNestedAnchorInXreftext = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Section{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.ShorthandAttribute{
          Style: nil,
          ID: &asciidoc.ShorthandID{
            Set: asciidoc.Set{
              &asciidoc.String{
                Value: "a",
              },
            },
          },
          Roles: nil,
          Options: nil,
        },
      },
      Set: asciidoc.Set{
        asciidoc.EmptyLine{
          Text: "",
        },
      },
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "See <<b>>",
        },
      },
      Level: 1,
    },
    &asciidoc.Section{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.ShorthandAttribute{
          Style: nil,
          ID: &asciidoc.ShorthandID{
            Set: asciidoc.Set{
              &asciidoc.String{
                Value: "b",
              },
            },
          },
          Roles: nil,
          Options: nil,
        },
      },
      Set: nil,
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "Consult https://google.com[Google]",
        },
      },
      Level: 1,
    },
  },
}

var shouldNotResolveForwardXrefEvaluatedDuringParsing = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Section{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.ShorthandAttribute{
          Style: nil,
          ID: &asciidoc.ShorthandID{
            Set: asciidoc.Set{
              &asciidoc.String{
                Value: "s1",
              },
            },
          },
          Roles: nil,
          Options: nil,
        },
      },
      Set: asciidoc.Set{
        asciidoc.EmptyLine{
          Text: "",
        },
      },
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "<<forward>>",
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
      },
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "<<s1>>",
        },
      },
      Level: 1,
    },
    &asciidoc.Section{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.ShorthandAttribute{
          Style: nil,
          ID: &asciidoc.ShorthandID{
            Set: asciidoc.Set{
              &asciidoc.String{
                Value: "forward",
              },
            },
          },
          Roles: nil,
          Options: nil,
        },
      },
      Set: nil,
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "Forward",
        },
      },
      Level: 1,
    },
  },
}

var shouldNotResolveForwardNaturalXrefEvaluatedDuringParsing = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "idprefix",
      Set: nil,
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Section{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.ShorthandAttribute{
          Style: nil,
          ID: &asciidoc.ShorthandID{
            Set: asciidoc.Set{
              &asciidoc.String{
                Value: "s1",
              },
            },
          },
          Roles: nil,
          Options: nil,
        },
      },
      Set: asciidoc.Set{
        asciidoc.EmptyLine{
          Text: "",
        },
      },
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "<<Forward>>",
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
      },
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "<<s1>>",
        },
      },
      Level: 1,
    },
    &asciidoc.Section{
      AttributeList: nil,
      Set: nil,
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "Forward",
        },
      },
      Level: 1,
    },
  },
}

var shouldResolveFirstMatchingNaturalXref = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "see ",
    },
    &asciidoc.CrossReference{
      Set: nil,
      ID: "Section Title",
    },
    &asciidoc.NewLine{},
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Section{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.ShorthandAttribute{
          Style: nil,
          ID: &asciidoc.ShorthandID{
            Set: asciidoc.Set{
              &asciidoc.String{
                Value: "s1",
              },
            },
          },
          Roles: nil,
          Options: nil,
        },
      },
      Set: asciidoc.Set{
        asciidoc.EmptyLine{
          Text: "",
        },
      },
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "Section Title",
        },
      },
      Level: 1,
    },
    &asciidoc.Section{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.ShorthandAttribute{
          Style: nil,
          ID: &asciidoc.ShorthandID{
            Set: asciidoc.Set{
              &asciidoc.String{
                Value: "s2",
              },
            },
          },
          Roles: nil,
          Options: nil,
        },
      },
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

var shouldNotMatchNumericCharacterReferencesWhileSearchingForFragmentInXrefTarget = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "see <<Cub => Tiger>>",
    },
    &asciidoc.NewLine{},
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Section{
      AttributeList: nil,
      Set: nil,
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "Cub => Tiger",
        },
      },
      Level: 1,
    },
  },
}

var shouldNotMatchNumericCharacterReferencesInPathOfInterdocumentXref = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "see xref:",
    },
    &asciidoc.CharacterReplacementReference{
      Value: "cpp",
    },
    &asciidoc.String{
      Value: "[",
    },
    &asciidoc.CharacterReplacementReference{
      Value: "cpp",
    },
    &asciidoc.String{
      Value: "].",
    },
    &asciidoc.NewLine{},
  },
}


