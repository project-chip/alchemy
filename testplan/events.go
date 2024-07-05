package testplan

import (
	"fmt"
	"strings"

	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
)

func renderEvents(doc *spec.Doc, cluster *matter.Cluster, b *strings.Builder) {
	if len(cluster.Events) == 0 {
		return
	}
	b.WriteString("==== Events\n\n")
	names := make([]string, 0, len(cluster.Events))
	var longest int
	for _, event := range cluster.Events {
		name := entityIdentifier(event)
		if len(name) > longest {
			longest = len(name)
		}
		names = append(names, name)
	}
	for i, name := range names {
		b.WriteString(":")
		b.WriteString(fmt.Sprintf("%-*s", longest, name))
		b.WriteString(": ")
		b.WriteString(cluster.Events[i].Name)
		b.WriteRune('\n')
	}
	b.WriteRune('\n')
	for i, name := range names {
		b.WriteString(fmt.Sprintf(":PICS_S%-*s : {PICS_S}.E%02x({%s})\n", longest, name, i, name))
	}
	b.WriteRune('\n')
	for i, name := range names {
		b.WriteString(fmt.Sprintf(":PICS_S%-*s_CONFORMANCE : {PICS_S}.E%02x\n", longest, name, i))
	}
	b.WriteString("\n\n|===\n")
	b.WriteString("| *Variable* | *Description* | *Mandatory/Optional* | *Notes/Additional Constraints*\n")
	for i, event := range cluster.Events {
		name := names[i]
		b.WriteString(fmt.Sprintf("| {PICS_S%s} | {devimp} sending the _{%s}_ event?| ", name, name))
		if len(event.Conformance) > 0 {
			renderPicsConformance(b, doc, cluster, event.Conformance)
		}
		b.WriteString(" |\n")
	}
	b.WriteString("|===\n\n")

}
