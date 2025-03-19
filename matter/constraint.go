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

	visitedReferences map[types.Entity]struct{}
}

func (cc *ConstraintContext) Nullable() bool {
	if cc.Field != nil {
		return cc.Field.Quality.Has(QualityNullable)
	}
	return false
}

func (cc *ConstraintContext) DataType() *types.DataType {
	if cc.Field != nil {
		return cc.Field.Type
	}
	return nil
}

func (cc *ConstraintContext) referenceAlreadyVisited(entity types.Entity) bool {
	if cc.visitedReferences == nil {
		cc.visitedReferences = make(map[types.Entity]struct{})
	} else if _, ok := cc.visitedReferences[entity]; ok {
		switch e := entity.(type) {
		case *Field:
			slog.Warn("Visited field twice when resolving constraint", slog.String("name", e.Name), log.Path("source", e))
		case *EnumValue:
			slog.Warn("Visited enum value twice when resolving constraint", slog.String("name", e.Name), log.Path("source", e))
		default:
			slog.Warn("Visited entity twice when resolving constraint", log.Type("entity", entity))

		}
		return true
	}
	cc.visitedReferences[entity] = struct{}{}
	return false
}

func (cc *ConstraintContext) MinEntityValue(entity types.Entity, field constraint.Limit) (min types.DataTypeExtreme) {
	if cc.referenceAlreadyVisited(entity) {
		return
	}
	switch entity := entity.(type) {
	case *Field:

		min = entity.Constraint.Min(&ConstraintContext{Field: entity, Fields: cc.Fields, visitedReferences: cc.visitedReferences})
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
	default:
		slog.Warn("Unexpected entity type on MinEntityValue", log.Type("entity", entity))
	}
	return
}

func (cc *ConstraintContext) MaxEntityValue(entity types.Entity, field constraint.Limit) (max types.DataTypeExtreme) {
	if cc.referenceAlreadyVisited(entity) {
		return
	}
	switch entity := entity.(type) {
	case *Field:
		max = entity.Constraint.Max(&ConstraintContext{Field: entity, Fields: cc.Fields, visitedReferences: cc.visitedReferences})
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
	default:
		slog.Warn("Unexpected entity type on MaxEntityValue", log.Type("entity", entity))
	}
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
