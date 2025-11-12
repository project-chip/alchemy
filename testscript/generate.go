package testscript

import (
	"context"
	"log/slog"
	"path/filepath"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
)

type TestScriptGenerator struct {
	spec       *spec.Specification
	sdkRoot    string
	picsLabels map[string]string
}

func NewTestScriptGenerator(spec *spec.Specification, sdkRoot string, picsLabels map[string]string) *TestScriptGenerator {
	return &TestScriptGenerator{spec: spec, sdkRoot: sdkRoot, picsLabels: picsLabels}
}

func (sp TestScriptGenerator) Name() string {
	return "Creating test script steps"
}

func (sp *TestScriptGenerator) Process(cxt context.Context, input *pipeline.Data[*asciidoc.Document], index int32, total int32) (outputs []*pipeline.Data[*Test], extras []*pipeline.Data[*asciidoc.Document], err error) {
	entities := sp.spec.EntitiesForDocument(input.Content)

	var clusters []*matter.Cluster
	for _, m := range entities {
		switch m := m.(type) {
		case *matter.ClusterGroup:
			clusters = append(clusters, m.Clusters...)
		case *matter.Cluster:
			clusters = append(clusters, m)
		}
	}
	for _, cluster := range clusters {
		if len(cluster.Attributes) > 0 {
			var t *Test
			t, err = sp.buildClusterTest(input.Content, cluster)
			if err != nil {
				slog.Error("Error generating test script", matter.LogEntity("cluster", cluster), slog.Any("error", err))
				err = nil
				continue
			}
			outputs = append(outputs, pipeline.NewData(getPath(sp.sdkRoot, t), t))
		}
	}
	return
}

func getPath(sdkRoot string, test *Test) string {

	path := getTestName(test)
	path = strings.ReplaceAll(path, "/", "")
	path += ".py"
	return filepath.Join(sdkRoot, "src/python_testing", path)
}

func getTestName(test *Test) string {
	if strings.HasPrefix(test.ID, "TC_") {
		return test.ID
	}
	return "TC_" + test.ID
}
