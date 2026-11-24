package sdk

import (
	"log/slog"
	"maps"
	"slices"

	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
)

// copyForceLocalTypes gives clusters their own copy of bitmaps, enums and structs that they reference from
// another cluster, when the referencing cluster's document marks them with force-local in its types errata.
// Global types marked with force-local are handled before data model resolution by applyErrataLocality.
//
// It must run after data type resolution, so field types point to their entities, and before renames, as
// the errata use spec names. The referencing document's types errata is applied to each copy at that point,
// so overrides that need resolving (e.g. an override-type naming a custom type) are not resolved on the copy;
// simple overrides such as constraints work as expected.
func copyForceLocalTypes(s *spec.Specification) (err error) {
	if s.Errata == nil {
		return
	}
	p := &patch{spec: s}
	// Copies can chain (a copied struct may use a type its owner copied), so process documents in a fixed order
	all := maps.Collect(s.Errata.All())
	for _, path := range slices.Sorted(maps.Keys(all)) {
		sdkErrata := &all[path].SDK
		if !hasForceLocalTypes(sdkErrata.Types) {
			continue
		}
		doc, ok := s.Docs[path]
		if !ok {
			slog.Warn("Errata refers to unknown document", "path", path)
			continue
		}
		var clusters []*matter.Cluster
		for _, entity := range s.EntitiesForDocument(doc) {
			switch entity := entity.(type) {
			case *matter.Cluster:
				clusters = append(clusters, entity)
			case *matter.ClusterGroup:
				clusters = append(clusters, entity.Clusters...)
			}
		}
		copied := make(map[types.Entity]bool)
		for _, cluster := range clusters {
			err = p.copyTypesToCluster(cluster, sdkErrata, copied)
			if err != nil {
				return
			}
		}
		warnUnmatchedForceLocalTypes(s, path, sdkErrata.Types, clusters, copied)
	}
	return
}

// copyTypesToCluster replaces references from fields owned by the cluster to force-local data types owned
// by another cluster with references to a copy owned by the cluster. Fields of other clusters' entities are
// never modified. Copied structs are scanned too, so force-local types nested inside them are also copied.
func (p *patch) copyTypesToCluster(cluster *matter.Cluster, sdkErrata *errata.SDK, copied map[types.Entity]bool) error {
	copies := make(map[types.Entity]types.Entity)
	for {
		var added bool
		for _, field := range clusterOwnedFields(cluster) {
			dataType := field.Type
			if dataType != nil && dataType.IsArray() {
				dataType = dataType.EntryType
			}
			if dataType == nil || dataType.Entity == nil || !isForceLocal(sdkErrata.Types, dataType.Entity) {
				continue
			}
			owner, ok := dataType.Entity.Parent().(*matter.Cluster)
			if !ok || owner == cluster {
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
				copied[clone] = true
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

// forceLocalEntries returns the types errata entries for the given entity's kind
func forceLocalEntries(typeOverrides *errata.SDKTypes, entity types.Entity) map[string]*errata.SDKType {
	if typeOverrides == nil {
		return nil
	}
	switch entity.(type) {
	case *matter.Bitmap:
		return typeOverrides.Bitmaps
	case *matter.Enum:
		return typeOverrides.Enums
	case *matter.Struct:
		return typeOverrides.Structs
	}
	return nil
}

func isForceLocal(typeOverrides *errata.SDKTypes, entity types.Entity) bool {
	entry, ok := forceLocalEntries(typeOverrides, entity)[matter.EntityName(entity)]
	return ok && entry != nil && entry.ForceLocal
}

func hasForceLocalTypes(typeOverrides *errata.SDKTypes) bool {
	if typeOverrides == nil {
		return false
	}
	for _, entries := range []map[string]*errata.SDKType{typeOverrides.Bitmaps, typeOverrides.Enums, typeOverrides.Structs} {
		for _, entry := range entries {
			if entry != nil && entry.ForceLocal {
				return true
			}
		}
	}
	return false
}

// warnUnmatchedForceLocalTypes warns about force-local entries that match nothing: not copied from another
// cluster, not a global type (handled by applyErrataLocality), and not already owned by the document's clusters
func warnUnmatchedForceLocalTypes(s *spec.Specification, path string, typeOverrides *errata.SDKTypes, clusters []*matter.Cluster, copied map[types.Entity]bool) {
	known := make(map[string]bool)
	add := func(e types.Entity) {
		known[matter.EntityName(e)] = true
	}
	for e := range copied {
		add(e)
	}
	for e := range s.GlobalObjects {
		add(e)
	}
	for _, c := range clusters {
		for _, e := range c.Bitmaps {
			add(e)
		}
		for _, e := range c.Enums {
			add(e)
		}
		for _, e := range c.Structs {
			add(e)
		}
	}
	for _, entries := range []map[string]*errata.SDKType{typeOverrides.Bitmaps, typeOverrides.Enums, typeOverrides.Structs} {
		for name, entry := range entries {
			if entry != nil && entry.ForceLocal && !known[name] {
				slog.Warn("Errata: force-local data type is not referenced from another cluster, global, or local", slog.String("path", path), slog.String("name", name))
			}
		}
	}
}
