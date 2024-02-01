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
		sdkRoot, _ := cmd.Flags().GetString("sdkRoot")
		return compare.Compare(context.Background(), specRoot, sdkRoot, common.AsciiDocAttributes(cmd))
	},
}

func init() {
	Command.Flags().String("specRoot", "", "the src root of your clone of CHIP-Specifications/connectedhomeip-spec")
	Command.Flags().String("sdkRoot", "", "the src root of your clone of project-chip/connectedhomeip")
	_ = Command.MarkFlagRequired("specRoot")
	_ = Command.MarkFlagRequired("sdkRoot")
}
