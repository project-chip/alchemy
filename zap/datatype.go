package zap

import (
	"log/slog"
	"strings"

	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/constraint"
	"github.com/hasty/alchemy/matter/types"
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

func DataTypeToZap(dataType *types.DataType) string {
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
		return "endpoint_id"
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
		return "ipadr"
	case types.BaseDataTypeIPv4Address:
		return "ipv4addr"
	case types.BaseDataTypeIPv6Address:
		return "ipv6addr"
	case types.BaseDataTypeIPv6Prefix:
		return "ipv6pre"
	case types.BaseDataTypeHardwareAddress:
		return "hwadr"
	case types.BaseDataTypeSemanticTag:
		return "semtag"
	case types.BaseDataTypeNamespace:
		return "namespace"
	case types.BaseDataTypeTag:
		return "tag"
	case types.BaseDataTypeMessageID:
		return "octet_string"
	}
	return dataType.Name
}

func ConvertDataTypeNameToZap(s string) string {
	if z, ok := matterToZapMap[strings.ToLower(s)]; ok {
		return z
	}
	return s
}

func ConvertZapToDataTypeName(s string) string {
	if z, ok := zapToMatterMap[strings.ToLower(s)]; ok {
		return z
	}
	return s
}

func FieldToZapDataType(fs matter.FieldSet, f *matter.Field) string {
	if f.Type == nil {
		return ""
	}
	if f.Type.BaseType == types.BaseDataTypeString && f.Constraint != nil {
		// Special case; needs to be long_char_string if over 255
		max := f.Constraint.Max(&matter.ConstraintContext{Field: f, Fields: fs})
		switch max.Type {
		case types.DataTypeExtremeTypeInt64:
			if max.Int64 > 255 {
				return "long_char_string"
			}
		case types.DataTypeExtremeTypeUInt64:
			if max.UInt64 > 255 {
				return "long_char_string"
			}
		}
	}
	if f.Type.IsArray() {
		return DataTypeToZap(f.Type.EntryType)
	}
	return DataTypeToZap(f.Type)
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
		isNullable := cc.Field.Quality.Has(matter.QualityNullable)
		from = cc.Field.Type.Min(isNullable)
		to = cc.Field.Type.Max(isNullable)

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
	if cc.Field.Type.IsArray() {
		switch cc.Field.Constraint.(type) {
		case *constraint.DescribedConstraint, *constraint.AllConstraint, *constraint.GenericConstraint:
			return
		case *constraint.ListConstraint, *constraint.MaxConstraint, *constraint.MinConstraint, *constraint.RangeConstraint, *constraint.ExactConstraint:
		default:
			slog.Warn("Array field has constraint not compatible with arrays", "field", cc.Field.Name, "constraint", cc.Field.Constraint)
			return
		}
	}
	from = c.Min(cc)
	to = c.Max(cc)
	return
}

func GetDefaultValue(cc *matter.ConstraintContext) (defaultValue types.DataTypeExtreme) {
	c, err := constraint.ParseString(cc.Field.Default)
	if err != nil {
		c = &constraint.GenericConstraint{Value: cc.Field.Default}
	}
	defaultValue = c.Default(cc)
	switch defaultValue.Type {
	case types.DataTypeExtremeTypeEmpty:
		if !cc.Field.Type.HasLength() {
			defaultValue = types.DataTypeExtreme{}
		}
	case types.DataTypeExtremeTypeNull:
		if cc.Field.Type.NullValue() == 0 {
			defaultValue = types.DataTypeExtreme{}
		}
	case types.DataTypeExtremeTypeInt64, types.DataTypeExtremeTypeUInt64:
		if cc.Field.Type != nil {
			switch cc.Field.Type.Entity.(type) {
			case *matter.Bitmap, *matter.Enum:
				defaultValue.Format = types.NumberFormatHex
			}
		}
	}

	return
}
