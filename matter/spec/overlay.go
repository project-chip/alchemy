package spec

import (
	"context"
	"path/filepath"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/asciidoc/overlay"
	"github.com/project-chip/alchemy/internal/pipeline"
)

type overlayContext struct {
	specRoot string
	library  *Library
}

func (o *overlayContext) AddDocument(doc *asciidoc.Document) {
	o.library.Docs = append(o.library.Docs, doc)
}

// IncludeFile implements overlay.OverlayContext.
func (o *overlayContext) IncludeFile(path asciidoc.Path, parent asciidoc.Parent) (doc *asciidoc.Document, err error) {
	return o.library.cache.include(path, parent)
}

// IncludePath implements overlay.OverlayContext.
func (o *overlayContext) ResolvePath(root string, path string) (asciidoc.Path, error) {
	linkPath := filepath.Join(root, path)
	return NewSpecPath(linkPath, o.specRoot)
}

// MakeSectionName implements overlay.OverlayContext.
func (o *overlayContext) MakeSectionName(reader asciidoc.Reader, section *asciidoc.Section, variables overlay.Variables) (string, error) {
	var title strings.Builder
	err := buildSectionTitle(variables, section, reader, &title, section.Title...)
	if err != nil {
		return "", err
	}
	return title.String(), nil
}

// SectionLevel implements overlay.OverlayContext.
func (o *overlayContext) SectionLevel(section *asciidoc.Section) int {
	return o.library.SectionLevel(section)
}

// SectionName implements overlay.OverlayContext.
func (o *overlayContext) SectionName(section *asciidoc.Section) string {
	return o.library.SectionName(section)
}

// SetParent implements overlay.OverlayContext.
func (o *overlayContext) SetParent(parent *asciidoc.Document, child *asciidoc.Document) {
	o.library.parents[child] = append(o.library.parents[child], parent)
	o.library.children[parent] = append(o.library.children[parent], child)

}

// SetSectionLevel implements overlay.OverlayContext.
func (o *overlayContext) SetSectionLevel(section *asciidoc.Section, level int) {
	o.library.SetSectionLevel(section, level)
}

// SetSectionName implements overlay.OverlayContext.
func (o *overlayContext) SetSectionName(section *asciidoc.Section, name string) {
	o.library.SetSectionName(section, name)
}

// ShouldIncludeFile implements overlay.OverlayContext.
func (o *overlayContext) ShouldIncludeFile(path asciidoc.Path) bool {
	switch path.Relative {
	case "templates/DiscoBallCluster.adoc",
		"templates/DiscoBallDeviceType.adoc":
		return false
	default:
		return true
	}
}

var _ overlay.OverlayContext = &overlayContext{}

func (library *Library) Preparse(specRoot string, attributes []asciidoc.AttributeName) (err error) {
	cxt := &overlayContext{
		specRoot: specRoot,
		library:  library,
	}
	library.Reader, err = overlay.Build(cxt, library.Root, specRoot, attributes)

	return
}

type LibraryParser struct {
	specRoot string

	attributes []asciidoc.AttributeName
}

func NewLibraryParser(specRoot string, attributes []asciidoc.AttributeName) (*LibraryParser, error) {

	return &LibraryParser{specRoot: specRoot, attributes: attributes}, nil
}

func (r LibraryParser) Name() string {
	return "Building library overlay"
}

func (r LibraryParser) Process(cxt context.Context, input *pipeline.Data[*Library], index int32, total int32) (outputs []*pipeline.Data[*asciidoc.Document], extras []*pipeline.Data[*Library], err error) {
	err = input.Content.Preparse(r.specRoot, r.attributes)
	return
}
