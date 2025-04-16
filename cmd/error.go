//go:build !github

package cmd

import (
	"github.com/alecthomas/kong"
)

func handleError(ctx *kong.Context, err error) {
	//fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	ctx.FatalIfErrorf(err)
}
