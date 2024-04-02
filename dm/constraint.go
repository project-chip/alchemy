package dm

import (
	"fmt"

	"github.com/beevik/etree"
	"github.com/hasty/alchemy/matter/constraint"
	"github.com/hasty/alchemy/matter/types"
)

func renderConstraint(con constraint.Constraint, dataType *types.DataType, parent *etree.Element) error {
	if con == nil {
		return nil
	}
	_, err := renderConstraintElement("constraint", con, dataType, parent)
	if err != nil {
		return fmt.Errorf("error rendering constraint element %s: %w", con.AsciiDocString(dataType), err)
	}
	return nil
}

func renderConstraintElement(name string, con constraint.Constraint, dataType *types.DataType, parent *etree.Element) (cx *etree.Element, err error) {
	if con == nil {
		return
	}
	switch con := con.(type) {
	case *constraint.AllConstraint, *constraint.GenericConstraint:
		return
	case *constraint.DescribedConstraint:
		cx = parent.CreateElement(name)
		cx.CreateAttr("type", "desc")
	case *constraint.ExactConstraint:
		cx = parent.CreateElement(name)
		cx.CreateAttr("type", "allowed")
		cx.CreateAttr("value", renderConstraintLimit(con.Value, dataType))
	case *constraint.RangeConstraint:
		cx = parent.CreateElement(name)
		if dataType.IsArray() {
			cx.CreateAttr("type", "countBetween")
		} else if dataType.HasLength() {
			cx.CreateAttr("type", "lengthBetween")
		} else {
			cx.CreateAttr("type", "between")
		}
		cx.CreateAttr("from", renderConstraintLimit(con.Minimum, dataType))
		cx.CreateAttr("to", renderConstraintLimit(con.Maximum, dataType))
	case *constraint.MinConstraint:
		cx = parent.CreateElement(name)
		if dataType.IsArray() {
			cx.CreateAttr("type", "minCount")
		} else if dataType.HasLength() {
			cx.CreateAttr("type", "minLength")
		} else {
			cx.CreateAttr("type", "min")
		}
		cx.CreateAttr("value", renderConstraintLimit(con.Minimum, dataType))
	case *constraint.MaxConstraint:
		cx = parent.CreateElement(name)
		if dataType.IsArray() {
			cx.CreateAttr("type", "maxCount")
		} else if dataType.HasLength() {
			cx.CreateAttr("type", "maxLength")
		} else {
			cx.CreateAttr("type", "max")
		}
		cx.CreateAttr("value", renderConstraintLimit(con.Maximum, dataType))
	case *constraint.ListConstraint:
		if mc, ok := con.Constraint.(*constraint.MaxConstraint); ok {
			cx = parent.CreateElement(name)
			cx.CreateAttr("type", "maxCount")
			cx.CreateAttr("value", renderConstraintLimit(mc.Maximum, dataType))
		}
	case constraint.Set:
		for _, cs := range con {
			_, err = renderConstraintElement("constraint", cs, dataType, parent)
			if err != nil {
				return
			}
		}
	default:
		err = fmt.Errorf("unknown constraint type: %T", con)
	}
	return
}

func renderConstraintLimit(limit constraint.Limit, dataType *types.DataType) string {
	s := limit.DataModelString(dataType)
	switch limit.(type) {
	case *constraint.MathExpressionLimit:
		if len(s) > 2 && s[0] == '(' && s[len(s)-1] == ')' {
			s = s[1 : len(s)-1]
		}
	}
	return s
}
