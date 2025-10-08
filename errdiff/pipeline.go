package errdiff

import (
	"context"
	"errors"
	"log/slog"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
)

type ComparableError interface {
	spec.Error
	ComparableEntity() types.ComparableEntity
}

type specs struct {
	Base           *spec.Specification
	BaseInProgress *spec.Specification
	Head           *spec.Specification
	HeadInProgress *spec.Specification
}

func Pipeline(cxt context.Context, baseRoot string, headRoot string, docPaths []string, pipelineOptions pipeline.ProcessingOptions) (err error) {
	var specs specs

	// Read Head specs
	specs.Head, specs.HeadInProgress, err = loadSpecs(cxt, pipelineOptions, headRoot)
	if err != nil {
		return
	}
	slog.Info("cluster count head", "count", len(specs.Head.Clusters))
	slog.Info("cluster count head in-progress", "count", len(specs.HeadInProgress.Clusters))

	// Read Base specs
	specs.Base, specs.BaseInProgress, err = loadSpecs(cxt, pipelineOptions, baseRoot)
	if err != nil {
		return
	}
	slog.Info("cluster count base", "count", len(specs.Base.Clusters))
	slog.Info("cluster count base in-progress", "count", len(specs.BaseInProgress.Clusters))

	// Compare the head and base
	slog.Info("Comparing head and base")
	err1 := compare(specs.Base, specs.Head)

	// Compare the in-progress head and base
	slog.Info("Comparing head and base (in-progress)")
	err2 := compare(specs.BaseInProgress, specs.HeadInProgress)

	err = errors.Join(err1, err2)

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

func loadSpecs(cxt context.Context, pipelineOptions pipeline.ProcessingOptions, specRoot string) (baseSpec *spec.Specification, inProgressSpec *spec.Specification, err error) {
	parserOptions := spec.ParserOptions{Root: specRoot}
	baseSpec, _, err = spec.Parse(cxt, parserOptions, pipelineOptions, nil, []asciidoc.AttributeName{})

	if err != nil {
		return
	}
	inProgressSpec, _, err = spec.Parse(cxt, parserOptions, pipelineOptions, nil, []asciidoc.AttributeName{"in-progress"})
	return
}

func compare(Base *spec.Specification, Head *spec.Specification) (err error) {
	groupedHeadErrors := groupErrorsByType(Head.Errors)
	groupedBaseErrors := groupErrorsByType(Base.Errors)
	for errorType, headErrors := range groupedHeadErrors {
		if baseErrors, ok := groupedBaseErrors[errorType]; ok {
			slog.Info("Comparing errors", "type", errorType, "headCount", len(headErrors), "baseCount", len(baseErrors))

			for _, headErr := range headErrors {
				if ce, ok := headErr.(ComparableError); ok {
					entityToFind := ce.ComparableEntity()
					if entityToFind == nil {
						continue
					}

					baseErrorsAsError := make([]error, len(baseErrors))
					for i, e := range baseErrors {
						baseErrorsAsError[i] = e
					}

					if matchingError := findMatchingError(entityToFind, baseErrorsAsError); matchingError != nil {
						slog.Info("Found matching error", "headError", headErr.Error(), "baseError", matchingError.Error())
					} else {
						slog.Info("Not found a matching error for", "headError", headErr.Error())
						err = errors.New("found one or more new errors on head")
					}
				}
			}
		}
	}

	return
}
