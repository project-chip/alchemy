package testplan

import (
	"fmt"
	"strings"

	"github.com/project-chip/alchemy/matter/spec"
)

func renderEvents(doc *spec.Doc, cut *clusterUnderTest, b *strings.Builder) {
	if len(cut.events) == 0 {
		return
	}
	b.WriteString("==== Events\n\n")
	names := make([]string, 0, len(cut.events))
	var longest int
	for _, event := range cut.events {
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
		b.WriteString(cut.events[i].Name)
		b.WriteRune('\n')
	}
	b.WriteRune('\n')
	for i, name := range names {
		b.WriteString(fmt.Sprintf(":PICS_S%-*s : {PICS_S}.E%s({%s})\n", longest, name, cut.events[i].ID.ShortHexString(), name))
	}
	b.WriteRune('\n')
	for i, name := range names {
		b.WriteString(fmt.Sprintf(":PICS_S%-*s_CONFORMANCE : {PICS_S}.E%s\n", longest, name, cut.events[i].ID.ShortHexString()))
	}
	b.WriteString("\n\n|===\n")
	b.WriteString("| *Variable* | *Description* | *Mandatory/Optional* | *Notes/Additional Constraints*\n")
	for i, event := range cut.events {
		name := names[i]
		b.WriteString(fmt.Sprintf("| {PICS_S%s} | {devimp} sending the _{%s}_ event?| ", name, name))
		if len(event.Conformance) > 0 {
			renderPicsConformance(b, doc, cut.cluster, event.Conformance)
		}
		b.WriteString(" |\n")
	}
	b.WriteString("|===\n\n")

}
