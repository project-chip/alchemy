package spec

import (
	"iter"
	"log/slog"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/internal/parse"
	"github.com/project-chip/alchemy/internal/suggest"
	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

type eventFactory struct{}

func (cf *eventFactory) New(spec *Specification, d *Doc, s *Section, ti *TableInfo, row *asciidoc.TableRow, name string, parent types.Entity) (e *matter.Event, err error) {

	e = matter.NewEvent(s.Base, parent)
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

func (cf *eventFactory) Details(spec *Specification, d *Doc, s *Section, pc *parseContext, e *matter.Event) (err error) {
	e.Description = getDescription(d, e, s.Set)
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
	var fieldMap map[string]*matter.Field
	e.Fields, fieldMap, err = d.readFields(spec, ti, types.EntityTypeEventField, e)
	if err != nil {
		return
	}
	err = s.mapFields(fieldMap, pc)
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

func (s *Section) toEvents(spec *Specification, d *Doc, pc *parseContext, parent types.Entity) (events matter.EventSet, err error) {
	t := FindFirstTable(s)
	if t == nil {
		return nil, nil
	}

	var ef eventFactory
	events, err = buildList(spec, d, s, t, pc, events, &ef, parent)

	if err != nil {
		return
	}

	eventMap := make(map[uint64]*matter.Event)
	eventNameMap := make(map[string]*matter.Event)
	for _, ev := range events {
		existing, ok := eventMap[ev.ID.Value()]
		if ok {
			slog.Error("Duplicate Event ID",
				slog.String("clusterName",
					matter.EntityName(parent)),
				slog.String("eventName", ev.Name),
				slog.String("eventId", ev.ID.IntString()),
				slog.String("previousEvent", existing.Name),
			)
			spec.addError(&DuplicateEntityIDError{Entity: ev, Previous: existing})
		}
		existing, ok = eventNameMap[ev.Name]
		if ok {
			slog.Error("Duplicate Event Name",
				slog.String("clusterName",
					matter.EntityName(parent)),
				slog.String("eventName", ev.Name),
				slog.String("eventId", ev.ID.IntString()),
				slog.String("previousEvent", existing.Name),
			)
			spec.addError(&DuplicateEntityNameError{Entity: ev, Previous: existing})
		}
		eventMap[ev.ID.Value()] = ev
		eventNameMap[ev.Name] = ev
	}
	return
}

type eventFinder struct {
	entityFinderCommon

	events []*matter.Event
}

func newEventFinder(events []*matter.Event, inner entityFinder) *eventFinder {
	return &eventFinder{entityFinderCommon: entityFinderCommon{inner: inner}, events: events}
}

func (ef *eventFinder) findEntityByIdentifier(identifier string, source log.Source) types.Entity {
	for _, c := range ef.events {
		if c.Name == identifier && c != ef.identity {
			return c
		}
	}
	if ef.inner != nil {
		return ef.inner.findEntityByIdentifier(identifier, source)
	}
	return nil
}

func (ef *eventFinder) suggestIdentifiers(identifier string, suggestions map[types.Entity]int) {
	suggest.PossibleEntities(identifier, suggestions, func(yield func(string, types.Entity) bool) {
		for _, f := range ef.events {

			if f == ef.identity {
				continue
			}
			if !yield(f.Name, f) {
				return
			}

		}
	})
	if ef.inner != nil {
		ef.inner.suggestIdentifiers(identifier, suggestions)
	}
	return
}

func validateEvents(spec *Specification) {
	for c := range spec.Clusters {
		for _, e := range c.Events {
			validateFields(spec, e, e.Fields)
		}
	}
	for obj := range spec.GlobalObjects {
		switch obj := obj.(type) {
		case *matter.Event:
			validateFields(spec, obj, obj.Fields)
		}
	}
}
