package tests

import (
	"testing"

	"github.com/hasty/alchemy/asciidoc"
)

func TestAttributes(t *testing.T) {
	attributesTests.run(t)
}

var attributesTests = parseTests{
	
	{ "creates an attribute by fusing a legacy multi-line value", "asciidoctor/attributes_test_creates_an_attribute_by_fusing_a_legacy_multi_line_value.adoc", createsAnAttributeByFusingALegacyMultiLineValue },
	
	{ "creates an attribute by fusing a multi-line value", "asciidoctor/attributes_test_creates_an_attribute_by_fusing_a_multi_line_value.adoc", createsAnAttributeByFusingAMultiLineValue },
	
	{ "honors line break characters in multi-line values", "asciidoctor/attributes_test_honors_line_break_characters_in_multi_line_values.adoc", honorsLineBreakCharactersInMultiLineValues },
	
	{ "should allow pass macro to surround a multi-line value that contains line breaks", "asciidoctor/attributes_test_should_allow_pass_macro_to_surround_a_multi_line_value_that_contains_line_breaks.adoc", shouldAllowPassMacroToSurroundAMultiLineValueThatContainsLineBreaks },
	
	{ "performs attribute substitution on attribute value", "asciidoctor/attributes_test_performs_attribute_substitution_on_attribute_value.adoc", performsAttributeSubstitutionOnAttributeValue },
	
	{ "resolves attributes inside attribute value within header", "asciidoctor/attributes_test_resolves_attributes_inside_attribute_value_within_header.adoc", resolvesAttributesInsideAttributeValueWithinHeader },
	
	{ "resolves attributes and pass macro inside attribute value outside header", "asciidoctor/attributes_test_resolves_attributes_and_pass_macro_inside_attribute_value_outside_header.adoc", resolvesAttributesAndPassMacroInsideAttributeValueOutsideHeader },
	
	{ "should handle multibyte characters when limiting attribute value size", "asciidoctor/attributes_test_should_handle_multibyte_characters_when_limiting_attribute_value_size.adoc", shouldHandleMultibyteCharactersWhenLimitingAttributeValueSize },
	
	{ "should not mangle multibyte characters when limiting attribute value size", "asciidoctor/attributes_test_should_not_mangle_multibyte_characters_when_limiting_attribute_value_size.adoc", shouldNotMangleMultibyteCharactersWhenLimitingAttributeValueSize },
	
	{ "resolves user-home attribute if safe mode is less than SERVER", "asciidoctor/attributes_test_resolves_user_home_attribute_if_safe_mode_is_less_than_server.adoc", resolvesUserHomeAttributeIfSafeModeIsLessThanServer },
	
	{ "user-home attribute resolves to . if safe mode is SERVER or greater", "asciidoctor/attributes_test_user_home_attribute_resolves_to___if_safe_mode_is_server_or_greater.adoc", userHomeAttributeResolvesToIfSafeModeIsServerOrGreater },
	
	{ "user-home attribute can be overridden by API if safe mode is less than SERVER", "asciidoctor/attributes_test_user_home_attribute_can_be_overridden_by_api_if_safe_mode_is_less_than_server.adoc", userHomeAttributeCanBeOverriddenByApiIfSafeModeIsLessThanServer },
	
	{ "user-home attribute can be overridden by API if safe mode is SERVER or greater", "asciidoctor/attributes_test_user_home_attribute_can_be_overridden_by_api_if_safe_mode_is_server_or_greater.adoc", userHomeAttributeCanBeOverriddenByApiIfSafeModeIsServerOrGreater },
	
	{ "should not recognize pass macro with invalid substitution list in attribute value", "asciidoctor/attributes_test_should_not_recognize_pass_macro_with_invalid_substitution_list_in_attribute_value.adoc", shouldNotRecognizePassMacroWithInvalidSubstitutionListInAttributeValue },
	
	{ "attribute set via API overrides attribute set in document", "asciidoctor/attributes_test_attribute_set_via_api_overrides_attribute_set_in_document.adoc", attributeSetViaApiOverridesAttributeSetInDocument },
	
	{ "backend and doctype attributes are set by default in custom configuration", "asciidoctor/attributes_test_backend_and_doctype_attributes_are_set_by_default_in_custom_configuration.adoc", backendAndDoctypeAttributesAreSetByDefaultInCustomConfiguration },
	
	{ "backend attributes are updated if backend attribute is defined in document and safe mode is less than SERVER", "asciidoctor/attributes_test_backend_attributes_are_updated_if_backend_attribute_is_defined_in_document_and_safe_mode_is_less_than_server.adoc", backendAttributesAreUpdatedIfBackendAttributeIsDefinedInDocumentAndSafeModeIsLessThanServer },
	
	{ "backend attributes defined in document options overrides backend attribute in document", "asciidoctor/attributes_test_backend_attributes_defined_in_document_options_overrides_backend_attribute_in_document.adoc", backendAttributesDefinedInDocumentOptionsOverridesBackendAttributeInDocument },
	
	{ "set_attribute should set attribute if key is not locked", "asciidoctor/attributes_test_set_attribute_should_set_attribute_if_key_is_not_locked.adoc", setAttributeShouldSetAttributeIfKeyIsNotLocked },
	
	{ "convert properly with simple names", "asciidoctor/attributes_test_convert_properly_with_simple_names.adoc", convertProperlyWithSimpleNames },
	
	{ "convert properly with single character name", "asciidoctor/attributes_test_convert_properly_with_single_character_name.adoc", convertProperlyWithSingleCharacterName },
	
	{ "ignores lines with bad attributes if attribute-missing is drop-line", "asciidoctor/attributes_test_ignores_lines_with_bad_attributes_if_attribute_missing_is_drop_line.adoc", ignoresLinesWithBadAttributesIfAttributeMissingIsDropLine },
	
	{ "attribute value gets interpreted when converting", "asciidoctor/attributes_test_attribute_value_gets_interpreted_when_converting.adoc", attributeValueGetsInterpretedWhenConverting },
	
	{ "should not drop line with reference to missing attribute by default", "asciidoctor/attributes_test_should_not_drop_line_with_reference_to_missing_attribute_by_default.adoc", shouldNotDropLineWithReferenceToMissingAttributeByDefault },
	
	{ "should drop line with attribute unassignment by default", "asciidoctor/attributes_test_should_drop_line_with_attribute_unassignment_by_default.adoc", shouldDropLineWithAttributeUnassignmentByDefault },
	
	{ "should not drop line with attribute unassignment if attribute-undefined is drop", "asciidoctor/attributes_test_should_not_drop_line_with_attribute_unassignment_if_attribute_undefined_is_drop.adoc", shouldNotDropLineWithAttributeUnassignmentIfAttributeUndefinedIsDrop },
	
	{ "should drop line that only contains attribute assignment", "asciidoctor/attributes_test_should_drop_line_that_only_contains_attribute_assignment.adoc", shouldDropLineThatOnlyContainsAttributeAssignment },
	
	{ "should drop line that only contains unresolved attribute when attribute-missing is drop", "asciidoctor/attributes_test_should_drop_line_that_only_contains_unresolved_attribute_when_attribute_missing_is_drop.adoc", shouldDropLineThatOnlyContainsUnresolvedAttributeWhenAttributeMissingIsDrop },
	
	{ "substitutes inside unordered list items", "asciidoctor/attributes_test_substitutes_inside_unordered_list_items.adoc", substitutesInsideUnorderedListItems },
	
	{ "interpolates author attribute inside attribute entry in header", "asciidoctor/attributes_test_interpolates_author_attribute_inside_attribute_entry_in_header.adoc", interpolatesAuthorAttributeInsideAttributeEntryInHeader },
	
	{ "interpolates revinfo attribute inside attribute entry in header", "asciidoctor/attributes_test_interpolates_revinfo_attribute_inside_attribute_entry_in_header.adoc", interpolatesRevinfoAttributeInsideAttributeEntryInHeader },
	
	{ "attribute entries can resolve previously defined attributes", "asciidoctor/attributes_test_attribute_entries_can_resolve_previously_defined_attributes.adoc", attributeEntriesCanResolvePreviouslyDefinedAttributes },
	
	{ "should warn if unterminated block comment is detected in document header", "asciidoctor/attributes_test_should_warn_if_unterminated_block_comment_is_detected_in_document_header.adoc", shouldWarnIfUnterminatedBlockCommentIsDetectedInDocumentHeader },
	
	{ "substitutes inside block title", "asciidoctor/attributes_test_substitutes_inside_block_title.adoc", substitutesInsideBlockTitle },
	
	{ "sets attribute until it is deleted", "asciidoctor/attributes_test_sets_attribute_until_it_is_deleted.adoc", setsAttributeUntilItIsDeleted },
	
	{ "should allow compat-mode to be set and unset in middle of document", "asciidoctor/attributes_test_should_allow_compat_mode_to_be_set_and_unset_in_middle_of_document.adoc", shouldAllowCompatModeToBeSetAndUnsetInMiddleOfDocument },
	
	{ "does not disturb attribute-looking things escaped with backslash", "asciidoctor/attributes_test_does_not_disturb_attribute_looking_things_escaped_with_backslash.adoc", doesNotDisturbAttributeLookingThingsEscapedWithBackslash },
	
	{ "does not substitute attributes inside literal blocks", "asciidoctor/attributes_test_does_not_substitute_attributes_inside_literal_blocks.adoc", doesNotSubstituteAttributesInsideLiteralBlocks },
	
	{ "does not show docdir and shows relative docfile if safe mode is SERVER or greater", "asciidoctor/attributes_test_does_not_show_docdir_and_shows_relative_docfile_if_safe_mode_is_server_or_greater.adoc", doesNotShowDocdirAndShowsRelativeDocfileIfSafeModeIsServerOrGreater },
	
	{ "shows absolute docdir and docfile paths if safe mode is less than SERVER", "asciidoctor/attributes_test_shows_absolute_docdir_and_docfile_paths_if_safe_mode_is_less_than_server.adoc", showsAbsoluteDocdirAndDocfilePathsIfSafeModeIsLessThanServer },
	
	{ "assigns attribute defined in attribute reference with set prefix and value", "asciidoctor/attributes_test_assigns_attribute_defined_in_attribute_reference_with_set_prefix_and_value.adoc", assignsAttributeDefinedInAttributeReferenceWithSetPrefixAndValue },
	
	{ "creates counter", "asciidoctor/attributes_test_creates_counter.adoc", createsCounter },
	
	{ "can seed counter to start at A", "asciidoctor/attributes_test_can_seed_counter_to_start_at_a.adoc", canSeedCounterToStartAtA },
	
	{ "increments counter with positive numeric value", "asciidoctor/attributes_test_increments_counter_with_positive_numeric_value.adoc", incrementsCounterWithPositiveNumericValue },
	
	{ "increments counter with negative numeric value", "asciidoctor/attributes_test_increments_counter_with_negative_numeric_value.adoc", incrementsCounterWithNegativeNumericValue },
	
	{ "increments counter with ASCII character value", "asciidoctor/attributes_test_increments_counter_with_ascii_character_value.adoc", incrementsCounterWithAsciiCharacterValue },
	
	{ "increments counter with non-ASCII character value", "asciidoctor/attributes_test_increments_counter_with_non_ascii_character_value.adoc", incrementsCounterWithNonAsciiCharacterValue },
	
	{ "increments counter with emoji character value", "asciidoctor/attributes_test_increments_counter_with_emoji_character_value.adoc", incrementsCounterWithEmojiCharacterValue },
	
	{ "increments counter with multi-character value", "asciidoctor/attributes_test_increments_counter_with_multi_character_value.adoc", incrementsCounterWithMultiCharacterValue },
	
	{ "counter uses 0 as seed value if seed attribute is nil", "asciidoctor/attributes_test_counter_uses_0_as_seed_value_if_seed_attribute_is_nil.adoc", counterUses0AsSeedValueIfSeedAttributeIsNil },
	
	{ "counter value can be reset by attribute entry", "asciidoctor/attributes_test_counter_value_can_be_reset_by_attribute_entry.adoc", counterValueCanBeResetByAttributeEntry },
	
	{ "counter value can be advanced by attribute entry", "asciidoctor/attributes_test_counter_value_can_be_advanced_by_attribute_entry.adoc", counterValueCanBeAdvancedByAttributeEntry },
	
	{ "nested document should use counter from parent document", "asciidoctor/attributes_test_nested_document_should_use_counter_from_parent_document.adoc", nestedDocumentShouldUseCounterFromParentDocument },
	
	{ "should not allow counter to modify locked attribute", "asciidoctor/attributes_test_should_not_allow_counter_to_modify_locked_attribute.adoc", shouldNotAllowCounterToModifyLockedAttribute },
	
	{ "should not allow counter2 to modify locked attribute", "asciidoctor/attributes_test_should_not_allow_counter_2_to_modify_locked_attribute.adoc", shouldNotAllowCounter2ToModifyLockedAttribute },
	
	{ "should not allow counter to modify built-in locked attribute", "asciidoctor/attributes_test_should_not_allow_counter_to_modify_built_in_locked_attribute.adoc", shouldNotAllowCounterToModifyBuiltInLockedAttribute },
	
	{ "should not allow counter2 to modify built-in locked attribute", "asciidoctor/attributes_test_should_not_allow_counter_2_to_modify_built_in_locked_attribute.adoc", shouldNotAllowCounter2ToModifyBuiltInLockedAttribute },
	
	{ "parses named attribute with valid name", "asciidoctor/attributes_test_parses_named_attribute_with_valid_name.adoc", parsesNamedAttributeWithValidName },
	
	{ "does not parse named attribute if name is invalid", "asciidoctor/attributes_test_does_not_parse_named_attribute_if_name_is_invalid.adoc", doesNotParseNamedAttributeIfNameIsInvalid },
	
	{ "positional attributes assigned to block", "asciidoctor/attributes_test_positional_attributes_assigned_to_block.adoc", positionalAttributesAssignedToBlock },
	
	{ "normal substitutions are performed on single-quoted positional attribute", "asciidoctor/attributes_test_normal_substitutions_are_performed_on_single_quoted_positional_attribute.adoc", normalSubstitutionsArePerformedOnSingleQuotedPositionalAttribute },
	
	{ "normal substitutions are performed on single-quoted named attribute", "asciidoctor/attributes_test_normal_substitutions_are_performed_on_single_quoted_named_attribute.adoc", normalSubstitutionsArePerformedOnSingleQuotedNamedAttribute },
	
	{ "normal substitutions are performed once on single-quoted named title attribute", "asciidoctor/attributes_test_normal_substitutions_are_performed_once_on_single_quoted_named_title_attribute.adoc", normalSubstitutionsArePerformedOnceOnSingleQuotedNamedTitleAttribute },
	
	{ "attribute list may not begin with space", "asciidoctor/attributes_test_attribute_list_may_not_begin_with_space.adoc", attributeListMayNotBeginWithSpace },
	
	{ "attribute list may begin with comma", "asciidoctor/attributes_test_attribute_list_may_begin_with_comma.adoc", attributeListMayBeginWithComma },
	
	{ "first attribute in list may be double quoted", "asciidoctor/attributes_test_first_attribute_in_list_may_be_double_quoted.adoc", firstAttributeInListMayBeDoubleQuoted },
	
	{ "first attribute in list may be single quoted", "asciidoctor/attributes_test_first_attribute_in_list_may_be_single_quoted.adoc", firstAttributeInListMayBeSingleQuoted },
	
	{ "attribute with value None without quotes is ignored", "asciidoctor/attributes_test_attribute_with_value_none_without_quotes_is_ignored.adoc", attributeWithValueNoneWithoutQuotesIsIgnored },
	
	{ "role? returns true if role is assigned", "asciidoctor/attributes_test_role?_returns_true_if_role_is_assigned.adoc", roleReturnsTrueIfRoleIsAssigned },
	
	{ "role? does not return true if role attribute is set on document", "asciidoctor/attributes_test_role?_does_not_return_true_if_role_attribute_is_set_on_document.adoc", roleDoesNotReturnTrueIfRoleAttributeIsSetOnDocument },
	
	{ "role? can check for exact role name match", "asciidoctor/attributes_test_role?_can_check_for_exact_role_name_match.adoc", roleCanCheckForExactRoleNameMatch },
	
	{ "has_role? can check for presence of role name", "asciidoctor/attributes_test_has_role?_can_check_for_presence_of_role_name.adoc", hasRoleCanCheckForPresenceOfRoleName },
	
	{ "has_role? does not look for role defined as document attribute", "asciidoctor/attributes_test_has_role?_does_not_look_for_role_defined_as_document_attribute.adoc", hasRoleDoesNotLookForRoleDefinedAsDocumentAttribute },
	
	{ "roles returns array of role names", "asciidoctor/attributes_test_roles_returns_array_of_role_names.adoc", rolesReturnsArrayOfRoleNames },
	
	{ "roles returns empty array if role attribute is not set", "asciidoctor/attributes_test_roles_returns_empty_array_if_role_attribute_is_not_set.adoc", rolesReturnsEmptyArrayIfRoleAttributeIsNotSet },
	
	{ "roles= sets the role attribute on the node", "asciidoctor/attributes_test_roles=_sets_the_role_attribute_on_the_node.adoc", rolesSetsTheRoleAttributeOnTheNode },
	
	{ "id, role and options attributes can be specified on block style using shorthand syntax", "asciidoctor/attributes_test_id_role_and_options_attributes_can_be_specified_on_block_style_using_shorthand_syntax.adoc", idRoleAndOptionsAttributesCanBeSpecifiedOnBlockStyleUsingShorthandSyntax },
	
	{ "id, role and options attributes can be specified using shorthand syntax on block style using multiple block attribute lines", "asciidoctor/attributes_test_id_role_and_options_attributes_can_be_specified_using_shorthand_syntax_on_block_style_using_multiple_block_attribute_lines.adoc", idRoleAndOptionsAttributesCanBeSpecifiedUsingShorthandSyntaxOnBlockStyleUsingMultipleBlockAttributeLines },
	
	{ "multiple roles and options can be specified in block style using shorthand syntax", "asciidoctor/attributes_test_multiple_roles_and_options_can_be_specified_in_block_style_using_shorthand_syntax.adoc", multipleRolesAndOptionsCanBeSpecifiedInBlockStyleUsingShorthandSyntax },
	
	{ "options specified using shorthand syntax on block style across multiple lines should be additive", "asciidoctor/attributes_test_options_specified_using_shorthand_syntax_on_block_style_across_multiple_lines_should_be_additive.adoc", optionsSpecifiedUsingShorthandSyntaxOnBlockStyleAcrossMultipleLinesShouldBeAdditive },
	
	{ "roles specified using shorthand syntax on block style across multiple lines should be additive", "asciidoctor/attributes_test_roles_specified_using_shorthand_syntax_on_block_style_across_multiple_lines_should_be_additive.adoc", rolesSpecifiedUsingShorthandSyntaxOnBlockStyleAcrossMultipleLinesShouldBeAdditive },
	
	{ "setting a role using the role attribute replaces any existing roles", "asciidoctor/attributes_test_setting_a_role_using_the_role_attribute_replaces_any_existing_roles.adoc", settingARoleUsingTheRoleAttributeReplacesAnyExistingRoles },
	
	{ "setting a role using the shorthand syntax on block style should not clear the ID", "asciidoctor/attributes_test_setting_a_role_using_the_shorthand_syntax_on_block_style_should_not_clear_the_id.adoc", settingARoleUsingTheShorthandSyntaxOnBlockStyleShouldNotClearTheId },
	
	{ "a role can be added using add_role when the node has no roles", "asciidoctor/attributes_test_a_role_can_be_added_using_add_role_when_the_node_has_no_roles.adoc", aRoleCanBeAddedUsingAddRoleWhenTheNodeHasNoRoles },
	
	{ "a role is not added using add_role if the node already has that role", "asciidoctor/attributes_test_a_role_is_not_added_using_add_role_if_the_node_already_has_that_role.adoc", aRoleIsNotAddedUsingAddRoleIfTheNodeAlreadyHasThatRole },
	
	{ "an existing role can be removed using remove_role", "asciidoctor/attributes_test_an_existing_role_can_be_removed_using_remove_role.adoc", anExistingRoleCanBeRemovedUsingRemoveRole },
	
	{ "roles are removed when last role is removed using remove_role", "asciidoctor/attributes_test_roles_are_removed_when_last_role_is_removed_using_remove_role.adoc", rolesAreRemovedWhenLastRoleIsRemovedUsingRemoveRole },
	
	{ "roles are not changed when a non-existent role is removed using remove_role", "asciidoctor/attributes_test_roles_are_not_changed_when_a_non_existent_role_is_removed_using_remove_role.adoc", rolesAreNotChangedWhenANonExistentRoleIsRemovedUsingRemoveRole },
	
	{ "roles are not changed when using remove_role if the node has no roles", "asciidoctor/attributes_test_roles_are_not_changed_when_using_remove_role_if_the_node_has_no_roles.adoc", rolesAreNotChangedWhenUsingRemoveRoleIfTheNodeHasNoRoles },
	
	{ "id and role attributes can be specified on section style using shorthand syntax", "asciidoctor/attributes_test_id_and_role_attributes_can_be_specified_on_section_style_using_shorthand_syntax.adoc", idAndRoleAttributesCanBeSpecifiedOnSectionStyleUsingShorthandSyntax },
	
	{ "id attribute specified using shorthand syntax should not create a special section", "asciidoctor/attributes_test_id_attribute_specified_using_shorthand_syntax_should_not_create_a_special_section.adoc", idAttributeSpecifiedUsingShorthandSyntaxShouldNotCreateASpecialSection },
	
	{ "Block attributes are additive", "asciidoctor/attributes_test_block_attributes_are_additive.adoc", blockAttributesAreAdditive },
	
	{ "Last wins for id attribute", "asciidoctor/attributes_test_last_wins_for_id_attribute.adoc", lastWinsForIdAttribute },
	
	{ "trailing block attributes transfer to the following section", "asciidoctor/attributes_test_trailing_block_attributes_transfer_to_the_following_section.adoc", trailingBlockAttributesTransferToTheFollowingSection },
	
}


