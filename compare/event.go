package compare

import (
	"strings"

	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/types"
)

func compareEvent(specEvent *matter.Event, zapEvent *matter.Event) (diffs []Diff) {
	if !namesEqual(specEvent.Name, zapEvent.Name) {
		diffs = append(diffs, &StringDiff{Type: DiffTypeMismatch, Property: DiffPropertyName, Spec: specEvent.Name, ZAP: zapEvent.Name})
	}
	if !strings.EqualFold(specEvent.Priority, zapEvent.Priority) {
		diffs = append(diffs, &StringDiff{Type: DiffTypeMismatch, Property: DiffPropertyPriority, Spec: specEvent.Priority, ZAP: zapEvent.Priority})
	}
	diffs = append(diffs, compareAccess(types.EntityTypeEvent, specEvent.Access, zapEvent.Access)...)
	diffs = append(diffs, compareConformance(types.EntityTypeEvent, specEvent.Conformance, zapEvent.Conformance)...)
	fieldDiffs, err := compareFields(types.EntityTypeField, specEvent.Fields, zapEvent.Fields)
	if err == nil && len(fieldDiffs) > 0 {
		diffs = append(diffs, fieldDiffs...)
	}
	return
}

func compareEvents(specEvents []*matter.Event, zapEvents []*matter.Event) (diffs []Diff) {
	specEventMap := make(map[string]*matter.Event)
	for _, f := range specEvents {
		specEventMap[f.ID.IntString()] = f
	}

	zapEventMap := make(map[string]*matter.Event)
	for _, f := range zapEvents {
		zapEventMap[f.ID.IntString()] = f
	}
	for name, zapEvent := range zapEventMap {
		specName := name
		specEvent, ok := specEventMap[specName]
		if !ok {
			specName += "event"
			specEvent, ok = specEventMap[specName]
			if !ok {
				continue
			}
		}
		delete(zapEventMap, name)
		delete(specEventMap, specName)
		eventDiffs := compareEvent(specEvent, zapEvent)
		if len(eventDiffs) > 0 {
			diffs = append(diffs, &IdentifiedDiff{Type: DiffTypeMismatch, Entity: types.EntityTypeEvent, ID: specEvent.ID, Name: specEvent.Name, Diffs: eventDiffs})
		}
	}
	for _, f := range specEventMap {
		diffs = append(diffs, newMissingDiff(f.Name, types.EntityTypeEvent, SourceZAP))
	}
	for _, f := range zapEventMap {
		diffs = append(diffs, newMissingDiff(f.Name, types.EntityTypeEvent, SourceSpec))
	}
	return
}
