package common

import (
	"strings"

	"github.com/hasty/adoc/asciidoc"
	"github.com/spf13/cobra"
)

func ASCIIDocAttributes(cmd *cobra.Command) (settings []asciidoc.AttributeName) {
	attributes, _ := cmd.Flags().GetStringSlice("attribute")
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
