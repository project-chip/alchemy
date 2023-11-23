package matter

import (
	"fmt"
	"strconv"
)

type Constraint interface {
	AsciiDocString(dataType *DataType) string
	Equal(o Constraint) bool
	Min(c *ConstraintContext) (min ConstraintExtreme)
	Max(c *ConstraintContext) (max ConstraintExtreme)
}

type ConstraintLimit interface {
	AsciiDocString(dataType *DataType) string
	Equal(o ConstraintLimit) bool
	Min(c *ConstraintContext) (min ConstraintExtreme)
	Max(c *ConstraintContext) (max ConstraintExtreme)
}

type ConstraintContext struct {
	Fields            FieldSet
	VisitedReferences map[string]struct{}
}

type ConstraintExtremeType uint8

const (
	ConstraintExtremeTypeUndefined ConstraintExtremeType = iota
	ConstraintExtremeTypeInt64
	ConstraintExtremeTypeUInt64
)

type ConstraintExtremeFormat uint8

const (
	ConstraintExtremeFormatUndefined ConstraintExtremeFormat = iota
	ConstraintExtremeFormatInt
	ConstraintExtremeFormatHex
	ConstraintExtremeFormatAuto
)

type ConstraintExtreme struct {
	Type   ConstraintExtremeType
	Format ConstraintExtremeFormat
	Int64  int64
	UInt64 uint64
}

func NewIntConstraintExtreme(i int64, f ConstraintExtremeFormat) ConstraintExtreme {
	return ConstraintExtreme{Type: ConstraintExtremeTypeInt64, Format: f, Int64: i}
}

func NewUintConstraintExtreme(u uint64, f ConstraintExtremeFormat) ConstraintExtreme {
	return ConstraintExtreme{Type: ConstraintExtremeTypeUInt64, Format: f, UInt64: u}
}

func (ce *ConstraintExtreme) Defined() bool {
	return ce.Type != ConstraintExtremeTypeUndefined
}

func (ce *ConstraintExtreme) Value() any {
	switch ce.Type {
	case ConstraintExtremeTypeInt64:
		return ce.Int64
	case ConstraintExtremeTypeUInt64:
		return ce.UInt64
	default:
		return nil
	}
}

func (ce *ConstraintExtreme) ZapString(dataType *DataType) string {
	switch ce.Type {
	case ConstraintExtremeTypeInt64:
		val := ce.Int64
		switch ce.Format {
		case ConstraintExtremeFormatHex:
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
		case ConstraintExtremeFormatAuto:
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
	case ConstraintExtremeTypeUInt64:
		val := ce.UInt64
		switch ce.Format {
		case ConstraintExtremeFormatHex:
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
		case ConstraintExtremeFormatAuto:
			if val > 0xFF {
				return fmt.Sprintf("0x%X", uint64(val))
			}
			return strconv.FormatUint(val, 10)
		default:
			return strconv.FormatUint(val, 10)
		}
	default:
		return ""
	}
}

func (ce ConstraintExtreme) Equals(o ConstraintExtreme) bool {
	if ce.Type != o.Type {
		return false
	}
	switch ce.Type {
	case ConstraintExtremeTypeInt64:
		return ce.Int64 != o.Int64
	case ConstraintExtremeTypeUInt64:
		return ce.UInt64 != o.UInt64
	default:
		return false
	}
}
