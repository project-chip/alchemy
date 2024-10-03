package testplan

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

func entityPICS(entity types.Entity) string {
	return fmt.Sprintf("{PICS_S%s}", entityIdentifier(entity))
}

func entityPICSConformance(entity types.Entity) string {
	return fmt.Sprintf("{PICS_S%s_CONFORMANCE}", entityIdentifier(entity))
}

func entityVariable(entity types.Entity) string {
	return fmt.Sprintf("{%s}", entityIdentifier(entity))
}

func entityIdentifier(entity types.Entity) string {
	switch entity := entity.(type) {
	case *matter.Field:
		switch entity.EntityType() {
		case types.EntityTypeAttribute:
			return fmt.Sprintf("A_%s", strings.ToUpper(strcase.ToCamel(entity.Name)))
		}
	case *matter.Feature:
		return fmt.Sprintf("F_%s", entity.Code)
	case *matter.Event:
		return fmt.Sprintf("E_%s", strings.ToUpper(strcase.ToCamel(entity.Name)))
	case *matter.Command:
		return fmt.Sprintf("C_%s", strings.ToUpper(strcase.ToCamel(entity.Name)))
	}
	return fmt.Sprintf("UNKNOWN_TYPE_%T", entity)
}
