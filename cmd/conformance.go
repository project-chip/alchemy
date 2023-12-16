package cmd

import (
	"fmt"
	"os"

	"github.com/hasty/alchemy/matter/conformance"
	"github.com/spf13/cobra"
)

var conformanceCommand = &cobra.Command{
	Use:   "conformance",
	Short: "test conformance values",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		if len(args) == 0 {
			return cmd.Usage()
		}
		c := conformance.ParseConformance(args[0])
		fmt.Fprintf(os.Stdout, "description: %s\n", c.String())
		if len(args) > 1 {
			var cxt conformance.ConformanceContext
			cxt.Values = make(map[string]any)
			for _, arg := range args[1:] {
				cxt.Values[arg] = true
			}
			crm, err := c.Eval(cxt)
			if err != nil {
				return err
			}
			fmt.Fprintf(os.Stdout, "conformance: %v\n", crm)
		}
		return nil
	},
}
