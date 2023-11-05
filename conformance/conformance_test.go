package conformance

import (
	"testing"

	"github.com/hasty/matterfmt/matter"
)

func TestOptional(t *testing.T) {
	conformance, err := ParseConformance("[!AB & (CD != EF)], O")
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
	conformance, err := ParseConformance(cts.Conformance)
	if err != nil {
		t.Errorf("failed parsing conformance %s: %v", cts.Conformance, err)
		return
	}
	for _, test := range cts.Tests {
		result, err := conformance.Eval(test.Context)
		if err != nil {
			t.Errorf("failed evaluating conformance status %v: %v", test.Context, err)
			return
		}
		if result != test.Expected {
			t.Errorf("failed checking conformance %s with %v: expected %v, got %v", conformance.String(), test.Context, test.Expected, result)
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
}

func TestOtherwise(t *testing.T) {
	for _, test := range otherwiseTests {
		test.run(t)
	}
}
