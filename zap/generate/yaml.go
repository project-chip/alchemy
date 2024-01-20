package generate

import (
	"fmt"
	"os"
	"path"
	"regexp"
	"slices"
	"strings"
)

var yamlFileLinkPattern = regexp.MustCompile(`(?m)^(?P<Indent>\s+)(?P<File>src/app/zap-templates/zcl/data-model/(?:[^/\.]+/)*(?:[^.]+\.)xml)\s\\\n`)

func patchTestsYaml(zclRoot string, files []string) error {
	testsYamlPath := path.Join(zclRoot, ".github/workflows/tests.yaml")
	yamlBytes, err := os.ReadFile(testsYamlPath)
	if err != nil {
		return err
	}

	yaml := string(yamlBytes)

	matches := yamlFileLinkPattern.FindAllStringSubmatch(yaml, -1)
	if len(matches) == 0 {
		return fmt.Errorf("could not find existing paths in tests.yaml")
	}
	var indent = matches[0][1]

	filesMap := make(map[string]struct{})
	for _, file := range files {
		path := fmt.Sprintf("%ssrc/app/zap-templates/zcl/data-model/chip/%s \\\n", indent, strings.TrimPrefix(file, "connectedhomeip/"))
		filesMap[path] = struct{}{}
	}

	var sb strings.Builder
	lines := make([]string, 0, len(matches))
	for _, m := range matches {
		line := m[0]
		delete(filesMap, line)
		lines = append(lines, line)
	}

	for file := range filesMap {
		var inserted bool
		for i, line := range lines {
			if i < 1 {
				// We skip the first line, "global-attributes.xml"
				continue
			}
			if strings.Compare(file, line) < 0 {
				lines = slices.Insert(lines, i, file)
				inserted = true
				break
			}
		}
		if !inserted {
			lines = append(lines, file)
		}
	}

	for _, line := range lines {
		sb.WriteString(line)
	}

	var replaced bool
	yaml = yamlFileLinkPattern.ReplaceAllStringFunc(yaml, func(s string) string {
		if replaced {
			return ""
		}
		replaced = true
		return sb.String()
	})

	return os.WriteFile(testsYamlPath, []byte(yaml), os.ModeAppend|0644)
}
