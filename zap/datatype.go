package zap

import (
	"log/slog"
	"strings"

	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/constraint"
	"github.com/project-chip/alchemy/matter/types"
)

var matterToZapMap = map[string]string{
	"bool": "boolean",

	"uint8":  "int8u",
	"uint16": "int16u",
	"uint24": "int24u",
	"uint32": "int32u",
	"uint40": "int40u",
	"uint48": "int48u",
	"uint56": "int56u",
	"uint64": "int64u",

	"int8":  "int8s",
	"int16": "int16s",
	"int24": "int24s",
	"int32": "int32s",
	"int40": "int40s",
	"int48": "int48s",
	"int56": "int56s",
	"int64": "int64s",

	"enum8":  "enum8",
	"enum16": "enum16",
	"enum32": "enum32",

	"map8":  "bitmap8",
	"map16": "bitmap16",
	"map32": "bitmap32",
	"map64": "bitmap64",

	"string":       "char_string",
	"octstr":       "octet_string",
	"ref_tempdiff": "int16s",
	"amperage-ma":  "amperage_ma",
	"voltage-mv":   "voltage_mv",
	"power-mw":     "power_mw",
	"energy-mwh":   "energy_mwh",

	"elapsed-s":  "elapsed_s",
	"epoch-s":    "epoch_s",
	"epoch-us":   "epoch_us",
	"systime-ms": "systime_ms",
	"systime-us": "systime_us",
	"posix-ms":   "posix_ms",
	"utc":        "epoch_s", //Deprecated

	"action-id":      "action_id",
	"attrib-id":      "attrib_id",
	"attribute-id":   "attrib_id",
	"cluster-id":     "cluster_id",
	"command-id":     "command_id",
	"data-ver":       "data_ver",
	"devtype-id":     "devtype_id",
	"entry-idx":      "entry_idx",
	"event-id":       "event_id",
	"event-no":       "event_no",
	"fabric-id":      "fabric_id",
	"fabric-idx":     "fabric_idx",
	"field-id":       "field_id",
	"group-id":       "group_id",
	"node-id":        "node_id",
	"transaction-id": "transaction_id",
	"vendor-id":      "vendor_id",
	"endpoint-id":    "endpoint_id",
	"endpoint-no":    "endpoint_no",
	"eui64":          "eui64",

	"unsignedtemperature":   "int8u",
	"signedtemperature":     "int8s",
	"temperaturedifference": "int16s",

	"subjectid": "int64u",

	/* Same on both sides:
	percent
	percent100ths
	tod
	date
	temperature

	Not sure:
	ipadr
	ipv4adr
	ipv6adr
	ipv6pre
	hwadr
	semtag
	namespace
	tag

	*/
}

var zapToMatterMap map[string]string

func init() {
	zapToMatterMap = make(map[string]string, len(matterToZapMap))
	for k, v := range matterToZapMap {
		zapToMatterMap[strings.ToLower(v)] = k
	}
}

