package regen

import (
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

func ifHasValidFieldsHelper(fs matter.FieldSet, options *raymond.Options) string {
	if len(filterEntities(fs)) > 0 {
		return options.Fn()
	} else {
		return options.Inverse()
	}
}

func structFieldsHelper(spec *spec.Specification) func(s matter.Struct, options *raymond.Options) raymond.SafeString {
	return func(s matter.Struct, options *raymond.Options) raymond.SafeString {
		fields := filterEntities(s.Fields)
		if s.FabricScoping == matter.FabricScopingScoped {
			fabricIndex := &matter.Field{ID: matter.NewNumber(254), Name: "FabricIndex", Type: types.NewDataType(types.BaseDataTypeFabricIndex, false), Conformance: conformance.Set{&conformance.Mandatory{}}}
			fabricIndex.SetParent(&s)
			fields = append(fields, fabricIndex)
		}
		return enumerateEntitiesHelper(fields, spec, options)
	}
}

func eventFieldsHelper(spec *spec.Specification) func(e matter.Event, options *raymond.Options) raymond.SafeString {
	return func(e matter.Event, options *raymond.Options) raymond.SafeString {
		fields := filterEntities(e.Fields)
		if e.Access.FabricSensitivity == matter.FabricSensitivitySensitive {
			fields = append(fields, &matter.Field{ID: matter.NewNumber(254), Name: "FabricIndex", Type: types.NewDataType(types.BaseDataTypeFabricIndex, false), Conformance: conformance.Set{&conformance.Mandatory{}}})
		}
		return enumerateEntitiesHelper(fields, spec, options)
	}
}
