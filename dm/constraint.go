package dm

import (
	"fmt"
	"log/slog"
	"strconv"

	"github.com/beevik/etree"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/constraint"
	"github.com/project-chip/alchemy/matter/types"
)

type identifierStore interface {
	Identifier(name string) (types.Entity, bool)
}

func renderConstraint(con constraint.Constraint, dataType *types.DataType, parent *etree.Element, parentEntity types.Entity) error {
	if con == nil {
		return nil
	}
	err := renderConstraintElement(con, dataType, parent, parentEntity)
	if err != nil {
		return fmt.Errorf("error rendering constraint element %s: %w", con.ASCIIDocString(dataType), err)
	}
	return nil
}

func renderConstraintElement(con constraint.Constraint, dataType *types.DataType, parent *etree.Element, parentEntity types.Entity) (err error) {
	if con == nil {
		return
	}
	var constraints []*etree.Element
	name := "constraint"
	switch con := con.(type) {
	case *constraint.AllConstraint, *constraint.GenericConstraint:
		return
	case *constraint.DescribedConstraint:
		cx := etree.NewElement(name)
		cx.CreateElement("desc")
		constraints = append(constraints, cx)
	case *constraint.ExactConstraint:
		cx := etree.NewElement(name)
		renderConstraintLimit(cx.CreateElement("allowed"), con.Value, dataType, "value", parentEntity)
		constraints = append(constraints, cx)
	case *constraint.RangeConstraint:
		cx := etree.NewElement(name)
		var rangeElement *etree.Element
		if dataType.IsArray() {
			rangeElement = cx.CreateElement("countBetween")
		} else if dataType.HasLength() {
			rangeElement = cx.CreateElement("lengthBetween")
		} else {
			rangeElement = cx.CreateElement("between")
		}
		fromRef, fromEntity, fromField := getLimitField(con.Minimum)
		toRef, toEntity, toField := getLimitField(con.Maximum)
		if fromEntity == nil && fromField == nil && toField == nil && toEntity == nil {
			rangeElement.CreateElement("from").CreateAttr("value", renderLimitValue(con.Minimum, dataType))
			rangeElement.CreateElement("to").CreateAttr("value", renderLimitValue(con.Maximum, dataType))
		} else {
			renderConstraintReferenceLimit(rangeElement, fromField, con.Minimum, dataType, "from", fromEntity, fromRef, parentEntity)
			renderConstraintReferenceLimit(rangeElement, toField, con.Maximum, dataType, "to", toEntity, toRef, parentEntity)
		}
		constraints = append(constraints, cx)
	case *constraint.MinConstraint:
		cx := etree.NewElement(name)
		var minElement *etree.Element
		if dataType.IsArray() {
			minElement = cx.CreateElement("minCount")
		} else if dataType.HasLength() {
			minElement = cx.CreateElement("minLength")
		} else {
			minElement = cx.CreateElement("min")
		}
		renderConstraintLimit(minElement, con.Minimum, dataType, "value", parentEntity)
		constraints = append(constraints, cx)
	case *constraint.MaxConstraint:
		cx := etree.NewElement(name)
		var maxElement *etree.Element
		if dataType.IsArray() {
			maxElement = cx.CreateElement("maxCount")
		} else if dataType.HasLength() {
			maxElement = cx.CreateElement("maxLength")
		} else {
			maxElement = cx.CreateElement("max")
		}
		renderConstraintLimit(maxElement, con.Maximum, dataType, "value", parentEntity)
		constraints = append(constraints, cx)
		if characterLimit, ok := con.Maximum.(*constraint.CharacterLimit); ok {
			cx := parent.CreateElement(name)
			maxElement = cx.CreateElement("maxCodePoints")
			renderConstraintLimit(maxElement, characterLimit.CodepointCount, dataType, "value", parentEntity)
			constraints = append(constraints, cx)
		}
	case *constraint.ListConstraint:
		err = renderConstraintElement(con.Constraint, dataType, parent, parentEntity)
	case *constraint.TagListConstraint:
		cx := etree.NewElement(name)
		tagsElement := cx.CreateElement("tags")
		renderConstraintLimit(tagsElement, con.Tags, dataType, "value", parentEntity)
		constraints = append(constraints, cx)
	case constraint.Set:
		for _, cs := range con {
			err = renderConstraintElement(cs, dataType, parent, parentEntity)
			if err != nil {
				return
			}
		}
	default:
		err = fmt.Errorf("unknown constraint type: %T", con)
	}
	for _, cx := range constraints {
		parent.AddChild(cx)

	}
	return
}

