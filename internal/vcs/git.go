package vcs

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

type NotARepositoryError struct {
	Path string
}

func (nar *NotARepositoryError) Error() string {
	return fmt.Sprintf("%s is not a Git repository", nar.Path)
}

type ShallowRepositoryError struct {
	Path string
}

func (nar *ShallowRepositoryError) Error() string {
	return fmt.Sprintf("%s is too shallow", nar.Path)
}

func GitDescribe(root string) (string, error) {
	cmd := exec.Command("git", "describe", "--dirty", "--broken", "--tags")
	cmd.Dir = root

	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	out, err := cmd.Output()
	if err != nil {
		return "", parseGitError(root, stderr.String())
	}
	gitDescription := strings.TrimSpace(string(out))
	if gitDescription == "" {
		return "", fmt.Errorf("git describe returned empty string")
	}
	return gitDescription, nil
}

func parseGitError(root string, e string) error {
	if strings.HasPrefix(e, "fatal: No names found, cannot describe anything.") {
		return &ShallowRepositoryError{Path: root}
	} else if strings.HasPrefix(e, "fatal: not a git repository") {
		return &NotARepositoryError{Path: root}
	}
	return fmt.Errorf("error running git: %s", e)
}
