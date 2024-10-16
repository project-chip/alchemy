package action

import (
	"github.com/project-chip/alchemy/cmd/action/disco"
	"github.com/project-chip/alchemy/cmd/action/zap"
	"github.com/spf13/cobra"
)

var Disco = &cobra.Command{
	Use:   "disco",
	Short: "GitHub action for Matter spec documents",
	Long:  ``,
	RunE:  disco.Ball,
}

var ZAP = &cobra.Command{
	Use:   "zap",
	Short: "GitHub action for Matter SDK ZAP XML",
	Long:  ``,
	RunE:  zap.CheckZAPXML,
}
