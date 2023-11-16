package constraint

import "github.com/hasty/alchemy/matter"

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

func (c *ReferenceLimit) MinMax(cc *matter.ConstraintContext) (min matter.ConstraintExtreme, max matter.ConstraintExtreme) {
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
	return r.Constraint.MinMax(cc)
}
