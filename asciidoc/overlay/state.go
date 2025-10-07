package overlay

import (
	"strconv"

	"github.com/project-chip/alchemy/asciidoc"
)

type conditionalBlock struct {
	open     asciidoc.Element
	suppress bool
}

type overlayState struct {
	attributes map[string]any
	counters   map[string]*asciidoc.CounterState

	overlays elementOverlays

	lastSection      *asciidoc.Section
	lastSectionLevel int
}

type overlayFileState struct {
	state *overlayState

	docPath  asciidoc.Path
	rootDoc  *asciidoc.Document
	rootPath string

	sectionLevelOffset int
}

func (ofs *overlayFileState) IsSet(name string) bool {
	_, ok := ofs.state.attributes[name]
	return ok
}

func (ofs *overlayFileState) Get(name string) any {
	return ofs.state.attributes[name]
}

func (ofs *overlayFileState) Set(name string, value any) {
	if ofs.state.attributes == nil {
		ofs.state.attributes = make(map[string]any)
	}
	ofs.state.attributes[name] = value
}

func (ofs *overlayFileState) Unset(name string) {
	delete(ofs.state.attributes, name)
}

func (ofs *overlayFileState) GetCounterState(name string, initialValue string) (*asciidoc.CounterState, error) {
	if ofs.state.counters == nil {
		ofs.state.counters = make(map[string]*asciidoc.CounterState)
	}
	cc, ok := ofs.state.counters[name]
	if ok {
		return cc, nil
	}
	cc = &asciidoc.CounterState{}
	ofs.state.counters[name] = cc
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

func (ofs *overlayFileState) ShouldIncludeFile(path asciidoc.Path) bool {
	switch path.Relative {
	case "templates/DiscoBallCluster.adoc", "templates/DiscoBallDeviceType.adoc":
		return false
	default:
		return true
	}
}
