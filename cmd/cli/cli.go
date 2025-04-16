package cli

import (
	"context"

	"github.com/alecthomas/kong"
)

type Alchemy struct {
	context.Context

	Kong *kong.Context
}
