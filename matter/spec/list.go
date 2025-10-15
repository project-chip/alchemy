package spec

import (
	"iter"
	"log/slog"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

type listIndex[T types.Entity] struct {
	byName      map[string]T
	byReference map[asciidoc.Element]T
}

type entityFactory[T types.Entity] interface {
	New(spec *Specification, library *Library, reader asciidoc.Reader, d *asciidoc.Document, s *asciidoc.Section, ti *TableInfo, row *asciidoc.TableRow, name string, parent types.Entity) (T, error)
	Details(spec *Specification, library *Library, reader asciidoc.Reader, d *asciidoc.Document, s *asciidoc.Section, e T) error
	EntityName(library *Library, reader asciidoc.Reader, d *asciidoc.Document, s *asciidoc.Section) string
	Children(library *Library, reader asciidoc.Reader, d *asciidoc.Document, s *asciidoc.Section) iter.Seq[*asciidoc.Section]
}

func buildList[T types.Entity, L ~[]T](spec *Specification, library *Library, reader asciidoc.Reader, doc *asciidoc.Document, section *asciidoc.Section, table *asciidoc.Table, list L, factory entityFactory[T], parent types.Entity) (L, error) {

	index := listIndex[T]{
		byName:      make(map[string]T),
		byReference: make(map[asciidoc.Element]T),
	}
	var err error
	var ti *TableInfo
	ti, err = ReadTable(doc, reader, table)
	if err != nil {
		return nil, err
	}
	for row := range ti.ContentRows() {

		var name string
		var xref *asciidoc.CrossReference
		name, xref, err = ti.ReadName(library, row, matter.TableColumnName)
		if err != nil {
			return nil, err
		}

		var entity T
		entity, err = factory.New(spec, library, reader, doc, section, ti, row, name, parent)
		if err != nil {
			return nil, err
		}

		list = append(list, entity)
		index.byName[strings.ToLower(name)] = entity
		if xref != nil {
			anchor := library.FindAnchorByID(xref.ID, xref, xref)
			if anchor != nil && anchor.Element != nil {
				index.byReference[anchor.Element] = entity
			}
		}
	}

	for subSection := range factory.Children(library, reader, doc, section) {
		e, ok := index.byReference[subSection]
		if !ok {
			name := factory.EntityName(library, reader, doc, subSection)
			e, ok = index.byName[strings.ToLower(name)]
			if !ok {
				slog.Warn("unknown entity attempting to match sub section to parent table", log.Path("sectionPath", subSection), log.Path("parentTablePath", table), slog.String("entityName", library.SectionName(subSection)), log.Element("source", doc.Path, subSection))
				continue
			}
		}
		err = factory.Details(spec, library, reader, doc, subSection, e)
		if err != nil {
			return nil, err
		}
		library.addEntity(subSection, e)
	}
	return list, nil
}
