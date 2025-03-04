package yaml2python

import (
	"context"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
	"github.com/project-chip/alchemy/testplan"
	"github.com/project-chip/alchemy/testplan/pics"
	"github.com/project-chip/alchemy/yaml2python/parse"
)

type YamlTestConverter struct {
	spec       *spec.Specification
	sdkRoot    string
	picsLabels map[string]string
}

func NewYamlTestConverter(spec *spec.Specification, sdkRoot string, picsLabels map[string]string) *YamlTestConverter {
	return &YamlTestConverter{spec: spec, sdkRoot: sdkRoot, picsLabels: picsLabels}
}

func (ytc YamlTestConverter) Name() string {
	return "Converting YAML tests"
}

var stepPattern = regexp.MustCompile(`(?s)^\s*[s|S]tep\s+([0-9a-zA-Z]+):\s*(.*)`)

func (ytc *YamlTestConverter) Process(cxt context.Context, input *pipeline.Data[*parse.Test], index int32, total int32) (outputs []*pipeline.Data[*testplan.Test], extras []*pipeline.Data[*parse.Test], err error) {

	t := &testplan.Test{
		Test: *input.Content,
		ID:   getTestName(input.Path),
	}
	for _, tp := range input.Content.PICS {
		var pe pics.Expression
		pe, err = pics.ParseString(tp)
		if err != nil {
			return
		}
		t.PICSList = append(t.PICSList, pe)
	}
	var currentGroup *testplan.Group
	for _, s := range input.Content.Tests {
		ts := &testplan.Step{
			TestStep: *s,
		}
		labelParts := stepPattern.FindStringSubmatch(s.Label)
		var label, description string
		if len(labelParts) > 0 {
			label = labelParts[1]
			description = labelParts[2]
		} else {
			description = s.Label
		}
		if len(label) > 0 {
			if currentGroup == nil || label != currentGroup.Name {
				currentGroup = &testplan.Group{Name: label, Description: description}
				t.Groups = append(t.Groups, currentGroup)
			}
		} else if currentGroup == nil {
			currentGroup = &testplan.Group{Parent: t, Name: label, Description: description}
		}
		if len(description) > 0 {

			ts.Description = strings.Split(description, "\n")
		}
		if len(s.Verification) > 0 {
			ts.UserVerification = strings.Split(s.Verification, "\n")
		}
		ts.PICSSet, err = pics.ParseString(s.PICS)
		if err != nil {
			return
		}
		currentGroup.Steps = append(currentGroup.Steps, ts)
	}
	t.Variables = getVariables(t)
	t.PICSAliases = testplan.BuildPicsMap(ytc.spec, t)
	picsAliases := testplan.BuildPicsAliasList(t.PICSAliases)
	var lastEntityType = types.EntityTypeUnknown
	var entityAliases []*testplan.PicsAlias
	for _, pa := range picsAliases {
		if pa.EntityType != lastEntityType && len(entityAliases) > 0 {
			t.PICSAliasList = append(t.PICSAliasList, entityAliases)
			entityAliases = nil
		}
		entityAliases = append(entityAliases, pa)
		label, ok := ytc.picsLabels[pa.Pics]
		if ok {
			pa.Comments = strings.Split(label, "\n")
		}
		lastEntityType = pa.EntityType
	}
	if len(entityAliases) > 0 {
		t.PICSAliasList = append(t.PICSAliasList, entityAliases)
	}
	outputs = append(outputs, pipeline.NewData(ytc.getPath(input.Path), t))
	return
}

func (ytc *YamlTestConverter) getPath(path string) string {

	path = getTestName(path)
	path += ".py"
	return filepath.Join(ytc.sdkRoot, "src/python_testing", path)
}

func getTestName(path string) string {
	path = filepath.Base(path)
	path = text.TrimCaseInsensitivePrefix(path, "test_")
	path = text.TrimCaseInsensitiveSuffix(path, ".yaml")
	return path
}
