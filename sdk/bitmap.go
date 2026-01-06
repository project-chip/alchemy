package sdk

import (
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

func applyErrataToBitmap(bitmap *matter.Bitmap, typeNames map[string]string, typeOverrides *errata.SDKTypes) {
	if typeOverrides != nil {
		override, ok := typeOverrides.Bitmaps[bitmap.Name]

		if ok {
			if override.OverrideName != "" {
				bitmap.Name = override.OverrideName
			}
			if override.OverrideType != "" {
				bitmap.Type = types.ParseDataType(override.OverrideType, false)
			}
			if len(override.Fields) == 0 {
				return
			}
			for _, f := range override.Fields {
				for _, b := range bitmap.Bits {
					if b.Name() == f.Name {
						if f.OverrideName != "" {
							b.SetName(f.OverrideName)
						}
						break
					}
				}
			}
		}
	}
	bitmap.Name = applyTypeName(typeNames, bitmap.Name)
}
