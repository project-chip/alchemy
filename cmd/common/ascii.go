package common

import (
	"strings"

	"github.com/hasty/adoc/elements"
	"github.com/spf13/cobra"
)

func ASCIIDocAttributes(cmd *cobra.Command) (settings []elements.Attribute) {
	attributes, _ := cmd.Flags().GetStringSlice("attribute")
	for _, a := range attributes {
		if len(a) == 0 {
			continue
		}
		for _, set := range strings.Split(a, ",") {
			settings = append(settings, elements.NewNamedAttribute(strings.TrimSpace(set), true))
		}
	}
	return
}
