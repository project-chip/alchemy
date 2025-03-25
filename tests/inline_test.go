package tests

import (
	"testing"
)

var inlineTests = parseTests{
	{"inline admonition", "inline_admonition.adoc", inlineAdmonition, nil},
	{"inline image", "inline_image.adoc", inlineImage, nil},
}

func TestInline(t *testing.T) {
	inlineTests.run(t)
}
