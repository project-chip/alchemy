package spec

import (
	"fmt"
	"log/slog"
	"sync"
	"unicode"
	"unicode/utf8"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/internal/parse"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

type Doc struct {
	sync.RWMutex

	Path string

	Base *asciidoc.Document
	asciidoc.Set

	docType matter.DocType

	Domain   matter.Domain
	parents  []*Doc
	children []*Doc

	referenceIndex
	attributes map[asciidoc.AttributeName]any

	entities       []types.Entity
	entitiesParsed bool

	entitiesBySection map[asciidoc.Attributable][]types.Entity

	spec  *Specification
	group *DocGroup
}

func NewDoc(d *asciidoc.Document, path string) (*Doc, error) {
	doc := &Doc{
		Base:           d,
		Path:           path,
		attributes:     make(map[asciidoc.AttributeName]any),
		referenceIndex: newReferenceIndex(),
	}
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
	if doc.entitiesParsed {
		return doc.entities, nil
	}
	doc.entitiesParsed = true

	var entitiesBySection = make(map[asciidoc.Attributable][]types.Entity)
	for _, top := range parse.Skim[*Section](doc.Elements()) {
		err := AssignSectionTypes(doc, top)
		if err != nil {
			return nil, err
		}

		var m []types.Entity
		m, err = top.toEntities(doc, entitiesBySection)
		if err != nil {
			return nil, fmt.Errorf("failed converting doc %s to entities: %w", doc.Path, err)
		}
		entities = append(entities, m...)
	}
	doc.entities = entities
	doc.entitiesBySection = entitiesBySection
	return
}

func (doc *Doc) Reference(ref string) (types.Entity, bool) {

	a := doc.FindAnchor(ref)

	if a == nil {
		slog.Warn("unknown reference", slog.String("path", doc.Path), slog.String("reference", ref))
		return nil, false
	}
	wa, ok := a.Element.(asciidoc.Attributable)
	if !ok {
		slog.Warn("reference to non-entity", slog.String("path", doc.Path), slog.String("reference", ref))
		return nil, false
	}
	entities, ok := doc.entitiesBySection[wa]
	if !ok {
		slog.Warn("unknown reference entity", slog.String("path", doc.Path), slog.String("reference", ref), slog.Any("count", len(doc.entitiesBySection)))
		for sec, e := range doc.entitiesBySection {
			slog.Warn("reference", slog.String("path", doc.Path), slog.String("reference", ref), slog.Any("sec", sec), slog.Any("entity", e))

		}
	}
	if len(entities) == 0 {
		slog.Warn("unknown reference entity", slog.String("path", doc.Path), slog.String("reference", ref))
		return nil, false
	}
	if len(entities) > 1 {
		slog.Warn("ambiguous reference", slog.String("path", doc.Path), slog.String("reference", ref))
		for _, e := range entities {
			slog.Warn("reference", slog.String("path", doc.Path), slog.String("reference", ref), slog.Any("entity", e))

		}
		return nil, false
	}
	return entities[0], true
}

func GithubSettings() []asciidoc.AttributeName {
	return []asciidoc.AttributeName{asciidoc.AttributeName("env-github")}
}
