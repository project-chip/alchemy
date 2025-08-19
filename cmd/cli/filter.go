package cli

import (
	"log/slog"
	"path/filepath"

	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/internal/paths"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter/spec"
)

func filterSpecDocs(cc *Context, specDocs spec.DocSet, s *spec.Specification, filterOptions spec.FilterOptions, processingOptions pipeline.ProcessingOptions) (filteredDocs spec.DocSet, err error) {
	filteredDocs = specDocs
	if len(filterOptions.Paths) > 0 { // Filter the spec by whatever extra args were passed
		filter := paths.NewIncludeFilter[*spec.Doc](s.Root, filterOptions.Paths)
		filteredDocs, err = pipeline.Collective(cc, processingOptions, filter, filteredDocs)
		if err != nil {
			return
		}
	}

	if len(filterOptions.Exclude) > 0 {
		filter := paths.NewExcludeFilter[*spec.Doc](s.Root, filterOptions.Exclude)
		filteredDocs, err = pipeline.Collective(cc, processingOptions, filter, filteredDocs)
		if err != nil {
			return
		}
	}
	return
}

func filterSpecErrors[T comparable](cc *Context, input pipeline.Map[string, *pipeline.Data[T]], s *spec.Specification, filterOptions spec.FilterOptions, processingOptions pipeline.ProcessingOptions) (output pipeline.Map[string, *pipeline.Data[T]], err error) {
	output = input
	if len(s.Errors) == 0 {
		return
	}
	if filterOptions.Force {
		return
	}
	if filterOptions.IgnoreErrored {
		var errorPaths []string
		for _, specError := range s.Errors {
			path, _ := specError.Origin()
			path = filepath.Join(s.Root, path)
			slog.Warn("Ignoring errored file", slog.String("file", path))
			errorPaths = append(errorPaths, path)
		}
		filter := paths.NewExcludeFilter[T](s.Root, errorPaths)
		output, err = pipeline.Collective(cc, processingOptions, filter, output)
		return
	}
	return
}

func checkSpecErrors[T comparable](cc *Context, s *spec.Specification, filterOptions spec.FilterOptions, inputs ...pipeline.Map[string, *pipeline.Data[T]]) (err error) {
	if len(s.Errors) == 0 {
		return
	}
	if filterOptions.Force {
		slog.Warn("Ignoring parse errors; proceed with caution")
		return
	}
	if filterOptions.IgnoreErrored {
		return
	}
	// If we are not forcing, then we need to check that all errors are in files we are ignoring
	inputDocs := make(map[string]struct{})
	for _, input := range inputs {
		input.Range(func(key string, value *pipeline.Data[T]) bool {
			inputDocs[key] = struct{}{}
			return true
		})
	}
	var errors spec.ParseErrors
	for _, specError := range s.Errors {
		path, _ := specError.Origin()
		if !filepath.IsAbs(path) {
			path = filepath.Join(s.Root, path)
		}
		_, ok := inputDocs[path]
		if ok {
			errors.Errors = append(errors.Errors, specError)
		} else {
			slog.Warn("ignoring error in file because it did not yield any entities", slog.Any("error", specError), log.Path("source", specError))
		}
	}

	if len(errors.Errors) > 0 {
		err = &errors
		return
	}
	return
}
