package matter

import (
	"log/slog"

	"github.com/project-chip/alchemy/internal/log"
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
		args = append(args, log.Path("source", entity))
	case *Features:
		args = append(args, slog.String("type", "features"))
		args = append(args, log.Path("source", entity))
	case *Bitmap:
		args = append(args, slog.String("type", "bitmap"))
		args = append(args, slog.String("name", entity.Name))
		args = append(args, log.Path("source", entity))
	case *Enum:
		args = append(args, slog.String("type", "enum"))
		args = append(args, slog.String("name", entity.Name))
		args = append(args, log.Path("source", entity))
	case *Struct:
		args = append(args, slog.String("type", "struct"))
		args = append(args, slog.String("name", entity.Name))
		args = append(args, log.Path("source", entity))
	case *Command:
		args = append(args, slog.String("type", "command"))
		args = append(args, slog.String("name", entity.Name))
		args = append(args, log.Path("source", entity))
	case *Event:
		args = append(args, slog.String("type", "event"))
		args = append(args, slog.String("name", entity.Name))
		args = append(args, log.Path("source", entity))
	case *Cluster:
		args = append(args, slog.String("type", "cluster"))
		args = append(args, slog.String("name", entity.Name))
		args = append(args, log.Path("source", entity))
	case *ClusterGroup:
		args = append(args, slog.String("type", "clusterGroup"))
		args = append(args, slog.String("name", entity.Name))
		args = append(args, log.Path("source", entity))
	case *TypeDef:
		args = append(args, slog.String("type", "typeDef"))
		args = append(args, slog.String("name", entity.Name))
		args = append(args, log.Path("source", entity))
	case *EnumValue:
		args = append(args, slog.String("type", "enumValue"))
		args = append(args, slog.String("name", entity.Name))
		args = append(args, log.Path("source", entity))
	case *Condition:
		args = append(args, slog.String("type", "condition"))
		args = append(args, slog.String("name", entity.Description))
		args = append(args, log.Path("source", entity))
	case *Feature:
		args = append(args, slog.String("type", "feature"))
		args = append(args, slog.String("name", entity.Code))
		args = append(args, log.Path("source", entity))
	case *Namespace:
		args = append(args, slog.String("type", "namespace"))
		args = append(args, slog.String("name", entity.Name))
		args = append(args, log.Path("source", entity))
	case Bit:
		args = append(args, slog.String("type", "bit"))
		args = append(args, slog.String("name", entity.Name()))
		args = append(args, log.Path("source", entity))
	case *ClusterRequirement:
		args = append(args, slog.String("type", "clusterRequirement"))
		args = append(args, slog.String("clusterId", entity.ClusterID.HexString()))
		args = append(args, slog.String("clusterName", entity.ClusterName))
		args = append(args, slog.String("interface", entity.Interface.String()))
		args = append(args, log.Path("source", entity))
	case *Constant:
		args = append(args, slog.String("type", "constant"))
		args = append(args, slog.String("name", entity.Name))
		args = append(args, slog.Any("value", entity.Value))
		args = append(args, log.Path("source", entity))
	case nil:
		args = append(args, slog.String("type", "nil"))
	default:
		args = append(args, log.Type("unknownEntityType", en))
	}
	if en != nil {
		parent := en.Parent()
		if parent != nil {
			args = append(args, LogEntity("parent", en.Parent()))
		}
	}

	return slog.Group(key, args...)
}
