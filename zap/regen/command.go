package regen

import (
	"slices"

	"github.com/mailgun/raymond/v2"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
	"github.com/project-chip/alchemy/zap"
)

func commandsHelper(spec *spec.Specification) func(commands matter.CommandSet, options *raymond.Options) raymond.SafeString {
	return func(commands matter.CommandSet, options *raymond.Options) raymond.SafeString {
		sortedCommands := make(matter.CommandSet, len(commands))
		copy(sortedCommands, commands)
		slices.SortStableFunc(sortedCommands, func(a *matter.Command, b *matter.Command) int {
			cmp := a.ID.Compare(b.ID)
			if cmp != 0 {
				return cmp
			}
			if a.Direction == matter.InterfaceServer {
				return 1
			} else {
				return 0
			}
		})
		/*var requests []*matter.Command
		responses := make(map[*matter.Command]struct{})
		for _, command := range commands {
			switch command.Direction {
			case matter.InterfaceServer:
				requests = append(requests, command)
			case matter.InterfaceClient:
				responses[command] = struct{}{}
			}
		}
		slices.SortStableFunc(requests, func(a *matter.Command, b *matter.Command) int {
			return a.ID.Compare(b.ID)
		})
		for _, req := range requests {
			sortedCommands = append(sortedCommands, req)
			if req.Response != nil {
				switch response := req.Response.Entity.(type) {
				case *matter.Command:
					if _, unused := responses[response]; unused {
						sortedCommands = append(sortedCommands, response)
						delete(responses, response)
					}
				case nil:
				}
			}
		}*/
		return enumerateEntitiesHelper(sortedCommands, spec, options)
	}

}

func commandFieldsHelper(spec *spec.Specification) func(matter.Command, *raymond.Options) raymond.SafeString {
	return func(cmd matter.Command, options *raymond.Options) raymond.SafeString {
		fields := filterFields(cmd.Fields)
		if cmd.Access.FabricSensitivity == matter.FabricSensitivitySensitive {
			fields = append(fields, &matter.Field{ID: matter.NewNumber(254), Name: "FabricIndex", Type: types.NewDataType(types.BaseDataTypeFabricIndex, false), Conformance: conformance.Set{&conformance.Mandatory{}}})
		}
		return enumerateEntitiesHelper(fields, spec, options)
	}
}

func requestsHelper(spec *spec.Specification) func(commands matter.CommandSet, options *raymond.Options) raymond.SafeString {
	return func(commands matter.CommandSet, options *raymond.Options) raymond.SafeString {
		var requests []*matter.Command
		for _, command := range commands {
			if conformance.IsZigbee(command.Conformance) || zap.IsDisallowed(command, command.Conformance) {
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
		return enumerateEntitiesHelper(requests, spec, options)
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