func renderConstraintLimit(parent *etree.Element, limit constraint.Limit, dataType *types.DataType, name string, parentEntity types.Entity) {
	switch limit := limit.(type) {
	case *constraint.TagIdentifierLimit:
		t := parent.CreateElement("tag")
		t.CreateAttr("name", limit.Tag)
	case *constraint.LogicalLimit:
		renderLogicalLimit(parent, limit, dataType, name, parentEntity)
	default:
		ref, entity, field := getLimitField(limit)
		if entity == nil && field == nil {
			parent.CreateAttr(name, renderLimitValue(limit, dataType))
			return
		}
		renderConstraintReferenceLimit(parent, field, limit, dataType, name, entity, ref, parentEntity)
	}
}

func renderLogicalLimit(parent *etree.Element, limit *constraint.LogicalLimit, dataType *types.DataType, name string, parentEntity types.Entity) {
	if limit.Not {
		parent = parent.CreateElement("notTerm")
	}
	var el *etree.Element
	switch limit.Operand {
	case "&", "and":
		el = parent.CreateElement("andTerm")
	case "|", "or":
		el = parent.CreateElement("orTerm")
	case "^":
		el = parent.CreateElement("xorTerm")
	default:
		el = parent
	}
	renderConstraintLimit(el, limit.Left, dataType, name, parentEntity)
	for _, r := range limit.Right {
		renderConstraintLimit(el, r, dataType, name, parentEntity)
	}
}

func renderConstraintReferenceLimit(parent *etree.Element, field constraint.Limit, limit constraint.Limit, dataType *types.DataType, name string, entity types.Entity, ref string, parentEntity types.Entity) bool {
	if name != "value" {
		parent = parent.CreateElement(name)
	}
	switch entity := entity.(type) {
	case *matter.EnumValue:
		el := parent.CreateElement("enum")
		el.CreateAttr(name, entity.Name)
		return true
	case matter.Bit:
		mask, err := entity.Mask()
		if err == nil {
			el := parent.CreateElement("bitmap")
			el.CreateAttr("mask", strconv.FormatUint(mask, 10))
		}
		return true
	case *matter.Field:
		renderFieldConstraint(parent, entity, ref, field, parentEntity)
	case *matter.Constant:
		el := parent.CreateElement("constant")
		el.CreateAttr("name", entity.Name)
		switch val := entity.Value.(type) {
		case int:
			el.CreateAttr("value", strconv.FormatInt(int64(val), 10))
		default:
			slog.Info("unknown constant value type", log.Type("type", val))
		}
	case nil:
		parent.CreateAttr("value", renderLimitValue(limit, dataType))
	default:
		slog.Info("unexpected constraint limit", "entity", entity, "limit", limit, "dataType", dataType, "field", field)
	}
	return false
}

