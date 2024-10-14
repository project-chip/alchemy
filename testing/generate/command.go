package generate

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/mailgun/raymond/v2"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
)

func getArg(arg any) (name string, value any, ok bool) {
	var m map[string]any
	m, ok = arg.(map[string]any)
	if !ok {
		return
	}
	var n any
	n, ok = m["name"]
	if !ok {
		return
	}
	name, ok = n.(string)
	if !ok {
		return
	}
	value, ok = m["value"]
	if !ok {
		return
	}
	return
}

func getCluster(spec *spec.Specification, t *test, ts *testStep) (clusterName string, cluster *matter.Cluster) {
	clusterName = ts.Cluster
	if clusterName == "" {
		clusterName = t.Config.Cluster
	}
	var ok bool
	cluster, ok = spec.ClustersByName[clusterName]
	if !ok {
		slog.Warn("Unknown cluster in test", slog.String("path", ts.Parent.Path), slog.String("clusterName", clusterName))
	}
	return
}

func wrapValue(name string, field *matter.Field, value any) string {
	if field == nil {
		return fmt.Sprintf("unknown field: %s", name)
	}
	if field.Type == nil {
		return fmt.Sprintf("%v=%v", strcase.ToLowerCamel(name), value)
	}
	if field.Type.Entity != nil {
		var parentEntity types.Entity
		var namespace string
		var objectGroup string
		var valName string
		var objectName string
		switch entity := field.Type.Entity.(type) {
		case *matter.Enum:
			parentEntity = entity.ParentEntity
			objectGroup = "Enums"
			objectName = entity.Name
			switch value := value.(type) {
			case uint64:
				for _, val := range entity.Values {
					if val.Value.Valid() && val.Value.Value() == value {
						valName = "k" + val.Name
						break
					}
				}
			default:
				slog.Info("unknown arg type", log.Type("type", value))
			}

		default:
			slog.Warn("unsupported argument entity", log.Type("type", entity))
		}
		if parentEntity != nil {
			switch parentEntity := parentEntity.(type) {
			case *matter.Cluster:
				namespace = fmt.Sprintf("Clusters.Objects.%s", spec.CanonicalName(parentEntity.Name))
			}
		}
		var sb strings.Builder
		sb.WriteString(strcase.ToLowerCamel(name))
		sb.WriteRune('=')
		sb.WriteString(namespace)
		sb.WriteRune('.')
		sb.WriteString(objectGroup)
		sb.WriteRune('.')
		sb.WriteString(objectName)
		sb.WriteRune('.')
		if valName != "" {
			sb.WriteString(valName)
		} else {
			sb.WriteRune('(')
			sb.WriteString(fmt.Sprintf("%v", value))
			sb.WriteRune(')')
		}
		return sb.String()
	}
	return fmt.Sprintf("%v=%v", strcase.ToLowerCamel(name), value)
}

func commandArgHelper(test test, step testStep, name string) raymond.SafeString {
	for _, arg := range step.Arguments.Values {
		argName, value, ok := getArg(arg)
		if !ok {
			slog.Warn("unable to cast arg", slog.String("testId", test.ID), log.Type("type", arg))
			continue
		}
		if strings.EqualFold(name, argName) {
			return comparisonValueHelper(value)
		}
	}
	return raymond.SafeString(fmt.Sprintf("unknown argument: %s", name))
}

func commandArgsHelper(spec *spec.Specification) func(test test, step testStep) raymond.SafeString {
	return func(test test, step testStep) raymond.SafeString {
		clusterName, cluster := getCluster(spec, &test, &step)
		if cluster == nil {
			return raymond.SafeString(fmt.Sprintf("error: unknown cluster: %s", clusterName))
		}
		var command *matter.Command
		for _, cmd := range cluster.Commands {
			if strings.EqualFold(cmd.Name, step.Command) {
				command = cmd
				break
			}
		}
		if command == nil {
			slog.Warn("Unknown command in test", slog.String("testId", test.ID), slog.String("step", step.Label), slog.String("commandName", step.Command))
		}
		var args []string
		if len(step.Arguments.Values) > 0 {
			for _, v := range step.Arguments.Values {
				if v, ok := v.(map[string]any); ok {
					name, value, ok := getArg(v)
					if !ok {
						continue
					}
					var field *matter.Field
					for _, f := range command.Fields {
						if strings.EqualFold(name, f.Name) {
							field = f
							break
						}
					}
					if field == nil {
						slog.Warn("Unknown command field in test", slog.String("fieldName", name))
					}
					args = append(args, wrapValue(name, field, value))
				}
			}
		}

		return raymond.SafeString(strings.Join(args, ", "))
	}
}
