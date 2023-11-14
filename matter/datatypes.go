package matter

import (
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
	BaseType BaseDataType `json:"baseType,omitempty"`
	Name     string       `json:"name,omitempty"`
	IsArray  bool         `json:"isArray,omitempty"`
}

func NewDataType(name string, isArray bool) *DataType {
	dt := &DataType{Name: name, IsArray: isArray}
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
	case "systime_ms":
		dt.BaseType = BaseDataTypeSystimeMilliseconds
	case "systime_us":
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
	}
	return dt
}

func (dt *DataType) IsString() bool {
	return dt != nil && dt.Name == "string" || dt.Name == "octstr"
}

type dataTypeRange struct {
	et    ConstraintExtremeType
	fromU uint64
	fromI int64
	toU   uint64
	toI   int64
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

var ranges = map[BaseDataType]dataTypeRange{
	BaseDataTypeBoolean:           {et: ConstraintExtremeTypeUInt64, fromU: 0, toU: 1},
	BaseDataTypeMap8:              {et: ConstraintExtremeTypeUInt64, fromU: 0, toU: math.MaxUint8},
	BaseDataTypeMap16:             {et: ConstraintExtremeTypeUInt64, fromU: 0, toU: math.MaxUint16},
	BaseDataTypeMap32:             {et: ConstraintExtremeTypeUInt64, fromU: 0, toU: math.MaxUint32},
	BaseDataTypeMap64:             {et: ConstraintExtremeTypeUInt64, fromU: 0, toU: math.MaxUint64},
	BaseDataTypeEnum8:             {et: ConstraintExtremeTypeUInt64, fromU: 0, toU: math.MaxUint8},
	BaseDataTypeEnum16:            {et: ConstraintExtremeTypeUInt64, fromU: 0, toU: math.MaxUint16},
	BaseDataTypeUInt8:             {et: ConstraintExtremeTypeUInt64, fromU: 0, toU: math.MaxUint8},
	BaseDataTypeUInt16:            {et: ConstraintExtremeTypeUInt64, fromU: 0, toU: math.MaxUint16},
	BaseDataTypeUInt24:            {et: ConstraintExtremeTypeUInt64, fromU: 0, toU: maxUint24},
	BaseDataTypeUInt32:            {et: ConstraintExtremeTypeUInt64, fromU: 0, toU: math.MaxUint32},
	BaseDataTypeUInt40:            {et: ConstraintExtremeTypeUInt64, fromU: 0, toU: maxUint40},
	BaseDataTypeUInt48:            {et: ConstraintExtremeTypeUInt64, fromU: 0, toU: maxUint48},
	BaseDataTypeUInt56:            {et: ConstraintExtremeTypeUInt64, fromU: 0, toU: maxUint56},
	BaseDataTypeUInt64:            {et: ConstraintExtremeTypeUInt64, fromU: 0, toU: math.MaxUint64},
	BaseDataTypeInt8:              {et: ConstraintExtremeTypeInt64, fromI: math.MinInt8, toI: math.MaxInt8},
	BaseDataTypeInt16:             {et: ConstraintExtremeTypeInt64, fromI: math.MinInt16, toI: math.MaxInt16},
	BaseDataTypeInt24:             {et: ConstraintExtremeTypeInt64, fromI: minInt24, toI: maxInt24},
	BaseDataTypeInt32:             {et: ConstraintExtremeTypeInt64, fromI: math.MinInt32, toI: math.MaxInt32},
	BaseDataTypeInt40:             {et: ConstraintExtremeTypeInt64, fromI: minInt40, toI: maxInt40},
	BaseDataTypeInt48:             {et: ConstraintExtremeTypeInt64, fromI: minInt48, toI: maxInt48},
	BaseDataTypeInt56:             {et: ConstraintExtremeTypeInt64, fromI: minInt56, toI: maxInt56},
	BaseDataTypeInt64:             {et: ConstraintExtremeTypeInt64, fromI: math.MinInt64, toI: math.MaxInt64},
	BaseDataTypePercent:           {et: ConstraintExtremeTypeUInt64, fromU: 0, toU: 100},
	BaseDataTypePercentHundredths: {et: ConstraintExtremeTypeUInt64, fromU: 0, toU: 10000},
	BaseDataTypeEpochMicroseconds: {et: ConstraintExtremeTypeUInt64, fromU: 0, toU: math.MaxUint64},
	BaseDataTypeEpochSeconds:      {et: ConstraintExtremeTypeUInt64, fromU: 0, toU: math.MaxUint32},
	BaseDataTypePosixMilliseconds: {et: ConstraintExtremeTypeUInt64, fromU: 0, toU: math.MaxUint64},
	BaseDataTypeTemperature:       {et: ConstraintExtremeTypeInt64, fromI: -27315, toI: math.MaxInt16},
	BaseDataTypeAmperage:          {et: ConstraintExtremeTypeInt64, fromI: minInt62, toI: maxInt62},
	BaseDataTypeVoltage:           {et: ConstraintExtremeTypeInt64, fromI: minInt62, toI: maxInt62},
	BaseDataTypePower:             {et: ConstraintExtremeTypeInt64, fromI: minInt62, toI: maxInt62},
	BaseDataTypeEnergy:            {et: ConstraintExtremeTypeInt64, fromI: minInt62, toI: maxInt62},
}

var nullableRanges = map[BaseDataType]dataTypeRange{
	BaseDataTypeMap8:              {et: ConstraintExtremeTypeUInt64, fromU: 0, toU: math.MaxUint8 & ^(1 << 7)},
	BaseDataTypeMap16:             {et: ConstraintExtremeTypeUInt64, fromU: 0, toU: math.MaxUint16 & ^(1 << 15)},
	BaseDataTypeMap32:             {et: ConstraintExtremeTypeUInt64, fromU: 0, toU: math.MaxUint32 & ^(1 << 31)},
	BaseDataTypeMap64:             {et: ConstraintExtremeTypeUInt64, fromU: 0, toU: math.MaxUint64 & ^(1 << 64)},
	BaseDataTypeEnum8:             {et: ConstraintExtremeTypeUInt64, fromU: 0, toU: math.MaxUint8 - 1},
	BaseDataTypeEnum16:            {et: ConstraintExtremeTypeUInt64, fromU: 0, toU: math.MaxUint16 - 1},
	BaseDataTypeUInt8:             {et: ConstraintExtremeTypeUInt64, fromU: 0, toU: math.MaxUint8 - 1},
	BaseDataTypeUInt16:            {et: ConstraintExtremeTypeUInt64, fromU: 0, toU: math.MaxUint16 - 1},
	BaseDataTypeUInt24:            {et: ConstraintExtremeTypeUInt64, fromU: 0, toU: maxUint24 - 1},
	BaseDataTypeUInt32:            {et: ConstraintExtremeTypeUInt64, fromU: 0, toU: math.MaxUint32 - 1},
	BaseDataTypeUInt40:            {et: ConstraintExtremeTypeUInt64, fromU: 0, toU: maxUint40 - 1},
	BaseDataTypeUInt48:            {et: ConstraintExtremeTypeUInt64, fromU: 0, toU: maxUint48 - 1},
	BaseDataTypeUInt56:            {et: ConstraintExtremeTypeUInt64, fromU: 0, toU: maxUint56 - 1},
	BaseDataTypeUInt64:            {et: ConstraintExtremeTypeUInt64, fromU: 0, toU: math.MaxUint64 - 1},
	BaseDataTypeInt8:              {et: ConstraintExtremeTypeInt64, fromI: math.MinInt8 + 1, toI: math.MaxInt8},
	BaseDataTypeInt16:             {et: ConstraintExtremeTypeInt64, fromI: math.MinInt16 + 1, toI: math.MaxInt16},
	BaseDataTypeInt24:             {et: ConstraintExtremeTypeInt64, fromI: minInt24 + 1, toI: maxInt24},
	BaseDataTypeInt32:             {et: ConstraintExtremeTypeInt64, fromI: math.MinInt32 + 1, toI: math.MaxInt32},
	BaseDataTypeInt40:             {et: ConstraintExtremeTypeInt64, fromI: minInt40 + 1, toI: maxInt40},
	BaseDataTypeInt48:             {et: ConstraintExtremeTypeInt64, fromI: minInt48 + 1, toI: maxInt48},
	BaseDataTypeInt56:             {et: ConstraintExtremeTypeInt64, fromI: minInt56 + 1, toI: maxInt56},
	BaseDataTypeInt64:             {et: ConstraintExtremeTypeInt64, fromI: math.MinInt64 + 1, toI: math.MaxInt64},
	BaseDataTypeEpochMicroseconds: {et: ConstraintExtremeTypeUInt64, fromU: 0, toU: math.MaxUint64 - 1},
	BaseDataTypeEpochSeconds:      {et: ConstraintExtremeTypeUInt64, fromU: 0, toU: math.MaxUint32 - 1},
	BaseDataTypePosixMilliseconds: {et: ConstraintExtremeTypeUInt64, fromU: 0, toU: math.MaxUint64 - 1},
	BaseDataTypeAmperage:          {et: ConstraintExtremeTypeInt64, fromI: minInt62 + 1, toI: maxInt62},
	BaseDataTypeVoltage:           {et: ConstraintExtremeTypeInt64, fromI: minInt62 + 1, toI: maxInt62},
	BaseDataTypePower:             {et: ConstraintExtremeTypeInt64, fromI: minInt62 + 1, toI: maxInt62},
	BaseDataTypeEnergy:            {et: ConstraintExtremeTypeInt64, fromI: minInt62 + 1, toI: maxInt62},
}

func (dt *DataType) MinMax(nullable bool) (from ConstraintExtreme, to ConstraintExtreme) {
	var r dataTypeRange
	var ok bool
	if nullable {
		r, ok = nullableRanges[dt.BaseType]
	}
	if !ok {
		r, ok = ranges[dt.BaseType]
	}
	if ok {
		from.Type = r.et
		to.Type = r.et
		switch r.et {
		case ConstraintExtremeTypeInt64:
			from.Int64 = r.fromI
			to.Int64 = r.toI
		case ConstraintExtremeTypeUInt64:
			from.UInt64 = r.fromU
			to.UInt64 = r.toU
		}
	}

	return
}
