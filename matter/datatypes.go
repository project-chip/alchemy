package matter

import (
	"log/slog"
	"strings"

	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter/types"
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

var typeSuffixes = []string{" Attribute", " Type", " Field", " Command", " Attribute", " Event"}

func StripTypeSuffixes(dataType string) string {
	for _, suffix := range typeSuffixes {
		if strings.HasSuffix(dataType, suffix) {
			dataType = dataType[0 : len(dataType)-len(suffix)]
			break
		}
	}
	return dataType
}

type AssociatedDataTypes struct {
	parentEntity types.Entity

	Bitmaps  BitmapSet  `json:"bitmaps,omitempty"`
	Enums    EnumSet    `json:"enums,omitempty"`
	Structs  StructSet  `json:"structs,omitempty"`
	TypeDefs TypeDefSet `json:"typedefs,omitempty"`
}

func (adt *AssociatedDataTypes) AddBitmaps(bitmaps ...*Bitmap) {
	for _, bm := range bitmaps {
		if bm.ParentEntity != nil {
			if _, ok := bm.ParentEntity.(*ClusterGroup); !ok {
				slog.Warn("Bitmap belongs to multiple parents", slog.String("name", bm.Name), log.Path("source", bm), LogEntity(adt.parentEntity))
			}
			continue
		}
		bm.ParentEntity = adt.parentEntity
	}
	adt.Bitmaps = append(adt.Bitmaps, bitmaps...)
}

func (adt *AssociatedDataTypes) AddEnums(enums ...*Enum) {
	for _, e := range enums {
		if e.ParentEntity != nil {
			if _, ok := e.ParentEntity.(*ClusterGroup); !ok {
				slog.Warn("Enum belongs to multiple parents", slog.String("name", e.Name), log.Path("source", e), LogEntity(adt.parentEntity))
			}
			continue
		}
		e.ParentEntity = adt.parentEntity
	}
	adt.Enums = append(adt.Enums, enums...)
}

func (adt *AssociatedDataTypes) AddStructs(structs ...*Struct) {
	for _, s := range structs {
		if s.ParentEntity != nil {
			if _, ok := s.ParentEntity.(*ClusterGroup); !ok {
				slog.Warn("Struct belongs to multiple parents", slog.String("name", s.Name), log.Path("source", s), LogEntity(adt.parentEntity))
			}
			continue
		}
		s.ParentEntity = adt.parentEntity
	}
	adt.Structs = append(adt.Structs, structs...)
}

func (adt *AssociatedDataTypes) AddTypeDefs(typeDefs ...*TypeDef) {
	for _, td := range typeDefs {
		if td.ParentEntity != nil {
			if _, ok := td.ParentEntity.(*ClusterGroup); !ok {
				slog.Warn("TypeDef belongs to multiple parents", slog.String("name", td.Name), log.Path("source", td), LogEntity(adt.parentEntity))
			}
			continue
		}
		td.ParentEntity = adt.parentEntity
	}
	adt.TypeDefs = append(adt.TypeDefs, typeDefs...)
}
