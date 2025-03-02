package testplan

import (
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/testplan/pics"
	"github.com/project-chip/alchemy/yaml2python/parse"
)

type Test struct {
	parse.Test
	ID       string
	Cluster  *matter.Cluster
	PICSList []pics.Expression

	Groups []*Group

	Variables     []string
	PICSAliases   map[string]string
	PICSAliasList [][]*PicsAlias
}

type Group struct {
	Parent *Test

	Name        string
	Description string
	Steps       []*Step
}

type Step struct {
	parse.TestStep
	Parent *Group

	Description      []string
	UserVerification []string
	PICSSet          pics.Expression

	Label string
	PICS  string
}
