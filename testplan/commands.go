package testplan

import (
	"fmt"
	"strings"

	"github.com/project-chip/alchemy/matter/spec"
)

func renderCommands(doc *spec.Doc, cut *clusterUnderTest, b *strings.Builder) {
	renderCommandsAccepted(doc, cut, b)
	renderCommandsGenerated(doc, cut, b)
}

func renderCommandsAccepted(doc *spec.Doc, cut *clusterUnderTest, b *strings.Builder) {
	if len(cut.commandsAccepted) == 0 {
		return
	}

	names := make([]string, 0, len(cut.commandsAccepted))
	var longest int
	for _, a := range cut.commandsAccepted {
		name := entityIdentifier(a)
		if len(name) > longest {
			longest = len(name)
		}
		names = append(names, name)
	}
	b.WriteString("==== Commands received\n\n\n")
	for i, name := range names {
		b.WriteString(":")
		b.WriteString(fmt.Sprintf("%-*s", longest, name))
		b.WriteString(" : ")
		b.WriteString(cut.commandsAccepted[i].Name)
		b.WriteRune('\n')
	}
	b.WriteString("\n\n")
	for i, name := range names {
		b.WriteString(fmt.Sprintf(":PICS_S%-*s : {PICS_S}.C%02X.Rsp({%s})\n", longest, name, cut.commandsAccepted[i].ID.Value(), name))
	}
	b.WriteString("\n\n|===\n")
	b.WriteString("| *Variable* | *Description* | *Mandatory/Optional* | *Notes/Additional Constraints*\n")
	for i, a := range cut.commandsAccepted {
		name := names[i]
		b.WriteString(fmt.Sprintf("| {PICS_S%s} | {devimp} the _{%s}_ command?| ", name, name))
		if len(a.Conformance) > 0 {
			b.WriteString("{PICS_S}: ")
			renderPicsConformance(b, doc, cut.cluster, a.Conformance)
		}
		b.WriteString(" |\n")
	}
	b.WriteString("|===\n\n")
}

func renderCommandsGenerated(doc *spec.Doc, cut *clusterUnderTest, b *strings.Builder) {
	if len(cut.commandsGenerated) == 0 {
		return
	}

	names := make([]string, 0, len(cut.commandsGenerated))
	var longest int
	for _, a := range cut.commandsGenerated {
		name := entityIdentifier(a)
		if len(name) > longest {
			longest = len(name)
		}
		names = append(names, name)
	}
	b.WriteString("==== Commands generated\n\n\n")
	for i, name := range names {
		b.WriteString(":")
		b.WriteString(fmt.Sprintf("%-*s", longest, name))
		b.WriteString(" : ")
		b.WriteString(cut.commandsGenerated[i].Name)
		b.WriteRune('\n')
	}
	b.WriteString("\n\n")
	for i, name := range names {
		b.WriteString(fmt.Sprintf(":PICS_S%-*s : {PICS_S}.C%02X.Tx({%s})\n", longest, name, cut.commandsGenerated[i].ID.Value(), name))
	}
	b.WriteString("\n\n|===\n")
	b.WriteString("| *Variable* | *Description* | *Mandatory/Optional* | *Notes/Additional Constraints*\n")
	for i, a := range cut.commandsGenerated {
		name := names[i]
		b.WriteString(fmt.Sprintf("| {PICS_S%s} | {devimp} sending the _{%s}_ command?| ", name, name))
		if len(a.Conformance) > 0 {
			b.WriteString("{PICS_S}: ")
			renderPicsConformance(b, doc, cut.cluster, a.Conformance)
		}
		b.WriteString(" |\n")
	}
	b.WriteString("|===\n\n")
}
