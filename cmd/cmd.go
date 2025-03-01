package cmd

import (
	"context"
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
var suppressVersionCheck bool

func Execute() {
	if len(defaultCommand) > 0 {
		cmd, _, err := rootCmd.Find(os.Args[1:])
		if err != nil || cmd.Use == rootCmd.Use {
			rootCmd.SetArgs(append([]string{defaultCommand}, os.Args[1:]...))
		}
	}

	cxt := context.Background()

	versionChan := make(chan string, 1)
	if !suppressVersionCheck {
		go checkVersion(cxt, versionChan)
	}

	err := rootCmd.ExecuteContext(cxt)
	if err != nil {
		handleError(err)
	}

	if !suppressVersionCheck {
		select {
		case version := <-versionChan:
			compareVersion(version)
		default:
		}
	}
}

func init() {
	rootCmd.PersistentFlags().Bool("verbose", false, "display verbose information")
	rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		verbose, _ := rootCmd.PersistentFlags().GetBool("verbose")
		level := slog.LevelInfo
		if verbose {
			level = slog.LevelDebug
		}
		slog.SetDefault(slog.New(tint.NewHandler(os.Stderr, &tint.Options{
			Level:      level,
			TimeFormat: time.StampMilli,
		})))
		suppressVersionCheck, _ = rootCmd.PersistentFlags().GetBool("suppressVersionCheck")
	}
	rootCmd.PersistentFlags().Bool("suppressVersionCheck", false, "")
	rootCmd.PersistentFlags().MarkHidden("suppressVersionCheck")
}
