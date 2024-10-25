package testplan

import (
	"fmt"
	"strings"

	"github.com/mailgun/raymond/v2"
	"github.com/project-chip/alchemy/internal/handlebars"
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

func entityIdentifierHelper(entity types.Entity) raymond.SafeString {
	return raymond.SafeString(entityIdentifier(entity))
}

func entityIdentifierPaddedHelper(list any, entity types.Entity) raymond.SafeString {
	var longest int
	for entity := range handlebars.Iterate[types.Entity](list) {
		id := entityIdentifier(entity)
		if len(id) > longest {
			longest = len(id)
		}
	}
	return raymond.SafeString(fmt.Sprintf("%-*s", longest, entityIdentifier(entity)))
}

func entityIdentifierPaddingHelper(list any, entity types.Entity) raymond.SafeString {
	var longest int
	for entity := range handlebars.Iterate[types.Entity](list) {
		id := entityIdentifier(entity)
		if len(id) > longest {
			longest = len(id)
		}
	}
	id := entityIdentifier(entity)
	return raymond.SafeString(strings.Repeat(" ", longest-len(id)))
}

func idHelper(id matter.Number, options *raymond.Options) raymond.SafeString {
	format := options.HashStr("format")
	if format == "" {
		format = "%04X"
	}
	return raymond.SafeString(fmt.Sprintf(format, id.Value()))
}

func shortIdHelper(id matter.Number) raymond.SafeString {
	return raymond.SafeString(fmt.Sprintf("%02X", id.Value()))
}
