package matter

import "testing"

type caseTest struct {
	input  string
	output string
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

func TestSuite(t *testing.T) {
	AddCaseAcronym("ID")
	for _, ct := range caseTests {
		out := Case(ct.input)
		if out != ct.output {
			t.Errorf("casing failed; input %s; expected %s, got %s", ct.input, ct.output, out)
		}
	}

}