var createsAnAttributeByFusingALegacyMultiLineValue = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "description",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "This is the first      ",
        },
        &asciidoc.LineBreak{},
        &asciidoc.NewLine{},
        &asciidoc.String{
          Value: "              Ruby implementation of ",
        },
        &asciidoc.LineBreak{},
        &asciidoc.NewLine{},
        &asciidoc.String{
          Value: "              AsciiDoc.",
        },
      },
    },
  },
}

var createsAnAttributeByFusingAMultiLineValue = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "description",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "This is the first \\",
        },
      },
    },
    &asciidoc.String{
      Value: "              Ruby implementation of \\",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "              AsciiDoc.",
    },
    &asciidoc.NewLine{},
  },
}

var honorsLineBreakCharactersInMultiLineValues = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "signature",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "Linus Torvalds + \\",
        },
      },
    },
    &asciidoc.String{
      Value: "Linux Hacker + \\",
    },
    &asciidoc.NewLine{},
    asciidoc.Email{
      Address: "linus.torvalds@example.com",
    },
    &asciidoc.NewLine{},
  },
}

var shouldAllowPassMacroToSurroundAMultiLineValueThatContainsLineBreaks = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "signature",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "pass:a[{author} + \\",
        },
      },
    },
    &asciidoc.UserAttributeReference{
      Value: "title",
    },
    &asciidoc.String{
      Value: " + \\",
    },
    &asciidoc.NewLine{},
    &asciidoc.UserAttributeReference{
      Value: "email",
    },
    &asciidoc.String{
      Value: "]",
    },
    &asciidoc.NewLine{},
  },
}

