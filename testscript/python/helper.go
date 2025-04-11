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
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
	"github.com/project-chip/alchemy/testplan"
	"github.com/project-chip/alchemy/testscript"
)

func registerSpecHelpers(t *raymond.Template, spec *spec.Specification) {
	t.RegisterHelper("pics", picsHelper)
	t.RegisterHelper("picsList", picsListHelper)
	t.RegisterHelper("picsGuard", picsGuardHelper)
	t.RegisterHelper("conformanceGuard", conformanceGuardHelper)
	t.RegisterHelper("actionIs", actionIsHelper)
	t.RegisterHelper("commandIs", commandIsHelper)
	t.RegisterHelper("pythonValue", pythonValueHelper)
	t.RegisterHelper("asUpperCamelCase", asUpperCamelCaseHelper)
	t.RegisterHelper("clusterVariable", clusterVariableHelper(spec))
	t.RegisterHelper("endpointVariable", endpointVariableHelper)
	t.RegisterHelper("commandArg", commandArgHelper)
	t.RegisterHelper("commandArgs", commandArgsHelper(spec))
	t.RegisterHelper("statusError", statusErrorHelper)
	t.RegisterHelper("octetString", octetStringHelper)
	t.RegisterHelper("pythonString", pythonStringHelper)
	t.RegisterHelper("ifIsEnum", ifEnumHelper)
	t.RegisterHelper("ifIsBitmap", ifBitmapHelper)
	t.RegisterHelper("ifNeedsConformanceCheck", needsConformanceCheckHelper)
	t.RegisterHelper("minConstraint", minConstraintHelper)
	t.RegisterHelper("maxConstraint", maxConstraintHelper)
	t.RegisterHelper("ifFieldIsNullable", ifFieldIsNullableHelper)
	t.RegisterHelper("ifFieldIsArray", ifFieldIsArrayHelper)
	t.RegisterHelper("ifFieldHasLength", ifFieldHasLengthHelper)
	t.RegisterHelper("typeCheckIs", typeCheckIsHelper)
	t.RegisterHelper("unimplementedTypeCheck", unimplementedTypeCheckHelper)
	t.RegisterHelper("entryTypeCheckIs", entryTypeCheckIsHelper)
	t.RegisterHelper("type", typeHelper)
	t.RegisterHelper("ifFieldIsOptional", ifFieldIsOptionalHelper)
}

func registerDocHelpers(t *raymond.Template, errata *errata.SDK) {
	t.RegisterHelper("bitmapName", bitmapNameHelper(errata))
	t.RegisterHelper("entityTypeName", entityTypeNameHelper(errata))
	t.RegisterHelper("entryTypeName", entryTypeNameHelper(errata))
	t.RegisterHelper("clusterName", clusterNameHelper(errata))
	t.RegisterHelper("attributeName", attributeNameHelper(errata))
	t.RegisterHelper("stepClusterName", stepClusterNameHelper(errata))
	t.RegisterHelper("enumName", enumNameHelper(errata))
	t.RegisterHelper("entityTypeQualifiedName", entityTypeQualifiedNameHelper(errata))
	t.RegisterHelper("entryTypeQualifiedName", entryTypeQualifiedNameHelper(errata))
	t.RegisterHelper("entityTypeFullName", entityTypeFullNameHelper(errata))
	t.RegisterHelper("entryTypeFullName", entryTypeFullNameHelper(errata))
	t.RegisterHelper("structName", structNameHelper(errata))
	t.RegisterHelper("structFullName", structFullNameHelper(errata))
	t.RegisterHelper("fieldName", fieldNameHelper(errata))
}

func typeHelper(t any) raymond.SafeString {
	return raymond.SafeString(fmt.Sprintf("%T", t))
}

func actionIsHelper(action testscript.TestAction, is string, options *raymond.Options) string {
	var ok bool
	switch is {
	case "readAttribute":
		_, ok = action.(*testscript.ReadAttribute)
		if !ok {
			_, ok = action.(testscript.ReadAttribute)
		}
	case "writeAttribute":
		slog.Info("writeAttribute", log.Type("action", action))
		_, ok = action.(*testscript.WriteAttribute)
		if !ok {
			_, ok = action.(testscript.WriteAttribute)
		}
	case "checkMinConstraint":
		_, ok = action.(*testscript.CheckMinConstraint)
		if !ok {
			_, ok = action.(testscript.CheckMinConstraint)
		}
	case "checkMaxConstraint":
		_, ok = action.(*testscript.CheckMaxConstraint)
		if !ok {
			_, ok = action.(testscript.CheckMaxConstraint)
		}
	case "checkRangeConstraint":
		_, ok = action.(*testscript.CheckRangeConstraint)
		if !ok {
			_, ok = action.(testscript.CheckRangeConstraint)
		}
	case "checkType":
		_, ok = action.(*testscript.CheckType)
		if !ok {
			_, ok = action.(testscript.CheckType)
		}
	case "checkStruct":
		_, ok = action.(*testscript.CheckStruct)
		if !ok {
			_, ok = action.(testscript.CheckStruct)
		}
	case "anyOf":
		_, ok = action.(*testscript.CheckAnyOfConstraint)
		if !ok {
			_, ok = action.(testscript.CheckAnyOfConstraint)
		}
	case "equals":
		_, ok = action.(*testscript.CheckValueConstraint)
		if !ok {
			_, ok = action.(testscript.CheckValueConstraint)
		}
	case "list":
		_, ok = action.(*testscript.CheckListEntries)
		if !ok {
			_, ok = action.(testscript.CheckListEntries)
		}
	}
	if ok {
		return options.Fn()
	}
	return options.Inverse()
}

