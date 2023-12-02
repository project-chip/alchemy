package constraint

import (
	"testing"

	"github.com/hasty/alchemy/matter"
)

type constraintTest struct {
	constraint string
	dataType   *matter.DataType
	min        matter.ConstraintExtreme
	max        matter.ConstraintExtreme
	asciiDoc   string
	zapMin     string
	zapMax     string
	fields     matter.FieldSet
	generic    bool
}

var constraintTests = []constraintTest{
	{
		constraint: "00000xxx",
		generic:    true,
	},
	{
		constraint: "0b0000 xxxx",
		generic:    true,
	},

	{
		constraint: "-2^62^ to 2^62^",
		min:        matter.NewIntConstraintExtreme(-4611686018427387904, matter.NumberFormatHex),
		max:        matter.NewIntConstraintExtreme(4611686018427387904, matter.NumberFormatHex),
		zapMin:     "0xC000000000000000",
		zapMax:     "0x4000000000000000",
	},
	{
		constraint: "0, MinMeasuredValue to MaxMeasuredValue",
		fields: matter.FieldSet{
			{Name: "MinMeasuredValue", Constraint: ParseConstraint("1 to MaxMeasuredValue-1")},
			{Name: "MaxMeasuredValue", Constraint: ParseConstraint("MinMeasuredValue+1 to 65534")},
		},
		min:    matter.NewIntConstraintExtreme(0, matter.NumberFormatInt),
		max:    matter.NewIntConstraintExtreme(65534, matter.NumberFormatInt),
		zapMin: "0",
		zapMax: "65534",
	},
	{
		constraint: "1 to MaxMeasuredValue-1",
		fields: matter.FieldSet{
			{Name: "MinMeasuredValue", Constraint: ParseConstraint("1 to MaxMeasuredValue-1")},
			{Name: "MaxMeasuredValue", Constraint: ParseConstraint("MinMeasuredValue+1 to 65534")},
		},
		min:      matter.NewIntConstraintExtreme(1, matter.NumberFormatInt),
		max:      matter.NewIntConstraintExtreme(65533, matter.NumberFormatInt),
		asciiDoc: "1 to (MaxMeasuredValue - 1)",
		zapMin:   "1",
		zapMax:   "65533",
	},
	{
		constraint: "MinMeasuredValue+1 to 65534",
		fields: matter.FieldSet{
			{Name: "MinMeasuredValue", Constraint: ParseConstraint("1 to MaxMeasuredValue-1")},
			{Name: "MaxMeasuredValue", Constraint: ParseConstraint("MinMeasuredValue+1 to 65534")},
		},
		min:      matter.NewIntConstraintExtreme(2, matter.NumberFormatInt),
		max:      matter.NewIntConstraintExtreme(65534, matter.NumberFormatInt),
		asciiDoc: "(MinMeasuredValue + 1) to 65534",
		zapMin:   "2",
		zapMax:   "65534",
	},
	{
		constraint: "-2^62 to 2^62",
		asciiDoc:   "-2^62^ to 2^62^",
		min:        matter.NewIntConstraintExtreme(-4611686018427387904, matter.NumberFormatHex),
		max:        matter.NewIntConstraintExtreme(4611686018427387904, matter.NumberFormatHex),
		zapMin:     "0xC000000000000000",
		zapMax:     "0x4000000000000000",
	},

	{
		constraint: "max 2^62 - 1",
		asciiDoc:   "max (2^62^ - 1)",
		max:        matter.NewIntConstraintExtreme(4611686018427387903, matter.NumberFormatAuto),
		zapMax:     "0x3FFFFFFFFFFFFFFF",
	},
	{
		constraint: "0 to 80000",
		min:        matter.NewIntConstraintExtreme(0, matter.NumberFormatInt),
		max:        matter.NewIntConstraintExtreme(80000, matter.NumberFormatInt),
		zapMin:     "0",
		zapMax:     "80000",
	},
	{
		constraint: "max (NumberOfEventsPerProgram * (1 + NumberOfLoadControlPrograms))",
	},
	{
		constraint: "InstalledOpenLimitLift to InstalledClosedLimitLift",
	},
	{
		constraint: "0x00 to 0x3C",
		asciiDoc:   "0x0 to 0x3C",
		min:        matter.NewUintConstraintExtreme(0, matter.NumberFormatHex),
		max:        matter.NewUintConstraintExtreme(60, matter.NumberFormatHex),
		zapMin:     "0x0",
		zapMax:     "0x3C",
	},
	{
		constraint: "-32767 to MaxScaledValue-1",
		asciiDoc:   "-32767 to (MaxScaledValue - 1)",
		min:        matter.NewIntConstraintExtreme(-32767, matter.NumberFormatInt),
		zapMin:     "-32767",
	},
	{
		constraint: "MaxScaledValue-1",
		asciiDoc:   "(MaxScaledValue - 1)",
	},
	{
		constraint: "-10000 to +10000",
		asciiDoc:   "-10000 to 10000",
		min:        matter.NewIntConstraintExtreme(-10000, matter.NumberFormatInt),
		max:        matter.NewIntConstraintExtreme(10000, matter.NumberFormatInt),
		zapMin:     "-10000",
		zapMax:     "10000",
	},
	{
		constraint: "-127 to 127",
		min:        matter.NewIntConstraintExtreme(-127, matter.NumberFormatInt),
		max:        matter.NewIntConstraintExtreme(127, matter.NumberFormatInt),
		zapMin:     "-127",
		zapMax:     "127",
	},
	{
		constraint: "-2.5°C to 2.5°C",
		dataType:   &matter.DataType{BaseType: matter.BaseDataTypeTemperature},
		min:        matter.NewIntConstraintExtreme(-250, matter.NumberFormatHex),
		max:        matter.NewIntConstraintExtreme(250, matter.NumberFormatHex),
		zapMin:     "0xFF06",
		zapMax:     "0x00FA",
	},
	{
		constraint: "0 to 0x001F",
		dataType:   &matter.DataType{BaseType: matter.BaseDataTypeMap16},
		asciiDoc:   "0 to 0x001F",
		min:        matter.NewIntConstraintExtreme(0, matter.NumberFormatInt),
		max:        matter.NewUintConstraintExtreme(31, matter.NumberFormatHex),
		zapMin:     "0",
		zapMax:     "0x001F",
	},
	{
		constraint: "0 to 0xFEFF",
		min:        matter.NewIntConstraintExtreme(0, matter.NumberFormatInt),
		max:        matter.NewUintConstraintExtreme(65279, matter.NumberFormatHex),
		zapMin:     "0",
		zapMax:     "0xFEFF",
	},
	{
		constraint: "0 to 1000000",
		min:        matter.NewIntConstraintExtreme(0, matter.NumberFormatInt),
		max:        matter.NewIntConstraintExtreme(1000000, matter.NumberFormatInt),
		zapMin:     "0",
		zapMax:     "1000000",
	},
	{
		constraint: "0 to MaxFrequency",
		min:        matter.NewIntConstraintExtreme(0, matter.NumberFormatInt),
		zapMin:     "0",
	},
	{
		constraint: "0% to 100%",
		min:        matter.NewIntConstraintExtreme(0, matter.NumberFormatInt),
		max:        matter.NewIntConstraintExtreme(100, matter.NumberFormatInt),
		zapMin:     "0",
		zapMax:     "100",
	},
	{
		constraint: "0% to 100%",
		dataType:   &matter.DataType{BaseType: matter.BaseDataTypePercentHundredths},
		min:        matter.NewIntConstraintExtreme(0, matter.NumberFormatInt),
		max:        matter.NewIntConstraintExtreme(100, matter.NumberFormatInt),
		zapMin:     "0",
		zapMax:     "10000",
	},

	{
		constraint: "0x954D to 0x7FFF",
		dataType:   &matter.DataType{BaseType: matter.BaseDataTypeTemperature},
		asciiDoc:   "0x954D to 0x7FFF",
		min:        matter.NewUintConstraintExtreme(38221, matter.NumberFormatHex),
		max:        matter.NewUintConstraintExtreme(32767, matter.NumberFormatHex),
		zapMin:     "0x954D",
		zapMax:     "0x7FFF",
	},
	{
		constraint: "0°C to 2.5°C",
		dataType:   &matter.DataType{BaseType: matter.BaseDataTypeTemperature},
		asciiDoc:   "0°C to 2.5°C",
		min:        matter.NewIntConstraintExtreme(0, matter.NumberFormatHex),
		max:        matter.NewIntConstraintExtreme(250, matter.NumberFormatHex),
		zapMin:     "0x0000",
		zapMax:     "0x00FA",
	},
	{
		constraint: "1 to 100",
		min:        matter.NewIntConstraintExtreme(1, matter.NumberFormatInt),
		max:        matter.NewIntConstraintExtreme(100, matter.NumberFormatInt),
		zapMin:     "1",
		zapMax:     "100",
	},
	{
		constraint: "1 to MaxLevel",
		min:        matter.NewIntConstraintExtreme(1, matter.NumberFormatInt),
		zapMin:     "1",
	},
	{
		constraint: "1 to MaxMeasuredValue-1",
		asciiDoc:   "1 to (MaxMeasuredValue - 1)",
		min:        matter.NewIntConstraintExtreme(1, matter.NumberFormatInt),
		zapMin:     "1",
	},
	{
		constraint: "100 to MS",
		min:        matter.NewIntConstraintExtreme(100, matter.NumberFormatInt),
		zapMin:     "100",
	},
	{
		constraint: "16",
		asciiDoc:   "16",
		min:        matter.NewIntConstraintExtreme(16, matter.NumberFormatInt),
		max:        matter.NewIntConstraintExtreme(16, matter.NumberFormatInt),
		zapMin:     "16",
		zapMax:     "16",
	},
	{
		constraint: "16[2]",
		asciiDoc:   "16[2]",
		min:        matter.NewIntConstraintExtreme(16, matter.NumberFormatInt),
		max:        matter.NewIntConstraintExtreme(16, matter.NumberFormatInt),
		zapMin:     "16",
		zapMax:     "16",
	},
	{
		constraint: "InstalledOpenLimitLift to InstalledClosedLimitLift",
	},
	{
		constraint: "MinMeasuredValue+1 to 10000",
		asciiDoc:   "(MinMeasuredValue + 1) to 10000",
		max:        matter.NewIntConstraintExtreme(10000, matter.NumberFormatInt),
		zapMax:     "10000",
	},
	{
		constraint: "MinPower to 100",
		max:        matter.NewIntConstraintExtreme(100, matter.NumberFormatInt),
		zapMax:     "100",
	},
	{
		constraint: "OccupiedEnabled, OccupiedDisabled",
	},
	{
		constraint: "OccupiedSetbackMin to 25.4°C",
		dataType:   &matter.DataType{BaseType: matter.BaseDataTypeTemperature},
		max:        matter.NewIntConstraintExtreme(2540, matter.NumberFormatHex),
		zapMax:     "0x09EC",
	},
	{
		constraint: "TODO",
		generic:    true,
	},
	{
		constraint: "all[min 1]",
	},
	{
		constraint: "any",
	},
	{
		constraint: "max MaxTemperature - 1",
		asciiDoc:   "max (MaxTemperature - 1)",
	},
	{
		constraint: "max MaxTemperature - MinTemperature",
		asciiDoc:   "max (MaxTemperature - MinTemperature)",
	},
	{
		constraint: "max 0xFFFE",
		max:        matter.NewUintConstraintExtreme(65534, matter.NumberFormatHex),
		zapMax:     "0xFFFE",
	},
	{
		constraint: "max 10",
		max:        matter.NewIntConstraintExtreme(10, matter.NumberFormatInt),
		zapMax:     "10",
	},
	{
		constraint: "max 10 [max 50]",
		asciiDoc:   "max 10[max 50]",
		max:        matter.NewIntConstraintExtreme(10, matter.NumberFormatInt),
		zapMax:     "10",
	},
	{
		constraint: "max 32 chars",
		asciiDoc:   "max 32",
		max:        matter.NewIntConstraintExtreme(32, matter.NumberFormatInt),
		zapMax:     "32",
	},
	{
		constraint: "max 604800",
		max:        matter.NewIntConstraintExtreme(604800, matter.NumberFormatInt),
		zapMax:     "604800",
	},
	{
		constraint: "max NumberOfPositions-1",
		asciiDoc:   "max (NumberOfPositions - 1)",
	},
	{
		constraint: "min -27315",
		min:        matter.NewIntConstraintExtreme(-27315, matter.NumberFormatInt),
		zapMin:     "-27315",
	},
	{
		constraint: "Min -27315",
		asciiDoc:   "min -27315",
		min:        matter.NewIntConstraintExtreme(-27315, matter.NumberFormatInt),
		zapMin:     "-27315",
	},
	{
		constraint: "min 0",
		min:        matter.NewIntConstraintExtreme(0, matter.NumberFormatInt),
		zapMin:     "0",
	},
	{
		constraint: "max MinFrequency",
	},
	{
		constraint: "percent",
		generic:    true,
	},
	{
		constraint: "null",
		min:        matter.ConstraintExtreme{},
		max:        matter.ConstraintExtreme{},
	},
}

