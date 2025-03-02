package python

import (
	"encoding/hex"
	"fmt"
	"log/slog"
	"math"
	"strconv"
	"strings"

	"github.com/goccy/go-yaml"
	"github.com/iancoleman/strcase"
	"github.com/mailgun/raymond/v2"
	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/testplan"
)

func registerHelpers(t *raymond.Template, spec *spec.Specification) {
	t.RegisterHelper("pics", picsHelper)
	t.RegisterHelper("picsGuard", picsGuardHelper)
	t.RegisterHelper("clusterIs", clusterIsHelper)
	t.RegisterHelper("commandIs", commandIsHelper)
	t.RegisterHelper("pythonValue", pythonValueHelper)
	t.RegisterHelper("asUpperCamelCase", asUpperCamelCaseHelper)
	t.RegisterHelper("clusterName", clusterNameHelper(spec))
	t.RegisterHelper("clusterVariable", clusterVariableHelper(spec))
	t.RegisterHelper("endpointVariable", endpointVariableHelper)
	t.RegisterHelper("stepClusterName", stepClusterNameHelper(spec))
	t.RegisterHelper("commandArg", commandArgHelper)
	t.RegisterHelper("commandArgs", commandArgsHelper(spec))
	t.RegisterHelper("statusError", statusErrorHelper)
	t.RegisterHelper("octetString", octetStringHelper)
	t.RegisterHelper("pythonString", pythonStringHelper)
}

func clusterIsHelper(step testplan.Step, is string, options *raymond.Options) string {
	if step.Cluster == is {
		return options.Fn()
	}
	return options.Inverse()
}

func clusterNameHelper(sp *spec.Specification) func(test testplan.Test) raymond.SafeString {
	return func(test testplan.Test) raymond.SafeString {
		clusterName := test.Config.Cluster
		_, ok := sp.ClustersByName[clusterName]
		if !ok {
			slog.Warn("Unknown cluster in test", slog.String("clusterName", clusterName))
		}
		return raymond.SafeString(spec.CanonicalName(clusterName))
	}
}