var performsAttributeSubstitutionOnAttributeValue = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "release",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "Asciidoctor ",
        },
        &asciidoc.LineBreak{},
        &asciidoc.NewLine{},
        &asciidoc.String{
          Value: "          {version}",
        },
      },
    },
  },
}

var resolvesAttributesInsideAttributeValueWithinHeader = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Section{
      AttributeList: nil,
      Set: asciidoc.Set{
        &asciidoc.AttributeEntry{
          Name: "big",
          Set: asciidoc.Set{
            &asciidoc.String{
              Value: "big",
            },
          },
        },
        &asciidoc.AttributeEntry{
          Name: "bigfoot",
          Set: asciidoc.Set{
            &asciidoc.String{
              Value: "{big}foot",
            },
          },
        },
        asciidoc.EmptyLine{
          Text: "",
        },
        &asciidoc.UserAttributeReference{
          Value: "bigfoot",
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

var resolvesAttributesAndPassMacroInsideAttributeValueOutsideHeader = &asciidoc.Document{
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
          Value: "content",
        },
        &asciidoc.NewLine{},
        asciidoc.EmptyLine{
          Text: "",
        },
        &asciidoc.AttributeEntry{
          Name: "big",
          Set: asciidoc.Set{
            &asciidoc.String{
              Value: "pass:a,q[_big_]",
            },
          },
        },
        &asciidoc.AttributeEntry{
          Name: "bigfoot",
          Set: asciidoc.Set{
            &asciidoc.String{
              Value: "{big}foot",
            },
          },
        },
        &asciidoc.UserAttributeReference{
          Value: "bigfoot",
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

var shouldHandleMultibyteCharactersWhenLimitingAttributeValueSize = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "name",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "日本語",
        },
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.UserAttributeReference{
      Value: "name",
    },
    &asciidoc.NewLine{},
  },
}

var shouldNotMangleMultibyteCharactersWhenLimitingAttributeValueSize = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "name",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "日本語",
        },
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.UserAttributeReference{
      Value: "name",
    },
    &asciidoc.NewLine{},
  },
}

