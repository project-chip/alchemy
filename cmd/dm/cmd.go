package dm

import (
	"context"

	"github.com/hasty/alchemy/cmd/common"
	"github.com/hasty/alchemy/cmd/files"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "dm",
	Short: "transmute the Matter spec into data model XML",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		specRoot, _ := cmd.Flags().GetString("specRoot")
		zclRoot, _ := cmd.Flags().GetString("zclRoot")
		return Render(context.Background(), specRoot, zclRoot, files.Flags(cmd), args, common.AsciiDocAttributes(cmd))
	},
}

func init() {
	Command.Flags().String("specRoot", "", "the root of your clone of CHIP-Specifications/connectedhomeip-spec")
	Command.Flags().String("zclRoot", "", "the root of your clone of project-chip/connectedhomeip")
	_ = Command.MarkFlagRequired("specRoot")
	_ = Command.MarkFlagRequired("zclRoot")
}
