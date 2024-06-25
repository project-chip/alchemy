package testplan

import (
	"strings"

	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
)

func renderServer(doc *spec.Doc, cluster *matter.Cluster, b *strings.Builder) (err error) {

	b.WriteString("=== Server\n\n")
	renderFeatures(doc, cluster, b)
	renderAttributes(doc, cluster, b)
	renderEvents(doc, cluster, b)

	return
}
