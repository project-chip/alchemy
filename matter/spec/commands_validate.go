package spec

import (
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

func validateCommands(spec *Specification, cluster *matter.Cluster, c []*matter.Command) {
	idsu := make(idUniqueness[*matter.Command])
	idcu := make(idUniqueness[*matter.Command])
	nu := make(nameUniqueness[*matter.Command])
	cv := make(conformanceValidation)

	for _, command := range c {
		switch command.Direction {
		case matter.InterfaceServer:
			idsu.check(spec, command.ID, command)
		case matter.InterfaceClient:
			idcu.check(spec, command.ID, command)
		}
		nu.check(spec, command)
		cv.add(command, command.Conformance)
		validateAccess(spec, command, command.Access)
		validateConformance(spec, command, command.Conformance)
		validateQuality(spec, command.Quality, types.EntityTypeCommand, command)
		validateFields(spec, cluster, command, command.Fields, types.EntityTypeCommandField)
	}

	cv.check(spec)
}
