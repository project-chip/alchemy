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
	default:
		args = append(args, slog.String("type", en.EntityType().String()))
	}
	parent := en.Parent()
	if parent != nil {
		args = append(args, LogEntity("parent", en.Parent()))
	}

	return slog.Group(key, args...)
}
