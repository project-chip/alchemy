package matter

import (
	"encoding/json"
	"math"
	"strings"
)

type BaseDataType uint16

const (
	BaseDataTypeUnknown BaseDataType = iota
	BaseDataTypeCustom
	BaseDataTypeBoolean
	BaseDataTypeMap8
	BaseDataTypeMap16
	BaseDataTypeMap32
	BaseDataTypeMap64
	BaseDataTypeUInt8
	BaseDataTypeUInt16
	BaseDataTypeUInt24
	BaseDataTypeUInt32
	BaseDataTypeUInt40
	BaseDataTypeUInt48
	BaseDataTypeUInt56
	BaseDataTypeUInt64
	BaseDataTypeInt8
	BaseDataTypeInt16
	BaseDataTypeInt24
	BaseDataTypeInt32
	BaseDataTypeInt40
	BaseDataTypeInt48
	BaseDataTypeInt56
	BaseDataTypeInt64
	BaseDataTypeSingle
	BaseDataTypeDouble
	BaseDataTypeOctStr
	BaseDataTypeList
	BaseDataTypeStruct

	BaseDataTypePercent
	BaseDataTypePercentHundredths
	BaseDataTypeTimeOfDay
	BaseDataTypeDate
	BaseDataTypeEpochMicroseconds
	BaseDataTypeEpochSeconds
	BaseDataTypePosixMilliseconds
	BaseDataTypeSystimeMicroseconds
	BaseDataTypeSystimeMilliseconds
	BaseDataTypeElapsedSeconds
	BaseDataTypeTemperature
	BaseDataTypeAmperage
	BaseDataTypeVoltage
	BaseDataTypePower
	BaseDataTypeEnergy

	BaseDataTypeTemperatureDifference
	BaseDataTypeUnsignedTemperature
	BaseDataTypeSignedTemperature

	BaseDataTypeEnum8
	BaseDataTypeEnum16
	BaseDataTypePriority
	BaseDataTypeStatus

	BaseDataTypeGroupID
	BaseDataTypeEndpointID
	BaseDataTypeEndpointNumber
	BaseDataTypeVendorID
	BaseDataTypeDeviceTypeID
	BaseDataTypeFabricID
	BaseDataTypeFabricIndex
	BaseDataTypeClusterID
	BaseDataTypeAttributeID
	BaseDataTypeFieldID
	BaseDataTypeEventID
	BaseDataTypeCommandID
	BaseDataTypeActionID
	BaseDataTypeTransactionID
	BaseDataTypeNodeID
	BaseDataTypeIeeeAddress
	BaseDataTypeEntryIndex
	BaseDataTypeDataVersion
	BaseDataTypeEventNumber

	BaseDataTypeString
	BaseDataTypeIPAddress
	BaseDataTypeIPv4Address
	BaseDataTypeIPv6Address
	BaseDataTypeIPv6Prefix
	BaseDataTypeHardwareAddress

	BaseDataTypeSemanticTag
	BaseDataTypeNamespace
	BaseDataTypeTag
)

type DataTypeCategory uint8

const (
	DataTypeCategoryUnknown DataTypeCategory = iota
	DataTypeCategoryBitmap
	DataTypeCategoryEnum
	DataTypeCategoryStruct
)

var DataTypeOrder = [...]DataTypeCategory{
	DataTypeCategoryBitmap,
	DataTypeCategoryEnum,
	DataTypeCategoryStruct,
}

var DataTypeSuffixes = map[DataTypeCategory]string{
	DataTypeCategoryBitmap: "Bitmap",
	DataTypeCategoryEnum:   "Enum",
	DataTypeCategoryStruct: "Struct",
}

var DataTypeIdentityColumn = map[DataTypeCategory]TableColumn{
	DataTypeCategoryBitmap: TableColumnBit,
	DataTypeCategoryEnum:   TableColumnValue,
	DataTypeCategoryStruct: TableColumnID,
}

