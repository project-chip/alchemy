package pipeline

import (
	"github.com/spf13/pflag"
)

type Options struct {
	Serial     bool
	NoProgress bool
}

func Flags(flags *pflag.FlagSet) (options Options) {
	options.Serial, _ = flags.GetBool("serial")
	return
}
