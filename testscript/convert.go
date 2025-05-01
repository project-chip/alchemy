package testscript

import (
	"context"
	"log/slog"
	"strconv"

	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/internal/suggest"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/constraint"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
	"github.com/project-chip/alchemy/testplan"
)

type TestScriptConverter struct {
	spec       *spec.Specification
	sdkRoot    string
	picsLabels map[string]string
}

func NewTestScriptConverter(spec *spec.Specification, sdkRoot string, picsLabels map[string]string) *TestScriptConverter {
	return &TestScriptConverter{spec: spec, sdkRoot: sdkRoot, picsLabels: picsLabels}
}

func (sp TestScriptConverter) Name() string {
	return "Converting test plan to test script"
}

func (sp *TestScriptConverter) Process(cxt context.Context, input *pipeline.Data[*testplan.Test], index int32, total int32) (outputs []*pipeline.Data[*Test], extras []*pipeline.Data[*testplan.Test], err error) {
	testPlan := input.Content
	cluster := testPlan.Cluster

	if cluster == nil {
		clusterName := testPlan.Config.Cluster
		if clusterName == "" {
			slog.Error("Missing cluster converting test plan to test script", slog.String("testPlanId", testPlan.ID))
			return
		}
		var ok bool
		cluster, ok = sp.spec.ClustersByName[clusterName]
		if !ok {
			slog.Error("Unknown cluster converting test plan to test script", slog.String("testPlanId", testPlan.ID), slog.String("clusterName", clusterName))
			possibilities := make(map[types.Entity]int)
			suggest.PossibleEntities(clusterName, possibilities, func(yield func(string, types.Entity) bool) {
				for name, cluster := range sp.spec.ClustersByName {
					if !yield(name, cluster) {
						return
					}
				}
			})
			suggest.ListPossibilities(clusterName, possibilities)
			return
		}
	}

	t := &Test{
		Cluster:       cluster,
		ID:            testPlan.ID,
		Name:          testPlan.Name,
		PICSList:      testPlan.PICSList,
		YamlVariables: testPlan.Config.Extras,
	}

	slog.Info("yaml", slog.Any("extras", t.YamlVariables))

	for _, g := range testPlan.Groups {
		step := &TestStep{
			Name:        g.Name,
			Description: g.Description,
		}
		for _, testPlanStep := range g.Steps {
			err = sp.convertCommand(testPlan, testPlanStep, cluster, step)
			if err != nil {
				return
			}
		}
		t.AddStep(step)
	}

	outputs = append(outputs, pipeline.NewData(getPath(sp.sdkRoot, t), t))

	return
}

func buildValidations(step *testplan.Step, field *matter.Field, variableName string) (validations []TestAction, err error) {
	if step.Response.Value != nil {
		slog.Warn("adding validation", "value", step.Response.Value, log.Type("type", step.Response.Value))
		validations = append(validations, &CheckValueConstraint{constraintAction: constraintAction{Field: field, Variable: variableName}, Value: step.Response.Value})
	}
	if step.Response.Constraints != nil {
		if step.Response.Constraints.Type != "" {
			validations = append(validations, &CheckType{constraintAction: constraintAction{Field: field, Variable: variableName}})
		}
		if step.Response.Constraints.MinValue != nil {
			validations = append(validations, &CheckMinConstraint{constraintAction: constraintAction{Field: field, Variable: variableName}, Constraint: &constraint.MinConstraint{Minimum: buildValidationLimit(step.Response.Constraints.MinValue)}})
		}
		if step.Response.Constraints.MaxValue != nil {
			validations = append(validations, &CheckMaxConstraint{constraintAction: constraintAction{Field: field, Variable: variableName}, Constraint: &constraint.MaxConstraint{Maximum: buildValidationLimit(step.Response.Constraints.MaxValue)}})
		}
		if step.Response.Constraints.MinLength != nil {
			validations = append(validations, &CheckMinConstraint{constraintAction: constraintAction{Field: field, Variable: variableName}, Constraint: &constraint.MinConstraint{Minimum: buildValidationLimit(step.Response.Constraints.MinLength)}})
		}
		if step.Response.Constraints.MaxLength != nil {
			validations = append(validations, &CheckMaxConstraint{constraintAction: constraintAction{Field: field, Variable: variableName}, Constraint: &constraint.MaxConstraint{Maximum: buildValidationLimit(step.Response.Constraints.MaxLength)}})
		}
		if step.Response.Constraints.AnyOf != nil {
			validations = append(validations, &CheckAnyOfConstraint{constraintAction: constraintAction{Field: field, Variable: variableName}, Values: step.Response.Constraints.AnyOf})
		}
		if step.Response.Constraints.NotValue != nil {
			validations = append(validations, &CheckNotValueConstraint{constraintAction: constraintAction{Field: field, Variable: variableName}, Value: step.Response.Value})
		}
	}
	return
}

func buildValidationLimit(val any) constraint.Limit {
	switch val := val.(type) {
	case uint64:
		return &constraint.HexLimit{Value: val}
	case int64:
		return &constraint.IntLimit{Value: val}
	case string:
		i, err := strconv.ParseInt(val, 10, 64)
		if err == nil {
			return &constraint.IntLimit{Value: i}
		}
	default:
		slog.Error("Unexpected limit value type", log.Type("type", val))
	}
	return nil
}
