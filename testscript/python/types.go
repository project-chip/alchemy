package python

import (
	"log/slog"

	"github.com/mailgun/raymond/v2"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
	"github.com/project-chip/alchemy/sdk"
	"github.com/project-chip/alchemy/testscript"
)

func typeCheckIsHelper(action testscript.CheckType, is string, options *raymond.Options) string {
	ok := checkUnderlyingType(action.Field.Type, is)
	if ok {
		return options.Fn()
	}
	return options.Inverse()
}

func entryTypeCheckIsHelper(action testscript.CheckType, is string, options *raymond.Options) string {
	if !action.Field.Type.IsArray() {
		return options.Inverse()
	}
	ok := checkUnderlyingType(action.Field.Type.EntryType, is)
	if ok {
		return options.Fn()
	}
	return options.Inverse()
}

func checkUnderlyingType(dataType *types.DataType, is string) bool {
	var ok bool
	underlyingType := sdk.ToUnderlyingType(dataType.BaseType)
	switch is {
	case "uint64":
		ok = underlyingType == types.BaseDataTypeUInt64
	case "uint32":
		ok = underlyingType == types.BaseDataTypeUInt32
	case "uint24":
		ok = underlyingType == types.BaseDataTypeUInt24
	case "uint16":
		ok = underlyingType == types.BaseDataTypeUInt16
	case "uint8":
		ok = underlyingType == types.BaseDataTypeUInt8
	case "int64":
		ok = underlyingType == types.BaseDataTypeInt64
	case "int32":
		ok = underlyingType == types.BaseDataTypeInt32
	case "int24":
		ok = underlyingType == types.BaseDataTypeInt24
	case "int16":
		ok = underlyingType == types.BaseDataTypeInt16
	case "int8":
		ok = underlyingType == types.BaseDataTypeInt8
	case "map64":
		ok = underlyingType == types.BaseDataTypeMap64
	case "map32":
		ok = underlyingType == types.BaseDataTypeMap32
	case "map16":
		ok = underlyingType == types.BaseDataTypeMap16
	case "map8":
		ok = underlyingType == types.BaseDataTypeMap8
	case "string":
		ok = underlyingType == types.BaseDataTypeString
	case "list":
		ok = underlyingType == types.BaseDataTypeList
	case "percent":
		ok = dataType.BaseType == types.BaseDataTypePercent
	case "percent100ths":
		ok = dataType.BaseType == types.BaseDataTypePercentHundredths
	case "octstr":
		ok = dataType.BaseType == types.BaseDataTypeOctStr
	case "single":
		ok = dataType.BaseType == types.BaseDataTypeSingle
	case "double":
		ok = dataType.BaseType == types.BaseDataTypeDouble
	case "bitmap":
		if dataType.Entity != nil {
			_, ok = dataType.Entity.(*matter.Bitmap)
		}
	case "enum":
		if dataType.Entity != nil {
			_, ok = dataType.Entity.(*matter.Enum)
		}
	case "struct":
		if dataType.Entity != nil {
			_, ok = dataType.Entity.(*matter.Struct)
		}
	case "bool":
		ok = underlyingType == types.BaseDataTypeBoolean
	case "typedef":
		ok = underlyingType == types.BaseDataTypeCustom
		if ok {
			ok = dataType.Entity != nil
			if ok {
				_, ok = dataType.Entity.(*matter.TypeDef)
			}
		}
	case "custom":
		ok = underlyingType == types.BaseDataTypeCustom
		if ok {
			ok = dataType.Entity != nil
			if ok {
				_, ok = dataType.Entity.(*matter.TypeDef)
				ok = !ok
			}
		}
	case "vendor-id":
		ok = underlyingType == types.BaseDataTypeVendorID
	case "node-id":
		ok = underlyingType == types.BaseDataTypeNodeID
	case "endpoint-no":
		ok = underlyingType == types.BaseDataTypeEndpointNumber
	case "devtype-id":
		ok = underlyingType == types.BaseDataTypeDeviceTypeID
	case "group-id":
		ok = underlyingType == types.BaseDataTypeGroupID
	case "subject-id":
		ok = underlyingType == types.BaseDataTypeSubjectID
	case "cluster-id":
		ok = underlyingType == types.BaseDataTypeClusterID
	case "fabric-id":
		ok = underlyingType == types.BaseDataTypeFabricID
	case "fabric-idx":
		ok = underlyingType == types.BaseDataTypeFabricIndex
	case "ipadr":
		ok = underlyingType == types.BaseDataTypeIPAddress
	case "ipv6pre":
		ok = underlyingType == types.BaseDataTypeIPv6Prefix
	case "ipv6adr":
		ok = underlyingType == types.BaseDataTypeIPv6Address
	case "ipv4adr":
		ok = underlyingType == types.BaseDataTypeIPv4Address
	case "hwaddr":
		ok = underlyingType == types.BaseDataTypeHardwareAddress
	case "message-id":
		ok = underlyingType == types.BaseDataTypeMessageID
	case "enum16":
		ok = underlyingType == types.BaseDataTypeEnum16
	case "enum8":
		ok = underlyingType == types.BaseDataTypeEnum8
	case "tag":
		ok = underlyingType == types.BaseDataTypeTag
	case "namespace":
		ok = underlyingType == types.BaseDataTypeNamespaceID
	}
	return ok
}

func unimplementedTypeCheckHelper(step testscript.TestStep, action testscript.CheckType) raymond.SafeString {
	slog.Error("Unimplemented type check", slog.String("fieldName", action.Field.Name), slog.String("type", action.Field.Type.BaseType.String()), log.Path("source", action.Field))
	return raymond.SafeString("Unimplemented type check")
}
