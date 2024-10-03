package conformance

import (
	"testing"
)

func TestOptional(t *testing.T) {
	conformance, err := tryParseConformance("[!AB & (CD != EF)], O")
	if err != nil {

		t.Errorf("failed parsing: %v", err)
	}
	t.Logf("conformance: %s", conformance.ASCIIDocString())
	t.Logf("description: %s", conformance.Description())
}

type conformanceTestSuite struct {
	Conformance        string
	InvalidConformance bool
	ASCIIDocString     string

	Tests []conformanceTest
}

func (cts *conformanceTestSuite) run(t *testing.T) {
	conformance, err := tryParseConformance(cts.Conformance)
	t.Logf("testing %s: %T", cts.Conformance, conformance)
	if err != nil {
		if cts.InvalidConformance {
			return
		}
		t.Errorf("failed parsing conformance %s: %v", cts.Conformance, err)
		return
	}
	t.Logf("\tconformance set: %d", len(conformance))
	for _, c := range conformance {
		t.Logf("\ttesting %s: %T %v", cts.Conformance, c, c)
	}
	for _, test := range cts.Tests {
		result, err := conformance.Eval(test.Context)
		if err != nil {
			t.Errorf("failed evaluating conformance status %v: %v", test.Context, err)
			return
		}
		if result != test.Expected {
			t.Errorf("failed checking conformance %s (parsed %s) with %v: expected %v, got %v", cts.Conformance, conformance.ASCIIDocString(), test.Context, test.Expected, result)
		}
	}
	if len(cts.ASCIIDocString) > 0 {
		ads := conformance.ASCIIDocString()
		if ads != cts.ASCIIDocString {
			t.Errorf("unexpected Asciidoc string for conformance \"%s\"; expected \"%s\", got \"%s\"", cts.Conformance, cts.ASCIIDocString, ads)
		}
	}
}

type conformanceTest struct {
	Context  Context
	Expected State
}

