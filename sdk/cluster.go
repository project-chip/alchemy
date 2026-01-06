package sdk

import (
	"log/slog"
	"slices"

	"github.com/project-chip/alchemy/asciidoc/parse"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
)

func applyErrataToCluster(spec *spec.Specification, cluster *matter.Cluster, errata errata.SDK) {
	if errata.Types != nil {
		ac, ok := errata.Types.Clusters[cluster.Name]
		if ok {
			if ac.Domain != "" {
				cluster.Domain = ac.Domain
			}
			if ac.Description != "" {
				cluster.Description = ac.Description
			}
		}
		for _, a := range cluster.Attributes {
			ao, ok := errata.Types.Attributes[a.Name]
			if !ok {
				continue
			}
			applyErrataToField(a, ao)
		}
		if cluster.Features != nil {
			fc, ok := errata.Types.Bitmaps["Features"]
			if ok {
				for _, feature := range cluster.Features.Bits {
					for _, f := range fc.Fields {
						if feature.Name() == f.Name {
							if f.OverrideName != "" {
								feature.SetName(f.OverrideName)
							}
							break
						}
					}
				}
			}
		}
	}

	for _, bm := range cluster.Bitmaps {
		applyErrataToBitmap(bm, errata.TypeNames, errata.Types)
	}
	for _, en := range cluster.Enums {
		applyErrataToEnum(en, errata.TypeNames, errata.Types)
	}
	for _, s := range cluster.Structs {
		applyErrataToStruct(s, errata.TypeNames, errata.Types)
	}
	for _, cmd := range cluster.Commands {
		applyErrataToCommand(cmd, errata.TypeNames, errata.Types)
	}
	for _, ev := range cluster.Events {
		applyErrataToEvent(ev, errata.TypeNames, errata.Types)
	}

	sharedEntities := make(map[types.Entity]types.Entity)
	for bitmapName := range errata.SharedBitmaps {
		if cluster.ParentCluster == nil {
			slog.Warn("Errata: separate bitmap on cluster without parent", slog.String("bitmapName", bitmapName), slog.String("clusterName", cluster.Name))
			continue
		}
		cluster.Bitmaps = replaceSharedEntity(spec, cluster, cluster.ParentCluster, bitmapName, cluster.ParentCluster.Bitmaps, cluster.Bitmaps, sharedEntities)
	}
	for enumName := range errata.SharedEnums {
		if cluster.ParentCluster == nil {
			slog.Warn("Errata: separate enum on cluster without parent", slog.String("enumName", enumName), slog.String("clusterName", cluster.Name))
			continue
		}
		cluster.Enums = replaceSharedEntity(spec, cluster, cluster.ParentCluster, enumName, cluster.ParentCluster.Enums, cluster.Enums, sharedEntities)
	}
	for structName := range errata.SharedStructs {
		if cluster.ParentCluster == nil {
			slog.Warn("Errata: separate struct on cluster without parent", slog.String("structName", structName), slog.String("clusterName", cluster.Name))
			continue
		}
		cluster.Structs = replaceSharedEntity(spec, cluster, cluster.ParentCluster, structName, cluster.ParentCluster.Structs, cluster.Structs, sharedEntities)
	}
	for enumName := range errata.SeparateEnums {
		for _, en := range cluster.Enums {
			if en.Name == enumName {
				spec.ClusterRefs.Add(cluster, en)
			}
		}
	}
	if len(sharedEntities) > 0 {
		cluster.TraverseDataTypes(func(parent, entity types.Entity) parse.SearchShould {
			target, ok := sharedEntities[entity]
			if !ok {
				return parse.SearchShouldContinue
			}
			switch parent := parent.(type) {
			case *matter.Field:
				fieldType := parent.Type
				if fieldType == nil {
					break
				}
				if fieldType.IsArray() {
					fieldType = fieldType.EntryType
				}
				if fieldType == nil {
					break
				}
				fieldType.Entity = target

			}
			return parse.SearchShouldContinue
		})

	}

}

func replaceSharedEntity[T types.Entity](spec *spec.Specification, cluster *matter.Cluster, parentCluster *matter.Cluster, entityName string, parentList []T, targetList []T, sharedEntities map[types.Entity]types.Entity) []T {
	var target, parentEntity T
	var found bool
	targetIndex := -1
	for i, s := range targetList {
		if matter.EntityName(s) == entityName {
			target = s
			found = true
			targetIndex = i
			break
		}
	}
	if !found {
		slog.Warn("Errata: separate entity not found", slog.String("entityName", entityName), slog.String("clusterName", cluster.Name))
		return targetList
	}
	found = false
	for _, s := range parentList {
		if matter.EntityName(s) == entityName {
			parentEntity = s
			found = true
			break
		}
	}
	if !found {
		slog.Warn("Errata: separate entity not found in parent", slog.String("entityName", entityName), slog.String("clusterName", cluster.Name), slog.String("parentClusterName", parentCluster.Name))
		return targetList
	}
	targetList = slices.Delete(targetList, targetIndex, targetIndex+1)
	spec.ClusterRefs.Remove(cluster, target)
	spec.ClusterRefs.Add(cluster, parentEntity)
	sharedEntities[target] = parentEntity
	return targetList
}
