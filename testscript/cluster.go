package testscript

import (
	"fmt"
	"log/slog"
	"math"
	"slices"

	"github.com/iancoleman/strcase"
	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
	"github.com/project-chip/alchemy/sdk"
	"github.com/project-chip/alchemy/testplan/pics"
)

func (*TestScriptGenerator) buildClusterTest(doc *asciidoc.Document, cluster *matter.Cluster) (t *Test, err error) {
	id := cluster.PICS

	if id == "" {
		slog.Warn("Missing PICS for cluster, substituting name", slog.String("clusterName", cluster.Name))
		id = strcase.ToScreamingSnake(spec.CanonicalName(cluster.Name))
	}
	t = &Test{
		Doc:     doc,
		Cluster: cluster,
		ID:      id + "_2_1",
		Name:    "Attributes with Server as DUT",
	}

	if cluster.PICS != "" {
		var p pics.Expression
		p, err = pics.ParseString(cluster.PICS + ".S")
		if err != nil {
			return
		}
		t.PICSList = append(t.PICSList, p)
	}

	variables := make(map[types.Entity]struct{})
	findReferencedEntities(cluster, variables)

	if len(variables) > 0 {
		t.GlobalVariables = make(map[string]types.Entity)
		for v := range variables {
			switch v := v.(type) {
			case *matter.Field:
				t.GlobalVariableNames = append(t.GlobalVariableNames, v.Name)
				t.GlobalVariables[v.Name] = v
			case *matter.Constant:
				t.GlobalVariableNames = append(t.GlobalVariableNames, v.Name)
				t.GlobalVariables[v.Name] = v
			case *matter.EnumValue:
				t.GlobalVariableNames = append(t.GlobalVariableNames, v.Name)
				t.GlobalVariables[v.Name] = v
			default:
				err = fmt.Errorf("unexpected variable type: %T", v)
				return
			}
		}
		slices.Sort(t.GlobalVariableNames)
	}

	structs := findStructs(cluster)

	t.StructChecks, err = buildTestsForStructs(structs)
	if err != nil {
		return
	}

	for _, a := range cluster.Attributes {
		if conformance.IsDeprecated(a.Conformance) || conformance.IsDisallowed(a.Conformance) {
			continue
		}
		step := &TestStep{
			Description: fmt.Sprintf("Read %s attribute", a.Name),
		}
		readAttribute := &ReadAttribute{
			remoteAction: remoteAction{
				Endpoint: math.MaxUint64,
			},
			Attribute:  a,
			Attributes: cluster.Attributes,
		}

		if !conformance.IsMandatory(a.Conformance) {
			attributeCheck := &conformance.Mandatory{
				Expression: &conformance.IdentifierExpression{
					ID:     a.Name,
					Entity: a,
				},
			}
			readAttribute.Conformance = conformance.Set{attributeCheck}
		}

		step.Actions = append(step.Actions, readAttribute)

		variableName := "val"
		_, ok := variables[a]
		if ok {
			variableName = "self." + a.Name
		}

		if CanCheckType(a) {
			readAttribute.Validations = append(readAttribute.Validations, &CheckType{constraintAction: constraintAction{Field: a, Variable: variableName}})
		}
		var actions []TestAction
		actions, err = addConstraintActions(a, cluster.Attributes, a.Constraint, variableName)
		if err != nil {
			return
		}
		actions = append(actions, checkBitmapRange(a, cluster.Attributes, variableName)...)
		readAttribute.Validations = append(readAttribute.Validations, actions...)

		if len(readAttribute.Validations) > 0 {
			readAttribute.Variable = variableName
		}
		t.AddStep(step)
	}
	return
}

func CanCheckType(field *matter.Field) bool {
	if field == nil || field.Type == nil {
		return false
	}
	underlyingType := sdk.ToUnderlyingType(field.Type.BaseType)
	switch underlyingType {
	case types.BaseDataTypeUInt64,
		types.BaseDataTypeInt64,
		types.BaseDataTypeUInt32,
		types.BaseDataTypeInt32,
		types.BaseDataTypeUInt24,
		types.BaseDataTypeInt24,
		types.BaseDataTypeUInt16,
		types.BaseDataTypeInt16,
		types.BaseDataTypeInt8,
		types.BaseDataTypeUInt8,
		types.BaseDataTypeEnum16,
		types.BaseDataTypeEnum8,
		types.BaseDataTypeMap64,
		types.BaseDataTypeMap32,
		types.BaseDataTypeMap16,
		types.BaseDataTypeMap8,
		types.BaseDataTypeOctStr,
		types.BaseDataTypeString,
		types.BaseDataTypeBoolean,
		types.BaseDataTypeSingle,
		types.BaseDataTypeDouble,
		types.BaseDataTypeCustom,
		types.BaseDataTypeList,
		types.BaseDataTypeVendorID,
		types.BaseDataTypeGroupID,
		types.BaseDataTypeDeviceTypeID,
		types.BaseDataTypeNodeID,
		types.BaseDataTypeMessageID,
		types.BaseDataTypeSubjectID,
		types.BaseDataTypeFabricID,
		types.BaseDataTypeFabricIndex,
		types.BaseDataTypeIPAddress,
		types.BaseDataTypeIPv4Address,
		types.BaseDataTypeIPv6Prefix,
		types.BaseDataTypeIPv6Address,
		types.BaseDataTypeHardwareAddress,
		types.BaseDataTypeClusterID,
		types.BaseDataTypeEndpointNumber,
		types.BaseDataTypeTag,
		types.BaseDataTypeNamespaceID:
		return true
	}
	slog.Warn("Unimplemented base type; no type check will be generated", slog.String("field", field.Name), slog.String("type", field.Type.BaseType.String()), log.Path("source", field))
	return false
}
