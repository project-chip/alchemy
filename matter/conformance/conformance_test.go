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
	if cts.InvalidConformance {
		t.Errorf("parsed conformance that should have been invalid %s: %v", cts.Conformance, err)
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

type conformanceTestContext struct {
	BasicContext
}

func (c *conformanceTestContext) Value(identifier string) (any, bool) {
	val, ok := c.Values[identifier]
	if !ok {
		return ConfidencePossible, true
	}
	return val, true
}

func makeTestContext(args ...any) *conformanceTestContext {
	ctc := conformanceTestContext{
		BasicContext{Values: map[string]any{}},
	}
	pairCount := len(args) / 2
	for i := 0; i < pairCount; i++ {
		key := args[i*2].(string)
		ctc.Values[key] = args[i*2+1]
	}
	return &ctc
}

var conformanceTests = []conformanceTestSuite{
	{
		Conformance: "[v2 < Rev]",
		Tests: []conformanceTest{
			{Context: makeTestContext("Revision", Revision(1)), Expected: ConformanceState{State: StateDisallowed, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("Revision", Revision(2)), Expected: ConformanceState{State: StateDisallowed, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("Revision", Revision(3)), Expected: ConformanceState{State: StateOptional, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("Matter", true), Expected: ConformanceState{State: StateOptional, Confidence: ConfidencePossible}},
		},
	},

	{
		Conformance: "v1 <= Rev",
		Tests: []conformanceTest{
			{Context: makeTestContext("Revision", Revision(1)), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("Revision", Revision(2)), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("Revision", Revision(3)), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("Matter", true), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidencePossible}},
		},
	},
	{
		Conformance: "[v1 <= Rev]",
		Tests: []conformanceTest{
			{Context: makeTestContext("Revision", Revision(1)), Expected: ConformanceState{State: StateOptional, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("Revision", Revision(2)), Expected: ConformanceState{State: StateOptional, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("Revision", Revision(3)), Expected: ConformanceState{State: StateOptional, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("Matter", true), Expected: ConformanceState{State: StateOptional, Confidence: ConfidencePossible}},
		},
	},
	{
		Conformance: "[Rev >= v1]",
		Tests: []conformanceTest{
			{Context: makeTestContext("Revision", Revision(1)), Expected: ConformanceState{State: StateOptional, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("Revision", Revision(2)), Expected: ConformanceState{State: StateOptional, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("Revision", Revision(3)), Expected: ConformanceState{State: StateOptional, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("Matter", true), Expected: ConformanceState{State: StateOptional, Confidence: ConfidencePossible}},
		},
	},
	{
		Conformance: "[v2 < Rev < v4]",
		Tests: []conformanceTest{
			{Context: makeTestContext("Revision", Revision(1)), Expected: ConformanceState{State: StateDisallowed, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("Revision", Revision(2)), Expected: ConformanceState{State: StateDisallowed, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("Revision", Revision(3)), Expected: ConformanceState{State: StateOptional, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("Revision", Revision(4)), Expected: ConformanceState{State: StateDisallowed, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("Matter", true), Expected: ConformanceState{State: StateOptional, Confidence: ConfidencePossible}},
		},
	},
	{
		Conformance: "[v2 <= Rev < v4]",
		Tests: []conformanceTest{
			{Context: makeTestContext("Revision", Revision(1)), Expected: ConformanceState{State: StateDisallowed, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("Revision", Revision(2)), Expected: ConformanceState{State: StateOptional, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("Revision", Revision(3)), Expected: ConformanceState{State: StateOptional, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("Revision", Revision(4)), Expected: ConformanceState{State: StateDisallowed, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("Matter", true), Expected: ConformanceState{State: StateOptional, Confidence: ConfidencePossible}},
		},
	},
	{
		Conformance: "[v4 >= Rev > v2]",
		Tests: []conformanceTest{
			{Context: makeTestContext("Revision", Revision(1)), Expected: ConformanceState{State: StateDisallowed, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("Revision", Revision(2)), Expected: ConformanceState{State: StateDisallowed, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("Revision", Revision(3)), Expected: ConformanceState{State: StateOptional, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("Revision", Revision(4)), Expected: ConformanceState{State: StateOptional, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("Matter", true), Expected: ConformanceState{State: StateOptional, Confidence: ConfidencePossible}},
		},
	},
	{
		Conformance: "Rev > v4, [Rev >= v2], D",
		Tests: []conformanceTest{
			{Context: makeTestContext("Revision", Revision(1)), Expected: ConformanceState{State: StateDeprecated, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("Revision", Revision(2)), Expected: ConformanceState{State: StateOptional, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("Revision", Revision(3)), Expected: ConformanceState{State: StateOptional, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("Revision", Revision(4)), Expected: ConformanceState{State: StateOptional, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("Revision", Revision(5)), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("Matter", true), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidencePossible}},
		},
	},
	{
		Conformance:        "v2 > v3",
		InvalidConformance: true,
	},
	{
		Conformance:        "Rev > Rev",
		InvalidConformance: true,
	},
	{
		Conformance:        "Rev > v0",
		InvalidConformance: true,
	},
	{
		Conformance:        "Rev > 0",
		InvalidConformance: true,
	},
	{
		Conformance:        "v2 > v3 > v5",
		InvalidConformance: true,
	},
	{
		Conformance:        "v2 > Rev > v5",
		InvalidConformance: true,
	},
	{
		Conformance:        "v2 < Rev == v5",
		InvalidConformance: true,
	},
	{
		Conformance:        "v2 < Rev != v5",
		InvalidConformance: true,
	},
	{
		Conformance: "[AB & CM].a2",
		Tests: []conformanceTest{
			{Context: makeTestContext("AB", true), Expected: ConformanceState{State: StateOptional, Confidence: ConfidencePossible}},
			{Context: makeTestContext("AB", true, "CM", true), Expected: ConformanceState{State: StateOptional, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("AB", true, "CM", false), Expected: ConformanceState{State: StateDisallowed, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("Matter", true), Expected: ConformanceState{State: StateOptional, Confidence: ConfidencePossible}},
		},
	},
	{
		Conformance: "RequiresEncodedPixels==true",
		Tests: []conformanceTest{
			{Context: makeTestContext("RequiresEncodedPixels", false), Expected: ConformanceState{State: StateDisallowed, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("RequiresEncodedPixels", true), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceDefinite}},
		},
	},
	{
		Conformance: `OperationalStateID >= 0x80 & OperationalStateID \<= 0xBF`,
		Tests: []conformanceTest{
			{Context: makeTestContext("OperationalStateID", uint64(0x7A)), Expected: ConformanceState{State: StateDisallowed, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("OperationalStateID", uint64(0x80)), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("OperationalStateID", uint64(0xBF)), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("OperationalStateID", uint64(0xC0)), Expected: ConformanceState{State: StateDisallowed, Confidence: ConfidenceDefinite}},
		},
	},
	{
		Conformance: "OperationalStateID >= 0x80 & OperationalStateID <= 0xBF",
		Tests: []conformanceTest{
			{Context: makeTestContext("OperationalStateID", uint64(0x7A)), Expected: ConformanceState{State: StateDisallowed, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("OperationalStateID", uint64(0x80)), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("OperationalStateID", uint64(0xBF)), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("OperationalStateID", uint64(0xC0)), Expected: ConformanceState{State: StateDisallowed, Confidence: ConfidenceDefinite}},
		},
	},

	{
		Conformance: "AB, O.a1+",
		Tests: []conformanceTest{
			{Context: makeTestContext("AB", true), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("Matter", true), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidencePossible}},
		},
	},
	{
		Conformance: "O.a1+",
		Tests: []conformanceTest{
			{Context: makeTestContext("Matter", true), Expected: ConformanceState{State: StateOptional, Confidence: ConfidenceDefinite}},
		},
	},
	{
		Conformance: "NumberOfPrimaries > 0, O",
		Tests: []conformanceTest{
			{Context: makeTestContext("NumberOfPrimaries", int64(1)), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("NumberOfPrimaries", int64(0)), Expected: ConformanceState{State: StateOptional, Confidence: ConfidenceDefinite}},
		},
	},
	{
		Conformance: "TwoDCART",
		Tests: []conformanceTest{
			{Context: makeTestContext("TwoDCART", true), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("TwoDCART", false), Expected: ConformanceState{State: StateDisallowed, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("SFR", true), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidencePossible}},
			{Context: makeTestContext("Matter", true), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidencePossible}},
		},
	},
	{
		Conformance: "(STA|PAU|FA|CON)&!SFR,!PA&!SFR,O.a-",
		Tests: []conformanceTest{
			{Context: makeTestContext("STA", true), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidencePossible}},
			{Context: makeTestContext("STA", true, "SFR", true), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidencePossible}},
			{Context: makeTestContext("STA", true, "PAU", true, "FA", true, "CON", true, "SFR", true, "PA", true), Expected: ConformanceState{State: StateOptional, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("STA", true, "PAU", true, "FA", true, "CON", true, "SFR", false, "PA", true), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("STA", false, "PAU", false, "FA", false, "CON", false, "SFR", false, "PA", false), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("PA", true), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidencePossible}},
			{Context: makeTestContext("CF", true), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidencePossible}},
			{Context: makeTestContext("SFR", true), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidencePossible}},
			{Context: makeTestContext("Matter", true), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidencePossible}},
		},
	},
	{

		Conformance: "[LT | DF & CF]",
		Tests: []conformanceTest{
			{Context: makeTestContext("AA", true), Expected: ConformanceState{State: StateOptional, Confidence: ConfidencePossible}},
			{Context: makeTestContext("LT", true), Expected: ConformanceState{State: StateOptional, Confidence: ConfidencePossible}},
			{Context: makeTestContext("LT", true, "DF", false, "CF", false), Expected: ConformanceState{State: StateOptional, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("LT", true, "DF", true, "CF", false), Expected: ConformanceState{State: StateOptional, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("LT", false, "DF", true, "CF", true), Expected: ConformanceState{State: StateOptional, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("LT", false, "DF", false, "CF", true), Expected: ConformanceState{State: StateDisallowed, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("LT", false, "DF", true, "CF", false), Expected: ConformanceState{State: StateDisallowed, Confidence: ConfidenceDefinite}},
		},
	},
	{
		Conformance: "AB, [CD]",
		Tests: []conformanceTest{
			{Context: makeTestContext("AB", true), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("CD", true), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidencePossible}},
			{Context: makeTestContext("AB", false, "CD", true), Expected: ConformanceState{State: StateOptional, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("AB", false, "CD", false), Expected: ConformanceState{State: StateDisallowed, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("Matter", true), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidencePossible}},
		},
	},
	{
		Conformance: "!AB, O",
		Tests: []conformanceTest{
			{Context: makeTestContext("AB", true), Expected: ConformanceState{State: StateOptional, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("AB", false), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("CD", true), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidencePossible}},
			{Context: makeTestContext("Matter", true), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidencePossible}},
		},
	},
	{
		Conformance: "[AA], [BB], [CC]",
		Tests: []conformanceTest{
			{Context: makeTestContext("AA", true), Expected: ConformanceState{State: StateOptional, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("BB", true), Expected: ConformanceState{State: StateOptional, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("CC", true), Expected: ConformanceState{State: StateOptional, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("AA", false, "BB", false, "CC", false), Expected: ConformanceState{State: StateDisallowed, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("Matter", true), Expected: ConformanceState{State: StateOptional, Confidence: ConfidencePossible}},
		},
	},
	{
		Conformance: "[!(LT | DF)]",
		Tests: []conformanceTest{
			{Context: makeTestContext("AA", true), Expected: ConformanceState{State: StateOptional, Confidence: ConfidencePossible}},
			{Context: makeTestContext("LT", true), Expected: ConformanceState{State: StateOptional, Confidence: ConfidencePossible}},
			{Context: makeTestContext("DF", true), Expected: ConformanceState{State: StateOptional, Confidence: ConfidencePossible}},
			{Context: makeTestContext("LT", true, "DF", true), Expected: ConformanceState{State: StateDisallowed, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("LT", false, "DF", true), Expected: ConformanceState{State: StateDisallowed, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("LT", true, "DF", false), Expected: ConformanceState{State: StateDisallowed, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("LT", false, "DF", false), Expected: ConformanceState{State: StateOptional, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("Matter", true), Expected: ConformanceState{State: StateOptional, Confidence: ConfidencePossible}},
		},
	},
	{
		Conformance: "[!(LT | DF | CF)]",
		Tests: []conformanceTest{
			{Context: makeTestContext("AA", true), Expected: ConformanceState{State: StateOptional, Confidence: ConfidencePossible}},
			{Context: makeTestContext("LT", true), Expected: ConformanceState{State: StateOptional, Confidence: ConfidencePossible}},
			{Context: makeTestContext("DF", true), Expected: ConformanceState{State: StateOptional, Confidence: ConfidencePossible}},
			{Context: makeTestContext("LT", true, "DF", true, "CF", true), Expected: ConformanceState{State: StateDisallowed, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("LT", false, "DF", true, "CF", true), Expected: ConformanceState{State: StateDisallowed, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("LT", true, "DF", false, "CF", true), Expected: ConformanceState{State: StateDisallowed, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("LT", false, "DF", false, "CF", true), Expected: ConformanceState{State: StateDisallowed, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("LT", true, "DF", true, "CF", false), Expected: ConformanceState{State: StateDisallowed, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("LT", false, "DF", true, "CF", false), Expected: ConformanceState{State: StateDisallowed, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("LT", true, "DF", false, "CF", false), Expected: ConformanceState{State: StateDisallowed, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("LT", false, "DF", false, "CF", false), Expected: ConformanceState{State: StateOptional, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("Matter", true), Expected: ConformanceState{State: StateOptional, Confidence: ConfidencePossible}},
		},
	},
	{
		Conformance: "[LT | DF]",
		Tests: []conformanceTest{
			{Context: makeTestContext("AA", true), Expected: ConformanceState{State: StateOptional, Confidence: ConfidencePossible}},
			{Context: makeTestContext("LT", true), Expected: ConformanceState{State: StateOptional, Confidence: ConfidencePossible}},
			{Context: makeTestContext("DF", true), Expected: ConformanceState{State: StateOptional, Confidence: ConfidencePossible}},
			{Context: makeTestContext("LT", true, "DF", true), Expected: ConformanceState{State: StateOptional, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("LT", true, "DF", false), Expected: ConformanceState{State: StateOptional, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("LT", false, "DF", true), Expected: ConformanceState{State: StateOptional, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("LT", false, "DF", false), Expected: ConformanceState{State: StateDisallowed, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("Matter", true), Expected: ConformanceState{State: StateOptional, Confidence: ConfidencePossible}},
		},
	},

	{
		Conformance: "UltrasonicUnoccupiedToOccupiedThreshold, O",
		Tests: []conformanceTest{
			{Context: makeTestContext("AA", true), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidencePossible}},
			{Context: makeTestContext("UltrasonicUnoccupiedToOccupiedThreshold", true), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("UltrasonicUnoccupiedToOccupiedThreshold", false), Expected: ConformanceState{State: StateOptional, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("Matter", true), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidencePossible}},
		},
	},
	{
		Conformance: "Zigbee",
		Tests: []conformanceTest{
			{Context: makeTestContext("AA", true), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidencePossible}},
			{Context: makeTestContext("Zigbee", true), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("Zigbee", false), Expected: ConformanceState{State: StateDisallowed, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("Matter", true), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidencePossible}},
		},
	},
	{
		Conformance: "[Zigbee]",
		Tests: []conformanceTest{
			{Context: makeTestContext("AA", true), Expected: ConformanceState{State: StateOptional, Confidence: ConfidencePossible}},
			{Context: makeTestContext("Zigbee", true), Expected: ConformanceState{State: StateOptional, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("Zigbee", false), Expected: ConformanceState{State: StateDisallowed, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("Matter", true), Expected: ConformanceState{State: StateOptional, Confidence: ConfidencePossible}},
		},
	},
	{
		Conformance: "MSCH",
		Tests: []conformanceTest{
			{Context: makeTestContext("AA", true), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidencePossible}},
			{Context: makeTestContext("MSCH", true), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("MSCH", false), Expected: ConformanceState{State: StateDisallowed, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("Matter", true), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidencePossible}},
		},
	},
	{
		Conformance: "M",
		Tests: []conformanceTest{
			{Context: makeTestContext("AA", true), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("MSCH", true), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("MSCH", false), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("Matter", true), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceDefinite}},
		},
	},
	{
		Conformance: "(VIS | AUD) & SPRS",
		Tests: []conformanceTest{
			{Context: makeTestContext("VIS", true), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidencePossible}},
			{Context: makeTestContext("AUD", true), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidencePossible}},
			{Context: makeTestContext("VIS", true, "AUD", true), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidencePossible}},
			{Context: makeTestContext("VIS", true, "AUD", true, "SPRS", true), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("VIS", true, "AUD", true, "SPRS", false), Expected: ConformanceState{State: StateDisallowed, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("VIS", false, "AUD", true, "SPRS", true), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("VIS", true, "AUD", false, "SPRS", false), Expected: ConformanceState{State: StateDisallowed, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("VIS", false, "AUD", false, "SPRS", true), Expected: ConformanceState{State: StateDisallowed, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("VIS", true, "SPRS", true), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidencePossible}},
			{Context: makeTestContext("AUD", true, "SPRS", true), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidencePossible}},
			{Context: makeTestContext("SPRS", true), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidencePossible}},
			{Context: makeTestContext("Matter", true), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidencePossible}},
		},
	},
	{
		Conformance: "UltrasonicUnoccupiedToOccupiedDelay, O",
		Tests: []conformanceTest{
			{Context: makeTestContext("AA", true), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidencePossible}},
			{Context: makeTestContext("UltrasonicUnoccupiedToOccupiedDelay", true), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("UltrasonicUnoccupiedToOccupiedDelay", false), Expected: ConformanceState{State: StateOptional, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("Matter", true), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidencePossible}},
		},
	},
	{
		Conformance: "PIRUnoccupiedToOccupiedThreshold, O",
		Tests: []conformanceTest{
			{Context: makeTestContext("AA", true), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidencePossible}},
			{Context: makeTestContext("PIRUnoccupiedToOccupiedThreshold", true), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("PIRUnoccupiedToOccupiedThreshold", false), Expected: ConformanceState{State: StateOptional, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("Matter", true), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidencePossible}},
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
			{Context: makeTestContext("AA", true), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("BB", true), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidencePossible}},
			{Context: makeTestContext("AA", false, "BB", true), Expected: ConformanceState{State: StateOptional, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("CC", false), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidencePossible}},
			{Context: makeTestContext("Matter", true), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidencePossible}},
		},
	},
	{
		Conformance: "[Wi-Fi]",
		Tests: []conformanceTest{
			{Context: makeTestContext("AA", true), Expected: ConformanceState{State: StateOptional, Confidence: ConfidencePossible}},
			{Context: makeTestContext("Wi-Fi", true), Expected: ConformanceState{State: StateOptional, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("Wi-Fi", false), Expected: ConformanceState{State: StateDisallowed, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("Matter", true), Expected: ConformanceState{State: StateOptional, Confidence: ConfidencePossible}},
		},
	},
	{
		Conformance:    "<<ref_Ranges>>",
		ASCIIDocString: "<<ref_Ranges>>",
		Tests: []conformanceTest{
			{Context: makeTestContext("AA", true), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceImpossible}},
			{Context: makeTestContext("Wi-Fi", true), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceImpossible}},
			{Context: makeTestContext("UltrasonicUnoccupiedToOccupiedDelay", false), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceImpossible}},
			{Context: makeTestContext("Matter", true), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceImpossible}},
		},
	},
	{
		Conformance:    "<<ref_Ranges, Ranges>>",
		ASCIIDocString: "<<ref_Ranges, Ranges>>",
		Tests: []conformanceTest{
			{Context: makeTestContext("AA", true), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceImpossible}},
			{Context: makeTestContext("ref_Ranges", true), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceImpossible}},
			{Context: makeTestContext("ref_Ranges", false), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceImpossible}},
			{Context: makeTestContext("Matter", true), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceImpossible}},
		},
	},
	{
		Conformance:    "P, WATTS",
		ASCIIDocString: "P, WATTS",
		Tests: []conformanceTest{
			{Context: makeTestContext("AA", true), Expected: ConformanceState{State: StateProvisional, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("Wi-Fi", true), Expected: ConformanceState{State: StateProvisional, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("WATTS", false), Expected: ConformanceState{State: StateProvisional, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("Matter", true), Expected: ConformanceState{State: StateProvisional, Confidence: ConfidenceDefinite}},
		},
	},
	{
		Conformance:    "M, D",
		ASCIIDocString: "M, D", // Preserve the deprecated after mandatory
		Tests: []conformanceTest{
			{Context: makeTestContext("AA", true), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("Wi-Fi", true), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("WATTS", false), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("Matter", true), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceDefinite}},
		},
	},
	{
		Conformance:    "O, D",
		ASCIIDocString: "O, D", // Preserve the deprecated after optional
		Tests: []conformanceTest{
			{Context: makeTestContext("AA", true), Expected: ConformanceState{State: StateOptional, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("Wi-Fi", true), Expected: ConformanceState{State: StateOptional, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("WATTS", false), Expected: ConformanceState{State: StateOptional, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("Matter", true), Expected: ConformanceState{State: StateOptional, Confidence: ConfidenceDefinite}},
		},
	},
	{
		Conformance:    "M, [AA], D",
		ASCIIDocString: "M, D", // Remove the nonsensical optional conformance after mandatory, but preserve the deprecated
		Tests: []conformanceTest{
			{Context: makeTestContext("AA", true), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("Wi-Fi", true), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("WATTS", false), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceDefinite}},
			{Context: makeTestContext("Matter", true), Expected: ConformanceState{State: StateMandatory, Confidence: ConfidenceDefinite}},
		},
	},
}

func TestConformance(t *testing.T) {
	for _, test := range conformanceTests {
		test.run(t)
	}
}
