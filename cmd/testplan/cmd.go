package testplan

import (
	"context"

	"github.com/hasty/alchemy/cmd/common"
	"github.com/hasty/alchemy/cmd/files"
	"github.com/hasty/alchemy/testplan"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "testplan",
	Short: "create an initial test plan from the spec",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		specRoot, _ := cmd.Flags().GetString("specRoot")
		sdkRoot, _ := cmd.Flags().GetString("testRoot")
		options := testplan.Options{
			Files: files.Flags(cmd),
			Ascii: common.AsciiDocAttributes(cmd),
		}
		options.Overwrite, _ = cmd.Flags().GetBool("overwrite")
		return testplan.Generate(context.Background(), specRoot, sdkRoot, args, options)
	},
}

func init() {
	Command.Flags().String("specRoot", "", "the root of your clone of CHIP-Specifications/connectedhomeip-spec")
	Command.Flags().String("testRoot", "", "the root of your clone of CHIP-Specifications/chip-test-plans")
	Command.Flags().Bool("overwrite", false, "overwrite existing test plans")
	_ = Command.MarkFlagRequired("specRoot")
	_ = Command.MarkFlagRequired("testRoot")
}
