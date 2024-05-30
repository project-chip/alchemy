package tests

import (
	"testing"
)

var blockTests = parseTests{
	{"block attributes", "block_attributes.adoc", blockAttributes},
	{"block comment", "block_comment.adoc", blockComment},
	{"block image", "block_image.adoc", blockImage},
}

func TestBlocks(t *testing.T) {
	blockTests.run(t)
}
