package spec

import (
	"fmt"
	"log/slog"
	"path/filepath"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/internal/files"
)

func NewSpecPath(path string, rootPath string) (asciidoc.Path, error) {
	var p asciidoc.Path
	if !filepath.IsAbs(path) {
		var err error
		p.Absolute, err = filepath.Abs(path)
		if err != nil {
			return p, err
		}
	} else {
		p.Absolute = path
	}
	if rootPath == "" {
		rootPath := deriveSpecPath(p.Absolute)
		if rootPath == "" {
			return p, fmt.Errorf("unable to determine root for path %s", p.Absolute)
		}
	}
	var err error
	p.Relative, err = filepath.Rel(rootPath, p.Absolute)
	return p, err
}

func NewDocPath(path string, rootPath string) (asciidoc.Path, error) {
	var p asciidoc.Path
	if !filepath.IsAbs(path) {
		var err error
		p.Absolute, err = filepath.Abs(path)
		if err != nil {
			return p, err
		}
	} else {
		p.Absolute = path
	}
	var err error
	var r string
	r, err = filepath.Abs(rootPath)
	if err != nil {
		return p, err
	}
	p.Relative, err = filepath.Rel(r, p.Absolute)
	return p, err
}

func deriveSpecPath(path string) string {
	if !filepath.IsAbs(path) {
		var err error
		path, err = filepath.Abs(path)
		if err != nil {
			return ""
		}
	}
	dir := filepath.Dir(path)
	for {
		alchemyConfigExists, err := files.Exists(filepath.Join(dir, ".github/workflows/alchemy.yml"))
		if err != nil {
			slog.Warn("error checking for alchemy config file", slog.Any("error", err))
			break
		}
		if alchemyConfigExists {
			return dir
		}
		lastSeparator := strings.LastIndex(dir, string(filepath.Separator))
		if lastSeparator == -1 {
			break
		}
		dir = dir[:lastSeparator]
	}

	// Fallback if we can't find the alchemy.yml file
	parts := strings.Split(dir, string(filepath.Separator))
	for i, part := range parts {
		if strings.EqualFold(part, "connectedhomeip-spec") {
			return strings.Join(parts[0:i+1], string(filepath.Separator))
		}
	}
	return ""
}

func DeriveSpecPathFromPaths(paths []string) string {
	for _, path := range paths {
		specPath := deriveSpecPath(path)
		if specPath != "" {
			return specPath
		}
	}
	return ""
}
