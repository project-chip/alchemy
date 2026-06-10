package sdk

import (
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/constraint"
	"github.com/project-chip/alchemy/matter/types"
)

func applyErrataToStruct(st *matter.Struct, typeNames map[string]string, typeOverrides *errata.SDKTypes) error {
	if typeOverrides != nil {
		ast, ok := typeOverrides.Structs[st.Name]
		if !ok {
			return nil
		}
		if ast.OverrideName != "" {
			st.Name = ast.OverrideName
		}
		switch ast.FabricScoping {
		case "none":
			st.FabricScoping = matter.FabricScopingUnscoped
		case "fabric-scoped":
			st.FabricScoping = matter.FabricScopingScoped
		}
		err := applyErrataToFields(st.Fields, ast)
		if err != nil {
			return err
		}
		err = injectExtraFields(st, &st.Fields, types.EntityTypeStructField, ast.ExtraFields)
		if err != nil {
			return err
		}
	}
	st.Name = applyTypeName(typeNames, st.Name)
	return nil
}

func injectExtraFields(parent types.Entity, fields *matter.FieldSet, entityType types.EntityType, extraFields []*errata.SDKType) error {
	for _, f := range extraFields {
		var found bool
		for _, field := range *fields {
			if field.Name == f.Name {
				found = true
				break
			}
		}
		if !found {
			field := matter.NewField(nil, parent, entityType)
			field.Name = f.Name
			if f.Type != "" {
				var rank types.DataTypeRank = types.DataTypeRankScalar
				if f.List {
					rank = types.DataTypeRankList
				}
				field.Type = types.ParseDataType(f.Type, rank)
			}
			err := applyErrataToField(field, f)
			if err != nil {
				return err
			}
			*fields = append(*fields, field)
		}
	}
	return nil
}

func applyErrataToFields(fs matter.FieldSet, override *errata.SDKType) error {
	if len(override.Fields) != 0 {
		for _, f := range override.Fields {
			for _, field := range fs {
				if field.Name == f.Name {
					err := applyErrataToField(field, f)
					if err != nil {
						return err
					}
					break
				}
			}
		}
	}
	return nil
}

func applyErrataToField(field *matter.Field, override *errata.SDKType) error {
	if override.OverrideName != "" {
		field.Name = override.OverrideName
	}
	if override.OverrideType != "" {
		rank := types.DataTypeRankScalar
		if override.List {
			rank = types.DataTypeRankList
		}
		field.Type = types.ParseDataType(override.OverrideType, rank)
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
	if override.Value != "" {
		field.ID = matter.ParseNumber(override.Value)
	}
	field.Quality = overrideQuality(override, field.Quality)
	access, err := overrideAccess(override, field.EntityType(), field.Access)
	if err != nil {
		return err
	}
	field.Access = access
	return nil
}
