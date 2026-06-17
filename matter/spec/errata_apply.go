package spec

import (
	"github.com/project-chip/alchemy/matter"
)

func applySdkErrata(spec *Specification) {
	globalEnums := make(map[string]*matter.Enum)
	globalBitmaps := make(map[string]*matter.Bitmap)
	globalStructs := make(map[string]*matter.Struct)

	for entity := range spec.GlobalObjects {
		switch e := entity.(type) {
		case *matter.Enum:
			globalEnums[e.Name] = e
		case *matter.Bitmap:
			globalBitmaps[e.Name] = e
		case *matter.Struct:
			globalStructs[e.Name] = e
		}
	}

	for originalC := range spec.Clusters {
		doc, ok := spec.DocRefs[originalC]
		if !ok || spec.Errata == nil {
			continue
		}
		errata := spec.Errata.Get(doc.Path.Relative)
		if errata == nil || errata.SDK.Types == nil {
			continue
		}

		// Process ForceLocal / Keep
		for name, entry := range errata.SDK.Types.Enums {
			if entry.ForceLocal || entry.Keep {
				if en, ok := globalEnums[name]; ok {
					exists := false
					for _, existing := range originalC.Enums {
						if existing.Name == name {
							exists = true
							break
						}
					}
					if !exists {
						originalC.Enums = append(originalC.Enums, en)
					}
				}
			}
		}
		for name, entry := range errata.SDK.Types.Bitmaps {
			if entry.ForceLocal || entry.Keep {
				if bm, ok := globalBitmaps[name]; ok {
					exists := false
					for _, existing := range originalC.Bitmaps {
						if existing.Name == name {
							exists = true
							break
						}
					}
					if !exists {
						originalC.Bitmaps = append(originalC.Bitmaps, bm)
					}
				}
			}
		}
		for name, entry := range errata.SDK.Types.Structs {
			if entry.ForceLocal || entry.Keep {
				if st, ok := globalStructs[name]; ok {
					exists := false
					for _, existing := range originalC.Structs {
						if existing.Name == name {
							exists = true
							break
						}
					}
					if !exists {
						originalC.Structs = append(originalC.Structs, st)
					}
				}
			}
		}

		// Process ForceGlobal
		for name, entry := range errata.SDK.Types.Enums {
			if entry.ForceGlobal {
				for i, existing := range originalC.Enums {
					if existing.Name == name {
						originalC.Enums = append(originalC.Enums[:i], originalC.Enums[i+1:]...)
						spec.addEntityByName(name, existing, nil)
						spec.GlobalObjects[existing] = doc
						break
					}
				}
			}
		}
		for name, entry := range errata.SDK.Types.Bitmaps {
			if entry.ForceGlobal {
				for i, existing := range originalC.Bitmaps {
					if existing.Name == name {
						originalC.Bitmaps = append(originalC.Bitmaps[:i], originalC.Bitmaps[i+1:]...)
						spec.addEntityByName(name, existing, nil)
						spec.GlobalObjects[existing] = doc
						break
					}
				}
			}
		}
		for name, entry := range errata.SDK.Types.Structs {
			if entry.ForceGlobal {
				for i, existing := range originalC.Structs {
					if existing.Name == name {
						originalC.Structs = append(originalC.Structs[:i], originalC.Structs[i+1:]...)
						spec.addEntityByName(name, existing, nil)
						spec.GlobalObjects[existing] = doc
						break
					}
				}
			}
		}
	}
}
