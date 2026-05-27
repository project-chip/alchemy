package sdk

import (
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/types"
)

func applyErrataToEnum(en *matter.Enum, typeNames map[string]string, typeOverrides *errata.SDKTypes) {
	if typeOverrides != nil {
		override, ok := typeOverrides.Enums[en.Name]
		if ok {
			if override.OverrideName != "" {
				en.Name = override.OverrideName
			}
			if override.OverrideType != "" {
				en.Type = types.ParseDataType(override.OverrideType, types.DataTypeRankScalar)
			}
			if len(override.Fields) == 0 {
				return
			}
			for _, f := range override.Fields {
				for _, ev := range en.Values {
					if ev.Name == f.Name {
						if f.OverrideName != "" {
							ev.Name = f.OverrideName
						}
						if f.Conformance != "" {
							ev.Conformance = conformance.ParseConformance(f.Conformance)
						}
						break
					}
				}
			}
		}
	}
	en.Name = applyTypeName(typeNames, en.Name)
}