var resolvesUserHomeAttributeIfSafeModeIsLessThanServer = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "imagesdir",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "{user-home}/etc/images",
        },
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.UserAttributeReference{
      Value: "imagesdir",
    },
    &asciidoc.NewLine{},
  },
}

var userHomeAttributeResolvesToIfSafeModeIsServerOrGreater = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "imagesdir",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "{user-home}/etc/images",
        },
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.UserAttributeReference{
      Value: "imagesdir",
    },
    &asciidoc.NewLine{},
  },
}

var userHomeAttributeCanBeOverriddenByApiIfSafeModeIsLessThanServer = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "Go ",
    },
    &asciidoc.UserAttributeReference{
      Value: "user-home",
    },
    &asciidoc.String{
      Value: "!",
    },
    &asciidoc.NewLine{},
  },
}

var userHomeAttributeCanBeOverriddenByApiIfSafeModeIsServerOrGreater = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "Go ",
    },
    &asciidoc.UserAttributeReference{
      Value: "user-home",
    },
    &asciidoc.String{
      Value: "!",
    },
    &asciidoc.NewLine{},
  },
}

var shouldNotRecognizePassMacroWithInvalidSubstitutionListInAttributeValue = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "holygrail",
      Set: nil,
    },
    &asciidoc.IfDef{
      Attributes: asciidoc.AttributeNames{
        "holygrail",
      },
      Union: 0,
    },
    &asciidoc.String{
      Value: "The holy grail has been found!",
    },
    &asciidoc.NewLine{},
    &asciidoc.EndIf{
      Attributes: asciidoc.AttributeNames{
        "holygrail",
      },
      Union: 0,
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeReset{
      Name: "holygrail",
    },
    &asciidoc.IfNDef{
      Attributes: asciidoc.AttributeNames{
        "holygrail",
      },
      Union: 0,
    },
    &asciidoc.String{
      Value: "Buggers! What happened to the grail?",
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

var attributeSetViaApiOverridesAttributeSetInDocument = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Section{
      AttributeList: nil,
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "Author Name",
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
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "Document Title",
        },
      },
      Level: 0,
    },
  },
}

var backendAndDoctypeAttributesAreSetByDefaultInCustomConfiguration = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Section{
      AttributeList: nil,
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "Author Name",
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
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "Document Title",
        },
      },
      Level: 0,
    },
  },
}

var backendAttributesAreUpdatedIfBackendAttributeIsDefinedInDocumentAndSafeModeIsLessThanServer = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Section{
      AttributeList: nil,
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "Author Name",
        },
        &asciidoc.NewLine{},
        &asciidoc.AttributeEntry{
          Name: "backend",
          Set: asciidoc.Set{
            &asciidoc.String{
              Value: "docbook",
            },
          },
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

var backendAttributesDefinedInDocumentOptionsOverridesBackendAttributeInDocument = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "uri",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "http://example.org",
        },
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.UserAttributeReference{
      Value: "uri",
    },
    &asciidoc.NewLine{},
  },
}

var setAttributeShouldSetAttributeIfKeyIsNotLocked = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "#attributes                               |toc|toc-position|toc-placement|toc-class",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "toc                                       |   |nil         |auto         |nil",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "toc=header                                |   |nil         |auto         |nil",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "toc=beeboo                                |   |nil         |auto         |nil",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "toc=left                                  |   |left        |auto         |toc2",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "toc2                                      |   |left        |auto         |toc2",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "toc=right                                 |   |right       |auto         |toc2",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "toc=preamble                              |   |content     |preamble     |nil",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "toc=macro                                 |   |content     |macro        |nil",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "toc toc-placement=macro toc-position=left |   |content     |macro        |nil",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "toc toc-placement!                        |   |content     |macro        |nil",
    },
    &asciidoc.NewLine{},
  },
}

var convertProperlyWithSimpleNames = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "He-Man",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "The most powerful man in the universe",
        },
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "He-Man: ",
    },
    &asciidoc.UserAttributeReference{
      Value: "He-Man",
    },
    &asciidoc.NewLine{},
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "She-Ra: ",
    },
    &asciidoc.UserAttributeReference{
      Value: "She-Ra",
    },
    &asciidoc.NewLine{},
  },
}

var convertProperlyWithSingleCharacterName = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "Main Header",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "===========",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: ":My frog: Tanglefoot",
    },
    &asciidoc.NewLine{},
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "Yo, ",
    },
    &asciidoc.UserAttributeReference{
      Value: "myfrog",
    },
    &asciidoc.String{
      Value: "!",
    },
    &asciidoc.NewLine{},
  },
}

var ignoresLinesWithBadAttributesIfAttributeMissingIsDropLine = &asciidoc.Document{
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
    &asciidoc.String{
      Value: "This is",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "blah blah ",
    },
    &asciidoc.UserAttributeReference{
      Value: "foobarbaz",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "all there is.",
    },
    &asciidoc.NewLine{},
  },
}

var attributeValueGetsInterpretedWhenConverting = &asciidoc.Document{
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
    &asciidoc.String{
      Value: "Line 1: This line should appear in the output.",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "Line 2: Oh no, a ",
    },
    &asciidoc.UserAttributeReference{
      Value: "bogus-attribute",
    },
    &asciidoc.String{
      Value: "! This line should not appear in the output.",
    },
    &asciidoc.NewLine{},
  },
}

var shouldNotDropLineWithReferenceToMissingAttributeByDefault = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "Line 1: This line should appear in the output.",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "Line 2: A ",
    },
    &asciidoc.UserAttributeReference{
      Value: "bogus-attribute",
    },
    &asciidoc.String{
      Value: "! This time, this line should appear in the output.",
    },
    &asciidoc.NewLine{},
  },
}

var shouldDropLineWithAttributeUnassignmentByDefault = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "a",
      Set: nil,
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "Line 1: This line should appear in the output.",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "Line 2: {set:a!}This line should not appear in the output.",
    },
    &asciidoc.NewLine{},
  },
}

var shouldNotDropLineWithAttributeUnassignmentIfAttributeUndefinedIsDrop = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "attribute-undefined",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "drop",
        },
      },
    },
    &asciidoc.AttributeEntry{
      Name: "a",
      Set: nil,
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "Line 1: This line should appear in the output.",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "Line 2: {set:a!}This line should appear in the output.",
    },
    &asciidoc.NewLine{},
  },
}

var shouldDropLineThatOnlyContainsAttributeAssignment = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "Line 1",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "{set:a}",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "Line 2",
    },
    &asciidoc.NewLine{},
  },
}

var shouldDropLineThatOnlyContainsUnresolvedAttributeWhenAttributeMissingIsDrop = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "Line 1",
    },
    &asciidoc.NewLine{},
    &asciidoc.UserAttributeReference{
      Value: "unresolved",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "Line 2",
    },
    &asciidoc.NewLine{},
  },
}

var substitutesInsideUnorderedListItems = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Section{
      AttributeList: nil,
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "Author Name",
        },
        &asciidoc.NewLine{},
        &asciidoc.AttributeEntry{
          Name: "attribute-a",
          Set: asciidoc.Set{
            &asciidoc.String{
              Value: "value",
            },
          },
        },
        &asciidoc.AttributeEntry{
          Name: "attribute-b",
          Set: asciidoc.Set{
            &asciidoc.String{
              Value: "{attribute-a}",
            },
          },
        },
        asciidoc.EmptyLine{
          Text: "",
        },
        &asciidoc.String{
          Value: "preamble",
        },
        &asciidoc.NewLine{},
      },
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "Title",
        },
      },
      Level: 0,
    },
  },
}

var interpolatesAuthorAttributeInsideAttributeEntryInHeader = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Section{
      AttributeList: nil,
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "Author Name",
        },
        &asciidoc.NewLine{},
        &asciidoc.AttributeEntry{
          Name: "name",
          Set: asciidoc.Set{
            &asciidoc.String{
              Value: "{author}",
            },
          },
        },
        asciidoc.EmptyLine{
          Text: "",
        },
        &asciidoc.String{
          Value: "preamble",
        },
        &asciidoc.NewLine{},
      },
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "Title",
        },
      },
      Level: 0,
    },
  },
}

var interpolatesRevinfoAttributeInsideAttributeEntryInHeader = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Section{
      AttributeList: nil,
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "Author Name",
        },
        &asciidoc.NewLine{},
        &asciidoc.String{
          Value: "2013-01-01",
        },
        &asciidoc.NewLine{},
        &asciidoc.AttributeEntry{
          Name: "date",
          Set: asciidoc.Set{
            &asciidoc.String{
              Value: "{revdate}",
            },
          },
        },
        asciidoc.EmptyLine{
          Text: "",
        },
        &asciidoc.String{
          Value: "preamble",
        },
        &asciidoc.NewLine{},
      },
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "Title",
        },
      },
      Level: 0,
    },
  },
}

