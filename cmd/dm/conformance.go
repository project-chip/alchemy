package dm

import (
	"fmt"
	"strconv"

	"github.com/beevik/etree"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/conformance"
)

func renderConformanceString(cluster conformance.ValueStore, c conformance.Conformance, parent *etree.Element) error {
	if c == nil {
		return nil
	}
	switch cs := c.(type) {
	case conformance.Set:
		if len(cs) > 1 {
			oc := parent.CreateElement("otherwiseConform")
			for _, c := range cs {
				err := renderConformance(cluster, c, oc)
				if err != nil {
					return fmt.Errorf("error rendering conformance %s: %w", c.String(), err)
				}
			}
		} else if len(cs) == 1 {
			return renderConformance(cluster, cs[0], parent)
		}
	case *conformance.Generic:
		return nil
	default:
		return renderConformance(cluster, c, parent)

	}

	return nil
}

func renderConformance(cluster conformance.ValueStore, con conformance.Conformance, parent *etree.Element) error {
	switch con := con.(type) {
	case *conformance.Mandatory:
		mc := parent.CreateElement("mandatoryConform")
		renderConformanceExpression(cluster, con.Expression, mc)
	case *conformance.Provisional:
		parent.CreateElement("provisionalConform")
	case *conformance.Optional:
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
		return renderConformanceExpression(cluster, con.Expression, oc)
	case *conformance.Disallowed:
		parent.CreateElement("disallowConform")
	case *conformance.Deprecated:
		parent.CreateElement("deprecateConform")
	case *conformance.Described:

	case conformance.Set:
		for _, con := range con {
			err := renderConformance(cluster, con, parent)
			if err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unknown conformance type: %T", con)
	}
	return nil
}

func renderConformanceExpression(cluster conformance.ValueStore, exp conformance.Expression, parent *etree.Element) error {
	if exp == nil {
		return nil
	}
	switch e := exp.(type) {
	case *conformance.FeatureExpression:
		if e.Not {
			parent = parent.CreateElement("notTerm")
		}
		parent.CreateElement("feature").CreateAttr("name", e.ID)
	case *conformance.IdentifierExpression:
		if cluster == nil {
			parent.CreateElement("condition").CreateAttr("name", e.ID)

		} else {
			id := cluster.Reference(e.ID)
			if id == nil {
				parent.CreateElement("condition").CreateAttr("name", e.ID)
			} else {
				model, ok := id.(matter.Model)
				if ok {
					switch model.Entity() {
					case matter.EntityAttribute, matter.EntityCondition:
						parent.CreateElement("attribute").CreateAttr("name", e.ID)
					default:
						parent.CreateElement("condition").CreateAttr("name", e.ID)
					}
				}
			}
		}
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
		err := renderConformanceExpression(cluster, e.Left, el)
		if err != nil {
			return fmt.Errorf("error rendering conformance expression %s: %w", e.Left.String(), err)
		}
		for _, r := range e.Right {
			err = renderConformanceExpression(cluster, r, el)
			if err != nil {
				return fmt.Errorf("error rendering conformance expression %s: %w", r.String(), err)
			}

		}

	case *conformance.EqualityExpression:
		return nil
	default:
		return fmt.Errorf("unknown conformance expression type: %T", exp)
	}
	return nil
}
