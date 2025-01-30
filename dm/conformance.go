package dm

import (
	"fmt"
	"log/slog"
	"strconv"

	"github.com/beevik/etree"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
)

func renderConformanceElement(doc *spec.Doc, c conformance.Conformance, parent *etree.Element, parentEntity types.Entity) error {
	conformanceElement, err := CreateConformanceElement(doc, c, parentEntity)
	if err != nil {
		return err
	}
	if conformanceElement != nil {
		parent.AddChild(conformanceElement)
	}
	return nil
}

func CreateConformanceElement(doc *spec.Doc, c conformance.Conformance, parentEntity types.Entity) (element *etree.Element, err error) {
	if c == nil {
		return
	}
	switch cs := c.(type) {
	case conformance.Set:
		if len(cs) > 1 {
			return renderConformanceSet(doc, cs, parentEntity)
		} else if len(cs) == 1 {
			return renderConformance(doc, cs[0], parentEntity)
		}
	case *conformance.Generic:
		return
	default:
		return renderConformance(doc, c, parentEntity)

	}
	return
}

func renderConformanceSet(doc *spec.Doc, set conformance.Set, parentEntity types.Entity) (element *etree.Element, err error) {
	element = etree.NewElement("otherwiseConform")
	for _, c := range set {
		var ce *etree.Element
		ce, err = renderConformance(doc, c, parentEntity)
		if err != nil {
			err = fmt.Errorf("error rendering conformance %s: %w", c.ASCIIDocString(), err)
			return
		}
		element.AddChild(ce)
	}
	return
}

