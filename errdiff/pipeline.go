package errdiff

import (
	"context"
	"errors"
	"log/slog"
	"os"

	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
)

type ComparableError interface {
	spec.Error
	ComparableEntity() types.ComparableEntity
}

func Pipeline(cxt context.Context, baseRoot string, headRoot string, docPaths []string, pipelineOptions pipeline.ProcessingOptions, outputFile string) (err error) {
	specs, err := spec.LoadSpecSet(cxt, baseRoot, headRoot, docPaths, pipelineOptions, nil)
	if err != nil {
		return
	}
	_, err = ProcessComparison(&specs, outputFile)
	return
}

func ProcessComparison(specs *spec.SpecSet, outputFile string) (violations map[string][]spec.Violation, err error) {
	slog.Info("Comparing head and base")
	err1 := compare(specs.Base, specs.Head)

	slog.Info("Comparing head and base (in-progress)")
	err2 := compare(specs.BaseInProgress, specs.HeadInProgress)

	err = errors.Join(err1, err2)

	if err != nil && len(outputFile) > 0 {
		fileErr := os.WriteFile(outputFile, []byte(err.Error()), 0666)
		if fileErr != nil {
			slog.Error("error writing to output file", "err", fileErr)
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

					if matchingError := findMatchingError(entityToFind, baseErrorsAsError); matchingError == nil {
						slog.Error("This error is introduced by the current PR: <", headErr.Error(), ">")
						err = errors.Join(err, errors.New(("This error is introduced by the current PR: <" + headErr.Error() + ">")))
					}
				}
			}
		}
	}

	return
}
