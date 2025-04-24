package testscript

import (
	"context"
	"log/slog"
	"strconv"

	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/internal/suggest"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
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
	var testPlan *testplan.Test
	testPlan = input.Content
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
		Cluster:  cluster,
		ID:       testPlan.ID,
		Name:     testPlan.Name,
		PICSList: testPlan.PICSList,
		/*Config: parse.TestConfig{
			Cluster: cluster.Name,
		},*/
		YamlVariables: testPlan.Config.Extras,
	}

	slog.Info("yaml", slog.Any("extras", t.YamlVariables))

	attributes := make(map[string]*matter.Field)
	for _, a := range cluster.Attributes {
		attributes[a.Name] = a
	}

	for _, g := range testPlan.Groups {
		step := &TestStep{
			Name:        g.Name,
			Description: g.Description,
		}
		for _, s := range g.Steps {
			switch s.Command {
			case "readAttribute":
				slog.Info("Reading attribute", slog.Any("comments", s.Comments))
				commandCluster := cluster
				a, ok := attributes[s.Attribute]
				if !ok {
					if s.Cluster == "" {
						slog.Error("Unknown attribute", slog.String("testPlanId", testPlan.ID), slog.String("attribute", s.Attribute))
						continue
					}
					c, ok := sp.spec.ClustersByName[s.Cluster]
					if !ok {
						slog.Error("Unknown cluster", slog.String("testPlanId", testPlan.ID), slog.String("cluster", s.Cluster))
						continue
					}
					for _, attr := range c.Attributes {
						if attr.Name == s.Attribute {
							a = attr
							commandCluster = c
							break
						}
					}
					if a == nil {
						slog.Error("Unknown attribute", slog.String("testPlanId", testPlan.ID), slog.String("attribute", s.Attribute))
						continue
					}
				}
				variableName := "val"
				if s.Response.SaveAs != "" {
					variableName = s.Response.SaveAs
				}
				readAttribute := &ReadAttribute{
					remoteAction: remoteAction{
						action: action{
							Description: s.Description,
							Conformance: a.Conformance,
						},
						Endpoint: s.Endpoint,
					},
					Attribute:  a,
					Attributes: cluster.Attributes,
					Variable:   variableName,
				}
				if commandCluster != cluster {
					readAttribute.Cluster = commandCluster
				}
				readAttribute.Validations, err = buildValidations(s, a, variableName)
				step.Actions = append(step.Actions, readAttribute)
			case "writeAttribute":
				a, ok := attributes[s.Attribute]
				if !ok {
					slog.Error("Unknown attribute", slog.String("testPlanId", testPlan.ID), slog.String("attribute", s.Attribute))
					continue
				}
				writeAttribute := &WriteAttribute{
					remoteAction: remoteAction{
						action: action{
							Comments: s.Comments,
						},

						Endpoint: s.Endpoint,
					},
					Attribute: a,
					Value:     s.Arguments.Value,
				}
				if !conformance.IsMandatory(a.Conformance) {
					attributeCheck := &conformance.Mandatory{
						Expression: &conformance.IdentifierExpression{
							ID:     a.Name,
							Entity: a,
						},
					}
					writeAttribute.Conformance = conformance.Set{attributeCheck}
				}
				if s.Response.Error != "" {
					slog.Info("expected error", "error", s.Response.Error)
					writeAttribute.ExpectedError = s.Response.Error
				}
				step.Actions = append(step.Actions, writeAttribute)
			default:
				slog.Error("Unknown command converting test plan to test script", slog.String("testPlanId", testPlan.ID), slog.String("command", s.Command))
			}
		}
		t.AddStep(step)
	}

	outputs = append(outputs, pipeline.NewData(getPath(sp.sdkRoot, t), t))

	return
}

func buildValidations(step *testplan.Step, field *matter.Field, variableName string) (validations []TestAction, err error) {
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
		if step.Response.Constraints.AnyOf != nil {
			validations = append(validations, &CheckAnyOfConstraint{constraintAction: constraintAction{Field: field, Variable: variableName}, Values: step.Response.Constraints.AnyOf})
		}
	}
	slog.Warn("validation", log.Type("value", step.Response.Value))
	if step.Response.Value != nil {
		slog.Warn("adding validation", "value", step.Response.Value)
		validations = append(validations, &CheckValueConstraint{constraintAction: constraintAction{Field: field, Variable: variableName}, Value: step.Response.Value})
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
