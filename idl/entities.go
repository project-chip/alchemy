package idl

import (
	"strings"

	"github.com/mailgun/raymond/v2"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
	"github.com/project-chip/alchemy/provisional"
	"github.com/project-chip/alchemy/zap"
)

type ProvisionalFilter struct {
	Mode string
}

func entityShouldBeIncluded(spec *spec.Specification, filter ProvisionalFilter, e types.Entity) bool {
	conf := matter.EntityConformance(e)
	if conformance.IsZigbee(conf) || zap.IsDisallowed(e, conf) || conformance.IsDeprecated(conf) {
		return false
	}

	if isKeptByErrata(spec, e) {
		return true
	}

	if filter.Mode != "none" && isProvisional(spec, e) {
		return false
	}
	return true
}

func filterEntities[T types.Entity](spec *spec.Specification, filter ProvisionalFilter, sets ...[]T) (set []T) {
	for _, s := range sets {
		for _, e := range s {
			if !entityShouldBeIncluded(spec, filter, e) {
				continue
			}
			set = append(set, e)
		}
	}
	return
}

func enumerateEntitiesHelper[T types.Entity](list []T, spec *spec.Specification, filter ProvisionalFilter, options *raymond.Options) raymond.SafeString {
	var result strings.Builder
	for i, en := range filterEntities(spec, filter, list) {
		df := options.DataFrame().Copy()
		df.Set("index", i)
		df.Set("key", nil)
		df.Set("first", i == 0)
		df.Set("last", i == len(list)-1)
		if spec != nil {
			if isShared(spec, en) {
				df.Set("shared", true)
			}
			df.Set("provisional", isProvisional(spec, en))
			dr, ok := spec.DataTypeRefs.Get(en)
			if ok {
				df.Set("refCount", dr.Size())
			}
		}
		result.WriteString(options.FnCtxData(en, df))
	}
	return raymond.SafeString(result.String())
}

func isShared(spec *spec.Specification, en types.Entity) bool {
	refs, ok := spec.ClusterRefs.Get(en)
	if ok && refs.Size() > 1 {
		return true
	}
	var cluster *matter.Cluster
	var name string
	var isTargetType bool
	switch entity := en.(type) {
	case *matter.Struct:
		cluster = entity.Cluster()
		name = entity.Name
		isTargetType = true
	case *matter.Enum:
		cluster = entity.Cluster()
		name = entity.Name
		isTargetType = true
	case *matter.Bitmap:
		cluster = entity.Cluster()
		name = entity.Name
		isTargetType = true
	}
	if !isTargetType || cluster == nil {
		return false
	}
	doc, ok := spec.DocRefs[cluster]
	if !ok {
		return false
	}
	if spec.Errata == nil {
		return false
	}
	errata := spec.Errata.Get(doc.Path.Relative)
	if errata == nil {
		return false
	}
	switch en.(type) {
	case *matter.Struct:
		_, ok = errata.SDK.SharedStructs[name]
		return ok
	case *matter.Enum:
		_, ok = errata.SDK.SharedEnums[name]
		return ok
	case *matter.Bitmap:
		_, ok = errata.SDK.SharedBitmaps[name]
		return ok
	}
	return false
}

func isKeptByErrata(spec *spec.Specification, entity types.Entity) bool {
	if spec != nil && spec.Errata != nil {
		var parentCluster *matter.Cluster
		curr := entity
		for curr != nil {
			if c, ok := curr.(*matter.Cluster); ok {
				parentCluster = c
				break
			}
			curr = curr.Parent()
		}
		for parentCluster != nil {
			if doc, ok := spec.DocRefs[parentCluster]; ok {
				if errata := spec.Errata.Get(doc.Path.Relative); errata != nil && errata.SDK.Types != nil {
					name := matter.EntityName(entity)
					switch entity.(type) {
					case *matter.Field:
						if entry, ok := errata.SDK.Types.Attributes[name]; ok && entry.Keep {
							return true
						}
					case *matter.Command:
						if entry, ok := errata.SDK.Types.Commands[name]; ok && entry.Keep {
							return true
						}
					case *matter.Event:
						if entry, ok := errata.SDK.Types.Events[name]; ok && entry.Keep {
							return true
						}
					case *matter.Enum:
						if entry, ok := errata.SDK.Types.Enums[name]; ok && entry.Keep {
							return true
						}
					case *matter.Bitmap:
						if entry, ok := errata.SDK.Types.Bitmaps[name]; ok && entry.Keep {
							return true
						}
					case *matter.Struct:
						if entry, ok := errata.SDK.Types.Structs[name]; ok && entry.Keep {
							return true
						}
					}
				}
			}
			parentCluster = parentCluster.ParentCluster
		}
	}
	return false
}

func isProvisional(spec *spec.Specification, entity types.Entity) bool {
	switch entity := entity.(type) {
	case *matter.Bitmap:
		if entity.Name == "Feature" || entity.Name == "Features" {
			return false
		}
	case *matter.Enum:
		if strings.HasSuffix(entity.Name, "Tag") {
			// This is a hacky workaround for namespaces being represented as enumerations in .matter files, but not being provisional
			return false
		}
	case *matter.Field:
		if entity.Name == "FabricIndex" {
			return false
		}
		return conformance.IsProvisional(matter.EntityConformance(entity))
	case nil:
		return false
	}
	is := provisional.Check(spec, entity, entity)

	switch is {
	case provisional.StateAllClustersProvisional,
		provisional.StateAllDataTypeReferencesProvisional,
		provisional.StateExplicitlyProvisional,
		provisional.StateUnreferenced:
		return true
	default:
		return false
	}
}

func entityPath(e types.Entity) string {
	var parts []string
	curr := e
	for curr != nil {
		switch node := curr.(type) {
		case *matter.Cluster:
			parts = append(parts, caseify(node.Name, false, false))
			curr = nil // Stop at cluster level
		case *matter.Enum:
			parts = append(parts, node.Name)
			curr = node.Parent()
		case *matter.Bitmap:
			if node.Name == "Features" {
				parts = append(parts, "Feature")
			} else {
				parts = append(parts, node.Name)
			}
			curr = node.Parent()
		case *matter.Features:
			parts = append(parts, "Feature")
			curr = node.Parent()
		case *matter.Struct:
			parts = append(parts, node.Name)
			curr = node.Parent()
		case *matter.Event:
			parts = append(parts, node.Name, "event")
			curr = node.Parent()
		case *matter.Command:
			parts = append(parts, node.Name, "command")
			curr = node.Parent()
		case *matter.Field:
			if _, isCluster := node.Parent().(*matter.Cluster); isCluster {
				parts = append(parts, node.Name, "attribute")
			} else {
				parts = append(parts, node.Name)
			}
			curr = node.Parent()
		case *matter.EnumValue:
			parts = append(parts, node.Name)
			curr = node.Parent()
		case *matter.BitmapBit:
			parts = append(parts, node.Name())
			curr = node.Parent()
		case *matter.Feature:
			parts = append(parts, node.Name())
			curr = node.Parent()
		default:
			curr = curr.Parent()
		}
	}
	for i, j := 0, len(parts)-1; i < j; i, j = i+1, j-1 {
		parts[i], parts[j] = parts[j], parts[i]
	}
	return strings.Join(parts, ".")
}
