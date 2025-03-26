package files

import (
	"github.com/spf13/pflag"
)

type Options struct {
	DryRun bool
	Patch  bool
}

func Flags(flags *pflag.FlagSet) (options Options) {
	options.Patch, _ = flags.GetBool("patch")
	options.DryRun, _ = flags.GetBool("dryrun")
	return
}
