package conformance

import (
	"testing"
)

func TestOptional(t *testing.T) {
	conformance, err := tryParseConformance("[!AB & (CD != EF)], O")
	if err != nil {

		t.Errorf("failed parsing: %v", err)
	}
	t.Logf("conformance: %s", conformance.String())
}

type conformanceTestSuite struct {
	Conformance string
	Tests       []conformanceTest
}

func (cts *conformanceTestSuite) run(t *testing.T) {
	conformance, err := tryParseConformance(cts.Conformance)
	t.Logf("testing %s: %T", cts.Conformance, conformance)
	if err != nil {
		t.Errorf("failed parsing conformance %s: %v", cts.Conformance, err)
		return
	}
	if cs, ok := conformance.(ConformanceSet); ok {
		t.Logf("\tconformance set: %d", len(cs))
		for _, c := range cs {
			t.Logf("\ttesting %s: %T %v", cts.Conformance, c, c)
		}
	}
	for _, test := range cts.Tests {
		result, err := conformance.Eval(test.Context)
		if err != nil {
			t.Errorf("failed evaluating conformance status %v: %v", test.Context, err)
			return
		}
		if result != test.Expected {
			t.Errorf("failed checking conformance %s (parsed %s) with %v: expected %v, got %v", cts.Conformance, conformance.String(), test.Context, test.Expected, result)
		}
	}
}

type conformanceTest struct {
	Context  ConformanceContext
	Expected ConformanceState
}

