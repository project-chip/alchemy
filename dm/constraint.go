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
	case constraint.Set:
		for _, cs := range con {
			err = renderConstraintElement(cs, dataType, parent, parentEntity)
			if err != nil {
				return
			}
		}
	case *constraint.AllConstraint, *constraint.GenericConstraint:
		return
	case *constraint.DescribedConstraint:
		cx := etree.NewElement(name)
		cx.CreateElement("desc")
		constraints = append(constraints, cx)
	case *constraint.ExactConstraint:
		cx := etree.NewElement(name)
		allowed := cx.CreateElement("allowed")
		err = renderConstraintLimit(allowed, allowed, con.Value, dataType, "value", parentEntity)
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
		if fromEntity == nil && fromField == nil {
			fromElement := rangeElement.CreateElement("from")
			err = renderConstraintLimit(fromElement, fromElement, con.Minimum, dataType, "value", parentEntity)
			if err != nil {
				return
			}
		} else {
			renderConstraintReferenceLimit(rangeElement, fromField, con.Minimum, dataType, "from", fromEntity, fromRef, parentEntity)
		}
		if toField == nil && toEntity == nil {
			toElement := rangeElement.CreateElement("to")
			err = renderConstraintLimit(toElement, toElement, con.Maximum, dataType, "value", parentEntity)
			if err != nil {
				return
			}
		} else {
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
		err = renderConstraintLimit(minElement, minElement, con.Minimum, dataType, "value", parentEntity)
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
		switch limit := con.Maximum.(type) {
		case *constraint.CharacterLimit:
			err = renderConstraintLimit(maxElement, maxElement, limit.ByteCount, dataType, "value", parentEntity)
			if err != nil {
				return
			}
			constraints = append(constraints, cx)
			cx := parent.CreateElement(name)
			maxElement = cx.CreateElement("maxCodePoints")
			err = renderConstraintLimit(maxElement, maxElement, limit.CodepointCount, dataType, "value", parentEntity)
			if err != nil {
				return
			}
			constraints = append(constraints, cx)
		default:
			err = renderConstraintLimit(maxElement, maxElement, con.Maximum, dataType, "value", parentEntity)
			if err != nil {
				return
			}
			constraints = append(constraints, cx)
		}
	case *constraint.ListConstraint:
		err = renderConstraintElement(con.Constraint, dataType, parent, parentEntity)
	case *constraint.TagListConstraint:
		cx := etree.NewElement(name)
		tagsElement := cx.CreateElement("tags")
		err = renderConstraintLimit(tagsElement, tagsElement, con.Tags, dataType, "value", parentEntity)
		constraints = append(constraints, cx)

	default:
		err = fmt.Errorf("unknown constraint type: %T", con)
	}
	if err != nil {
		return
	}
	for _, cx := range constraints {
		parent.AddChild(cx)
	}
	return
}

func renderLimit(parent *etree.Element, valueElement *etree.Element, name string, limit string) {
	if valueElement != nil {
		valueElement.CreateAttr(name, limit)
	} else {
		valueElement = parent.CreateElement(name)
		valueElement.SetText(limit)
	}
}

func renderConstraintLimit(parent *etree.Element, valueElement *etree.Element, limit constraint.Limit, dataType *types.DataType, name string, parentEntity types.Entity) (err error) {

	switch limit := limit.(type) {
	case *constraint.TagIdentifierLimit:
		t := parent.CreateElement("tag")
		t.CreateAttr("name", limit.Tag)
	case *constraint.LogicalLimit:
		renderLogicalLimit(parent, limit, dataType, name, parentEntity)
	case *constraint.MaxOfLimit:
		t := parent.CreateElement("maxOf")
		for _, l := range limit.Maximums {
			err = renderConstraintLimit(t, nil, l, dataType, name, parentEntity)
			if err != nil {
				return
			}
		}
	case *constraint.MinOfLimit:
		t := parent.CreateElement("minOf")
		for _, l := range limit.Minimums {
			err = renderConstraintLimit(t, nil, l, dataType, name, parentEntity)
			if err != nil {
				return
			}
		}
	case *constraint.IdentifierLimit, *constraint.ReferenceLimit, *constraint.LengthLimit:
		ref, entity, field := getLimitField(limit)
		if entity == nil {
			//err = fmt.Errorf("missing entity on identifier or reference limit")
			slog.Warn("Missing entity on identifier or reference limit")
			renderLimit(parent, valueElement, name, ref)
			return
		}
		renderConstraintReferenceLimit(parent, field, limit, dataType, name, entity, ref, parentEntity)
	case *constraint.NullLimit:
		renderLimit(parent, valueElement, name, "null")
	case *constraint.IntLimit,
		*constraint.PercentLimit,
		*constraint.StringLimit,
		*constraint.BooleanLimit,
		*constraint.HexLimit,
		*constraint.ExpLimit,
		*constraint.StatusCodeLimit,
		*constraint.GenericLimit,
		*constraint.TemperatureLimit:
		renderLimit(parent, valueElement, name, limit.DataModelString(dataType))
	case *constraint.MathExpressionLimit:
		compute := parent.CreateElement("compute")
		operation := compute.CreateElement("operation")
		switch limit.Operand {
		case "+":
			operation.SetText("add")
		case "-":
			operation.SetText("subtract")
		case "*":
			operation.SetText("multiply")
		case "/":
			operation.SetText("divide")
		default:
			err = fmt.Errorf("unknown math expression operand: %s", limit.Operand)
			return
		}
		left := compute.CreateElement("left")
		err = renderConstraintLimit(left, left, limit.Left, dataType, "value", parentEntity)
		if err != nil {
			return
		}
		right := compute.CreateElement("right")
		err = renderConstraintLimit(right, right, limit.Right, dataType, "value", parentEntity)
		if err != nil {
			return
		}
	case *constraint.EmptyLimit:
		renderLimit(parent, valueElement, name, "empty")
	case *constraint.UnspecifiedLimit:
	case *constraint.ManufacturerLimit:
		renderLimit(parent, valueElement, name, "MS")
	default:
		err = fmt.Errorf("unknown constraint limit type: %T", limit)
	}
	return
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
	renderConstraintLimit(el, el, limit.Left, dataType, name, parentEntity)
	for _, r := range limit.Right {
		renderConstraintLimit(el, el, r, dataType, name, parentEntity)
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
		slog.Info("unexpected nil constraint limit", log.Path("source", entity), "entity", entity, "limit", limit, "dataType", dataType, "field", field)
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
	if hasCluster, ok := parentEntity.(interface{ Cluster() *matter.Cluster }); ok {
		entityCluster := entity.Cluster()
		parentCluster := hasCluster.Cluster()
		if entityCluster != nil && parentCluster != nil && entityCluster != parentCluster {
			parent.CreateAttr("cluster", entityCluster.Name)
		}
	}
	if field != nil {
		renderConstraintLimit(parent, parent, field, nil, "value", entity.Type.Entity)
	}
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
