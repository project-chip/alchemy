//go:build !db

package cmd

import (
	"fmt"
	"os"

	"github.com/project-chip/alchemy/cmd/cli"
	"github.com/project-chip/alchemy/matter/conformance"
)

type Conformance struct {
	Conformance string   `arg:"" help:"conformance string" required:""`
	Params      []string `arg:"" help:"parameters to use to evaluate conformance" optional:""`
}

func (cmd *Conformance) Run(alchemy *cli.Alchemy) (err error) {
	if len(cmd.Conformance) == 0 {
		// TODO: re-add usage
		return nil
	}
	c := conformance.ParseConformance(cmd.Conformance)
	fmt.Fprintf(os.Stdout, "description: %s\n", c.Description())
	if len(cmd.Params) > 0 {
		var cxt conformance.Context
		cxt.Values = make(map[string]any)
		for _, arg := range cmd.Params {
			cxt.Values[arg] = true
		}
		crm, err := c.Eval(cxt)
		if err != nil {
			return err
		}
		fmt.Fprintf(os.Stdout, "conformance: %v\n", crm)
	}
	return nil
}