var otherwiseTests = []conformanceTestSuite{
	{
		Conformance: "[LT | DF & CF]",
		Tests: []conformanceTest{
			{Context: ConformanceContext{Values: map[string]any{"AA": true}}, Expected: ConformanceStateDisallowed},
			{Context: ConformanceContext{Values: map[string]any{"LT": true}}, Expected: ConformanceStateOptional},
			{Context: ConformanceContext{Values: map[string]any{"DF": true}}, Expected: ConformanceStateOptional},
			{Context: ConformanceContext{Values: map[string]any{"Matter": true}}, Expected: ConformanceStateDisallowed},
		},
	},
	{
		Conformance: "AB, [CD]",
		Tests: []conformanceTest{
			{Context: ConformanceContext{Values: map[string]any{"AB": true}}, Expected: ConformanceStateMandatory},
			{Context: ConformanceContext{Values: map[string]any{"CD": true}}, Expected: ConformanceStateOptional},
			{Context: ConformanceContext{Values: map[string]any{"Matter": true}}, Expected: ConformanceStateDisallowed},
		},
	},
	{
		Conformance: "!AB, O",
		Tests: []conformanceTest{
			{Context: ConformanceContext{Values: map[string]any{"AB": true}}, Expected: ConformanceStateOptional},
			{Context: ConformanceContext{Values: map[string]any{"CD": true}}, Expected: ConformanceStateMandatory},
			{Context: ConformanceContext{Values: map[string]any{"Matter": true}}, Expected: ConformanceStateMandatory},
		},
	},
	{
		Conformance: "[AA], [BB], [CC]",
		Tests: []conformanceTest{
			{Context: ConformanceContext{Values: map[string]any{"AA": true}}, Expected: ConformanceStateOptional},
			{Context: ConformanceContext{Values: map[string]any{"BB": true}}, Expected: ConformanceStateOptional},
			{Context: ConformanceContext{Values: map[string]any{"CC": true}}, Expected: ConformanceStateOptional},
			{Context: ConformanceContext{Values: map[string]any{"Matter": true}}, Expected: ConformanceStateDisallowed},
		},
	},
	{
		Conformance: "[!(LT | DF)]",
		Tests: []conformanceTest{
			{Context: ConformanceContext{Values: map[string]any{"AA": true}}, Expected: ConformanceStateOptional},
			{Context: ConformanceContext{Values: map[string]any{"LT": true}}, Expected: ConformanceStateDisallowed},
			{Context: ConformanceContext{Values: map[string]any{"DF": true}}, Expected: ConformanceStateDisallowed},
			{Context: ConformanceContext{Values: map[string]any{"Matter": true}}, Expected: ConformanceStateOptional},
		},
	},
	{
		Conformance: "[!(LT | DF | CF)]",
		Tests: []conformanceTest{
			{Context: ConformanceContext{Values: map[string]any{"AA": true}}, Expected: ConformanceStateOptional},
			{Context: ConformanceContext{Values: map[string]any{"LT": true}}, Expected: ConformanceStateDisallowed},
			{Context: ConformanceContext{Values: map[string]any{"DF": true}}, Expected: ConformanceStateDisallowed},
			{Context: ConformanceContext{Values: map[string]any{"CF": true}}, Expected: ConformanceStateDisallowed},
			{Context: ConformanceContext{Values: map[string]any{"Matter": true}}, Expected: ConformanceStateOptional},
		},
	},
	{
		Conformance: "[LT | DF]",
		Tests: []conformanceTest{
			{Context: ConformanceContext{Values: map[string]any{"AA": true}}, Expected: ConformanceStateDisallowed},
			{Context: ConformanceContext{Values: map[string]any{"LT": true}}, Expected: ConformanceStateOptional},
			{Context: ConformanceContext{Values: map[string]any{"DF": true}}, Expected: ConformanceStateOptional},
			{Context: ConformanceContext{Values: map[string]any{"Matter": true}}, Expected: ConformanceStateDisallowed},
		},
	},

	{
		Conformance: "UltrasonicUnoccupiedToOccupiedThreshold, O",
		Tests: []conformanceTest{
			{Context: ConformanceContext{Values: map[string]any{"AA": true}}, Expected: ConformanceStateOptional},
			{Context: ConformanceContext{Values: map[string]any{"UltrasonicUnoccupiedToOccupiedThreshold": true}}, Expected: ConformanceStateMandatory},
			{Context: ConformanceContext{Values: map[string]any{"UltrasonicUnoccupiedToOccupiedThreshold": false}}, Expected: ConformanceStateOptional},
			{Context: ConformanceContext{Values: map[string]any{"Matter": true}}, Expected: ConformanceStateOptional},
		},
	},
	{
		Conformance: "Zigbee",
		Tests: []conformanceTest{
			{Context: ConformanceContext{Values: map[string]any{"AA": true}}, Expected: ConformanceStateDisallowed},
			{Context: ConformanceContext{Values: map[string]any{"Zigbee": true}}, Expected: ConformanceStateMandatory},
			{Context: ConformanceContext{Values: map[string]any{"Zigbee": false}}, Expected: ConformanceStateDisallowed},
			{Context: ConformanceContext{Values: map[string]any{"Matter": true}}, Expected: ConformanceStateDisallowed},
		},
	},
	{
		Conformance: "[Zigbee]",
		Tests: []conformanceTest{
			{Context: ConformanceContext{Values: map[string]any{"AA": true}}, Expected: ConformanceStateDisallowed},
			{Context: ConformanceContext{Values: map[string]any{"Zigbee": true}}, Expected: ConformanceStateOptional},
			{Context: ConformanceContext{Values: map[string]any{"Zigbee": false}}, Expected: ConformanceStateDisallowed},
			{Context: ConformanceContext{Values: map[string]any{"Matter": true}}, Expected: ConformanceStateDisallowed},
		},
	},
	{
		Conformance: "MSCH",
		Tests: []conformanceTest{
			{Context: ConformanceContext{Values: map[string]any{"AA": true}}, Expected: ConformanceStateDisallowed},
			{Context: ConformanceContext{Values: map[string]any{"MSCH": true}}, Expected: ConformanceStateMandatory},
			{Context: ConformanceContext{Values: map[string]any{"MSCH": false}}, Expected: ConformanceStateDisallowed},
			{Context: ConformanceContext{Values: map[string]any{"Matter": true}}, Expected: ConformanceStateDisallowed},
		},
	},
	{
		Conformance: "M",
		Tests: []conformanceTest{
			{Context: ConformanceContext{Values: map[string]any{"AA": true}}, Expected: ConformanceStateMandatory},
			{Context: ConformanceContext{Values: map[string]any{"MSCH": true}}, Expected: ConformanceStateMandatory},
			{Context: ConformanceContext{Values: map[string]any{"MSCH": false}}, Expected: ConformanceStateMandatory},
			{Context: ConformanceContext{Values: map[string]any{"Matter": true}}, Expected: ConformanceStateMandatory},
		},
	},
	{
		Conformance: "(VIS | AUD) & SPRS",
		Tests: []conformanceTest{
			{Context: ConformanceContext{Values: map[string]any{"VIS": true}}, Expected: ConformanceStateDisallowed},
			{Context: ConformanceContext{Values: map[string]any{"AUD": true}}, Expected: ConformanceStateDisallowed},
			{Context: ConformanceContext{Values: map[string]any{"VIS": true, "AUD": true}}, Expected: ConformanceStateDisallowed},
			{Context: ConformanceContext{Values: map[string]any{"VIS": true, "AUD": true, "SPRS": true}}, Expected: ConformanceStateMandatory},
			{Context: ConformanceContext{Values: map[string]any{"VIS": true, "SPRS": true}}, Expected: ConformanceStateMandatory},
			{Context: ConformanceContext{Values: map[string]any{"AUD": true, "SPRS": true}}, Expected: ConformanceStateMandatory},
			{Context: ConformanceContext{Values: map[string]any{"SPRS": true}}, Expected: ConformanceStateDisallowed},
			{Context: ConformanceContext{Values: map[string]any{"Matter": true}}, Expected: ConformanceStateDisallowed},
		},
	},
	{
		Conformance: "UltrasonicUnoccupiedToOccupiedDelay, O",
		Tests: []conformanceTest{
			{Context: ConformanceContext{Values: map[string]any{"AA": true}}, Expected: ConformanceStateOptional},
			{Context: ConformanceContext{Values: map[string]any{"UltrasonicUnoccupiedToOccupiedDelay": true}}, Expected: ConformanceStateMandatory},
			{Context: ConformanceContext{Values: map[string]any{"UltrasonicUnoccupiedToOccupiedDelay": false}}, Expected: ConformanceStateOptional},
			{Context: ConformanceContext{Values: map[string]any{"Matter": true}}, Expected: ConformanceStateOptional},
		},
	},
	{
		Conformance: "PIRUnoccupiedToOccupiedThreshold, O",
		Tests: []conformanceTest{
			{Context: ConformanceContext{Values: map[string]any{"AA": true}}, Expected: ConformanceStateOptional},
			{Context: ConformanceContext{Values: map[string]any{"PIRUnoccupiedToOccupiedThreshold": true}}, Expected: ConformanceStateMandatory},
			{Context: ConformanceContext{Values: map[string]any{"UltrasonicUnoccupiedToOccupiedDelay": false}}, Expected: ConformanceStateOptional},
			{Context: ConformanceContext{Values: map[string]any{"Matter": true}}, Expected: ConformanceStateOptional},
		},
	},
}

func TestOtherwise(t *testing.T) {
	for _, test := range otherwiseTests {
		test.run(t)
		break
	}
}