var attributeEntriesCanResolvePreviouslyDefinedAttributes = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Section{
      AttributeList: nil,
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "Author Name",
        },
        &asciidoc.NewLine{},
        &asciidoc.String{
          Value: "v1.0, 2010-01-01: First release!",
        },
        &asciidoc.NewLine{},
        &asciidoc.AttributeEntry{
          Name: "a",
          Set: asciidoc.Set{
            &asciidoc.String{
              Value: "value",
            },
          },
        },
        &asciidoc.AttributeEntry{
          Name: "a2",
          Set: asciidoc.Set{
            &asciidoc.String{
              Value: "{a}",
            },
          },
        },
        &asciidoc.AttributeEntry{
          Name: "revdate2",
          Set: asciidoc.Set{
            &asciidoc.String{
              Value: "{revdate}",
            },
          },
        },
        asciidoc.EmptyLine{
          Text: "",
        },
        &asciidoc.UserAttributeReference{
          Value: "a",
        },
        &asciidoc.String{
          Value: " == ",
        },
        &asciidoc.UserAttributeReference{
          Value: "a2",
        },
        &asciidoc.NewLine{},
        asciidoc.EmptyLine{
          Text: "",
        },
        &asciidoc.UserAttributeReference{
          Value: "revdate",
        },
        &asciidoc.String{
          Value: " == ",
        },
        &asciidoc.UserAttributeReference{
          Value: "revdate2",
        },
        &asciidoc.NewLine{},
      },
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "Title",
        },
      },
      Level: 0,
    },
  },
}

var shouldWarnIfUnterminatedBlockCommentIsDetectedInDocumentHeader = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Section{
      AttributeList: nil,
      Set: asciidoc.Set{
        &asciidoc.AttributeEntry{
          Name: "foo",
          Set: asciidoc.Set{
            &asciidoc.String{
              Value: "bar",
            },
          },
        },
        &asciidoc.String{
          Value: "////",
        },
        &asciidoc.NewLine{},
        &asciidoc.AttributeEntry{
          Name: "hey",
          Set: asciidoc.Set{
            &asciidoc.String{
              Value: "there",
            },
          },
        },
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
          Value: "Document Title",
        },
      },
      Level: 0,
    },
  },
}

var substitutesInsideBlockTitle = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "gem_name",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "asciidoctor",
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
              Value: "Require the +",
            },
            &asciidoc.UserAttributeReference{
              Value: "gem_name",
            },
            &asciidoc.String{
              Value: "+ gem",
            },
          },
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "To use ",
        },
        &asciidoc.UserAttributeReference{
          Value: "gem_name",
        },
        &asciidoc.String{
          Value: ", the first thing to do is to import it in your Ruby source file.",
        },
        &asciidoc.NewLine{},
      },
      Admonition: 0,
    },
  },
}

var setsAttributeUntilItIsDeleted = &asciidoc.Document{
  Set: asciidoc.Set{
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
    &asciidoc.String{
      Value: "Crossing the ",
    },
    &asciidoc.UserAttributeReference{
      Value: "foo",
    },
    &asciidoc.String{
      Value: ".",
    },
    &asciidoc.NewLine{},
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeReset{
      Name: "foo",
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "Belly up to the ",
    },
    &asciidoc.UserAttributeReference{
      Value: "foo",
    },
    &asciidoc.String{
      Value: ".",
    },
    &asciidoc.NewLine{},
  },
}

var shouldAllowCompatModeToBeSetAndUnsetInMiddleOfDocument = &asciidoc.Document{
  Set: asciidoc.Set{
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
        &asciidoc.AnchorAttribute{
          ID: &asciidoc.String{
            Value: "paragraph-a",
          },
          Label: nil,
        },
      },
      Set: asciidoc.Set{
        &asciidoc.Monospace{
          AttributeList: nil,
          Set: asciidoc.Set{
            &asciidoc.UserAttributeReference{
              Value: "foo",
            },
          },
        },
        &asciidoc.NewLine{},
      },
      Admonition: 0,
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeReset{
      Name: "compat-mode",
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Paragraph{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.AnchorAttribute{
          ID: &asciidoc.String{
            Value: "paragraph-b",
          },
          Label: nil,
        },
      },
      Set: asciidoc.Set{
        &asciidoc.Monospace{
          AttributeList: nil,
          Set: asciidoc.Set{
            &asciidoc.UserAttributeReference{
              Value: "foo",
            },
          },
        },
        &asciidoc.NewLine{},
      },
      Admonition: 0,
    },
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
    &asciidoc.Paragraph{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.AnchorAttribute{
          ID: &asciidoc.String{
            Value: "paragraph-c",
          },
          Label: nil,
        },
      },
      Set: asciidoc.Set{
        &asciidoc.Monospace{
          AttributeList: nil,
          Set: asciidoc.Set{
            &asciidoc.UserAttributeReference{
              Value: "foo",
            },
          },
        },
        &asciidoc.NewLine{},
      },
      Admonition: 0,
    },
  },
}

var doesNotDisturbAttributeLookingThingsEscapedWithBackslash = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "forecast",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "snow",
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
        "puts 'The forecast for today is {forecast}'",
      },
    },
  },
}

var doesNotSubstituteAttributesInsideLiteralBlocks = &asciidoc.Document{
  Set: asciidoc.Set{
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
    &asciidoc.LiteralBlock{
      AttributeList: nil,
      Delimiter: asciidoc.Delimiter{
        Type: 6,
        Length: 4,
      },
      LineList: asciidoc.LineList{
        "You insert the text {foo} to expand the value",
        "of the attribute named foo in your document.",
      },
    },
  },
}

var doesNotShowDocdirAndShowsRelativeDocfileIfSafeModeIsServerOrGreater = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.UnorderedListItem{
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "docdir: ",
        },
        &asciidoc.UserAttributeReference{
          Value: "docdir",
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
          Value: "docfile: ",
        },
        &asciidoc.UserAttributeReference{
          Value: "docfile",
        },
      },
      AttributeList: nil,
      Indent: "",
      Marker: "*",
      Checklist: 0,
    },
  },
}

var showsAbsoluteDocdirAndDocfilePathsIfSafeModeIsLessThanServer = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.UnorderedListItem{
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "docdir: ",
        },
        &asciidoc.UserAttributeReference{
          Value: "docdir",
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
          Value: "docfile: ",
        },
        &asciidoc.UserAttributeReference{
          Value: "docfile",
        },
      },
      AttributeList: nil,
      Indent: "",
      Marker: "*",
      Checklist: 0,
    },
  },
}

var assignsAttributeDefinedInAttributeReferenceWithSetPrefixAndValue = &asciidoc.Document{
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
    &asciidoc.AttributeEntry{
      Name: "foo",
      Set: nil,
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "{set:foo!}",
    },
    &asciidoc.NewLine{},
    &asciidoc.UserAttributeReference{
      Value: "foo",
    },
    &asciidoc.String{
      Value: "yes",
    },
    &asciidoc.NewLine{},
  },
}

var createsCounter = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "mycounter",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "0",
        },
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Counter{
      Name: "mycounter",
      InitialValue: "",
      Display: true,
    },
    &asciidoc.NewLine{},
  },
}

var canSeedCounterToStartAtA = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "mycounter",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "@",
        },
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Counter{
      Name: "mycounter",
      InitialValue: "",
      Display: true,
    },
    &asciidoc.NewLine{},
  },
}

var incrementsCounterWithPositiveNumericValue = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.StemBlock{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.NamedAttribute{
          Name: "subs",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "attributes",
            },
          },
          Quote: 0,
        },
      },
      Delimiter: asciidoc.Delimiter{
        Type: 9,
        Length: 4,
      },
      LineList: asciidoc.LineList{
        "{counter:mycounter:1}",
        "{counter:mycounter}",
        "{counter:mycounter}",
        "{mycounter}",
      },
    },
  },
}

var incrementsCounterWithNegativeNumericValue = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.StemBlock{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.NamedAttribute{
          Name: "subs",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "attributes",
            },
          },
          Quote: 0,
        },
      },
      Delimiter: asciidoc.Delimiter{
        Type: 9,
        Length: 4,
      },
      LineList: asciidoc.LineList{
        "{counter:mycounter:-2}",
        "{counter:mycounter}",
        "{counter:mycounter}",
        "{mycounter}",
      },
    },
  },
}

