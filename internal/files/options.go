package files

import "github.com/spf13/cobra"

type Options struct {
	DryRun bool
	Patch  bool
}

func Flags(cmd *cobra.Command) (options Options) {
	options.Patch, _ = cmd.Flags().GetBool("patch")
	options.DryRun, _ = cmd.Flags().GetBool("dryrun")
	return
}
