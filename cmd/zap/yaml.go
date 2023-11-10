package zap

import (
	"fmt"
	"os"
	"path"
	"regexp"
	"slices"
	"strings"

	"gopkg.in/yaml.v3"
)

var appTemplates = regexp.MustCompile(`(?m)^ +(src/app/zap-templates/[^\. ]+.xml) +\\+\n`)

func patchTestsYaml(zclRoot string, files []string) error {
	testsYamlPath := path.Join(zclRoot, ".github/workflows/tests.yaml")
	yamlBytes, err := os.ReadFile(testsYamlPath)
	if err != nil {
		return err
	}

	var doc yaml.Node
	err = yaml.Unmarshal(yamlBytes, &doc)

	if err != nil {
		return err
	}

	root := doc.Content[0]

	jobs := getYamlMapEntry(root, "jobs", yaml.ScalarNode)
	if jobs == nil {
		return fmt.Errorf("could not find jobs node")
	}
	testSuites := getYamlMapEntry(jobs, "test_suites_linux", yaml.ScalarNode)
	if testSuites == nil {
		return fmt.Errorf("could not find test_suites_linux node")
	}
	steps := getYamlMapEntry(testSuites, "steps", yaml.ScalarNode)
	if testSuites == nil {
		return fmt.Errorf("could not find steps node")
	}
	var parserRun *yaml.Node
	for _, step := range steps.Content {
		run := getYamlMapEntry(step, "run", yaml.ScalarNode)
		if run == nil {
			continue
		}
		if strings.Contains(run.Value, "zapxml_parser.py") {
			parserRun = run
			break

		}
	}
	if parserRun == nil {
		return fmt.Errorf("could not find zapxml_parser step node")
	}
	paths := make([]string, 0, len(files))
	for _, f := range files {
		paths = append(paths, strings.TrimPrefix(f, "connectedhomeip/")+" \\")
	}
	matches := appTemplates.FindAllStringSubmatch(parserRun.Value, -1)
	for _, m := range matches {
		paths = append(paths, strings.TrimSpace(m[0]))
	}
	slices.Sort(paths)
	slices.Compact(paths)
	slices.SortFunc(paths, func(a string, b string) int {
		if strings.HasPrefix(a, "src/app/zap-templates/zcl/data-model/chip/global-attributes") {
			return -1
		} else if strings.HasPrefix(b, "src/app/zap-templates/zcl/data-model/chip/global-attributes") {
			return 1
		}
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
		return 0
	})
	var sb strings.Builder
	for _, p := range paths {
		sb.WriteString("    ")
		sb.WriteString(p)
		sb.WriteRune('\n')
	}
	var replaced bool
	s := appTemplates.ReplaceAllStringFunc(parserRun.Value, func(s string) string {
		if replaced {
			return ""
		}
		replaced = true
		return sb.String()
	})
	parserRun.Value = s
	yamlBytes, err = yaml.Marshal(&doc)
	if err != nil {
		return err
	}
	return os.WriteFile(testsYamlPath, yamlBytes, os.ModeAppend|0644)
}

func getYamlNode(root *yaml.Node, name string, kind yaml.Kind) *yaml.Node {

	for _, key := range root.Content {
		if key.Kind == kind && key.Value == name {
			return key

		}
	}
	return nil
}

func getYamlMapEntry(root *yaml.Node, name string, kind yaml.Kind) *yaml.Node {

	for i := 0; i < len(root.Content)/2; i++ {
		key := root.Content[i*2]
		if key.Kind == kind && key.Value == name {
			return root.Content[i*2+1]

		}
	}
	return nil
}
