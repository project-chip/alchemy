package conformance

import (
	"testing"

	"github.com/hasty/alchemy/matter/types"
)

func TestOptional(t *testing.T) {
	conformance, err := tryParseConformance("[!AB & (CD != EF)], O")
	if err != nil {

		t.Errorf("failed parsing: %v", err)
	}
	t.Logf("conformance: %s", conformance.String())
}

type conformanceTestSuite struct {
	Conformance        string
	InvalidConformance bool

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
			t.Errorf("failed checking conformance %s (parsed %s) with %v: expected %v, got %v", cts.Conformance, conformance.String(), test.Context, test.Expected, result)
		}
	}
}

type conformanceTest struct {
	Context  Context
	Expected State
}

type referenceStore struct {
	references map[string]types.Entity
}

var otherwiseTests = []conformanceTestSuite{
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
}

func TestOtherwise(t *testing.T) {
	for _, test := range otherwiseTests {
		test.run(t)
	}
}
