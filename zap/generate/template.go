package generate

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/beevik/etree"
	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/internal/files"
	"github.com/hasty/alchemy/internal/pipeline"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/types"
	"github.com/hasty/alchemy/zap"
)

type TemplateGenerator struct {
	spec     *matter.Spec
	file     files.Options
	pipeline pipeline.Options
	sdkRoot  string

	ProvisionalZclFiles pipeline.Map[string, *pipeline.Data[struct{}]]
}

func NewTemplateGenerator(spec *matter.Spec, fileOptions files.Options, pipelineOptions pipeline.Options, sdkRoot string) *TemplateGenerator {
	return &TemplateGenerator{
		spec:                spec,
		file:                fileOptions,
		pipeline:            pipelineOptions,
		sdkRoot:             sdkRoot,
		ProvisionalZclFiles: pipeline.NewConcurrentMap[string, *pipeline.Data[struct{}]](),
	}
}

func (p TemplateGenerator) Name() string {
	return "Generating ZAP XML"
}

func (p TemplateGenerator) Type() pipeline.ProcessorType {
	return pipeline.ProcessorTypeIndividual
}

func (p TemplateGenerator) Process(cxt context.Context, input *pipeline.Data[*ascii.Doc], index int32, total int32) (outputs []*pipeline.Data[string], extra []*pipeline.Data[*ascii.Doc], err error) {
	return p.render(cxt, input)
}

func (tg *TemplateGenerator) render(cxt context.Context, input *pipeline.Data[*ascii.Doc]) (outputs []*pipeline.Data[string], extra []*pipeline.Data[*ascii.Doc], err error) {
	var entities []types.Entity
	entities, err = input.Content.Entities()
	if err != nil {
		return
	}

	errata, ok := zap.Erratas[filepath.Base(input.Content.Path)]
	if !ok {
		errata = zap.DefaultErrata
	}

	destinations := ZAPTemplateDestinations(tg.sdkRoot, input.Content.Path, entities, errata)

	dependencies := pipeline.NewConcurrentMap[string, bool]()

	dependencies.Store(input.Content.Path, true)

	for newPath, entities := range destinations {

		if len(entities) == 0 {
			continue
		}

		findDependencies(tg.spec, entities, dependencies)

		input.Content.Domain = getDocDomain(input.Content)

		if input.Content.Domain == matter.DomainUnknown {
			if errata.Domain != matter.DomainUnknown {
				input.Content.Domain = errata.Domain
			} else {
				input.Content.Domain = matter.DomainGeneral
			}
		}

		if len(entities) == 0 {
			slog.WarnContext(cxt, "Skipped spec file with no entities", "from", input.Content.Path, "to", newPath)
			return
		}

		var configurator *zap.Configurator
		configurator, err = zap.NewConfigurator(tg.spec, input.Content, entities)
		if err != nil {
			return
		}

		var result string

		var doc *etree.Document
		var provisional bool

		var existing []byte
		existing, err = os.ReadFile(newPath)
		if errors.Is(err, os.ErrNotExist) || tg.file.DryRun {
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
		result, err = renderZapTemplate(configurator, doc, errata)
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

func SplitZAPDocs(cxt context.Context, inputs pipeline.Map[string, *pipeline.Data[*ascii.Doc]]) (clusters pipeline.Map[string, *pipeline.Data[*ascii.Doc]], deviceTypes pipeline.Map[string, *pipeline.Data[[]*matter.DeviceType]], err error) {
	clusters = pipeline.NewMap[string, *pipeline.Data[*ascii.Doc]]()
	deviceTypes = pipeline.NewMap[string, *pipeline.Data[[]*matter.DeviceType]]()
	inputs.Range(func(path string, data *pipeline.Data[*ascii.Doc]) bool {
		var hasCluster bool
		var dts []*matter.DeviceType
		var entities []types.Entity
		entities, err = data.Content.Entities()
		if err != nil {
			slog.ErrorContext(cxt, "error converting doc to entities", "doc", data.Content.Path, "error", err)
			err = nil
			return true
		}
		for _, e := range entities {
			switch e := e.(type) {
			case *matter.Cluster:
				hasCluster = true
			case *matter.DeviceType:
				dts = append(dts, e)
			}
		}
		if hasCluster {
			clusters.Store(path, data)
		}
		if len(dts) > 0 {
			deviceTypes.Store(path, pipeline.NewData[[]*matter.DeviceType](path, dts))
		}
		return true
	})
	return
}
