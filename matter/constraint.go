package matter

import (
	"log/slog"

	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter/constraint"
	"github.com/project-chip/alchemy/matter/types"
)

type ConstraintContext struct {
	Field  *Field
	Fields FieldSet

	visitedMinReferences map[types.Entity]struct{}
	visitedMaxReferences map[types.Entity]struct{}

	parent *ConstraintContext
}

func NewConstraintContext(field *Field, fields FieldSet) *ConstraintContext {
	return &ConstraintContext{
		Field:                field,
		Fields:               fields,
		visitedMinReferences: make(map[types.Entity]struct{}),
		visitedMaxReferences: make(map[types.Entity]struct{}),
	}
}

func (cc *ConstraintContext) Child(field *Field, fields FieldSet) *ConstraintContext {
	return &ConstraintContext{
		Field:                field,
		Fields:               cc.Fields,
		parent:               cc,
		visitedMinReferences: make(map[types.Entity]struct{}),
		visitedMaxReferences: make(map[types.Entity]struct{}),
	}
}

func (cc *ConstraintContext) Nullability() types.Nullability {
	if cc.Field != nil && cc.Field.Quality.Has(QualityNullable) {
		return types.NullabilityNullable
	}
	return types.NullabilityNonNull
}

func (cc *ConstraintContext) DataType() *types.DataType {
	if cc.Field != nil {
		return cc.Field.Type
	}
	return nil
}

func (cc *ConstraintContext) MinEntityValue(entity types.Entity, field constraint.Limit) (min types.DataTypeExtreme) {

	switch entity := entity.(type) {
	case *Field:
		child := cc.Child(entity, cc.Fields)
		if child.recordMinReference(entity) {
			slog.Error("Circular reference detected", LogEntity("entity", entity))
			return
		}
		min = entity.Constraint.Min(child)
	case Bit:
		from, _, err := entity.Bits()
		if err != nil {
			slog.Warn("Error getting minimum value from bitmap bit", slog.String("name", entity.Name()), slog.Any("error", err))
		} else {
			min = types.NewUintDataTypeExtreme(from, types.NumberFormatHex)
		}
	case *EnumValue:
		min = types.NewUintDataTypeExtreme(entity.Value.Value(), types.NumberFormatHex)
	case *Constant:
		min = getConstantValue(entity)
	case *TypeDef:
		min = types.Min(entity.Type.BaseType, cc.Nullability())
	default:
		slog.Warn("Unexpected entity type on MinEntityValue", log.Type("entity", entity))
	}
	return
}

func (cc *ConstraintContext) MaxEntityValue(entity types.Entity, field constraint.Limit) (max types.DataTypeExtreme) {

	switch entity := entity.(type) {
	case *Field:
		child := cc.Child(entity, cc.Fields)
		if child.recordMaxReference(entity) {
			slog.Error("Circular reference detected", LogEntity("entity", entity))
			return
		}
		max = entity.Constraint.Max(child)
	case Bit:
		_, to, err := entity.Bits()
		if err != nil {
			slog.Warn("Error getting maximum value from bitmap bit", slog.String("name", entity.Name()), slog.Any("error", err))
		} else {
			max = types.NewUintDataTypeExtreme(to, types.NumberFormatHex)
		}
	case *EnumValue:
		max = types.NewUintDataTypeExtreme(entity.Value.Value(), types.NumberFormatHex)
	case *Constant:
		max = getConstantValue(entity)
	case *TypeDef:
		slog.Warn("MaxEntityValue on type definition")
		max = types.Max(entity.Type.BaseType, cc.Nullability())
	default:
		slog.Warn("Unexpected entity type on MaxEntityValue", log.Type("entity", entity))
	}
	return
}

func (cc *ConstraintContext) recordMinReference(entity types.Entity) (previouslyVisited bool) {
	c := cc
	for {
		if c == nil {
			break
		}
		if _, previouslyVisited = c.visitedMinReferences[entity]; previouslyVisited {
			return
		}
		c = c.parent
	}
	cc.visitedMinReferences[entity] = struct{}{}
	return
}

func (cc *ConstraintContext) recordMaxReference(entity types.Entity) (previouslyVisited bool) {
	c := cc
	for {
		if c == nil {
			break
		}
		if _, previouslyVisited = c.visitedMaxReferences[entity]; previouslyVisited {
			return
		}
		c = c.parent
	}
	cc.visitedMaxReferences[entity] = struct{}{}
	return
}

func getConstantValue(c *Constant) (val types.DataTypeExtreme) {
	switch v := c.Value.(type) {
	case uint64:
		val = types.NewUintDataTypeExtreme(v, types.NumberFormatHex)
	case uint:
		val = types.NewUintDataTypeExtreme(uint64(v), types.NumberFormatHex)
	case int64:
		val = types.NewIntDataTypeExtreme(v, types.NumberFormatHex)
	case int:
		val = types.NewIntDataTypeExtreme(int64(v), types.NumberFormatHex)
	default:
	}
	return
}

func (cc *ConstraintContext) Fallback(entity types.Entity, field constraint.Limit) (def types.DataTypeExtreme) {
	switch entity := entity.(type) {
	case *Field:
		if !constraint.IsGenericLimit(entity.Fallback) {
			def = entity.Fallback.Fallback(cc)
			return
		}
	case *EnumValue:
		if entity.Value.Valid() {
			def = types.NewUintDataTypeExtreme(entity.Value.Value(), types.NumberFormatInt)
			def.Entity = entity
			return
		}
	case Bit:
		val, err := entity.Mask()
		if err != nil {
			def = types.NewUintDataTypeExtreme(val, types.NumberFormatInt)
			def.Entity = entity
			return
		}
	case nil:
	default:
		slog.Warn("Unexpected entity type on Fallback", log.Type("entity", entity))
	}

	return
}

func EntityConstraint(entity types.Entity) constraint.Constraint {
	switch entity := entity.(type) {
	case *Field:
		return entity.Constraint
	case *DeviceTypeRequirement:
		return entity.Constraint
	case *ElementRequirement:
		return entity.Constraint
	case nil:
		slog.Warn("Unexpected nil entity fetching constraint")
	default:
		slog.Warn("Unexpected entity fetching constraint", LogEntity("entity", entity))
	}
	return nil
}

func EntityFallback(entity types.Entity) constraint.Limit {
	switch entity := entity.(type) {
	case *Field:
		return entity.Fallback
	case nil:
		slog.Warn("Unexpected nil entity fetching fallback")
	default:
		slog.Warn("Unexpected entity fetching fallback", LogEntity("entity", entity))
	}
	return nil
}