func StripDataTypeSuffixes(dataType string) string {
	for _, suffix := range DataTypeSuffixes {
		if strings.HasSuffix(dataType, suffix) {
			dataType = dataType[0 : len(dataType)-len(suffix)]
			break
		}
	}
	return dataType
}

type DataType struct {
	BaseType BaseDataType
	Name     string
	Model    any

	isArray bool
}

func (w *DataType) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"baseType": w.BaseType,
		"name":     w.Name,
		"isArray":  w.isArray,
	})
}

func NewDataType(name string, isArray bool) *DataType {
	dt := &DataType{Name: name, isArray: isArray}
	switch strings.ToLower(name) {
	case "bool", "boolean":
		dt.BaseType = BaseDataTypeBoolean
	case "uint8":
		dt.BaseType = BaseDataTypeUInt8
	case "uint16":
		dt.BaseType = BaseDataTypeUInt16
	case "uint24":
		dt.BaseType = BaseDataTypeUInt24
	case "uint32":
		dt.BaseType = BaseDataTypeUInt32
	case "uint40":
		dt.BaseType = BaseDataTypeUInt40
	case "uint48":
		dt.BaseType = BaseDataTypeUInt48
	case "uint56":
		dt.BaseType = BaseDataTypeUInt56
	case "uint64":
		dt.BaseType = BaseDataTypeUInt64

	case "int8":
		dt.BaseType = BaseDataTypeInt8
	case "int16":
		dt.BaseType = BaseDataTypeInt16
	case "int24":
		dt.BaseType = BaseDataTypeInt24
	case "int32":
		dt.BaseType = BaseDataTypeInt32
	case "int40":
		dt.BaseType = BaseDataTypeInt40
	case "int48":
		dt.BaseType = BaseDataTypeInt48
	case "int56":
		dt.BaseType = BaseDataTypeInt56
	case "int64":
		dt.BaseType = BaseDataTypeInt64

	case "single":
		dt.BaseType = BaseDataTypeSingle
	case "double":
		dt.BaseType = BaseDataTypeDouble

	case "enum8":
		dt.BaseType = BaseDataTypeEnum8
	case "enum16":
		dt.BaseType = BaseDataTypeEnum16

	case "map8":
		dt.BaseType = BaseDataTypeMap8
	case "map16":
		dt.BaseType = BaseDataTypeMap16
	case "map32":
		dt.BaseType = BaseDataTypeMap32
	case "map64":
		dt.BaseType = BaseDataTypeMap64

	case "string":
		dt.BaseType = BaseDataTypeString
	case "octstr":
		dt.BaseType = BaseDataTypeOctStr
	case "percent":
		dt.BaseType = BaseDataTypePercent
	case "percent100ths":
		dt.BaseType = BaseDataTypePercentHundredths
	case "temperature":
		dt.BaseType = BaseDataTypeTemperature
	case "amperage-ma":
		dt.BaseType = BaseDataTypeAmperage
	case "voltage-mv":
		dt.BaseType = BaseDataTypeVoltage
	case "power-mw":
		dt.BaseType = BaseDataTypePower
	case "energy-mwh":
		dt.BaseType = BaseDataTypeEnergy
	case "elapsed-s":
		dt.BaseType = BaseDataTypeElapsedSeconds
	case "epoch-s", "utc": // utc is deprecated
		dt.BaseType = BaseDataTypeEpochSeconds
	case "epoch-us":
		dt.BaseType = BaseDataTypeEpochMicroseconds
	case "systime_ms", "systime-ms":
		dt.BaseType = BaseDataTypeSystimeMilliseconds
	case "systime_us", "systime-us":
		dt.BaseType = BaseDataTypeSystimeMicroseconds
	case "posix-ms":
		dt.BaseType = BaseDataTypePosixMilliseconds
	case "action-id":
		dt.BaseType = BaseDataTypeActionID
	case "attrib-id", "attribute-id":
		dt.BaseType = BaseDataTypeAttributeID
	case "cluster-id":
		dt.BaseType = BaseDataTypeClusterID
	case "command-id":
		dt.BaseType = BaseDataTypeCommandID
	case "data-ver":
		dt.BaseType = BaseDataTypeDataVersion
	case "devtype-id":
		dt.BaseType = BaseDataTypeDeviceTypeID
	case "entry-idx":
		dt.BaseType = BaseDataTypeEntryIndex
	case "event-id":
		dt.BaseType = BaseDataTypeEventID
	case "event-no":
		dt.BaseType = BaseDataTypeEventNumber
	case "fabric-id":
		dt.BaseType = BaseDataTypeFabricID
	case "fabric-idx":
		dt.BaseType = BaseDataTypeFabricIndex
	case "field-id":
		dt.BaseType = BaseDataTypeFieldID
	case "group-id":
		dt.BaseType = BaseDataTypeGroupID
	case "node-id":
		dt.BaseType = BaseDataTypeNodeID
	case "transaction-id":
		dt.BaseType = BaseDataTypeTransactionID
	case "vendor-id":
		dt.BaseType = BaseDataTypeVendorID
	case "endpoint-id":
		dt.BaseType = BaseDataTypeEndpointID
	case "endpoint-no":
		dt.BaseType = BaseDataTypeEndpointNumber
	case "eui64":
		dt.BaseType = BaseDataTypeIeeeAddress
	case "temperaturedifference":
		dt.BaseType = BaseDataTypeTemperatureDifference
	case "unsignedtemperature":
		dt.BaseType = BaseDataTypeUnsignedTemperature
	case "signedtemperature":
		dt.BaseType = BaseDataTypeSignedTemperature
	default:
		dt.BaseType = BaseDataTypeCustom
	}
	return dt
}

