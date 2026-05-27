package constraint

import (
	"fmt"
	"testing"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/matter/types"
)

type field struct {
	Name       string
	Constraint Constraint
	Type       *types.DataType
	Nullable   bool
}

func (f *field) EntityType() types.EntityType {
	return types.EntityTypeAttribute
}

func (f *field) Source() asciidoc.Element {
	return nil
}

func (f *field) Origin() (path string, line int) {
	return
}

func (f *field) Parent() types.Entity {
	return nil
}

func (f *field) Equals(e types.Entity) bool {
	oe, ok := e.(*field)
	if !ok {
		return false
	}
	return f.Name == oe.Name
}

type fieldSet []*field

func (fs fieldSet) getField(name string) *field {
	for _, f := range fs {
		if f.Name == name {
			return f
		}
	}
	return nil
}

type constraintTestContext struct {
	field  *field
	fields fieldSet
}

func (cc *constraintTestContext) child(field *field, fields fieldSet) *constraintTestContext {
	return &constraintTestContext{
		field:  field,
		fields: cc.fields,
	}
}

func (cc *constraintTestContext) Nullability() types.Nullability {
	if cc.field != nil && cc.field.Nullable {
		return types.NullabilityNullable
	}
	return types.NullabilityNonNull
}

func (cc *constraintTestContext) DataType() *types.DataType {
	if cc.field != nil {
		return cc.field.Type
	}
	return nil
}

func (cc *constraintTestContext) MinEntityValue(entity types.Entity, f Limit) (min types.DataTypeExtreme) {
	switch entity := entity.(type) {
	case *field:
		if entity != nil {
			min = entity.Constraint.Min(cc.child(entity, cc.fields))
		}
	}
	return
}

func (cc *constraintTestContext) MaxEntityValue(entity types.Entity, f Limit) (max types.DataTypeExtreme) {
	switch entity := entity.(type) {
	case *field:
		if entity != nil {
			max = entity.Constraint.Max(cc.child(entity, cc.fields))
		}
	}
	return
}

func (cc *constraintTestContext) Fallback(entity types.Entity, field Limit) (def types.DataTypeExtreme) {
	return
}

func mustParseConstraint(s string) Constraint {
	constraint, err := TryParseString(s)
	if err != nil {
		panic(fmt.Errorf("failed parsing constraint \"%s\": %w", s, err))
	}
	return constraint
}

func stitchFieldSet(fs fieldSet) fieldSet {
	for _, f := range fs {
		stitchFieldConstraint(fs, f.Constraint)
	}
	return fs
}

func stitchFieldConstraint(fs fieldSet, cons Constraint) {
	switch cons := cons.(type) {
	case Set:
		for _, c := range cons {
			stitchFieldConstraint(fs, c)
		}
	case *ExactConstraint:
		stitchFieldLimit(fs, cons.Value)
	case *ListConstraint:
		stitchFieldConstraint(fs, cons.Constraint)
		stitchFieldConstraint(fs, cons.EntryConstraint)
	case *MaxConstraint:
		stitchFieldLimit(fs, cons.Maximum)
	case *MinConstraint:
		stitchFieldLimit(fs, cons.Minimum)
	case *RangeConstraint:
		stitchFieldLimit(fs, cons.Maximum)
		stitchFieldLimit(fs, cons.Minimum)
	case *LogicalConstraint:
		stitchFieldConstraint(fs, cons.Left)
		for _, r := range cons.Right {
			stitchFieldConstraint(fs, r)
		}
	case *AllConstraint, *TagListConstraint:
	default:
		fmt.Printf("unknown field constraint type: %T\n", cons)
	}
}

