package compare

import (
	"strings"

	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/types"
)

func compareStruct(specStruct *matter.Struct, zapStruct *matter.Struct) (diffs []any) {
	switch zapStruct.FabricScoping {
	case matter.FabricScopingScoped:
		if specStruct.FabricScoping != matter.FabricScopingScoped {
			diffs = append(diffs, &StringDiff{Type: DiffTypeMismatch, Property: DiffPropertyFabricScoping, Spec: specStruct.FabricScoping.String(), ZAP: zapStruct.FabricScoping.String()})
		}
	case matter.FabricScopingUnscoped:
		if specStruct.FabricScoping == matter.FabricScopingScoped {
			diffs = append(diffs, &StringDiff{Type: DiffTypeMismatch, Property: DiffPropertyFabricScoping, Spec: specStruct.FabricScoping.String(), ZAP: zapStruct.FabricScoping.String()})
		}
	}

	fieldDiffs, err := compareFields(specStruct.Fields, zapStruct.Fields)
	if err == nil && len(fieldDiffs) > 0 {
		diffs = append(diffs, fieldDiffs)
	}
	return
}

func compareStructs(specStructs []*matter.Struct, zapStructs []*matter.Struct) (diffs []any) {
	specStructMap := make(map[string]*matter.Struct)
	for _, f := range specStructs {
		specStructMap[strings.ToLower(f.Name)] = f
	}

	zapStructMap := make(map[string]*matter.Struct)
	for _, f := range zapStructs {
		zapStructMap[strings.ToLower(f.Name)] = f
	}
	for name, zapStruct := range zapStructMap {
		specName := name
		specStruct, ok := specStructMap[specName]
		if !ok {
			specName += "command"
			specStruct, ok = specStructMap[specName]
			if !ok {
				continue
			}
		}
		delete(zapStructMap, name)
		delete(specStructMap, specName)
		structDiffs := compareStruct(specStruct, zapStruct)
		if len(structDiffs) > 0 {
			diffs = append(diffs, &IdentifiedDiff{Type: DiffTypeMismatch, Entity: types.EntityTypeStruct, Name: specStruct.Name, Diffs: structDiffs})
		}
	}
	for _, f := range specStructMap {
		diffs = append(diffs, newMissingDiff(f.Name, types.EntityTypeStruct, SourceZAP))
	}
	for _, f := range zapStructMap {
		diffs = append(diffs, newMissingDiff(f.Name, types.EntityTypeStruct, SourceSpec))
	}
	return
}
