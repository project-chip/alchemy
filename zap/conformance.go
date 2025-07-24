package zap

import (
	"log/slog"

	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/types"
)

func IsDisallowed(entity types.Entity, c conformance.Set) bool {
	if len(c) == 0 {
		return false
	}
	switch entity := entity.(type) {
	case *matter.EnumValue:
		if entity.Name == "LocalOptimization" {
			slog.Info("conformance", matter.LogEntity("entity", entity))

		}
	}

	cxt := conformance.Context{}
	conf, err := c.Eval(cxt)
	if err != nil {
		slog.Error("failed evaluating conformance", matter.LogEntity("entity", entity), slog.Any("error", err), log.Path("source", entity))
		return true
	}
	return conf.State == conformance.StateDisallowed && conf.Confidence == conformance.ConfidenceDefinite
}
