package testscript

import (
	"log/slog"

	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/testplan"
)

func (sp *TestScriptConverter) convertCommand(testPlan *testplan.Test, s *testplan.Step, cluster *matter.Cluster, step *TestStep) (err error) {
	commandCluster := cluster
	if s.Cluster != "" && s.Cluster != "DelayCommands" {
		c, ok := sp.spec.ClustersByName[s.Cluster]
		if !ok {
			slog.Error("Unknown cluster", slog.String("testPlanId", testPlan.ID), slog.String("cluster", s.Cluster))
			return
		}
		commandCluster = c
	}
	switch s.Command {
	case "readAttribute":
		slog.Info("Reading attribute", slog.Any("comments", s.Comments))
		a := commandCluster.Attributes.Get(s.Attribute)
		if a == nil {
			slog.Error("Unknown attribute", slog.String("testPlanId", testPlan.ID), slog.String("attribute", s.Attribute))
			return
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
		if err != nil {
			return
		}
		step.Actions = append(step.Actions, readAttribute)
	case "writeAttribute":
		a := cluster.Attributes.Get(s.Attribute)
		if a == nil {
			slog.Error("Unknown attribute", slog.String("testPlanId", testPlan.ID), slog.String("attribute", s.Attribute))
			return
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
	case "subscribeAttribute":
		a := cluster.Attributes.Get(s.Attribute)
		if a == nil {
			slog.Error("Unknown attribute to subscribe to", slog.String("testPlanId", testPlan.ID), slog.String("attribute", s.Attribute))
			return
		}
		subscribeAttribute := &SubscribeAttribute{
			remoteAction: remoteAction{
				action: action{
					Comments: s.Comments,
				},

				Endpoint: s.Endpoint,
			},
			Attribute:   a,
			MinInterval: s.MinInterval,
			MaxInterval: s.MaxInterval,
			Timeout:     s.Timeout,
		}
		step.Actions = append(step.Actions, subscribeAttribute)
	case "TestEventTrigger":
		testEventTrigger := &TestEventTrigger{}
		args := s.Arguments.Values.ToValues()
		for name, value := range args {
			switch name {
			case "EventTrigger":
				testEventTrigger.EventTrigger = value.(string)
			case "EnableKey":
				testEventTrigger.EnableKey = value.(string)
			}
		}
		step.Actions = append(step.Actions, testEventTrigger)
	case "WaitForCommissionee":
		step.Actions = append(step.Actions, &WaitForCommissionee{})
	default:
		err = sp.convertClusterCommand(testPlan, s, cluster, step)
	}
	return
}

func (sp *TestScriptConverter) convertClusterCommand(testPlan *testplan.Test, s *testplan.Step, cluster *matter.Cluster, step *TestStep) (err error) {
	commandName := s.Command
	var command *matter.Command
	for _, c := range cluster.Commands {
		if commandName == c.Name {
			command = c
			break
		}
	}
	if command == nil {
		slog.Error("Unknown command converting test plan to test script", slog.String("testPlanId", testPlan.ID), slog.String("command", s.Command))
		return
	}
	slog.Info("Adding command", "name", command.Name)
	callCommand := &CallCommand{Cluster: cluster, Command: command}
	args := s.Arguments.Values.ToValues()
	for name, value := range args {
		var field *matter.Field
		for _, f := range command.Fields {
			if name == f.Name {
				field = f
				break
			}
		}
		if field == nil {
			slog.Error("Unknown command argument", slog.String("testPlanId", testPlan.ID), slog.String("command", s.Command), slog.String("argument", name))
			return
		}
		callCommand.Arguments = append(callCommand.Arguments, &CommandArgument{Field: field, Value: value})
	}
	step.Actions = append(step.Actions, callCommand)
	return
}