func clusterNameHelper(errata *errata.SDK) func(test testscript.Test) raymond.SafeString {
	return func(test testscript.Test) raymond.SafeString {
		return raymond.SafeString(errata.OverrideName(test.Cluster, spec.CanonicalName(test.Cluster.Name)))
	}
}

func stepClusterNameHelper(errata *errata.SDK) func(test testscript.Test, step testscript.TestStep) raymond.SafeString {
	return func(test testscript.Test, step testscript.TestStep) raymond.SafeString {
		clusterName := errata.OverrideName(test.Cluster, test.Cluster.Name)
		if step.Cluster != nil {
			clusterName = errata.OverrideName(step.Cluster, step.Cluster.Name)
		}
		return raymond.SafeString(spec.CanonicalName(clusterName))
	}
}

func attributeNameHelper(errata *errata.SDK) func(step testscript.TestStep, action testscript.TestAction) raymond.SafeString {
	return func(step testscript.TestStep, action testscript.TestAction) raymond.SafeString {
		switch action := action.(type) {
		case testscript.ReadAttribute:
			if action.Attribute != nil {
				return raymond.SafeString(errata.OverrideName(action.Attribute, action.Attribute.Name))
			}
			return raymond.SafeString(action.AttributeName)
		case *testscript.ReadAttribute:
			if action.Attribute != nil {
				return raymond.SafeString(errata.OverrideName(action.Attribute, action.Attribute.Name))
			}
			return raymond.SafeString(action.AttributeName)
		case testscript.WriteAttribute:
			if action.Attribute != nil {
				return raymond.SafeString(errata.OverrideName(action.Attribute, action.Attribute.Name))
			}
			return raymond.SafeString(action.AttributeName)
		case *testscript.WriteAttribute:
			if action.Attribute != nil {
				return raymond.SafeString(errata.OverrideName(action.Attribute, action.Attribute.Name))
			}
			return raymond.SafeString(action.AttributeName)
		case testscript.CheckType:
			return raymond.SafeString(errata.OverrideName(action.Field, action.Field.Name))
		case testscript.CheckMaxConstraint:
			return raymond.SafeString(errata.OverrideName(action.Field, action.Field.Name))
		case testscript.CheckMinConstraint:
			return raymond.SafeString(errata.OverrideName(action.Field, action.Field.Name))
		case testscript.CheckRangeConstraint:
			return raymond.SafeString(errata.OverrideName(action.Field, action.Field.Name))
		default:
			slog.Error("Unexpected action type in attribute name helper", log.Type("type", action))
		}
		return raymond.SafeString("UnknownAttribute")
	}
}

func clusterVariableHelper(sp *spec.Specification) func(test testscript.Test, step testscript.TestStep, action testscript.TestAction) raymond.SafeString {
	return func(test testscript.Test, step testscript.TestStep, action testscript.TestAction) raymond.SafeString {
		var cluster *matter.Cluster
		switch action := action.(type) {
		case testscript.ReadAttribute:
			cluster = action.Cluster
		case testscript.WriteAttribute:
			cluster = action.Cluster
		}
		if cluster == nil {
			cluster = step.Cluster
		}
		if cluster == nil {
			cluster = test.Cluster
		}
		if cluster == test.Cluster {
			return raymond.SafeString("cluster")
		}
		clusterName := cluster.Name
		return raymond.SafeString("Clusters." + spec.CanonicalName(clusterName))
	}
}

