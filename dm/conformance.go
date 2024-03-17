package dm

import (
	"fmt"
	"strconv"

	"github.com/beevik/etree"
	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/matter/conformance"
	"github.com/hasty/alchemy/matter/types"
)

func renderConformanceString(doc *ascii.Doc, identifierStore conformance.IdentifierStore, c conformance.Conformance, parent *etree.Element) error {
	if c == nil {
		return nil
	}
	switch cs := c.(type) {
	case conformance.Set:
		if len(cs) > 1 {
			oc := parent.CreateElement("otherwiseConform")
			for _, c := range cs {
				err := renderConformance(doc, identifierStore, c, oc)
				if err != nil {
					return fmt.Errorf("error rendering conformance %s: %w", c.AsciiDocString(), err)
				}
			}
		} else if len(cs) == 1 {
			return renderConformance(doc, identifierStore, cs[0], parent)
		}
	case *conformance.Generic:
		return nil
	default:
		return renderConformance(doc, identifierStore, c, parent)

	}

	return nil
}

func renderConformance(doc *ascii.Doc, identifierStore conformance.IdentifierStore, con conformance.Conformance, parent *etree.Element) (err error) {
	switch con := con.(type) {
	case *conformance.Mandatory:
		_, isEquality := con.Expression.(*conformance.EqualityExpression)
		if !isEquality {
			mc := parent.CreateElement("mandatoryConform")
			err = renderConformanceExpression(doc, identifierStore, con.Expression, mc)
			if err != nil {
				return
			}
		}
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
					oc.CreateAttr("more", "true") // Existing data model does this for some reason
					oc.CreateAttr("min", strconv.Itoa(l.Min))
				case *conformance.ChoiceMaxLimit:
					oc.CreateAttr("max", strconv.Itoa(l.Max))
				case *conformance.ChoiceRangeLimit:
					oc.CreateAttr("more", "true") // Existing data model does this for some reason
					oc.CreateAttr("min", strconv.Itoa(l.Min))
					oc.CreateAttr("max", strconv.Itoa(l.Max))
				}
			}
		}
		return renderConformanceExpression(doc, identifierStore, con.Expression, oc)
	case *conformance.Disallowed:
		parent.CreateElement("disallowConform")
	case *conformance.Deprecated:
		parent.CreateElement("deprecateConform")
	case *conformance.Described:
	case *conformance.Generic:
	case conformance.Set:
		for _, con := range con {
			err = renderConformance(doc, identifierStore, con, parent)
			if err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unknown conformance type: %T", con)
	}
	return nil
}

func renderConformanceExpression(doc *ascii.Doc, identifierStore conformance.IdentifierStore, exp conformance.Expression, parent *etree.Element) error {
	if exp == nil {
		return nil
	}
	switch e := exp.(type) {
	case *conformance.FeatureExpression:
		if e.Not {
			parent = parent.CreateElement("notTerm")
		}
		parent.CreateElement("feature").CreateAttr("name", e.Feature)
	case *conformance.IdentifierExpression:
		if e.Not {
			parent = parent.CreateElement("notTerm")
		}
		if identifierStore == nil {
			parent.CreateElement("condition").CreateAttr("name", e.ID)
		} else {
			entity, ok := identifierStore.Identifier(e.ID)
			if !ok {
				parent.CreateElement("condition").CreateAttr("name", e.ID)
			} else {
				switch entity.EntityType() {
				case types.EntityTypeAttribute, types.EntityTypeCondition:
					parent.CreateElement("attribute").CreateAttr("name", e.ID)
				case types.EntityTypeCommand:
					parent.CreateElement("command").CreateAttr("name", e.ID)
				case types.EntityTypeField:
					parent.CreateElement("field").CreateAttr("name", e.ID)
				default:
					parent.CreateElement("condition").CreateAttr("name", e.ID)
				}
			}
		}
	case *conformance.EqualityExpression:
		return renderConformanceEqualityExpression(doc, identifierStore, e, parent)
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
		err := renderConformanceExpression(doc, identifierStore, e.Left, el)
		if err != nil {
			return fmt.Errorf("error rendering conformance expression %s: %w", e.Left.AsciiDocString(), err)
		}
		for _, r := range e.Right {
			err = renderConformanceExpression(doc, identifierStore, r, el)
			if err != nil {
				return fmt.Errorf("error rendering conformance expression %s: %w", r.AsciiDocString(), err)
			}

		}

	case *conformance.ReferenceExpression:
		if e.Not {
			parent = parent.CreateElement("notTerm")
		}
		if doc == nil {
			parent.CreateElement("condition").CreateAttr("name", e.Reference)
		} else {
			entity, ok := doc.Reference(e.Reference)
			if !ok {
				parent.CreateElement("condition").CreateAttr("name", e.Reference)
			} else {
				switch entity.EntityType() {
				case types.EntityTypeAttribute, types.EntityTypeCondition:
					parent.CreateElement("attribute").CreateAttr("name", e.Reference)
				case types.EntityTypeCommand:
					parent.CreateElement("command").CreateAttr("name", e.Reference)
				case types.EntityTypeField:
					parent.CreateElement("field").CreateAttr("name", e.Reference)
				default:
					parent.CreateElement("condition").CreateAttr("name", e.Reference)
				}
			}
		}
	default:
		return fmt.Errorf("unimplemented conformance expression type: %T", exp)
	}
	return nil
}

func renderConformanceEqualityExpression(doc *ascii.Doc, cluster conformance.IdentifierStore, exp *conformance.EqualityExpression, parent *etree.Element) (err error) {
	var e *etree.Element
	if exp.Not {
		e = parent.CreateElement("notEqualTerm")
	} else {
		e = parent.CreateElement("equalTerm")
	}
	err = renderConformanceExpression(doc, cluster, exp.Left, e)
	if err != nil {
		return
	}
	err = renderConformanceExpression(doc, cluster, exp.Right, e)
	return
}
