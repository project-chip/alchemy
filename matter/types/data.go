package types

import "strings"

type DataType struct {
	BaseType BaseDataType `json:"baseType"`
	Name     string       `json:"name"`
	Entity   Entity       `json:"-"`

	EntryType *DataType `json:"entryType,omitempty"`
}

func NewDataType(typeName string, isArray bool) *DataType {
	if len(typeName) == 0 {
		return nil
	}
	if isArray {
		return &DataType{Name: "list", BaseType: BaseDataTypeList, EntryType: NewDataType(typeName, false)}
	}
	typeName = strings.TrimPrefix(typeName, "ref_")
	typeName = strings.TrimPrefix(typeName, "DataType")
	dt := &DataType{Name: typeName}

	switch strings.ToLower(typeName) {
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
	case "epoch-us", "epochus":
		dt.BaseType = BaseDataTypeEpochMicroseconds
	case "systime_ms", "systime-ms":
		dt.BaseType = BaseDataTypeSystimeMilliseconds
	case "systime_us", "systime-us", "systemtimeus":
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
	case "entry-idx", "entryidx":
		dt.BaseType = BaseDataTypeEntryIndex
	case "event-id", "eventid":
		dt.BaseType = BaseDataTypeEventID
	case "event-no", "eventnumber":
		dt.BaseType = BaseDataTypeEventNumber
	case "fabric-id", "fabricid":
		dt.BaseType = BaseDataTypeFabricID
	case "fabric-idx", "fabricidx":
		dt.BaseType = BaseDataTypeFabricIndex
	case "field-id", "fieldid":
		dt.BaseType = BaseDataTypeFieldID
	case "group-id", "groupid":
		dt.BaseType = BaseDataTypeGroupID
	case "node-id", "nodeid":
		dt.BaseType = BaseDataTypeNodeID
	case "subject-id", "subjectid":
		dt.BaseType = BaseDataTypeSubjectID
	case "transaction-id":
		dt.BaseType = BaseDataTypeTransactionID
	case "vendor-id", "vendorid":
		dt.BaseType = BaseDataTypeVendorID
	case "endpoint-id":
		dt.BaseType = BaseDataTypeEndpointID
	case "endpoint-no", "endpointnumber":
		dt.BaseType = BaseDataTypeEndpointNumber
	case "eui64":
		dt.BaseType = BaseDataTypeIeeeAddress
	case "temperaturedifference", "tempdiff":
		dt.BaseType = BaseDataTypeTemperatureDifference
	case "unsignedtemperature":
		dt.BaseType = BaseDataTypeUnsignedTemperature
	case "signedtemperature":
		dt.BaseType = BaseDataTypeSignedTemperature
	case "hwadr":
		dt.BaseType = BaseDataTypeHardwareAddress
	case "ipv4adr":
		dt.BaseType = BaseDataTypeIPv4Address
	case "ipv6adr":
		dt.BaseType = BaseDataTypeIPv6Address
	case "semtag":
		dt.BaseType = BaseDataTypeSemanticTag
	case "status":
		dt.BaseType = BaseDataTypeStatus
	case "priority":
		dt.BaseType = BaseDataTypePriority
	case "messageid":
		dt.BaseType = BaseDataTypeMessageID
	default:
		dt.BaseType = BaseDataTypeCustom
	}
	return dt
}

func (dt *DataType) Clone() *DataType {
	ndt := &DataType{Name: dt.Name, BaseType: dt.BaseType, Entity: dt.Entity}
	if dt.EntryType != nil {
		ndt.EntryType = dt.EntryType.Clone()
	}
	ndt.Entity = dt.Entity
	return ndt
}

func (dt *DataType) HasLength() bool {
	return dt != nil && (dt.BaseType == BaseDataTypeString || dt.BaseType == BaseDataTypeOctStr || dt.BaseType == BaseDataTypeMessageID)
}

func (dt *DataType) IsArray() bool {
	return dt != nil && dt.BaseType == BaseDataTypeList
}

func (dt *DataType) IsMap() bool {
	if dt == nil {
		return false
	}
	switch dt.BaseType {
	case BaseDataTypeMap8, BaseDataTypeMap16, BaseDataTypeMap32, BaseDataTypeMap64:
		return true
	}
	return false
}

func (dt *DataType) IsEnum() bool {
	if dt == nil {
		return false
	}
	switch dt.BaseType {
	case BaseDataTypeEnum8, BaseDataTypeEnum16:
		return true
	}
	return false
}
