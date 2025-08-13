package tests

import (
	"testing"

	"github.com/project-chip/alchemy/asciidoc"
)

func TestConverter(t *testing.T) {
	converterTests.run(t)
}

var converterTests = parseTests{

	{"should set Haml format to html5 for html5 backend", "asciidoctor/converter_test_should_set_haml_format_to_html_5_for_html_5_backend.adoc", converterTestShouldSetHamlFormatToHtml5ForHtml5Backend},

	{"should use built-in global cache to cache templates", "asciidoctor/converter_test_should_use_built_in_global_cache_to_cache_templates.adoc", converterTestShouldUseBuiltInGlobalCacheToCacheTemplates},

	{"should load ERB templates using ErubiTemplate if eruby is set to erubi", "asciidoctor/converter_test_should_load_erb_templates_using_erubi_template_if_eruby_is_set_to_erubi.adoc", converterTestShouldLoadErbTemplatesUsingErubiTemplateIfErubyIsSetToErubi},

	{"should be able to override the outline using a custom template", "asciidoctor/converter_test_should_be_able_to_override_the_outline_using_a_custom_template.adoc", converterTestShouldBeAbleToOverrideTheOutlineUsingACustomTemplate},

	{"should not expose included method on Converter class", "asciidoctor/converter_test_should_not_expose_included_method_on_converter_class.adoc", converterTestShouldNotExposeIncludedMethodOnConverterClass},

	{"should use specified converter for specified backend", "asciidoctor/converter_test_should_use_specified_converter_for_specified_backend.adoc", converterTestShouldUseSpecifiedConverterForSpecifiedBackend},

	{"should get converter from specified converter factory", "asciidoctor/converter_test_should_get_converter_from_specified_converter_factory.adoc", converterTestShouldGetConverterFromSpecifiedConverterFactory},
}

var converterTestShouldSetHamlFormatToHtml5ForHtml5Backend = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
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
				&asciidoc.Section{
					AttributeList: nil,
					Elements: asciidoc.Elements{
						&asciidoc.EmptyLine{
							Text: "",
						},
						&asciidoc.String{
							Value: "Sample paragraph",
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
							AttributeList: asciidoc.AttributeList{
								&asciidoc.TitleAttribute{
									Val: asciidoc.Elements{
										&asciidoc.String{
											Value: "Related",
										},
									},
								},
							},
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "Sidebar content",
								},
								&asciidoc.NewLine{},
							},
						},
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

var converterTestShouldUseBuiltInGlobalCacheToCacheTemplates = &asciidoc.Document{
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
					Value: "      <div class=\"openblock wrapper\">",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "      <div class=\"content\">",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "      <div class=\"paragraph\">",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "      <p>foobar</p>",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "      </div>",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "      </div>",
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

var converterTestShouldLoadErbTemplatesUsingErubiTemplateIfErubyIsSetToErubi = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
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
				&asciidoc.Section{
					AttributeList: nil,
					Elements: asciidoc.Elements{
						&asciidoc.EmptyLine{
							Text: "",
						},
						&asciidoc.String{
							Value: "Sample paragraph",
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
							AttributeList: asciidoc.AttributeList{
								&asciidoc.TitleAttribute{
									Val: asciidoc.Elements{
										&asciidoc.String{
											Value: "Related",
										},
									},
								},
							},
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "Sidebar content",
								},
								&asciidoc.NewLine{},
							},
						},
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

var converterTestShouldBeAbleToOverrideTheOutlineUsingACustomTemplate = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.AttributeEntry{
			Name:     "toc",
			Elements: nil,
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
					},
					Title: asciidoc.Elements{
						&asciidoc.String{
							Value: "Section One",
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
					},
					Title: asciidoc.Elements{
						&asciidoc.String{
							Value: "Section Two",
						},
					},
					Level: 1,
				},
				&asciidoc.Section{
					AttributeList: nil,
					Elements:      nil,
					Title: asciidoc.Elements{
						&asciidoc.String{
							Value: "Section Three",
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

var converterTestShouldNotExposeIncludedMethodOnConverterClass = &asciidoc.Document{
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

var converterTestShouldUseSpecifiedConverterForSpecifiedBackend = &asciidoc.Document{
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

var converterTestShouldGetConverterFromSpecifiedConverterFactory = &asciidoc.Document{
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
