package generate

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/goccy/go-yaml"
	"github.com/iancoleman/strcase"
	"github.com/mailgun/raymond/v2"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
	"github.com/project-chip/alchemy/testing/parse"
)

func getArg(arg any) (name string, value any, ok bool) {
	var n any
	switch v := arg.(type) {
	case map[string]any:

		n, ok = v["name"]
		if !ok {
			return
		}
		name, ok = n.(string)
		if !ok {
			return
		}
	case yaml.MapSlice:
		n, ok = parse.ValueFromMapSlice(v, "name")
		if !ok {
			return
		}
		name, ok = n.(string)
		if !ok {
			return
		}
		value, ok = parse.ValueFromMapSlice(v, "value")

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
		return fmt.Sprintf("%v=%s", strcase.ToLowerCamel(name), string(pythonValueHelper(value)))
	}
	if field.Type.Entity == nil {
		return fmt.Sprintf("%v=%s", strcase.ToLowerCamel(name), string(pythonValueHelper(value)))

	}
	var parentEntity types.Entity
	var namespace string
	var objectGroup string
	var valName string
	var objectName string
	switch entity := field.Type.Entity.(type) {
	case *matter.Enum:
		parentEntity = entity.Parent()
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
			slog.Info("unknown arg type", slog.String("name", name), log.Type("type", value))
		}
	case *matter.Bitmap:
		parentEntity = entity.Parent()
		objectGroup = "Bitmaps"
		objectName = entity.Name
		switch value := value.(type) {
		case uint64:
			for _, val := range entity.Bits {
				from, to, err := val.Bits()
				if err != nil || from != to {
					continue
				}
				if from == value {
					valName = "k" + val.Name()
					break
				}
			}

		default:
			slog.Info("unknown arg type", slog.String("name", name), log.Type("type", value))
		}
	case *matter.Struct:
		slog.Warn("struct argument entity", log.Type("type", value))
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
	if namespace != "" || objectGroup != "" || objectName != "" {
		if namespace != "" {
			sb.WriteString(namespace)
			sb.WriteRune('.')
		}
		if objectGroup != "" {
			sb.WriteString(objectGroup)
			sb.WriteRune('.')
		}
		if objectName != "" {
			sb.WriteString(objectName)
			sb.WriteRune('.')
		}
		if valName != "" {
			sb.WriteString(valName)
		} else {
			sb.WriteRune('(')
			sb.WriteString(string(pythonValueHelper(value)))
			sb.WriteRune(')')
		}
		return sb.String()
	}
	sb.WriteString(string(pythonValueHelper(value)))
	return sb.String()

}

func commandArgHelper(test test, step testStep, name string) raymond.SafeString {
	for _, arg := range step.Arguments.Values {
		argName, value, ok := getArg(arg)
		if !ok {
			slog.Warn("unable to cast arg", slog.String("testId", test.ID), log.Type("type", arg))
			continue
		}
		if strings.EqualFold(name, argName) {
			return pythonValueHelper(value)
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
				var name string
				var value any
				switch v := v.(type) {
				case map[string]any:
					var ok bool
					name, value, ok = getArg(v)
					if !ok {
						continue
					}

				case yaml.MapSlice:
					n, ok := parse.ValueFromMapSlice(v, "name")
					if !ok {
						continue
					}
					name, ok = n.(string)
					if !ok {
						continue
					}
					value, ok = parse.ValueFromMapSlice(v, "value")
					if !ok {
						continue
					}
				}
				if name == "" {
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

		return raymond.SafeString(strings.Join(args, ", "))
	}
}
