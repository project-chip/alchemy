package generate

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
)

func (sp *PythonTestGenerator) writeCommand(t *test, ts *testStep, indent string, sb *strings.Builder) error {
	if ts.Disabled {
		sb.WriteString(indent)
		sb.WriteString("if False: # Disabled")
		indent = "    " + indent
	}
	switch ts.Command {
	case "readAttribute":
		readAttribute(ts, indent, sb)
	case "writeAttribute":
		writeAttribute(ts, indent, sb)
	case "UserPrompt":
		writeUserPrompt(ts, indent, sb)
	default:
		sp.clusterCommand(t, ts, indent, sb)
	}
	return nil
}

func (sp *PythonTestGenerator) clusterCommand(t *test, ts *testStep, indent string, sb *strings.Builder) error {

	clusterName := ts.Cluster
	if clusterName == "" {
		clusterName = t.cluster
	}
	cluster, ok := sp.spec.ClustersByName[clusterName]
	if !ok {
		slog.Warn("Unknown cluster in test", slog.String("clusterName", clusterName))
	}

	var command *matter.Command
	for _, cmd := range cluster.Commands {
		if strings.EqualFold(cmd.Name, ts.Command) {
			command = cmd
			break
		}
	}
	if command == nil {
		slog.Warn("Unknown command in test", slog.String("commandName", ts.Command))
	}
	sb.WriteString(indent)
	sb.WriteString("await self.send_single_cmd(cmd=Clusters.Objects.")
	sb.WriteString(clusterName)
	sb.WriteString(".Commands.")
	sb.WriteString(ts.Command)
	sb.WriteString("(")
	var args []string
	if len(ts.Arguments.Values) > 0 {
		for _, v := range ts.Arguments.Values {
			if v, ok := v.(map[string]any); ok {
				n, ok := v["name"]
				if !ok {
					continue

				}
				name, ok := n.(string)
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
				value, ok := v["value"]
				if !ok {
					continue
				}
				if field.Type.IsEnum() {

				}
				args = append(args, wrapValue(name, field, value))
			}
		}
	}
	sb.WriteString(strings.Join(args, ", "))
	sb.WriteString("))\n")
	return nil
}

func wrapValue(name string, field *matter.Field, value any) string {
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

func writeUserPrompt(ts *testStep, indent string, sb *strings.Builder) {
	sb.WriteString(indent)
	sb.WriteString("pass\n")
	sb.WriteString(indent)
	sb.WriteString("# TODO: Rewrite this user prompt test\n")
	sb.WriteString(indent)
	sb.WriteString("#\n")
	verification := strings.Split(ts.Verification, "\n")
	for _, v := range verification {
		sb.WriteString(indent)
		sb.WriteString("# ")
		sb.WriteString(v)
		sb.WriteString("\n")
	}
}
