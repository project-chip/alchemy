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
	globalObjects     []types.Entity
	entitiesByElement map[asciidoc.Attributable][]types.Entity

	//preparseState map[asciidoc.Element]*elementOverlay

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
		children:           map[*asciidoc.Document][]*asciidoc.Document{},
		entitiesByElement:  make(map[asciidoc.Attributable][]types.Entity),
		strings:            make(map[asciidoc.Parent]string),
		cache:              cache,
		sectionNames:       pipeline.NewConcurrentMap[*asciidoc.Section, string](),
		sectionTypes:       pipeline.NewConcurrentMap[*asciidoc.Section, matter.Section](),
		sectionLevels:      pipeline.NewConcurrentMap[*asciidoc.Section, int](),
		docTypes:           map[*asciidoc.Document]matter.DocType{},
	}
}

/*
func (library *Library) Iterate(parent asciidoc.Parent, elements asciidoc.Elements) asciidoc.ElementIterator {
	return func(yield func(asciidoc.Element) bool) {
		for _, el := range elements {
			action, ok := library.preparseState[el]
			if !ok {
				if !yield(el) {
					return
				}
			} else {
				if action.remove {
					continue
				}
				if len(action.replace) > 0 {
					for _, el := range action.replace {
						if !yield(el) {
							return
						}
					}
				}
			}
		}
		switch parent := parent.(type) {
		case asciidoc.Element:
			if action, ok := library.preparseState[parent]; ok {
				if len(action.append) > 0 {
					for _, el := range action.append {
						if !yield(el) {
							return
						}
					}
				}
			}
		}
	}
}

func (library *Library) StringValue(parent asciidoc.Parent, elements asciidoc.Elements) (string, error) {
	val, ok := library.strings[parent]
	if ok {
		return val, nil
	}
	var s strings.Builder
	for el := range library.Iterate(parent, elements) {
		switch el := el.(type) {
		case *asciidoc.String:
			s.WriteString(el.Value)
		default:
			return "", fmt.Errorf("unexpected type in anchor id: %T", el)
		}
	}
	library.strings[parent] = s.String()
	return s.String(), nil
}
*/

func setDocGroup(d *asciidoc.Document, docGroup *Library) {
	/*if d.group != nil {
		if d.group.Root.Path.Relative != docGroup.Root.Path.Relative && d.Path.Base() != "matter-defines.adoc" {
			slog.Warn("multiple doc group roots", "path", d.Path.String(), "root", d.group.Root, "newRoot", docGroup.Root)
		}
		return
	}*/
	docGroup.Docs = append(docGroup.Docs, d)
	docGroup.index[d.Path.Relative] = d
	//d.group = docGroup
	/*for _, c := range d.children {
		setDocGroup(c, docGroup)
	}*/
}

/*
func setSpec(d *asciidoc.Document, si *Specification) {

		d.spec = si
		for _, c := range d.children {
			setSpec(c, si)
		}
	}
*/
/*func (dg *DocGroup) Anchors(id string) []*Anchor {
	return dg.anchors[id]
}*/

func (library *Library) CrossReferences(id string) []*CrossReference {
	return library.crossReferences[id]
}

func (library *Library) Parents(doc *asciidoc.Document) []*asciidoc.Document {
	return library.parents[doc]
}

func (library *Library) indexAnchors() (err error) {
	//var anchors map[string][]*Anchor
	_, err = library.Anchors(library)
	if err != nil {
		return
	}
	/*for id, anchor := range anchors {
		library.anchorsByLabel[id] = append(library.anchorsByLabel[id], anchor...)
	}
	*/
	return
}

func (library *Library) indexCrossReferences() {
	parse.Traverse(library.Root, library, library.Root, library.Children(library.Root), func(doc *asciidoc.Document, cr *asciidoc.CrossReference, parent asciidoc.ParentElement, offset int) parse.SearchShould {
		referenceID := library.anchorId(library, cr, cr, cr.ID)
		//slog.Info("adding cross reference", "id", referenceID, log.Path("source", cr))
		library.crossReferences[referenceID] = append(library.crossReferences[referenceID], &CrossReference{Document: doc, Reference: cr, Parent: parent, Source: NewSource(doc, cr)})
		return parse.SearchShouldContinue
	})
	/*for _, d := range library.Docs {
		if errata.GetSpec(d.Path.Relative).UtilityInclude {
			continue
		}
		crossReferences := d.CrossReferences(library.Reader)
		for id, xrefs := range crossReferences {
			d.group.crossReferences[id] = append(d.group.crossReferences[id], xrefs...)
			for _, xref := range xrefs {
				d.group.crossReferenceDocs[xref.Reference] = d
			}
		}
	}*/
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
