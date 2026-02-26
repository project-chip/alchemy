package sdk

import (
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/types"
)

func applyErrataToBitmap(bitmap *matter.Bitmap, typeNames map[string]string, typeOverrides *errata.SDKTypes) {
	if typeOverrides != nil {
		override, ok := typeOverrides.Bitmaps[bitmap.Name]

		if ok {
			applyBitmapOverride(bitmap, override)
		}
	}
	bitmap.Name = applyTypeName(typeNames, bitmap.Name)
}

func applyBitmapOverride(bitmap *matter.Bitmap, override *errata.SDKType) {
	if override.OverrideName != "" {
		bitmap.Name = override.OverrideName
	}
	if override.OverrideType != "" {
		bitmap.Type = types.ParseDataType(override.OverrideType, types.DataTypeRankScalar)
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
	for _, f := range override.ExtraFields {
		var found bool
		for _, b := range bitmap.Bits {
			if b.Name() == f.Name {
				found = true
				break
			}
		}
		if !found {
			bit := f.Bit
			if bit == "" {
				bit = f.Value
			}
			var conf conformance.Set
			if f.Conformance != "" {
				conf = conformance.ParseConformance(f.Conformance)
			}
			if f.Code != "" || bitmap.Name == "Features" {
				code := f.Code
				if code == "" {
					code = f.Name
				}
				bitmap.AddBit(matter.NewFeature(nil, bit, f.Name, code, f.Description, conf))
			} else {
				bitmap.AddBit(matter.NewBitmapBit(nil, bitmap, bit, f.Name, f.Description, conf))
			}
		}
	}
}

