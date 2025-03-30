package testscript

import (
	"fmt"

	"github.com/project-chip/alchemy/matter"
)

type Test struct {
	ID   string
	Name string

	Cluster *matter.Cluster

	GlobalVariables []string
	Steps           []*TestStep

	StructChecks []*TestStep
}

func (t *Test) AddStep(s *TestStep) {
	s.Parent = t
	t.Steps = append(t.Steps, s)
	s.Name = fmt.Sprintf("%d", len(t.Steps))
}

type TestStep struct {
	Parent *Test

	Name        string
	Description string

	Cluster *matter.Cluster
	Actions []TestAction
}
