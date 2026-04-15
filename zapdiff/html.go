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
	Level   string   `json:"level"`
	Type    string   `json:"type"`
	File    string   `json:"file"`
	XPath   string   `json:"xpath"`
	Details string   `json:"details"`
	Diff    string   `json:"diff"`
	HasDiff bool     `json:"hasDiff"`
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

func getCustomDiffLines(path string, targetID string) ([]string, int, error) {
	lines, startLine, err := findElementLines(path, targetID)
	if err != nil {
		return nil, 0, err
	}
	parentID := getParentID(targetID)
	if parentID == "" || parentID == "configurator" {
		return lines, startLine, nil
	}
	parentLines, parentStartLine, err := findElementLines(path, parentID)
	if err != nil || len(parentLines) == 0 {
		return lines, startLine, nil
	}

	var customLines []string
	customLines = append(customLines, parentLines[0])
	customLines = append(customLines, lines...)
	if len(parentLines) > 1 {
		customLines = append(customLines, parentLines[len(parentLines)-1])
	}

	return customLines, parentStartLine, nil
}

func getMajorParentID(id string) string {
	// IDs look like: configurator/cluster[name='...']/attribute[@code='...']/description
	// We want to truncate to the major parent: cluster, attribute, command, event, feature.
	
	segments := strings.Split(id, "/")
	for i := len(segments) - 1; i >= 0; i-- {
		seg := segments[i]
		if strings.HasPrefix(seg, "cluster[") || 
		   strings.HasPrefix(seg, "attribute[") || 
		   strings.HasPrefix(seg, "command[") || 
		   strings.HasPrefix(seg, "event[") || 
		   strings.HasPrefix(seg, "feature[") {
			return strings.Join(segments[:i+1], "/")
		}
	}
	return id
}

func WriteMismatchesToHTML(w io.Writer, mm []XmlMismatch, l XmlMismatchLevel, root1, root2 string) error {
	var jmm []jsonMismatch
	for _, m := range mm {
		if m.Level() >= l {
			path1 := filepath.Join(root1, m.Path)
			path2 := filepath.Join(root2, m.Path)

			targetID := getMajorParentID(m.EntityUniqueIdentifier)

			lines1, start1, _ := getCustomDiffLines(path1, targetID)
			lines2, start2, _ := getCustomDiffLines(path2, targetID)

			hasDiff := len(lines1) > 0 && len(lines2) > 0

			diffStr := ""
			if hasDiff {
				var err error
				diffStr, err = GenerateUnifiedDiff(lines1, lines2, start1, start2)
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
