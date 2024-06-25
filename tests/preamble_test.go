package tests

import (
	"testing"

	"github.com/project-chip/alchemy/asciidoc"
)

func TestPreamble(t *testing.T) {
	preambleTests.run(t)
}

var preambleTests = parseTests{

	{"title and single paragraph preamble before section", "asciidoctor/preamble_test_title_and_single_paragraph_preamble_before_section.adoc", titleAndSingleParagraphPreambleBeforeSection},

	{"title of preface is blank by default in DocBook output", "asciidoctor/preamble_test_title_of_preface_is_blank_by_default_in_doc_book_output.adoc", titleOfPrefaceIsBlankByDefaultInDocBookOutput},

	{"preface-title attribute is assigned as title of preface in DocBook output", "asciidoctor/preamble_test_preface_title_attribute_is_assigned_as_title_of_preface_in_doc_book_output.adoc", prefaceTitleAttributeIsAssignedAsTitleOfPrefaceInDocBookOutput},

	{"title and multi-paragraph preamble before section", "asciidoctor/preamble_test_title_and_multi_paragraph_preamble_before_section.adoc", titleAndMultiParagraphPreambleBeforeSection},

	{"should not wrap content in preamble if document has title but no sections", "asciidoctor/preamble_test_should_not_wrap_content_in_preamble_if_document_has_title_but_no_sections.adoc", shouldNotWrapContentInPreambleIfDocumentHasTitleButNoSections},

	{"title and section without preamble", "asciidoctor/preamble_test_title_and_section_without_preamble.adoc", titleAndSectionWithoutPreamble},

	{"no title with preamble and section", "asciidoctor/preamble_test_no_title_with_preamble_and_section.adoc", noTitleWithPreambleAndSection},

	{"preamble in book doctype", "asciidoctor/preamble_test_preamble_in_book_doctype.adoc", preambleInBookDoctype},

	{"should output table of contents in preamble if toc-placement attribute value is preamble", "asciidoctor/preamble_test_should_output_table_of_contents_in_preamble_if_toc_placement_attribute_value_is_preamble.adoc", shouldOutputTableOfContentsInPreambleIfTocPlacementAttributeValueIsPreamble},

	{"should move abstract in implicit preface to info tag when converting to DocBook", "asciidoctor/preamble_test_should_move_abstract_in_implicit_preface_to_info_tag_when_converting_to_doc_book.adoc", shouldMoveAbstractInImplicitPrefaceToInfoTagWhenConvertingToDocBook},

	{"should move abstract as first section to info tag when converting to DocBook", "asciidoctor/preamble_test_should_move_abstract_as_first_section_to_info_tag_when_converting_to_doc_book.adoc", shouldMoveAbstractAsFirstSectionToInfoTagWhenConvertingToDocBook},

	{"should move abstract in preface section to info tag when converting to DocBook", "asciidoctor/preamble_test_should_move_abstract_in_preface_section_to_info_tag_when_converting_to_doc_book.adoc", shouldMoveAbstractInPrefaceSectionToInfoTagWhenConvertingToDocBook},
}

var titleAndSingleParagraphPreambleBeforeSection = &asciidoc.Document{
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
					Value: "Preamble paragraph 1.",
				},
				&asciidoc.NewLine{},
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
							Value: "Section paragraph 1.",
						},
						&asciidoc.NewLine{},
					},
					Title: asciidoc.Set{
						&asciidoc.String{
							Value: "First Section",
						},
					},
					Level: 1,
				},
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

var titleOfPrefaceIsBlankByDefaultInDocBookOutput = &asciidoc.Document{
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
				&asciidoc.String{
					Value: "Preface content.",
				},
				&asciidoc.NewLine{},
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
							Value: "Section content.",
						},
						&asciidoc.NewLine{},
					},
					Title: asciidoc.Set{
						&asciidoc.String{
							Value: "First Section",
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

var prefaceTitleAttributeIsAssignedAsTitleOfPrefaceInDocBookOutput = &asciidoc.Document{
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
				&asciidoc.AttributeEntry{
					Name: "preface-title",
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "Preface",
						},
					},
				},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "Preface content.",
				},
				&asciidoc.NewLine{},
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
							Value: "Section content.",
						},
						&asciidoc.NewLine{},
					},
					Title: asciidoc.Set{
						&asciidoc.String{
							Value: "First Section",
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

var titleAndMultiParagraphPreambleBeforeSection = &asciidoc.Document{
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
					Value: "Preamble paragraph 1.",
				},
				&asciidoc.NewLine{},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "Preamble paragraph 2.",
				},
				&asciidoc.NewLine{},
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
							Value: "Section paragraph 1.",
						},
						&asciidoc.NewLine{},
					},
					Title: asciidoc.Set{
						&asciidoc.String{
							Value: "First Section",
						},
					},
					Level: 1,
				},
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

var shouldNotWrapContentInPreambleIfDocumentHasTitleButNoSections = &asciidoc.Document{
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
					Value: "paragraph",
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

var titleAndSectionWithoutPreamble = &asciidoc.Document{
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
						&asciidoc.String{
							Value: "Section paragraph 1.",
						},
						&asciidoc.NewLine{},
					},
					Title: asciidoc.Set{
						&asciidoc.String{
							Value: "First Section",
						},
					},
					Level: 1,
				},
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

