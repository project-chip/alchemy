package zap

import "strings"

var matterToZapMap = map[string]string{
	"bool":        "boolean",
	"uint8":       "INT8U",
	"uint16":      "INT16U",
	"uint32":      "INT32U",
	"uint64":      "INT64U",
	"enum8":       "ENUM8",
	"enum16":      "ENUM16",
	"enum32":      "ENUM32",
	"map8":        "BITMAP8",
	"map16":       "BITMAP16",
	"map32":       "BITMAP32",
	"map64":       "BITMAP64",
	"string":      "CHAR_STRING",
	"octstr":      "OCTET_STRING",
	"elapsed-s":   "elapsed_s",
	"epoch-s":     "epoch_s",
	"epoch-us":    "epoch_us",
	"fabric-idx":  "fabric_idx",
	"node-id":     "NODE_ID",
	"vendor-id":   "VENDOR_ID",
	"group-id":    "group_id",
	"endpoint-id": "endpoint_id",
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
