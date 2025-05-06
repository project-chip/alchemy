//go:build !github

package cmd

import (
	"errors"
	"fmt"
	"os"
	"slices"

	"github.com/alecthomas/kong"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/internal/vcs"
	"github.com/project-chip/alchemy/matter/spec"
)

func handleError(ctx *kong.Context, err error) {
	var shallowError *vcs.ShallowRepositoryError
	var notARepoError *vcs.NotARepositoryError
	var parseErrors *spec.ParseErrors
	if errors.As(err, &shallowError) {
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "Alchemy was unable to determine a Git tag in your repository.\n\n")
		fmt.Fprintf(os.Stderr, "Some common reasons this might happen:\n\n")
		fmt.Fprintf(os.Stderr, "1. Your clone of the connectedhomeip-spec is too shallow.\n")
		fmt.Fprintf(os.Stderr, "2. You cloned a GitHub fork of the connectedhomeip-spec which was created with the option to not copy branches; this also prevents copying tags.\n")
		fmt.Fprintf(os.Stderr, "3. Your clone is on a branch which does not have any commits with tags on them.\n")
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "Please make a deep clone of the connectedhomeip-spec repository and ensure you are on a branch with at least one tag in its history.\n\n")
	} else if errors.As(err, &notARepoError) {
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "Alchemy was unable to determine a Git tag connectedhomeip-spec because it is not a Git repository.\n\n")
		fmt.Fprintf(os.Stderr, "Some common reasons this might happen:\n\n")
		fmt.Fprintf(os.Stderr, "1. You downloaded the repository as a ZIP file from GitHub and unpacked it.\n")
		fmt.Fprintf(os.Stderr, "2. You copied the files from a clone of the repository, but did not include the .git directory.\n")
		fmt.Fprintf(os.Stderr, "\n\n")
		fmt.Fprintf(os.Stderr, "Please make a deep clone of the connectedhomeip-spec repository and ensure you are on a branch with at least one tag in its history.\n\n")
	} else if errors.As(err, &parseErrors) {
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "Alchemy was unable to proceed due to the following fatal errors in parsing the spec:\n\n")
		var errors []string
		for _, e := range parseErrors.Errors {
			errors = append(errors, fmt.Sprintf("%s - %s", log.Origin(e), e.Error()))
		}
		slices.Sort(errors)
		for _, e := range errors {
			fmt.Fprintf(os.Stderr, "%s\n", e)
		}
		fmt.Fprintf(os.Stderr, "\n\n")

	} else {
		ctx.FatalIfErrorf(err)
	}
}
