package cmd

import (
	"os"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/hasty/alchemy/cmd/files"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:          "alchemy",
	Short:        "",
	Long:         ``,
	SilenceUsage: true,
}

func Execute() {
	logrus.SetLevel(logrus.ErrorLevel)

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolP("dryrun", "d", false, "whether or not to actually output files")
	rootCmd.PersistentFlags().Bool("serial", false, "process files one-by-one")
	rootCmd.PersistentFlags().StringSliceP("attribute", "a", []string{}, "attribute for pre-processing asciidoc; this flag can be provided more than once")
}

func getOptions() (options []Option) {
	dry, _ := rootCmd.Flags().GetBool("dryrun")
	if dry {
		options = append(options, DryRun(dry))
	}
	serial, _ := rootCmd.Flags().GetBool("serial")
	if dry {
		options = append(options, Serial(serial))
	}

	attributes, _ := rootCmd.Flags().GetStringSlice("attribute")
	if len(attributes) > 0 {
		options = append(options, AsciiAttributes(attributes))
	}

	return
}

func getFilesOptions() (options files.Options) {
	options.DryRun, _ = rootCmd.Flags().GetBool("dryrun")
	options.Serial, _ = rootCmd.Flags().GetBool("serial")
	return
}

func getAsciiAttributes() (settings []configuration.Setting) {
	attributes, _ := rootCmd.Flags().GetStringSlice("attribute")
	for _, a := range attributes {
		if len(a) == 0 {
			continue
		}
		for _, set := range strings.Split(a, ",") {
			settings = append(settings, configuration.WithAttribute(strings.TrimSpace(set), true))
		}
	}
	return
}

func applyOptions(target any, options []Option) (err error) {
	for _, o := range options {
		err = o(target)
		if err != nil {
			return
		}
	}
	return
}
