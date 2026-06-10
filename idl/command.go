package idl

import (
	"slices"
	"strings"

	"github.com/mailgun/raymond/v2"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
)

func commandsHelper(spec *spec.Specification, filter ProvisionalFilter) func(commands matter.CommandSet, options *raymond.Options) raymond.SafeString {
	return func(commands matter.CommandSet, options *raymond.Options) raymond.SafeString {
		var sortedCommands matter.CommandSet
		for _, cmd := range commands {
			if !entityShouldBeIncluded(spec, filter, cmd) {
				continue
			}
			sortedCommands = append(sortedCommands, cmd)
		}
		serverCommandIDs := make(map[string]*matter.Number)
		for _, c := range commands {
			if c.Direction == matter.InterfaceServer && c.Response != nil && c.Response.Name != "" {
				serverCommandIDs[c.Response.Name] = c.ID
			}
		}
		slices.SortStableFunc(sortedCommands, func(a *matter.Command, b *matter.Command) int {
			aID := getSortID(a, serverCommandIDs)
			bID := getSortID(b, serverCommandIDs)
			cmp := aID.Compare(bID)
			if cmp != 0 {
				return cmp
			}
			return strings.Compare(a.Name, b.Name)
		})
		return enumerateEntitiesHelper(sortedCommands, spec, filter, options)
	}
}

func getSortID(cmd *matter.Command, serverCommandIDs map[string]*matter.Number) *matter.Number {
	if cmd.Direction == matter.InterfaceClient {
		if id, ok := serverCommandIDs[cmd.Name]; ok {
			return id
		}
	}
	return cmd.ID
}

func commandFieldsHelper(spec *spec.Specification, filter ProvisionalFilter) func(matter.Command, *raymond.Options) raymond.SafeString {
	return func(cmd matter.Command, options *raymond.Options) raymond.SafeString {
		fields := filterEntities(spec, filter, cmd.Fields)
		if cmd.Access.FabricSensitivity == matter.FabricSensitivitySensitive {
			fields = append(fields, &matter.Field{ID: matter.NewNumber(254), Name: "FabricIndex", Type: types.NewDataType(types.BaseDataTypeFabricIndex, types.DataTypeRankScalar), Conformance: conformance.Set{&conformance.Mandatory{}}})
		}
		slices.SortStableFunc(fields, func(a *matter.Field, b *matter.Field) int {
			return a.ID.Compare(b.ID)
		})
		return enumerateEntitiesHelper(fields, spec, filter, options)
	}
}

func requestsHelper(spec *spec.Specification, filter ProvisionalFilter) func(commands matter.CommandSet, options *raymond.Options) raymond.SafeString {
	return func(commands matter.CommandSet, options *raymond.Options) raymond.SafeString {
		var requests []*matter.Command
		for _, command := range commands {
			if !entityShouldBeIncluded(spec, filter, command) {
				continue
			}
			switch command.Direction {
			case matter.InterfaceServer:
				requests = append(requests, command)
			}
		}
		slices.SortStableFunc(requests, func(a *matter.Command, b *matter.Command) int {
			return a.ID.Compare(b.ID)
		})
		return enumerateEntitiesHelper(requests, spec, filter, options)
	}
}

func isTimedHelper(command matter.Command, options *raymond.Options) string {
	if command.Access.Timing == matter.TimingTimed {
		return options.Fn()
	} else {
		return options.Inverse()
	}
}

func requestNameHelper(command matter.Command) raymond.SafeString {
	return raymond.SafeString(command.Name)
}

func responseNameHelper(command matter.Command) raymond.SafeString {
	if command.Response != nil {
		switch response := command.Response.Entity.(type) {
		case *matter.Command:
			return raymond.SafeString(response.Name)
		default:
		}
	}
	return raymond.SafeString("DefaultSuccess")
}
