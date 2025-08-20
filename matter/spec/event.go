package spec

import (
	"iter"
	"log/slog"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/asciidoc/parse"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/internal/suggest"
	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

type eventFactory struct{}

func (cf *eventFactory) New(spec *Specification, library *Library, reader asciidoc.Reader, d *asciidoc.Document, s *asciidoc.Section, ti *TableInfo, row *asciidoc.TableRow, name string, parent types.Entity) (e *matter.Event, err error) {

	e = matter.NewEvent(s, parent)
	e.Name = matter.StripTypeSuffixes(name)
	e.ID, err = ti.ReadID(reader, row, matter.TableColumnEventID, matter.TableColumnID)
	if err != nil {
		return
	}
	e.Priority, err = ti.ReadString(reader, row, matter.TableColumnPriority)
	if err != nil {
		return
	}
	e.Conformance = ti.ReadConformance(library, row, matter.TableColumnConformance)
	var a string
	a, err = ti.ReadString(reader, row, matter.TableColumnAccess)
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

func (cf *eventFactory) Details(spec *Specification, library *Library, reader asciidoc.Reader, doc *asciidoc.Document, section *asciidoc.Section, e *matter.Event) (err error) {
	e.Description = library.getDescription(reader, doc, e, section, reader.Children(section))
	var ti *TableInfo
	ti, err = parseFirstTable(reader, doc, section)
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
			cv, rowErr := RenderTableCell(reader, tableCells[0])
			if rowErr == nil {
				cv = strings.ToLower(cv)
				if strings.Contains(cv, "fabric sensitive") || strings.Contains(cv, "fabric-sensitive") {
					e.Access.FabricSensitivity = matter.FabricSensitivitySensitive
				}
			}
		}
	}
	var fieldMap map[string]*matter.Field
	e.Fields, fieldMap, err = library.readFields(spec, reader, ti, types.EntityTypeEventField, e)
	if err != nil {
		return
	}
	err = library.mapFields(reader, doc, section, fieldMap)
	if err != nil {
		return
	}
	for _, f := range e.Fields {
		f.Name = CanonicalName(f.Name)
	}
	return
}

func (cf *eventFactory) EntityName(library *Library, reader asciidoc.Reader, doc *asciidoc.Document, s *asciidoc.Section) string {
	return strings.ToLower(text.TrimCaseInsensitiveSuffix(library.SectionName(s), " Event"))
}

func (cf *eventFactory) Children(library *Library, reader asciidoc.Reader, d *asciidoc.Document, s *asciidoc.Section) iter.Seq[*asciidoc.Section] {
	return func(yield func(*asciidoc.Section) bool) {
		parse.SkimFunc(library, s, reader.Children(s), func(s *asciidoc.Section) bool {
			if library.SectionType(s) != matter.SectionEvent {
				return false
			}
			return !yield(s)
		})
	}
}

func (library *Library) toEvents(spec *Specification, reader asciidoc.Reader, d *asciidoc.Document, s *asciidoc.Section, parent types.Entity) (events matter.EventSet, err error) {
	t := FindFirstTable(reader, s)
	if t == nil {
		return nil, nil
	}

	var ef eventFactory
	events, err = buildList(spec, library, reader, d, s, t, events, &ef, parent)

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
}

func validateEvents(spec *Specification) {
	for c := range spec.Clusters {
		idu := make(idUniqueness[*matter.Event])
		nu := make(nameUniqueness[*matter.Event])
		cv := make(conformanceValidation)
		for _, e := range c.Events {
			idu.check(spec, e.ID, e)
			nu.check(spec, e)
			cv.add(e, e.Conformance)
			validateFields(spec, e, e.Fields)
		}
		cv.check(spec)
	}
	idu := make(idUniqueness[*matter.Event])
	nu := make(nameUniqueness[*matter.Event])
	for obj := range spec.GlobalObjects {
		switch e := obj.(type) {
		case *matter.Event:
			idu.check(spec, e.ID, e)
			nu.check(spec, e)
			validateFields(spec, e, e.Fields)
		}
	}
}
