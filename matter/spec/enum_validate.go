package spec

import (
	"log/slog"

	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter"
)

func validateEnums(spec *Specification, enums []*matter.Enum) {
	for _, e := range enums {
		validateEnum(spec, e)
	}
}

func validateEnum(spec *Specification, en *matter.Enum) {
	idu := make(idUniqueness[*matter.EnumValue])
	nu := make(nameUniqueness[*matter.EnumValue])
	cv := make(conformanceValidation)
	for _, ev := range en.Values {
		if !ev.Value.Valid() {
			slog.Warn("Enum value has invalid ID", log.Path("source", ev), matter.LogEntity("parent", en), slog.String("valueName", en.Name))
			continue
		}
		cv.add(ev, ev.Conformance)
		idu.check(spec, ev.Value, ev)
		nu.check(spec, ev)
	}
	cv.check(spec)
}
