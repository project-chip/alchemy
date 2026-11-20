package sdk

import (
	"log/slog"
	"slices"

	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
)

// copyErrataTypes gives clusters their own copy of bitmaps, enums and structs that they reference from
// another cluster, as listed in the copy-bitmaps, copy-enums and copy-structs errata of the referencing
// cluster's document. It must run after data type resolution, so field types point to their entities,
// and before renames, as the errata lists use spec names.
func copyErrataTypes(s *spec.Specification) (err error) {
	if s.Errata == nil {
		return
	}
	p := &patch{
		spec:          s,
		globalEnums:   make(map[string]*matter.Enum),
		globalBitmaps: make(map[string]*matter.Bitmap),
		globalStructs: make(map[string]*matter.Struct),
	}
	for path, e := range s.Errata.All() {
		if !e.SDK.HasCopies() {
			continue
		}
		doc, ok := s.Docs[path]
		if !ok {
			slog.Warn("Errata refers to unknown document", "path", path)
			continue
		}
		copied := make(map[string]bool)
		for _, entity := range s.EntitiesForDocument(doc) {
			switch entity := entity.(type) {
			case *matter.Cluster:
				err = p.copyTypesToCluster(entity, &e.SDK, copied)
			case *matter.ClusterGroup:
				for _, c := range entity.Clusters {
					err = p.copyTypesToCluster(c, &e.SDK, copied)
					if err != nil {
						return
					}
				}
			}
			if err != nil {
				return
			}
		}
		warnUncopiedTypes(path, "copy-bitmaps", e.SDK.CopyBitmaps, copied)
		warnUncopiedTypes(path, "copy-enums", e.SDK.CopyEnums, copied)
		warnUncopiedTypes(path, "copy-structs", e.SDK.CopyStructs, copied)
	}
	return
}

// copyTypesToCluster replaces references from fields owned by the cluster to listed data types owned by
// another cluster with references to a copy owned by the cluster. Fields of other clusters' entities are
// never modified. Copied structs are scanned too, so listed types nested inside them are also copied.
func (p *patch) copyTypesToCluster(cluster *matter.Cluster, sdkErrata *errata.SDK, copied map[string]bool) error {
	copies := make(map[types.Entity]types.Entity)
	for {
		var added bool
		for _, field := range clusterOwnedFields(cluster) {
			dataType := field.Type
			if dataType != nil && dataType.IsArray() {
				dataType = dataType.EntryType
			}
			if dataType == nil || dataType.Entity == nil || !isListedForCopy(sdkErrata, dataType.Entity) {
				continue
			}
			owner, ok := dataType.Entity.Parent().(*matter.Cluster)
			if !ok {
				slog.Warn("Errata: unable to copy data type not owned by a cluster", matter.LogEntity("entity", dataType.Entity), slog.String("clusterName", cluster.Name))
				continue
			}
			if owner == cluster {
				continue
			}
			clone, ok := copies[dataType.Entity]
			if !ok {
				var err error
				clone, err = p.cloneIntoCluster(cluster, dataType.Entity, sdkErrata)
				if err != nil {
					return err
				}
				copies[dataType.Entity] = clone
				copied[matter.EntityName(dataType.Entity)] = true
				added = true
			}
			dataType.Entity = clone
		}
		if !added {
			return nil
		}
	}
}

func (p *patch) cloneIntoCluster(cluster *matter.Cluster, original types.Entity, sdkErrata *errata.SDK) (clone types.Entity, err error) {
	switch original := original.(type) {
	case *matter.Bitmap:
		bm := original.CloneTo(cluster)
		cluster.Bitmaps = append(cluster.Bitmaps, bm)
		p.applyErrataToBitmap(bm, sdkErrata.Types)
		clone = bm
	case *matter.Enum:
		en := original.CloneTo(cluster)
		cluster.Enums = append(cluster.Enums, en)
		p.applyErrataToEnum(en, sdkErrata.Types)
		clone = en
	case *matter.Struct:
		st := original.CloneTo(cluster)
		cluster.Structs = append(cluster.Structs, st)
		err = p.applyErrataToStruct(st, sdkErrata.Types)
		clone = st
	}
	p.noteCopyRefs(cluster, clone)
	return
}

// noteCopyRefs records the copy and its children as belonging to the cluster's document and library,
// so that later steps (renames, rendering) treat it like any other entity of the cluster
func (p *patch) noteCopyRefs(cluster *matter.Cluster, clone types.Entity) {
	doc, hasDoc := p.spec.DocRefs[cluster]
	library, hasLibrary := p.spec.LibraryRefs[cluster]
	note := func(e types.Entity) {
		if hasDoc {
			p.spec.DocRefs[e] = doc
		}
		if hasLibrary {
			p.spec.LibraryRefs[e] = library
		}
	}
	note(clone)
	switch clone := clone.(type) {
	case *matter.Bitmap:
		for _, b := range clone.Bits {
			note(b)
		}
	case *matter.Enum:
		for _, v := range clone.Values {
			note(v)
		}
	case *matter.Struct:
		for _, f := range clone.Fields {
			note(f)
		}
	}
}

func clusterOwnedFields(cluster *matter.Cluster) (fields []*matter.Field) {
	fields = slices.Clone(cluster.Attributes)
	for _, cmd := range cluster.Commands {
		fields = append(fields, cmd.Fields...)
	}
	for _, ev := range cluster.Events {
		fields = append(fields, ev.Fields...)
	}
	for _, s := range cluster.Structs {
		fields = append(fields, s.Fields...)
	}
	return
}

func isListedForCopy(sdkErrata *errata.SDK, entity types.Entity) bool {
	var list errata.UniqueStringList
	switch entity.(type) {
	case *matter.Bitmap:
		list = sdkErrata.CopyBitmaps
	case *matter.Enum:
		list = sdkErrata.CopyEnums
	case *matter.Struct:
		list = sdkErrata.CopyStructs
	default:
		return false
	}
	_, ok := list[matter.EntityName(entity)]
	return ok
}

func warnUncopiedTypes(path string, key string, list errata.UniqueStringList, copied map[string]bool) {
	for name := range list {
		if !copied[name] {
			slog.Warn("Errata: data type listed for copying is not referenced from another cluster", slog.String("path", path), slog.String("errata", key), slog.String("name", name))
		}
	}
}