func stitchFieldLimit(fs fieldSet, l Limit) {
	switch l := l.(type) {
	case *CharacterLimit:
		stitchFieldLimit(fs, l.ByteCount)
		stitchFieldLimit(fs, l.CodepointCount)
	case *LengthLimit:
		stitchFieldLimit(fs, l.Reference)
	case *ReferenceLimit:
		l.Entity = fs.getField(l.Reference)
	case *IdentifierLimit:
		l.Entity = fs.getField(l.ID)
	case *MathExpressionLimit:
		stitchFieldLimit(fs, l.Left)
		stitchFieldLimit(fs, l.Right)
	case *LogicalLimit:
		stitchFieldLimit(fs, l.Left)
		for _, r := range l.Right {
			stitchFieldLimit(fs, r)
		}
	case *BooleanLimit, *IntLimit, *ExpLimit, *HexLimit, *NullLimit, *PercentLimit, *TemperatureLimit, *ManufacturerLimit:
	case *MaxOfLimit:
		for _, l := range l.Maximums {
			stitchFieldLimit(fs, l)
		}
	case *MinOfLimit:
		for _, l := range l.Minimums {
			stitchFieldLimit(fs, l)
		}
	default:
		fmt.Printf("unknown field limit type: %T\n", l)
	}
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
	invalid    bool
}

