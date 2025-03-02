package render

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/mailgun/raymond/v2"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
)

type GeneratorOption func(g *Renderer)

type Renderer struct {
	testPlanRoot string
	templateRoot string
	overwrite    bool
}

func NewRenderer(testPlanRoot string, overwrite bool, options ...GeneratorOption) *Renderer {
	g := &Renderer{testPlanRoot: testPlanRoot, overwrite: overwrite}
	for _, o := range options {
		o(g)
	}
	return g
}

func TemplateRoot(templateRoot string) func(*Renderer) {
	return func(g *Renderer) {
		g.templateRoot = templateRoot
	}
}

func (sp Renderer) Name() string {
	return "Generating test plans"
}

func (sp *Renderer) Process(cxt context.Context, input *pipeline.Data[*spec.Doc], index int32, total int32) (outputs []*pipeline.Data[string], extras []*pipeline.Data[*spec.Doc], err error) {
	doc := input.Content
	path := doc.Path

	var entities []types.Entity
	entities, err = doc.Entities()
	if err != nil {
		return
	}

	destinations := buildDestinations(sp.testPlanRoot, entities, doc.Errata().TestPlan)

	var t *raymond.Template
	t, err = sp.loadTemplate()
	if err != nil {
		return
	}

	for newPath, cluster := range destinations {

		_, err = os.ReadFile(newPath)
		if (err == nil || !errors.Is(err, os.ErrNotExist)) && !sp.overwrite {
			slog.InfoContext(cxt, "Skipping existing test plan", slog.String("path", newPath))
			continue
		}

		var cut *clusterUnderTest
		cut, err = filterCluster(doc, cluster)
		if err != nil {
			err = fmt.Errorf("failed filtering %s: %w", path, err)
			return
		}

		tc := templateContext{ReferenceStore: conformance.ReferenceStore(doc)}

		args := map[string]any{
			"cluster":           cluster,
			"context":           tc,
			"doc":               doc,
			"features":          cut.features,
			"attributes":        cut.attributes,
			"commandsAccepted":  cut.commandsAccepted,
			"commandsGenerated": cut.commandsGenerated,
			"events":            cut.events,
		}

		if len(cluster.Revisions) > 0 {
			args["lastRevision"] = cluster.Revisions[len(cluster.Revisions)-1]
		}
		args["attributeOptions"] = getOptionality(cut.attributes, func(a *matter.Field) string { return a.Name }, func(a *matter.Field) conformance.Set { return a.Conformance })
		args["eventOptions"] = getOptionality(cut.events, func(a *matter.Event) string { return a.Name }, func(a *matter.Event) conformance.Set { return a.Conformance })
		args["commandsAcceptedOptions"] = getOptionality(cut.commandsAccepted, func(a *matter.Command) string { return a.Name }, func(a *matter.Command) conformance.Set { return a.Conformance })
		args["commandsGeneratedOptions"] = getOptionality(cut.commandsGenerated, func(a *matter.Command) string { return a.Name }, func(a *matter.Command) conformance.Set { return a.Conformance })

		var result string
		result, err = t.Exec(args)
		if err != nil {
			err = fmt.Errorf("failed rendering %s: %w", path, err)
			return
		}

		outputs = append(outputs, pipeline.NewData[string](newPath, result))
	}
	return
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

func buildDestinations(testplanRoot string, entities []types.Entity, errata errata.TestPlan) (destinations map[string]*matter.Cluster) {
	destinations = make(map[string]*matter.Cluster)

	for _, e := range entities {
		switch e := e.(type) {
		case *matter.ClusterGroup:
			for _, c := range e.Clusters {
				destinations[getTestPlanPathForCluster(testplanRoot, c, errata)] = c
			}
		case *matter.Cluster:
			destinations[getTestPlanPathForCluster(testplanRoot, e, errata)] = e
		}
	}
	return

}

func getTestPlanPathForCluster(testplanRoot string, cluster *matter.Cluster, errata errata.TestPlan) string {
	if len(errata.TestPlanPaths) > 0 {
		tpp, ok := errata.TestPlanPaths[cluster.Name]
		if ok {
			return filepath.Join(testplanRoot, tpp.Path)
		}
	}
	if errata.TestPlanPath != "" {
		return filepath.Join(testplanRoot, errata.TestPlanPath)
	}
	return getTestPlanPath(testplanRoot, strings.ToLower(strcase.ToSnake(cluster.Name)))
}
