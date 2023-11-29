package zap

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"slices"
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

	filesMap := make(map[string]struct{})
	for _, doc := range docs {
		path := filepath.Base(doc.Path)
		path = strings.TrimSuffix(path, filepath.Ext(path))
		filesMap[fmt.Sprintf("jni/%sClient-InvokeSubscribeImpl.cpp", path)] = struct{}{}
		filesMap[fmt.Sprintf("jni/%sClient-ReadImpl.cpp", path)] = struct{}{}
	}

	matches := gnFileListPattern.FindStringSubmatch(gn)
	if matches == nil {
		return fmt.Errorf("failed to find file list in BUILD.gn")
	}

	files := gnFileLinkPattern.FindAllStringSubmatch(matches[2], -1)
	if files == nil {
		return fmt.Errorf("failed to parse file list in BUILD.gn")
	}
	var combined []string
	var indent string
	for _, fileMatch := range files {
		if indent == "" {
			indent = fileMatch[1]
		}
		delete(filesMap, fileMatch[2])
		combined = append(combined, fileMatch[0])
	}
	for p := range filesMap {
		//fmt.Printf("adding %s...\n", p)
		combined = append(combined, fmt.Sprintf("%s\"%s\",\n", indent, p))
	}
	slices.Sort(combined)

	var replaced bool
	gn = gnFileListPattern.ReplaceAllStringFunc(gn, func(s string) string {
		if replaced {
			return ""
		}
		replaced = true
		return fmt.Sprintf("%s%s%s", matches[1], strings.Join(combined, ""), matches[3])
	})
	return os.WriteFile(buildGNPath, []byte(gn), os.ModeAppend|0644)
}
