package errata

import (
	_ "embed"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/goccy/go-yaml"
	"github.com/project-chip/alchemy/internal/files"
)

//go:embed default.yaml
var defaultErrata []byte

func LoadErrataConfig(specRoot string) {
	b := loadConfig(specRoot)
	if b == nil {
		b = defaultErrata
	}
	var errataOverlay errataOverlay
	err := yaml.Unmarshal(b, &errataOverlay)
	if err != nil {
		slog.Warn("error parsing errata file", slog.Any("error", err))
		return
	}
	Erratas = errataOverlay.Errata
	for p := range Erratas {
		path := filepath.Join(specRoot, p)
		exists, _ := files.Exists(path)
		if !exists {
			slog.Warn("errata points to non-existent file", "path", p)
		}
	}
}

type errataOverlay struct {
	Errata map[string]*Errata `yaml:"errata"`
}

func loadConfig(specRoot string) []byte {
	if specRoot == "" {
		return nil
	}
	errataPath := filepath.Join(specRoot, ".github/alchemy/errata.yaml")
	exists, err := files.Exists(errataPath)
	if err != nil {
		slog.Warn("error checking for errata path", slog.Any("error", err))
		return nil
	}
	if !exists {
		return nil
	}
	b, err := os.ReadFile(errataPath)
	slog.Debug("Using errata overlay", slog.Any("path", errataPath))
	if err != nil {
		slog.Warn("error reading errata file", slog.Any("error", err))
		return nil
	}
	return b
}

func dumpConfig(errataPath string) {
	errataOverlay := errataOverlay{Errata: Erratas}
	d, err := yaml.Marshal(&errataOverlay)
	if err != nil {
		slog.Warn("error marshalling yaml", slog.Any("error", err))
	}
	dir := filepath.Dir(errataPath)
	if _, de := os.Stat(dir); os.IsNotExist(de) {
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return
		}
	}
	err = os.WriteFile(errataPath, d, os.ModeAppend|0644)
	if err != nil {
		slog.Warn("error writing errata", slog.Any("error", err))
		return
	}
}
