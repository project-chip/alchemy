package sdk

import (
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/constraint"
	"github.com/project-chip/alchemy/matter/types"
)

func applyErrataToStruct(st *matter.Struct, typeNames map[string]string, typeOverrides *errata.SDKTypes) {
	if typeOverrides != nil {
		ast, ok := typeOverrides.Structs[st.Name]
		if !ok {
			return
		}
		if ast.OverrideName != "" {
			st.Name = ast.OverrideName
		}
		applyErrataToFields(st.Fields, ast)
	}
	st.Name = applyTypeName(typeNames, st.Name)

}

func applyErrataToFields(fs matter.FieldSet, override *errata.SDKType) {
	if len(override.Fields) != 0 {
		for _, f := range override.Fields {
			for _, field := range fs {
				if field.Name == f.Name {
					applyErrataToField(field, f)
					break
				}
			}
		}
	}
}

func applyErrataToField(field *matter.Field, override *errata.SDKType) {
	if override.OverrideName != "" {
		field.Name = override.OverrideName
	}
	if override.OverrideType != "" {
		field.Type = types.ParseDataType(override.OverrideType, false)
	}
	if override.Conformance != "" {
		field.Conformance = conformance.ParseConformance(override.Conformance)
	}
	if override.Constraint != "" {
		field.Constraint = constraint.ParseString(override.Constraint)
	}
	if override.Fallback != "" {
		field.Fallback = constraint.ParseLimit(override.Fallback)
	}
	field.Quality = overrideQuality(override, field.Quality)
	field.Access = overrideAccess(override, field.EntityType(), field.Access)
}
