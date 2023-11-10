package matter

import (
	"strings"
)

type BaseDataType uint8

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

	BaseDataTypeEnum8
	BaseDataTypeEnum16
	BaseDataTypePriority
	BaseDataTypeStatus

	BaseDataTypeGroupID
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

func (dt *DataType) IsString() bool {
	return dt.Name == "string" || dt.Name == "octstr"
}

func (dt *DataType) MinMax(nullable bool) (from ConstraintExtreme, to ConstraintExtreme) {
	switch dt.BaseType {
	case BaseDataTypeBoolean:
		from = ConstraintExtreme{Type: ConstraintExtremeTypeUInt64, UInt64: 0}
		to = ConstraintExtreme{Type: ConstraintExtremeTypeUInt64, UInt64: 1}
	case BaseDataTypeMap8:
		from = ConstraintExtreme{Type: ConstraintExtremeTypeUInt64, UInt64: 0}
		to = ConstraintExtreme{Type: ConstraintExtremeTypeUInt64, UInt64: 0xFF}
	case BaseDataTypeMap16:
		from = ConstraintExtreme{Type: ConstraintExtremeTypeUInt64, UInt64: 0}
		to = ConstraintExtreme{Type: ConstraintExtremeTypeUInt64, UInt64: 0xFFFF}
	case BaseDataTypeMap32:
		from = ConstraintExtreme{Type: ConstraintExtremeTypeUInt64, UInt64: 0}
		to = ConstraintExtreme{Type: ConstraintExtremeTypeUInt64, UInt64: 0xFFFFFFFF}
	case BaseDataTypeMap64:
		from = ConstraintExtreme{Type: ConstraintExtremeTypeUInt64, UInt64: 0}
		to = ConstraintExtreme{Type: ConstraintExtremeTypeUInt64, UInt64: 0xFFFFFFFFFFFFFFFF}
	case BaseDataTypeUInt8:
	case BaseDataTypeUInt16:
	case BaseDataTypeUInt24:
	case BaseDataTypeUInt32:
	case BaseDataTypeUInt40:
	case BaseDataTypeUInt48:
	case BaseDataTypeUInt56:
	case BaseDataTypeUInt64:
	case BaseDataTypeInt8:
	case BaseDataTypeInt16:
	case BaseDataTypeInt24:
	case BaseDataTypeInt32:
	case BaseDataTypeInt40:
	case BaseDataTypeInt48:
	case BaseDataTypeInt56:
	case BaseDataTypeInt64:
	case BaseDataTypeSingle:
	case BaseDataTypeDouble:
	case BaseDataTypeList:
	case BaseDataTypeStruct:
	}
	return
}