func DataTypeName(dataType *types.DataType) string {
	if dataType.Entity != nil {
		switch e := dataType.Entity.(type) {
		case *matter.Bitmap:
			return e.Name
		case *matter.Enum:
			return e.Name
		case *matter.Struct:
			return e.Name
		case *matter.TypeDef:
			dataType = e.Type
		}
	}
	switch dataType.BaseType {
	case types.BaseDataTypeBoolean:
		return "boolean"
	case types.BaseDataTypeMap8:
		return "bitmap8"
	case types.BaseDataTypeMap16:
		return "bitmap16"
	case types.BaseDataTypeMap32:
		return "bitmap32"
	case types.BaseDataTypeMap64:
		return "bitmap64"
	case types.BaseDataTypeUInt8:
		return "int8u"
	case types.BaseDataTypeUInt16:
		return "int16u"
	case types.BaseDataTypeUInt24:
		return "int24u"
	case types.BaseDataTypeUInt32:
		return "int32u"
	case types.BaseDataTypeUInt40:
		return "int40u"
	case types.BaseDataTypeUInt48:
		return "int48u"
	case types.BaseDataTypeUInt56:
		return "int56u"
	case types.BaseDataTypeUInt64:
		return "int64u"
	case types.BaseDataTypeInt8:
		return "int8s"
	case types.BaseDataTypeInt16:
		return "int16s"
	case types.BaseDataTypeInt24:
		return "int24s"
	case types.BaseDataTypeInt32:
		return "int32s"
	case types.BaseDataTypeInt40:
		return "int40s"
	case types.BaseDataTypeInt48:
		return "int48s"
	case types.BaseDataTypeInt56:
		return "int56s"
	case types.BaseDataTypeInt64:
		return "int64s"
	case types.BaseDataTypeSingle:
		return "single"
	case types.BaseDataTypeDouble:
		return "double"
	case types.BaseDataTypeOctStr:
		return "octet_string"
	case types.BaseDataTypePercent:
		return "percent"
	case types.BaseDataTypePercentHundredths:
		return "percent100ths"
	case types.BaseDataTypeTimeOfDay:
		return "tod"
	case types.BaseDataTypeDate:
		return "date"
	case types.BaseDataTypeEpochMicroseconds:
		return "epoch_us"
	case types.BaseDataTypeEpochSeconds:
		return "epoch_s"
	case types.BaseDataTypePosixMilliseconds:
		return "posix_ms"
	case types.BaseDataTypeSystimeMicroseconds:
		return "systime_us"
	case types.BaseDataTypeSystimeMilliseconds:
		return "systime_ms"
	case types.BaseDataTypeElapsedSeconds:
		return "elapsed_s"
	case types.BaseDataTypeTemperature:
		return "temperature"
	case types.BaseDataTypeAmperage:
		return "amperage_ma"
	case types.BaseDataTypeVoltage:
		return "voltage_mv"
	case types.BaseDataTypePower:
		return "power_mw"
	case types.BaseDataTypeEnergy:
		return "energy_mwh"
	case types.BaseDataTypeApparentPower:
		return "power_mva"
	case types.BaseDataTypeApparentEnergy:
		return "energy_mvah"
	case types.BaseDataTypeReactivePower:
		return "power_mvar"
	case types.BaseDataTypeReactiveEnergy:
		return "energy_mvarh"
	case types.BaseDataTypeMoney:
		return "money"
	case types.BaseDataTypeTemperatureDifference:
		return "int16s"
	case types.BaseDataTypeUnsignedTemperature:
		return "int8u"
	case types.BaseDataTypeSignedTemperature:
		return "int8s"
	case types.BaseDataTypeEnum8:
		return "enum8"
	case types.BaseDataTypeEnum16:
		return "enum16"
	case types.BaseDataTypeGroupID:
		return "group_id"
	case types.BaseDataTypeEndpointID:
		return "endpoint_no"
	case types.BaseDataTypeEndpointNumber:
		return "endpoint_no"
	case types.BaseDataTypeVendorID:
		return "vendor_id"
	case types.BaseDataTypeDeviceTypeID:
		return "devtype_id"
	case types.BaseDataTypeFabricID:
		return "fabric_id"
	case types.BaseDataTypeFabricIndex:
		return "fabric_idx"
	case types.BaseDataTypeClusterID:
		return "cluster_id"
	case types.BaseDataTypeAttributeID:
		return "attrib_id"
	case types.BaseDataTypeFieldID:
		return "field_id"
	case types.BaseDataTypeEventID:
		return "event_id"
	case types.BaseDataTypeCommandID:
		return "command_id"
	case types.BaseDataTypeActionID:
		return "action_id"
	case types.BaseDataTypeSubjectID:
		return "int64u"
	case types.BaseDataTypeTransactionID:
		return "transaction_id"
	case types.BaseDataTypeNodeID:
		return "node_id"
	case types.BaseDataTypeIeeeAddress:
		return "ieee_address"
	case types.BaseDataTypeEntryIndex:
		return "entry_idx"
	case types.BaseDataTypeDataVersion:
		return "data_ver"
	case types.BaseDataTypeEventNumber:
		return "event_no"
	case types.BaseDataTypeString:
		return "char_string"
	case types.BaseDataTypeIPAddress:
		return "octet_string"
	case types.BaseDataTypeIPv4Address:
		return "octet_string"
	case types.BaseDataTypeIPv6Address:
		return "octet_string"
	case types.BaseDataTypeIPv6Prefix:
		return "octet_string"
	case types.BaseDataTypeHardwareAddress:
		return "octet_string"
	case types.BaseDataTypeSemanticTag:
		return "semtag"
	case types.BaseDataTypeNamespaceID:
		return "namespace"
	case types.BaseDataTypeTag:
		return "tag"
	case types.BaseDataTypeMessageID:
		return "octet_string"
	case types.BaseDataTypeStatus:
		return "status"
	}
	return dataType.Name
}

