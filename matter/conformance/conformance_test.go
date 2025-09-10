package conformance

import (
	"testing"
)

func TestOptional(t *testing.T) {
	_, err := TryParseConformance("[!AB & (CD != EF)], O")
	if err != nil {
		t.Errorf("failed parsing: %v", err)
	}
}

type conformanceTestSuite struct {
	Conformance        string
	InvalidConformance bool
	ASCIIDocString     string

	Tests []conformanceTest
}

func (cts *conformanceTestSuite) run(t *testing.T) {
	conformance, err := TryParseConformance(cts.Conformance)
	//t.Logf("testing %s: %T", cts.Conformance, conformance)
	if err != nil {
		if cts.InvalidConformance {
			return
		}
		t.Errorf("failed parsing conformance %s: %v", cts.Conformance, err)
		return
	}
	//t.Logf("\tconformance set: %d", len(conformance))
	/*for _, c := range conformance {
		t.Logf("\ttesting %s: %T %v", cts.Conformance, c, c)
	}*/
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
	Expected ConformanceState
}

var conformanceTests = []conformanceTestSuite{
	{
		Conformance: "RequiresEncodedPixels==true",
		Tests: []conformanceTest{
			{Context: Context{Values: map[string]any{"RequiresEncodedPixels": false}}, Expected: ConformanceState{State: StateDisallowed, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"RequiresEncodedPixels": true}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceDefinite}},
		},
	},
	{
		Conformance: `OperationalStateID >= 0x80 & OperationalStateID \<= 0xBF`,
		Tests: []conformanceTest{
			{Context: Context{Values: map[string]any{"OperationalStateID": uint64(0x7A)}}, Expected: ConformanceState{State: StateDisallowed, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"OperationalStateID": uint64(0x80)}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"OperationalStateID": uint64(0xBF)}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"OperationalStateID": uint64(0xC0)}}, Expected: ConformanceState{State: StateDisallowed, Confidence: ConfidenceDefinite}},
		},
	},
	{
		Conformance: "OperationalStateID >= 0x80 & OperationalStateID <= 0xBF",
		Tests: []conformanceTest{
			{Context: Context{Values: map[string]any{"OperationalStateID": uint64(0x7A)}}, Expected: ConformanceState{State: StateDisallowed, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"OperationalStateID": uint64(0x80)}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"OperationalStateID": uint64(0xBF)}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"OperationalStateID": uint64(0xC0)}}, Expected: ConformanceState{State: StateDisallowed, Confidence: ConfidenceDefinite}},
		},
	},
	{
		Conformance: "[AB & CM].a2",
		Tests: []conformanceTest{
			{Context: Context{Values: map[string]any{"AB": true}}, Expected: ConformanceState{State: StateOptional, Confidence: ConfidencePossible}},
			{Context: Context{Values: map[string]any{"AB": true, "CM": true}}, Expected: ConformanceState{State: StateOptional, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"AB": true, "CM": false}}, Expected: ConformanceState{State: StateDisallowed, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"Matter": true}}, Expected: ConformanceState{State: StateOptional, Confidence: ConfidencePossible}},
		},
	},
	{
		Conformance: "AB, O.a1+",
		Tests: []conformanceTest{
			{Context: Context{Values: map[string]any{"AB": true}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"Matter": true}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidencePossible}},
		},
	},
	{
		Conformance: "O.a1+",
		Tests: []conformanceTest{
			{Context: Context{Values: map[string]any{"Matter": true}}, Expected: ConformanceState{State: StateOptional, Confidence: ConfidenceDefinite}},
		},
	},
	{
		Conformance: "NumberOfPrimaries > 0, O",
		Tests: []conformanceTest{
			{Context: Context{Values: map[string]any{"NumberOfPrimaries": int64(1)}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"NumberOfPrimaries": int64(0)}}, Expected: ConformanceState{State: StateOptional, Confidence: ConfidenceDefinite}},
		},
	},
	{
		Conformance: "TwoDCART",
		Tests: []conformanceTest{
			{Context: Context{Values: map[string]any{"TwoDCART": true}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"TwoDCART": false}}, Expected: ConformanceState{State: StateDisallowed, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"SFR": true}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidencePossible}},
			{Context: Context{Values: map[string]any{"Matter": true}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidencePossible}},
		},
	},
	{
		Conformance: "(STA|PAU|FA|CON)&!SFR,!PA&!SFR,O.a-",
		Tests: []conformanceTest{
			{Context: Context{Values: map[string]any{"STA": true}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidencePossible}},
			{Context: Context{Values: map[string]any{"STA": true, "SFR": true}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidencePossible}},
			{Context: Context{Values: map[string]any{"STA": true, "PAU": true, "FA": true, "CON": true, "SFR": true, "PA": true}}, Expected: ConformanceState{State: StateOptional, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"STA": true, "PAU": true, "FA": true, "CON": true, "SFR": false, "PA": true}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"STA": false, "PAU": false, "FA": false, "CON": false, "SFR": false, "PA": false}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"PA": true}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidencePossible}},
			{Context: Context{Values: map[string]any{"CF": true}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidencePossible}},
			{Context: Context{Values: map[string]any{"SFR": true}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidencePossible}},
			{Context: Context{Values: map[string]any{"Matter": true}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidencePossible}},
		},
	},
	{

		Conformance: "[LT | DF & CF]",
		Tests: []conformanceTest{
			{Context: Context{Values: map[string]any{"AA": true}}, Expected: ConformanceState{State: StateOptional, Confidence: ConfidencePossible}},
			{Context: Context{Values: map[string]any{"LT": true}}, Expected: ConformanceState{State: StateOptional, Confidence: ConfidencePossible}},
			{Context: Context{Values: map[string]any{"LT": true, "DF": false, "CF": false}}, Expected: ConformanceState{State: StateOptional, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"LT": true, "DF": true, "CF": false}}, Expected: ConformanceState{State: StateOptional, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"LT": false, "DF": true, "CF": true}}, Expected: ConformanceState{State: StateOptional, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"LT": false, "DF": false, "CF": true}}, Expected: ConformanceState{State: StateDisallowed, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"LT": false, "DF": true, "CF": false}}, Expected: ConformanceState{State: StateDisallowed, Confidence: ConfidenceDefinite}},
		},
	},
	{
		Conformance: "AB, [CD]",
		Tests: []conformanceTest{
			{Context: Context{Values: map[string]any{"AB": true}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"CD": true}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidencePossible}},
			{Context: Context{Values: map[string]any{"AB": false, "CD": true}}, Expected: ConformanceState{State: StateOptional, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"AB": false, "CD": false}}, Expected: ConformanceState{State: StateDisallowed, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"Matter": true}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidencePossible}},
		},
	},
	{
		Conformance: "!AB, O",
		Tests: []conformanceTest{
			{Context: Context{Values: map[string]any{"AB": true}}, Expected: ConformanceState{State: StateOptional, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"AB": false}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"CD": true}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidencePossible}},
			{Context: Context{Values: map[string]any{"Matter": true}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidencePossible}},
		},
	},
	{
		Conformance: "[AA], [BB], [CC]",
		Tests: []conformanceTest{
			{Context: Context{Values: map[string]any{"AA": true}}, Expected: ConformanceState{State: StateOptional, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"BB": true}}, Expected: ConformanceState{State: StateOptional, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"CC": true}}, Expected: ConformanceState{State: StateOptional, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"AA": false, "BB": false, "CC": false}}, Expected: ConformanceState{State: StateDisallowed, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"Matter": true}}, Expected: ConformanceState{State: StateOptional, Confidence: ConfidencePossible}},
		},
	},
	{
		Conformance: "[!(LT | DF)]",
		Tests: []conformanceTest{
			{Context: Context{Values: map[string]any{"AA": true}}, Expected: ConformanceState{State: StateOptional, Confidence: ConfidencePossible}},
			{Context: Context{Values: map[string]any{"LT": true}}, Expected: ConformanceState{State: StateOptional, Confidence: ConfidencePossible}},
			{Context: Context{Values: map[string]any{"DF": true}}, Expected: ConformanceState{State: StateOptional, Confidence: ConfidencePossible}},
			{Context: Context{Values: map[string]any{"LT": true, "DF": true}}, Expected: ConformanceState{State: StateDisallowed, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"LT": false, "DF": true}}, Expected: ConformanceState{State: StateDisallowed, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"LT": true, "DF": false}}, Expected: ConformanceState{State: StateDisallowed, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"LT": false, "DF": false}}, Expected: ConformanceState{State: StateOptional, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"Matter": true}}, Expected: ConformanceState{State: StateOptional, Confidence: ConfidencePossible}},
		},
	},
	{
		Conformance: "[!(LT | DF | CF)]",
		Tests: []conformanceTest{
			{Context: Context{Values: map[string]any{"AA": true}}, Expected: ConformanceState{State: StateOptional, Confidence: ConfidencePossible}},
			{Context: Context{Values: map[string]any{"LT": true}}, Expected: ConformanceState{State: StateOptional, Confidence: ConfidencePossible}},
			{Context: Context{Values: map[string]any{"DF": true}}, Expected: ConformanceState{State: StateOptional, Confidence: ConfidencePossible}},
			{Context: Context{Values: map[string]any{"LT": true, "DF": true, "CF": true}}, Expected: ConformanceState{State: StateDisallowed, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"LT": false, "DF": true, "CF": true}}, Expected: ConformanceState{State: StateDisallowed, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"LT": true, "DF": false, "CF": true}}, Expected: ConformanceState{State: StateDisallowed, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"LT": false, "DF": false, "CF": true}}, Expected: ConformanceState{State: StateDisallowed, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"LT": true, "DF": true, "CF": false}}, Expected: ConformanceState{State: StateDisallowed, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"LT": false, "DF": true, "CF": false}}, Expected: ConformanceState{State: StateDisallowed, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"LT": true, "DF": false, "CF": false}}, Expected: ConformanceState{State: StateDisallowed, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"LT": false, "DF": false, "CF": false}}, Expected: ConformanceState{State: StateOptional, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"Matter": true}}, Expected: ConformanceState{State: StateOptional, Confidence: ConfidencePossible}},
		},
	},
	{
		Conformance: "[LT | DF]",
		Tests: []conformanceTest{
			{Context: Context{Values: map[string]any{"AA": true}}, Expected: ConformanceState{State: StateOptional, Confidence: ConfidencePossible}},
			{Context: Context{Values: map[string]any{"LT": true}}, Expected: ConformanceState{State: StateOptional, Confidence: ConfidencePossible}},
			{Context: Context{Values: map[string]any{"DF": true}}, Expected: ConformanceState{State: StateOptional, Confidence: ConfidencePossible}},
			{Context: Context{Values: map[string]any{"LT": true, "DF": true}}, Expected: ConformanceState{State: StateOptional, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"LT": true, "DF": false}}, Expected: ConformanceState{State: StateOptional, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"LT": false, "DF": true}}, Expected: ConformanceState{State: StateOptional, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"LT": false, "DF": false}}, Expected: ConformanceState{State: StateDisallowed, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"Matter": true}}, Expected: ConformanceState{State: StateOptional, Confidence: ConfidencePossible}},
		},
	},

	{
		Conformance: "UltrasonicUnoccupiedToOccupiedThreshold, O",
		Tests: []conformanceTest{
			{Context: Context{Values: map[string]any{"AA": true}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidencePossible}},
			{Context: Context{Values: map[string]any{"UltrasonicUnoccupiedToOccupiedThreshold": true}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"UltrasonicUnoccupiedToOccupiedThreshold": false}}, Expected: ConformanceState{State: StateOptional, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"Matter": true}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidencePossible}},
		},
	},
	{
		Conformance: "Zigbee",
		Tests: []conformanceTest{
			{Context: Context{Values: map[string]any{"AA": true}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidencePossible}},
			{Context: Context{Values: map[string]any{"Zigbee": true}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"Zigbee": false}}, Expected: ConformanceState{State: StateDisallowed, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"Matter": true}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidencePossible}},
		},
	},
	{
		Conformance: "[Zigbee]",
		Tests: []conformanceTest{
			{Context: Context{Values: map[string]any{"AA": true}}, Expected: ConformanceState{State: StateOptional, Confidence: ConfidencePossible}},
			{Context: Context{Values: map[string]any{"Zigbee": true}}, Expected: ConformanceState{State: StateOptional, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"Zigbee": false}}, Expected: ConformanceState{State: StateDisallowed, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"Matter": true}}, Expected: ConformanceState{State: StateOptional, Confidence: ConfidencePossible}},
		},
	},
	{
		Conformance: "MSCH",
		Tests: []conformanceTest{
			{Context: Context{Values: map[string]any{"AA": true}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidencePossible}},
			{Context: Context{Values: map[string]any{"MSCH": true}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"MSCH": false}}, Expected: ConformanceState{State: StateDisallowed, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"Matter": true}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidencePossible}},
		},
	},
	{
		Conformance: "M",
		Tests: []conformanceTest{
			{Context: Context{Values: map[string]any{"AA": true}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"MSCH": true}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"MSCH": false}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"Matter": true}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceDefinite}},
		},
	},
	{
		Conformance: "(VIS | AUD) & SPRS",
		Tests: []conformanceTest{
			{Context: Context{Values: map[string]any{"VIS": true}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidencePossible}},
			{Context: Context{Values: map[string]any{"AUD": true}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidencePossible}},
			{Context: Context{Values: map[string]any{"VIS": true, "AUD": true}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidencePossible}},
			{Context: Context{Values: map[string]any{"VIS": true, "AUD": true, "SPRS": true}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"VIS": true, "AUD": true, "SPRS": false}}, Expected: ConformanceState{State: StateDisallowed, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"VIS": false, "AUD": true, "SPRS": true}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"VIS": true, "AUD": false, "SPRS": false}}, Expected: ConformanceState{State: StateDisallowed, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"VIS": false, "AUD": false, "SPRS": true}}, Expected: ConformanceState{State: StateDisallowed, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"VIS": true, "SPRS": true}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidencePossible}},
			{Context: Context{Values: map[string]any{"AUD": true, "SPRS": true}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidencePossible}},
			{Context: Context{Values: map[string]any{"SPRS": true}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidencePossible}},
			{Context: Context{Values: map[string]any{"Matter": true}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidencePossible}},
		},
	},
	{
		Conformance: "UltrasonicUnoccupiedToOccupiedDelay, O",
		Tests: []conformanceTest{
			{Context: Context{Values: map[string]any{"AA": true}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidencePossible}},
			{Context: Context{Values: map[string]any{"UltrasonicUnoccupiedToOccupiedDelay": true}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"UltrasonicUnoccupiedToOccupiedDelay": false}}, Expected: ConformanceState{State: StateOptional, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"Matter": true}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidencePossible}},
		},
	},
	{
		Conformance: "PIRUnoccupiedToOccupiedThreshold, O",
		Tests: []conformanceTest{
			{Context: Context{Values: map[string]any{"AA": true}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidencePossible}},
			{Context: Context{Values: map[string]any{"PIRUnoccupiedToOccupiedThreshold": true}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"PIRUnoccupiedToOccupiedThreshold": false}}, Expected: ConformanceState{State: StateOptional, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"Matter": true}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidencePossible}},
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
			{Context: Context{Values: map[string]any{"AA": true}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"BB": true}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidencePossible}},
			{Context: Context{Values: map[string]any{"AA": false, "BB": true}}, Expected: ConformanceState{State: StateOptional, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"CC": false}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidencePossible}},
			{Context: Context{Values: map[string]any{"Matter": true}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidencePossible}},
		},
	},
	{
		Conformance: "[Wi-Fi]",
		Tests: []conformanceTest{
			{Context: Context{Values: map[string]any{"AA": true}}, Expected: ConformanceState{State: StateOptional, Confidence: ConfidencePossible}},
			{Context: Context{Values: map[string]any{"Wi-Fi": true}}, Expected: ConformanceState{State: StateOptional, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"Wi-Fi": false}}, Expected: ConformanceState{State: StateDisallowed, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"Matter": true}}, Expected: ConformanceState{State: StateOptional, Confidence: ConfidencePossible}},
		},
	},
	{
		Conformance:    "<<ref_Ranges>>",
		ASCIIDocString: "<<ref_Ranges>>",
		Tests: []conformanceTest{
			{Context: Context{Values: map[string]any{"AA": true}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceImpossible}},
			{Context: Context{Values: map[string]any{"Wi-Fi": true}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceImpossible}},
			{Context: Context{Values: map[string]any{"UltrasonicUnoccupiedToOccupiedDelay": false}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceImpossible}},
			{Context: Context{Values: map[string]any{"Matter": true}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceImpossible}},
		},
	},
	{
		Conformance:    "<<ref_Ranges, Ranges>>",
		ASCIIDocString: "<<ref_Ranges, Ranges>>",
		Tests: []conformanceTest{
			{Context: Context{Values: map[string]any{"AA": true}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceImpossible}},
			{Context: Context{Values: map[string]any{"ref_Ranges": true}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceImpossible}},
			{Context: Context{Values: map[string]any{"ref_Ranges": false}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceImpossible}},
			{Context: Context{Values: map[string]any{"Matter": true}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceImpossible}},
		},
	},
	{
		Conformance:    "P, WATTS",
		ASCIIDocString: "P, WATTS",
		Tests: []conformanceTest{
			{Context: Context{Values: map[string]any{"AA": true}}, Expected: ConformanceState{State: StateProvisional, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"Wi-Fi": true}}, Expected: ConformanceState{State: StateProvisional, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"WATTS": false}}, Expected: ConformanceState{State: StateProvisional, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"Matter": true}}, Expected: ConformanceState{State: StateProvisional, Confidence: ConfidenceDefinite}},
		},
	},
	{
		Conformance:    "M, D",
		ASCIIDocString: "M, D", // Preserve the deprecated after mandatory
		Tests: []conformanceTest{
			{Context: Context{Values: map[string]any{"AA": true}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"Wi-Fi": true}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"WATTS": false}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"Matter": true}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceDefinite}},
		},
	},
	{
		Conformance:    "O, D",
		ASCIIDocString: "O, D", // Preserve the deprecated after optional
		Tests: []conformanceTest{
			{Context: Context{Values: map[string]any{"AA": true}}, Expected: ConformanceState{State: StateOptional, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"Wi-Fi": true}}, Expected: ConformanceState{State: StateOptional, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"WATTS": false}}, Expected: ConformanceState{State: StateOptional, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"Matter": true}}, Expected: ConformanceState{State: StateOptional, Confidence: ConfidenceDefinite}},
		},
	},
	{
		Conformance:    "M, [AA], D",
		ASCIIDocString: "M, D", // Remove the nonsensical optional conformance after mandatory, but preserve the deprecated
		Tests: []conformanceTest{
			{Context: Context{Values: map[string]any{"AA": true}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"Wi-Fi": true}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"WATTS": false}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceDefinite}},
			{Context: Context{Values: map[string]any{"Matter": true}}, Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceDefinite}},
		},
	},
}

func TestConformance(t *testing.T) {
	for _, test := range conformanceTests {
		test.run(t)
	}
}
