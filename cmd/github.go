//go:build github

package cmd

import (
	"github.com/alecthomas/kong"
	"github.com/project-chip/alchemy/cmd/disco"
	"github.com/project-chip/alchemy/cmd/zap"
	"github.com/sethvargo/go-githubactions"
)

var commands struct {
	Disco   disco.Command `cmd:"" default:"1" help:"disco ball Matter spec documents specified by the filename_pattern" group:"Spec Commands:"`
	ZAP     zap.Command   `cmd:"" help:"transmute the Matter spec into ZAP templates, optionally filtered to the files specified by filename_pattern" group:"SDK Commands:"`
	Version Version       `cmd:"" hidden:"" name:"version" help:"display version number"`

	globalFlags `embed:""`
}

func handleError(ctx *kong.Context, err error) {
	githubactions.Fatalf("failed disco ball action: %v\n", err)
}
