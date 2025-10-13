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

func LoadErrata(config *config.Config) (*Collection, error) {
	b := loadErrataFile(config.Root())
	if b == nil {
		b = defaultErrata
	}
	var errataOverlay errataOverlay
	err := yaml.Unmarshal(b, &errataOverlay)
	if err != nil {
		slog.Warn("error parsing errata file", slog.Any("error", err))
		return nil, err
	}
	err = checkMinimumVersion(errataOverlay)
	if err != nil {
		return nil, err
	}
	// TODO: once alchemy has migrated to using sdk: instead of zap: in errata.yaml, we can drop this indirection
	// This is only here because YAML can't have two tags aliasing the same value
	overlayErrata := make(map[string]*Errata)
	for path, oe := range errataOverlay.Errata {
		var e Errata
		if oe.Disco != nil {
			e.Disco = *oe.Disco
		}
		if oe.Spec != nil {
			e.Spec = *oe.Spec
		}
		if oe.TestPlan != nil {
			e.TestPlan = *oe.TestPlan
		}
		if oe.SDK != nil {
			e.SDK = *oe.SDK
		} else if oe.ZAP != nil {
			e.SDK = *oe.ZAP
		}
		overlayErrata[path] = &e
	}
	c := &Collection{errata: overlayErrata}

	for p := range c.errata {
		path := filepath.Join(config.Root(), p)
		exists, err := files.Exists(path)
		if err != nil {
			slog.Error("error checking if file exists", slog.Any("error", err))
		}
		if !exists {
			slog.Warn("errata points to non-existent file", "path", p)
		}

	}

	return c, nil
}

type errataOverlay struct {
	MinimumVersion string                    `yaml:"minimum-version"`
	Errata         map[string]*overlayErrata `yaml:"errata"`
}

type overlayErrata struct {
	Disco    *Disco    `yaml:"disco,omitempty"`
	Spec     *Spec     `yaml:"spec,omitempty"`
	TestPlan *TestPlan `yaml:"test-plan,omitempty"`
	SDK      *SDK      `yaml:"sdk,omitempty"`
	ZAP      *SDK      `yaml:"zap,omitempty"`
}

func loadErrataFile(specRoot string) []byte {
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

func dumpConfig(c *Collection, errataPath string) {
	errataOverlay := errataOverlay{Errata: make(map[string]*overlayErrata)}
	for path, oe := range c.errata {
		errataOverlay.Errata[path] = &overlayErrata{Disco: &oe.Disco, Spec: &oe.Spec, TestPlan: &oe.TestPlan, SDK: &oe.SDK}
	}
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
