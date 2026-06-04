package idl

import (
	"slices"

	"github.com/mailgun/raymond/v2"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
	"github.com/project-chip/alchemy/zap"
)

func fieldTypeHelper(field matter.Field, fs matter.FieldSet, options *raymond.Options) raymond.SafeString {
	return raymond.SafeString(zap.FieldToZapDataType(fs, &field, field.Constraint))
}

func fieldIsArrayHelper(a any, options *raymond.Options) string {
	var t *types.DataType
	switch a := a.(type) {
	case types.DataType:
		t = &a
	case *types.DataType:
		t = a
	default:
		return options.Inverse()
	}
	if t == nil {
		return options.Inverse()
	}
	if t.IsArray() {
		return options.Fn()
	} else {
		return options.Inverse()
	}
}

func ifHasValidFieldsHelper(spec *spec.Specification, filter ProvisionalFilter) func(fs matter.FieldSet, options *raymond.Options) string {
	return func(fs matter.FieldSet, options *raymond.Options) string {
		if len(filterEntities(spec, filter, fs)) > 0 {
			return options.Fn()
		} else {
			return options.Inverse()
		}
	}
}

func structFieldsHelper(spec *spec.Specification, filter ProvisionalFilter) func(s matter.Struct, options *raymond.Options) raymond.SafeString {
	return func(s matter.Struct, options *raymond.Options) raymond.SafeString {
		fields := filterEntities(spec, filter, s.Fields)
		if s.FabricScoping == matter.FabricScopingScoped {
			fabricIndex := &matter.Field{ID: matter.NewNumber(254), Name: "FabricIndex", Type: types.NewDataType(types.BaseDataTypeFabricIndex, types.DataTypeRankScalar), Conformance: conformance.Set{&conformance.Mandatory{}}}
			fabricIndex.SetParent(&s)
			fields = append(fields, fabricIndex)
		}
		slices.SortStableFunc(fields, func(a *matter.Field, b *matter.Field) int {
			return a.ID.Compare(b.ID)
		})
		return enumerateEntitiesHelper(fields, spec, filter, options)
	}
}

func eventFieldsHelper(spec *spec.Specification, filter ProvisionalFilter) func(e matter.Event, options *raymond.Options) raymond.SafeString {
	return func(e matter.Event, options *raymond.Options) raymond.SafeString {
		fields := filterEntities(spec, filter, e.Fields)
		if e.Access.FabricSensitivity == matter.FabricSensitivitySensitive {
			fabricIndex := &matter.Field{ID: matter.NewNumber(254), Name: "FabricIndex", Type: types.NewDataType(types.BaseDataTypeFabricIndex, types.DataTypeRankScalar), Conformance: conformance.Set{&conformance.Mandatory{}}}
			fabricIndex.SetParent(&e)
			fields = append(fields, fabricIndex)
		}
		slices.SortStableFunc(fields, func(a *matter.Field, b *matter.Field) int {
			return a.ID.Compare(b.ID)
		})
		return enumerateEntitiesHelper(fields, spec, filter, options)
	}
}
