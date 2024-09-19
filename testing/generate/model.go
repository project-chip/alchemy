package generate

import (
	"regexp"

	"github.com/project-chip/alchemy/testing/parse"
	"github.com/project-chip/alchemy/testing/pics"
)

type test struct {
	name    string
	id      string
	cluster string
	pics    []pics.Expression
	steps   []*testStep
}

type testStep struct {
	parse.TestStep
	label       string
	description string
	pics        pics.Expression
}

var stepPattern = regexp.MustCompile(`^\s*[s|S]tep\s+([0-9a-zA-Z]+):\s*(.*)`)

func convert(tst *parse.Test, path string) (t *test, err error) {
	t = &test{
		id:      getTestName(path),
		name:    tst.Name,
		cluster: tst.Config.Cluster,
	}
	for _, tp := range tst.PICS {
		var pe pics.Expression
		pe, err = pics.ParseString(tp)
		if err != nil {
			return
		}
		t.pics = append(t.pics, pe)
	}
	for _, s := range tst.Tests {
		ts := &testStep{
			TestStep: *s,
		}
		label := stepPattern.FindStringSubmatch(s.Label)
		if len(label) > 0 {
			ts.label = label[1]
			ts.description = label[2]
		} else {
			ts.description = s.Label
		}
		ts.pics, err = pics.ParseString(s.PICS)
		if err != nil {
			return
		}
		t.steps = append(t.steps, ts)
	}
	return
}
