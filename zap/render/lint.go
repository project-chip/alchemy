package render

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/project-chip/alchemy/internal/files"
)

var templatePathPattern = regexp.MustCompile(`(?m)^load "../src/app/zap-templates/zcl/data-model/chip/[^.]+\.xml";\n`)

func patchLintBytes(sdkRoot string, xmlFiles []string) (lintPath string, lintBytes []byte, err error) {
	lintPath = filepath.Join(sdkRoot, "/.matterlint")
	exists, _ := files.Exists(lintPath)
	if !exists {
		lintPath = filepath.Join(sdkRoot, "/scripts/rules.matterlint")
		exists, _ = files.Exists(lintPath)
		if !exists {
			err = fmt.Errorf("unable to find rules.matterlint in %s or %s", filepath.Join(sdkRoot, "/rules.matterlint"), filepath.Join(sdkRoot, "/scripts/rules.matterlint"))
			return
		}
	}

	lintBytes, err = os.ReadFile(lintPath)
	if err != nil {
		return
	}
	lint := string(lintBytes)

	newPathMap := make(map[string]struct{})

	for _, f := range xmlFiles {
		newPathMap[`load "../src/app/zap-templates/zcl/data-model/chip/`+filepath.Base(f)+"\";\n"] = struct{}{}
	}

	matches := templatePathPattern.FindAllStringSubmatch(lint, -1)
	paths := make([]string, 0, len(xmlFiles)+len(matches))
	for _, m := range matches {
		path := m[0]
		delete(newPathMap, path)
		paths = append(paths, path)
	}

	paths = mergeLines(paths, newPathMap, 0)

	var sb strings.Builder
	for _, p := range paths {
		sb.WriteString(p)
	}

	var replaced bool
	s := templatePathPattern.ReplaceAllStringFunc(lint, func(s string) string {
		if replaced {
			return ""
		}
		replaced = true
		return sb.String()
	})
	lintBytes = []byte(s)
	return
}
