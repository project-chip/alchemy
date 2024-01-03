package constraint

import (
	"testing"

	"github.com/hasty/alchemy/matter/types"
)

type field struct {
	Name       string
	Constraint Constraint
	Type       *types.DataType
}

type fieldSet []*field

type constraintTestContext struct {
	field  *field
	fields fieldSet
}

func (cc *constraintTestContext) DataType() *types.DataType {
	if cc.field != nil {
		return cc.field.Type
	}
	return nil
}

func (cc *constraintTestContext) getReference(ref string) *field {
	for _, f := range cc.fields {
		if f.Name == ref {
			return f
		}
	}
	return nil
}

func (cc *constraintTestContext) ReferenceConstraint(ref string) Constraint {
	f := cc.getReference(ref)
	if f == nil {
		return nil
	}
	return f.Constraint
}

func (cc *constraintTestContext) Default(name string) (def types.DataTypeExtreme) {

	return
}

type constraintTest struct {
	constraint string
	dataType   *types.DataType
	min        types.DataTypeExtreme
	max        types.DataTypeExtreme
	asciiDoc   string
	zapMin     string
	zapMax     string
	fields     fieldSet
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
		min:        types.NewIntDataTypeExtreme(-4611686018427387904, types.NumberFormatHex),
		max:        types.NewIntDataTypeExtreme(4611686018427387904, types.NumberFormatHex),
		zapMin:     "0xC000000000000000",
		zapMax:     "0x4000000000000000",
	},
	{
		constraint: "0, MinMeasuredValue to MaxMeasuredValue",
		fields: fieldSet{
			{Name: "MinMeasuredValue", Constraint: ParseString("1 to MaxMeasuredValue-1")},
			{Name: "MaxMeasuredValue", Constraint: ParseString("MinMeasuredValue+1 to 65534")},
		},
		min:    types.NewIntDataTypeExtreme(0, types.NumberFormatInt),
		max:    types.NewIntDataTypeExtreme(65534, types.NumberFormatInt),
		zapMin: "0",
		zapMax: "65534",
	},
	{
		constraint: "1 to MaxMeasuredValue-1",
		fields: fieldSet{
			{Name: "MinMeasuredValue", Constraint: ParseString("1 to MaxMeasuredValue-1")},
			{Name: "MaxMeasuredValue", Constraint: ParseString("MinMeasuredValue+1 to 65534")},
		},
		min:      types.NewIntDataTypeExtreme(1, types.NumberFormatInt),
		max:      types.NewIntDataTypeExtreme(65533, types.NumberFormatInt),
		asciiDoc: "1 to (MaxMeasuredValue - 1)",
		zapMin:   "1",
		zapMax:   "65533",
	},
	{
		constraint: "MinMeasuredValue+1 to 65534",
		fields: fieldSet{
			{Name: "MinMeasuredValue", Constraint: ParseString("1 to MaxMeasuredValue-1")},
			{Name: "MaxMeasuredValue", Constraint: ParseString("MinMeasuredValue+1 to 65534")},
		},
		min:      types.NewIntDataTypeExtreme(2, types.NumberFormatInt),
		max:      types.NewIntDataTypeExtreme(65534, types.NumberFormatInt),
		asciiDoc: "(MinMeasuredValue + 1) to 65534",
		zapMin:   "2",
		zapMax:   "65534",
	},
	{
		constraint: "-2^62 to 2^62",
		asciiDoc:   "-2^62^ to 2^62^",
		min:        types.NewIntDataTypeExtreme(-4611686018427387904, types.NumberFormatHex),
		max:        types.NewIntDataTypeExtreme(4611686018427387904, types.NumberFormatHex),
		zapMin:     "0xC000000000000000",
		zapMax:     "0x4000000000000000",
	},

	{
		constraint: "max 2^62 - 1",
		asciiDoc:   "max (2^62^ - 1)",
		max:        types.NewIntDataTypeExtreme(4611686018427387903, types.NumberFormatAuto),
		zapMax:     "0x3FFFFFFFFFFFFFFF",
	},
	{
		constraint: "0 to 80000",
		min:        types.NewIntDataTypeExtreme(0, types.NumberFormatInt),
		max:        types.NewIntDataTypeExtreme(80000, types.NumberFormatInt),
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
		min:        types.NewUintDataTypeExtreme(0, types.NumberFormatHex),
		max:        types.NewUintDataTypeExtreme(60, types.NumberFormatHex),
		zapMin:     "0x0",
		zapMax:     "0x3C",
	},
	{
		constraint: "-32767 to MaxScaledValue-1",
		asciiDoc:   "-32767 to (MaxScaledValue - 1)",
		min:        types.NewIntDataTypeExtreme(-32767, types.NumberFormatInt),
		zapMin:     "-32767",
	},
	{
		constraint: "MaxScaledValue-1",
		asciiDoc:   "(MaxScaledValue - 1)",
	},
	{
		constraint: "-10000 to +10000",
		asciiDoc:   "-10000 to 10000",
		min:        types.NewIntDataTypeExtreme(-10000, types.NumberFormatInt),
		max:        types.NewIntDataTypeExtreme(10000, types.NumberFormatInt),
		zapMin:     "-10000",
		zapMax:     "10000",
	},
	{
		constraint: "-127 to 127",
		min:        types.NewIntDataTypeExtreme(-127, types.NumberFormatInt),
		max:        types.NewIntDataTypeExtreme(127, types.NumberFormatInt),
		zapMin:     "-127",
		zapMax:     "127",
	},
	{
		constraint: "-2.5°C to 2.5°C",
		dataType:   &types.DataType{BaseType: types.BaseDataTypeTemperature},
		min:        types.NewIntDataTypeExtreme(-250, types.NumberFormatInt),
		max:        types.NewIntDataTypeExtreme(250, types.NumberFormatInt),
		zapMin:     "-250",
		zapMax:     "250",
	},
	{
		constraint: "0 to 0x001F",
		dataType:   &types.DataType{BaseType: types.BaseDataTypeMap16},
		asciiDoc:   "0 to 0x001F",
		min:        types.NewIntDataTypeExtreme(0, types.NumberFormatInt),
		max:        types.NewUintDataTypeExtreme(31, types.NumberFormatHex),
		zapMin:     "0",
		zapMax:     "0x001F",
	},
	{
		constraint: "0 to 0xFEFF",
		min:        types.NewIntDataTypeExtreme(0, types.NumberFormatInt),
		max:        types.NewUintDataTypeExtreme(65279, types.NumberFormatHex),
		zapMin:     "0",
		zapMax:     "0xFEFF",
	},
	{
		constraint: "0 to 1000000",
		min:        types.NewIntDataTypeExtreme(0, types.NumberFormatInt),
		max:        types.NewIntDataTypeExtreme(1000000, types.NumberFormatInt),
		zapMin:     "0",
		zapMax:     "1000000",
	},
	{
		constraint: "0 to MaxFrequency",
		min:        types.NewIntDataTypeExtreme(0, types.NumberFormatInt),
		zapMin:     "0",
	},
	{
		constraint: "0% to 100%",
		min:        types.NewIntDataTypeExtreme(0, types.NumberFormatInt),
		max:        types.NewIntDataTypeExtreme(100, types.NumberFormatInt),
		zapMin:     "0",
		zapMax:     "100",
	},
	{
		constraint: "0% to 100%",
		dataType:   &types.DataType{BaseType: types.BaseDataTypePercentHundredths},
		min:        types.NewIntDataTypeExtreme(0, types.NumberFormatInt),
		max:        types.NewIntDataTypeExtreme(100, types.NumberFormatInt),
		zapMin:     "0",
		zapMax:     "10000",
	},

	{
		constraint: "0x954D to 0x7FFF",
		dataType:   &types.DataType{BaseType: types.BaseDataTypeTemperature},
		asciiDoc:   "0x954D to 0x7FFF",
		min:        types.NewUintDataTypeExtreme(38221, types.NumberFormatHex),
		max:        types.NewUintDataTypeExtreme(32767, types.NumberFormatHex),
		zapMin:     "0x954D",
		zapMax:     "0x7FFF",
	},
	{
		constraint: "0°C to 2.5°C",
		dataType:   &types.DataType{BaseType: types.BaseDataTypeTemperature},
		asciiDoc:   "0°C to 2.5°C",
		min:        types.NewIntDataTypeExtreme(0, types.NumberFormatInt),
		max:        types.NewIntDataTypeExtreme(250, types.NumberFormatInt),
		zapMin:     "0",
		zapMax:     "250",
	},
	{
		constraint: "1 to 100",
		min:        types.NewIntDataTypeExtreme(1, types.NumberFormatInt),
		max:        types.NewIntDataTypeExtreme(100, types.NumberFormatInt),
		zapMin:     "1",
		zapMax:     "100",
	},
	{
		constraint: "1 to MaxLevel",
		min:        types.NewIntDataTypeExtreme(1, types.NumberFormatInt),
		zapMin:     "1",
	},
	{
		constraint: "1 to MaxMeasuredValue-1",
		asciiDoc:   "1 to (MaxMeasuredValue - 1)",
		min:        types.NewIntDataTypeExtreme(1, types.NumberFormatInt),
		zapMin:     "1",
	},
	{
		constraint: "100 to MS",
		min:        types.NewIntDataTypeExtreme(100, types.NumberFormatInt),
		zapMin:     "100",
	},
	{
		constraint: "16",
		asciiDoc:   "16",
		min:        types.NewIntDataTypeExtreme(16, types.NumberFormatInt),
		max:        types.NewIntDataTypeExtreme(16, types.NumberFormatInt),
		zapMin:     "16",
		zapMax:     "16",
	},
	{
		constraint: "16[2]",
		asciiDoc:   "16[2]",
		min:        types.NewIntDataTypeExtreme(16, types.NumberFormatInt),
		max:        types.NewIntDataTypeExtreme(16, types.NumberFormatInt),
		zapMin:     "16",
		zapMax:     "16",
	},
	{
		constraint: "InstalledOpenLimitLift to InstalledClosedLimitLift",
	},
	{
		constraint: "MinMeasuredValue+1 to 10000",
		asciiDoc:   "(MinMeasuredValue + 1) to 10000",
		max:        types.NewIntDataTypeExtreme(10000, types.NumberFormatInt),
		zapMax:     "10000",
	},
	{
		constraint: "MinPower to 100",
		max:        types.NewIntDataTypeExtreme(100, types.NumberFormatInt),
		zapMax:     "100",
	},
	{
		constraint: "OccupiedEnabled, OccupiedDisabled",
	},
	{
		constraint: "OccupiedSetbackMin to 25.4°C",
		dataType:   &types.DataType{BaseType: types.BaseDataTypeTemperature},
		max:        types.NewIntDataTypeExtreme(2540, types.NumberFormatInt),
		zapMax:     "2540",
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
		max:        types.NewUintDataTypeExtreme(65534, types.NumberFormatHex),
		zapMax:     "0xFFFE",
	},
	{
		constraint: "max 10",
		max:        types.NewIntDataTypeExtreme(10, types.NumberFormatInt),
		zapMax:     "10",
	},
	{
		constraint: "max 10 [max 50]",
		asciiDoc:   "max 10[max 50]",
		max:        types.NewIntDataTypeExtreme(10, types.NumberFormatInt),
		zapMax:     "10",
	},
	{
		constraint: "max 32 chars",
		asciiDoc:   "max 32",
		max:        types.NewIntDataTypeExtreme(32, types.NumberFormatInt),
		zapMax:     "32",
	},
	{
		constraint: "max 604800",
		max:        types.NewIntDataTypeExtreme(604800, types.NumberFormatInt),
		zapMax:     "604800",
	},
	{
		constraint: "max NumberOfPositions-1",
		asciiDoc:   "max (NumberOfPositions - 1)",
	},
	{
		constraint: "min -27315",
		min:        types.NewIntDataTypeExtreme(-27315, types.NumberFormatInt),
		zapMin:     "-27315",
	},
	{
		constraint: "Min -27315",
		asciiDoc:   "min -27315",
		min:        types.NewIntDataTypeExtreme(-27315, types.NumberFormatInt),
		zapMin:     "-27315",
	},
	{
		constraint: "min 0",
		min:        types.NewIntDataTypeExtreme(0, types.NumberFormatInt),
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
		min:        types.DataTypeExtreme{Type: types.DataTypeExtremeTypeNull, Format: types.NumberFormatInt},
		max:        types.DataTypeExtreme{Type: types.DataTypeExtremeTypeNull, Format: types.NumberFormatInt},
	},
}

func TestSuite(t *testing.T) {
	for _, ct := range constraintTests {
		c := ParseString(ct.constraint)
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
		minField := &field{}
		minField.Type = ct.dataType
		min := c.Min(&constraintTestContext{fields: ct.fields, field: minField})
		if min != ct.min {
			t.Errorf("incorrect min value for \"%s\": expected %d, got %d", ct.constraint, ct.min, min)
		}
		maxField := &field{}
		maxField.Type = ct.dataType
		max := c.Max(&constraintTestContext{fields: ct.fields, field: maxField})
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
