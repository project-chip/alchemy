package zap

import "strings"

var matterToZapMap = map[string]string{
	"bool":        "boolean",
	"uint8":       "int8u",
	"uint16":      "int16u",
	"uint32":      "int32u",
	"uint64":      "int64u",
	"int8":        "int8s",
	"int16":       "int16s",
	"int32":       "int32s",
	"int64":       "in64s",
	"enum8":       "enum8",
	"enum16":      "enum16",
	"enum32":      "enum32",
	"map8":        "bitmap8",
	"map16":       "bitmap16",
	"map32":       "bitmap32",
	"map64":       "bitmap64",
	"string":      "char_string",
	"octstr":      "octet_string",
	"elapsed-s":   "elapsed_s",
	"epoch-s":     "epoch_s",
	"epoch-us":    "epoch_us",
	"fabric-idx":  "fabric_idx",
	"node-id":     "node_id",
	"vendor-id":   "vendor_id",
	"group-id":    "group_id",
	"endpoint-id": "endpoint_id",
	"endpoint-no": "endpoint_no",
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
