package constraint

import (
	"github.com/hasty/alchemy/matter"
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
	r := cc.Fields.GetField(c.Value)
	if cc.VisitedReferences == nil {
		cc.VisitedReferences = make(map[string]struct{})
	}
	if _, ok := cc.VisitedReferences[c.Value]; ok {
		return
	}
	cc.VisitedReferences[c.Value] = struct{}{}
	if r == nil || r.Constraint == nil {
		return
	}
	return r.Constraint.Min(cc)
}

func (c *ReferenceLimit) Max(cc *matter.ConstraintContext) (max matter.ConstraintExtreme) {
	r := cc.Fields.GetField(c.Value)
	if cc.VisitedReferences == nil {
		cc.VisitedReferences = make(map[string]struct{})
	}
	if _, ok := cc.VisitedReferences[c.Value]; ok {
		return
	}
	cc.VisitedReferences[c.Value] = struct{}{}
	if r == nil || r.Constraint == nil {
		return
	}
	return r.Constraint.Max(cc)
}
