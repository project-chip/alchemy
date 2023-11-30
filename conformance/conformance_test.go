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
			if mc, ok := c.(*MandatoryConformance); ok {
				t.Logf("\t\ttesting %s: %T %v", cts.Conformance, mc.Expression, IsZigbee(conformance))
			}
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
		Conformance: "AB, [CD]",
		Tests: []conformanceTest{
			{Context: matter.ConformanceContext{"AB": true}, Expected: matter.ConformanceStateMandatory},
			{Context: matter.ConformanceContext{"CD": true}, Expected: matter.ConformanceStateOptional},
			{Context: matter.ConformanceContext{"Matter": true}, Expected: matter.ConformanceStateDisallowed},
		},
	},
	{
		Conformance: "!AB, O",
		Tests: []conformanceTest{
			{Context: matter.ConformanceContext{"AB": true}, Expected: matter.ConformanceStateOptional},
			{Context: matter.ConformanceContext{"CD": true}, Expected: matter.ConformanceStateMandatory},
			{Context: matter.ConformanceContext{"Matter": true}, Expected: matter.ConformanceStateMandatory},
		},
	},
	{
		Conformance: "[AA], [BB], [CC]",
		Tests: []conformanceTest{
			{Context: matter.ConformanceContext{"AA": true}, Expected: matter.ConformanceStateOptional},
			{Context: matter.ConformanceContext{"BB": true}, Expected: matter.ConformanceStateOptional},
			{Context: matter.ConformanceContext{"CC": true}, Expected: matter.ConformanceStateOptional},
			{Context: matter.ConformanceContext{"Matter": true}, Expected: matter.ConformanceStateDisallowed},
		},
	},
	{
		Conformance: "[!(LT | DF)]",
		Tests: []conformanceTest{
			{Context: matter.ConformanceContext{"AA": true}, Expected: matter.ConformanceStateOptional},
			{Context: matter.ConformanceContext{"LT": true}, Expected: matter.ConformanceStateDisallowed},
			{Context: matter.ConformanceContext{"DF": true}, Expected: matter.ConformanceStateDisallowed},
			{Context: matter.ConformanceContext{"Matter": true}, Expected: matter.ConformanceStateOptional},
		},
	},
	{
		Conformance: "UltrasonicUnoccupiedToOccupiedThreshold, O",
		Tests: []conformanceTest{
			{Context: matter.ConformanceContext{"AA": true}, Expected: matter.ConformanceStateOptional},
			{Context: matter.ConformanceContext{"UltrasonicUnoccupiedToOccupiedThreshold": true}, Expected: matter.ConformanceStateMandatory},
			{Context: matter.ConformanceContext{"UltrasonicUnoccupiedToOccupiedThreshold": false}, Expected: matter.ConformanceStateOptional},
			{Context: matter.ConformanceContext{"Matter": true}, Expected: matter.ConformanceStateOptional},
		},
	},
	{
		Conformance: "Zigbee",
		Tests: []conformanceTest{
			{Context: matter.ConformanceContext{"AA": true}, Expected: matter.ConformanceStateDisallowed},
			{Context: matter.ConformanceContext{"Zigbee": true}, Expected: matter.ConformanceStateMandatory},
			{Context: matter.ConformanceContext{"Zigbee": false}, Expected: matter.ConformanceStateDisallowed},
			{Context: matter.ConformanceContext{"Matter": true}, Expected: matter.ConformanceStateDisallowed},
		},
	},
	{
		Conformance: "[Zigbee]",
		Tests: []conformanceTest{
			{Context: matter.ConformanceContext{"AA": true}, Expected: matter.ConformanceStateDisallowed},
			{Context: matter.ConformanceContext{"Zigbee": true}, Expected: matter.ConformanceStateOptional},
			{Context: matter.ConformanceContext{"Zigbee": false}, Expected: matter.ConformanceStateDisallowed},
			{Context: matter.ConformanceContext{"Matter": true}, Expected: matter.ConformanceStateDisallowed},
		},
	},
	{
		Conformance: "MSCH",
		Tests: []conformanceTest{
			{Context: matter.ConformanceContext{"AA": true}, Expected: matter.ConformanceStateDisallowed},
			{Context: matter.ConformanceContext{"MSCH": true}, Expected: matter.ConformanceStateMandatory},
			{Context: matter.ConformanceContext{"MSCH": false}, Expected: matter.ConformanceStateDisallowed},
			{Context: matter.ConformanceContext{"Matter": true}, Expected: matter.ConformanceStateDisallowed},
		},
	},
	{
		Conformance: "M",
		Tests: []conformanceTest{
			{Context: matter.ConformanceContext{"AA": true}, Expected: matter.ConformanceStateMandatory},
			{Context: matter.ConformanceContext{"MSCH": true}, Expected: matter.ConformanceStateMandatory},
			{Context: matter.ConformanceContext{"MSCH": false}, Expected: matter.ConformanceStateMandatory},
			{Context: matter.ConformanceContext{"Matter": true}, Expected: matter.ConformanceStateMandatory},
		},
	},
	{
		Conformance: "(VIS | AUD) & SPRS",
		Tests: []conformanceTest{
			{Context: matter.ConformanceContext{"VIS": true}, Expected: matter.ConformanceStateMandatory},
			{Context: matter.ConformanceContext{"AUD": true}, Expected: matter.ConformanceStateMandatory},
			{Context: matter.ConformanceContext{"VIS": true, "AUD": true}, Expected: matter.ConformanceStateMandatory},
			{Context: matter.ConformanceContext{"Matter": true}, Expected: matter.ConformanceStateMandatory},
		},
	},
}

func TestOtherwise(t *testing.T) {
	for _, test := range otherwiseTests {
		test.run(t)
	}
}
