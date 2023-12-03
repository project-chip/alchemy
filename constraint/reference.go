package constraint

import (
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/parse"
)

type ReferenceLimit struct {
	Value string
}

func (c *ReferenceLimit) AsciiDocString(dataType *matter.DataType) string {
	return c.Value
}

func (c *ReferenceLimit) Equal(o matter.ConstraintLimit) bool {
	if oc, ok := o.(*ReferenceLimit); ok {
		return oc.Value == c.Value
	}
	return false
}

func (c *ReferenceLimit) Min(cc *matter.ConstraintContext) (min matter.ConstraintExtreme) {
	r := c.getReference(cc)
	if r == nil || r.Constraint == nil {
		return
	}
	return r.Constraint.Min(cc)
}

func (c *ReferenceLimit) Max(cc *matter.ConstraintContext) (max matter.ConstraintExtreme) {
	r := c.getReference(cc)
	if r == nil || r.Constraint == nil {
		return
	}
	return r.Constraint.Max(cc)
}

func (c *ReferenceLimit) getReference(cc *matter.ConstraintContext) *matter.Field {
	r := cc.Fields.GetField(c.Value)
	if cc.VisitedReferences == nil {
		cc.VisitedReferences = make(map[string]struct{})
	} else if _, ok := cc.VisitedReferences[c.Value]; ok {
		return nil
	}
	cc.VisitedReferences[c.Value] = struct{}{}
	return r
}

func (c *ReferenceLimit) getEnumValue(cc *matter.ConstraintContext) (def matter.ConstraintExtreme) {
	if cc.Field.Type == nil || cc.Field.Type.Model == nil {
		return
	}
	en, ok := cc.Field.Type.Model.(*matter.Enum)
	if !ok {
		return
	}
	for _, v := range en.Values {
		if v.Name == c.Value {
			val, err := parse.HexOrDec(v.Value)
			if err == nil {
				def = matter.NewUintConstraintExtreme(val, matter.NumberFormatInt)
				return
			}
		}
	}
	return
}

func (c *ReferenceLimit) getBitmapValue(cc *matter.ConstraintContext) (def matter.ConstraintExtreme) {
	if cc.Field.Type == nil || cc.Field.Type.Model == nil {
		return
	}
	en, ok := cc.Field.Type.Model.(*matter.Bitmap)
	if !ok {
		return
	}
	for _, v := range en.Bits {
		if v.Name == c.Value {
			val, err := v.Mask()
			if err == nil {
				def = matter.NewUintConstraintExtreme(val, matter.NumberFormatHex)
				return
			}
		}
	}
	return
}

func (c *ReferenceLimit) Default(cc *matter.ConstraintContext) (max matter.ConstraintExtreme) {
	r := c.getReference(cc)
	if r != nil && r.Default != "" {
		cons := ParseConstraint(r.Default)
		return cons.Default(cc)
	}
	// Couldn't find it as a reference; let's check other possibilities
	max = c.getEnumValue(cc)
	if max.Defined() {
		return
	}
	max = c.getBitmapValue(cc)
	if max.Defined() {
		return
	}
	return
}
