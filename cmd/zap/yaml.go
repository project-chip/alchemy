package zap

import (
	"fmt"
	"os"
	"path"
	"regexp"
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

	filesMap := make(map[string]struct{})
	for _, file := range files {
		path := fmt.Sprintf("src/app/zap-templates/zcl/data-model/chip/%s", strings.TrimPrefix(file, "connectedhomeip/"))
		//fmt.Printf("adding file %s\n", path)
		filesMap[path] = struct{}{}
	}

	matches := yamlFileLinkPattern.FindAllStringSubmatch(yaml, -1)
	var indent string
	var sb strings.Builder
	for _, m := range matches {
		if len(indent) == 0 && len(m[1]) > 0 {
			indent = m[1]
		}
		if _, ok := filesMap[m[2]]; ok {
			delete(filesMap, m[2])
		}
		sb.WriteString(m[0])
	}
	for file := range filesMap {
		sb.WriteString(indent)
		sb.WriteString(file)
		sb.WriteString(" \\\n")
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
