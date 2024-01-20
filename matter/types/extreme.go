package types

import (
	"fmt"
	"strconv"
)

type DataTypeExtremeType uint8

const (
	DataTypeExtremeTypeUndefined DataTypeExtremeType = iota
	DataTypeExtremeTypeInt64
	DataTypeExtremeTypeUInt64
	DataTypeExtremeTypeNull
	DataTypeExtremeTypeEmpty
)

type DataTypeExtreme struct {
	Type   DataTypeExtremeType
	Format NumberFormat
	Int64  int64
	UInt64 uint64
}

func NewIntDataTypeExtreme(i int64, f NumberFormat) DataTypeExtreme {
	return DataTypeExtreme{Type: DataTypeExtremeTypeInt64, Format: f, Int64: i}
}

func NewUintDataTypeExtreme(u uint64, f NumberFormat) DataTypeExtreme {
	return DataTypeExtreme{Type: DataTypeExtremeTypeUInt64, Format: f, UInt64: u}
}

func (ce *DataTypeExtreme) Defined() bool {
	return ce.Type != DataTypeExtremeTypeUndefined
}

func (ce *DataTypeExtreme) IsNull() bool {
	return ce.Type == DataTypeExtremeTypeNull
}

func (ce *DataTypeExtreme) Value() any {
	switch ce.Type {
	case DataTypeExtremeTypeInt64:
		return ce.Int64
	case DataTypeExtremeTypeUInt64:
		return ce.UInt64
	default:
		return nil
	}
}

func (ce *DataTypeExtreme) ZapString(dataType *DataType) string {
	switch ce.Type {
	case DataTypeExtremeTypeInt64:
		val := ce.Int64
		switch ce.Format {
		case NumberFormatHex:
			if dataType != nil {
				switch dataType.Size() {
				case 1:
					return fmt.Sprintf("0x%02X", uint8(val))
				case 2:
					return fmt.Sprintf("0x%04X", uint16(val))
				case 4:
					return fmt.Sprintf("0x%08X", uint32(val))
				case 8:
					return fmt.Sprintf("0x%16X", uint64(val))

				}
			}
			return fmt.Sprintf("0x%X", uint64(val))
		case NumberFormatAuto:
			if val > 255 || val < 256 {
				return fmt.Sprintf("0x%X", uint64(val))
			}
			return strconv.FormatInt(val, 10)
		default:
			if dataType != nil && dataType.BaseType == BaseDataTypePercentHundredths {
				return strconv.FormatInt(ce.Int64*100, 10)
			}
			return strconv.FormatInt(ce.Int64, 10)
		}
	case DataTypeExtremeTypeUInt64:
		return ce.formatUint64(dataType, ce.UInt64)
	case DataTypeExtremeTypeNull, DataTypeExtremeTypeEmpty:
		return ""
	}
	return ""
}

func (ce *DataTypeExtreme) DataModelString(dataType *DataType) string {
	switch ce.Type {
	case DataTypeExtremeTypeInt64:
		val := ce.Int64
		switch ce.Format {
		case NumberFormatHex:
			if dataType != nil {
				switch dataType.Size() {
				case 1:
					return fmt.Sprintf("0x%02X", uint8(val))
				case 2:
					return fmt.Sprintf("0x%04X", uint16(val))
				case 4:
					return fmt.Sprintf("0x%08X", uint32(val))
				case 8:
					return fmt.Sprintf("0x%16X", uint64(val))

				}
			}
			return fmt.Sprintf("0x%X", uint64(val))
		case NumberFormatAuto:
			if val > 255 || val < 256 {
				return fmt.Sprintf("0x%X", uint64(val))
			}
			return strconv.FormatInt(val, 10)
		default:
			if dataType != nil {
				switch dataType.BaseType {
				case BaseDataTypePercentHundredths:
					return strconv.FormatInt(ce.Int64*100, 10)
				case BaseDataTypeBoolean:
					return strconv.FormatBool(ce.Int64 == 1)
				}
			}
			return strconv.FormatInt(ce.Int64, 10)
		}
	case DataTypeExtremeTypeUInt64:
		if dataType != nil {
			switch dataType.BaseType {
			case BaseDataTypeBoolean:
				return strconv.FormatBool(ce.Int64 == 1)
			}
		}
		return ce.formatUint64(dataType, ce.UInt64)
	case DataTypeExtremeTypeNull:
		return "null"
	case DataTypeExtremeTypeEmpty:
		return "empty"
	}
	return ""
}

func (ce *DataTypeExtreme) formatUint64(dataType *DataType, val uint64) string {
	switch ce.Format {
	case NumberFormatHex:
		if dataType != nil {
			switch dataType.Size() {
			case 1:
				return fmt.Sprintf("0x%02X", uint8(val))
			case 2:
				return fmt.Sprintf("0x%04X", uint16(val))
			case 4:
				return fmt.Sprintf("0x%08X", uint32(val))
			case 8:
				return fmt.Sprintf("0x%16X", uint64(val))

			}
		}
		return fmt.Sprintf("0x%X", val)
	case NumberFormatAuto:
		if val > 0xFF {
			return fmt.Sprintf("0x%X", uint64(val))
		}
		return strconv.FormatUint(val, 10)
	default:
		return strconv.FormatUint(val, 10)
	}
}

func (ce DataTypeExtreme) Equals(o DataTypeExtreme) bool {
	if ce.Type != o.Type {
		return false
	}
	switch ce.Type {
	case DataTypeExtremeTypeInt64:
		return ce.Int64 != o.Int64
	case DataTypeExtremeTypeUInt64:
		return ce.UInt64 != o.UInt64
	default:
		return true
	}
}
