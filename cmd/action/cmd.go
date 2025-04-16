package action

import (
	"github.com/project-chip/alchemy/cmd/action/disco"
	"github.com/project-chip/alchemy/cmd/action/zap"
)

type Action struct {
	Disco disco.Command `cmd:"" help:"GitHub action for Matter spec documents"`
	ZAP   zap.Command   `cmd:"" help:"GitHub action for Matter SDK ZAP XML"`
}
