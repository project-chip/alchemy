package dm

import (
	"fmt"
	"log/slog"
	"strconv"

	"github.com/beevik/etree"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/constraint"
	"github.com/project-chip/alchemy/matter/types"
)

type identifierStore interface {
	Identifier(name string) (types.Entity, bool)
}

func renderConstraint(con constraint.Constraint, dataType *types.DataType, parent *etree.Element) error {
	if con == nil {
		return nil
	}
	err := renderConstraintElement(con, dataType, parent)
	if err != nil {
		return fmt.Errorf("error rendering constraint element %s: %w", con.ASCIIDocString(dataType), err)
	}
	return nil
}

func renderConstraintElement(con constraint.Constraint, dataType *types.DataType, parent *etree.Element) (err error) {
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
		cx.CreateAttr("type", "desc")
		constraints = append(constraints, cx)
	case *constraint.ExactConstraint:
		cx := etree.NewElement(name)
		cx.CreateAttr("type", "allowed")
		renderConstraintLimit(cx, con.Value, dataType, "value")
		constraints = append(constraints, cx)
	case *constraint.RangeConstraint:
		cx := etree.NewElement(name)
		if dataType.IsArray() {
			cx.CreateAttr("type", "countBetween")
		} else if dataType.HasLength() {
			cx.CreateAttr("type", "lengthBetween")
		} else {
			cx.CreateAttr("type", "between")
		}
		fromRef, fromEntity, fromField := getLimitField(con.Minimum)
		toRef, toEntity, toField := getLimitField(con.Maximum)
		if fromEntity == nil && fromField == nil && toField == nil && toEntity == nil {
			cx.CreateAttr("from", renderLimitValue(con.Minimum, dataType))
			cx.CreateAttr("to", renderLimitValue(con.Maximum, dataType))
		} else {
			renderConstraintReferenceLimit(cx, fromField, con.Minimum, dataType, "from", fromEntity, fromRef)
			renderConstraintReferenceLimit(cx, toField, con.Maximum, dataType, "to", toEntity, toRef)
		}
		constraints = append(constraints, cx)
	case *constraint.MinConstraint:
		cx := etree.NewElement(name)
		if dataType.IsArray() {
			cx.CreateAttr("type", "minCount")
		} else if dataType.HasLength() {
			cx.CreateAttr("type", "minLength")
		} else {
			cx.CreateAttr("type", "min")
		}
		renderConstraintLimit(cx, con.Minimum, dataType, "value")
		constraints = append(constraints, cx)
	case *constraint.MaxConstraint:
		cx := etree.NewElement(name)
		if dataType.IsArray() {
			cx.CreateAttr("type", "maxCount")
		} else if dataType.HasLength() {
			cx.CreateAttr("type", "maxLength")
		} else {
			cx.CreateAttr("type", "max")
		}
		renderConstraintLimit(cx, con.Maximum, dataType, "value")
		constraints = append(constraints, cx)
		if characterLimit, ok := con.Maximum.(*constraint.CharacterLimit); ok {
			cx := parent.CreateElement(name)
			cx.CreateAttr("type", "maxCodePoints")
			renderConstraintLimit(cx, characterLimit.CodepointCount, dataType, "value")
			constraints = append(constraints, cx)
		}
	case *constraint.ListConstraint:
		if mc, ok := con.Constraint.(*constraint.MaxConstraint); ok {
			cx := etree.NewElement(name)
			cx.CreateAttr("type", "maxCount")
			renderConstraintLimit(cx, mc.Maximum, dataType, "value")
			constraints = append(constraints, cx)
		}
	case constraint.Set:
		for _, cs := range con {
			err = renderConstraintElement(cs, dataType, parent)
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

func renderConstraintLimit(parent *etree.Element, limit constraint.Limit, dataType *types.DataType, name string) {
	ref, entity, field := getLimitField(limit)
	if entity == nil && field == nil {
		parent.CreateAttr(name, renderLimitValue(limit, dataType))
		return
	}
	renderConstraintReferenceLimit(parent, field, limit, dataType, name, entity, ref)
}

func renderConstraintReferenceLimit(parent *etree.Element, field constraint.Limit, limit constraint.Limit, dataType *types.DataType, name string, entity types.Entity, ref string) bool {
	switch entity := entity.(type) {
	case *matter.EnumValue:
		parent.CreateAttr(name, entity.Name)
		return true
	case matter.Bit:
		mask, err := entity.Mask()
		if err == nil {
			parent.CreateAttr(name, strconv.FormatUint(mask, 10))
		}
		return true
	case *matter.Field:
		el := parent.CreateElement(name)
		switch entity.EntityType() {
		case types.EntityTypeAttribute:
			el.CreateAttr("attribute", entity.Name)
		default:
			slog.Info("Unexpected entity type on reference limit", "entity", entity)
			el.CreateAttr("reference", ref)
		}
		if field != nil {
			switch field := field.(type) {
			case *constraint.IdentifierLimit:
				el.CreateAttr("field", field.ID)
			}
		}
	case nil:
		el := parent.CreateElement(name)
		el.CreateAttr("value", renderLimitValue(limit, dataType))
	default:
		slog.Info("unexpected constraint limit", "entity", entity, "limit", limit, "dataType", dataType, "field", field)
	}
	return false
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