func (dt *DataType) IsString() bool {
	return dt != nil && dt.Name == "string" || dt.Name == "octstr"
}

func (dt *DataType) IsArray() bool {
	return dt != nil && dt.isArray
}

func (dt *DataType) Size() int {

	switch dt.BaseType {
	case BaseDataTypeBoolean, BaseDataTypeMap8, BaseDataTypeUInt8, BaseDataTypeInt8, BaseDataTypeEnum8, BaseDataTypePercent, BaseDataTypePriority, BaseDataTypeStatus:
		return 1
	case BaseDataTypeMap16, BaseDataTypeUInt16, BaseDataTypeInt16, BaseDataTypeEnum16, BaseDataTypePercentHundredths, BaseDataTypeGroupID, BaseDataTypeEndpointID, BaseDataTypeEndpointNumber, BaseDataTypeVendorID:
		return 2
	case BaseDataTypeUInt24, BaseDataTypeInt24:
		return 3
	case BaseDataTypeMap32, BaseDataTypeUInt32, BaseDataTypeInt32, BaseDataTypeSingle, BaseDataTypeEpochSeconds, BaseDataTypeElapsedSeconds:
		return 4
	case BaseDataTypeUInt40, BaseDataTypeInt40:
		return 5
	case BaseDataTypeUInt48, BaseDataTypeInt48:
		return 6
	case BaseDataTypeUInt56, BaseDataTypeInt56:
		return 7
	case BaseDataTypeMap64, BaseDataTypeUInt64, BaseDataTypeInt64, BaseDataTypeDouble:
		return 8
	case BaseDataTypeTimeOfDay:
		return 4
	case BaseDataTypeDate:
		return 8
	case BaseDataTypeEpochMicroseconds, BaseDataTypePosixMilliseconds, BaseDataTypeSystimeMicroseconds, BaseDataTypeSystimeMilliseconds:
		return 8
	case BaseDataTypeTemperature, BaseDataTypeTemperatureDifference:
		return 2
	case BaseDataTypeSignedTemperature, BaseDataTypeUnsignedTemperature:
		return 1
	case BaseDataTypeAmperage, BaseDataTypeVoltage, BaseDataTypePower, BaseDataTypeEnergy:
		return 8
	case BaseDataTypeDeviceTypeID:
		return 4
	case BaseDataTypeFabricID, BaseDataTypeNodeID, BaseDataTypeIeeeAddress:
		return 8
	case BaseDataTypeFabricIndex:
		return 1
	case BaseDataTypeClusterID, BaseDataTypeAttributeID, BaseDataTypeFieldID, BaseDataTypeEventID, BaseDataTypeCommandID, BaseDataTypeTransactionID:
		return 4
	case BaseDataTypeActionID:
		return 1
	case BaseDataTypeEntryIndex:
		return 2
	case BaseDataTypeDataVersion:
		return 4
	case BaseDataTypeEventNumber:
		return 8
	case BaseDataTypeIPv4Address:
		return 4
	case BaseDataTypeIPv6Address:
		return 16
	case BaseDataTypeSemanticTag:
		return 4
	case BaseDataTypeNamespace:
		return 1
	case BaseDataTypeTag:
		return 1
	}
	return 0
}

