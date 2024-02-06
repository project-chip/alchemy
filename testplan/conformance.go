package testplan

import (
	"fmt"
	"strings"

	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/conformance"
	"github.com/hasty/alchemy/matter/types"
)

type conformanceEntityFormatter func(entity types.Entity) string

func renderPicsConformance(b *strings.Builder, doc *ascii.Doc, cluster *matter.Cluster, cs conformance.Set) {
	if len(cs) == 0 {
		return
	}
	renderConformance(cs, b, doc, cluster, entityPICS)
}

func renderFeatureConformance(b *strings.Builder, doc *ascii.Doc, cluster *matter.Cluster, cs conformance.Set) {
	if len(cs) == 0 {
		return
	}
	b.WriteString("{PICS_S}: ")
	renderConformance(cs, b, doc, cluster, entityVariable)
}

func renderConformance(cs conformance.Set, b *strings.Builder, doc *ascii.Doc, cluster *matter.Cluster, formatter conformanceEntityFormatter) {
	for _, c := range cs {
		switch c := c.(type) {
		case *conformance.Mandatory:
			if c.Expression == nil {
				b.WriteString("M")
				continue
			}
			renderExpression(b, doc, cluster, c.Expression, formatter)
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
			renderExpression(b, doc, cluster, c.Expression, formatter)
			b.WriteRune(']')
			if c.Choice != nil {
				b.WriteRune('.')
				b.WriteString(c.Choice.AsciiDocString())
			}
		default:
			b.WriteString(fmt.Sprintf("unknown conformance: %T", c))
		}
	}
}

func renderExpression(b *strings.Builder, doc *ascii.Doc, cluster *matter.Cluster, exp conformance.Expression, formatter conformanceEntityFormatter) {
	switch exp := exp.(type) {
	case *conformance.EqualityExpression:
		b.WriteRune('(')
		renderExpression(b, doc, cluster, exp.Left, formatter)
		b.WriteString(" == ")
		renderExpression(b, doc, cluster, exp.Right, formatter)
		b.WriteRune(')')
	case *conformance.FeatureExpression:
		b.WriteString(renderIdentifier(cluster.Features, exp.Feature, formatter))
	case *conformance.IdentifierExpression:
		b.WriteString(renderIdentifier(cluster, exp.ID, formatter))
	case *conformance.ReferenceExpression:
		b.WriteString(renderReference(doc, exp.Reference, formatter))
	case *conformance.LogicalExpression:
		if exp.Not {
			b.WriteRune('!')
		}
		b.WriteRune('(')
		renderExpression(b, doc, cluster, exp.Left, formatter)
		for _, e := range exp.Right {
			b.WriteString(" ")
			b.WriteString(exp.Operand)
			b.WriteString(" ")
			renderExpression(b, doc, cluster, e, formatter)
		}
		b.WriteRune(')')
	default:
		b.WriteString(fmt.Sprintf("ERROR: unknown expression type: %T", exp))
	}
}

func renderIdentifier(store conformance.IdentifierStore, id string, formatter conformanceEntityFormatter) string {
	if store == nil {
		return ""
	}
	entity, ok := store.Identifier(id)
	if !ok {
		return fmt.Sprintf("UNKNOWN_ID_%s", id)
	}
	return formatter(entity)
}

func renderReference(store conformance.ReferenceStore, id string, formatter conformanceEntityFormatter) string {
	if store == nil {
		return ""
	}
	entity, ok := store.Reference(id)
	if !ok {
		return fmt.Sprintf("UNKNOWN_ID_%s", id)
	}
	return formatter(entity)
}
