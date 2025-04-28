package validate

import (
	"log/slog"

	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
)

func validateStructs(spec *spec.Specification) {
	for c := range spec.Clusters {
		for _, s := range c.Structs {
			validateFields(s)
		}
	}
	for obj := range spec.GlobalObjects {
		switch obj := obj.(type) {
		case *matter.Struct:
			validateFields(obj)
		}
	}
}

func validateFields(s *matter.Struct) {
	fieldIds := make(map[uint64]*matter.Field)
	for _, f := range s.Fields {
		if !f.ID.Valid() {
			slog.Warn("Field has invalid ID", log.Path("source", f), slog.String("structName", s.Name), slog.String("fieldName", f.Name))
		}
		fieldId := f.ID.Value()
		existing, ok := fieldIds[fieldId]
		if ok {
			slog.Warn("Duplicate field ID", log.Path("source", f), slog.String("structName", s.Name), slog.String("fieldId", f.ID.HexString()), slog.String("fieldName", f.Name), slog.String("previousFieldName", existing.Name))
		} else {
			fieldIds[fieldId] = f
		}
		if fieldId >= 0xFE {
			slog.Warn("Struct is using global field ID", log.Path("source", f), slog.String("structName", s.Name), slog.String("fieldName", f.Name), slog.String("fieldId", f.ID.HexString()))
		}
	}
}
