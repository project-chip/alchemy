package cmd

import (
	"context"

	"github.com/hasty/alchemy/compare"
	"github.com/spf13/cobra"
)

type zclConparer struct {
	processor
	asciiParser

	serial bool
	dryRun bool
}

var compareCommand = &cobra.Command{
	Use:   "compare",
	Short: "compare the spec to zap-templates and output a JSON diff",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		specRoot, _ := cmd.Flags().GetString("specRoot")
		zclRoot, _ := cmd.Flags().GetString("zclRoot")
		return compare.Compare(context.Background(), specRoot, zclRoot, getAsciiAttributes())
	},
}

func init() {
	rootCmd.AddCommand(compareCommand)
	compareCommand.Flags().String("specRoot", "", "the src root of your clone of CHIP-Specifications/connectedhomeip-spec")
	compareCommand.Flags().String("zclRoot", "", "the src root of your clone of project-chip/connectedhomeip")
	compareCommand.MarkFlagRequired("specRoot")
	compareCommand.MarkFlagRequired("zclRoot")
}
