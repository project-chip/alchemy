package testplan

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
)

type Generator struct {
	testPlanRoot string
	overwrite    bool
}

func (sp Generator) Name() string {
	return "Generating test plans"
}

func (sp Generator) Type() pipeline.ProcessorType {
	return pipeline.ProcessorTypeIndividual
}

func (sp *Generator) Process(cxt context.Context, input *pipeline.Data[*spec.Doc], index int32, total int32) (outputs []*pipeline.Data[string], extras []*pipeline.Data[*spec.Doc], err error) {
	doc := input.Content
	path := doc.Path

	var entities []types.Entity
	entities, err = doc.Entities()
	if err != nil {
		return
	}

	destinations := buildDestinations(sp.testPlanRoot, entities, doc.Errata().TestPlan)

	for newPath, cluster := range destinations {

		_, err = os.ReadFile(newPath)
		if (err == nil || !errors.Is(err, os.ErrNotExist)) && !sp.overwrite {
			slog.InfoContext(cxt, "Skipping existing test plan", slog.String("path", newPath))
			continue
		}

		var result string
		result, err = renderClusterTestPlan(doc, cluster)
		if err != nil {
			err = fmt.Errorf("failed rendering %s: %w", path, err)
			return
		}

		outputs = append(outputs, pipeline.NewData[string](newPath, result))
	}
	return
}

func NewGenerator(testPlanRoot string, overwrite bool) *Generator {
	return &Generator{testPlanRoot: testPlanRoot, overwrite: overwrite}
}

func getTestPlanPath(testplanRoot string, name string) string {
	return filepath.Join(testplanRoot, "src/cluster/", name+".adoc")
}

func testPlanName(path string, entities []types.Entity) string {

	path = filepath.Base(path)
	name := text.TrimCaseInsensitiveSuffix(path, filepath.Ext(path))

	var suffix string
	for _, m := range entities {
		switch m.(type) {
		case *matter.Cluster:
			suffix = "Cluster"
		}
	}
	if !strings.HasSuffix(name, suffix) {
		name += " " + suffix
	}
	return strcase.ToKebab(name)
}
