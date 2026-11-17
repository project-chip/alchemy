package spec

import (
	"log/slog"

	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/types"
)

func validateConformance(spec *Specification, source types.Entity, c conformance.Conformance) {
	for c := range conformance.Traverse(c) {
		switch c := c.(type) {
		case *conformance.Generic:
			if conformance.IsBlank(c) {
				spec.addError(&InvalidConformanceError{Source: source, Conformance: c.RawText()})
			}
		default:
			for _, exp := range conformance.TraverseExpressions(c) {
				switch exp := exp.(type) {
				case *conformance.ReferenceExpression:
					if exp.Entity != nil {
						continue
					}
					slog.Error("Failed to resolve conformance expression reference", "ref", exp.Reference, log.Path("source", source))
					spec.addError(&UnknownConformanceReferenceError{Entity: source, Reference: exp.Reference})
				case *conformance.IdentifierExpression:
					if exp.Entity != nil {
						continue
					}
					slog.Error("Failed to resolve conformance expression identifier", "ref", exp.ID, log.Path("source", source))
					spec.addError(&UnknownConformanceIdentifierError{Entity: source, Identifier: exp.ID})
				//failure.finder.suggestIdentifiers(exp.ID, suggestions)
				//suggest.ListPossibilities(exp.ID, suggestions)
				case *conformance.ComparisonExpression:
					validateComparisonValue(spec, source, exp.Left)
					validateComparisonValue(spec, source, exp.Right)
				}
			}
		}
	}
}

func validateComparisonValue(spec *Specification, source types.Entity, cv conformance.ComparisonValue) {
	for cv := range conformance.TraverseComparisonValue(cv) {
		switch cv := cv.(type) {
		case *conformance.IdentifierValue:
			if cv.Entity != nil {
				continue
			}
			spec.addError(&UnknownConformanceIdentifierError{Entity: source, Identifier: cv.ID})
			//suggestions := make(map[types.Entity]int)
			//failure.finder.suggestIdentifiers(exp.ID, suggestions)
			//suggest.ListPossibilities(exp.ID, suggestions)
		case *conformance.ReferenceValue:
			if cv.Entity != nil {
				continue
			}
			slog.Error("failed to resolve conformance value reference", "ref", cv.Reference, log.Path("source", source))
			spec.addError(&UnknownConformanceReferenceError{Entity: source, Reference: cv.Reference})
		}
	}
}
