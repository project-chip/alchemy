package testplan

import (
	"fmt"
	"strings"

	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/matter"
)

func renderFeatures(doc *ascii.Doc, cluster *matter.Cluster, b *strings.Builder) {
	if cluster.Features != nil && len(cluster.Features.Bits) > 0 {
		b.WriteString("==== Features\n\n// FeatureMap defined macros\n")
		for _, bit := range cluster.Features.Bits {
			f := bit.(*matter.Feature)
			b.WriteString(fmt.Sprintf(":F_%s: %s\n", f.Code, f.Code))
		}
		b.WriteRune('\n')
		for i, bit := range cluster.Features.Bits {
			f := bit.(*matter.Feature)
			b.WriteString(fmt.Sprintf(":PICS_SF_%s: {PICS_S}.F%02d({F_%s})\n", f.Code, i, f.Code))
		}
		b.WriteRune('\n')
		b.WriteString("|===\n")
		b.WriteString("| *Variable* | *Description* | *Mandatory/Optional* | *Notes/Additional Constraints*\n")
		for _, bit := range cluster.Features.Bits {
			f := bit.(*matter.Feature)
			b.WriteString("| {PICS_SF_")
			b.WriteString(f.Code)
			b.WriteString("} | {devsup} ")
			b.WriteString(f.Summary())
			b.WriteString(" | ")
			renderFeatureConformance(b, doc, cluster, f.Conformance())
			b.WriteString(" | \n")
		}
		b.WriteString("|===\n\n\n")
	}
}