var noTitleWithPreambleAndSection = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "Preamble paragraph 1.",
		},
		&asciidoc.NewLine{},
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
					Value: "Section paragraph 1.",
				},
				&asciidoc.NewLine{},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "First Section",
				},
			},
			Level: 1,
		},
	},
}

var preambleInBookDoctype = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
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
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "Back then...",
				},
				&asciidoc.NewLine{},
				asciidoc.EmptyLine{
					Text: "",
				},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Book",
				},
			},
			Level: 0,
		},
		&asciidoc.Section{ // p0
			AttributeList: nil,
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
										Value: "partintro",
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
							Value: "It was a dark and stormy night...",
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
					Set: asciidoc.Set{
						asciidoc.EmptyLine{
							Text: "",
						},
						&asciidoc.String{
							Value: "Someone's gonna get axed.",
						},
						&asciidoc.NewLine{},
						asciidoc.EmptyLine{
							Text: "",
						},
					},
					Title: asciidoc.Set{
						&asciidoc.String{
							Value: "Scene One",
						},
					},
					Level: 1,
				},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Chapter One",
				},
			},
			Level: 0,
		},
		&asciidoc.Section{ // p1
			AttributeList: nil,
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
										Value: "partintro",
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
							Value: "They couldn't believe their eyes when...",
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
					Set: asciidoc.Set{
						asciidoc.EmptyLine{
							Text: "",
						},
						&asciidoc.String{
							Value: "The axe came swinging.",
						},
						&asciidoc.NewLine{},
					},
					Title: asciidoc.Set{
						&asciidoc.String{
							Value: "Scene One",
						},
					},
					Level: 1,
				},
			},
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Chapter Two",
				},
			},
			Level: 0,
		},
	},
}

var shouldOutputTableOfContentsInPreambleIfTocPlacementAttributeValueIsPreamble = &asciidoc.Document{
	Set: asciidoc.Set{
		asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{ // p0
			AttributeList: nil,
			Set: asciidoc.Set{
				&asciidoc.AttributeEntry{
					Name: "toc",
					Set:  nil,
				},
				&asciidoc.AttributeEntry{
					Name: "toc-placement",
					Set: asciidoc.Set{
						&asciidoc.String{
							Value: "preamble",
						},
					},
				},
				asciidoc.EmptyLine{
					Text: "",
				},
				&asciidoc.String{
					Value: "Once upon a time...",
				},
				&asciidoc.NewLine{},
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
							Value: "It was a dark and stormy night...",
						},
						&asciidoc.NewLine{},
						asciidoc.EmptyLine{
							Text: "",
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
							Value: "They couldn't believe their eyes when...",
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
			Title: asciidoc.Set{
				&asciidoc.String{
					Value: "Article",
				},
			},
			Level: 0,
		},
	},
}

var shouldMoveAbstractInImplicitPrefaceToInfoTagWhenConvertingToDocBook = &asciidoc.Document{
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
				&asciidoc.Paragraph{
					AttributeList: asciidoc.AttributeList{
						&asciidoc.ShorthandAttribute{
							Style: &asciidoc.ShorthandStyle{
								Set: asciidoc.Set{
									&asciidoc.String{
										Value: "abstract",
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
							Value: "This is the abstract.",
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
					Set:           nil,
					Title: asciidoc.Set{
						&asciidoc.String{
							Value: "Fin",
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

var shouldMoveAbstractAsFirstSectionToInfoTagWhenConvertingToDocBook = &asciidoc.Document{
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
					AttributeList: asciidoc.AttributeList{
						&asciidoc.ShorthandAttribute{
							Style: &asciidoc.ShorthandStyle{
								Set: asciidoc.Set{
									&asciidoc.String{
										Value: "abstract",
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
						&asciidoc.String{
							Value: "This is the abstract.",
						},
						&asciidoc.NewLine{},
						asciidoc.EmptyLine{
							Text: "",
						},
					},
					Title: asciidoc.Set{
						&asciidoc.String{
							Value: "Abstract",
						},
					},
					Level: 1,
				},
				&asciidoc.Section{
					AttributeList: nil,
					Set:           nil,
					Title: asciidoc.Set{
						&asciidoc.String{
							Value: "Fin",
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

var shouldMoveAbstractInPrefaceSectionToInfoTagWhenConvertingToDocBook = &asciidoc.Document{
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
							Style: &asciidoc.ShorthandStyle{
								Set: asciidoc.Set{
									&asciidoc.String{
										Value: "preface",
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
						&asciidoc.Paragraph{
							AttributeList: asciidoc.AttributeList{
								&asciidoc.ShorthandAttribute{
									Style: &asciidoc.ShorthandStyle{
										Set: asciidoc.Set{
											&asciidoc.String{
												Value: "abstract",
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
									Value: "This is the abstract.",
								},
								&asciidoc.NewLine{},
							},
							Admonition: 0,
						},
						asciidoc.EmptyLine{
							Text: "",
						},
					},
					Title: asciidoc.Set{
						&asciidoc.String{
							Value: "Preface",
						},
					},
					Level: 1,
				},
				&asciidoc.Section{
					AttributeList: nil,
					Set:           nil,
					Title: asciidoc.Set{
						&asciidoc.String{
							Value: "Fin",
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
