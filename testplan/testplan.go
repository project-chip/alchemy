package testplan

import (
	"path/filepath"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

func buildDestinations(testplanRoot string, entities []types.Entity, errata errata.TestPlan) (destinations map[string]*matter.Cluster) {
	destinations = make(map[string]*matter.Cluster)

	for _, e := range entities {
		switch e := e.(type) {
		case *matter.ClusterGroup:
			for _, c := range e.Clusters {
				fileName := strings.ToLower(strcase.ToSnake(c.Name))
				newPath := getTestPlanPath(testplanRoot, fileName)
				destinations[newPath] = c
			}
		case *matter.Cluster:
			var newPath string
			if errata.TestPlanPath != "" {
				newPath = filepath.Join(testplanRoot, errata.TestPlanPath)
			} else {
				newPath = getTestPlanPath(testplanRoot, strings.ToLower(strcase.ToSnake(e.Name)))
			}
			destinations[newPath] = e
		}
	}
	return

}
