package testscript

import (
	"slices"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/constraint"
	"github.com/project-chip/alchemy/matter/types"
)

func findStructs(cluster *matter.Cluster) (structs map[*matter.Struct]struct{}) {
	structs = make(map[*matter.Struct]struct{})

	for _, a := range cluster.Attributes {
		findStructsForField(a, structs)
	}
	return
}

func findStructsForField(field *matter.Field, structs map[*matter.Struct]struct{}) {
	if field.Type == nil {
		return
	}
	findStructsForEntity(field.Type.Entity, structs)
	if field.Type.EntryType != nil {
		findStructsForEntity(field.Type.EntryType.Entity, structs)
	}
}

func findStructsForEntity(entity types.Entity, structs map[*matter.Struct]struct{}) {
	if entity == nil {
		return
	}
	switch e := entity.(type) {
	case *matter.Struct:
		structs[e] = struct{}{}
		for _, f := range e.Fields {
			findStructsForField(f, structs)
		}
	default:
		return
	}
}

func buildTestsForStructs(structs map[*matter.Struct]struct{}) (tests []*TestStep, err error) {
	sl := make([]*matter.Struct, 0, len(structs))
	for s := range structs {
		sl = append(sl, s)
	}
	slices.SortStableFunc(sl, func(a, b *matter.Struct) int {
		return strings.Compare(a.Name, b.Name)
	})
	for _, s := range sl {
		step := &TestStep{
			Name:   s.Name,
			Entity: s,
		}
		checkStruct := &CheckStruct{
			action: action{},

			Struct: s,
		}
		for _, f := range s.Fields {
			if conformance.IsBlank(f.Conformance) && constraint.IsBlank(f.Constraint) {
				continue
			}

			checkStructField := &CheckStructField{
				Field: f,
			}

			variableName := "struct." + strcase.ToLowerCamel(f.Name)

			var actions []TestAction
			if canCheckType(f) {
				checkStructField.Validations = append(checkStructField.Validations, &CheckType{constraintAction: constraintAction{Field: f, Variable: variableName}})
			}
			actions, err = addConstraintActions(f, s.Fields, f.Constraint, variableName)
			if err != nil {
				return
			}
			checkStructField.Validations = append(checkStructField.Validations, actions...)
			if len(checkStructField.Validations) > 0 {
				checkStruct.Fields = append(checkStruct.Fields, checkStructField)
			}
		}
		if len(checkStruct.Fields) > 0 {
			step.Actions = append(step.Actions, checkStruct)
			tests = append(tests, step)
		}
	}
	return
}
