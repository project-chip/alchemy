package types

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
	BaseDataTypeSubjectID
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

	BaseDataTypeMessageID
)

type HasBaseDataType interface {
	BaseDataType() BaseDataType
}

func (bdt BaseDataType) IsUnsigned() bool {
	switch bdt {
	case BaseDataTypeBoolean, BaseDataTypeMap8,
		BaseDataTypeMap16, BaseDataTypeMap32, BaseDataTypeMap64,
		BaseDataTypeUInt8, BaseDataTypeUInt16, BaseDataTypeUInt24,
		BaseDataTypeUInt32, BaseDataTypeUInt40, BaseDataTypeUInt48,
		BaseDataTypeUInt56, BaseDataTypeUInt64, BaseDataTypePercent,
		BaseDataTypePercentHundredths, BaseDataTypeEpochMicroseconds,
		BaseDataTypeEpochSeconds, BaseDataTypePosixMilliseconds,
		BaseDataTypeSystimeMicroseconds, BaseDataTypeSystimeMilliseconds,
		BaseDataTypeElapsedSeconds, BaseDataTypeGroupID, BaseDataTypeEndpointNumber,
		BaseDataTypeVendorID, BaseDataTypeDeviceTypeID, BaseDataTypeFabricID,
		BaseDataTypeFabricIndex, BaseDataTypeClusterID, BaseDataTypeAttributeID,
		BaseDataTypeFieldID, BaseDataTypeEventID, BaseDataTypeCommandID, BaseDataTypeActionID,
		BaseDataTypeTransactionID, BaseDataTypeNodeID, BaseDataTypeEntryIndex,
		BaseDataTypeDataVersion, BaseDataTypeEventNumber:
		return true
	}
	return false
}
