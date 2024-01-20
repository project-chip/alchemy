package generate

import (
	"fmt"
	"log/slog"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/hasty/alchemy/ascii"
)

var gnFileListPattern = regexp.MustCompile(`(?m)(?P<Preface> *generator\s*=\s*\"java-jni\"\s*\n+\s*outputs\s*=\s*\[\n)(?P<List>[^\]]+)(?P<Suffix>\s+\])`)
var gnFileLinkPattern = regexp.MustCompile(`(?m)^(?P<Indent>\s+)\"(?P<File>[^"]+)\",\n`)

func patchBuildGN(zclRoot string, docs []*ascii.Doc) error {
	buildGNPath := path.Join(zclRoot, "/src/controller/data_model/BUILD.gn")
	buildGNBytes, err := os.ReadFile(buildGNPath)
	if err != nil {
		return err
	}

	gn := string(buildGNBytes)

	matches := gnFileListPattern.FindStringSubmatch(gn)
	if matches == nil {
		return fmt.Errorf("failed to find file list in BUILD.gn")
	}

	files := gnFileLinkPattern.FindAllStringSubmatch(matches[2], -1)
	if len(files) == 0 {
		return fmt.Errorf("failed to parse file list in BUILD.gn")
	}
	var lines []string
	var indent = files[0][1]

	filesMap := make(map[string]struct{})
	for _, doc := range docs {
		path := filepath.Base(doc.Path)
		path = strings.TrimSuffix(path, filepath.Ext(path))
		filesMap[fmt.Sprintf("%s\"jni/%sClient-InvokeSubscribeImpl.cpp\",\n", indent, path)] = struct{}{}
		filesMap[fmt.Sprintf("%s\"jni/%sClient-ReadImpl.cpp\",\n", indent, path)] = struct{}{}
	}

	for _, fileMatch := range files {
		line := fileMatch[0]
		delete(filesMap, line)
		lines = append(lines, line)
	}

	lines = mergeLines(lines, filesMap, 1)

	slog.Info("Patching src/controller/data_model/BUILD.gn...")

	var replaced bool
	gn = gnFileListPattern.ReplaceAllStringFunc(gn, func(s string) string {
		if replaced {
			return ""
		}
		replaced = true
		return fmt.Sprintf("%s%s%s", matches[1], strings.Join(lines, ""), matches[3])
	})
	return os.WriteFile(buildGNPath, []byte(gn), os.ModeAppend|0644)
}
