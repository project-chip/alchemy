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
	New(spec *Specification, d *Doc, s *Section, ti *TableInfo, row *asciidoc.TableRow, name string, parent types.Entity) (T, error)
	Details(spec *Specification, d *Doc, s *Section, pc *parseContext, e T) error
	EntityName(s *Section) string
	Children(d *Doc, s *Section) iter.Seq[*Section]
}

func buildList[T types.Entity, L ~[]T](spec *Specification, d *Doc, s *Section, t *asciidoc.Table, pc *parseContext, list L, factory entityFactory[T], parent types.Entity) (L, error) {

	index := listIndex[T]{
		byName:      make(map[string]T),
		byReference: make(map[asciidoc.Element]T),
	}
	var err error
	var ti *TableInfo
	ti, err = ReadTable(d, t)
	if err != nil {
		return nil, err
	}
	for row := range ti.ContentRows() {

		var name string
		var xref *asciidoc.CrossReference
		name, xref, err = ti.ReadName(row, matter.TableColumnName)
		if err != nil {
			return nil, err
		}

		var entity T
		entity, err = factory.New(spec, d, s, ti, row, name, parent)
		if err != nil {
			return nil, err
		}

		list = append(list, entity)
		index.byName[strings.ToLower(name)] = entity
		if xref != nil {
			anchor := d.FindAnchor(xref.ID, xref)
			if anchor != nil && anchor.Element != nil {
				index.byReference[anchor.Element] = entity
			}
		}
	}

	for s := range factory.Children(d, s) {
		e, ok := index.byReference[s.Base]
		if !ok {
			name := factory.EntityName(s)
			e, ok = index.byName[strings.ToLower(name)]
			if !ok {
				slog.Warn("unknown entity", log.Element("source", d.Path, s.Base), "entityName", s.Name)
				continue
			}
		}
		err = factory.Details(spec, d, s, pc, e)
		if err != nil {
			return nil, err
		}
		pc.orderedEntities = append(pc.orderedEntities, e)
		pc.entitiesByElement[s.Base] = append(pc.entitiesByElement[s.Base], e)
	}
	return list, nil
}
