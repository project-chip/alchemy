package render

import (
	"fmt"
	"strings"

	"github.com/mailgun/raymond/v2"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/types"
)

type conformanceEntityFormatter func(entity types.Entity) string

func picsConformanceHelper(cluster matter.Cluster, cs conformance.Set) raymond.SafeString {
	var b strings.Builder
	renderConformance(cs, &b, &cluster, entityPICS)
	return raymond.SafeString(b.String())
}

func conformanceHelper(cluster matter.Cluster, cs conformance.Set) raymond.SafeString {
	var b strings.Builder
	renderConformance(cs, &b, &cluster, entityVariable)
	return raymond.SafeString(b.String())
}

func renderChoice(c *conformance.Optional, b *strings.Builder) {
	// PICS tool does not support + style conformances, so unless this is a "pick one" choice,
	//render as fully optional, we'll check the choice conformance properly in the tests.
	o := conformance.ChoiceExactLimit{Limit: 1}
	if c.Choice != nil && o.Equal(c.Choice.Limit) {
		b.WriteRune('.')
		b.WriteString(c.Choice.ASCIIDocString())
	}
}

func renderConformance(cs conformance.Set, b *strings.Builder, cluster *matter.Cluster, formatter conformanceEntityFormatter) {
	// PICS tool can't handle otherwise conformances, so render anything with an otherwise conformance as optional for the purposes of the
	// test plan PICS. This can be fully evaluated in the tests.
	// The only exception is if it is provisional, which should be rendered as X.
	if len(cs) != 1 {
		switch cs[0].(type) {
		case *conformance.Provisional:
			b.WriteRune('X')
		default:
			b.WriteString("{PICS_S}: O")
		}
		return
	}
	switch c := cs[0].(type) {
	case *conformance.Mandatory:
		if c.Expression == nil {
			b.WriteString("{PICS_S}: M")
			return
		}
		renderExpression(b, cluster, c.Expression, formatter)
	case *conformance.Optional:
		if c.Expression == nil {
			b.WriteString("{PICS_S}: O")
			renderChoice(c, b)
			return
		}
		renderExpression(b, cluster, c.Expression, formatter)
		b.WriteString(": O")
		renderChoice(c, b)
	case *conformance.Provisional:
		b.WriteRune('X')
	case *conformance.Disallowed:
		b.WriteRune('X')
	case *conformance.Deprecated:
		b.WriteRune('X')
	case *conformance.Described:
		b.WriteString("{PICS_S}: O")
	default:
		b.WriteString(fmt.Sprintf("unknown conformance: %T", c))
	}
}

func renderExpression(b *strings.Builder, cluster *matter.Cluster, exp conformance.Expression, formatter conformanceEntityFormatter) {
	switch exp := exp.(type) {
	case *conformance.EqualityExpression:
		b.WriteRune('(')
		renderExpression(b, cluster, exp.Left, formatter)
		b.WriteString(" == ")
		renderExpression(b, cluster, exp.Right, formatter)
		b.WriteRune(')')
	case *conformance.IdentifierExpression:
		b.WriteString(renderIdentifier(exp, formatter))
	case *conformance.ReferenceExpression:
		b.WriteString(renderReference(exp, formatter))
	case *conformance.LogicalExpression:
		if exp.Not {
			b.WriteString("NOT")
		}
		b.WriteRune('(')
		renderExpression(b, cluster, exp.Left, formatter)
		for _, e := range exp.Right {
			b.WriteString(" ")
			b.WriteString(exp.Operand)
			b.WriteString(" ")
			renderExpression(b, cluster, e, formatter)
		}
		b.WriteRune(')')
	case *conformance.ComparisonExpression:
		b.WriteRune('(')
		b.WriteString(exp.Left.ASCIIDocString())
		b.WriteString(" ")
		b.WriteString(exp.Op.String())
		b.WriteString(" ")
		b.WriteString(exp.Right.ASCIIDocString())
		b.WriteRune(')')
	default:
		b.WriteString(fmt.Sprintf("ERROR: unknown expression type: %T", exp))
	}
}

func renderIdentifier(id *conformance.IdentifierExpression, formatter conformanceEntityFormatter) string {
	if id.Entity == nil {
		return fmt.Sprintf("UNKNOWN_ID_%s", id.ID)
	}
	return formatter(id.Entity)
}

func renderReference(ref *conformance.ReferenceExpression, formatter conformanceEntityFormatter) string {
	if ref.Entity == nil {
		return fmt.Sprintf("UNKNOWN_ID_%s", ref.Reference)
	}
	return formatter(ref.Entity)
}
