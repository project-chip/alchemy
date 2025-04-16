package common

import (
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
)

type ASCIIDocAttributes struct {
	Attribute []string `help:"attribute for pre-processing AsciiDoc; this flag can be provided more than once" group:"Spec:"`
}

func (aa *ASCIIDocAttributes) ToList() (settings []asciidoc.AttributeName) {
	for _, a := range aa.Attribute {
		if len(a) == 0 {
			continue
		}
		for _, set := range strings.Split(a, ",") {
			settings = append(settings, asciidoc.AttributeName(strings.TrimSpace(set)))
		}
	}
	return
}
