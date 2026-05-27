package zapdiff

import (
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"path/filepath"
	"strings"

	"github.com/mailgun/raymond/v2"
	"github.com/project-chip/alchemy/internal/handlebars"
	"github.com/project-chip/alchemy/internal/pipeline"
)

//go:embed template
var templateFiles embed.FS

type jsonMismatch struct {
	Level             string `json:"level"`
	Type              string `json:"type"`
	File              string `json:"file"`
	XPath             string `json:"xpath"`
	Details           string `json:"details"`
	RefContent        string `json:"refContent"`
	GenContent        string `json:"genContent"`
	RefHighlightStart int    `json:"refHighlightStart"`
	RefHighlightEnd   int    `json:"refHighlightEnd"`
	GenHighlightStart int    `json:"genHighlightStart"`
	GenHighlightEnd   int    `json:"genHighlightEnd"`
	HasDiff           bool   `json:"hasDiff"`
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



func WriteMismatchesToHTML(w io.Writer, mm []XmlMismatch, l XmlMismatchLevel, root1, root2 string) error {
	var jmm []jsonMismatch
	for _, m := range mm {
		if m.Level() >= l {
			path1 := filepath.Join(root1, m.Path)
			path2 := filepath.Join(root2, m.Path)

			targetID := m.EntityUniqueIdentifier

			lines1, relStart1, relEnd1, _ := getCustomDiffLines(path1, targetID)
			lines2, relStart2, relEnd2, _ := getCustomDiffLines(path2, targetID)

			refContent := ""
			if len(lines1) > 0 && relStart1 != -1 {
				refContent = strings.Join(lines1, "\n")
			}
			genContent := ""
			if len(lines2) > 0 && relStart2 != -1 {
				genContent = strings.Join(lines2, "\n")
			}

			hasDiff := len(lines1) > 0 || len(lines2) > 0

			jmm = append(jmm, jsonMismatch{
				Level:             m.Level().String(),
				Type:              m.Type.String(),
				File:              m.Path,
				XPath:             m.EntityUniqueIdentifier,
				Details:           m.Details,
				RefContent:        refContent,
				GenContent:        genContent,
				RefHighlightStart: relStart1,
				RefHighlightEnd:   relEnd1,
				GenHighlightStart: relStart2,
				GenHighlightEnd:   relEnd2,
				HasDiff:           hasDiff,
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
