package spec

import (
	"fmt"
	"iter"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/internal/parse"
	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

type eventFactory struct{}

func (cf *eventFactory) New(d *Doc, s *Section, ti *TableInfo, row *asciidoc.TableRow, name string) (e *matter.Event, err error) {

	e = matter.NewEvent(s.Base)
	e.Name = matter.StripTypeSuffixes(name)
	e.ID, err = ti.ReadID(row, matter.TableColumnID)
	if err != nil {
		return
	}
	e.Priority, err = ti.ReadString(row, matter.TableColumnPriority)
	if err != nil {
		return
	}
	e.Conformance = ti.ReadConformance(row, matter.TableColumnConformance)
	var a string
	a, err = ti.ReadString(row, matter.TableColumnAccess)
	if err != nil {
		return
	}
	e.Access, _ = ParseAccess(a, types.EntityTypeEvent)
	if e.Access.Read == matter.PrivilegeUnknown {
		// Sometimes the invoke access is omitted; we assume it's view
		e.Access.Read = matter.PrivilegeView
	}
	e.Name = CanonicalName(e.Name)
	return
}

func (cf *eventFactory) Details(d *Doc, s *Section, entityMap map[asciidoc.Attributable][]types.Entity, e *matter.Event) (err error) {
	e.Description = getDescription(d, s.Set)
	var ti *TableInfo
	ti, err = parseFirstTable(d, s)
	if err != nil {
		if err == ErrNoTableFound {
			err = nil
			return
		}
		return
	}
	if ti.HeaderRowIndex > 0 {
		firstRow := ti.Rows[0]
		tableCells := firstRow.TableCells()
		if len(tableCells) > 0 {
			cv, rowErr := RenderTableCell(tableCells[0])
			if rowErr == nil {
				cv = strings.ToLower(cv)
				if strings.Contains(cv, "fabric sensitive") || strings.Contains(cv, "fabric-sensitive") {
					e.Access.FabricSensitivity = matter.FabricSensitivitySensitive
				}
			}
		}
	}
	if err != nil {
		if err == ErrNoTableFound {
			err = nil
			return
		}
		err = fmt.Errorf("failed reading %s event fields: %w", s.Name, err)
		return
	}
	e.Fields, err = d.readFields(ti, types.EntityTypeEventField)
	if err != nil {
		return
	}
	entityMap[s.Base] = append(entityMap[s.Base], e)
	fieldMap := make(map[string]*matter.Field, len(e.Fields))
	for _, f := range e.Fields {
		fieldMap[f.Name] = f
	}
	err = s.mapFields(fieldMap, entityMap)
	if err != nil {
		return
	}
	for _, f := range e.Fields {
		f.Name = CanonicalName(f.Name)
	}
	return
}

func (cf *eventFactory) EntityName(s *Section) string {
	return strings.ToLower(text.TrimCaseInsensitiveSuffix(s.Name, " Event"))
}

func (cf *eventFactory) Children(d *Doc, s *Section) iter.Seq[*Section] {
	return func(yield func(*Section) bool) {
		parse.SkimFunc(s.Elements(), func(s *Section) bool {
			if s.SecType != matter.SectionEvent {
				return false
			}
			return !yield(s)
		})
	}
}

func (s *Section) toEvents(d *Doc, entityMap map[asciidoc.Attributable][]types.Entity) (events matter.EventSet, err error) {
	t := FindFirstTable(s)
	if t == nil {
		return nil, nil
	}

	var ef eventFactory
	events, err = buildList(d, s, t, entityMap, events, &ef)

	/*var rows []*asciidoc.TableRow
	var headerRowIndex int
	var columnMap ColumnIndex
	rows, headerRowIndex, columnMap, _, err = parseFirstTable(d, s)
	if err != nil {
		return nil, fmt.Errorf("failed reading events: %w", err)
	}

	eventMap := make(map[string]*matter.Event)
	for i := headerRowIndex + 1; i < len(rows); i++ {
		row := rows[i]
		e := matter.NewEvent(s.Base)
		e.Name, err = ReadRowValue(d, row, columnMap, matter.TableColumnName)
		if err != nil {
			return
		}
		e.Name = matter.StripTypeSuffixes(e.Name)
		e.ID, err = readRowID(row, columnMap, matter.TableColumnID)
		if err != nil {
			return
		}
		e.Priority, err = readRowASCIIDocString(row, columnMap, matter.TableColumnPriority)
		if err != nil {
			return
		}
		e.Conformance = d.getRowConformance(row, columnMap, matter.TableColumnConformance)
		var a string
		a, err = readRowASCIIDocString(row, columnMap, matter.TableColumnAccess)
		if err != nil {
			return
		}
		e.Access, _ = ParseAccess(a, types.EntityTypeEvent)
		if e.Access.Read == matter.PrivilegeUnknown {
			// Sometimes the invoke access is omitted; we assume it's view
			e.Access.Read = matter.PrivilegeView
		}
		events = append(events, e)

		eventMap[e.Name] = e
		e.Name = CanonicalName(e.Name)
	}

	for _, s := range parse.Skim[*Section](s.Elements()) {
		switch s.SecType {
		case matter.SectionEvent:

			name := text.TrimCaseInsensitiveSuffix(s.Name, " Event")
			e, ok := eventMap[name]
			if !ok {
				slog.Debug("unknown event", "event", name)
				continue
			}

		}
	}*/
	return
}
