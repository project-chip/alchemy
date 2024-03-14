//go:build db

package cmd

import (
	"os"

	"github.com/hasty/alchemy/cmd/database"
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
	verbose, _ := rootCmd.Flags().GetBool("verbose")
	if verbose {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.ErrorLevel)
	}

	commandFound := false
	for _, a := range os.Args[1:] {
		if a == "db" {
			commandFound = true
			break
		}
	}

	if !commandFound {
		args := append([]string{"db"}, os.Args[1:]...)
		rootCmd.SetArgs(args)
	}

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().Bool("verbose", false, "display verbose information")
	rootCmd.PersistentFlags().StringSliceP("attribute", "a", []string{}, "attribute for pre-processing asciidoc; this flag can be provided more than once")

	rootCmd.AddCommand(database.Command)

}
