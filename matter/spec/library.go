package spec

import (
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/asciidoc/parse"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

type Library struct {
	Root *asciidoc.Document
	Spec *Specification

	cache *DocCache

	referenceIndex
	asciidoc.Reader

	Docs []*asciidoc.Document

	index map[string]*asciidoc.Document

	crossReferenceDocs map[*asciidoc.CrossReference]*asciidoc.Document

	parents  map[*asciidoc.Document][]*asciidoc.Document
	children map[*asciidoc.Document][]*asciidoc.Document

	entities          []types.Entity
	orderedEntities   []types.Entity
	entitiesByElement map[asciidoc.Attributable][]types.Entity

	strings map[asciidoc.Parent]string

	sectionNames  pipeline.Map[*asciidoc.Section, string]
	sectionTypes  pipeline.Map[*asciidoc.Section, matter.Section]
	sectionLevels pipeline.Map[*asciidoc.Section, int]

	docTypes map[*asciidoc.Document]matter.DocType
}

type LibrarySet pipeline.Map[string, *pipeline.Data[*Library]]

func NewLibrary(root *asciidoc.Document, cache *DocCache) *Library {
	return &Library{
		Root:               root,
		referenceIndex:     newReferenceIndex(),
		crossReferenceDocs: make(map[*asciidoc.CrossReference]*asciidoc.Document),
		index:              map[string]*asciidoc.Document{},
		parents:            make(map[*asciidoc.Document][]*asciidoc.Document),
		children:           make(map[*asciidoc.Document][]*asciidoc.Document),
		entitiesByElement:  make(map[asciidoc.Attributable][]types.Entity),
		strings:            make(map[asciidoc.Parent]string),
		cache:              cache,
		sectionNames:       pipeline.NewConcurrentMap[*asciidoc.Section, string](),
		sectionTypes:       pipeline.NewConcurrentMap[*asciidoc.Section, matter.Section](),
		sectionLevels:      pipeline.NewConcurrentMap[*asciidoc.Section, int](),
		docTypes:           map[*asciidoc.Document]matter.DocType{},
	}
}

func (library *Library) CrossReferences(id string) []*CrossReference {
	return library.crossReferencesByID[id]
}

func (library *Library) CrossReferencesForDoc(doc *asciidoc.Document) map[string][]*CrossReference {
	return library.crossReferencesByDoc[doc]
}

func (library *Library) EntitiesForElement(element asciidoc.Attributable) (entities []types.Entity, ok bool) {
	entities, ok = library.entitiesByElement[element]
	return
}

func (library *Library) Parents(doc *asciidoc.Document) []*asciidoc.Document {
	return library.parents[doc]
}

func (library *Library) indexAnchors() (err error) {
	_, err = library.Anchors(library)
	if err != nil {
		return
	}

	return
}

func (library *Library) indexCrossReferences() {
	parse.Traverse(library.Root, library, library.Root, library.Children(library.Root), func(doc *asciidoc.Document, cr *asciidoc.CrossReference, parent asciidoc.ParentElement, offset int) parse.SearchShould {
		referenceID := library.elementIdentifier(library, cr, cr, cr.ID)
		c := &CrossReference{Document: doc, Reference: cr, Parent: parent, Source: NewSource(doc, cr)}
		library.crossReferencesByID[referenceID] = append(library.crossReferencesByID[referenceID], c)
		docReferences, ok := library.crossReferencesByDoc[doc]
		if !ok {
			docReferences = make(map[string][]*CrossReference, 0)
			library.crossReferencesByDoc[doc] = docReferences
		}
		docReferences[referenceID] = append(docReferences[referenceID], c)
		return parse.SearchShouldContinue
	})
}

func (library *Library) addEntity(element *asciidoc.Section, entity types.Entity) {
	library.entities = append(library.entities, entity)
	library.orderedEntities = append(library.orderedEntities, entity)
	library.entitiesByElement[element] = append(library.entitiesByElement[element], entity)
}

func (si *Specification) addEntity(entity types.Entity, cluster *matter.Cluster) {
	switch entity := entity.(type) {
	case *matter.Bitmap:
		si.addEntityByName(entity.Name, entity, cluster)
	case *matter.Enum:
		si.addEntityByName(entity.Name, entity, cluster)
	case *matter.Struct:
		si.addEntityByName(entity.Name, entity, cluster)
	case *matter.TypeDef:
		si.addEntityByName(entity.Name, entity, cluster)
	case *matter.Namespace:
		si.addEntityByName(entity.Name, entity, cluster)
	default:
		slog.Warn("Unexpected type adding entity to spec", log.Type("type", entity))
	}
}

func (si *Specification) addEntityByName(name string, entity types.Entity, cluster *matter.Cluster) {
	m, ok := si.entities[name]
	if !ok {
		m = make(map[types.Entity]map[*matter.Cluster]struct{})
		si.entities[name] = m
	}
	clusters, ok := m[entity]
	if !ok {
		clusters = make(map[*matter.Cluster]struct{})
		m[entity] = clusters
	}
	_, ok = clusters[cluster]
	if ok {
		slog.Debug("Registering same entity twice", "cluster", cluster.Name, "name", name, "address", fmt.Sprintf("%p", cluster))
		return
	}
	clusters[cluster] = struct{}{}
}

func (library *Library) SectionName(s *asciidoc.Section) (name string) {
	name, _ = library.sectionNames.Load(s)
	if name == "" {
		var title strings.Builder
		buildSectionTitle(&emptyVariableStore{}, s, library, &title, s.Title...)
		name = title.String()
		library.sectionNames.Store(s, name)
	}
	return
}

func (library *Library) SetSectionName(s *asciidoc.Section, name string) {
	library.sectionNames.Store(s, name)
}

func (library *Library) SectionType(s *asciidoc.Section) (st matter.Section) {
	st, _ = library.sectionTypes.Load(s)
	return
}

func (library *Library) SetSectionType(s *asciidoc.Section, st matter.Section) {
	//slog.Info("Setting section", "name", library.SectionName(s), "type", st.String(), log.Path("source", s))
	library.sectionTypes.Store(s, st)
}

func (library *Library) SectionLevel(s *asciidoc.Section) (level int) {
	var ok bool
	level, ok = library.sectionLevels.Load(s)
	if !ok {
		level = s.Level
	}
	return
}

func (library *Library) SetSectionLevel(s *asciidoc.Section, level int) {
	library.sectionLevels.Store(s, level)
}

func dumpLibrary(library *Library) {
	fmt.Fprintf(os.Stderr, "Library root: %s", library.Root.Path.Relative)
	for section := range parse.Skim[*asciidoc.Section](library, library.Root, library.Children(library.Root)) {
		dumpLibrarySection(library, section, 0)
	}
}

func dumpLibrarySection(library *Library, section *asciidoc.Section, indent int) {
	sectionName := library.SectionName(section)
	path, line := section.Origin()
	doc := section.Document()
	var docType matter.DocType
	if doc != nil {
		path = doc.Path.Relative
		docType, _ = library.DocType(doc)
	}
	fmt.Print(strings.Repeat("\t", indent))
	fmt.Printf("section %s: %s (%s:%d - %s)\n", sectionName, library.SectionType(section).String(), path, line, docType.String())
	for section := range parse.Skim[*asciidoc.Section](library, section, library.Children(section)) {
		dumpLibrarySection(library, section, indent+1)
	}

}
