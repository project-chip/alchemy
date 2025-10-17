package regen

import (
	"context"
	"embed"
	"fmt"
	"log/slog"
	"path/filepath"
	"slices"
	"strconv"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/mailgun/raymond/v2"
	"github.com/project-chip/alchemy/internal/handlebars"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
	"github.com/project-chip/alchemy/zap"
)

//go:embed templates
var templateFiles embed.FS

type IdlRenderer struct {
	spec *spec.Specification
}

func NewIdlRenderer(spec *spec.Specification) (IdlRenderer, error) {
	return IdlRenderer{spec: spec}, nil
}

func (p IdlRenderer) Name() string {
	return "Writing Matter files"
}

func (p IdlRenderer) Process(cxt context.Context, input *pipeline.Data[*zap.File], index int32, total int32) (outputs []*pipeline.Data[string], extras []*pipeline.Data[*zap.File], err error) {

	dir := filepath.Dir(input.Path)
	base := filepath.Base(input.Path)
	extension := filepath.Ext(base)
	file := strings.TrimSuffix(base, extension)
	path := filepath.Join(dir, file+".matter")

	slog.Info("converting zap path", "path", input.Path, "matter", path)

	var t *raymond.Template
	t, err = p.loadTemplate(p.spec)
	if err != nil {
		return
	}
	/*variables := make(map[string]struct{})
	t.RegisterHelper("variable", variableHelper(variables))
	t.RegisterHelper("globalVariable", globalVariableHelper(test.GlobalVariables))
	t.RegisterHelper("value", valueHelper(variables))
	registerDocHelpers(t, sdkErrata)*/

	clusters := make(map[*matter.Cluster]struct{})

	for _, endpoint := range input.Content.EndpointTypes {
		for _, clusterId := range endpoint.Clusters {
			cluster, ok := p.spec.ClustersByID[uint64(clusterId.Code)]
			if !ok {
				//err = fmt.Errorf("unrecognized cluster id in %s: %d", input.Path, clusterId.Code)
				return
			}
			clusters[cluster] = struct{}{}
		}
	}

	clusterList := make([]*matter.Cluster, 0, len(clusters))
	for cluster := range clusters {
		clusterList = append(clusterList, cluster)
	}

	slices.SortFunc(clusterList, func(a *matter.Cluster, b *matter.Cluster) int {
		return a.ID.Compare(b.ID)
	})

	tc := map[string]any{
		"clusters": clusterList,
	}
	var out string
	out, err = t.Exec(tc)
	if err != nil {
		slog.Error("error", "err", err)
		return
	}
	outputs = append(outputs, pipeline.NewData(path, out))
	return
}

var template pipeline.Once[*raymond.Template]

func (sp *IdlRenderer) loadTemplate(spec *spec.Specification) (*raymond.Template, error) {
	t, err := template.Do(func() (*raymond.Template, error) {

		ov := handlebars.NewOverlay("", templateFiles, "templates")
		err := ov.Flush()
		if err != nil {
			slog.Error("Error flushing embedded templates", slog.Any("error", err))
		}
		t, err := handlebars.LoadTemplate("{{> matter}}", ov)
		if err != nil {
			return nil, err
		}

		handlebars.RegisterCommonHelpers(t)

		registerIdlHelpers(t, spec)

		return t, nil
	})
	if err != nil {
		return nil, err
	}
	return t.Clone(), nil
}

func registerIdlHelpers(t *raymond.Template, spec *spec.Specification) {
	t.RegisterHelper("number", numberHelper)
	t.RegisterHelper("currentRevision", currentRevisionHelper)
	t.RegisterHelper("asUpperCamelCase", asUpperCamelCaseHelper)
	t.RegisterHelper("asLowerCamelCase", asLowerCamelCaseHelper)
	t.RegisterHelper("descriptionComment", descriptionCommentHelper)
	t.RegisterHelper("enums", clusterEnumsHelper)
	t.RegisterHelper("structs", clusterStructsHelper)
	t.RegisterHelper("structFields", clusterStructFieldsHelper)
	t.RegisterHelper("fieldType", fieldTypeHelper)
	t.RegisterHelper("fieldIsArray", fieldIsArrayHelper)
	t.RegisterHelper("ifFabricScoped", ifFabricScopedHelper)
	t.RegisterHelper("ifOptional", ifOptionalHelper)
	t.RegisterHelper("ifNullable", ifNullableHelper)
	t.RegisterHelper("ifReadOnly", ifReadOnlyHelper)
	/*t.RegisterHelper("picsGuard", clusterEnumsHelper)
	t.RegisterHelper("endpointVariable", endpointVariableHelper)
	t.RegisterHelper("commandArgs", commandArgsHelper(spec))
	t.RegisterHelper("statusError", statusErrorHelper)
	t.RegisterHelper("octetString", octetStringHelper)
	t.RegisterHelper("pythonString", pythonStringHelper)
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
	t.RegisterHelper("ifFieldIsOptional", ifFieldIsOptionalHelper)*/
}

func numberHelper(value any) raymond.SafeString {
	switch value := value.(type) {
	case uint64:
		return raymond.SafeString(strconv.FormatUint(value, 10))
	case string:
		return raymond.SafeString(value)
	case int64:
		return raymond.SafeString(strconv.FormatInt(value, 10))
	case matter.Number:
		return raymond.SafeString(value.IntString())
	case *matter.Number:
		return raymond.SafeString(value.IntString())
	default:
		return raymond.SafeString(fmt.Sprintf("unknown number type: %T", value))
	}
}

