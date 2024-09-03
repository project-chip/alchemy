package spec

import (
	"log/slog"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

var AllowedQualities = map[types.EntityType]matter.Quality{
	types.EntityTypeAttribute:    matter.QualitySourceAttribution | matter.QualityChangedOmitted | matter.QualityFixed | matter.QualityNonVolatile | matter.QualityReportable | matter.QualityQuieterReporting | matter.QualityScene | matter.QualityAtomicWrite | matter.QualityNullable,
	types.EntityTypeCommand:      matter.QualityLargeMessage,
	types.EntityTypeCommandField: matter.QualitySourceAttribution | matter.QualityChangedOmitted | matter.QualityFixed | matter.QualityNonVolatile | matter.QualityReportable | matter.QualityQuieterReporting | matter.QualityScene | matter.QualityAtomicWrite | matter.QualityNullable,
	types.EntityTypeStructField:  matter.QualitySourceAttribution | matter.QualityChangedOmitted | matter.QualityFixed | matter.QualityNonVolatile | matter.QualityReportable | matter.QualityQuieterReporting | matter.QualityScene | matter.QualityAtomicWrite | matter.QualityNullable,
	types.EntityTypeEvent:        matter.QualitySourceAttribution,
	types.EntityTypeEventField:   matter.QualitySourceAttribution | matter.QualityChangedOmitted | matter.QualityFixed | matter.QualityNonVolatile | matter.QualityReportable | matter.QualityQuieterReporting | matter.QualityScene | matter.QualityAtomicWrite | matter.QualityNullable,
	types.EntityTypeCluster:      matter.QualitySingleton | matter.QualityDiagnostics,
}

func parseQuality(s string, entityType types.EntityType, doc *Doc, element asciidoc.Element) matter.Quality {

	q := matter.ParseQuality(s)
	if allowed, ok := AllowedQualities[entityType]; ok {
		if q&allowed != q {
			slog.Warn("Invalid quality on entity", slog.String("quality", q.String()), log.Path("path", NewSource(doc, element)))
		}
	}
	return q
}
