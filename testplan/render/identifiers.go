package render

import (
	"fmt"
	"strings"

	"github.com/mailgun/raymond/v2"
	"github.com/project-chip/alchemy/internal/handlebars"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
	"github.com/project-chip/alchemy/testplan/pics"
)

func entityPICS(entity types.Entity) string {
	return fmt.Sprintf("{PICS_S%s}", pics.EntityIdentifier(entity))
}

func entityVariable(entity types.Entity) string {
	return fmt.Sprintf("{%s}", pics.EntityIdentifier(entity))
}

func entityIdentifierHelper(entity types.Entity) raymond.SafeString {
	return raymond.SafeString(pics.EntityIdentifier(entity))
}

func entityIdentifierPaddedHelper(list any, entity types.Entity) raymond.SafeString {
	var longest int
	for entity := range handlebars.Iterate[types.Entity](list) {
		id := pics.EntityIdentifier(entity)
		if len(id) > longest {
			longest = len(id)
		}
	}
	return raymond.SafeString(fmt.Sprintf("%-*s", longest, pics.EntityIdentifier(entity)))
}

func entityIdentifierPaddingHelper(list any, entity types.Entity) raymond.SafeString {
	var longest int
	for entity := range handlebars.Iterate[types.Entity](list) {
		id := pics.EntityIdentifier(entity)
		if len(id) > longest {
			longest = len(id)
		}
	}
	id := pics.EntityIdentifier(entity)
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
