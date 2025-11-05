package regen

import (
	"strings"

	"github.com/mailgun/raymond/v2"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
	"github.com/project-chip/alchemy/provisional"
)

func enumerateEntitiesHelper[T types.Entity](list []T, spec *spec.Specification, options *raymond.Options) raymond.SafeString {
	var result strings.Builder
	for i, en := range list {
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