func TestSuite(t *testing.T) {
	for _, ct := range constraintTests {
		c := ParseConstraint(ct.constraint)
		_, isGeneric := c.(*GenericConstraint)
		if ct.generic {
			if !isGeneric {
				t.Errorf("expected generic constraint for %s, got %T", ct.constraint, c)
			}
			continue
		} else if isGeneric {
			t.Errorf("failed to parse constraint %s", ct.constraint)
			continue
		}
		min := c.Min(&matter.ConstraintContext{Fields: ct.fields, Field: &matter.Field{Type: ct.dataType}})
		if min != ct.min {
			t.Errorf("incorrect min value for \"%s\": expected %d, got %d", ct.constraint, ct.min, min)
		}
		max := c.Max(&matter.ConstraintContext{Fields: ct.fields, Field: &matter.Field{Type: ct.dataType}})
		if max != ct.max {
			t.Errorf("incorrect max value for \"%s\": expected %d, got %d", ct.constraint, ct.max, max)
		}
		as := c.AsciiDocString(ct.dataType)
		es := ct.constraint
		if len(ct.asciiDoc) > 0 {
			es = ct.asciiDoc
		}
		if as != es {
			t.Errorf("incorrect AsciiDoc value for \"%s\": expected %s, got %s", ct.constraint, es, as)
		}

		if min.ZapString(ct.dataType) != ct.zapMin {
			t.Errorf("incorrect ZAP min value for \"%s\": expected %s, got %s", ct.constraint, ct.zapMin, min.ZapString(ct.dataType))

		}
		if max.ZapString(ct.dataType) != ct.zapMax {
			t.Errorf("incorrect ZAP max value for \"%s\": expected %s, got %s", ct.constraint, ct.zapMax, max.ZapString(ct.dataType))
		}
	}

}
