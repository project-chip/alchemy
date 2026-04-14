package zapdiff

import (
	"embed"
	"encoding/json"
	"fmt"
	"io"

	"github.com/mailgun/raymond/v2"
	"github.com/project-chip/alchemy/internal/handlebars"
	"github.com/project-chip/alchemy/internal/pipeline"
)

//go:embed template
var templateFiles embed.FS

type jsonMismatch struct {
	Level   string `json:"level"`
	Type    string `json:"type"`
	File    string `json:"file"`
	XPath   string `json:"xpath"`
	Details string `json:"details"`
}

var htmlReportTemplate pipeline.Once[*raymond.Template]

func LoadHTMLReportTemplate() (*raymond.Template, error) {
	t, err := htmlReportTemplate.Do(func() (*raymond.Template, error) {
		t, err := handlebars.LoadTemplate("{{> template/mismatches}}", templateFiles)
		if err != nil {
			return nil, err
		}
		handlebars.RegisterCommonHelpers(t)
		return t, nil
	})
	if err != nil {
		return nil, err
	}
	return t.Clone(), nil
}

func WriteMismatchesToHTML(w io.Writer, mm []XmlMismatch, l XmlMismatchLevel) error {
	var jmm []jsonMismatch
	for _, m := range mm {
		if m.Level() >= l {
			jmm = append(jmm, jsonMismatch{
				Level:   m.Level().String(),
				Type:    m.Type.String(),
				File:    m.Path,
				XPath:   m.ElementID,
				Details: m.Details,
			})
		}
	}

	jsonData, err := json.Marshal(jmm)
	if err != nil {
		return fmt.Errorf("failed to marshal mismatches to JSON: %w", err)
	}

	t, err := LoadHTMLReportTemplate()
	if err != nil {
		return fmt.Errorf("failed to load HTML template: %w", err)
	}

	res, err := t.Exec(map[string]interface{}{
		"jsonData": string(jsonData),
	})
	if err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	_, err = io.WriteString(w, res)
	return err
}
