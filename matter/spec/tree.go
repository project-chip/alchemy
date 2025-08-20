package spec

import (
	"context"
	"log/slog"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal/pipeline"
)

type LibraryBuilder struct {
	specRoot string
}

func NewLibraryBuilder(specRoot string) *LibraryBuilder {
	b := &LibraryBuilder{
		specRoot: specRoot,
	}
	return b
}

func (dg LibraryBuilder) Name() string {
	return "Grouping spec documents"
}

func (dg *LibraryBuilder) Process(cxt context.Context, inputs []*pipeline.Data[*asciidoc.Document]) (outputs []*pipeline.Data[*Library], err error) {
	/*docs := make([]*asciidoc.Document, 0, len(inputs))
	for _, i := range inputs {
		docs = append(docs, i.Content)
	}

	buildTree(asciidoc.RawReader, dg.specRoot, docs)

	docGroups := buildDocumentGroups(docs)
	for _, g := range docGroups {
		outputs = append(outputs, pipeline.NewData(g.Root.Path.Relative, g))
	}*/

	docCache := cacheFromPipeline(dg.specRoot, inputs)

	for _, docRoot := range errata.DocRoots {
		root, ok := docCache.cache.Load(docRoot)
		if !ok {
			slog.Warn("doc root not found", "root", docRoot)
			continue
		}
		outputs = append(outputs, pipeline.NewData(root.Path.Relative, NewLibrary(root, docCache)))
	}
	return
}

/*func buildTree(reader asciidoc.Reader, specRoot string, docs []*asciidoc.Document) error {

	tree := make(map[*asciidoc.Document][]*asciidoc.FileInclude)
	docPaths := make(map[string]*asciidoc.Document)

	for _, doc := range docs {

		path := doc.Path
		docPaths[path.Absolute] = doc

		parse.Search(doc, reader, doc, doc.Children(), func(doc *asciidoc.Document, link *asciidoc.FileInclude, parent asciidoc.ParentElement, index int) parse.SearchShould {
			tree[doc] = append(tree[doc], link)
			return parse.SearchShouldContinue
		})
	}

	for doc, children := range tree {
		for _, link := range children {
			var p strings.Builder
			buildDataTypeString(reader, doc, link.Elements, &p)
			linkFullPath := filepath.Join(doc.Path.Dir(), p.String())
			linkPath, err := asciidoc.NewPath(linkFullPath, specRoot)
			if err != nil {
				return err
			}
			slog.Debug("Link path", log.Path("from", doc.Path), slog.String("to", p.String()), log.Path("linkPath", linkPath))
			if cd, ok := docPaths[linkPath.Absolute]; ok {
				cd.addParent(doc)
				doc.addChild(cd)
			} else {
				if strings.HasPrefix(linkPath.Relative, "src/") {
					slog.Warn("unknown child path", log.Element("parent", doc.Path, link), "child", linkPath.Relative)
				}
			}
		}
	}
	return nil
}

func buildDataTypeString(reader asciidoc.Reader, d *asciidoc.Document, cellElements asciidoc.Elements, sb *strings.Builder) (source asciidoc.Element) {
	for _, el := range cellElements {
		switch v := el.(type) {
		case *asciidoc.String:
			sb.WriteString(v.Value)

		case *asciidoc.SpecialCharacter:
		case *asciidoc.Paragraph:
			source = buildDataTypeString(reader, d, v.Children(), sb)
		default:
			slog.Warn("unknown path value element", log.Element("source", d.Path, el), "type", fmt.Sprintf("%T", v))
		}
	}
	return
}

func dumpTree(r *asciidoc.Document, indent int) {
	fmt.Print(strings.Repeat("\t", indent))
	fmt.Printf("%s (%s)\n", r.Path.Absolute, r.Path.Relative)
	for _, c := range r.children {
		dumpTree(c, indent+1)
	}
}

func buildDocumentGroups(docs []*Doc) (docGroups []*Library) {
	for _, d := range docs {
		if len(d.parents) > 0 {
			continue
		}

		var isDocRoot bool
		path := d.Path.Relative
		for _, docRoot := range errata.DocRoots {
			if strings.EqualFold(path, docRoot) {
				isDocRoot = true
				break
			}
		}

		if !isDocRoot {
			continue
		}

		dg := NewLibrary(d)
		docGroups = append(docGroups, dg)
		setDocGroup(d, dg)
	}
	return
}
*/
