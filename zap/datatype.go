package zap

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/hasty/alchemy/matter"
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

	"string":        "char_string",
	"octstr":        "octet_string",
	"percent":       "int8u",
	"percent100ths": "int16u",
	"ref_tempdiff":  "int16s",
	"temperature":   "int16s",
	"amperage-ma":   "int64",
	"voltage-mv":    "int64",
	"power-mw":      "int64",
	"energy-mwh":    "int64",

	"elapsed-s":  "elapsed_s",
	"epoch-s":    "epoch_s",
	"epoch-us":   "epoch_us",
	"systime_ms": "systime_ms",
	"systime_us": "systime_us",
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

func ConvertDataTypeToZap(s string) string {
	if z, ok := matterToZapMap[strings.ToLower(s)]; ok {
		return z
	}
	return s
}

func ConvertZapToDataType(s string) string {
	if z, ok := zapToMatterMap[strings.ToLower(s)]; ok {
		return z
	}
	return s
}

func GetMinMax(fs matter.FieldSet, f *matter.Field) (from matter.ConstraintExtreme, to matter.ConstraintExtreme) {
	if f.Type == nil || f.Type.IsArray {
		return
	}
	if f.Constraint != nil {
		from, to = f.Constraint.MinMax(&matter.ConstraintContext{Fields: fs})
		return
	}
	if !from.Defined() || !to.Defined() {
		f, t := f.Type.MinMax(f.Quality.Has(matter.QualityNullable))
		if !from.Defined() && f.Defined() {
			from = f
		}
		if !to.Defined() && t.Defined() {
			to = t
		}
	}
	return
}

func FormatConstraintValue(val any) string {
	switch v := val.(type) {
	case uint64:
		if v > 0xFF {
			return "0x" + strconv.FormatUint(v, 16)
		}
		return strconv.FormatUint(v, 10)
	case int64:
		if v > 255 || v < 256 {
			return "0x" + strconv.FormatUint(uint64(v), 16)
		}
		return strconv.FormatInt(v, 10)
	default:
		return fmt.Sprintf("%d", val)
	}
}
