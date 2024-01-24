package testplan

import (
	"fmt"
	"strings"

	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/conformance"
)

func renderFeatures(cluster *matter.Cluster, b *strings.Builder) {
	if cluster.Features != nil && len(cluster.Features.Bits) > 0 {
		b.WriteString("==== Features\n\n// FeatureMap defined macros\n")
		for _, bit := range cluster.Features.Bits {
			b.WriteString(fmt.Sprintf(":F_%s: %s\n", bit.Code, bit.Code))
		}
		b.WriteRune('\n')
		for i, bit := range cluster.Features.Bits {
			b.WriteString(fmt.Sprintf(":PICS_SF_%s: {PICS_S}.F%02d({F_%s})\n", bit.Code, i, bit.Code))
		}
		b.WriteRune('\n')
		b.WriteString("|===\n")
		b.WriteString("| *Variable* | *Description* | *Mandatory/Optional* | *Notes/Additional Constraints*\n")
		for _, bit := range cluster.Features.Bits {
			b.WriteString("| {PICS_SF_")
			b.WriteString(bit.Code)
			b.WriteString("} | {devsup} ")
			b.WriteString(bit.Summary)
			b.WriteString(" | ")
			renderConformance(b, cluster, cluster.Features, bit.Conformance, "{F_%s}")
			b.WriteString(" | \n")
		}
		b.WriteString("|===\n\n")
	}
}

func renderConformance(b *strings.Builder, cluster *matter.Cluster, features *matter.Bitmap, cs conformance.Set, featureFormat string) {
	if len(cs) == 0 {
		return
	}
	b.WriteString("{PICS_S}: ")
	for _, c := range cs {
		switch c := c.(type) {
		case *conformance.Mandatory:
			if c.Expression == nil {
				b.WriteString("M")
				continue
			}
			renderExpression(b, cluster, c.Expression, featureFormat)
		case *conformance.Optional:
			if c.Expression == nil {
				b.WriteString("O")
				if c.Choice != nil {
					b.WriteRune('.')
					b.WriteString(c.Choice.AsciiDocString())
				}
				continue
			}
			b.WriteRune('[')
			renderExpression(b, cluster, c.Expression, featureFormat)
			b.WriteRune(']')
			if c.Choice != nil {
				b.WriteRune('.')
				b.WriteString(c.Choice.AsciiDocString())
			}
		}
	}
}

func renderExpression(b *strings.Builder, cluster *matter.Cluster, exp conformance.Expression, featureFormat string) {
	switch exp := exp.(type) {
	case *conformance.EqualityExpression:
		b.WriteRune('(')
		renderExpression(b, cluster, exp.Left, featureFormat)
		b.WriteString(" == ")
		renderExpression(b, cluster, exp.Right, featureFormat)
		b.WriteRune(')')
	case *conformance.FeatureExpression:
		b.WriteString(renderIdentifier(cluster.Features, exp.ID, featureFormat))
	case *conformance.IdentifierExpression:
		b.WriteString("identifier")
	case *conformance.LogicalExpression:
		if exp.Not {
			b.WriteRune('!')
		}
		b.WriteRune('(')
		renderExpression(b, cluster, exp.Left, featureFormat)
		for _, e := range exp.Right {
			b.WriteString(" ")
			b.WriteString(exp.Operand)
			b.WriteString(" ")
			renderExpression(b, cluster, e, featureFormat)
		}
		b.WriteRune(')')
	}
}

func renderIdentifier(features *matter.Bitmap, id string, featureFormat string) string {
	if features == nil {
		return ""
	}
	for _, f := range features.Bits {
		if strings.EqualFold(f.Code, id) {
			return fmt.Sprintf(featureFormat, f.Code)
		}
	}
	return ""
}
