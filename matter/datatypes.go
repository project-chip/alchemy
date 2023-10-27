package matter

import (
	"strings"
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

type Enum struct {
	Name        string
	Description string
	Type        string
	Values      []*EnumValue
}

type EnumValue struct {
	Value       string
	Name        string
	Summary     string
	Conformance string
}

type Bitmap struct {
	Name        string
	Description string
	Type        string
	Bits        []*BitmapValue
}

type BitmapValue struct {
	Bit         string
	Name        string
	Summary     string
	Conformance string
}

type Struct struct {
	Name        string
	Description string
	Fields      []*Field
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
	Name    string
	IsArray bool
}
