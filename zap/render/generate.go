package render

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"

	"github.com/beevik/etree"
	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/internal/vcs"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
	"github.com/project-chip/alchemy/zap"
)

type TemplateGenerator struct {
	spec        *spec.Specification
	pipeline    pipeline.ProcessingOptions
	attributes  []asciidoc.AttributeName
	sdkRoot     string
	specVersion string

	options TemplateOptions

	globalObjectDependencies pipeline.Map[types.Entity, struct{}]

	ClusterAliases pipeline.Map[string, []string]
}

func NewTemplateGenerator(spec *spec.Specification, pipelineOptions pipeline.ProcessingOptions, sdkRoot string, options TemplateOptions) (*TemplateGenerator, error) {
	tg := &TemplateGenerator{
		spec:                     spec,
		pipeline:                 pipelineOptions,
		sdkRoot:                  sdkRoot,
		globalObjectDependencies: pipeline.NewConcurrentMap[types.Entity, struct{}](),
		ClusterAliases:           pipeline.NewConcurrentMap[string, []string](),
		options:                  options,
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

func (tg TemplateGenerator) Process(cxt context.Context, input *pipeline.Data[*spec.Doc], index int32, total int32) (outputs []*pipeline.Data[string], extra []*pipeline.Data[*spec.Doc], err error) {
	var entities []types.Entity
	entities, err = input.Content.Entities()
	if err != nil {
		return
	}

	errata := errata.GetSDK(input.Content.Path.Relative)

	if errata.SkipFile {
		return
	}

	destinations := ZAPTemplateDestinations(tg.sdkRoot, input.Content.Path.Relative, entities, errata)

	dependencies := pipeline.NewConcurrentMap[string, bool]()

	dependencies.Store(input.Content.Path.Relative, true)

	for newPath, entities := range destinations {

		if len(entities) == 0 {
			slog.WarnContext(cxt, "Skipped spec file with no entities", "from", input.Content.Path, "to", newPath)
			continue
		}

		tg.findDependencies(tg.spec, entities, dependencies)

		input.Content.Domain = getDocDomain(input.Content)

		if input.Content.Domain == matter.DomainUnknown {
			if errata.Domain != matter.DomainUnknown {
				input.Content.Domain = errata.Domain
			} else {
				input.Content.Domain = matter.DomainGeneral
			}
		}

		var configurator *zap.Configurator
		configurator, err = zap.NewConfigurator(tg.spec, []*spec.Doc{input.Content}, entities, newPath, errata, false)
		if err != nil {
			return
		}

		var result string

		var doc *etree.Document
		doc, err = tg.openConfigurator(configurator)
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

		tg.buildClusterAliases(configurator)

		result, err = tg.renderZapTemplate(configurator, doc)
		if err != nil {
			err = fmt.Errorf("failed rendering %s: %w", input.Content.Path, err)
			return
		}
		outputs = append(outputs, &pipeline.Data[string]{Path: newPath, Content: result})

	}
	return
}

func (tg *TemplateGenerator) openConfigurator(configurator *zap.Configurator) (doc *etree.Document, err error) {
	var existing []byte
	existing, err = os.ReadFile(configurator.OutPath)
	if errors.Is(err, os.ErrNotExist) {
		if tg.pipeline.Serial {
			slog.Info("Rendering new ZAP template", configurator.DocLogs(), "to", configurator.OutPath)
		}
		doc = newZapTemplate()
		err = nil
	} else if err != nil {
		return
	} else {
		if tg.pipeline.Serial {
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

func SplitZAPDocs(cxt context.Context, inputs spec.DocSet) (clusters spec.DocSet, deviceTypes spec.DocSet, namespaces pipeline.Map[string, *pipeline.Data[[]*matter.Namespace]], err error) {
	clusters = spec.NewDocSet()
	deviceTypes = spec.NewDocSet()
	namespaces = pipeline.NewMap[string, *pipeline.Data[[]*matter.Namespace]]()
	inputs.Range(func(path string, data *pipeline.Data[*spec.Doc]) bool {
		var hasCluster bool
		var dts []*matter.DeviceType
		var ns []*matter.Namespace
		var entities []types.Entity
		entities, err = data.Content.Entities()
		if err != nil {
			slog.ErrorContext(cxt, "error converting doc to entities", "doc", data.Content.Path, "error", err)
			err = nil
			return true
		}
		for _, e := range entities {
			switch e := e.(type) {
			case *matter.Cluster, *matter.ClusterGroup:
				hasCluster = true
			case *matter.DeviceType:
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
			namespaces.Store(path, pipeline.NewData(path, ns))
		}
		return true
	})
	return
}

func getDocDomain(doc *spec.Doc) matter.Domain {
	if doc.Domain != matter.DomainUnknown {
		return doc.Domain
	}
	for _, p := range doc.Parents() {
		d := getDocDomain(p)
		if d != matter.DomainUnknown {
			return d
		}
	}
	return matter.DomainUnknown
}