var maxUint24 uint64 = math.MaxUint16 | (math.MaxUint8 << 16)
var maxUint40 uint64 = math.MaxUint32 | (math.MaxUint8 << 32)
var maxUint48 uint64 = math.MaxUint32 | (math.MaxUint16 << 32)
var maxUint56 uint64 = math.MaxUint32 | (maxUint24 << 32)

var maxInt24 int64 = math.MaxUint16 | (math.MaxInt8 << 16)
var maxInt40 int64 = math.MaxUint32 | (math.MaxInt8 << 32)
var maxInt48 int64 = math.MaxUint32 | (math.MaxInt16 << 32)
var maxInt56 int64 = math.MaxUint32 | (maxInt24 << 32)
var maxInt62 int64 = math.MaxUint64 & ^(3 << 62)

var minInt24 int64 = ^maxInt24
var minInt40 int64 = ^maxInt40
var minInt48 int64 = ^maxInt48
var minInt56 int64 = ^maxInt56
var minInt62 int64 = ^maxInt62

var fromRanges = map[BaseDataType]ConstraintExtreme{
	BaseDataTypeInt8:                  {Type: ConstraintExtremeTypeInt64, Int64: math.MinInt8, Format: ConstraintExtremeFormatAuto},
	BaseDataTypeInt16:                 {Type: ConstraintExtremeTypeInt64, Int64: math.MinInt16, Format: ConstraintExtremeFormatAuto},
	BaseDataTypeInt24:                 {Type: ConstraintExtremeTypeInt64, Int64: minInt24, Format: ConstraintExtremeFormatAuto},
	BaseDataTypeInt32:                 {Type: ConstraintExtremeTypeInt64, Int64: math.MinInt32, Format: ConstraintExtremeFormatAuto},
	BaseDataTypeInt40:                 {Type: ConstraintExtremeTypeInt64, Int64: minInt40, Format: ConstraintExtremeFormatAuto},
	BaseDataTypeInt48:                 {Type: ConstraintExtremeTypeInt64, Int64: minInt48, Format: ConstraintExtremeFormatAuto},
	BaseDataTypeInt56:                 {Type: ConstraintExtremeTypeInt64, Int64: minInt56, Format: ConstraintExtremeFormatAuto},
	BaseDataTypeInt64:                 {Type: ConstraintExtremeTypeInt64, Int64: math.MinInt64, Format: ConstraintExtremeFormatAuto},
	BaseDataTypeTemperature:           {Type: ConstraintExtremeTypeInt64, Int64: -27315, Format: ConstraintExtremeFormatInt},
	BaseDataTypeAmperage:              {Type: ConstraintExtremeTypeInt64, Int64: minInt62},
	BaseDataTypeVoltage:               {Type: ConstraintExtremeTypeInt64, Int64: minInt62},
	BaseDataTypePower:                 {Type: ConstraintExtremeTypeInt64, Int64: minInt62},
	BaseDataTypeEnergy:                {Type: ConstraintExtremeTypeInt64, Int64: minInt62},
	BaseDataTypeTemperatureDifference: {Type: ConstraintExtremeTypeInt64, Int64: math.MinInt16},
	BaseDataTypeSignedTemperature:     {Type: ConstraintExtremeTypeInt64, Int64: math.MinInt8},
	BaseDataTypeUnsignedTemperature:   {Type: ConstraintExtremeTypeUInt64, UInt64: 0},
}

