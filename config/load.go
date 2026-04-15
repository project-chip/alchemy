package config

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/Masterminds/semver"
	"github.com/goccy/go-yaml"
	"github.com/project-chip/alchemy/internal/files"
)

func Load(specRoot string) (*Config, error) {
	b := loadConfigFile(specRoot)
	if b == nil {
		b = defaultConfig
	}
	var cfg Config
	err := yaml.Unmarshal(b, &cfg)
	if err != nil {
		slog.Warn("error parsing config file", slog.Any("error", err))
		return nil, err
	}
	err = checkMinimumVersion(cfg)
	if err != nil {
		return nil, err
	}
	cfg.root = specRoot
	cfg.libraryRoots = make(map[string]struct{})
	for _, l := range cfg.Libraries {
		if _, ok := cfg.libraryRoots[l.Root]; ok {
			return nil, fmt.Errorf("duplicate library root: %s", l.Root)
		}
		cfg.libraryRoots[l.Root] = struct{}{}
	}
	return &cfg, nil
}

func loadConfigFile(specRoot string) []byte {
	if specRoot == "" {
		return nil
	}
	configPath := filepath.Join(specRoot, ".github/alchemy/config.yaml")
	exists, err := files.Exists(configPath)
	if err != nil {
		slog.Warn("error checking for config path", slog.Any("error", err))
		return nil
	}
	if !exists {
		return nil
	}
	b, err := os.ReadFile(configPath)
	if err != nil {
		slog.Warn("error reading config file", slog.Any("error", err))
		return nil
	}
	return b
}

func checkMinimumVersion(cfg Config) error {
	minVersion, err := semver.NewVersion(cfg.MinimumVersion)
	if err != nil {
		slog.Debug("error parsing minimum version", slog.Any("error", err))
		return nil
	}
	bv, err := semver.NewVersion(Version())
	if err != nil {
		slog.Debug("error parsing local version", slog.Any("error", err), slog.String("version", Version()))
		return nil
	}
	if minVersion.GreaterThan(bv) {
		return fmt.Errorf("this version of the Matter specification requires Alchemy %s", minVersion.String())
	}
	return nil
}
