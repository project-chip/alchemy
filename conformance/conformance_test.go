package conformance

import (
	"testing"

	"github.com/hasty/alchemy/matter"
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
	Context  matter.ConformanceContext
	Expected matter.ConformanceState
}

var otherwiseTests = []conformanceTestSuite{
	{
		Conformance: "[LT | DF & CF]",
		Tests: []conformanceTest{
			{Context: matter.ConformanceContext{Values: map[string]any{"AA": true}}, Expected: matter.ConformanceStateDisallowed},
			{Context: matter.ConformanceContext{Values: map[string]any{"LT": true}}, Expected: matter.ConformanceStateOptional},
			{Context: matter.ConformanceContext{Values: map[string]any{"DF": true}}, Expected: matter.ConformanceStateOptional},
			{Context: matter.ConformanceContext{Values: map[string]any{"Matter": true}}, Expected: matter.ConformanceStateDisallowed},
		},
	},
	{
		Conformance: "AB, [CD]",
		Tests: []conformanceTest{
			{Context: matter.ConformanceContext{Values: map[string]any{"AB": true}}, Expected: matter.ConformanceStateMandatory},
			{Context: matter.ConformanceContext{Values: map[string]any{"CD": true}}, Expected: matter.ConformanceStateOptional},
			{Context: matter.ConformanceContext{Values: map[string]any{"Matter": true}}, Expected: matter.ConformanceStateDisallowed},
		},
	},
	{
		Conformance: "!AB, O",
		Tests: []conformanceTest{
			{Context: matter.ConformanceContext{Values: map[string]any{"AB": true}}, Expected: matter.ConformanceStateOptional},
			{Context: matter.ConformanceContext{Values: map[string]any{"CD": true}}, Expected: matter.ConformanceStateMandatory},
			{Context: matter.ConformanceContext{Values: map[string]any{"Matter": true}}, Expected: matter.ConformanceStateMandatory},
		},
	},
	{
		Conformance: "[AA], [BB], [CC]",
		Tests: []conformanceTest{
			{Context: matter.ConformanceContext{Values: map[string]any{"AA": true}}, Expected: matter.ConformanceStateOptional},
			{Context: matter.ConformanceContext{Values: map[string]any{"BB": true}}, Expected: matter.ConformanceStateOptional},
			{Context: matter.ConformanceContext{Values: map[string]any{"CC": true}}, Expected: matter.ConformanceStateOptional},
			{Context: matter.ConformanceContext{Values: map[string]any{"Matter": true}}, Expected: matter.ConformanceStateDisallowed},
		},
	},
	{
		Conformance: "[!(LT | DF)]",
		Tests: []conformanceTest{
			{Context: matter.ConformanceContext{Values: map[string]any{"AA": true}}, Expected: matter.ConformanceStateOptional},
			{Context: matter.ConformanceContext{Values: map[string]any{"LT": true}}, Expected: matter.ConformanceStateDisallowed},
			{Context: matter.ConformanceContext{Values: map[string]any{"DF": true}}, Expected: matter.ConformanceStateDisallowed},
			{Context: matter.ConformanceContext{Values: map[string]any{"Matter": true}}, Expected: matter.ConformanceStateOptional},
		},
	},
	{
		Conformance: "[!(LT | DF | CF)]",
		Tests: []conformanceTest{
			{Context: matter.ConformanceContext{Values: map[string]any{"AA": true}}, Expected: matter.ConformanceStateOptional},
			{Context: matter.ConformanceContext{Values: map[string]any{"LT": true}}, Expected: matter.ConformanceStateDisallowed},
			{Context: matter.ConformanceContext{Values: map[string]any{"DF": true}}, Expected: matter.ConformanceStateDisallowed},
			{Context: matter.ConformanceContext{Values: map[string]any{"CF": true}}, Expected: matter.ConformanceStateDisallowed},
			{Context: matter.ConformanceContext{Values: map[string]any{"Matter": true}}, Expected: matter.ConformanceStateOptional},
		},
	},
	{
		Conformance: "[LT | DF]",
		Tests: []conformanceTest{
			{Context: matter.ConformanceContext{Values: map[string]any{"AA": true}}, Expected: matter.ConformanceStateDisallowed},
			{Context: matter.ConformanceContext{Values: map[string]any{"LT": true}}, Expected: matter.ConformanceStateOptional},
			{Context: matter.ConformanceContext{Values: map[string]any{"DF": true}}, Expected: matter.ConformanceStateOptional},
			{Context: matter.ConformanceContext{Values: map[string]any{"Matter": true}}, Expected: matter.ConformanceStateDisallowed},
		},
	},

	{
		Conformance: "UltrasonicUnoccupiedToOccupiedThreshold, O",
		Tests: []conformanceTest{
			{Context: matter.ConformanceContext{Values: map[string]any{"AA": true}}, Expected: matter.ConformanceStateOptional},
			{Context: matter.ConformanceContext{Values: map[string]any{"UltrasonicUnoccupiedToOccupiedThreshold": true}}, Expected: matter.ConformanceStateMandatory},
			{Context: matter.ConformanceContext{Values: map[string]any{"UltrasonicUnoccupiedToOccupiedThreshold": false}}, Expected: matter.ConformanceStateOptional},
			{Context: matter.ConformanceContext{Values: map[string]any{"Matter": true}}, Expected: matter.ConformanceStateOptional},
		},
	},
	{
		Conformance: "Zigbee",
		Tests: []conformanceTest{
			{Context: matter.ConformanceContext{Values: map[string]any{"AA": true}}, Expected: matter.ConformanceStateDisallowed},
			{Context: matter.ConformanceContext{Values: map[string]any{"Zigbee": true}}, Expected: matter.ConformanceStateMandatory},
			{Context: matter.ConformanceContext{Values: map[string]any{"Zigbee": false}}, Expected: matter.ConformanceStateDisallowed},
			{Context: matter.ConformanceContext{Values: map[string]any{"Matter": true}}, Expected: matter.ConformanceStateDisallowed},
		},
	},
	{
		Conformance: "[Zigbee]",
		Tests: []conformanceTest{
			{Context: matter.ConformanceContext{Values: map[string]any{"AA": true}}, Expected: matter.ConformanceStateDisallowed},
			{Context: matter.ConformanceContext{Values: map[string]any{"Zigbee": true}}, Expected: matter.ConformanceStateOptional},
			{Context: matter.ConformanceContext{Values: map[string]any{"Zigbee": false}}, Expected: matter.ConformanceStateDisallowed},
			{Context: matter.ConformanceContext{Values: map[string]any{"Matter": true}}, Expected: matter.ConformanceStateDisallowed},
		},
	},
	{
		Conformance: "MSCH",
		Tests: []conformanceTest{
			{Context: matter.ConformanceContext{Values: map[string]any{"AA": true}}, Expected: matter.ConformanceStateDisallowed},
			{Context: matter.ConformanceContext{Values: map[string]any{"MSCH": true}}, Expected: matter.ConformanceStateMandatory},
			{Context: matter.ConformanceContext{Values: map[string]any{"MSCH": false}}, Expected: matter.ConformanceStateDisallowed},
			{Context: matter.ConformanceContext{Values: map[string]any{"Matter": true}}, Expected: matter.ConformanceStateDisallowed},
		},
	},
	{
		Conformance: "M",
		Tests: []conformanceTest{
			{Context: matter.ConformanceContext{Values: map[string]any{"AA": true}}, Expected: matter.ConformanceStateMandatory},
			{Context: matter.ConformanceContext{Values: map[string]any{"MSCH": true}}, Expected: matter.ConformanceStateMandatory},
			{Context: matter.ConformanceContext{Values: map[string]any{"MSCH": false}}, Expected: matter.ConformanceStateMandatory},
			{Context: matter.ConformanceContext{Values: map[string]any{"Matter": true}}, Expected: matter.ConformanceStateMandatory},
		},
	},
	{
		Conformance: "(VIS | AUD) & SPRS",
		Tests: []conformanceTest{
			{Context: matter.ConformanceContext{Values: map[string]any{"VIS": true}}, Expected: matter.ConformanceStateDisallowed},
			{Context: matter.ConformanceContext{Values: map[string]any{"AUD": true}}, Expected: matter.ConformanceStateDisallowed},
			{Context: matter.ConformanceContext{Values: map[string]any{"VIS": true, "AUD": true}}, Expected: matter.ConformanceStateDisallowed},
			{Context: matter.ConformanceContext{Values: map[string]any{"VIS": true, "AUD": true, "SPRS": true}}, Expected: matter.ConformanceStateMandatory},
			{Context: matter.ConformanceContext{Values: map[string]any{"VIS": true, "SPRS": true}}, Expected: matter.ConformanceStateMandatory},
			{Context: matter.ConformanceContext{Values: map[string]any{"AUD": true, "SPRS": true}}, Expected: matter.ConformanceStateMandatory},
			{Context: matter.ConformanceContext{Values: map[string]any{"SPRS": true}}, Expected: matter.ConformanceStateDisallowed},
			{Context: matter.ConformanceContext{Values: map[string]any{"Matter": true}}, Expected: matter.ConformanceStateDisallowed},
		},
	},
	{
		Conformance: "UltrasonicUnoccupiedToOccupiedDelay, O",
		Tests: []conformanceTest{
			{Context: matter.ConformanceContext{Values: map[string]any{"AA": true}}, Expected: matter.ConformanceStateOptional},
			{Context: matter.ConformanceContext{Values: map[string]any{"UltrasonicUnoccupiedToOccupiedDelay": true}}, Expected: matter.ConformanceStateMandatory},
			{Context: matter.ConformanceContext{Values: map[string]any{"UltrasonicUnoccupiedToOccupiedDelay": false}}, Expected: matter.ConformanceStateOptional},
			{Context: matter.ConformanceContext{Values: map[string]any{"Matter": true}}, Expected: matter.ConformanceStateOptional},
		},
	},
	{
		Conformance: "PIRUnoccupiedToOccupiedThreshold, O",
		Tests: []conformanceTest{
			{Context: matter.ConformanceContext{Values: map[string]any{"AA": true}}, Expected: matter.ConformanceStateOptional},
			{Context: matter.ConformanceContext{Values: map[string]any{"PIRUnoccupiedToOccupiedThreshold": true}}, Expected: matter.ConformanceStateMandatory},
			{Context: matter.ConformanceContext{Values: map[string]any{"UltrasonicUnoccupiedToOccupiedDelay": false}}, Expected: matter.ConformanceStateOptional},
			{Context: matter.ConformanceContext{Values: map[string]any{"Matter": true}}, Expected: matter.ConformanceStateOptional},
		},
	},
}

func TestOtherwise(t *testing.T) {
	for _, test := range otherwiseTests {
		test.run(t)
		break
	}
}
