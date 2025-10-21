package dump

import (
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/matter"
)

type Reader interface {
	asciidoc.Reader
	SectionType(section *asciidoc.Section) matter.Section
}

type dumpInfoCache struct {
	asciidoc.Reader
	types    map[*asciidoc.Section]matter.Section
	names    map[*asciidoc.Section]string
	docTypes map[*asciidoc.Document]matter.DocType
}

func (dic *dumpInfoCache) SectionName(section *asciidoc.Section) string {
	name, ok := dic.names[section]
	if ok {
		return name
	}
	var sb strings.Builder
	for _, e := range section.Title.Children() {
		switch e := e.(type) {
		case *asciidoc.String:
			sb.WriteString(e.Value)
		}
	}
	dic.SetSectionName(section, sb.String())
	return sb.String()
}

func (dic *dumpInfoCache) SetSectionName(section *asciidoc.Section, name string) {
	dic.names[section] = name
}

func (dic *dumpInfoCache) SectionType(section *asciidoc.Section) matter.Section {
	return dic.types[section]
}

func (dic *dumpInfoCache) SetSectionType(section *asciidoc.Section, sectionType matter.Section) {
	dic.types[section] = sectionType
}

func (dic *dumpInfoCache) DocType(document *asciidoc.Document) (matter.DocType, error) {
	return dic.docTypes[document], nil
}

func (dic *dumpInfoCache) SetDocType(document *asciidoc.Document, dt matter.DocType) {
	dic.docTypes[document] = dt
}

func (dic *dumpInfoCache) Parents(document *asciidoc.Document) []*asciidoc.Document {
	return nil
}

func (dic *dumpInfoCache) ErrataForPath(docPath string) *errata.Errata {
	return errata.DefaultErrata
}

func newDumpInfoCache(reader asciidoc.Reader) *dumpInfoCache {
	return &dumpInfoCache{
		Reader:   reader,
		types:    make(map[*asciidoc.Section]matter.Section),
		names:    make(map[*asciidoc.Section]string),
		docTypes: map[*asciidoc.Document]matter.DocType{},
	}
}
