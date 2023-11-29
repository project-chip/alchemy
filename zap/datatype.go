package zap

import (
	"log/slog"
	"strings"

	"github.com/hasty/alchemy/constraint"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/parse"
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
	"amperage-ma":  "int64s",
	"voltage-mv":   "int64s",
	"power-mw":     "int64s",
	"energy-mwh":   "int64s",

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
	"EUI64":          "eui64",

	"unsignedtemperature":   "int8u",
	"signedtemperature":     "int8s",
	"temperaturedifference": "int16s",

	// Hacky workaround for Thermostat
	"ref_occupancybitmap": "bitmap8",

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
	if f.Type.BaseType == matter.BaseDataTypeString && f.Constraint != nil {
		// Special case; needs to be long_char_string if over 255
		max := f.Constraint.Max(&matter.ConstraintContext{Field: f, Fields: fs})
		switch max.Type {
		case matter.ConstraintExtremeTypeInt64:
			if max.Int64 > 255 {
				return "long_char_string"
			}
		case matter.ConstraintExtremeTypeUInt64:
			if max.UInt64 > 255 {
				return "long_char_string"
			}
		}
	}
	return ConvertDataTypeNameToZap(f.Type.Name)
}

func GetMinMax(cc *matter.ConstraintContext) (from matter.ConstraintExtreme, to matter.ConstraintExtreme) {
	if cc.Field.Type == nil {
		return
	}
	from, to = minMaxFromConstraint(cc)

	if from.Defined() || to.Defined() {
		return
	}

	from, to = minMaxFromModel(cc)
	return
}

func minMaxFromModel(cc *matter.ConstraintContext) (from matter.ConstraintExtreme, to matter.ConstraintExtreme) {
	if cc.Field.Type.Model != nil {
		switch m := cc.Field.Type.Model.(type) {
		case *matter.Enum:
			if len(m.Values) > 0 {
				var f, t uint64
				for _, v := range m.Values {
					val, err := parse.HexOrDec(v.Value)
					if err == nil {
						f = min(f, val)
						t = max(t, val)
					}
				}
				from = matter.NewUintConstraintExtreme(f, matter.ConstraintExtremeFormatHex)
				to = matter.NewUintConstraintExtreme(t, matter.ConstraintExtremeFormatHex)
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
				from = matter.NewUintConstraintExtreme(0, matter.ConstraintExtremeFormatHex)
				to = matter.NewUintConstraintExtreme(t, matter.ConstraintExtremeFormatHex)
				return
			}
		}
	}
	return
}

func minMaxFromConstraint(cc *matter.ConstraintContext) (from matter.ConstraintExtreme, to matter.ConstraintExtreme) {
	if cc.Field.Constraint != nil {
		if cc.Field.Type.IsArray() {
			switch cc.Field.Constraint.(type) {
			case *constraint.DescribedConstraint, *constraint.AllConstraint:
				return
			case *constraint.ListConstraint, *constraint.MaxConstraint, *constraint.MinConstraint, *constraint.RangeConstraint, *constraint.ExactConstraint:
			default:
				slog.Warn("Array field has constraint not compatible with arrays", "field", cc.Field.Name)
				return
			}
		}
		from = cc.Field.Constraint.Min(cc)
		to = cc.Field.Constraint.Max(cc)
		return
	}
	return
}
