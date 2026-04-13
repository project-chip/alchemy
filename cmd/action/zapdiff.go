package action

import (
	"log/slog"

	"github.com/project-chip/alchemy/cmd/cli"
)

type ZAPDiff struct {
}

func (z *ZAPDiff) Run(cc *cli.Context) (err error) {
	slog.Info("zapdiff test string")
	return nil
}