var constraintTests = []constraintTest{
	{
		constraint: "0 to (FragmentDuration/2)",
		asciiDoc:   "0 to (FragmentDuration / 2)",
		fields: stitchFieldSet(fieldSet{
			{Name: "FragmentDuration", Type: &types.DataType{BaseType: types.BaseDataTypeUInt16}, Constraint: mustParseConstraint("all")},
		}),
		min:    types.NewIntDataTypeExtreme(0, types.NumberFormatInt),
		max:    types.NewUintDataTypeExtreme(32767, types.NumberFormatInt),
		zapMin: "0",
		zapMax: "32767",
	},
	{
		constraint: "HoldTimeMin & min 10",
		asciiDoc:   "(HoldTimeMin & min 10)",
		min:        types.DataTypeExtreme{Type: types.DataTypeExtremeTypeInt64, Int64: 10, Format: types.NumberFormatInt},
		max:        types.DataTypeExtreme{Type: types.DataTypeExtremeTypeUndefined},
		zapMin:     "10",
		fields: stitchFieldSet(fieldSet{
			{Name: "HoldTimeMin", Constraint: mustParseConstraint("min 1")},
		}),
	},
	{
		constraint: "Includes `Grid` and `Battery`",
		asciiDoc:   "Includes (`Grid` and `Battery`)",
		min:        types.DataTypeExtreme{Type: types.DataTypeExtremeTypeUndefined, Format: types.NumberFormatAuto},
		max:        types.DataTypeExtreme{Type: types.DataTypeExtremeTypeUndefined, Format: types.NumberFormatAuto},
	},
	{
		constraint: "1 to 3[20]",
		min:        types.NewIntDataTypeExtreme(1, types.NumberFormatInt),
		max:        types.NewIntDataTypeExtreme(3, types.NumberFormatInt),
		zapMin:     "1",
		zapMax:     "3",
	},
	{
		constraint: "Includes `Grid`",
		min:        types.DataTypeExtreme{Type: types.DataTypeExtremeTypeUndefined, Format: types.NumberFormatAuto},
		max:        types.DataTypeExtreme{Type: types.DataTypeExtremeTypeUndefined, Format: types.NumberFormatAuto},
	},
	{
		constraint: "kWh | kVAh",
		asciiDoc:   `kWh \| kVAh`,
		min:        types.DataTypeExtreme{Type: types.DataTypeExtremeTypeUndefined},
		max:        types.DataTypeExtreme{Type: types.DataTypeExtremeTypeUndefined},
	},
	{
		constraint: `kWh \| kVAh`,
		asciiDoc:   `kWh \| kVAh`,
		min:        types.DataTypeExtreme{Type: types.DataTypeExtremeTypeUndefined},
		max:        types.DataTypeExtreme{Type: types.DataTypeExtremeTypeUndefined},
	},
	{
		constraint: "min <<ref_NumberOfScheduleTransitions>>",
	},
	{
		constraint: "1 to <<ref_NumberOfScheduleTransitions>>",
		min:        types.NewIntDataTypeExtreme(1, types.NumberFormatInt),
		zapMin:     "1",
	},
	{
		constraint: "0, MinMeasuredValue to MaxMeasuredValue",
		fields: stitchFieldSet(fieldSet{
			{Name: "MinMeasuredValue", Constraint: mustParseConstraint("1 to MaxMeasuredValue-1")},
			{Name: "MaxMeasuredValue", Constraint: mustParseConstraint("MinMeasuredValue+1 to 65534")},
		}),
		min:    types.NewIntDataTypeExtreme(0, types.NumberFormatInt),
		max:    types.NewIntDataTypeExtreme(65534, types.NumberFormatInt),
		zapMin: "0",
		zapMax: "65534",
	},
	{
		constraint: "max 128{32}",
		asciiDoc:   "max 128{32}",
		dataType:   &types.DataType{BaseType: types.BaseDataTypeString},
		max:        types.NewIntDataTypeExtreme(128, types.NumberFormatInt),
		zapMax:     "128",
	},
	{
		constraint: "True",
		asciiDoc:   "true",
		dataType:   &types.DataType{BaseType: types.BaseDataTypeBoolean},
		min:        types.NewUintDataTypeExtreme(1, types.NumberFormatInt),
		max:        types.NewUintDataTypeExtreme(1, types.NumberFormatInt),
		zapMin:     "1",
		zapMax:     "1",
	},
	{
		constraint: "False",
		asciiDoc:   "false",
		dataType:   &types.DataType{BaseType: types.BaseDataTypeBoolean},
		min:        types.NewUintDataTypeExtreme(0, types.NumberFormatInt),
		max:        types.NewUintDataTypeExtreme(0, types.NumberFormatInt),
		zapMin:     "0",
		zapMax:     "0",
	},
	{
		constraint: "1 to <<ref_NumberOfScheduleTransitions>>",
		min:        types.NewIntDataTypeExtreme(1, types.NumberFormatInt),
		zapMin:     "1",
	},
	{
		constraint: "max <<ref_RespMaxConstant,RESP_MAX>>",
		asciiDoc:   "max <<ref_RespMaxConstant, RESP_MAX>>",
	},
	{
		constraint: "1 to NumberOfPINUsersSupported, 0xFFFE",
		min:        types.NewIntDataTypeExtreme(1, types.NumberFormatInt),
		max:        types.NewUintDataTypeExtreme(65534, types.NumberFormatHex),
		zapMin:     "1",
		zapMax:     "0xFFFE",
	},
	{
		constraint: "00000xxx",
		invalid:    true,
	},
	{
		constraint: "0b0000 xxxx",
		invalid:    true,
	},

	{
		constraint: "-2^62^ to 2^62^",
		min:        types.NewIntDataTypeExtreme(-4611686018427387904, types.NumberFormatInt),
		max:        types.NewIntDataTypeExtreme(4611686018427387904, types.NumberFormatInt),
		zapMin:     "-4611686018427387904",
		zapMax:     "4611686018427387904",
	},

	{
		constraint: "1 to MaxMeasuredValue-1",
		fields: stitchFieldSet(fieldSet{
			{Name: "MinMeasuredValue", Constraint: mustParseConstraint("1 to MaxMeasuredValue-1")},
			{Name: "MaxMeasuredValue", Constraint: mustParseConstraint("MinMeasuredValue+1 to 65534")},
		}),
		min:      types.NewIntDataTypeExtreme(1, types.NumberFormatInt),
		max:      types.NewUintDataTypeExtreme(65533, types.NumberFormatAuto),
		asciiDoc: "1 to (MaxMeasuredValue - 1)",
		zapMin:   "1",
		zapMax:   "65533",
	},
	{
		constraint: "MinMeasuredValue+1 to 65534",
		fields: stitchFieldSet(fieldSet{
			{Name: "MinMeasuredValue", Constraint: mustParseConstraint("1 to MaxMeasuredValue-1")},
			{Name: "MaxMeasuredValue", Constraint: mustParseConstraint("MinMeasuredValue+1 to 65534")},
		}),
		min:      types.NewUintDataTypeExtreme(2, types.NumberFormatInt),
		max:      types.NewIntDataTypeExtreme(65534, types.NumberFormatInt),
		asciiDoc: "(MinMeasuredValue + 1) to 65534",
		zapMin:   "2",
		zapMax:   "65534",
	},
	{
		constraint: "-2^62 to 2^62",
		asciiDoc:   "-2^62^ to 2^62^",
		min:        types.NewIntDataTypeExtreme(-4611686018427387904, types.NumberFormatInt),
		max:        types.NewIntDataTypeExtreme(4611686018427387904, types.NumberFormatInt),
		zapMin:     "-4611686018427387904",
		zapMax:     "4611686018427387904",
	},

	{
		constraint: "-2^62^ to 2^62^",
		min:        types.NewIntDataTypeExtreme(-4611686018427387904, types.NumberFormatInt),
		max:        types.NewIntDataTypeExtreme(4611686018427387904, types.NumberFormatInt),
		zapMin:     "-4611686018427387904",
		zapMax:     "4611686018427387904",
	},
	{
		constraint: "max 2^62 - 1",
		asciiDoc:   "max (2^62^ - 1)",
		max:        types.NewUintDataTypeExtreme(4611686018427387903, types.NumberFormatInt),
		zapMax:     "4611686018427387903",
	},
	{
		constraint: "0 to 2^62^",
		asciiDoc:   "0 to 2^62^",
		min:        types.NewIntDataTypeExtreme(0, types.NumberFormatInt),
		max:        types.NewIntDataTypeExtreme(4611686018427387904, types.NumberFormatInt),
		zapMin:     "0",
		zapMax:     "4611686018427387904",
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
		asciiDoc:   "MaxScaledValue - 1",
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
		min:        types.NewUintDataTypeExtreme(0, types.NumberFormatInt),
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
		invalid:    true,
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
		invalid:    true,
	},
	{
		constraint: "null",
		min:        types.DataTypeExtreme{Type: types.DataTypeExtremeTypeNull, Format: types.NumberFormatAuto},
		max:        types.DataTypeExtreme{Type: types.DataTypeExtremeTypeNull, Format: types.NumberFormatAuto},
	},
}

