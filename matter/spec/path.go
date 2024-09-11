package spec

import (
	"fmt"
	"log/slog"
	"path/filepath"
	"strings"

	"github.com/project-chip/alchemy/internal/files"
)

type Path struct {
	Absolute string
	Relative string
}

func (p Path) String() string {
	return p.Absolute
}

func (p Path) Base() string {
	return filepath.Base(p.Absolute)
}

func (p Path) Ext() string {
	return filepath.Ext(p.Absolute)
}

func (p Path) Dir() string {
	return filepath.Dir(p.Absolute)
}

func (p Path) Origin() (path string, line int) {
	return p.Relative, -1
}

func NewSpecPath(path string, rootPath string) (Path, error) {
	var p Path
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

func NewDocPath(path string, rootPath string) (Path, error) {
	var p Path
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
	p.Relative, err = filepath.Rel(rootPath, p.Absolute)
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
