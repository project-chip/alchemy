package generate

import (
	"fmt"
	"log/slog"
	"strconv"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/mailgun/raymond/v2"
	"github.com/project-chip/alchemy/matter/spec"
)

func registerHelpers(t *raymond.Template, spec *spec.Specification) {
	t.RegisterHelper("quote", quoteHelper)
	t.RegisterHelper("pics", picsHelper)
	t.RegisterHelper("picsGuard", picsGuardHelper)
	t.RegisterHelper("clusterIs", clusterIsHelper)
	t.RegisterHelper("commandIs", commandIsHelper)
	t.RegisterHelper("comparisonValue", comparisonValueHelper)
	t.RegisterHelper("asUpperCamelCase", asUpperCamelCaseHelper)
	t.RegisterHelper("raw", rawHelper)
	t.RegisterHelper("ifSet", ifSetHelper)
	t.RegisterHelper("clusterName", clusterNameHelper(spec))
	t.RegisterHelper("stepClusterName", stepClusterNameHelper(spec))
	t.RegisterHelper("commandArg", commandArgHelper)
	t.RegisterHelper("commandArgs", commandArgsHelper(spec))
	t.RegisterHelper("statusError", statusErrorHelper)
}

func quoteHelper(s string) raymond.SafeString {
	return raymond.SafeString(strconv.Quote(s))
}

func clusterIsHelper(step testStep, is string, options *raymond.Options) string {
	if step.Cluster == is {
		return options.Fn()
	}
	return options.Inverse()
}

func clusterNameHelper(sp *spec.Specification) func(test test) raymond.SafeString {
	return func(test test) raymond.SafeString {
		clusterName := test.Config.Cluster
		_, ok := sp.ClustersByName[clusterName]
		if !ok {
			slog.Warn("Unknown cluster in test", slog.String("clusterName", clusterName))
		}
		return raymond.SafeString(spec.CanonicalName(clusterName))
	}
}

func stepClusterNameHelper(sp *spec.Specification) func(test test, step testStep) raymond.SafeString {
	return func(test test, step testStep) raymond.SafeString {
		clusterName := test.Config.Cluster
		if step.Cluster != "" {
			clusterName = step.Cluster
		}
		_, ok := sp.ClustersByName[clusterName]
		if !ok {
			slog.Warn("Unknown cluster in test", slog.String("clusterName", clusterName))
		}
		return raymond.SafeString(spec.CanonicalName(clusterName))
	}
}

func commandIsHelper(step testStep, is string, options *raymond.Options) string {
	if step.Command == is {
		return options.Fn()
	}
	return options.Inverse()
}

func comparisonValueHelper(value any) raymond.SafeString {
	switch value := value.(type) {
	case uint64:
		return raymond.SafeString(strconv.FormatUint(value, 10))
	case string:
		return raymond.SafeString(value)
	case int64:
		return raymond.SafeString(strconv.FormatInt(value, 10))
	case map[string]any:
		var sb strings.Builder
		sb.WriteRune('{')
		var count int
		for key, val := range value {
			if count > 0 {
				sb.WriteString(", ")
			}
			sb.WriteString(key)
			sb.WriteString(": ")
			sb.WriteString(string(comparisonValueHelper(val)))
			count++
		}
		sb.WriteRune('}')
		return raymond.SafeString(sb.String())
	case []any:
		var elements []string
		for _, e := range value {
			elements = append(elements, string(comparisonValueHelper(e)))
		}
		return raymond.SafeString("[" + strings.Join(elements, ", ") + "]")
	case nil:
		return raymond.SafeString("None")
	default:
		return raymond.SafeString(fmt.Sprintf("unknown comparisonValue type: %T", value))
	}
}

func rawHelper(value string) raymond.SafeString {
	return raymond.SafeString(value)
}

func ifSetHelper(value any, options *raymond.Options) string {
	switch value.(type) {
	case nil:
		return options.Inverse()
	default:
		return options.Fn()
	}
}

func statusErrorHelper(value string) raymond.SafeString {
	return raymond.SafeString("Status." + strcase.ToCamel(value))
}

func asUpperCamelCaseHelper(value string) raymond.SafeString {
	return raymond.SafeString(strcase.ToCamel(value))
}
