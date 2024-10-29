package matter

import "testing"

type caseTest struct {
	input  string
	output string
}

type caseSeperatorTest struct {
	caseTest
	separator rune
}

var caseTests = []caseTest{
	{"TCP Wait", "TCPWait"},
	{"VendorID", "VendorID"},
	{"VendorID", "VendorID"},
	{"Vendor Attribute", "VendorAttribute"},
	{"HTTP Request", "HTTPRequest"},
	{"TV Set", "TVSet"},
	{"Set TV", "SetTV"},
	{"Set Tv", "SetTv"},
}

var caseSeparatorTests = []caseSeperatorTest{
	{caseTest: caseTest{"TCP Wait", "TCP_Wait"}, separator: '_'},
	{caseTest: caseTest{"VendorID", "VendorID"}, separator: '_'},
	{caseTest: caseTest{"VendorID", "VendorID"}, separator: '_'},
	{caseTest: caseTest{"Vendor Attribute", "Vendor_Attribute"}, separator: '_'},
	{caseTest: caseTest{"HTTP Request", "HTTP_Request"}, separator: '_'},
	{caseTest: caseTest{"TV Set", "TV_Set"}, separator: '_'},
	{caseTest: caseTest{"Set TV", "Set_TV"}, separator: '_'},
	{caseTest: caseTest{"Set Tv", "Set_Tv"}, separator: '_'},
	{caseTest: caseTest{"LoggedOut", "LoggedOut"}, separator: '_'},
}

func TestSuite(t *testing.T) {
	AddCaseAcronym("ID")
	for _, ct := range caseTests {
		out := Case(ct.input)
		if out != ct.output {
			t.Errorf("casing failed; input %s; expected %s, got %s", ct.input, ct.output, out)
		}
	}
	for _, ct := range caseSeparatorTests {
		out := CaseWithSeparator(ct.input, ct.separator)
		if out != ct.output {
			t.Errorf("casing with separator failed; input %s; expected %s, got %s", ct.input, ct.output, out)
		}
	}
}
