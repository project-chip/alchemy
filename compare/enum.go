package compare

import (
	"log/slog"
	"strings"

	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
)

func compareEnum(specEnum *matter.Enum, zapEnum *matter.Enum) (diffs []Diff) {
	specEnumMap := make(map[uint64]*matter.EnumValue)
	for _, f := range specEnum.Values {
		if !f.Value.Valid() {
			slog.Warn("invalid spec enum value", slog.String("enum", specEnum.Name), slog.String("val", f.Name), slog.String("value", f.Value.Text()))
			continue
		}
		specEnumMap[f.Value.Value()] = f
	}

	zapEnumMap := make(map[uint64]*matter.EnumValue)
	for _, f := range zapEnum.Values {
		if !f.Value.Valid() {
			slog.Warn("invalid ZAP enum value", slog.String("enum", specEnum.Name), slog.String("val", f.Name), slog.String("value", f.Value.Text()))
			continue
		}

		zapEnumMap[f.Value.Value()] = f
	}

	for val := range zapEnumMap {
		_, ok := specEnumMap[val]
		if !ok {
			continue
		}
		delete(zapEnumMap, val)
		delete(specEnumMap, val)
	}
	for _, f := range specEnumMap {
		diffs = append(diffs, newMissingDiff(f.Name, types.EntityTypeEnumValue, SourceZAP))
	}
	for _, f := range zapEnumMap {
		diffs = append(diffs, newMissingDiff(f.Name, types.EntityTypeEnumValue, SourceSpec))
	}
	return
}

func compareEnums(spec *spec.Specification, specCluster *matter.Cluster, zapEnums []*matter.Enum) (diffs []Diff) {
	specEnumMap := make(map[string]*matter.Enum)
	for _, f := range specCluster.Enums {
		specEnumMap[strings.ToLower(f.Name)] = f
	}

	zapEnumMap := make(map[string]*matter.Enum)
	for _, f := range zapEnums {
		zapEnumMap[strings.ToLower(f.Name)] = f
	}
	for name, zapEnum := range zapEnumMap {
		specName := name
		specEnum, ok := specEnumMap[specName]
		if !ok {
			specName += "enum"
			specEnum, ok = specEnumMap[specName]
			if !ok {
				continue
			}
		}
		delete(zapEnumMap, name)
		delete(specEnumMap, specName)
		enumDiffs := compareEnum(specEnum, zapEnum)
		if len(enumDiffs) > 0 {
			diffs = append(diffs, &IdentifiedDiff{Entity: types.EntityTypeEnum, Name: specEnum.Name, Diffs: enumDiffs})
		}
	}
	for _, f := range specEnumMap {
		clusters, ok := spec.ClusterRefs[f]
		if ok {
			var externalReference bool
			for c := range clusters {
				slog.Info("checking enum reference", "name", f.Name, "clusterName", c.Name)
				if c == specCluster {
					externalReference = true
					break
				}
			}
			if externalReference {
				slog.Warn("Enum referred by different spec", "name", f.Name)
				continue
			}
		}
		diffs = append(diffs, newMissingDiff(f.Name, types.EntityTypeEnum, SourceZAP))
	}
	for _, f := range zapEnumMap {
		diffs = append(diffs, newMissingDiff(f.Name, types.EntityTypeEnum, SourceSpec))
	}
	return
}
