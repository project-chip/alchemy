package compare

import (
	"strings"

	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/conformance"
	mattertypes "github.com/hasty/alchemy/matter/types"
	"github.com/hasty/alchemy/zap"
)

func compareConformance(entityType mattertypes.EntityType, spec conformance.Set, zap conformance.Set) (diffs []Diff) {
	if len(spec) == 0 {
		if len(zap) > 0 {
			diffs = append(diffs, newMissingDiff("", entityType, DiffPropertyConformance, SourceSpec))
		}
		return
	} else if len(zap) == 0 {
		diffs = append(diffs, newMissingDiff("", entityType, DiffPropertyConformance, SourceZAP))
		return
	}

	var specState = conformance.StateOptional
	var zapState = conformance.StateOptional
	if conformance.IsMandatory(spec) {
		specState = conformance.StateMandatory
	}
	if conformance.IsMandatory(zap) {
		zapState = conformance.StateMandatory
	}

	if specState != zapState {
		diffs = append(diffs, &ConformanceDiff{Type: DiffTypeMismatch, Property: DiffPropertyConformance, Spec: specState, ZAP: zapState})
	}

	return
}
func compareConstraint(entityType mattertypes.EntityType, specFieldSet matter.FieldSet, specField *matter.Field, zapFieldSet matter.FieldSet, zapField *matter.Field) (diffs []Diff) {
	if specField.Constraint == nil && zapField.Constraint == nil {
		return
	}

	var maxProp = DiffPropertyMax
	var minProp = DiffPropertyMin
	if specField.Type != nil && (specField.Type.HasLength() || specField.Type.IsArray()) {
		maxProp = DiffPropertyLength
		minProp = DiffPropertyMinLength
	}

	specFrom, specTo := zap.GetMinMax(&matter.ConstraintContext{Field: specField, Fields: specFieldSet}, specField.Constraint)
	zapFrom, zapTo := zap.GetMinMax(&matter.ConstraintContext{Field: zapField, Fields: zapFieldSet}, zapField.Constraint)
	if specFrom.Defined() {
		if !zapFrom.Defined() {
			diffs = append(diffs, &StringDiff{Type: DiffTypeMismatch, Property: minProp, Spec: specFrom.ZapString(specField.Type)})
		} else if !specFrom.ValueEquals(zapFrom) {
			diffs = append(diffs, &StringDiff{Type: DiffTypeMismatch, Property: minProp, Spec: specFrom.ZapString(specField.Type), ZAP: zapFrom.ZapString(zapField.Type)})
		}
	} else if zapFrom.Defined() {
		diffs = append(diffs, &StringDiff{Type: DiffTypeMismatch, Property: minProp, ZAP: zapFrom.ZapString(zapField.Type)})
	}
	if specTo.Defined() {
		if !zapTo.Defined() {
			diffs = append(diffs, &StringDiff{Type: DiffTypeMismatch, Property: maxProp, Spec: specTo.ZapString(specField.Type)})
		} else if !specTo.ValueEquals(zapTo) {
			diffs = append(diffs, &StringDiff{Type: DiffTypeMismatch, Property: maxProp, Spec: specTo.ZapString(specField.Type), ZAP: zapTo.ZapString(zapField.Type)})
		}
	} else if zapTo.Defined() {
		diffs = append(diffs, &StringDiff{Type: DiffTypeMismatch, Property: maxProp, ZAP: zapTo.ZapString(zapField.Type)})
	}
	return
}

func namesEqual(specName string, zapName string) bool {
	if strings.EqualFold(specName, zapName) {
		return true
	}
	specName = matter.Case(specName)
	zapName = matter.Case(zapName)
	return strings.EqualFold(specName, zapName)
}
