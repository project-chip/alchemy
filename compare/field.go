package compare

import (
	"log/slog"
	"strings"

	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/conformance"
	"github.com/hasty/alchemy/matter/types"
	"github.com/hasty/alchemy/zap"
)

type FieldDiffs struct {
	ID    *matter.Number `json:"id"`
	Name  string         `json:"name"`
	Diffs []any          `json:"diffs"`
}

func compareFieldTypes(specFieldName string, specFieldType *types.DataType, zapFieldName string, zapFieldType *types.DataType) (diffs []any) {
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
		diffs = append(diffs, compareFieldTypes(specFieldName, specFieldType.EntryType, zapFieldName, zapFieldType.EntryType)...)
		return
	}
	if specFieldType.BaseType != zapFieldType.BaseType {
		diffs = append(diffs, &StringDiff{Type: DiffTypeMismatch, Property: DiffPropertyType, Spec: specFieldType.Name, ZAP: zapFieldType.Name})
		return
	}
	if specFieldType.BaseType == types.BaseDataTypeCustom {
		specFieldTypeName := specFieldType.Name
		zapFieldTypeName := zapFieldType.Name
		if !strings.HasPrefix(specFieldTypeName, zapFieldTypeName) {
			diffs = append(diffs, &StringDiff{Type: DiffTypeMismatch, Property: DiffPropertyType, Spec: specFieldTypeName, ZAP: zapFieldTypeName})
		}
	}
	return
}

func compareField(specFields matter.FieldSet, specField *matter.Field, zapField *matter.Field) (diffs []any) {
	if !namesEqual(specField.Name, zapField.Name) {
		diffs = append(diffs, &StringDiff{Type: DiffTypeMismatch, Property: DiffPropertyName, Spec: specField.Name, ZAP: zapField.Name})
	}
	diffs = append(diffs, compareFieldTypes(specField.Name, specField.Type, zapField.Name, zapField.Type)...)

	if specField.Quality != zapField.Quality {
		if (specField.Quality.Has(matter.QualityNullable) && !zapField.Quality.Has(matter.QualityNullable)) || (!specField.Quality.Has(matter.QualityNullable) && zapField.Quality.Has(matter.QualityNullable)) {
			diffs = append(diffs, &BoolDiff{Type: DiffTypeMismatch, Property: DiffPropertyNullable, Spec: specField.Quality.Has(matter.QualityNullable), ZAP: zapField.Quality.Has(matter.QualityNullable)})
		}
	}
	if !specField.Access.Equal(zapField.Access) {
		if specField.Access.Read != matter.PrivilegeView && specField.Access.Write != matter.PrivilegeUnknown {
			diffs = append(diffs, &AccessDiff{Type: DiffTypeMismatch, Property: DiffPropertyAccess, Spec: specField.Access, ZAP: zapField.Access})
		}
	}
	defaultValue := zap.GetDefaultValue(&matter.ConstraintContext{Field: specField, Fields: specFields})
	if defaultValue.Defined() {
		specDefault := defaultValue.ZapString(specField.Type)
		if specDefault != zapField.Default && !(specField.Default == "null" && len(zapField.Default) == 0) { // ZAP frequently omits default null
			specDefaultVal := matter.ParseNumber(specDefault)
			zapDefault := matter.ParseNumber(zapField.Default)
			if !specDefaultVal.Equals(zapDefault) {
				diffs = append(diffs, &StringDiff{Type: DiffTypeMismatch, Property: DiffPropertyDefault, Spec: specDefault, ZAP: zapField.Default})
			}
		}
	} else if len(zapField.Default) > 0 {
		if len(specField.Default) > 0 {
			diffs = append(diffs, &StringDiff{Type: DiffTypeMismatch, Property: DiffPropertyDefault, Spec: specField.Default, ZAP: zapField.Default})
		} else {
			diffs = append(diffs, newMissingDiff(specField.Name, DiffPropertyDefault, types.EntityTypeField, SourceSpec))
		}
	}
	diffs = append(diffs, compareConformance(specField.Conformance, zapField.Conformance)...)
	return
}

func compareFields(specFields matter.FieldSet, zapFields matter.FieldSet) (diffs []any, err error) {
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
				delete(specFieldMap, specField.ID.Value())
				delete(specFieldNameMap, strings.ToLower(zapField.Name))
			}
		}
		if !ok {
			continue
		}
		fieldDiffs := compareField(specFields, specField, zapField)
		if len(fieldDiffs) > 0 {
			diffs = append(diffs, &FieldDiffs{ID: specField.ID, Name: specField.Name, Diffs: fieldDiffs})
		}
	}

	for _, f := range specFieldMap {
		diffs = append(diffs, newMissingDiff(f.Name, f.ID, SourceZAP))
		delete(specFieldNameMap, strings.ToLower(f.Name))
	}
	for _, f := range zapFieldMap {
		diffs = append(diffs, newMissingDiff(f.Name, f.ID, SourceSpec))
	}
	for _, f := range specFieldNameMap {
		diffs = append(diffs, newMissingDiff(f.Name, f.ID, SourceZAP))
	}
	return
}