var otherwiseTests = []conformanceTestSuite{
	{
		Conformance: "NumberOfPrimaries > 0, O",
		Tests: []conformanceTest{
			{Context: Context{Values: map[string]any{"NumberOfPrimaries": int64(1)}}, Expected: StateMandatory},
			{Context: Context{Values: map[string]any{"NumberOfPrimaries": int64(0)}}, Expected: StateOptional},
		},
	},
	{
		Conformance: "TwoDCART",
		Tests: []conformanceTest{
			{Context: Context{Values: map[string]any{"TwoDCART": true}}, Expected: StateMandatory},
			{Context: Context{Values: map[string]any{"SFR": true}}, Expected: StateDisallowed},
			{Context: Context{Values: map[string]any{"Matter": true}}, Expected: StateDisallowed},
		},
	},
	{
		Conformance: "(STA|PAU|FA|CON)&!SFR,!PA&!SFR,O.a-",
		Tests: []conformanceTest{
			{Context: Context{Values: map[string]any{"STA": true}}, Expected: StateMandatory},
			{Context: Context{Values: map[string]any{"STA": true, "SFR": true}}, Expected: StateOptional},
			{Context: Context{Values: map[string]any{"PA": true}}, Expected: StateOptional},
			{Context: Context{Values: map[string]any{"CF": true}}, Expected: StateMandatory},
			{Context: Context{Values: map[string]any{"SFR": true}}, Expected: StateOptional},
			{Context: Context{Values: map[string]any{"Matter": true}}, Expected: StateMandatory},
		},
	},
	{

		Conformance: "[LT | DF & CF]",
		Tests: []conformanceTest{
			{Context: Context{Values: map[string]any{"AA": true}}, Expected: StateDisallowed},
			{Context: Context{Values: map[string]any{"LT": true}}, Expected: StateOptional},
			{Context: Context{Values: map[string]any{"DF": true}}, Expected: StateDisallowed},
			{Context: Context{Values: map[string]any{"CF": true}}, Expected: StateDisallowed},
			{Context: Context{Values: map[string]any{"DF": true, "CF": true}}, Expected: StateOptional},
			{Context: Context{Values: map[string]any{"Matter": true}}, Expected: StateDisallowed},
		},
	},
	{
		Conformance: "AB, [CD]",
		Tests: []conformanceTest{
			{Context: Context{Values: map[string]any{"AB": true}}, Expected: StateMandatory},
			{Context: Context{Values: map[string]any{"CD": true}}, Expected: StateOptional},
			{Context: Context{Values: map[string]any{"Matter": true}}, Expected: StateDisallowed},
		},
	},
	{
		Conformance: "!AB, O",
		Tests: []conformanceTest{
			{Context: Context{Values: map[string]any{"AB": true}}, Expected: StateOptional},
			{Context: Context{Values: map[string]any{"CD": true}}, Expected: StateMandatory},
			{Context: Context{Values: map[string]any{"Matter": true}}, Expected: StateMandatory},
		},
	},
	{
		Conformance: "[AA], [BB], [CC]",
		Tests: []conformanceTest{
			{Context: Context{Values: map[string]any{"AA": true}}, Expected: StateOptional},
			{Context: Context{Values: map[string]any{"BB": true}}, Expected: StateOptional},
			{Context: Context{Values: map[string]any{"CC": true}}, Expected: StateOptional},
			{Context: Context{Values: map[string]any{"Matter": true}}, Expected: StateDisallowed},
		},
	},
	{
		Conformance: "[!(LT | DF)]",
		Tests: []conformanceTest{
			{Context: Context{Values: map[string]any{"AA": true}}, Expected: StateOptional},
			{Context: Context{Values: map[string]any{"LT": true}}, Expected: StateDisallowed},
			{Context: Context{Values: map[string]any{"DF": true}}, Expected: StateDisallowed},
			{Context: Context{Values: map[string]any{"Matter": true}}, Expected: StateOptional},
		},
	},
	{
		Conformance: "[!(LT | DF | CF)]",
		Tests: []conformanceTest{
			{Context: Context{Values: map[string]any{"AA": true}}, Expected: StateOptional},
			{Context: Context{Values: map[string]any{"LT": true}}, Expected: StateDisallowed},
			{Context: Context{Values: map[string]any{"DF": true}}, Expected: StateDisallowed},
			{Context: Context{Values: map[string]any{"CF": true}}, Expected: StateDisallowed},
			{Context: Context{Values: map[string]any{"Matter": true}}, Expected: StateOptional},
		},
	},
	{
		Conformance: "[LT | DF]",
		Tests: []conformanceTest{
			{Context: Context{Values: map[string]any{"AA": true}}, Expected: StateDisallowed},
			{Context: Context{Values: map[string]any{"LT": true}}, Expected: StateOptional},
			{Context: Context{Values: map[string]any{"DF": true}}, Expected: StateOptional},
			{Context: Context{Values: map[string]any{"Matter": true}}, Expected: StateDisallowed},
		},
	},

	{
		Conformance: "UltrasonicUnoccupiedToOccupiedThreshold, O",
		Tests: []conformanceTest{
			{Context: Context{Values: map[string]any{"AA": true}}, Expected: StateOptional},
			{Context: Context{Values: map[string]any{"UltrasonicUnoccupiedToOccupiedThreshold": true}}, Expected: StateMandatory},
			{Context: Context{Values: map[string]any{"UltrasonicUnoccupiedToOccupiedThreshold": false}}, Expected: StateOptional},
			{Context: Context{Values: map[string]any{"Matter": true}}, Expected: StateOptional},
		},
	},
	{
		Conformance: "Zigbee",
		Tests: []conformanceTest{
			{Context: Context{Values: map[string]any{"AA": true}}, Expected: StateDisallowed},
			{Context: Context{Values: map[string]any{"Zigbee": true}}, Expected: StateMandatory},
			{Context: Context{Values: map[string]any{"Zigbee": false}}, Expected: StateDisallowed},
			{Context: Context{Values: map[string]any{"Matter": true}}, Expected: StateDisallowed},
		},
	},
	{
		Conformance: "[Zigbee]",
		Tests: []conformanceTest{
			{Context: Context{Values: map[string]any{"AA": true}}, Expected: StateDisallowed},
			{Context: Context{Values: map[string]any{"Zigbee": true}}, Expected: StateOptional},
			{Context: Context{Values: map[string]any{"Zigbee": false}}, Expected: StateDisallowed},
			{Context: Context{Values: map[string]any{"Matter": true}}, Expected: StateDisallowed},
		},
	},
	{
		Conformance: "MSCH",
		Tests: []conformanceTest{
			{Context: Context{Values: map[string]any{"AA": true}}, Expected: StateDisallowed},
			{Context: Context{Values: map[string]any{"MSCH": true}}, Expected: StateMandatory},
			{Context: Context{Values: map[string]any{"MSCH": false}}, Expected: StateDisallowed},
			{Context: Context{Values: map[string]any{"Matter": true}}, Expected: StateDisallowed},
		},
	},
	{
		Conformance: "M",
		Tests: []conformanceTest{
			{Context: Context{Values: map[string]any{"AA": true}}, Expected: StateMandatory},
			{Context: Context{Values: map[string]any{"MSCH": true}}, Expected: StateMandatory},
			{Context: Context{Values: map[string]any{"MSCH": false}}, Expected: StateMandatory},
			{Context: Context{Values: map[string]any{"Matter": true}}, Expected: StateMandatory},
		},
	},
	{
		Conformance: "(VIS | AUD) & SPRS",
		Tests: []conformanceTest{
			{Context: Context{Values: map[string]any{"VIS": true}}, Expected: StateDisallowed},
			{Context: Context{Values: map[string]any{"AUD": true}}, Expected: StateDisallowed},
			{Context: Context{Values: map[string]any{"VIS": true, "AUD": true}}, Expected: StateDisallowed},
			{Context: Context{Values: map[string]any{"VIS": true, "AUD": true, "SPRS": true}}, Expected: StateMandatory},
			{Context: Context{Values: map[string]any{"VIS": true, "SPRS": true}}, Expected: StateMandatory},
			{Context: Context{Values: map[string]any{"AUD": true, "SPRS": true}}, Expected: StateMandatory},
			{Context: Context{Values: map[string]any{"SPRS": true}}, Expected: StateDisallowed},
			{Context: Context{Values: map[string]any{"Matter": true}}, Expected: StateDisallowed},
		},
	},
	{
		Conformance: "UltrasonicUnoccupiedToOccupiedDelay, O",
		Tests: []conformanceTest{
			{Context: Context{Values: map[string]any{"AA": true}}, Expected: StateOptional},
			{Context: Context{Values: map[string]any{"UltrasonicUnoccupiedToOccupiedDelay": true}}, Expected: StateMandatory},
			{Context: Context{Values: map[string]any{"UltrasonicUnoccupiedToOccupiedDelay": false}}, Expected: StateOptional},
			{Context: Context{Values: map[string]any{"Matter": true}}, Expected: StateOptional},
		},
	},
	{
		Conformance: "PIRUnoccupiedToOccupiedThreshold, O",
		Tests: []conformanceTest{
			{Context: Context{Values: map[string]any{"AA": true}}, Expected: StateOptional},
			{Context: Context{Values: map[string]any{"PIRUnoccupiedToOccupiedThreshold": true}}, Expected: StateMandatory},
			{Context: Context{Values: map[string]any{"UltrasonicUnoccupiedToOccupiedDelay": false}}, Expected: StateOptional},
			{Context: Context{Values: map[string]any{"Matter": true}}, Expected: StateOptional},
		},
	},
	{
		Conformance:        "[AA] & BB",
		InvalidConformance: true,
	},
	{
		Conformance:        "AA | [BB]",
		InvalidConformance: true,
	},
	{
		Conformance: "AA, [BB]",
		Tests: []conformanceTest{
			{Context: Context{Values: map[string]any{"AA": true}}, Expected: StateMandatory},
			{Context: Context{Values: map[string]any{"BB": true}}, Expected: StateOptional},
			{Context: Context{Values: map[string]any{"CC": false}}, Expected: StateDisallowed},
			{Context: Context{Values: map[string]any{"Matter": true}}, Expected: StateDisallowed},
		},
	},
	{
		Conformance: "<<ref_Ranges>>",
		Tests: []conformanceTest{
			{Context: Context{Values: map[string]any{"AA": true}}, Expected: StateDisallowed},
			{Context: Context{Values: map[string]any{"BB": true}}, Expected: StateDisallowed},
			{Context: Context{Values: map[string]any{"CC": false}}, Expected: StateDisallowed},
			{Context: Context{Values: map[string]any{"Matter": true}}, Expected: StateDisallowed},
		},
	},
	{
		Conformance: "[Wi-Fi]",
		Tests: []conformanceTest{
			{Context: Context{Values: map[string]any{"AA": true}}, Expected: StateDisallowed},
			{Context: Context{Values: map[string]any{"Wi-Fi": true}}, Expected: StateOptional},
			{Context: Context{Values: map[string]any{"UltrasonicUnoccupiedToOccupiedDelay": false}}, Expected: StateDisallowed},
			{Context: Context{Values: map[string]any{"Matter": true}}, Expected: StateDisallowed},
		},
	},
	{
		Conformance: "[Wi-Fi]",
		Tests: []conformanceTest{
			{Context: Context{Values: map[string]any{"AA": true}}, Expected: StateDisallowed},
			{Context: Context{Values: map[string]any{"Wi-Fi": true}}, Expected: StateOptional},
			{Context: Context{Values: map[string]any{"UltrasonicUnoccupiedToOccupiedDelay": false}}, Expected: StateDisallowed},
			{Context: Context{Values: map[string]any{"Matter": true}}, Expected: StateDisallowed},
		},
	},
	{
		Conformance:    "<<ref_Ranges>>",
		ASCIIDocString: "<<ref_Ranges>>",
		Tests: []conformanceTest{
			{Context: Context{Values: map[string]any{"AA": true}}, Expected: StateDisallowed},
			{Context: Context{Values: map[string]any{"Wi-Fi": true}}, Expected: StateDisallowed},
			{Context: Context{Values: map[string]any{"UltrasonicUnoccupiedToOccupiedDelay": false}}, Expected: StateDisallowed},
			{Context: Context{Values: map[string]any{"Matter": true}}, Expected: StateDisallowed},
		},
	},
	{
		Conformance:    "<<ref_Ranges, Ranges>>",
		ASCIIDocString: "<<ref_Ranges, Ranges>>",
		Tests: []conformanceTest{
			{Context: Context{Values: map[string]any{"AA": true}}, Expected: StateDisallowed},
			{Context: Context{Values: map[string]any{"Wi-Fi": true}}, Expected: StateDisallowed},
			{Context: Context{Values: map[string]any{"UltrasonicUnoccupiedToOccupiedDelay": false}}, Expected: StateDisallowed},
			{Context: Context{Values: map[string]any{"Matter": true}}, Expected: StateDisallowed},
		},
	},
	{
		Conformance:    "P, WATTS",
		ASCIIDocString: "P, WATTS",
		Tests: []conformanceTest{
			{Context: Context{Values: map[string]any{"AA": true}}, Expected: StateProvisional},
			{Context: Context{Values: map[string]any{"Wi-Fi": true}}, Expected: StateProvisional},
			{Context: Context{Values: map[string]any{"WATTS": false}}, Expected: StateProvisional},
			{Context: Context{Values: map[string]any{"Matter": true}}, Expected: StateProvisional},
		},
	},
	{
		Conformance:    "M, D",
		ASCIIDocString: "M, D", // Preserve the deprecated after mandatory
		Tests: []conformanceTest{
			{Context: Context{Values: map[string]any{"AA": true}}, Expected: StateMandatory},
			{Context: Context{Values: map[string]any{"Wi-Fi": true}}, Expected: StateMandatory},
			{Context: Context{Values: map[string]any{"WATTS": false}}, Expected: StateMandatory},
			{Context: Context{Values: map[string]any{"Matter": true}}, Expected: StateMandatory},
		},
	},
	{
		Conformance:    "O, D",
		ASCIIDocString: "O, D", // Preserve the deprecated after optional
		Tests: []conformanceTest{
			{Context: Context{Values: map[string]any{"AA": true}}, Expected: StateOptional},
			{Context: Context{Values: map[string]any{"Wi-Fi": true}}, Expected: StateOptional},
			{Context: Context{Values: map[string]any{"WATTS": false}}, Expected: StateOptional},
			{Context: Context{Values: map[string]any{"Matter": true}}, Expected: StateOptional},
		},
	},
	{
		Conformance:    "M, [AA], D",
		ASCIIDocString: "M, D", // Remove the nonsensical optional conformance after mandatory, but preserve the deprecated
		Tests: []conformanceTest{
			{Context: Context{Values: map[string]any{"AA": true}}, Expected: StateMandatory},
			{Context: Context{Values: map[string]any{"Wi-Fi": true}}, Expected: StateMandatory},
			{Context: Context{Values: map[string]any{"WATTS": false}}, Expected: StateMandatory},
			{Context: Context{Values: map[string]any{"Matter": true}}, Expected: StateMandatory},
		},
	},
}

func TestOtherwise(t *testing.T) {
	for _, test := range otherwiseTests {
		test.run(t)
	}
}
