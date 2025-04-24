package suggest

import (
	"cmp"
	"fmt"
	"iter"
	"log/slog"
	"slices"
	"strings"

	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

func PossibleEntities(incorrect string, possibilities map[types.Entity]int, list iter.Seq2[string, types.Entity]) {
	for possibility, entity := range list {
		l := levenshtein(incorrect, possibility)
		if strings.HasPrefix(possibility, incorrect) || l <= 2 {
			existing, ok := possibilities[entity]
			if !ok || l < existing {
				possibilities[entity] = l
			}
		}
	}
}

type suggestionScore struct {
	entity types.Entity
	score  int
}

func ListPossibilities(identifier string, possibilities map[types.Entity]int) {
	switch len(possibilities) {
	case 0:
		return
	default:
		slog.Info(fmt.Sprintf("By \"%s\", did you mean:", identifier))
		entities := make([]suggestionScore, 0, len(possibilities))
		for entity, score := range possibilities {
			entities = append(entities, suggestionScore{entity: entity, score: score})
		}
		slices.SortFunc(entities, func(a, b suggestionScore) int {
			return cmp.Compare(a.score, b.score)
		})

		for _, entitySore := range entities {
			entity := entitySore.entity
			parent := entity.Parent()
			switch entity := entity.(type) {
			case *matter.Feature:
				slog.Info(fmt.Sprintf("\tThe feature \"%s\" (%s)", entity.Code, entity.Name()), log.Path("source", entity))
			case matter.Bit:
				if parent != nil {
					slog.Info(fmt.Sprintf("\tThe \"%s\" bit on the %s bitmap", entity.Name(), matter.EntityName(parent)), log.Path("source", entity))
				} else {
					slog.Info(fmt.Sprintf("\tThe \"%s\" bit", entity.Name()), log.Path("source", entity))
				}
			case *matter.EnumValue:
				if parent != nil {
					slog.Info(fmt.Sprintf("\tThe \"%s\" enum value on the %s enum", entity.Name, matter.EntityName(parent)), log.Path("source", entity))
				} else {
					slog.Info(fmt.Sprintf("\tThe \"%s\" enum value", entity.Name), log.Path("source", entity))
				}
			case *matter.Field:
				switch entity.EntityType() {
				case types.EntityTypeAttribute:
					if parent != nil {
						slog.Info(fmt.Sprintf("\tThe \"%s\" attribute on the %s cluster", entity.Name, matter.EntityName(parent)), log.Path("source", entity))
					} else {
						slog.Info(fmt.Sprintf("\tThe \"%s\" attribute", entity.Name), log.Path("source", entity))
					}
				default:
					if parent != nil {
						slog.Info(fmt.Sprintf("\tThe \"%s\" field on the %s %s", entity.Name, matter.EntityName(parent), parent.EntityType().String()), log.Path("source", entity))
					} else {
						slog.Info(fmt.Sprintf("\tThe \"%s\" field", entity.Name), log.Path("source", entity))
					}
				}
			case *matter.Condition:
				if parent != nil {
					slog.Info(fmt.Sprintf("\tThe \"%s\" condition (%s) on the %s device type", entity.Feature, entity.Description, matter.EntityName(parent)), log.Path("source", entity))
				} else {
					slog.Info(fmt.Sprintf("\tThe \"%s\" condition (%s)", entity.Feature, entity.Description), log.Path("source", entity))
				}
			default:
				slog.Info("\tThe entity", matter.LogEntity("entity", entity))
			}
		}
	}
}
