package spec

import (
	"log/slog"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/config"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/internal/suggest"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

type Specification struct {
	Root   string
	Config *config.Config
	Errata *errata.Collection

	Clusters       map[*matter.Cluster]struct{}
	ClustersByID   map[uint64]*matter.Cluster
	ClustersByName map[string]*matter.Cluster

	DeviceTypes       []*matter.DeviceType
	DeviceTypesByID   map[uint64]*matter.DeviceType
	DeviceTypesByName map[string]*matter.DeviceType

	BaseDeviceType     *matter.DeviceType
	RootNodeDeviceType *matter.DeviceType

	deviceTypeCompositionCache map[*matter.DeviceType]*matter.DeviceTypeComposition

	Namespaces []*matter.Namespace

	ClusterRefs  EntityRefs[*matter.Cluster]
	DataTypeRefs EntityRefs[types.Entity]
	DocRefs      map[types.Entity]*asciidoc.Document
	LibraryRefs  map[types.Entity]*Library
	entityRefs   map[*asciidoc.Document][]types.Entity

	GlobalObjects types.EntitySet[*asciidoc.Document]

	entitiesByID map[string]map[types.Entity]map[*matter.Cluster]struct{}

	Docs       map[string]*asciidoc.Document
	UnusedDocs []*asciidoc.Document
	Libraries  []*Library

	Errors []Error

	libraryIndex map[*asciidoc.Document]*Library
}

func newSpec(specRoot string, config *config.Config, errata *errata.Collection) *Specification {
	return &Specification{
		Root:   specRoot,
		Config: config,
		Errata: errata,

		Clusters:          make(map[*matter.Cluster]struct{}),
		ClustersByID:      make(map[uint64]*matter.Cluster),
		ClustersByName:    make(map[string]*matter.Cluster),
		ClusterRefs:       NewEntityRefs[*matter.Cluster](),
		DataTypeRefs:      NewEntityRefs[types.Entity](),
		DeviceTypesByID:   make(map[uint64]*matter.DeviceType),
		DeviceTypesByName: make(map[string]*matter.DeviceType),
		Docs:              make(map[string]*asciidoc.Document),
		DocRefs:           make(map[types.Entity]*asciidoc.Document),
		LibraryRefs:       make(map[types.Entity]*Library),
		entityRefs:        make(map[*asciidoc.Document][]types.Entity),

		GlobalObjects: make(types.EntitySet[*asciidoc.Document]),

		entitiesByID:               make(map[string]map[types.Entity]map[*matter.Cluster]struct{}),
		deviceTypeCompositionCache: make(map[*matter.DeviceType]*matter.DeviceTypeComposition),
		libraryIndex:               make(map[*asciidoc.Document]*Library),
	}
}

func (s *Specification) LibraryForDocument(doc *asciidoc.Document) (library *Library, ok bool) {
	library, ok = s.libraryIndex[doc]
	return
}

func (s *Specification) EntitiesForDocument(doc *asciidoc.Document) []types.Entity {
	return s.entityRefs[doc]
}

type specEntityFinder struct {
	entityFinderCommon

	spec    *Specification
	cluster *matter.Cluster
}

func newSpecEntityFinder(spec *Specification, cluster *matter.Cluster, inner entityFinder) *specEntityFinder {
	return &specEntityFinder{entityFinderCommon: entityFinderCommon{inner: inner}, spec: spec, cluster: cluster}
}

func (sef *specEntityFinder) findEntityByIdentifier(identifier string, source log.Source) types.Entity {
	entities := sef.spec.entitiesByID[identifier]
	if len(entities) == 0 {
		canonicalName := CanonicalName(identifier)
		if canonicalName != identifier {
			return sef.findEntityByIdentifier(canonicalName, source)
		}
		// These are special-cased, as they are references to attributes from a different rendered document
		// Once we combine clusters into a single document, this can be removed.
		switch identifier {
		case "SupportedFabrics":
			if opCredentialsCluster, ok := sef.spec.ClustersByID[0x003E]; ok {
				return opCredentialsCluster.Attributes.Get(identifier)
			}
		case "SubjectsPerAccessControlEntry", "TargetsPerAccessControlEntry":
			if aclCluster, ok := sef.spec.ClustersByID[0x001F]; ok {
				return aclCluster.Attributes.Get(identifier)
			}
		}
	} else if len(entities) == 1 {
		for m := range entities {
			return m
		}
	} else {
		return disambiguateDataType(entities, sef.cluster, identifier, source)
	}
	return nil
}

func (tf *specEntityFinder) suggestIdentifiers(identifier string, suggestions map[types.Entity]int) {
	suggest.PossibleEntities(identifier, suggestions, func(yield func(string, types.Entity) bool) {
		for identifier, entities := range tf.spec.entitiesByID {
			for e := range entities {
				if e == tf.identity {
					continue
				}
				if !yield(identifier, e) {
					return
				}
			}
		}
	})
	if tf.inner != nil {
		tf.inner.suggestIdentifiers(identifier, suggestions)
	}
}

func (sef *specEntityFinder) findEntityByReference(reference string, label string, source log.Source) (e types.Entity) {
	e = sef.findSpecEntityByReference(reference, label, source)
	if e != nil {
		return
	}
	if sef.inner != nil {
		e = sef.inner.findEntityByReference(reference, label, source)
	}
	return
}

func (sef *specEntityFinder) findSpecEntityByReference(reference string, label string, source log.Source) (e types.Entity) {
	var referenceDoc *asciidoc.Document
	if sef.cluster != nil {
		referenceDoc = sef.spec.DocRefs[sef.cluster]
	}
	if referenceDoc == nil {
		path, _ := source.Origin()
		specPath, err := NewSpecPath(path, sef.spec.Root)
		if err != nil {
			slog.Warn("failed to create spec path for reference", "ref", reference, log.Path("source", source), slog.Any("error", err))
		} else {
			ld := sef.spec.Docs[specPath.Relative]
			if ld != nil {
				referenceDoc = ld
			}
		}
	}
	if referenceDoc == nil {
		slog.Warn("failed to find document for reference", "ref", reference, log.Path("source", source), slog.Any("cluster", sef.cluster))
		return
	}
	var anchors []*Anchor

	library, ok := sef.spec.libraryIndex[referenceDoc]
	if !ok {
		slog.Warn("failed to find library for reference", "ref", reference, log.Path("source", source), slog.Any("cluster", sef.cluster))
		return
	}

	anchors = library.anchors[reference]
	if len(anchors) == 0 {
		slog.Warn("failed to find anchors for data type reference", "ref", reference, log.Path("source", source), slog.String("cluster", clusterName(sef.cluster)), slog.String("docPath", referenceDoc.Path.Relative))
		return
	}

	var discoveredEntities []types.Entity
	for _, anchor := range anchors {
		switch el := anchor.Element.(type) {
		case *asciidoc.Section:

			entities := library.entitiesByElement[el]
			discoveredEntities = append(discoveredEntities, entities...)
		default:
			slog.Warn("unexpected type of anchor element", log.Type("type", el))
		}
	}
	switch len(discoveredEntities) {
	case 0:
		slog.Warn("no entities found for reference", "ref", reference, log.Path("source", source))
		for _, anchor := range anchors {
			switch el := anchor.Element.(type) {
			case *asciidoc.Section:
				slog.Warn("anchor element", log.Path("type", el), log.Address("address", el))
			}
		}
	case 1:
		e = discoveredEntities[0]
	default:
		slog.Warn("ambiguous reference", "ref", reference, log.Path("source", source))
		for _, m := range discoveredEntities {
			slog.Warn("ambiguous reference", matter.LogEntity("entity", m), log.Path("source", source))
		}
	}
	if e != nil && label != "" {

		switch entity := e.(type) {
		case *matter.Enum:
			if strings.EqualFold(entity.Name, label) {
				return
			}
			for _, ev := range entity.Values {
				if strings.EqualFold(label, ev.Name) {
					e = ev
					return
				}
			}
		case *matter.Bitmap:
			if strings.EqualFold(entity.Name, label) {
				return
			}
			for _, ev := range entity.Bits {
				if strings.EqualFold(label, ev.Name()) {
					e = ev
					return
				}
			}
		case *matter.Struct:
			if strings.EqualFold(entity.Name, label) {
				return
			}
			for _, f := range entity.Fields {
				if strings.EqualFold(label, f.Name) {
					e = f
					return
				}
			}
		case *matter.Command:
			if strings.EqualFold(entity.Name, label) {
				return
			}
			for _, f := range entity.Fields {
				if strings.EqualFold(label, f.Name) {
					e = f
					return
				}
			}
		case *matter.Field:
			if strings.EqualFold(entity.Name, label) {
				return
			}
			slog.Warn("Unhandled reference field with label", slog.String("clusterName", sef.cluster.Name), slog.String("field", entity.Name), slog.String("label", label), matter.LogEntity("entity", e), log.Path("source", source))
		case *matter.Constant:
			if strings.EqualFold(entity.Name, label) {
				return
			}
			slog.Warn("Unhandled reference constant with label", slog.String("clusterName", sef.cluster.Name), slog.String("constant", entity.Name), slog.String("label", label), matter.LogEntity("entity", e), log.Path("source", source))
		default:
			slog.Warn("Unhandled reference type with label", slog.String("clusterName", sef.cluster.Name), log.Type("entityType", e), slog.String("label", label), matter.LogEntity("entity", e), log.Path("source", source))
		}
	}
	return
}

func disambiguateDataType(entities map[types.Entity]map[*matter.Cluster]struct{}, cluster *matter.Cluster, identifier string, source log.Source) types.Entity {
	// If there are multiple entities with the same name, prefer the one on the current cluster
	for m, clusters := range entities {
		for c := range clusters {
			if c == cluster {
				return m
			}
		}
	}

	// OK, if the data type is defined on the direct parent of this cluster, take that one
	if cluster != nil && cluster.ParentCluster != nil {
		for m, clusters := range entities {
			for c := range clusters {
				if c != nil && c == cluster.ParentCluster {
					return m
				}
			}
		}
	}

	var nakedEntities []types.Entity
	for m, clusters := range entities {
		for c := range clusters {
			if c == nil {
				nakedEntities = append(nakedEntities, m)
			}
		}
	}
	if len(nakedEntities) == 1 {
		return nakedEntities[0]
	}

	// Can't disambiguate out this data model
	slog.Warn("ambiguous data type", "cluster", clusterName(cluster), "identifier", identifier, log.Path("source", source))
	for m, clusters := range entities {
		for c := range clusters {
			var clusterName string
			if c != nil {
				clusterName = c.Name
			} else {
				clusterName = "naked"
			}
			slog.Warn("ambiguous data type candidate", matter.LogEntity("source", m), "cluster", clusterName)
		}
	}
	return nil
}
