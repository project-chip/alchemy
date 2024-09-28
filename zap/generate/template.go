package generate

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/beevik/etree"
	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal/files"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
	"github.com/project-chip/alchemy/zap"
)

type TemplateGenerator struct {
	spec        *spec.Specification
	file        files.Options
	pipeline    pipeline.Options
	attributes  []asciidoc.AttributeName
	sdkRoot     string
	specRoot    string
	specVersion string

	generateFeaturesXML bool

	ProvisionalZclFiles      pipeline.Map[string, *pipeline.Data[struct{}]]
	globalObjectDependencies pipeline.Map[types.Entity, struct{}]

	ClusterAliases pipeline.Map[string, []string]
}

type TemplateOption func(tg *TemplateGenerator)

func GenerateFeatureXML(generate bool) TemplateOption {
	return func(tg *TemplateGenerator) {
		tg.generateFeaturesXML = generate
	}
}

func AsciiAttributes(attributes []asciidoc.AttributeName) TemplateOption {
	return func(tg *TemplateGenerator) {
		tg.attributes = attributes
	}
}

func SpecRoot(specRoot string) TemplateOption {
	return func(tg *TemplateGenerator) {
		tg.specRoot = specRoot
	}
}

func NewTemplateGenerator(spec *spec.Specification, fileOptions files.Options, pipelineOptions pipeline.Options, sdkRoot string, options ...TemplateOption) *TemplateGenerator {
	tg := &TemplateGenerator{
		spec:                     spec,
		file:                     fileOptions,
		pipeline:                 pipelineOptions,
		sdkRoot:                  sdkRoot,
		ProvisionalZclFiles:      pipeline.NewConcurrentMap[string, *pipeline.Data[struct{}]](),
		globalObjectDependencies: pipeline.NewConcurrentMap[types.Entity, struct{}](),
		ClusterAliases:           pipeline.NewConcurrentMap[string, []string](),
	}
	for _, o := range options {
		o(tg)
	}
	if tg.specRoot != "" {
		var err error
		tg.specVersion, err = gitDescribe(tg.specRoot)
		if err != nil {
			slog.Warn("Unable to determine spec git tag", slog.Any("error", err))
		}
	}
	return tg
}

func (tg TemplateGenerator) Name() string {
	return "Generating ZAP XML"
}

func (tg TemplateGenerator) Type() pipeline.ProcessorType {
	return pipeline.ProcessorTypeIndividual
}

func (tg TemplateGenerator) Process(cxt context.Context, input *pipeline.Data[*spec.Doc], index int32, total int32) (outputs []*pipeline.Data[string], extra []*pipeline.Data[*spec.Doc], err error) {
	var entities []types.Entity
	entities, err = input.Content.Entities()
	if err != nil {
		return
	}

	errata := errata.GetZAP(input.Content.Path.Relative)

	destinations := ZAPTemplateDestinations(tg.sdkRoot, input.Content.Path.Relative, entities, errata)

	dependencies := pipeline.NewConcurrentMap[string, bool]()

	dependencies.Store(input.Content.Path.Relative, true)

	for newPath, entities := range destinations {

		if len(entities) == 0 {
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
		configurator, err = zap.NewConfigurator(tg.spec, input.Content, entities, newPath, errata)
		if err != nil {
			return
		}

		var result string

		var doc *etree.Document
		var provisional bool

		var existing []byte
		existing, err = os.ReadFile(newPath)
		if errors.Is(err, os.ErrNotExist) {
			if tg.pipeline.Serial {
				slog.InfoContext(cxt, "Rendering new ZAP template", "from", input.Content.Path, "to", newPath)
			}
			provisional = true
			doc = newZapTemplate()
		} else if err != nil {
			return
		} else {
			if tg.pipeline.Serial {
				slog.InfoContext(cxt, "Rendering existing ZAP template", "from", input.Content.Path, "to", newPath)
			}
			doc = etree.NewDocument()
			err = doc.ReadFromBytes(existing)
			if err != nil {
				err = fmt.Errorf("failed reading ZAP template %s: %w", input.Content.Path, err)
				return
			}

		}
		result, err = tg.renderZapTemplate(configurator, doc, errata)
		if err != nil {
			err = fmt.Errorf("failed rendering %s: %w", input.Content.Path, err)
			return
		}
		outputs = append(outputs, &pipeline.Data[string]{Path: newPath, Content: result})
		if provisional {
			tg.ProvisionalZclFiles.Store(filepath.Base(newPath), pipeline.NewData[struct{}](newPath, struct{}{}))
		}
	}
	return
}

func SplitZAPDocs(cxt context.Context, inputs pipeline.Map[string, *pipeline.Data[*spec.Doc]]) (clusters pipeline.Map[string, *pipeline.Data[*spec.Doc]], deviceTypes pipeline.Map[string, *pipeline.Data[[]*matter.DeviceType]], namespaces pipeline.Map[string, *pipeline.Data[[]*matter.Namespace]], err error) {
	clusters = pipeline.NewMap[string, *pipeline.Data[*spec.Doc]]()
	deviceTypes = pipeline.NewMap[string, *pipeline.Data[[]*matter.DeviceType]]()
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
			deviceTypes.Store(path, pipeline.NewData[[]*matter.DeviceType](path, dts))
		}
		if len(ns) > 0 {
			namespaces.Store(path, pipeline.NewData[[]*matter.Namespace](path, ns))
		}
		return true
	})
	return
}
