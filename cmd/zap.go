package cmd

import (
	"context"

	"github.com/hasty/alchemy/cmd/zap"
	"github.com/spf13/cobra"
)

var zapCommand = &cobra.Command{
	Use:   "zap",
	Short: "transmute the Matter spec into ZAP templates",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		specRoot, _ := cmd.Flags().GetString("specRoot")
		zclRoot, _ := cmd.Flags().GetString("zclRoot")
		return zap.Migrate(context.Background(), specRoot, zclRoot, getFilesOptions(), getAsciiAttributes())
	},
}

func init() {
	rootCmd.AddCommand(zapCommand)
	zapCommand.Flags().String("specRoot", "", "the root of your clone of CHIP-Specifications/connectedhomeip-spec")
	zapCommand.Flags().String("zclRoot", "", "the root of your clone of project-chip/connectedhomeip")
	zapCommand.MarkFlagRequired("specRoot")
	zapCommand.MarkFlagRequired("zclRoot")

}
