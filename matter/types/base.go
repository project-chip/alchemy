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
	BaseDataTypeApparentPower
	BaseDataTypeApparentEnergy
	BaseDataTypeReactivePower
	BaseDataTypeReactiveEnergy

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
	BaseDataTypeNamespaceID
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

func (bdt BaseDataType) String() string {
	return BaseDataTypeName(bdt)
}

func (bdt BaseDataType) IsSimple() bool {
	switch bdt {
	case BaseDataTypeUInt8, BaseDataTypeUInt16, BaseDataTypeUInt24, BaseDataTypeUInt32,
		BaseDataTypeUInt40, BaseDataTypeUInt48, BaseDataTypeUInt56, BaseDataTypeUInt64,
		BaseDataTypeInt8, BaseDataTypeInt16, BaseDataTypeInt24, BaseDataTypeInt32,
		BaseDataTypeInt40, BaseDataTypeInt48, BaseDataTypeInt56, BaseDataTypeInt64,
		BaseDataTypeString:
		return true
	}
	return false
}

func BaseDataTypeName(baseDataType BaseDataType) string {
	switch baseDataType {
	case BaseDataTypeUnknown:
		return "unknown"
	case BaseDataTypeCustom:
		return "custom"
	case BaseDataTypeBoolean:
		return "bool"
	case BaseDataTypeMap8:
		return "map8"
	case BaseDataTypeMap16:
		return "map16"
	case BaseDataTypeMap32:
		return "map32"
	case BaseDataTypeMap64:
		return "map64"
	case BaseDataTypeUInt8:
		return "uint8"
	case BaseDataTypeUInt16:
		return "uint16"
	case BaseDataTypeUInt24:
		return "uint24"
	case BaseDataTypeUInt32:
		return "uint32"
	case BaseDataTypeUInt40:
		return "uint40"
	case BaseDataTypeUInt48:
		return "uint48"
	case BaseDataTypeUInt56:
		return "uint56"
	case BaseDataTypeUInt64:
		return "uint64"
	case BaseDataTypeInt8:
		return "int8"
	case BaseDataTypeInt16:
		return "int16"
	case BaseDataTypeInt24:
		return "int24"
	case BaseDataTypeInt32:
		return "int32"
	case BaseDataTypeInt40:
		return "int40"
	case BaseDataTypeInt48:
		return "int48"
	case BaseDataTypeInt56:
		return "int56"
	case BaseDataTypeInt64:
		return "int64"
	case BaseDataTypeSingle:
		return "single"
	case BaseDataTypeDouble:
		return "double"
	case BaseDataTypeOctStr:
		return "octstr"
	case BaseDataTypeList:
		return "list"
	case BaseDataTypePercent:
		return "percent"
	case BaseDataTypePercentHundredths:
		return "percent100ths"
	case BaseDataTypeTimeOfDay:
		return "tod"
	case BaseDataTypeDate:
		return "date"
	case BaseDataTypeEpochMicroseconds:
		return "epoch-us"
	case BaseDataTypeEpochSeconds:
		return "epoch-s"
	case BaseDataTypePosixMilliseconds:
		return "posix-ms"
	case BaseDataTypeSystimeMicroseconds:
		return "systime-us"
	case BaseDataTypeSystimeMilliseconds:
		return "systime-ms"
	case BaseDataTypeElapsedSeconds:
		return "elapsed-s"
	case BaseDataTypeTemperature:
		return "temperature"
	case BaseDataTypeAmperage:
		return "amperage-mA"
	case BaseDataTypeVoltage:
		return "voltage-mW"
	case BaseDataTypePower:
		return "power-mW"
	case BaseDataTypeEnergy:
		return "energy-mWh"
	case BaseDataTypeApparentPower:
		return "power-mVA"
	case BaseDataTypeApparentEnergy:
		return "energy-mVAh"
	case BaseDataTypeReactivePower:
		return "power-mVAR"
	case BaseDataTypeReactiveEnergy:
		return "energy-mVARh"
	case BaseDataTypeTemperatureDifference:
		return "tempdiff"
	case BaseDataTypeUnsignedTemperature:
		return "utemp"
	case BaseDataTypeSignedTemperature:
		return "stemp"
	case BaseDataTypeEnum8:
		return "enum8"
	case BaseDataTypeEnum16:
		return "enum16"
	case BaseDataTypePriority:
		return "priority"
	case BaseDataTypeStatus:
		return "status"
	case BaseDataTypeGroupID:
		return "group-id"
	case BaseDataTypeEndpointID:
		return "endpoint-id"
	case BaseDataTypeEndpointNumber:
		return "endpoint-no"
	case BaseDataTypeVendorID:
		return "vendor-id"
	case BaseDataTypeDeviceTypeID:
		return "devtype-id"
	case BaseDataTypeFabricID:
		return "fabric-id"
	case BaseDataTypeFabricIndex:
		return "fabric-idx"
	case BaseDataTypeClusterID:
		return "cluster-id"
	case BaseDataTypeAttributeID:
		return "attrib-id"
	case BaseDataTypeFieldID:
		return "field-id"
	case BaseDataTypeEventID:
		return "event-id"
	case BaseDataTypeCommandID:
		return "command-id"
	case BaseDataTypeActionID:
		return "action-id"
	case BaseDataTypeSubjectID:
		return "subject-id"
	case BaseDataTypeTransactionID:
		return "trans-id"
	case BaseDataTypeNodeID:
		return "node-id"
	case BaseDataTypeIeeeAddress:
		return "EUI64"
	case BaseDataTypeEntryIndex:
		return "entry-idx"
	case BaseDataTypeDataVersion:
		return "data-ver"
	case BaseDataTypeEventNumber:
		return "event-no"
	case BaseDataTypeString:
		return "string"
	case BaseDataTypeIPAddress:
		return "ipadr"
	case BaseDataTypeIPv4Address:
		return "ipv4adr"
	case BaseDataTypeIPv6Address:
		return "ipv6adr"
	case BaseDataTypeIPv6Prefix:
		return "ipv6pre"
	case BaseDataTypeHardwareAddress:
		return "hwadr"
	case BaseDataTypeSemanticTag:
		return "semtag"
	case BaseDataTypeNamespaceID:
		return "namespace"
	case BaseDataTypeTag:
		return "tag"
	case BaseDataTypeMessageID:
		return "message-id"
	}
	return "unknown"
}
