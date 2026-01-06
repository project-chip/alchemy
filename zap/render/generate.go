package render

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"

	"github.com/beevik/etree"
	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/internal/vcs"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/zap"
)

type TemplateGenerator struct {
	spec        *spec.Specification
	pipeline    pipeline.ProcessingOptions
	sdkRoot     string
	specVersion string

	options TemplateOptions

	ClusterAliases pipeline.Map[string, []string]
}

func NewTemplateGenerator(spec *spec.Specification, pipelineOptions pipeline.ProcessingOptions, sdkRoot string, options TemplateOptions) (*TemplateGenerator, error) {
	tg := &TemplateGenerator{
		spec:           spec,
		pipeline:       pipelineOptions,
		sdkRoot:        sdkRoot,
		ClusterAliases: pipeline.NewConcurrentMap[string, []string](),
		options:        options,
	}
	if spec.Root != "" {
		var err error
		tg.specVersion, err = vcs.GitDescribe(spec.Root)
		if err != nil {
			slog.Error("Unable to determine spec git tag", slog.Any("error", err))
			return nil, err
		}
	} else {
		slog.Warn("Skipping Git due to empty specification root")
	}
	return tg, nil
}

func (tg TemplateGenerator) Name() string {
	return "Generating ZAP XML"
}

func (tg TemplateGenerator) Process(cxt context.Context, input *pipeline.Data[*asciidoc.Document], index int32, total int32) (outputs []*pipeline.Data[string], extra []*pipeline.Data[*asciidoc.Document], err error) {
	d := input.Content
	entities := tg.spec.EntitiesForDocument(d)

	library, ok := tg.spec.LibraryForDocument(input.Content)
	if !ok {
		err = fmt.Errorf("unable to find library for doc %s", d.Path.Relative)
		return
	}
	errata := library.ErrataForPath(d.Path.Relative)

	if errata.SDK.SkipFile {
		return
	}

	destinations := ZAPTemplateDestinations(tg.sdkRoot, d.Path.Relative, entities, &errata.SDK)

	dependencies := pipeline.NewConcurrentMap[string, bool]()

	dependencies.Store(d.Path.Relative, true)

	for newPath, entities := range destinations {

		if len(entities) == 0 {
			slog.WarnContext(cxt, "Skipped spec file with no entities", "from", d.Path, "to", newPath)
			continue
		}

		var configurator *zap.Configurator
		configurator, err = zap.NewConfigurator(tg.spec, []*asciidoc.Document{d}, entities, newPath, &errata.SDK, false)
		if err != nil {
			return
		}

		var result string

		var doc *etree.Document
		doc, err = openConfigurator(configurator, tg.pipeline)
		if err != nil {
			return
		}

		for e := range configurator.ExternalEntities {
			externalDoc, ok := tg.spec.DocRefs[e]
			if !ok {
				slog.Error("External entity with no associated document", matter.LogEntity("entity", e), "path", newPath)
				continue
			}
			extra = append(extra, pipeline.NewData(externalDoc.Path.Absolute, externalDoc))
		}

		if configurator.IsEmpty() {
			continue
		}

		tg.buildClusterAliases(configurator)

		result, err = tg.renderZapTemplate(configurator, doc)
		if err != nil {
			err = fmt.Errorf("failed rendering %s: %w", d.Path.Relative, err)
			return
		}
		outputs = append(outputs, &pipeline.Data[string]{Path: newPath, Content: result})

	}
	return
}

func openConfigurator(configurator *zap.Configurator, options pipeline.ProcessingOptions) (doc *etree.Document, err error) {
	var existing []byte
	existing, err = os.ReadFile(configurator.OutPath)
	if errors.Is(err, os.ErrNotExist) {
		if options.Serial {
			slog.Info("Rendering new ZAP template", configurator.DocLogs(), "to", configurator.OutPath)
		}
		doc = newZapTemplate()
		err = nil
	} else if err != nil {
		return
	} else {
		if options.Serial {
			slog.Info("Rendering existing ZAP template", configurator.DocLogs(), "to", configurator.OutPath)
		}
		doc = etree.NewDocument()
		err = doc.ReadFromBytes(existing)
		if err != nil {
			err = fmt.Errorf("failed reading ZAP template %v: %w", configurator.Docs, err)
			return
		}
	}
	return
}

func SplitZAPDocs(cxt context.Context, spec *spec.Specification, inputs spec.DocSet) (clusters spec.DocSet, deviceTypes spec.DocSet, namespaces spec.DocSet, err error) {
	clusters = pipeline.NewMap[string, *pipeline.Data[*asciidoc.Document]]()
	deviceTypes = pipeline.NewMap[string, *pipeline.Data[*asciidoc.Document]]()
	namespaces = pipeline.NewMap[string, *pipeline.Data[*asciidoc.Document]]()
	inputs.Range(func(path string, data *pipeline.Data[*asciidoc.Document]) bool {
		var hasCluster bool
		var dts []*matter.DeviceType
		var ns []*matter.Namespace
		entities := spec.EntitiesForDocument(data.Content)
		for _, e := range entities {
			switch e := e.(type) {
			case *matter.Cluster, *matter.ClusterGroup:
				hasCluster = true
			case *matter.DeviceType:
				if e.Name == "Base Device Type" {
					continue
				}
				dts = append(dts, e)
			case *matter.Namespace:
				ns = append(ns, e)
			}
		}
		if hasCluster {
			clusters.Store(path, data)
		}
		if len(dts) > 0 {
			deviceTypes.Store(path, data)
		}
		if len(ns) > 0 {
			namespaces.Store(path, data)
		}
		return true
	})
	return
}
