package pipeline

import "github.com/spf13/cobra"

type Options struct {
	Serial     bool
	NoProgress bool
}

func Flags(cmd *cobra.Command) (options Options) {
	options.Serial, _ = cmd.Flags().GetBool("serial")
	return
}
