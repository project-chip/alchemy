package errdiff

import (
	"context"
	"log/slog"

	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
)

type ComparableError interface {
	spec.Error
	ComparableEntity() types.ComparableEntity
}

func Pipeline(cxt context.Context, baseRoot string, headRoot string, docPaths []string, pipelineOptions pipeline.ProcessingOptions) (violations map[string][]spec.Violation, err error) {
	specs, err := spec.LoadSpecPullRequest(cxt, baseRoot, headRoot, docPaths, pipelineOptions, nil)
	if err != nil {
		return
	}

	violations = ProcessComparison(&specs)
	return
}

func ProcessComparison(specs *spec.SpecPullRequest) (violations map[string][]spec.Violation) {
	slog.Info("Comparing head and base")
	v1 := compareErrors(specs.Base, specs.Head)

	slog.Info("Comparing head and base (in-progress)")
	v2 := compareErrors(specs.BaseInProgress, specs.HeadInProgress)

	violations = spec.MergeViolations(v1, v2)
	return
}

func groupErrorsByType(errors []spec.Error) map[spec.ErrorType][]spec.Error {
	errorBuckets := make(map[spec.ErrorType][]spec.Error)

	for _, err := range errors {
		errorType := err.Type()
		errorBuckets[errorType] = append(errorBuckets[errorType], err)
	}

	return errorBuckets
}

func findMatchingError(entity types.ComparableEntity, errors []error) error {
	for _, e := range errors {
		switch e := e.(type) {
		case ComparableError:
			ce := e.ComparableEntity()
			if ce != nil && entity.Equals(ce) {
				return e
			}
		}
	}
	return nil
}

func compareErrors(Base *spec.Specification, Head *spec.Specification) (violations map[string][]spec.Violation) {
	violations = make(map[string][]spec.Violation)

	groupedHeadErrors := groupErrorsByType(Head.Errors)
	groupedBaseErrors := groupErrorsByType(Base.Errors)
	for errorType, headErrors := range groupedHeadErrors {
		if baseErrors, ok := groupedBaseErrors[errorType]; ok {
			slog.Info("Comparing errors", "type", errorType, "headCount", len(headErrors), "baseCount", len(baseErrors))

			baseErrorsAsError := make([]error, len(baseErrors))
			for i, e := range baseErrors {
				baseErrorsAsError[i] = e
			}

			for _, headErr := range headErrors {
				if ce, ok := headErr.(ComparableError); ok {
					entityToFind := ce.ComparableEntity()
					if entityToFind == nil {
						continue
					}

					if matchingError := findMatchingError(entityToFind, baseErrorsAsError); matchingError == nil {
						slog.Error("This error is introduced by the current PR: <", headErr.Error(), ">")
						v := spec.Violation{Entity: entityToFind, Type: spec.ViolationNewParseError, Text: headErr.Error()}
						source, ok := entityToFind.(log.Source)
						if ok {
							v.Path, v.Line = source.Origin()
						}
						violations[v.Path] = append(violations[v.Path], v)
					}
				}
			}
		} else {
			for _, headErr := range headErrors {
				if ce, ok := headErr.(ComparableError); ok {
					entityToFind := ce.ComparableEntity()
					if entityToFind == nil {
						continue
					}
					slog.Error("This error is introduced by the current PR: <", headErr.Error(), ">")
					v := spec.Violation{Entity: entityToFind, Type: spec.ViolationNewParseError, Text: headErr.Error()}
					source, ok := entityToFind.(log.Source)
					if ok {
						v.Path, v.Line = source.Origin()
					}
					violations[v.Path] = append(violations[v.Path], v)
				}
			}
		}
	}

	return
}
