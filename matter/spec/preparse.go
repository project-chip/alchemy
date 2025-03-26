package spec

import (
	"path/filepath"
	"strconv"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/asciidoc/parse"
	"github.com/project-chip/alchemy/errata"
)

func NewPreparseContext(path asciidoc.Path, specRoot string, attributes ...asciidoc.AttributeName) parse.PreParseContext {

	ac := &preparseContext{
		docPath:  path,
		rootPath: specRoot,
	}

	for _, a := range attributes {
		ac.Set(string(a), nil)
	}
	return ac
}

type preparseContext struct {
	docPath    asciidoc.Path
	rootPath   string
	attributes map[string]any
	counters   map[string]*asciidoc.CounterState
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

func (ac *preparseContext) GetCounterState(name string, initialValue string) (*asciidoc.CounterState, error) {
	if ac.counters == nil {
		ac.counters = make(map[string]*asciidoc.CounterState)
	}
	cc, ok := ac.counters[name]
	if ok {
		return cc, nil
	}
	cc = &asciidoc.CounterState{}
	ac.counters[name] = cc
	switch len(initialValue) {
	case 0:
		cc.Value = 1
		cc.CounterType = asciidoc.CounterTypeInteger
	case 1:
		r := initialValue[0]
		if r >= 'a' && r <= 'z' {
			cc.Value = int(r) - int('a')
			cc.CounterType = asciidoc.CounterTypeLowerCase
		} else if r >= 'A' && r <= 'Z' {
			cc.Value = int(r) - int('A')
			cc.CounterType = asciidoc.CounterTypeUpperCase
		} else {
			var err error
			cc.Value, err = strconv.Atoi(initialValue)
			if err != nil {
				return nil, err
			}
			cc.CounterType = asciidoc.CounterTypeInteger
		}
	default:
		var err error
		cc.Value, err = strconv.Atoi(initialValue)
		if err != nil {
			return nil, err
		}
		cc.CounterType = asciidoc.CounterTypeInteger
	}
	return cc, nil

}

func (ac *preparseContext) ResolvePath(path string) (asciidoc.Path, error) {
	linkPath := filepath.Join(ac.docPath.Dir(), path)
	return NewSpecPath(linkPath, ac.rootPath)
}

func (ac *preparseContext) ShouldIncludeFile(path asciidoc.Path) bool {
	return errata.GetSpec(path.Relative).UtilityInclude
}
