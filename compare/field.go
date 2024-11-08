package compare

import (
	"log/slog"
	"strings"

	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/types"
	"github.com/project-chip/alchemy/zap"
)

func compareFieldTypes(specFieldSet matter.FieldSet, specField *matter.Field, specFieldName string, specFieldType *types.DataType, zapFieldSet matter.FieldSet, zapField *matter.Field, zapFieldName string, zapFieldType *types.DataType) (diffs []Diff) {
	if specFieldType == nil && zapFieldType != nil {
		diffs = append(diffs, newMissingDiff(zapFieldName, DiffPropertyType, SourceSpec))
		return
	}
	if specFieldType != nil && zapFieldType == nil {
		diffs = append(diffs, newMissingDiff(specFieldName, DiffPropertyType, SourceZAP))
		return
	}
	if specFieldType.IsArray() != zapFieldType.IsArray() {
		diffs = append(diffs, &BoolDiff{Type: DiffTypeMismatch, Property: DiffPropertyIsArray, Spec: specFieldType.IsArray(), ZAP: zapFieldType.IsArray()})
		return
	}
	if specFieldType.IsArray() {
		diffs = append(diffs, compareFieldTypes(specFieldSet, specField, specFieldName, specFieldType.EntryType, zapFieldSet, zapField, zapFieldName, zapFieldType.EntryType)...)
		return
	}
	specFieldTypeName := zap.FieldToZapDataType(specFieldSet, specField)
	var zapFieldTypeName string
	if zapField.Type.IsArray() {
		zapFieldTypeName = zapField.Type.EntryType.Name
	} else {
		zapFieldTypeName = zapField.Type.Name
	}
	switch specFieldType.BaseType {
	case types.BaseDataTypeCustom:
		if !strings.HasPrefix(specFieldTypeName, zapFieldTypeName) {
			diffs = append(diffs, &StringDiff{Type: DiffTypeMismatch, Property: DiffPropertyType, Spec: specFieldTypeName, ZAP: zapFieldTypeName})
		}
	default:
		if !strings.EqualFold(specFieldTypeName, zapFieldTypeName) {
			diffs = append(diffs, &StringDiff{Type: DiffTypeMismatch, Property: DiffPropertyType, Spec: specFieldTypeName, ZAP: zapFieldTypeName})
		}
	}
	return
}

func compareField(entityType types.EntityType, specFields matter.FieldSet, specField *matter.Field, zapFields matter.FieldSet, zapField *matter.Field) (diffs []Diff) {
	if !namesEqual(specField.Name, zapField.Name) {
		diffs = append(diffs, &StringDiff{Type: DiffTypeMismatch, Property: DiffPropertyName, Spec: specField.Name, ZAP: zapField.Name})
	}
	diffs = append(diffs, compareFieldTypes(specFields, specField, specField.Name, specField.Type, zapFields, zapField, zapField.Name, zapField.Type)...)

	if specField.Quality != zapField.Quality {
		if (specField.Quality.Has(matter.QualityNullable) && !zapField.Quality.Has(matter.QualityNullable)) || (!specField.Quality.Has(matter.QualityNullable) && zapField.Quality.Has(matter.QualityNullable)) {
			diffs = append(diffs, &BoolDiff{Type: DiffTypeMismatch, Property: DiffPropertyNullable, Spec: specField.Quality.Has(matter.QualityNullable), ZAP: zapField.Quality.Has(matter.QualityNullable)})
		}
	}
	diffs = append(diffs, compareAccess(entityType, specField.Access, zapField.Access)...)
	defaultValue := zap.GetFallbackValue(&matter.ConstraintContext{Field: specField, Fields: specFields})
	if defaultValue.Defined() {
		specDefault := defaultValue.ZapString(specField.Type)
		if specDefault != zapField.Fallback && !(specField.Fallback == "null" && len(zapField.Fallback) == 0) { // ZAP frequently omits default null
			specDefaultVal := matter.ParseNumber(specDefault)
			zapDefault := matter.ParseNumber(zapField.Fallback)
			if !specDefaultVal.Equals(zapDefault) {
				diffs = append(diffs, &StringDiff{Type: DiffTypeMismatch, Property: DiffPropertyDefault, Spec: specDefault, ZAP: zapField.Fallback})
			}
		}

	} else if len(zapField.Fallback) > 0 {

		if len(specField.Fallback) > 0 {
			z := matter.ParseNumber(zapField.Fallback)
			if specField.Fallback != "null" || !z.Equals(matter.NewNumber(specField.Type.NullValue())) {
				s := matter.ParseNumber(specField.Fallback)
				if !z.Equals(s) {
					diffs = append(diffs, &StringDiff{Type: DiffTypeMismatch, Property: DiffPropertyDefault, Spec: specField.Fallback, ZAP: zapField.Fallback})
				}
			}
		} else {
			z := matter.ParseNumber(zapField.Fallback)
			if !z.Valid() || z.Value() != 0 {
				diffs = append(diffs, newMissingDiff(specField.Name, DiffPropertyDefault, entityType, SourceSpec))
			}
		}
	}
	diffs = append(diffs, compareConformance(entityType, specField.Conformance, zapField.Conformance)...)
	diffs = append(diffs, compareConstraint(entityType, specFields, specField, zapFields, zapField)...)
	return
}

