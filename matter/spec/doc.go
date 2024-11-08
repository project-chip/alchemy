package spec

import (
	"log/slog"
	"sync"
	"unicode"
	"unicode/utf8"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

type Doc struct {
	sync.RWMutex

	Path asciidoc.Path

	Base *asciidoc.Document
	asciidoc.Set

	docType matter.DocType

	Domain   matter.Domain
	parents  []*Doc
	children []*Doc

	referenceIndex
	attributes map[asciidoc.AttributeName]any

	parsed            bool // Tracks whether this doc was parsed vs. just read (i.e. were file substituions done)
	entities          []types.Entity
	orderedEntities   []types.Entity
	globalObjects     []types.Entity
	entitiesBySection map[asciidoc.Attributable][]types.Entity
	entitiesParsed    bool

	spec  *Specification
	group *DocGroup

	errata *errata.Errata
}

func newDoc(d *asciidoc.Document, path asciidoc.Path) (*Doc, error) {
	doc := &Doc{
		Base:           d,
		Path:           path,
		attributes:     make(map[asciidoc.AttributeName]any),
		referenceIndex: newReferenceIndex(),
	}
	doc.errata = errata.GetErrata(path.Relative)
	for _, e := range d.Elements() {
		switch el := e.(type) {
		case *asciidoc.AttributeEntry:
			doc.attributes[el.Name] = el.Elements()
			doc.Append(NewElement(doc, e))
		case *asciidoc.Section:
			s, err := NewSection(doc, doc, el)
			if err != nil {
				return nil, err
			}
			doc.Append(s)
		default:
			doc.Append(NewElement(doc, e))
		}
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

func (doc *Doc) Parents() []*Doc {
	doc.RLock()
	p := make([]*Doc, len(doc.parents))
	copy(p, doc.parents)
	doc.RUnlock()
	return p
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

func (doc *Doc) Entities() (entities []types.Entity, err error) {
	if !doc.entitiesParsed {
		err = doc.parseEntities()
		if err != nil {
			return nil, err
		}
	}
	return doc.entities, nil
}

func (doc *Doc) GlobalObjects() (entities []types.Entity, err error) {
	if !doc.entitiesParsed {
		err = doc.parseEntities()
		if err != nil {
			return nil, err
		}
	}
	return doc.globalObjects, nil
}

func (doc *Doc) OrderedEntities() (entities []types.Entity, err error) {
	if !doc.entitiesParsed {
		err = doc.parseEntities()
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

func (doc *Doc) Reference(ref string) (types.Entity, bool) {

	a := doc.FindAnchor(ref)

	if a == nil {
		slog.Warn("unknown reference", slog.String("path", doc.Path.String()), slog.String("reference", ref))
		return nil, false
	}
	wa, ok := a.Element.(asciidoc.Attributable)
	if !ok {
		slog.Warn("reference to non-entity", slog.String("path", doc.Path.String()), slog.String("reference", ref))
		return nil, false
	}
	entities, ok := doc.entitiesBySection[wa]
	if !ok {
		slog.Warn("unknown reference entity", slog.String("path", doc.Path.String()), slog.String("reference", ref), slog.Any("count", len(doc.entitiesBySection)))
		for sec, e := range doc.entitiesBySection {
			slog.Warn("reference", slog.String("path", doc.Path.String()), slog.String("reference", ref), slog.Any("sec", sec), slog.Any("entity", e))

		}
	}
	if len(entities) == 0 {
		slog.Warn("unknown reference entity", slog.String("path", doc.Path.String()), slog.String("reference", ref))
		return nil, false
	}
	if len(entities) > 1 {
		slog.Warn("ambiguous reference", slog.String("path", doc.Path.String()), slog.String("reference", ref))
		for _, e := range entities {
			slog.Warn("reference", slog.String("path", doc.Path.String()), slog.String("reference", ref), slog.Any("entity", e))

		}
		return nil, false
	}
	return entities[0], true
}
