//go:build !db && !github

package cmd

import (
	"github.com/hasty/alchemy/cmd/compare"
	"github.com/hasty/alchemy/cmd/disco"
	"github.com/hasty/alchemy/cmd/dm"
	"github.com/hasty/alchemy/cmd/dump"
	"github.com/hasty/alchemy/cmd/format"
	"github.com/hasty/alchemy/cmd/testplan"
	"github.com/hasty/alchemy/cmd/zap"
)

func init() {
	rootCmd.PersistentFlags().BoolP("dryrun", "d", false, "whether or not to actually output files")
	rootCmd.PersistentFlags().BoolP("patch", "p", false, "generate patch file")
	rootCmd.PersistentFlags().Bool("serial", false, "process files one-by-one")
	rootCmd.PersistentFlags().StringSliceP("attribute", "a", []string{}, "attribute for pre-processing asciidoc; this flag can be provided more than once")

	rootCmd.AddCommand(format.Command)
	rootCmd.AddCommand(disco.Command)
	rootCmd.AddCommand(zap.Command)
	rootCmd.AddCommand(compare.Command)
	rootCmd.AddCommand(conformanceCommand)
	rootCmd.AddCommand(dump.Command)
	rootCmd.AddCommand(dm.Command)
	rootCmd.AddCommand(testplan.Command)
}
