package zap

import (
	"context"

	"github.com/hasty/alchemy/cmd/common"
	"github.com/hasty/alchemy/cmd/files"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "zap",
	Short: "transmute the Matter spec into ZAP templates",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		specRoot, _ := cmd.Flags().GetString("specRoot")
		zclRoot, _ := cmd.Flags().GetString("zclRoot")
		options := migrateOptions{
			filesOptions:  files.Flags(cmd),
			asciiSettings: common.AsciiDocAttributes(cmd),
		}
		return Migrate(context.Background(), specRoot, zclRoot, args, options)
	},
}

func init() {
	Command.Flags().String("specRoot", "", "the root of your clone of CHIP-Specifications/connectedhomeip-spec")
	Command.Flags().String("zclRoot", "", "the root of your clone of project-chip/connectedhomeip")
	Command.MarkFlagRequired("specRoot")
	Command.MarkFlagRequired("zclRoot")
}
