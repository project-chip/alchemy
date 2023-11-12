package cmd

import (
	"os"

	"github.com/hasty/alchemy/cmd/compare"
	"github.com/hasty/alchemy/cmd/database"
	"github.com/hasty/alchemy/cmd/disco"
	"github.com/hasty/alchemy/cmd/dump"
	"github.com/hasty/alchemy/cmd/format"
	"github.com/hasty/alchemy/cmd/zap"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:          "alchemy",
	Short:        "",
	Long:         ``,
	SilenceUsage: true,
}

func Execute() {
	logrus.SetLevel(logrus.ErrorLevel)

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolP("dryrun", "d", false, "whether or not to actually output files")
	rootCmd.PersistentFlags().Bool("serial", false, "process files one-by-one")
	rootCmd.PersistentFlags().StringSliceP("attribute", "a", []string{}, "attribute for pre-processing asciidoc; this flag can be provided more than once")

	rootCmd.AddCommand(format.Command)
	rootCmd.AddCommand(disco.Command)
	rootCmd.AddCommand(zap.Command)
	rootCmd.AddCommand(database.Command)
	rootCmd.AddCommand(compare.Command)
	rootCmd.AddCommand(conformanceCommand)
	rootCmd.AddCommand(dump.Command)
}
