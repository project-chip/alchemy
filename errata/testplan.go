package errata

type TestPlan struct {
	TestPlanPath  string                  `yaml:"testplan-path,omitempty"`
	TestPlanPaths map[string]TestPlanPath `yaml:"testplan-paths,omitempty"`
}

type TestPlanPath struct {
	Path string `yaml:"path,omitempty"`
}
