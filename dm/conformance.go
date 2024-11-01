package dm

import (
	"fmt"
	"strconv"

	"github.com/beevik/etree"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
)

func RenderConformanceElement(doc *spec.Doc, identifierStore conformance.IdentifierStore, c conformance.Conformance, parent *etree.Element) error {
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
					return fmt.Errorf("error rendering conformance %s: %w", c.ASCIIDocString(), err)
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

func renderConformance(doc *spec.Doc, identifierStore conformance.IdentifierStore, con conformance.Conformance, parent *etree.Element) (err error) {
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
		renderChoice(con.Choice, oc)
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

func renderChoice(choice *conformance.Choice, parent *etree.Element) {
	if choice == nil {
		return
	}
	parent.CreateAttr("choice", choice.Set)
	if choice.Limit == nil {
		return
	}
	switch l := choice.Limit.(type) {
	case *conformance.ChoiceExactLimit:
		parent.CreateAttr("min", strconv.Itoa(l.Limit))
		parent.CreateAttr("max", strconv.Itoa(l.Limit))
	case *conformance.ChoiceMinLimit:
		parent.CreateAttr("more", "true") // Existing data model does this for some reason
		parent.CreateAttr("min", strconv.Itoa(l.Min))
	case *conformance.ChoiceMaxLimit:
		parent.CreateAttr("max", strconv.Itoa(l.Max))
	case *conformance.ChoiceRangeLimit:
		parent.CreateAttr("more", "true") // Existing data model does this for some reason
		parent.CreateAttr("min", strconv.Itoa(l.Min))
		parent.CreateAttr("max", strconv.Itoa(l.Max))
	}
}

func renderConformanceExpression(doc *spec.Doc, identifierStore conformance.IdentifierStore, exp conformance.Expression, parent *etree.Element) error {
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
		renderIdentifier(identifierStore, parent, e.ID)
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
			return fmt.Errorf("error rendering conformance expression %s: %w", e.Left.ASCIIDocString(), err)
		}
		for _, r := range e.Right {
			err = renderConformanceExpression(doc, identifierStore, r, el)
			if err != nil {
				return fmt.Errorf("error rendering conformance expression %s: %w", r.ASCIIDocString(), err)
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
				switch entity := entity.(type) {
				case *matter.Field:
					switch entity.EntityType() {
					case types.EntityTypeAttribute:
						parent.CreateElement("attribute").CreateAttr("name", entity.Name)
					case types.EntityTypeStructField:
						parent.CreateElement("field").CreateAttr("name", entity.Name)
					}
				case *matter.Command:
					parent.CreateElement("command").CreateAttr("name", entity.Name)
				default:
					switch entity.EntityType() {
					case types.EntityTypeCondition:
						parent.CreateElement("attribute").CreateAttr("name", e.Reference)
					default:
						parent.CreateElement("condition").CreateAttr("name", e.Reference)
					}
				}
			}
		}
	case *conformance.ComparisonExpression:
		switch e.Op {
		case conformance.ComparisonOperatorEqual:
			parent = parent.CreateElement("equalTerm")
		case conformance.ComparisonOperatorNotEqual:
			parent = parent.CreateElement("notEqualTerm")
		case conformance.ComparisonOperatorGreaterThan:
			parent = parent.CreateElement("greaterTerm")
		case conformance.ComparisonOperatorGreaterThanOrEqual:
			parent = parent.CreateElement("greaterOrEqualTerm")
		case conformance.ComparisonOperatorLessThan:
			parent = parent.CreateElement("lessTerm")
		case conformance.ComparisonOperatorLessThanOrEqual:
			parent = parent.CreateElement("lessOrEqualTerm")
		default:
			return fmt.Errorf("unexpected comparison expression operator: %s", e.Op.String())
		}
		err := renderComparisonValue(identifierStore, e.Left, parent)
		if err != nil {
			return err
		}
		err = renderComparisonValue(identifierStore, e.Right, parent)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("unimplemented conformance expression type: %T", exp)
	}
	return nil
}

func renderIdentifier(identifierStore conformance.IdentifierStore, parent *etree.Element, name string) {
	if identifierStore == nil {
		parent.CreateElement("condition").CreateAttr("name", name)
	} else {
		entity, ok := identifierStore.Identifier(name)
		if !ok {
			parent.CreateElement("condition").CreateAttr("name", name)
		} else {
			switch entity.EntityType() {
			case types.EntityTypeAttribute, types.EntityTypeCondition:
				parent.CreateElement("attribute").CreateAttr("name", name)
			case types.EntityTypeCommand:
				parent.CreateElement("command").CreateAttr("name", name)
			case types.EntityTypeStructField:
				parent.CreateElement("field").CreateAttr("name", name)
			default:
				parent.CreateElement("condition").CreateAttr("name", name)
			}
		}
	}
}

func renderComparisonValue(identifierStore conformance.IdentifierStore, value conformance.ComparisonValue, parent *etree.Element) (err error) {
	switch value := value.(type) {
	case *conformance.FeatureValue:
		parent.CreateElement("feature").CreateAttr("name", value.Feature)
	case *conformance.IdentifierValue:
		renderIdentifier(identifierStore, parent, value.ID)
	case *conformance.IntValue:
		parent.CreateElement("literal").CreateAttr("value", strconv.FormatInt(value.Int, 10))
	case *conformance.FloatValue:
		parent.CreateElement("literal").CreateAttr("value", value.Float.String())
	case *conformance.HexValue:
		parent.CreateElement("literal").CreateAttr("value", value.ASCIIDocString())
	default:
		return fmt.Errorf("unexpected type in comparison value: %T", value)
	}
	return nil
}

func renderConformanceEqualityExpression(doc *spec.Doc, cluster conformance.IdentifierStore, exp *conformance.EqualityExpression, parent *etree.Element) (err error) {
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
