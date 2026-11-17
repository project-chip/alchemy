package spec

import (
	"log/slog"

	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

func validateFields(spec *Specification, cluster *matter.Cluster, parent types.Entity, fields matter.FieldSet, entityType types.EntityType) {
	idu := make(idUniqueness[*matter.Field])
	nu := make(nameUniqueness[*matter.Field])
	cv := make(conformanceValidation)
	for _, f := range fields {
		if !f.ID.Valid() {
			slog.Warn("Field has invalid ID", log.Path("source", f), matter.LogEntity("parent", parent), slog.String("fieldName", f.Name))
		}
		fieldId := f.ID.Value()
		idu.check(spec, f.ID, f)
		nu.check(spec, f)
		if fieldId >= 0xFE && fieldId <= 0x100 {
			slog.Warn("Struct is using global field ID", log.Path("source", f), matter.LogEntity("parent", parent), slog.String("fieldName", f.Name), slog.String("fieldId", f.ID.HexString()))
		}
		cv.add(f, f.Conformance)

		validateAccess(spec, f, f.Access)
		validateFieldDataType(spec, cluster, f)
		validateQuality(spec, f.Quality, entityType, f)
		validateFieldConstraint(spec, f)
		validateConformance(spec, f, f.Conformance)
		validateFallback(spec, f)
	}
	cv.check(spec)
}
