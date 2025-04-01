package testscript

import (
	"fmt"

	"github.com/goccy/go-yaml"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
	"github.com/project-chip/alchemy/testplan/pics"
)

type Test struct {
	ID   string
	Name string

	Cluster *matter.Cluster

	GlobalVariableNames []string
	GlobalVariables     map[string]types.Entity
	Steps               []*TestStep

	StructChecks []*TestStep

	PICSList      []pics.Expression
	YamlVariables yaml.MapSlice
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

	Entity types.Entity
}
