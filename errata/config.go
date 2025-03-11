package errata

import (
	_ "embed"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/Masterminds/semver"
	"github.com/goccy/go-yaml"
	"github.com/project-chip/alchemy/config"
	"github.com/project-chip/alchemy/internal/files"
)

//go:embed default.yaml
var defaultErrata []byte

func LoadErrataConfig(specRoot string) error {
	b := loadConfig(specRoot)
	if b == nil {
		b = defaultErrata
	}
	var errataOverlay errataOverlay
	err := yaml.Unmarshal(b, &errataOverlay)
	if err != nil {
		slog.Warn("error parsing errata file", slog.Any("error", err))
		return nil
	}
	err = checkMinimumVersion(errataOverlay)
	if err != nil {
		return err
	}
	Erratas = errataOverlay.Errata
	for p := range Erratas {
		path := filepath.Join(specRoot, p)
		exists, _ := files.Exists(path)
		if !exists {
			slog.Warn("errata points to non-existent file", "path", p)
		}
	}
	return nil
}

type errataOverlay struct {
	MinimumVersion string             `yaml:"minimumVersion"`
	Errata         map[string]*Errata `yaml:"errata"`
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

func checkMinimumVersion(errataOverlay errataOverlay) error {
	minVersion, err := semver.NewVersion(errataOverlay.MinimumVersion)
	if err != nil {
		slog.Debug("error parsing minimum version", slog.Any("error", err))
		return nil
	}
	bv, err := semver.NewVersion(config.Version())
	if err != nil {
		slog.Debug("error parsing local version", slog.Any("error", err), slog.String("version", config.Version()))
		return nil
	}
	if minVersion.GreaterThan(bv) {
		return fmt.Errorf("this version of the Matter specification requires Alchemy %s", minVersion.String())
	}
	return nil
}
