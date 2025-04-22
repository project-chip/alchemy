package matter

import (
	"log/slog"

	"github.com/project-chip/alchemy/matter/types"
)

func LogEntity(key string, en types.Entity) slog.Attr {
	var args []any
	switch entity := en.(type) {
	case *Field:
		switch en.EntityType() {
		case types.EntityTypeAttribute:
			args = append(args, slog.String("type", "attribute"))
		case types.EntityTypeCommandField:
			args = append(args, slog.String("type", "field"))
		default:
			args = append(args, slog.String("type", en.EntityType().String()))
		}
		args = append(args, slog.String("name", entity.Name))
	case *Bitmap:
		args = append(args, slog.String("type", "bitmap"))
		args = append(args, slog.String("name", entity.Name))
	case *Enum:
		args = append(args, slog.String("type", "enum"))
		args = append(args, slog.String("name", entity.Name))
	case *Struct:
		args = append(args, slog.String("type", "struct"))
		args = append(args, slog.String("name", entity.Name))
	case *Command:
		args = append(args, slog.String("type", "command"))
		args = append(args, slog.String("name", entity.Name))
	case *Event:
		args = append(args, slog.String("type", "event"))
		args = append(args, slog.String("name", entity.Name))
	case *Cluster:
		args = append(args, slog.String("type", "cluster"))
		args = append(args, slog.String("name", entity.Name))
	case *TypeDef:
		args = append(args, slog.String("type", "typeDef"))
		args = append(args, slog.String("name", entity.Name))
	case *EnumValue:
		args = append(args, slog.String("type", "enumValue"))
		args = append(args, slog.String("name", entity.Name))
	case *Condition:
		args = append(args, slog.String("type", "condition"))
		args = append(args, slog.String("name", entity.Description))
	case *Feature:
		args = append(args, slog.String("type", "feature"))
		args = append(args, slog.String("name", entity.Code))
	case Bit:
		args = append(args, slog.String("type", "bit"))
		args = append(args, slog.String("name", entity.Name()))
	case nil:
		args = append(args, slog.String("type", "nil"))
	default:
		args = append(args, slog.String("type", en.EntityType().String()))
	}
	if en != nil {
		parent := en.Parent()
		if parent != nil {
			args = append(args, LogEntity("parent", en.Parent()))
		}
	}

	return slog.Group(key, args...)
}