func endpointVariableHelper(test testscript.Test, endpoint uint64) raymond.SafeString {
	if endpoint != math.MaxUint64 {
		return raymond.SafeString(strconv.FormatUint(endpoint, 10))
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

func globalVariableHelper(variables map[string]types.Entity) func(variableName string) raymond.SafeString {
	return func(variableName string) raymond.SafeString {
		e, ok := variables[variableName]
		if ok {
			switch e := e.(type) {
			case *matter.Field:
				return raymond.SafeString("None")
			case *matter.Constant:
				switch v := e.Value.(type) {
				case int:
					return raymond.SafeString(strconv.Itoa(v))
				default:
					slog.Warn("Unexpected constant value type", log.Type("type", v), matter.LogEntity("entity", e))
				}
				return raymond.SafeString("None")
			}
		}

		return raymond.SafeString("None")
	}
}

func ifEnumHelper(e types.Entity, options *raymond.Options) string {
	switch e := e.(type) {
	case *matter.Field:
		if e.Type.IsEnum() {
			return options.Fn()
		}
		switch e.Type.Entity.(type) {
		case *matter.Enum:
			return options.Fn()
		}
	default:
		slog.Error("Unexpected type checking isEnum", log.Type("type", e))
	}
	return options.Inverse()
}

func enumNameHelper(errata *errata.SDK) func(action testscript.CheckType) raymond.SafeString {
	return func(action testscript.CheckType) raymond.SafeString {
		if action.Field.Type.Entity == nil {
			slog.Error("Missing enum entity on field", slog.String("fieldName", action.Field.Name))
			return raymond.SafeString("")
		}
		switch entity := action.Field.Type.Entity.(type) {
		case *matter.Enum:
			return raymond.SafeString(errata.OverrideName(entity, entity.Name))
		default:
			slog.Error("Unexpected type getting enum name", log.Type("type", entity))
			return raymond.SafeString("unknown")
		}
	}
}

func ifBitmapHelper(e types.Entity, options *raymond.Options) string {
	switch e := e.(type) {
	case *matter.Field:
		switch e.Type.Entity.(type) {
		case *matter.Bitmap:
			return options.Fn()
		}
	default:
		slog.Error("Unexpected type checking ifBitmap", log.Type("type", e))
	}
	return options.Inverse()
}

func bitmapNameHelper(errata *errata.SDK) func(action testscript.CheckType) raymond.SafeString {
	return func(action testscript.CheckType) raymond.SafeString {
		if action.Field.Type.Entity == nil {
			slog.Error("Missing bitmap entity on field", slog.String("fieldName", action.Field.Name))
			return raymond.SafeString("")
		}
		switch entity := action.Field.Type.Entity.(type) {
		case *matter.Bitmap:
			return raymond.SafeString(errata.OverrideName(entity, entity.Name))
		default:
			slog.Error("Unexpected type getting bitmap name", log.Type("type", entity))
			return raymond.SafeString("unknown")
		}
	}
}

func entityTypeNameHelper(errata *errata.SDK) func(test testscript.Test, step testscript.TestStep, action testscript.CheckType) raymond.SafeString {
	return func(test testscript.Test, step testscript.TestStep, action testscript.CheckType) raymond.SafeString {
		return customTypeHelper(test, step, *action.Field, action.Field.Type, action.Field.Type.Entity, customTypeNameTypeShort, errata)
	}
}

func entryTypeNameHelper(errata *errata.SDK) func(test testscript.Test, step testscript.TestStep, field matter.Field) raymond.SafeString {
	return func(test testscript.Test, step testscript.TestStep, field matter.Field) raymond.SafeString {
		return customTypeHelper(test, step, field, field.Type.EntryType, field.Type.EntryType.Entity, customTypeNameTypeShort, errata)
	}
}

func entityTypeQualifiedNameHelper(errata *errata.SDK) func(test testscript.Test, step testscript.TestStep, action testscript.CheckType) raymond.SafeString {
	return func(test testscript.Test, step testscript.TestStep, action testscript.CheckType) raymond.SafeString {
		return customTypeHelper(test, step, *action.Field, action.Field.Type, action.Field.Type.Entity, customTypeNameTypeQualified, errata)
	}
}

func entryTypeQualifiedNameHelper(errata *errata.SDK) func(test testscript.Test, step testscript.TestStep, field matter.Field) raymond.SafeString {
	return func(test testscript.Test, step testscript.TestStep, field matter.Field) raymond.SafeString {
		return customTypeHelper(test, step, field, field.Type.EntryType, field.Type.EntryType.Entity, customTypeNameTypeQualified, errata)
	}
}

func entityTypeFullNameHelper(errata *errata.SDK) func(test testscript.Test, step testscript.TestStep, action testscript.CheckType) raymond.SafeString {
	return func(test testscript.Test, step testscript.TestStep, action testscript.CheckType) raymond.SafeString {
		return customTypeHelper(test, step, *action.Field, action.Field.Type, action.Field.Type.Entity, customTypeNameTypeFull, errata)
	}
}

func entryTypeFullNameHelper(errata *errata.SDK) func(test testscript.Test, step testscript.TestStep, field matter.Field) raymond.SafeString {
	return func(test testscript.Test, step testscript.TestStep, field matter.Field) raymond.SafeString {
		return customTypeHelper(test, step, field, field.Type.EntryType, field.Type.EntryType.Entity, customTypeNameTypeFull, errata)
	}
}

func structNameHelper(errata *errata.SDK) func(test testscript.Test, step testscript.TestStep, s *matter.Struct) raymond.SafeString {
	return func(test testscript.Test, step testscript.TestStep, s *matter.Struct) raymond.SafeString {
		return customTypeHelper(test, step, matter.Field{}, nil, s, customTypeNameTypeShort, errata)
	}
}

func structFullNameHelper(errata *errata.SDK) func(test testscript.Test, step testscript.TestStep, s *matter.Struct) raymond.SafeString {
	return func(test testscript.Test, step testscript.TestStep, s *matter.Struct) raymond.SafeString {
		return customTypeHelper(test, step, matter.Field{}, nil, s, customTypeNameTypeFull, errata)
	}
}

type customTypeNameType int

const (
	customTypeNameTypeShort customTypeNameType = iota
	customTypeNameTypeQualified
	customTypeNameTypeFull
)

func customTypeHelper(test testscript.Test, step testscript.TestStep, field matter.Field, dataType *types.DataType, entity types.Entity, nameType customTypeNameType, errata *errata.SDK) raymond.SafeString {
	if dataType != nil && !dataType.IsCustom() && testscript.CanCheckType(&field) {
		return raymond.SafeString(toPythonType(dataType.BaseType))
	}
	var name string
	var collection string
	var cluster *matter.Cluster
	switch entryEntity := entity.(type) {
	case *matter.Bitmap:
		name = errata.OverrideName(entryEntity, entryEntity.Name)
		cluster = entryEntity.Cluster()
		collection = "Bitmaps"
	case *matter.Command:
		name = errata.OverrideName(entryEntity, entryEntity.Name)
		cluster = entryEntity.Cluster()
		collection = "Commands"
	case *matter.Struct:
		name = errata.OverrideName(entryEntity, entryEntity.Name)
		cluster = entryEntity.Cluster()
		collection = "Structs"
	case *matter.Enum:
		name = errata.OverrideName(entryEntity, entryEntity.Name)
		cluster = entryEntity.Cluster()
		collection = "Enums"
	case nil:
		slog.Error("Missing entry type entity on list field", slog.String("fieldName", field.Name), slog.String("baseDataType", dataType.BaseType.String()))
		return raymond.SafeString("")
	case *matter.TypeDef:
		return customTypeHelper(test, step, field, entryEntity.Type, nil, nameType, errata)
	default:
		slog.Error("Unknown entry type entity on list field", slog.String("fieldName", field.Name), log.Type("type", entryEntity))
		return raymond.SafeString("")
	}
	var clusterPrefix string
	switch nameType {
	case customTypeNameTypeShort:
	case customTypeNameTypeQualified:
		var localCluster = test.Cluster
		if step.Cluster != nil {
			localCluster = step.Cluster
		}
		if cluster == nil {
			clusterPrefix = fmt.Sprintf("Globals.%s.", collection)
		} else if localCluster != cluster {
			clusterPrefix = fmt.Sprintf("Clusters.%s.%s.", spec.CanonicalName(cluster.Name), collection)
		} else {
			clusterPrefix = fmt.Sprintf("cluster.%s.", collection)
		}
	case customTypeNameTypeFull:
		if cluster == nil {
			clusterPrefix = fmt.Sprintf("Globals.%s.", collection)
		} else {
			clusterPrefix = fmt.Sprintf("Clusters.%s.%s.", spec.CanonicalName(cluster.Name), collection)
		}
	}
	return raymond.SafeString(clusterPrefix + name)
}

func fieldNameHelper(errata *errata.SDK) func(test testscript.Test, step testscript.TestStep, field matter.Field) raymond.SafeString {
	return func(test testscript.Test, step testscript.TestStep, field matter.Field) raymond.SafeString {
		return raymond.SafeString(strcase.ToLowerCamel(errata.OverrideName(&field, field.Name)))
	}
}

func ifFieldIsNullableHelper(field matter.Field, options *raymond.Options) string {
	if field.Quality.Has(matter.QualityNullable) {
		return options.Fn()
	}
	return options.Inverse()
}

func ifFieldIsOptionalHelper(field matter.Field, options *raymond.Options) string {
	if !conformance.IsRequired(field.Conformance) {
		return options.Fn()
	}
	return options.Inverse()
}

func ifFieldIsArrayHelper(field matter.Field, options *raymond.Options) string {
	if field.Type.IsArray() {
		return options.Fn()
	}
	return options.Inverse()
}

func ifFieldHasLengthHelper(field matter.Field, options *raymond.Options) string {
	if field.Type.HasLength() {
		return options.Fn()
	}
	return options.Inverse()
}
