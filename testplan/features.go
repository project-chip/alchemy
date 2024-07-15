package testplan

import (
	"fmt"
	"strings"

	"github.com/project-chip/alchemy/matter/spec"
)

func renderFeatures(doc *spec.Doc, cut *clusterUnderTest, b *strings.Builder) {
	if cut.features == nil || len(cut.features) == 0 {
		return
	}

	b.WriteString("==== Features\n\n// FeatureMap defined macros\n")
	for _, f := range cut.features {
		b.WriteString(fmt.Sprintf(":F_%s: %s\n", f.Code, f.Code))
	}
	b.WriteRune('\n')
	for i, f := range cut.features {
		b.WriteString(fmt.Sprintf(":PICS_SF_%s: {PICS_S}.F%02d({F_%s})\n", f.Code, i, f.Code))
	}
	b.WriteRune('\n')
	b.WriteString("|===\n")
	b.WriteString("| *Variable* | *Description* | *Mandatory/Optional* | *Notes/Additional Constraints*\n")
	for _, f := range cut.features {
		b.WriteString("| {PICS_SF_")
		b.WriteString(f.Code)
		b.WriteString("} | {devsup} ")
		b.WriteString(f.Summary())
		b.WriteString(" | ")
		renderFeatureConformance(b, doc, cut.cluster, f.Conformance())
		b.WriteString(" | \n")
	}
	b.WriteString("|===\n\n\n")

}
