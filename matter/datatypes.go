package matter

import (
	"iter"
	"log/slog"
	"strings"

	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/internal/text"
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
		if text.HasCaseInsensitiveSuffix(dataType, suffix) {
			dataType = dataType[0 : len(dataType)-len(suffix)]
			break
		}
	}
	return dataType
}

type AssociatedDataTypes struct {
	parentEntity types.Entity

	Bitmaps   BitmapSet   `json:"bitmaps,omitempty"`
	Enums     EnumSet     `json:"enums,omitempty"`
	Structs   StructSet   `json:"structs,omitempty"`
	TypeDefs  TypeDefSet  `json:"typedefs,omitempty"`
	Constants ConstantSet `json:"constants,omitempty"`
}

func (adt *AssociatedDataTypes) AddDataTypes(entities ...types.Entity) {
	for _, entity := range entities {
		switch entity := entity.(type) {
		case *Bitmap:
			adt.AddBitmaps(entity)
		case *Enum:
			adt.AddEnums(entity)
		case *Struct:
			adt.AddStructs(entity)
		case *TypeDef:
			adt.AddTypeDefs(entity)
		case *Constant:
			adt.AddConstants(entity)
		}
	}
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

func (adt *AssociatedDataTypes) AddConstants(constants ...*Constant) {
	for _, c := range constants {
		if c.parent != nil {
			if _, ok := c.parent.(*ClusterGroup); !ok && c.parent != adt.parentEntity {
				slog.Warn("Constant belongs to multiple parents", slog.String("name", c.Name), log.Path("source", c), LogEntity("parent", adt.parentEntity))
			}
			continue
		}
		c.parent = adt.parentEntity
	}
	adt.Constants = append(adt.Constants, constants...)
}

// This really exists to allow patching ZAP
func (adt *AssociatedDataTypes) MoveEnum(en *Enum) {
	adt.Enums = append(adt.Enums, en)
	en.parent = adt.parentEntity
}

// This really exists to allow patching ZAP
func (adt *AssociatedDataTypes) MoveStruct(s *Struct) {
	adt.Structs = append(adt.Structs, s)
	s.parent = adt.parentEntity
}

type DataTypeIterator iter.Seq2[types.Entity, types.Entity]

func iterateOverDataTypes(c *Cluster) DataTypeIterator {
	return func(yield func(types.Entity, types.Entity) bool) {
		if c.Features != nil {
			if !yield(c, c.Features) {
				return
			}
		}
		for _, en := range c.Bitmaps {
			if !yield(c, en) {
				return
			}
		}
		for _, en := range c.Enums {
			if !yield(c, en) {
				return
			}
		}
		for _, en := range c.Structs {
			if !yield(c, en) {
				return
			}
			for _, f := range en.Fields {
				for p, e := range iterateOverFieldDataTypes(f) {
					if !yield(p, e) {
						return
					}
				}
			}
		}
		for _, cmd := range c.Commands {
			if !yield(c, cmd) {
				return
			}
			for _, f := range cmd.Fields {
				for p, e := range iterateOverFieldDataTypes(f) {
					if !yield(p, e) {
						return
					}
				}
			}
			if cmd.Response != nil && cmd.Response.Entity != nil && cmd.Response.Name == "" {
				if !yield(cmd, cmd.Response.Entity) {
					return
				}
			}
			for _, ev := range c.Events {
				if !yield(c, ev) {
					return
				}
				for _, f := range ev.Fields {
					for p, e := range iterateOverFieldDataTypes(f) {
						if !yield(p, e) {
							return
						}
					}
				}
			}
			for _, a := range c.Attributes {

				if !yield(c, a) {
					return
				}
				for p, e := range iterateOverFieldDataTypes(a) {
					if !yield(p, e) {
						return
					}
				}
			}
		}
	}
}

func iterateOverFieldDataTypes(field *Field) DataTypeIterator {
	return func(yield func(types.Entity, types.Entity) bool) {
		if field.Type == nil {
			return
		}
		var entity types.Entity
		if field.Type.IsArray() {
			if field.Type.EntryType == nil {
				return
			}
			if field.Type.EntryType.Entity == nil {
				return
			}
			entity = field.Type.EntryType.Entity
		} else {
			entity = field.Type.Entity
		}
		if entity == nil {
			return
		}
		if !yield(field, entity) {
			return
		}
		switch entity := entity.(type) {
		case *Struct:
			for _, f := range entity.Fields {
				for p, e := range iterateOverFieldDataTypes(f) {
					if !yield(p, e) {
						return
					}
				}
			}
		}

	}
}
