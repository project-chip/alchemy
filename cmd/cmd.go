package cmd

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/alecthomas/kong"
	"github.com/project-chip/alchemy/cmd/cli"
	"github.com/project-chip/alchemy/config"
)

type globalFlags struct {
	Verbose              bool   `default:"false" help:"display verbose information" group:"Logging:"`
	LogLevel             string `default:"info" aliases:"loglevel" help:"changes level of log; 'debug', 'info', 'warn' or 'error'" group:"Logging:"`
	Log                  string `default:"console"  help:"changes format of log; 'console' or 'json'" group:"Logging:"`
	SuppressVersionCheck bool   `default:"false" aliases:"suppressVersionCheck" hidden:""`
	ErrorExitCode        bool   `default:"false" aliases:"errorExitCode" hidden:""`

	Version kong.VersionFlag ` help:"display version"`
}

func Execute() {

	if len(os.Args) < 2 {
		os.Args = append(os.Args, "--help")
	}

	k := kong.Parse(&commands,
		kong.Name("alchemy"),
		kong.Description("A transmuter of Matter"),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact:   true,
			FlagsLast: true,
		}),
		kong.UsageOnError(),
		kong.Vars{"version": fmt.Sprintf("version: %s", config.Version())})

	alchemy := cli.Alchemy{
		Context: context.Background(),
		Kong:    k,
	}

	configureLogging()

	versionChan := make(chan string, 1)
	if !commands.SuppressVersionCheck {
		go checkVersion(alchemy, versionChan)
	}

	err := k.Run(&alchemy)

	if err != nil {
		handleError(k, err)
	}

	if logHadErrors {
		os.Exit(1)
	}

	if !commands.SuppressVersionCheck {
		select {
		case version := <-versionChan:
			compareVersion(version)
		default:
		}
	}

}

var logHadErrors bool

type errorHandler struct {
	slog.Handler
}

func (er *errorHandler) Handle(cxt context.Context, record slog.Record) error {
	if !logHadErrors && record.Level >= slog.LevelError {
		logHadErrors = true
	}
	return er.Handler.Handle(cxt, record)
}
