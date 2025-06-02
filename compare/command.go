package compare

import (
	"log/slog"

	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/types"
)

func compareCommand(specCommand *matter.Command, zapCommand *matter.Command) (diffs []Diff) {
	if !namesEqual(specCommand.Name, zapCommand.Name) {
		diffs = append(diffs, &StringDiff{Type: DiffTypeMismatch, Property: DiffPropertyName, Spec: specCommand.Name, ZAP: zapCommand.Name})
	}
	if !specCommand.Response.Equals(zapCommand.Response) && !(specCommand.Response == nil && zapCommand.Response.Name == "N") {
		var specResponse, zapResponse string
		if specCommand.Response != nil {
			specResponse = specCommand.Response.Name
		}
		if zapCommand.Response != nil {
			zapResponse = zapCommand.Response.Name
		}
		diffs = append(diffs, &StringDiff{Type: DiffTypeMismatch, Property: DiffPropertyCommandResponse, Spec: specResponse, ZAP: zapResponse})
	}
	if specCommand.Direction != zapCommand.Direction {
		diffs = append(diffs, &StringDiff{Type: DiffTypeMismatch, Property: DiffPropertyCommandDirection, Spec: specCommand.Direction.String(), ZAP: zapCommand.Direction.String()})
	}
	diffs = append(diffs, compareAccess(types.EntityTypeCommand, specCommand.Access, zapCommand.Access)...)
	diffs = append(diffs, compareConformance(types.EntityTypeCommand, specCommand.Conformance, zapCommand.Conformance)...)
	fieldDiffs, err := compareFields(types.EntityTypeCommandField, specCommand.Fields, zapCommand.Fields)
	if err == nil && len(fieldDiffs) > 0 {
		diffs = append(diffs, fieldDiffs...)
	}
	return
}

func compareCommands(specCommands matter.CommandSet, zapCommands []*matter.Command) (diffs []Diff) {
	specCommandMap := make(map[uint64]*matter.Command)
	specResponseMap := make(map[uint64]*matter.Command)
	for _, f := range specCommands {
		if conformance.IsZigbee(f.Conformance) || conformance.IsDisallowed(f.Conformance) {
			continue
		}
		if !f.ID.Valid() {
			slog.Warn("invalid spec command ID", slog.String("name", f.Name), slog.String("id", f.ID.Text()))
			continue
		}
		switch f.Direction {
		case matter.InterfaceServer:
			specCommandMap[f.ID.Value()] = f
		case matter.InterfaceClient:
			specResponseMap[f.ID.Value()] = f
		default:
			slog.Warn("invalid spec command direction", slog.String("name", f.Name), slog.Any("direction", f.Direction))
		}
	}

	zapCommandMap := make(map[uint64]*matter.Command)
	zapResponseMap := make(map[uint64]*matter.Command)
	for _, f := range zapCommands {
		if !f.ID.Valid() {
			slog.Warn("invalid ZAP command ID", slog.String("name", f.Name), slog.String("id", f.ID.Text()))
			continue
		}
		switch f.Direction {
		case matter.InterfaceServer:
			zapCommandMap[f.ID.Value()] = f
		case matter.InterfaceClient:
			zapResponseMap[f.ID.Value()] = f
		default:
			slog.Warn("invalid ZAP command direction", slog.String("name", f.Name), slog.Any("direction", f.Direction))
		}
	}
	diffs = append(diffs, compareCommandSets(specCommandMap, zapCommandMap)...)
	diffs = append(diffs, compareCommandSets(specResponseMap, zapResponseMap)...)
	return
}

func compareCommandSets(specCommandMap map[uint64]*matter.Command, zapCommandMap map[uint64]*matter.Command) (diffs []Diff) {
	for commandID, zapCommand := range zapCommandMap {
		specCommand, ok := specCommandMap[commandID]
		if !ok {
			continue
		}
		delete(zapCommandMap, commandID)
		delete(specCommandMap, commandID)
		commandDiffs := compareCommand(specCommand, zapCommand)
		if len(commandDiffs) > 0 {
			diffs = append(diffs, &IdentifiedDiff{Type: DiffTypeMismatch, Entity: types.EntityTypeCommand, ID: specCommand.ID, Name: specCommand.Name, Diffs: commandDiffs})
		}
	}
	for _, f := range specCommandMap {
		diffs = append(diffs, newMissingDiff(f.Name, types.EntityTypeCommand, SourceZAP))
	}
	for _, f := range zapCommandMap {
		diffs = append(diffs, newMissingDiff(f.Name, types.EntityTypeCommand, SourceSpec))
	}
	return
}
