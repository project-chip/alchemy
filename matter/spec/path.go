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

func NewPath(path string) (Path, error) {
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
	dir := filepath.Dir(p.Absolute)
	parts := strings.Split(dir, string(filepath.Separator))

	rootIndex := -1
	for i := len(parts) - 1; i >= 0; i-- {
		part := parts[i]
		if strings.EqualFold(part, "connectedhomeip-spec") {
			rootIndex = i
			break
		}
	}
	if rootIndex == -1 {
		for i := len(parts) - 1; i >= 0; i-- {
			alchemyYmlPath := filepath.Join(strings.Join(append(parts[i+1:], ".github", "workflows"), string(filepath.Separator)), "alchemy.yml")
			alchemyYmlPathExists, err := files.Exists(alchemyYmlPath)
			if err == nil && alchemyYmlPathExists {
				rootIndex = i
				break
			}
		}
	}
	if rootIndex >= 0 {
		p.Relative = filepath.Join(strings.Join(parts[rootIndex+1:], string(filepath.Separator)), filepath.Base(path))
	} else {
		return p, fmt.Errorf("unable to determine spec root for path %s", p.Absolute)
	}

	return p, nil
}

func DeriveSpecPath(path string) string {
	slog.Info("deriving spec path", slog.String("path", path))
	if !filepath.IsAbs(path) {
		var err error
		path, err = filepath.Abs(path)
		if err != nil {
			return ""
		}
	}
	slog.Info("deriving spec path", slog.String("abs path", path))
	dir := filepath.Dir(path)
	parts := strings.Split(dir, string(filepath.Separator))
	for i, part := range parts {
		slog.Info("spec path part", slog.String("part", part))
		if strings.EqualFold(part, "connectedhomeip-spec") {
			return strings.Join(parts[0:i+1], string(filepath.Separator))
		}
	}
	return ""
}

func DeriveSpecPathFromPaths(paths []string) string {
	for _, path := range paths {
		specPath := DeriveSpecPath(path)
		if specPath != "" {
			slog.Info("derived spec path", slog.String("path", path))
			return specPath
		}
	}
	return ""
}
