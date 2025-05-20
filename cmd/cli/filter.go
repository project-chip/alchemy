package cli

import (
	"log/slog"
	"path/filepath"

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

func checkSpecErrors[T any](cc *Context, input pipeline.Map[string, *pipeline.Data[T]], s *spec.Specification, filterOptions spec.FilterOptions, processingOptions pipeline.ProcessingOptions) (output pipeline.Map[string, *pipeline.Data[T]], err error) {
	output = input
	if len(s.Errors) == 0 {
		return
	}
	if filterOptions.Force {
		slog.Warn("Ignoring parse errors; proceed with caution")
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

	var errors spec.ParseErrors
	for _, specError := range s.Errors {
		path, _ := specError.Origin()
		path = filepath.Join(s.Root, path)
		_, ok := output.Load(path)
		if ok {
			errors.Errors = append(errors.Errors, specError)
		}
	}
	if len(errors.Errors) > 0 {
		err = &errors
		return
	}
	return
}
