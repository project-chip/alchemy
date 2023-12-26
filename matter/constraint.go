package matter

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type ConstraintType uint8

const (
	ConstraintTypeUnknown ConstraintType = iota
	ConstraintTypeAll                    // Special section type for everything that comes before any known sections
	ConstraintTypeDescribed
	ConstraintTypeExact
	ConstraintTypeGeneric
	ConstraintTypeList
	ConstraintTypeMax
	ConstraintTypeMin
	ConstraintTypeRange
	ConstraintTypeSet
)

var nameConstraintType map[string]ConstraintType

var constraintTypeNames = map[ConstraintType]string{
	ConstraintTypeUnknown:   "unknown",
	ConstraintTypeAll:       "all",
	ConstraintTypeDescribed: "described",
	ConstraintTypeExact:     "exact",
	ConstraintTypeGeneric:   "generic",
	ConstraintTypeList:      "list",
	ConstraintTypeMax:       "max",
	ConstraintTypeMin:       "min",
	ConstraintTypeRange:     "range",
	ConstraintTypeSet:       "set",
}

func init() {
	nameConstraintType = make(map[string]ConstraintType, len(constraintTypeNames))
	for i, q := range constraintTypeNames {
		nameConstraintType[q] = i
	}
}

func (ct ConstraintType) MarshalJSON() ([]byte, error) {
	v, ok := constraintTypeNames[ct]
	if !ok {
		return nil, fmt.Errorf("unknown constraint type %d", ct)
	}
	return json.Marshal(v)
}

func (c *ConstraintType) UnmarshalJSON(bytes []byte) error {
	var name string
	err := json.Unmarshal(bytes, &name)
	if err != nil {
		return err
	}
	v, ok := nameConstraintType[name]
	if !ok {
		return fmt.Errorf("unknown constraint type %s", name)
	}
	*c = v
	return nil
}

type Constraint interface {
	Type() ConstraintType
	AsciiDocString(dataType *DataType) string
	Equal(o Constraint) bool
	Min(c *ConstraintContext) (min ConstraintExtreme)
	Max(c *ConstraintContext) (max ConstraintExtreme)
	Default(c *ConstraintContext) (max ConstraintExtreme)
	Clone() Constraint
}

type ConstraintLimit interface {
	AsciiDocString(dataType *DataType) string
	Equal(o ConstraintLimit) bool
	Min(c *ConstraintContext) (min ConstraintExtreme)
	Max(c *ConstraintContext) (max ConstraintExtreme)
	Default(c *ConstraintContext) (max ConstraintExtreme)
	Clone() ConstraintLimit
}

type ConstraintContext struct {
	Field             *Field
	Fields            FieldSet
	VisitedReferences map[string]struct{}
}

type ConstraintExtremeType uint8

const (
	ConstraintExtremeTypeUndefined ConstraintExtremeType = iota
	ConstraintExtremeTypeInt64
	ConstraintExtremeTypeUInt64
	ConstraintExtremeTypeNull
	ConstraintExtremeTypeEmpty
)

type NumberFormat uint8

const (
	NumberFormatUndefined NumberFormat = iota
	NumberFormatInt
	NumberFormatHex
	NumberFormatAuto
)

type ConstraintExtreme struct {
	Type   ConstraintExtremeType
	Format NumberFormat
	Int64  int64
	UInt64 uint64
}

func NewIntConstraintExtreme(i int64, f NumberFormat) ConstraintExtreme {
	return ConstraintExtreme{Type: ConstraintExtremeTypeInt64, Format: f, Int64: i}
}

func NewUintConstraintExtreme(u uint64, f NumberFormat) ConstraintExtreme {
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
	case ConstraintExtremeTypeUInt64:
		return ce.formatUint64(dataType, ce.UInt64)
	case ConstraintExtremeTypeNull:
		val := dataType.NullValue()
		if val > 0 {
			return ce.formatUint64(dataType, val)
		}
	case ConstraintExtremeTypeEmpty:
		return ""
	}
	return ""
}

func (ce *ConstraintExtreme) DataModelString(dataType *DataType) string {
	switch ce.Type {
	case ConstraintExtremeTypeInt64:
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
	case ConstraintExtremeTypeUInt64:
		if dataType != nil {
			switch dataType.BaseType {
			case BaseDataTypeBoolean:
				return strconv.FormatBool(ce.Int64 == 1)
			}
		}
		return ce.formatUint64(dataType, ce.UInt64)
	case ConstraintExtremeTypeNull:
		return "null"
	case ConstraintExtremeTypeEmpty:
		return "empty"
	}
	return ""
}

func (ce *ConstraintExtreme) formatUint64(dataType *DataType, val uint64) string {
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
		return true
	}
}
