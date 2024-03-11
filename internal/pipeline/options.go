package pipeline

import "github.com/spf13/cobra"

type Options struct {
	Serial bool
}

func Flags(cmd *cobra.Command) (options Options) {
	options.Serial, _ = cmd.Flags().GetBool("serial")
	return
}

func (o Options) DefaultProcessorType() ProcessorType {
	if o.Serial {
		return ProcessorTypeSerial
	}
	return ProcessorTypeParallel
}
