package errata

import (
	"log/slog"
	"os"
	"path/filepath"

	"github.com/goccy/go-yaml"
	"github.com/project-chip/alchemy/internal/files"
)

func OverlayErrataConfig(specRoot string) {
	if specRoot == "" {
		return
	}
	errataPath := filepath.Join(specRoot, ".github/alchemy/errata.yaml")
	exists, err := files.Exists(errataPath)
	if err != nil {
		slog.Warn("error checking for errata path", slog.Any("error", err))
		return
	}
	if !exists {
		slog.Warn("errata file does not exist", slog.Any("path", errataPath))
	} else {
		b, err := os.ReadFile(errataPath)
		if err != nil {
			slog.Warn("error reading errata file", slog.Any("error", err))
			return
		}
		var errataOverlay errataOverlay
		err = yaml.Unmarshal(b, &errataOverlay)
		if err != nil {
			slog.Warn("error parsing errata file", slog.Any("error", err))
			return
		}
		slog.Warn("Using errata overlay", slog.Any("path", errataPath), slog.Any("count", len(errataOverlay.Errata)))
		Erratas = errataOverlay.Errata
	}
	/*eo := errataOverlay{Errata: Erratas}
	d, err := yaml.Marshal(&eo)
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
	}*/

}

type errataOverlay struct {
	Errata map[string]*Errata `yaml:"errata"`
}