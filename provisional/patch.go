package provisional

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/asciidoc/render"
	"github.com/project-chip/alchemy/internal/files"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/internal/paths"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
)

func patchProvisional(cxt context.Context, pipelineOptions pipeline.ProcessingOptions, s *spec.Specification, violations map[string][]Violation, writer files.Writer[string]) (err error) {
	var docPaths []string
	for path := range violations {
		docPaths = append(docPaths, path)
	}
	var inputs pipeline.Paths
	inputs, err = pipeline.Start(cxt, paths.NewTargeter(docPaths...))
	if err != nil {
		return err
	}

	docReader, err := spec.NewReader(spec.ParserOptions{Root: s.Root})
	if err != nil {
		return err
	}
	docs, err := pipeline.Parallel(cxt, pipelineOptions, docReader, inputs)
	if err != nil {
		return err
	}

	for path, vs := range violations {
		doc, ok := docs.Load(path)
		if !ok {
			err = fmt.Errorf("failed to load doc")
			return
		}
		for _, v := range vs {
			err = patchViolation(doc.Content, v)
			if err != nil {
				return
			}
		}
	}

	renderDocs := pipeline.NewConcurrentMapPresized[string, *pipeline.Data[*asciidoc.Document]](docs.Size())
	pipeline.Cast(docs, renderDocs)

	renderer := render.NewRenderer()
	var renders pipeline.StringSet
	renders, err = pipeline.Parallel(cxt, pipelineOptions, renderer, renderDocs)
	if err != nil {
		return err
	}

	err = writer.Write(cxt, renders, pipelineOptions)
	return
}

func patchViolation(doc *asciidoc.Document, v Violation) (err error) {
	switch e := v.Entity.(type) {
	case *matter.EnumValue:
		source := e.Source()
		err = addProvisionalConformance(doc, e, source)
	default:
		slog.Error("Unexpected provisional entity", matter.LogEntity("entity", e), slog.String("path", v.Path))
	}
	return
}

func addProvisionalConformance(doc *asciidoc.Document, e types.Entity, source asciidoc.Element) (err error) {
	switch source := source.(type) {
	case *asciidoc.TableRow:
		var table *spec.TableInfo
		table, err = spec.ReadTable(doc, asciidoc.RawReader, source.Parent)
		if err != nil {
			return
		}
		conf := table.ReadConformance(asciidoc.RawReader, source, matter.TableColumnConformance)
		if conformance.IsProvisional(conf) {
			slog.Error("Already provisional!")
			return
		}
		conf = append(conformance.Set{&conformance.Provisional{}}, conf)
		conformanceIndex, ok := table.ColumnMap[matter.TableColumnConformance]
		if !ok {
			return
		}
		setCellString(source.TableCells()[conformanceIndex], conf.ASCIIDocString())
	default:
		slog.Error("Unexpected provisional conformance source", matter.LogEntity("entity", e), log.Type("sourceType", source))
	}
	return
}

func setCellString(cell *asciidoc.TableCell, v string) {
	se := asciidoc.NewString(v)
	cell.SetChildren(asciidoc.Elements{se})
}
