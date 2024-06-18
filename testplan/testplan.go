package testplan

import (
	"strings"

	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/types"
	"github.com/iancoleman/strcase"
)

func buildDestinations(sdkRoot string, entities []types.Entity) (destinations map[string]*matter.Cluster) {
	destinations = make(map[string]*matter.Cluster)

	for _, e := range entities {
		switch e := e.(type) {
		case *matter.ClusterGroup:
			for _, c := range e.Clusters {
				fileName := strings.ToLower(strcase.ToSnake(c.Name))
				newPath := getTestPlanPath(sdkRoot, fileName)
				destinations[newPath] = c
			}
		case *matter.Cluster:
			fileName := strings.ToLower(strcase.ToSnake(e.Name))
			newPath := getTestPlanPath(sdkRoot, fileName)
			destinations[newPath] = e
		}
	}
	return

}
