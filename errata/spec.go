package errata

import "path/filepath"

type Spec struct {
	IgnoreSections map[string]struct{}
}

func (spec *Spec) IgnoreSection(sectionName string) bool {
	if spec == nil {
		return false
	}
	if spec.IgnoreSections == nil {
		return false
	}
	_, ignore := spec.IgnoreSections[sectionName]
	return ignore
}

func GetSpec(path string) *Spec {
	errata, ok := Erratas[filepath.Base(path)]
	if !ok {
		return &DefaultErrata.Spec
	}
	return &errata.Spec
}
