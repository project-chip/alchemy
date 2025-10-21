package regen

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/mailgun/raymond/v2"
	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
	"github.com/project-chip/alchemy/provisional"
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
	t.RegisterHelper("isProvisional", ifProvisionalHelper(sp.spec))
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
	t.RegisterHelper("ifRequest", ifRequestHelper)

	t.RegisterHelper("storageOption", storageOptionHelper)

	t.RegisterHelper("accessString", accessStringHelper)
	t.RegisterHelper("ifHasLengthConstraint", ifHasLengthConstraintHelper)
	t.RegisterHelper("lengthConstraint", lengthConstraintHelper)
	t.RegisterHelper("bitName", bitNameHelper)
	t.RegisterHelper("bitMask", bitMaskHelper)
	t.RegisterHelper("bitmapType", bitmapTypeHelper)
	t.RegisterHelper("deviceTypeName", deviceTypeNameHelper)
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

func currentRevisionHelper(revisions []*matter.Revision) raymond.SafeString {
	var lastRevision uint64
	for _, rev := range revisions {
		revNumber := matter.ParseNumber(rev.Number)
		if revNumber.Valid() && revNumber.Value() > lastRevision {
			lastRevision = revNumber.Value()
		}
	}
	return raymond.SafeString(strconv.FormatUint(lastRevision, 10))

}

func asUpperCamelCaseHelper(value string) raymond.SafeString {
	return raymond.SafeString(caseify(value, false, true))
}

func asLowerCamelCaseHelper(value string) raymond.SafeString {
	if len(value) > 1 && text.IsUpperCase(value) {
		return raymond.SafeString(strings.ToLower(value))
	}
	return raymond.SafeString(caseify(value, true, true))
}

func upperCamelCaseDiffersHelper(value string, options *raymond.Options) string {
	if string(asUpperCamelCaseHelper(value)) != value {
		return options.Fn()
	}
	return options.Inverse()
}

func lowerCamelCaseDiffersHelper(value string, options *raymond.Options) string {
	if string(asLowerCamelCaseHelper(value)) != value {
		return options.Fn()
	}
	return options.Inverse()
}

func enumerateHelper[T types.Entity](list []T, spec *spec.Specification, options *raymond.Options) raymond.SafeString {
	var result strings.Builder
	for i, en := range list {
		df := options.DataFrame().Copy()
		df.Set("index", i)
		df.Set("key", nil)
		df.Set("first", i == 0)
		df.Set("last", i == len(list)-1)
		if spec != nil {
			refs, ok := spec.ClusterRefs.Get(en)
			if ok && refs.Size() > 1 {
				df.Set("shared", true)
			}
			is := provisional.Check(spec, en, en)
			switch is {
			case provisional.StateAllClustersProvisional,
				provisional.StateAllDataTypeReferencesProvisional,
				provisional.StateExplicitlyProvisional,
				provisional.StateUnreferenced:
				df.Set("provisional", true)
			default:
				df.Set("provisional", false)
			}

		}
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
		return enumerateHelper(sortedEnums, spec, options)
	}
}

func clusterBitmapsHelper(spec *spec.Specification) func(cluster matter.Cluster, options *raymond.Options) raymond.SafeString {
	return func(cluster matter.Cluster, options *raymond.Options) raymond.SafeString {
		bitmaps := make(matter.BitmapSet, 0, len(cluster.Bitmaps))
		bitmaps = append(bitmaps, cluster.Bitmaps...)
		slices.SortStableFunc(bitmaps, func(a *matter.Bitmap, b *matter.Bitmap) int {
			return strings.Compare(a.Name, b.Name)
		})
		if cluster.Features != nil {
			features := cluster.Features.Bitmap.Clone()
			// ZAP renames this for some reason
			features.Name = "Feature"
			bitmaps = append(matter.BitmapSet{features}, bitmaps...)
		}
		return enumerateHelper(bitmaps, spec, options)

	}
}

