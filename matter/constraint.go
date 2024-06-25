package matter

import (
	"github.com/project-chip/alchemy/matter/constraint"
	"github.com/project-chip/alchemy/matter/types"
)

type ConstraintContext struct {
	Field  *Field
	Fields FieldSet

	visitedReferences map[string]struct{}
}

func (cc *ConstraintContext) DataType() *types.DataType {
	if cc.Field != nil {
		return cc.Field.Type
	}
	return nil
}

func (cc *ConstraintContext) getReference(ref string) *Field {
	r := cc.Fields.GetField(ref)
	if cc.visitedReferences == nil {
		cc.visitedReferences = make(map[string]struct{})
	} else if _, ok := cc.visitedReferences[ref]; ok {
		return nil
	}
	cc.visitedReferences[ref] = struct{}{}
	return r
}

func (cc *ConstraintContext) ReferenceConstraint(ref string) constraint.Constraint {
	f := cc.getReference(ref)
	if f == nil {
		return nil
	}
	return f.Constraint
}

func (cc *ConstraintContext) Default(name string) (def types.DataTypeExtreme) {
	f := cc.getReference(name)
	if f != nil && f.Default != "" {
		cons, err := constraint.ParseString(f.Default)
		if err != nil {
			cons = &constraint.GenericConstraint{Value: f.Default}
		}
		return cons.Default(cc)
	}
	// Couldn't find it as a reference; let's check other possibilities
	def = cc.getEnumValue(name)
	if def.Defined() {
		return
	}
	def = cc.getBitmapValue(name)
	if def.Defined() {
		return
	}
	return
}

func (cc *ConstraintContext) getEnumValue(name string) (def types.DataTypeExtreme) {
	if cc.Field.Type == nil || cc.Field.Type.Entity == nil {
		return
	}
	en, ok := cc.Field.Type.Entity.(*Enum)
	if !ok {
		return
	}
	for _, v := range en.Values {
		if v.Name == name {
			if v.Value.Valid() {
				def = types.NewUintDataTypeExtreme(v.Value.Value(), types.NumberFormatInt)
			}
		}
	}
	return
}

func (cc *ConstraintContext) getBitmapValue(name string) (def types.DataTypeExtreme) {
	if cc.Field.Type == nil || cc.Field.Type.Entity == nil {
		return
	}
	en, ok := cc.Field.Type.Entity.(*Bitmap)
	if !ok {
		return
	}
	for _, v := range en.Bits {
		if v.Name() == name {
			val, err := v.Mask()
			if err == nil {
				def = types.NewUintDataTypeExtreme(val, types.NumberFormatHex)
				return
			}
		}
	}
	return
}
