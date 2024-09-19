package generate

import (
	"context"
	"fmt"
	"log/slog"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/testing/parse"
)

type PythonTestGenerator struct {
	sdkRoot   string
	overwrite bool

	spec   *spec.Specification
	labels map[string]string
}

func NewPythonTestGenerator(spec *spec.Specification, sdkRoot string, overwrite bool, labels map[string]string) *PythonTestGenerator {
	return &PythonTestGenerator{spec: spec, sdkRoot: sdkRoot, overwrite: overwrite, labels: labels}
}

func (sp PythonTestGenerator) Name() string {
	return "Generating test plans"
}

func (sp PythonTestGenerator) Type() pipeline.ProcessorType {
	return pipeline.ProcessorTypeIndividual
}

func (sp *PythonTestGenerator) Process(cxt context.Context, input *pipeline.Data[*parse.Test], index int32, total int32) (outputs []*pipeline.Data[string], extras []*pipeline.Data[*parse.Test], err error) {
	outPath := sp.getPath(input.Path)
	slog.Info("generating", "in", input.Path, "out", outPath)

	var script strings.Builder
	script.WriteString(preamble)

	var test *test
	test, err = convert(input.Content, input.Path)
	if err != nil {
		return
	}

	picsAliases := sp.buildPicsMap(test)

	script.WriteString(fmt.Sprintf("\ncluster = Clusters.%s\n\n", test.cluster))

	pics := make([]string, 0, len(test.pics))
	for _, r := range test.pics {
		pics = append(pics, r.PythonString())
	}

	script.WriteString(fmt.Sprintf("class %s(MatterBaseTest):\n", test.id))

	script.WriteString("\n")
	script.WriteString(fmt.Sprintf("    def desc_%s(self) -> str:\n", test.id))
	script.WriteString("        \"\"\"Returns a description of this test\"\"\"\n")
	script.WriteString(fmt.Sprintf("        return \"%s\"\n", test.name))

	script.WriteString("\n")
	script.WriteString(fmt.Sprintf("    def pics_%s(self):\n", test.id))
	script.WriteString("        \"\"\"This function returns a list of PICS for this test case that must be True for the test to be run\"\"\"\n")
	script.WriteString(fmt.Sprintf("        return [%s]\n", strings.Join(pics, ",")))

	script.WriteString("\n")
	script.WriteString(fmt.Sprintf("    def steps_%s(self) -> list[TestStep]:\n", test.id))
	script.WriteString("        steps = [\n")
	var lastLabel string
	for _, step := range test.steps {
		if step.label != "" && lastLabel != step.label {
			script.WriteString("            TestStep(")
			script.WriteString(strconv.Quote(step.label))
			script.WriteString(", ")
			script.WriteString(strconv.Quote(step.description))
			script.WriteString("),\n")
			lastLabel = step.label
		}

	}
	script.WriteString("        ]\n\n        return steps\n\n")

	script.WriteString("    @ async_test_body\n")
	script.WriteString(fmt.Sprintf("    async def test_%s(self):\n", test.id))
	script.WriteString("        endpoint = self.matter_test_config.endpoint if self.matter_test_config.endpoint is not None else 1\n")

	writeVariables(test, &script)

	sp.writePicsAliases(picsAliases, &script)

	lastLabel = ""
	for _, step := range test.steps {
		if step.label != "" && step.label != lastLabel {
			script.WriteString("        self.step(")
			script.WriteString(strconv.Quote(step.label))
			script.WriteString(")\n\n")
			lastLabel = step.label
		}
		switch step.Cluster {
		case "DelayCommands":
		default:
			indent := "        "
			if len(step.description) > 0 {
				script.WriteString(indent)
				script.WriteString("# ")
				script.WriteString(strings.ReplaceAll(step.description, "\n", " "))
				script.WriteRune('\n')
			}
			for _, comment := range step.Comments {
				script.WriteString(indent)
				script.WriteString("# ")
				script.WriteString(comment)
				script.WriteRune('\n')
			}
			if step.pics != nil {
				sp.writePicsGuard(step, picsAliases, &script)
				indent = "            "
			}
			err = sp.writeCommand(test, step, indent, &script)
			if err != nil {
				return
			}
			script.WriteRune('\n')
		}
	}

	script.WriteString("if __name__ == \"__main__\":\n")
	script.WriteString("    default_matter_test_main()\n")

	outputs = append(outputs, pipeline.NewData(outPath, script.String()))
	return
}

func (sp *PythonTestGenerator) getPath(path string) string {

	path = getTestName(path)
	path += ".py"
	return filepath.Join(sp.sdkRoot, "src/python_testing", path)
}

func getTestName(path string) string {
	path = filepath.Base(path)
	path = text.TrimCaseInsensitivePrefix(path, "test_")
	path = text.TrimCaseInsensitiveSuffix(path, ".yaml")
	return path
}