var incrementsCounterWithAsciiCharacterValue = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.StemBlock{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.NamedAttribute{
          Name: "subs",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "attributes",
            },
          },
          Quote: 0,
        },
      },
      Delimiter: asciidoc.Delimiter{
        Type: 9,
        Length: 4,
      },
      LineList: asciidoc.LineList{
        "{counter:mycounter:A}",
        "{counter:mycounter}",
        "{counter:mycounter}",
        "{mycounter}",
      },
    },
  },
}

var incrementsCounterWithNonAsciiCharacterValue = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.StemBlock{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.NamedAttribute{
          Name: "subs",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "attributes",
            },
          },
          Quote: 0,
        },
      },
      Delimiter: asciidoc.Delimiter{
        Type: 9,
        Length: 4,
      },
      LineList: asciidoc.LineList{
        "{counter:mycounter:é}",
        "{counter:mycounter}",
        "{counter:mycounter}",
        "{mycounter}",
      },
    },
  },
}

var incrementsCounterWithEmojiCharacterValue = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.StemBlock{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.NamedAttribute{
          Name: "subs",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "attributes",
            },
          },
          Quote: 0,
        },
      },
      Delimiter: asciidoc.Delimiter{
        Type: 9,
        Length: 4,
      },
      LineList: asciidoc.LineList{
        "{counter:smiley:😋}",
        "{counter:smiley}",
        "{counter:smiley}",
        "{smiley}",
      },
    },
  },
}

var incrementsCounterWithMultiCharacterValue = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.StemBlock{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.NamedAttribute{
          Name: "subs",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "attributes",
            },
          },
          Quote: 0,
        },
      },
      Delimiter: asciidoc.Delimiter{
        Type: 9,
        Length: 4,
      },
      LineList: asciidoc.LineList{
        "{counter:math:1x}",
        "{counter:math}",
        "{counter:math}",
        "{math}",
      },
    },
  },
}

var counterUses0AsSeedValueIfSeedAttributeIsNil = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "mycounter",
      Set: nil,
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Counter{
      Name: "mycounter",
      InitialValue: "",
      Display: true,
    },
    &asciidoc.NewLine{},
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.UserAttributeReference{
      Value: "mycounter",
    },
    &asciidoc.NewLine{},
  },
}

var counterValueCanBeResetByAttributeEntry = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "mycounter",
      Set: nil,
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "before: ",
    },
    &asciidoc.Counter{
      Name: "mycounter",
      InitialValue: "",
      Display: true,
    },
    &asciidoc.String{
      Value: " ",
    },
    &asciidoc.Counter{
      Name: "mycounter",
      InitialValue: "",
      Display: true,
    },
    &asciidoc.String{
      Value: " ",
    },
    &asciidoc.Counter{
      Name: "mycounter",
      InitialValue: "",
      Display: true,
    },
    &asciidoc.NewLine{},
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeReset{
      Name: "mycounter",
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "after: ",
    },
    &asciidoc.Counter{
      Name: "mycounter",
      InitialValue: "",
      Display: true,
    },
    &asciidoc.NewLine{},
  },
}

var counterValueCanBeAdvancedByAttributeEntry = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "before: ",
    },
    &asciidoc.Counter{
      Name: "mycounter",
      InitialValue: "",
      Display: true,
    },
    &asciidoc.NewLine{},
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "mycounter",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "10",
        },
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "after: ",
    },
    &asciidoc.Counter{
      Name: "mycounter",
      InitialValue: "",
      Display: true,
    },
    &asciidoc.NewLine{},
  },
}

var nestedDocumentShouldUseCounterFromParentDocument = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.BlockImage{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TitleAttribute{
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Title for Foo",
            },
          },
        },
      },
      Path: asciidoc.Set{
        &asciidoc.String{
          Value: "foo.jpg",
        },
      },
    },
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
                &asciidoc.NewLine{},
                &asciidoc.BlockImage{
                  AttributeList: asciidoc.AttributeList{
                    &asciidoc.TitleAttribute{
                      Val: asciidoc.Set{
                        &asciidoc.String{
                          Value: "Title for Bar",
                        },
                      },
                    },
                  },
                  Path: asciidoc.Set{
                    &asciidoc.String{
                      Value: "bar.jpg",
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
                &asciidoc.NewLine{},
                &asciidoc.String{
                  Value: ".Title for Baz",
                },
                &asciidoc.NewLine{},
                &asciidoc.String{
                  Value: "image::baz.jpg[]",
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
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.BlockImage{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.TitleAttribute{
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "Title for Qux",
            },
          },
        },
      },
      Path: asciidoc.Set{
        &asciidoc.String{
          Value: "qux.jpg",
        },
      },
    },
  },
}

var shouldNotAllowCounterToModifyLockedAttribute = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "{counter:foo:ignored} is not ",
    },
    &asciidoc.UserAttributeReference{
      Value: "foo",
    },
    &asciidoc.NewLine{},
  },
}

var shouldNotAllowCounter2ToModifyLockedAttribute = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "{counter2:foo:ignored}",
    },
    &asciidoc.UserAttributeReference{
      Value: "foo",
    },
    &asciidoc.NewLine{},
  },
}

var shouldNotAllowCounterToModifyBuiltInLockedAttribute = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Counter{
      Name: "max-include-depth",
      InitialValue: "",
      Display: true,
    },
    &asciidoc.String{
      Value: " is one more than ",
    },
    &asciidoc.UserAttributeReference{
      Value: "max-include-depth",
    },
    &asciidoc.NewLine{},
  },
}

var shouldNotAllowCounter2ToModifyBuiltInLockedAttribute = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Counter{
      Name: "max-include-depth",
      InitialValue: "",
      Display: false,
    },
    &asciidoc.UserAttributeReference{
      Value: "max-include-depth",
    },
    &asciidoc.NewLine{},
  },
}

var parsesNamedAttributeWithValidName = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "[normal,foo=\"bar\",_foo=\"_bar\",foo1=\"bar1\",foo-foo=\"bar-bar\",foo.foo=\"bar.bar\"]",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "content",
    },
    &asciidoc.NewLine{},
  },
}

var doesNotParseNamedAttributeIfNameIsInvalid = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "[normal,foo.foo=\"bar.bar\",-foo-foo=\"-bar-bar\"]",
    },
    &asciidoc.NewLine{},
    &asciidoc.String{
      Value: "content",
    },
    &asciidoc.NewLine{},
  },
}

var positionalAttributesAssignedToBlock = &asciidoc.Document{
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
        &asciidoc.ShorthandAttribute{
          Style: &asciidoc.ShorthandStyle{
            Val: asciidoc.Set{
              &asciidoc.String{
                Value: "quote",
              },
            },
          },
          ID: nil,
          Roles: nil,
          Options: nil,
        },
        &asciidoc.PositionalAttribute{
          Offset: 1,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "author",
            },
          },
        },
        &asciidoc.PositionalAttribute{
          Offset: 2,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "source",
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

var normalSubstitutionsArePerformedOnSingleQuotedPositionalAttribute = &asciidoc.Document{
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
        &asciidoc.ShorthandAttribute{
          Style: &asciidoc.ShorthandStyle{
            Val: asciidoc.Set{
              &asciidoc.String{
                Value: "quote",
              },
            },
          },
          ID: nil,
          Roles: nil,
          Options: nil,
        },
        &asciidoc.PositionalAttribute{
          Offset: 1,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "author",
            },
          },
        },
        &asciidoc.PositionalAttribute{
          Offset: 2,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "http://wikipedia.org[source]",
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

var normalSubstitutionsArePerformedOnSingleQuotedNamedAttribute = &asciidoc.Document{
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
        &asciidoc.ShorthandAttribute{
          Style: &asciidoc.ShorthandStyle{
            Val: asciidoc.Set{
              &asciidoc.String{
                Value: "quote",
              },
            },
          },
          ID: nil,
          Roles: nil,
          Options: nil,
        },
        &asciidoc.PositionalAttribute{
          Offset: 1,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "author",
            },
          },
        },
        &asciidoc.NamedAttribute{
          Name: "citetitle",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "http://wikipedia.org[source]",
            },
          },
          Quote: 1,
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

var normalSubstitutionsArePerformedOnceOnSingleQuotedNamedTitleAttribute = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Paragraph{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.NamedAttribute{
          Name: "title",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "*title*",
            },
          },
          Quote: 1,
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

