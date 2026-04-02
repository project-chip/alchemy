package matter

import (
	"iter"
	"log/slog"
	"strings"

	"github.com/project-chip/alchemy/asciidoc/parse"
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

type EntityIterator iter.Seq2[types.Entity, types.Entity]

func iterateOverEntities(c *Cluster) EntityIterator {
	return func(yield func(types.Entity, types.Entity) bool) {
		traverseEntities(c, func(parent, entity types.Entity) parse.SearchShould {
			if !yield(parent, entity) {
				return parse.SearchShouldStop
			}
			return parse.SearchShouldContinue
		})
	}
}

func iterateOverFieldDataTypes(field *Field) EntityIterator {
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

type EntityCallback func(parent types.Entity, entity types.Entity) parse.SearchShould

func traverseYieldEntity(parent types.Entity, entity types.Entity, callback EntityCallback) bool {
	switch callback(parent, entity) {
	case parse.SearchShouldStop:
		return false
	}
	return true
}

func traverseYieldParentEntity(parent types.Entity, entity types.Entity, children iter.Seq[types.Entity], callback EntityCallback) bool {
	if entity != nil && !traverseYieldEntity(parent, entity, callback) {
		return false
	}
	for child := range children {
		var entityFields FieldSet
		switch child := child.(type) {
		case *Field:
			if child.Type == nil {
				continue
			}
			var fieldEntity types.Entity
			if child.Type.IsArray() {
				if child.Type.EntryType == nil {
					continue
				}
				if child.Type.EntryType.Entity == nil {
					continue
				}
				fieldEntity = child.Type.EntryType.Entity
			} else {
				fieldEntity = child.Type.Entity
			}
			if fieldEntity == nil {
				continue
			}
			switch entity := entity.(type) {
			case *Struct:
				entityFields = entity.Fields
			case *Command:
				entityFields = entity.Fields
			case *Event:
				entityFields = entity.Fields
			}
			if !traverseYieldParentEntity(child, fieldEntity, entityFields.Iterate(), callback) {
				return false
			}
		}

	}
	return true
}

func traverseEntities(c *Cluster, callback EntityCallback) parse.SearchShould {
	yieldEntity := func(parent types.Entity, entity types.Entity) bool {
		switch callback(parent, entity) {
		case parse.SearchShouldStop:
			return false
		}
		return true
	}

	if c.Features != nil {
		if !traverseYieldEntity(c, c.Features, callback) {
			return parse.SearchShouldStop
		}
	}
	for _, en := range c.Bitmaps {
		if !traverseYieldEntity(c, en, callback) {
			return parse.SearchShouldStop
		}
	}
	for _, en := range c.Enums {
		if !traverseYieldEntity(c, en, callback) {
			return parse.SearchShouldStop
		}
	}
	for _, en := range c.Structs {
		if !traverseYieldParentEntity(c, en, en.Fields.Iterate(), callback) {
			return parse.SearchShouldStop
		}
	}
	for _, cmd := range c.Commands {
		if !traverseYieldParentEntity(c, cmd, cmd.Fields.Iterate(), callback) {
			return parse.SearchShouldStop
		}
		if cmd.Response != nil && cmd.Response.Entity != nil && cmd.Response.Name == "" {
			if !yieldEntity(cmd, cmd.Response.Entity) {
				return parse.SearchShouldStop
			}
		}
	}
	for _, ev := range c.Events {
		if !traverseYieldParentEntity(c, ev, ev.Fields.Iterate(), callback) {
			return parse.SearchShouldStop
		}
	}
	for _, a := range c.Attributes {
		if !yieldEntity(c, a) {
			return parse.SearchShouldStop
		}
		traverseFieldDataTypes(a, callback)
	}
	return parse.SearchShouldContinue
}

func traverseFieldDataTypes(field *Field, callback EntityCallback) parse.SearchShould {
	if field.Type == nil {
		return parse.SearchShouldContinue
	}
	var entity types.Entity
	if field.Type.IsArray() {
		if field.Type.EntryType == nil {
			return parse.SearchShouldContinue
		}
		if field.Type.EntryType.Entity == nil {
			return parse.SearchShouldContinue
		}
		entity = field.Type.EntryType.Entity
	} else {
		entity = field.Type.Entity
	}
	if entity == nil {
		return parse.SearchShouldContinue
	}
	switch callback(field, entity) {
	case parse.SearchShouldSkip:
		return parse.SearchShouldContinue
	case parse.SearchShouldContinue:

	case parse.SearchShouldStop:
		return parse.SearchShouldStop
	}
	var fields FieldSet
	switch entity := entity.(type) {
	case *Struct:
		fields = entity.Fields
	case *Command:
		fields = entity.Fields
	case *Event:
		fields = entity.Fields
	}
	for _, f := range fields {
		switch traverseFieldDataTypes(f, callback) {
		case parse.SearchShouldStop:
			return parse.SearchShouldStop
		}
	}
	return parse.SearchShouldContinue
}
