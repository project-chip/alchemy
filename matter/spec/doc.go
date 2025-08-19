package spec

import (
	"unicode"
	"unicode/utf8"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

type Doc struct {
	Path asciidoc.Path

	Base *asciidoc.Document
	//asciidoc.Elements

	docType matter.DocType

	Domain   matter.Domain
	parents  []*Doc
	children []*Doc

	referenceIndex

	parsed            bool // Tracks whether this doc was parsed vs. just read (i.e. were file substituions done)
	entities          []types.Entity
	orderedEntities   []types.Entity
	globalObjects     []types.Entity
	entitiesBySection map[asciidoc.Attributable][]types.Entity
	entitiesParsed    bool

	spec  *Specification
	group *DocGroup

	errata *errata.Errata

	reader asciidoc.Reader
}

type DocSet pipeline.Map[string, *pipeline.Data[*Doc]]

func NewDocSet() DocSet {
	return pipeline.NewMap[string, *pipeline.Data[*Doc]]()
}

func newDoc(d *asciidoc.Document, path asciidoc.Path) (*Doc, error) {
	doc := &Doc{
		Base:           d,
		Path:           path,
		referenceIndex: newReferenceIndex(),
		errata:         errata.GetErrata(path.Relative),
	}
	return doc, nil
}

func firstLetterIsLower(s string) bool {
	firstLetter, _ := utf8.DecodeRuneInString(s)
	return unicode.IsLower(firstLetter)
}

func (doc *Doc) Footnotes() []*asciidoc.Footnote {
	return nil
}

func (doc *Doc) Errata() *errata.Errata {
	return doc.errata
}

func (doc *Doc) Reader() asciidoc.Reader {
	if doc.reader != nil {
		return doc.reader
	}
	if doc.group != nil {
		return doc.group.Reader
	}
	return asciidoc.NewRawReader()
}

func (doc *Doc) Parents() []*Doc {
	doc.RLock()
	p := make([]*Doc, len(doc.parents))
	copy(p, doc.parents)
	doc.RUnlock()
	return p
}

func (doc *Doc) Children() asciidoc.Elements {
	return doc.Base.Children()
}

func (doc *Doc) SetChildren(e asciidoc.Elements) {
	doc.Base.SetChildren(e)
}

func (doc *Doc) Append(e ...asciidoc.Element) {
	doc.Base.Append(e...)
}

func (doc *Doc) Type() asciidoc.ElementType {
	return asciidoc.ElementTypeDocument
}

func (doc *Doc) Equals(other asciidoc.Element) bool {
	if other, ok := other.(*Doc); ok {
		return doc.Path.Absolute == other.Path.Absolute
	}
	return false
}

func (doc *Doc) Group() *DocGroup {
	return doc.group
}

func (doc *Doc) addParent(parent *Doc) {
	doc.Lock()
	doc.parents = append(doc.parents, parent)
	doc.Unlock()
}

func (doc *Doc) addChild(child *Doc) {
	doc.Lock()
	doc.children = append(doc.children, child)
	doc.Unlock()
}

func (doc *Doc) SectionName(s *asciidoc.Section) (name string) {
	var ok bool
	name, ok = doc.sectionName(s)
	if !ok && doc.group != nil {
		name, _ = doc.group.sectionName(s)
	}
	return
}

func (doc *Doc) SetSectionName(s *asciidoc.Section, name string) {
	doc.setSectionName(s, name)
	if doc.group != nil {
		doc.group.setSectionName(s, name)
	}
}

func (doc *Doc) SectionType(s *asciidoc.Section) (st matter.Section) {
	var ok bool
	st, ok = doc.sectionType(s)
	if !ok && doc.group != nil {
		st, _ = doc.group.sectionType(s)
	}
	return
}

func (doc *Doc) SetSectionType(s *asciidoc.Section, st matter.Section) {
	doc.setSectionType(s, st)
	if doc.group != nil {
		doc.group.setSectionType(s, st)
	}
}

func (doc *Doc) Entities() (entities []types.Entity, err error) {
	if !doc.entitiesParsed {
		err = doc.parseEntities(nil)
		if err != nil {
			return nil, err
		}
	}
	return doc.entities, nil
}

func (doc *Doc) GlobalObjects() (entities []types.Entity, err error) {
	if !doc.entitiesParsed {
		err = doc.parseEntities(nil)
		if err != nil {
			return nil, err
		}
	}
	return doc.globalObjects, nil
}

func (doc *Doc) OrderedEntities() (entities []types.Entity, err error) {
	if !doc.entitiesParsed {
		err = doc.parseEntities(nil)
		if err != nil {
			return nil, err
		}
	}
	return doc.orderedEntities, nil
}

func (d *Doc) EntitiesForSection(section *asciidoc.Section) ([]types.Entity, bool) {
	e, ok := d.entitiesBySection[section]
	return e, ok
}
