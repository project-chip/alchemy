package testplan

import (
	"fmt"
	"strings"

	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/spec"
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
		b.WriteString(fmt.Sprintf(":PICS_S%-*s : {PICS_S}.A%04x({%s})\n", longest, name, i, name))
	}
	b.WriteString("\n\n|===\n")
	b.WriteString("| *Variable* | *Description* | *Mandatory/Optional* | *Notes/Additional Constraints*\n")
	for i, event := range cluster.Events {
		name := names[i]
		b.WriteString(fmt.Sprintf("| {PICS_S%s} | {devimp} sending the _{%s}_ event?| ", name, name))
		if len(event.Conformance) > 0 {
			b.WriteString("{PICS_S}: ")
			renderPicsConformance(b, doc, cluster, event.Conformance)
		}
		b.WriteString(" |\n")
	}
	b.WriteString("|===\n\n")

}
