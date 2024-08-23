package errata

type Spec struct {
	IgnoreSections map[string]Purpose `yaml:"ignore-sections,omitempty"`
}

func (spec *Spec) IgnoreSection(sectionName string, purpose Purpose) bool {
	if spec == nil {
		return false
	}
	if spec.IgnoreSections == nil {
		return false
	}
	if p, ok := spec.IgnoreSections[sectionName]; ok {
		return (p & purpose) != PurposeNone
	}
	return false
}

func GetSpec(path string) *Spec {
	e := GetErrata(path)
	return &e.Spec
}
