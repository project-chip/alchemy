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

	visitedReferences map[string]struct{}
}

func (cc *ConstraintContext) DataType() *types.DataType {
	if cc.Field != nil {
		return cc.Field.Type
	}
	return nil
}

func (cc *ConstraintContext) getReference(ref string) *Field {
	r := cc.Fields.Get(ref)
	if cc.visitedReferences == nil {
		cc.visitedReferences = make(map[string]struct{})
	} else if _, ok := cc.visitedReferences[ref]; ok {
		return nil
	}
	cc.visitedReferences[ref] = struct{}{}
	return r
}

func (cc *ConstraintContext) getReferencedField(ref string, field constraint.Limit) *Field {
	f := cc.getReference(ref)
	if f == nil {
		slog.Warn("Unknown reference when fetching constraint", slog.String("reference", ref))
		return nil
	}
	if field == nil {
		return f
	}
	if f.Type == nil {
		slog.Warn("Referenced constraint field has no type information for child field", log.Path("source", f), slog.String("reference", ref), slog.Any("field", field))
		return nil
	}
	var fieldSet FieldSet
	switch e := f.Type.Entity.(type) {
	case *Struct:
		fieldSet = e.Fields
	case *Command:
		fieldSet = e.Fields
	case *Event:
		fieldSet = e.Fields
	default:
		slog.Warn("referenced constraint field has a type without fields", log.Path("source", f), slog.String("reference", ref), slog.Any("field", field), log.Type("type", e))
		return nil
	}
	childField := fieldSet.Get(ref)
	ccc := &ConstraintContext{Field: childField, Fields: fieldSet}
	switch f := field.(type) {
	case *constraint.IdentifierLimit:
		return ccc.getReferencedField(f.ID, f.Field)
	case *constraint.ReferenceLimit:
		return ccc.getReferencedField(f.Reference, f.Field)
	}
	slog.Warn("referenced constraint field has unrecognized type", log.Path("source", f), slog.String("reference", ref), slog.Any("field", field), log.Type("type", field))
	return nil
}

func (cc *ConstraintContext) MinEntityValue(entity types.Entity, field constraint.Limit) (min types.DataTypeExtreme) {
	switch entity := entity.(type) {
	case *Field:
		min = entity.Constraint.Min(cc)
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
	switch entity := entity.(type) {
	case *Field:
		max = entity.Constraint.Max(cc)
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
