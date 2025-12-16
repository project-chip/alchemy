package regen

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/mailgun/raymond/v2"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
	"github.com/project-chip/alchemy/sdk"
	"github.com/project-chip/alchemy/zap"
)

func (sp *IdlRenderer) registerIdlHelpers(t *raymond.Template, spec *spec.Specification) {
	t.RegisterHelper("number", numberHelper)
	t.RegisterHelper("currentRevision", currentRevisionHelper)
	t.RegisterHelper("asUpperCamelCase", asUpperCamelCaseHelper)
	t.RegisterHelper("asLowerCamelCase", asLowerCamelCaseHelper)
	t.RegisterHelper("upperCamelCaseDiffers", upperCamelCaseDiffersHelper)
	t.RegisterHelper("lowerCamelCaseDiffers", lowerCamelCaseDiffersHelper)
	t.RegisterHelper("descriptionComment", descriptionCommentHelper)
	t.RegisterHelper("enums", enumsHelper(sp.spec))
	t.RegisterHelper("structs", clusterStructsHelper(sp.spec))
	t.RegisterHelper("events", clusterEventsHelper(sp.spec))
	t.RegisterHelper("commands", commandsHelper(sp.spec))
	t.RegisterHelper("requests", requestsHelper(sp.spec))
	t.RegisterHelper("isTimed", isTimedHelper)
	t.RegisterHelper("requestName", requestNameHelper)
	t.RegisterHelper("responseName", responseNameHelper)
	t.RegisterHelper("attributes", clusterAttributesHelper(sp.spec, sp.commonAttributes))
	t.RegisterHelper("bitmaps", clusterBitmapsHelper(sp.spec))
	t.RegisterHelper("structFields", structFieldsHelper(sp.spec))
	t.RegisterHelper("eventFields", eventFieldsHelper(sp.spec))
	t.RegisterHelper("commandFields", commandFieldsHelper(sp.spec))
	t.RegisterHelper("fieldType", fieldTypeHelper)
	t.RegisterHelper("fieldIsArray", fieldIsArrayHelper)
	t.RegisterHelper("ifFabricScoped", ifFabricScopedHelper)
	t.RegisterHelper("ifFabricSensitive", ifFabricSensitiveHelper)
	t.RegisterHelper("ifOptional", ifOptionalHelper)
	t.RegisterHelper("ifNullable", ifNullableHelper)
	t.RegisterHelper("ifReadOnly", ifReadOnlyHelper)
	t.RegisterHelper("ifHasAccess", ifHasAccessHelper)
	t.RegisterHelper("ifHasValidFields", ifHasValidFieldsHelper)
	t.RegisterHelper("ifRequest", ifRequestHelper)

	t.RegisterHelper("storageOption", storageOptionHelper)

	t.RegisterHelper("accessString", accessStringHelper)
	t.RegisterHelper("ifHasLengthConstraint", ifHasLengthConstraintHelper)
	t.RegisterHelper("lengthConstraint", lengthConstraintHelper)
	t.RegisterHelper("bitName", bitNameHelper)
	t.RegisterHelper("bitMask", bitMaskHelper)
	t.RegisterHelper("bitmapType", bitmapTypeHelper)
	t.RegisterHelper("deviceTypeName", deviceTypeNameHelper)
	t.RegisterHelper("endpointServers", endpointServersHelper(sp.spec))
	t.RegisterHelper("endpointClients", endpointClientsHelper(spec))

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

func currentRevisionHelper(revisions matter.Revisions) raymond.SafeString {
	lastRevision := revisions.MostRecent()
	if lastRevision != nil {
		return raymond.SafeString(lastRevision.Number.IntString())
	}
	return raymond.SafeString("")

}

func enumerateHelper[T any](list []T, spec *spec.Specification, options *raymond.Options) raymond.SafeString {
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

func enumsHelper(spec *spec.Specification) func(enums matter.EnumSet, options *raymond.Options) raymond.SafeString {
	return func(enums matter.EnumSet, options *raymond.Options) raymond.SafeString {
		sortedEnums := make(matter.EnumSet, len(enums))
		copy(sortedEnums, enums)
		slices.SortStableFunc(sortedEnums, func(a *matter.Enum, b *matter.Enum) int {
			return strings.Compare(a.Name, b.Name)
		})
		return enumerateEntitiesHelper(sortedEnums, spec, options)
	}
}

func ifHasValidFieldsHelper(fs matter.FieldSet, options *raymond.Options) string {
	if len(filterFields(fs)) > 0 {
		return options.Fn()
	} else {
		return options.Inverse()
	}
}

func structFieldsHelper(spec *spec.Specification) func(s matter.Struct, options *raymond.Options) raymond.SafeString {
	return func(s matter.Struct, options *raymond.Options) raymond.SafeString {
		fields := filterFields(s.Fields)
		if s.FabricScoping == matter.FabricScopingScoped {
			fabricIndex := &matter.Field{ID: matter.NewNumber(254), Name: "FabricIndex", Type: types.NewDataType(types.BaseDataTypeFabricIndex, false), Conformance: conformance.Set{&conformance.Mandatory{}}}
			fabricIndex.SetParent(&s)
			fields = append(fields, fabricIndex)
		}
		return enumerateEntitiesHelper(fields, spec, options)

	}
}

func eventFieldsHelper(spec *spec.Specification) func(e matter.Event, options *raymond.Options) raymond.SafeString {
	return func(e matter.Event, options *raymond.Options) raymond.SafeString {
		fields := filterFields(e.Fields)
		if e.Access.FabricSensitivity == matter.FabricSensitivitySensitive {
			fields = append(fields, &matter.Field{ID: matter.NewNumber(254), Name: "FabricIndex", Type: types.NewDataType(types.BaseDataTypeFabricIndex, false), Conformance: conformance.Set{&conformance.Mandatory{}}})
		}
		return enumerateEntitiesHelper(fields, spec, options)
	}
}

func ifFabricScopedHelper(a matter.FabricScoping, options *raymond.Options) string {
	if a == matter.FabricScopingScoped {
		return options.Fn()
	}
	return options.Inverse()
}

func ifFabricSensitiveHelper(access matter.Access, options *raymond.Options) string {
	if access.FabricSensitivity == matter.FabricSensitivitySensitive {
		return options.Fn()
	}
	return options.Inverse()
}

func maxValue(field matter.Field, fs matter.FieldSet) (max types.DataTypeExtreme, ok bool) {
	if field.Type == nil || field.Type.IsArray() {
		return
	}
	if field.Constraint == nil {
		return
	}
	if !field.Type.HasLength() {
		switch field.Type.Entity.(type) {
		case *matter.Enum, *matter.Bitmap:
			return
		}
		if field.Type.BaseType.IsSimple() {
			return
		}
		if sdk.ToUnderlyingType(field.Type.BaseType).IsSimple() {
			return
		}
	}
	_, max = zap.GetMinMax(matter.NewConstraintContext(&field, fs), field.Constraint)
	hasNumericMax := max.IsNumeric()
	if !hasNumericMax {
		return
	}

	var maxDueToNullable bool
	if hasNumericMax {
		maxDueToNullable = types.Max(sdk.ToUnderlyingType(sdk.FindBaseType(field.Type)), true).ValueEquals(max)
	}
	if maxDueToNullable {
		return
	}
	var redundant bool
	max, redundant = sdk.CheckUnderlyingType(&field, max, types.DataExtremePurposeMaximum)
	if redundant {
		return
	}
	ok = true
	return
}

func ifHasLengthConstraintHelper(field matter.Field, fs matter.FieldSet, options *raymond.Options) string {
	if _, ok := maxValue(field, fs); ok {
		return options.Fn()
	}
	return options.Inverse()
}

func lengthConstraintHelper(field matter.Field, fs matter.FieldSet) raymond.SafeString {
	if max, ok := maxValue(field, fs); ok {
		return raymond.SafeString(max.ZapString(field.Type))
	}
	return raymond.SafeString("")
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
	if conf.Write != matter.PrivilegeUnknown {
		return options.Inverse()
	}
	return options.Fn()
}

func ifHasAccessHelper(a matter.Access, options *raymond.Options) string {
	if (a.Read != matter.PrivilegeUnknown && a.Read != matter.PrivilegeView) ||
		(a.Write != matter.PrivilegeUnknown && a.Write != matter.PrivilegeOperate) ||
		(a.Invoke != matter.PrivilegeUnknown && a.Invoke != matter.PrivilegeOperate) {
		return options.Fn()
	}
	return options.Inverse()
}

func accessStringHelper(a matter.Access) raymond.SafeString {
	var list strings.Builder
	if a.Read != matter.PrivilegeUnknown && a.Read != matter.PrivilegeView {
		list.WriteString("read: ")
		list.WriteString(strings.ToLower(a.Read.String()))
	}
	if a.Write != matter.PrivilegeUnknown && a.Write != matter.PrivilegeView {
		if list.Len() > 0 {
			list.WriteString(", ")
		}
		list.WriteString("write: ")
		list.WriteString(strings.ToLower(a.Write.String()))
	}
	if a.Invoke != matter.PrivilegeUnknown && a.Invoke != matter.PrivilegeOperate {
		if list.Len() > 0 {
			list.WriteString(", ")
		}
		list.WriteString("invoke: ")
		list.WriteString(strings.ToLower(a.Invoke.String()))
	}

	return raymond.SafeString(list.String())
}

func ifRequestHelper(s *matter.Command, options *raymond.Options) string {
	if s.Direction == matter.InterfaceServer {
		return options.Fn()
	}
	return options.Inverse()
}

func storageOptionHelper(s string) raymond.SafeString {
	switch s {
	case "External":
		return "callback"
	case "RAM":
		return "ram     "
	case "NVM":
		return "persist "
	default:
		return raymond.SafeString(s)
	}
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

func deviceTypeNameHelper(s string) raymond.SafeString {
	return raymond.SafeString(strcase.ToSnake(s))
}