var attributeListMayNotBeginWithSpace = &asciidoc.Document{
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
        &asciidoc.ShorthandAttribute{
          Style: &asciidoc.ShorthandStyle{
            Val: asciidoc.Set{
              &asciidoc.String{
                Value: "quote",
              },
            },
          },
          ID: nil,
          Roles: nil,
          Options: nil,
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

var attributeListMayBeginWithComma = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "[, author, source]",
    },
    &asciidoc.NewLine{},
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

var firstAttributeInListMayBeDoubleQuoted = &asciidoc.Document{
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
        &asciidoc.ShorthandAttribute{
          Style: &asciidoc.ShorthandStyle{
            Val: asciidoc.Set{
              &asciidoc.String{
                Value: "quote",
              },
            },
          },
          ID: nil,
          Roles: nil,
          Options: nil,
        },
        &asciidoc.PositionalAttribute{
          Offset: 1,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "author",
            },
          },
        },
        &asciidoc.PositionalAttribute{
          Offset: 2,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "source",
            },
          },
        },
        &asciidoc.NamedAttribute{
          Name: "role",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "famous",
            },
          },
          Quote: 2,
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

var firstAttributeInListMayBeSingleQuoted = &asciidoc.Document{
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
        &asciidoc.ShorthandAttribute{
          Style: &asciidoc.ShorthandStyle{
            Val: asciidoc.Set{
              &asciidoc.String{
                Value: "quote",
              },
            },
          },
          ID: nil,
          Roles: nil,
          Options: nil,
        },
        &asciidoc.PositionalAttribute{
          Offset: 1,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "author",
            },
          },
        },
        &asciidoc.PositionalAttribute{
          Offset: 2,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "source",
            },
          },
        },
        &asciidoc.NamedAttribute{
          Name: "role",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "famous",
            },
          },
          Quote: 1,
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

var attributeWithValueNoneWithoutQuotesIsIgnored = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Paragraph{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.NamedAttribute{
          Name: "id",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "None",
            },
          },
          Quote: 0,
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "paragraph",
        },
        &asciidoc.NewLine{},
      },
      Admonition: 0,
    },
  },
}

var roleReturnsTrueIfRoleIsAssigned = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Paragraph{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.NamedAttribute{
          Name: "role",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "lead",
            },
          },
          Quote: 2,
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "A paragraph",
        },
        &asciidoc.NewLine{},
      },
      Admonition: 0,
    },
  },
}

var roleDoesNotReturnTrueIfRoleAttributeIsSetOnDocument = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "role",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "lead",
        },
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "A paragraph",
    },
    &asciidoc.NewLine{},
  },
}

var roleCanCheckForExactRoleNameMatch = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Paragraph{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.NamedAttribute{
          Name: "role",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "lead",
            },
          },
          Quote: 2,
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "A paragraph",
        },
        &asciidoc.NewLine{},
      },
      Admonition: 0,
    },
  },
}

var hasRoleCanCheckForPresenceOfRoleName = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Paragraph{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.NamedAttribute{
          Name: "role",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "lead abstract",
            },
          },
          Quote: 2,
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "A paragraph",
        },
        &asciidoc.NewLine{},
      },
      Admonition: 0,
    },
  },
}

var hasRoleDoesNotLookForRoleDefinedAsDocumentAttribute = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "role",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "lead abstract",
        },
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "A paragraph",
    },
    &asciidoc.NewLine{},
  },
}

var rolesReturnsArrayOfRoleNames = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Paragraph{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.NamedAttribute{
          Name: "role",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "story lead",
            },
          },
          Quote: 2,
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "A paragraph",
        },
        &asciidoc.NewLine{},
      },
      Admonition: 0,
    },
  },
}

var rolesReturnsEmptyArrayIfRoleAttributeIsNotSet = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "role",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "story lead",
        },
      },
    },
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.String{
      Value: "A paragraph",
    },
    &asciidoc.NewLine{},
  },
}

var rolesSetsTheRoleAttributeOnTheNode = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.AttributeEntry{
      Name: "lead",
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "role=\"lead\"",
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
            Val: asciidoc.Set{
              &asciidoc.String{
                Value: "{lead}",
              },
            },
          },
          ID: nil,
          Roles: nil,
          Options: nil,
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "A paragraph",
        },
        &asciidoc.NewLine{},
      },
      Admonition: 0,
    },
  },
}

var idRoleAndOptionsAttributesCanBeSpecifiedOnBlockStyleUsingShorthandSyntax = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Paragraph{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.ShorthandAttribute{
          Style: &asciidoc.ShorthandStyle{
            Val: asciidoc.Set{
              &asciidoc.String{
                Value: "literal",
              },
            },
          },
          ID: &asciidoc.ShorthandID{
            Val: asciidoc.Set{
              &asciidoc.String{
                Value: "first",
              },
            },
          },
          Roles: []*asciidoc.ShorthandRole{
            &asciidoc.ShorthandRole{
              Val: asciidoc.Set{
                &asciidoc.String{
                  Value: "lead",
                },
              },
            },
          },
          Options: []*asciidoc.ShorthandOption{
            &asciidoc.ShorthandOption{
              Val: asciidoc.Set{
                &asciidoc.String{
                  Value: "step",
                },
              },
            },
          },
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "A literal paragraph.",
        },
        &asciidoc.NewLine{},
      },
      Admonition: 0,
    },
  },
}

var idRoleAndOptionsAttributesCanBeSpecifiedUsingShorthandSyntaxOnBlockStyleUsingMultipleBlockAttributeLines = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Paragraph{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.ShorthandAttribute{
          Style: &asciidoc.ShorthandStyle{
            Val: asciidoc.Set{
              &asciidoc.String{
                Value: "literal",
              },
            },
          },
          ID: nil,
          Roles: nil,
          Options: nil,
        },
        &asciidoc.PositionalAttribute{
          Offset: 1,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "#first",
            },
          },
        },
        &asciidoc.PositionalAttribute{
          Offset: 2,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: ".lead",
            },
          },
        },
        &asciidoc.PositionalAttribute{
          Offset: 3,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "%step",
            },
          },
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "A literal paragraph.",
        },
        &asciidoc.NewLine{},
      },
      Admonition: 0,
    },
  },
}

var multipleRolesAndOptionsCanBeSpecifiedInBlockStyleUsingShorthandSyntax = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Paragraph{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.ShorthandAttribute{
          Style: nil,
          ID: nil,
          Roles: []*asciidoc.ShorthandRole{
            &asciidoc.ShorthandRole{
              Val: asciidoc.Set{
                &asciidoc.String{
                  Value: "role1",
                },
              },
            },
            &asciidoc.ShorthandRole{
              Val: asciidoc.Set{
                &asciidoc.String{
                  Value: "role2",
                },
              },
            },
          },
          Options: []*asciidoc.ShorthandOption{
            &asciidoc.ShorthandOption{
              Val: asciidoc.Set{
                &asciidoc.String{
                  Value: "option1",
                },
              },
            },
            &asciidoc.ShorthandOption{
              Val: asciidoc.Set{
                &asciidoc.String{
                  Value: "option2",
                },
              },
            },
          },
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "Text",
        },
        &asciidoc.NewLine{},
      },
      Admonition: 0,
    },
  },
}

var optionsSpecifiedUsingShorthandSyntaxOnBlockStyleAcrossMultipleLinesShouldBeAdditive = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Paragraph{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.ShorthandAttribute{
          Style: nil,
          ID: nil,
          Roles: nil,
          Options: []*asciidoc.ShorthandOption{
            &asciidoc.ShorthandOption{
              Val: asciidoc.Set{
                &asciidoc.String{
                  Value: "option1",
                },
              },
            },
          },
        },
        &asciidoc.PositionalAttribute{
          Offset: 1,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "%option2",
            },
          },
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "Text",
        },
        &asciidoc.NewLine{},
      },
      Admonition: 0,
    },
  },
}

var rolesSpecifiedUsingShorthandSyntaxOnBlockStyleAcrossMultipleLinesShouldBeAdditive = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Paragraph{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.ShorthandAttribute{
          Style: nil,
          ID: nil,
          Roles: []*asciidoc.ShorthandRole{
            &asciidoc.ShorthandRole{
              Val: asciidoc.Set{
                &asciidoc.String{
                  Value: "role1",
                },
              },
            },
          },
          Options: nil,
        },
        &asciidoc.PositionalAttribute{
          Offset: 1,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: ".role2.role3",
            },
          },
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "Text",
        },
        &asciidoc.NewLine{},
      },
      Admonition: 0,
    },
  },
}

var settingARoleUsingTheRoleAttributeReplacesAnyExistingRoles = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Paragraph{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.ShorthandAttribute{
          Style: nil,
          ID: nil,
          Roles: []*asciidoc.ShorthandRole{
            &asciidoc.ShorthandRole{
              Val: asciidoc.Set{
                &asciidoc.String{
                  Value: "role1",
                },
              },
            },
          },
          Options: nil,
        },
        &asciidoc.NamedAttribute{
          Name: "role",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "role2",
            },
          },
          Quote: 0,
        },
        &asciidoc.PositionalAttribute{
          Offset: 1,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: ".role3",
            },
          },
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "Text",
        },
        &asciidoc.NewLine{},
      },
      Admonition: 0,
    },
  },
}

