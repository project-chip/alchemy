package spec

import (
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

func validateEvents(spec *Specification, cluster *matter.Cluster, events []*matter.Event) {
	idu := make(idUniqueness[*matter.Event])
	nu := make(nameUniqueness[*matter.Event])
	cv := make(conformanceValidation)
	for _, e := range events {
		cv.add(e, e.Conformance)
		validateEvent(spec, cluster, e, idu, nu)
	}
	cv.check(spec)
}

func validateEvent(spec *Specification, cluster *matter.Cluster, e *matter.Event, idu idUniqueness[*matter.Event], nu nameUniqueness[*matter.Event]) {
	idu.check(spec, e.ID, e)
	nu.check(spec, e)
	validateFields(spec, cluster, e, e.Fields, types.EntityTypeEventField)
}