func clusterAttributesHelper(spec *spec.Specification, commonAttributes matter.FieldSet) func(cluster matter.Cluster, options *raymond.Options) raymond.SafeString {
	return func(cluster matter.Cluster, options *raymond.Options) raymond.SafeString {
		attributes := make(matter.FieldSet, 0, len(cluster.Attributes)+len(commonAttributes))
		attributes = append(attributes, cluster.Attributes...)
		for _, ca := range commonAttributes {
			ca = ca.Clone()
			ca.SetParent(&cluster)
			attributes = append(attributes, ca)
		}
		slices.SortStableFunc(attributes, func(a *matter.Field, b *matter.Field) int {
			return a.ID.Compare(b.ID)
		})
		return enumerateHelper(attributes, spec, options)
	}
}

func clusterStructsHelper(spec *spec.Specification) func(s matter.StructSet, options *raymond.Options) raymond.SafeString {
	return func(s matter.StructSet, options *raymond.Options) raymond.SafeString {
		structs := make(matter.StructSet, len(s))
		copy(structs, s)
		slices.SortStableFunc(structs, func(a *matter.Struct, b *matter.Struct) int {
			return strings.Compare(a.Name, b.Name)
		})
		return enumerateHelper(structs, spec, options)

	}
}

func clusterEventsHelper(spec *spec.Specification) func(cluster matter.Cluster, options *raymond.Options) raymond.SafeString {
	return func(cluster matter.Cluster, options *raymond.Options) raymond.SafeString {
		events := make(matter.EventSet, len(cluster.Events))
		copy(events, cluster.Events)
		return enumerateHelper(events, spec, options)
	}
}

func structFieldsHelper(spec *spec.Specification) func(s matter.Struct, options *raymond.Options) raymond.SafeString {
	return func(s matter.Struct, options *raymond.Options) raymond.SafeString {
		fields := make(matter.FieldSet, len(s.Fields))
		copy(fields, s.Fields)
		if s.FabricScoping == matter.FabricScopingScoped {
			fabricIndex := &matter.Field{ID: matter.NewNumber(254), Name: "FabricIndex", Type: types.NewDataType(types.BaseDataTypeFabricIndex, false), Conformance: conformance.Set{&conformance.Mandatory{}}}
			fabricIndex.SetParent(&s)
			fields = append(fields, fabricIndex)
		}
		return enumerateHelper(fields, spec, options)

	}
}

func eventFieldsHelper(spec *spec.Specification) func(e matter.Event, options *raymond.Options) raymond.SafeString {
	return func(e matter.Event, options *raymond.Options) raymond.SafeString {
		fields := make(matter.FieldSet, len(e.Fields))
		copy(fields, e.Fields)
		if e.Access.FabricSensitivity == matter.FabricSensitivitySensitive {
			fields = append(fields, &matter.Field{ID: matter.NewNumber(254), Name: "FabricIndex", Type: types.NewDataType(types.BaseDataTypeFabricIndex, false), Conformance: conformance.Set{&conformance.Mandatory{}}})
		}
		return enumerateHelper(fields, spec, options)
	}
}

func commandsHelper(spec *spec.Specification) func(commands matter.CommandSet, options *raymond.Options) raymond.SafeString {
	return func(commands matter.CommandSet, options *raymond.Options) raymond.SafeString {
		sortedCommands := make(matter.CommandSet, 0, len(commands))
		var requests []*matter.Command
		responses := make(map[*matter.Command]struct{})
		for _, command := range commands {
			switch command.Direction {
			case matter.InterfaceServer:
				requests = append(requests, command)
			case matter.InterfaceClient:
				responses[command] = struct{}{}
			}
		}
		slices.SortStableFunc(requests, func(a *matter.Command, b *matter.Command) int {
			return a.ID.Compare(b.ID)
		})
		for _, req := range requests {
			sortedCommands = append(sortedCommands, req)
			if req.Response != nil {
				switch response := req.Response.Entity.(type) {
				case *matter.Command:
					if _, unused := responses[response]; unused {
						sortedCommands = append(sortedCommands, response)
						delete(responses, response)
					}
				case nil:
				}
			}
		}
		return enumerateHelper(sortedCommands, spec, options)
	}

}

func commandFieldsHelper(spec *spec.Specification) func(e matter.Command, options *raymond.Options) raymond.SafeString {
	return func(e matter.Command, options *raymond.Options) raymond.SafeString {
		fields := make(matter.FieldSet, len(e.Fields))
		copy(fields, e.Fields)
		if e.Access.FabricSensitivity == matter.FabricSensitivitySensitive {
			fields = append(fields, &matter.Field{ID: matter.NewNumber(254), Name: "FabricIndex", Type: types.NewDataType(types.BaseDataTypeFabricIndex, false), Conformance: conformance.Set{&conformance.Mandatory{}}})
		}
		return enumerateHelper(fields, spec, options)
	}
}

