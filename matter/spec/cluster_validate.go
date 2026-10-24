package spec

import "github.com/project-chip/alchemy/matter"

func validateClusters(spec *Specification) {
	idu := make(idUniqueness[*matter.Cluster])
	nu := make(nameUniqueness[*matter.Cluster])
	for c := range spec.Clusters {
		idu.check(spec, c.ID, c)
		nu.check(spec, c)
		validateCluster(spec, c)
	}
}

func validateCluster(spec *Specification, c *matter.Cluster) {
	if c.Hierarchy != "Base" && c.ParentCluster == nil {
		spec.addError(&UnknownBaseClusterError{Cluster: c})
	}
	if c.Features != nil {
		validateBitmap(spec, &c.Features.Bitmap)
	}
	validateAttributes(spec, c)
	validateCommands(spec, c, c.Commands)
	validateEvents(spec, c, c.Events)
	validateBitmaps(spec, c.Bitmaps)
	validateEnums(spec, c.Enums)
	validateStructs(spec, c, c.Structs)
}
