package testscript

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/iancoleman/strcase"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/constraint"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
	"github.com/project-chip/alchemy/testplan"
	"github.com/project-chip/alchemy/testplan/pics"
	"github.com/project-chip/alchemy/testscript/yaml2python/parse"
)

type TestPlanGenerator struct {
	spec       *spec.Specification
	sdkRoot    string
	picsLabels map[string]string
}

func NewTestPlanGenerator(spec *spec.Specification, sdkRoot string, picsLabels map[string]string) *TestPlanGenerator {
	return &TestPlanGenerator{spec: spec, sdkRoot: sdkRoot, picsLabels: picsLabels}
}

func (sp TestPlanGenerator) Name() string {
	return "Creating test script steps"
}

func (sp *TestPlanGenerator) Process(cxt context.Context, input *pipeline.Data[*spec.Doc], index int32, total int32) (outputs []*pipeline.Data[*testplan.Test], extras []*pipeline.Data[*spec.Doc], err error) {
	var entities []types.Entity
	entities, err = input.Content.Entities()
	if err != nil {
		return
	}

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
			var t *testplan.Test
			t, err = sp.buildAttributesTest(cluster)
			outputs = append(outputs, pipeline.NewData(sp.getPath(t), t))
		}
	}
	return
}

func (*TestPlanGenerator) buildAttributesTest(cluster *matter.Cluster) (t *testplan.Test, err error) {
	t = &testplan.Test{
		Cluster: cluster,
		ID:      strcase.ToScreamingSnake(spec.CanonicalName(cluster.Name)) + "_2_1",
		Test: parse.Test{
			Name: "Attributes with Server as DUT",
			Config: parse.TestConfig{
				Cluster: cluster.Name,
			},
		},
	}
	t.PICS = []string{cluster.PICS}
	counter := 1
	for _, a := range cluster.Attributes {
		g := &testplan.Group{Parent: t, Name: fmt.Sprintf("%d", counter), Description: fmt.Sprintf("Read %s attribute", a.Name)}
		counter++
		step := &testplan.Step{
			TestStep: parse.TestStep{
				Command:   "readAttribute",
				Attribute: a.Name,
			},
		}
		if a.Type != nil {
			ensureConstraints(step).Type = a.Type.Name
		}
		err = setPicsSet(step, a)
		if err != nil {
			return
		}
		err = setConstraintChecks(step, a)
		if err != nil {
			return
		}
		g.Steps = append(g.Steps, step)
		t.Groups = append(t.Groups, g)
	}
	return
}

func setPicsSet(step *testplan.Step, field *matter.Field) (err error) {
	if isExcludedConformance(field.Conformance) {
		return
	}
	var picsSet pics.Expression
	picsSet, err = pics.ConvertConformance(field, field.Conformance)
	if err != nil {
		return
	} else {
		step.PICSSet = picsSet
	}
	return
}

func setConstraintChecks(step *testplan.Step, field *matter.Field) (err error) {
	if constraint.IsBlank(field.Constraint) {
		return
	}
	switch c := field.Constraint.(type) {
	case *constraint.MinConstraint:
		if field.Type.IsArray() {
			ensureConstraints(step).MinLength = c.Minimum.DataModelString(field.Type)
		} else {
			ensureConstraints(step).MinValue = c.Minimum.DataModelString(field.Type)
		}
	case *constraint.MaxConstraint:
		if field.Type.IsArray() {
			ensureConstraints(step).MaxLength = c.Maximum.DataModelString(field.Type)
		} else {
			ensureConstraints(step).MaxValue = c.Maximum.DataModelString(field.Type)
		}
	case *constraint.RangeConstraint:
		if field.Type.IsArray() {
			ensureConstraints(step).MinLength = c.Minimum.DataModelString(field.Type)
			ensureConstraints(step).MaxLength = c.Maximum.DataModelString(field.Type)
		} else {
			ensureConstraints(step).MinValue = c.Minimum.DataModelString(field.Type)
			ensureConstraints(step).MaxValue = c.Maximum.DataModelString(field.Type)
		}
	case *constraint.AllConstraint, *constraint.DescribedConstraint:
		return
	default:
		err = fmt.Errorf("unexpected constraint type setting test step response checks: %T", c)
		return
	}
	return
}

func ensureConstraints(step *testplan.Step) *parse.StepResponseConstraints {
	if step.Response.Constraints == nil {
		step.Response.Constraints = &parse.StepResponseConstraints{}
	}
	return step.Response.Constraints
}

func (ytc *TestPlanGenerator) getPath(test *testplan.Test) string {

	path := getTestName(test)
	path += ".py"
	return filepath.Join(ytc.sdkRoot, "src/python_testing", path)
}

func getTestName(test *testplan.Test) string {
	return "TC_" + test.ID
}

func isExcludedConformance(c conformance.Conformance) bool {
	if conformance.IsDeprecated(c) {
		return true
	}
	if conformance.IsDisallowed(c) {
		return true
	}
	if conformance.IsProvisional(c) {
		return true
	}
	return false
}
