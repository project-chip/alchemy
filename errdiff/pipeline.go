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
	Base *spec.Specification
	Head *spec.Specification
}

func Pipeline(cxt context.Context, baseRoot string, headRoot string, docPaths []string, pipelineOptions pipeline.ProcessingOptions) (err error) {
	var specs specs

	parserOptions := spec.ParserOptions{Root: headRoot}
	specs.Head, _, err = spec.Parse(cxt, parserOptions, pipelineOptions, nil, []asciidoc.AttributeName{})
	if err != nil {
		return
	}

	parserOptions = spec.ParserOptions{Root: baseRoot}
	specs.Base, _, err = spec.Parse(cxt, parserOptions, pipelineOptions, nil, []asciidoc.AttributeName{})
	if err != nil {
		return
	}

	groupedHeadErrors := groupErrorsByType(specs.Head.Errors)
	groupedBaseErrors := groupErrorsByType(specs.Base.Errors)
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
