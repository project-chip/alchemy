package pics

import (
	"fmt"
	"strings"

	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

func EntityIdentifier(entity types.Entity) string {
	switch entity := entity.(type) {
	case *matter.Field:
		switch entity.EntityType() {
		case types.EntityTypeAttribute:
			return fmt.Sprintf("A_%s", strings.ToUpper(matter.CaseWithSeparator(entity.Name, '_')))
		}
	case *matter.Feature:
		return fmt.Sprintf("F_%s", entity.Code)
	case *matter.Event:
		return fmt.Sprintf("E_%s", strings.ToUpper(matter.CaseWithSeparator(entity.Name, '_')))
	case *matter.Command:
		return fmt.Sprintf("C_%s", strings.ToUpper(matter.CaseWithSeparator(entity.Name, '_')))
	}
	return fmt.Sprintf("UNKNOWN_TYPE_%T", entity)
}
