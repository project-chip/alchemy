package spec

import (
	"log/slog"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

func parseQuality(s string, entityType types.EntityType, doc *Doc, element asciidoc.Element) matter.Quality {

	q := matter.ParseQuality(s)
	if allowed, ok := matter.AllowedQualities[entityType]; ok {
		if q&allowed != q {
			slog.Warn("Invalid quality on entity", slog.String("quality", q.String()), log.Path("path", NewSource(doc, element)))
		}
	}
	return q
}
