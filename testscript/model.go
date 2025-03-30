package testscript

import (
	"fmt"

	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/testplan/pics"
)

type Test struct {
	ID   string
	Name string

	Cluster *matter.Cluster

	GlobalVariables []string
	Steps           []*TestStep

	StructChecks []*TestStep

	PICSList []pics.Expression
}

func (t *Test) AddStep(s *TestStep) {
	s.Parent = t
	t.Steps = append(t.Steps, s)
	if s.Name == "" {
		s.Name = fmt.Sprintf("%d", len(t.Steps))
	}
}

type TestStep struct {
	Parent *Test

	Name        string
	Description string

	Cluster *matter.Cluster
	Actions []TestAction
}
