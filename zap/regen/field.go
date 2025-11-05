package regen

import (
	"github.com/mailgun/raymond/v2"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
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

func filterFields(fieldSets ...matter.FieldSet) (fields matter.FieldSet) {
	count := 0
	for _, fieldSet := range fieldSets {
		count += len(fieldSet)
	}
	fields = make(matter.FieldSet, 0, count)
	for _, fieldSet := range fieldSets {
		for _, f := range fieldSet {
			if conformance.IsZigbee(f.Conformance) || zap.IsDisallowed(f, f.Conformance) {
				continue
			}
			fields = append(fields, f)
		}
	}
	return
}
