package dm

import (
	"fmt"

	"github.com/beevik/etree"
	"github.com/hasty/alchemy/constraint"
	"github.com/hasty/alchemy/matter"
)

func renderConstraint(con matter.Constraint, dataType *matter.DataType, parent *etree.Element) error {
	if con == nil {
		return nil
	}
	_, err := renderConstraintElement("constraint", con, dataType, parent)
	if err != nil {
		return fmt.Errorf("error rendering constraint element %s: %w", con.AsciiDocString(dataType), err)
	}
	return nil
}

func renderConstraintElement(name string, con matter.Constraint, dataType *matter.DataType, parent *etree.Element) (cx *etree.Element, err error) {
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
		cx.CreateAttr("value", con.Value.AsciiDocString(dataType))
	case *constraint.RangeConstraint:
		cx = parent.CreateElement(name)
		if dataType.IsArray() {
			cx.CreateAttr("type", "countBetween")
		} else if dataType.HasLength() {
			cx.CreateAttr("type", "lengthBetween")
		} else {
			cx.CreateAttr("type", "between")
		}
		cx.CreateAttr("from", con.Minimum.AsciiDocString(dataType))
		cx.CreateAttr("to", con.Maximum.AsciiDocString(dataType))
	case *constraint.MinConstraint:
		cx = parent.CreateElement(name)
		if dataType.IsArray() {
			cx.CreateAttr("type", "minCount")
		} else if dataType.HasLength() {
			cx.CreateAttr("type", "minLength")
		} else {
			cx.CreateAttr("type", "min")
		}
		cx.CreateAttr("value", con.Minimum.AsciiDocString(dataType))
	case *constraint.MaxConstraint:
		cx = parent.CreateElement(name)
		if dataType.IsArray() {
			cx.CreateAttr("type", "maxCount")
		} else if dataType.HasLength() {
			cx.CreateAttr("type", "maxLength")
		} else {
			cx.CreateAttr("type", "max")
		}
		cx.CreateAttr("value", con.Maximum.AsciiDocString(dataType))
	case *constraint.ListConstraint:
		if mc, ok := con.Constraint.(*constraint.MaxConstraint); ok {
			cx = parent.CreateElement(name)
			cx.CreateAttr("type", "maxCount")
			cx.CreateAttr("value", mc.Maximum.AsciiDocString(dataType))
		}
	case constraint.ConstraintSet:
		for _, cs := range con {
			renderConstraintElement("constraint", cs, dataType, parent)
		}
	default:
		err = fmt.Errorf("unknown constraint type: %T", con)
	}
	return
}
