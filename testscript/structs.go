package testscript

import (
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
	for s := range structs {
		step := &TestStep{
			Name: s.Name,
		}
		for _, f := range s.Fields {
			if conformance.IsBlank(f.Conformance) && constraint.IsBlank(f.Constraint) {
				continue
			}

			var actions []TestAction
			actions, err = addConstraintActions(f, s.Fields, f.Constraint, "struct."+f.Name)
			if err != nil {
				return
			}
			step.Actions = append(step.Actions, actions...)
		}
		if len(step.Actions) > 0 {
			tests = append(tests, step)
		}
	}
	return
}
