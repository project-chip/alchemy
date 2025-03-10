package sdk

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/Masterminds/semver"
	"github.com/goccy/go-yaml"
	"github.com/project-chip/alchemy/config"
	"github.com/project-chip/alchemy/internal/files"
)

func CheckAlchemyVersion(sdkRoot string) error {
	alchemySettings := filepath.Join(sdkRoot, ".github/alchemy.yaml")
	exists, err := files.Exists(alchemySettings)
	if err != nil {
		slog.Error("Error checking for Alchemy SDK settings", slog.String("path", alchemySettings), slog.Any("error", err))
		return nil
	}
	if !exists {
		return nil
	}
	b, err := os.ReadFile(alchemySettings)
	if err != nil {
		slog.Error("Error reading Alchemy SDK settings", slog.String("path", alchemySettings), slog.Any("error", err))
		return nil
	}
	var alchemySdkSettings alchemySdkSettings
	err = yaml.Unmarshal(b, &alchemySdkSettings)
	if err != nil {
		slog.Warn("error parsing Alchemy SDK settings", slog.Any("error", err))
		return nil
	}
	minVersion, err := semver.NewVersion(alchemySdkSettings.MinimumVersion)
	if err != nil {
		slog.Debug("error parsing remote version", slog.Any("error", err))
		return nil
	}
	bv, err := semver.NewVersion(config.Version())
	if err != nil {
		slog.Debug("error parsing local version", slog.Any("error", err), slog.String("version", config.Version()))
		return nil
	}
	if minVersion.GreaterThan(bv) {
		return fmt.Errorf("this version of the SDK requires Alchemy %s", minVersion.String())
	}
	return nil
}

type alchemySdkSettings struct {
	MinimumVersion string `yaml:"minimumVersion"`
}
