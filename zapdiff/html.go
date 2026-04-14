package zapdiff

import (
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"path/filepath"
	"strings"

	"github.com/mailgun/raymond/v2"
	"github.com/pmezard/go-difflib/difflib"
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
	Diff    string `json:"diff"`
	HasDiff bool   `json:"hasDiff"`
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

func getParentID(id string) string {
	i := strings.LastIndex(id, "/")
	if i < 0 {
		return ""
	}
	return id[:i]
}

func WriteMismatchesToHTML(w io.Writer, mm []XmlMismatch, l XmlMismatchLevel, root1, root2 string) error {
	var jmm []jsonMismatch
	for _, m := range mm {
		if m.Level() >= l {
			path1 := filepath.Join(root1, m.Path)
			path2 := filepath.Join(root2, m.Path)

			lines1, _ := findElementLines(path1, m.EntityUniqueIdentifier)
			lines2, _ := findElementLines(path2, m.EntityUniqueIdentifier)

			hasDiff := len(lines1) > 0 && len(lines2) > 0

			if !hasDiff {
				// Fallback to parent if not found by specific ID in both files
				parentID := getParentID(m.EntityUniqueIdentifier)
				if parentID != "" {
					lines1, _ = findElementLines(path1, parentID)
					lines2, _ = findElementLines(path2, parentID)
					hasDiff = len(lines1) > 0 && len(lines2) > 0
				}
			}

			diffStr := ""
			if hasDiff {
				diff := difflib.UnifiedDiff{
					A:        lines1,
					B:        lines2,
					FromFile: "Ref",
					ToFile:   "Generated",
					Context:  3,
				}
				var err error
				diffStr, err = difflib.GetUnifiedDiffString(diff)
				if err != nil {
					diffStr = fmt.Sprintf("Failed to generate diff: %v", err)
				}
			}

			jmm = append(jmm, jsonMismatch{
				Level:   m.Level().String(),
				Type:    m.Type.String(),
				File:    m.Path,
				XPath:   m.EntityUniqueIdentifier,
				Details: m.Details,
				Diff:    diffStr,
				HasDiff: hasDiff,
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
