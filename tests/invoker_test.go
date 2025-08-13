package tests

import (
	"testing"

	"github.com/project-chip/alchemy/asciidoc"
)

func TestInvoker(t *testing.T) {
	invokerTests.run(t)
}

var invokerTests = parseTests{

	{"should allow Options to be passed as first argument of constructor", "asciidoctor/invoker_test_should_allow_options_to_be_passed_as_first_argument_of_constructor.adoc", invokerTestShouldAllowOptionsToBePassedAsFirstArgumentOfConstructor},

	{"should change level on logger when --log-level is specified", "asciidoctor/invoker_test_should_change_level_on_logger_when___log_level_is_specified.adoc", invokerTestShouldChangeLevelOnLoggerWhenLogLevelIsSpecified},

	{"should not log when --log-level and -q are both specified", "asciidoctor/invoker_test_should_not_log_when___log_level_and__q_are_both_specified.adoc", invokerTestShouldNotLogWhenLogLevelAndQAreBothSpecified},

	{"should use specified log level when --log-level and -v are both specified", "asciidoctor/invoker_test_should_use_specified_log_level_when___log_level_and__v_are_both_specified.adoc", invokerTestShouldUseSpecifiedLogLevelWhenLogLevelAndVAreBothSpecified},

	{"should enable script warnings if -w flag is specified", "asciidoctor/invoker_test_should_enable_script_warnings_if__w_flag_is_specified.adoc", invokerTestShouldEnableScriptWarningsIfWFlagIsSpecified},

	{"should not fail to check log level when -q flag is specified", "asciidoctor/invoker_test_should_not_fail_to_check_log_level_when__q_flag_is_specified.adoc", invokerTestShouldNotFailToCheckLogLevelWhenQFlagIsSpecified},

	{"should return non-zero exit code if failure level is reached", "asciidoctor/invoker_test_should_return_non_zero_exit_code_if_failure_level_is_reached.adoc", invokerTestShouldReturnNonZeroExitCodeIfFailureLevelIsReached},

	{"should report usage if no input file given", "asciidoctor/invoker_test_should_report_usage_if_no_input_file_given.adoc", invokerTestShouldReportUsageIfNoInputFileGiven},
}

var invokerTestShouldAllowOptionsToBePassedAsFirstArgumentOfConstructor = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "second",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "2.",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "third",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "3.",
		},
	},
}

var invokerTestShouldChangeLevelOnLoggerWhenLogLevelIsSpecified = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "skip to ",
		},
		&asciidoc.CrossReference{
			AttributeList: nil,
			Elements:      nil,
			ID:            "install",
			Format:        0,
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "download",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "install",
				},
				&asciidoc.Anchor{
					ID:       "install",
					Elements: nil,
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "run",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
	},
}

var invokerTestShouldNotLogWhenLogLevelAndQAreBothSpecified = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "skip to ",
		},
		&asciidoc.CrossReference{
			AttributeList: nil,
			Elements:      nil,
			ID:            "install",
			Format:        0,
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "download",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "install",
				},
				&asciidoc.Anchor{
					ID:       "install",
					Elements: nil,
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "run",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
	},
}

var invokerTestShouldUseSpecifiedLogLevelWhenLogLevelAndVAreBothSpecified = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "skip to ",
		},
		&asciidoc.CrossReference{
			AttributeList: nil,
			Elements:      nil,
			ID:            "install",
			Format:        0,
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "download",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "install",
				},
				&asciidoc.Anchor{
					ID:       "install",
					Elements: nil,
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "run",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
	},
}

var invokerTestShouldEnableScriptWarningsIfWFlagIsSpecified = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "second",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "2.",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "third",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "3.",
		},
	},
}

var invokerTestShouldNotFailToCheckLogLevelWhenQFlagIsSpecified = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.String{
			Value: "skip to ",
		},
		&asciidoc.CrossReference{
			AttributeList: nil,
			Elements:      nil,
			ID:            "install",
			Format:        0,
		},
		&asciidoc.NewLine{},
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "download",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "install",
				},
				&asciidoc.Anchor{
					ID:       "install",
					Elements: nil,
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "run",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        ".",
		},
	},
}

var invokerTestShouldReturnNonZeroExitCodeIfFailureLevelIsReached = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "second",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "2.",
		},
		&asciidoc.OrderedListItem{
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "third",
				},
			},
			AttributeList: nil,
			Indent:        "",
			Marker:        "3.",
		},
	},
}

var invokerTestShouldReportUsageIfNoInputFileGiven = &asciidoc.Document{
	Elements: asciidoc.Elements{
		&asciidoc.EmptyLine{
			Text: "",
		},
		&asciidoc.Section{
			AttributeList: nil,
			Elements: asciidoc.Elements{
				&asciidoc.String{
					Value: "Andrew Stanton",
				},
				&asciidoc.NewLine{},
				&asciidoc.String{
					Value: "v1.0.0",
				},
				&asciidoc.NewLine{},
				&asciidoc.AttributeEntry{
					Name: "doctype",
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "manpage",
						},
					},
				},
				&asciidoc.AttributeEntry{
					Name: "manmanual",
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "EVE",
						},
					},
				},
				&asciidoc.AttributeEntry{
					Name: "mansource",
					Elements: asciidoc.Elements{
						&asciidoc.String{
							Value: "EVE",
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
						&asciidoc.String{
							Value: "eve, islifeform - analyzes an image to determine if it's a picture of a life form",
						},
						&asciidoc.NewLine{},
						&asciidoc.EmptyLine{
							Text: "",
						},
					},
					Title: asciidoc.Elements{
						&asciidoc.String{
							Value: "NAME",
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
						&asciidoc.Bold{
							AttributeList: nil,
							Elements: asciidoc.Elements{
								&asciidoc.String{
									Value: "eve",
								},
							},
						},
						&asciidoc.String{
							Value: " ['OPTION']... 'FILE'...",
						},
						&asciidoc.NewLine{},
					},
					Title: asciidoc.Elements{
						&asciidoc.String{
							Value: "SYNOPSIS",
						},
					},
					Level: 1,
				},
			},
			Title: asciidoc.Elements{
				&asciidoc.String{
					Value: "eve(1)",
				},
			},
			Level: 0,
		},
	},
}