func renderConformance(doc *spec.Doc, con conformance.Conformance, parentEntity types.Entity) (element *etree.Element, err error) {
	switch con := con.(type) {
	case *conformance.Mandatory:
		_, isEquality := con.Expression.(*conformance.EqualityExpression)
		if !isEquality {
			element = etree.NewElement("mandatoryConform")
			err = renderConformanceExpression(doc, con.Expression, element, parentEntity)
			if err != nil {
				return
			}
		}
	case *conformance.Provisional:
		element = etree.NewElement("provisionalConform")
	case *conformance.Optional:
		element = etree.NewElement("optionalConform")
		renderChoice(con.Choice, element)
		err = renderConformanceExpression(doc, con.Expression, element, parentEntity)
		return
	case *conformance.Disallowed:
		element = etree.NewElement("disallowConform")
	case *conformance.Deprecated:
		element = etree.NewElement("deprecateConform")
	case *conformance.Described:
	case *conformance.Generic:
	case conformance.Set:
		return renderConformanceSet(doc, con, parentEntity)
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

func renderConformanceExpression(doc *spec.Doc, exp conformance.Expression, parent *etree.Element, parentEntity types.Entity) error {
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
		renderConformanceEntity(parent, e.Entity, e.ID, e.Field, parentEntity)
	case *conformance.EqualityExpression:
		return renderConformanceEqualityExpression(doc, e, parent, parentEntity)
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
		err := renderConformanceExpression(doc, e.Left, el, parentEntity)
		if err != nil {
			return fmt.Errorf("error rendering conformance expression %s: %w", e.Left.ASCIIDocString(), err)
		}
		for _, r := range e.Right {
			err = renderConformanceExpression(doc, r, el, parentEntity)
			if err != nil {
				return fmt.Errorf("error rendering conformance expression %s: %w", r.ASCIIDocString(), err)
			}

		}

	case *conformance.ReferenceExpression:
		if e.Not {
			parent = parent.CreateElement("notTerm")
		}
		renderConformanceEntity(parent, e.Entity, e.Reference, e.Field, parentEntity)

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
		err := renderComparisonValue(e.Left, parent, parentEntity)
		if err != nil {
			return err
		}
		err = renderComparisonValue(e.Right, parent, parentEntity)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("unimplemented conformance expression type: %T", exp)
	}
	return nil
}

func renderConformanceEntity(parent *etree.Element, entity types.Entity, name string, field conformance.ComparisonValue, parentEntity types.Entity) {
	var el *etree.Element
	switch entity := entity.(type) {
	case *matter.Command:
		el = parent.CreateElement("command")
		el.CreateAttr("name", entity.Name)
	case *matter.Field:
		switch entity.EntityType() {
		case types.EntityTypeAttribute:
			el = parent.CreateElement("attribute")
			el.CreateAttr("name", entity.Name)
		case types.EntityTypeStructField, types.EntityTypeCommandField, types.EntityTypeEventField:
			el = writeOptionalParentElement(parent, entity, parentEntity)
		default:
			slog.Warn("Unexpected field entity type when rendering conformance", slog.String("entityType", entity.EntityType().String()))
		}
	case *matter.Condition:
		el = parent.CreateElement("condition")
		el.CreateAttr("name", entity.Feature)
	case *matter.TypeDef:
		el = parent.CreateElement("typeDef")
		el.CreateAttr("name", entity.Name)
	case *matter.Enum:
		el = parent.CreateElement("enum")
		el.CreateAttr("name", entity.Name)
	case nil:
		el = parent.CreateElement("condition")
		el.CreateAttr("name", name)
	case *matter.EnumValue:
		enumParent := entity.Parent()
		if enumParent != nil && enumParent != parentEntity {
			switch enumParent := enumParent.(type) {
			case *matter.Enum:
				el = parent.CreateElement("enum")
				el.CreateAttr("name", enumParent.Name)
				el.CreateAttr("value", entity.Name)
			default:
				slog.Warn("Unexpected enum value parent entity type", log.Type("parentEntityType", enumParent))
			}
		} else {
			el = parent.CreateElement("value")
			el.CreateAttr("name", entity.Name)
		}
	default:
		slog.Warn("Unexpected conformance entity type", log.Type("type", entity))
	}

	if field != nil && el != nil {
		renderComparisonValue(field, el, entity)
	}
}

func writeOptionalParentElement(parent *etree.Element, entity *matter.Field, parentEntity types.Entity) (el *etree.Element) {
	fieldParent := entity.Parent()
	if fieldParent != nil && fieldParent != parentEntity {
		switch fieldParent := fieldParent.(type) {
		case *matter.Event:
			el = parent.CreateElement("event")
			el.CreateAttr("name", fieldParent.Name)
			el.CreateAttr("field", entity.Name)
		case *matter.Command:
			el = parent.CreateElement("command")
			el.CreateAttr("name", fieldParent.Name)
			el.CreateAttr("field", entity.Name)
		case *matter.Struct:
			el = parent.CreateElement("struct")
			el.CreateAttr("name", fieldParent.Name)
			el.CreateAttr("field", entity.Name)
		default:
			slog.Warn("Unexpected struct field parent entity type", log.Type("parentEntityType", fieldParent))
		}
	} else {
		el = parent.CreateElement("field")
		el.CreateAttr("name", entity.Name)
	}
	return
}

func renderComparisonValue(value conformance.ComparisonValue, parent *etree.Element, parentEntity types.Entity) (err error) {
	switch value := value.(type) {
	case *conformance.FeatureValue:
		parent.CreateElement("feature").CreateAttr("name", value.Feature)
	case *conformance.IdentifierValue:
		renderConformanceEntity(parent, value.Entity, value.ID, value.Field, parentEntity)
	case *conformance.ReferenceValue:
		renderConformanceEntity(parent, value.Entity, value.Reference, value.Field, parentEntity)
	case *conformance.IntValue:
		parent.CreateElement("literal").CreateAttr("value", strconv.FormatInt(value.Int, 10))
	case *conformance.FloatValue:
		parent.CreateElement("literal").CreateAttr("value", value.Float.String())
	case *conformance.HexValue:
		parent.CreateElement("literal").CreateAttr("value", value.ASCIIDocString())
	case *conformance.StatusCodeValue:
		parent.CreateElement("status").CreateAttr("name", value.StatusCode.String())
	default:
		slog.Warn("Unexpected type in comparison value", log.Type("type", value))
	}
	return nil
}

func renderConformanceEqualityExpression(doc *spec.Doc, exp *conformance.EqualityExpression, parent *etree.Element, parentEntity types.Entity) (err error) {
	var e *etree.Element
	if exp.Not {
		e = parent.CreateElement("notEqualTerm")
	} else {
		e = parent.CreateElement("equalTerm")
	}
	err = renderConformanceExpression(doc, exp.Left, e, parentEntity)
	if err != nil {
		return
	}
	err = renderConformanceExpression(doc, exp.Right, e, parentEntity)
	return
}
