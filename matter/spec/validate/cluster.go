package validate

import (
	"log/slog"

	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
)

func validateClusters(spec *spec.Specification) {
	for c := range spec.Clusters {
		validateAttributes(c, c.Attributes)
	}

}

func validateAttributes(cluster *matter.Cluster, as matter.FieldSet) {
	fieldIds := make(map[uint64]*matter.Field)
	for _, f := range as {
		if !f.ID.Valid() {
			slog.Warn("Attribute has invalid ID", log.Path("path", f), slog.String("clusterName", cluster.Name), slog.String("fieldName", f.Name))
		}
		fieldId := f.ID.Value()
		existing, ok := fieldIds[fieldId]
		if ok {
			slog.Warn("Duplicate field ID", log.Path("path", f), slog.String("clusterName", cluster.Name), slog.String("fieldId", f.ID.HexString()), slog.String("fieldName", f.Name), slog.String("previousFieldName", existing.Name))
		} else {
			fieldIds[fieldId] = f
		}
		if fieldId >= 0xFE {
			slog.Warn("Attribute is using global field ID", log.Path("path", f), slog.String("clusterName", cluster.Name), slog.String("fieldName", f.Name), slog.String("fieldId", f.ID.HexString()))
		}
	}
}