func compareFields(entityType types.EntityType, specFields matter.FieldSet, zapFields matter.FieldSet) (diffs []Diff, err error) {
	specFieldMap := make(map[uint64]*matter.Field)
	specFieldNameMap := make(map[string]*matter.Field)
	for _, f := range specFields {
		if conformance.IsZigbee(specFields, f.Conformance) {
			continue
		}
		specFieldNameMap[strings.ToLower(f.Name)] = f
		if !f.ID.Valid() {
			slog.Warn("unable to parse spec field ID", slog.String("id", f.ID.Text()), slog.String("name", f.Name))
			continue
		}
		specFieldMap[f.ID.Value()] = f
	}

	zapFieldMap := make(map[uint64]*matter.Field)
	for _, f := range zapFields {
		if !f.ID.Valid() {
			continue
		}
		zapFieldMap[f.ID.Value()] = f
	}

	for _, zapField := range zapFields {
		var specField *matter.Field
		var ok bool
		if zapField.ID.Valid() {
			specField, ok = specFieldMap[zapField.ID.Value()]
			if ok {
				delete(zapFieldMap, zapField.ID.Value())
				delete(specFieldMap, zapField.ID.Value())
				delete(specFieldNameMap, strings.ToLower(specField.Name))
			}
		}
		if !ok {
			specField, ok = specFieldNameMap[strings.ToLower(zapField.Name)]
			if ok {
				if zapField.ID.Valid() {
					delete(zapFieldMap, zapField.ID.Value())
				}
				delete(specFieldMap, specField.ID.Value())
				delete(specFieldNameMap, strings.ToLower(zapField.Name))
			}
		}
		if !ok {
			continue
		}
		fieldDiffs := compareField(entityType, specFields, specField, zapFields, zapField)
		if len(fieldDiffs) > 0 {
			diffs = append(diffs, &IdentifiedDiff{Entity: entityType, ID: specField.ID, Name: specField.Name, Diffs: fieldDiffs})
		}
	}

	for _, f := range specFieldMap {
		if !conformance.IsDeprecated(f.Conformance) {
			diffs = append(diffs, newMissingDiff(f.Name, entityType, f.ID, SourceZAP))
		}
		delete(specFieldNameMap, strings.ToLower(f.Name))
	}
	for _, f := range zapFieldMap {
		diffs = append(diffs, newMissingDiff(f.Name, entityType, f.ID, SourceSpec))
	}
	for _, f := range specFieldNameMap {
		if !conformance.IsDeprecated(f.Conformance) {
			diffs = append(diffs, newMissingDiff(f.Name, entityType, f.ID, SourceZAP))
		}
	}
	return
}
