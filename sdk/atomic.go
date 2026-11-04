package sdk

import (
	"log/slog"

	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
)

func resolveAtomicOperations(spec *spec.Specification) {
	for m := range spec.DocRefs {
		switch v := m.(type) {
		case *matter.ClusterGroup:
			for _, cl := range v.Clusters {
				if hasAtomicAttributes(cl) {
					addAtomicOperations(spec, cl)
				}
			}
		case *matter.Cluster:
			if hasAtomicAttributes(v) {
				addAtomicOperations(spec, v)
			}
		}
	}
}

func hasAtomicAttributes(cluster *matter.Cluster) bool {
	for _, f := range cluster.Attributes {
		if f.Quality.Has(matter.QualityAtomicWrite) {
			return true
		}
	}
	return false
}

func addAtomicOperations(spec *spec.Specification, cluster *matter.Cluster) {
	var atomicRequest, atomicResponse *matter.Command
	for o := range spec.GlobalObjects {
		switch o := o.(type) {
		case *matter.Command:
			switch o.Name {
			case "AtomicRequest":
				atomicRequest = o
			case "AtomicResponse":
				atomicResponse = o
			}
		}
	}

	if atomicRequest == nil {
		slog.Warn("Could not find AtomicRequest command")
		return
	}
	if atomicResponse == nil {
		slog.Warn("Could not find AtomicResponse command")
		return
	}
	cluster.Commands = append(cluster.Commands, atomicRequest)
	cluster.Commands = append(cluster.Commands, atomicResponse)
	spec.ClusterRefs.Add(cluster, atomicRequest)
	spec.ClusterRefs.Add(cluster, atomicResponse)
}
