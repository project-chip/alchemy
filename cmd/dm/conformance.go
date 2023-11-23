package dm

import (
	"fmt"
	"strconv"

	"github.com/beevik/etree"
	"github.com/hasty/alchemy/conformance"
	"github.com/hasty/alchemy/matter"
)

func renderConformanceString(c matter.Conformance, parent *etree.Element) error {
	if c == nil {
		return nil
	}
	switch cs := c.(type) {
	case conformance.ConformanceSet:
		if len(cs) > 1 {
			oc := parent.CreateElement("otherwiseConform")
			for _, c := range cs {
				err := renderConformance(c, oc)
				if err != nil {
					return err
				}
			}
		} else if len(cs) == 1 {
			return renderConformance(cs[0], parent)
		}
	case *conformance.GenericConformance:
		return nil

	}

	return nil
}

func renderConformance(con matter.Conformance, parent *etree.Element) error {
	switch con := con.(type) {
	case *conformance.MandatoryConformance:
		mc := parent.CreateElement("mandatoryConform")
		renderConformanceExpression(con.Expression, mc)
	case *conformance.ProvisionalConformance:
		parent.CreateElement("provisionalConform")
	case *conformance.OptionalConformance:
		oc := parent.CreateElement("optionalConform")
		if con.Choice != nil {
			oc.CreateAttr("choice", con.Choice.Set)
			if con.Choice.Limit != nil {
				switch l := con.Choice.Limit.(type) {
				case *conformance.ChoiceExactLimit:
					oc.CreateAttr("min", strconv.Itoa(l.Limit))
					oc.CreateAttr("max", strconv.Itoa(l.Limit))
				case *conformance.ChoiceMinLimit:
					oc.CreateAttr("min", strconv.Itoa(l.Min))
				case *conformance.ChoiceMaxLimit:
					oc.CreateAttr("max", strconv.Itoa(l.Max))
				case *conformance.ChoiceRangeLimit:
					oc.CreateAttr("min", strconv.Itoa(l.Min))
					oc.CreateAttr("max", strconv.Itoa(l.Max))
				}
			}
		}
		return renderConformanceExpression(con.Expression, oc)
	case *conformance.DisallowedConformance:
		parent.CreateElement("disallowConform")
	case *conformance.DeprecatedConformance:
		parent.CreateElement("deprecateConform")
	case *conformance.DescribedConformance:

	case conformance.ConformanceSet:
		for _, con := range con {
			err := renderConformance(con, parent)
			if err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unknown conformance type: %T", con)
	}
	return nil
}

func renderConformanceExpression(exp conformance.ConformanceExpression, parent *etree.Element) error {
	if exp == nil {
		return nil
	}
	switch e := exp.(type) {
	case *conformance.FeatureExpression:
		if e.Not {
			parent = parent.CreateElement("notTerm")
		}
		parent.CreateElement("feature").CreateAttr("name", e.ID)
	case *conformance.LogicalExpression:
		if e.Not {
			parent = parent.CreateElement("notTerm")
		}
		var el *etree.Element
		switch e.Operand {
		case "&":
			el = parent.CreateElement("andTerm")
		case "|":
			el = parent.CreateElement("orTerm")
		case "^":
			el = parent.CreateElement("xorTerm")
		default:
			return fmt.Errorf("unknown operand: %s", e.Operand)
		}
		err := renderConformanceExpression(e.Left, el)
		if err != nil {
			return err
		}
		err = renderConformanceExpression(e.Right, el)
		if err != nil {
			return err
		}

	case *conformance.EqualityExpression:
		return nil
		return fmt.Errorf("unknown conformance expression type: %s", e.String())
	default:
		return fmt.Errorf("unknown conformance expression type: %T", exp)
	}
	return nil
}