var fromRangesNullable = map[BaseDataType]ConstraintExtreme{
	BaseDataTypeInt8:  {Type: ConstraintExtremeTypeInt64, Int64: math.MinInt8 + 1, Format: ConstraintExtremeFormatAuto},
	BaseDataTypeInt16: {Type: ConstraintExtremeTypeInt64, Int64: math.MinInt16 + 1, Format: ConstraintExtremeFormatAuto},
	BaseDataTypeInt24: {Type: ConstraintExtremeTypeInt64, Int64: minInt24 + 1, Format: ConstraintExtremeFormatAuto},
	BaseDataTypeInt32: {Type: ConstraintExtremeTypeInt64, Int64: math.MinInt32 + 1, Format: ConstraintExtremeFormatAuto},
	BaseDataTypeInt40: {Type: ConstraintExtremeTypeInt64, Int64: minInt40 + 1, Format: ConstraintExtremeFormatAuto},
	BaseDataTypeInt48: {Type: ConstraintExtremeTypeInt64, Int64: minInt48 + 1, Format: ConstraintExtremeFormatAuto},
	BaseDataTypeInt56: {Type: ConstraintExtremeTypeInt64, Int64: minInt56 + 1, Format: ConstraintExtremeFormatAuto},
	BaseDataTypeInt64: {Type: ConstraintExtremeTypeInt64, Int64: math.MinInt64 + 1, Format: ConstraintExtremeFormatAuto},

	BaseDataTypeAmperage: {Type: ConstraintExtremeTypeInt64, Int64: minInt62 + 1},
	BaseDataTypeVoltage:  {Type: ConstraintExtremeTypeInt64, Int64: minInt62 + 1},
	BaseDataTypePower:    {Type: ConstraintExtremeTypeInt64, Int64: minInt62 + 1},
	BaseDataTypeEnergy:   {Type: ConstraintExtremeTypeInt64, Int64: minInt62 + 1},

	BaseDataTypeTemperatureDifference: {Type: ConstraintExtremeTypeInt64, Int64: math.MinInt16 + 1},
	BaseDataTypeSignedTemperature:     {Type: ConstraintExtremeTypeInt64, Int64: math.MinInt8 + 1},
	BaseDataTypeUnsignedTemperature:   {Type: ConstraintExtremeTypeUInt64, UInt64: 0},
}

