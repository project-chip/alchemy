//go:build db

package cmd

import (
	"github.com/project-chip/alchemy/cmd/database"
)

var commands struct {
	DB database.Command `cmd:"" default:"1" help:"run a local MySQL DB containing the contents of the Matter spec or the ZAP templates"`

	globalFlags `embed:""`
}
