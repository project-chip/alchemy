package errata

type Errata struct {
	Disco    Disco    `yaml:"disco,omitempty"`
	Spec     Spec     `yaml:"spec,omitempty"`
	TestPlan TestPlan `yaml:"test-plan,omitempty"`
	SDK      SDK      `yaml:"sdk,omitempty"`
}

var DefaultErrata = &Errata{}

type Collection struct {
	errata   map[string]*Errata
	docRoots []string
}

func (c *Collection) Get(path string) *Errata {
	errata, ok := c.errata[path]
	if ok {
		return errata
	}
	return DefaultErrata
}

var defaultDocRoots = []string{"src/main.adoc",
	"src/appclusters.adoc",
	"src/device_library.adoc",
	"src/standard_namespaces.adoc"}

func (c *Collection) DocRoots() []string {
	if len(c.docRoots) == 0 {
		return defaultDocRoots
	}
	return c.docRoots
}
