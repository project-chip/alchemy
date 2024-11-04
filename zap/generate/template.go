package generate

import (
	"fmt"
	"time"

	"github.com/beevik/etree"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/zap"
)

func newZapTemplate() (x *etree.Document) {
	x = etree.NewDocument()

	x.CreateProcInst("xml", `version="1.0"`)
	x.CreateComment(fmt.Sprintf(license, time.Now().Year()))
	return
}

func (tg *TemplateGenerator) renderZapTemplate(configurator *zap.Configurator, x *etree.Document) (result string, err error) {

	var exampleCluster *matter.Cluster
	for c := range configurator.Clusters {
		if c != nil {
			if exampleCluster == nil {
				exampleCluster = c
			}

			if len(configurator.Errata.ClusterAliases) > 0 {
				if aliases, ok := configurator.Errata.ClusterAliases[c.Name]; ok {
					tg.ClusterAliases.Store(c.Name, aliases)
				}
			}
		}
	}

	cr := newConfiguratorRenderer(tg, configurator)
	return cr.render(x, exampleCluster)
}
