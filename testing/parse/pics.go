package parse

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/goccy/go-yaml"
)

type picsFile struct {
	Name    string      `yaml:"name,omitempty"`
	Entries []picsEntry `yaml:"PICS,omitempty"`
}

type picsEntry struct {
	ID    string `yaml:"id,omitempty"`
	Label string `yaml:"label,omitempty"`
}

func LoadPICSLabels(sdkRoot string) (labels map[string]string, err error) {

	path := filepath.Join(sdkRoot, "src/app/tests/suites/certification/PICS.yaml")

	var r []byte
	r, err = os.ReadFile(path)
	if err != nil {
		return
	}

	var t picsFile
	err = yaml.UnmarshalWithOptions(r, &t, yaml.DisallowUnknownField())
	if err != nil {
		err = fmt.Errorf("error parsing %s: %w", path, err)
		return
	}

	labels = make(map[string]string)

	for _, e := range t.Entries {
		_, ok := labels[e.ID]
		if ok {
			slog.Warn("duplicate PICS id", slog.String("id", e.ID))
			continue
		}
		labels[e.ID] = e.Label
	}

	return
}
