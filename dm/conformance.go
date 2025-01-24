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

func renderConformanceElement(doc *spec.Doc, identifierStore conformance.IdentifierStore, c conformance.Conformance, parent *etree.Element) error {
	conformanceElement, err := CreateConformanceElement(doc, identifierStore, c)
	if err != nil {
		return err
	}
	if conformanceElement != nil {
		parent.AddChild(conformanceElement)
	}
	return nil
}

func CreateConformanceElement(doc *spec.Doc, identifierStore conformance.IdentifierStore, c conformance.Conformance) (element *etree.Element, err error) {
	if c == nil {
		return
	}
	switch cs := c.(type) {
	case conformance.Set:
		if len(cs) > 1 {
			return renderConformanceSet(doc, identifierStore, cs)
		} else if len(cs) == 1 {
			return renderConformance(doc, identifierStore, cs[0])
		}
	case *conformance.Generic:
		return
	default:
		return renderConformance(doc, identifierStore, c)

	}
	return
}

func renderConformanceSet(doc *spec.Doc, identifierStore conformance.IdentifierStore, set conformance.Set) (element *etree.Element, err error) {
	element = etree.NewElement("otherwiseConform")
	for _, c := range set {
		var ce *etree.Element
		ce, err = renderConformance(doc, identifierStore, c)
		if err != nil {
			err = fmt.Errorf("error rendering conformance %s: %w", c.ASCIIDocString(), err)
			return
		}
		element.AddChild(ce)
	}
	return
}

func renderConformance(doc *spec.Doc, identifierStore conformance.IdentifierStore, con conformance.Conformance) (element *etree.Element, err error) {
	switch con := con.(type) {
	case *conformance.Mandatory:
		_, isEquality := con.Expression.(*conformance.EqualityExpression)
		if !isEquality {
			element = etree.NewElement("mandatoryConform")
			err = renderConformanceExpression(doc, identifierStore, con.Expression, element)
			if err != nil {
				return
			}
		}
	case *conformance.Provisional:
		element = etree.NewElement("provisionalConform")
	case *conformance.Optional:
		element = etree.NewElement("optionalConform")
		renderChoice(con.Choice, element)
		err = renderConformanceExpression(doc, identifierStore, con.Expression, element)
		return
	case *conformance.Disallowed:
		element = etree.NewElement("disallowConform")
	case *conformance.Deprecated:
		element = etree.NewElement("deprecateConform")
	case *conformance.Described:
	case *conformance.Generic:
	case conformance.Set:
		return renderConformanceSet(doc, identifierStore, con)
	default:
		err = fmt.Errorf("unknown conformance type: %T", con)
	}
	return
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
		parent.CreateAttr("min", strconv.FormatInt(l.Limit, 10))
		parent.CreateAttr("max", strconv.FormatInt(l.Limit, 10))
	case *conformance.ChoiceMinLimit:
		parent.CreateAttr("more", "true") // Existing data model does this for some reason
		parent.CreateAttr("min", strconv.FormatInt(l.Min, 10))
	case *conformance.ChoiceMaxLimit:
		parent.CreateAttr("max", strconv.FormatInt(l.Max, 10))
	case *conformance.ChoiceRangeLimit:
		parent.CreateAttr("more", "true") // Existing data model does this for some reason
		parent.CreateAttr("min", strconv.FormatInt(l.Min, 10))
		parent.CreateAttr("max", strconv.FormatInt(l.Max, 10))
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
		var el *etree.Element
		if doc == nil {
			el = parent.CreateElement("condition")
			el.CreateAttr("name", e.Reference)
		} else {
			entity, ok := doc.Reference(e.Reference)
			if !ok {
				el = parent.CreateElement("condition")
				el.CreateAttr("name", e.Reference)
			} else {
				switch entity := entity.(type) {
				case *matter.Field:
					switch entity.EntityType() {
					case types.EntityTypeAttribute:
						el = parent.CreateElement("attribute")
						el.CreateAttr("name", entity.Name)
					case types.EntityTypeStructField:
						el = parent.CreateElement("field")
						el.CreateAttr("name", entity.Name)
					}
				case *matter.Command:
					el = parent.CreateElement("command")
					el.CreateAttr("name", entity.Name)
				default:
					switch entity.EntityType() {
					case types.EntityTypeCondition:
						el = parent.CreateElement("attribute")
					default:
						parent.CreateElement("condition")
					}
					el.CreateAttr("name", e.Reference)
				}
			}
		}
		if e.Field != "" && el != nil {
			el.CreateAttr("property", e.Field)
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
	case *conformance.ReferenceValue:
		parent.CreateElement("reference").CreateAttr("name", value.Reference)
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
