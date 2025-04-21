package cli

import (
	"context"

	"github.com/alecthomas/kong"
)

type Context struct {
	context.Context

	Kong *kong.Context
}
