package tests

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/asciidoc/parse"
	"github.com/project-chip/alchemy/errata"
	"github.com/sanity-io/litter"
)

type parseTest struct {
	name   string
	input  string
	output *asciidoc.Document

	parseOutput *asciidoc.Document
}

func (pt *parseTest) run() error {
	in, err := os.ReadFile(pt.input)
	if err != nil {
		return fmt.Errorf("error reading %s: %v", pt.input, err)
	}
	err = pt.testParser(in)
	if err != nil {
		return err
	}
	if pt.parseOutput != nil {
		err = pt.testInlineParser(in)
		if err != nil {
			return err
		}
	}
	return nil
}

func (pt *parseTest) testParser(in []byte) error {
	out, err := parse.Reader("", bytes.NewReader(in))
	if err != nil {
		return fmt.Errorf("error parsing %s: %v", pt.name, err)
	}
	if !out.Equals(pt.output) {
		return fmt.Errorf("unexpected result for test %s: %s expected %s", pt.name, cleanObjectDump(litter.Sdump(out)), cleanObjectDump(litter.Sdump(pt.output)))
	}
	return nil
}

func (pt *parseTest) testInlineParser(in []byte) error {
	path, err := asciidoc.NewPath(pt.name, ".")
	if err != nil {
		return err
	}
	pc := &preparseContext{
		docPath: path,
	}
	out, err := parse.Inline(pc, "", bytes.NewReader(in))
	if err != nil {
		return fmt.Errorf("error parsing %s: %v", pt.name, err)
	}
	if !out.Equals(pt.parseOutput) {
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

type preparseContext struct {
	docPath    asciidoc.Path
	rootPath   string
	attributes map[string]any
	counters   map[string]*parse.CounterState
}

func (ac *preparseContext) IsSet(name string) bool {
	_, ok := ac.attributes[name]
	return ok
}

func (ac *preparseContext) Get(name string) any {
	return ac.attributes[name]
}

func (ac *preparseContext) Set(name string, value any) {
	if ac.attributes == nil {
		ac.attributes = make(map[string]any)
	}
	ac.attributes[name] = value
}

func (ac *preparseContext) Unset(name string) {
	delete(ac.attributes, name)
}

func (ac *preparseContext) GetCounterState(name string, initialValue string) (*parse.CounterState, error) {
	if ac.counters == nil {
		ac.counters = make(map[string]*parse.CounterState)
	}
	cc, ok := ac.counters[name]
	if ok {
		return cc, nil
	}
	cc = &parse.CounterState{}
	ac.counters[name] = cc
	switch len(initialValue) {
	case 0:
		cc.Value = 1
		cc.CounterType = parse.CounterTypeInteger
	case 1:
		r := initialValue[0]
		if r >= 'a' && r <= 'z' {
			cc.Value = int(r) - int('a')
			cc.CounterType = parse.CounterTypeLowerCase
		} else if r >= 'A' && r <= 'Z' {
			cc.Value = int(r) - int('A')
			cc.CounterType = parse.CounterTypeUpperCase
		} else {
			var err error
			cc.Value, err = strconv.Atoi(initialValue)
			if err != nil {
				return nil, err
			}
			cc.CounterType = parse.CounterTypeInteger
		}
	default:
		var err error
		cc.Value, err = strconv.Atoi(initialValue)
		if err != nil {
			return nil, err
		}
		cc.CounterType = parse.CounterTypeInteger
	}
	return cc, nil

}

func (ac *preparseContext) ResolvePath(path string) (asciidoc.Path, error) {
	linkPath := filepath.Join(ac.docPath.Dir(), path)
	return asciidoc.NewPath(linkPath, ac.rootPath)
}

func (ac *preparseContext) ShouldIncludeFile(path asciidoc.Path) bool {
	return errata.GetSpec(path.Relative).UtilityInclude
}