func requestsHelper(spec *spec.Specification) func(commands matter.CommandSet, options *raymond.Options) raymond.SafeString {
	return func(commands matter.CommandSet, options *raymond.Options) raymond.SafeString {
		var requests []*matter.Command
		for _, command := range commands {
			switch command.Direction {
			case matter.InterfaceServer:
				requests = append(requests, command)
			}
		}
		slices.SortStableFunc(requests, func(a *matter.Command, b *matter.Command) int {
			return a.ID.Compare(b.ID)
		})
		return enumerateHelper(requests, spec, options)
	}

}

func isTimedHelper(command matter.Command, options *raymond.Options) string {
	if command.Access.Timing == matter.TimingTimed {
		return options.Fn()
	} else {
		return options.Inverse()
	}
}

func ifProvisionalHelper(spec *spec.Specification) func(entity types.Entity, options *raymond.Options) string {
	return func(entity types.Entity, options *raymond.Options) string {
		is := provisional.Check(spec, entity, entity)
		switch is {
		case provisional.StateAllClustersProvisional,
			provisional.StateAllDataTypeReferencesProvisional,
			provisional.StateExplicitlyProvisional,
			provisional.StateUnreferenced:
			return options.Fn()
		default:
			return options.Inverse()
		}
	}

}

func requestNameHelper(command matter.Command) raymond.SafeString {
	return raymond.SafeString(command.Name)
}

func responseNameHelper(command matter.Command) raymond.SafeString {
	if command.Response != nil {
		switch response := command.Response.Entity.(type) {
		case *matter.Command:
			return raymond.SafeString(response.Name)
		default:
		}
	}
	return raymond.SafeString("DefaultSuccess")
}

func fieldTypeHelper(field matter.Field, fs matter.FieldSet, options *raymond.Options) raymond.SafeString {
	return raymond.SafeString(zap.FieldToZapDataType(fs, &field, field.Constraint))
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
	if field.Type == nil {
		return
	}
	if !field.Type.HasLength() {
		return
	}

	if field.Constraint == nil {
		return
	}
	_, max = zap.GetMinMax(matter.NewConstraintContext(&field, fs), field.Constraint)
	if max.IsNumeric() {
		ok = true
	}
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

func bitmapTypeHelper(bitmap matter.Bitmap) raymond.SafeString {
	switch bitmap.Type.BaseType {
	case types.BaseDataTypeMap8:
		return raymond.SafeString("bitmap8")
	case types.BaseDataTypeMap16:
		return raymond.SafeString("bitmap16")
	case types.BaseDataTypeMap32:
		return raymond.SafeString("bitmap32")
	case types.BaseDataTypeMap64:
		return raymond.SafeString("bitmap24")
	default:
		return raymond.SafeString("unknown bitmap type")

	}
}

func bitNameHelper(bit any) raymond.SafeString {
	switch bit := bit.(type) {
	case matter.BitmapBit:
		return raymond.SafeString(bit.Name())
	case matter.Feature:
		return raymond.SafeString(bit.Name())
	default:
		return raymond.SafeString(fmt.Sprintf("unexpected bitName type: %T", bit))
	}
}

func bitMaskHelper(bit any) raymond.SafeString {
	switch bit := bit.(type) {
	case matter.BitmapBit:
		mask, err := bit.Mask()
		if err != nil {
			return raymond.SafeString(fmt.Sprintf("error converting bitmap mask: %v", err))
		}
		return raymond.SafeString(fmt.Sprintf("0x%s", strconv.FormatUint(mask, 16)))
	case matter.Feature:
		mask, err := bit.Mask()
		if err != nil {
			return raymond.SafeString(fmt.Sprintf("error converting feature mask: %v", err))
		}
		return raymond.SafeString(fmt.Sprintf("0x%s", strconv.FormatUint(mask, 16)))
	default:
		return raymond.SafeString(fmt.Sprintf("unexpected bitName type: %T", bit))
	}
}

func deviceTypeNameHelper(s string) raymond.SafeString {
	return raymond.SafeString(strcase.ToSnake(s))
}
