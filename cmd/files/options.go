package files

import "github.com/spf13/cobra"

type Options struct {
	Serial bool
	DryRun bool
}

func Flags(cmd *cobra.Command) (options Options) {
	options.DryRun, _ = cmd.Flags().GetBool("dryrun")
	options.Serial, _ = cmd.Flags().GetBool("serial")
	return
}
