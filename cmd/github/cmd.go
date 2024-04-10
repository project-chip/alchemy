package github

import "github.com/spf13/cobra"

var Command = &cobra.Command{
	Use:   "github",
	Short: "GitHub action for Matter spec documents",
	Long:  ``,
	RunE:  Action,
}
