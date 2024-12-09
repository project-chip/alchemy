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

var typeSuffixes = []string{" Attribute", " Device Type", " Type", " Field", " Command", " Attribute", " Event"}

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
		parentEntity := bm.Parent()
		if parentEntity != nil {
			if _, ok := parentEntity.(*ClusterGroup); !ok && bm.parent != adt.parentEntity {
				slog.Warn("Bitmap belongs to multiple parents", slog.String("name", bm.Name), log.Path("source", bm), LogEntity("parent", adt.parentEntity))
			}
			continue
		}
		bm.parent = adt.parentEntity
	}
	adt.Bitmaps = append(adt.Bitmaps, bitmaps...)
}

func (adt *AssociatedDataTypes) AddEnums(enums ...*Enum) {
	for _, e := range enums {
		if e.parent != nil {
			if _, ok := e.parent.(*ClusterGroup); !ok && e.parent != adt.parentEntity {
				slog.Warn("Enum belongs to multiple parents", slog.String("name", e.Name), log.Path("source", e), LogEntity("parent", adt.parentEntity))
			}
			continue
		}
		e.parent = adt.parentEntity
	}
	adt.Enums = append(adt.Enums, enums...)
}

func (adt *AssociatedDataTypes) AddStructs(structs ...*Struct) {
	for _, s := range structs {
		if s.parent != nil {
			if _, ok := s.parent.(*ClusterGroup); !ok && s.parent != adt.parentEntity {
				slog.Warn("Struct belongs to multiple parents", slog.String("name", s.Name), log.Path("source", s), LogEntity("parent", adt.parentEntity))
			}
			continue
		}
		s.parent = adt.parentEntity
	}
	adt.Structs = append(adt.Structs, structs...)
}

func (adt *AssociatedDataTypes) AddTypeDefs(typeDefs ...*TypeDef) {
	for _, td := range typeDefs {
		if td.parent != nil {
			if _, ok := td.parent.(*ClusterGroup); !ok && td.parent != adt.parentEntity {
				slog.Warn("TypeDef belongs to multiple parents", slog.String("name", td.Name), log.Path("source", td), LogEntity("parent", adt.parentEntity))
			}
			continue
		}
		td.parent = adt.parentEntity
	}
	adt.TypeDefs = append(adt.TypeDefs, typeDefs...)
}

// This really exists to allow patching ZAP
func (adt *AssociatedDataTypes) MoveStruct(s *Struct) {
	adt.Structs = append(adt.Structs, s)
	s.parent = adt.parentEntity
}