func ToBaseDataType(s string) types.BaseDataType {
	switch strings.ToLower(s) {
	case "boolean":
		return types.BaseDataTypeBoolean
	case "bitmap8":
		return types.BaseDataTypeMap8
	case "bitmap16":
		return types.BaseDataTypeMap16
	case "bitmap32":
		return types.BaseDataTypeMap32
	case "bitmap64":
		return types.BaseDataTypeMap64
	case "int8u":
		return types.BaseDataTypeUInt8
	case "int16u":
		return types.BaseDataTypeUInt16
	case "int24u":
		return types.BaseDataTypeUInt24
	case "int32u":
		return types.BaseDataTypeUInt32
	case "int40u":
		return types.BaseDataTypeUInt40
	case "int48u":
		return types.BaseDataTypeUInt48
	case "int56u":
		return types.BaseDataTypeUInt56
	case "int64u":
		return types.BaseDataTypeUInt64
	case "int8s":
		return types.BaseDataTypeInt8
	case "int16s":
		return types.BaseDataTypeInt16
	case "int24s":
		return types.BaseDataTypeInt24
	case "int32s":
		return types.BaseDataTypeInt32
	case "int40s":
		return types.BaseDataTypeInt40
	case "int48s":
		return types.BaseDataTypeInt48
	case "int56s":
		return types.BaseDataTypeInt56
	case "int64s":
		return types.BaseDataTypeInt64
	case "single":
		return types.BaseDataTypeSingle
	case "double":
		return types.BaseDataTypeDouble
	case "octet_string":
		return types.BaseDataTypeOctStr
	case "percent":
		return types.BaseDataTypePercent
	case "percent100ths":
		return types.BaseDataTypePercentHundredths
	case "tod":
		return types.BaseDataTypeTimeOfDay
	case "date":
		return types.BaseDataTypeDate
	case "epoch_us":
		return types.BaseDataTypeEpochMicroseconds
	case "epoch_s":
		return types.BaseDataTypeEpochSeconds
	case "posix_ms":
		return types.BaseDataTypePosixMilliseconds
	case "systime_us":
		return types.BaseDataTypeSystimeMicroseconds
	case "systime_ms":
		return types.BaseDataTypeSystimeMilliseconds
	case "elapsed_s":
		return types.BaseDataTypeElapsedSeconds
	case "temperature":
		return types.BaseDataTypeTemperature
	case "amperage_ma":
		return types.BaseDataTypeAmperage
	case "voltage_mv":
		return types.BaseDataTypeVoltage
	case "power_mw":
		return types.BaseDataTypePower
	case "energy_mwh":
		return types.BaseDataTypeEnergy
	case "money":
		return types.BaseDataTypeMoney
	case "enum8":
		return types.BaseDataTypeEnum8
	case "enum16":
		return types.BaseDataTypeEnum16
	case "group_id":
		return types.BaseDataTypeGroupID
	case "endpoint_id":
		return types.BaseDataTypeEndpointID
	case "endpoint_no":
		return types.BaseDataTypeEndpointNumber
	case "vendor_id":
		return types.BaseDataTypeVendorID
	case "devtype_id":
		return types.BaseDataTypeDeviceTypeID
	case "fabric_id":
		return types.BaseDataTypeFabricID
	case "fabric_idx":
		return types.BaseDataTypeFabricIndex
	case "cluster_id":
		return types.BaseDataTypeClusterID
	case "attrib_id":
		return types.BaseDataTypeAttributeID
	case "field_id":
		return types.BaseDataTypeFieldID
	case "event_id":
		return types.BaseDataTypeEventID
	case "command_id":
		return types.BaseDataTypeCommandID
	case "action_id":
		return types.BaseDataTypeActionID
	case "transaction_id":
		return types.BaseDataTypeTransactionID
	case "node_id":
		return types.BaseDataTypeNodeID
	case "ieee_address":
		return types.BaseDataTypeIeeeAddress
	case "entry_idx":
		return types.BaseDataTypeEntryIndex
	case "data_ver":
		return types.BaseDataTypeDataVersion
	case "event_no":
		return types.BaseDataTypeEventNumber
	case "char_string", "long_char_string":
		return types.BaseDataTypeString
	case "ipadr":
		return types.BaseDataTypeIPAddress
	case "ipv4addr":
		return types.BaseDataTypeIPv4Address
	case "ipv6addr":
		return types.BaseDataTypeIPv6Address
	case "ipv6pre":
		return types.BaseDataTypeIPv6Prefix
	case "hwadr":
		return types.BaseDataTypeHardwareAddress
	case "semtag":
		return types.BaseDataTypeSemanticTag
	case "namespace":
		return types.BaseDataTypeNamespaceID
	case "tag":
		return types.BaseDataTypeTag
	case "array":
		return types.BaseDataTypeList
	case "status":
		return types.BaseDataTypeStatus
	}
	if len(s) > 0 {
		return types.BaseDataTypeCustom
	}
	return types.BaseDataTypeUnknown
}

func maxOver255Bytes(fs matter.FieldSet, f *matter.Field, constraint constraint.Constraint) bool {
	if constraint == nil {
		return false
	}
	max := constraint.Max(matter.NewConstraintContext(f, fs))

	switch max.Type {
	case types.DataTypeExtremeTypeInt64:
		if max.Int64 > 255 {
			return true
		}
	case types.DataTypeExtremeTypeUInt64:
		if max.UInt64 > 255 {
			return true
		}
	}
	return false
}