func renderFieldConstraint(parent *etree.Element, entity *matter.Field, ref string, field constraint.Limit, parentEntity types.Entity) {
	switch entity.EntityType() {
	case types.EntityTypeAttribute:
		parent = parent.CreateElement("attribute")
		parent.CreateAttr("name", entity.Name)
	case types.EntityTypeStructField:
		fieldParent := entity.Parent()
		if fieldParent != nil && fieldParent != parentEntity {
			switch fieldParent := fieldParent.(type) {
			case *matter.Struct:
				parent = parent.CreateElement("struct")
				parent.CreateAttr("name", fieldParent.Name)
			default:
				slog.Warn("Unexpected struct field parent entity type", log.Path("source", entity), log.Type("parentEntityType", fieldParent))
			}
		}
		parent = parent.CreateElement("field")
		parent.CreateAttr("name", entity.Name)
	case types.EntityTypeCommandField:
		fieldParent := entity.Parent()
		if fieldParent != nil && fieldParent != parentEntity {
			switch fieldParent := fieldParent.(type) {
			case *matter.Command:
				parent = parent.CreateElement("command")
				parent.CreateAttr("name", fieldParent.Name)
			default:
				slog.Warn("Unexpected command field parent entity type", log.Path("source", entity), log.Type("parentEntityType", fieldParent))
			}
		}
		parent = parent.CreateElement("field")
		parent.CreateAttr("name", entity.Name)
	case types.EntityTypeEventField:
		fieldParent := entity.Parent()
		if fieldParent != nil && fieldParent != parentEntity {
			switch fieldParent := fieldParent.(type) {
			case *matter.Event:
				parent = parent.CreateElement("event")
				parent.CreateAttr("name", fieldParent.Name)
			default:
				slog.Warn("Unexpected event field parent entity type", log.Path("source", entity), log.Type("parentEntityType", fieldParent))
			}
		}
		parent = parent.CreateElement("field")
		parent.CreateAttr("name", entity.Name)
	default:
		slog.Warn("Unexpected entity type on reference limit", log.Type("type", entity))
		parent.CreateAttr("reference", ref)
	}
	if field != nil {
		renderConstraintLimit(parent, field, nil, "value", entity.Type.Entity)
	}
}

func renderLimitValue(limit constraint.Limit, dataType *types.DataType) string {
	limitString := limit.DataModelString(dataType)
	switch limit.(type) {
	case *constraint.MathExpressionLimit:
		if len(limitString) > 2 && limitString[0] == '(' && limitString[len(limitString)-1] == ')' {
			limitString = limitString[1 : len(limitString)-1]
		}

	}
	return limitString
}

func getLimitField(limit constraint.Limit) (ref string, entity types.Entity, field constraint.Limit) {
	switch limit := limit.(type) {
	case *constraint.IdentifierLimit:
		ref = limit.ID
		entity = limit.Entity
		field = limit.Field
	case *constraint.ReferenceLimit:
		ref = limit.Reference
		entity = limit.Entity
		field = limit.Field
	case *constraint.LengthLimit:
		ref, entity, field = getLimitField(limit.Reference)
	}
	return
}

/*func renderConstraintRange(parent *etree.Element, from constraint.Limit, to constraint.Limit, dataType *types.DataType) {
	fromLimitString, fromField := renderLimitValue(from, dataType)
	toLimitString, toField := renderLimitValue(to, dataType)
	limitString := from.DataModelString(dataType)
	var field string
	switch limit := from.(type) {
	case *constraint.MathExpressionLimit:
		if len(limitString) > 2 && limitString[0] == '(' && limitString[len(limitString)-1] == ')' {
			limitString = limitString[1 : len(limitString)-1]
		}
	case *constraint.IdentifierLimit:
		field = limit.Field
	case *constraint.ReferenceLimit:
		field = limit.Field
	case *constraint.LengthLimit:
		field = limit.Field
	}
	if field == "" {
		parent.CreateAttr("value", limitString)
		return
	}
}*/

/*func renderConstraintLimit(limit constraint.Limit, dataType *types.DataType) string {
	s := limit.DataModelString(dataType)
	switch limit.(type) {
	case *constraint.MathExpressionLimit:
		if len(s) > 2 && s[0] == '(' && s[len(s)-1] == ')' {
			s = s[1 : len(s)-1]
		}
	}
	return s
}*/
