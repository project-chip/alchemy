package regen

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

func entityShouldBeIncluded(e types.Entity) bool {
	conf := matter.EntityConformance(e)
	if conformance.IsZigbee(conf) || zap.IsDisallowed(e, conf) || conformance.IsDeprecated(conf) {
		return false
	}
	return true
}

func filterEntities[T types.Entity](sets ...[]T) (set []T) {
	for _, s := range sets {
		for _, e := range s {
			if !entityShouldBeIncluded(e) {
				continue
			}
			set = append(set, e)
		}
	}
	return
}

func enumerateEntitiesHelper[T types.Entity](list []T, spec *spec.Specification, options *raymond.Options) raymond.SafeString {
	var result strings.Builder
	for i, en := range filterEntities(list) {
		df := options.DataFrame().Copy()
		df.Set("index", i)
		df.Set("key", nil)
		df.Set("first", i == 0)
		df.Set("last", i == len(list)-1)
		if spec != nil {
			refs, ok := spec.ClusterRefs.Get(en)
			if ok && refs.Size() > 1 {
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

func isProvisional(spec *spec.Specification, entity types.Entity) bool {
	switch entity := entity.(type) {
	case *matter.Bitmap:
		if entity.Name == "Feature" {
			return false
		}
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
