package errata

import "path/filepath"

type Errata struct {
	Disco    Disco    `yaml:"disco,omitempty"`
	Spec     Spec     `yaml:"spec,omitempty"`
	TestPlan TestPlan `yaml:"test-plan,omitempty"`
	ZAP      ZAP      `yaml:"zap,omitempty"`
}

var DefaultErrata = &Errata{}

func GetErrata(path string) *Errata {
	path = filepath.ToSlash(path)
	errata, ok := Erratas[path]
	if ok {
		return errata
	}
	return DefaultErrata
}

var Erratas = map[string]*Errata{}