func TestConstraints(t *testing.T) {
	for _, ct := range constraintTests {
		if ct.invalid {
			c, err := TryParseString(ct.constraint)
			if err == nil {
				t.Errorf("expected error parsing %s; got %T", ct.constraint, c)
				if c, ok := c.(Set); ok {
					for _, c := range c {
						t.Errorf("expected error parsing %s; got %T: %s", ct.constraint, c, c.ASCIIDocString(ct.dataType))
					}
				}
			}
			continue
		}
		c := mustParseConstraint(ct.constraint)
		_, isGeneric := c.(*GenericConstraint)
		if isGeneric {
			t.Errorf("failed to parse constraint %s", ct.constraint)
			continue
		}
		stitchFieldConstraint(ct.fields, c)
		minField := &field{}
		minField.Type = ct.dataType
		min := c.Min(&constraintTestContext{fields: ct.fields, field: minField})
		if !min.Equals(ct.min) {
			//t.Errorf("incorrect min value for \"%s\": expected %s, got %s", ct.constraint, ct.min.DataModelString(ct.dataType), min.DataModelString(ct.dataType))
			t.Errorf("incorrect min value for \"%v\": expected %v, got %v", ct.constraint, ct.min, min)
		}
		maxField := &field{}
		maxField.Type = ct.dataType
		max := c.Max(&constraintTestContext{fields: ct.fields, field: maxField})
		if !max.Equals(ct.max) {
			t.Errorf("incorrect max value for \"%s\": expected %v, got %v", ct.constraint, ct.max, max)
		}
		as := c.ASCIIDocString(ct.dataType)
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