var toRanges = map[BaseDataType]ConstraintExtreme{
	BaseDataTypeMap8:              {Type: ConstraintExtremeTypeUInt64, UInt64: math.MaxUint8},
	BaseDataTypeMap16:             {Type: ConstraintExtremeTypeUInt64, UInt64: math.MaxUint16},
	BaseDataTypeMap32:             {Type: ConstraintExtremeTypeUInt64, UInt64: math.MaxUint32},
	BaseDataTypeMap64:             {Type: ConstraintExtremeTypeUInt64, UInt64: math.MaxUint64},
	BaseDataTypeEnum8:             {Type: ConstraintExtremeTypeUInt64, UInt64: math.MaxUint8},
	BaseDataTypeEnum16:            {Type: ConstraintExtremeTypeUInt64, UInt64: math.MaxUint16},
	BaseDataTypeUInt8:             {Type: ConstraintExtremeTypeUInt64, UInt64: math.MaxUint8, Format: ConstraintExtremeFormatAuto},
	BaseDataTypeUInt16:            {Type: ConstraintExtremeTypeUInt64, UInt64: math.MaxUint16, Format: ConstraintExtremeFormatAuto},
	BaseDataTypeUInt24:            {Type: ConstraintExtremeTypeUInt64, UInt64: maxUint24, Format: ConstraintExtremeFormatAuto},
	BaseDataTypeUInt32:            {Type: ConstraintExtremeTypeUInt64, UInt64: math.MaxUint32, Format: ConstraintExtremeFormatAuto},
	BaseDataTypeUInt40:            {Type: ConstraintExtremeTypeUInt64, UInt64: maxUint40, Format: ConstraintExtremeFormatAuto},
	BaseDataTypeUInt48:            {Type: ConstraintExtremeTypeUInt64, UInt64: maxUint48, Format: ConstraintExtremeFormatAuto},
	BaseDataTypeUInt56:            {Type: ConstraintExtremeTypeUInt64, UInt64: maxUint56, Format: ConstraintExtremeFormatAuto},
	BaseDataTypeUInt64:            {Type: ConstraintExtremeTypeUInt64, UInt64: math.MaxUint64, Format: ConstraintExtremeFormatAuto},
	BaseDataTypeInt8:              {Type: ConstraintExtremeTypeInt64, Int64: math.MaxInt8, Format: ConstraintExtremeFormatAuto},
	BaseDataTypeInt16:             {Type: ConstraintExtremeTypeInt64, Int64: math.MaxInt16, Format: ConstraintExtremeFormatAuto},
	BaseDataTypeInt24:             {Type: ConstraintExtremeTypeInt64, Int64: maxInt24, Format: ConstraintExtremeFormatAuto},
	BaseDataTypeInt32:             {Type: ConstraintExtremeTypeInt64, Int64: math.MaxInt32, Format: ConstraintExtremeFormatAuto},
	BaseDataTypeInt40:             {Type: ConstraintExtremeTypeInt64, Int64: maxInt40, Format: ConstraintExtremeFormatAuto},
	BaseDataTypeInt48:             {Type: ConstraintExtremeTypeInt64, Int64: maxInt48, Format: ConstraintExtremeFormatAuto},
	BaseDataTypeInt56:             {Type: ConstraintExtremeTypeInt64, Int64: maxInt56, Format: ConstraintExtremeFormatAuto},
	BaseDataTypeInt64:             {Type: ConstraintExtremeTypeInt64, Int64: math.MaxInt64, Format: ConstraintExtremeFormatAuto},
	BaseDataTypePercent:           {Type: ConstraintExtremeTypeUInt64, UInt64: 100},
	BaseDataTypePercentHundredths: {Type: ConstraintExtremeTypeUInt64, UInt64: 10000},
	BaseDataTypeEpochMicroseconds: {Type: ConstraintExtremeTypeUInt64, UInt64: math.MaxUint64, Format: ConstraintExtremeFormatHex},
	BaseDataTypeEpochSeconds:      {Type: ConstraintExtremeTypeUInt64, UInt64: math.MaxUint32, Format: ConstraintExtremeFormatHex},
	BaseDataTypePosixMilliseconds: {Type: ConstraintExtremeTypeUInt64, UInt64: math.MaxUint64, Format: ConstraintExtremeFormatHex},

	BaseDataTypeAmperage: {Type: ConstraintExtremeTypeInt64, Int64: maxInt62},
	BaseDataTypeVoltage:  {Type: ConstraintExtremeTypeInt64, Int64: maxInt62},
	BaseDataTypePower:    {Type: ConstraintExtremeTypeInt64, Int64: maxInt62},
	BaseDataTypeEnergy:   {Type: ConstraintExtremeTypeInt64, Int64: maxInt62},

	BaseDataTypeTemperature:           {Type: ConstraintExtremeTypeInt64, Int64: math.MaxInt16, Format: ConstraintExtremeFormatHex},
	BaseDataTypeTemperatureDifference: {Type: ConstraintExtremeTypeInt64, Int64: math.MaxInt16},
	BaseDataTypeSignedTemperature:     {Type: ConstraintExtremeTypeInt64, Int64: math.MinInt8},
	BaseDataTypeUnsignedTemperature:   {Type: ConstraintExtremeTypeUInt64, UInt64: math.MaxUint8},
}

