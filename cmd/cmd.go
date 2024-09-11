package cmd

import (
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:          "alchemy",
	Short:        "",
	Long:         ``,
	SilenceUsage: true,
}

var defaultCommand string

func Execute() {
	if len(defaultCommand) > 0 {
		cmd, _, err := rootCmd.Find(os.Args[1:])
		if err != nil || cmd.Use == rootCmd.Use {
			rootCmd.SetArgs(append([]string{defaultCommand}, os.Args[1:]...))
		}
	}

	err := rootCmd.Execute()
	if err != nil {
		handleError(err)
	}
}

func init() {
	rootCmd.PersistentFlags().Bool("verbose", false, "display verbose information")
	rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		verbose, _ := rootCmd.Flags().GetBool("verbose")
		level := slog.LevelInfo
		if verbose {
			level = slog.LevelDebug
		}
		slog.SetDefault(slog.New(tint.NewHandler(os.Stderr, &tint.Options{
			Level:      level,
			TimeFormat: time.StampMilli,
		})))
	}
}