func stepClusterNameHelper(sp *spec.Specification) func(test testplan.Test, step testplan.Step) raymond.SafeString {
	return func(test testplan.Test, step testplan.Step) raymond.SafeString {
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

func clusterVariableHelper(sp *spec.Specification) func(test testplan.Test, step testplan.Step) raymond.SafeString {
	return func(test testplan.Test, step testplan.Step) raymond.SafeString {
		if step.Cluster == "" || step.Cluster == test.Config.Cluster {
			return raymond.SafeString("cluster")
		}
		clusterName := step.Cluster
		_, ok := sp.ClustersByName[clusterName]
		if !ok {
			slog.Warn("Unknown cluster in test", slog.String("clusterName", clusterName))
		}
		return raymond.SafeString("Clusters." + spec.CanonicalName(clusterName))
	}
}

func endpointVariableHelper(test testplan.Test, step testplan.Step) raymond.SafeString {
	if step.Endpoint != math.MaxUint64 {
		return raymond.SafeString(strconv.FormatUint(step.Endpoint, 10))
	}
	return raymond.SafeString("endpoint")
}

func commandIsHelper(step testplan.Step, is string, options *raymond.Options) string {
	if step.Command == is {
		return options.Fn()
	}
	return options.Inverse()
}

func pythonValueHelper(value any) raymond.SafeString {
	switch value := value.(type) {
	case uint64:
		return raymond.SafeString(strconv.FormatUint(value, 10))
	case string:
		return raymond.SafeString(value)
	case int64:
		return raymond.SafeString(strconv.FormatInt(value, 10))
	case yaml.MapSlice:
		var sb strings.Builder
		sb.WriteRune('{')
		var count int
		for _, val := range value {
			key, ok := val.Key.(string)
			if !ok {
				continue
			}
			if count > 0 {
				sb.WriteString(", ")
			}
			sb.WriteString(key)
			sb.WriteString(": ")
			sb.WriteString(string(pythonValueHelper(val.Value)))
			count++
		}
		sb.WriteRune('}')
		return raymond.SafeString(sb.String())
	case []uint64:
		return numberArray(value, strconv.FormatUint)
	case []int64:
		return numberArray(value, strconv.FormatInt)
	case []any:
		var elements []string
		for _, e := range value {
			elements = append(elements, string(pythonValueHelper(e)))
		}
		return raymond.SafeString("[" + strings.Join(elements, ", ") + "]")
	case nil:
		return raymond.SafeString("None")
	case bool:
		if value {
			return raymond.SafeString("True")
		}
		return raymond.SafeString("False")
	default:
		return raymond.SafeString(fmt.Sprintf("unknown pythonValue type: %T", value))
	}
}

func numberArray[T any](value []T, formatter func(T, int) string) raymond.SafeString {
	var sb strings.Builder
	var count int
	sb.WriteRune('[')
	for _, v := range value {
		if count > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(formatter(v, 10))
		count++
	}
	sb.WriteRune(']')
	return raymond.SafeString(sb.String())
}

func statusErrorHelper(value string) raymond.SafeString {
	return raymond.SafeString("Status." + strcase.ToCamel(value))
}

func asUpperCamelCaseHelper(value string) raymond.SafeString {
	return raymond.SafeString(strcase.ToCamel(value))
}

func octetStringHelper(value string) raymond.SafeString {
	if text.HasCaseInsensitivePrefix(value, "hex:") {
		bytes, err := hex.DecodeString(text.TrimCaseInsensitivePrefix(value, "hex:"))
		if err != nil {
			slog.Warn("Error parsing hexadecimal value", slog.Any("error", err), slog.String("value", value))
		} else {
			var hexstring strings.Builder
			hexstring.WriteString("b'")
			for _, b := range bytes {
				hexstring.WriteRune('\\')
				hexstring.WriteString(fmt.Sprintf("%x", b))
			}
			hexstring.WriteRune('\'')
			return raymond.SafeString(hexstring.String())
		}
	}
	return raymond.SafeString(`""`)
}

func pythonStringHelper(s string) raymond.SafeString {
	var sb strings.Builder
	quoteCharacter := '\''
	if strings.ContainsRune(s, '\'') && !strings.ContainsRune(s, '"') {
		quoteCharacter = '"'
	}
	sb.WriteRune(quoteCharacter)
	for _, r := range s {
		switch {
		case r < ' ':
			// Control characters
			switch r {
			// See https://www.w3schools.com/python/gloss_python_escape_characters.asp
			case '\n':
				sb.WriteString(`\n`)
			case '\r':
				sb.WriteString(`\r`)
			case '\t':
				sb.WriteString(`\t`)
			case '\b':
				sb.WriteString(`\b`)
			case '\f':
				sb.WriteString(`\f`)
			default:
				sb.WriteString(fmt.Sprintf(`\x%02x`, r))
			}
		case strconv.IsPrint(r):
			if r == '\\' || r == quoteCharacter {
				sb.WriteRune('\\')
			}
			sb.WriteRune(r)
		default:
			sb.WriteString(fmt.Sprintf(`\x%02x`, r))
		}
	}
	sb.WriteRune(quoteCharacter)
	return raymond.SafeString(sb.String())
}

func valueHelper(variables map[string]struct{}) func(value any) raymond.SafeString {

	return func(value any) raymond.SafeString {
		switch value := value.(type) {
		case string:
			_, ok := variables[value]
			if ok {
				return raymond.SafeString(value)
			}
			return pythonStringHelper(value)
		default:
			return pythonValueHelper(value)
		}
	}
}

func variableHelper(variables map[string]struct{}) func(variableName string) raymond.SafeString {
	return func(variableName string) raymond.SafeString {
		variables[variableName] = struct{}{}
		return raymond.SafeString(variableName)
	}
}
