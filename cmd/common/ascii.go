package common

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/spf13/cobra"
)

func ASCIIDocAttributes(cmd *cobra.Command) (settings []configuration.Setting) {
	attributes, _ := cmd.Flags().GetStringSlice("attribute")
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
