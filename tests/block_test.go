package tests

import (
	"testing"
)

var blockTests = parseTests{
	{"block attributes", "block_attributes.adoc", blockAttributes, nil},
	{"block comment", "block_comment.adoc", blockComment, nil},
	{"block image", "block_image.adoc", blockImage, nil},
}

func TestBlockElements(t *testing.T) {
	blockTests.run(t)
}
