//go:build !db && !github

package cmd

import (
	"github.com/project-chip/alchemy/cmd/compare"
	"github.com/project-chip/alchemy/cmd/disco"
	"github.com/project-chip/alchemy/cmd/dm"
	"github.com/project-chip/alchemy/cmd/dump"
	"github.com/project-chip/alchemy/cmd/format"
	"github.com/project-chip/alchemy/cmd/testplan"
	"github.com/project-chip/alchemy/cmd/testscript"
	"github.com/project-chip/alchemy/cmd/validate"
	"github.com/project-chip/alchemy/cmd/yaml2python"
	"github.com/project-chip/alchemy/cmd/zap"
)

func init() {
	flags := rootCmd.PersistentFlags()
	flags.BoolP("dryrun", "d", false, "whether or not to actually output files")
	flags.BoolP("patch", "p", false, "generate patch file")
	flags.Bool("serial", false, "process files one-by-one")
	flags.StringSliceP("attribute", "a", []string{}, "attribute for pre-processing asciidoc; this flag can be provided more than once")

	rootCmd.AddCommand(format.Command)
	rootCmd.AddCommand(disco.Command)
	rootCmd.AddCommand(zap.Command)
	rootCmd.AddCommand(compare.Command)
	rootCmd.AddCommand(conformanceCommand)
	rootCmd.AddCommand(dump.Command)
	rootCmd.AddCommand(dm.Command)
	rootCmd.AddCommand(testplan.Command)
	rootCmd.AddCommand(testscript.Command)
	rootCmd.AddCommand(validate.Command)
	rootCmd.AddCommand(yaml2python.Command)
}
