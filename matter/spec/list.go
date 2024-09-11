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
	New(d *Doc, s *Section, row *asciidoc.TableRow, columnMap ColumnIndex, name string) (T, error)
	Details(d *Doc, s *Section, entityMap map[asciidoc.Attributable][]types.Entity, e T) error
	EntityName(s *Section) string
	Children(d *Doc, s *Section) iter.Seq[*Section]
}

func buildList[T types.Entity, L ~[]T](d *Doc, s *Section, t *asciidoc.Table, entityMap map[asciidoc.Attributable][]types.Entity, list L, factory entityFactory[T]) (L, error) {

	index := listIndex[T]{
		byName:      make(map[string]T),
		byReference: make(map[asciidoc.Element]T),
	}
	var rows []*asciidoc.TableRow
	var headerRowIndex int
	var columnMap ColumnIndex
	var err error
	rows, headerRowIndex, columnMap, _, err = parseTable(d, s, t)
	if err != nil {
		return nil, err
	}
	for i := headerRowIndex + 1; i < len(rows); i++ {
		row := rows[i]

		var name string
		var xref *asciidoc.CrossReference
		name, xref, err = readRowName(d, row, columnMap, matter.TableColumnName)
		if err != nil {
			return nil, err
		}

		var entity T
		entity, err = factory.New(d, s, row, columnMap, name)
		if err != nil {
			return nil, err
		}

		list = append(list, entity)
		index.byName[strings.ToLower(name)] = entity
		if xref != nil {
			anchor := d.FindAnchor(xref.ID)
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
				slog.Warn("unknown entity", log.Element("path", d.Path, s.Base), "entityName", s.Name)
				continue
			}
		}
		err = factory.Details(d, s, entityMap, e)
		if err != nil {
			return nil, err
		}
		entityMap[s.Base] = append(entityMap[s.Base], e)
	}
	return list, nil
}
