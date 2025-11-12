package errata

import "iter"

type Errata struct {
	Disco    Disco    `yaml:"disco,omitempty"`
	Spec     Spec     `yaml:"spec,omitempty"`
	TestPlan TestPlan `yaml:"test-plan,omitempty"`
	SDK      SDK      `yaml:"sdk,omitempty"`
}

var DefaultErrata = &Errata{}

type Collection struct {
	errata map[string]*Errata
}

func (c *Collection) Get(path string) *Errata {
	errata, ok := c.errata[path]
	if ok {
		return errata
	}
	return DefaultErrata
}

func (c *Collection) All() iter.Seq2[string, *Errata] {
	return func(yield func(string, *Errata) bool) {
		for path, errata := range c.errata {
			if !yield(path, errata) {
				return
			}
		}
	}
}