var settingARoleUsingTheShorthandSyntaxOnBlockStyleShouldNotClearTheId = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Paragraph{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.ShorthandAttribute{
          Style: nil,
          ID: &asciidoc.ShorthandID{
            Val: asciidoc.Set{
              &asciidoc.String{
                Value: "id",
              },
            },
          },
          Roles: nil,
          Options: nil,
        },
        &asciidoc.PositionalAttribute{
          Offset: 1,
          ImpliedName: "",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: ".role",
            },
          },
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "Text",
        },
        &asciidoc.NewLine{},
      },
      Admonition: 0,
    },
  },
}

var aRoleCanBeAddedUsingAddRoleWhenTheNodeHasNoRoles = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Paragraph{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.ShorthandAttribute{
          Style: nil,
          ID: nil,
          Roles: []*asciidoc.ShorthandRole{
            &asciidoc.ShorthandRole{
              Val: asciidoc.Set{
                &asciidoc.String{
                  Value: "role1",
                },
              },
            },
          },
          Options: nil,
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "A normal paragraph",
        },
        &asciidoc.NewLine{},
      },
      Admonition: 0,
    },
  },
}

var aRoleIsNotAddedUsingAddRoleIfTheNodeAlreadyHasThatRole = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Paragraph{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.ShorthandAttribute{
          Style: nil,
          ID: nil,
          Roles: []*asciidoc.ShorthandRole{
            &asciidoc.ShorthandRole{
              Val: asciidoc.Set{
                &asciidoc.String{
                  Value: "role1",
                },
              },
            },
          },
          Options: nil,
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "A normal paragraph",
        },
        &asciidoc.NewLine{},
      },
      Admonition: 0,
    },
  },
}

var anExistingRoleCanBeRemovedUsingRemoveRole = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Paragraph{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.ShorthandAttribute{
          Style: nil,
          ID: nil,
          Roles: []*asciidoc.ShorthandRole{
            &asciidoc.ShorthandRole{
              Val: asciidoc.Set{
                &asciidoc.String{
                  Value: "role1",
                },
              },
            },
            &asciidoc.ShorthandRole{
              Val: asciidoc.Set{
                &asciidoc.String{
                  Value: "role2",
                },
              },
            },
          },
          Options: nil,
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "A normal paragraph",
        },
        &asciidoc.NewLine{},
      },
      Admonition: 0,
    },
  },
}

var rolesAreRemovedWhenLastRoleIsRemovedUsingRemoveRole = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Paragraph{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.ShorthandAttribute{
          Style: nil,
          ID: nil,
          Roles: []*asciidoc.ShorthandRole{
            &asciidoc.ShorthandRole{
              Val: asciidoc.Set{
                &asciidoc.String{
                  Value: "role1",
                },
              },
            },
          },
          Options: nil,
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "A normal paragraph",
        },
        &asciidoc.NewLine{},
      },
      Admonition: 0,
    },
  },
}

var rolesAreNotChangedWhenANonExistentRoleIsRemovedUsingRemoveRole = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Paragraph{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.ShorthandAttribute{
          Style: nil,
          ID: nil,
          Roles: []*asciidoc.ShorthandRole{
            &asciidoc.ShorthandRole{
              Val: asciidoc.Set{
                &asciidoc.String{
                  Value: "role1",
                },
              },
            },
          },
          Options: nil,
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "A normal paragraph",
        },
        &asciidoc.NewLine{},
      },
      Admonition: 0,
    },
  },
}

var rolesAreNotChangedWhenUsingRemoveRoleIfTheNodeHasNoRoles = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.UnorderedListItem{
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "checked",
        },
      },
      AttributeList: asciidoc.AttributeList{
        &asciidoc.ShorthandAttribute{
          Style: nil,
          ID: nil,
          Roles: nil,
          Options: []*asciidoc.ShorthandOption{
            &asciidoc.ShorthandOption{
              Val: asciidoc.Set{
                &asciidoc.String{
                  Value: "interactive",
                },
              },
            },
          },
        },
      },
      Indent: "",
      Marker: "-",
      Checklist: 2,
    },
  },
}

var idAndRoleAttributesCanBeSpecifiedOnSectionStyleUsingShorthandSyntax = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Section{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.ShorthandAttribute{
          Style: &asciidoc.ShorthandStyle{
            Val: asciidoc.Set{
              &asciidoc.String{
                Value: "dedication",
              },
            },
          },
          ID: &asciidoc.ShorthandID{
            Val: asciidoc.Set{
              &asciidoc.String{
                Value: "dedication",
              },
            },
          },
          Roles: []*asciidoc.ShorthandRole{
            &asciidoc.ShorthandRole{
              Val: asciidoc.Set{
                &asciidoc.String{
                  Value: "small",
                },
              },
            },
          },
          Options: nil,
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "Content.",
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

var idAttributeSpecifiedUsingShorthandSyntaxShouldNotCreateASpecialSection = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Section{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.ShorthandAttribute{
          Style: nil,
          ID: &asciidoc.ShorthandID{
            Val: asciidoc.Set{
              &asciidoc.String{
                Value: "idname",
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
          Value: "content",
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

var blockAttributesAreAdditive = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Paragraph{
      AttributeList: asciidoc.AttributeList{
        &asciidoc.NamedAttribute{
          Name: "id",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "foo",
            },
          },
          Quote: 1,
        },
        &asciidoc.NamedAttribute{
          Name: "role",
          Val: asciidoc.Set{
            &asciidoc.String{
              Value: "lead",
            },
          },
          Quote: 1,
        },
      },
      Set: asciidoc.Set{
        &asciidoc.String{
          Value: "A paragraph.",
        },
        &asciidoc.NewLine{},
      },
      Admonition: 0,
    },
  },
}

var lastWinsForIdAttribute = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Section{ // p0
      AttributeList: asciidoc.AttributeList{
        &asciidoc.AnchorAttribute{
          ID: &asciidoc.String{
            Value: "bar",
          },
          Label: nil,
        },
        &asciidoc.AnchorAttribute{
          ID: &asciidoc.String{
            Value: "foo",
          },
          Label: nil,
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
        asciidoc.EmptyLine{
          Text: "",
        },
        &asciidoc.Section{
          AttributeList: asciidoc.AttributeList{
            &asciidoc.AnchorAttribute{
              ID: &asciidoc.String{
                Value: "baz",
              },
              Label: nil,
            },
            &asciidoc.NamedAttribute{
              Name: "id",
              Val: asciidoc.Set{
                &asciidoc.String{
                  Value: "coolio",
                },
              },
              Quote: 1,
            },
          },
          Set: nil,
          Title: asciidoc.Set{
            &asciidoc.String{
              Value: "Section",
            },
          },
          Level: 2,
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

var trailingBlockAttributesTransferToTheFollowingSection = &asciidoc.Document{
  Set: asciidoc.Set{
    asciidoc.EmptyLine{
      Text: "",
    },
    &asciidoc.Section{ // p0
      AttributeList: asciidoc.AttributeList{
        &asciidoc.AnchorAttribute{
          ID: &asciidoc.String{
            Value: "one",
          },
          Label: nil,
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
        asciidoc.EmptyLine{
          Text: "",
        },
        &asciidoc.Paragraph{
          AttributeList: asciidoc.AttributeList{
            &asciidoc.AnchorAttribute{
              ID: &asciidoc.String{
                Value: "sub",
              },
              Label: nil,
            },
          },
          Set: asciidoc.Set{},
          Admonition: 0,
        },
        &asciidoc.SingleLineComment{
          Value: " try to mess this up!",
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
              Value: "paragraph",
            },
            &asciidoc.NewLine{},
            asciidoc.EmptyLine{
              Text: "",
            },
            &asciidoc.Paragraph{
              AttributeList: asciidoc.AttributeList{
                &asciidoc.NamedAttribute{
                  Name: "role",
                  Val: asciidoc.Set{
                    &asciidoc.String{
                      Value: "classy",
                    },
                  },
                  Quote: 1,
                },
              },
              Set: asciidoc.Set{},
              Admonition: 0,
            },
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
          },
          Title: asciidoc.Set{
            &asciidoc.String{
              Value: "Sub-section",
            },
          },
          Level: 2,
        },
      },
      Title: asciidoc.Set{
        &asciidoc.String{
          Value: "Section One",
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
          Value: "Section Two",
        },
      },
      Level: 1,
    },
  },
}