func currentRevisionHelper(cluster matter.Cluster) raymond.SafeString {
	var lastRevision uint64
	for _, rev := range cluster.Revisions {
		revNumber := matter.ParseNumber(rev.Number)
		if revNumber.Valid() && revNumber.Value() > lastRevision {
			lastRevision = revNumber.Value()
		}
	}
	return raymond.SafeString(strconv.FormatUint(lastRevision, 10))

}

func asUpperCamelCaseHelper(value string) raymond.SafeString {
	return raymond.SafeString(strcase.ToCamel(value))
}

func asLowerCamelCaseHelper(value string) raymond.SafeString {
	return raymond.SafeString(strcase.ToLowerCamel(value))
}

func enumerateHelper[T any](list []T, options *raymond.Options) raymond.SafeString {
	var result strings.Builder
	for i, en := range list {
		df := options.DataFrame().Copy()
		df.Set("index", i)
		df.Set("key", nil)
		df.Set("first", i == 0)
		df.Set("last", i == len(list)-1)
		result.WriteString(options.FnCtxData(en, df))
	}
	return raymond.SafeString(result.String())
}

func enumsHelper(enums matter.EnumSet, options *raymond.Options) raymond.SafeString {

	slices.SortStableFunc(enums, func(a *matter.Enum, b *matter.Enum) int {
		return strings.Compare(a.Name, b.Name)
	})
	return enumerateHelper(enums, options)
}

func clusterEnumsHelper(cluster matter.Cluster, options *raymond.Options) raymond.SafeString {
	slog.Info("rendering enums", "count", len(cluster.Enums))
	enums := make(matter.EnumSet, len(cluster.Enums))
	copy(enums, cluster.Enums)
	return enumsHelper(enums, options)
}

func clusterStructsHelper(cluster matter.Cluster, options *raymond.Options) raymond.SafeString {
	structs := make(matter.StructSet, len(cluster.Structs))
	copy(structs, cluster.Structs)
	slices.SortStableFunc(structs, func(a *matter.Struct, b *matter.Struct) int {
		return strings.Compare(a.Name, b.Name)
	})
	return enumerateHelper(structs, options)
}

func clusterStructFieldsHelper(cluster matter.Struct, options *raymond.Options) raymond.SafeString {
	fields := make(matter.FieldSet, len(cluster.Fields))
	copy(fields, cluster.Fields)
	if cluster.FabricScoping == matter.FabricScopingScoped {
		fields = append(fields, &matter.Field{ID: matter.NewNumber(254), Name: "FabricIndex", Type: types.NewDataType(types.BaseDataTypeFabricIndex, false), Conformance: conformance.Set{&conformance.Mandatory{}}})
	}
	return enumerateHelper(fields, options)
}

func fieldTypeHelper(a any, options *raymond.Options) raymond.SafeString {
	var t *types.DataType
	switch a := a.(type) {
	case types.DataType:
		t = &a
	case *types.DataType:
		t = a
	default:
		return raymond.SafeString("notDataType")
	}
	if t == nil {
		return raymond.SafeString("missingType")

	}
	if t.IsArray() {
		if t.EntryType != nil {
			return raymond.SafeString(zap.DataTypeName(t.EntryType))
		}
		return raymond.SafeString("missingType")
	} else {
		return raymond.SafeString(zap.DataTypeName(t))
	}

}

func fieldIsArrayHelper(a any, options *raymond.Options) string {
	var t *types.DataType
	switch a := a.(type) {
	case types.DataType:
		t = &a
	case *types.DataType:
		t = a
	default:
		return options.Inverse()
	}
	if t == nil {
		return options.Inverse()
	}
	if t.IsArray() {
		return options.Fn()
	} else {
		return options.Inverse()
	}

}

func ifFabricScopedHelper(s matter.Struct, options *raymond.Options) string {
	if s.FabricScoping == matter.FabricScopingScoped {
		return options.Fn()
	}
	return options.Inverse()
}

func ifOptionalHelper(conf conformance.Conformance, options *raymond.Options) string {
	if !conformance.IsMandatory(conf) {
		return options.Fn()
	}
	return options.Inverse()
}

func ifNullableHelper(conf matter.Quality, options *raymond.Options) string {
	if conf.Has(matter.QualityNullable) {
		return options.Fn()
	}
	return options.Inverse()
}

func ifReadOnlyHelper(conf matter.Access, options *raymond.Options) string {
	if conf.Write != matter.PrivilegeUnknown && conf.Write != matter.PrivilegeOperate {
		return options.Inverse()
	}
	return options.Fn()
}

func descriptionCommentHelper(description string) raymond.SafeString {
	if len(description) == 0 {
		return raymond.SafeString("")
	}
	var comment strings.Builder
	var line strings.Builder
	line.WriteString("/**")
	words := strings.Split(description, " ")
	for _, word := range words {
		if line.Len() > 100 {
			comment.WriteString(line.String())
			comment.WriteRune('\n')
			line.Reset()
			line.WriteString("      ")
		}
		line.WriteRune(' ')
		line.WriteString(word)
	}
	line.WriteString(" */")
	comment.WriteString(line.String())
	comment.WriteRune('\n')
	return raymond.SafeString(comment.String())
}