var toRangesNullable = map[BaseDataType]ConstraintExtreme{
	BaseDataTypeMap8:              {Type: ConstraintExtremeTypeUInt64, UInt64: math.MaxUint8 & ^(1 << 7)},
	BaseDataTypeMap16:             {Type: ConstraintExtremeTypeUInt64, UInt64: math.MaxUint16 & ^(1 << 15)},
	BaseDataTypeMap32:             {Type: ConstraintExtremeTypeUInt64, UInt64: math.MaxUint32 & ^(1 << 31)},
	BaseDataTypeMap64:             {Type: ConstraintExtremeTypeUInt64, UInt64: math.MaxUint64 & ^(1 << 64)},
	BaseDataTypeEnum8:             {Type: ConstraintExtremeTypeUInt64, UInt64: math.MaxUint8 - 1},
	BaseDataTypeEnum16:            {Type: ConstraintExtremeTypeUInt64, UInt64: math.MaxUint16 - 1},
	BaseDataTypeUInt8:             {Type: ConstraintExtremeTypeUInt64, UInt64: math.MaxUint8 - 1, Format: ConstraintExtremeFormatAuto},
	BaseDataTypeUInt16:            {Type: ConstraintExtremeTypeUInt64, UInt64: math.MaxUint16 - 1, Format: ConstraintExtremeFormatAuto},
	BaseDataTypeUInt24:            {Type: ConstraintExtremeTypeUInt64, UInt64: maxUint24 - 1, Format: ConstraintExtremeFormatAuto},
	BaseDataTypeUInt32:            {Type: ConstraintExtremeTypeUInt64, UInt64: math.MaxUint32 - 1, Format: ConstraintExtremeFormatAuto},
	BaseDataTypeUInt40:            {Type: ConstraintExtremeTypeUInt64, UInt64: maxUint40 - 1, Format: ConstraintExtremeFormatAuto},
	BaseDataTypeUInt48:            {Type: ConstraintExtremeTypeUInt64, UInt64: maxUint48 - 1, Format: ConstraintExtremeFormatAuto},
	BaseDataTypeUInt56:            {Type: ConstraintExtremeTypeUInt64, UInt64: maxUint56 - 1, Format: ConstraintExtremeFormatAuto},
	BaseDataTypeUInt64:            {Type: ConstraintExtremeTypeUInt64, UInt64: math.MaxUint64 - 1, Format: ConstraintExtremeFormatAuto},
	BaseDataTypeEpochMicroseconds: {Type: ConstraintExtremeTypeUInt64, UInt64: math.MaxUint64 - 1, Format: ConstraintExtremeFormatHex},
	BaseDataTypeEpochSeconds:      {Type: ConstraintExtremeTypeUInt64, UInt64: math.MaxUint32 - 1, Format: ConstraintExtremeFormatHex},
	BaseDataTypePosixMilliseconds: {Type: ConstraintExtremeTypeUInt64, UInt64: math.MaxUint64 - 1, Format: ConstraintExtremeFormatHex},
	BaseDataTypeAmperage:          {Type: ConstraintExtremeTypeInt64, Int64: maxInt62},
	BaseDataTypeVoltage:           {Type: ConstraintExtremeTypeInt64, Int64: maxInt62},
	BaseDataTypePower:             {Type: ConstraintExtremeTypeInt64, Int64: maxInt62},
	BaseDataTypeEnergy:            {Type: ConstraintExtremeTypeInt64, Int64: maxInt62},

	BaseDataTypeUnsignedTemperature: {Type: ConstraintExtremeTypeUInt64, UInt64: math.MaxUint8 - 1},
}

func (dt *DataType) Min(nullable bool) (from ConstraintExtreme) {
	var ok bool
	if nullable {
		from, ok = fromRangesNullable[dt.BaseType]
	}
	if !ok {
		from = fromRanges[dt.BaseType]
	}
	return
}

func (dt *DataType) Max(nullable bool) (to ConstraintExtreme) {
	var ok bool
	if nullable {
		to, ok = toRangesNullable[dt.BaseType]
	}
	if !ok {
		to = toRanges[dt.BaseType]
	}
	return
}
