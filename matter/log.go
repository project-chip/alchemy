package matter

import (
	"log/slog"

	"github.com/project-chip/alchemy/matter/types"
)

func LogEntity(en types.Entity) slog.Attr {
	var args []any
	switch entity := en.(type) {
	case *Bitmap:
		args = append(args, slog.String("type", "bitmap"))
		args = append(args, slog.String("name", entity.Name))
	case *Enum:
		args = append(args, slog.String("type", "enum"))
		args = append(args, slog.String("name", entity.Name))
	case *Struct:
		args = append(args, slog.String("type", "struct"))
		args = append(args, slog.String("name", entity.Name))
	default:
		args = append(args, slog.String("type", en.EntityType().String()))
	}
	return slog.Group("entity", args...)
}
