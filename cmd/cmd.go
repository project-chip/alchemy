package cmd

import (
	"context"
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/lmittmann/tint"
	"github.com/project-chip/alchemy/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:          "alchemy",
	Short:        "",
	Long:         ``,
	SilenceUsage: true,
	Version:      config.Version(),
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

	if errored {
		os.Exit(1)
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
	flags := rootCmd.PersistentFlags()
	flags.Bool("verbose", false, "display verbose information")
	flags.String("log", "console", "changes format of log; 'console' or 'json'")
	flags.String("loglevel", "info", "changes level of log; 'debug', 'info', 'warn' or 'error'")
	flags.Bool("suppressVersionCheck", false, "")
	flags.MarkHidden("suppressVersionCheck")
	flags.Bool("errorExitCode", false, "")
	flags.MarkHidden("errorExitCode")
	rootCmd.SetVersionTemplate(`{{printf "version: %s" .Version}}
`)
	rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		flags := rootCmd.PersistentFlags()
		verbose, _ := flags.GetBool("verbose")
		level := slog.LevelInfo
		loglevel, _ := flags.GetString("loglevel")
		switch strings.ToLower(loglevel) {
		case "error":
			level = slog.LevelError
		case "warn":
			level = slog.LevelWarn
		case "info":
			level = slog.LevelInfo
		case "debug":
			level = slog.LevelDebug
		}
		if verbose {
			level = slog.LevelDebug
		}
		var handler slog.Handler
		logType, _ := flags.GetString("log")
		switch logType {
		case "json":
			handler = slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: level})
		default:
			handler = tint.NewHandler(os.Stderr, &tint.Options{
				Level:      level,
				TimeFormat: time.StampMilli,
			})
		}
		errorExitCode, _ := flags.GetBool("errorExitCode")
		if errorExitCode {
			handler = &errorHandler{Handler: handler}
		}
		slog.SetDefault(slog.New(handler))
		suppressVersionCheck, _ = flags.GetBool("suppressVersionCheck")
	}

}

var errored bool

type errorHandler struct {
	slog.Handler
}

func (er *errorHandler) Handle(cxt context.Context, record slog.Record) error {
	if !errored && record.Level >= slog.LevelError {
		errored = true
	}
	return er.Handler.Handle(cxt, record)
}
