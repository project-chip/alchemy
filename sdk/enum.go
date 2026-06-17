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
			for _, f := range override.Fields {
				for _, ev := range en.Values {
					if ev.Name == f.Name {
						if f.OverrideName != "" {
							ev.Name = f.OverrideName
						}
						if f.Conformance != "" {
							ev.Conformance = conformance.ParseConformance(f.Conformance)
						}
						if f.Value != "" {
							ev.Value = matter.ParseNumber(f.Value)
						}
						break
					}
				}
			}
			for _, f := range override.ExtraFields {
				ev := matter.NewEnumValue(en.Source(), en)
				ev.Name = f.Name
				if f.OverrideName != "" {
					ev.Name = f.OverrideName
				}
				if f.Conformance != "" {
					ev.Conformance = conformance.ParseConformance(f.Conformance)
				} else {
					ev.Conformance = conformance.Set{&conformance.Mandatory{}}
				}
				if f.Value != "" {
					ev.Value = matter.ParseNumber(f.Value)
				}
				en.Values = append(en.Values, ev)
			}
		}
	}
	en.Name = applyTypeName(typeNames, en.Name)
}
