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

	if len(s.Errors) > 0 {
		if filterOptions.Force {
			slog.Warn("Ignoring parse errors; proceed with caution")
		} else if filterOptions.IgnoreErrored {
			var errorPaths []string
			for _, specError := range s.Errors {
				path, _ := specError.Origin()
				path = filepath.Join(s.Root, path)
				slog.Warn("Ignoring errored file", slog.String("file", path))
				errorPaths = append(errorPaths, path)
			}
			filter := paths.NewExcludeFilter[*spec.Doc](s.Root, errorPaths)
			filteredDocs, err = pipeline.Collective(cc, processingOptions, filter, filteredDocs)
			if err != nil {
				return
			}
		} else {
			var errors spec.ParseErrors
			for _, specError := range s.Errors {
				path, _ := specError.Origin()
				path = filepath.Join(s.Root, path)
				_, ok := specDocs.Load(path)
				if ok {
					errors.Errors = append(errors.Errors, specError)
				}
			}
			if len(errors.Errors) > 0 {
				err = &errors
				return
			}
		}
	}
	return
}
