package common

import (
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/spf13/pflag"
)

func AttributeFlags(flags *pflag.FlagSet) {
	flags.StringSliceP("attribute", "a", []string{}, "attribute for pre-processing asciidoc; this flag can be provided more than once")
}

func ASCIIDocAttributes(flags *pflag.FlagSet) (settings []asciidoc.AttributeName) {
	attributes, _ := flags.GetStringSlice("attribute")
	for _, a := range attributes {
		if len(a) == 0 {
			continue
		}
		for _, set := range strings.Split(a, ",") {
			settings = append(settings, asciidoc.AttributeName(strings.TrimSpace(set)))
		}
	}
	return
}
