package zap

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var deviceType = regexp.MustCompile(`(?m)^load "../src/app/zap-templates/zcl/data-model/chip/[^.]+\.xml";\n`)

func patchLint(zclRoot string, files []string) error {

	lintPath := filepath.Join(zclRoot, "/scripts/rules.matterlint")
	lintBytes, err := os.ReadFile(lintPath)
	if err != nil {
		return err
	}
	lint := string(lintBytes)

	newPathMap := make(map[string]struct{})

	for _, f := range files {
		newPathMap[`load "../src/app/zap-templates/zcl/data-model/chip/`+filepath.Base(f)+"\";\n"] = struct{}{}
	}

	matches := deviceType.FindAllStringSubmatch(lint, -1)
	paths := make([]string, 0, len(files)+len(matches))
	for _, m := range matches {
		path := m[0]
		delete(newPathMap, path)
		paths = append(paths, path)
	}
	for path := range newPathMap {
		paths = append(paths, path)
	}

	var sb strings.Builder
	for _, p := range paths {
		sb.WriteString(p)
	}

	var replaced bool
	s := deviceType.ReplaceAllStringFunc(lint, func(s string) string {
		if replaced {
			return ""
		}
		replaced = true
		return sb.String()
	})

	return os.WriteFile(lintPath, []byte(s), os.ModeAppend|0644)
}