func FieldToZapDataType(fs matter.FieldSet, f *matter.Field, constraint constraint.Constraint) string {
	if f.Type == nil {
		return ""
	}

	if f.Type.BaseType == types.BaseDataTypeString && maxOver255Bytes(fs, f, constraint) {
		// Special case; needs to be long_char_string if over 255
		return "long_char_string"
	}
	if f.Type.BaseType == types.BaseDataTypeOctStr && maxOver255Bytes(fs, f, constraint) {
		// Special case; needs to be long_octet_string if over 255
		return "long_octet_string"
	}
	if f.Type.IsArray() {
		return DataTypeName(f.Type.EntryType)
	}

	return DataTypeName(f.Type)
}

func GetMinMax(cc *matter.ConstraintContext, c constraint.Constraint) (from types.DataTypeExtreme, to types.DataTypeExtreme) {
	if cc.Field.Type == nil {
		return
	}

	from, to = minMaxFromConstraint(cc, c)

	if from.Defined() || to.Defined() {
		return
	}

	from, to = minMaxFromEntity(cc)

	if from.Defined() || to.Defined() {
		return
	}

	if cc.Field.Access.Write != matter.PrivilegeUnknown {
		// Writable fields get default min/max
		var nullability types.Nullability
		if cc.Field.Quality.Has(matter.QualityNullable) {
			nullability = types.NullabilityNullable
		}
		from = cc.Field.Type.Min(nullability)
		to = cc.Field.Type.Max(nullability)

	}
	return
}

func minMaxFromEntity(cc *matter.ConstraintContext) (from types.DataTypeExtreme, to types.DataTypeExtreme) {
	if cc.Field.Type.Entity != nil {
		switch m := cc.Field.Type.Entity.(type) {
		case *matter.Enum:
			if len(m.Values) > 0 {
				var f, t uint64
				for _, v := range m.Values {
					if v.Value.Valid() {
						val := v.Value.Value()
						f = min(f, val)
						t = max(t, val)
					}
				}
				from = types.NewUintDataTypeExtreme(f, types.NumberFormatHex)
				to = types.NewUintDataTypeExtreme(t, types.NumberFormatHex)
				return
			}
		case *matter.Bitmap:
			if len(m.Bits) > 0 {
				var t uint64
				for _, b := range m.Bits {
					mask, err := b.Mask()
					if err != nil {
						return
					}
					t |= mask
				}
				from = types.NewUintDataTypeExtreme(0, types.NumberFormatHex)
				to = types.NewUintDataTypeExtreme(t, types.NumberFormatHex)
				return
			}
		}
	}
	return
}

func minMaxFromConstraint(cc *matter.ConstraintContext, c constraint.Constraint) (from types.DataTypeExtreme, to types.DataTypeExtreme) {
	if c == nil {
		return
	}
	if cc.Field.Type.IsArray() && !isValidArrayConstraint(cc.Field, c) {
		return
	}
	from = c.Min(cc)
	to = c.Max(cc)
	return
}

func isValidArrayConstraint(field *matter.Field, cons constraint.Constraint) bool {
	switch cons := cons.(type) {
	case constraint.Set:
		for _, c := range cons {
			if !isValidArrayConstraint(field, c) {
				return false
			}
		}
	case *constraint.DescribedConstraint, *constraint.AllConstraint, *constraint.GenericConstraint:
		return false
	case *constraint.ListConstraint, *constraint.MaxConstraint, *constraint.MinConstraint, *constraint.RangeConstraint, *constraint.ExactConstraint:
	default:
		slog.Warn("Array field has constraint not compatible with arrays", "field", field.Name, "constraint", cons)
		return false
	}
	return true
}

func GetFallbackValue(cc *matter.ConstraintContext, fallback constraint.Limit) (fallbackValue types.DataTypeExtreme) {
	fallbackValue = fallback.Fallback(cc)
	switch fallbackValue.Type {
	case types.DataTypeExtremeTypeEmptyList:
		if !cc.Field.Type.HasLength() {
			fallbackValue = types.DataTypeExtreme{}
		}
	case types.DataTypeExtremeTypeNull:
		if cc.Field.Type.NullValue() == 0 {
			fallbackValue = types.DataTypeExtreme{}
		}
	case types.DataTypeExtremeTypeInt64, types.DataTypeExtremeTypeUInt64:
		if cc.Field.Type != nil {
			switch cc.Field.Type.Entity.(type) {
			case *matter.Bitmap, *matter.Enum:
				fallbackValue.Format = types.NumberFormatHex
			}
		}
	}

	return
}
