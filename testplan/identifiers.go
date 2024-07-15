package testplan

import (
	"fmt"

	"github.com/iancoleman/strcase"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

func entityPICS(entity types.Entity) string {
	return fmt.Sprintf("{PICS_S%s}", entityIdentifier(entity))
}

func entityVariable(entity types.Entity) string {
	return fmt.Sprintf("{%s}", entityIdentifier(entity))
}

func entityIdentifier(entity types.Entity) string {
	switch entity := entity.(type) {
	case *matter.Field:
		switch entity.EntityType() {
		case types.EntityTypeAttribute:
			return fmt.Sprintf("A_%s", strcase.ToScreamingSnake(entity.Name))
		}
	case *matter.Feature:
		return fmt.Sprintf("F_%s", entity.Code)
	case *matter.Event:
		return fmt.Sprintf("E_%s", strcase.ToScreamingSnake(entity.Name))
	case *matter.Command:
		return fmt.Sprintf("C_%s", strcase.ToScreamingSnake(entity.Name))
	}
	return fmt.Sprintf("UNKNOWN_TYPE_%T", entity)
}
