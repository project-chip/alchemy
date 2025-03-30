package pipeline

import (
	"github.com/spf13/pflag"
)

type Options struct {
	Serial     bool
	NoProgress bool
}

func PipelineOptions(flags *pflag.FlagSet) (options Options) {
	options.Serial, _ = flags.GetBool("serial")
	options.NoProgress, _ = flags.GetBool("hide-progress-bar")
	return
}

func Flags(flags *pflag.FlagSet) {
	flags.Bool("serial", false, "process files one-by-one")
	flags.Bool("hide-progress-bar", false, "hide the progress bar")
}
