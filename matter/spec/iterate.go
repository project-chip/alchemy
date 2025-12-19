package spec

import (
	"iter"

	"github.com/project-chip/alchemy/asciidoc/parse"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

type EntityCallback func(parent types.Entity, entity types.Entity) parse.SearchShould

func traverseYieldEntity(parent types.Entity, entity types.Entity, callback EntityCallback) bool {
	switch callback(parent, entity) {
	case parse.SearchShouldStop:
		return false
	}
	return true
}

/*
func traverseYieldParentEntity(parent types.Entity, entity types.Entity, children iter.Seq[types.Entity], callback EntityCallback) bool {
	if parent != nil && !traverseYieldEntity(parent, entity, callback) {
		return false
	}
	for child := range children {
		if !traverseYieldEntity(entity, child, callback) {
			return false
		}
		var entityFields matter.FieldSet
		var fieldEntity types.Entity
		switch child := child.(type) {
		case *matter.Field:
			if child.Type == nil {
				continue
			}
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
			switch entity := fieldEntity.(type) {
			case *matter.Struct:
				entityFields = entity.Fields
			case *matter.Command:
				entityFields = entity.Fields
			case *matter.Event:
				entityFields = entity.Fields
			}
		case *matter.Command:
		}
		if !traverseYieldParentEntity(child, fieldEntity, entityFields.Iterate(), callback) {
			return false
		}

	}
	return true
}*/

func traverseEntityList(parent types.Entity, root types.Entity, children iter.Seq[types.Entity], callback EntityCallback) bool {
	if parent != nil {
		switch callback(parent, root) {
		case parse.SearchShouldStop:
			return false
		case parse.SearchShouldSkip:
			return true
		case parse.SearchShouldContinue:
		}
	}
	for child := range children {
		switch callback(root, child) {
		case parse.SearchShouldStop:
			return false
		case parse.SearchShouldSkip:
			continue
		case parse.SearchShouldContinue:
		}
		switch child := child.(type) {
		case *matter.Field:
			var entityFields matter.FieldSet
			var fieldEntity types.Entity
			if child.Type == nil {
				continue
			}
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
			switch entity := fieldEntity.(type) {
			case *matter.Struct:
				entityFields = entity.Fields
			case *matter.Command:
				entityFields = entity.Fields
			case *matter.Event:
				entityFields = entity.Fields
			}
			if !traverseEntityList(child, fieldEntity, entityFields.Iterate(), callback) {
				return false
			}
		case *matter.Bitmap:
			if !traverseEntityList(nil, child, child.Bits.Iterate(), callback) {
				return false
			}
		case *matter.Enum:
			if !traverseEntityList(nil, child, child.Values.Iterate(), callback) {
				return false
			}
		case *matter.Struct:
			if !traverseEntityList(nil, child, child.Fields.Iterate(), callback) {
				return false
			}
		case *matter.Command:
			if !traverseEntityList(nil, child, child.Fields.Iterate(), callback) {
				return false
			}
		}

	}
	return true
}

func traverseClusterEntities(c *matter.Cluster, callback EntityCallback) parse.SearchShould {
	if c.Features != nil {
		if !traverseEntityList(c, c.Features, c.Features.Iterate(), callback) {
			return parse.SearchShouldStop
		}
	}
	for _, bm := range c.Bitmaps {
		if !traverseEntityList(c, bm, bm.Bits.Iterate(), callback) {
			return parse.SearchShouldStop
		}
	}
	for _, en := range c.Enums {
		if !traverseEntityList(c, en, en.Values.Iterate(), callback) {
			return parse.SearchShouldStop
		}
	}
	for _, s := range c.Structs {
		if !traverseEntityList(c, s, s.Fields.Iterate(), callback) {
			return parse.SearchShouldStop
		}
	}
	for _, cmd := range c.Commands {
		if !traverseEntityList(c, cmd, cmd.Fields.Iterate(), callback) {
			return parse.SearchShouldStop
		}
		if cmd.Response != nil && cmd.Response.Entity != nil && cmd.Response.Name == "" {
			if !traverseYieldEntity(cmd, cmd.Response.Entity, callback) {
				return parse.SearchShouldStop
			}
		}
	}
	for _, ev := range c.Events {
		if !traverseEntityList(c, ev, ev.Fields.Iterate(), callback) {
			return parse.SearchShouldStop
		}
	}
	if !traverseEntityList(nil, c, c.Attributes.Iterate(), callback) {
		return parse.SearchShouldStop
	}
	return parse.SearchShouldContinue
}
