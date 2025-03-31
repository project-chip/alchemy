package testscript

import (
	"fmt"
	"log/slog"
	"math"

	"github.com/iancoleman/strcase"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
	"github.com/project-chip/alchemy/sdk"
)

func (*TestScriptGenerator) buildClusterTest(cluster *matter.Cluster) (t *Test, err error) {
	t = &Test{
		Cluster: cluster,
		ID:      strcase.ToScreamingSnake(spec.CanonicalName(cluster.Name)) + "_2_1",
		Name:    "Attributes with Server as DUT",
		/*Config: parse.TestConfig{
			Cluster: cluster.Name,
		},*/
	}

	/*t.AddStep(&TestStep{
		Description: "Read feature map",
		Actions: []TestAction{
			&ReadAttribute{
				remoteAction: remoteAction{
					Endpoint: math.MaxUint64,
				},
				AttributeName: "FeatureMap",
				Variable:      "feature_map",
			},
		},
	})*/

	variables := findVariables(cluster)
	for v := range variables {
		switch v := v.(type) {
		case *matter.Field:
			t.GlobalVariables = append(t.GlobalVariables, v.Name)
		default:
			err = fmt.Errorf("unexpected variable type: %T", v)
			return
		}
	}
	//slices.Sort(t.Variables)

	structs := findStructs(cluster)

	t.StructChecks, err = buildTestsForStructs(structs)
	if err != nil {
		return
	}

	for _, a := range cluster.Attributes {
		if conformance.IsDeprecated(a.Conformance) {
			continue
		}
		step := &TestStep{
			Description: fmt.Sprintf("Read %s attribute", a.Name),
		}
		readAttribute := &ReadAttribute{
			remoteAction: remoteAction{
				action: action{
					Conformance: a.Conformance,
				},
				Endpoint: math.MaxUint64,
			},
			Attribute:  a,
			Attributes: cluster.Attributes,
		}
		step.Actions = append(step.Actions, readAttribute)

		variableName := "val"
		_, ok := variables[a]
		if ok {
			variableName = "self." + a.Name
		}

		if canCheckType(a) {
			readAttribute.Validations = append(readAttribute.Validations, &CheckType{constraintAction: constraintAction{Field: a, Variable: variableName}})
		}
		var actions []TestAction
		actions, err = addConstraintActions(a, cluster.Attributes, a.Constraint, variableName)
		if err != nil {
			return
		}
		readAttribute.Validations = append(readAttribute.Validations, actions...)
		if len(readAttribute.Validations) > 0 {
			readAttribute.Variable = variableName
		}
		t.AddStep(step)
	}
	return
}

func canCheckType(field *matter.Field) bool {
	underlyingType := sdk.ToUnderlyingType(field.Type.BaseType)
	switch underlyingType {
	case types.BaseDataTypeUInt64,
		types.BaseDataTypeInt64,
		types.BaseDataTypeUInt32,
		types.BaseDataTypeInt32,
		types.BaseDataTypeUInt16,
		types.BaseDataTypeInt16,
		types.BaseDataTypeInt8,
		types.BaseDataTypeUInt8,
		types.BaseDataTypeOctStr,
		types.BaseDataTypeString,
		types.BaseDataTypeBoolean,
		types.BaseDataTypeCustom,
		types.BaseDataTypeList:
		return true
	}
	slog.Info("can't check type", slog.String("field", field.Name), slog.String("type", field.Type.BaseType.String()))
	return false
}
