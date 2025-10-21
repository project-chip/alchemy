package spec

import (
	"log/slog"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

func parseQuality(s string, entityType types.EntityType, doc *asciidoc.Document, element asciidoc.Element) matter.Quality {

	q := matter.ParseQuality(s)
	if allowed, ok := matter.AllowedQualities[entityType]; ok {
		if q&allowed != q {
			disallowed := q ^ (q & allowed)
			if disallowed != matter.QualityReportable {
				// This is still in use in a bunch of places to support Zigbee, so we'll ignore it for now
				slog.Warn("Invalid quality on entity", slog.String("disallowed", disallowed.String()), slog.String("entityType", entityType.String()), log.Path("source", NewSource(doc, element)))
			}
		}
	}
	return q
}
