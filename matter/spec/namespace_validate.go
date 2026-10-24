package spec

import "github.com/project-chip/alchemy/matter"

func validateNamespaces(spec *Specification) {
	iu := make(idUniqueness[*matter.Namespace])
	nu := make(nameUniqueness[*matter.Namespace])
	for _, ns := range spec.Namespaces {
		iu.check(spec, ns.ID, ns)
		nu.check(spec, ns)
		stiu := make(idUniqueness[*matter.SemanticTag])
		stnu := make(nameUniqueness[*matter.SemanticTag])
		for _, tag := range ns.SemanticTags {
			stiu.check(spec, tag.ID, tag)
			stnu.check(spec, tag)
		}
	}
}
