package compare

import (
	"context"

	"github.com/hasty/alchemy/cmd/common"
	"github.com/hasty/alchemy/compare"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "compare",
	Short: "compare the spec to zap-templates and output a JSON diff",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		specRoot, _ := cmd.Flags().GetString("specRoot")
		zclRoot, _ := cmd.Flags().GetString("zclRoot")
		return compare.Compare(context.Background(), specRoot, zclRoot, common.AsciiDocAttributes(cmd))
	},
}

func init() {
	Command.Flags().String("specRoot", "", "the src root of your clone of CHIP-Specifications/connectedhomeip-spec")
	Command.Flags().String("zclRoot", "", "the src root of your clone of project-chip/connectedhomeip")
	Command.MarkFlagRequired("specRoot")
	Command.MarkFlagRequired("zclRoot")
}
