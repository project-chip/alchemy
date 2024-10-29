package generate

import (
	"regexp"
	"strings"

	"github.com/project-chip/alchemy/matter/types"
	"github.com/project-chip/alchemy/testing/parse"
	"github.com/project-chip/alchemy/testing/pics"
)

type test struct {
	parse.Test
	//Name    string
	ID string
	//Cluster string
	PICSList []pics.Expression
	Groups   []*testGroup

	Variables     []string
	PICSAliases   map[string]string
	PICSAliasList [][]*picsAlias
}

type testGroup struct {
	Step        string
	Description string
	Steps       []*testStep
}

type testStep struct {
	parse.TestStep
	Description      []string
	UserVerification []string
	PICSSet          pics.Expression
}

var stepPattern = regexp.MustCompile(`(?s)^\s*[s|S]tep\s+([0-9a-zA-Z]+):\s*(.*)`)

func (sp *PythonTestGenerator) convert(tst *parse.Test, path string) (t *test, err error) {
	t = &test{
		ID:   getTestName(path),
		Test: *tst,
	}
	for _, tp := range tst.PICS {
		var pe pics.Expression
		pe, err = pics.ParseString(tp)
		if err != nil {
			return
		}
		t.PICSList = append(t.PICSList, pe)
	}
	var currentGroup *testGroup
	for _, s := range tst.Tests {
		ts := &testStep{
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
			if currentGroup == nil || label != currentGroup.Step {
				currentGroup = &testGroup{Step: label, Description: description}
				t.Groups = append(t.Groups, currentGroup)
			}
		} else if currentGroup == nil {
			currentGroup = &testGroup{Step: label, Description: description}
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
	t.PICSAliases = sp.buildPicsMap(t)
	picsAliases := buildPicsAliasList(t.PICSAliases)
	var lastEntityType = types.EntityTypeUnknown
	var entityAliases []*picsAlias
	for _, pa := range picsAliases {
		if pa.entityType != lastEntityType && len(entityAliases) > 0 {
			t.PICSAliasList = append(t.PICSAliasList, entityAliases)
			entityAliases = nil
		}
		entityAliases = append(entityAliases, pa)
		label, ok := sp.picsLabels[pa.Pics]
		if ok {
			pa.Comments = strings.Split(label, "\n")
		}
		lastEntityType = pa.entityType
	}
	if len(entityAliases) > 0 {
		t.PICSAliasList = append(t.PICSAliasList, entityAliases)
	}
	return
}
