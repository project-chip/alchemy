package tests

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/asciidoc/parse"
	"github.com/sanity-io/litter"
)

type parseTest struct {
	name   string
	input  string
	output *asciidoc.Document
}

func (pt *parseTest) run() error {
	in, err := os.ReadFile(pt.input)
	if err != nil {
		return fmt.Errorf("error reading %s: %v", pt.input, err)
	}
	out, err := parse.Reader("", bytes.NewReader(in))
	if err != nil {
		return fmt.Errorf("error parsing %s: %v", pt.name, err)
	}
	if !out.Equals(pt.output) {
		return fmt.Errorf("unexpected result for test %s: %s expected %s", pt.name, cleanObjectDump(litter.Sdump(out)), cleanObjectDump(litter.Sdump(pt.output)))
	}
	return nil
}

func cleanObjectDump(s string) string {
	return strings.ReplaceAll(s, "[github.com/project-chip/alchemy/", "[")
}

type parseTests []parseTest

func (pts parseTests) run(t *testing.T) {
	for _, pt := range pts {
		err := pt.run()
		if err != nil {
			t.Errorf("test %s failed: %v", pt.name, err)
		}
	}
}
