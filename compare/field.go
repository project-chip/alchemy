package compare

import (
	"fmt"

	"github.com/hasty/alchemy/matter"
)

type FieldDiffs struct {
	ID    *matter.Number `json:"id"`
	Diffs []any          `json:"diffs"`
}

type MissingDiff struct {
	Type     DiffType       `json:"type"`
	Source   Source         `json:"source,omitempty"`
	Property DiffProperty   `json:"property,omitempty"`
	ID       *matter.Number `json:"id,omitempty"`
	Code     string         `json:"code,omitempty"`
}

type StringDiff struct {
	Type     DiffType     `json:"type"`
	Property DiffProperty `json:"property"`
	Spec     string       `json:"spec"`
	ZAP      string       `json:"zap"`
}

type BoolDiff struct {
	Type     DiffType     `json:"type"`
	Property DiffProperty `json:"property"`
	Spec     bool         `json:"spec"`
	ZAP      bool         `json:"zap"`
}

type Missing struct {
	Type     DiffType     `json:"type"`
	Property DiffProperty `json:"property"`
	Spec     string       `json:"spec"`
	ZAP      string       `json:"zap"`
}

type ConstraintDiff struct {
	Type     DiffType          `json:"type"`
	Property DiffProperty      `json:"property"`
	Spec     matter.Constraint `json:"spec"`
	ZAP      matter.Constraint `json:"zap"`
}

type QualityDiff struct {
	Type     DiffType       `json:"type"`
	Property DiffProperty   `json:"property"`
	Spec     matter.Quality `json:"spec"`
	ZAP      matter.Quality `json:"zap"`
}

type AccessDiff struct {
	Type     DiffType      `json:"type"`
	Property DiffProperty  `json:"property"`
	Spec     matter.Access `json:"spec"`
	ZAP      matter.Access `json:"zap"`
}

func compareField(specField *matter.Field, zapField *matter.Field) (diffs []any) {
	if specField.Name != zapField.Name {
		diffs = append(diffs, &StringDiff{Type: DiffTypeMismatch, Property: DiffPropertyName, Spec: specField.Name, ZAP: zapField.Name})
	}
	if specField.Type == nil && zapField.Type != nil {
		diffs = append(diffs, &MissingDiff{Type: DiffTypeMissing, Property: DiffPropertyType, Source: SourceSpec})

	} else if specField.Type != nil && zapField.Type == nil {
		diffs = append(diffs, &MissingDiff{Type: DiffTypeMissing, Property: DiffPropertyType, Source: SourceZAP})

	} else if specField.Type != nil && zapField.Type != nil {
		if specField.Type.Name != zapField.Type.Name {
			diffs = append(diffs, &StringDiff{Type: DiffTypeMismatch, Property: DiffPropertyType, Spec: specField.Type.Name, ZAP: zapField.Type.Name})
		}
		if specField.Type.IsArray() != zapField.Type.IsArray() {
			diffs = append(diffs, &BoolDiff{Type: DiffTypeMismatch, Property: DiffPropertyIsArray, Spec: specField.Type.IsArray(), ZAP: zapField.Type.IsArray()})
		}
	}
	if specField.Constraint == nil && zapField.Constraint != nil {
		diffs = append(diffs, &MissingDiff{Type: DiffTypeMissing, Property: DiffPropertyConstraint, Source: SourceSpec})
	} else if specField.Constraint != nil && zapField.Constraint == nil {
		diffs = append(diffs, &MissingDiff{Type: DiffTypeMissing, Property: DiffPropertyConstraint, Source: SourceZAP})
	} else if specField.Constraint != nil && zapField.Constraint != nil && !specField.Constraint.Equal(zapField.Constraint) {
		diffs = append(diffs, &ConstraintDiff{Type: DiffTypeMismatch, Property: DiffPropertyConstraint, Spec: specField.Constraint, ZAP: zapField.Constraint})
	}
	if specField.Quality != zapField.Quality {
		if (specField.Quality.Has(matter.QualityNullable) && !zapField.Quality.Has(matter.QualityNullable)) || (!specField.Quality.Has(matter.QualityNullable) && zapField.Quality.Has(matter.QualityNullable)) {
			diffs = append(diffs, &QualityDiff{Type: DiffTypeMismatch, Property: DiffPropertyQuality, Spec: specField.Quality, ZAP: zapField.Quality})
		}
	}
	if !specField.Access.Equal(zapField.Access) {
		if specField.Access.Read != matter.PrivilegeView && specField.Access.Write != matter.PrivilegeUnknown {
			diffs = append(diffs, &AccessDiff{Type: DiffTypeMismatch, Property: DiffPropertyAccess, Spec: specField.Access, ZAP: zapField.Access})
		}
	}
	if specField.Default != zapField.Default {
		diffs = append(diffs, &StringDiff{Type: DiffTypeMismatch, Property: DiffPropertyDefault, Spec: specField.Default, ZAP: zapField.Default})
	}
	if specField.Conformance != zapField.Conformance {
		diffs = append(diffs, &StringDiff{Type: DiffTypeMismatch, Property: DiffPropertyConformance, Spec: specField.Conformance.String(), ZAP: zapField.Conformance.String()})
	}
	return
}

func compareFields(specFields []*matter.Field, zapFields []*matter.Field) (diffs []any, err error) {
	specFieldMap := make(map[uint64]*matter.Field)
	for _, f := range specFields {
		if !f.ID.Valid() {
			err = fmt.Errorf("unable to parse spec field ID: %s; %w", f.ID.IntString(), err)
		}
		specFieldMap[f.ID.Value()] = f
	}

	zapFieldMap := make(map[uint64]*matter.Field)
	for _, f := range zapFields {
		if !f.ID.Valid() {
			err = fmt.Errorf("unable to parse ZAP field ID: %s; %w", f.ID.IntString(), err)
		}
		zapFieldMap[f.ID.Value()] = f
	}

	for code, zapField := range zapFieldMap {
		specField, ok := specFieldMap[code]
		if !ok {
			continue
		}
		delete(zapFieldMap, code)
		delete(specFieldMap, code)
		fieldDiffs := compareField(specField, zapField)
		if len(fieldDiffs) > 0 {
			diffs = append(diffs, &FieldDiffs{ID: zapField.ID, Diffs: fieldDiffs})
		}
	}
	for _, f := range specFieldMap {
		diffs = append(diffs, &MissingDiff{Type: DiffTypeMissing, ID: f.ID, Source: SourceZAP})
	}
	for _, f := range zapFieldMap {
		diffs = append(diffs, &MissingDiff{Type: DiffTypeMissing, ID: f.ID, Source: SourceSpec})
	}
	return
}
