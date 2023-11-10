package zap

import (
	"os"
	"path/filepath"
	"regexp"
	"slices"
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

	paths := make([]string, 0, len(files))
	for _, f := range files {
		paths = append(paths, `load "../src/app/zap-templates/zcl/data-model/chip/`+filepath.Base(f)+"\";\n")
	}
	matches := deviceType.FindAllStringSubmatch(lint, -1)
	for _, m := range matches {
		paths = append(paths, m[0])
	}

	slices.Sort(paths)

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
