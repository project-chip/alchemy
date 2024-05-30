package tests

import (
	"testing"
)

var inlineTests = parseTests{
	{"inline admonition", "inline_admonition.adoc", inlineAdmonition},
	{"inline image", "inline_image.adoc", inlineImage},
}

func TestInline(t *testing.T) {
	inlineTests.run(t)
}
