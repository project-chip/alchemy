package cli

import (
	"fmt"
	"log/slog"

	"github.com/project-chip/alchemy/cmd/common"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/mle"
)

type MLE struct {
	common.ASCIIDocAttributes  `embed:""`
	pipeline.ProcessingOptions `embed:""`
	spec.ParserOptions         `embed:""`
}

func (m *MLE) Run(cc *Context) (err error) {
	builderOptions := []spec.BuilderOption{spec.IgnoreHierarchy(true)}

	var specification *spec.Specification
	specification, _, err = spec.Parse(cc, m.ParserOptions, m.ProcessingOptions, builderOptions, m.ASCIIDocAttributes.ToList())
	if err != nil {
		return err
	}

	var violations map[string][]spec.Violation
	violations, err = mle.Process(m.ParserOptions.Root, specification)

	for _, vv := range violations {
		for _, v := range vv {
			source := fmt.Sprintf("%s:%d", v.Path, v.Line)
			slog.Warn(v.Text, slog.String("source", source))
		}
	}

	return
}
