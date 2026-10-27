package sdk

import (
	"fmt"
	"log/slog"
	"slices"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/asciidoc/parse"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
)

func (p *patch) applyErrataToCluster(spec *spec.Specification, cluster *matter.Cluster, errata errata.SDK, options applyErrataOptions) (err error) {
	if errata.Types != nil {
		ac, ok := errata.Types.Clusters[cluster.Name]
		if ok {
			if ac.Domain != "" {
				cluster.Domain = ac.Domain
			}
			if ac.Description != "" {
				cluster.Description = ac.Description
			}
			if ac.Conformance != "" {
				cluster.Conformance = conformance.ParseConformance(ac.Conformance)
			}
		}
		for _, a := range cluster.Attributes {
			ao, ok := errata.Types.Attributes[a.Name]
			if !ok {
				continue
			}
			err := p.applyErrataToField(a, ao)
			if err != nil {
				return err
			}
		}
		if cluster.Features != nil {
			fc, ok := errata.Types.Bitmaps["Features"]
			if ok {
				p.applyBitmapOverride(&cluster.Features.Bitmap, fc)
			}
		}
	}

	doc, ok := spec.DocRefs[cluster]
	if !ok {
		return fmt.Errorf("Unable to apply errata for cluster %s; unable to find cluster in spec", cluster.Name)
	}

	for _, bm := range cluster.Bitmaps {
		p.applyErrataToBitmap(bm, errata.TypeNames, errata.Types)
	}
	if errata.Types != nil && errata.Types.Bitmaps != nil {
		cluster.Bitmaps = p.applyErrataLocality(doc, errata.Types.Bitmaps, p.globalBitmaps, cluster.Bitmaps)
	}

	for _, en := range cluster.Enums {
		p.applyErrataToEnum(en, errata.TypeNames, errata.Types)
	}
	if errata.Types != nil && errata.Types.Enums != nil {
		cluster.Enums = p.applyErrataLocality(doc, errata.Types.Enums, p.globalEnums, cluster.Enums)
	}

	for _, s := range cluster.Structs {
		err := p.applyErrataToStruct(s, errata.TypeNames, errata.Types)
		if err != nil {
			return err
		}
	}
	if errata.Types != nil && errata.Types.Structs != nil {
		cluster.Structs = p.applyErrataLocality(doc, errata.Types.Structs, p.globalStructs, cluster.Structs)
	}

	for _, cmd := range cluster.Commands {
		err := p.applyErrataToCommand(cmd, errata.TypeNames, errata.Types)
		if err != nil {
			return err
		}
	}
	for _, ev := range cluster.Events {
		err := p.applyErrataToEvent(ev, errata.TypeNames, errata.Types)
		if err != nil {
			return err
		}
	}

	if !options.skipSharedEntities {
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
	} else {
		if cluster.ParentCluster != nil {
			doc, ok := spec.DocRefs[cluster.ParentCluster]
			if ok {
				parentErrata := spec.Errata.Get(doc.Path.Relative)
				if parentErrata != nil {
					for bitmapName := range errata.SharedBitmaps {
						for _, bm := range cluster.Bitmaps {
							if matter.EntityName(bm) == bitmapName {
								p.applyErrataToBitmap(bm, parentErrata.SDK.TypeNames, parentErrata.SDK.Types)
								break
							}
						}
					}
					for enumName := range errata.SharedEnums {
						for _, en := range cluster.Enums {
							if matter.EntityName(en) == enumName {
								p.applyErrataToEnum(en, parentErrata.SDK.TypeNames, parentErrata.SDK.Types)
								break
							}
						}
					}
					for structName := range errata.SharedStructs {
						for _, st := range cluster.Structs {
							if matter.EntityName(st) == structName {
								err = p.applyErrataToStruct(st, parentErrata.SDK.TypeNames, parentErrata.SDK.Types)
								if err != nil {
									return
								}
								break
							}
						}
					}
				}
			}
		}
	}
	for enumName := range errata.SeparateEnums {
		for _, en := range cluster.Enums {
			if en.Name == enumName {
				spec.ClusterRefs.Add(cluster, en)
			}
		}
	}

	return
}

func (p *patch) applyErrataLocality[T types.Entity](doc *asciidoc.Document, types map[string]*errata.SDKType, globals map[string]T, list []T) []T {
	for name, entry := range types {
		var local T
		var localIndex int
		global := globals[name]
		for i, e := range list {
			if matter.EntityName(e) == name {
				local = e
				localIndex = i
				break
			}
		}
		if entry.ForceLocal || entry.Keep {
			if any(local) == nil && any(global) != nil {
				list = append(list, global)
			}
		} else if entry.ForceGlobal {
			if any(global) == nil && any(local) != nil {
				list = slices.Delete(list, localIndex, localIndex+1)
				p.spec.GlobalObjects[local] = doc
			}
		}
	}
	return list
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
