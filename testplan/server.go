package testplan

import (
	"strings"

	"github.com/hasty/alchemy/matter"
)

func renderServer(cluster *matter.Cluster, b *strings.Builder) (err error) {

	b.WriteString("=== Server\n\n")
	renderFeatures(cluster, b)
	renderAttributes(cluster, b)
	return
}
